/**
 * @time: 2019-08-21 22:34
 * @author: louis
 */
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux() //自己实现handler,返回一个Mux
	mux.HandleFunc("/hello",sayHello)
	mux.Handle("/",&myHandler{})//注册服务

	wd,err := os.Getwd()
	if err !=nil {
		log.Fatal(err)
	}

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":9090",mux)
	if err!=nil {
		log.Fatal("ListenAndServe: ",err)
	}
}

type myHandler struct {

}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "URL: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Hello world,this is version 2.")
}
