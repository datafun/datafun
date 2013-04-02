package main

import (
	"fmt"
	"github.com/datafun/datafun"
)

type Accumulator struct {
	sum float64
}

func main() {
	program := datafun.Init()
	datafun.SetupAggregateDef(program)
	program.Parse()

	create := func () interface{} {
		acc := new(Accumulator)
		acc.sum = 0.0
		return acc
	}

	output := func (acc interface{}) string {
		val := acc.(*Accumulator)
		return fmt.Sprintf("%g", val.sum)
	}

	each := func (num float64, acc interface{}) {
		val := acc.(*Accumulator)
		val.sum += num
	}

	datafun.ProcessEach(program, create, output, each)
}


