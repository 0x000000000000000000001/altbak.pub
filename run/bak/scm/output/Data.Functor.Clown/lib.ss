#!r6rs
#!chezscheme
(library
  (Data.Functor.Clown lib)
  (export
    Clown
    biapplicativeClown
    biapplyClown
    bifunctorClown
    eqClown
    functorClown
    hoistClown
    newtypeClown
    ordClown
    profunctorClown
    showClown)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Clown
    (scm:lambda (x0)
      x0))

  (scm:define showClown
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Clown ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define profunctorClown
    (scm:lambda (dictContravariant0)
      (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f1)
        (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictContravariant0 (scm:string->symbol "cmap")) f1) v13))))))))

  (scm:define ordClown
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define newtypeClown
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define hoistClown
    (scm:lambda (f0)
      (scm:lambda (v1)
        (f0 v1))))

  (scm:define functorClown
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (v0)
      (scm:lambda (v11)
        v11)))))

  (scm:define eqClown
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define bifunctorClown
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f1)
        (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1) v13))))))))

  (scm:define biapplyClown
    (scm:lambda (dictApply0)
      (scm:let*
        ([_1 ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [bifunctorClown12 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref _1 (scm:string->symbol "map")) f2) v14))))))])
          (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref dictApply0 (scm:string->symbol "apply")) v3) v14)))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorClown12))))))

  (scm:define biapplicativeClown
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [biapplyClown13 (scm:let ([bifunctorClown13 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f3)
          (scm:lambda (v4)
            (scm:lambda (v15)
              (((rt:record-ref _2 (scm:string->symbol "map")) f3) v15))))))])
          (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v4)
            (scm:lambda (v15)
              (((rt:record-ref _1 (scm:string->symbol "apply")) v4) v15)))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorClown13))))])
          (scm:list (scm:cons (scm:string->symbol "bipure") (scm:lambda (a4)
            (scm:lambda (v5)
              ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) a4)))) (scm:cons (scm:string->symbol "Biapply0") (scm:lambda (_)
            biapplyClown13)))))))
