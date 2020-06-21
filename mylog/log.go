package mylog

import (
	"log"
	"os"
)

var myLogger *log.Logger

func GetLogger() *log.Logger {
	if myLogger == nil {
		myLogger = log.New(os.Stdout, "[chat]", log.Ldate|log.Ltime)
	}

	return myLogger
}
