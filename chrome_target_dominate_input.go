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
