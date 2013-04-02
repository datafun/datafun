package main

import (
	"fmt"
	"github.com/datafun/datafun"
)

type Accumulator struct {
	product float64
}

func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	datafun.AddHorizontalOption(program)
	datafun.AddChunkOption(program)
	program.Parse()

	create := func () interface{} {
		acc := new(Accumulator)
		acc.product = 1.0
		return acc
	}

	output := func (acc interface{}) string {
		val := acc.(*Accumulator)
		return fmt.Sprintf("%g", val.product)
	}

	each := func (num float64, acc interface{}) {
		val := acc.(*Accumulator)
		val.product *= num
	}

	datafun.ProcessEach(program, create, output, each)
}


