package action

import "net/http"

type Create struct {}

func(route Create) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func(route Create) Name() string {
	return "create action"
}

func(route Create) Method() string {
	return "POST"
}

func(route Create) Path() string {
	return ""
}
