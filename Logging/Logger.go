package Logging

import (
	"fmt"
	"time"
)

var _Logger_Type_String = []string{"ERROR", "INFO", "WARNING", "DEBUG"}

type LoggerType int

func (l LoggerType) String() string {
	return _Logger_Type_String[int(l)]
}

const ERROR LoggerType = 0
const INFO LoggerType = 1
const WARNING LoggerType = 2
const DEBUG LoggerType = 3

var CurrentLevel LoggerType = 0

type Logger struct {
	LoggerLevel LoggerType
}

func print(Type LoggerType, val ...interface{}) {
	if Type <= CurrentLevel {
		t := time.Now()
		var str string
		for _, elt := range val {
			str += fmt.Sprintf("%v ", elt)
		}
		fmt.Printf("[%s] (%s) %s\n", t.Format("2006-01-02 15:04:05.999999"), Type.String(), str)
	}
}

func SetLevel(Type LoggerType) {
	CurrentLevel = Type
}
func Error(val ...interface{}) {
	print(ERROR, val)
}
func Warning(val ...interface{}) {
	print(WARNING, val)
}
func Info(val ...interface{}) {
	print(INFO, val)
}
func Debug(val ...interface{}) {
	print(DEBUG, val)
}
