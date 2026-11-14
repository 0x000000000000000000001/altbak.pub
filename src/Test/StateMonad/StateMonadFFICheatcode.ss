(library (Test.StateMonadFFICheatcode foreign)
  (export runStateMonadFFICheatcode)
  (import (chezscheme))

  (define (runStateMonadFFICheatcode limit)
    (let loop ([n (* limit 20)] [state 0])
      (if (<= n 0)
          state
          (loop (- n 1) (+ state 1)))))
)
