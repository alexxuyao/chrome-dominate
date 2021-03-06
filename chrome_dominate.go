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
	"strconv"
	"time"
)

type ChromeDominate struct {
	config     DominateConfig
	process    *os.Process
	mainTarget *ChromeTargetDominate
	targets    map[TargetID]*ChromeTargetDominate
}

type AfterNewChromeDominateTarget interface {
	AfterNewChromeDominateTargetCreate(c *ChromeTargetDominate)
}

type DominateConfig struct {
	ChromePath                   string
	RemoteDebuggingPort          int
	AfterNewChromeDominateTarget []AfterNewChromeDominateTarget
}

func NewChromeDominate(config DominateConfig) (*ChromeDominate, error) {

	if config.RemoteDebuggingPort == 0 {
		config.RemoteDebuggingPort = 9222
	}

	dominate := &ChromeDominate{
		config:  config,
		targets: make(map[TargetID]*ChromeTargetDominate),
	}

	// 启动 chrome 进程
	cmd := exec.Command(config.ChromePath, "--remote-debugging-port="+strconv.Itoa(config.RemoteDebuggingPort))
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
			return nil, errors.New("can not dial tcp localhost:" + strconv.Itoa(config.RemoteDebuggingPort))
		}

		c, err := net.Dial("tcp", "localhost:"+strconv.Itoa(config.RemoteDebuggingPort))
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if c.Close() != nil {
			log.Print(err)
		}

		break
	}

	return dominate, nil
}

func NewServerAgentChromeDominate(config DominateConfig) (*ChromeDominate, error) {

	dominate := &ChromeDominate{
		config:  config,
		targets: make(map[TargetID]*ChromeTargetDominate),
	}

	return dominate, nil
}

func (c *ChromeDominate) RemoveTargetById(id TargetID) {
	// TODO 是否需要锁
	delete(c.targets, id)
}

func (c *ChromeDominate) GetOneTarget() (*ChromeTargetDominate, error) {

	if len(c.targets) == 0 {

		// 拿target
		if err := c.InitPageTargets(); err != nil {
			return nil, err
		}

	}

	for _, v := range c.targets {
		if v.IsAlive {
			return v, nil
		}
	}

	return nil, errors.New("not target found")
}

func (c *ChromeDominate) GetMainTarget() (*ChromeTargetDominate, error) {
	if nil == c.mainTarget {
		if err := c.initMainTarget(); err != nil {
			return nil, err
		}
	}

	if nil != c.mainTarget {
		return c.mainTarget, nil
	}

	return nil, errors.New("can not init main target")
}

func (c *ChromeDominate) InitPageTargets() error {
	url := "http://localhost:" + strconv.Itoa(c.config.RemoteDebuggingPort) + "/json/list"
	resp, err := http.Get(url)

	if err != nil {
		log.Print(err, url)
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if 200 != resp.StatusCode {
		return errors.New("httpStatusCode not 200:" + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err, url)
		return err
	}

	targets := make([]ChromeTargetType, 0)
	err = json.Unmarshal(body, &targets)

	if err != nil {
		log.Print(err, url)
		return err
	}

	for _, target := range targets {
		if _, find := c.targets[target.Id]; !find && target.Type == "page" {
			newItem, err := NewChromeTarget(target, c)
			if err != nil {
				log.Print(err, target.WebSocketDebuggerUrl)
				continue
			}

			for _, v := range c.config.AfterNewChromeDominateTarget {
				v.AfterNewChromeDominateTargetCreate(newItem)
			}

			c.targets[target.Id] = newItem
		}
	}

	return nil
}

func (c *ChromeDominate) initMainTarget() error {

	url := "http://localhost:" + strconv.Itoa(c.config.RemoteDebuggingPort) + "/json/version"
	resp, err := http.Get(url)

	if err != nil {
		log.Print(err, url)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err, url)
		return err
	}

	err = resp.Body.Close()

	if err != nil {
		log.Print(err, url)
		return err
	}

	target := ChromeMainTargetType{}
	err = json.Unmarshal(body, &target)

	if err != nil {
		log.Print(err, url)
		return err
	}

	info := ChromeTargetType{
		WebSocketDebuggerUrl: target.WebSocketDebuggerUrl,
	}

	newItem, err := NewChromeTarget(info, c)
	if err != nil {
		log.Print(err, target.WebSocketDebuggerUrl)
		return err
	}

	c.mainTarget = newItem
	return nil
}
