package responser

import "isso0424/racion-api/router"

type ErrorPayload struct {
	Status int
	Message string
}

type DonePayload struct {
	Status int
	Data interface{}
	Route router.Route
}
