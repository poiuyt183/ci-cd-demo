package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {

}

var (
	ServiceName = "golang-docker-demo"
	port        = "8080"
)

func main() {
	http.HandleFunc("/ngonlun", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("ping vcl ok %s", ServiceName)))
	})
	fmt.Println("start service with port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}


