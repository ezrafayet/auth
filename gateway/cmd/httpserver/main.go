package main

import (
	"gateway/src/infra/httpserver"
	"log"
)

func main() {
	log.Fatal(httpserver.Start())
}
