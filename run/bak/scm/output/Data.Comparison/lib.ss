#!r6rs
#!chezscheme
(library
  (Data.Comparison lib)
  (export
    Comparison
    contravariantComparison
    defaultComparison
    monoidComparison
    newtypeComparison
    semigroupComparison)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define Comparison
    (scm:lambda (x0)
      x0))

  (scm:define semigroupComparison
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (x2)
          (scm:let*
            ([_3 (v0 x2)]
             [_4 (v11 x2)])
              (scm:lambda (x5)
                (scm:let*
                  ([_6 (_3 x5)]
                   [_7 (_4 x5)])
                    (scm:cond
                      [(Data.Ordering.LT? _6) Data.Ordering.LT]
                      [(Data.Ordering.GT? _6) Data.Ordering.GT]
                      [(Data.Ordering.EQ? _6) _7]
                      [scm:else (rt:fail)]))))))))))

  (scm:define newtypeComparison
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidComparison
    (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Ordering.EQ))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      semigroupComparison))))

  (scm:define defaultComparison
    (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare"))))

  (scm:define contravariantComparison
    (scm:list (scm:cons (scm:string->symbol "cmap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((v1 (f0 x2)) (f0 y3))))))))))
