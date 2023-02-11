package golog

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// LogMessage() - Takes a log level, which is then compared to the configured log level. If the level provided
// is lower than or equal to the configured log level, the message is logged.
func (g GoLogger) LogMessage(messageLogLevel int, message ...interface{}) {
	if messageLogLevel <= g.LogLevel {
		fmt.Println(message...)
	}
}

// CreateError() - Returns an error with a stack trace
func (e ErrorLogger) CreateError(message ...interface{}) error {
	return errors.New(fmt.Sprint(message...))
}

// LogError() - Log error, regardless of logging level. Include stack trace. Provide the error type - ERROR or FATAL.
// ERROR will allow app to continue, FATAL will cause app to shut down.
func (e ErrorLogger) LogError(errorType string, err error) {
	var finalError string = err.Error() + "\n"

	if err, ok := err.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			str := fmt.Sprintf("%+s", f)

			// Don't include golog or runtime frame in the stack trace
			if strings.Contains(str, "golog") || strings.Contains(str, "runtime") {
				continue
			}

			// Add new line if error string doesn't end in one
			if !strings.HasSuffix(str, "\n") {
				str += "\n"
			}

			finalError += str
		}
	}

	if errorType == ERROR {
		log.Print(finalError)
	} else if errorType == FATAL {
		log.Fatal(finalError)
	} else {
		log.Printf("%s%s\n", "Invalid errorType:", errorType)
	}
}
