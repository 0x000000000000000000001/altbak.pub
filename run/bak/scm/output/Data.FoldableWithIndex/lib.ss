#!r6rs
#!chezscheme
(library
  (Data.FoldableWithIndex lib)
  (export
    allWithIndex
    anyWithIndex
    findMapWithIndex
    findWithIndex
    foldMapDefault
    foldMapWithIndex
    foldMapWithIndexDefaultL
    foldMapWithIndexDefaultR
    foldWithIndexM
    foldableWithIndexAdditive
    foldableWithIndexApp
    foldableWithIndexArray
    foldableWithIndexCompose
    foldableWithIndexConj
    foldableWithIndexConst
    foldableWithIndexCoproduct
    foldableWithIndexDisj
    foldableWithIndexDual
    foldableWithIndexEither
    foldableWithIndexFirst
    foldableWithIndexIdentity
    foldableWithIndexLast
    foldableWithIndexMaybe
    foldableWithIndexMultiplicative
    foldableWithIndexProduct
    foldableWithIndexTuple
    foldlDefault
    foldlWithIndex
    foldlWithIndexDefault
    foldrDefault
    foldrWithIndex
    foldrWithIndexDefault
    forWithIndex_
    monoidDual
    monoidEndo
    surroundMapWithIndex
    traverseWithIndex_)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Apply lib) Control.Apply.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Foldable lib) Data.Foldable.)
    (prefix (Data.FunctorWithIndex lib) Data.FunctorWithIndex.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define monoidEndo
    (scm:let ([semigroupEndo10 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (x2)
          (v0 (v11 x2)))))))])
      (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (x1)
        x1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
        semigroupEndo10)))))

  (scm:define monoidDual
    (scm:let*
      ([_0 ((rt:record-ref monoidEndo (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
       [semigroupDual11 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref _0 (scm:string->symbol "append")) v12) v1)))))])
        (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref monoidEndo (scm:string->symbol "mempty"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
          semigroupDual11)))))

  (scm:define foldrWithIndex
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldrWithIndex"))))

  (scm:define traverseWithIndex_
    (scm:lambda (dictApplicative0)
      (scm:let ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))])
        (scm:lambda (dictFoldableWithIndex2)
          (scm:lambda (f3)
            (((rt:record-ref dictFoldableWithIndex2 (scm:string->symbol "foldrWithIndex")) (scm:lambda (i4)
              (scm:let ([_5 (f3 i4)])
                (scm:lambda (x6)
                  (scm:let ([_7 (_5 x6)])
                    (scm:lambda (b8)
                      (((rt:record-ref _1 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (v9)
                        Control.Apply.identity)) _7)) b8))))))) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) Data.Unit.unit)))))))

  (scm:define forWithIndex_
    (scm:lambda (dictApplicative0)
      (scm:let ([traverseWithIndex_11 (traverseWithIndex_ dictApplicative0)])
        (scm:lambda (dictFoldableWithIndex2)
          (scm:let ([_3 (traverseWithIndex_11 dictFoldableWithIndex2)])
            (scm:lambda (b4)
              (scm:lambda (a5)
                ((_3 a5) b4))))))))

  (scm:define foldrDefault
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (f1)
        ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldrWithIndex")) (scm:lambda (v2)
          f1)))))

  (scm:define foldlWithIndex
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldlWithIndex"))))

  (scm:define foldlDefault
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (f1)
        ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (v2)
          f1)))))

  (scm:define foldableWithIndexTuple
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (((f0 Data.Unit.unit) (Data.Tuple.Tuple-value1 v2)) z1))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (((f0 Data.Unit.unit) z1) (Data.Tuple.Tuple-value1 v2)))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          ((f1 Data.Unit.unit) (Data.Tuple.Tuple-value1 v2)))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableTuple))))

  (scm:define foldableWithIndexMultiplicative
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            ((_1 v3) z2)))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (f1 Data.Unit.unit)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableMultiplicative))))

  (scm:define foldableWithIndexMaybe
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (v12)
          (scm:lambda (v23)
            (scm:cond
              [(Data.Maybe.Nothing? v23) v12]
              [(Data.Maybe.Just? v23) ((_1 (Data.Maybe.Just-value0 v23)) v12)]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (v12)
          (scm:lambda (v23)
            (scm:cond
              [(Data.Maybe.Nothing? v23) v12]
              [(Data.Maybe.Just? v23) ((_1 v12) (Data.Maybe.Just-value0 v23))]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (f2)
          (scm:let ([_3 (f2 Data.Unit.unit)])
            (scm:lambda (v14)
              (scm:cond
                [(Data.Maybe.Nothing? v14) mempty1]
                [(Data.Maybe.Just? v14) (_3 (Data.Maybe.Just-value0 v14))]
                [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableMaybe))))

  (scm:define foldableWithIndexLast
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            (scm:cond
              [(Data.Maybe.Nothing? v3) z2]
              [(Data.Maybe.Just? v3) ((_1 (Data.Maybe.Just-value0 v3)) z2)]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            (scm:cond
              [(Data.Maybe.Nothing? v3) z2]
              [(Data.Maybe.Just? v3) ((_1 z2) (Data.Maybe.Just-value0 v3))]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (f2)
          (scm:let ([_3 (f2 Data.Unit.unit)])
            (scm:lambda (v14)
              (scm:cond
                [(Data.Maybe.Nothing? v14) mempty1]
                [(Data.Maybe.Just? v14) (_3 (Data.Maybe.Just-value0 v14))]
                [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableLast))))

  (scm:define foldableWithIndexIdentity
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (((f0 Data.Unit.unit) v2) z1))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (((f0 Data.Unit.unit) z1) v2))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          ((f1 Data.Unit.unit) v2))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableIdentity))))

  (scm:define foldableWithIndexFirst
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            (scm:cond
              [(Data.Maybe.Nothing? v3) z2]
              [(Data.Maybe.Just? v3) ((_1 (Data.Maybe.Just-value0 v3)) z2)]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            (scm:cond
              [(Data.Maybe.Nothing? v3) z2]
              [(Data.Maybe.Just? v3) ((_1 z2) (Data.Maybe.Just-value0 v3))]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (f2)
          (scm:let ([_3 (f2 Data.Unit.unit)])
            (scm:lambda (v14)
              (scm:cond
                [(Data.Maybe.Nothing? v14) mempty1]
                [(Data.Maybe.Just? v14) (_3 (Data.Maybe.Just-value0 v14))]
                [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableFirst))))

  (scm:define foldableWithIndexEither
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Either.Left? v22) v11]
            [(Data.Either.Right? v22) (((v0 Data.Unit.unit) (Data.Either.Right-value0 v22)) v11)]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Either.Left? v22) v11]
            [(Data.Either.Right? v22) (((v0 Data.Unit.unit) v11) (Data.Either.Right-value0 v22))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Either.Left? v13) mempty1]
              [(Data.Either.Right? v13) ((v2 Data.Unit.unit) (Data.Either.Right-value0 v13))]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableEither))))

  (scm:define foldableWithIndexDual
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            ((_1 v3) z2)))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (f1 Data.Unit.unit)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableDual))))

  (scm:define foldableWithIndexDisj
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            ((_1 v3) z2)))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (f1 Data.Unit.unit)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableDisj))))

  (scm:define foldableWithIndexConst
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (v0)
      (scm:lambda (z1)
        (scm:lambda (v12)
          z1)))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (v0)
      (scm:lambda (z1)
        (scm:lambda (v12)
          z1)))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            mempty1))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableConst))))

  (scm:define foldableWithIndexConj
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            ((_1 v3) z2)))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (f1 Data.Unit.unit)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableConj))))

  (scm:define foldableWithIndexAdditive
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:let ([_1 (f0 Data.Unit.unit)])
        (scm:lambda (z2)
          (scm:lambda (v3)
            ((_1 v3) z2)))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (f0 Data.Unit.unit))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (f1 Data.Unit.unit)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableAdditive))))

  (scm:define foldWithIndexM
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (dictMonad1)
        (scm:lambda (f2)
          (scm:lambda (a03)
            (((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (i4)
              (scm:lambda (ma5)
                (scm:lambda (b6)
                  (((rt:record-ref ((rt:record-ref dictMonad1 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) ma5) (scm:let ([_7 (f2 i4)])
                    (scm:lambda (a8)
                      ((_7 a8) b6)))))))) ((rt:record-ref ((rt:record-ref dictMonad1 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) a03)))))))

  (scm:define foldMapWithIndexDefaultR
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (dictMonoid1)
        (scm:let ([mempty2 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))])
          (scm:lambda (f3)
            (((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldrWithIndex")) (scm:lambda (i4)
              (scm:lambda (x5)
                (scm:lambda (acc6)
                  (((rt:record-ref ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) ((f3 i4) x5)) acc6))))) mempty2))))))

  (rt:define-lazy $lazy-foldableWithIndexArray "foldableWithIndexArray" "Data.FoldableWithIndex"
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:let*
          ([_2 ((Data.Foldable.foldrArray (scm:lambda (v2)
            (scm:let*
              ([_3 (Data.Tuple.Tuple-value0 v2)]
               [_4 (Data.Tuple.Tuple-value1 v2)])
                (scm:lambda (y5)
                  (((f0 _3) _4) y5))))) z1)]
           [_3 (Data.FunctorWithIndex.mapWithIndexArray Data.Tuple.Tuple)])
            (scm:lambda (x4)
              (_2 (_3 x4))))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:let*
          ([_2 ((Data.Foldable.foldlArray (scm:lambda (y2)
            (scm:lambda (v3)
              (((f0 (Data.Tuple.Tuple-value0 v3)) y2) (Data.Tuple.Tuple-value1 v3))))) z1)]
           [_3 (Data.FunctorWithIndex.mapWithIndexArray Data.Tuple.Tuple)])
            (scm:lambda (x4)
              (_2 (_3 x4))))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (f2)
          (((rt:record-ref ($lazy-foldableWithIndexArray) (scm:string->symbol "foldrWithIndex")) (scm:lambda (i3)
            (scm:lambda (x4)
              (scm:lambda (acc5)
                (((rt:record-ref ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) ((f2 i3) x4)) acc5))))) mempty1))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableArray))))

  (scm:define foldableWithIndexArray
    ($lazy-foldableWithIndexArray))

  (scm:define foldMapWithIndexDefaultL
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (dictMonoid1)
        (scm:let ([mempty2 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))])
          (scm:lambda (f3)
            (((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (i4)
              (scm:lambda (acc5)
                (scm:lambda (x6)
                  (((rt:record-ref ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) acc5) ((f3 i4) x6)))))) mempty2))))))

  (scm:define foldMapWithIndex
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldMapWithIndex"))))

  (scm:define foldableWithIndexApp
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let*
        ([_1 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "Foldable0")) (scm:quote undefined))]
         [foldableApp2 (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f2)
          (scm:lambda (i3)
            (scm:lambda (v4)
              ((((rt:record-ref _1 (scm:string->symbol "foldr")) f2) i3) v4))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f2)
          (scm:lambda (i3)
            (scm:lambda (v4)
              ((((rt:record-ref _1 (scm:string->symbol "foldl")) f2) i3) v4))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid2)
          ((rt:record-ref _1 (scm:string->symbol "foldMap")) dictMonoid2))))])
          (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f3)
            (scm:lambda (z4)
              (scm:lambda (v5)
                ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldrWithIndex")) f3) z4) v5))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f3)
            (scm:lambda (z4)
              (scm:lambda (v5)
                ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) f3) z4) v5))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid3)
            ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) dictMonoid3))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
            foldableApp2))))))

  (scm:define foldableWithIndexCompose
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let ([_1 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "Foldable0")) (scm:quote undefined))])
        (scm:lambda (dictFoldableWithIndex12)
          (scm:let*
            ([_3 ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "Foldable0")) (scm:quote undefined))]
             [foldableCompose14 (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f4)
              (scm:lambda (i5)
                (scm:lambda (v6)
                  ((((rt:record-ref _1 (scm:string->symbol "foldr")) (scm:let ([_7 ((rt:record-ref _3 (scm:string->symbol "foldr")) f4)])
                    (scm:lambda (b8)
                      (scm:lambda (a9)
                        ((_7 a9) b8))))) i5) v6))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f4)
              (scm:lambda (i5)
                (scm:lambda (v6)
                  ((((rt:record-ref _1 (scm:string->symbol "foldl")) ((rt:record-ref _3 (scm:string->symbol "foldl")) f4)) i5) v6))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid4)
              (scm:let*
                ([foldMap45 ((rt:record-ref _1 (scm:string->symbol "foldMap")) dictMonoid4)]
                 [foldMap56 ((rt:record-ref _3 (scm:string->symbol "foldMap")) dictMonoid4)])
                  (scm:lambda (f7)
                    (scm:lambda (v8)
                      ((foldMap45 (foldMap56 f7)) v8)))))))])
              (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f5)
                (scm:lambda (i6)
                  (scm:lambda (v7)
                    ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldrWithIndex")) (scm:lambda (a8)
                      (scm:let ([_9 ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldrWithIndex")) (scm:lambda (b9)
                        (f5 (Data.Tuple.Tuple* a8 b9))))])
                        (scm:lambda (b10)
                          (scm:lambda (a11)
                            ((_9 a11) b10)))))) i6) v7))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f5)
                (scm:lambda (i6)
                  (scm:lambda (v7)
                    ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (x8)
                      ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldlWithIndex")) (scm:lambda (b9)
                        (f5 (Data.Tuple.Tuple* x8 b9)))))) i6) v7))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid5)
                (scm:let*
                  ([foldMapWithIndex36 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) dictMonoid5)]
                   [foldMapWithIndex47 ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldMapWithIndex")) dictMonoid5)])
                    (scm:lambda (f8)
                      (scm:lambda (v9)
                        ((foldMapWithIndex36 (scm:lambda (x10)
                          (foldMapWithIndex47 (scm:lambda (b11)
                            (f8 (Data.Tuple.Tuple* x10 b11)))))) v9)))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
                foldableCompose14))))))))

  (scm:define foldableWithIndexCoproduct
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let ([foldableCoproduct1 (Data.Foldable.foldableCoproduct ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "Foldable0")) (scm:quote undefined)))])
        (scm:lambda (dictFoldableWithIndex12)
          (scm:let ([foldableCoproduct13 (foldableCoproduct1 ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "Foldable0")) (scm:quote undefined)))])
            (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f4)
              (scm:lambda (z5)
                (scm:let*
                  ([_6 (((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldrWithIndex")) (scm:lambda (x6)
                    (f4 (Data.Either.Left x6)))) z5)]
                   [_7 (((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldrWithIndex")) (scm:lambda (x7)
                    (f4 (Data.Either.Right x7)))) z5)])
                    (scm:lambda (v28)
                      (scm:cond
                        [(Data.Either.Left? v28) (_6 (Data.Either.Left-value0 v28))]
                        [(Data.Either.Right? v28) (_7 (Data.Either.Right-value0 v28))]
                        [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f4)
              (scm:lambda (z5)
                (scm:let*
                  ([_6 (((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (x6)
                    (f4 (Data.Either.Left x6)))) z5)]
                   [_7 (((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldlWithIndex")) (scm:lambda (x7)
                    (f4 (Data.Either.Right x7)))) z5)])
                    (scm:lambda (v28)
                      (scm:cond
                        [(Data.Either.Left? v28) (_6 (Data.Either.Left-value0 v28))]
                        [(Data.Either.Right? v28) (_7 (Data.Either.Right-value0 v28))]
                        [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid4)
              (scm:let*
                ([foldMapWithIndex35 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) dictMonoid4)]
                 [foldMapWithIndex46 ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldMapWithIndex")) dictMonoid4)])
                  (scm:lambda (f7)
                    (scm:let*
                      ([_8 (foldMapWithIndex35 (scm:lambda (x8)
                        (f7 (Data.Either.Left x8))))]
                       [_9 (foldMapWithIndex46 (scm:lambda (x9)
                        (f7 (Data.Either.Right x9))))])
                        (scm:lambda (v210)
                          (scm:cond
                            [(Data.Either.Left? v210) (_8 (Data.Either.Left-value0 v210))]
                            [(Data.Either.Right? v210) (_9 (Data.Either.Right-value0 v210))]
                            [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
              foldableCoproduct13))))))))

  (scm:define foldableWithIndexProduct
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let ([foldableProduct1 (Data.Foldable.foldableProduct ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "Foldable0")) (scm:quote undefined)))])
        (scm:lambda (dictFoldableWithIndex12)
          (scm:let ([foldableProduct13 (foldableProduct1 ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "Foldable0")) (scm:quote undefined)))])
            (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f4)
              (scm:lambda (z5)
                (scm:lambda (v6)
                  ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldrWithIndex")) (scm:lambda (x7)
                    (f4 (Data.Either.Left x7)))) ((((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldrWithIndex")) (scm:lambda (x7)
                    (f4 (Data.Either.Right x7)))) z5) (Data.Tuple.Tuple-value1 v6))) (Data.Tuple.Tuple-value0 v6)))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f4)
              (scm:lambda (z5)
                (scm:lambda (v6)
                  ((((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldlWithIndex")) (scm:lambda (x7)
                    (f4 (Data.Either.Right x7)))) ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (x7)
                    (f4 (Data.Either.Left x7)))) z5) (Data.Tuple.Tuple-value0 v6))) (Data.Tuple.Tuple-value1 v6)))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid4)
              (scm:let*
                ([foldMapWithIndex35 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) dictMonoid4)]
                 [foldMapWithIndex46 ((rt:record-ref dictFoldableWithIndex12 (scm:string->symbol "foldMapWithIndex")) dictMonoid4)])
                  (scm:lambda (f7)
                    (scm:lambda (v8)
                      (((rt:record-ref ((rt:record-ref dictMonoid4 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) ((foldMapWithIndex35 (scm:lambda (x9)
                        (f7 (Data.Either.Left x9)))) (Data.Tuple.Tuple-value0 v8))) ((foldMapWithIndex46 (scm:lambda (x9)
                        (f7 (Data.Either.Right x9)))) (Data.Tuple.Tuple-value1 v8)))))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
              foldableProduct13))))))))

  (scm:define foldlWithIndexDefault
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let ([foldMapWithIndex11 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) monoidDual)])
        (scm:lambda (c2)
          (scm:lambda (u3)
            (scm:lambda (xs4)
              (((foldMapWithIndex11 (scm:lambda (i5)
                (scm:let ([_6 (c2 i5)])
                  (scm:lambda (x7)
                    (scm:lambda (a8)
                      ((_6 a8) x7)))))) xs4) u3)))))))

  (scm:define foldrWithIndexDefault
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let ([foldMapWithIndex11 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) monoidEndo)])
        (scm:lambda (c2)
          (scm:lambda (u3)
            (scm:lambda (xs4)
              (((foldMapWithIndex11 (scm:lambda (i5)
                (c2 i5))) xs4) u3)))))))

  (scm:define surroundMapWithIndex
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let ([foldMapWithIndex11 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) monoidEndo)])
        (scm:lambda (dictSemigroup2)
          (scm:lambda (d3)
            (scm:lambda (t4)
              (scm:lambda (f5)
                (((foldMapWithIndex11 (scm:lambda (i6)
                  (scm:lambda (a7)
                    (scm:lambda (m8)
                      (((rt:record-ref dictSemigroup2 (scm:string->symbol "append")) d3) (((rt:record-ref dictSemigroup2 (scm:string->symbol "append")) ((t4 i6) a7)) m8)))))) f5) d3))))))))

  (scm:define foldMapDefault
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (dictMonoid1)
        (scm:let ([foldMapWithIndex22 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) dictMonoid1)])
          (scm:lambda (f3)
            (foldMapWithIndex22 (scm:lambda (v4)
              f3)))))))

  (scm:define findWithIndex
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (p1)
        (((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (v24)
              (scm:cond
                [(scm:and (Data.Maybe.Nothing? v13) ((p1 v2) v24)) (Data.Maybe.Just (scm:list (scm:cons (scm:string->symbol "index") v2) (scm:cons (scm:string->symbol "value") v24)))]
                [scm:else v13]))))) Data.Maybe.Nothing))))

  (scm:define findMapWithIndex
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (f1)
        (((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (v24)
              (scm:cond
                [(Data.Maybe.Nothing? v13) ((f1 v2) v24)]
                [scm:else v13]))))) Data.Maybe.Nothing))))

  (scm:define anyWithIndex
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (dictHeytingAlgebra1)
        (scm:let ([foldMapWithIndex22 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) (scm:let ([semigroupDisj12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "disj")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "ff"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupDisj12)))))])
          (scm:lambda (t3)
            (foldMapWithIndex22 (scm:lambda (i4)
              (t3 i4))))))))

  (scm:define allWithIndex
    (scm:lambda (dictFoldableWithIndex0)
      (scm:lambda (dictHeytingAlgebra1)
        (scm:let ([foldMapWithIndex22 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) (scm:let ([semigroupConj12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "conj")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "tt"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupConj12)))))])
          (scm:lambda (t3)
            (foldMapWithIndex22 (scm:lambda (i4)
              (t3 i4)))))))))
