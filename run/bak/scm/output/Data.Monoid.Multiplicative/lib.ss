#!r6rs
#!chezscheme
(library
  (Data.Monoid.Multiplicative lib)
  (export
    Multiplicative
    applicativeMultiplicative
    applyMultiplicative
    bindMultiplicative
    boundedMultiplicative
    eq1Multiplicative
    eqMultiplicative
    functorMultiplicative
    monadMultiplicative
    monoidMultiplicative
    ord1Multiplicative
    ordMultiplicative
    semigroupMultiplicative
    showMultiplicative)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Multiplicative
    (scm:lambda (x0)
      x0))

  (scm:define showMultiplicative
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Multiplicative ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupMultiplicative
    (scm:lambda (dictSemiring0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemiring0 (scm:string->symbol "mul")) v1) v12)))))))

  (scm:define ordMultiplicative
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define monoidMultiplicative
    (scm:lambda (dictSemiring0)
      (scm:let ([semigroupMultiplicative11 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemiring0 (scm:string->symbol "mul")) v1) v12)))))])
        (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictSemiring0 (scm:string->symbol "one"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
          semigroupMultiplicative11))))))

  (scm:define functorMultiplicative
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1))))))

  (scm:define eqMultiplicative
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Multiplicative
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))))

  (scm:define ord1Multiplicative
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Multiplicative))))

  (scm:define boundedMultiplicative
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define applyMultiplicative
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (v0 v11)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorMultiplicative))))

  (scm:define bindMultiplicative
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (f1 v0)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyMultiplicative))))

  (scm:define applicativeMultiplicative
    (scm:list (scm:cons (scm:string->symbol "pure") Multiplicative) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyMultiplicative))))

  (scm:define monadMultiplicative
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeMultiplicative)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindMultiplicative)))))
