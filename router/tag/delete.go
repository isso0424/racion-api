package tag

import (
	"encoding/json"
	"io/ioutil"
	"isso0424/racion-api/router/handler"
	"isso0424/racion-api/router/logger"
	"isso0424/racion-api/router/responser"
	"isso0424/racion-api/router/variables"
	"net/http"
)

type Delete struct {}

func(route Delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		ID string
	}
	params := Param{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	err = json.Unmarshal(data, &params)
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	if params.ID == "" {
		handler.HandleError("invalid arguments", "invalid arguments", http.StatusBadRequest, route, w)

		return
	}

	tag, err := variables.TagController.Delete(params.ID)
	if err != nil {
		handler.HandleError(err.Error(), "target not found", http.StatusNotFound, route, w)

		return
	}

	err = responser.Success(
		responser.DonePayload{
			Data: tag,
			Route: route,
			Status: http.StatusOK,
		},
		w,
	)

	if err != nil {
		logger.LoggingError(route, err.Error())
	}
}

func(route Delete) Name() string {
	return "delete tag"
}

func(route Delete) Method() string {
	return "DELETE"
}

func(route Delete) Path() string {
	return "/tag"
}
