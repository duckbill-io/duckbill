// duckbill 一个简单的博客引擎
package main

import (
	"github.com/duckbill-io/duckbill/router"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router.New(),
	}
	log.Println("starting duckbill ...")
	log.Fatal(server.ListenAndServe())
}
