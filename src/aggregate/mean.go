package main

import (
	"fmt"
	//"os"
	//"strconv"
	//"log"
	//"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
)



/*func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	datafun.AddChunkOption(program)
	program.Parse()

	sum := 0.0
	count := 0.0

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)

	chunkLines, _ := strconv.ParseInt(program.Opts["chunk"].StringValue, 10, 32)


	readline.ReadLine(fin, func(line string) {
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += val
		count += 1
	})

	fmt.Fprintf(fout, "%g\n", (sum / count))
}*/

type Accumulator struct {
	sum, count float64
}


func main() {
	program := datafun.Init()
	/*datafun.AddInputOutputOptions(program)
	datafun.AddHorizontalOption(program)
	datafun.AddCsvOption(program)
	datafun.AddChunkOption(program)*/
	program.Parse()

	//sum := 0.0
	//count := 0.0

	create := func () interface{} {
		return new(Accumulator)
	}

	output := func (acc interface{}) string {
		val := acc.(*Accumulator)
		return fmt.Sprintf("%g", val.sum / val.count)
	}

	reset := func (acc interface{}) {
		val := acc.(*Accumulator)
		val.sum = 0.0
		val.count = 0.0
	}

	datafun.ProcessEach(program, create, output, reset, func (num float64, acc interface{}) {
		val := acc.(*Accumulator)
		val.count += 1.0
		val.sum += num
	})
}


