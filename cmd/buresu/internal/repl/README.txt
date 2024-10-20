usage: buresu repl

The `buresu repl` command starts an interactive Read-Eval-Print Loop, also
known as a REPL, that reads Buresu Lisp code and evaluates it on the fly.

When started, this command displays the following prompt:

    >>>

To exit the REPL, use the `Ctrl-D` key combination.

If the statement you type is incomplete, for example, if you have only
typed half of a `lambda` expression, the prompt changes to:

    ...

to indicate that the REPL is waiting for more input.

Use `Ctrl-C` to cancel the current statement and start a new one.

Also, you can use `Ctrl-C` to interrupt the evaluation of an
expression that is taking too long to complete.

Apart from this, the REPL behaves as if you typed the code in a file and
executed it with the `buresu run` command.

Currently, no command line flags are supported.

This command exits with `0` on success and `1` on failure.
