package chromedominate

func (c *ChromeTargetDominate) DispatchMouseEvent(param DispatchMouseEventParam) error {
	cmd := CmdRootType{
		Method: "Input.dispatchMouseEvent",
		Params: param,
	}

	ret := make(map[string]string)
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChromeTargetDominate) DispatchKeyEvent(param DispatchKeyEventParam) error {
	cmd := CmdRootType{
		Method: "Input.dispatchKeyEvent",
		Params: param,
	}

	ret := make(map[string]string)
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChromeTargetDominate) InsertText(param string) error {
	cmd := CmdRootType{
		Method: "Input.insertText",
		Params: map[string]string{
			"text": param,
		},
	}

	ret := make(map[string]string)
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return err
	}

	return nil
}
