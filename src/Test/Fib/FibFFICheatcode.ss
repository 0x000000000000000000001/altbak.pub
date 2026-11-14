(library (Test.FibFFICheatcode foreign)
  (export runFibFFICheatcode)
  (import (chezscheme))

  (define (fib n)
    (if (< n 2)
        n
        (+ (fib (- n 1)) (fib (- n 2)))))

  (define (runFibFFICheatcode limit)
    (fib limit))
)
