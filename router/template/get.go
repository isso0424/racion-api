package template

import (
	"isso0424/racion-api/router/handler"
	"isso0424/racion-api/router/logger"
	"isso0424/racion-api/router/responser"
	"isso0424/racion-api/router/variables"
	"isso0424/racion-api/types/client_error"
	"net/http"
)

type Get struct{}

func (route Get) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		Name string
		ID   string
	}
	param := Query{}
	err := variables.Decoder.Decode(&param, r.URL.Query())
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)
		return
	}

	if param.ID != "" && param.Name != "" {
		errMessage := "Cannot use more than one args"
		handler.HandleError(errMessage, errMessage, http.StatusBadRequest, route, w)
	}

	if param.ID != "" {
		template, err := variables.TemplateController.GetByID(param.ID)
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
				Status: http.StatusOK,
				Route:  route,
			},
			w,
		)
		if err != nil {
			logger.LoggingError(route, err.Error())
		}

		return
	}

	if param.Name != "" {
		templates, err := variables.TemplateController.GetByName(param.Name)
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
				Data:   templates,
				Status: http.StatusOK,
				Route:  route,
			},
			w,
		)
		if err != nil {
			logger.LoggingError(route, err.Error())
		}

		return
	}
	templates, err := variables.TemplateController.GetAll()
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
			Data:   templates,
			Status: http.StatusOK,
			Route:  route,
		},
		w,
	)
	if err != nil {
		logger.LoggingError(route, err.Error())
	}
}

func (route Get) Name() string {
	return "get template"
}

func (route Get) Method() string {
	return "GET"
}

func (route Get) Path() string {
	return "/template"
}
