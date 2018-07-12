package Logging

import (
	"fmt"
	"time"
)

var _Logger_Type_String = []string{"ERROR","INFO","WARNING","DEBUG"}

type LoggerType int
func (l LoggerType) String() string{
	return _Logger_Type_String[int(l)]
}

const ERROR LoggerType = 0
const INFO LoggerType = 1
const WARNING LoggerType = 2
const DEBUG LoggerType = 3

var _Logger_Instance *Logger

type Logger struct {
	LoggerLevel LoggerType
}

func GetLogger() *Logger{
	return _Logger_Instance
}

func NewLogger(Type LoggerType) *Logger{
	if _Logger_Instance == nil {
		_Logger_Instance = &Logger{
			LoggerLevel:Type,
		}
	}
	return _Logger_Instance
}

func (l *Logger) Print(Type LoggerType, val... interface{}){
	if Type <= l.LoggerLevel{
		t := time.Now()
		var str string
		for _, elt := range val{
			str += fmt.Sprintf("%v ",elt)
		}
		fmt.Printf("[%s] (%s) %s\n",t.Format("2006-01-02 15:04:05.999999"),Type.String(),str)
	}
}

func (l *Logger) SetLevel(Type LoggerType){
	l.LoggerLevel = Type
}
