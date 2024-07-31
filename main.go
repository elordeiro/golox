package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/lox"
)

type Config struct {
	Filename string
	RunRepl  bool
	Mode     int
}

func main() {
	config := parseArgs()

	if config.Mode == lox.ModeHelp {
		fmt.Fprintln(os.Stderr, "Usage: ")
		fmt.Fprintln(os.Stderr, "\t./golox.sh                     # Repl Mode - Currently default opt produces AST")
		fmt.Fprintln(os.Stderr, "\t./golox.sh tokenize            # Repl Mode - Produces tokens")
		fmt.Fprintln(os.Stderr, "\t./golox.sh tokenize <filename> # Tokenize file")
		fmt.Fprintln(os.Stderr, "\t./golox.sh parse               # Parse Mode - Produces AST")
		fmt.Fprintln(os.Stderr, "\t./golox.sh parse <filename>    # Parse file")
		fmt.Fprintln(os.Stderr, "\t./golox.sh help                # Display this help message")
		os.Exit(1)
	}

	if config.Mode == lox.ModeUnknown {
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", os.Args[1])
		os.Exit(1)
	}

	if config.RunRepl {
		runPrompt(config)
	} else {
		runFile(config)
	}
}

func runFile(config *Config) {
	lox := &lox.Lox{HadError: false, Mode: config.Mode}
	fileContents, err := os.ReadFile(config.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	lox.Run(string(fileContents))
	if lox.HadError {
		os.Exit(65)
	}
	if lox.HadRuntimeError {
		os.Exit(70)
	}
}

func runPrompt(config *Config) {
	lox := &lox.Lox{HadError: false, Mode: lox.ModeParse}
	if config.Mode > 0 {
		lox.Mode = config.Mode
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		var input string
		input, _ = reader.ReadString('\n')
		if input == "exit()\n" {
			os.Exit(1)
		}
		lox.Run(input)
	}
}

func parseArgs() *Config {
	config := &Config{
		Filename: "",
		Mode:     lox.ModeInterpret,
	}

	if len(os.Args) == 1 {
		config.Mode = lox.ModeRepl
		config.RunRepl = true
		return config
	}

	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "tokenize":
			config.Mode = lox.ModeTokenize
		case "parse":
			config.Mode = lox.ModeParse
		case "evaluate":
			config.Mode = lox.ModeEvaluate
		case "help":
			config.Mode = lox.ModeHelp
		default:
			config.Mode = lox.ModeUnknown
		}
		config.RunRepl = true
		return config
	}

	if len(os.Args) == 3 {
		switch os.Args[1] {
		case "tokenize":
			config.Mode = lox.ModeTokenize
		case "parse":
			config.Mode = lox.ModeParse
		case "evaluate":
			config.Mode = lox.ModeEvaluate
		default:
			config.Mode = lox.ModeUnknown
		}
		config.Filename = os.Args[2]
		return config
	}

	config.Mode = lox.ModeUnknown
	return config
}
