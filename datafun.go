package datafun

import (
	"os"
	"fmt"
	"path"
	"strconv"
	"log"
	"strings"
	"encoding/csv"
	"github.com/jprichardson/commander-go"
	//"github.com/jprichardson/readline-go"
)

func Init() *commander.Commander {
	return commander.Init(Version)
}

func AddInputOutputOptions(commander *commander.Commander) {
	commander.Option("-i, --input <file>", "input file, defaults to stdin")
	commander.Option("-o, --output <file>", "output file, defaults to stdout")

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
	commander.Option("-c, --chunk <number>", "chunk the input to treat as multiple separate input files with <number> of chunks")
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


func ProcessEach(program *commander.Commander, create func() interface{}, output func(interface{}) string, each func(float64, interface{})) {
	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)
	isHorizontal := program.Opts["horizontal"].Value.(bool)
	chunkCount, _ := strconv.ParseInt(program.Opts["chunk"].StringValue, 10, 64)

	csvReader := csv.NewReader(fin)
	records, err := csvReader.Read()
	var count int64 = 1 

	convertEach := func (str string, acc interface{}) {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatal(err)
		}
		each(val, acc)
	}

	//fmt.Printf("ARGS: %#v", os.Args)
	//fmt.Printf("HORIZONTAL: %#v\n", program.Opts["horizontal"])
	//fmt.Printf("LEN: %d\n", len(records))
	if !isHorizontal {
		accs := make([]interface{}, csvReader.FieldsPerRecord)
		
		initArr := func () {
			for i, _ := range records {
				accs[i] = create()
			}
		}
		initArr() //initialize empty array

		outputResults := func () {
			for i, _ := range accs {
				result := output(accs[i])
				if i < len(accs) - 1 {
					fmt.Fprintf(fout, "%s%c", result, csvReader.Comma)
				} else {
					fmt.Fprintf(fout, "%s\n", result)
				}
			}
		}
		
		for err == nil {
			for i, _ := range records {
				convertEach(records[i], accs[i])
			}

			if (chunkCount > 0) {
				if (count == chunkCount) {
					outputResults()
					initArr()
					count = 0
				}
			}
			
			records, err = csvReader.Read()
			if err == nil {
				count += 1
			}
		}

		if count > 0 { //output leftovers
			outputResults()
		}
		
	} else { //going horizontal... across... by columns
		for err == nil {
			acc := create()
			count = 0 //here, we're counting columns instead of lines (rows)
			outp := ""

			for i, _ := range records {
				convertEach(records[i], acc)
				count +=1 

				if (chunkCount > 0) {
					if (count == chunkCount) {
						result := output(acc)
						outp += result + ","
						count = 0
						acc = create()
					}
				}
			}

			if (chunkCount > 0) {
				//left overs
				if (count > 0) {
					result := output(acc)
					outp += result + ","
				}

				outp = strings.TrimRight(outp, ",")
				fmt.Fprintf(fout, "%s\n", outp)
			} else {
				result := output(acc)
				fmt.Fprintf(fout, "%s\n", result)
			}

			records, err = csvReader.Read()
		}
	}

}



