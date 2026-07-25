#!r6rs
#!chezscheme
(library
  (Data.Tuple lib)
  (export
    Tuple
    Tuple*
    Tuple-value0
    Tuple-value1
    Tuple?
    applicativeTuple
    applyTuple
    bindTuple
    booleanAlgebraTuple
    boundedTuple
    commutativeRingTuple
    comonadTuple
    curry
    eq1Tuple
    eqTuple
    extendTuple
    fst
    functorTuple
    genericTuple
    heytingAlgebraTuple
    invariantTuple
    lazyTuple
    monadTuple
    monoidTuple
    ord1Tuple
    ordTuple
    ringTuple
    semigroupTuple
    semigroupoidTuple
    semiringTuple
    showTuple
    snd
    swap
    uncurry)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define-record-type (Tuple$ Tuple* Tuple?)
    (scm:fields (scm:immutable value0 Tuple-value0) (scm:immutable value1 Tuple-value1)))

  (scm:define Tuple
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Tuple* value0 value1))))

  (scm:define uncurry
    (scm:lambda (f0)
      (scm:lambda (v1)
        ((f0 (Tuple-value0 v1)) (Tuple-value1 v1)))))

  (scm:define swap
    (scm:lambda (v0)
      (Tuple* (Tuple-value1 v0) (Tuple-value0 v0))))

  (scm:define snd
    (scm:lambda (v0)
      (Tuple-value1 v0)))

  (scm:define showTuple
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Tuple ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Tuple-value0 v2))) (rt:string->pstring " ")) ((rt:record-ref dictShow11 (scm:string->symbol "show")) (Tuple-value1 v2))) (rt:string->pstring ")"))))))))

  (scm:define semiringTuple
    (scm:lambda (dictSemiring0)
      (scm:let*
        ([one1 (rt:record-ref dictSemiring0 (scm:string->symbol "one"))]
         [zero2 (rt:record-ref dictSemiring0 (scm:string->symbol "zero"))])
          (scm:lambda (dictSemiring13)
            (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (v4)
              (scm:lambda (v15)
                (Tuple* (((rt:record-ref dictSemiring0 (scm:string->symbol "add")) (Tuple-value0 v4)) (Tuple-value0 v15)) (((rt:record-ref dictSemiring13 (scm:string->symbol "add")) (Tuple-value1 v4)) (Tuple-value1 v15)))))) (scm:cons (scm:string->symbol "one") (Tuple* one1 (rt:record-ref dictSemiring13 (scm:string->symbol "one")))) (scm:cons (scm:string->symbol "mul") (scm:lambda (v4)
              (scm:lambda (v15)
                (Tuple* (((rt:record-ref dictSemiring0 (scm:string->symbol "mul")) (Tuple-value0 v4)) (Tuple-value0 v15)) (((rt:record-ref dictSemiring13 (scm:string->symbol "mul")) (Tuple-value1 v4)) (Tuple-value1 v15)))))) (scm:cons (scm:string->symbol "zero") (Tuple* zero2 (rt:record-ref dictSemiring13 (scm:string->symbol "zero")))))))))

  (scm:define semigroupoidTuple
    (scm:list (scm:cons (scm:string->symbol "compose") (scm:lambda (v0)
      (scm:lambda (v11)
        (Tuple* (Tuple-value0 v11) (Tuple-value1 v0)))))))

  (scm:define semigroupTuple
    (scm:lambda (dictSemigroup0)
      (scm:lambda (dictSemigroup11)
        (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v2)
          (scm:lambda (v13)
            (Tuple* (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (Tuple-value0 v2)) (Tuple-value0 v13)) (((rt:record-ref dictSemigroup11 (scm:string->symbol "append")) (Tuple-value1 v2)) (Tuple-value1 v13))))))))))

  (scm:define ringTuple
    (scm:lambda (dictRing0)
      (scm:let*
        ([_1 ((rt:record-ref dictRing0 (scm:string->symbol "Semiring0")) (scm:quote undefined))]
         [one2 (rt:record-ref _1 (scm:string->symbol "one"))]
         [semiringTuple13 (scm:let ([zero3 (rt:record-ref _1 (scm:string->symbol "zero"))])
          (scm:lambda (dictSemiring14)
            (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (v5)
              (scm:lambda (v16)
                (Tuple* (((rt:record-ref _1 (scm:string->symbol "add")) (Tuple-value0 v5)) (Tuple-value0 v16)) (((rt:record-ref dictSemiring14 (scm:string->symbol "add")) (Tuple-value1 v5)) (Tuple-value1 v16)))))) (scm:cons (scm:string->symbol "one") (Tuple* one2 (rt:record-ref dictSemiring14 (scm:string->symbol "one")))) (scm:cons (scm:string->symbol "mul") (scm:lambda (v5)
              (scm:lambda (v16)
                (Tuple* (((rt:record-ref _1 (scm:string->symbol "mul")) (Tuple-value0 v5)) (Tuple-value0 v16)) (((rt:record-ref dictSemiring14 (scm:string->symbol "mul")) (Tuple-value1 v5)) (Tuple-value1 v16)))))) (scm:cons (scm:string->symbol "zero") (Tuple* zero3 (rt:record-ref dictSemiring14 (scm:string->symbol "zero")))))))])
          (scm:lambda (dictRing14)
            (scm:let ([semiringTuple25 (semiringTuple13 ((rt:record-ref dictRing14 (scm:string->symbol "Semiring0")) (scm:quote undefined)))])
              (scm:list (scm:cons (scm:string->symbol "sub") (scm:lambda (v6)
                (scm:lambda (v17)
                  (Tuple* (((rt:record-ref dictRing0 (scm:string->symbol "sub")) (Tuple-value0 v6)) (Tuple-value0 v17)) (((rt:record-ref dictRing14 (scm:string->symbol "sub")) (Tuple-value1 v6)) (Tuple-value1 v17)))))) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
                semiringTuple25))))))))

  (scm:define monoidTuple
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))])
          (scm:lambda (dictMonoid13)
            (scm:let*
              ([_4 ((rt:record-ref dictMonoid13 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
               [semigroupTuple25 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v5)
                (scm:lambda (v16)
                  (Tuple* (((rt:record-ref _2 (scm:string->symbol "append")) (Tuple-value0 v5)) (Tuple-value0 v16)) (((rt:record-ref _4 (scm:string->symbol "append")) (Tuple-value1 v5)) (Tuple-value1 v16)))))))])
                (scm:list (scm:cons (scm:string->symbol "mempty") (Tuple* mempty1 (rt:record-ref dictMonoid13 (scm:string->symbol "mempty")))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
                  semigroupTuple25))))))))

  (scm:define heytingAlgebraTuple
    (scm:lambda (dictHeytingAlgebra0)
      (scm:let*
        ([tt1 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "tt"))]
         [ff2 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "ff"))])
          (scm:lambda (dictHeytingAlgebra13)
            (scm:list (scm:cons (scm:string->symbol "tt") (Tuple* tt1 (rt:record-ref dictHeytingAlgebra13 (scm:string->symbol "tt")))) (scm:cons (scm:string->symbol "ff") (Tuple* ff2 (rt:record-ref dictHeytingAlgebra13 (scm:string->symbol "ff")))) (scm:cons (scm:string->symbol "implies") (scm:lambda (v4)
              (scm:lambda (v15)
                (Tuple* (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "implies")) (Tuple-value0 v4)) (Tuple-value0 v15)) (((rt:record-ref dictHeytingAlgebra13 (scm:string->symbol "implies")) (Tuple-value1 v4)) (Tuple-value1 v15)))))) (scm:cons (scm:string->symbol "conj") (scm:lambda (v4)
              (scm:lambda (v15)
                (Tuple* (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "conj")) (Tuple-value0 v4)) (Tuple-value0 v15)) (((rt:record-ref dictHeytingAlgebra13 (scm:string->symbol "conj")) (Tuple-value1 v4)) (Tuple-value1 v15)))))) (scm:cons (scm:string->symbol "disj") (scm:lambda (v4)
              (scm:lambda (v15)
                (Tuple* (((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "disj")) (Tuple-value0 v4)) (Tuple-value0 v15)) (((rt:record-ref dictHeytingAlgebra13 (scm:string->symbol "disj")) (Tuple-value1 v4)) (Tuple-value1 v15)))))) (scm:cons (scm:string->symbol "not") (scm:lambda (v4)
              (Tuple* ((rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "not")) (Tuple-value0 v4)) ((rt:record-ref dictHeytingAlgebra13 (scm:string->symbol "not")) (Tuple-value1 v4))))))))))

  (scm:define genericTuple
    (scm:list (scm:cons (scm:string->symbol "to") (scm:lambda (x0)
      (Tuple* (Data.Generic.Rep.Product-value0 x0) (Data.Generic.Rep.Product-value1 x0)))) (scm:cons (scm:string->symbol "from") (scm:lambda (x0)
      (Data.Generic.Rep.Product* (Tuple-value0 x0) (Tuple-value1 x0))))))

  (scm:define functorTuple
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (Tuple* (Tuple-value0 m1) (f0 (Tuple-value1 m1))))))))

  (scm:define invariantTuple
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        (scm:lambda (m2)
          (Tuple* (Tuple-value0 m2) (f0 (Tuple-value1 m2)))))))))

  (scm:define fst
    (scm:lambda (v0)
      (Tuple-value0 v0)))

  (scm:define lazyTuple
    (scm:lambda (dictLazy0)
      (scm:lambda (dictLazy11)
        (scm:list (scm:cons (scm:string->symbol "defer") (scm:lambda (f2)
          (Tuple* ((rt:record-ref dictLazy0 (scm:string->symbol "defer")) (scm:lambda (v3)
            (Tuple-value0 (f2 Data.Unit.unit)))) ((rt:record-ref dictLazy11 (scm:string->symbol "defer")) (scm:lambda (v3)
            (Tuple-value1 (f2 Data.Unit.unit)))))))))))

  (scm:define extendTuple
    (scm:list (scm:cons (scm:string->symbol "extend") (scm:lambda (f0)
      (scm:lambda (v1)
        (Tuple* (Tuple-value0 v1) (f0 v1))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorTuple))))

  (scm:define eqTuple
    (scm:lambda (dictEq0)
      (scm:lambda (dictEq11)
        (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:and (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Tuple-value0 x2)) (Tuple-value0 y3)) (((rt:record-ref dictEq11 (scm:string->symbol "eq")) (Tuple-value1 x2)) (Tuple-value1 y3))))))))))

  (scm:define ordTuple
    (scm:lambda (dictOrd0)
      (scm:let ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))])
        (scm:lambda (dictOrd12)
          (scm:let*
            ([_3 ((rt:record-ref dictOrd12 (scm:string->symbol "Eq0")) (scm:quote undefined))]
             [eqTuple24 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x4)
              (scm:lambda (y5)
                (scm:and (((rt:record-ref _1 (scm:string->symbol "eq")) (Tuple-value0 x4)) (Tuple-value0 y5)) (((rt:record-ref _3 (scm:string->symbol "eq")) (Tuple-value1 x4)) (Tuple-value1 y5)))))))])
              (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x5)
                (scm:lambda (y6)
                  (scm:let ([v7 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Tuple-value0 x5)) (Tuple-value0 y6))])
                    (scm:cond
                      [(Data.Ordering.LT? v7) Data.Ordering.LT]
                      [(Data.Ordering.GT? v7) Data.Ordering.GT]
                      [scm:else (((rt:record-ref dictOrd12 (scm:string->symbol "compare")) (Tuple-value1 x5)) (Tuple-value1 y6))]))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                eqTuple24))))))))

  (scm:define eq1Tuple
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq11)
        (scm:lambda (x2)
          (scm:lambda (y3)
            (scm:and (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (Tuple-value0 x2)) (Tuple-value0 y3)) (((rt:record-ref dictEq11 (scm:string->symbol "eq")) (Tuple-value1 x2)) (Tuple-value1 y3))))))))))

  (scm:define ord1Tuple
    (scm:lambda (dictOrd0)
      (scm:let*
        ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))]
         [eq1Tuple12 (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq12)
          (scm:lambda (x3)
            (scm:lambda (y4)
              (scm:and (((rt:record-ref _1 (scm:string->symbol "eq")) (Tuple-value0 x3)) (Tuple-value0 y4)) (((rt:record-ref dictEq12 (scm:string->symbol "eq")) (Tuple-value1 x3)) (Tuple-value1 y4))))))))])
          (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd13)
            (scm:lambda (x4)
              (scm:lambda (y5)
                (scm:let ([v6 (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (Tuple-value0 x4)) (Tuple-value0 y5))])
                  (scm:cond
                    [(Data.Ordering.LT? v6) Data.Ordering.LT]
                    [(Data.Ordering.GT? v6) Data.Ordering.GT]
                    [scm:else (((rt:record-ref dictOrd13 (scm:string->symbol "compare")) (Tuple-value1 x4)) (Tuple-value1 y5))])))))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
            eq1Tuple12))))))

  (scm:define curry
    (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (f0 (Tuple* a1 b2))))))

  (scm:define comonadTuple
    (scm:list (scm:cons (scm:string->symbol "extract") snd) (scm:cons (scm:string->symbol "Extend0") (scm:lambda (_)
      extendTuple))))

  (scm:define commutativeRingTuple
    (scm:lambda (dictCommutativeRing0)
      (scm:let ([ringTuple11 (ringTuple ((rt:record-ref dictCommutativeRing0 (scm:string->symbol "Ring0")) (scm:quote undefined)))])
        (scm:lambda (dictCommutativeRing12)
          (scm:let ([ringTuple23 (ringTuple11 ((rt:record-ref dictCommutativeRing12 (scm:string->symbol "Ring0")) (scm:quote undefined)))])
            (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
              ringTuple23))))))))

  (scm:define boundedTuple
    (scm:lambda (dictBounded0)
      (scm:let*
        ([top1 (rt:record-ref dictBounded0 (scm:string->symbol "top"))]
         [bottom2 (rt:record-ref dictBounded0 (scm:string->symbol "bottom"))]
         [_3 ((rt:record-ref dictBounded0 (scm:string->symbol "Ord0")) (scm:quote undefined))]
         [_4 ((rt:record-ref _3 (scm:string->symbol "Eq0")) (scm:quote undefined))])
          (scm:lambda (dictBounded15)
            (scm:let*
              ([_6 ((rt:record-ref dictBounded15 (scm:string->symbol "Ord0")) (scm:quote undefined))]
               [_7 ((rt:record-ref _6 (scm:string->symbol "Eq0")) (scm:quote undefined))]
               [ordTuple28 (scm:let ([eqTuple28 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x8)
                (scm:lambda (y9)
                  (scm:and (((rt:record-ref _4 (scm:string->symbol "eq")) (Tuple-value0 x8)) (Tuple-value0 y9)) (((rt:record-ref _7 (scm:string->symbol "eq")) (Tuple-value1 x8)) (Tuple-value1 y9)))))))])
                (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x9)
                  (scm:lambda (y10)
                    (scm:let ([v11 (((rt:record-ref _3 (scm:string->symbol "compare")) (Tuple-value0 x9)) (Tuple-value0 y10))])
                      (scm:cond
                        [(Data.Ordering.LT? v11) Data.Ordering.LT]
                        [(Data.Ordering.GT? v11) Data.Ordering.GT]
                        [scm:else (((rt:record-ref _6 (scm:string->symbol "compare")) (Tuple-value1 x9)) (Tuple-value1 y10))]))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
                  eqTuple28))))])
                (scm:list (scm:cons (scm:string->symbol "top") (Tuple* top1 (rt:record-ref dictBounded15 (scm:string->symbol "top")))) (scm:cons (scm:string->symbol "bottom") (Tuple* bottom2 (rt:record-ref dictBounded15 (scm:string->symbol "bottom")))) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
                  ordTuple28))))))))

  (scm:define booleanAlgebraTuple
    (scm:lambda (dictBooleanAlgebra0)
      (scm:let ([heytingAlgebraTuple11 (heytingAlgebraTuple ((rt:record-ref dictBooleanAlgebra0 (scm:string->symbol "HeytingAlgebra0")) (scm:quote undefined)))])
        (scm:lambda (dictBooleanAlgebra12)
          (scm:let ([heytingAlgebraTuple23 (heytingAlgebraTuple11 ((rt:record-ref dictBooleanAlgebra12 (scm:string->symbol "HeytingAlgebra0")) (scm:quote undefined)))])
            (scm:list (scm:cons (scm:string->symbol "HeytingAlgebra0") (scm:lambda (_)
              heytingAlgebraTuple23))))))))

  (scm:define applyTuple
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (v1)
        (scm:lambda (v12)
          (Tuple* (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (Tuple-value0 v1)) (Tuple-value0 v12)) ((Tuple-value1 v1) (Tuple-value1 v12)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
        functorTuple)))))

  (scm:define bindTuple
    (scm:lambda (dictSemigroup0)
      (scm:let ([applyTuple11 (applyTuple dictSemigroup0)])
        (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (v2)
          (scm:lambda (f3)
            (scm:let ([v14 (f3 (Tuple-value1 v2))])
              (Tuple* (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (Tuple-value0 v2)) (Tuple-value0 v14)) (Tuple-value1 v14)))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
          applyTuple11))))))

  (scm:define applicativeTuple
    (scm:lambda (dictMonoid0)
      (scm:let ([applyTuple11 (applyTuple ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)))])
        (scm:list (scm:cons (scm:string->symbol "pure") (Tuple (rt:record-ref dictMonoid0 (scm:string->symbol "mempty")))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
          applyTuple11))))))

  (scm:define monadTuple
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([applicativeTuple11 (applicativeTuple dictMonoid0)]
         [bindTuple12 (bindTuple ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)))])
          (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
            applicativeTuple11)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
            bindTuple12)))))))
