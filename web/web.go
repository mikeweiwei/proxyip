package web

import (
	"net/http"
	"encoding/json"
	"github.com/mikeweiwei/proxyip/put"
	"log"
	"strings"
)

const VERSION  = "/v1"

	func WebRun()  {
	
	http.HandleFunc(VERSION + "/all",GetAllHandler)
	http.HandleFunc(VERSION + "/type",GetTypeHandler)
	log.Println("Starting server", "9999")
	http.ListenAndServe(":9999", nil)


}

func GetAllHandler(w http.ResponseWriter, r *http.Request)  {

	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(put.FindAll())
		if err != nil {
			return
		}
		w.Write(b)
	}
	return
	
}

func GetTypeHandler(w http.ResponseWriter, r *http.Request)  {

	r.ParseForm()  //解析参数，默认是不会解析的
	iptype := strings.Join(r.Form["type"], "")

	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		data, error := json.Marshal(put.FindType(iptype))
		if error != nil {
			return
		}
		w.Write(data)
	}


	return

}
