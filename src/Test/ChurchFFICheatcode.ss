(library (Test.ChurchFFICheatcode foreign)
  (export runChurchFFICheatcode)
  (import (chezscheme))

  (define (runChurchFFICheatcode limit)
    (let loop ([i 0] [acc 0])
      (if (>= i (* limit 10000))
          acc
          (loop (+ i 1) (+ acc 1)))))
)
