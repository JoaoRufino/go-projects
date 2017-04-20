package main

import (
	"log"
	"net/http"

	"./common"
	"./routers"
	"github.com/codegangsta/negroni"
)

func main() {
	common.StartUp()
	r := routers.InitRouters()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Listening...")
	http.ListenAndServe(":8080", n)
}
