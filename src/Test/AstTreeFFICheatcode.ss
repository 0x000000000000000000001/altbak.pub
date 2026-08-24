(library (Test.AstTreeFFICheatcode foreign)
  (export runAstTreeFFICheatcode)
  (import (chezscheme))

  ;; Scheme is naturally good at evaluating lists as ASTs. We can use vectors for a slight speedup,
  ;; or just plain lists since chezscheme optimizes them well. Let's use vectors for CC.
  (define (eval-ast e)
    (cond
      [(eq? (vector-ref e 0) 0) (vector-ref e 1)]
      [(eq? (vector-ref e 0) 1) (+ (eval-ast (vector-ref e 1)) (eval-ast (vector-ref e 2)))]
      [(eq? (vector-ref e 0) 2) (* (eval-ast (vector-ref e 1)) (eval-ast (vector-ref e 2)))]
      [(eq? (vector-ref e 0) 3) (- (eval-ast (vector-ref e 1)) (eval-ast (vector-ref e 2)))]
      [else 0]))

  (define (buildTree depth)
    (if (= depth 0)
        (vector 0 1)
        (vector 1 (vector 2 (vector 0 depth) (buildTree (- depth 1)))
                  (vector 3 (buildTree (- depth 1)) (vector 0 1)))))

  (define (runAstTreeFFICheatcode limit)
    (let loop ([i limit] [res 0])
      (if (= i 0)
          res
          (loop (- i 1) (eval-ast (buildTree 6))))))
)
