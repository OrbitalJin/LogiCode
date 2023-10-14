# LogiCode

LogiCode is a scripting language that is designed to be used for [boolean algebra](https://en.wikipedia.org/wiki/Boolean_algebra) and [logic circuit](https://en.wikipedia.org/?title=Logic_circuit&redirect=no) design. It is a simple language that is easy to learn and use. The compiler is written in [Go](https://en.wikipedia.org/wiki/Go_(programming_language)) which implements a hand-crafted [lexer](https://en.wikipedia.org/wiki/Lexical_analysis) and [parser](https://en.wikipedia.org/wiki/Parsing).

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

The `Lexer` returns an array of [`Tokens`](https://bits.netbeans.org/11.1/javadoc/org-netbeans-modules-lexer/index.html?org/netbeans/api/lexer/Token.html) from the [HLL](https://en.wikipedia.org/wiki/High-level_programming_language) source code, which are then passed to the `Parser`. The `Parser` generates an Abstract Syntax Tree ([`AST`](https://en.wikipedia.org/wiki/Abstract_syntax_tree)) which is then passed to the Interpreter. The Interpreter then evaluates the `AST` using evalutation rules and produces an output. Here is a simple diagram that illustrates the architecture:

```mermaid
graph LR;
  HLL --> Lexer;
  Lexer --> Parser;
  Parser --> AST[AST Evaluation];
  AST --> O(Output);
```

## Roadmap

-   [x] Lexer Implementation (_Hasn't been thoroughly tested, yet_)
-   [ ] Error Handling Infrastructure across all modules
-   [ ] Parser Implementation
-   [ ] REPL Implementation
-   [ ] Packaging

## Lexer

LogiCode embraces a simple syntax that is easy to learn and use.
Here are some of the supported Lexable tokens:

| Token Type | Description           | Token Type     | Description           |
| ---------- | --------------------- | -------------- | --------------------- |
| `ASSIGN`   | Assignment operator   | `EOF`          | End of file           |
| `IDENT`    | Identifier            | `DECLARESTART` | Declare start keyword |
| ~~`INT`~~  | ~~Integer~~           | `DECLAREEND`   | Declare end keyword   |
| `SIGN`     | Signal                | `PROGRAMSTART` | Program start keyword |
| `AND`      | BitWise and operator  | `PROGRAMEND`   | Program end keyword   |
| `OR`       | BitWise or operator   | `BEGIN`        | Begin keyword         |
| `XOR`      | BitWise xor operator  | `END`          | End keyword           |
| `NOT`      | BitWise not operator  | `SEMICOLON`    | Line Seperator        |
| `NAND`     | BitWise nand operator | `WRITE`        | Write keyword         |
| `NOR`      | BitWise nor operator  | `READ`         | Read keyword          |
| `XNOR`     | BitWise xnor operator | `LET`          | Let keyword           |
> **Note** These are reserved keywords that cannot be used as identifiers.
> Eventually I would like to migrate from the the key-per-key string streaming approach to a buffer scanning approach.

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
