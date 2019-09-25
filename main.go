package chromedominate

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func sendCmd(c *websocket.Conn, cmdMap map[string]interface{}) {
	cmd, err := json.Marshal(cmdMap)

	if err != nil {
		fmt.Println("dial:", err)
	}

	err = c.WriteMessage(websocket.TextMessage, cmd)

	if err != nil {
		fmt.Println("dial:", err)
	}
}

type QuerySelector struct {
	Id     int                 `json:"id"`
	Result QuerySelectorResult `json:"result"`
}

type GetContentQuads struct {
	Id     int                   `json:"id"`
	Result GetContentQuadsResult `json:"result"`
}

func getResult(msgChan chan []byte, obj interface{}) {
	msg := <-msgChan
	err := json.Unmarshal(msg, &obj)
	if err != nil {
		fmt.Println("get result err")
	}
}

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:9222/devtools/page/DAF3CA252E4FB5705C480B297657ABD9", nil)
	if err != nil {
		fmt.Println("dial:", err)
	}
	defer c.Close()

	msgChan := make(chan []byte)

	go func() {
		for {
			_, message, err := c.ReadMessage()
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
			}

			if _, exist := ret["id"]; exist {
				msgChan <- message
			}

		}
	}()

	//sendCmd(c, map[string]interface{}{
	//	"id":     1,
	//	"method": "Network.enable",
	//})
	//
	//sendCmd(c, map[string]interface{}{
	//	"id":     2,
	//	"method": "Network.getAllCookies",
	//})

	sendCmd(c, map[string]interface{}{
		"id":     3,
		"method": "DOM.getDocument",
	})

	msg := <-msgChan

	fmt.Println("--->", string(msg))

	sendCmd(c, map[string]interface{}{
		"id":     4,
		"method": "DOM.querySelector",
		"params": map[string]interface{}{
			"nodeId":   1,
			"selector": "#user-mobile",
		},
	})

	ret := QuerySelector{}
	getResult(msgChan, &ret)

	fmt.Println(ret.Result.NodeId)

	sendCmd(c, map[string]interface{}{
		"id":     5,
		"method": "DOM.getContentQuads",
		"params": map[string]interface{}{
			"nodeId": ret.Result.NodeId,
		},
	})

	ret2 := GetContentQuads{}
	getResult(msgChan, &ret2)

	fmt.Println(ret2.Result.Quads[0][0])

	sendCmd(c, map[string]interface{}{
		"id":     6,
		"method": "Input.dispatchMouseEvent",
		"params": map[string]interface{}{
			"type":        "mousePressed",
			"x":           ret2.Result.Quads[0][0] + 10,
			"y":           ret2.Result.Quads[0][1] + 10,
			"button":      "left",
			"buttons":     1,
			"clickCount":  1,
			"deltaX":      ret2.Result.Quads[0][0] + 10,
			"deltaY":      ret2.Result.Quads[0][1] + 10,
			"pointerType": "mouse",
		},
	})

	msg = <-msgChan

	mobile := "13510508340"
	for i := 0; i < len(mobile); i++ {
		sendCmd(c, map[string]interface{}{
			"id":     6,
			"method": "Input.dispatchKeyEvent",
			"params": map[string]interface{}{
				"type": "keyDown",
				"text": mobile[i : i+1], // 这个不能长于三个中文字符
			},
		})

		msg = <-msgChan
	}

	done := make(chan struct{})
	<-done

	fmt.Println("finish")
}
