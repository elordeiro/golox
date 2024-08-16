# Lox Interpreter in Go

Welcome to my implementation of the Lox programming language in Go! Lox is a dynamically-typed, object-oriented programming language originally created by Bob Nystrom as a part of his book [_Crafting Interpreters_](https://craftinginterpreters.com).

This project is an ongoing effort to build a fully-functional Lox interpreter from the ground up, following the structure and guidance provided by the book, but written entirely in Go. The goal is to not only understand how interpreters work but also to leverage Go's unique features to create an efficient and clean implementation.

## Features

### Completed

-   **Scanning (Lexical Analysis)**

    -   Converts raw source code into a stream of tokens, which are the basic elements (keywords, identifiers, symbols, etc.) that the interpreter understands.

-   **Parsing (Syntactic Analysis)**
    -   The tokens are organized into a hierarchical structure, forming an abstract syntax tree (AST) that represents the logical structure of the code.
-   **Evaluating**
    -   The interpreter traverses the AST and executes the code according to the rules of the Lox language.

### In Progress

-   **Statements**
    -   Implementation of Lox's control flow constructs such as `if`, `else`, `while`, and `for` loops.
-   **Expressions**

    -   Handling arithmetic operations, comparisons, logical operations, and more complex expressions.

### Upcoming Features

-   **Control Flow**

    -   Proper management of the program's execution flow, including branching, looping, and more.

-   **Functions**
    -   Definition and invocation of Lox functions, including support for recursion and closures.
-   **Resolving and Binding**

    -   Static analysis to resolve variable bindings, ensuring correct scoping and access to variables.

-   **Classes**

    -   Object-oriented features, including class definitions, instantiation, and method calls.

-   **Inheritance**
    -   Implementation of class inheritance to allow for more complex object hierarchies.

## Getting Started

### Prerequisites

-   Go 1.20 or higher
-   Basic understanding of Go and interpreter design concepts

### Installation

Clone this repository and navigate to the project directory:

```bash
git clone https://github.com/elordeiro/GoLox.git
cd GoLox
```

### Running the Interpreter

To run the Lox interpreter on a Lox source file, use:

```bash
./golox.sh <mode> <file>
```

Where `<mode>` is either `tokenize`. `parse` or `evaluate` and `<file>` is the path to the Lox source file.

Alternatively, you can run the interpreter in REPL mode by simply executing:

```bash
./golox.sh <mode>
```

Where `<mode>` is either `tokenize`, `parse`, or `evaluate`.

### Contributing

Contributions are welcome! Feel free to submit issues, fork the repository, and open pull requests.

## Resources

-   [Crafting Interpreters](https://craftinginterpreters.com) - The book by Bob Nystrom that inspired this project.
-   [Go Programming Language](https://golang.org) - Official documentation for Go.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---
