#!r6rs
#!chezscheme
(library
  (Data.Op lib)
  (export
    Op
    categoryOp
    contravariantOp
    monoidOp
    newtypeOp
    semigroupOp
    semigroupoidOp)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Op
    (scm:lambda (x0)
      x0))

  (scm:define semigroupoidOp
    (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (x2)
          (v11 (v0 x2))))))))

  (scm:define semigroupOp
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (x3)
            (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (f1 x3)) (g2 x3)))))))))

  (scm:define newtypeOp
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidOp
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([mempty11 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupFn3 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f3)
          (scm:lambda (g4)
            (scm:lambda (x5)
              (((rt:record-ref _2 (scm:string->symbol "append")) (f3 x5)) (g4 x5)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4)
            mempty11)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupFn3))))))

  (scm:define contravariantOp
    (scm:list (scm:cons (scm:string->symbol "cmap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (x2)
          (v1 (f0 x2))))))))

  (scm:define categoryOp
    (scm:list (scm:cons (scm:string->symbol "identity") (scm:lambda (x0)
      x0)) (scm:cons (scm:string->symbol "Semigroupoid0") (scm:lambda (_)
      semigroupoidOp)))))
