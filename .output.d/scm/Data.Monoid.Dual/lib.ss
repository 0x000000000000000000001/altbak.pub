#!r6rs
#!chezscheme
(library
  (Data.Monoid.Dual lib)
  (export
    Dual
    applicativeDual
    applyDual
    bindDual
    boundedDual
    eq1Dual
    eqDual
    functorDual
    monadDual
    monoidDual
    ord1Dual
    ordDual
    semigroupDual
    showDual)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Dual
    (scm:lambda (x0)
      x0))

  (scm:define showDual
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Dual ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupDual
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) v12) v1)))))))

  (scm:define ordDual
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define monoidDual
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupDual12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref _1 (scm:string->symbol "append")) v13) v2)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupDual12))))))

  (scm:define functorDual
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1))))))

  (scm:define eqDual
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Dual
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))))

  (scm:define ord1Dual
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Dual))))

  (scm:define boundedDual
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define applyDual
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (v0 v11)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorDual))))

  (scm:define bindDual
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (f1 v0)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyDual))))

  (scm:define applicativeDual
    (scm:list (scm:cons (scm:string->symbol "pure") Dual) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyDual))))

  (scm:define monadDual
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeDual)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindDual)))))
