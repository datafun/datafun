package main

import (
	"fmt"
	//"github.com/jprichardson/commander-go"
	"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
	"log"
	"os"
	"strconv"
)


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	program.Parse()

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)

	var sum = 0.0

	readline.ReadLine(fin, func(line string) {
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += val
	})

	fmt.Fprintf(fout, "%g\n", sum)
}
