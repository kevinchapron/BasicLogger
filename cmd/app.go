package main

import (
	"github.com/kevinchapron/BasicLogger"
	"fmt"
)

var Logger BasicLogger.Logger;

func main(){
	Logger.Print(fmt.Sprintf("Enabled : %v\n",Logger.GetEnabled()),BasicLogger.INFO)
	Logger.Print(fmt.Sprintf("Disabled : %v\n",Logger.GetDisabled()),BasicLogger.INFO)

	Logger.Print(fmt.Sprintf("Error 1\n"),BasicLogger.ERROR)
	Logger.Disable(BasicLogger.ERROR)
	Logger.Print(fmt.Sprintf("Error 2\n"),BasicLogger.ERROR)
	Logger.Enable(BasicLogger.ERROR)
	Logger.Print(fmt.Sprintf("Error 3\n"),BasicLogger.ERROR)

	Logger.Disable(BasicLogger.DEBUG)

	Logger.Print(fmt.Sprintf("Enabled : %v\n",Logger.GetEnabled()),BasicLogger.INFO)
	Logger.Print(fmt.Sprintf("Disabled : %v\n",Logger.GetDisabled()),BasicLogger.INFO)
}