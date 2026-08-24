(library (Test.PolymorphismFFI foreign)
  (export runPolymorphismFFI)
  (import (chezscheme))

  (define (runPolymorphismFFI limit)
    (let loop ([i limit] [acc 0])
      (if (<= i 0)
          acc
          (loop (- i 1) (+ acc 1)))))
)
