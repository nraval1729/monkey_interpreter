package repl

import (
	"../lexer"
	"../token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		instruction, err := reader.ReadString('\n')

		if err != nil {
			panic(fmt.Errorf("repl.Start() threw %v\n", err))
		}
		l := lexer.New(instruction)

		for nt := l.NextToken(); nt.Type != token.EOF; nt = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%+v\n", nt)
		}
	}
}
