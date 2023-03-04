package hypolaslogger

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

// VarDebug is a function for variable debugging
type VarDebug func(info interface{}, name string)

// HypolasLogger struct for logger
type HypolasLogger struct {
	Info     *log.Logger
	Warn     *log.Logger
	Err      *log.Logger
	Debug    *log.Logger
	VarDebug VarDebug // Debug variable. Print value, name and struct in log file
	LogFile  *os.File
	LogDebug bool // Enable debug variable
}

// NewLogger create logger. PATH is the file path where logs will be store
func NewLogger(path string) HypolasLogger {
	var err error
	var l HypolasLogger
	l = HypolasLogger{
		VarDebug: func(info interface{}, name string) {
			var deb string
			deb = fmt.Sprintf("Name: %s | ", name)
			deb = deb + fmt.Sprintf("Type: %+v | ", reflect.TypeOf(info))

			varType := fmt.Sprintf("%+v", reflect.TypeOf(info))
			if varType == "[]uint8" {
				varValue := string(reflect.ValueOf(info).Bytes())
				deb = deb + fmt.Sprintf("Value: %s | ", varValue)
			} else {
				deb = deb + fmt.Sprintf("Value: %s | ", reflect.ValueOf(info))
			}
			l.Debug.Println(deb)
		},
	}

	l.LogFile, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l.Info = log.New(l.LogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Warn = log.New(l.LogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Err = log.New(l.LogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Debug = log.New(l.LogFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return l
}
