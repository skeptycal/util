// Package chromedriver implements Chromedriver
package chromedriver

import (
	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/webtools/getpage"
)

const (
	sampleChromeDriverPort = `http://localhost:9515/`
)

func SampleScript() string {
	s, err := getpage.GetPage(sampleChromeDriverPort)
	if err != nil {
		log.Error(err)
	}
	return s
}
