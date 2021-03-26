package main

import (
	"isso0424/racion-api/mock/repository/action"
	"isso0424/racion-api/mock/repository/tag"
	"isso0424/racion-api/mock/repository/template"
	"isso0424/racion-api/router/variables"
	"isso0424/racion-api/server"
	"isso0424/racion-api/types/domain"
	"log"
)

func main() {
	variables.New(
		&action.MockActionDB{Data: make([]domain.Action, 0)},
		&tag.MockTagDB{Data: make([]domain.Tag, 0)},
		&template.MockTemplateDB{Data: make([]domain.Template, 0)},
	)
	log.Println(server.Serve(&server.Config{ListenIP: "localhost", ListenPort: 8000}))
}
