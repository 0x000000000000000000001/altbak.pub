#!r6rs
#!chezscheme
(library
  (Data.Functor.Compose lib)
  (export
    Compose
    altCompose
    alternativeCompose
    applicativeCompose
    applyCompose
    bihoistCompose
    eq1Compose
    eqCompose
    functorCompose
    newtypeCompose
    ord1Compose
    ordCompose
    plusCompose
    showCompose)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define Compose
    (scm:lambda (x0)
      x0))

  (scm:define showCompose
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Compose ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define newtypeCompose
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define functorCompose
    (scm:lambda (dictFunctor0)
      (scm:lambda (dictFunctor11)
        (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) ((rt:record-ref dictFunctor11 (scm:string->symbol "map")) f2)) v3))))))))

  (scm:define eqCompose
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq111)
        (scm:lambda (dictEq2)
          (scm:let ([eq113 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) (scm:let ([eq113 ((rt:record-ref dictEq111 (scm:string->symbol "eq1")) dictEq2)])
            (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x4)
              (scm:lambda (y5)
                ((eq113 x4) y5)))))))])
            (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (v4)
              (scm:lambda (v15)
                ((eq113 v4) v15))))))))))

  (scm:define ordCompose
    (scm:lambda (dictOrd10)
      (scm:let ([_1 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
        (scm:lambda (dictOrd112)
          (scm:let*
            ([_3 ((rt:record-ref dictOrd112 (scm:string->symbol "Eq10")) (scm:quote undefined))]
             [_4 ((rt:record-ref dictOrd112 (scm:string->symbol "Eq10")) (scm:quote undefined))])
              (scm:lambda (dictOrd5)
                (scm:let*
                  ([compare116 ((rt:record-ref dictOrd10 (scm:string->symbol "compare1")) (scm:let*
                    ([compare116 ((rt:record-ref dictOrd112 (scm:string->symbol "compare1")) dictOrd5)]
                     [eq117 ((rt:record-ref _3 (scm:string->symbol "eq1")) ((rt:record-ref dictOrd5 (scm:string->symbol "Eq0")) (scm:quote undefined)))]
                     [eqApp28 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x8)
                      (scm:lambda (y9)
                        ((eq117 x8) y9)))))])
                      (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x9)
                        (scm:lambda (y10)
                          ((compare116 x9) y10)))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                        eqApp28)))))]
                   [eq117 ((rt:record-ref _1 (scm:string->symbol "eq1")) (scm:let ([eq117 ((rt:record-ref _4 (scm:string->symbol "eq1")) ((rt:record-ref dictOrd5 (scm:string->symbol "Eq0")) (scm:quote undefined)))])
                    (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x8)
                      (scm:lambda (y9)
                        ((eq117 x8) y9)))))))]
                   [eqCompose38 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (v8)
                    (scm:lambda (v19)
                      ((eq117 v8) v19)))))])
                    (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (v9)
                      (scm:lambda (v110)
                        ((compare116 v9) v110)))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                      eqCompose38))))))))))

  (scm:define eq1Compose
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq111)
        (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq2)
          ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) (scm:let ([eq113 ((rt:record-ref dictEq111 (scm:string->symbol "eq1")) dictEq2)])
            (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x4)
              (scm:lambda (y5)
                ((eq113 x4) y5)))))))))))))

  (scm:define ord1Compose
    (scm:lambda (dictOrd10)
      (scm:let*
        ([ordCompose11 (ordCompose dictOrd10)]
         [_2 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
          (scm:lambda (dictOrd113)
            (scm:let*
              ([ordCompose24 (ordCompose11 dictOrd113)]
               [_5 ((rt:record-ref dictOrd113 (scm:string->symbol "Eq10")) (scm:quote undefined))]
               [eq1Compose26 (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq6)
                ((rt:record-ref _2 (scm:string->symbol "eq1")) (scm:let ([eq117 ((rt:record-ref _5 (scm:string->symbol "eq1")) dictEq6)])
                  (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x8)
                    (scm:lambda (y9)
                      ((eq117 x8) y9))))))))))])
                (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd7)
                  (rt:record-ref (ordCompose24 dictOrd7) (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
                  eq1Compose26))))))))

  (scm:define bihoistCompose
    (scm:lambda (dictFunctor0)
      (scm:lambda (natF1)
        (scm:lambda (natG2)
          (scm:lambda (v3)
            (natF1 (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) natG2) v3)))))))

  (scm:define applyCompose
    (scm:lambda (dictApply0)
      (scm:let ([Functor01 ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (dictApply12)
          (scm:let*
            ([apply13 (rt:record-ref dictApply12 (scm:string->symbol "apply"))]
             [_4 ((rt:record-ref dictApply12 (scm:string->symbol "Functor0")) (scm:quote undefined))]
             [functorCompose25 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f5)
              (scm:lambda (v6)
                (((rt:record-ref Functor01 (scm:string->symbol "map")) ((rt:record-ref _4 (scm:string->symbol "map")) f5)) v6)))))])
              (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v6)
                (scm:lambda (v17)
                  (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref Functor01 (scm:string->symbol "map")) apply13) v6)) v17)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                functorCompose25))))))))

  (scm:define applicativeCompose
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [Functor02 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))])
          (scm:lambda (dictApplicative13)
            (scm:let*
              ([_4 ((rt:record-ref dictApplicative13 (scm:string->symbol "Apply0")) (scm:quote undefined))]
               [apply15 (rt:record-ref _4 (scm:string->symbol "apply"))]
               [applyCompose26 (scm:let*
                ([_6 ((rt:record-ref _4 (scm:string->symbol "Functor0")) (scm:quote undefined))]
                 [functorCompose27 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f7)
                  (scm:lambda (v8)
                    (((rt:record-ref Functor02 (scm:string->symbol "map")) ((rt:record-ref _6 (scm:string->symbol "map")) f7)) v8)))))])
                  (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v8)
                    (scm:lambda (v19)
                      (((rt:record-ref _1 (scm:string->symbol "apply")) (((rt:record-ref Functor02 (scm:string->symbol "map")) apply15) v8)) v19)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                    functorCompose27))))])
                (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (x7)
                  ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) ((rt:record-ref dictApplicative13 (scm:string->symbol "pure")) x7)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
                  applyCompose26))))))))

  (scm:define altCompose
    (scm:lambda (dictAlt0)
      (scm:let ([_1 ((rt:record-ref dictAlt0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (dictFunctor2)
          (scm:let ([functorCompose23 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f3)
            (scm:lambda (v4)
              (((rt:record-ref _1 (scm:string->symbol "map")) ((rt:record-ref dictFunctor2 (scm:string->symbol "map")) f3)) v4)))))])
            (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v4)
              (scm:lambda (v15)
                (((rt:record-ref dictAlt0 (scm:string->symbol "alt")) v4) v15)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
              functorCompose23))))))))

  (scm:define plusCompose
    (scm:lambda (dictPlus0)
      (scm:let*
        ([empty1 (rt:record-ref dictPlus0 (scm:string->symbol "empty"))]
         [_2 ((rt:record-ref dictPlus0 (scm:string->symbol "Alt0")) (scm:quote undefined))]
         [_3 ((rt:record-ref _2 (scm:string->symbol "Functor0")) (scm:quote undefined))])
          (scm:lambda (dictFunctor4)
            (scm:let*
              ([functorCompose25 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f5)
                (scm:lambda (v6)
                  (((rt:record-ref _3 (scm:string->symbol "map")) ((rt:record-ref dictFunctor4 (scm:string->symbol "map")) f5)) v6)))))]
               [altCompose26 (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v6)
                (scm:lambda (v17)
                  (((rt:record-ref _2 (scm:string->symbol "alt")) v6) v17)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                functorCompose25)))])
                (scm:list (scm:cons (scm:string->symbol "empty") empty1) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
                  altCompose26))))))))

  (scm:define alternativeCompose
    (scm:lambda (dictAlternative0)
      (scm:let*
        ([applicativeCompose11 (applicativeCompose ((rt:record-ref dictAlternative0 (scm:string->symbol "Applicative0")) (scm:quote undefined)))]
         [_2 ((rt:record-ref dictAlternative0 (scm:string->symbol "Plus1")) (scm:quote undefined))]
         [empty3 (rt:record-ref _2 (scm:string->symbol "empty"))]
         [plusCompose14 (scm:let*
          ([_4 ((rt:record-ref _2 (scm:string->symbol "Alt0")) (scm:quote undefined))]
           [_5 ((rt:record-ref _4 (scm:string->symbol "Functor0")) (scm:quote undefined))])
            (scm:lambda (dictFunctor6)
              (scm:let*
                ([functorCompose27 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f7)
                  (scm:lambda (v8)
                    (((rt:record-ref _5 (scm:string->symbol "map")) ((rt:record-ref dictFunctor6 (scm:string->symbol "map")) f7)) v8)))))]
                 [altCompose28 (scm:list (scm:cons (scm:string->symbol "alt") (scm:lambda (v8)
                  (scm:lambda (v19)
                    (((rt:record-ref _4 (scm:string->symbol "alt")) v8) v19)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorCompose27)))])
                  (scm:list (scm:cons (scm:string->symbol "empty") empty3) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
                    altCompose28))))))])
          (scm:lambda (dictApplicative5)
            (scm:let*
              ([applicativeCompose26 (applicativeCompose11 dictApplicative5)]
               [plusCompose27 (plusCompose14 ((rt:record-ref ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
                  applicativeCompose26)) (scm:cons (scm:string->symbol "Plus1") (scm:lambda (_)
                  plusCompose27)))))))))
