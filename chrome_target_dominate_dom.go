package chromedominate

import (
	"errors"
	"log"
	"time"
)

func (c *ChromeTargetDominate) GetRootDOM() (*ResultDOMNode, error) {
	cmd := CmdRootType{
		Method: "DOM.getDocument",
		Params: make(map[string]interface{}),
	}

	ret := &GetRootResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return &ret.Root, nil
}

func (c *ChromeTargetDominate) QuerySelector(param QuerySelectorParam) (*QuerySelectorResult, error) {
	cmd := CmdRootType{
		Method: "DOM.querySelector",
		Params: param,
	}

	ret := &QuerySelectorResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *ChromeTargetDominate) QuerySelectorFromRoot(selector string) (*QuerySelectorResult, error) {
	root, err := c.GetRootDOM()

	if err != nil {
		return nil, err
	}

	param := QuerySelectorParam{
		NodeId:   root.NodeId,
		Selector: selector,
	}

	return c.QuerySelector(param)
}

func (c *ChromeTargetDominate) QuerySelectorTimeout(selector string, t time.Duration) (*QuerySelectorResult, error) {

	st := time.Now().Unix()
	tryTime := 0
	for {
		ret, err := c.QuerySelectorFromRoot(selector)
		//if err != nil {
		//	return nil, err
		//}

		if err == nil && ret.NodeId > 0 {

			if tryTime > 0 {
				log.Print("tryTime is ", tryTime)
			}

			return ret, nil
		}

		n := time.Now().Unix()

		if n-st > int64(t.Seconds()) {
			return nil, errors.New("query timeout:" + err.Error())
		}

		time.Sleep(time.Millisecond * 50)
		tryTime++
	}
}

func (c *ChromeTargetDominate) GetContentQuads(param GetContentQuadsParam) (*GetContentQuadsResult, error) {
	cmd := CmdRootType{
		Method: "DOM.getContentQuads",
		Params: param,
	}

	ret := &GetContentQuadsResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *ChromeTargetDominate) GetBoxModel(param GetBoxModelParam) (*BoxModel, error) {
	cmd := CmdRootType{
		Method: "DOM.getBoxModel",
		Params: param,
	}

	ret := &GetBoxModelResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return &ret.Model, nil
}

func (c *ChromeTargetDominate) QueryDom(selector string) (*ChromeDOM, error) {
	ret, err := c.QuerySelectorFromRoot(selector)
	if err != nil {
		return nil, err
	}

	if ret.NodeId <= 0 {
		return nil, errors.New("node not found")
	}

	return NewChromeDom(ret.NodeId, c), nil
}

func (c *ChromeTargetDominate) QueryDomTimeout(selector string, t time.Duration) (*ChromeDOM, error) {
	ret, err := c.QuerySelectorTimeout(selector, t)
	if err != nil {
		return nil, err
	}

	if ret.NodeId <= 0 {
		return nil, errors.New("node not found")
	}

	return NewChromeDom(ret.NodeId, c), nil
}

func (c *ChromeTargetDominate) SetAttributeValue(nodeId NodeId, name string, value string) error {
	param := SetAttributeValueParam{
		NodeId: nodeId,
		Name:   name,
		Value:  value,
	}

	cmd := CmdRootType{
		Method: "DOM.setAttributeValue",
		Params: param,
	}

	ret := make(map[string]string)
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChromeTargetDominate) GetOuterHTML(param GetOuterHTMLParam) (string, error) {

	cmd := CmdRootType{
		Method: "DOM.getOuterHTML",
		Params: param,
	}

	ret := &GetOuterHTMLResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return "", err
	}

	return ret.OuterHTML, nil
}
