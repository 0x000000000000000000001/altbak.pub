#!r6rs
#!chezscheme
(library
  (Data.Monoid.Conj lib)
  (export
    Conj
    applicativeConj
    applyConj
    bindConj
    boundedConj
    eq1Conj
    eqConj
    functorConj
    monadConj
    monoidConj
    ord1Conj
    ordConj
    semigroupConj
    semiringConj
    showConj)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Conj
    (scm:lambda (x0)
      x0))

  (scm:define showConj
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Conj ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semiringConj
    (scm:lambda (dictHeytingAlgebra0)
      (scm:list (scm:cons (scm:string->symbol "zero") (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "tt"))) (scm:cons (scm:string->symbol "one") (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "ff"))) (scm:cons (scm:string->symbol "add") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "conj")) v1) v12)))) (scm:cons (scm:string->symbol "mul") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "disj")) v1) v12)))))))

  (scm:define semigroupConj
    (scm:lambda (dictHeytingAlgebra0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "conj")) v1) v12)))))))

  (scm:define ordConj
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define monoidConj
    (scm:lambda (dictHeytingAlgebra0)
      (scm:let ([semigroupConj11 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "conj")) v1) v12)))))])
        (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "tt"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
          semigroupConj11))))))

  (scm:define functorConj
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1))))))

  (scm:define eqConj
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Conj
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))))

  (scm:define ord1Conj
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Conj))))

  (scm:define boundedConj
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define applyConj
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (v0 v11)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorConj))))

  (scm:define bindConj
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (f1 v0)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyConj))))

  (scm:define applicativeConj
    (scm:list (scm:cons (scm:string->symbol "pure") Conj) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyConj))))

  (scm:define monadConj
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeConj)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindConj)))))
