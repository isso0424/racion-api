package action

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
		Title string
		ID    string
	}
	query := Query{}
	err := variables.Decoder.Decode(&query, r.URL.Query())
	if err != nil {
		handler.HandleError(err.Error(), "internal server error", http.StatusInternalServerError, route, w)

		return
	}

	if query.ID != "" && query.Title != "" {
		handler.HandleError("Can use one or less args", "Can use one or less args", http.StatusInternalServerError, route, w)

		return
	}

	if query.ID != "" {
		action, err := variables.ActionController.GetByID(query.ID)
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
				Status: http.StatusOK,
				Route:  route,
				Data:   action,
			},
			w,
		)

		if err != nil {
			logger.LoggingError(route, err.Error())
		}

		return
	}

	if query.Title != "" {
		action, err := variables.ActionController.GetByTitle(query.Title)
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
				Status: http.StatusOK,
				Route:  route,
				Data:   action,
			},
			w,
		)

		if err != nil {
			logger.LoggingError(route, err.Error())
		}

		return
	}

	action, err := variables.ActionController.GetAll()
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
			Status: http.StatusOK,
			Route:  route,
			Data:   action,
		},
		w,
	)

	if err != nil {
		logger.LoggingError(route, err.Error())
	}
}

func (route Get) Method() string {
	return "GET"
}

func (route Get) Name() string {
	return "Get actions"
}

func (route Get) Path() string {
	return "/action"
}
