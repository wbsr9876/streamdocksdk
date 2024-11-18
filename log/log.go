package log

import (
	"fmt"
	"log"
)

type Writer interface {
	LogMessage(string)
}

var connection Writer

func SetConnection(conn Writer) {
	connection = conn
}

func Message(format string, args ...interface{}) {
	if format == "" {
		return
	}
	message := fmt.Sprintf(format, args...)
	if connection == nil {
		log.Println(message)
		return
	}
	connection.LogMessage(message)
}
