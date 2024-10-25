# Buresu

Buresu (ブレース, in Japanese, meaning "brace") is an intepreter for
the namesake Lisp-like, hobby programming language. This project is
primarily for educational purposes and personal enjoyment.


## Features

- **Abstract Syntax Tree (AST)**: The core of the language is
represented using an AST.
- **Scanner**: Tokenizes the input source code (lexical analysis).
- **Parser**: Converts a sequence of tokens into an AST.
- **Includer**: Includes external scripts in the main script.
- **Type checker**: Checks the types of the AST nodes.
- **Evaluator**: Evaluates the AST nodes to execute the program.
- **Built-in Functions**: Includes basic built-in functions like addition,
multiplication, and display.


## Getting Started

### Prerequisites

- Go 1.23 or later

### Installation

1. Clone the repository:

```sh
git clone https://github.com/bassosimone/buresu
cd buresu
```

2. Build the project:

```sh
go build -v ./cmd/buresu
```

### Running Examples

You can run the provided examples to see the language in action:

```sh
./buresu run example/fact.brs
./buresu run example/fib.brs
```

### Interactive Shell

Use

```sh
./buresu repl
```

to run an interactive shell.

## Project Structure

- `cmd`: Contains the source code for the command-line interface.
- `internal`: Contains the internal packages.
- `pkg/ast`: Contains the AST definitions.
- `pkg/dumper`: Contains the AST dumper.
- `pkg/includer`: Contains the includer that includes external scripts in the main script.
- `pkg/legacy`: Contains the legacy evaluator that executes the AST nodes.
- `pkg/parser`: Contains the parser that converts tokens into AST nodes.
- `pkg/evaluator`: Contains the evaluator that executes the AST nodes.
- `pkg/scanner`: Contains the scanner that tokenizes the input source code.
- `pkg/tolen`: Interface between the scanner and the parser.
- `example`: Contains example scripts written in the Buresu language.
- `example/lib`: Contains library scripts that can be included in example scripts.

## Contributing

This is a personal hobby project, but contributions and suggestions are
welcome. Feel free to open issues or submit pull requests.

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```
