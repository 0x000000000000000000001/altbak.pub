#!r6rs
#!chezscheme
(library
  (Data.Monoid.Alternate lib)
  (export
    Alternate
    altAlternate
    alternativeAlternate
    applicativeAlternate
    applyAlternate
    bindAlternate
    boundedAlternate
    comonadAlternate
    eq1Alternate
    eqAlternate
    extendAlternate
    functorAlternate
    monadAlternate
    monoidAlternate
    newtypeAlternate
    ord1Alternate
    ordAlternate
    plusAlternate
    semigroupAlternate
    showAlternate)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Alternate
    (scm:lambda (x0)
      x0))

  (scm:define showAlternate
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Alternate ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupAlternate
    (scm:lambda (dictAlt0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictAlt0 (scm:string->symbol "alt")) v1) v12)))))))

  (scm:define plusAlternate
    (scm:lambda (dictPlus0)
      dictPlus0))

  (scm:define ordAlternate
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define ord1Alternate
    (scm:lambda (dictOrd10)
      dictOrd10))

  (scm:define newtypeAlternate
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidAlternate
    (scm:lambda (dictPlus0)
      (scm:let*
        ([_1 ((rt:record-ref dictPlus0 (scm:string->symbol "Alt0")) (scm:quote undefined))]
         [semigroupAlternate12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref _1 (scm:string->symbol "alt")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictPlus0 (scm:string->symbol "empty"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupAlternate12))))))

  (scm:define monadAlternate
    (scm:lambda (dictMonad0)
      dictMonad0))

  (scm:define functorAlternate
    (scm:lambda (dictFunctor0)
      dictFunctor0))

  (scm:define extendAlternate
    (scm:lambda (dictExtend0)
      dictExtend0))

  (scm:define eqAlternate
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define eq1Alternate
    (scm:lambda (dictEq10)
      dictEq10))

  (scm:define comonadAlternate
    (scm:lambda (dictComonad0)
      dictComonad0))

  (scm:define boundedAlternate
    (scm:lambda (dictBounded0)
      dictBounded0))

  (scm:define bindAlternate
    (scm:lambda (dictBind0)
      dictBind0))

  (scm:define applyAlternate
    (scm:lambda (dictApply0)
      dictApply0))

  (scm:define applicativeAlternate
    (scm:lambda (dictApplicative0)
      dictApplicative0))

  (scm:define alternativeAlternate
    (scm:lambda (dictAlternative0)
      dictAlternative0))

  (scm:define altAlternate
    (scm:lambda (dictAlt0)
      dictAlt0)))
