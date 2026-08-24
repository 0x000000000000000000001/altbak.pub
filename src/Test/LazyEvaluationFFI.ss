(library (Test.LazyEvaluationFFI foreign)
  (export runLazyEvaluationFFI)
  (import (chezscheme))

  (define (defer f) (lambda () (f)))
  (define (force-thunk t) (t))
  
  (define (build-thunks n acc)
    (if (= n 0)
        acc
        (build-thunks (- n 1) (defer (lambda () (+ 1 (force-thunk acc)))))))

  (define (runLazyEvaluationFFI limit)
    (let loop ([n limit] [acc 0])
      (if (<= n 0)
          acc
          (let* ([t (build-thunks 1000 (defer (lambda () 0)))]
                 [val (force-thunk t)])
            (loop (- n 1) (+ acc val))))))
)
