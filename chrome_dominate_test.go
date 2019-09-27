package chromedominate

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {

	c, err := NewChromeDominate(DominateConfig{
		ChromePath: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	})

	if err != nil {
		log.Println(err, "new chrome dominate error")
	}

	target, err := c.GetOneTarget()

	if err != nil {
		log.Println(err, "new chrome dominate error")
	}

	ret, err := target.NavigateLink("https://www.alipay.com/")

	if err != nil {
		log.Println(err, "open baidu error")
	}

	log.Println(ret.FrameId)

}
