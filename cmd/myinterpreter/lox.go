package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ModeInterpret = iota
	ModeRepl
	ModeHelp
	ModeTokenize
	ModeUnknown
)

type Config struct {
	Filename string
	Mode     int
}

type Lox struct {
	HadError bool
}

func main() {
	config := parseArgs()

	if config.Mode == ModeHelp {
		fmt.Fprintln(os.Stderr, "Usage: ")
		fmt.Fprintln(os.Stderr, "\t./golox.sh tokenize <filename>")
		fmt.Fprintln(os.Stderr, "\t./golox.sh # Repl Not implemented yet")
		fmt.Fprintln(os.Stderr, "\t./golox.sh <filename> # Interpret File Not implemented yet")
		os.Exit(1)
	}

	if config.Mode == ModeUnknown {
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", os.Args[1])
		os.Exit(1)
	}

	if config.Mode == ModeRepl {
		runPrompt()
	} else {
		runFile(config)
	}
}

func runFile(config *Config) {
	lox := &Lox{HadError: false}
	fileContents, err := os.ReadFile(config.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	lox.run(string(fileContents))
}

func runPrompt() {
	lox := &Lox{HadError: false}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		var input string
		input, _ = reader.ReadString('\n')
		lox.run(input)
	}
}

func (lox *Lox) run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens(lox)

	for _, token := range tokens {
		fmt.Println(token)
		if lox.HadError {
			os.Exit(65)
		}
	}
}

func (lox *Lox) error(line int, message string) {
	lox.report(line, "", message)
}

func (lox *Lox) report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s\n", line, where, message)
	lox.HadError = true
}

func parseArgs() *Config {
	config := &Config{
		Filename: "",
		Mode:     ModeInterpret,
	}

	if len(os.Args) == 1 {
		config.Mode = ModeRepl
		return config
	}

	if len(os.Args) == 2 {
		if os.Args[2] == "help" {
			return config
		}
		config.Filename = os.Args[1]
		return config
	}

	if len(os.Args) == 3 {
		if os.Args[1] == "tokenize" {
			config.Mode = ModeTokenize
			config.Filename = os.Args[2]
			return config
		} else {
			config.Mode = ModeUnknown
			return config
		}
	}

	return config
}