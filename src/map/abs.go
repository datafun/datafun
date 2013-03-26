package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
	"math"
)


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	program.Parse()

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)

	readline.ReadLine(fin, func(line string) {
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(fout, "%g", math.Abs(val))
	})	
}
