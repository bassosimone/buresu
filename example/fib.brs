(include! "example/lib/fib.brs")

(define fib (fibgen))

(define idx 0)
(while (< idx 11) (block
    (display (fib))
    (set! idx (+ idx 1)
)))
