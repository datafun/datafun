package main

import (
	"fmt"
	"github.com/datafun/datafun"
	"math"
)

type Accumulator struct {
	current float64
}

func main() {
	program := datafun.Init()
	datafun.SetupMapDef(program)

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


