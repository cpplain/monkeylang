package main

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func startRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := newLexer(line)

		for tok := l.nextToken(); tok.Type != EOF; tok = l.nextToken() {
			fmt.Fprint(out, tok, "\n")
		}
	}
}
