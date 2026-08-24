(library (Test.ArrayOpsFFI foreign)
  (export runArrayOpsFFI)
  (import (chezscheme))

  (define (runArrayOpsFFI limit)
    (let loop ([i 1] [sum 0])
      (if (> i limit)
          sum
          (if (= (modulo i 2) 0)
              (loop (+ i 1) (+ sum i))
              (loop (+ i 1) sum)))))
)
