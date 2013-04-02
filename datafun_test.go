package datafun

import (
	"testing"
	"path"
	//"io"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	"github.com/jprichardson/goatee-go"
	"github.com/jprichardson/readline-go"
)

func TestAddInputOutputOptionsExist (*testing.T) {
	return;
	program := Init()
	
	_, exists := program.Opts["input"]
	t.F (exists)
	_, exists = program.Opts["output"]
	t.F (exists)
	
	AddInputOutputOptions(program)

	_, exists = program.Opts["input"]
	t.T (exists)
	_, exists = program.Opts["output"]
	t.T (exists)
}

func TestAddInputOutputOptionsInput (*testing.T) {
	program := Init()
	
	AddInputOutputOptions(program)

	tempFile := path.Join(os.TempDir(), "temp-input")
	ioutil.WriteFile(tempFile, []byte("line1\nline2"), 0644)
	//println(tempFile)

	os.Args = []string{"program", "-i", tempFile}
	program.Parse()

	expect := ""
	readline.ReadLine(program.Opts["input"].Value.(*os.File), func (line string) {
		expect += line + "\n"
	})

	t.EQ ("line1\nline2", strings.TrimRight(expect, "\n"))
}

func TestAddInputOutputOptionsInputDefault (*testing.T) {
	program := Init()
	
	AddInputOutputOptions(program)


	os.Args = []string{"program"}
	program.Parse()

	t.EQ (program.Opts["input"].Value, os.Stdin)
}

func TestAddInputOutputOptionsOutput (*testing.T) {
	program := Init()
	
	AddInputOutputOptions(program)

	tempFile := path.Join(os.TempDir(), "temp-output")

	os.Args = []string{"program", "-o", tempFile}
	program.Parse()

	fo := program.Opts["output"].Value.(*os.File)
	fmt.Fprintf(fo, "line4\nline5")
	fo.Close()

	data, _ := ioutil.ReadFile(tempFile)

	t.EQ ("line4\nline5", string(data))
}

func TestAddInputOutputOptionsOutputDefault (*testing.T) {
	program := Init()
	
	AddInputOutputOptions(program)

	os.Args = []string{"program"}
	program.Parse()

	t.EQ (program.Opts["output"].Value, os.Stdout)
}

func TestAddChunkOption (*testing.T) {
	program := Init()
	
	AddChunkOption(program)

	os.Args = []string{"program", "-c", "500"}
	program.Parse()

	t.EQ (program.Opts["chunk"].StringValue, "500")
}


