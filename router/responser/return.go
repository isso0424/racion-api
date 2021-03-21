package responser

import (
	"encoding/json"
	"isso0424/racion-api/router/logger"
	"net/http"
)

func Success(payload DonePayload, writer http.ResponseWriter) (err error) {
	result, err := json.Marshal(payload.Data)
	if err != nil {
		Fail(
			ErrorPayload{
				Status: http.StatusInternalServerError,
				Message: "Value parsing error",
			},
			writer,
		)
		return
	}
	logger.LoggingInfo(payload.Route, string(result))
	return returnValue(result, payload.Status, writer)
}

func Fail(payload ErrorPayload, writer http.ResponseWriter) (err error) {
	result, err := json.Marshal(map[string]string{ "message": payload.Message })
	if err != nil {
		Fail(
			ErrorPayload{
				Status: http.StatusInternalServerError,
				Message: "Value parsing error",
			},
			writer,
		)
		return
	}
	return returnValue(result, payload.Status, writer)
}

func returnValue(data []byte, status int, writer http.ResponseWriter) (err error) {
	writer.WriteHeader(status)
	_, err = writer.Write(data)

	return err
}
