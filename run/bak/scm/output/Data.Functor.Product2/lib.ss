#!r6rs
#!chezscheme
(library
  (Data.Functor.Product2 lib)
  (export
    Product2
    Product2*
    Product2-value0
    Product2-value1
    Product2?
    biapplicativeProduct2
    biapplyProduct2
    bifunctorProduct2
    eqProduct2
    functorProduct2
    ordProduct2
    profunctorProduct2
    showProduct2)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ordering lib) Data.Ordering.))

  (scm:define-record-type (Product2$ Product2* Product2?)
    (scm:fields (scm:immutable value0 Product2-value0) (scm:immutable value1 Product2-value1)))

  (scm:define Product2
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Product2* value0 value1))))

  (scm:define showProduct2
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Product2 ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Product2-value0 v2))) (rt:string->pstring " ")) ((rt:record-ref dictShow11 (scm:string->symbol "show")) (Product2-value1 v2))) (rt:string->pstring ")"))))))))

  (scm:define profunctorProduct2
    (scm:lambda (dictProfunctor0)
      (scm:lambda (dictProfunctor11)
        (scm:list (scm:cons (scm:string->symbol "dimap") (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (v4)
              (Product2* ((((rt:record-ref dictProfunctor0 (scm:string->symbol "dimap")) f2) g3) (Product2-value0 v4)) ((((rt:record-ref dictProfunctor11 (scm:string->symbol "dimap")) f2) g3) (Product2-value1 v4)))))))))))

  (scm:define functorProduct2
    (scm:lambda (dictFunctor0)
      (scm:lambda (dictFunctor11)
        (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (v3)
            (Product2* (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f2) (Product2-value0 v3)) (((rt:record-ref dictFunctor11 (scm:string->symbol "map")) f2) (Product2-value1 v3))))))))))

  (scm:define eqProduct2
    (scm:lambda (dictEq0)
      (scm:lambda (dictEq11)
        (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:and (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Product2-value0 x2)) (Product2-value0 y3)) (((rt:record-ref dictEq11 (scm:string->symbol "eq")) (Product2-value1 x2)) (Product2-value1 y3))))))))))

  (scm:define ordProduct2
    (scm:lambda (dictOrd0)
      (scm:let ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))])
        (scm:lambda (dictOrd12)
          (scm:let*
            ([_3 ((rt:record-ref dictOrd12 (scm:string->symbol "Eq0")) (scm:quote undefined))]
             [eqProduct224 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x4)
              (scm:lambda (y5)
                (scm:and (((rt:record-ref _1 (scm:string->symbol "eq")) (Product2-value0 x4)) (Product2-value0 y5)) (((rt:record-ref _3 (scm:string->symbol "eq")) (Product2-value1 x4)) (Product2-value1 y5)))))))])
              (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x5)
                (scm:lambda (y6)
                  (scm:let ([v7 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Product2-value0 x5)) (Product2-value0 y6))])
                    (scm:cond
                      [(Data.Ordering.LT? v7) Data.Ordering.LT]
                      [(Data.Ordering.GT? v7) Data.Ordering.GT]
                      [scm:else (((rt:record-ref dictOrd12 (scm:string->symbol "compare")) (Product2-value1 x5)) (Product2-value1 y6))]))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                eqProduct224))))))))

  (scm:define bifunctorProduct2
    (scm:lambda (dictBifunctor0)
      (scm:lambda (dictBifunctor11)
        (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (v4)
              (Product2* ((((rt:record-ref dictBifunctor0 (scm:string->symbol "bimap")) f2) g3) (Product2-value0 v4)) ((((rt:record-ref dictBifunctor11 (scm:string->symbol "bimap")) f2) g3) (Product2-value1 v4)))))))))))

  (scm:define biapplyProduct2
    (scm:lambda (dictBiapply0)
      (scm:let ([_1 ((rt:record-ref dictBiapply0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))])
        (scm:lambda (dictBiapply12)
          (scm:let*
            ([_3 ((rt:record-ref dictBiapply12 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
             [bifunctorProduct224 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f4)
              (scm:lambda (g5)
                (scm:lambda (v6)
                  (Product2* ((((rt:record-ref _1 (scm:string->symbol "bimap")) f4) g5) (Product2-value0 v6)) ((((rt:record-ref _3 (scm:string->symbol "bimap")) f4) g5) (Product2-value1 v6))))))))])
              (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v5)
                (scm:lambda (v16)
                  (Product2* (((rt:record-ref dictBiapply0 (scm:string->symbol "biapply")) (Product2-value0 v5)) (Product2-value0 v16)) (((rt:record-ref dictBiapply12 (scm:string->symbol "biapply")) (Product2-value1 v5)) (Product2-value1 v16)))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
                bifunctorProduct224))))))))

  (scm:define biapplicativeProduct2
    (scm:lambda (dictBiapplicative0)
      (scm:let*
        ([_1 ((rt:record-ref dictBiapplicative0 (scm:string->symbol "Biapply0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))])
          (scm:lambda (dictBiapplicative13)
            (scm:let*
              ([_4 ((rt:record-ref dictBiapplicative13 (scm:string->symbol "Biapply0")) (scm:quote undefined))]
               [_5 ((rt:record-ref _4 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
               [biapplyProduct226 (scm:let ([bifunctorProduct226 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f6)
                (scm:lambda (g7)
                  (scm:lambda (v8)
                    (Product2* ((((rt:record-ref _2 (scm:string->symbol "bimap")) f6) g7) (Product2-value0 v8)) ((((rt:record-ref _5 (scm:string->symbol "bimap")) f6) g7) (Product2-value1 v8))))))))])
                (scm:list (scm:cons (scm:string->symbol "biapply") (scm:lambda (v7)
                  (scm:lambda (v18)
                    (Product2* (((rt:record-ref _1 (scm:string->symbol "biapply")) (Product2-value0 v7)) (Product2-value0 v18)) (((rt:record-ref _4 (scm:string->symbol "biapply")) (Product2-value1 v7)) (Product2-value1 v18)))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
                  bifunctorProduct226))))])
                (scm:list (scm:cons (scm:string->symbol "bipure") (scm:lambda (a7)
                  (scm:lambda (b8)
                    (Product2* (((rt:record-ref dictBiapplicative0 (scm:string->symbol "bipure")) a7) b8) (((rt:record-ref dictBiapplicative13 (scm:string->symbol "bipure")) a7) b8))))) (scm:cons (scm:string->symbol "Biapply0") (scm:lambda (_)
                  biapplyProduct226)))))))))
