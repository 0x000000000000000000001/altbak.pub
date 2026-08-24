(library (Test.StateMonadFFI foreign)
  (export runStateMonadFFI)
  (import (chezscheme))

  (define (run-state s init)
    (s init))
  (define (bind-state s g)
    (lambda (state)
      (let* ([r1 (s state)]
             [v (car r1)]
             [s1 (cdr r1)]
             [g-prime (g v)])
        (g-prime s1))))
  (define (pure-state a)
    (lambda (s) (cons a s)))
  (define (get)
    (lambda (s) (cons s s)))
  (define (put s)
    (lambda (_) (cons 0 s)))
  (define (modify f)
    (bind-state (get) (lambda (s) (put (f s)))))
    
  (define (chain n)
    (if (= n 0)
        (pure-state 0)
        (bind-state (modify (lambda (x) (+ x 1)))
                    (lambda (_) (chain (- n 1))))))

  (define (runStateMonadFFI limit)
    (let loop ([n limit] [acc 0])
      (if (<= n 0)
          acc
          (loop (- n 1) (+ acc (cdr (run-state (chain 60) 0)))))))
)
