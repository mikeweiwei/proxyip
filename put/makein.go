package put

import (
	"github.com/mikeweiwei/proxyip/model"
	"github.com/mikeweiwei/proxyip/check"
	"log"
)

var db = DB()

func Check(ip *model.Ip) {
	if check.CheckIP(ip.Ip) == true {
		insertOne(db,ip.Ip,ip.IpType)
	}else {
		log.Printf("ip false --->" + ip.Ip)
	}
}

func CheckDB() {

	sum := count()
	log.Printf("befor checkDB size:(%v)",sum)
	all := FindAll()

	for _,v := range all{


			if check.CheckIP(v.Ip) == false {
				deleteOne(v.Ip,v.IpType)
			}


	}


	sum = count()
	log.Printf("after checkDB size:(%v)",sum)

}

