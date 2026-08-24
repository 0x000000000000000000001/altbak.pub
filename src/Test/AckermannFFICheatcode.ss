(library (Test.AckermannFFICheatcode foreign)
  (export runAckermannFFICheatcode)
  (import (chezscheme))

  (define (ack m n)
    (cond
      [(= m 0) (+ n 1)]
      [(and (> m 0) (= n 0)) (ack (- m 1) 1)]
      [else (ack (- m 1) (ack m (- n 1)))]))

  (define (runAckermannFFICheatcode limit)
    (ack 3 4))
)
