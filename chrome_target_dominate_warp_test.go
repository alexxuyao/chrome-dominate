package chromedominate

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestChromeTargetDominate_GetAllCookies(t *testing.T) {
	c, err := NewChromeDominate("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome")

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetDefaultTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.OpenPage("https://www.alipay.com/")

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
	c, err := NewChromeDominate("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome")

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetDefaultTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.OpenPage("https://www.alipay.com/")

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

	fmt.Println(len(cookies))
	for _, v := range cookies {
		if strings.Index(v.Name, "hello") >= 0 {
			fmt.Println(v)
		}
	}
}

func TestChromeTargetDominate_SetCookie(t *testing.T) {
	c, err := NewChromeDominate("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome")

	if err != nil {
		t.Error(err)
		return
	}

	target, err := c.GetDefaultTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
		t.Error(err)
		return
	}

	ret, err := target.OpenPage("https://www.alipay.com/")

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
		Value:    "world",
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
