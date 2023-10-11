# LogiCode

LogiCode is a scripting language that is designed to be used for boolean algebra and logic circuit design. It is a simple language that is easy to learn and use. It is designed to be used on Linux systems. It is written in Go which implements a hand-crafted lexer and parser.

## Grammar

Here is an example snippet of what a LogiCode program would look like:

`foo.lc`

```
!Program

!Declare
  let SIG x <- 001;
  let SIG y <- 010;
  let SIG z <- 011;
!EndDeclare

!Begin
  -- This is a comment
  let SIG foo <- x AND y;
  let SIG bar <- NOT z;
  let SIG baz <- foo OR bar;
  -- This is should output `100` to stdout
  read baz;
!End

!EndProgram
```

## Architecture

The `Lexer` produces an array of `Tokens` from the source code, which are then passed to the `Parser`. The `Parser` produces an Abstract Syntax Tree (`AST`) which is then passed to the Interpreter. The Interpreter then evaluates the `AST` using evalutation rules and produces an output. Here is a simple diagram that illustrates the architecture:

```mermaid
graph LR;
  I(Src) --> Lexer;
  Lexer --> Parser;
  Parser --> Interpreter;
  Interpreter --> O(Output);
```

## Lexer

LogiCode embraces a simple syntax that is easy to learn and use.
Here are some of the soon-to-be Lexable tokens:

| Token Type | Description          | Token Type     | Description           |
| ---------- | -------------------- | -------------- | --------------------- |
| `ASSIGN`   | Assignment operator  | `LET`          | Let keyword           |
| `IDENT`    | Identifier           | `EOF`          | End of file           |
| `INT`      | Integer              | `DECLARESTART` | Declare start keyword |
| `SIGN`     | Signal               | `DECLAREEND`   | Declare end keyword   |
| `LPAREN`   | Left parenthesis?    | `PROGRAMSTART` | Program start keyword |
| `RPAREN`   | Right parenthesis?   | `PROGRAMEND`   | Program end keyword   |
| `AND`      | Boolean and keyword  | `BEGIN`        | Begin keyword         |
| `OR`       | Boolean or keyword   | `END`          | End keyword           |
| `XOR`      | Boolean xor keyword  | `SEMICOLON`    | Semicolon             |
| `NOT`      | Boolean not keyword  | `WRITE`        | Write keyword         |
| `NAND`     | Boolean nand keyword | `READ`         | Read keyword          |
| `NOR`      | Boolean nor keyword  |
| `XNOR`     | Boolean xnor keyword |

> **Note** These are reserved keywords that cannot be used as identifiers.

## Parser

The `Parser` is a recursive descent parser that produces an `AST` from the `Tokens` that are produced by the `Lexer`.

## AST

The `AST` is a tree data structure that represents the source code. It is used to evaluate the source code in a recursive manner which
is defined by the grammar of the language. This approach makes it trivial to respect operator precedence and associativity. Here is a simple diagram that illustrates the `AST`:

```mermaid
graph TD;
    P[Program] --> S1[Stmt 1];
    P[Program] --> S2[Stmt 2];
    S1 --> E1[Expr 1];
    S1 --> E2[Expr 2];
    E2 --> T1[Term 1];
    E2 --> F1[Factor];
    F1[Factor] --> T2[Term 2];
    F1[Factor] --> T3[Term 3];

```

Here is a simple example of the `AST` representation of an expression:

Expression: `a & b | c`

```mermaid
graph TD;
    F1[&] --> T1[a];
    F1 --> F2["|"];
    F2 --> T2[b];
    F2 --> T3[c];
```

> **Warning** The `AST` is not a binary tree. It is a tree data structure that can have any number of children.
