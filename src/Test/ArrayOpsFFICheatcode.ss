(library (Test.ArrayOpsFFICheatcode foreign)
  (export runArrayOpsFFICheatcode)
  (import (chezscheme))

  (define (runArrayOpsFFICheatcode limit)
    (let loop ([i 1] [sum 0] [step (if (>= limit 1) 1 -1)])
      (if (= i (+ limit step))
          sum
          (if (= (modulo i 2) 0)
              (loop (+ i step) (+ sum i) step)
              (loop (+ i step) sum step)))))
)
