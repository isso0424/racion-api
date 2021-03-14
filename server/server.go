package server

import (
	"fmt"
	"isso0424/gorilla-template/router"
	"isso0424/gorilla-template/router/routes"
	"isso0424/gorilla-template/server/midleware"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Serve(c *Config) error {
	r := mux.NewRouter()

	for _, route := range routes.Routes {
		handler := createHandler(route)

		r.Methods(route.Method()).Path(route.Path()).Name(route.Name()).Handler(handler)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", c.ListenPort), handlers.CompressHandler(r))
}

func createHandler(route router.Route) (handler http.Handler) {
	return midleware.RegisterLogger(route)
}
