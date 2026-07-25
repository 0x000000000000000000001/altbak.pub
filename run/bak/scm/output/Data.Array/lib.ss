#!r6rs
#!chezscheme
(library
  (Data.Array lib)
  (export
    _deleteAt
    _insertAt
    _updateAt
    all
    allImpl
    alterAt
    any
    anyImpl
    catMaybes
    concat
    concatMap
    cons
    delete
    deleteAt
    deleteBy
    difference
    drop
    dropEnd
    dropWhile
    elem
    elemIndex
    elemLastIndex
    filter
    filterA
    filterImpl
    find
    findIndex
    findIndexImpl
    findLastIndex
    findLastIndexImpl
    findMap
    findMapImpl
    fold
    foldM
    foldMap
    foldRecM
    foldl
    foldr
    fromFoldable
    fromFoldableImpl
    group
    groupAll
    groupAllBy
    groupBy
    head
    index
    indexImpl
    init
    insert
    insertAt
    insertBy
    intercalate
    intercalate1
    intersect
    intersectBy
    intersperse
    last
    length
    many
    mapMaybe
    mapWithIndex
    modifyAt
    modifyAtIndices
    notElem
    nub
    nubBy
    nubByEq
    nubEq
    null
    partition
    partitionImpl
    range
    rangeImpl
    replicate
    replicateImpl
    reverse
    scanl
    scanlImpl
    scanr
    scanrImpl
    singleton
    slice
    sliceImpl
    snoc
    some
    sort
    sortBy
    sortByImpl
    sortWith
    span
    splitAt
    tail
    take
    takeEnd
    takeWhile
    toUnfoldable
    transpose
    traverse_
    uncons
    unconsImpl
    union
    unionBy
    unsafeIndex
    unsafeIndexImpl
    unsnoc
    unzip
    updateAt
    updateAtIndices
    zip
    zipWith
    zipWithA
    zipWithImpl)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Bind lib) Control.Bind.)
    (prefix (Control.Monad.Rec.Class lib) Control.Monad.Rec.Class.)
    (prefix (Control.Monad.ST.Internal lib) Control.Monad.ST.Internal.)
    (prefix (Data.Array.ST lib) Data.Array.ST.)
    (prefix (Data.Array.ST.Iterator lib) Data.Array.ST.Iterator.)
    (prefix (Data.Foldable lib) Data.Foldable.)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.FunctorWithIndex lib) Data.FunctorWithIndex.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Ord lib) Data.Ord.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Semigroup lib) Data.Semigroup.)
    (prefix (Data.Traversable lib) Data.Traversable.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.)
    (Data.Array foreign))

  (scm:define traverse_
    (Data.Foldable.traverse_ Control.Monad.ST.Internal.applicativeST))

  (scm:define intercalate1
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [mempty2 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
          (scm:lambda (sep3)
            (scm:lambda (xs4)
              (rt:record-ref (((Data.Foldable.foldlArray (scm:lambda (v5)
                (scm:lambda (v16)
                  (scm:cond
                    [(rt:record-ref v5 (scm:string->symbol "init")) (scm:list (scm:cons (scm:string->symbol "init") #f) (scm:cons (scm:string->symbol "acc") v16))]
                    [scm:else (scm:list (scm:cons (scm:string->symbol "init") #f) (scm:cons (scm:string->symbol "acc") (((rt:record-ref _1 (scm:string->symbol "append")) (rt:record-ref v5 (scm:string->symbol "acc"))) (((rt:record-ref _1 (scm:string->symbol "append")) sep3) v16))))])))) (scm:list (scm:cons (scm:string->symbol "init") #t) (scm:cons (scm:string->symbol "acc") mempty2))) xs4) (scm:string->symbol "acc")))))))

  (scm:define zipWith
    (scm:lambda (_0)
      (scm:lambda (_1)
        (scm:lambda (_2)
          (zipWithImpl _0 _1 _2)))))

  (scm:define zipWithA
    (scm:lambda (dictApplicative0)
      (scm:let ([sequence11 (((rt:record-ref Data.Traversable.traversableArray (scm:string->symbol "traverse")) dictApplicative0) Data.Traversable.identity)])
        (scm:lambda (f2)
          (scm:lambda (xs3)
            (scm:lambda (ys4)
              (sequence11 (zipWithImpl f2 xs3 ys4))))))))

  (scm:define zip
    (zipWith Data.Tuple.Tuple))

  (scm:define updateAtIndices
    (scm:lambda (dictFoldable0)
      (scm:let ([traverse_11 (traverse_ dictFoldable0)])
        (scm:lambda (us2)
          (scm:lambda (xs3)
            (Control.Monad.ST.Internal.run (scm:lambda ()
              (scm:let*
                ([result4 (Data.Array.ST.thawImpl xs3)]
                 [_ (((traverse_11 (scm:lambda (v5)
                  (scm:let*
                    ([_6 (Data.Tuple.Tuple-value1 v5)]
                     [_7 (Data.Tuple.Tuple-value0 v5)])
                      (scm:lambda ()
                        (Data.Array.ST.pokeImpl _7 _6 result4))))) us2))])
                  ((scm:lambda ()
                    (Data.Array.ST.unsafeFreezeImpl result4)))))))))))

  (scm:define updateAt
    (scm:lambda (_0)
      (scm:lambda (_1)
        (scm:lambda (_2)
          (_updateAt Data.Maybe.Just Data.Maybe.Nothing _0 _1 _2)))))

  (scm:define unsafeIndex
    (scm:lambda (_)
      (scm:lambda (_1)
        (scm:lambda (_2)
          (rt:array-ref _1 _2)))))

  (scm:define uncons
    (scm:lambda (_0)
      (unconsImpl (scm:lambda (v1)
        Data.Maybe.Nothing) (scm:lambda (x1)
        (scm:lambda (xs2)
          (Data.Maybe.Just (scm:list (scm:cons (scm:string->symbol "head") x1) (scm:cons (scm:string->symbol "tail") xs2))))) _0)))

  (scm:define toUnfoldable
    (scm:lambda (dictUnfoldable0)
      (scm:lambda (xs1)
        (scm:let ([len2 (rt:array-length xs1)])
          (((rt:record-ref dictUnfoldable0 (scm:string->symbol "unfoldr")) (scm:lambda (i3)
            (scm:cond
              [(scm:fx<? i3 len2) (Data.Maybe.Just (Data.Tuple.Tuple* (rt:array-ref xs1 i3) (scm:fx+ i3 1)))]
              [scm:else Data.Maybe.Nothing]))) 0)))))

  (scm:define tail
    (scm:lambda (_0)
      (unconsImpl (scm:lambda (v1)
        Data.Maybe.Nothing) (scm:lambda (v1)
        (scm:lambda (xs2)
          (Data.Maybe.Just xs2))) _0)))

  (scm:define sortBy
    (scm:lambda (comp0)
      (scm:lambda (_1)
        (sortByImpl comp0 (scm:lambda (v2)
          (scm:cond
            [(Data.Ordering.GT? v2) 1]
            [(Data.Ordering.EQ? v2) 0]
            [(Data.Ordering.LT? v2) -1]
            [scm:else (rt:fail)])) _1))))

  (scm:define sortWith
    (scm:lambda (dictOrd0)
      (scm:lambda (f1)
        (sortBy (scm:lambda (x2)
          (scm:lambda (y3)
            (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (f1 x2)) (f1 y3))))))))

  (scm:define sort
    (scm:lambda (dictOrd0)
      (scm:let ([compare1 (rt:record-ref dictOrd0 (scm:string->symbol "compare"))])
        (scm:lambda (xs2)
          ((sortBy compare1) xs2)))))

  (scm:define snoc
    (scm:lambda (xs0)
      (scm:lambda (x1)
        (Control.Monad.ST.Internal.run (scm:let ([_2 (Data.Array.ST.push x1)])
          (scm:lambda ()
            (scm:let*
              ([result3 (Data.Array.ST.thawImpl xs0)]
               [_ ((_2 result3))])
                ((scm:lambda ()
                  (Data.Array.ST.unsafeFreezeImpl result3))))))))))

  (scm:define slice
    (scm:lambda (_0)
      (scm:lambda (_1)
        (scm:lambda (_2)
          (sliceImpl _0 _1 _2)))))

  (scm:define splitAt
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx<=? v0 0) (scm:list (scm:cons (scm:string->symbol "before") (rt:make-array)) (scm:cons (scm:string->symbol "after") v11))]
          [scm:else (scm:list (scm:cons (scm:string->symbol "before") (sliceImpl 0 v0 v11)) (scm:cons (scm:string->symbol "after") (sliceImpl v0 (rt:array-length v11) v11)))]))))

  (scm:define take
    (scm:lambda (n0)
      (scm:lambda (xs1)
        (scm:cond
          [(scm:fx<? n0 1) (rt:make-array)]
          [scm:else (sliceImpl 0 n0 xs1)]))))

  (scm:define singleton
    (scm:lambda (a0)
      (rt:make-array a0)))

  (scm:define scanr
    (scm:lambda (_0)
      (scm:lambda (_1)
        (scm:lambda (_2)
          (scanrImpl _0 _1 _2)))))

  (scm:define scanl
    (scm:lambda (_0)
      (scm:lambda (_1)
        (scm:lambda (_2)
          (scanlImpl _0 _1 _2)))))

  (scm:define replicate
    (scm:lambda (_0)
      (scm:lambda (_1)
        (replicateImpl _0 _1))))

  (scm:define range
    (scm:lambda (_0)
      (scm:lambda (_1)
        (rangeImpl _0 _1))))

  (scm:define partition
    (scm:lambda (_0)
      (scm:lambda (_1)
        (partitionImpl _0 _1))))

  (scm:define null
    (scm:lambda (xs0)
      (scm:fx=? (rt:array-length xs0) 0)))

  (scm:define modifyAtIndices
    (scm:lambda (dictFoldable0)
      (scm:let ([traverse_11 (traverse_ dictFoldable0)])
        (scm:lambda (is2)
          (scm:lambda (f3)
            (scm:lambda (xs4)
              (Control.Monad.ST.Internal.run (scm:lambda ()
                (scm:let*
                  ([result5 (Data.Array.ST.thawImpl xs4)]
                   [_ (((traverse_11 (scm:lambda (i6)
                    (((Data.Array.ST.modify i6) f3) result5))) is2))])
                    ((scm:lambda ()
                      (Data.Array.ST.unsafeFreezeImpl result5))))))))))))

  (scm:define mapWithIndex
    Data.FunctorWithIndex.mapWithIndexArray)

  (scm:define intersperse
    (scm:lambda (a0)
      (scm:lambda (arr1)
        (scm:let ([v2 (rt:array-length arr1)])
          (scm:cond
            [(scm:fx<? v2 2) arr1]
            [scm:else (Control.Monad.ST.Internal.run (scm:lambda ()
              (scm:let*
                ([out3 (Data.Array.ST.new)]
                 [_ (Data.Array.ST.pushImpl (rt:array-ref arr1 0) out3)]
                 [_ ((((Control.Monad.ST.Internal.for 1) v2) (scm:lambda (idx5)
                  (scm:lambda ()
                    (scm:let*
                      ([_ (Data.Array.ST.pushImpl a0 out3)]
                       [_7 (Data.Array.ST.pushImpl (rt:array-ref arr1 idx5) out3)])
                        Data.Unit.unit)))))])
                  ((scm:lambda ()
                    (Data.Array.ST.unsafeFreezeImpl out3))))))])))))

  (scm:define intercalate
    (scm:lambda (dictMonoid0)
      (intercalate1 dictMonoid0)))

  (scm:define insertAt
    (scm:lambda (_0)
      (scm:lambda (_1)
        (scm:lambda (_2)
          (_insertAt Data.Maybe.Just Data.Maybe.Nothing _0 _1 _2)))))

  (scm:define init
    (scm:lambda (xs0)
      (scm:cond
        [(scm:fx=? (rt:array-length xs0) 0) Data.Maybe.Nothing]
        [scm:else (Data.Maybe.Just (sliceImpl 0 (scm:fx- (rt:array-length xs0) 1) xs0))])))

  (scm:define index
    (scm:lambda (_0)
      (scm:lambda (_1)
        (indexImpl Data.Maybe.Just Data.Maybe.Nothing _0 _1))))

  (scm:define last
    (scm:lambda (xs0)
      (indexImpl Data.Maybe.Just Data.Maybe.Nothing xs0 (scm:fx- (rt:array-length xs0) 1))))

  (scm:define unsnoc
    (scm:lambda (xs0)
      (scm:cond
        [(scm:fx=? (rt:array-length xs0) 0) Data.Maybe.Nothing]
        [scm:else (scm:let ([_1 (indexImpl Data.Maybe.Just Data.Maybe.Nothing xs0 (scm:fx- (rt:array-length xs0) 1))])
          (scm:cond
            [(Data.Maybe.Just? _1) (Data.Maybe.Just (scm:list (scm:cons (scm:string->symbol "init") (sliceImpl 0 (scm:fx- (rt:array-length xs0) 1) xs0)) (scm:cons (scm:string->symbol "last") (Data.Maybe.Just-value0 _1))))]
            [scm:else Data.Maybe.Nothing]))])))

  (scm:define modifyAt
    (scm:lambda (i0)
      (scm:lambda (f1)
        (scm:lambda (xs2)
          (scm:let ([_3 (indexImpl Data.Maybe.Just Data.Maybe.Nothing xs2 i0)])
            (scm:cond
              [(Data.Maybe.Nothing? _3) Data.Maybe.Nothing]
              [(Data.Maybe.Just? _3) (_updateAt Data.Maybe.Just Data.Maybe.Nothing i0 (f1 (Data.Maybe.Just-value0 _3)) xs2)]
              [scm:else (rt:fail)]))))))

  (scm:define span
    (scm:lambda (p0)
      (scm:lambda (arr1)
        (scm:letrec ([go2 (scm:lambda (i3)
          (scm:let ([v4 (indexImpl Data.Maybe.Just Data.Maybe.Nothing arr1 i3)])
            (scm:cond
              [(Data.Maybe.Just? v4) (scm:cond
                [(p0 (Data.Maybe.Just-value0 v4)) (go2 (scm:fx+ i3 1))]
                [scm:else (Data.Maybe.Just i3)])]
              [(Data.Maybe.Nothing? v4) Data.Maybe.Nothing]
              [scm:else (rt:fail)])))])
          (scm:let ([breakIndex3 (go2 0)])
            (scm:cond
              [(Data.Maybe.Just? breakIndex3) (scm:cond
                [(scm:fx=? (Data.Maybe.Just-value0 breakIndex3) 0) (scm:list (scm:cons (scm:string->symbol "init") (rt:make-array)) (scm:cons (scm:string->symbol "rest") arr1))]
                [scm:else (scm:list (scm:cons (scm:string->symbol "init") (sliceImpl 0 (Data.Maybe.Just-value0 breakIndex3) arr1)) (scm:cons (scm:string->symbol "rest") (sliceImpl (Data.Maybe.Just-value0 breakIndex3) (rt:array-length arr1) arr1)))])]
              [(Data.Maybe.Nothing? breakIndex3) (scm:list (scm:cons (scm:string->symbol "init") arr1) (scm:cons (scm:string->symbol "rest") (rt:make-array)))]
              [scm:else (rt:fail)]))))))

  (scm:define takeWhile
    (scm:lambda (p0)
      (scm:lambda (xs1)
        (rt:record-ref ((span p0) xs1) (scm:string->symbol "init")))))

  (scm:define unzip
    (scm:lambda (xs0)
      (Control.Monad.ST.Internal.run (scm:lambda ()
        (scm:let*
          ([fsts1 (Data.Array.ST.new)]
           [snds2 (Data.Array.ST.new)]
           [_3 (scm:box 0)]
           [_ (((Data.Array.ST.Iterator.iterate (Data.Array.ST.Iterator.Iterator* (scm:lambda (v4)
            (indexImpl Data.Maybe.Just Data.Maybe.Nothing xs0 v4)) _3)) (scm:lambda (v4)
            (scm:let*
              ([_5 (Data.Tuple.Tuple-value0 v4)]
               [_6 (Data.Tuple.Tuple-value1 v4)])
                (scm:lambda ()
                  (scm:let*
                    ([_7 (Data.Array.ST.pushImpl _5 fsts1)]
                     [_8 (Data.Array.ST.pushImpl _6 snds2)])
                      Data.Unit.unit))))))]
           [fsts$p5 (Data.Array.ST.unsafeFreezeImpl fsts1)]
           [snds$p6 (Data.Array.ST.unsafeFreezeImpl snds2)])
            (Data.Tuple.Tuple* fsts$p5 snds$p6))))))

  (scm:define head
    (scm:lambda (xs0)
      (indexImpl Data.Maybe.Just Data.Maybe.Nothing xs0 0)))

  (scm:define nubBy
    (scm:lambda (comp0)
      (scm:lambda (xs1)
        (scm:let*
          ([indexedAndSorted2 ((sortBy (scm:lambda (x2)
            (scm:lambda (y3)
              ((comp0 (Data.Tuple.Tuple-value1 x2)) (Data.Tuple.Tuple-value1 y3))))) ((Data.FunctorWithIndex.mapWithIndexArray Data.Tuple.Tuple) xs1))]
           [v3 (indexImpl Data.Maybe.Just Data.Maybe.Nothing indexedAndSorted2 0)])
            (scm:cond
              [(Data.Maybe.Nothing? v3) (rt:make-array)]
              [(Data.Maybe.Just? v3) ((Data.Functor.arrayMap Data.Tuple.snd) (((sortWith Data.Ord.ordInt) Data.Tuple.fst) (Control.Monad.ST.Internal.run (scm:lambda ()
                (scm:let*
                  ([result4 (Data.Array.ST.unsafeThawImpl (rt:make-array (Data.Maybe.Just-value0 v3)))]
                   [_ (((Control.Monad.ST.Internal.foreach indexedAndSorted2) (scm:lambda (v15)
                    (scm:let ([_6 (Data.Tuple.Tuple-value1 v15)])
                      (scm:lambda ()
                        (scm:let*
                          ([_7 (Data.Array.ST.unsafeFreezeImpl result4)]
                           [_8 ((comp0 (scm:let ([_8 (indexImpl Data.Maybe.Just Data.Maybe.Nothing _7 (scm:fx- (rt:array-length _7) 1))])
                            (scm:cond
                              [(Data.Maybe.Just? _8) (Data.Tuple.Tuple-value1 (Data.Maybe.Just-value0 _8))]
                              [scm:else (rt:fail)]))) _6)])
                            ((scm:cond
                              [(scm:or (Data.Ordering.LT? _8) (scm:or (Data.Ordering.GT? _8) (scm:not (Data.Ordering.EQ? _8)))) (scm:lambda ()
                                (scm:let ([_9 (Data.Array.ST.pushImpl v15 result4)])
                                  Data.Unit.unit))]
                              [scm:else (scm:lambda ()
                                Data.Unit.unit)]))))))))])
                    ((scm:lambda ()
                      (Data.Array.ST.unsafeFreezeImpl result4))))))))]
              [scm:else (rt:fail)])))))

  (scm:define nub
    (scm:lambda (dictOrd0)
      (nubBy (rt:record-ref dictOrd0 (scm:string->symbol "compare")))))

  (scm:define groupBy
    (scm:lambda (op0)
      (scm:lambda (xs1)
        (Control.Monad.ST.Internal.run (scm:lambda ()
          (scm:let*
            ([result2 (Data.Array.ST.new)]
             [_3 (scm:box 0)]
             [iter4 (Data.Array.ST.Iterator.Iterator* (scm:lambda (v4)
              (indexImpl Data.Maybe.Just Data.Maybe.Nothing xs1 v4)) _3)]
             [_ (((Data.Array.ST.Iterator.iterate iter4) (scm:lambda (x5)
              (scm:lambda ()
                (scm:let*
                  ([sub16 (Data.Array.ST.new)]
                   [_ (Data.Array.ST.pushImpl x5 sub16)]
                   [_ ((((Data.Array.ST.Iterator.pushWhile (op0 x5)) iter4) sub16))]
                   [grp9 (Data.Array.ST.unsafeFreezeImpl sub16)]
                   [_10 (Data.Array.ST.pushImpl grp9 result2)])
                    Data.Unit.unit)))))])
              ((scm:lambda ()
                (Data.Array.ST.unsafeFreezeImpl result2)))))))))

  (scm:define groupAllBy
    (scm:lambda (cmp0)
      (scm:let ([_1 (groupBy (scm:lambda (x1)
        (scm:lambda (y2)
          (Data.Ordering.EQ? ((cmp0 x1) y2)))))])
        (scm:lambda (x2)
          (_1 ((sortBy cmp0) x2))))))

  (scm:define groupAll
    (scm:lambda (dictOrd0)
      (groupAllBy (rt:record-ref dictOrd0 (scm:string->symbol "compare")))))

  (scm:define group
    (scm:lambda (dictEq0)
      (scm:let ([eq21 (rt:record-ref dictEq0 (scm:string->symbol "eq"))])
        (scm:lambda (xs2)
          ((groupBy eq21) xs2)))))

  (scm:define fromFoldable
    (scm:lambda (dictFoldable0)
      (scm:let ([_1 (rt:record-ref dictFoldable0 (scm:string->symbol "foldr"))])
        (scm:lambda (_2)
          (fromFoldableImpl _1 _2)))))

  (scm:define foldr
    Data.Foldable.foldrArray)

  (scm:define foldl
    Data.Foldable.foldlArray)

  (scm:define transpose
    (scm:lambda (xs0)
      (scm:letrec ([go1 (scm:lambda (idx2)
        (scm:lambda (allArrays3)
          (scm:let ([v4 (((Data.Foldable.foldlArray (scm:lambda (acc4)
            (scm:lambda (nextArr5)
              (scm:let ([_6 (indexImpl Data.Maybe.Just Data.Maybe.Nothing nextArr5 idx2)])
                (scm:cond
                  [(Data.Maybe.Nothing? _6) acc4]
                  [(Data.Maybe.Just? _6) (Data.Maybe.Just (scm:cond
                    [(Data.Maybe.Nothing? acc4) (rt:make-array (Data.Maybe.Just-value0 _6))]
                    [(Data.Maybe.Just? acc4) ((snoc (Data.Maybe.Just-value0 acc4)) (Data.Maybe.Just-value0 _6))]
                    [scm:else (rt:fail)]))]
                  [scm:else (rt:fail)]))))) Data.Maybe.Nothing) xs0)])
            (scm:cond
              [(Data.Maybe.Nothing? v4) allArrays3]
              [(Data.Maybe.Just? v4) ((go1 (scm:fx+ idx2 1)) ((snoc allArrays3) (Data.Maybe.Just-value0 v4)))]
              [scm:else (rt:fail)]))))])
        ((go1 0) (rt:make-array)))))

  (scm:define foldRecM
    (scm:lambda (dictMonadRec0)
      (scm:let*
        ([Monad01 ((rt:record-ref dictMonadRec0 (scm:string->symbol "Monad0")) (scm:quote undefined))]
         [_2 ((rt:record-ref Monad01 (scm:string->symbol "Applicative0")) (scm:quote undefined))])
          (scm:lambda (f3)
            (scm:lambda (b4)
              (scm:lambda (array5)
                (((rt:record-ref dictMonadRec0 (scm:string->symbol "tailRecM")) (scm:lambda (o6)
                  (scm:cond
                    [(scm:fx>=? (rt:record-ref o6 (scm:string->symbol "b")) (rt:array-length array5)) ((rt:record-ref _2 (scm:string->symbol "pure")) (Control.Monad.Rec.Class.Done (rt:record-ref o6 (scm:string->symbol "a"))))]
                    [scm:else (((rt:record-ref ((rt:record-ref Monad01 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) ((f3 (rt:record-ref o6 (scm:string->symbol "a"))) (rt:array-ref array5 (rt:record-ref o6 (scm:string->symbol "b"))))) (scm:lambda (res$p7)
                      ((rt:record-ref _2 (scm:string->symbol "pure")) (Control.Monad.Rec.Class.Loop (scm:list (scm:cons (scm:string->symbol "a") res$p7) (scm:cons (scm:string->symbol "b") (scm:fx+ (rt:record-ref o6 (scm:string->symbol "b")) 1)))))))]))) (scm:list (scm:cons (scm:string->symbol "a") b4) (scm:cons (scm:string->symbol "b") 0)))))))))

  (scm:define foldMap
    (scm:lambda (dictMonoid0)
      ((rt:record-ref Data.Foldable.foldableArray (scm:string->symbol "foldMap")) dictMonoid0)))

  (scm:define foldM
    (scm:lambda (dictMonad0)
      (scm:lambda (f1)
        (scm:lambda (b2)
          (scm:lambda (_3)
            (unconsImpl (scm:lambda (v4)
              ((rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) b2)) (scm:lambda (a4)
              (scm:lambda (as5)
                (((rt:record-ref ((rt:record-ref dictMonad0 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) ((f1 b2) a4)) (scm:lambda (b$p6)
                  ((((foldM dictMonad0) f1) b$p6) as5))))) _3))))))

  (scm:define fold
    (scm:lambda (dictMonoid0)
      (((rt:record-ref Data.Foldable.foldableArray (scm:string->symbol "foldMap")) dictMonoid0) Data.Foldable.identity)))

  (scm:define findMap
    (scm:lambda (_0)
      (scm:lambda (_1)
        (findMapImpl Data.Maybe.Nothing Data.Maybe.isJust _0 _1))))

  (scm:define findLastIndex
    (scm:lambda (_0)
      (scm:lambda (_1)
        (findLastIndexImpl Data.Maybe.Just Data.Maybe.Nothing _0 _1))))

  (scm:define insertBy
    (scm:lambda (cmp0)
      (scm:lambda (x1)
        (scm:lambda (ys2)
          (scm:let ([_3 (_insertAt Data.Maybe.Just Data.Maybe.Nothing (scm:let ([_3 (findLastIndexImpl Data.Maybe.Just Data.Maybe.Nothing (scm:lambda (y3)
            (Data.Ordering.GT? ((cmp0 x1) y3))) ys2)])
            (scm:cond
              [(Data.Maybe.Nothing? _3) 0]
              [(Data.Maybe.Just? _3) (scm:fx+ (Data.Maybe.Just-value0 _3) 1)]
              [scm:else (rt:fail)])) x1 ys2)])
            (scm:cond
              [(Data.Maybe.Just? _3) (Data.Maybe.Just-value0 _3)]
              [scm:else (rt:fail)]))))))

  (scm:define insert
    (scm:lambda (dictOrd0)
      (insertBy (rt:record-ref dictOrd0 (scm:string->symbol "compare")))))

  (scm:define findIndex
    (scm:lambda (_0)
      (scm:lambda (_1)
        (findIndexImpl Data.Maybe.Just Data.Maybe.Nothing _0 _1))))

  (scm:define find
    (scm:lambda (f0)
      (scm:lambda (xs1)
        (scm:let ([_2 (findIndexImpl Data.Maybe.Just Data.Maybe.Nothing f0 xs1)])
          (scm:cond
            [(Data.Maybe.Just? _2) (Data.Maybe.Just (rt:array-ref xs1 (Data.Maybe.Just-value0 _2)))]
            [scm:else Data.Maybe.Nothing])))))

  (scm:define filter
    (scm:lambda (_0)
      (scm:lambda (_1)
        (filterImpl _0 _1))))

  (scm:define intersectBy
    (scm:lambda (eq20)
      (scm:lambda (xs1)
        (scm:lambda (ys2)
          (filterImpl (scm:lambda (x3)
            (scm:let ([_4 (findIndexImpl Data.Maybe.Just Data.Maybe.Nothing (eq20 x3) ys2)])
              (scm:cond
                [(Data.Maybe.Nothing? _4) #f]
                [(Data.Maybe.Just? _4) #t]
                [scm:else (rt:fail)]))) xs1)))))

  (scm:define intersect
    (scm:lambda (dictEq0)
      (intersectBy (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define elemLastIndex
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (findLastIndex (scm:lambda (v2)
          (((rt:record-ref dictEq0 (scm:string->symbol "eq")) v2) x1))))))

  (scm:define elemIndex
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (findIndex (scm:lambda (v2)
          (((rt:record-ref dictEq0 (scm:string->symbol "eq")) v2) x1))))))

  (scm:define notElem
    (scm:lambda (dictEq0)
      (scm:lambda (a1)
        (scm:lambda (arr2)
          (scm:let ([_3 (findIndexImpl Data.Maybe.Just Data.Maybe.Nothing (scm:lambda (v3)
            (((rt:record-ref dictEq0 (scm:string->symbol "eq")) v3) a1)) arr2)])
            (scm:cond
              [(Data.Maybe.Nothing? _3) #t]
              [(Data.Maybe.Just? _3) #f]
              [scm:else (rt:fail)]))))))

  (scm:define elem
    (scm:lambda (dictEq0)
      (scm:lambda (a1)
        (scm:lambda (arr2)
          (scm:let ([_3 (findIndexImpl Data.Maybe.Just Data.Maybe.Nothing (scm:lambda (v3)
            (((rt:record-ref dictEq0 (scm:string->symbol "eq")) v3) a1)) arr2)])
            (scm:cond
              [(Data.Maybe.Nothing? _3) #f]
              [(Data.Maybe.Just? _3) #t]
              [scm:else (rt:fail)]))))))

  (scm:define dropWhile
    (scm:lambda (p0)
      (scm:lambda (xs1)
        (rt:record-ref ((span p0) xs1) (scm:string->symbol "rest")))))

  (scm:define dropEnd
    (scm:lambda (n0)
      (scm:lambda (xs1)
        (scm:let ([_2 (scm:fx- (rt:array-length xs1) n0)])
          (scm:cond
            [(scm:fx<? _2 1) (rt:make-array)]
            [scm:else (sliceImpl 0 _2 xs1)])))))

  (scm:define drop
    (scm:lambda (n0)
      (scm:lambda (xs1)
        (scm:cond
          [(scm:fx<? n0 1) xs1]
          [scm:else (sliceImpl n0 (rt:array-length xs1) xs1)]))))

  (scm:define takeEnd
    (scm:lambda (n0)
      (scm:lambda (xs1)
        (scm:let ([_2 (scm:fx- (rt:array-length xs1) n0)])
          (scm:cond
            [(scm:fx<? _2 1) xs1]
            [scm:else (sliceImpl _2 (rt:array-length xs1) xs1)])))))

  (scm:define deleteAt
    (scm:lambda (_0)
      (scm:lambda (_1)
        (_deleteAt Data.Maybe.Just Data.Maybe.Nothing _0 _1))))

  (scm:define deleteBy
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(scm:fx=? (rt:array-length v22) 0) (rt:make-array)]
            [scm:else (scm:let ([_3 (findIndexImpl Data.Maybe.Just Data.Maybe.Nothing (v0 v11) v22)])
              (scm:cond
                [(Data.Maybe.Nothing? _3) v22]
                [(Data.Maybe.Just? _3) (scm:let ([_4 (_deleteAt Data.Maybe.Just Data.Maybe.Nothing (Data.Maybe.Just-value0 _3) v22)])
                  (scm:cond
                    [(Data.Maybe.Just? _4) (Data.Maybe.Just-value0 _4)]
                    [scm:else (rt:fail)]))]
                [scm:else (rt:fail)]))])))))

  (scm:define delete
    (scm:lambda (dictEq0)
      (deleteBy (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define difference
    (scm:lambda (dictEq0)
      (Data.Foldable.foldrArray (delete dictEq0))))

  (scm:define cons
    (scm:lambda (x0)
      (scm:lambda (xs1)
        ((Data.Semigroup.concatArray (rt:make-array x0)) xs1))))

  (scm:define some
    (scm:lambda (dictAlternative0)
      (scm:lambda (dictLazy1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref ((rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Plus1")) (scm:quote undefined)) (scm:string->symbol "Alt0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) cons) v2)) ((rt:record-ref dictLazy1 (scm:string->symbol "defer")) (scm:lambda (v13)
            (((many dictAlternative0) dictLazy1) v2))))))))

  (scm:define many
    (scm:lambda (dictAlternative0)
      (scm:lambda (dictLazy1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Plus1")) (scm:quote undefined)) (scm:string->symbol "Alt0")) (scm:quote undefined)) (scm:string->symbol "alt")) (((some dictAlternative0) dictLazy1) v2)) ((rt:record-ref ((rt:record-ref dictAlternative0 (scm:string->symbol "Applicative0")) (scm:quote undefined)) (scm:string->symbol "pure")) (rt:make-array)))))))

  (scm:define concatMap
    (scm:lambda (b0)
      (scm:lambda (a1)
        ((Control.Bind.arrayBind a1) b0))))

  (scm:define mapMaybe
    (scm:lambda (f0)
      (concatMap (scm:lambda (x1)
        (scm:let ([_2 (f0 x1)])
          (scm:cond
            [(Data.Maybe.Nothing? _2) (rt:make-array)]
            [(Data.Maybe.Just? _2) (rt:make-array (Data.Maybe.Just-value0 _2))]
            [scm:else (rt:fail)]))))))

  (scm:define filterA
    (scm:lambda (dictApplicative0)
      (scm:let*
        ([traverse11 ((rt:record-ref Data.Traversable.traversableArray (scm:string->symbol "traverse")) dictApplicative0)]
         [_2 ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))])
          (scm:lambda (p3)
            (scm:let*
              ([_4 (traverse11 (scm:lambda (x4)
                (((rt:record-ref _2 (scm:string->symbol "map")) (Data.Tuple.Tuple x4)) (p3 x4))))]
               [_5 ((rt:record-ref _2 (scm:string->symbol "map")) (mapMaybe (scm:lambda (v5)
                (scm:cond
                  [(Data.Tuple.Tuple-value1 v5) (Data.Maybe.Just (Data.Tuple.Tuple-value0 v5))]
                  [scm:else Data.Maybe.Nothing]))))])
                (scm:lambda (x6)
                  (_5 (_4 x6))))))))

  (scm:define catMaybes
    (mapMaybe (scm:lambda (x0)
      x0)))

  (scm:define any
    (scm:lambda (_0)
      (scm:lambda (_1)
        (anyImpl _0 _1))))

  (scm:define nubByEq
    (scm:lambda (eq20)
      (scm:lambda (xs1)
        (Control.Monad.ST.Internal.run (scm:lambda ()
          (scm:let*
            ([arr2 (Data.Array.ST.new)]
             [_ (((Control.Monad.ST.Internal.foreach xs1) (scm:lambda (x3)
              (scm:let ([_4 (any (scm:lambda (v4)
                ((eq20 v4) x3)))])
                (scm:lambda ()
                  (scm:let ([_5 (Data.Array.ST.unsafeFreezeImpl arr2)])
                    ((scm:cond
                      [(scm:not (_4 _5)) (scm:lambda ()
                        (scm:let ([_6 (Data.Array.ST.pushImpl x3 arr2)])
                          Data.Unit.unit))]
                      [scm:else (scm:lambda ()
                        Data.Unit.unit)]))))))))])
              ((scm:lambda ()
                (Data.Array.ST.unsafeFreezeImpl arr2)))))))))

  (scm:define nubEq
    (scm:lambda (dictEq0)
      (nubByEq (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define unionBy
    (scm:lambda (eq20)
      (scm:lambda (xs1)
        (scm:lambda (ys2)
          ((Data.Semigroup.concatArray xs1) (((Data.Foldable.foldlArray (scm:lambda (b3)
            (scm:lambda (a4)
              (((deleteBy eq20) a4) b3)))) ((nubByEq eq20) ys2)) xs1))))))

  (scm:define union
    (scm:lambda (dictEq0)
      (unionBy (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define alterAt
    (scm:lambda (i0)
      (scm:lambda (f1)
        (scm:lambda (xs2)
          (scm:let ([_3 (indexImpl Data.Maybe.Just Data.Maybe.Nothing xs2 i0)])
            (scm:cond
              [(Data.Maybe.Nothing? _3) Data.Maybe.Nothing]
              [(Data.Maybe.Just? _3) (scm:let ([v4 (f1 (Data.Maybe.Just-value0 _3))])
                (scm:cond
                  [(Data.Maybe.Nothing? v4) (_deleteAt Data.Maybe.Just Data.Maybe.Nothing i0 xs2)]
                  [(Data.Maybe.Just? v4) (_updateAt Data.Maybe.Just Data.Maybe.Nothing i0 (Data.Maybe.Just-value0 v4) xs2)]
                  [scm:else (rt:fail)]))]
              [scm:else (rt:fail)]))))))

  (scm:define all
    (scm:lambda (_0)
      (scm:lambda (_1)
        (allImpl _0 _1)))))
