# calkbud

Calkbud is a fast, lightweight command-line calculator built in Go for seamless terminal math operations. Designed with a modular Cobra architecture, it brings intuitive, developer-friendly calculations straight to your shell.

## Features

- **Basic Arithmetic Operators**: Add, subtract, multiply, and divide with instant results
- **Expression Evaluation**: Evaluate complex mathematical expressions with support for operator precedence and parentheses
- **Precision Control**: Adjust decimal precision on-the-fly using the `-p` or `--precision` flag
- **Robust Error Handling**: Non-zero exit codes for invalid input ensure reliable scripting and pipeline integration
- **Modular Architecture**: Built on Cobra CLI framework for extensibility and maintainability

## Installation

### Prerequisites

- Go 1.26.5 or higher

### Build from Source

```bash
git clone https://github.com/your-org/calkbud.git
cd calkbud
go build -o calkbud .
```

Then add the binary to your PATH or run it directly:

```bash
./calkbud [command] [args]
```

## Usage

### Basic Commands

**Add numbers:**
```bash
calkbud add 5 3 2 10
# Output: 20.00
```

**Subtract:**
```bash
calkbud sub 100 20 10
# Output: 70.00
```

**Multiply:**
```bash
calkbud mul 2 3 4
# Output: 24.00
```

**Divide:**
```bash
calkbud div 100 2 5
# Output: 10.00
```

### Expression Evaluation

Evaluate complex mathematical expressions with automatic operator precedence and bracket handling:

```bash
calkbud eval "2 + 3 * 4"
# Output: 14.00

calkbud eval "(2 + 3) * 4"
# Output: 20.00

calkbud eval "10 / 2 - 3"
# Output: 2.00
```

**With Precision Control:**
```bash
calkbud eval "10 / 3" -p 4
# Output: 3.3333

calkbud eval "22 / 7" --precision 5
# Output: 3.14286
```

**Implicit Multiplication:**
```bash
calkbud eval "3(2+4)"
# Output: 18.00

calkbud eval "(5)(6)"
# Output: 30.00
```

## Architecture

### Commands Structure

```
cmd/
├── root.go       # Main command dispatcher (Cobra setup)
├── add.go        # Addition command
├── subtract.go   # Subtraction command
├── multiply.go   # Multiplication command
├── divide.go     # Division command
└── eval.go       # Expression evaluation with govaluate
```

### Key Design Decisions

- **DisableFlagParsing in eval**: The eval subcommand disables Cobra's default flag parsing to preserve negative operators and minus signs in expressions (e.g., `-7-9` is correctly interpreted as "negative seven minus nine")
- **Manual Flag Parsing**: Custom flag parsing loop in eval.go handles precision flags flexibly (`-p 4`, `-p=4`, `--precision=4`)
- **Bracket Normalization**: Square brackets `[]` and curly braces `{}` are normalized to parentheses for consistent expression evaluation
- **Error Handling**: All commands exit with code 1 on invalid input for proper error signaling in shell scripts and pipelines


### Dependencies

- **[Cobra](https://github.com/spf13/cobra)** - CLI framework
- **[govaluate](https://github.com/Knetic/govaluate)** - Expression evaluation engine

## License

See [LICENSE](LICENSE) file for details.
