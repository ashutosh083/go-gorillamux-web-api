package main

import (
	"amazon/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("[ INFO ] Listening at port 4000")

}
