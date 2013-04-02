package main

import (
	"fmt"
	"github.com/datafun/datafun"
)

type Accumulator struct {
	quotient float64
}

func main() {
	program := datafun.Init()
	datafun.SetupAggregateDef(program)
	program.Parse()

	create := func () interface{} {
		acc := new(Accumulator)
		acc.quotient = 0.0
		return acc
	}

	output := func (acc interface{}) string {
		val := acc.(*Accumulator)
		return fmt.Sprintf("%g", val.quotient)
	}

	each := func (num float64, acc interface{}) {
		val := acc.(*Accumulator)
		if val.quotient == 0 { //obviously can't be 0, so it's the beginning
			val.quotient = num
		} else {
			val.quotient /= num
		}
	}

	datafun.ProcessEach(program, create, output, each)
}


