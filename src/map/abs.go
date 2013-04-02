package main

import (
	"fmt"
	"github.com/datafun/datafun"
	"math"
	"os"
)

type Accumulator struct {
	current float64
}

func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	datafun.AddHorizontalOption(program)
	datafun.AddChunkOption(program)

	//don't show on --help
	program.Opts["horizontal"].Hidden = true
	program.Opts["chunk"].Hidden = true

	//force ProcessEach to iterate and output at every element
	os.Args = append(os.Args, "-x", "-c", "1")

	program.Parse()

	create := func () interface{} {
		return new(Accumulator)
	}

	output := func (acc interface{}) string {
		val := acc.(*Accumulator)
		return fmt.Sprintf("%g", val.current)
	}

	each := func (num float64, acc interface{}) {
		val := acc.(*Accumulator)
		val.current = math.Abs(num)
	}

	datafun.ProcessEach(program, create, output, each)
}


