package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/0xsuk/golox/parse_error"
	"github.com/0xsuk/golox/parser"
	"github.com/0xsuk/golox/runtime_error"
	"github.com/0xsuk/golox/scanner"
	"github.com/0xsuk/golox/semantic_error"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func runFile(file string) {
	dat, err := ioutil.ReadFile(file)
	check(err)
	run(string(dat))
	if parse_error.HadError {
		os.Exit(65)
	} else if runtime_error.HadError || semantic_error.HadError {
		os.Exit(70)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		dat, err := reader.ReadBytes('\n')
		check(err)
		run(string(dat))
		parse_error.HadError = false
		runtime_error.HadError = false
		semantic_error.HadError = false
	}
}

func run(src string) {
	scanner := scanner.New(src)
	tokens := scanner.ScanTokens()

	fmt.Println("Tokens:")
	for _, token := range tokens {
		fmt.Println("\t" + token.String())
	}

	parser := parser.New(tokens)
	_ = parser.Parse()
	if parse_error.HadError {
		return
	}
}

func main() {
	flag.String("file", "", "the script file to execute")
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		fmt.Println("usage: ./golox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}
