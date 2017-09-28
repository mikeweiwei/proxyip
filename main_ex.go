package main

import (



	"proxyip/get"
	"proxyip/web"
	"sync"
	"proxyip/model"
	"proxyip/put"
	"log"
	"runtime"
	//"time"
	"time"
	"fmt"
)



func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	ipChan := make(chan *model.Ip, 2000)
	put.Init()

	go func() {
		web.WebRun()

	}()

	go func() {
		put.CheckDB()
	}()

	//go func() {
	//	pr(ipChan)
	//}()

	for i := 0; i < 50; i++ {
		go func() {
			for {
				put.Check(<-ipChan)
			}
		}()
	}

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


