(library (Test.TCOFFICheatcode foreign)
  (export runTCOFFICheatcode)
  (import (chezscheme))

  (define (tco n acc)
    (if (<= n 0)
        acc
        (tco (- n 1) (+ acc (modulo n 3)))))

  (define (runTCOFFICheatcode limit)
    (tco limit 0))
)
