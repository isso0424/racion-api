package tag

import (
	"encoding/json"
	"io/ioutil"
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
		logger.LoggingError(route, err.Error())
		responser.Fail(
			responser.ErrorPayload{
				Message: "internal server error",
				Status: http.StatusInternalServerError,
			},
			w,
		)

		return
	}

	err = json.Unmarshal(data, &params)
	if err != nil {
		logger.LoggingError(route, err.Error())
		responser.Fail(
			responser.ErrorPayload{
				Message: "internal server error",
				Status: http.StatusInternalServerError,
			},
			w,
		)

		return
	}

	if params.Name == "" || params.Color == "" || params.Description == "" {
		logger.LoggingError(route, "invalid arguments")
		responser.Fail(
			responser.ErrorPayload{
				Message: "invalid arguments",
				Status: http.StatusBadRequest,
			},
			w,
		)

		return
	}

	tag, err := variables.TagController.Create(params.Name, params.Description, params.Color)
	if err != nil {
		logger.LoggingError(route, err.Error())
		responser.Fail(
			responser.ErrorPayload{
				Message: "internal server error",
				Status: http.StatusInternalServerError,
			},
			w,
		)

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
