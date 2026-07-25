#!r6rs
#!chezscheme
(library
  (Data.TraversableWithIndex lib)
  (export
    forWithIndex
    mapAccumLWithIndex
    mapAccumRWithIndex
    scanlWithIndex
    scanrWithIndex
    traversableWithIndexAdditive
    traversableWithIndexApp
    traversableWithIndexArray
    traversableWithIndexCompose
    traversableWithIndexConj
    traversableWithIndexConst
    traversableWithIndexCoproduct
    traversableWithIndexDisj
    traversableWithIndexDual
    traversableWithIndexEither
    traversableWithIndexFirst
    traversableWithIndexIdentity
    traversableWithIndexLast
    traversableWithIndexMaybe
    traversableWithIndexMultiplicative
    traversableWithIndexProduct
    traversableWithIndexTuple
    traverseDefault
    traverseWithIndex
    traverseWithIndexDefault)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.FoldableWithIndex lib) Data.FoldableWithIndex.)
    (prefix (Data.Functor.App lib) Data.Functor.App.)
    (prefix (Data.Functor.Compose lib) Data.Functor.Compose.)
    (prefix (Data.Functor.Product lib) Data.Functor.Product.)
    (prefix (Data.FunctorWithIndex lib) Data.FunctorWithIndex.)
    (prefix (Data.Identity lib) Data.Identity.)
    (prefix (Data.Traversable lib) Data.Traversable.)
    (prefix (Data.Traversable.Accum.Internal lib) Data.Traversable.Accum.Internal.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define traverseWithIndexDefault
    (scm:lambda (dictTraversableWithIndex0)
      (scm:lambda (dictApplicative1)
        (scm:let ([sequence12 ((rt:record-ref ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "Traversable2")) (scm:quote undefined)) (scm:string->symbol "sequence")) dictApplicative1)])
          (scm:lambda (f3)
            (scm:let ([_4 ((rt:record-ref ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined)) (scm:string->symbol "mapWithIndex")) f3)])
              (scm:lambda (x5)
                (sequence12 (_4 x5)))))))))

  (scm:define traverseWithIndex
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "traverseWithIndex"))))

  (scm:define traverseDefault
    (scm:lambda (dictTraversableWithIndex0)
      (scm:lambda (dictApplicative1)
        (scm:let ([traverseWithIndex22 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) dictApplicative1)])
          (scm:lambda (f3)
            (traverseWithIndex22 (scm:lambda (v4)
              f3)))))))

  (scm:define traversableWithIndexTuple
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (Data.Tuple.Tuple (Data.Tuple.Tuple-value0 v2))) ((f1 Data.Unit.unit) (Data.Tuple.Tuple-value1 v2))))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexTuple)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexTuple)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableTuple))))

  (scm:define traversableWithIndexProduct
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [foldableWithIndexProduct3 (Data.FoldableWithIndex.foldableWithIndexProduct ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined)))]
         [traversableProduct4 (Data.Traversable.traversableProduct ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
          (scm:lambda (dictTraversableWithIndex15)
            (scm:let*
              ([_6 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined))]
               [_7 ((rt:record-ref _6 (scm:string->symbol "Functor0")) (scm:quote undefined))]
               [functorWithIndexProduct18 (scm:let ([functorProduct18 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f8)
                (scm:lambda (v9)
                  (Data.Tuple.Tuple* (((rt:record-ref _2 (scm:string->symbol "map")) f8) (Data.Tuple.Tuple-value0 v9)) (((rt:record-ref _7 (scm:string->symbol "map")) f8) (Data.Tuple.Tuple-value1 v9)))))))])
                (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f9)
                  (scm:lambda (v10)
                    (Data.Tuple.Tuple* (((rt:record-ref _1 (scm:string->symbol "mapWithIndex")) (scm:lambda (x11)
                      (f9 (Data.Either.Left x11)))) (Data.Tuple.Tuple-value0 v10)) (((rt:record-ref _6 (scm:string->symbol "mapWithIndex")) (scm:lambda (x11)
                      (f9 (Data.Either.Right x11)))) (Data.Tuple.Tuple-value1 v10)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorProduct18))))]
               [foldableWithIndexProduct19 (foldableWithIndexProduct3 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined)))]
               [traversableProduct110 (traversableProduct4 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative11)
                  (scm:let*
                    ([_12 ((rt:record-ref dictApplicative11 (scm:string->symbol "Apply0")) (scm:quote undefined))]
                     [traverseWithIndex313 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) dictApplicative11)]
                     [traverseWithIndex414 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "traverseWithIndex")) dictApplicative11)])
                      (scm:lambda (f15)
                        (scm:lambda (v16)
                          (((rt:record-ref _12 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref _12 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Product.product) ((traverseWithIndex313 (scm:lambda (x17)
                            (f15 (Data.Either.Left x17)))) (Data.Tuple.Tuple-value0 v16)))) ((traverseWithIndex414 (scm:lambda (x17)
                            (f15 (Data.Either.Right x17)))) (Data.Tuple.Tuple-value1 v16)))))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
                  functorWithIndexProduct18)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
                  foldableWithIndexProduct19)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
                  traversableProduct110))))))))

  (scm:define traversableWithIndexMultiplicative
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableMultiplicative (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexMultiplicative)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexMultiplicative)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableMultiplicative))))

  (scm:define traversableWithIndexMaybe
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableMaybe (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexMaybe)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexMaybe)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableMaybe))))

  (scm:define traversableWithIndexLast
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableLast (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexLast)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexLast)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableLast))))

  (scm:define traversableWithIndexIdentity
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Identity.Identity) ((f1 Data.Unit.unit) v2)))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexIdentity)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexIdentity)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableIdentity))))

  (scm:define traversableWithIndexFirst
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableFirst (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexFirst)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexFirst)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableFirst))))

  (scm:define traversableWithIndexEither
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(Data.Either.Left? v12) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) (Data.Either.Left (Data.Either.Left-value0 v12)))]
            [(Data.Either.Right? v12) (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Either.Right) ((v1 Data.Unit.unit) (Data.Either.Right-value0 v12)))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexEither)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexEither)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableEither))))

  (scm:define traversableWithIndexDual
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableDual (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexDual)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexDual)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableDual))))

  (scm:define traversableWithIndexDisj
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableDisj (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexDisj)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexDisj)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableDisj))))

  (scm:define traversableWithIndexCoproduct
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let*
        ([functorWithIndexCoproduct1 (Data.FunctorWithIndex.functorWithIndexCoproduct ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined)))]
         [foldableWithIndexCoproduct2 (Data.FoldableWithIndex.foldableWithIndexCoproduct ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined)))]
         [traversableCoproduct3 (Data.Traversable.traversableCoproduct ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
          (scm:lambda (dictTraversableWithIndex14)
            (scm:let*
              ([functorWithIndexCoproduct15 (functorWithIndexCoproduct1 ((rt:record-ref dictTraversableWithIndex14 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined)))]
               [foldableWithIndexCoproduct16 (foldableWithIndexCoproduct2 ((rt:record-ref dictTraversableWithIndex14 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined)))]
               [traversableCoproduct17 (traversableCoproduct3 ((rt:record-ref dictTraversableWithIndex14 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative8)
                  (scm:let*
                    ([_9 ((rt:record-ref ((rt:record-ref dictApplicative8 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))]
                     [traverseWithIndex310 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) dictApplicative8)]
                     [traverseWithIndex411 ((rt:record-ref dictTraversableWithIndex14 (scm:string->symbol "traverseWithIndex")) dictApplicative8)])
                      (scm:lambda (f12)
                        (scm:let*
                          ([_13 ((rt:record-ref _9 (scm:string->symbol "map")) (scm:lambda (x13)
                            (Data.Either.Left x13)))]
                           [_14 (traverseWithIndex310 (scm:lambda (x14)
                            (f12 (Data.Either.Left x14))))]
                           [_15 ((rt:record-ref _9 (scm:string->symbol "map")) (scm:lambda (x15)
                            (Data.Either.Right x15)))]
                           [_16 (traverseWithIndex411 (scm:lambda (x16)
                            (f12 (Data.Either.Right x16))))])
                            (scm:lambda (v217)
                              (scm:cond
                                [(Data.Either.Left? v217) (_13 (_14 (Data.Either.Left-value0 v217)))]
                                [(Data.Either.Right? v217) (_15 (_16 (Data.Either.Right-value0 v217)))]
                                [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
                  functorWithIndexCoproduct15)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
                  foldableWithIndexCoproduct16)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
                  traversableCoproduct17))))))))

  (scm:define traversableWithIndexConst
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) v12))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexConst)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexConst)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableConst))))

  (scm:define traversableWithIndexConj
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableConj (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexConj)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexConj)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableConj))))

  (scm:define traversableWithIndexCompose
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [foldableWithIndexCompose3 (Data.FoldableWithIndex.foldableWithIndexCompose ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined)))]
         [traversableCompose4 (Data.Traversable.traversableCompose ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
          (scm:lambda (dictTraversableWithIndex15)
            (scm:let*
              ([_6 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined))]
               [_7 ((rt:record-ref _6 (scm:string->symbol "Functor0")) (scm:quote undefined))]
               [functorWithIndexCompose18 (scm:let ([functorCompose18 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f8)
                (scm:lambda (v9)
                  (((rt:record-ref _2 (scm:string->symbol "map")) ((rt:record-ref _7 (scm:string->symbol "map")) f8)) v9)))))])
                (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f9)
                  (scm:lambda (v10)
                    (((rt:record-ref _1 (scm:string->symbol "mapWithIndex")) (scm:lambda (x11)
                      ((rt:record-ref _6 (scm:string->symbol "mapWithIndex")) (scm:lambda (b12)
                        (f9 (Data.Tuple.Tuple* x11 b12)))))) v10)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorCompose18))))]
               [foldableWithIndexCompose19 (foldableWithIndexCompose3 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined)))]
               [traversableCompose110 (traversableCompose4 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative11)
                  (scm:let*
                    ([traverseWithIndex312 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) dictApplicative11)]
                     [traverseWithIndex413 ((rt:record-ref dictTraversableWithIndex15 (scm:string->symbol "traverseWithIndex")) dictApplicative11)])
                      (scm:lambda (f14)
                        (scm:lambda (v15)
                          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative11 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Compose.Compose) ((traverseWithIndex312 (scm:lambda (x16)
                            (traverseWithIndex413 (scm:lambda (b17)
                              (f14 (Data.Tuple.Tuple* x16 b17)))))) v15))))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
                  functorWithIndexCompose18)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
                  foldableWithIndexCompose19)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
                  traversableCompose110))))))))

  (rt:define-lazy $lazy-traversableWithIndexArray "traversableWithIndexArray" "Data.TraversableWithIndex"
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:let ([sequence11 ((rt:record-ref ((rt:record-ref ($lazy-traversableWithIndexArray) (scm:string->symbol "Traversable2")) (scm:quote undefined)) (scm:string->symbol "sequence")) dictApplicative0)])
        (scm:lambda (f2)
          (scm:let ([_3 ((rt:record-ref ((rt:record-ref ($lazy-traversableWithIndexArray) (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined)) (scm:string->symbol "mapWithIndex")) f2)])
            (scm:lambda (x4)
              (sequence11 (_3 x4)))))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexArray)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexArray)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableArray))))

  (scm:define traversableWithIndexArray
    ($lazy-traversableWithIndexArray))

  (scm:define traversableWithIndexApp
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorWithIndexApp3 (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f3)
          (scm:lambda (v4)
            (((rt:record-ref _1 (scm:string->symbol "mapWithIndex")) f3) v4)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
          _2)))]
         [_4 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined))]
         [_5 ((rt:record-ref _4 (scm:string->symbol "Foldable0")) (scm:quote undefined))]
         [foldableWithIndexApp6 (scm:let ([foldableApp6 (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f6)
          (scm:lambda (i7)
            (scm:lambda (v8)
              ((((rt:record-ref _5 (scm:string->symbol "foldr")) f6) i7) v8))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f6)
          (scm:lambda (i7)
            (scm:lambda (v8)
              ((((rt:record-ref _5 (scm:string->symbol "foldl")) f6) i7) v8))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid6)
          ((rt:record-ref _5 (scm:string->symbol "foldMap")) dictMonoid6))))])
          (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f7)
            (scm:lambda (z8)
              (scm:lambda (v9)
                ((((rt:record-ref _4 (scm:string->symbol "foldrWithIndex")) f7) z8) v9))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f7)
            (scm:lambda (z8)
              (scm:lambda (v9)
                ((((rt:record-ref _4 (scm:string->symbol "foldlWithIndex")) f7) z8) v9))))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid7)
            ((rt:record-ref _4 (scm:string->symbol "foldMapWithIndex")) dictMonoid7))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
            foldableApp6))))]
         [traversableApp7 (Data.Traversable.traversableApp ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
          (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative8)
            (scm:let ([traverseWithIndex29 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) dictApplicative8)])
              (scm:lambda (f10)
                (scm:lambda (v11)
                  (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative8 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.App.App) ((traverseWithIndex29 f10) v11))))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
            functorWithIndexApp3)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
            foldableWithIndexApp6)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
            traversableApp7))))))

  (scm:define traversableWithIndexAdditive
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref Data.Traversable.traversableAdditive (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      Data.FunctorWithIndex.functorWithIndexAdditive)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      Data.FoldableWithIndex.foldableWithIndexAdditive)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      Data.Traversable.traversableAdditive))))

  (scm:define mapAccumRWithIndex
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let ([traverseWithIndex11 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) Data.Traversable.Accum.Internal.applicativeStateR)])
        (scm:lambda (f2)
          (scm:lambda (s03)
            (scm:lambda (xs4)
              (((traverseWithIndex11 (scm:lambda (i5)
                (scm:lambda (a6)
                  (scm:lambda (s7)
                    (((f2 i5) s7) a6))))) xs4) s03)))))))

  (scm:define scanrWithIndex
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let ([mapAccumRWithIndex11 (mapAccumRWithIndex dictTraversableWithIndex0)])
        (scm:lambda (f2)
          (scm:lambda (b03)
            (scm:lambda (xs4)
              (rt:record-ref (((mapAccumRWithIndex11 (scm:lambda (i5)
                (scm:lambda (b6)
                  (scm:lambda (a7)
                    (scm:let ([b$p8 (((f2 i5) a7) b6)])
                      (scm:list (scm:cons (scm:string->symbol "accum") b$p8) (scm:cons (scm:string->symbol "value") b$p8))))))) b03) xs4) (scm:string->symbol "value"))))))))

  (scm:define mapAccumLWithIndex
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let ([traverseWithIndex11 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) Data.Traversable.Accum.Internal.applicativeStateL)])
        (scm:lambda (f2)
          (scm:lambda (s03)
            (scm:lambda (xs4)
              (((traverseWithIndex11 (scm:lambda (i5)
                (scm:lambda (a6)
                  (scm:lambda (s7)
                    (((f2 i5) s7) a6))))) xs4) s03)))))))

  (scm:define scanlWithIndex
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let ([mapAccumLWithIndex11 (mapAccumLWithIndex dictTraversableWithIndex0)])
        (scm:lambda (f2)
          (scm:lambda (b03)
            (scm:lambda (xs4)
              (rt:record-ref (((mapAccumLWithIndex11 (scm:lambda (i5)
                (scm:lambda (b6)
                  (scm:lambda (a7)
                    (scm:let ([b$p8 (((f2 i5) b6) a7)])
                      (scm:list (scm:cons (scm:string->symbol "accum") b$p8) (scm:cons (scm:string->symbol "value") b$p8))))))) b03) xs4) (scm:string->symbol "value"))))))))

  (scm:define forWithIndex
    (scm:lambda (dictApplicative0)
      (scm:lambda (dictTraversableWithIndex1)
        (scm:let ([_2 ((rt:record-ref dictTraversableWithIndex1 (scm:string->symbol "traverseWithIndex")) dictApplicative0)])
          (scm:lambda (b3)
            (scm:lambda (a4)
              ((_2 a4) b3))))))))
