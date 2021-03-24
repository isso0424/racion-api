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

type Put struct {}

func(route Put) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		ID string
		Title string
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

	if params.Title == "" || params.Color == "" || params.Description == "" {
		handler.HandleError("invalid arguments", "invalid arguments", http.StatusBadRequest, route, w)

		return
	}

	tag, err := variables.TagController.Edit(params.ID, params.Title, params.Description, params.Color)
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

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

func(route Put) Name() string {
	return "update tag"
}

func(route Put) Method() string {
	return "PUT"
}

func(route Put) Path() string {
	return "/tag"
}
