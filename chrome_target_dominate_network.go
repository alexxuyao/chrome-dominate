package chromedominate

func (c *ChromeTargetDominate) GetAllCookies() ([]Cookie, error) {

	cmd := CmdRootType{
		Method: "Network.getAllCookies",
		Params: make(map[string]interface{}),
	}

	ret := &GetCookieResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret.Cookies, nil
}

func (c *ChromeTargetDominate) GetCookies(urls []string) ([]Cookie, error) {

	cmd := CmdRootType{
		Method: "Network.getCookies",
		Params: map[string]interface{}{
			"urls": urls,
		},
	}

	ret := &GetCookieResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return nil, err
	}

	return ret.Cookies, nil
}

func (c *ChromeTargetDominate) SetCookie(cookie CookieParam) (bool, error) {

	cmd := CmdRootType{
		Method: "Network.setCookie",
		Params: cookie,
	}

	ret := &SetCookieResult{}
	_, err := c.SendCmdWithResult(cmd, ret)

	if err != nil {
		return false, err
	}

	return ret.Success, nil
}

func (c *ChromeTargetDominate) SetCookies(cookies []CookieParam) error {
	cmd := CmdRootType{
		Method: "Network.setCookies",
		Params: cookies,
	}

	ret := make(map[string]interface{})
	_, err := c.SendCmdWithResult(cmd, &ret)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChromeTargetDominate) EnableNetwork(param NetworkEnableParam) error {
	cmd := CmdRootType{
		Method: "Network.enable",
		Params: param,
	}

	ret := make(map[string]interface{})
	_, err := c.SendCmdWithResult(cmd, &ret)

	if err != nil {
		return err
	}

	return nil
}

func (c *ChromeTargetDominate) GetResponseBody(requestId RequestId) (GetResponseBodyResult, error) {
	cmd := CmdRootType{
		Method: "Network.getResponseBody",
		Params: map[string]string{
			"requestId": string(requestId),
		},
	}

	ret := GetResponseBodyResult{}
	_, err := c.SendCmdWithResult(cmd, &ret)

	if err != nil {
		return ret, err
	}

	return ret, nil
}
