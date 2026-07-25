#!r6rs
#!chezscheme
(library
  (Data.Foldable lib)
  (export
    Append
    Append*
    Append-value0
    Append-value1
    Append?
    Empty
    Empty?
    Node
    Node-value0
    Node?
    all
    and
    any
    elem
    find
    findMap
    fold
    foldM
    foldMap
    foldMapDefaultL
    foldMapDefaultR
    foldableAdditive
    foldableApp
    foldableArray
    foldableCompose
    foldableConj
    foldableConst
    foldableCoproduct
    foldableDisj
    foldableDual
    foldableEither
    foldableFirst
    foldableFreeMonoidTree
    foldableIdentity
    foldableLast
    foldableMaybe
    foldableMultiplicative
    foldableProduct
    foldableTuple
    foldl
    foldlArray
    foldlDefault
    foldr
    foldrArray
    foldrDefault
    for_
    identity
    indexl
    indexr
    intercalate
    length
    lookup
    maximum
    maximumBy
    minimum
    minimumBy
    monoidEndo
    monoidFreeMonoidTree
    notElem
    null
    oneOf
    oneOfMap
    or
    product
    semigroupFreeMonoidTree
    sequence_
    sum
    surround
    surroundMap
    traverse_)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Apply lib) Control.Apply.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Maybe.First lib) Data.Maybe.First.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.)
    (Data.Foldable foreign))

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

  (scm:define Empty
    (scm:quote Empty))

  (scm:define Empty?
    (scm:lambda (v)
      (scm:eq? (scm:quote Empty) v)))

  (scm:define-record-type (Node$ Node Node?)
    (scm:fields (scm:immutable value0 Node-value0)))

  (scm:define-record-type (Append$ Append* Append?)
    (scm:fields (scm:immutable value0 Append-value0) (scm:immutable value1 Append-value1)))

  (scm:define Append
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Append* value0 value1))))

  (scm:define semigroupFreeMonoidTree
    (scm:list (scm:cons (scm:string->symbol "append") Append)))

  (scm:define monoidFreeMonoidTree
    (scm:list (scm:cons (scm:string->symbol "mempty") Empty) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
      semigroupFreeMonoidTree))))

  (scm:define foldr
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldr"))))

  (scm:define indexr
    (scm:lambda (dictFoldable0)
      (scm:lambda (idx1)
        (scm:let ([_2 (((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) (scm:lambda (a2)
          (scm:lambda (cursor3)
            (scm:cond
              [(Data.Maybe.Just? (rt:record-ref cursor3 (scm:string->symbol "elem"))) cursor3]
              [(scm:fx=? (rt:record-ref cursor3 (scm:string->symbol "pos")) idx1) (scm:list (scm:cons (scm:string->symbol "elem") (Data.Maybe.Just a2)) (scm:cons (scm:string->symbol "pos") (rt:record-ref cursor3 (scm:string->symbol "pos"))))]
              [scm:else (scm:list (scm:cons (scm:string->symbol "pos") (scm:fx+ (rt:record-ref cursor3 (scm:string->symbol "pos")) 1)) (scm:cons (scm:string->symbol "elem") (rt:record-ref cursor3 (scm:string->symbol "elem"))))])))) (scm:list (scm:cons (scm:string->symbol "elem") Data.Maybe.Nothing) (scm:cons (scm:string->symbol "pos") 0)))])
          (scm:lambda (x3)
            (rt:record-ref (_2 x3) (scm:string->symbol "elem")))))))

  (scm:define null
    (scm:lambda (dictFoldable0)
      (((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) (scm:lambda (v1)
        (scm:lambda (v12)
          #f))) #t)))

  (scm:define oneOf
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictPlus1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) (rt:record-ref ((rt:record-ref dictPlus1 (scm:string->symbol "Alt0")) (scm:quote undefined)) (scm:string->symbol "alt"))) (rt:record-ref dictPlus1 (scm:string->symbol "empty"))))))

  (scm:define oneOfMap
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictPlus1)
        (scm:let ([empty2 (rt:record-ref dictPlus1 (scm:string->symbol "empty"))])
          (scm:lambda (f3)
            (((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) (scm:lambda (x4)
              ((rt:record-ref ((rt:record-ref dictPlus1 (scm:string->symbol "Alt0")) (scm:quote undefined)) (scm:string->symbol "alt")) (f3 x4)))) empty2))))))

  (scm:define traverse_
    (scm:lambda (dictApplicative0)
      (scm:let ([_1 ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined))])
        (scm:lambda (dictFoldable2)
          (scm:lambda (f3)
            (((rt:record-ref dictFoldable2 (scm:string->symbol "foldr")) (scm:lambda (x4)
              (scm:let ([_5 (f3 x4)])
                (scm:lambda (b6)
                  (((rt:record-ref _1 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref _1 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (v7)
                    Control.Apply.identity)) _5)) b6))))) ((rt:record-ref dictApplicative0 (scm:string->symbol "pure")) Data.Unit.unit)))))))

  (scm:define for_
    (scm:lambda (dictApplicative0)
      (scm:let ([traverse_11 (traverse_ dictApplicative0)])
        (scm:lambda (dictFoldable2)
          (scm:let ([_3 (traverse_11 dictFoldable2)])
            (scm:lambda (b4)
              (scm:lambda (a5)
                ((_3 a5) b4))))))))

  (scm:define sequence_
    (scm:lambda (dictApplicative0)
      (scm:let ([traverse_11 (traverse_ dictApplicative0)])
        (scm:lambda (dictFoldable2)
          ((traverse_11 dictFoldable2) identity)))))

  (scm:define foldl
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldl"))))

  (scm:define indexl
    (scm:lambda (dictFoldable0)
      (scm:lambda (idx1)
        (scm:let ([_2 (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (cursor2)
          (scm:lambda (a3)
            (scm:cond
              [(Data.Maybe.Just? (rt:record-ref cursor2 (scm:string->symbol "elem"))) cursor2]
              [(scm:fx=? (rt:record-ref cursor2 (scm:string->symbol "pos")) idx1) (scm:list (scm:cons (scm:string->symbol "elem") (Data.Maybe.Just a3)) (scm:cons (scm:string->symbol "pos") (rt:record-ref cursor2 (scm:string->symbol "pos"))))]
              [scm:else (scm:list (scm:cons (scm:string->symbol "pos") (scm:fx+ (rt:record-ref cursor2 (scm:string->symbol "pos")) 1)) (scm:cons (scm:string->symbol "elem") (rt:record-ref cursor2 (scm:string->symbol "elem"))))])))) (scm:list (scm:cons (scm:string->symbol "elem") Data.Maybe.Nothing) (scm:cons (scm:string->symbol "pos") 0)))])
          (scm:lambda (x3)
            (rt:record-ref (_2 x3) (scm:string->symbol "elem")))))))

  (scm:define intercalate
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictMonoid1)
        (scm:let*
          ([_2 ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
           [mempty3 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))])
            (scm:lambda (sep4)
              (scm:lambda (xs5)
                (rt:record-ref ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (v6)
                  (scm:lambda (v17)
                    (scm:cond
                      [(rt:record-ref v6 (scm:string->symbol "init")) (scm:list (scm:cons (scm:string->symbol "init") #f) (scm:cons (scm:string->symbol "acc") v17))]
                      [scm:else (scm:list (scm:cons (scm:string->symbol "init") #f) (scm:cons (scm:string->symbol "acc") (((rt:record-ref _2 (scm:string->symbol "append")) (rt:record-ref v6 (scm:string->symbol "acc"))) (((rt:record-ref _2 (scm:string->symbol "append")) sep4) v17))))])))) (scm:list (scm:cons (scm:string->symbol "init") #t) (scm:cons (scm:string->symbol "acc") mempty3))) xs5) (scm:string->symbol "acc"))))))))

  (scm:define length
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictSemiring1)
        (scm:let ([one2 (rt:record-ref dictSemiring1 (scm:string->symbol "one"))])
          (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (c3)
            (scm:lambda (v4)
              (((rt:record-ref dictSemiring1 (scm:string->symbol "add")) one2) c3)))) (rt:record-ref dictSemiring1 (scm:string->symbol "zero")))))))

  (scm:define maximumBy
    (scm:lambda (dictFoldable0)
      (scm:lambda (cmp1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v2) (Data.Maybe.Just v13)]
              [(Data.Maybe.Just? v2) (Data.Maybe.Just (scm:cond
                [(Data.Ordering.GT? ((cmp1 (Data.Maybe.Just-value0 v2)) v13)) (Data.Maybe.Just-value0 v2)]
                [scm:else v13]))]
              [scm:else (rt:fail)])))) Data.Maybe.Nothing))))

  (scm:define maximum
    (scm:lambda (dictOrd0)
      (scm:lambda (dictFoldable1)
        (((rt:record-ref dictFoldable1 (scm:string->symbol "foldl")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v2) (Data.Maybe.Just v13)]
              [(Data.Maybe.Just? v2) (Data.Maybe.Just (scm:cond
                [(Data.Ordering.GT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Data.Maybe.Just-value0 v2)) v13)) (Data.Maybe.Just-value0 v2)]
                [scm:else v13]))]
              [scm:else (rt:fail)])))) Data.Maybe.Nothing))))

  (scm:define minimumBy
    (scm:lambda (dictFoldable0)
      (scm:lambda (cmp1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v2) (Data.Maybe.Just v13)]
              [(Data.Maybe.Just? v2) (Data.Maybe.Just (scm:cond
                [(Data.Ordering.LT? ((cmp1 (Data.Maybe.Just-value0 v2)) v13)) (Data.Maybe.Just-value0 v2)]
                [scm:else v13]))]
              [scm:else (rt:fail)])))) Data.Maybe.Nothing))))

  (scm:define minimum
    (scm:lambda (dictOrd0)
      (scm:lambda (dictFoldable1)
        (((rt:record-ref dictFoldable1 (scm:string->symbol "foldl")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v2) (Data.Maybe.Just v13)]
              [(Data.Maybe.Just? v2) (Data.Maybe.Just (scm:cond
                [(Data.Ordering.LT? (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Data.Maybe.Just-value0 v2)) v13)) (Data.Maybe.Just-value0 v2)]
                [scm:else v13]))]
              [scm:else (rt:fail)])))) Data.Maybe.Nothing))))

  (scm:define product
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictSemiring1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (rt:record-ref dictSemiring1 (scm:string->symbol "mul"))) (rt:record-ref dictSemiring1 (scm:string->symbol "one"))))))

  (scm:define sum
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictSemiring1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (rt:record-ref dictSemiring1 (scm:string->symbol "add"))) (rt:record-ref dictSemiring1 (scm:string->symbol "zero"))))))

  (scm:define foldableTuple
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 (Data.Tuple.Tuple-value1 v2)) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 z1) (Data.Tuple.Tuple-value1 v2)))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 (Data.Tuple.Tuple-value1 v2))))))))

  (scm:define foldableMultiplicative
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 v2) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 z1) v2))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2)))))))

  (scm:define foldableMaybe
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Maybe.Nothing? v22) v11]
            [(Data.Maybe.Just? v22) ((v0 (Data.Maybe.Just-value0 v22)) v11)]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Maybe.Nothing? v22) v11]
            [(Data.Maybe.Just? v22) ((v0 v11) (Data.Maybe.Just-value0 v22))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v13) mempty1]
              [(Data.Maybe.Just? v13) (v2 (Data.Maybe.Just-value0 v13))]
              [scm:else (rt:fail)]))))))))

  (scm:define foldableIdentity
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 v2) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 z1) v2))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2)))))))

  (scm:define foldableEither
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Either.Left? v22) v11]
            [(Data.Either.Right? v22) ((v0 (Data.Either.Right-value0 v22)) v11)]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Data.Either.Left? v22) v11]
            [(Data.Either.Right? v22) ((v0 v11) (Data.Either.Right-value0 v22))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Either.Left? v13) mempty1]
              [(Data.Either.Right? v13) (v2 (Data.Either.Right-value0 v13))]
              [scm:else (rt:fail)]))))))))

  (scm:define foldableDual
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 v2) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 z1) v2))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2)))))))

  (scm:define foldableDisj
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 v2) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 z1) v2))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2)))))))

  (scm:define foldableConst
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (v0)
      (scm:lambda (z1)
        (scm:lambda (v12)
          z1)))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (v0)
      (scm:lambda (z1)
        (scm:lambda (v12)
          z1)))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            mempty1)))))))

  (scm:define foldableConj
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 v2) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 z1) v2))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2)))))))

  (scm:define foldableAdditive
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 v2) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          ((f0 z1) v2))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2)))))))

  (scm:define foldMapDefaultR
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictMonoid1)
        (scm:let ([mempty2 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))])
          (scm:lambda (f3)
            (((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) (scm:lambda (x4)
              (scm:lambda (acc5)
                (((rt:record-ref ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f3 x4)) acc5)))) mempty2))))))

  (rt:define-lazy $lazy-foldableArray "foldableArray" "Data.Foldable"
    (scm:list (scm:cons (scm:string->symbol "foldr") foldrArray) (scm:cons (scm:string->symbol "foldl") foldlArray) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (f2)
          (((rt:record-ref ($lazy-foldableArray) (scm:string->symbol "foldr")) (scm:lambda (x3)
            (scm:lambda (acc4)
              (((rt:record-ref ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f2 x3)) acc4)))) mempty1)))))))

  (scm:define foldableArray
    ($lazy-foldableArray))

  (rt:define-lazy $lazy-foldableFreeMonoidTree "foldableFreeMonoidTree" "Data.Foldable"
    (scm:list (scm:cons (scm:string->symbol "foldl") (scm:lambda (fn0)
      (scm:letrec ([go1 (scm:lambda (acc2)
        (scm:lambda (lhs3)
          (scm:lambda (rhs4)
            (scm:cond
              [(Node? lhs3) (((go1 ((fn0 acc2) (Node-value0 lhs3))) rhs4) Empty)]
              [(Append? lhs3) (scm:cond
                [(Empty? (Append-value1 lhs3)) (((go1 acc2) (Append-value0 lhs3)) rhs4)]
                [(Empty? rhs4) (((go1 acc2) (Append-value0 lhs3)) (Append-value1 lhs3))]
                [scm:else (((go1 acc2) (Append-value0 lhs3)) (Append* (Append-value1 lhs3) rhs4))])]
              [(Empty? lhs3) (scm:cond
                [(Empty? rhs4) acc2]
                [scm:else (((go1 acc2) rhs4) Empty)])]
              [scm:else (rt:fail)]))))])
        (scm:lambda (a2)
          (scm:lambda (b3)
            (((go1 a2) b3) Empty)))))) (scm:cons (scm:string->symbol "foldr") (scm:lambda (fn0)
      (scm:letrec ([go1 (scm:lambda (acc2)
        (scm:lambda (lhs3)
          (scm:lambda (rhs4)
            (scm:cond
              [(Node? rhs4) (((go1 ((fn0 (Node-value0 rhs4)) acc2)) Empty) lhs3)]
              [(Append? rhs4) (scm:cond
                [(Empty? (Append-value0 rhs4)) (((go1 acc2) lhs3) (Append-value1 rhs4))]
                [(Empty? lhs3) (((go1 acc2) (Append-value0 rhs4)) (Append-value1 rhs4))]
                [scm:else (((go1 acc2) (Append* lhs3 (Append-value0 rhs4))) (Append-value1 rhs4))])]
              [(Empty? rhs4) (scm:cond
                [(Empty? lhs3) acc2]
                [scm:else (((go1 acc2) Empty) lhs3)])]
              [scm:else (rt:fail)]))))])
        (scm:lambda (a2)
          (scm:lambda (b3)
            (((go1 a2) Empty) b3)))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (f2)
          (((rt:record-ref ($lazy-foldableFreeMonoidTree) (scm:string->symbol "foldr")) (scm:lambda (x3)
            (scm:lambda (acc4)
              (((rt:record-ref ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) (f2 x3)) acc4)))) mempty1)))))))

  (scm:define foldableFreeMonoidTree
    ($lazy-foldableFreeMonoidTree))

  (scm:define foldMapDefaultL
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictMonoid1)
        (scm:let ([mempty2 (rt:record-ref dictMonoid1 (scm:string->symbol "mempty"))])
          (scm:lambda (f3)
            (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (acc4)
              (scm:lambda (x5)
                (((rt:record-ref ((rt:record-ref dictMonoid1 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) acc4) (f3 x5))))) mempty2))))))

  (scm:define foldMap
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldMap"))))

  (scm:define foldableApp
    (scm:lambda (dictFoldable0)
      (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f1)
        (scm:lambda (i2)
          (scm:lambda (v3)
            ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) f1) i2) v3))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f1)
        (scm:lambda (i2)
          (scm:lambda (v3)
            ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) f1) i2) v3))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid1)
        ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid1))))))

  (scm:define foldableCompose
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictFoldable11)
        (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f2)
          (scm:lambda (i3)
            (scm:lambda (v4)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) (scm:let ([_5 ((rt:record-ref dictFoldable11 (scm:string->symbol "foldr")) f2)])
                (scm:lambda (b6)
                  (scm:lambda (a7)
                    ((_5 a7) b6))))) i3) v4))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f2)
          (scm:lambda (i3)
            (scm:lambda (v4)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) ((rt:record-ref dictFoldable11 (scm:string->symbol "foldl")) f2)) i3) v4))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid2)
          (scm:let*
            ([foldMap43 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid2)]
             [foldMap54 ((rt:record-ref dictFoldable11 (scm:string->symbol "foldMap")) dictMonoid2)])
              (scm:lambda (f5)
                (scm:lambda (v6)
                  ((foldMap43 (foldMap54 f5)) v6))))))))))

  (scm:define foldableCoproduct
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictFoldable11)
        (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f2)
          (scm:lambda (z3)
            (scm:let*
              ([_4 (((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) f2) z3)]
               [_5 (((rt:record-ref dictFoldable11 (scm:string->symbol "foldr")) f2) z3)])
                (scm:lambda (v26)
                  (scm:cond
                    [(Data.Either.Left? v26) (_4 (Data.Either.Left-value0 v26))]
                    [(Data.Either.Right? v26) (_5 (Data.Either.Right-value0 v26))]
                    [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f2)
          (scm:lambda (z3)
            (scm:let*
              ([_4 (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) f2) z3)]
               [_5 (((rt:record-ref dictFoldable11 (scm:string->symbol "foldl")) f2) z3)])
                (scm:lambda (v26)
                  (scm:cond
                    [(Data.Either.Left? v26) (_4 (Data.Either.Left-value0 v26))]
                    [(Data.Either.Right? v26) (_5 (Data.Either.Right-value0 v26))]
                    [scm:else (rt:fail)])))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid2)
          (scm:let*
            ([foldMap43 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid2)]
             [foldMap54 ((rt:record-ref dictFoldable11 (scm:string->symbol "foldMap")) dictMonoid2)])
              (scm:lambda (f5)
                (scm:let*
                  ([_6 (foldMap43 f5)]
                   [_7 (foldMap54 f5)])
                    (scm:lambda (v28)
                      (scm:cond
                        [(Data.Either.Left? v28) (_6 (Data.Either.Left-value0 v28))]
                        [(Data.Either.Right? v28) (_7 (Data.Either.Right-value0 v28))]
                        [scm:else (rt:fail)])))))))))))

  (scm:define foldableFirst
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (scm:cond
            [(Data.Maybe.Nothing? v2) z1]
            [(Data.Maybe.Just? v2) ((f0 (Data.Maybe.Just-value0 v2)) z1)]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (scm:cond
            [(Data.Maybe.Nothing? v2) z1]
            [(Data.Maybe.Just? v2) ((f0 z1) (Data.Maybe.Just-value0 v2))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v13) mempty1]
              [(Data.Maybe.Just? v13) (v2 (Data.Maybe.Just-value0 v13))]
              [scm:else (rt:fail)]))))))))

  (scm:define foldableLast
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (scm:cond
            [(Data.Maybe.Nothing? v2) z1]
            [(Data.Maybe.Just? v2) ((f0 (Data.Maybe.Just-value0 v2)) z1)]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (v2)
          (scm:cond
            [(Data.Maybe.Nothing? v2) z1]
            [(Data.Maybe.Just? v2) ((f0 z1) (Data.Maybe.Just-value0 v2))]
            [scm:else (rt:fail)]))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v13) mempty1]
              [(Data.Maybe.Just? v13) (v2 (Data.Maybe.Just-value0 v13))]
              [scm:else (rt:fail)]))))))))

  (scm:define foldableProduct
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictFoldable11)
        (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f2)
          (scm:lambda (z3)
            (scm:lambda (v4)
              ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldr")) f2) ((((rt:record-ref dictFoldable11 (scm:string->symbol "foldr")) f2) z3) (Data.Tuple.Tuple-value1 v4))) (Data.Tuple.Tuple-value0 v4)))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f2)
          (scm:lambda (z3)
            (scm:lambda (v4)
              ((((rt:record-ref dictFoldable11 (scm:string->symbol "foldl")) f2) ((((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) f2) z3) (Data.Tuple.Tuple-value0 v4))) (Data.Tuple.Tuple-value1 v4)))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid2)
          (scm:let*
            ([foldMap43 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid2)]
             [foldMap54 ((rt:record-ref dictFoldable11 (scm:string->symbol "foldMap")) dictMonoid2)])
              (scm:lambda (f5)
                (scm:lambda (v6)
                  (((rt:record-ref ((rt:record-ref dictMonoid2 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) ((foldMap43 f5) (Data.Tuple.Tuple-value0 v6))) ((foldMap54 f5) (Data.Tuple.Tuple-value1 v6))))))))))))

  (scm:define foldlDefault
    (scm:lambda (dictFoldable0)
      (scm:let ([foldMap21 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) monoidFreeMonoidTree)])
        (scm:lambda (c2)
          (scm:lambda (u3)
            (scm:lambda (xs4)
              ((((rt:record-ref foldableFreeMonoidTree (scm:string->symbol "foldl")) c2) u3) ((foldMap21 Node) xs4))))))))

  (scm:define foldrDefault
    (scm:lambda (dictFoldable0)
      (scm:let ([foldMap21 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) monoidFreeMonoidTree)])
        (scm:lambda (c2)
          (scm:lambda (u3)
            (scm:lambda (xs4)
              ((((rt:record-ref foldableFreeMonoidTree (scm:string->symbol "foldr")) c2) u3) ((foldMap21 Node) xs4))))))))

  (scm:define lookup
    (scm:lambda (dictFoldable0)
      (scm:let ([foldMap21 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) Data.Maybe.First.monoidFirst)])
        (scm:lambda (dictEq2)
          (scm:lambda (a3)
            (foldMap21 (scm:lambda (v4)
              (scm:cond
                [(((rt:record-ref dictEq2 (scm:string->symbol "eq")) a3) (Data.Tuple.Tuple-value0 v4)) (Data.Maybe.Just (Data.Tuple.Tuple-value1 v4))]
                [scm:else Data.Maybe.Nothing]))))))))

  (scm:define surroundMap
    (scm:lambda (dictFoldable0)
      (scm:let ([foldMap21 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) monoidEndo)])
        (scm:lambda (dictSemigroup2)
          (scm:lambda (d3)
            (scm:lambda (t4)
              (scm:lambda (f5)
                (((foldMap21 (scm:lambda (a6)
                  (scm:lambda (m7)
                    (((rt:record-ref dictSemigroup2 (scm:string->symbol "append")) d3) (((rt:record-ref dictSemigroup2 (scm:string->symbol "append")) (t4 a6)) m7))))) f5) d3))))))))

  (scm:define surround
    (scm:lambda (dictFoldable0)
      (scm:let ([surroundMap11 (surroundMap dictFoldable0)])
        (scm:lambda (dictSemigroup2)
          (scm:let ([surroundMap23 (surroundMap11 dictSemigroup2)])
            (scm:lambda (d4)
              ((surroundMap23 d4) identity)))))))

  (scm:define foldM
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictMonad1)
        (scm:lambda (f2)
          (scm:lambda (b03)
            (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (b4)
              (scm:lambda (a5)
                (((rt:record-ref ((rt:record-ref dictMonad1 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) b4) (scm:lambda (a6)
                  ((f2 a6) a5)))))) ((rt:record-ref ((rt:record-ref dictMonad1 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) b03)))))))

  (scm:define fold
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictMonoid1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) dictMonoid1) identity))))

  (scm:define findMap
    (scm:lambda (dictFoldable0)
      (scm:lambda (p1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(Data.Maybe.Nothing? v2) (p1 v13)]
              [scm:else v2])))) Data.Maybe.Nothing))))

  (scm:define find
    (scm:lambda (dictFoldable0)
      (scm:lambda (p1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldl")) (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:cond
              [(scm:and (Data.Maybe.Nothing? v2) (p1 v13)) (Data.Maybe.Just v13)]
              [scm:else v2])))) Data.Maybe.Nothing))))

  (scm:define any
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictHeytingAlgebra1)
        ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) (scm:let ([semigroupDisj12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "disj")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "ff"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupDisj12))))))))

  (scm:define elem
    (scm:lambda (dictFoldable0)
      (scm:let ([any11 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) (scm:let ([semigroupDisj11 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:or v1 v12)))))])
        (scm:list (scm:cons (scm:string->symbol "mempty") #f) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
          semigroupDisj11)))))])
        (scm:lambda (dictEq2)
          (scm:lambda (x3)
            (any11 ((rt:record-ref dictEq2 (scm:string->symbol "eq")) x3)))))))

  (scm:define notElem
    (scm:lambda (dictFoldable0)
      (scm:let ([any11 ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) (scm:let ([semigroupDisj11 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:or v1 v12)))))])
        (scm:list (scm:cons (scm:string->symbol "mempty") #f) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
          semigroupDisj11)))))])
        (scm:lambda (dictEq2)
          (scm:lambda (x3)
            (scm:let ([_4 (any11 ((rt:record-ref dictEq2 (scm:string->symbol "eq")) x3))])
              (scm:lambda (x5)
                (scm:not (_4 x5)))))))))

  (scm:define or
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictHeytingAlgebra1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) (scm:let ([semigroupDisj12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "disj")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "ff"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupDisj12))))) identity))))

  (scm:define all
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictHeytingAlgebra1)
        ((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) (scm:let ([semigroupConj12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "conj")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "tt"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupConj12))))))))

  (scm:define and
    (scm:lambda (dictFoldable0)
      (scm:lambda (dictHeytingAlgebra1)
        (((rt:record-ref dictFoldable0 (scm:string->symbol "foldMap")) (scm:let ([semigroupConj12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (((rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "conj")) v2) v13)))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (rt:record-ref dictHeytingAlgebra1 (scm:string->symbol "tt"))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupConj12))))) identity)))))
