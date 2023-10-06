package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/JulienR1/monkey/internal/pkg/lexer"
	"github.com/JulienR1/monkey/internal/pkg/token"
)

const PROMPT = ">>"

func Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for t := lexer.NextToken(); t.Type != token.EOF; t = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}
