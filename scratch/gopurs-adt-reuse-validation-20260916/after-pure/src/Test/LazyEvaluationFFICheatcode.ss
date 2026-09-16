(library (Test.LazyEvaluationFFICheatcode foreign)
  (export runLazyEvaluationFFICheatcode)
  (import (chezscheme))

  (define (runLazyEvaluationFFICheatcode limit)
    (let loop ([i limit] [acc 0])
      (if (<= i 0)
          acc
          (loop (- i 1) (+ acc 1000)))))
)
