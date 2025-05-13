package main

import (
	"MONKE/evaluator"
	"MONKE/lexer"
	"MONKE/object"
	"MONKE/parser"
	"MONKE/repl"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) > 1 {
		runFile(os.Args[1])
		return
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the MONKE Programming language! \n", user.Username)
	fmt.Printf("Type commands now!!")
	repl.Start(os.Stdin, os.Stdout)
}

func runFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Could not read file %s: %s\n", filename, err)
		return
	}

	env := object.NewEnvironment()
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Println("Parser error:", msg)
		}
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}
