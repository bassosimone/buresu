usage: buresu run [flags] FILE

The `buresu run` command interprets and runs a Buresu program.

By default, `buresu run` interprets and executes the given file. However,
the `-E, --emit` flag can be used to stop the build process earlier
and dump intermediate representations.

The compilation pipeline is roughly as follows:

1. *scanner*: takes the source code as input and emits tokens,
which you can see by using `--emit tokens`.

2. *parser*: takes the tokens as input and emits an abstract syntax tree,
or AST, which you can see by using `--emit ast`.

3. *interpreter*: takes the AST as input and executes the program.

We support the following flags:

    -E, --emit
            Emit specific output (tokens, ast).

    -h, --help
            Show this help message and exit.

This command exits with `0` on success and `1` on failure.
