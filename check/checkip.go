package check

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

func CheckIP(ip string) bool {
	pollURL := "http://httpbin.org/get"
	resp, _, errs := gorequest.New().Proxy("http://" + ip).Get(pollURL).End()
	if errs != nil {
		return false
	}
	if resp.StatusCode == 200 {
		log.Printf("ip true --->" + ip)
		return true
	}
	return false
}
