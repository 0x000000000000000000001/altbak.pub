#!r6rs
#!chezscheme
(library
  (Data.Decide lib)
  (export
    choose
    chooseComparison
    chooseEquivalence
    chooseOp
    choosePredicate
    chosen
    identity)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Divide lib) Data.Divide.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define choosePredicate
    (scm:list (scm:cons (scm:string->symbol "choose") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (x3)
            (scm:let ([_4 (f0 x3)])
              (scm:cond
                [(Data.Either.Left? _4) (v1 (Data.Either.Left-value0 _4))]
                [(Data.Either.Right? _4) (v12 (Data.Either.Right-value0 _4))]
                [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
      Data.Divide.dividePredicate))))

  (scm:define chooseOp
    (scm:lambda (dictSemigroup0)
      (scm:let ([divideOp1 (Data.Divide.divideOp dictSemigroup0)])
        (scm:list (scm:cons (scm:string->symbol "choose") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:lambda (v14)
              (scm:lambda (x5)
                (scm:let ([_6 (f2 x5)])
                  (scm:cond
                    [(Data.Either.Left? _6) (v3 (Data.Either.Left-value0 _6))]
                    [(Data.Either.Right? _6) (v14 (Data.Either.Right-value0 _6))]
                    [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
          divideOp1))))))

  (scm:define chooseEquivalence
    (scm:list (scm:cons (scm:string->symbol "choose") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:let ([v25 (f0 a3)])
                (scm:cond
                  [(Data.Either.Left? v25) (scm:let ([v36 (f0 b4)])
                    (scm:cond
                      [(Data.Either.Left? v36) ((v1 (Data.Either.Left-value0 v25)) (Data.Either.Left-value0 v36))]
                      [(Data.Either.Right? v36) #f]
                      [scm:else (rt:fail)]))]
                  [(Data.Either.Right? v25) (scm:let ([v36 (f0 b4)])
                    (scm:cond
                      [(Data.Either.Left? v36) #f]
                      [(Data.Either.Right? v36) ((v12 (Data.Either.Right-value0 v25)) (Data.Either.Right-value0 v36))]
                      [scm:else (rt:fail)]))]
                  [scm:else (rt:fail)])))))))) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
      Data.Divide.divideEquivalence))))

  (scm:define chooseComparison
    (scm:list (scm:cons (scm:string->symbol "choose") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:let ([v25 (f0 a3)])
                (scm:cond
                  [(Data.Either.Left? v25) (scm:let ([v36 (f0 b4)])
                    (scm:cond
                      [(Data.Either.Left? v36) ((v1 (Data.Either.Left-value0 v25)) (Data.Either.Left-value0 v36))]
                      [(Data.Either.Right? v36) Data.Ordering.LT]
                      [scm:else (rt:fail)]))]
                  [(Data.Either.Right? v25) (scm:let ([v36 (f0 b4)])
                    (scm:cond
                      [(Data.Either.Left? v36) Data.Ordering.GT]
                      [(Data.Either.Right? v36) ((v12 (Data.Either.Right-value0 v25)) (Data.Either.Right-value0 v36))]
                      [scm:else (rt:fail)]))]
                  [scm:else (rt:fail)])))))))) (scm:cons (scm:string->symbol "Divide0") (scm:lambda (_)
      Data.Divide.divideComparison))))

  (scm:define choose
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "choose"))))

  (scm:define chosen
    (scm:lambda (dictDecide0)
      ((rt:record-ref dictDecide0 (scm:string->symbol "choose")) identity))))
