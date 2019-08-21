/**
 * @time: 2019-08-21 22:42
 * @author: louis
 */
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHand{},
		ReadTimeout: 5 * time.Second,
	}
	mux := make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayH //路由注册
	mux["/bye"] = sayB

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

type myHand struct {
}

func (*myHand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)  //转发
		return
	}
	io.WriteString(w, "URL: "+r.URL.String())
}

func sayH(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world,this is version 3.")
}

func sayB(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Say Bye.")
}
