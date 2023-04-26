package main

import (
	"iam/src/infra/httpserver"
	"log"
)

func main() {
	log.Fatal(httpserver.Start())
}
