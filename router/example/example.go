package example

import (
	"isso0424/racion-api/router/logger"
	"net/http"
)

type Example struct{}

func (route Example) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!!!"))

	logger.LoggingInfo(route, "Success!!!")
}

func (route Example) Method() string {
	return "GET"
}

func (route Example) Name() string {
	return "example"
}

func (route Example) Path() string {
	return "/example"
}
