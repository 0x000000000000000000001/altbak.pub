#!r6rs
#!chezscheme
(library
  (Data.Functor.Joker lib)
  (export
    Joker
    applicativeJoker
    applyJoker
    biapplicativeJoker
    biapplyJoker
    bifunctorJoker
    bindJoker
    choiceJoker
    eqJoker
    functorJoker
    hoistJoker
    monadJoker
    newtypeJoker
    ordJoker
    profunctorJoker
    showJoker)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.))

  (scm:define Joker
    (scm:lambda (x0)
      x0))

  (scm:define showJoker
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Joker ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define profunctorJoker
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (v1)
        (scm:lambda (g2)
          (scm:lambda (v13)
            (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) g2) v13))))))))

  (scm:define ordJoker
    (scm:lambda (dictOrd0)
      dictOrd0))

  (scm:define newtypeJoker
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define hoistJoker
    (scm:lambda (f0)
      (scm:lambda (v1)
        (f0 v1))))

  (scm:define functorJoker
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1) v2)))))))

  (scm:define eqJoker
    (scm:lambda (dictEq0)
      dictEq0))

  (scm:define choiceJoker
    (scm:lambda (dictFunctor0)
      (scm:let ([profunctorJoker11 (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (v1)
        (scm:lambda (g2)
          (scm:lambda (v13)
            (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) g2) v13))))))])
        (scm:list (scm:cons (scm:string->symbol "left") (scm:lambda (v2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) Data.Either.Left) v2))) (scm:cons (scm:string->symbol "right") (scm:lambda (v2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) Data.Either.Right) v2))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
          profunctorJoker11))))))

  (scm:define bifunctorJoker
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (v1)
        (scm:lambda (g2)
          (scm:lambda (v13)
            (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) g2) v13))))))))

  (scm:define biapplyJoker
    (scm:lambda (dictApply0)
      (scm:let*
        ([_1 ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [bifunctorJoker12 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (v2)
          (scm:lambda (g3)
            (scm:lambda (v14)
              (((rt:record-ref _1 (scm:string->symbol "map")) g3) v14))))))])
          (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref dictApply0 (scm:string->symbol "apply")) v3) v14)))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorJoker12))))))

  (scm:define biapplicativeJoker
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [biapplyJoker13 (scm:let ([bifunctorJoker13 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (v3)
          (scm:lambda (g4)
            (scm:lambda (v15)
              (((rt:record-ref _2 (scm:string->symbol "map")) g4) v15))))))])
          (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v4)
            (scm:lambda (v15)
              (((rt:record-ref _1 (scm:string->symbol "apply")) v4) v15)))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorJoker13))))])
          (scm:list (scm:cons (scm:string->symbol "bipure") (scm:lambda (v4)
            (scm:lambda (b5)
              ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) b5)))) (scm:cons (scm:string->symbol "Biapply0") (scm:lambda (_)
            biapplyJoker13))))))

  (scm:define applyJoker
    (scm:lambda (dictApply0)
      (scm:let*
        ([_1 ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorJoker12 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (((rt:record-ref _1 (scm:string->symbol "map")) f2) v3)))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref dictApply0 (scm:string->symbol "apply")) v3) v14)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorJoker12))))))

  (scm:define bindJoker
    (scm:lambda (dictBind0)
      (scm:let*
        ([_1 ((rt:record-ref dictBind0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [applyJoker13 (scm:let ([functorJoker13 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f3)
          (scm:lambda (v4)
            (((rt:record-ref _2 (scm:string->symbol "map")) f3) v4)))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v4)
            (scm:lambda (v15)
              (((rt:record-ref _1 (scm:string->symbol "apply")) v4) v15)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorJoker13))))])
          (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v4)
            (scm:lambda (amb5)
              (((rt:record-ref dictBind0 (scm:string->symbol "bind")) v4) (scm:lambda (x6)
                (amb5 x6)))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
            applyJoker13))))))

  (scm:define applicativeJoker
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [applyJoker13 (scm:let ([functorJoker13 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f3)
          (scm:lambda (v4)
            (((rt:record-ref _2 (scm:string->symbol "map")) f3) v4)))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v4)
            (scm:lambda (v15)
              (((rt:record-ref _1 (scm:string->symbol "apply")) v4) v15)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorJoker13))))])
          (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (x4)
            ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) x4))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
            applyJoker13))))))

  (scm:define monadJoker
    (scm:lambda (dictMonad0)
      (scm:let*
        ([_1 ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [applicativeJoker13 (scm:let*
          ([_3 ((rt:record-ref _2 (scm:string->symbol "Functor0")) (scm:quote undefined))]
           [functorJoker14 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
            (scm:lambda (v5)
              (((rt:record-ref _3 (scm:string->symbol "map")) f4) v5)))))]
           [applyJoker15 (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v5)
            (scm:lambda (v16)
              (((rt:record-ref _2 (scm:string->symbol "apply")) v5) v16)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorJoker14)))])
            (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (x6)
              ((rt:record-ref _1 (scm:string->symbol "pure")) x6))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
              applyJoker15))))]
         [_4 ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined))]
         [_5 ((rt:record-ref _4 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [bindJoker16 (scm:let*
          ([_6 ((rt:record-ref _5 (scm:string->symbol "Functor0")) (scm:quote undefined))]
           [functorJoker17 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f7)
            (scm:lambda (v8)
              (((rt:record-ref _6 (scm:string->symbol "map")) f7) v8)))))]
           [applyJoker18 (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v8)
            (scm:lambda (v19)
              (((rt:record-ref _5 (scm:string->symbol "apply")) v8) v19)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorJoker17)))])
            (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v9)
              (scm:lambda (amb10)
                (((rt:record-ref _4 (scm:string->symbol "bind")) v9) (scm:lambda (x11)
                  (amb10 x11)))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
              applyJoker18))))])
          (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
            applicativeJoker13)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
            bindJoker16)))))))
