(library (Test.AstTreeFFI foreign)
  (export runAstTreeFFI)
  (import (chezscheme))

  (define (eval-ast e)
    (cond
      [(eq? (car e) 'Val) (cadr e)]
      [(eq? (car e) 'Add) (+ (eval-ast (cadr e)) (eval-ast (caddr e)))]
      [(eq? (car e) 'Mul) (* (eval-ast (cadr e)) (eval-ast (caddr e)))]
      [(eq? (car e) 'Sub) (- (eval-ast (cadr e)) (eval-ast (caddr e)))]
      [else 0]))

  (define (buildTree depth)
    (if (= depth 0)
        '(Val 1)
        `(Add (Mul (Val ,depth) ,(buildTree (- depth 1)))
              (Sub ,(buildTree (- depth 1)) (Val 1)))))

  (define (runAstTreeFFI limit)
    (let loop ([i limit] [res 0])
      (if (= i 0)
          res
          (loop (- i 1) (eval-ast (buildTree 6))))))
)
