package main

import (
	Log "github.com/kevinchapron/BasicLogger/Logging"
)

// Here is the hierarchy of Logging Levels :
// 		DEBUG > INFO > WARNING > ERROR
// 	A higher level includes every inferior levels
//	E.G.: Debug activates Info, Warning & Error
//	E.G.: Warning activates Error

func main() {

	Log.Info("Current Logging Level => ", Log.CurrentLevel)

	Log.SetLevel(Log.ERROR)
	Log.Debug("I'm a debug info which won't be displayed because Logging Level has been set to \"Error\".")

	Log.SetLevel(Log.DEBUG)
	Log.Debug("I'm a debug info which will be displayed !")
}
