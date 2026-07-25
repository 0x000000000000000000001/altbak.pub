#!r6rs
#!chezscheme
(library
  (Data.Bitraversable lib)
  (export
    bifor
    bisequence
    bisequenceDefault
    bitraversableClown
    bitraversableConst
    bitraversableEither
    bitraversableFlip
    bitraversableJoker
    bitraversableProduct2
    bitraversableTuple
    bitraverse
    bitraverseDefault
    identity
    lfor
    ltraverse
    rfor
    rtraverse)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Bifoldable lib) Data.Bifoldable.)
    (prefix (Data.Bifunctor lib) Data.Bifunctor.)
    (prefix (Data.Const lib) Data.Const.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Functor.Clown lib) Data.Functor.Clown.)
    (prefix (Data.Functor.Flip lib) Data.Functor.Flip.)
    (prefix (Data.Functor.Joker lib) Data.Functor.Joker.)
    (prefix (Data.Functor.Product2 lib) Data.Functor.Product2.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define bitraverse
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bitraverse"))))

  (scm:define lfor
    (scm:lambda (dictBitraversable0)
      (scm:lambda (dictApplicative1)
        (scm:let*
          ([bitraverse22 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative1)]
           [pure3 (rt:record-ref dictApplicative1 (scm:string->symbol "pure"))])
            (scm:lambda (t4)
              (scm:lambda (f5)
                (((bitraverse22 f5) pure3) t4)))))))

  (scm:define ltraverse
    (scm:lambda (dictBitraversable0)
      (scm:lambda (dictApplicative1)
        (scm:let*
          ([bitraverse22 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative1)]
           [pure3 (rt:record-ref dictApplicative1 (scm:string->symbol "pure"))])
            (scm:lambda (f4)
              ((bitraverse22 f4) pure3))))))

  (scm:define rfor
    (scm:lambda (dictBitraversable0)
      (scm:lambda (dictApplicative1)
        (scm:let*
          ([bitraverse22 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative1)]
           [pure3 (rt:record-ref dictApplicative1 (scm:string->symbol "pure"))])
            (scm:lambda (t4)
              (scm:lambda (f5)
                (((bitraverse22 pure3) f5) t4)))))))

  (scm:define rtraverse
    (scm:lambda (dictBitraversable0)
      (scm:lambda (dictApplicative1)
        (((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative1) (rt:record-ref dictApplicative1 (scm:string->symbol "pure"))))))

  (scm:define bitraversableTuple
    (scm:list (scm:cons (scm:string->symbol "bitraverse") (scm:lambda (dictApplicative0)
      (scm:let ([Apply01 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))])
        (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (v4)
              (((rt:record-ref Apply01 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref Apply01 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Tuple.Tuple) (f2 (Data.Tuple.Tuple-value0 v4)))) (g3 (Data.Tuple.Tuple-value1 v4))))))))) (scm:cons (scm:string->symbol "bisequence") (scm:lambda (dictApplicative0)
      (scm:let ([Apply01 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))])
        (scm:lambda (v2)
          (((rt:record-ref Apply01 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref Apply01 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Tuple.Tuple) (Data.Tuple.Tuple-value0 v2))) (Data.Tuple.Tuple-value1 v2)))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
      Data.Bifunctor.bifunctorTuple)) (scm:cons (scm:string->symbol "Bifoldable1") (scm:lambda (_)
      Data.Bifoldable.bifoldableTuple))))

  (scm:define bitraversableJoker
    (scm:lambda (dictTraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [bifunctorJoker2 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (v2)
          (scm:lambda (g3)
            (scm:lambda (v14)
              (((rt:record-ref _1 (scm:string->symbol "map")) g3) v14))))))]
         [_3 ((rt:record-ref dictTraversable0 (scm:string->symbol "Foldable1")) (scm:quote undefined))]
         [bifoldableJoker4 (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (v4)
          (scm:lambda (r5)
            (scm:lambda (u6)
              (scm:lambda (v17)
                ((((rt:record-ref _3 (scm:string->symbol "foldr")) r5) u6) v17)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (v4)
          (scm:lambda (r5)
            (scm:lambda (u6)
              (scm:lambda (v17)
                ((((rt:record-ref _3 (scm:string->symbol "foldl")) r5) u6) v17)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid4)
          (scm:let ([foldMap15 ((rt:record-ref _3 (scm:string->symbol "foldMap")) dictMonoid4)])
            (scm:lambda (v6)
              (scm:lambda (r7)
                (scm:lambda (v18)
                  ((foldMap15 r7) v18))))))))])
          (scm:list (scm:cons (scm:string->symbol "bitraverse") (scm:lambda (dictApplicative5)
            (scm:let ([traverse16 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative5)])
              (scm:lambda (v7)
                (scm:lambda (r8)
                  (scm:lambda (v19)
                    (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Joker.Joker) ((traverse16 r8) v19)))))))) (scm:cons (scm:string->symbol "bisequence") (scm:lambda (dictApplicative5)
            (scm:let ([sequence16 ((rt:record-ref dictTraversable0 (scm:string->symbol "sequence")) dictApplicative5)])
              (scm:lambda (v7)
                (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Joker.Joker) (sequence16 v7)))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorJoker2)) (scm:cons (scm:string->symbol "Bifoldable1") (scm:lambda (_)
            bifoldableJoker4))))))

  (scm:define bitraversableEither
    (scm:list (scm:cons (scm:string->symbol "bitraverse") (scm:lambda (dictApplicative0)
      (scm:let ([_1 ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (v24)
              (scm:cond
                [(Data.Either.Left? v24) (((rt:record-ref _1 (scm:string->symbol "map")) Data.Either.Left) (v2 (Data.Either.Left-value0 v24)))]
                [(Data.Either.Right? v24) (((rt:record-ref _1 (scm:string->symbol "map")) Data.Either.Right) (v13 (Data.Either.Right-value0 v24)))]
                [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "bisequence") (scm:lambda (dictApplicative0)
      (scm:let ([_1 ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (v2)
          (scm:cond
            [(Data.Either.Left? v2) (((rt:record-ref _1 (scm:string->symbol "map")) Data.Either.Left) (Data.Either.Left-value0 v2))]
            [(Data.Either.Right? v2) (((rt:record-ref _1 (scm:string->symbol "map")) Data.Either.Right) (Data.Either.Right-value0 v2))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
      Data.Bifunctor.bifunctorEither)) (scm:cons (scm:string->symbol "Bifoldable1") (scm:lambda (_)
      Data.Bifoldable.bifoldableEither))))

  (scm:define bitraversableConst
    (scm:list (scm:cons (scm:string->symbol "bitraverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Const.Const) (f1 v13))))))) (scm:cons (scm:string->symbol "bisequence") (scm:lambda (dictApplicative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Const.Const) v1)))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
      Data.Bifunctor.bifunctorConst)) (scm:cons (scm:string->symbol "Bifoldable1") (scm:lambda (_)
      Data.Bifoldable.bifoldableConst))))

  (scm:define bitraversableClown
    (scm:lambda (dictTraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [bifunctorClown2 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref _1 (scm:string->symbol "map")) f2) v14))))))]
         [_3 ((rt:record-ref dictTraversable0 (scm:string->symbol "Foldable1")) (scm:quote undefined))]
         [bifoldableClown4 (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (l4)
          (scm:lambda (v5)
            (scm:lambda (u6)
              (scm:lambda (v17)
                ((((rt:record-ref _3 (scm:string->symbol "foldr")) l4) u6) v17)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (l4)
          (scm:lambda (v5)
            (scm:lambda (u6)
              (scm:lambda (v17)
                ((((rt:record-ref _3 (scm:string->symbol "foldl")) l4) u6) v17)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid4)
          (scm:let ([foldMap15 ((rt:record-ref _3 (scm:string->symbol "foldMap")) dictMonoid4)])
            (scm:lambda (l6)
              (scm:lambda (v7)
                (scm:lambda (v18)
                  ((foldMap15 l6) v18))))))))])
          (scm:list (scm:cons (scm:string->symbol "bitraverse") (scm:lambda (dictApplicative5)
            (scm:let ([traverse16 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative5)])
              (scm:lambda (l7)
                (scm:lambda (v8)
                  (scm:lambda (v19)
                    (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Clown.Clown) ((traverse16 l7) v19)))))))) (scm:cons (scm:string->symbol "bisequence") (scm:lambda (dictApplicative5)
            (scm:let ([sequence16 ((rt:record-ref dictTraversable0 (scm:string->symbol "sequence")) dictApplicative5)])
              (scm:lambda (v7)
                (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Clown.Clown) (sequence16 v7)))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorClown2)) (scm:cons (scm:string->symbol "Bifoldable1") (scm:lambda (_)
            bifoldableClown4))))))

  (scm:define bisequenceDefault
    (scm:lambda (dictBitraversable0)
      (scm:lambda (dictApplicative1)
        ((((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative1) identity) identity))))

  (scm:define bisequence
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bisequence"))))

  (scm:define bitraversableFlip
    (scm:lambda (dictBitraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictBitraversable0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
         [bifunctorFlip2 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (v4)
              ((((rt:record-ref _1 (scm:string->symbol "bimap")) g3) f2) v4))))))]
         [_3 ((rt:record-ref dictBitraversable0 (scm:string->symbol "Bifoldable1")) (scm:quote undefined))]
         [bifoldableFlip4 (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (r4)
          (scm:lambda (l5)
            (scm:lambda (u6)
              (scm:lambda (v7)
                (((((rt:record-ref _3 (scm:string->symbol "bifoldr")) l5) r4) u6) v7)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (r4)
          (scm:lambda (l5)
            (scm:lambda (u6)
              (scm:lambda (v7)
                (((((rt:record-ref _3 (scm:string->symbol "bifoldl")) l5) r4) u6) v7)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid4)
          (scm:let ([bifoldMap25 ((rt:record-ref _3 (scm:string->symbol "bifoldMap")) dictMonoid4)])
            (scm:lambda (r6)
              (scm:lambda (l7)
                (scm:lambda (v8)
                  (((bifoldMap25 l7) r6) v8))))))))])
          (scm:list (scm:cons (scm:string->symbol "bitraverse") (scm:lambda (dictApplicative5)
            (scm:let ([bitraverse26 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative5)])
              (scm:lambda (r7)
                (scm:lambda (l8)
                  (scm:lambda (v9)
                    (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Flip.Flip) (((bitraverse26 l8) r7) v9)))))))) (scm:cons (scm:string->symbol "bisequence") (scm:lambda (dictApplicative5)
            (scm:let ([bisequence26 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bisequence")) dictApplicative5)])
              (scm:lambda (v7)
                (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Flip.Flip) (bisequence26 v7)))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
            bifunctorFlip2)) (scm:cons (scm:string->symbol "Bifoldable1") (scm:lambda (_)
            bifoldableFlip4))))))

  (scm:define bitraversableProduct2
    (scm:lambda (dictBitraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictBitraversable0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
         [bifoldableProduct22 (Data.Bifoldable.bifoldableProduct2 ((rt:record-ref dictBitraversable0 (scm:string->symbol "Bifoldable1")) (scm:quote undefined)))])
          (scm:lambda (dictBitraversable13)
            (scm:let*
              ([_4 ((rt:record-ref dictBitraversable13 (scm:string->symbol "Bifunctor0")) (scm:quote undefined))]
               [bifunctorProduct215 (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (v7)
                    (Data.Functor.Product2.Product2* ((((rt:record-ref _1 (scm:string->symbol "bimap")) f5) g6) (Data.Functor.Product2.Product2-value0 v7)) ((((rt:record-ref _4 (scm:string->symbol "bimap")) f5) g6) (Data.Functor.Product2.Product2-value1 v7))))))))]
               [bifoldableProduct216 (bifoldableProduct22 ((rt:record-ref dictBitraversable13 (scm:string->symbol "Bifoldable1")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "bitraverse") (scm:lambda (dictApplicative7)
                  (scm:let*
                    ([Apply08 ((rt:record-ref dictApplicative7 (scm:string->symbol "Apply0")) (scm:quote undefined))]
                     [bitraverse39 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative7)]
                     [bitraverse410 ((rt:record-ref dictBitraversable13 (scm:string->symbol "bitraverse")) dictApplicative7)])
                      (scm:lambda (l11)
                        (scm:lambda (r12)
                          (scm:lambda (v13)
                            (((rt:record-ref Apply08 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref Apply08 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Product2.Product2) (((bitraverse39 l11) r12) (Data.Functor.Product2.Product2-value0 v13)))) (((bitraverse410 l11) r12) (Data.Functor.Product2.Product2-value1 v13))))))))) (scm:cons (scm:string->symbol "bisequence") (scm:lambda (dictApplicative7)
                  (scm:let*
                    ([Apply08 ((rt:record-ref dictApplicative7 (scm:string->symbol "Apply0")) (scm:quote undefined))]
                     [bisequence39 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bisequence")) dictApplicative7)]
                     [bisequence410 ((rt:record-ref dictBitraversable13 (scm:string->symbol "bisequence")) dictApplicative7)])
                      (scm:lambda (v11)
                        (((rt:record-ref Apply08 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref Apply08 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Functor.Product2.Product2) (bisequence39 (Data.Functor.Product2.Product2-value0 v11)))) (bisequence410 (Data.Functor.Product2.Product2-value1 v11))))))) (scm:cons (scm:string->symbol "Bifunctor0") (scm:lambda (_)
                  bifunctorProduct215)) (scm:cons (scm:string->symbol "Bifoldable1") (scm:lambda (_)
                  bifoldableProduct216))))))))

  (scm:define bitraverseDefault
    (scm:lambda (dictBitraversable0)
      (scm:lambda (dictApplicative1)
        (scm:let ([bisequence22 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bisequence")) dictApplicative1)])
          (scm:lambda (f3)
            (scm:lambda (g4)
              (scm:lambda (t5)
                (bisequence22 ((((rt:record-ref ((rt:record-ref dictBitraversable0 (scm:string->symbol "Bifunctor0")) (scm:quote undefined)) (scm:string->symbol "bimap")) f3) g4) t5)))))))))

  (scm:define bifor
    (scm:lambda (dictBitraversable0)
      (scm:lambda (dictApplicative1)
        (scm:let ([bitraverse22 ((rt:record-ref dictBitraversable0 (scm:string->symbol "bitraverse")) dictApplicative1)])
          (scm:lambda (t3)
            (scm:lambda (f4)
              (scm:lambda (g5)
                (((bitraverse22 f4) g5) t3)))))))))
