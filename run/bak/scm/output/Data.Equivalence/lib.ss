#!r6rs
#!chezscheme
(library
  (Data.Equivalence lib)
  (export
    Equivalence
    comparisonEquivalence
    contravariantEquivalence
    defaultEquivalence
    monoidEquivalence
    newtypeEquivalence
    semigroupEquivalence)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define Equivalence
    (scm:lambda (x0)
      x0))

  (scm:define semigroupEquivalence
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:and ((v0 a2) b3) ((v11 a2) b3)))))))))

  (scm:define newtypeEquivalence
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidEquivalence
    (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v0)
      (scm:lambda (v11)
        #t))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      semigroupEquivalence))))

  (scm:define defaultEquivalence
    (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))

  (scm:define contravariantEquivalence
    (scm:list (scm:cons (scm:string->symbol "cmap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (x2)
          (scm:lambda (y3)
            ((v1 (f0 x2)) (f0 y3)))))))))

  (scm:define comparisonEquivalence
    (scm:lambda (v0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (Data.Ordering.EQ? ((v0 a1) b2)))))))
