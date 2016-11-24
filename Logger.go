package BasicLogger

import "fmt"

type Logger struct {
	disabled []LoggerType
}

func (l *Logger) Disable(Type LoggerType){
	if(l.isDisabled(Type)){	return }
	l.disabled = append(l.disabled,Type)
}
func (l *Logger) Enable(Type LoggerType){
	if(!l.isDisabled(Type)){return }
	var d []LoggerType
	for _, a := range l.disabled{
		if a != Type{
			d = append(d,a)
		}
	}
	l.disabled = d
}

func (l *Logger) Print(s string, Type LoggerType){
	if(!l.isDisabled(Type)){
		fmt.Print(s)
	}
}
func (l *Logger) GetEnabled() []string{
	var s []string
	for i := 0 ; i<len(ALL) ; i++{
		if l.isDisabled(ALL[i]){
			continue
		}
		s = append(s,ALL[i].String());
	}
	return s
}
func (l *Logger) GetDisabled() []string{
	var s []string
	for i := 0 ; i<len(l.disabled) ; i++{
		s = append(s,l.disabled[i].String());
	}
	return s
}


func (l *Logger) isDisabled(Type LoggerType) bool{
	return l.getIndex(Type) != -1
}
func (l *Logger) getIndex(Type LoggerType) int {
	for index, v := range l.disabled{
		if v == Type{
			return index
		}
	}
	return -1
}