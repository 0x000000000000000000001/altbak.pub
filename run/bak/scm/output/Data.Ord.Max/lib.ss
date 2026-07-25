#!r6rs
#!chezscheme
(library
  (Data.Ord.Max lib)
  (export
    Max
    eqMax
    monoidMax
    newtypeMax
    ordMax
    semigroupMax
    showMax)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define Max
    (scm:lambda (x0)
      x0))

  (scm:define showMax
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Max ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupMax
    (scm:lambda (dictOrd0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:let ([v3 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) v1) v12)])
            (scm:cond
              [(Data.Ordering.LT? v3) v12]
              [(Data.Ordering.EQ? v3) v1]
              [(Data.Ordering.GT? v3) v1]
              [scm:else (rt:fail)]))))))))

  (scm:define newtypeMax
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidMax
    (scm:lambda (dictBounded0)
      (scm:let*
        ([_1 ((rt:record-ref dictBounded0 (scm:string->symbol "Ord0")) (scm:quote undefined))]
         [semigroupMax12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:let ([v4 (((rt:record-ref _1 (scm:string->symbol "compare")) v2) v13)])
              (scm:cond
                [(Data.Ordering.LT? v4) v13]
                [(Data.Ordering.EQ? v4) v2]
                [(Data.Ordering.GT? v4) v2]
                [scm:else (rt:fail)]))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictBounded0 (scm:string->symbol "bottom"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupMax12))))))

  (scm:define eqMax
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define ordMax
    (scm:lambda (dictOrd0)
      (scm:let ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))])
        (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) v2) v13)))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
          _1)))))))
