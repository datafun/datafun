package main

import (
	"fmt"
	//"os"
	//"strconv"
	//"log"
	//"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
)

type Accumulator struct {
	sum, count float64
}


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	datafun.AddHorizontalOption(program)
	datafun.AddChunkOption(program)
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


