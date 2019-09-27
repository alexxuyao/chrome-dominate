package chromedominate

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestChromeTargetDominate_GetAllCookies(t *testing.T) {
	c, err := NewChromeDominate(DominateConfig{
		ChromePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	})

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetOneTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.NavigateLink("https://www.alipay.com/")

	if err != nil {
		log.Println(err, "open baidu error")
		t.Error(err)
		return
	}

	log.Println(ret.FrameId)

	time.Sleep(2 * time.Second)

	cookies, err := target.GetAllCookies()
	if err != nil {
		log.Println(err, "get cookie error")
		t.Error(err)
		return
	}

	if len(cookies) == 0 {
		t.Error("cookies is empty")
		return
	}

	fmt.Println(cookies)
}

func TestChromeTargetDominate_GetCookies(t *testing.T) {
	c, err := NewChromeDominate(DominateConfig{
		ChromePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	})

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetOneTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.NavigateLink("https://www.alipay.com/")

	if err != nil {
		log.Println(err, "open baidu error")
		t.Error(err)
		return
	}

	log.Println(ret.FrameId)

	time.Sleep(2 * time.Second)

	cookies, err := target.GetCookies([]string{
		"https://www.alipay.com/",
	})
	if err != nil {
		log.Println(err, "get cookie error")
		t.Error(err)
		return
	}

	if len(cookies) == 0 {
		t.Error("cookies is empty")
		return
	}

	for _, v := range cookies {
		if strings.Index(v.Name, "hello") >= 0 {
			fmt.Println(v)
		}
	}
}

func TestChromeTargetDominate_SetCookie(t *testing.T) {
	c, err := NewChromeDominate(DominateConfig{
		ChromePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	})

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetOneTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.NavigateLink("https://www.alipay.com/")

	if err != nil {
		log.Println(err, "open baidu error")
		t.Error(err)
		return
	}

	log.Println(ret.FrameId)

	time.Sleep(2 * time.Second)

	cUrl := "http://www.alipay.com"
	cDomain := ".alipay.com"
	cPath := "/"
	cSecure := false
	cHttpOnly := false
	cSameSite := "Lax"
	cExpires := -1.0

	cookie := CookieParam{
		Name:     "hello",
		Value:    "alex",
		Url:      &cUrl,
		Domain:   &cDomain,
		Path:     &cPath,
		Secure:   &cSecure,
		HttpOnly: &cHttpOnly,
		SameSite: &cSameSite,
		Expires:  &cExpires,
	}

	s, err := target.SetCookie(cookie)

	if err != nil {
		log.Println(err, "set cookie error")
		t.Error(err)
		return
	}

	fmt.Println(s)
}

type TestListener struct {
	Name string
}

func (c *TestListener) OnMessage(msgType string, message []byte) {
	fmt.Println("i am " + c.Name + ", i get message :" + string(message))
}

func TestChromeTargetDominate_ListenTarget(t *testing.T) {
	c, err := NewChromeDominate(DominateConfig{
		ChromePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	})

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetOneTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.NavigateLink("https://www.alipay.com/")

	if err != nil {
		log.Println(err, "open baidu error")
		t.Error(err)
		return
	}

	log.Println(ret.FrameId)

	time.Sleep(2 * time.Second)
	//mainTarget, err := c.GetMainTarget()
	//if err != nil {
	//	log.Println(err, "GetMainTarget error")
	//	t.Error(err)
	//	return
	//}

	err = target.EnableNetwork(NetworkEnableParam{})
	if err != nil {
		log.Println(err, "EnableNetwork error")
		t.Error(err)
		return
	}

	l1 := &TestListener{
		Name: "l1",
	}
	target.AddListener(l1)

	time.Sleep(3 * time.Minute)
}

func TestChromeTargetDominate_EnablePage(t *testing.T) {
	c, err := NewChromeDominate(DominateConfig{
		ChromePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	})

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetOneTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.NavigateLink("https://www.alipay.com/")

	if err != nil {
		log.Println(err, "open baidu error")
		t.Error(err)
		return
	}

	log.Println(ret.FrameId)

	time.Sleep(2 * time.Second)
	//mainTarget, err := c.GetMainTarget()
	//if err != nil {
	//	log.Println(err, "GetMainTarget error")
	//	t.Error(err)
	//	return
	//}

	err = target.EnablePage()
	if err != nil {
		log.Println(err, "EnablePage error")
		t.Error(err)
		return
	}

	l1 := &TestListener{
		Name: "l1",
	}
	target.AddListener(l1)

	time.Sleep(3 * time.Minute)
}
