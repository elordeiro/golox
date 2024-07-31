package lox

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	ModeInterpret = iota
	ModeRepl
	ModeTokenize
	ModeParse
	ModeEvaluate
	ModeHelp
	ModeUnknown
)

type Lox struct {
	HadError        bool
	HadRuntimeError bool
	Mode            int
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
	case ModeEvaluate:
		parser := NewParser(tokens)
		expression := parser.parse()

		if lox.HadError {
			return
		}

		interpreter := NewInterpreter(*lox)
		interpreter.Interpret(expression)
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

func (lox *Lox) RuntimeError(err RuntimeError) {
	fmt.Printf("%v\n[line %d]", err.Message, err.Token.Line)
	lox.HadRuntimeError = true
}

func FormatNumber(num float64) string {
	if math.Floor(num) == num {
		return fmt.Sprintf("%.1f", num)
	}
	return strconv.FormatFloat(num, 'g', -1, 64)
}
