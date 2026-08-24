(library (Test.RecordsFFICheatcode foreign)
  (export runRecordsFFICheatcode)
  (import (chezscheme))

  ;; Using mutable variables for cheatcode
  (define (runRecordsFFICheatcode limit)
    (let ([a 0] [b_c 0] [b_d_e 0] [b_d_f 0])
      (let loop ([n limit])
        (if (<= n 0)
            b_d_f
            (begin
              (set! a (+ a 1))
              (set! b_c (+ b_c 2))
              (set! b_d_e (+ b_d_e 3))
              (set! b_d_f (+ b_d_f (modulo n 5)))
              (loop (- n 1)))))))
)
