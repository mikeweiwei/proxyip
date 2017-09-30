package get

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/mikeweiwei/proxyip/model"
)

// IP181 get ip from ip181.com
func IP181() (result []*model.Ip) {
	pollURL := "http://www.ip181.com/"
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
	doc.Find("tr.warning").Each(func(i int, s *goquery.Selection) {
		ss := s.Find("td:nth-child(1)").Text()
		sss := s.Find("td:nth-child(2)").Text()
		ssss := s.Find("td:nth-child(4)").Text()
		ip := model.Newip()
		ip.Ip = ss + ":" + sss
		ip.IpType = ssss
		result = append(result, ip)
	})

	log.Println("IP181 done.")
	return
}
