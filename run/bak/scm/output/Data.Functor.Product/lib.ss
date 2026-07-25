#!r6rs
#!chezscheme
(library
  (Data.Functor.Product lib)
  (export
    Product
    applicativeProduct
    applyProduct
    bihoistProduct
    bindProduct
    eq1Product
    eqProduct
    functorProduct
    monadProduct
    newtypeProduct
    ord1Product
    ordProduct
    product
    showProduct)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define Product
    (scm:lambda (x0)
      x0))

  (scm:define showProduct
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(product ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Data.Tuple.Tuple-value0 v2))) (rt:string->pstring " ")) ((rt:record-ref dictShow11 (scm:string->symbol "show")) (Data.Tuple.Tuple-value1 v2))) (rt:string->pstring ")"))))))))

  (scm:define product
    (scm:lambda (fa0)
      (scm:lambda (ga1)
        (Data.Tuple.Tuple* fa0 ga1))))

  (scm:define newtypeProduct
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define functorProduct
    (scm:lambda (dictFunctor0)
      (scm:lambda (dictFunctor11)
        (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (Data.Tuple.Tuple* (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f2) (Data.Tuple.Tuple-value0 v3)) (((rt:record-ref dictFunctor11 (scm:string->symbol "map")) f2) (Data.Tuple.Tuple-value1 v3))))))))))

  (scm:define eq1Product
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq111)
        (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq2)
          (scm:let*
            ([eq123 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq2)]
             [eq134 ((rt:record-ref dictEq111 (scm:string->symbol "eq1")) dictEq2)])
              (scm:lambda (v5)
                (scm:lambda (v16)
                  (scm:and ((eq123 (Data.Tuple.Tuple-value0 v5)) (Data.Tuple.Tuple-value0 v16)) ((eq134 (Data.Tuple.Tuple-value1 v5)) (Data.Tuple.Tuple-value1 v16))))))))))))

  (scm:define eqProduct
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq111)
        (scm:lambda (dictEq2)
          (scm:list (scm:cons (scm:string->symbol "eq") (scm:let*
            ([eq123 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq2)]
             [eq134 ((rt:record-ref dictEq111 (scm:string->symbol "eq1")) dictEq2)])
              (scm:lambda (v5)
                (scm:lambda (v16)
                  (scm:and ((eq123 (Data.Tuple.Tuple-value0 v5)) (Data.Tuple.Tuple-value0 v16)) ((eq134 (Data.Tuple.Tuple-value1 v5)) (Data.Tuple.Tuple-value1 v16))))))))))))

  (scm:define ord1Product
    (scm:lambda (dictOrd10)
      (scm:let ([_1 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
        (scm:lambda (dictOrd112)
          (scm:let*
            ([_3 ((rt:record-ref dictOrd112 (scm:string->symbol "Eq10")) (scm:quote undefined))]
             [eq1Product24 (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq4)
              (scm:let*
                ([eq125 ((rt:record-ref _1 (scm:string->symbol "eq1")) dictEq4)]
                 [eq136 ((rt:record-ref _3 (scm:string->symbol "eq1")) dictEq4)])
                  (scm:lambda (v7)
                    (scm:lambda (v18)
                      (scm:and ((eq125 (Data.Tuple.Tuple-value0 v7)) (Data.Tuple.Tuple-value0 v18)) ((eq136 (Data.Tuple.Tuple-value1 v7)) (Data.Tuple.Tuple-value1 v18)))))))))])
              (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd5)
                (scm:let*
                  ([compare126 ((rt:record-ref dictOrd10 (scm:string->symbol "compare1")) dictOrd5)]
                   [compare137 ((rt:record-ref dictOrd112 (scm:string->symbol "compare1")) dictOrd5)])
                    (scm:lambda (v8)
                      (scm:lambda (v19)
                        (scm:let ([v210 ((compare126 (Data.Tuple.Tuple-value0 v8)) (Data.Tuple.Tuple-value0 v19))])
                          (scm:cond
                            [(Data.Ordering.EQ? v210) ((compare137 (Data.Tuple.Tuple-value1 v8)) (Data.Tuple.Tuple-value1 v19))]
                            [scm:else v210]))))))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
                eq1Product24))))))))

  (scm:define ordProduct
    (scm:lambda (dictOrd10)
      (scm:let*
        ([ord1Product11 (ord1Product dictOrd10)]
         [_2 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
          (scm:lambda (dictOrd113)
            (scm:let ([_4 ((rt:record-ref dictOrd113 (scm:string->symbol "Eq10")) (scm:quote undefined))])
              (scm:lambda (dictOrd5)
                (scm:let*
                  ([_6 ((rt:record-ref dictOrd5 (scm:string->symbol "Eq0")) (scm:quote undefined))]
                   [eqProduct37 (scm:list (scm:cons (scm:string->symbol "eq") (scm:let*
                    ([eq127 ((rt:record-ref _2 (scm:string->symbol "eq1")) _6)]
                     [eq138 ((rt:record-ref _4 (scm:string->symbol "eq1")) _6)])
                      (scm:lambda (v9)
                        (scm:lambda (v110)
                          (scm:and ((eq127 (Data.Tuple.Tuple-value0 v9)) (Data.Tuple.Tuple-value0 v110)) ((eq138 (Data.Tuple.Tuple-value1 v9)) (Data.Tuple.Tuple-value1 v110))))))))])
                    (scm:list (scm:cons (scm:string->symbol "compare") ((rt:record-ref (ord1Product11 dictOrd113) (scm:string->symbol "compare1")) dictOrd5)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                      eqProduct37))))))))))

  (scm:define bihoistProduct
    (scm:lambda (natF0)
      (scm:lambda (natG1)
        (scm:lambda (v2)
          (Data.Tuple.Tuple* (natF0 (Data.Tuple.Tuple-value0 v2)) (natG1 (Data.Tuple.Tuple-value1 v2)))))))

  (scm:define applyProduct
    (scm:lambda (dictApply0)
      (scm:let ([_1 ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (dictApply12)
          (scm:let*
            ([_3 ((rt:record-ref dictApply12 (scm:string->symbol "Functor0")) (scm:quote undefined))]
             [functorProduct24 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
              (scm:lambda (v5)
                (Data.Tuple.Tuple* (((rt:record-ref _1 (scm:string->symbol "map")) f4) (Data.Tuple.Tuple-value0 v5)) (((rt:record-ref _3 (scm:string->symbol "map")) f4) (Data.Tuple.Tuple-value1 v5)))))))])
              (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v5)
                (scm:lambda (v16)
                  (Data.Tuple.Tuple* (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (Data.Tuple.Tuple-value0 v5)) (Data.Tuple.Tuple-value0 v16)) (((rt:record-ref dictApply12 (scm:string->symbol "apply")) (Data.Tuple.Tuple-value1 v5)) (Data.Tuple.Tuple-value1 v16)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                functorProduct24))))))))

  (scm:define bindProduct
    (scm:lambda (dictBind0)
      (scm:let*
        ([_1 ((rt:record-ref dictBind0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))])
          (scm:lambda (dictBind13)
            (scm:let*
              ([_4 ((rt:record-ref dictBind13 (scm:string->symbol "Apply0")) (scm:quote undefined))]
               [_5 ((rt:record-ref _4 (scm:string->symbol "Functor0")) (scm:quote undefined))]
               [applyProduct26 (scm:let ([functorProduct26 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f6)
                (scm:lambda (v7)
                  (Data.Tuple.Tuple* (((rt:record-ref _2 (scm:string->symbol "map")) f6) (Data.Tuple.Tuple-value0 v7)) (((rt:record-ref _5 (scm:string->symbol "map")) f6) (Data.Tuple.Tuple-value1 v7)))))))])
                (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v7)
                  (scm:lambda (v18)
                    (Data.Tuple.Tuple* (((rt:record-ref _1 (scm:string->symbol "apply")) (Data.Tuple.Tuple-value0 v7)) (Data.Tuple.Tuple-value0 v18)) (((rt:record-ref _4 (scm:string->symbol "apply")) (Data.Tuple.Tuple-value1 v7)) (Data.Tuple.Tuple-value1 v18)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorProduct26))))])
                (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v7)
                  (scm:lambda (f8)
                    (Data.Tuple.Tuple* (((rt:record-ref dictBind0 (scm:string->symbol "bind")) (Data.Tuple.Tuple-value0 v7)) (scm:lambda (x9)
                      (Data.Tuple.Tuple-value0 (f8 x9)))) (((rt:record-ref dictBind13 (scm:string->symbol "bind")) (Data.Tuple.Tuple-value1 v7)) (scm:lambda (x9)
                      (Data.Tuple.Tuple-value1 (f8 x9)))))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
                  applyProduct26))))))))

  (scm:define applicativeProduct
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))])
          (scm:lambda (dictApplicative13)
            (scm:let*
              ([_4 ((rt:record-ref dictApplicative13 (scm:string->symbol "Apply0")) (scm:quote undefined))]
               [_5 ((rt:record-ref _4 (scm:string->symbol "Functor0")) (scm:quote undefined))]
               [applyProduct26 (scm:let ([functorProduct26 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f6)
                (scm:lambda (v7)
                  (Data.Tuple.Tuple* (((rt:record-ref _2 (scm:string->symbol "map")) f6) (Data.Tuple.Tuple-value0 v7)) (((rt:record-ref _5 (scm:string->symbol "map")) f6) (Data.Tuple.Tuple-value1 v7)))))))])
                (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v7)
                  (scm:lambda (v18)
                    (Data.Tuple.Tuple* (((rt:record-ref _1 (scm:string->symbol "apply")) (Data.Tuple.Tuple-value0 v7)) (Data.Tuple.Tuple-value0 v18)) (((rt:record-ref _4 (scm:string->symbol "apply")) (Data.Tuple.Tuple-value1 v7)) (Data.Tuple.Tuple-value1 v18)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorProduct26))))])
                (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a7)
                  (Data.Tuple.Tuple* ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) a7) ((rt:record-ref dictApplicative13 (scm:string->symbol "pure")) a7)))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
                  applyProduct26))))))))

  (scm:define monadProduct
    (scm:lambda (dictMonad0)
      (scm:let*
        ([applicativeProduct11 (applicativeProduct ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined)))]
         [bindProduct12 (bindProduct ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined)))])
          (scm:lambda (dictMonad13)
            (scm:let*
              ([applicativeProduct24 (applicativeProduct11 ((rt:record-ref dictMonad13 (scm:string->symbol "Applicative0")) (scm:quote undefined)))]
               [bindProduct25 (bindProduct12 ((rt:record-ref dictMonad13 (scm:string->symbol "Bind1")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
                  applicativeProduct24)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
                  bindProduct25)))))))))
