package datafun

import (
	"os"
	"fmt"
	"github.com/jprichardson/commander-go"
)

func Init() *commander.Commander {
	return commander.Init(Version)
}

func AddInputOutputOptions(commander *commander.Commander) {
	commander.Option("-i, --input [file]", "input file, defaults to stdin")
	commander.Option("-o, --output [file]", "output file, defaults to stdout")

	inputOption := commander.Opts["input"]
	inputOption.Value = os.Stdin; //default
	inputOption.Callback = func(args ...string) {
		if len(args) > 0 {
			val, err := os.Open(args[0])
			inputOption.Value = val
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Must specify a parameter for -i or --input")
		}
	}

	outputOption := commander.Opts["output"]
	outputOption.Value = os.Stdout; //default
	outputOption.Callback = func(args ...string) {
		if len(args) > 0 {
			val, err := os.Create(args[0])
			outputOption.Value = val
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Must specify a parameter for -i or --input")
		}
	}
}

