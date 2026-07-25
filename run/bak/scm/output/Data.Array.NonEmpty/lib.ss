#!r6rs
#!chezscheme
(library
  (Data.Array.NonEmpty lib)
  (export
    all
    alterAt
    any
    appendArray
    catMaybes
    concat
    concatMap
    cons
    cons$p
    delete
    deleteAt
    deleteBy
    difference
    difference$p
    drop
    dropEnd
    dropWhile
    elem
    elemIndex
    elemLastIndex
    filter
    filterA
    find
    findIndex
    findLastIndex
    findMap
    fold1
    foldM
    foldMap1
    foldRecM
    foldl1
    foldr1
    fromArray
    fromFoldable
    fromFoldable1
    fromNonEmpty
    group
    groupAll
    groupAllBy
    groupBy
    head
    index
    init
    insert
    insertAt
    insertBy
    intercalate
    intercalate1
    intersect
    intersect$p
    intersectBy
    intersectBy$p
    intersperse
    last
    length
    mapMaybe
    mapWithIndex
    max
    modifyAt
    modifyAtIndices
    notElem
    nub
    nubBy
    nubByEq
    nubEq
    partition
    prependArray
    range
    replicate
    reverse
    scanl
    scanr
    singleton
    slice
    snoc
    snoc$p
    some
    sort
    sortBy
    sortWith
    span
    splitAt
    tail
    take
    takeEnd
    takeWhile
    toArray
    toNonEmpty
    toUnfoldable
    toUnfoldable1
    transpose
    transpose$p
    uncons
    union
    union$p
    unionBy
    unionBy$p
    unsafeIndex
    unsnoc
    unzip
    updateAt
    updateAtIndices
    zip
    zipWith
    zipWithA)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Bind lib) Control.Bind.)
    (prefix (Control.Monad.Rec.Class lib) Control.Monad.Rec.Class.)
    (prefix (Data.Array lib) Data.Array.)
    (prefix (Data.Array.NonEmpty.Internal lib) Data.Array.NonEmpty.Internal.)
    (prefix (Data.Foldable lib) Data.Foldable.)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.FunctorWithIndex lib) Data.FunctorWithIndex.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.NonEmpty lib) Data.NonEmpty.)
    (prefix (Data.Ord lib) Data.Ord.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Semigroup lib) Data.Semigroup.)
    (prefix (Data.Semigroup.Foldable lib) Data.Semigroup.Foldable.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define max
    (scm:lambda (x0)
      (scm:lambda (y1)
        (scm:let ([v2 (((rt:record-ref Data.Ord.ordInt (scm:string->symbol "compare")) x0) y1)])
          (scm:cond
            [(Data.Ordering.LT? v2) y1]
            [(Data.Ordering.EQ? v2) x0]
            [(Data.Ordering.GT? v2) x0]
            [scm:else (rt:fail)])))))

  (scm:define intercalate1
    (scm:lambda (dictSemigroup0)
      (scm:let ([foldMap121 ((rt:record-ref Data.Array.NonEmpty.Internal.foldable1NonEmptyArray (scm:string->symbol "foldMap1")) (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:lambda (j3)
            (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (v1 j3)) (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) j3) (v12 j3)))))))))])
        (scm:lambda (a2)
          (scm:lambda (foldable3)
            (((foldMap121 (scm:lambda (x4)
              (scm:lambda (v5)
                x4))) foldable3) a2))))))

  (scm:define transpose
    (scm:lambda (x0)
      (Data.Array.transpose x0)))

  (scm:define toArray
    (scm:lambda (v0)
      v0))

  (scm:define unionBy$p
    (scm:lambda (eq0)
      (scm:lambda (xs1)
        (scm:lambda (x2)
          (((Data.Array.unionBy eq0) xs1) x2)))))

  (scm:define union$p
    (scm:lambda (dictEq0)
      (unionBy$p (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define unionBy
    (scm:lambda (eq0)
      (scm:lambda (xs1)
        (scm:lambda (x2)
          (((Data.Array.unionBy eq0) xs1) x2)))))

  (scm:define union
    (scm:lambda (dictEq0)
      (unionBy (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define unzip
    (scm:lambda (x0)
      (scm:let ([_1 (Data.Array.unzip x0)])
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 _1) (Data.Tuple.Tuple-value1 _1)))))

  (scm:define updateAt
    (scm:lambda (i0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (Data.Array._updateAt Data.Maybe.Just Data.Maybe.Nothing i0 x1 x2)))))

  (scm:define zip
    (scm:lambda (xs0)
      (scm:lambda (ys1)
        (Data.Array.zipWithImpl Data.Tuple.Tuple xs0 ys1))))

  (scm:define zipWith
    (scm:lambda (f0)
      (scm:lambda (xs1)
        (scm:lambda (ys2)
          (Data.Array.zipWithImpl f0 xs1 ys2)))))

  (scm:define zipWithA
    (scm:lambda (dictApplicative0)
      (Data.Array.zipWithA dictApplicative0)))

  (scm:define splitAt
    (scm:lambda (i0)
      (scm:lambda (xs1)
        ((Data.Array.splitAt i0) xs1))))

  (scm:define some
    (scm:lambda (dictAlternative0)
      (scm:lambda (dictLazy1)
        (scm:lambda (x2)
          (((Data.Array.some dictAlternative0) dictLazy1) x2)))))

  (scm:define snoc$p
    (scm:lambda (xs0)
      (scm:lambda (x1)
        ((Data.Array.snoc xs0) x1))))

  (scm:define snoc
    (scm:lambda (xs0)
      (scm:lambda (x1)
        ((Data.Array.snoc xs0) x1))))

  (scm:define singleton
    (scm:lambda (x0)
      (rt:make-array x0)))

  (scm:define replicate
    (scm:lambda (i0)
      (scm:lambda (x1)
        (Data.Array.replicateImpl ((max 1) i0) x1))))

  (scm:define range
    (scm:lambda (x0)
      (scm:lambda (y1)
        (Data.Array.rangeImpl x0 y1))))

  (scm:define prependArray
    (scm:lambda (xs0)
      (scm:lambda (ys1)
        ((Data.Semigroup.concatArray xs0) ys1))))

  (scm:define modifyAt
    (scm:lambda (i0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          (((Data.Array.modifyAt i0) f1) x2)))))

  (scm:define intersectBy$p
    (scm:lambda (eq0)
      (scm:lambda (xs1)
        ((Data.Array.intersectBy eq0) xs1))))

  (scm:define intersectBy
    (scm:lambda (eq0)
      (scm:lambda (xs1)
        (scm:lambda (x2)
          (((Data.Array.intersectBy eq0) xs1) x2)))))

  (scm:define intersect$p
    (scm:lambda (dictEq0)
      (intersectBy$p (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define intersect
    (scm:lambda (dictEq0)
      (intersectBy (rt:record-ref dictEq0 (scm:string->symbol "eq")))))

  (scm:define intercalate
    (scm:lambda (dictSemigroup0)
      (intercalate1 dictSemigroup0)))

  (scm:define insertAt
    (scm:lambda (i0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (Data.Array._insertAt Data.Maybe.Just Data.Maybe.Nothing i0 x1 x2)))))

  (scm:define fromFoldable1
    (scm:lambda (dictFoldable10)
      (scm:let ([_1 (rt:record-ref ((rt:record-ref dictFoldable10 (scm:string->symbol "Foldable0")) (scm:quote undefined)) (scm:string->symbol "foldr"))])
        (scm:lambda (x2)
          (Data.Array.fromFoldableImpl _1 x2)))))

  (scm:define fromArray
    (scm:lambda (xs0)
      (scm:cond
        [(scm:fx>? (rt:array-length xs0) 0) (Data.Maybe.Just xs0)]
        [scm:else Data.Maybe.Nothing])))

  (scm:define fromFoldable
    (scm:lambda (dictFoldable0)
      (scm:let ([_1 (rt:record-ref dictFoldable0 (scm:string->symbol "foldr"))])
        (scm:lambda (x2)
          (scm:let ([_3 (Data.Array.fromFoldableImpl _1 x2)])
            (scm:cond
              [(scm:fx>? (rt:array-length _3) 0) (Data.Maybe.Just _3)]
              [scm:else Data.Maybe.Nothing]))))))

  (scm:define transpose$p
    (scm:lambda (x0)
      (scm:let ([_1 (Data.Array.transpose x0)])
        (scm:cond
          [(scm:fx>? (rt:array-length _1) 0) (Data.Maybe.Just _1)]
          [scm:else Data.Maybe.Nothing]))))

  (scm:define foldr1
    (rt:record-ref Data.Array.NonEmpty.Internal.foldable1NonEmptyArray (scm:string->symbol "foldr1")))

  (scm:define foldl1
    (rt:record-ref Data.Array.NonEmpty.Internal.foldable1NonEmptyArray (scm:string->symbol "foldl1")))

  (scm:define foldMap1
    (scm:lambda (dictSemigroup0)
      ((rt:record-ref Data.Array.NonEmpty.Internal.foldable1NonEmptyArray (scm:string->symbol "foldMap1")) dictSemigroup0)))

  (scm:define fold1
    (scm:lambda (dictSemigroup0)
      (((rt:record-ref Data.Array.NonEmpty.Internal.foldable1NonEmptyArray (scm:string->symbol "foldMap1")) dictSemigroup0) Data.Semigroup.Foldable.identity)))

  (scm:define difference$p
    (scm:lambda (dictEq0)
      (Data.Foldable.foldrArray (Data.Array.delete dictEq0))))

  (scm:define cons$p
    (scm:lambda (x0)
      (scm:lambda (xs1)
        ((Data.Semigroup.concatArray (rt:make-array x0)) xs1))))

  (scm:define fromNonEmpty
    (scm:lambda (v0)
      ((Data.Semigroup.concatArray (rt:make-array (Data.NonEmpty.NonEmpty-value0 v0))) (Data.NonEmpty.NonEmpty-value1 v0))))

  (scm:define concatMap
    (scm:lambda (b0)
      (scm:lambda (a1)
        ((Control.Bind.arrayBind a1) b0))))

  (scm:define concat
    (scm:let ([_0 (Data.Functor.arrayMap toArray)])
      (scm:lambda (x1)
        (Data.Array.concat (_0 x1)))))

  (scm:define appendArray
    (scm:lambda (xs0)
      (scm:lambda (ys1)
        ((Data.Semigroup.concatArray xs0) ys1))))

  (scm:define alterAt
    (scm:lambda (i0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          (((Data.Array.alterAt i0) f1) x2)))))

  (scm:define head
    (scm:lambda (x0)
      (scm:let ([_1 (Data.Array.indexImpl Data.Maybe.Just Data.Maybe.Nothing x0 0)])
        (scm:cond
          [(Data.Maybe.Just? _1) (Data.Maybe.Just-value0 _1)]
          [scm:else (rt:fail)]))))

  (scm:define init
    (scm:lambda (x0)
      (scm:cond
        [(scm:fx=? (rt:array-length x0) 0) (rt:fail)]
        [scm:else (Data.Array.sliceImpl 0 (scm:fx- (rt:array-length x0) 1) x0)])))

  (scm:define last
    (scm:lambda (x0)
      (scm:let ([_1 (Data.Array.indexImpl Data.Maybe.Just Data.Maybe.Nothing x0 (scm:fx- (rt:array-length x0) 1))])
        (scm:cond
          [(Data.Maybe.Just? _1) (Data.Maybe.Just-value0 _1)]
          [scm:else (rt:fail)]))))

  (scm:define tail
    (scm:lambda (x0)
      (scm:let ([_1 (Data.Array.unconsImpl (scm:lambda (v1)
        Data.Maybe.Nothing) (scm:lambda (v1)
        (scm:lambda (xs2)
          (Data.Maybe.Just xs2))) x0)])
        (scm:cond
          [(Data.Maybe.Just? _1) (Data.Maybe.Just-value0 _1)]
          [scm:else (rt:fail)]))))

  (scm:define uncons
    (scm:lambda (x0)
      (scm:let ([_1 (Data.Array.unconsImpl (scm:lambda (v1)
        Data.Maybe.Nothing) (scm:lambda (x1)
        (scm:lambda (xs2)
          (Data.Maybe.Just (scm:list (scm:cons (scm:string->symbol "head") x1) (scm:cons (scm:string->symbol "tail") xs2))))) x0)])
        (scm:cond
          [(Data.Maybe.Just? _1) (Data.Maybe.Just-value0 _1)]
          [scm:else (rt:fail)]))))

  (scm:define toNonEmpty
    (scm:lambda (x0)
      (scm:let ([_1 (uncons x0)])
        (Data.NonEmpty.NonEmpty* (rt:record-ref _1 (scm:string->symbol "head")) (rt:record-ref _1 (scm:string->symbol "tail"))))))

  (scm:define unsnoc
    (scm:lambda (x0)
      (scm:let ([_1 (Data.Array.unsnoc x0)])
        (scm:cond
          [(Data.Maybe.Just? _1) (Data.Maybe.Just-value0 _1)]
          [scm:else (rt:fail)]))))

  (scm:define all
    (scm:lambda (p0)
      (scm:lambda (x1)
        (Data.Array.allImpl p0 x1))))

  (scm:define any
    (scm:lambda (p0)
      (scm:lambda (x1)
        (Data.Array.anyImpl p0 x1))))

  (scm:define catMaybes
    (scm:lambda (x0)
      ((Data.Array.mapMaybe (scm:lambda (x1)
        x1)) x0)))

  (scm:define delete
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (((Data.Array.deleteBy (rt:record-ref dictEq0 (scm:string->symbol "eq"))) x1) x2)))))

  (scm:define deleteAt
    (scm:lambda (i0)
      (scm:lambda (x1)
        (Data.Array._deleteAt Data.Maybe.Just Data.Maybe.Nothing i0 x1))))

  (scm:define deleteBy
    (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (((Data.Array.deleteBy f0) x1) x2)))))

  (scm:define difference
    (scm:lambda (dictEq0)
      (Data.Foldable.foldrArray (Data.Array.delete dictEq0))))

  (scm:define drop
    (scm:lambda (i0)
      (scm:lambda (x1)
        (scm:cond
          [(scm:fx<? i0 1) x1]
          [scm:else (Data.Array.sliceImpl i0 (rt:array-length x1) x1)]))))

  (scm:define dropEnd
    (scm:lambda (i0)
      (scm:lambda (x1)
        (scm:let ([_2 (scm:fx- (rt:array-length x1) i0)])
          (scm:cond
            [(scm:fx<? _2 1) (rt:make-array)]
            [scm:else (Data.Array.sliceImpl 0 _2 x1)])))))

  (scm:define dropWhile
    (scm:lambda (f0)
      (scm:lambda (x1)
        (rt:record-ref ((Data.Array.span f0) x1) (scm:string->symbol "rest")))))

  (scm:define elem
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (((Data.Array.elem dictEq0) x1) x2)))))

  (scm:define elemIndex
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (Data.Array.findIndexImpl Data.Maybe.Just Data.Maybe.Nothing (scm:lambda (v3)
            (((rt:record-ref dictEq0 (scm:string->symbol "eq")) v3) x1)) x2)))))

  (scm:define elemLastIndex
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (Data.Array.findLastIndexImpl Data.Maybe.Just Data.Maybe.Nothing (scm:lambda (v3)
            (((rt:record-ref dictEq0 (scm:string->symbol "eq")) v3) x1)) x2)))))

  (scm:define filter
    (scm:lambda (f0)
      (scm:lambda (x1)
        (Data.Array.filterImpl f0 x1))))

  (scm:define filterA
    (scm:lambda (dictApplicative0)
      (Data.Array.filterA dictApplicative0)))

  (scm:define find
    (scm:lambda (p0)
      (scm:lambda (x1)
        ((Data.Array.find p0) x1))))

  (scm:define findIndex
    (scm:lambda (p0)
      (scm:lambda (x1)
        (Data.Array.findIndexImpl Data.Maybe.Just Data.Maybe.Nothing p0 x1))))

  (scm:define findLastIndex
    (scm:lambda (x0)
      (scm:lambda (x1)
        (Data.Array.findLastIndexImpl Data.Maybe.Just Data.Maybe.Nothing x0 x1))))

  (scm:define findMap
    (scm:lambda (p0)
      (scm:lambda (x1)
        (Data.Array.findMapImpl Data.Maybe.Nothing Data.Maybe.isJust p0 x1))))

  (scm:define foldM
    (scm:lambda (dictMonad0)
      (scm:lambda (f1)
        (scm:lambda (acc2)
          (scm:lambda (x3)
            ((((Data.Array.foldM dictMonad0) f1) acc2) x3))))))

  (scm:define foldRecM
    (scm:lambda (dictMonadRec0)
      (scm:let*
        ([Monad01 ((rt:record-ref dictMonadRec0 (scm:string->symbol "Monad0")) (scm:quote undefined))]
         [_2 ((rt:record-ref Monad01 (scm:string->symbol "Applicative0")) (scm:quote undefined))])
          (scm:lambda (f3)
            (scm:lambda (acc4)
              (scm:lambda (array5)
                (((rt:record-ref dictMonadRec0 (scm:string->symbol "tailRecM")) (scm:lambda (o6)
                  (scm:cond
                    [(scm:fx>=? (rt:record-ref o6 (scm:string->symbol "b")) (rt:array-length array5)) ((rt:record-ref _2 (scm:string->symbol "pure")) (Control.Monad.Rec.Class.Done (rt:record-ref o6 (scm:string->symbol "a"))))]
                    [scm:else (((rt:record-ref ((rt:record-ref Monad01 (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "bind")) ((f3 (rt:record-ref o6 (scm:string->symbol "a"))) (rt:array-ref array5 (rt:record-ref o6 (scm:string->symbol "b"))))) (scm:lambda (res$p7)
                      ((rt:record-ref _2 (scm:string->symbol "pure")) (Control.Monad.Rec.Class.Loop (scm:list (scm:cons (scm:string->symbol "a") res$p7) (scm:cons (scm:string->symbol "b") (scm:fx+ (rt:record-ref o6 (scm:string->symbol "b")) 1)))))))]))) (scm:list (scm:cons (scm:string->symbol "a") acc4) (scm:cons (scm:string->symbol "b") 0)))))))))

  (scm:define index
    (scm:lambda (x0)
      (Data.Array.index x0)))

  (scm:define length
    (scm:lambda (x0)
      (rt:array-length x0)))

  (scm:define mapMaybe
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((Data.Array.mapMaybe f0) x1))))

  (scm:define notElem
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (((Data.Array.notElem dictEq0) x1) x2)))))

  (scm:define partition
    (scm:lambda (f0)
      (scm:lambda (x1)
        (Data.Array.partitionImpl f0 x1))))

  (scm:define slice
    (scm:lambda (start0)
      (scm:lambda (end1)
        (scm:lambda (x2)
          (Data.Array.sliceImpl start0 end1 x2)))))

  (scm:define span
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((Data.Array.span f0) x1))))

  (scm:define take
    (scm:lambda (i0)
      (scm:lambda (x1)
        (scm:cond
          [(scm:fx<? i0 1) (rt:make-array)]
          [scm:else (Data.Array.sliceImpl 0 i0 x1)]))))

  (scm:define takeEnd
    (scm:lambda (i0)
      (scm:lambda (x1)
        ((Data.Array.takeEnd i0) x1))))

  (scm:define takeWhile
    (scm:lambda (f0)
      (scm:lambda (x1)
        (rt:record-ref ((Data.Array.span f0) x1) (scm:string->symbol "init")))))

  (scm:define toUnfoldable
    (scm:lambda (dictUnfoldable0)
      (scm:lambda (x1)
        (scm:let ([len2 (rt:array-length x1)])
          (((rt:record-ref dictUnfoldable0 (scm:string->symbol "unfoldr")) (scm:lambda (i3)
            (scm:cond
              [(scm:fx<? i3 len2) (Data.Maybe.Just (Data.Tuple.Tuple* (rt:array-ref x1 i3) (scm:fx+ i3 1)))]
              [scm:else Data.Maybe.Nothing]))) 0)))))

  (scm:define cons
    (scm:lambda (x0)
      (scm:lambda (x1)
        ((Data.Semigroup.concatArray (rt:make-array x0)) x1))))

  (scm:define group
    (scm:lambda (dictEq0)
      (scm:let ([eq21 (rt:record-ref dictEq0 (scm:string->symbol "eq"))])
        (scm:lambda (x2)
          ((Data.Array.groupBy eq21) x2)))))

  (scm:define groupAllBy
    (scm:lambda (op0)
      (Data.Array.groupAllBy op0)))

  (scm:define groupAll
    (scm:lambda (dictOrd0)
      (Data.Array.groupAllBy (rt:record-ref dictOrd0 (scm:string->symbol "compare")))))

  (scm:define groupBy
    (scm:lambda (op0)
      (scm:lambda (x1)
        ((Data.Array.groupBy op0) x1))))

  (scm:define insert
    (scm:lambda (dictOrd0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (((Data.Array.insertBy (rt:record-ref dictOrd0 (scm:string->symbol "compare"))) x1) x2)))))

  (scm:define insertBy
    (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (((Data.Array.insertBy f0) x1) x2)))))

  (scm:define intersperse
    (scm:lambda (x0)
      (scm:lambda (x1)
        ((Data.Array.intersperse x0) x1))))

  (scm:define mapWithIndex
    (scm:lambda (f0)
      (Data.FunctorWithIndex.mapWithIndexArray f0)))

  (scm:define modifyAtIndices
    (scm:lambda (dictFoldable0)
      (Data.Array.modifyAtIndices dictFoldable0)))

  (scm:define nub
    (scm:lambda (dictOrd0)
      (scm:lambda (x1)
        ((Data.Array.nubBy (rt:record-ref dictOrd0 (scm:string->symbol "compare"))) x1))))

  (scm:define nubBy
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((Data.Array.nubBy f0) x1))))

  (scm:define nubByEq
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((Data.Array.nubByEq f0) x1))))

  (scm:define nubEq
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        ((Data.Array.nubByEq (rt:record-ref dictEq0 (scm:string->symbol "eq"))) x1))))

  (scm:define reverse
    (scm:lambda (x0)
      (Data.Array.reverse x0)))

  (scm:define scanl
    (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (Data.Array.scanlImpl f0 x1 x2)))))

  (scm:define scanr
    (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (x2)
          (Data.Array.scanrImpl f0 x1 x2)))))

  (scm:define sort
    (scm:lambda (dictOrd0)
      (scm:let ([compare1 (rt:record-ref dictOrd0 (scm:string->symbol "compare"))])
        (scm:lambda (x2)
          ((Data.Array.sortBy compare1) x2)))))

  (scm:define sortBy
    (scm:lambda (f0)
      (scm:lambda (x1)
        ((Data.Array.sortBy f0) x1))))

  (scm:define sortWith
    (scm:lambda (dictOrd0)
      (scm:lambda (f1)
        (scm:lambda (x2)
          (((Data.Array.sortWith dictOrd0) f1) x2)))))

  (scm:define updateAtIndices
    (scm:lambda (dictFoldable0)
      (Data.Array.updateAtIndices dictFoldable0)))

  (scm:define unsafeIndex
    (scm:lambda (_)
      (scm:lambda (x1)
        (scm:lambda (_2)
          (rt:array-ref x1 _2)))))

  (scm:define toUnfoldable1
    (scm:lambda (dictUnfoldable10)
      (scm:lambda (xs1)
        (scm:let ([len2 (rt:array-length xs1)])
          (((rt:record-ref dictUnfoldable10 (scm:string->symbol "unfoldr1")) (scm:lambda (i3)
            (Data.Tuple.Tuple* (rt:array-ref xs1 i3) (scm:cond
              [(scm:fx<? i3 (scm:fx- len2 1)) (Data.Maybe.Just (scm:fx+ i3 1))]
              [scm:else Data.Maybe.Nothing])))) 0))))))
