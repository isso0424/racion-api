package action

import (
	"encoding/json"
	"io/ioutil"
	"isso0424/racion-api/router/handler"
	"isso0424/racion-api/router/responser"
	"isso0424/racion-api/router/variables"
	"isso0424/racion-api/types/client_error"
	"net/http"
	"time"
)

type Update struct {}

func(route Update) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		Title string
		Tags []string
		Color string
		ID string
		StartAt int64
		EndAt int64
	}
	param := Param{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	err = json.Unmarshal(data, &param)
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	if param.ID == "" || param.Title == "" || param.Color == "" || param.StartAt == 0 || param.EndAt == 0 {
		handler.HandleError("invalid arguments", "invalid arguments", http.StatusBadRequest, route, w)

		return
	}

	action, err := variables.ActionController.Edit(param.ID, param.Title, param.Color, param.Tags, time.Unix(param.StartAt, 0), time.Unix(param.EndAt, 0))
	if err != nil {
		if client_error.IsNotFound(err) {
			handler.HandleError(err.Error(), err.Error(), http.StatusBadRequest, route, w)

			return
		}
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	err = responser.Success(
		responser.DonePayload{
			Data: action,
			Route: route,
			Status: http.StatusCreated,
		},
		w,
	)
	if err != nil {
		logger.LoggingError(route, err.Error())
	}
}

func(route Update) Method() string {
	return "PUT"
}

func(route Update) Path() string {
	return "/action"
}

func(route Update) Name() string {
	return "Update action"
}
