package chromedominate

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

func (c *ChromeTargetDominate) NavigateLink(link string) (*ResultPageNavigate, error) {
	param := NavigateParam{
		Url: link,
	}

	return c.Navigate(param)
}

func (c *ChromeTargetDominate) Navigate(param NavigateParam) (*ResultPageNavigate, error) {

	cmd := CmdRootType{
		Method: "Page.navigate",
		Params: param,
	}

	ret := &ResultPageNavigate{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *ChromeTargetDominate) EnablePage() error {
	cmd := CmdRootType{
		Method: "Page.enable",
		Params: map[string]string{},
	}

	ret := make(map[string]interface{})
	_, err := c.SendCmdWithResult(cmd, &ret)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChromeTargetDominate) AddScriptToEvaluateOnNewDocument(script string) (ScriptIdentifier, error) {

	param := AddScriptToEvaluateOnNewDocumentParam{
		Source: script,
	}

	cmd := CmdRootType{
		Method: "Page.addScriptToEvaluateOnNewDocument",
		Params: param,
	}

	ret := &AddScriptToEvaluateOnNewDocumentResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return "", err
	}

	return ret.Identifier, nil
}

// ============================================= 需要 EnablePage
func (c *ChromeTargetDominate) Navigated(param NavigateParam, t time.Duration) (*Frame, error) {
	l := NewPageEventListener()
	c.AddListener(l)
	defer func() {
		c.RemoveListener(l)
	}()

	ret, err := c.Navigate(param)
	if err != nil {
		return nil, err
	}

	return l.WaitFrameNavigated(ret.FrameId, t)
}

func (c *ChromeTargetDominate) NavigatedLink(link string, t time.Duration) (*Frame, error) {
	param := NavigateParam{
		Url: link,
	}
	return c.Navigated(param, t)
}

type PageEventListener struct {
	NavigateMap map[FrameId]*Frame
}

func NewPageEventListener() *PageEventListener {
	l := PageEventListener{
		NavigateMap: make(map[FrameId]*Frame),
	}

	return &l
}

func (p *PageEventListener) OnMessage(target *ChromeTargetDominate, method string, message []byte) {
	if EventPageFrameNavigated == method {

		param := &FrameNavigatedParam{}
		ret := CmdRootType{
			Params: param,
		}

		if err := json.Unmarshal(message, &ret); err != nil {
			log.Fatal(err)
			return
		}

		frame := &param.Frame

		p.NavigateMap[frame.Id] = frame
	}
}

func (p *PageEventListener) WaitFrameNavigated(frameId FrameId, t time.Duration) (*Frame, error) {
	st := time.Now().UnixNano()
	for {
		if ret, ok := p.NavigateMap[frameId]; ok {
			return ret, nil
		}

		nt := time.Now().UnixNano()
		if nt-st > t.Nanoseconds() {
			break
		}

		time.Sleep(10 * time.Millisecond)
	}

	return nil, errors.New("wait timeout")
}
