package get

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/mikeweiwei/proxyip/model"
)

// Data5u get ip from data5u.com
func Data5u() (result []*model.Ip) {
	pollURL := "http://www.data5u.com/free/index.shtml"
	resp, _, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		log.Println(errs)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find("body > div.wlist > li:nth-child(2) > ul").Each(func(i int, s *goquery.Selection) {
		node := strconv.Itoa(i + 1)
		ss := s.Find("ul:nth-child(" + node + ") > span:nth-child(1) > li").Text()
		sss := s.Find("ul:nth-child(" + node + ") > span:nth-child(2) > li").Text()
		ssss := s.Find("ul:nth-child(" + node + ") > span:nth-child(4) > li").Text()
		ip := model.Newip()
		ip.Ip = ss + ":" + sss
		ip.IpType = ssss
		println(ss + ":" + sss)
	})
	log.Println("5u done")
	return
}
