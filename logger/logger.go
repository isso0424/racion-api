package logger

import (
	"log"
	"os"
)

const Template = "%s %s %s"

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
)

func init() {
	warningLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LoggingWarn(method, path, message string) {
	warningLogger.Printf(Template, method, path, message)
}

func LoggingInfo(method, path, message string) {
	infoLogger.Printf(Template, method, path, message)
}

func LoggingDebug(method, path, message string) {
	debugLogger.Printf(Template, method, path, message)
}

func LoggingError(method, path, message string) {
	errorLogger.Printf(Template, method, path, message)
}
