(library (Test.ListOpsFFICheatcode foreign)
  (export runListOpsFFICheatcode)
  (import (chezscheme))

  (define (runListOpsFFICheatcode limit)
    (let loop ([i 1] [sum 0])
      (if (> i limit)
          sum
          (if (= (modulo i 2) 0)
              (loop (+ i 1) (+ sum i))
              (loop (+ i 1) sum)))))
)
