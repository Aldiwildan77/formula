# Formula

A formula engine implemented in Go that parses and evaluates mathematical and logical expressions.

## Overview

Formula is a Go-based expression evaluation engine that parses formulas into an Abstract Syntax Tree (AST) and evaluates them. It supports various mathematical operations, logical expressions, array functions, variable access, JSON data integration, and more.

## Features

- Mathematical operations
- Logical expressions
- Comparison operators
- Array functions
- Variable support
- JSON data integration

## Examples

The project includes several examples demonstrating different use cases:

- One-line expressions
- Multi-line formulas
- JSON array processing
- JSON map processing
- Compound interest calculations

## Sample Formula

This is a sample formula for Compound Interest Calculation:

```formula
MULTIPLY(
  VAR('P'), 
  POW(
    ADD(
      1, 
      DIVIDE(
        VAR('r'), 
        VAR('n')
      )
    ), 
    MULTIPLY(
      VAR('n'), 
      VAR('t')
    )
  )
)
```

## Getting Started

### Prerequisites

- Go 1.21.1 or higher

### Building

```bash
make build
```

### Running Tests

```bash
make test
```

For function-specific tests with coverage:

```bash
make test-functions
make test-functions-coverage
```

## Project Structure

- `core/`: Core types and AST definitions
- `engine/`: Formula parsing and evaluation engine
- `functions/`: Implementation of supported functions and operators
- `examples/`: Usage examples

## Available Make Commands

Run `make help` to see all available commands:

```bash
make help
```

## Author

- [Muhammad Wildan Aldiansyah](https://github.com/Aldiwildan77)

## Contributing

Contributions are welcome! Please feel free to submit an issue or pull request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
