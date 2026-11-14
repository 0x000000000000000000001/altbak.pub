(library (Test.TCOFFI foreign)
  (export runTCOFFI)
  (import (chezscheme))

  (define (tco n acc)
    (if (<= n 0)
        acc
        (tco (- n 1) (+ acc (modulo n 3)))))

  (define (runTCOFFI limit)
    (tco limit 0))
)
