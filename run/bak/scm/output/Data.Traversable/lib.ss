#!r6rs
#!chezscheme
(library
  (Data.Traversable lib)
  (export
    for
    identity
    mapAccumL
    mapAccumR
    scanl
    scanr
    sequence
    sequenceDefault
    traversableAdditive
    traversableApp
    traversableArray
    traversableCompose
    traversableConj
    traversableConst
    traversableCoproduct
    traversableDisj
    traversableDual
    traversableEither
    traversableFirst
    traversableIdentity
    traversableLast
    traversableMaybe
    traversableMultiplicative
    traversableProduct
    traversableTuple
    traverse
    traverseArrayImpl
    traverseDefault)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Const lib) Data.Const.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Foldable lib) Data.Foldable.)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.Functor.App lib) Data.Functor.App.)
    (prefix (Data.Functor.Compose lib) Data.Functor.Compose.)
    (prefix (Data.Functor.Product lib) Data.Functor.Product.)
    (prefix (Data.Identity lib) Data.Identity.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Maybe.First lib) Data.Maybe.First.)
    (prefix (Data.Maybe.Last lib) Data.Maybe.Last.)
    (prefix (Data.Monoid.Additive lib) Data.Monoid.Additive.)
    (prefix (Data.Monoid.Conj lib) Data.Monoid.Conj.)
    (prefix (Data.Monoid.Disj lib) Data.Monoid.Disj.)
    (prefix (Data.Monoid.Dual lib) Data.Monoid.Dual.)
    (prefix (Data.Monoid.Multiplicative lib) Data.Monoid.Multiplicative.)
    (prefix (Data.Traversable.Accum.Internal lib) Data.Traversable.Accum.Internal.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (Data.Traversable foreign))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define traverse
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "traverse"))))

  (scm:define traversableTuple
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (Data.Tuple.Tuple (Data.Tuple.Tuple-value0 v2))) (f1 (Data.Tuple.Tuple-value1 v2))))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (Data.Tuple.Tuple (Data.Tuple.Tuple-value0 v1))) (Data.Tuple.Tuple-value1 v1))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Tuple.functorTuple)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableTuple))))

  (scm:define traversableMultiplicative
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Multiplicative.Multiplicative) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Multiplicative.Multiplicative) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Multiplicative.functorMultiplicative)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableMultiplicative))))

  (scm:define traversableMaybe
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(Data.Maybe.Nothing? v12) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) Data.Maybe.Nothing)]
            [(Data.Maybe.Just? v12) (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Maybe.Just) (v1 (Data.Maybe.Just-value0 v12)))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:cond
          [(Data.Maybe.Nothing? v1) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) Data.Maybe.Nothing)]
          [(Data.Maybe.Just? v1) (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Maybe.Just) (Data.Maybe.Just-value0 v1))]
          [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableMaybe))))

  (scm:define traversableIdentity
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Identity.Identity) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Identity.Identity) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Identity.functorIdentity)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableIdentity))))

  (scm:define traversableEither
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:cond
            [(Data.Either.Left? v12) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) (Data.Either.Left (Data.Either.Left-value0 v12)))]
            [(Data.Either.Right? v12) (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Either.Right) (v1 (Data.Either.Right-value0 v12)))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:cond
          [(Data.Either.Left? v1) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) (Data.Either.Left (Data.Either.Left-value0 v1)))]
          [(Data.Either.Right? v1) (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Either.Right) (Data.Either.Right-value0 v1))]
          [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Either.functorEither)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableEither))))

  (scm:define traversableDual
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Dual.Dual) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Dual.Dual) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Dual.functorDual)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableDual))))

  (scm:define traversableDisj
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Disj.Disj) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Disj.Disj) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Disj.functorDisj)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableDisj))))

  (scm:define traversableConst
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) v12))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Const.functorConst)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableConst))))

  (scm:define traversableConj
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Conj.Conj) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Conj.Conj) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Conj.functorConj)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableConj))))

  (scm:define traversableCompose
    (scm:lambda (dictTraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [_2 ((rt:record-ref dictTraversable0 (scm:string->symbol "Foldable1")) (scm:quote undefined))])
          (scm:lambda (dictTraversable13)
            (scm:let*
              ([_4 ((rt:record-ref dictTraversable13 (scm:string->symbol "Functor0")) (scm:quote undefined))]
               [functorCompose15 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f5)
                (scm:lambda (v6)
                  (((rt:record-ref _1 (scm:string->symbol "map")) ((rt:record-ref _4 (scm:string->symbol "map")) f5)) v6)))))]
               [_6 ((rt:record-ref dictTraversable13 (scm:string->symbol "Foldable1")) (scm:quote undefined))]
               [foldableCompose17 (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f7)
                (scm:lambda (i8)
                  (scm:lambda (v9)
                    ((((rt:record-ref _2 (scm:string->symbol "foldr")) (scm:let ([_10 ((rt:record-ref _6 (scm:string->symbol "foldr")) f7)])
                      (scm:lambda (b11)
                        (scm:lambda (a12)
                          ((_10 a12) b11))))) i8) v9))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f7)
                (scm:lambda (i8)
                  (scm:lambda (v9)
                    ((((rt:record-ref _2 (scm:string->symbol "foldl")) ((rt:record-ref _6 (scm:string->symbol "foldl")) f7)) i8) v9))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid7)
                (scm:let*
                  ([foldMap48 ((rt:record-ref _2 (scm:string->symbol "foldMap")) dictMonoid7)]
                   [foldMap59 ((rt:record-ref _6 (scm:string->symbol "foldMap")) dictMonoid7)])
                    (scm:lambda (f10)
                      (scm:lambda (v11)
                        ((foldMap48 (foldMap59 f10)) v11)))))))])
                (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative8)
                  (scm:let*
                    ([traverse49 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative8)]
                     [traverse510 ((rt:record-ref dictTraversable13 (scm:string->symbol "traverse")) dictApplicative8)])
                      (scm:lambda (f11)
                        (scm:lambda (v12)
                          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative8 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Compose.Compose) ((traverse49 (traverse510 f11)) v12))))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative8)
                  (((rt:record-ref ((traversableCompose dictTraversable0) dictTraversable13) (scm:string->symbol "traverse")) dictApplicative8) identity))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorCompose15)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
                  foldableCompose17))))))))

  (scm:define traversableAdditive
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Additive.Additive) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Additive.Additive) v1)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Monoid.Additive.functorAdditive)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableAdditive))))

  (scm:define sequenceDefault
    (scm:lambda (dictTraversable0)
      (scm:lambda (dictApplicative1)
        (((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative1) identity))))

  (rt:define-lazy $lazy-traversableArray "traversableArray" "Data.Traversable"
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:let ([Apply01 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))])
        (((traverseArrayImpl (rt:record-ref Apply01 (scm:string->symbol "apply"))) (rt:record-ref ((rt:record-ref Apply01 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map"))) (rt:record-ref dictApplicative0 (scm:string->symbol "pure")))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (((rt:record-ref ($lazy-traversableArray) (scm:string->symbol "traverse")) dictApplicative0) identity))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorArray)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableArray))))

  (scm:define traversableArray
    ($lazy-traversableArray))

  (scm:define sequence
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "sequence"))))

  (scm:define traversableApp
    (scm:lambda (dictTraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [_2 ((rt:record-ref dictTraversable0 (scm:string->symbol "Foldable1")) (scm:quote undefined))]
         [foldableApp3 (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f3)
          (scm:lambda (i4)
            (scm:lambda (v5)
              ((((rt:record-ref _2 (scm:string->symbol "foldr")) f3) i4) v5))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f3)
          (scm:lambda (i4)
            (scm:lambda (v5)
              ((((rt:record-ref _2 (scm:string->symbol "foldl")) f3) i4) v5))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid3)
          ((rt:record-ref _2 (scm:string->symbol "foldMap")) dictMonoid3))))])
          (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative4)
            (scm:let ([traverse35 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative4)])
              (scm:lambda (f6)
                (scm:lambda (v7)
                  (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative4 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.App.App) ((traverse35 f6) v7))))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative4)
            (scm:let ([sequence35 ((rt:record-ref dictTraversable0 (scm:string->symbol "sequence")) dictApplicative4)])
              (scm:lambda (v6)
                (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative4 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.App.App) (sequence35 v6)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            _1)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
            foldableApp3))))))

  (scm:define traversableCoproduct
    (scm:lambda (dictTraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [foldableCoproduct2 (Data.Foldable.foldableCoproduct ((rt:record-ref dictTraversable0 (scm:string->symbol "Foldable1")) (scm:quote undefined)))])
          (scm:lambda (dictTraversable13)
            (scm:let*
              ([_4 ((rt:record-ref dictTraversable13 (scm:string->symbol "Functor0")) (scm:quote undefined))]
               [functorCoproduct15 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f5)
                (scm:lambda (v6)
                  (scm:let*
                    ([_7 ((rt:record-ref _1 (scm:string->symbol "map")) f5)]
                     [_8 ((rt:record-ref _4 (scm:string->symbol "map")) f5)])
                      (scm:cond
                        [(Data.Either.Left? v6) (Data.Either.Left (_7 (Data.Either.Left-value0 v6)))]
                        [(Data.Either.Right? v6) (Data.Either.Right (_8 (Data.Either.Right-value0 v6)))]
                        [scm:else (rt:fail)]))))))]
               [foldableCoproduct16 (foldableCoproduct2 ((rt:record-ref dictTraversable13 (scm:string->symbol "Foldable1")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative7)
                  (scm:let*
                    ([_8 ((rt:record-ref ((rt:record-ref dictApplicative7 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))]
                     [traverse49 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative7)]
                     [traverse510 ((rt:record-ref dictTraversable13 (scm:string->symbol "traverse")) dictApplicative7)])
                      (scm:lambda (f11)
                        (scm:let*
                          ([_12 ((rt:record-ref _8 (scm:string->symbol "map")) (scm:lambda (x12)
                            (Data.Either.Left x12)))]
                           [_13 (traverse49 f11)]
                           [_14 ((rt:record-ref _8 (scm:string->symbol "map")) (scm:lambda (x14)
                            (Data.Either.Right x14)))]
                           [_15 (traverse510 f11)])
                            (scm:lambda (v216)
                              (scm:cond
                                [(Data.Either.Left? v216) (_12 (_13 (Data.Either.Left-value0 v216)))]
                                [(Data.Either.Right? v216) (_14 (_15 (Data.Either.Right-value0 v216)))]
                                [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative7)
                  (scm:let*
                    ([_8 ((rt:record-ref ((rt:record-ref dictApplicative7 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))]
                     [_9 ((rt:record-ref _8 (scm:string->symbol "map")) (scm:lambda (x9)
                      (Data.Either.Left x9)))]
                     [_10 ((rt:record-ref dictTraversable0 (scm:string->symbol "sequence")) dictApplicative7)]
                     [_11 ((rt:record-ref _8 (scm:string->symbol "map")) (scm:lambda (x11)
                      (Data.Either.Right x11)))]
                     [_12 ((rt:record-ref dictTraversable13 (scm:string->symbol "sequence")) dictApplicative7)])
                      (scm:lambda (v213)
                        (scm:cond
                          [(Data.Either.Left? v213) (_9 (_10 (Data.Either.Left-value0 v213)))]
                          [(Data.Either.Right? v213) (_11 (_12 (Data.Either.Right-value0 v213)))]
                          [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorCoproduct15)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
                  foldableCoproduct16))))))))

  (scm:define traversableFirst
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Maybe.First.First) ((((rt:record-ref traversableMaybe (scm:string->symbol "traverse")) dictApplicative0) f1) v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Maybe.First.First) (((rt:record-ref traversableMaybe (scm:string->symbol "sequence")) dictApplicative0) v1))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableFirst))))

  (scm:define traversableLast
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Maybe.Last.Last) ((((rt:record-ref traversableMaybe (scm:string->symbol "traverse")) dictApplicative0) f1) v2)))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Maybe.Last.Last) (((rt:record-ref traversableMaybe (scm:string->symbol "sequence")) dictApplicative0) v1))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Maybe.functorMaybe)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      Data.Foldable.foldableLast))))

  (scm:define traversableProduct
    (scm:lambda (dictTraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [foldableProduct2 (Data.Foldable.foldableProduct ((rt:record-ref dictTraversable0 (scm:string->symbol "Foldable1")) (scm:quote undefined)))])
          (scm:lambda (dictTraversable13)
            (scm:let*
              ([_4 ((rt:record-ref dictTraversable13 (scm:string->symbol "Functor0")) (scm:quote undefined))]
               [functorProduct15 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f5)
                (scm:lambda (v6)
                  (Data.Tuple.Tuple* (((rt:record-ref _1 (scm:string->symbol "map")) f5) (Data.Tuple.Tuple-value0 v6)) (((rt:record-ref _4 (scm:string->symbol "map")) f5) (Data.Tuple.Tuple-value1 v6)))))))]
               [foldableProduct16 (foldableProduct2 ((rt:record-ref dictTraversable13 (scm:string->symbol "Foldable1")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative7)
                  (scm:let*
                    ([_8 ((rt:record-ref dictApplicative7 (scm:string->symbol "Apply0")) (scm:quote undefined))]
                     [traverse49 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative7)]
                     [traverse510 ((rt:record-ref dictTraversable13 (scm:string->symbol "traverse")) dictApplicative7)])
                      (scm:lambda (f11)
                        (scm:lambda (v12)
                          (((rt:record-ref _8 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref _8 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Product.product) ((traverse49 f11) (Data.Tuple.Tuple-value0 v12)))) ((traverse510 f11) (Data.Tuple.Tuple-value1 v12)))))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative7)
                  (scm:let*
                    ([_8 ((rt:record-ref dictApplicative7 (scm:string->symbol "Apply0")) (scm:quote undefined))]
                     [sequence49 ((rt:record-ref dictTraversable0 (scm:string->symbol "sequence")) dictApplicative7)]
                     [sequence510 ((rt:record-ref dictTraversable13 (scm:string->symbol "sequence")) dictApplicative7)])
                      (scm:lambda (v11)
                        (((rt:record-ref _8 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref _8 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Product.product) (sequence49 (Data.Tuple.Tuple-value0 v11)))) (sequence510 (Data.Tuple.Tuple-value1 v11))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
                  functorProduct15)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
                  foldableProduct16))))))))

  (scm:define traverseDefault
    (scm:lambda (dictTraversable0)
      (scm:lambda (dictApplicative1)
        (scm:let ([sequence32 ((rt:record-ref dictTraversable0 (scm:string->symbol "sequence")) dictApplicative1)])
          (scm:lambda (f3)
            (scm:lambda (ta4)
              (sequence32 (((rt:record-ref ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) f3) ta4))))))))

  (scm:define mapAccumR
    (scm:lambda (dictTraversable0)
      (scm:let ([traverse21 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) Data.Traversable.Accum.Internal.applicativeStateR)])
        (scm:lambda (f2)
          (scm:lambda (s03)
            (scm:lambda (xs4)
              (((traverse21 (scm:lambda (a5)
                (scm:lambda (s6)
                  ((f2 s6) a5)))) xs4) s03)))))))

  (scm:define scanr
    (scm:lambda (dictTraversable0)
      (scm:let ([mapAccumR11 (mapAccumR dictTraversable0)])
        (scm:lambda (f2)
          (scm:lambda (b03)
            (scm:lambda (xs4)
              (rt:record-ref (((mapAccumR11 (scm:lambda (b5)
                (scm:lambda (a6)
                  (scm:let ([b$p7 ((f2 a6) b5)])
                    (scm:list (scm:cons (scm:string->symbol "accum") b$p7) (scm:cons (scm:string->symbol "value") b$p7)))))) b03) xs4) (scm:string->symbol "value"))))))))

  (scm:define mapAccumL
    (scm:lambda (dictTraversable0)
      (scm:let ([traverse21 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) Data.Traversable.Accum.Internal.applicativeStateL)])
        (scm:lambda (f2)
          (scm:lambda (s03)
            (scm:lambda (xs4)
              (((traverse21 (scm:lambda (a5)
                (scm:lambda (s6)
                  ((f2 s6) a5)))) xs4) s03)))))))

  (scm:define scanl
    (scm:lambda (dictTraversable0)
      (scm:let ([mapAccumL11 (mapAccumL dictTraversable0)])
        (scm:lambda (f2)
          (scm:lambda (b03)
            (scm:lambda (xs4)
              (rt:record-ref (((mapAccumL11 (scm:lambda (b5)
                (scm:lambda (a6)
                  (scm:let ([b$p7 ((f2 b5) a6)])
                    (scm:list (scm:cons (scm:string->symbol "accum") b$p7) (scm:cons (scm:string->symbol "value") b$p7)))))) b03) xs4) (scm:string->symbol "value"))))))))

  (scm:define for
    (scm:lambda (dictApplicative0)
      (scm:lambda (dictTraversable1)
        (scm:let ([traverse22 ((rt:record-ref dictTraversable1 (scm:string->symbol "traverse")) dictApplicative0)])
          (scm:lambda (x3)
            (scm:lambda (f4)
              ((traverse22 f4) x3))))))))
