(library (Test.RBTreeFFI foreign)
  (export runRBTreeFFI)
  (import (chezscheme))

  (define (make-node color left value right)
    (list color left value right))

  (define (color n) (car n))
  (define (left n) (cadr n))
  (define (val n) (caddr n))
  (define (right n) (cadddr n))

  (define (balance c a x b)
    (if (eq? c 'B)
        (cond
          [(and (not (null? a)) (eq? (color a) 'R))
           (cond
             [(and (not (null? (left a))) (eq? (color (left a)) 'R))
              (make-node 'R 
                         (make-node 'B (left (left a)) (val (left a)) (right (left a)))
                         (val a)
                         (make-node 'B (right a) x b))]
             [(and (not (null? (right a))) (eq? (color (right a)) 'R))
              (make-node 'R
                         (make-node 'B (left a) (val a) (left (right a)))
                         (val (right a))
                         (make-node 'B (right (right a)) x b))]
             [else (make-node c a x b)])]
          [(and (not (null? b)) (eq? (color b) 'R))
           (cond
             [(and (not (null? (left b))) (eq? (color (left b)) 'R))
              (make-node 'R
                         (make-node 'B a x (left (left b)))
                         (val (left b))
                         (make-node 'B (right (left b)) (val b) (right b)))]
             [(and (not (null? (right b))) (eq? (color (right b)) 'R))
              (make-node 'R
                         (make-node 'B a x (left b))
                         (val b)
                         (make-node 'B (left (right b)) (val (right b)) (right (right b))))]
             [else (make-node c a x b)])]
          [else (make-node c a x b)])
        (make-node c a x b)))

  (define (insert x t)
    (letrec ([ins (lambda (t)
                    (if (null? t)
                        (make-node 'R '() x '())
                        (cond
                          [(< x (val t)) (balance (color t) (ins (left t)) (val t) (right t))]
                          [(> x (val t)) (balance (color t) (left t) (val t) (ins (right t)))]
                          [else t])))])
      (let ([res (ins t)])
        (make-node 'B (left res) (val res) (right res)))))

  (define (depth t)
    (if (null? t)
        0
        (+ 1 (max (depth (left t)) (depth (right t))))))

  (define (runRBTreeFFI limit)
    (let loop ([i limit] [t '()])
      (if (<= i 0)
          (depth t)
          (loop (- i 1) (insert i t)))))
)
