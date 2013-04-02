package main

import (
	"fmt"
	"github.com/datafun/datafun"
)

type Accumulator struct {
	sum, count float64
}

func main() {
	program := datafun.Init()
	datafun.SetupAggregateDef(program)
	
	program.Parse()

	create := func () interface{} {
		return new(Accumulator)
	}

	output := func (acc interface{}) string {
		val := acc.(*Accumulator)
		return fmt.Sprintf("%g", val.sum / val.count)
	}

	each := func (num float64, acc interface{}) {
		val := acc.(*Accumulator)
		val.count += 1.0
		val.sum += num
	}

	datafun.ProcessEach(program, create, output, each)
}


