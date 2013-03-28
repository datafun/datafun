package main

import (
	"fmt"
	//"github.com/jprichardson/commander-go"
	"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
	"log"
	"os"
	"strconv"
	"encoding/csv"
)


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	datafun.AddHorizontalOption(program)
	datafun.AddCsvOption(program)
	program.Parse()

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)
	//isHorizontal := program.Opts["horizontal"].Value.(bool)
	isCsv := program.Opts["csv"].Value.(bool) 

	iProduct := 1.0

	if !isCsv {
		readline.ReadLine(fin, func(line string) {
			val, err := strconv.ParseFloat(line, 64)
			if err != nil {
				log.Fatal(err)
			}
			iProduct *= val
		})
		fmt.Fprintf(fout, "%g\n", iProduct)
	} else {
		csvReader := csv.NewReader(fin)
		records, err := csvReader.Read()

		products := make([]float64, csvReader.FieldsPerRecord)
		//initialize
		for i, _ := range products {
			products[i] = 1.0
		}

		for err == nil {
			for i, _ := range records {
				val, e := strconv.ParseFloat(records[i], 64)
				if e != nil {
					log.Fatal(e)
				}
				products[i] *= val
			}

			records, err = csvReader.Read()
		}

		for i, _ := range products {
			if i < len(products) - 1 {
				fmt.Fprintf(fout, "%g,", products[i])
			} else {
				fmt.Fprintf(fout, "%g\n", products[i])
			}
		}
	}
}
