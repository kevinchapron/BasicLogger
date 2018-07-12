package main

import (
	"github.com/kevinchapron/BasicLogger/Logging"
)

// Here is the hierarchy of Logging Levels :
// 		DEBUG > INFO > WARNING > ERROR
// 	A higher level includes every inferior levels
//	E.G.: Debug activates Info, Warning & Error
//	E.G.: Warning activates Error
var Log *Logging.Logger = Logging.NewLogger(Logging.DEBUG);

func main(){
	Log.Print(Logging.INFO, "Current Logging Level => ", Log.LoggerLevel.String())

	Log.SetLevel(Logging.ERROR)
	Log.Print(Logging.DEBUG, "I'm a debug info which won't be displayed because Logging Level has been set to \"Error\".")

	Log.SetLevel(Logging.DEBUG)
	Log.Print(Logging.DEBUG, "I'm a debug info which will be displayed !")
}