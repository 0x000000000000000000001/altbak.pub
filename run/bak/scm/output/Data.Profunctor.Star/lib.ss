#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Star lib)
  (export
    Star
    altStar
    alternativeStar
    applicativeStar
    applyStar
    bindStar
    categoryStar
    choiceStar
    closedStar
    distributiveStar
    functorStar
    hoistStar
    invariantStar
    monadPlusStar
    monadStar
    newtypeStar
    plusStar
    profunctorStar
    semigroupoidStar
    strongStar)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define Star
    (scm:lambda (x0)
      x0))

  (scm:define semigroupoidStar
    (scm:lambda (dictBind0)
      (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (x3)
            (((rt:record-ref dictBind0 (scm:string->symbol "bind")) (v12 x3)) v1))))))))

  (scm:define profunctorStar
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) g2)])
              (scm:lambda (x5)
                (_4 (v3 (f1 x5))))))))))))

  (scm:define strongStar
    (scm:lambda (dictFunctor0)
      (scm:let ([profunctorStar11 (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) g2)])
              (scm:lambda (x5)
                (_4 (v3 (f1 x5))))))))))])
        (scm:list (scm:cons (scm:string->symbol "first") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:let ([_4 (Data.Tuple.Tuple-value1 v13)])
              (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (v25)
                (Data.Tuple.Tuple* v25 _4))) (v2 (Data.Tuple.Tuple-value0 v13))))))) (scm:cons (scm:string->symbol "second") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (Data.Tuple.Tuple (Data.Tuple.Tuple-value0 v13))) (v2 (Data.Tuple.Tuple-value1 v13)))))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
          profunctorStar11))))))

  (scm:define newtypeStar
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define invariantStar
    (scm:lambda (dictInvariant0)
      (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            (scm:let ([_4 (((rt:record-ref dictInvariant0 (scm:string->symbol "imap")) f1) g2)])
              (scm:lambda (x5)
                (_4 (v3 x5)))))))))))

  (scm:define hoistStar
    (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (x2)
          (f0 (v1 x2))))))

  (scm:define functorStar
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f1)
        (scm:lambda (v2)
          (scm:let ([_3 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1)])
            (scm:lambda (x4)
              (_3 (v2 x4))))))))))

  (scm:define distributiveStar
    (scm:lambda (dictDistributive0)
      (scm:let*
        ([_1 ((rt:record-ref dictDistributive0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorStar12 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref _1 (scm:string->symbol "map")) f2)])
              (scm:lambda (x5)
                (_4 (v3 x5))))))))])
          (scm:list (scm:cons (scm:string->symbol "distribute") (scm:lambda (dictFunctor3)
            (scm:let ([collect14 ((rt:record-ref dictDistributive0 (scm:string->symbol "collect")) dictFunctor3)])
              (scm:lambda (f5)
                (scm:lambda (a6)
                  ((collect14 (scm:lambda (v7)
                    (v7 a6))) f5)))))) (scm:cons (scm:string->symbol "collect") (scm:lambda (dictFunctor3)
            (scm:lambda (f4)
              (scm:let*
                ([_5 ((rt:record-ref (distributiveStar dictDistributive0) (scm:string->symbol "distribute")) dictFunctor3)]
                 [_6 ((rt:record-ref dictFunctor3 (scm:string->symbol "map")) f4)])
                  (scm:lambda (x7)
                    (_5 (_6 x7))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar12))))))

  (scm:define closedStar
    (scm:lambda (dictDistributive0)
      (scm:let*
        ([distribute1 ((rt:record-ref dictDistributive0 (scm:string->symbol "distribute")) Data.Functor.functorFn)]
         [_2 ((rt:record-ref dictDistributive0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [profunctorStar13 (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f3)
          (scm:lambda (g4)
            (scm:lambda (v5)
              (scm:let ([_6 ((rt:record-ref _2 (scm:string->symbol "map")) g4)])
                (scm:lambda (x7)
                  (_6 (v5 (f3 x7))))))))))])
          (scm:list (scm:cons (scm:string->symbol "closed") (scm:lambda (v4)
            (scm:lambda (g5)
              (distribute1 (scm:lambda (x6)
                (v4 (g5 x6))))))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
            profunctorStar13))))))

  (scm:define choiceStar
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([Functor01 ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [profunctorStar12 (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (v4)
              (scm:let ([_5 ((rt:record-ref Functor01 (scm:string->symbol "map")) g3)])
                (scm:lambda (x6)
                  (_5 (v4 (f2 x6))))))))))])
          (scm:list (scm:cons (scm:string->symbol "left") (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref Functor01 (scm:string->symbol "map")) Data.Either.Left)])
              (scm:lambda (v25)
                (scm:cond
                  [(Data.Either.Left? v25) (_4 (v3 (Data.Either.Left-value0 v25)))]
                  [(Data.Either.Right? v25) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) (Data.Either.Right (Data.Either.Right-value0 v25)))]
                  [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "right") (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref Functor01 (scm:string->symbol "map")) Data.Either.Right)])
              (scm:lambda (v25)
                (scm:cond
                  [(Data.Either.Left? v25) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) (Data.Either.Left (Data.Either.Left-value0 v25)))]
                  [(Data.Either.Right? v25) (_4 (v3 (Data.Either.Right-value0 v25)))]
                  [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
            profunctorStar12))))))

  (scm:define categoryStar
    (scm:lambda (dictMonad0)
      (scm:let*
        ([_1 ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined))]
         [semigroupoidStar12 (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (x4)
              (((rt:record-ref _1 (scm:string->symbol "bind")) (v13 x4)) v2))))))])
          (scm:list (scm:cons (scm:string->symbol "identity") (rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure"))) (scm:cons (scm:string->symbol "Semigroupoid0") (scm:lambda (_)
            semigroupoidStar12))))))

  (scm:define applyStar
    (scm:lambda (dictApply0)
      (scm:let*
        ([_1 ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorStar12 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref _1 (scm:string->symbol "map")) f2)])
              (scm:lambda (x5)
                (_4 (v3 x5))))))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v3)
            (scm:lambda (v14)
              (scm:lambda (a5)
                (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (v3 a5)) (v14 a5)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar12))))))

  (scm:define bindStar
    (scm:lambda (dictBind0)
      (scm:let*
        ([_1 ((rt:record-ref dictBind0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [applyStar13 (scm:let ([functorStar13 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f3)
          (scm:lambda (v4)
            (scm:let ([_5 ((rt:record-ref _2 (scm:string->symbol "map")) f3)])
              (scm:lambda (x6)
                (_5 (v4 x6))))))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v4)
            (scm:lambda (v15)
              (scm:lambda (a6)
                (((rt:record-ref _1 (scm:string->symbol "apply")) (v4 a6)) (v15 a6)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar13))))])
          (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v4)
            (scm:lambda (f5)
              (scm:lambda (x6)
                (((rt:record-ref dictBind0 (scm:string->symbol "bind")) (v4 x6)) (scm:lambda (a7)
                  ((f5 a7) x6))))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
            applyStar13))))))

  (scm:define applicativeStar
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [applyStar13 (scm:let ([functorStar13 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f3)
          (scm:lambda (v4)
            (scm:let ([_5 ((rt:record-ref _2 (scm:string->symbol "map")) f3)])
              (scm:lambda (x6)
                (_5 (v4 x6))))))))])
          (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v4)
            (scm:lambda (v15)
              (scm:lambda (a6)
                (((rt:record-ref _1 (scm:string->symbol "apply")) (v4 a6)) (v15 a6)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar13))))])
          (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a4)
            (scm:lambda (v5)
              ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) a4)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
            applyStar13))))))

  (scm:define monadStar
    (scm:lambda (dictMonad0)
      (scm:let*
        ([_1 ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [applicativeStar13 (scm:let*
          ([_3 ((rt:record-ref _2 (scm:string->symbol "Functor0")) (scm:quote undefined))]
           [functorStar14 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
            (scm:lambda (v5)
              (scm:let ([_6 ((rt:record-ref _3 (scm:string->symbol "map")) f4)])
                (scm:lambda (x7)
                  (_6 (v5 x7))))))))]
           [applyStar15 (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v5)
            (scm:lambda (v16)
              (scm:lambda (a7)
                (((rt:record-ref _2 (scm:string->symbol "apply")) (v5 a7)) (v16 a7)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar14)))])
            (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a6)
              (scm:lambda (v7)
                ((rt:record-ref _1 (scm:string->symbol "pure")) a6)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
              applyStar15))))]
         [bindStar14 (bindStar ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined)))])
          (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
            applicativeStar13)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
            bindStar14))))))

  (scm:define altStar
    (scm:lambda (dictAlt0)
      (scm:let*
        ([_1 ((rt:record-ref dictAlt0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorStar12 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:let ([_4 ((rt:record-ref _1 (scm:string->symbol "map")) f2)])
              (scm:lambda (x5)
                (_4 (v3 x5))))))))])
          (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v3)
            (scm:lambda (v14)
              (scm:lambda (a5)
                (((rt:record-ref dictAlt0 (scm:string->symbol "alt")) (v3 a5)) (v14 a5)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar12))))))

  (scm:define plusStar
    (scm:lambda (dictPlus0)
      (scm:let*
        ([empty1 (rt:record-ref dictPlus0 (scm:string->symbol "empty"))]
         [_2 ((rt:record-ref dictPlus0 (scm:string->symbol "Alt0")) (scm:quote undefined))]
         [_3 ((rt:record-ref _2 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [altStar14 (scm:let ([functorStar14 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
          (scm:lambda (v5)
            (scm:let ([_6 ((rt:record-ref _3 (scm:string->symbol "map")) f4)])
              (scm:lambda (x7)
                (_6 (v5 x7))))))))])
          (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v5)
            (scm:lambda (v16)
              (scm:lambda (a7)
                (((rt:record-ref _2 (scm:string->symbol "alt")) (v5 a7)) (v16 a7)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar14))))])
          (scm:list (scm:cons (scm:string->symbol "empty") (scm:lambda (v5)
            empty1)) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
            altStar14))))))

  (scm:define alternativeStar
    (scm:lambda (dictAlternative0)
      (scm:let*
        ([_1 ((rt:record-ref dictAlternative0 (scm:string->symbol "Applicative0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [applicativeStar13 (scm:let*
          ([_3 ((rt:record-ref _2 (scm:string->symbol "Functor0")) (scm:quote undefined))]
           [functorStar14 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
            (scm:lambda (v5)
              (scm:let ([_6 ((rt:record-ref _3 (scm:string->symbol "map")) f4)])
                (scm:lambda (x7)
                  (_6 (v5 x7))))))))]
           [applyStar15 (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v5)
            (scm:lambda (v16)
              (scm:lambda (a7)
                (((rt:record-ref _2 (scm:string->symbol "apply")) (v5 a7)) (v16 a7)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar14)))])
            (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a6)
              (scm:lambda (v7)
                ((rt:record-ref _1 (scm:string->symbol "pure")) a6)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
              applyStar15))))]
         [_4 ((rt:record-ref dictAlternative0 (scm:string->symbol "Plus1")) (scm:quote undefined))]
         [empty5 (rt:record-ref _4 (scm:string->symbol "empty"))]
         [plusStar16 (scm:let*
          ([_6 ((rt:record-ref _4 (scm:string->symbol "Alt0")) (scm:quote undefined))]
           [_7 ((rt:record-ref _6 (scm:string->symbol "Functor0")) (scm:quote undefined))]
           [functorStar18 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f8)
            (scm:lambda (v9)
              (scm:let ([_10 ((rt:record-ref _7 (scm:string->symbol "map")) f8)])
                (scm:lambda (x11)
                  (_10 (v9 x11))))))))]
           [altStar19 (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v9)
            (scm:lambda (v110)
              (scm:lambda (a11)
                (((rt:record-ref _6 (scm:string->symbol "alt")) (v9 a11)) (v110 a11)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorStar18)))])
            (scm:list (scm:cons (scm:string->symbol "empty") (scm:lambda (v10)
              empty5)) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
              altStar19))))])
          (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
            applicativeStar13)) (scm:cons (scm:string->symbol "Plus1") (scm:lambda (_)
            plusStar16))))))

  (scm:define monadPlusStar
    (scm:lambda (dictMonadPlus0)
      (scm:let*
        ([monadStar11 (monadStar ((rt:record-ref dictMonadPlus0 (scm:string->symbol "Monad0")) (scm:quote undefined)))]
         [alternativeStar12 (alternativeStar ((rt:record-ref dictMonadPlus0 (scm:string->symbol "Alternative1")) (scm:quote undefined)))])
          (scm:list (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
            monadStar11)) (scm:cons (scm:string->symbol "Alternative1") (scm:lambda (_)
            alternativeStar12)))))))
