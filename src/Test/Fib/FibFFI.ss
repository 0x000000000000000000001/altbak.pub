(library (Test.FibFFI foreign)
  (export runFibFFI)
  (import (chezscheme))

  (define (fib n)
    (if (< n 2)
        n
        (+ (fib (- n 1)) (fib (- n 2)))))

  (define (runFibFFI limit)
    (fib limit))
)
