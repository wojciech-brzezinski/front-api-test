package main

import (
	"front-test-api/internal/server"
	"log"
)

func main() {
	if err := server.Start(80); nil != err {
		log.Fatal(err)
	}
}
