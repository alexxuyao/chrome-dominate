package chromedominate

func (c *ChromeTargetDominate) CompileScript(param CompileScriptParam) (*CompileScriptResult, error) {
	cmd := CmdRootType{
		Method: "Runtime.compileScript",
		Params: param,
	}

	ret := &CompileScriptResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *ChromeTargetDominate) RunScript(param RunScriptParam) (*RunScriptResult, error) {
	cmd := CmdRootType{
		Method: "Runtime.runScript",
		Params: param,
	}

	ret := &RunScriptResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *ChromeTargetDominate) Evaluate(param EvaluateParam) (*EvaluateResult, error) {
	cmd := CmdRootType{
		Method: "Runtime.evaluate",
		Params: param,
	}

	ret := &EvaluateResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret, nil
}
