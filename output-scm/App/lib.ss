#!r6rs
#!chezscheme
(library
  (App lib)
  (export
    Add
    Add*
    Add-value0
    Add-value1
    Add?
    Mul
    Mul*
    Mul-value0
    Mul-value1
    Mul?
    Sub
    Sub*
    Sub-value0
    Sub-value1
    Sub?
    Val
    Val-value0
    Val?
    buildTree
    eval
    fib
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define-record-type (Val$ Val Val?)
    (scm:fields (scm:immutable value0 Val-value0)))

  (scm:define-record-type (Add$ Add* Add?)
    (scm:fields (scm:immutable value0 Add-value0) (scm:immutable value1 Add-value1)))

  (scm:define Add
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Add* value0 value1))))

  (scm:define-record-type (Mul$ Mul* Mul?)
    (scm:fields (scm:immutable value0 Mul-value0) (scm:immutable value1 Mul-value1)))

  (scm:define Mul
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Mul* value0 value1))))

  (scm:define-record-type (Sub$ Sub* Sub?)
    (scm:fields (scm:immutable value0 Sub-value0) (scm:immutable value1 Sub-value1)))

  (scm:define Sub
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Sub* value0 value1))))

  (scm:define fib
    (scm:lambda (v0)
      (scm:cond
        [(scm:fx=? v0 0) 0]
        [(scm:fx=? v0 1) 1]
        [scm:else (scm:fx+ (fib (scm:fx- v0 1)) (fib (scm:fx- v0 2)))])))

  (scm:define eval
    (scm:lambda (v0)
      (scm:cond
        [(Val? v0) (Val-value0 v0)]
        [(Add? v0) (scm:fx+ (eval (Add-value0 v0)) (eval (Add-value1 v0)))]
        [(Mul? v0) (scm:fx* (eval (Mul-value0 v0)) (eval (Mul-value1 v0)))]
        [(Sub? v0) (scm:fx- (eval (Sub-value0 v0)) (eval (Sub-value1 v0)))]
        [scm:else (rt:fail)])))

  (scm:define buildTree
    (scm:lambda (v0)
      (scm:cond
        [(scm:fx=? v0 0) (Val 1)]
        [scm:else (Add* (Mul* (Val v0) (buildTree (scm:fx- v0 1))) (Sub* (buildTree (scm:fx- v0 1)) (Val 1)))])))

  (scm:define main
    (scm:let ([_0 (Effect.Console.log (rt:string->pstring "=== AST Evaluation ==="))])
      (scm:lambda ()
        (scm:let*
          ([_ (_0)]
           [_ ((Effect.Console.log (Data.Show.showIntImpl (eval (buildTree 12)))))]
           [_ ((Effect.Console.log (rt:string->pstring "=== Fibonacci ===")))])
            ((Effect.Console.log (Data.Show.showIntImpl (fib 40)))))))))
