package pubtools

import (
	"math/rand"
	"time"
)

var (
	agents = []string{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.11 (KHTML, like Gecko) Ubuntu/11.10 Chromium/27.0.1453.93 Chrome/27.0.1453.93 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.94 Safari/537.36",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:21.0) Gecko/20130331 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.2; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0",
		"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv 11.0) like Gecko",
	}
)

// RandomUserAgent returns random useragent
func RandomUserAgent() string {
	rand.Seed(time.Now().Unix())
	return agents[rand.Intn(len(agents))]
}
