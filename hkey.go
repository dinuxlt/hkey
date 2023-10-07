package main

import (
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/jessevdk/go-flags"
)

type CommandLine struct {
	Message string  `short:"m"  long:"message"  description:"Message to type" required:"true"`
	Wait    int     `short:"w"  long:"wait"     description:"Time in seconds to wait before typing the message" default:"3"`
	Time    float64 `short:"t"  long:"time"     description:"Time in ms between letters" default:"500"`
	Return  bool    `short:"r"  long:"return"   description:"Send ENTER key after message"`
}

var commandLine CommandLine
var parser = flags.NewParser(&commandLine, flags.Default)

func main() {
	_, err := parser.Parse()
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	time.Sleep(time.Duration(commandLine.Wait) * time.Second)

	robotgo.TypeStr(commandLine.Message, commandLine.Time)

	if commandLine.Return {
		robotgo.MilliSleep(int(commandLine.Time))
		robotgo.KeyTap("enter")
	}
}
