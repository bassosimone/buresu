(define fibgen (lambda ()
    ":: (Callable () (Callable () Int))

    Returns an iterator that computes the next Fibonacci number each time it is called."
(block
    (define prev 0)
    (define cur 1)
    (lambda ()
        ":: (Callable () Int)"
    (block
        (define oprev prev)
        (set! prev cur)
        (set! cur (+ cur oprev))
        (return! cur)
    ))
)))

(define fib (fibgen))

(define idx 0)
(while (< idx 11) (block
    (display (fib))
    (set! idx (+ idx 1)
)))
