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

type TagCreating struct {}

func(route TagCreating) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		Name string
		Color string
		Description string
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

	if params.Name == "" || params.Color == "" || params.Description == "" {
		handler.HandleError("invalid arguments", "invalid arguments", http.StatusBadRequest, route, w)

		return
	}

	tag, err := variables.TagController.Create(params.Name, params.Description, params.Color)
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	err = responser.Success(
		responser.DonePayload{
			Data: tag,
			Route: route,
			Status: http.StatusCreated,
		},
		w,
	)

	if err != nil {
		logger.LoggingError(route, err.Error())
	}
}

func(route TagCreating) Name() string {
	return "create tag"
}

func(route TagCreating) Method() string {
	return "POST"
}

func(route TagCreating) Path() string {
	return "/tag"
}