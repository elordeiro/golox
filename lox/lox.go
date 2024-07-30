package lox

import (
	"fmt"
	"os"
)

const (
	ModeInterpret = iota
	ModeRepl
	ModeHelp
	ModeTokenize
	ModeParse
	ModeUnknown
)

type Lox struct {
	HadError bool
	Mode     int
}

func (lox *Lox) Run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(lox)

	switch lox.Mode {
	case ModeTokenize:
		for _, token := range tokens {
			fmt.Println(token)
		}
	case ModeParse:
		parser := NewParser(tokens)
		expression := parser.parse()

		if lox.HadError {
			return
		}

		fmt.Println(PrintAst(expression))
	}
}

func (lox *Lox) Error(line int, message string) {
	lox.report(line, "", message)
}

func (lox *Lox) report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	lox.HadError = true
}

func (lox *Lox) ErrorToken(token Token, message string) {
	if token.Type == EOF {
		lox.report(token.Line, " at end", message)
	} else {
		lox.report(token.Line, " at '"+token.Lexeme+"'", message)
	}
}
