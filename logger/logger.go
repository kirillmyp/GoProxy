package logger

import (
	"fmt"
	"os"
)

//Logger do somehting
type Logger struct {
	logFile        string
	logDefaultFile string
	logDefaultPath string
	logEnd         []byte
	Enter          func(logInfo []byte, machineNumber string) bool
	doLog          func(logInfo []byte) bool
	getPath        func() string
}

//GetDefaultLogger it's do something
func GetDefaultLogger() *Logger {
	var logger Logger
	logger.logDefaultFile = "./logs/1"
	logger.logDefaultPath = "./logs/"
	logger.logEnd = []byte("\n")
	logger.getPath = func() string {
		var logFile string
		if logger.logFile != "" {
			logFile = logger.logDefaultPath + logger.logFile
		} else {
			logFile = logger.logDefaultFile
		}
		if _, err := os.Stat(logFile); os.IsNotExist(err) {
			_, err := os.Create(logFile)
			if err != nil {
				fmt.Println("Error with create file")
				fmt.Println(err)
				return "Error with create file"
			}
			return logFile
		} else {
			file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0777)
			if err != nil {
				fmt.Println("Error dont find file")
				fmt.Println(err)
				return "Error dont find file"
			}
			defer file.Close()
			return logFile
		}
	}
	logger.doLog = func(logInfo []byte) bool {
		logFile := logger.getPath()
		fmt.Println(logFile)
		fmt.Println(logInfo)
		file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0777)
		defer file.Close()
		if err != nil {
			return false
		}
		n, _ := file.Write(logInfo)
		file.Write(logger.logEnd)
		fmt.Println(n)
		return true
	}
	logger.Enter = func(logInfo []byte, machineNumber string) bool {
		logger.logFile = machineNumber
		return logger.doLog(logInfo)
	}
	return &logger
}
