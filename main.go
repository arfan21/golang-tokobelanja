package main

import (
	"log"

	"github.com/arfan21/golang-tokobelanja/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
