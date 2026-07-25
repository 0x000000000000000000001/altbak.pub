#!r6rs
#!chezscheme
(library
  (Data.Maybe.First lib)
  (export
    First
    altFirst
    alternativeFirst
    applicativeFirst
    applyFirst
    bindFirst
    boundedFirst
    eq1First
    eqFirst
    extendFirst
    functorFirst
    invariantFirst
    monadFirst
    monoidFirst
    newtypeFirst
    ord1First
    ordFirst
    plusFirst
    semigroupFirst
    showFirst)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define First
    (scm:lambda (x0)
      x0))

  (scm:define showFirst
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (scm:cond
          [(Data.Maybe.Just? v1) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "First ((Just ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Data.Maybe.Just-value0 v1))) (rt:string->pstring "))"))]
          [(Data.Maybe.Nothing? v1) (rt:string->pstring "First (Nothing)")]
          [scm:else (rt:fail)]))))))

  (scm:define semigroupFirst
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Data.Maybe.Just? v0) v0]
          [scm:else v11]))))))

  (scm:define ordFirst
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

  (scm:define ord1First
    Data.Maybe.ord1Maybe)

  (scm:define newtypeFirst
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidFirst
    (scm:list (scm:cons (scm:string->symbol "mempty") Data.Maybe.Nothing) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      semigroupFirst))))

  (scm:define monadFirst
    Data.Maybe.monadMaybe)

  (scm:define invariantFirst
    Data.Maybe.invariantMaybe)

  (scm:define functorFirst
    Data.Maybe.functorMaybe)

  (scm:define extendFirst
    Data.Maybe.extendMaybe)

  (scm:define eqFirst
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:cond
            [(Data.Maybe.Nothing? x1) (Data.Maybe.Nothing? y2)]
            [scm:else (scm:and (Data.Maybe.Just? x1) (scm:and (Data.Maybe.Just? y2) (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Data.Maybe.Just-value0 x1)) (Data.Maybe.Just-value0 y2))))])))))))

  (scm:define eq1First
    Data.Maybe.eq1Maybe)

  (scm:define boundedFirst
    (scm:lambda (dictBounded0)
      (Data.Maybe.boundedMaybe dictBounded0)))

  (scm:define bindFirst
    Data.Maybe.bindMaybe)

  (scm:define applyFirst
    Data.Maybe.applyMaybe)

  (scm:define applicativeFirst
    Data.Maybe.applicativeMaybe)

  (scm:define altFirst
    (scm:list (scm:cons (scm:string->symbol "alt") (rt:record-ref semigroupFirst (scm:string->symbol "append"))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe))))

  (scm:define plusFirst
    (scm:list (scm:cons (scm:string->symbol "empty") Data.Maybe.Nothing) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
      altFirst))))

  (scm:define alternativeFirst
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      Data.Maybe.applicativeMaybe)) (scm:cons (scm:string->symbol "Plus1") (scm:lambda (_)
      plusFirst)))))
