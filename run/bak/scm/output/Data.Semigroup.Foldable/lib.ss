#!r6rs
#!chezscheme
(library
  (Data.Semigroup.Foldable lib)
  (export
    FoldRight1
    FoldRight1*
    FoldRight1-value0
    FoldRight1-value1
    FoldRight1?
    fold1
    foldMap1
    foldMap1DefaultL
    foldMap1DefaultR
    foldRight1Semigroup
    foldableDual
    foldableIdentity
    foldableMultiplicative
    foldableTuple
    foldl1
    foldl1Default
    foldr1
    foldr1Default
    for1_
    identity
    intercalate
    intercalateMap
    maximum
    maximumBy
    minimum
    minimumBy
    mkFoldRight1
    semigroupAct
    semigroupDual
    sequence1_
    traverse1_)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Apply lib) Control.Apply.)
    (prefix (Data.Foldable lib) Data.Foldable.)
    (prefix (Data.Function lib) Data.Function.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define-record-type (FoldRight1$ FoldRight1* FoldRight1?)
    (scm:fields (scm:immutable value0 FoldRight1-value0) (scm:immutable value1 FoldRight1-value1)))

  (scm:define FoldRight1
    (scm:lambda (value0)
      (scm:lambda (value1)
        (FoldRight1* value0 value1))))

  (scm:define semigroupAct
    (scm:lambda (dictApply0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (((rt:record-ref dictApply0 (scm:string->symbol "apply")) (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (v3)
            Control.Apply.identity)) v1)) v12)))))))

  (scm:define mkFoldRight1
    (FoldRight1 Data.Function.const))

  (scm:define foldr1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldr1"))))

  (scm:define foldl1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldl1"))))

  (scm:define maximumBy
    (scm:lambda (dictFoldable10)
      (scm:lambda (cmp1)
        ((rt:record-ref dictFoldable10 (scm:string->symbol "foldl1")) (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:cond
              [(Data.Ordering.GT? ((cmp1 x2) y3)) x2]
              [scm:else y3])))))))

  (scm:define minimumBy
    (scm:lambda (dictFoldable10)
      (scm:lambda (cmp1)
        ((rt:record-ref dictFoldable10 (scm:string->symbol "foldl1")) (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:cond
              [(Data.Ordering.LT? ((cmp1 x2) y3)) x2]
              [scm:else y3])))))))

  (scm:define foldableTuple
    (scm:list (scm:cons (scm:string->symbol "foldMap1") (scm:lambda (dictSemigroup0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 (Data.Tuple.Tuple-value1 v2)))))) (scm:cons (scm:string->symbol "foldr1") (scm:lambda (v0)
      (scm:lambda (v11)
        (Data.Tuple.Tuple-value1 v11)))) (scm:cons (scm:string->symbol "foldl1") (scm:lambda (v0)
      (scm:lambda (v11)
        (Data.Tuple.Tuple-value1 v11)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableTuple))))

  (scm:define foldableMultiplicative
    (scm:list (scm:cons (scm:string->symbol "foldr1") (scm:lambda (v0)
      (scm:lambda (v11)
        v11))) (scm:cons (scm:string->symbol "foldl1") (scm:lambda (v0)
      (scm:lambda (v11)
        v11))) (scm:cons (scm:string->symbol "foldMap1") (scm:lambda (dictSemigroup0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableMultiplicative))))

  (scm:define foldableIdentity
    (scm:list (scm:cons (scm:string->symbol "foldMap1") (scm:lambda (dictSemigroup0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2))))) (scm:cons (scm:string->symbol "foldl1") (scm:lambda (v0)
      (scm:lambda (v11)
        v11))) (scm:cons (scm:string->symbol "foldr1") (scm:lambda (v0)
      (scm:lambda (v11)
        v11))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableIdentity))))

  (scm:define foldableDual
    (scm:list (scm:cons (scm:string->symbol "foldr1") (scm:lambda (v0)
      (scm:lambda (v11)
        v11))) (scm:cons (scm:string->symbol "foldl1") (scm:lambda (v0)
      (scm:lambda (v11)
        v11))) (scm:cons (scm:string->symbol "foldMap1") (scm:lambda (dictSemigroup0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (f1 v2))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableDual))))

  (scm:define foldRight1Semigroup
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:let ([_2 (FoldRight1-value1 v0)])
          (FoldRight1* (scm:lambda (a3)
            (scm:lambda (f4)
              (((FoldRight1-value0 v0) ((f4 _2) (((FoldRight1-value0 v11) a3) f4))) f4))) (FoldRight1-value1 v11))))))))

  (scm:define semigroupDual
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:let ([_2 (FoldRight1-value1 v11)])
          (FoldRight1* (scm:lambda (a3)
            (scm:lambda (f4)
              (((FoldRight1-value0 v11) ((f4 _2) (((FoldRight1-value0 v0) a3) f4))) f4))) (FoldRight1-value1 v0))))))))

  (scm:define foldMap1DefaultR
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictFunctor1)
        (scm:lambda (dictSemigroup2)
          (scm:let ([append3 (rt:record-ref dictSemigroup2 (scm:string->symbol "append"))])
            (scm:lambda (f4)
              (scm:let*
                ([_5 ((rt:record-ref dictFunctor1 (scm:string->symbol "map")) f4)]
                 [_6 ((rt:record-ref dictFoldable10 (scm:string->symbol "foldr1")) append3)])
                  (scm:lambda (x7)
                    (_6 (_5 x7))))))))))

  (scm:define foldMap1DefaultL
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictFunctor1)
        (scm:lambda (dictSemigroup2)
          (scm:let ([append3 (rt:record-ref dictSemigroup2 (scm:string->symbol "append"))])
            (scm:lambda (f4)
              (scm:let*
                ([_5 ((rt:record-ref dictFunctor1 (scm:string->symbol "map")) f4)]
                 [_6 ((rt:record-ref dictFoldable10 (scm:string->symbol "foldl1")) append3)])
                  (scm:lambda (x7)
                    (_6 (_5 x7))))))))))

  (scm:define foldMap1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "foldMap1"))))

  (scm:define foldl1Default
    (scm:lambda (dictFoldable10)
      (scm:let ([_1 (((rt:record-ref dictFoldable10 (scm:string->symbol "foldMap1")) semigroupDual) mkFoldRight1)])
        (scm:lambda (x2)
          (scm:lambda (a3)
            (scm:let ([_4 (_1 a3)])
              (((FoldRight1-value0 _4) (FoldRight1-value1 _4)) (scm:lambda (b5)
                (scm:lambda (a6)
                  ((x2 a6) b5))))))))))

  (scm:define foldr1Default
    (scm:lambda (dictFoldable10)
      (scm:let ([_1 (((rt:record-ref dictFoldable10 (scm:string->symbol "foldMap1")) foldRight1Semigroup) mkFoldRight1)])
        (scm:lambda (b2)
          (scm:lambda (a3)
            (scm:let ([_4 (_1 a3)])
              (((FoldRight1-value0 _4) (FoldRight1-value1 _4)) b2)))))))

  (scm:define intercalateMap
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictSemigroup1)
        (scm:let ([foldMap122 ((rt:record-ref dictFoldable10 (scm:string->symbol "foldMap1")) (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (j4)
              (((rt:record-ref dictSemigroup1 (scm:string->symbol "append")) (v2 j4)) (((rt:record-ref dictSemigroup1 (scm:string->symbol "append")) j4) (v13 j4)))))))))])
          (scm:lambda (j3)
            (scm:lambda (f4)
              (scm:lambda (foldable5)
                (((foldMap122 (scm:lambda (x6)
                  (scm:let ([_7 (f4 x6)])
                    (scm:lambda (v8)
                      _7)))) foldable5) j3))))))))

  (scm:define intercalate
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictSemigroup1)
        (scm:let ([foldMap122 ((rt:record-ref dictFoldable10 (scm:string->symbol "foldMap1")) (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (scm:lambda (j4)
              (((rt:record-ref dictSemigroup1 (scm:string->symbol "append")) (v2 j4)) (((rt:record-ref dictSemigroup1 (scm:string->symbol "append")) j4) (v13 j4)))))))))])
          (scm:lambda (a3)
            (scm:lambda (foldable4)
              (((foldMap122 (scm:lambda (x5)
                (scm:lambda (v6)
                  x5))) foldable4) a3)))))))

  (scm:define maximum
    (scm:lambda (dictOrd0)
      (scm:let ([semigroupMax1 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:let ([v3 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) v1) v12)])
            (scm:cond
              [(Data.Ordering.LT? v3) v12]
              [(Data.Ordering.EQ? v3) v1]
              [(Data.Ordering.GT? v3) v1]
              [scm:else (rt:fail)]))))))])
        (scm:lambda (dictFoldable12)
          (((rt:record-ref dictFoldable12 (scm:string->symbol "foldMap1")) semigroupMax1) Unsafe.Coerce.unsafeCoerce)))))

  (scm:define minimum
    (scm:lambda (dictOrd0)
      (scm:let ([semigroupMin1 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v1)
        (scm:lambda (v12)
          (scm:let ([v3 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) v1) v12)])
            (scm:cond
              [(Data.Ordering.LT? v3) v1]
              [(Data.Ordering.EQ? v3) v1]
              [(Data.Ordering.GT? v3) v12]
              [scm:else (rt:fail)]))))))])
        (scm:lambda (dictFoldable12)
          (((rt:record-ref dictFoldable12 (scm:string->symbol "foldMap1")) semigroupMin1) Unsafe.Coerce.unsafeCoerce)))))

  (scm:define traverse1_
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictApply1)
        (scm:let*
          ([_2 ((rt:record-ref dictApply1 (scm:string->symbol "Functor0")) (scm:quote undefined))]
           [foldMap123 ((rt:record-ref dictFoldable10 (scm:string->symbol "foldMap1")) (semigroupAct dictApply1))])
            (scm:lambda (f4)
              (scm:lambda (t5)
                (((rt:record-ref _2 (scm:string->symbol "map")) (scm:lambda (v6)
                  Data.Unit.unit)) ((foldMap123 (scm:lambda (x6)
                  (f4 x6))) t5))))))))

  (scm:define for1_
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictApply1)
        (scm:let ([_2 ((traverse1_ dictFoldable10) dictApply1)])
          (scm:lambda (b3)
            (scm:lambda (a4)
              ((_2 a4) b3)))))))

  (scm:define sequence1_
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictApply1)
        (((traverse1_ dictFoldable10) dictApply1) identity))))

  (scm:define fold1
    (scm:lambda (dictFoldable10)
      (scm:lambda (dictSemigroup1)
        (((rt:record-ref dictFoldable10 (scm:string->symbol "foldMap1")) dictSemigroup1) identity)))))
