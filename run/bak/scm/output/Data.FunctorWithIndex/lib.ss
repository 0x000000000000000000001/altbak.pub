#!r6rs
#!chezscheme
(library
  (Data.FunctorWithIndex lib)
  (export
    functorWithIndexAdditive
    functorWithIndexApp
    functorWithIndexArray
    functorWithIndexCompose
    functorWithIndexConj
    functorWithIndexConst
    functorWithIndexCoproduct
    functorWithIndexDisj
    functorWithIndexDual
    functorWithIndexEither
    functorWithIndexFirst
    functorWithIndexIdentity
    functorWithIndexLast
    functorWithIndexMaybe
    functorWithIndexMultiplicative
    functorWithIndexProduct
    functorWithIndexTuple
    mapDefault
    mapWithIndex
    mapWithIndexArray)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Const lib) Data.Const.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.Identity lib) Data.Identity.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Monoid.Additive lib) Data.Monoid.Additive.)
    (prefix (Data.Monoid.Conj lib) Data.Monoid.Conj.)
    (prefix (Data.Monoid.Disj lib) Data.Monoid.Disj.)
    (prefix (Data.Monoid.Dual lib) Data.Monoid.Dual.)
    (prefix (Data.Monoid.Multiplicative lib) Data.Monoid.Multiplicative.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.)
    (Data.FunctorWithIndex foreign))

  (scm:define mapWithIndex
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "mapWithIndex"))))

  (scm:define mapDefault
    (scm:lambda (dictFunctorWithIndex0)
      (scm:lambda (f1)
        ((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "mapWithIndex")) (scm:lambda (v2)
          f1)))))

  (scm:define functorWithIndexTuple
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (m2)
          (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 m2) (_1 (Data.Tuple.Tuple-value1 m2))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Tuple.functorTuple))))

  (scm:define functorWithIndexProduct
    (scm:lambda (dictFunctorWithIndex0)
      (scm:let ([_1 ((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (dictFunctorWithIndex12)
          (scm:let*
            ([_3 ((rt:record-ref dictFunctorWithIndex12 (scm:string->symbol "Functor0")) (scm:quote undefined))]
             [functorProduct14 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
              (scm:lambda (v5)
                (Data.Tuple.Tuple* (((rt:record-ref _1 (scm:string->symbol "map")) f4) (Data.Tuple.Tuple-value0 v5)) (((rt:record-ref _3 (scm:string->symbol "map")) f4) (Data.Tuple.Tuple-value1 v5)))))))])
              (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f5)
                (scm:lambda (v6)
                  (Data.Tuple.Tuple* (((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "mapWithIndex")) (scm:lambda (x7)
                    (f5 (Data.Either.Left x7)))) (Data.Tuple.Tuple-value0 v6)) (((rt:record-ref dictFunctorWithIndex12 (scm:string->symbol "mapWithIndex")) (scm:lambda (x7)
                    (f5 (Data.Either.Right x7)))) (Data.Tuple.Tuple-value1 v6)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                functorProduct14))))))))

  (scm:define functorWithIndexMultiplicative
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Multiplicative.functorMultiplicative))))

  (scm:define functorWithIndexMaybe
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (v12)
          (scm:cond
            [(Data.Maybe.Just? v12) (Data.Maybe.Just (_1 (Data.Maybe.Just-value0 v12)))]
            [scm:else Data.Maybe.Nothing]))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe))))

  (scm:define functorWithIndexLast
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (v12)
          (scm:cond
            [(Data.Maybe.Just? v12) (Data.Maybe.Just (_1 (Data.Maybe.Just-value0 v12)))]
            [scm:else Data.Maybe.Nothing]))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe))))

  (scm:define functorWithIndexIdentity
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (scm:lambda (v1)
        ((f0 Data.Unit.unit) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Identity.functorIdentity))))

  (scm:define functorWithIndexFirst
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (v12)
          (scm:cond
            [(Data.Maybe.Just? v12) (Data.Maybe.Just (_1 (Data.Maybe.Just-value0 v12)))]
            [scm:else Data.Maybe.Nothing]))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe))))

  (scm:define functorWithIndexEither
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (m2)
          (scm:cond
            [(Data.Either.Left? m2) (Data.Either.Left (Data.Either.Left-value0 m2))]
            [(Data.Either.Right? m2) (Data.Either.Right (_1 (Data.Either.Right-value0 m2)))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Either.functorEither))))

  (scm:define functorWithIndexDual
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Dual.functorDual))))

  (scm:define functorWithIndexDisj
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Disj.functorDisj))))

  (scm:define functorWithIndexCoproduct
    (scm:lambda (dictFunctorWithIndex0)
      (scm:let ([_1 ((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (dictFunctorWithIndex12)
          (scm:let*
            ([_3 ((rt:record-ref dictFunctorWithIndex12 (scm:string->symbol "Functor0")) (scm:quote undefined))]
             [functorCoproduct14 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
              (scm:lambda (v5)
                (scm:let*
                  ([_6 ((rt:record-ref _1 (scm:string->symbol "map")) f4)]
                   [_7 ((rt:record-ref _3 (scm:string->symbol "map")) f4)])
                    (scm:cond
                      [(Data.Either.Left? v5) (Data.Either.Left (_6 (Data.Either.Left-value0 v5)))]
                      [(Data.Either.Right? v5) (Data.Either.Right (_7 (Data.Either.Right-value0 v5)))]
                      [scm:else (rt:fail)]))))))])
              (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f5)
                (scm:lambda (v6)
                  (scm:let*
                    ([_7 ((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "mapWithIndex")) (scm:lambda (x7)
                      (f5 (Data.Either.Left x7))))]
                     [_8 ((rt:record-ref dictFunctorWithIndex12 (scm:string->symbol "mapWithIndex")) (scm:lambda (x8)
                      (f5 (Data.Either.Right x8))))])
                      (scm:cond
                        [(Data.Either.Left? v6) (Data.Either.Left (_7 (Data.Either.Left-value0 v6)))]
                        [(Data.Either.Right? v6) (Data.Either.Right (_8 (Data.Either.Right-value0 v6)))]
                        [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                functorCoproduct14))))))))

  (scm:define functorWithIndexConst
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (v0)
      (scm:lambda (v11)
        v11))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Const.functorConst))))

  (scm:define functorWithIndexConj
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Conj.functorConj))))

  (scm:define functorWithIndexCompose
    (scm:lambda (dictFunctorWithIndex0)
      (scm:let ([_1 ((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (dictFunctorWithIndex12)
          (scm:let*
            ([_3 ((rt:record-ref dictFunctorWithIndex12 (scm:string->symbol "Functor0")) (scm:quote undefined))]
             [functorCompose14 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f4)
              (scm:lambda (v5)
                (((rt:record-ref _1 (scm:string->symbol "map")) ((rt:record-ref _3 (scm:string->symbol "map")) f4)) v5)))))])
              (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f5)
                (scm:lambda (v6)
                  (((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "mapWithIndex")) (scm:lambda (x7)
                    ((rt:record-ref dictFunctorWithIndex12 (scm:string->symbol "mapWithIndex")) (scm:lambda (b8)
                      (f5 (Data.Tuple.Tuple* x7 b8)))))) v6)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                functorCompose14))))))))

  (scm:define functorWithIndexArray
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") mapWithIndexArray) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorArray))))

  (scm:define functorWithIndexApp
    (scm:lambda (dictFunctorWithIndex0)
      (scm:let ([_1 ((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f2)
          (scm:lambda (v3)
            (((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "mapWithIndex")) f2) v3)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
          _1))))))

  (scm:define functorWithIndexAdditive
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Additive.functorAdditive)))))
