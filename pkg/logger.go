package pkg

import (
	"fmt"
	"log"
	"os"
)

var (
	ErrorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(msg string) {
	InfoLog.Println(msg)
}

func Error(msg string, err error) {
	ErrorLog.Printf("%s: %v\n", msg, err)
}


// to easily render HTTP error with their status in the service functions
type HTTPErr struct {
	Status int
	Message string
}

func (e *HTTPErr) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewHTTPErr(status int, msg string) *HTTPErr {
	return &HTTPErr{
		Status: status,
		Message: msg,
	}
}