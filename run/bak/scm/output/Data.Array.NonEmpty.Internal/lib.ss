#!r6rs
#!chezscheme
(library
  (Data.Array.NonEmpty.Internal lib)
  (export
    NonEmptyArray
    altNonEmptyArray
    applicativeNonEmptyArray
    applyNonEmptyArray
    bindNonEmptyArray
    eq1NonEmptyArray
    eqNonEmptyArray
    foldable1NonEmptyArray
    foldableNonEmptyArray
    foldableWithIndexNonEmptyArray
    foldl1Impl
    foldr1Impl
    functorNonEmptyArray
    functorWithIndexNonEmptyArray
    monadNonEmptyArray
    ord1NonEmptyArray
    ordNonEmptyArray
    semigroupNonEmptyArray
    showNonEmptyArray
    traversable1NonEmptyArray
    traversableNonEmptyArray
    traversableWithIndexNonEmptyArray
    traverse1Impl
    unfoldable1NonEmptyArray)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Alt lib) Control.Alt.)
    (prefix (Control.Applicative lib) Control.Applicative.)
    (prefix (Control.Apply lib) Control.Apply.)
    (prefix (Control.Bind lib) Control.Bind.)
    (prefix (Control.Monad lib) Control.Monad.)
    (prefix (Data.Eq lib) Data.Eq.)
    (prefix (Data.Foldable lib) Data.Foldable.)
    (prefix (Data.FoldableWithIndex lib) Data.FoldableWithIndex.)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.FunctorWithIndex lib) Data.FunctorWithIndex.)
    (prefix (Data.Ord lib) Data.Ord.)
    (prefix (Data.Semigroup lib) Data.Semigroup.)
    (prefix (Data.Semigroup.Traversable lib) Data.Semigroup.Traversable.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Data.Traversable lib) Data.Traversable.)
    (prefix (Data.TraversableWithIndex lib) Data.TraversableWithIndex.)
    (prefix (Data.Unfoldable1 lib) Data.Unfoldable1.)
    (Data.Array.NonEmpty.Internal foreign))

  (scm:define NonEmptyArray
    (scm:lambda (x0)
      x0))

  (scm:define unfoldable1NonEmptyArray
    Data.Unfoldable1.unfoldable1Array)

  (scm:define traversableWithIndexNonEmptyArray
    Data.TraversableWithIndex.traversableWithIndexArray)

  (scm:define traversableNonEmptyArray
    Data.Traversable.traversableArray)

  (scm:define showNonEmptyArray
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(NonEmptyArray ") ((Data.Show.showArrayImpl (rt:record-ref dictShow0 (scm:string->symbol "show"))) v1)) (rt:string->pstring ")")))))))

  (scm:define semigroupNonEmptyArray
    Data.Semigroup.semigroupArray)

  (scm:define ordNonEmptyArray
    (scm:lambda (dictOrd0)
      (Data.Ord.ordArray dictOrd0)))

  (scm:define ord1NonEmptyArray
    Data.Ord.ord1Array)

  (scm:define monadNonEmptyArray
    Control.Monad.monadArray)

  (scm:define functorWithIndexNonEmptyArray
    Data.FunctorWithIndex.functorWithIndexArray)

  (scm:define functorNonEmptyArray
    Data.Functor.functorArray)

  (scm:define foldableWithIndexNonEmptyArray
    Data.FoldableWithIndex.foldableWithIndexArray)

  (scm:define foldableNonEmptyArray
    Data.Foldable.foldableArray)

  (rt:define-lazy $lazy-foldable1NonEmptyArray "foldable1NonEmptyArray" "Data.Array.NonEmpty.Internal"
    (scm:list (scm:cons (scm:string->symbol "foldMap1") (scm:lambda (dictSemigroup0)
      (scm:let ([append1 (rt:record-ref dictSemigroup0 (scm:string->symbol "append"))])
        (scm:lambda (f2)
          (scm:let*
            ([_3 (Data.Functor.arrayMap f2)]
             [_4 ((rt:record-ref ($lazy-foldable1NonEmptyArray) (scm:string->symbol "foldl1")) append1)])
              (scm:lambda (x5)
                (_4 (_3 x5)))))))) (scm:cons (scm:string->symbol "foldr1") (scm:lambda (_0)
      (scm:lambda (_1)
        (foldr1Impl _0 _1)))) (scm:cons (scm:string->symbol "foldl1") (scm:lambda (_0)
      (scm:lambda (_1)
        (foldl1Impl _0 _1)))) (scm:cons (scm:string->symbol "Foldable0") (scm:lambda (_)
      Data.Foldable.foldableArray))))

  (scm:define foldable1NonEmptyArray
    ($lazy-foldable1NonEmptyArray))

  (rt:define-lazy $lazy-traversable1NonEmptyArray "traversable1NonEmptyArray" "Data.Array.NonEmpty.Internal"
    (scm:list (scm:cons (scm:string->symbol "traverse1") (scm:lambda (dictApply0)
      (scm:let*
        ([apply1 (rt:record-ref dictApply0 (scm:string->symbol "apply"))]
         [map2 (rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map"))])
          (scm:lambda (f3)
            (traverse1Impl apply1 map2 f3))))) (scm:cons (scm:string->symbol "sequence1") (scm:lambda (dictApply0)
      (((rt:record-ref ($lazy-traversable1NonEmptyArray) (scm:string->symbol "traverse1")) dictApply0) Data.Semigroup.Traversable.identity))) (scm:cons (scm:string->symbol "Foldable10") (scm:lambda (_)
      foldable1NonEmptyArray)) (scm:cons (scm:string->symbol "Traversable1") (scm:lambda (_)
      Data.Traversable.traversableArray))))

  (scm:define traversable1NonEmptyArray
    ($lazy-traversable1NonEmptyArray))

  (scm:define eqNonEmptyArray
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq") (Data.Eq.eqArrayImpl (rt:record-ref dictEq0 (scm:string->symbol "eq")))))))

  (scm:define eq1NonEmptyArray
    Data.Eq.eq1Array)

  (scm:define bindNonEmptyArray
    Control.Bind.bindArray)

  (scm:define applyNonEmptyArray
    Control.Apply.applyArray)

  (scm:define applicativeNonEmptyArray
    Control.Applicative.applicativeArray)

  (scm:define altNonEmptyArray
    Control.Alt.altArray))
