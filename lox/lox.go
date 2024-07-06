package lox

import (
	"fmt"
	"os"
)

type Lox struct {
	HadError bool
}

func (lox *Lox) Run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(lox)

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func (lox *Lox) Error(line int, message string) {
	lox.report(line, "", message)
}

func (lox *Lox) report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	lox.HadError = true
}
