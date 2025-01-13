# Tiny Compiler in Go

A lightweight, educational compiler written in Go, designed to help understand the basics of parsing and compiling a simplified programming language.

## Features

- Tokenizes input strings into meaningful components (lexer).
- Parses tokens to create an Abstract Syntax Tree (AST).
- Traverses and transforms the AST.
- Generates code for a target language from the transformed AST.

## How It Works

1. **Lexing (Tokenization):**
   Converts the raw input string into a series of tokens.

2. **Parsing:**
   Organizes the tokens into an Abstract Syntax Tree (AST), which represents the structure of the code.

3. **Traversing:**
   Visits each node of the AST, performing actions like transforming or analyzing the structure.

4. **Code Generation:**
   Transforms the AST into output code in another language.

### Example Input and Output

#### Input (Simple Expression in Your Language):
```javascript
(add 2 (subtract 4 2))
```

#### Output (Transpiled/Compiled to Target Language):
```javascript
add(2, subtract(4, 2));
