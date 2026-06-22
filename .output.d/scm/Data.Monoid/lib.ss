#!r6rs
#!chezscheme
(library
  (Data.Monoid lib)
  (export
    guard
    mempty
    memptyRecord
    monoidArray
    monoidFn
    monoidOrdering
    monoidRecord
    monoidRecordCons
    monoidRecordNil
    monoidString
    monoidUnit
    power)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.EuclideanRing lib) Data.EuclideanRing.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Semigroup lib) Data.Semigroup.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.))

  (scm:define monoidUnit
    (scm:list (scm:cons (scm:string->symbol "mempty") Data.Unit.unit) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      Data.Semigroup.semigroupUnit))))

  (scm:define monoidString
    (scm:list (scm:cons (scm:string->symbol "mempty") (rt:string->pstring "")) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      Data.Semigroup.semigroupString))))

  (scm:define monoidRecordNil
    (scm:list (scm:cons (scm:string->symbol "memptyRecord") (scm:lambda (v0)
      (scm:list))) (scm:cons (scm:string->symbol "SemigroupRecord0") (scm:lambda (_)
      Data.Semigroup.semigroupRecordNil))))

  (scm:define monoidOrdering
    (scm:list (scm:cons (scm:string->symbol "mempty") Data.Ordering.EQ) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      Data.Ordering.semigroupOrdering))))

  (scm:define monoidArray
    (scm:list (scm:cons (scm:string->symbol "mempty") (rt:make-array)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      Data.Semigroup.semigroupArray))))

  (scm:define memptyRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "memptyRecord"))))

  (scm:define monoidRecord
    (scm:lambda (_)
      (scm:lambda (dictMonoidRecord1)
        (scm:let ([semigroupRecord12 (scm:list (scm:cons (scm:string->symbol "append") ((rt:record-ref ((rt:record-ref dictMonoidRecord1 (scm:string->symbol "SemigroupRecord0")) (scm:quote undefined)) (scm:string->symbol "appendRecord")) Type.Proxy.Proxy)))])
          (scm:list (scm:cons (scm:string->symbol "mempty") ((rt:record-ref dictMonoidRecord1 (scm:string->symbol "memptyRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupRecord12)))))))

  (scm:define mempty
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "mempty"))))

  (scm:define monoidFn
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

  (scm:define monoidRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (dictMonoid1)
        (scm:let*
          ([mempty12 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))]
           [Semigroup03 ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined))])
            (scm:lambda (_)
              (scm:lambda (dictMonoidRecord5)
                (scm:let*
                  ([_6 ((rt:record-ref dictMonoidRecord5 (scm:string->symbol "SemigroupRecord0")) (scm:quote undefined))]
                   [semigroupRecordCons17 (scm:list (scm:cons (scm:string->symbol "appendRecord") (scm:lambda (v7)
                    (scm:lambda (ra8)
                      (scm:lambda (rb9)
                        (scm:let*
                          ([key10 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                           [get11 (Record.Unsafe.unsafeGet key10)])
                            (((Record.Unsafe.unsafeSet key10) (((rt:record-ref Semigroup03 (scm:string->symbol "append")) (get11 ra8)) (get11 rb9))) ((((rt:record-ref _6 (scm:string->symbol "appendRecord")) Type.Proxy.Proxy) ra8) rb9))))))))])
                    (scm:list (scm:cons (scm:string->symbol "memptyRecord") (scm:lambda (v8)
                      (((Record.Unsafe.unsafeSet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)) mempty12) ((rt:record-ref dictMonoidRecord5 (scm:string->symbol "memptyRecord")) Type.Proxy.Proxy)))) (scm:cons (scm:string->symbol "SemigroupRecord0") (scm:lambda (_)
                      semigroupRecordCons17))))))))))

  (scm:define power
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([mempty11 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))])
          (scm:lambda (x3)
            (scm:letrec ([go4 (scm:lambda (p5)
              (scm:cond
                [(scm:fx<=? p5 0) mempty11]
                [(scm:fx=? p5 1) x3]
                [(scm:fx=? ((Data.EuclideanRing.intMod p5) 2) 0) (scm:let ([x$p6 (go4 (scm:fx/ p5 2))])
                  (((rt:record-ref _2 (scm:string->symbol "append")) x$p6) x$p6))]
                [scm:else (scm:let ([x$p6 (go4 (scm:fx/ p5 2))])
                  (((rt:record-ref _2 (scm:string->symbol "append")) x$p6) (((rt:record-ref _2 (scm:string->symbol "append")) x$p6) x3)))]))])
              go4)))))

  (scm:define guard
    (scm:lambda (dictMonoid0)
      (scm:let ([mempty11 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [v2 v13]
              [scm:else mempty11])))))))
