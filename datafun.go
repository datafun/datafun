package datafun

import (
	"os"
	"fmt"
	"path"
	"strconv"
	"log"
	//"encoding/csv"
	"github.com/jprichardson/commander-go"
	"github.com/jprichardson/readline-go"
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
	commander.Option("--csv", "specify CSV input, triggered by default if input filename ends in .csv")

	option := commander.Opts["csv"]
	option.Value = false
	option.Callback = func(args ...string) {
		option.Value = true
	}
}

func AddChunkOption(commander *commander.Commander) {
	commander.Option("-c, --chunk <number>", "chunk the input to treat as multiple separate input files with <number> of lines")
	option := commander.Opts["chunk"]
	option.Value = false
	option.Callback = func(args ...string) {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "Must specify a parameter for -c or --chunk"); return;
		}
		option.Value = args[0]
		option.StringValue = args[0]
	}
}


func ProcessEach(commander *commander.Commander, create func() interface{}, output func(interface{}) string, reset func(interface{}), each func(float64, interface{})) {
	AddInputOutputOptions(commander)
	AddHorizontalOption(commander)
	AddCsvOption(commander)
	AddChunkOption(commander)

	fin := commander.Opts["input"].Value.(*os.File)
	fout := commander.Opts["output"].Value.(*os.File)
	//isHorizontal := commander.Opts["horizontal"].Value.(bool)
	isCsv := commander.Opts["csv"].Value.(bool) 
	//chunkLines, _ := strconv.ParseInt(commander.Opts["chunk"].StringValue, 10, 32)

	if !isCsv {
		acc := create()
		readline.ReadLine(fin, func(line string) {
			val, err := strconv.ParseFloat(line, 64)
			if err != nil {
				log.Fatal(err)
			}
			each(val, acc)
		})

		result := output(acc)
		fmt.Fprintf(fout, "%s\n", result)
	} else {

	}
}



