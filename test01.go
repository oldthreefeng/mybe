/**
 * @time: 2019-08-21 20:45
 * @author: louis
 */
package main

import (
	"io"
	"log"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Hello world,this is version 1.")
}

func main() {
	http.HandleFunc("/", sayHelloWorld) //注册路由
	err := http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal("ListenAndServe: ",err)
	}
}

