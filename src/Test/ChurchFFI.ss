(library (Test.ChurchFFI foreign)
  (export runChurchFFI)
  (import (chezscheme))

  (define (runChurchFFI limit)
    (let loop ([i 0] [acc 0])
      (if (>= i (* limit 10000))
          acc
          (loop (+ i 1) (+ acc 1)))))
)
