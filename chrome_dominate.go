package chromedominate

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type ChromeDominate struct {
	process *os.Process
	targets []*ChromeTargetDominate
}

func NewChromeDominate(chromePath string) (*ChromeDominate, error) {

	dominate := &ChromeDominate{
		targets: make([]*ChromeTargetDominate, 0),
	}

	// 启动 chrome 进程
	cmd := exec.Command(chromePath, "--remote-debugging-port=9222")
	var cmdErr error = nil
	go func() {
		_, err := cmd.Output()
		if err != nil {
			cmdErr = err
		}
	}()

	for {
		if nil != cmd.Process {
			dominate.process = cmd.Process
			break
		}

		if cmdErr != nil {
			return nil, cmdErr
		}

		time.Sleep(10 * time.Millisecond)
	}

	// 启动成功，监听失败
	// 可能是原来就已经启动

	// 检查端口是否已经监听
	startTime := time.Now().Unix()

	for {

		now := time.Now().Unix()
		if now-startTime > 10 {
			return nil, errors.New("can not dial tcp localhost:9222")
		}

		c, err := net.Dial("tcp", "localhost:9222")
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if c.Close() != nil {
			log.Print(err)
		}

		break
	}

	// 拿target
	url := "http://localhost:9222/json/list"
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err, url)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err, url)
		return nil, err
	}

	err = resp.Body.Close()

	if err != nil {
		log.Print(err, url)
		return nil, err
	}

	targets := make([]ChromeTargetType, 0)
	err = json.Unmarshal(body, &targets)

	if err != nil {
		log.Print(err, url)
		return nil, err
	}

	for _, target := range targets {
		if target.Type == "page" {
			newItem, err := NewChromeTarget(target.WebSocketDebuggerUrl)
			if err != nil {
				log.Print(err, target.WebSocketDebuggerUrl)
				continue
			}

			dominate.targets = append(dominate.targets, newItem)
		}
	}

	return dominate, nil
}

func (c *ChromeDominate) GetDefaultTarget() (*ChromeTargetDominate, error) {

	if len(c.targets) > 0 {
		return c.targets[0], nil
	}

	return nil, errors.New("targets not found")
}
