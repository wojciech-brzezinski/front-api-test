package main

import (
	"front-test-api/internal/server"
	"log"
)

func main() {
	if err := server.Start(12345); nil != err {
		log.Fatal(err)
	}
}
