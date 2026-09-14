(library (Test.ArrayOpsFFI foreign)
  (export runArrayOpsFFI)
  (import (chezscheme))
  ;; Materialize the range, filtered array, then fold, as in Test.ArrayOps.
  (define (runArrayOpsFFI limit)
    (let* ([input (make-vector limit)]
           [evens (make-vector (quotient limit 2))])
      (do ([i 0 (+ i 1)]) ((= i limit))
        (vector-set! input i (+ i 1)))
      (let filter-loop ([i 0] [j 0])
        (unless (= i limit)
          (let ([x (vector-ref input i)])
            (if (= (modulo x 2) 0)
                (begin (vector-set! evens j x) (filter-loop (+ i 1) (+ j 1)))
                (filter-loop (+ i 1) j)))))
      (let fold ([i 0] [sum 0])
        (if (= i (vector-length evens)) sum
            (fold (+ i 1) (+ sum (vector-ref evens i))))))))
