#!r6rs
#!chezscheme
(library
  (Data.Bifunctor lib)
  (export
    bifunctorConst
    bifunctorEither
    bifunctorTuple
    bimap
    identity
    lmap
    rmap)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define bimap
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bimap"))))

  (scm:define lmap
    (scm:lambda (dictBifunctor0)
      (scm:lambda (f1)
        (((rt:record-ref dictBifunctor0 (scm:string->symbol "bimap")) f1) identity))))

  (scm:define rmap
    (scm:lambda (dictBifunctor0)
      ((rt:record-ref dictBifunctor0 (scm:string->symbol "bimap")) identity)))

  (scm:define bifunctorTuple
    (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (v2)
          (Data.Tuple.Tuple* (f0 (Data.Tuple.Tuple-value0 v2)) (g1 (Data.Tuple.Tuple-value1 v2)))))))))

  (scm:define bifunctorEither
    (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Either.Left? v22) (Data.Either.Left (v0 (Data.Either.Left-value0 v22)))]
            [(Data.Either.Right? v22) (Data.Either.Right (v11 (Data.Either.Right-value0 v22)))]
            [scm:else (rt:fail)])))))))

  (scm:define bifunctorConst
    (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (f0 v12))))))))
