package hypolaslogger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
func NewLogger(pathToLogFile string) *HypolasLogger {
	if pathToLogFile == "" {
		pathToLogFile = os.Getenv("HYPOLAS_LOGS_FILE")
	}

	if pathToLogFile == "" {
		log.Println("No path defined for NewLogger(pathToLogFile string).")
		log.Fatalln("Define environnement variable HYPOLAS_LOGS_FILE or call function with pathToLogFile not empty.")
	}

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

	// Create logs directory if not exist
	createLogsFolder(filepath.Dir(pathToLogFile))

	l.LogFile, err = os.OpenFile(pathToLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l.Info = log.New(l.LogFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Warn = log.New(l.LogFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Err = log.New(l.LogFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	l.Debug = log.New(l.LogFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &l
}

func createLogsFolder(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println("ERROR create directory:", err)
		}
	}
}
