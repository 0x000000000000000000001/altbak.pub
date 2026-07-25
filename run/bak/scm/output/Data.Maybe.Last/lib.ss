#!r6rs
#!chezscheme
(library
  (Data.Maybe.Last lib)
  (export
    Last
    altLast
    alternativeLast
    applicativeLast
    applyLast
    bindLast
    boundedLast
    eq1Last
    eqLast
    extendLast
    functorLast
    invariantLast
    monadLast
    monoidLast
    newtypeLast
    ord1Last
    ordLast
    plusLast
    semigroupLast
    showLast)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define Last
    (scm:lambda (x0)
      x0))

  (scm:define showLast
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (scm:cond
          [(Data.Maybe.Just? v1) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Last (Just ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Data.Maybe.Just-value0 v1))) (rt:string->pstring "))"))]
          [(Data.Maybe.Nothing? v1) (rt:string->pstring "(Last Nothing)")]
          [scm:else (rt:fail)]))))))

  (scm:define semigroupLast
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Data.Maybe.Just? v11) v11]
          [(Data.Maybe.Nothing? v11) v0]
          [scm:else (rt:fail)]))))))

  (scm:define ordLast
    (scm:lambda (dictOrd0)
      (scm:let*
        ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))]
         [eqMaybe12 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:cond
              [(Data.Maybe.Nothing? x2) (Data.Maybe.Nothing? y3)]
              [scm:else (scm:and (Data.Maybe.Just? x2) (scm:and (Data.Maybe.Just? y3) (((rt:record-ref _1 (scm:string->symbol "eq")) (Data.Maybe.Just-value0 x2)) (Data.Maybe.Just-value0 y3))))])))))])
          (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x3)
            (scm:lambda (y4)
              (scm:cond
                [(Data.Maybe.Nothing? x3) (scm:cond
                  [(Data.Maybe.Nothing? y4) Data.Ordering.EQ]
                  [scm:else Data.Ordering.LT])]
                [(Data.Maybe.Nothing? y4) Data.Ordering.GT]
                [(scm:and (Data.Maybe.Just? x3) (Data.Maybe.Just? y4)) (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Data.Maybe.Just-value0 x3)) (Data.Maybe.Just-value0 y4))]
                [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
            eqMaybe12))))))

  (scm:define ord1Last
    Data.Maybe.ord1Maybe)

  (scm:define newtypeLast
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidLast
    (scm:list (scm:cons (scm:string->symbol "mempty") Data.Maybe.Nothing) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      semigroupLast))))

  (scm:define monadLast
    Data.Maybe.monadMaybe)

  (scm:define invariantLast
    Data.Maybe.invariantMaybe)

  (scm:define functorLast
    Data.Maybe.functorMaybe)

  (scm:define extendLast
    Data.Maybe.extendMaybe)

  (scm:define eqLast
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:cond
            [(Data.Maybe.Nothing? x1) (Data.Maybe.Nothing? y2)]
            [scm:else (scm:and (Data.Maybe.Just? x1) (scm:and (Data.Maybe.Just? y2) (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Data.Maybe.Just-value0 x1)) (Data.Maybe.Just-value0 y2))))])))))))

  (scm:define eq1Last
    Data.Maybe.eq1Maybe)

  (scm:define boundedLast
    (scm:lambda (dictBounded0)
      (Data.Maybe.boundedMaybe dictBounded0)))

  (scm:define bindLast
    Data.Maybe.bindMaybe)

  (scm:define applyLast
    Data.Maybe.applyMaybe)

  (scm:define applicativeLast
    Data.Maybe.applicativeMaybe)

  (scm:define altLast
    (scm:list (scm:cons (scm:string->symbol "alt") (rt:record-ref semigroupLast (scm:string->symbol "append"))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe))))

  (scm:define plusLast
    (scm:list (scm:cons (scm:string->symbol "empty") Data.Maybe.Nothing) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
      altLast))))

  (scm:define alternativeLast
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      Data.Maybe.applicativeMaybe)) (scm:cons (scm:string->symbol "Plus1") (scm:lambda (_)
      plusLast)))))
