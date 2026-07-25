#!r6rs
#!chezscheme
(library
  (Data.Functor.Coproduct lib)
  (export
    Coproduct
    bihoistCoproduct
    comonadCoproduct
    coproduct
    eq1Coproduct
    eqCoproduct
    extendCoproduct
    functorCoproduct
    left
    newtypeCoproduct
    ord1Coproduct
    ordCoproduct
    right
    showCoproduct)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define Coproduct
    (scm:lambda (x0)
      x0))

  (scm:define showCoproduct
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (scm:cond
            [(Data.Either.Left? v2) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(left ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Data.Either.Left-value0 v2))) (rt:string->pstring ")"))]
            [(Data.Either.Right? v2) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(right ") ((rt:record-ref dictShow11 (scm:string->symbol "show")) (Data.Either.Right-value0 v2))) (rt:string->pstring ")"))]
            [scm:else (rt:fail)])))))))

  (scm:define right
    (scm:lambda (ga0)
      (Data.Either.Right ga0)))

  (scm:define newtypeCoproduct
    (scm:list (scm:cons (scm:string->symbol "Coercible0") (scm:lambda (_)
      (scm:quote undefined)))))

  (scm:define left
    (scm:lambda (fa0)
      (Data.Either.Left fa0)))

  (scm:define functorCoproduct
    (scm:lambda (dictFunctor0)
      (scm:lambda (dictFunctor11)
        (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:let*
              ([_4 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f2)]
               [_5 ((rt:record-ref dictFunctor11 (scm:string->symbol "map")) f2)])
                (scm:cond
                  [(Data.Either.Left? v3) (Data.Either.Left (_4 (Data.Either.Left-value0 v3)))]
                  [(Data.Either.Right? v3) (Data.Either.Right (_5 (Data.Either.Right-value0 v3)))]
                  [scm:else (rt:fail)])))))))))

  (scm:define eq1Coproduct
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq111)
        (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq2)
          (scm:let*
            ([eq123 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq2)]
             [eq134 ((rt:record-ref dictEq111 (scm:string->symbol "eq1")) dictEq2)])
              (scm:lambda (v5)
                (scm:lambda (v16)
                  (scm:cond
                    [(Data.Either.Left? v5) (scm:and (Data.Either.Left? v16) ((eq123 (Data.Either.Left-value0 v5)) (Data.Either.Left-value0 v16)))]
                    [scm:else (scm:and (Data.Either.Right? v5) (scm:and (Data.Either.Right? v16) ((eq134 (Data.Either.Right-value0 v5)) (Data.Either.Right-value0 v16))))]))))))))))

  (scm:define eqCoproduct
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq111)
        (scm:lambda (dictEq2)
          (scm:list (scm:cons (scm:string->symbol "eq") (scm:let*
            ([eq123 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq2)]
             [eq134 ((rt:record-ref dictEq111 (scm:string->symbol "eq1")) dictEq2)])
              (scm:lambda (v5)
                (scm:lambda (v16)
                  (scm:cond
                    [(Data.Either.Left? v5) (scm:and (Data.Either.Left? v16) ((eq123 (Data.Either.Left-value0 v5)) (Data.Either.Left-value0 v16)))]
                    [scm:else (scm:and (Data.Either.Right? v5) (scm:and (Data.Either.Right? v16) ((eq134 (Data.Either.Right-value0 v5)) (Data.Either.Right-value0 v16))))]))))))))))

  (scm:define ord1Coproduct
    (scm:lambda (dictOrd10)
      (scm:let ([_1 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
        (scm:lambda (dictOrd112)
          (scm:let*
            ([_3 ((rt:record-ref dictOrd112 (scm:string->symbol "Eq10")) (scm:quote undefined))]
             [eq1Coproduct24 (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq4)
              (scm:let*
                ([eq125 ((rt:record-ref _1 (scm:string->symbol "eq1")) dictEq4)]
                 [eq136 ((rt:record-ref _3 (scm:string->symbol "eq1")) dictEq4)])
                  (scm:lambda (v7)
                    (scm:lambda (v18)
                      (scm:cond
                        [(Data.Either.Left? v7) (scm:and (Data.Either.Left? v18) ((eq125 (Data.Either.Left-value0 v7)) (Data.Either.Left-value0 v18)))]
                        [scm:else (scm:and (Data.Either.Right? v7) (scm:and (Data.Either.Right? v18) ((eq136 (Data.Either.Right-value0 v7)) (Data.Either.Right-value0 v18))))])))))))])
              (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd5)
                (scm:let*
                  ([compare126 ((rt:record-ref dictOrd10 (scm:string->symbol "compare1")) dictOrd5)]
                   [compare137 ((rt:record-ref dictOrd112 (scm:string->symbol "compare1")) dictOrd5)])
                    (scm:lambda (v8)
                      (scm:lambda (v19)
                        (scm:cond
                          [(Data.Either.Left? v8) (scm:cond
                            [(Data.Either.Left? v19) ((compare126 (Data.Either.Left-value0 v8)) (Data.Either.Left-value0 v19))]
                            [scm:else Data.Ordering.LT])]
                          [(Data.Either.Left? v19) Data.Ordering.GT]
                          [(scm:and (Data.Either.Right? v8) (Data.Either.Right? v19)) ((compare137 (Data.Either.Right-value0 v8)) (Data.Either.Right-value0 v19))]
                          [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
                eq1Coproduct24))))))))

  (scm:define ordCoproduct
    (scm:lambda (dictOrd10)
      (scm:let*
        ([ord1Coproduct11 (ord1Coproduct dictOrd10)]
         [_2 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
          (scm:lambda (dictOrd113)
            (scm:let ([_4 ((rt:record-ref dictOrd113 (scm:string->symbol "Eq10")) (scm:quote undefined))])
              (scm:lambda (dictOrd5)
                (scm:let*
                  ([_6 ((rt:record-ref dictOrd5 (scm:string->symbol "Eq0")) (scm:quote undefined))]
                   [eqCoproduct37 (scm:list (scm:cons (scm:string->symbol "eq") (scm:let*
                    ([eq127 ((rt:record-ref _2 (scm:string->symbol "eq1")) _6)]
                     [eq138 ((rt:record-ref _4 (scm:string->symbol "eq1")) _6)])
                      (scm:lambda (v9)
                        (scm:lambda (v110)
                          (scm:cond
                            [(Data.Either.Left? v9) (scm:and (Data.Either.Left? v110) ((eq127 (Data.Either.Left-value0 v9)) (Data.Either.Left-value0 v110)))]
                            [scm:else (scm:and (Data.Either.Right? v9) (scm:and (Data.Either.Right? v110) ((eq138 (Data.Either.Right-value0 v9)) (Data.Either.Right-value0 v110))))]))))))])
                    (scm:list (scm:cons (scm:string->symbol "compare") ((rt:record-ref (ord1Coproduct11 dictOrd113) (scm:string->symbol "compare1")) dictOrd5)) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                      eqCoproduct37))))))))))

  (scm:define coproduct
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Either.Left? v22) (v0 (Data.Either.Left-value0 v22))]
            [(Data.Either.Right? v22) (v11 (Data.Either.Right-value0 v22))]
            [scm:else (rt:fail)])))))

  (scm:define extendCoproduct
    (scm:lambda (dictExtend0)
      (scm:let ([_1 ((rt:record-ref dictExtend0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (dictExtend12)
          (scm:let*
            ([_3 ((rt:record-ref dictExtend12 (scm:string->symbol "Functor0")) (scm:quote undefined))]
             [functorCoproduct24 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
              (scm:lambda (v5)
                (scm:let*
                  ([_6 ((rt:record-ref _1 (scm:string->symbol "map")) f4)]
                   [_7 ((rt:record-ref _3 (scm:string->symbol "map")) f4)])
                    (scm:cond
                      [(Data.Either.Left? v5) (Data.Either.Left (_6 (Data.Either.Left-value0 v5)))]
                      [(Data.Either.Right? v5) (Data.Either.Right (_7 (Data.Either.Right-value0 v5)))]
                      [scm:else (rt:fail)]))))))])
              (scm:list (scm:cons (scm:string->symbol "extend") (scm:lambda (f5)
                (scm:let*
                  ([_6 ((rt:record-ref dictExtend0 (scm:string->symbol "extend")) (scm:lambda (x6)
                    (f5 (Data.Either.Left x6))))]
                   [_7 ((rt:record-ref dictExtend12 (scm:string->symbol "extend")) (scm:lambda (x7)
                    (f5 (Data.Either.Right x7))))])
                    (scm:lambda (v28)
                      (scm:cond
                        [(Data.Either.Left? v28) (Data.Either.Left (_6 (Data.Either.Left-value0 v28)))]
                        [(Data.Either.Right? v28) (Data.Either.Right (_7 (Data.Either.Right-value0 v28)))]
                        [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                functorCoproduct24))))))))

  (scm:define comonadCoproduct
    (scm:lambda (dictComonad0)
      (scm:let ([extendCoproduct11 (extendCoproduct ((rt:record-ref dictComonad0 (scm:string->symbol "Extend0")) (scm:quote undefined)))])
        (scm:lambda (dictComonad12)
          (scm:let ([extendCoproduct23 (extendCoproduct11 ((rt:record-ref dictComonad12 (scm:string->symbol "Extend0")) (scm:quote undefined)))])
            (scm:list (scm:cons (scm:string->symbol "extract") (scm:lambda (v24)
              (scm:cond
                [(Data.Either.Left? v24) ((rt:record-ref dictComonad0 (scm:string->symbol "extract")) (Data.Either.Left-value0 v24))]
                [(Data.Either.Right? v24) ((rt:record-ref dictComonad12 (scm:string->symbol "extract")) (Data.Either.Right-value0 v24))]
                [scm:else (rt:fail)]))) (scm:cons (scm:string->symbol "Extend0") (scm:lambda (_)
              extendCoproduct23))))))))

  (scm:define bihoistCoproduct
    (scm:lambda (natF0)
      (scm:lambda (natG1)
        (scm:lambda (v2)
          (scm:cond
            [(Data.Either.Left? v2) (Data.Either.Left (natF0 (Data.Either.Left-value0 v2)))]
            [(Data.Either.Right? v2) (Data.Either.Right (natG1 (Data.Either.Right-value0 v2)))]
            [scm:else (rt:fail)]))))))
