package main

import (
	"flag"
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

	port := flag.Int("port", 8000, "server listen port")
	host := flag.String("host", "localhost", "server listen host")
	flag.Parse()

	log.Printf("Server running on http://%s:%d", *host, *port)
	log.Println(server.Serve(&server.Config{ListenIP: *host, ListenPort: *port}))
}
