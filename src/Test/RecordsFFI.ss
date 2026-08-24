(library (Test.RecordsFFI foreign)
  (export runRecordsFFI)
  (import (chezscheme))

  ;; Whitebox FFI for records means using purely functional nested updates
  (define (runRecordsFFI limit)
    (let loop ([n limit] [a 0] [b_c 0] [b_d_e 0] [b_d_f 0])
      (if (<= n 0)
          b_d_f
          (loop (- n 1) 
                (+ a 1) 
                (+ b_c 2) 
                (+ b_d_e 3) 
                (+ b_d_f (modulo n 5))))))
)
