package chromedominate

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"sync"
	"time"
)

type ChromeEventListener interface {
	OnMessage(target *ChromeTargetDominate, method string, message []byte)
}

type ChromeTargetDominate struct {
	TargetInfo  ChromeTargetType
	IsAlive     bool // 是否活着，用于表示此target的websocket是否还活着
	listeners   []ChromeEventListener
	conn        *websocket.Conn
	resultCache *ResultCache
	cmdCache    *ResultCache
	mutex       *sync.RWMutex
	chanCmd     chan CmdRootType
	chanTmpId   chan int64
	chanClose   chan int
	rootDom     *ChromeDOM
	parent      *ChromeDominate
}

func NewChromeTarget(info ChromeTargetType, parent *ChromeDominate) (*ChromeTargetDominate, error) {

	target := &ChromeTargetDominate{
		TargetInfo:  info,
		listeners:   make([]ChromeEventListener, 0),
		resultCache: NewResultCache(1*time.Minute, 10*time.Second),
		cmdCache:    NewResultCache(1*time.Minute, 3*time.Second),
		mutex:       new(sync.RWMutex),
		parent:      parent,
		chanCmd:     make(chan CmdRootType),
		chanTmpId:   make(chan int64),
		chanClose:   make(chan int),
	}

	if err := target.InitWebSocket(); err != nil {
		return nil, err
	}

	// id 生成器
	go func(target *ChromeTargetDominate) {
		tmpId := int64(0)
		for {
			tmpId++
			target.chanTmpId <- tmpId
		}
	}(target)

	return target, nil
}

func (c *ChromeTargetDominate) SetAlive(alive bool) {
	c.IsAlive = alive
	if !alive {
		c.parent.RemoveTargetById(c.TargetInfo.Id)
	}
}

func (c *ChromeTargetDominate) InitWebSocket() error {
	dialer := websocket.DefaultDialer
	dialer.WriteBufferSize = 1024 * 1024 * 20
	conn, _, err := dialer.Dial(c.TargetInfo.WebSocketDebuggerUrl, nil)

	if err != nil {
		return err
	}

	c.conn = conn
	c.SetAlive(true)

	// 处理接收消息
	go func(target *ChromeTargetDominate) {

		defer func() {
			if err := recover(); err != nil {
				log.Println(err)

				if err = target.Close(); err != nil {
					log.Println(err)
				}
			}
		}()

		for {

			if !target.IsAlive {
				log.Println("target not alive, break read message goroutine", target.TargetInfo.Id)
				break
			}

			_, message, err := target.conn.ReadMessage()
			if err != nil {
				log.Println("websocket read message error:", err)
				continue
			}

			// log.Println("websocket recv:", string(message))

			ret := make(map[string]interface{})
			err = json.Unmarshal(message, &ret)

			if err != nil {
				log.Println("unmarshal error", string(message))
				continue
			}

			if _, exist := ret["id"]; exist {
				target.resultCache.Put(int64(ret["id"].(float64)), message)
			} else {
				s := string(message)
				s = s[:strings.Index(s, ",")]
				s = s[strings.Index(s, ":")+1:]
				method := s[1 : len(s)-1]

				target.fireMessage(method, message)
			}

		}
	}(c)

	// 处理写
	go func(target *ChromeTargetDominate) {
		for {

			//
			select {
			case cmd := <-target.chanCmd:

				// 序列化消息
				msg, err := json.Marshal(cmd)

				if err != nil {
					target.cmdCache.Put(cmd.Id, []byte(err.Error()))
					continue
				}

				log.Println("sendCmd:", cmd.Id, cmd.Method)

				if c.conn == nil {
					target.cmdCache.Put(cmd.Id, []byte("conn is nil"))
					continue
				}

				err = c.conn.WriteMessage(websocket.TextMessage, msg)

				if err != nil {
					target.cmdCache.Put(cmd.Id, []byte(err.Error()))
					continue
				}

				target.cmdCache.Put(cmd.Id, []byte(""))
			case <-target.chanClose:
				log.Println("target not alive, break write message goroutine:", target.TargetInfo.Id)
				break
			}

		}
	}(c)

	return nil
}

func (c *ChromeTargetDominate) Close() error {

	log.Println("target close:", c.TargetInfo.Id)

	c.chanClose <- 1
	c.SetAlive(false)

	if nil != c.conn {
		return c.conn.Close()
	}

	return nil
}

func (c *ChromeTargetDominate) fireMessage(method string, message []byte) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for _, listener := range c.listeners {
		go listener.OnMessage(c, method, message)
	}
}

func (c *ChromeTargetDominate) SendCmd(cmd CmdRootType) (int64, error) {
	cmd.Id = <-c.chanTmpId
	c.chanCmd <- cmd

	item, find := c.cmdCache.Pop(cmd.Id, 3*time.Second)
	if !find {
		return cmd.Id, errors.New("result not found")
	}

	msg := string(item.Data)
	if "" != msg {
		return cmd.Id, errors.New(msg)
	}

	return cmd.Id, nil
}

func (c *ChromeTargetDominate) SendCmdWithResult(cmd CmdRootType, result interface{}) (int64, error) {
	return c.SendCmdWithResultWait(cmd, result, 3*time.Second)
}

func (c *ChromeTargetDominate) SendCmdWithResultWait(cmd CmdRootType, result interface{}, timeWait time.Duration) (int64, error) {

	id, err := c.SendCmd(cmd)
	if err != nil {
		return id, err
	}

	item, find := c.resultCache.Pop(id, timeWait)
	if !find {
		return id, errors.New("result not found")
	}

	// log.Println("get result:", string(item.Data))

	ret := &ResultRootType{
		Result: result,
	}

	err = json.Unmarshal(item.Data, ret)

	if err != nil {
		return id, err
	}

	if ret.Error.Code != 0 {
		return id, errors.New(ret.Error.Message)
	}

	return id, nil
}

func (c *ChromeTargetDominate) AddListener(listener ChromeEventListener) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.listeners = append(c.listeners, listener)
}

func (c *ChromeTargetDominate) RemoveListener(listener ChromeEventListener) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	length := len(c.listeners)
	for i := 0; i < length; i++ {
		item := c.listeners[i]
		if item == listener {
			c.listeners = append(c.listeners[:i], c.listeners[i+1:]...)
			i--
			length--
		}
	}
}
