package tag

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
		errMessage := "Cannot use more than one args"
		handler.HandleError(errMessage, errMessage, http.StatusBadRequest, route, w)

		return
	}

	if query.ID != "" {
		tag, err := variables.TagController.GetByID(query.ID)
		if err != nil {
			if client_error.IsNotFound(err) {
				handler.HandleError(err.Error(), err.Error(), http.StatusNotFound, route, w)

				return
			}
			handler.HandleError(err.Error(), "not found", http.StatusNotFound, route, w)

			return
		}

		err = responser.Success(
			responser.DonePayload{
				Status: http.StatusOK,
				Data:   tag,
				Route:  route,
			},
			w,
		)
		if err != nil {
			logger.LoggingError(route, err.Error())
		}

		return
	}

	if query.Title != "" {
		tags, err := variables.TagController.GetByTitle(query.Title)
		if err != nil {
			if client_error.IsNotFound(err) {
				handler.HandleError(err.Error(), err.Error(), http.StatusNotFound, route, w)

				return
			}
			handler.HandleError(err.Error(), "not found", http.StatusNotFound, route, w)

			return
		}

		err = responser.Success(
			responser.DonePayload{
				Status: http.StatusOK,
				Data:   tags,
				Route:  route,
			},
			w,
		)
		if err != nil {
			logger.LoggingError(route, err.Error())
		}

		return
	}

	tags, err := variables.TagController.GetAll()
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
			Data:   tags,
			Route:  route,
		},
		w,
	)
	if err != nil {
		logger.LoggingError(route, err.Error())

		return
	}
}

func (route Get) Name() string {
	return "get tags"
}

func (route Get) Path() string {
	return "/tag"
}

func (route Get) Method() string {
	return "GET"
}
