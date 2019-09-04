package chromedominate

import (
	"encoding/json"
	"errors"
	"fmt"
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
	TargetInfo  ChromeTargetType
	listeners   []ChromeEventListener
	conn        *websocket.Conn
	resultCache *ResultCache
	tmpId       int64
	mutex       *sync.RWMutex
	rootDom     *ChromeDOM
}

func NewChromeTarget(info ChromeTargetType) (*ChromeTargetDominate, error) {

	target := &ChromeTargetDominate{
		TargetInfo:  info,
		listeners:   make([]ChromeEventListener, 0),
		resultCache: NewResultCache(1*time.Minute, 10*time.Second),
		tmpId:       0,
		mutex:       new(sync.RWMutex),
	}

	conn, _, err := websocket.DefaultDialer.Dial(info.WebSocketDebuggerUrl, nil)

	if err != nil {
		return nil, err
	}

	target.conn = conn

	go func(target *ChromeTargetDominate) {
		for {
			_, message, err := target.conn.ReadMessage()
			if err != nil {
				fmt.Println("error:", err)
				return
			}

			msg := string(message)
			fmt.Println("recv:", msg)

			ret := make(map[string]interface{})
			err = json.Unmarshal(message, &ret)

			if err != nil {
				fmt.Println("unmarshal error")
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
	}(target)

	return target, nil
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
		}
	}
}
