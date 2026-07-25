#!r6rs
#!chezscheme
(library
  (Data.Identity lib)
  (export
    Identity
    altIdentity
    applicativeIdentity
    applyIdentity
    bindIdentity
    booleanAlgebraIdentity
    boundedIdentity
    commutativeRingIdentity
    comonadIdentity
    eq1Identity
    eqIdentity
    euclideanRingIdentity
    extendIdentity
    functorIdentity
    heytingAlgebraIdentity
    invariantIdentity
    lazyIdentity
    monadIdentity
    monoidIdentity
    newtypeIdentity
    ord1Identity
    ordIdentity
    ringIdentity
    semigroupIdentity
    semiringIdentity
    showIdentity)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Identity
    (scm:lambda (x0)
      x0))

  (scm:define showIdentity
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Identity ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semiringIdentity
    (scm:lambda (dictSemiring0)
      dictSemiring0))

  (scm:define semigroupIdentity
    (scm:lambda (dictSemigroup0)
      dictSemigroup0))

  (scm:define ringIdentity
    (scm:lambda (dictRing0)
      dictRing0))

  (scm:define ordIdentity
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define newtypeIdentity
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidIdentity
    (scm:lambda (dictMonoid0)
      dictMonoid0))

  (scm:define lazyIdentity
    (scm:lambda (dictLazy0)
      dictLazy0))

  (scm:define heytingAlgebraIdentity
    (scm:lambda (dictHeytingAlgebra0)
      dictHeytingAlgebra0))

  (scm:define functorIdentity
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1))))))

  (scm:define invariantIdentity
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (m2)
          (f0 m2)))))))

  (scm:define extendIdentity
    (scm:list (scm:cons (scm:string->symbol "extend") (scm:lambda (f0)
      (scm:lambda (m1)
        (f0 m1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorIdentity))))

  (scm:define euclideanRingIdentity
    (scm:lambda (dictEuclideanRing0)
      dictEuclideanRing0))

  (scm:define eqIdentity
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Identity
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (rt:record-ref dictEq0 (scm:string->symbol "eq"))))))

  (scm:define ord1Identity
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref dictOrd0 (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Identity))))

  (scm:define comonadIdentity
    (scm:list (scm:cons (scm:string->symbol "extract") (scm:lambda (v0)
      v0)) (scm:cons (scm:string->symbol "Extend0") (scm:lambda (_)
      extendIdentity))))

  (scm:define commutativeRingIdentity
    (scm:lambda (dictCommutativeRing0)
      dictCommutativeRing0))

  (scm:define boundedIdentity
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define booleanAlgebraIdentity
    (scm:lambda (dictBooleanAlgebra0)
      dictBooleanAlgebra0))

  (scm:define applyIdentity
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v0)
      (scm:lambda (v11)
        (v0 v11)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorIdentity))))

  (scm:define bindIdentity
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v0)
      (scm:lambda (f1)
        (f1 v0)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyIdentity))))

  (scm:define applicativeIdentity
    (scm:list (scm:cons (scm:string->symbol "pure") Identity) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyIdentity))))

  (scm:define monadIdentity
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeIdentity)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindIdentity))))

  (scm:define altIdentity
    (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (x0)
      (scm:lambda (v1)
        x0))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorIdentity)))))
