#!r6rs
#!chezscheme
(library
  (Data.Predicate lib)
  (export
    Predicate
    booleanAlgebraPredicate
    contravariantPredicate
    heytingAlgebraPredicate
    newtypePredicate)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.BooleanAlgebra lib) Data.BooleanAlgebra.))

  (scm:define Predicate
    (scm:lambda (x0)
      x0))

  (scm:define newtypePredicate
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define heytingAlgebraPredicate
    (scm:list (scm:cons (scm:string->symbol "ff") (scm:lambda (v0)
      #f)) (scm:cons (scm:string->symbol "tt") (scm:lambda (v0)
      #t)) (scm:cons (scm:string->symbol "implies") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (a2)
          (scm:or (scm:not (f0 a2)) (g1 a2)))))) (scm:cons (scm:string->symbol "conj") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (a2)
          (scm:and (f0 a2) (g1 a2)))))) (scm:cons (scm:string->symbol "disj") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (a2)
          (scm:or (f0 a2) (g1 a2)))))) (scm:cons (scm:string->symbol "not") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:not (f0 a1)))))))

  (scm:define contravariantPredicate
    (scm:list (scm:cons (scm:string->symbol "cmap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (x2)
          (v1 (f0 x2))))))))

  (scm:define booleanAlgebraPredicate
    (Data.BooleanAlgebra.booleanAlgebraFn Data.BooleanAlgebra.booleanAlgebraBoolean)))
