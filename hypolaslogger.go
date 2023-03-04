package hypolaslogger

import (
	"log"
	"os"
)

// HypolasLogger: struct for logger
type HypolasLogger struct {
	Info    *log.Logger
	Warn    *log.Logger
	Err     *log.Logger
	LogFile *os.File
}

// NewLogger: create logs. PATH is the file path where logs will be store
func NewLogger(path string) HypolasLogger {
	var err error
	l := HypolasLogger{}
	l.LogFile, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l.Info = log.New(l.LogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Warn = log.New(l.LogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Err = log.New(l.LogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return l
}
