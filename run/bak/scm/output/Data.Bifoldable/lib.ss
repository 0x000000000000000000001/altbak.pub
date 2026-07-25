#!r6rs
#!chezscheme
(library
  (Data.Bifoldable lib)
  (export
    biall
    biany
    bifold
    bifoldMap
    bifoldMapDefaultL
    bifoldMapDefaultR
    bifoldableClown
    bifoldableConst
    bifoldableEither
    bifoldableFlip
    bifoldableJoker
    bifoldableProduct2
    bifoldableTuple
    bifoldl
    bifoldlDefault
    bifoldr
    bifoldrDefault
    bifor_
    bisequence_
    bitraverse_
    identity
    monoidDual
    monoidEndo)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Apply lib) Control.Apply.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Functor.Product2 lib) Data.Functor.Product2.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

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

  (scm:define bifoldr
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bifoldr"))))

  (scm:define bitraverse_
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictApplicative1)
        (scm:let*
          ([_2 ((rt:record-ref dictApplicative1 (scm:string->symbol "Apply0")) (scm:quote undefined))]
           [applySecond3 (scm:lambda (a3)
            (scm:lambda (b4)
              (((rt:record-ref _2 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref _2 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (v5)
                Control.Apply.identity)) a3)) b4)))])
            (scm:lambda (f4)
              (scm:lambda (g5)
                ((((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldr")) (scm:lambda (x6)
                  (applySecond3 (f4 x6)))) (scm:lambda (x6)
                  (applySecond3 (g5 x6)))) ((rt:record-ref dictApplicative1 (scm:string->symbol "pure")) Data.Unit.unit))))))))

  (scm:define bifor_
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictApplicative1)
        (scm:let ([bitraverse_22 ((bitraverse_ dictBifoldable0) dictApplicative1)])
          (scm:lambda (t3)
            (scm:lambda (f4)
              (scm:lambda (g5)
                (((bitraverse_22 f4) g5) t3))))))))

  (scm:define bisequence_
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictApplicative1)
        ((((bitraverse_ dictBifoldable0) dictApplicative1) identity) identity))))

  (scm:define bifoldl
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bifoldl"))))

  (scm:define bifoldableTuple
    (scm:list (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (v3)
            (((rt:record-ref ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f1 (Data.Tuple.Tuple-value0 v3))) (g2 (Data.Tuple.Tuple-value1 v3)))))))) (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (z2)
          (scm:lambda (v3)
            ((f0 (Data.Tuple.Tuple-value0 v3)) ((g1 (Data.Tuple.Tuple-value1 v3)) z2))))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (f0)
      (scm:lambda (g1)
        (scm:lambda (z2)
          (scm:lambda (v3)
            ((g1 ((f0 z2) (Data.Tuple.Tuple-value0 v3))) (Data.Tuple.Tuple-value1 v3)))))))))

  (scm:define bifoldableJoker
    (scm:lambda (dictFoldable0)
      (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (v1)
        (scm:lambda (r2)
          (scm:lambda (u3)
            (scm:lambda (v14)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) r2) u3) v14)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (v1)
        (scm:lambda (r2)
          (scm:lambda (u3)
            (scm:lambda (v14)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) r2) u3) v14)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid1)
        (scm:let ([foldMap12 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid1)])
          (scm:lambda (v3)
            (scm:lambda (r4)
              (scm:lambda (v15)
                ((foldMap12 r4) v15))))))))))

  (scm:define bifoldableEither
    (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:lambda (v33)
            (scm:cond
              [(Data.Either.Left? v33) ((v0 (Data.Either.Left-value0 v33)) v22)]
              [(Data.Either.Right? v33) ((v11 (Data.Either.Right-value0 v33)) v22)]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:lambda (v33)
            (scm:cond
              [(Data.Either.Left? v33) ((v0 v22) (Data.Either.Left-value0 v33))]
              [(Data.Either.Right? v33) ((v11 v22) (Data.Either.Right-value0 v33))]
              [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (v23)
            (scm:cond
              [(Data.Either.Left? v23) (v1 (Data.Either.Left-value0 v23))]
              [(Data.Either.Right? v23) (v12 (Data.Either.Right-value0 v23))]
              [scm:else (rt:fail)]))))))))

  (scm:define bifoldableConst
    (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (z2)
          (scm:lambda (v13)
            ((f0 v13) z2)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (z2)
          (scm:lambda (v13)
            ((f0 z2) v13)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (scm:lambda (v13)
            (f1 v13))))))))

  (scm:define bifoldableClown
    (scm:lambda (dictFoldable0)
      (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (l1)
        (scm:lambda (v2)
          (scm:lambda (u3)
            (scm:lambda (v14)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) l1) u3) v14)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (l1)
        (scm:lambda (v2)
          (scm:lambda (u3)
            (scm:lambda (v14)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) l1) u3) v14)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid1)
        (scm:let ([foldMap12 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid1)])
          (scm:lambda (l3)
            (scm:lambda (v4)
              (scm:lambda (v15)
                ((foldMap12 l3) v15))))))))))

  (scm:define bifoldMapDefaultR
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictMonoid1)
        (scm:let*
          ([_2 ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
           [mempty3 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))])
            (scm:lambda (f4)
              (scm:lambda (g5)
                ((((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldr")) (scm:lambda (x6)
                  ((rt:record-ref _2 (scm:string->symbol "append")) (f4 x6)))) (scm:lambda (x6)
                  ((rt:record-ref _2 (scm:string->symbol "append")) (g5 x6)))) mempty3)))))))

  (scm:define bifoldMapDefaultL
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictMonoid1)
        (scm:let*
          ([_2 ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
           [mempty3 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))])
            (scm:lambda (f4)
              (scm:lambda (g5)
                ((((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldl")) (scm:lambda (m6)
                  (scm:lambda (a7)
                    (((rt:record-ref _2 (scm:string->symbol "append")) m6) (f4 a7))))) (scm:lambda (m6)
                  (scm:lambda (b7)
                    (((rt:record-ref _2 (scm:string->symbol "append")) m6) (g5 b7))))) mempty3)))))))

  (scm:define bifoldMap
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "bifoldMap"))))

  (scm:define bifoldableFlip
    (scm:lambda (dictBifoldable0)
      (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (r1)
        (scm:lambda (l2)
          (scm:lambda (u3)
            (scm:lambda (v4)
              (((((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldr")) l2) r1) u3) v4)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (r1)
        (scm:lambda (l2)
          (scm:lambda (u3)
            (scm:lambda (v4)
              (((((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldl")) l2) r1) u3) v4)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid1)
        (scm:let ([bifoldMap22 ((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldMap")) dictMonoid1)])
          (scm:lambda (r3)
            (scm:lambda (l4)
              (scm:lambda (v5)
                (((bifoldMap22 l4) r3) v5))))))))))

  (scm:define bifoldlDefault
    (scm:lambda (dictBifoldable0)
      (scm:let ([bifoldMap11 ((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldMap")) monoidDual)])
        (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (z4)
              (scm:lambda (p5)
                ((((bifoldMap11 (scm:lambda (x6)
                  (scm:lambda (a7)
                    ((f2 a7) x6)))) (scm:lambda (x6)
                  (scm:lambda (a7)
                    ((g3 a7) x6)))) p5) z4))))))))

  (scm:define bifoldrDefault
    (scm:lambda (dictBifoldable0)
      (scm:let ([bifoldMap11 ((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldMap")) monoidEndo)])
        (scm:lambda (f2)
          (scm:lambda (g3)
            (scm:lambda (z4)
              (scm:lambda (p5)
                ((((bifoldMap11 (scm:lambda (x6)
                  (f2 x6))) (scm:lambda (x6)
                  (g3 x6))) p5) z4))))))))

  (scm:define bifoldableProduct2
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictBifoldable11)
        (scm:list (scm:cons (scm:string->symbol "bifoldr") (scm:lambda (l2)
          (scm:lambda (r3)
            (scm:lambda (u4)
              (scm:lambda (m5)
                (((((bifoldrDefault ((bifoldableProduct2 dictBifoldable0) dictBifoldable11)) l2) r3) u4) m5)))))) (scm:cons (scm:string->symbol "bifoldl") (scm:lambda (l2)
          (scm:lambda (r3)
            (scm:lambda (u4)
              (scm:lambda (m5)
                (((((bifoldlDefault ((bifoldableProduct2 dictBifoldable0) dictBifoldable11)) l2) r3) u4) m5)))))) (scm:cons (scm:string->symbol "bifoldMap") (scm:lambda (dictMonoid2)
          (scm:let*
            ([bifoldMap33 ((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldMap")) dictMonoid2)]
             [bifoldMap44 ((rt:record-ref dictBifoldable11 (scm:string->symbol "bifoldMap")) dictMonoid2)])
              (scm:lambda (l5)
                (scm:lambda (r6)
                  (scm:lambda (v7)
                    (((rt:record-ref ((rt:record-ref dictMonoid2 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (((bifoldMap33 l5) r6) (Data.Functor.Product2.Product2-value0 v7))) (((bifoldMap44 l5) r6) (Data.Functor.Product2.Product2-value1 v7)))))))))))))

  (scm:define bifold
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictMonoid1)
        ((((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldMap")) dictMonoid1) identity) identity))))

  (scm:define biany
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictBooleanAlgebra1)
        (scm:let ([bifoldMap22 ((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldMap")) (scm:let*
          ([_2 ((rt:record-ref dictBooleanAlgebra1 (scm:string->symbol "HeytingAlgebra0")) (scm:quote undefined))]
           [semigroupDisj13 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref _2 (scm:string->symbol "disj")) v3) v14)))))])
            (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref _2 (scm:string->symbol "ff"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
              semigroupDisj13)))))])
          (scm:lambda (p3)
            (scm:lambda (q4)
              ((bifoldMap22 (scm:lambda (x5)
                (p3 x5))) (scm:lambda (x5)
                (q4 x5)))))))))

  (scm:define biall
    (scm:lambda (dictBifoldable0)
      (scm:lambda (dictBooleanAlgebra1)
        (scm:let ([bifoldMap22 ((rt:record-ref dictBifoldable0 (scm:string->symbol "bifoldMap")) (scm:let*
          ([_2 ((rt:record-ref dictBooleanAlgebra1 (scm:string->symbol "HeytingAlgebra0")) (scm:quote undefined))]
           [semigroupConj13 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v3)
            (scm:lambda (v14)
              (((rt:record-ref _2 (scm:string->symbol "conj")) v3) v14)))))])
            (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref _2 (scm:string->symbol "tt"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
              semigroupConj13)))))])
          (scm:lambda (p3)
            (scm:lambda (q4)
              ((bifoldMap22 (scm:lambda (x5)
                (p3 x5))) (scm:lambda (x5)
                (q4 x5))))))))))
