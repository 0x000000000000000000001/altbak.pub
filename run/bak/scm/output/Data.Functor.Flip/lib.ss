#!r6rs
#!chezscheme
(library
  (Data.Functor.Flip lib)
  (export
    Flip
    biapplicativeFlip
    biapplyFlip
    bifunctorFlip
    categoryFlip
    contravariantFlip
    eqFlip
    functorFlip
    newtypeFlip
    ordFlip
    semigroupoidFlip
    showFlip)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Bifunctor lib) Data.Bifunctor.)
    (prefix (Data.Profunctor lib) Data.Profunctor.))

  (scm:define Flip
    (scm:lambda (x0)
      x0))

  (scm:define showFlip
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Flip ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupoidFlip
    (scm:lambda (dictSemigroupoid0)
      (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictSemigroupoid0 (scm:string->symbol "compose")) v12) v1)))))))

  (scm:define ordFlip
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define newtypeFlip
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define functorFlip
    (scm:lambda (dictBifunctor0)
      (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f1)
        (scm:lambda (v2)
          ((((rt:record-ref dictBifunctor0 (scm:string->symbol "bimap")) f1) Data.Bifunctor.identity) v2)))))))

  (scm:define eqFlip
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define contravariantFlip
    (scm:lambda (dictProfunctor0)
      (scm:list (scm:cons (scm:string->symbol "cmap") (scm:lambda (f1)
        (scm:lambda (v2)
          ((((rt:record-ref dictProfunctor0 (scm:string->symbol "dimap")) f1) Data.Profunctor.identity) v2)))))))

  (scm:define categoryFlip
    (scm:lambda (dictCategory0)
      (scm:let*
        ([_1 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))]
         [semigroupoidFlip12 (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref _1 (scm:string->symbol "compose")) v13) v2)))))])
          (scm:list (scm:cons (scm:string->symbol "identity") (rt:record-ref dictCategory0 (scm:string->symbol "identity"))) (scm:cons (scm:string->symbol "Semigroupoid0") (scm:lambda (_)
            semigroupoidFlip12))))))

  (scm:define bifunctorFlip
    (scm:lambda (dictBifunctor0)
      (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            ((((rt:record-ref dictBifunctor0 (scm:string->symbol "bimap")) g2) f1) v3))))))))

  (scm:define biapplyFlip
    (scm:lambda (dictBiapply0)
      (scm:let*
        ([_1 ((rt:record-ref dictBiapply0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
         [bifunctorFlip12 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (v4)
              ((((rt:record-ref _1 (scm:string->symbol "bimap")) g3) f2) v4))))))])
          (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) v3) v14)))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorFlip12))))))

  (scm:define biapplicativeFlip
    (scm:lambda (dictBiapplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictBiapplicative0 (scm:string->symbol "Biapply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
         [biapplyFlip13 (scm:let ([bifunctorFlip13 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f3)
          (scm:lambda (g4)
            (scm:lambda (v5)
              ((((rt:record-ref _2 (scm:string->symbol "bimap")) g4) f3) v5))))))])
          (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v4)
            (scm:lambda (v15)
              (((rt:record-ref _1 (scm:string->symbol "biapply")) v4) v15)))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorFlip13))))])
          (scm:list (scm:cons (scm:string->symbol "bipure") (scm:lambda (a4)
            (scm:lambda (b5)
              (((rt:record-ref dictBiapplicative0 (scm:string->symbol "bipure")) b5) a4)))) (scm:cons (scm:string->symbol "Biapply0") (scm:lambda (_)
            biapplyFlip13)))))))
