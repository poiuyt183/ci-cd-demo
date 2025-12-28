package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

var (
	ServiceName = "golang-docker-demo"
	port        = "8080"
)

func main() {
	http.HandleFunc("/ngonluon", func(w http.ResponseWriter, req *http.Request) {
		apiRes := os.Getenv("API_RES")
		w.Write([]byte(apiRes))
	})
	fmt.Println("start service with port: ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}


