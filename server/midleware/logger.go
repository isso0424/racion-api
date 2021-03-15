package midleware

import (
	"isso0424/racion-api/logger"
	"isso0424/racion-api/router"
	"net/http"
)

func RegisterLogger(route router.Route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.LoggingInfo(route.Method(), route.Path(), route.Name())

		route.ServeHTTP(w, r)
	})
}
