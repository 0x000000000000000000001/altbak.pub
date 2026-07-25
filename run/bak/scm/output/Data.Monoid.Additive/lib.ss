#!r6rs
#!chezscheme
(library
  (Data.Monoid.Additive lib)
  (export
    Additive
    applicativeAdditive
    applyAdditive
    bindAdditive
    boundedAdditive
    eq1Additive
    eqAdditive
    functorAdditive
    monadAdditive
    monoidAdditive
    ord1Additive
    ordAdditive
    semigroupAdditive
    showAdditive)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Additive
    (scm:lambda (x0)
      x0))

  (scm:define showAdditive
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Additive ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupAdditive
    (scm:lambda (dictSemiring0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemiring0 (scm:string->symbol "add")) v1) v12)))))))

  (scm:define ordAdditive
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define monoidAdditive
    (scm:lambda (dictSemiring0)
      (scm:let ([semigroupAdditive11 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemiring0 (scm:string->symbol "add")) v1) v12)))))])
        (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictSemiring0 (scm:string->symbol "zero"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
          semigroupAdditive11))))))

  (scm:define functorAdditive
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1))))))

  (scm:define eqAdditive
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Additive
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))))

  (scm:define ord1Additive
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Additive))))

  (scm:define boundedAdditive
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define applyAdditive
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (v0 v11)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorAdditive))))

  (scm:define bindAdditive
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (f1 v0)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyAdditive))))

  (scm:define applicativeAdditive
    (scm:list (scm:cons (scm:string->symbol "pure") Additive) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyAdditive))))

  (scm:define monadAdditive
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeAdditive)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindAdditive)))))
