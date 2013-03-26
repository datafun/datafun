package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/jprichardson/readline-go"
	"github.com/datafun/datafun"
	"path"
	"log"
	//"strings"
)


func main() {
	program := datafun.Init()
	datafun.AddInputOutputOptions(program)
	program.OptionWithDefault("-n, --number [lines]", "number of lines to split by, defaults to 1000", "1000")
	
	program.Parse()

	fin := program.Opts["input"].Value.(*os.File)
	fout := program.Opts["output"].Value.(*os.File)
	
	program.Parse()

	max, _ := strconv.ParseInt(program.Opts["number"].StringValue, 10, 64)
	var count int64 = 0

	fileName := os.Args[len(os.Args) - 1]
	var currentFileCount int64 = 0
	newFileName := generateFileName(fileName, currentFileCount)
	newFileFd, err := os.Create(newFileName)

	if err != nil {
		log.Fatal(err)
	}

	readline.ReadLine(fin, func(line string) {
		fmt.Fprintf(newFileFd, "%s\n", line)
		
		count += 1
		if (count == max) {
			count = 0
			defer newFileFd.Close()

			fmt.Fprintf(fout, "%s\n", newFileName) //output file name of generated file

			currentFileCount += 1
			newFileName = generateFileName(fileName, currentFileCount)
			newFileFd, err = os.Create(newFileName)
		}
	})

	if (count > 0) { //odd number, didn't get printed out
		fmt.Fprintf(fout, "%s\n", newFileName)
		defer newFileFd.Close()
	}
}

func generateFileName(fileName string, fileCount int64) string {
	base := path.Base(fileName)
	dir := path.Dir(fileName)
	newFileName := path.Join(dir, fmt.Sprintf("%d-%s", fileCount, base))
	return newFileName
}
