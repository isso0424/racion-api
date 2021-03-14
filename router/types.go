package router

import "net/http"

type Route interface {
	Name() string
	Path() string
	Method() string
	http.Handler
}
