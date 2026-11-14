(library (Test.PolymorphismFFICheatcode foreign)
  (export runPolymorphismFFICheatcode)
  (import (chezscheme))

  (define (runPolymorphismFFICheatcode limit)
    (let loop ([i limit] [acc 0])
      (if (<= i 0)
          acc
          (loop (- i 1) (+ acc 1)))))
)
