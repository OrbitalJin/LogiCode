# LogiCode

LogiCode is a scripting language that is designed to be used for boolean algebra and logic circuit design. It is a simple language that is easy to learn and use. It is designed to be used on Linux systems. It is written in Go which implements a hand-crafted lexer and parser.

## Grammar

Here is an example snippet of what a LogiCode program would look like:

`foo.lc`

```
!Program;

!Declare;
  LET x <- 001;
  LET y <- 010;
  LET z <- 011;
!EndDeclare;

!Begin;
  LET foo <- x AND y;
  LET bar <- NOT z;
  LET baz <- foo OR bar;
  READ baz;
!End;

!EndProgram;
```

This program will output `100` to stdout.

## Architecture

The [`Lexer`](https://en.wikipedia.org/wiki/Lexical_analysis) produces an array of [`Tokens`](https://bits.netbeans.org/11.1/javadoc/org-netbeans-modules-lexer/index.html?org/netbeans/api/lexer/Token.html) from the source code, which are then passed to the [`Parser`](https://en.wikipedia.org/wiki/Parsing). The `Parser` produces an Abstract Syntax Tree ([`AST`](https://en.wikipedia.org/wiki/Abstract_syntax_tree)) which is then passed to the Interpreter. The Interpreter then evaluates the `AST` using evalutation rules and produces an output. Here is a simple diagram that illustrates the architecture:

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
| `ASLETN`   | Assignment operator  | `EOF`          | End of file           |
| `IDENT`    | Identifier           | `DECLARESTART` | Declare start keyword |
| `INT`      | Integer              | `DECLAREEND`   | Declare end keyword   |
| `LETN`     | Signal               | `PROGRAMSTART` | Program start keyword |
| `LPAREN`   | Left parenthesis?    | `PROGRAMEND`   | Program end keyword   |
| `RPAREN`   | Right parenthesis?   | `BEGIN`        | Begin keyword         |
| `AND`      | Boolean and keyword  | `END`          | End keyword           |
| `OR`       | Boolean or keyword   | `SEMICOLON`    | Semicolon             |
| `XOR`      | Boolean xor keyword  | `WRITE`        | Write keyword         |
| `NOT`      | Boolean not keyword  | `READ`         | Read keyword          |
| `NAND`     | Boolean nand keyword | `LET`          | Let keyword           |
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

## Todo

-   [ ] Lex Comments (i.e `--`) (Filter out)
-   [x] Lex Write keyword (i.e. `WRITE`)
-   [x] Lex Read keyword (i.e. `READ`)
-   [x] Lex Let keyword (i.e. `LET`)
-   [x] Lex Identifier (i.e. `IDENT`)
-   [x] Lex Signal (i.e. `LETNAL`)
-   [x] Lex Assignment operator (i.e. `ASLETN`)
-   [x] Lex Boolean and keyword (i.e. `AND`)
-   [x] Lex Boolean or keyword (i.e. `OR`)
-   [x] Lex Boolean xor keyword (i.e. `XOR`)
-   [x] Lex Boolean not keyword (i.e. `NOT`)
-   [x] Lex Boolean nand keyword (i.e. `NAND`)
-   [x] Lex Boolean nor keyword (i.e. `NOR`)
-   [x] Lex Boolean xnor keyword (i.e. `XNOR`)
-   [x] Lex Entry Point delimiters (i.e. `BEGIN` and `END`)
-   [x] Lex Declare delimiters (i.e. `DECLARESTART` and `DECLAREEND`)
-   [x] Lex Program delimiters (i.e. `PROGRAMSTART` and `PROGRAMEND`)
