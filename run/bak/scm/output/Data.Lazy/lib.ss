#!r6rs
#!chezscheme
(library
  (Data.Lazy lib)
  (export
    Lazy
    applicativeLazy
    applyLazy
    bindLazy
    booleanAlgebraLazy
    boundedLazy
    commutativeRingLazy
    comonadLazy
    defer
    eq1Lazy
    eqLazy
    euclideanRingLazy
    extendLazy
    foldable1Lazy
    foldableLazy
    foldableWithIndexLazy
    force
    functorLazy
    functorWithIndexLazy
    heytingAlgebraLazy
    invariantLazy
    lazyLazy
    monadLazy
    monoidLazy
    ord1Lazy
    ordLazy
    ringLazy
    semigroupLazy
    semiringLazy
    showLazy
    traversable1Lazy
    traversableLazy
    traversableWithIndexLazy)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define Lazy
    (scm:lambda (x0)
      x0))

  (scm:define force
    (scm:lambda (v0)
      (v0 Data.Unit.unit)))

  (scm:define showLazy
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (x1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(defer \\_ -> ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (x1 Data.Unit.unit))) (rt:string->pstring ")")))))))

  (scm:define foldableLazy
    (scm:list (scm:cons (scm:string->symbol "foldr") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (l2)
          ((f0 (l2 Data.Unit.unit)) z1))))) (scm:cons (scm:string->symbol "foldl") (scm:lambda (f0)
      (scm:lambda (z1)
        (scm:lambda (l2)
          ((f0 z1) (l2 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "foldMap") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (scm:lambda (l2)
          (f1 (l2 Data.Unit.unit))))))))

  (scm:define foldableWithIndexLazy
    (scm:list (scm:cons (scm:string->symbol "foldrWithIndex") (scm:lambda (f0)
      ((rt:record-ref foldableLazy (scm:string->symbol "foldr")) (f0 Data.Unit.unit)))) (scm:cons (scm:string->symbol "foldlWithIndex") (scm:lambda (f0)
      ((rt:record-ref foldableLazy (scm:string->symbol "foldl")) (f0 Data.Unit.unit)))) (scm:cons (scm:string->symbol "foldMapWithIndex") (scm:lambda (dictMonoid0)
      (scm:lambda (f1)
        (((rt:record-ref foldableLazy (scm:string->symbol "foldMap")) dictMonoid0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      foldableLazy))))

  (scm:define foldable1Lazy
    (scm:list (scm:cons (scm:string->symbol "foldMap1") (scm:lambda (dictSemigroup0)
      (scm:lambda (f1)
        (scm:lambda (l2)
          (f1 (l2 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "foldr1") (scm:lambda (v0)
      (scm:lambda (l1)
        (l1 Data.Unit.unit)))) (scm:cons (scm:string->symbol "foldl1") (scm:lambda (v0)
      (scm:lambda (l1)
        (l1 Data.Unit.unit)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      foldableLazy))))

  (scm:define eqLazy
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x1)
        (scm:lambda (y2)
          (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (x1 Data.Unit.unit)) (y2 Data.Unit.unit))))))))

  (scm:define ordLazy
    (scm:lambda (dictOrd0)
      (scm:let*
        ([_1 ((rt:record-ref dictOrd0 (scm:string->symbol "Eq0")) (scm:quote undefined))]
         [eqLazy12 (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (x2)
          (scm:lambda (y3)
            (((rt:record-ref _1 (scm:string->symbol "eq")) (x2 Data.Unit.unit)) (y3 Data.Unit.unit))))))])
          (scm:list (scm:cons (scm:string->symbol "compare") (scm:lambda (x3)
            (scm:lambda (y4)
              (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (x3 Data.Unit.unit)) (y4 Data.Unit.unit))))) (scm:cons (scm:string->symbol "Eq0") (scm:lambda (_)
            eqLazy12))))))

  (scm:define eq1Lazy
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (y2)
          (((rt:record-ref dictEq0 (scm:string->symbol "eq")) (x1 Data.Unit.unit)) (y2 Data.Unit.unit))))))))

  (scm:define ord1Lazy
    (scm:list (scm:cons (scm:string->symbol "compare1") (scm:lambda (dictOrd0)
      (rt:record-ref (ordLazy dictOrd0) (scm:string->symbol "compare")))) (scm:cons (scm:string->symbol "Eq10") (scm:lambda (_)
      eq1Lazy))))

  (scm:define defer
    (scm:lambda (f0)
      f0))

  (scm:define functorLazy
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (l1)
        (scm:lambda (v2)
          (f0 (l1 Data.Unit.unit))))))))

  (scm:define extendLazy
    (scm:list (scm:cons (scm:string->symbol "extend") (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (v2)
          (f0 x1))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorLazy))))

  (scm:define functorWithIndexLazy
    (scm:list (scm:cons (scm:string->symbol "mapWithIndex") (scm:lambda (f0)
      ((rt:record-ref functorLazy (scm:string->symbol "map")) (f0 Data.Unit.unit)))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorLazy))))

  (scm:define invariantLazy
    (scm:list (scm:cons (scm:string->symbol "imap") (scm:lambda (f0)
      (scm:lambda (v1)
        ((rt:record-ref functorLazy (scm:string->symbol "map")) f0))))))

  (scm:define lazyLazy
    (scm:list (scm:cons (scm:string->symbol "defer") (scm:lambda (f0)
      (scm:lambda (v1)
        ((f0 Data.Unit.unit) Data.Unit.unit))))))

  (scm:define semigroupLazy
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (a1)
        (scm:lambda (b2)
          (scm:lambda (v3)
            (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (a1 Data.Unit.unit)) (b2 Data.Unit.unit)))))))))

  (scm:define monoidLazy
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))]
         [_2 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupLazy13 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (a3)
          (scm:lambda (b4)
            (scm:lambda (v5)
              (((rt:record-ref _2 (scm:string->symbol "append")) (a3 Data.Unit.unit)) (b4 Data.Unit.unit)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:lambda (v4)
            mempty1)) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupLazy13))))))

  (scm:define semiringLazy
    (scm:lambda (dictSemiring0)
      (scm:let*
        ([zero1 (rt:record-ref dictSemiring0 (scm:string->symbol "zero"))]
         [one2 (rt:record-ref dictSemiring0 (scm:string->symbol "one"))])
          (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:lambda (v5)
                (((rt:record-ref dictSemiring0 (scm:string->symbol "add")) (a3 Data.Unit.unit)) (b4 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "zero") (scm:lambda (v3)
            zero1)) (scm:cons (scm:string->symbol "mul") (scm:lambda (a3)
            (scm:lambda (b4)
              (scm:lambda (v5)
                (((rt:record-ref dictSemiring0 (scm:string->symbol "mul")) (a3 Data.Unit.unit)) (b4 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "one") (scm:lambda (v3)
            one2))))))

  (scm:define ringLazy
    (scm:lambda (dictRing0)
      (scm:let ([semiringLazy11 (semiringLazy ((rt:record-ref dictRing0 (scm:string->symbol "Semiring0")) (scm:quote undefined)))])
        (scm:list (scm:cons (scm:string->symbol "sub") (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (v4)
              (((rt:record-ref dictRing0 (scm:string->symbol "sub")) (a2 Data.Unit.unit)) (b3 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
          semiringLazy11))))))

  (scm:define traversableLazy
    (scm:list (scm:cons (scm:string->symbol "traverse") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (scm:lambda (l2)
          (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (x3)
            (scm:lambda (v4)
              x3))) (f1 (l2 Data.Unit.unit))))))) (scm:cons (scm:string->symbol "sequence") (scm:lambda (dictApplicative0)
      (scm:lambda (l1)
        (((rt:record-ref ((rt:record-ref ((rt:record-ref dictApplicative0 (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (x2)
          (scm:lambda (v3)
            x2))) (l1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorLazy)) (scm:cons (scm:string->symbol "Foldable1") (scm:lambda (_)
      foldableLazy))))

  (scm:define traversable1Lazy
    (scm:list (scm:cons (scm:string->symbol "traverse1") (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (l2)
          (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (x3)
            (scm:lambda (v4)
              x3))) (f1 (l2 Data.Unit.unit))))))) (scm:cons (scm:string->symbol "sequence1") (scm:lambda (dictApply0)
      (scm:lambda (l1)
        (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (scm:lambda (x2)
          (scm:lambda (v3)
            x2))) (l1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "Foldable10") (scm:lambda (_)
      foldable1Lazy)) (scm:cons (scm:string->symbol "Traversable1") (scm:lambda (_)
      traversableLazy))))

  (scm:define traversableWithIndexLazy
    (scm:list (scm:cons (scm:string->symbol "traverseWithIndex") (scm:lambda (dictApplicative0)
      (scm:lambda (f1)
        (((rt:record-ref traversableLazy (scm:string->symbol "traverse")) dictApplicative0) (f1 Data.Unit.unit))))) (scm:cons (scm:string->symbol "FunctorWithIndex0") (scm:lambda (_)
      functorWithIndexLazy)) (scm:cons (scm:string->symbol "FoldableWithIndex1") (scm:lambda (_)
      foldableWithIndexLazy)) (scm:cons (scm:string->symbol "Traversable2") (scm:lambda (_)
      traversableLazy))))

  (scm:define comonadLazy
    (scm:list (scm:cons (scm:string->symbol "extract") force) (scm:cons (scm:string->symbol "Extend0") (scm:lambda (_)
      extendLazy))))

  (scm:define commutativeRingLazy
    (scm:lambda (dictCommutativeRing0)
      (scm:let ([ringLazy11 (ringLazy ((rt:record-ref dictCommutativeRing0 (scm:string->symbol "Ring0")) (scm:quote undefined)))])
        (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
          ringLazy11))))))

  (scm:define euclideanRingLazy
    (scm:lambda (dictEuclideanRing0)
      (scm:let ([ringLazy11 (ringLazy ((rt:record-ref ((rt:record-ref dictEuclideanRing0 (scm:string->symbol "CommutativeRing0")) (scm:quote undefined)) (scm:string->symbol "Ring0")) (scm:quote undefined)))])
        (scm:list (scm:cons (scm:string->symbol "degree") (scm:lambda (x2)
          ((rt:record-ref dictEuclideanRing0 (scm:string->symbol "degree")) (x2 Data.Unit.unit)))) (scm:cons (scm:string->symbol "div") (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (v4)
              (((rt:record-ref dictEuclideanRing0 (scm:string->symbol "div")) (a2 Data.Unit.unit)) (b3 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "mod") (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (v4)
              (((rt:record-ref dictEuclideanRing0 (scm:string->symbol "mod")) (a2 Data.Unit.unit)) (b3 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "CommutativeRing0") (scm:lambda (_)
          (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
            ringLazy11)))))))))

  (scm:define boundedLazy
    (scm:lambda (dictBounded0)
      (scm:let*
        ([top1 (rt:record-ref dictBounded0 (scm:string->symbol "top"))]
         [bottom2 (rt:record-ref dictBounded0 (scm:string->symbol "bottom"))]
         [ordLazy13 (ordLazy ((rt:record-ref dictBounded0 (scm:string->symbol "Ord0")) (scm:quote undefined)))])
          (scm:list (scm:cons (scm:string->symbol "top") (scm:lambda (v4)
            top1)) (scm:cons (scm:string->symbol "bottom") (scm:lambda (v4)
            bottom2)) (scm:cons (scm:string->symbol "Ord0") (scm:lambda (_)
            ordLazy13))))))

  (scm:define applyLazy
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (f0)
      (scm:lambda (x1)
        (scm:lambda (v2)
          ((f0 Data.Unit.unit) (x1 Data.Unit.unit)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorLazy))))

  (scm:define bindLazy
    (scm:list (scm:cons (scm:string->symbol "bind") (scm:lambda (l0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          ((f1 (l0 Data.Unit.unit)) Data.Unit.unit))))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyLazy))))

  (scm:define heytingAlgebraLazy
    (scm:lambda (dictHeytingAlgebra0)
      (scm:let*
        ([ff1 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "ff"))]
         [tt2 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "tt"))]
         [implies3 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "implies"))]
         [conj4 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "conj"))]
         [disj5 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "disj"))]
         [not6 (rt:record-ref dictHeytingAlgebra0 (scm:string->symbol "not"))])
          (scm:list (scm:cons (scm:string->symbol "ff") (scm:lambda (v7)
            ff1)) (scm:cons (scm:string->symbol "tt") (scm:lambda (v7)
            tt2)) (scm:cons (scm:string->symbol "implies") (scm:lambda (a7)
            (scm:lambda (b8)
              (((rt:record-ref applyLazy (scm:string->symbol "apply")) (((rt:record-ref functorLazy (scm:string->symbol "map")) implies3) a7)) b8)))) (scm:cons (scm:string->symbol "conj") (scm:lambda (a7)
            (scm:lambda (b8)
              (((rt:record-ref applyLazy (scm:string->symbol "apply")) (((rt:record-ref functorLazy (scm:string->symbol "map")) conj4) a7)) b8)))) (scm:cons (scm:string->symbol "disj") (scm:lambda (a7)
            (scm:lambda (b8)
              (((rt:record-ref applyLazy (scm:string->symbol "apply")) (((rt:record-ref functorLazy (scm:string->symbol "map")) disj5) a7)) b8)))) (scm:cons (scm:string->symbol "not") (scm:lambda (a7)
            (((rt:record-ref functorLazy (scm:string->symbol "map")) not6) a7)))))))

  (scm:define booleanAlgebraLazy
    (scm:lambda (dictBooleanAlgebra0)
      (scm:let ([heytingAlgebraLazy11 (heytingAlgebraLazy ((rt:record-ref dictBooleanAlgebra0 (scm:string->symbol "HeytingAlgebra0")) (scm:quote undefined)))])
        (scm:list (scm:cons (scm:string->symbol "HeytingAlgebra0") (scm:lambda (_)
          heytingAlgebraLazy11))))))

  (scm:define applicativeLazy
    (scm:list (scm:cons (scm:string->symbol "pure") (scm:lambda (a0)
      (scm:lambda (v1)
        a0))) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      applyLazy))))

  (scm:define monadLazy
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      applicativeLazy)) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      bindLazy)))))
