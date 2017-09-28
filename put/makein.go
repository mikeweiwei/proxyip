package put

import (
	"proxyip/model"
	"proxyip/check"
	"log"
	"sync"
)

var db = DB()

func Check(ip *model.Ip) {
	if check.CheckIP(ip.Ip) {
		insertOne(db,ip.Ip,ip.IpType)
	}else {
		log.Printf("ip false --->" + ip.Ip)
	}
}

func CheckDB() {
	sum := count()
	log.Printf("befor checkDB size:(%v)",sum)
	all := FindAll()
	var wg sync.WaitGroup
	for _,v := range all{
		wg.Add(1)
		go func() {
			if !check.CheckIP(v.Ip) {
				deleteOne(v.Ip,v.IpType)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	sum = count()
	log.Printf("after checkDB size:(%v)",sum)

}

