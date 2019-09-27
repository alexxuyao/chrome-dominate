package chromedominate

func (c *ChromeTargetDominate) GetTargetInfo(targetId TargetID) (*TargetInfo, error) {
	cmd := CmdRootType{
		Method: "Target.getTargetInfo",
		Params: map[string]interface{}{
			"targetId": targetId,
		},
	}

	ret := &GetTargetInfoResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return &ret.TargetInfo, nil
}
