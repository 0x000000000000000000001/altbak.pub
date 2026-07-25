#!r6rs
#!chezscheme
(library
  (Data.Bifunctor.Join lib)
  (export
    Join
    biapplicativeJoin
    biapplyJoin
    bifunctorJoin
    eqJoin
    newtypeJoin
    ordJoin
    showJoin)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Join
    (scm:lambda (x0)
      x0))

  (scm:define showJoin
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Join ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define ordJoin
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define newtypeJoin
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define eqJoin
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define bifunctorJoin
    (scm:lambda (dictBifunctor0)
      (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f1)
        (scm:lambda (v2)
          ((((rt:record-ref dictBifunctor0 (scm:string->symbol "bimap")) f1) f1) v2)))))))

  (scm:define biapplyJoin
    (scm:lambda (dictBiapply0)
      (scm:let*
        ([_1 ((rt:record-ref dictBiapply0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
         [bifunctorJoin12 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            ((((rt:record-ref _1 (scm:string->symbol "bimap")) f2) f2) v3)))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) v3) v14)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            bifunctorJoin12))))))

  (scm:define biapplicativeJoin
    (scm:lambda (dictBiapplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictBiapplicative0 (scm:string->symbol "Biapply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
         [biapplyJoin13 (scm:let ([bifunctorJoin13 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f3)
          (scm:lambda (v4)
            ((((rt:record-ref _2 (scm:string->symbol "bimap")) f3) f3) v4)))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v4)
            (scm:lambda (v15)
              (((rt:record-ref _1 (scm:string->symbol "biapply")) v4) v15)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            bifunctorJoin13))))])
          (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a4)
            (((rt:record-ref dictBiapplicative0 (scm:string->symbol "bipure")) a4) a4))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
            biapplyJoin13)))))))
