#!r6rs
#!chezscheme
(library
  (Data.Const lib)
  (export
    Const
    applicativeConst
    applyConst
    booleanAlgebraConst
    boundedConst
    commutativeRingConst
    eq1Const
    eqConst
    euclideanRingConst
    functorConst
    heytingAlgebraConst
    invariantConst
    monoidConst
    newtypeConst
    ord1Const
    ordConst
    ringConst
    semigroupConst
    semigroupoidConst
    semiringConst
    showConst)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Const
    (scm:lambda (x0)
      x0))

  (scm:define showConst
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Const ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semiringConst
    (scm:lambda (dictSemiring0)
      dictSemiring0))

  (scm:define semigroupoidConst
    (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v0)
      (scm:lambda (v11)
        v11)))))

  (scm:define semigroupConst
    (scm:lambda (dictSemigroup0)
      dictSemigroup0))

  (scm:define ringConst
    (scm:lambda (dictRing0)
      dictRing0))

  (scm:define ordConst
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define newtypeConst
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidConst
    (scm:lambda (dictMonoid0)
      dictMonoid0))

  (scm:define heytingAlgebraConst
    (scm:lambda (dictHeytingAlgebra0)
      dictHeytingAlgebra0))

  (scm:define functorConst
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        m1)))))

  (scm:define invariantConst
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (m2)
          m2))))))

  (scm:define euclideanRingConst
    (scm:lambda (dictEuclideanRing0)
      dictEuclideanRing0))

  (scm:define eqConst
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Const
    (scm:lambda (dictEq0)
      (scm:let ([eq1 (rt:record-ref dictEq0 (scm:string->symbol "eq"))])
        (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq12)
          eq1))))))

  (scm:define ord1Const
    (scm:lambda (dictOrd0)
      (scm:let*
        ([compare1 (rt:record-ref dictOrd0 (scm:string->symbol "compare"))]
         [eq2 (rt:record-ref ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined)) (scm:string->symbol "eq"))])
          (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd13)
            compare1)) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
            (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq14)
              eq2)))))))))

  (scm:define commutativeRingConst
    (scm:lambda (dictCommutativeRing0)
      dictCommutativeRing0))

  (scm:define boundedConst
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define booleanAlgebraConst
    (scm:lambda (dictBooleanAlgebra0)
      dictBooleanAlgebra0))

  (scm:define applyConst
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) v1) v12)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
        functorConst)))))

  (scm:define applicativeConst
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [applyConst13 (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v3)
          (scm:lambda (v14)
            (((rt:record-ref _2 (scm:string->symbol "append")) v3) v14)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
          functorConst)))])
          (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (v4)
            mempty1)) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
            applyConst13)))))))
