package main

import (
	"os"
	"fmt"
	"strconv"
	"log"
	"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
)


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	program.OptionWithDefault("-n, --number [count]", "number of points to insert between each point", "0")
	
	program.Parse()

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)

	n, _ := strconv.ParseFloat(program.Opts["number"].StringValue, 64)
	var count int64 = 0

	lastNum := 0.0

	readline.ReadLine(fin, func(line string) {
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatal(err)
		}

		if count > 0 {
			base := lastNum
			step := (num - lastNum) / (n + 1)

			for i := base; i < num; i += step {
				//fmt.Printf("base: %g, lastnum: %g, num: %g, n: %g, i: %g\n", base, lastNum, num, n, i)
				fmt.Fprintf(fout, "%g\n", i)
			}

			//fmt.Fprintf(fout, "%g\n", num)
			lastNum = num
		}

		count += 1 
		lastNum = num
	})

	fmt.Fprintf(fout, "%g\n", lastNum)
}

