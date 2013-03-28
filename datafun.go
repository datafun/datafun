package datafun

import (
	"os"
	"fmt"
	"path"
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
			fileName := args[0]
			if (path.Ext(fileName) == ".csv") {
				//hacky
				AddCsvOption(commander)
				commander.Opts["csv"].Value = true
			}

			val, err := os.Open(fileName)
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


func AddHorizontalOption(commander *commander.Commander) {
	commander.Option("-x, --horizontal", "perform operation horizontally instead of vertically")

	horizontalOption := commander.Opts["horizontal"]
	horizontalOption.Value = false
	horizontalOption.Callback = func(args ...string) {
		horizontalOption.Value = true
	}
}

func AddCsvOption(commander *commander.Commander) {
	commander.Option("-c, --csv", "specify CSV input, triggered by default if input filename ends in .csv")

	option := commander.Opts["csv"]
	option.Value = false
	option.Callback = func(args ...string) {
		option.Value = true
	}
}

