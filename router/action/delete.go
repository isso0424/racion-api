package action

import (
	"encoding/json"
	"io/ioutil"
	"isso0424/racion-api/router/handler"
	"isso0424/racion-api/router/logger"
	"isso0424/racion-api/router/responser"
	"isso0424/racion-api/router/variables"
	"isso0424/racion-api/types/client_error"
	"net/http"
)

type Delete struct {}

func(route Delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		ID string
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

	action, err := variables.ActionController.Delete(param.ID)
	if err != nil {
		if client_error.IsNotFound(err) {
			handler.HandleError(err.Error(), err.Error(), http.StatusNotFound, route, w)

			return
		}

		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	err = responser.Success(
		responser.DonePayload{
			Data: action,
			Status: http.StatusOK,
			Route: route,
		},
		w,
	)

	if err != nil {
		logger.LoggingError(route, err.Error())

		return
	}
}

func(route Delete) Name() string {
	return "delete action"
}

func(route Delete) Method() string {
	return "DELETE"
}

func(route Delete) Path() string {
	return "/action"
}
