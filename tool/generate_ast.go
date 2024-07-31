package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate_ast <output directory>")
		os.Exit(64)
	}
	outputDir := os.Args[1]
	defineAst(outputDir, "Expr", []string{
		"Ternary  : Condition Expr, TrueExpr Expr, FalseExpr Expr",
		"Binary   : Left Expr, Operator Token, Right Expr",
		"Grouping : Expression Expr",
		"Literal  : Value any",
		"Unary    : Operator Token, Right Expr",
	})
}

func defineAst(outputDir string, baseName string, types []string) {
	path := outputDir + "/" + strings.ToLower(baseName) + ".go"
	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintf(file, "package lox\n\n")
	fmt.Fprintf(file, "type %s interface {\n", baseName)
	fmt.Fprintf(file, "\tAccept(visitor Visitor) any\n")
	fmt.Fprintf(file, "}\n\n")

	defineVisitor(file, baseName, types)

	for _, t := range types {
		className := strings.Trim(strings.Split(t, ":")[0], " ")
		fields := strings.Trim(strings.Split(t, ":")[1], " ")
		defineType(file, baseName, className, fields)
	}
}

func defineType(file *os.File, baseName string, className string, fieldList string) {
	fmt.Fprintf(file, "type %s struct {\n", className)
	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		fieldName := strings.Split(field, " ")[0]
		fieldType := strings.Split(field, " ")[1]
		fmt.Fprintf(file, "\t%s %s\n", fieldName, fieldType)
	}
	fmt.Fprintf(file, "}\n")
	fmt.Fprintf(file, "\n")
	fmt.Fprintf(file, "func (t %s) Accept(visitor Visitor) any {\n", className)
	fmt.Fprintf(file, "\treturn visitor.Visit%s%s(t)\n", baseName, className)
	fmt.Fprintf(file, "}\n\n")
}

func defineVisitor(file *os.File, baseName string, types []string) {
	fmt.Fprintf(file, "type Visitor interface {\n")
	for _, t := range types {
		className := strings.Trim(strings.Split(t, ":")[0], " ")
		fmt.Fprintf(file, "\tVisit%s%s(%s %s) any\n", baseName, className, strings.ToLower(className), className)
	}
	fmt.Fprintf(file, "}\n\n")
}
