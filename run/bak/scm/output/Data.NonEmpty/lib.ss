#!r6rs
#!chezscheme
(library
  (Data.NonEmpty lib)
  (export
    NonEmpty
    NonEmpty*
    NonEmpty-value0
    NonEmpty-value1
    NonEmpty?
    eq1NonEmpty
    eqNonEmpty
    foldable1NonEmpty
    foldableNonEmpty
    foldableWithIndexNonEmpty
    foldl1
    fromNonEmpty
    functorNonEmpty
    functorWithIndex
    head
    oneOf
    ord1NonEmpty
    ordNonEmpty
    semigroupNonEmpty
    showNonEmpty
    singleton
    tail
    traversableNonEmpty
    traversableWithIndexNonEmpty
    unfoldable1NonEmpty)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define-record-type (NonEmpty$ NonEmpty* NonEmpty?)
    (scm:fields (scm:immutable value0 NonEmpty-value0) (scm:immutable value1 NonEmpty-value1)))

  (scm:define NonEmpty
    (scm:lambda (value0)
      (scm:lambda (value1)
        (NonEmpty* value0 value1))))

  (scm:define unfoldable1NonEmpty
    (scm:lambda (dictUnfoldable0)
      (scm:list (scm:cons (scm:string->symbol "unfoldr1") (scm:lambda (f1)
        (scm:lambda (b2)
          (scm:let ([_3 (f1 b2)])
            (NonEmpty* (Data.Tuple.Tuple-value0 _3) (((rt:record-ref dictUnfoldable0 (scm:string->symbol "unfoldr")) (scm:lambda (v14)
              (scm:cond
                [(Data.Maybe.Just? v14) (Data.Maybe.Just (f1 (Data.Maybe.Just-value0 v14)))]
                [scm:else Data.Maybe.Nothing]))) (Data.Tuple.Tuple-value1 _3))))))))))

  (scm:define tail
    (scm:lambda (v0)
      (NonEmpty-value1 v0)))

  (scm:define singleton
    (scm:lambda (dictPlus0)
      (scm:let ([empty1 (rt:record-ref dictPlus0 (scm:string->symbol "empty"))])
        (scm:lambda (a2)
          (NonEmpty* a2 empty1)))))

  (scm:define showNonEmpty
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(NonEmpty ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (NonEmpty-value0 v2))) (rt:string->pstring " ")) ((rt:record-ref dictShow11 (scm:string->symbol "show")) (NonEmpty-value1 v2))) (rt:string->pstring ")"))))))))

  (scm:define semigroupNonEmpty
    (scm:lambda (dictApplicative0)
      (scm:lambda (dictSemigroup1)
        (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (NonEmpty* (NonEmpty-value0 v2) (((rt:record-ref dictSemigroup1 (scm:string->symbol "append")) (NonEmpty-value1 v2)) (((rt:record-ref dictSemigroup1 (scm:string->symbol "append")) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) (NonEmpty-value0 v13))) (NonEmpty-value1 v13)))))))))))

  (scm:define oneOf
    (scm:lambda (dictAlternative0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Plus1")) (scm:quote undefined)) (scm:string->symbol "Alt0")) (scm:quote undefined)) (scm:string->symbol "alt")) ((rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) (NonEmpty-value0 v1))) (NonEmpty-value1 v1)))))

  (scm:define head
    (scm:lambda (v0)
      (NonEmpty-value0 v0)))

  (scm:define functorNonEmpty
    (scm:lambda (dictFunctor0)
      (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f1)
        (scm:lambda (m2)
          (NonEmpty* (f1 (NonEmpty-value0 m2)) (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1) (NonEmpty-value1 m2)))))))))

  (scm:define functorWithIndex
    (scm:lambda (dictFunctorWithIndex0)
      (scm:let*
        ([_1 ((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorNonEmpty12 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (m3)
            (NonEmpty* (f2 (NonEmpty-value0 m3)) (((rt:record-ref _1 (scm:string->symbol "map")) f2) (NonEmpty-value1 m3)))))))])
          (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f3)
            (scm:lambda (v4)
              (NonEmpty* ((f3 Data.Maybe.Nothing) (NonEmpty-value0 v4)) (((rt:record-ref dictFunctorWithIndex0 (scm:string->symbol "mapWithIndex")) (scm:lambda (x5)
                (f3 (Data.Maybe.Just x5)))) (NonEmpty-value1 v4)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorNonEmpty12))))))

  (scm:define fromNonEmpty
    (scm:lambda (f0)
      (scm:lambda (v1)
        ((f0 (NonEmpty-value0 v1)) (NonEmpty-value1 v1)))))

  (scm:define foldableNonEmpty
    (scm:lambda (dictFoldable0)
      (scm:list (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid1)
        (scm:let ([foldMap12 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid1)])
          (scm:lambda (f3)
            (scm:lambda (v4)
              (((rt:record-ref ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f3 (NonEmpty-value0 v4))) ((foldMap12 f3) (NonEmpty-value1 v4)))))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f1)
        (scm:lambda (b2)
          (scm:lambda (v3)
            ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) f1) ((f1 b2) (NonEmpty-value0 v3))) (NonEmpty-value1 v3)))))) (scm:cons (scm:string->symbol "foldr") (scm:lambda (f1)
        (scm:lambda (b2)
          (scm:lambda (v3)
            ((f1 (NonEmpty-value0 v3)) ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) f1) b2) (NonEmpty-value1 v3))))))))))

  (scm:define foldableWithIndexNonEmpty
    (scm:lambda (dictFoldableWithIndex0)
      (scm:let*
        ([_1 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "Foldable0")) (scm:quote undefined))]
         [foldableNonEmpty12 (scm:list (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid2)
          (scm:let ([foldMap13 ((rt:record-ref _1 (scm:string->symbol "foldMap")) dictMonoid2)])
            (scm:lambda (f4)
              (scm:lambda (v5)
                (((rt:record-ref ((rt:record-ref dictMonoid2 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f4 (NonEmpty-value0 v5))) ((foldMap13 f4) (NonEmpty-value1 v5)))))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f2)
          (scm:lambda (b3)
            (scm:lambda (v4)
              ((((rt:record-ref _1 (scm:string->symbol "foldl")) f2) ((f2 b3) (NonEmpty-value0 v4))) (NonEmpty-value1 v4)))))) (scm:cons (scm:string->symbol "foldr") (scm:lambda (f2)
          (scm:lambda (b3)
            (scm:lambda (v4)
              ((f2 (NonEmpty-value0 v4)) ((((rt:record-ref _1 (scm:string->symbol "foldr")) f2) b3) (NonEmpty-value1 v4))))))))])
          (scm:list (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid3)
            (scm:let ([foldMapWithIndex14 ((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldMapWithIndex")) dictMonoid3)])
              (scm:lambda (f5)
                (scm:lambda (v6)
                  (((rt:record-ref ((rt:record-ref dictMonoid3 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) ((f5 Data.Maybe.Nothing) (NonEmpty-value0 v6))) ((foldMapWithIndex14 (scm:lambda (x7)
                    (f5 (Data.Maybe.Just x7)))) (NonEmpty-value1 v6)))))))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f3)
            (scm:lambda (b4)
              (scm:lambda (v5)
                ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldlWithIndex")) (scm:lambda (x6)
                  (f3 (Data.Maybe.Just x6)))) (((f3 Data.Maybe.Nothing) b4) (NonEmpty-value0 v5))) (NonEmpty-value1 v5)))))) (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f3)
            (scm:lambda (b4)
              (scm:lambda (v5)
                (((f3 Data.Maybe.Nothing) (NonEmpty-value0 v5)) ((((rt:record-ref dictFoldableWithIndex0 (scm:string->symbol "foldrWithIndex")) (scm:lambda (x6)
                  (f3 (Data.Maybe.Just x6)))) b4) (NonEmpty-value1 v5))))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
            foldableNonEmpty12))))))

  (scm:define traversableNonEmpty
    (scm:lambda (dictTraversable0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversable0 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorNonEmpty12 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f2)
          (scm:lambda (m3)
            (NonEmpty* (f2 (NonEmpty-value0 m3)) (((rt:record-ref _1 (scm:string->symbol "map")) f2) (NonEmpty-value1 m3)))))))]
         [_3 ((rt:record-ref dictTraversable0 (scm:string->symbol "Foldable1")) (scm:quote undefined))]
         [foldableNonEmpty14 (scm:list (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid4)
          (scm:let ([foldMap15 ((rt:record-ref _3 (scm:string->symbol "foldMap")) dictMonoid4)])
            (scm:lambda (f6)
              (scm:lambda (v7)
                (((rt:record-ref ((rt:record-ref dictMonoid4 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f6 (NonEmpty-value0 v7))) ((foldMap15 f6) (NonEmpty-value1 v7)))))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f4)
          (scm:lambda (b5)
            (scm:lambda (v6)
              ((((rt:record-ref _3 (scm:string->symbol "foldl")) f4) ((f4 b5) (NonEmpty-value0 v6))) (NonEmpty-value1 v6)))))) (scm:cons (scm:string->symbol "foldr") (scm:lambda (f4)
          (scm:lambda (b5)
            (scm:lambda (v6)
              ((f4 (NonEmpty-value0 v6)) ((((rt:record-ref _3 (scm:string->symbol "foldr")) f4) b5) (NonEmpty-value1 v6))))))))])
          (scm:list (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative5)
            (scm:let*
              ([Apply06 ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined))]
               [sequence17 ((rt:record-ref dictTraversable0 (scm:string->symbol "sequence")) dictApplicative5)])
                (scm:lambda (v8)
                  (((rt:record-ref Apply06 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref Apply06 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) NonEmpty) (NonEmpty-value0 v8))) (sequence17 (NonEmpty-value1 v8))))))) (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative5)
            (scm:let*
              ([Apply06 ((rt:record-ref dictApplicative5 (scm:string->symbol "Apply0")) (scm:quote undefined))]
               [traverse17 ((rt:record-ref dictTraversable0 (scm:string->symbol "traverse")) dictApplicative5)])
                (scm:lambda (f8)
                  (scm:lambda (v9)
                    (((rt:record-ref Apply06 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref Apply06 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) NonEmpty) (f8 (NonEmpty-value0 v9)))) ((traverse17 f8) (NonEmpty-value1 v9)))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorNonEmpty12)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
            foldableNonEmpty14))))))

  (scm:define traversableWithIndexNonEmpty
    (scm:lambda (dictTraversableWithIndex0)
      (scm:let*
        ([_1 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FunctorWithIndex0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
         [functorWithIndex13 (scm:let ([functorNonEmpty13 (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f3)
          (scm:lambda (m4)
            (NonEmpty* (f3 (NonEmpty-value0 m4)) (((rt:record-ref _2 (scm:string->symbol "map")) f3) (NonEmpty-value1 m4)))))))])
          (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f4)
            (scm:lambda (v5)
              (NonEmpty* ((f4 Data.Maybe.Nothing) (NonEmpty-value0 v5)) (((rt:record-ref _1 (scm:string->symbol "mapWithIndex")) (scm:lambda (x6)
                (f4 (Data.Maybe.Just x6)))) (NonEmpty-value1 v5)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
            functorNonEmpty13))))]
         [foldableWithIndexNonEmpty14 (foldableWithIndexNonEmpty ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "FoldableWithIndex1")) (scm:quote undefined)))]
         [traversableNonEmpty15 (traversableNonEmpty ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "Traversable2")) (scm:quote undefined)))])
          (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative6)
            (scm:let*
              ([Apply07 ((rt:record-ref dictApplicative6 (scm:string->symbol "Apply0")) (scm:quote undefined))]
               [traverseWithIndex18 ((rt:record-ref dictTraversableWithIndex0 (scm:string->symbol "traverseWithIndex")) dictApplicative6)])
                (scm:lambda (f9)
                  (scm:lambda (v10)
                    (((rt:record-ref Apply07 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref Apply07 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) NonEmpty) ((f9 Data.Maybe.Nothing) (NonEmpty-value0 v10)))) ((traverseWithIndex18 (scm:lambda (x11)
                      (f9 (Data.Maybe.Just x11)))) (NonEmpty-value1 v10)))))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
            functorWithIndex13)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
            foldableWithIndexNonEmpty14)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
            traversableNonEmpty15))))))

  (scm:define foldable1NonEmpty
    (scm:lambda (dictFoldable0)
      (scm:let ([foldableNonEmpty11 (scm:list (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid1)
        (scm:let ([foldMap12 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid1)])
          (scm:lambda (f3)
            (scm:lambda (v4)
              (((rt:record-ref ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f3 (NonEmpty-value0 v4))) ((foldMap12 f3) (NonEmpty-value1 v4)))))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f1)
        (scm:lambda (b2)
          (scm:lambda (v3)
            ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) f1) ((f1 b2) (NonEmpty-value0 v3))) (NonEmpty-value1 v3)))))) (scm:cons (scm:string->symbol "foldr") (scm:lambda (f1)
        (scm:lambda (b2)
          (scm:lambda (v3)
            ((f1 (NonEmpty-value0 v3)) ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) f1) b2) (NonEmpty-value1 v3))))))))])
        (scm:list (scm:cons (scm:string->symbol "foldMap1") (scm:lambda (dictSemigroup2)
          (scm:lambda (f3)
            (scm:lambda (v4)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (s5)
                (scm:lambda (a16)
                  (((rt:record-ref dictSemigroup2 (scm:string->symbol "append")) s5) (f3 a16))))) (f3 (NonEmpty-value0 v4))) (NonEmpty-value1 v4)))))) (scm:cons (scm:string->symbol "foldr1") (scm:lambda (f2)
          (scm:lambda (v3)
            (scm:let*
              ([_4 (f2 (NonEmpty-value0 v3))]
               [_5 ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) (scm:lambda (a15)
                (scm:let ([_6 (f2 a15)])
                  (scm:lambda (x7)
                    (Data.Maybe.Just (scm:cond
                      [(Data.Maybe.Nothing? x7) a15]
                      [(Data.Maybe.Just? x7) (_6 (Data.Maybe.Just-value0 x7))]
                      [scm:else (rt:fail)])))))) Data.Maybe.Nothing) (NonEmpty-value1 v3))])
                (scm:cond
                  [(Data.Maybe.Nothing? _5) (NonEmpty-value0 v3)]
                  [(Data.Maybe.Just? _5) (_4 (Data.Maybe.Just-value0 _5))]
                  [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldl1") (scm:lambda (f2)
          (scm:lambda (v3)
            ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) f2) (NonEmpty-value0 v3)) (NonEmpty-value1 v3))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
          foldableNonEmpty11))))))

  (scm:define foldl1
    (scm:lambda (dictFoldable0)
      (rt:record-ref (foldable1NonEmpty dictFoldable0) (scm:string->symbol "foldl1"))))

  (scm:define eqNonEmpty
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq1)
        (scm:let ([eq112 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq1)])
          (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x3)
            (scm:lambda (y4)
              (scm:and (((rt:record-ref dictEq1 (scm:string->symbol "eq")) (NonEmpty-value0 x3)) (NonEmpty-value0 y4)) ((eq112 (NonEmpty-value1 x3)) (NonEmpty-value1 y4)))))))))))

  (scm:define ordNonEmpty
    (scm:lambda (dictOrd10)
      (scm:let ([_1 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))])
        (scm:lambda (dictOrd2)
          (scm:let*
            ([compare113 ((rt:record-ref dictOrd10 (scm:string->symbol "compare1")) dictOrd2)]
             [_4 ((rt:record-ref dictOrd2 (scm:string->symbol "Eq0")) (scm:quote undefined))]
             [eq115 ((rt:record-ref _1 (scm:string->symbol "eq1")) _4)]
             [eqNonEmpty26 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x6)
              (scm:lambda (y7)
                (scm:and (((rt:record-ref _4 (scm:string->symbol "eq")) (NonEmpty-value0 x6)) (NonEmpty-value0 y7)) ((eq115 (NonEmpty-value1 x6)) (NonEmpty-value1 y7)))))))])
              (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x7)
                (scm:lambda (y8)
                  (scm:let ([v9 (((rt:record-ref dictOrd2 (scm:string->symbol "compare")) (NonEmpty-value0 x7)) (NonEmpty-value0 y8))])
                    (scm:cond
                      [(Data.Ordering.LT? v9) Data.Ordering.LT]
                      [(Data.Ordering.GT? v9) Data.Ordering.GT]
                      [scm:else ((compare113 (NonEmpty-value1 x7)) (NonEmpty-value1 y8))]))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                eqNonEmpty26))))))))

  (scm:define eq1NonEmpty
    (scm:lambda (dictEq10)
      (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq1)
        (scm:let ([eq112 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq1)])
          (scm:lambda (x3)
            (scm:lambda (y4)
              (scm:and (((rt:record-ref dictEq1 (scm:string->symbol "eq")) (NonEmpty-value0 x3)) (NonEmpty-value0 y4)) ((eq112 (NonEmpty-value1 x3)) (NonEmpty-value1 y4)))))))))))

  (scm:define ord1NonEmpty
    (scm:lambda (dictOrd10)
      (scm:let*
        ([ordNonEmpty11 (ordNonEmpty dictOrd10)]
         [_2 ((rt:record-ref dictOrd10 (scm:string->symbol "Eq10")) (scm:quote undefined))]
         [eq1NonEmpty13 (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq3)
          (scm:let ([eq114 ((rt:record-ref _2 (scm:string->symbol "eq1")) dictEq3)])
            (scm:lambda (x5)
              (scm:lambda (y6)
                (scm:and (((rt:record-ref dictEq3 (scm:string->symbol "eq")) (NonEmpty-value0 x5)) (NonEmpty-value0 y6)) ((eq114 (NonEmpty-value1 x5)) (NonEmpty-value1 y6)))))))))])
          (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd4)
            (rt:record-ref (ordNonEmpty11 dictOrd4) (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
            eq1NonEmpty13)))))))
