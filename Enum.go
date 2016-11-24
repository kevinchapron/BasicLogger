package BasicLogger

type LoggerType int;
func (l LoggerType) String() string{
	return STRINGS[int(l)]
}

var STRINGS []string = []string{"","ERROR","WARNING","INFORMATION","DEBUG"}

const ERROR LoggerType = 1
const WARNING LoggerType = 2
const INFO LoggerType = 3
const DEBUG LoggerType = 4

var ALL = []LoggerType{ERROR,WARNING,INFO,DEBUG}