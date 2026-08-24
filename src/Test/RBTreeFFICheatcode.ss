(library (Test.RBTreeFFICheatcode foreign)
  (export runRBTreeFFICheatcode)
  (import (chezscheme))

  ;; Arena allocation with vectors for tree nodes
  (define pool (make-vector 10000000))
  (define pool-idx 0)
  
  (define (make-node color left value right)
    (let ([idx pool-idx])
      (set! pool-idx (+ pool-idx 1))
      (vector-set! pool idx (vector color left value right))
      idx))

  (define (color n) (vector-ref (vector-ref pool n) 0))
  (define (left n) (vector-ref (vector-ref pool n) 1))
  (define (val n) (vector-ref (vector-ref pool n) 2))
  (define (right n) (vector-ref (vector-ref pool n) 3))

  (define (balance c a x b)
    (if (= c 1) ; 1 = B, 0 = R
        (cond
          [(and (not (= a -1)) (= (color a) 0))
           (cond
             [(and (not (= (left a) -1)) (= (color (left a)) 0))
              (make-node 0 
                         (make-node 1 (left (left a)) (val (left a)) (right (left a)))
                         (val a)
                         (make-node 1 (right a) x b))]
             [(and (not (= (right a) -1)) (= (color (right a)) 0))
              (make-node 0
                         (make-node 1 (left a) (val a) (left (right a)))
                         (val (right a))
                         (make-node 1 (right (right a)) x b))]
             [else (make-node c a x b)])]
          [(and (not (= b -1)) (= (color b) 0))
           (cond
             [(and (not (= (left b) -1)) (= (color (left b)) 0))
              (make-node 0
                         (make-node 1 a x (left (left b)))
                         (val (left b))
                         (make-node 1 (right (left b)) (val b) (right b)))]
             [(and (not (= (right b) -1)) (= (color (right b)) 0))
              (make-node 0
                         (make-node 1 a x (left b))
                         (val b)
                         (make-node 1 (left (right b)) (val (right b)) (right (right b))))]
             [else (make-node c a x b)])]
          [else (make-node c a x b)])
        (make-node c a x b)))

  (define (insert x t)
    (letrec ([ins (lambda (t)
                    (if (= t -1)
                        (make-node 0 -1 x -1)
                        (cond
                          [(< x (val t)) (balance (color t) (ins (left t)) (val t) (right t))]
                          [(> x (val t)) (balance (color t) (left t) (val t) (ins (right t)))]
                          [else t])))])
      (let ([res (ins t)])
        (make-node 1 (left res) (val res) (right res)))))

  (define (depth t)
    (if (= t -1)
        0
        (+ 1 (max (depth (left t)) (depth (right t))))))

  (define (runRBTreeFFICheatcode limit)
    (set! pool-idx 0)
    (let loop ([i limit] [t -1])
      (if (<= i 0)
          (depth t)
          (loop (- i 1) (insert i t)))))
)
