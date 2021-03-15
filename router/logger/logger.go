package logger

import (
	"isso0424/racion-api/logger"
	"isso0424/racion-api/router"
)

func LoggingInfo(route router.Route, message string) {
	logger.LoggingInfo(route.Method(), route.Path(), message)
}

func LoggingWarn(route router.Route, message string) {
	logger.LoggingWarn(route.Method(), route.Path(), message)
}

func LoggingDebug(route router.Route, message string) {
	logger.LoggingDebug(route.Method(), route.Path(), message)
}

func LoggingError(route router.Route, message string) {
	logger.LoggingError(route.Method(), route.Path(), message)
}
