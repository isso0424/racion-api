package main

import (
	"isso0424/racion-api/mock/repository/action"
	"isso0424/racion-api/mock/repository/tag"
	"isso0424/racion-api/mock/repository/template"
	"isso0424/racion-api/router/variables"
	"isso0424/racion-api/server"
	"log"
)

func main() {
	variables.New(&action.MockActionDB{}, &tag.MockTagDB{}, &template.MockTemplateDB{})
	log.Println(server.Serve(&server.Config{ListenIP: "localhost", ListenPort: 8000}))
}
