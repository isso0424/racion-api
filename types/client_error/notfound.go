package client_error

import "fmt"

type notFound struct {
	Target    string
	Type      string
	Condition string
}

func (err notFound) Error() string {
	return fmt.Sprintf("%s(%s, %s) not found", err.Type, err.Target, err.Condition)
}

func CreateNotFound(t, target, condition string) notFound {
	return notFound{Target: target, Type: t, Condition: condition}
}

func IsNotFound(err error) bool {
	errMsg := err.Error()
	length := len(errMsg)
	if length < 9 {
		return false
	}

	return errMsg[length-9:length] == "not found"
}
