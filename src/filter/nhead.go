package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
)


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	program.OptionWithDefault("-n, --number [count]", "number of lines to skip, defaults to 0", "0")
	
	program.Parse()

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)

	max, _ := strconv.ParseInt(program.Opts["number"].StringValue, 10, 64)
	var count int64 = 0

	readline.ReadLine(fin, func(line string) {
		if (count < max) {
			count += 1 
		} else {
			fmt.Fprintf(fout, "%s\n", line)
		}
	})
}

