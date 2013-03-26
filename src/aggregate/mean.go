package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
)


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	program.Parse()

	sum := 0.0
	count := 0.0

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)

	readline.ReadLine(fin, func(line string) {
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += val
		count += 1
	})

	fmt.Fprintf(fout, "%g", (sum / count))
}
