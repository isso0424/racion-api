package main

import (
	"isso0424/gorilla-template/server"
	"log"
)

func main() {
	log.Println(server.Serve(&server.Config{ListenIP: "localhost", ListenPort: 8000}))
}
