package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/lox"
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
	lox := &lox.Lox{HadError: false}
	fileContents, err := os.ReadFile(config.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	lox.Run(string(fileContents))
	if lox.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	lox := &lox.Lox{HadError: false}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		var input string
		input, _ = reader.ReadString('\n')
		lox.Run(input)
	}
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
