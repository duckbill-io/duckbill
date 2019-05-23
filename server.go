// 博客引擎的服务器模块
package main

import (
	"github.com/duckbill-io/duckbill/router"
	"net/http"
	"log"
)

func main() {
	server := http.Server {
		Addr: "localhost:8080",
		Handler: router.New(),
	}
	log.Println("starting duckbill ...")
	log.Fatal(server.ListenAndServe())
}
