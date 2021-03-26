package template

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

type Create struct{}

func (route Create) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Param struct {
		Tags  []string
		Name  string
		Color string
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

	if param.Name == "" || param.Color == "" {
		handler.HandleError("invalid arguments", "invalid arguments", http.StatusBadRequest, route, w)

		return
	}

	template, err := variables.TemplateController.Create(param.Name, param.Color, param.Tags)
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
			Data:   template,
			Status: http.StatusCreated,
			Route:  route,
		},
		w,
	)
	if err != nil {
		logger.LoggingError(route, err.Error())
	}
}

func (route Create) Name() string {
	return "create template"
}

func (route Create) Path() string {
	return "/template"
}

func (route Create) Method() string {
	return "POST"
}
