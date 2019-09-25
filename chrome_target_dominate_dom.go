package chromedominate

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

func (c *ChromeTargetDominate) QueryDom(selector string) (*ChromeDOM, error) {
	ret, err := c.QuerySelectorFromRoot(selector)
	if err != nil {
		return nil, err
	}

	return NewChromeDom(ret.NodeId, c), nil
}
