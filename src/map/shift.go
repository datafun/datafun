package main

import (
	"fmt"
	"strconv"
	"github.com/datafun/datafun"
)

type Accumulator struct {
	current float64
}

func main() {
	program := datafun.Init()
	datafun.SetupMapDef(program)
	
	program.Option("-n, --number <num>", "the number to shift by")
	program.Parse()

	shift, _ := strconv.ParseFloat(program.Opts["number"].StringValue, 64)

	create := func () interface{} {
		return new(Accumulator)
	}

	output := func (acc interface{}) string {
		val := acc.(*Accumulator)
		return fmt.Sprintf("%g", val.current)
	}

	each := func (num float64, acc interface{}) {
		val := acc.(*Accumulator)
		val.current = num + shift
	}

	datafun.ProcessEach(program, create, output, each)
}


