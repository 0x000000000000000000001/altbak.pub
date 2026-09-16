(library (Test.RBTreeFFICheatcode foreign)
  (export runRBTreeFFICheatcode)
  (import (chezscheme))

  (define-record-type (node make-node node?)
    (fields color left value right))

  (define (balance c a x b)
    (if (eq? c 'B)
        (cond
          [(and (not (null? a)) (eq? (node-color a) 'R))
           (cond
             [(and (not (null? (node-left a))) (eq? (node-color (node-left a)) 'R))
              (make-node 'R 
                         (make-node 'B (node-left (node-left a)) (node-value (node-left a)) (node-right (node-left a)))
                         (node-value a)
                         (make-node 'B (node-right a) x b))]
             [(and (not (null? (node-right a))) (eq? (node-color (node-right a)) 'R))
              (make-node 'R
                         (make-node 'B (node-left a) (node-value a) (node-left (node-right a)))
                         (node-value (node-right a))
                         (make-node 'B (node-right (node-right a)) x b))]
             [else (make-node c a x b)])]
          [(and (not (null? b)) (eq? (node-color b) 'R))
           (cond
             [(and (not (null? (node-left b))) (eq? (node-color (node-left b)) 'R))
              (make-node 'R
                         (make-node 'B a x (node-left (node-left b)))
                         (node-value (node-left b))
                         (make-node 'B (node-right (node-left b)) (node-value b) (node-right b)))]
             [(and (not (null? (node-right b))) (eq? (node-color (node-right b)) 'R))
              (make-node 'R
                         (make-node 'B a x (node-left b))
                         (node-value b)
                         (make-node 'B (node-left (node-right b)) (node-value (node-right b)) (node-right (node-right b))))]
             [else (make-node c a x b)])]
          [else (make-node c a x b)])
        (make-node c a x b)))

  (define (insert x t)
    (letrec ([ins (lambda (t)
                    (if (null? t)
                        (make-node 'R '() x '())
                        (cond
                          [(< x (node-value t)) (balance (node-color t) (ins (node-left t)) (node-value t) (node-right t))]
                          [(> x (node-value t)) (balance (node-color t) (node-left t) (node-value t) (ins (node-right t)))]
                          [else t])))])
      (let ([res (ins t)])
        (make-node 'B (node-left res) (node-value res) (node-right res)))))

  (define (depth t)
    (if (null? t)
        0
        (+ 1 (max (depth (node-left t)) (depth (node-right t))))))

  (define (runRBTreeFFICheatcode limit)
    (let loop ([i limit] [t '()])
      (if (<= i 0)
          (depth t)
          (loop (- i 1) (insert i t)))))
)
