package chromedominate

import "time"

func (c *ChromeTargetDominate) FetchDisable() error {
	cmd := CmdRootType{
		Method: "Fetch.disable",
		Params: make(map[string]interface{}),
	}

	_, err := c.SendCmd(cmd)

	return err
}

func (c *ChromeTargetDominate) FetchEnable(param FetchEnableParam) error {

	if nil != param.Patterns {

		i2r := map[InterceptionStage]RequestStage{
			"Request":         "Request",
			"HeadersReceived": "Response",
		}

		r2i := map[RequestStage]InterceptionStage{
			"Request":  "Request",
			"Response": "HeadersReceived",
		}

		for index, _ := range param.Patterns {
			pattern := &param.Patterns[index]
			if nil != pattern.InterceptionStage {
				tmp := i2r[*pattern.InterceptionStage]
				pattern.RequestStage = &tmp
			} else if nil != pattern.RequestStage {
				tmp := r2i[*pattern.RequestStage]
				pattern.InterceptionStage = &tmp
			}
		}
	}

	cmd := CmdRootType{
		Method: "Fetch.enable",
		Params: param,
	}

	_, err := c.SendCmd(cmd)

	return err
}

func (c *ChromeTargetDominate) FetchFailRequest(param FailRequestParam) error {
	cmd := CmdRootType{
		Method: "Fetch.failRequest",
		Params: param,
	}

	_, err := c.SendCmd(cmd)

	return err
}

func (c *ChromeTargetDominate) FetchFulfillRequest(param FulfillRequestParam) error {
	cmd := CmdRootType{
		Method: "Fetch.fulfillRequest",
		Params: param,
	}

	_, err := c.SendCmd(cmd)

	return err
}

func (c *ChromeTargetDominate) FetchContinueRequest(param ContinueRequestParam) error {
	cmd := CmdRootType{
		Method: "Fetch.continueRequest",
		Params: param,
	}

	_, err := c.SendCmd(cmd)

	return err
}

func (c *ChromeTargetDominate) FetchGetResponseBody(requestId RequestId) (GetResponseBodyResult, error) {
	p := GetResponseBodyParam{
		RequestId: requestId,
	}

	cmd := CmdRootType{
		Method: "Fetch.getResponseBody",
		Params: p,
	}

	ret := GetResponseBodyResult{}
	_, err := c.SendCmdWithResultWait(cmd, &ret, 20*time.Second)

	if err != nil {
		return ret, err
	}

	return ret, nil
}
