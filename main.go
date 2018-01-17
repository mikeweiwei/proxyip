package main

import (

	"github.com/mikeweiwei/proxyip/get"
	"github.com/mikeweiwei/proxyip/web"
	"sync"
	"github.com/mikeweiwei/proxyip/model"
	"github.com/mikeweiwei/proxyip/put"
	"github.com/robfig/cron"
	"log"
	"runtime"
	//"time"
	"time"
	"fmt"
)



func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	ipChan := make(chan *model.Ip, 2000)

	//connection DB and creat table
	put.Init()

	//open web api
	go func() {
		web.WebRun()

	}()


	//check in DB
	go func() {
		c := cron.New()
		i := 0
		spec := "0 */1 * * * ?"
		c.AddFunc(spec, func() {
			i++
			log.Println("cron running:", i)
			put.CheckDB()
		})
		c.Start()
		select{}
	}()

	//go func() {
	//	pr(ipChan)
	//}()


	//check chan to DB
	for i := 0; i < 100; i++ {
		go func() {
			for {
				put.Check(<-ipChan)
			}
		}()
	}

	//spider to chan
	for {
		log.Printf("Chan: %v", len(ipChan))
		if len(ipChan) < 100 {
			go run(ipChan)
		}
		time.Sleep(10 * time.Minute)
	}


}

func run(ipChan chan *model.Ip)  {

	var wg sync.WaitGroup

	funs := []func() []*model.Ip{
		get.GetIP336,
		get.Data5u,
		get.PLP,
		get.IP181,
	}
	for _,f := range funs{
		wg.Add(1)

		go func(f func() []*model.Ip) {
			temp := f()
			for _,v := range temp{
				ipChan <- v
			}
			wg.Done()
		}(f)
	}

	wg.Wait()

	log.Println("get done")
}
func pr(ipChan chan *model.Ip) {

	for value := range ipChan {

		fmt.Println(value)
	}
}


