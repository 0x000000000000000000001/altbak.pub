#!r6rs
#!chezscheme
(library
  (Data.Functor.App lib)
  (export
    App
    altApp
    alternativeApp
    applicativeApp
    applyApp
    bindApp
    comonadApp
    eq1App
    eqApp
    extendApp
    functorApp
    hoistApp
    hoistLiftApp
    hoistLowerApp
    lazyApp
    monadApp
    monadPlusApp
    monoidApp
    newtypeApp
    ord1App
    ordApp
    plusApp
    semigroupApp
    showApp)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define App
    (scm:lambda (x0)
      x0))

  (scm:define showApp
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(App ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupApp
    (scm:lambda (dictApply0)
      (scm:lambda (dictSemigroup1)
        (scm:let ([append12 (rt:record-ref dictSemigroup1 (scm:string->symbol "append"))])
          (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) append12) v3)) v14)))))))))

  (scm:define plusApp
    (scm:lambda (dictPlus0)
      dictPlus0))

  (scm:define newtypeApp
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define monoidApp
    (scm:lambda (dictApplicative0)
      (scm:let ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))])
        (scm:lambda (dictMonoid2)
          (scm:let*
            ([append13 (rt:record-ref ((rt:record-ref dictMonoid2 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append"))]
             [semigroupApp24 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v4)
              (scm:lambda (v15)
                (((rt:record-ref _1 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) append13) v4)) v15)))))])
              (scm:list (scm:cons (scm:string->symbol "mempty") ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) (rt:record-ref dictMonoid2 (scm:string->symbol "mempty")))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
                semigroupApp24))))))))

  (scm:define monadPlusApp
    (scm:lambda (dictMonadPlus0)
      dictMonadPlus0))

  (scm:define monadApp
    (scm:lambda (dictMonad0)
      dictMonad0))

  (scm:define lazyApp
    (scm:lambda (dictLazy0)
      dictLazy0))

  (scm:define hoistLowerApp
    Unsafe.Coerce.unsafeCoerce)

  (scm:define hoistLiftApp
    Unsafe.Coerce.unsafeCoerce)

  (scm:define hoistApp
    (scm:lambda (f0)
      (scm:lambda (v1)
        (f0 v1))))

  (scm:define functorApp
    (scm:lambda (dictFunctor0)
      dictFunctor0))

  (scm:define extendApp
    (scm:lambda (dictExtend0)
      dictExtend0))

  (scm:define eqApp
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq1)
        (scm:let ([eq112 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq1)])
          (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x3)
            (scm:lambda (y4)
              ((eq112 x3) y4)))))))))

  (scm:define ordApp
    (scm:lambda (dictOrd10)
      (scm:let ([_1 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
        (scm:lambda (dictOrd2)
          (scm:let*
            ([compare113 ((rt:record-ref dictOrd10 (scm:string->symbol "compare1")) dictOrd2)]
             [eq114 ((rt:record-ref _1 (scm:string->symbol "eq1")) ((rt:record-ref dictOrd2 (scm:string->symbol "Eq0")) (scm:quote undefined)))]
             [eqApp25 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x5)
              (scm:lambda (y6)
                ((eq114 x5) y6)))))])
              (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x6)
                (scm:lambda (y7)
                  ((compare113 x6) y7)))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                eqApp25))))))))

  (scm:define eq1App
    (scm:lambda (dictEq10)
      (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq1)
        ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq1))))))

  (scm:define ord1App
    (scm:lambda (dictOrd10)
      (scm:let*
        ([_1 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))]
         [_2 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))]
         [eq1App13 (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq3)
          ((rt:record-ref _2 (scm:string->symbol "eq1")) dictEq3))))])
          (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd4)
            (scm:let*
              ([compare115 ((rt:record-ref dictOrd10 (scm:string->symbol "compare1")) dictOrd4)]
               [eq116 ((rt:record-ref _1 (scm:string->symbol "eq1")) ((rt:record-ref dictOrd4 (scm:string->symbol "Eq0")) (scm:quote undefined)))])
                (rt:record-ref (scm:let ([eqApp27 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x7)
                  (scm:lambda (y8)
                    ((eq116 x7) y8)))))])
                  (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x8)
                    (scm:lambda (y9)
                      ((compare115 x8) y9)))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                    eqApp27)))) (scm:string->symbol "compare"))))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
            eq1App13))))))

  (scm:define comonadApp
    (scm:lambda (dictComonad0)
      dictComonad0))

  (scm:define bindApp
    (scm:lambda (dictBind0)
      dictBind0))

  (scm:define applyApp
    (scm:lambda (dictApply0)
      dictApply0))

  (scm:define applicativeApp
    (scm:lambda (dictApplicative0)
      dictApplicative0))

  (scm:define alternativeApp
    (scm:lambda (dictAlternative0)
      dictAlternative0))

  (scm:define altApp
    (scm:lambda (dictAlt0)
      dictAlt0)))
