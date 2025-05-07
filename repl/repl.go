package repl

import (
	"MONKE/evaluator"
	"MONKE/lexer"
	"MONKE/object"
	"MONKE/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKE = `          __,__         
 .--.  .-"     "-. .--.
/ .. \/ .-. .-.  \/ ..  \
| |  '| /   Y   \  |'  | |
| \   \ \ 0 | 0 /  /   / |
\ '- ,\.-"""""""-./, -' /
 ''-' /_   ^ ^   _\ '-''
     |  \._   _./  |
     \   \ '~' /   /
      '._ '-=-' _.'
         '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
