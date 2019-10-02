package chromedominate

func (c *ChromeTargetDominate) CanEmulate() (bool, error) {
	cmd := CmdRootType{
		Method: "Emulation.canEmulate",
		Params: make(map[string]interface{}),
	}

	ret := &CanEmulateResult{}
	_, err := c.SendCmdWithResult(cmd, ret)
	return ret.Result, err
}

func (c *ChromeTargetDominate) SetUserAgentOverride(param SetUserAgentOverrideParam) error {
	cmd := CmdRootType{
		Method: "Emulation.setUserAgentOverride",
		Params: param,
	}

	ret := make(map[string]interface{})
	_, err := c.SendCmdWithResult(cmd, ret)

	return err
}
