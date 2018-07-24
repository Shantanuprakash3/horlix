package logger

import (
	"bufio"
	"log"
	"os"
)

var aLogger *log.Logger
var appLog *os.File

func InitAppLogger(appLogDir string) error {
	var err error
	appLog, err = os.Create(appLogDir + "/horlix.log")
	if err != nil {
		return err
	}
	buf := bufio.NewWriterSize(appLog, 10)
	aLogger = log.New(buf, "", log.Lshortfile)
	return nil
}

func LogInfo(msg string) {
	aLogger.Println(msg)
}

func LogInfof(format string, v ...interface{}) {
	aLogger.Printf(format, v...)
}

func LogFatal(msg string) {
	aLogger.Fatalln(msg)
}

func LogFatalf(format string, v ...interface{}) {
	aLogger.Fatalf(format, v)
}
