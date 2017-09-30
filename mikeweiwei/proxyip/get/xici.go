package get

import (
	"log"
	"github.com/parnurzeal/gorequest"
	"github.com/PuerkitoBio/goquery"
	"github.com/mikeweiwei/proxyip/model"
	"strconv"
)

func GetIP336() (result []*model.Ip) {

	for i := 1; i < 7; i++ {
		pollURL := "http://www.ip3366.net/free/?stype=1&page=" + strconv.Itoa(i)

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

		tr := doc.Find("#list").Find("tbody").Find("tr")
		tr.Each(func(i int, s *goquery.Selection) {

			ip := s.Find("td").Eq(0).Text()
			port := s.Find("td").Eq(1).Text()
			ipType := s.Find("td").Eq(3).Text()
			//println(ip + port + ipType)
			data := model.Newip()
			data.Ip = ip + ":" + port
			data.IpType = ipType
			result = append(result, data)
		})
	}
	log.Println("336 done")
	return
}
