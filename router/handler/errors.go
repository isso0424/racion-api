package handler

import (
	"isso0424/racion-api/router"
	"isso0424/racion-api/router/logger"
	"isso0424/racion-api/router/responser"
	"net/http"
)

func HandleError(
	serverMessage,
	clientMessage string,
	status int,
	route router.Route,
	writer http.ResponseWriter,
) {
	logger.LoggingError(route, serverMessage)
	err := responser.Fail(
		responser.ErrorPayload{
			Message: clientMessage,
			Status:  status,
		},
		writer,
	)

	if err != nil {
		logger.LoggingError(route, err.Error())
	}
}
