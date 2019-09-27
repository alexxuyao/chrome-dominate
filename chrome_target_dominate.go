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
	OnMessage(method string, message []byte)
}

type ChromeTargetDominate struct {
	TargetInfo ChromeTargetType
	IsAlive    bool // 是否活着，用于表示此target的websocket是否还活着

	listeners   []ChromeEventListener
	conn        *websocket.Conn
	resultCache *ResultCache
	tmpId       int64
	mutex       *sync.RWMutex
	rootDom     *ChromeDOM
	parent      *ChromeDominate
}

func NewChromeTarget(info ChromeTargetType, parent *ChromeDominate) (*ChromeTargetDominate, error) {

	target := &ChromeTargetDominate{
		TargetInfo:  info,
		listeners:   make([]ChromeEventListener, 0),
		resultCache: NewResultCache(1*time.Minute, 10*time.Second),
		tmpId:       0,
		mutex:       new(sync.RWMutex),
		parent:      parent,
	}

	if err := target.InitWebSocket(); err != nil {
		return nil, err
	}

	return target, nil
}

func (c *ChromeTargetDominate) SetAlive(alive bool) {
	c.IsAlive = alive
	if !alive {
		c.parent.RemoveTargetById(c.TargetInfo.Id)
	}
}

func (c *ChromeTargetDominate) InitWebSocket() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.TargetInfo.WebSocketDebuggerUrl, nil)

	if err != nil {
		return err
	}

	c.conn = conn
	c.SetAlive(true)

	conn.SetCloseHandler(func(code int, text string) error {

		log.Println("websocket close.", code, text)

		c.SetAlive(false)
		return nil
	})

	// 处理接收消息
	go func(target *ChromeTargetDominate) {
		for {
			_, message, err := target.conn.ReadMessage()
			if err != nil {
				log.Println("websocket read message error:", err)
				c.SetAlive(false)
				return
			}

			msg := string(message)
			log.Println("websocket recv:", msg)

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

	return nil
}

func (c *ChromeTargetDominate) Close() error {

	if nil != c.conn {
		return c.conn.Close()
	}

	return nil
}

func (c *ChromeTargetDominate) fireMessage(method string, message []byte) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for _, listener := range c.listeners {
		listener.OnMessage(method, message)
	}
}

func (c *ChromeTargetDominate) newReqId() int64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.tmpId++
	return c.tmpId
}

func (c *ChromeTargetDominate) SendCmd(cmd CmdRootType) (int64, error) {

	cmd.Id = c.newReqId()

	// 序列化消息
	msg, err := json.Marshal(cmd)

	if err != nil {
		return -1, err
	}

	log.Println(string(msg))

	if c.conn == nil {
		return -1, errors.New("conn is nil")
	}

	err = c.conn.WriteMessage(websocket.TextMessage, msg)

	if err != nil {
		return -1, err
	}

	return cmd.Id, nil
}

func (c *ChromeTargetDominate) SendCmdWithResult(cmd CmdRootType, result interface{}) (int64, error) {

	id, err := c.SendCmd(cmd)
	if err != nil {
		return id, err
	}

	item, find := c.resultCache.Pop(id, time.Second*3)
	if !find {
		return id, errors.New("result not found")
	}

	log.Println("get result:", string(item.Data))

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
