(define fact (lambda (x)
    ":: (Callable (Int) Int)

    Calculates the factorial of x."
(block
    (if (< x 1) (block (return! 0)))

    (define total 1)
    (while (> x 1) (block
        (set! total (* total x))
        (set! x (+ x -1))
    ))

    (return! total)
)))

(display (fact 5))
