#!r6rs
#!chezscheme
(library
  (Data.Semigroup.Traversable lib)
  (export
    identity
    sequence1
    sequence1Default
    traversableDual
    traversableIdentity
    traversableMultiplicative
    traversableTuple
    traverse1
    traverse1Default)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Identity lib) Data.Identity.)
    (prefix (Data.Monoid.Dual lib) Data.Monoid.Dual.)
    (prefix (Data.Monoid.Multiplicative lib) Data.Monoid.Multiplicative.)
    (prefix (Data.Semigroup.Foldable lib) Data.Semigroup.Foldable.)
    (prefix (Data.Traversable lib) Data.Traversable.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define traverse1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "traverse1"))))

  (scm:define traversableTuple
    (scm:list (scm:cons (scm:string->symbol "traverse1") (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (Data.Tuple.Tuple (Data.Tuple.Tuple-value0 v2))) (f1 (Data.Tuple.Tuple-value1 v2))))))) (scm:cons (scm:string->symbol "sequence1") (scm:lambda (dictApply0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) (Data.Tuple.Tuple (Data.Tuple.Tuple-value0 v1))) (Data.Tuple.Tuple-value1 v1))))) (scm:cons (scm:string->symbol "Foldable10") (scm:lambda (_)
      Data.Semigroup.Foldable.foldableTuple)) (scm:cons (scm:string->symbol "Traversable1") (scm:lambda (_)
      Data.Traversable.traversableTuple))))

  (scm:define traversableIdentity
    (scm:list (scm:cons (scm:string->symbol "traverse1") (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Identity.Identity) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence1") (scm:lambda (dictApply0)
      (scm:lambda (v1)
        (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Identity.Identity) v1)))) (scm:cons (scm:string->symbol "Foldable10") (scm:lambda (_)
      Data.Semigroup.Foldable.foldableIdentity)) (scm:cons (scm:string->symbol "Traversable1") (scm:lambda (_)
      Data.Traversable.traversableIdentity))))

  (scm:define sequence1Default
    (scm:lambda (dictTraversable10)
      (scm:lambda (dictApply1)
        (((rt:record-ref dictTraversable10 (scm:string->symbol "traverse1")) dictApply1) identity))))

  (rt:define-lazy $lazy-traversableDual "traversableDual" "Data.Semigroup.Traversable"
    (scm:list (scm:cons (scm:string->symbol "traverse1") (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Dual.Dual) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence1") (scm:lambda (dictApply0)
      (((rt:record-ref ($lazy-traversableDual) (scm:string->symbol "traverse1")) dictApply0) identity))) (scm:cons (scm:string->symbol "Foldable10") (scm:lambda (_)
      Data.Semigroup.Foldable.foldableDual)) (scm:cons (scm:string->symbol "Traversable1") (scm:lambda (_)
      Data.Traversable.traversableDual))))

  (scm:define traversableDual
    ($lazy-traversableDual))

  (rt:define-lazy $lazy-traversableMultiplicative "traversableMultiplicative" "Data.Semigroup.Traversable"
    (scm:list (scm:cons (scm:string->symbol "traverse1") (scm:lambda (dictApply0)
      (scm:lambda (f1)
        (scm:lambda (v2)
          (((rt:record-ref ((rt:record-ref dictApply0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) Data.Monoid.Multiplicative.Multiplicative) (f1 v2)))))) (scm:cons (scm:string->symbol "sequence1") (scm:lambda (dictApply0)
      (((rt:record-ref ($lazy-traversableMultiplicative) (scm:string->symbol "traverse1")) dictApply0) identity))) (scm:cons (scm:string->symbol "Foldable10") (scm:lambda (_)
      Data.Semigroup.Foldable.foldableMultiplicative)) (scm:cons (scm:string->symbol "Traversable1") (scm:lambda (_)
      Data.Traversable.traversableMultiplicative))))

  (scm:define traversableMultiplicative
    ($lazy-traversableMultiplicative))

  (scm:define sequence1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "sequence1"))))

  (scm:define traverse1Default
    (scm:lambda (dictTraversable10)
      (scm:lambda (dictApply1)
        (scm:let ([sequence122 ((rt:record-ref dictTraversable10 (scm:string->symbol "sequence1")) dictApply1)])
          (scm:lambda (f3)
            (scm:lambda (ta4)
              (sequence122 (((rt:record-ref ((rt:record-ref ((rt:record-ref dictTraversable10 (scm:string->symbol "Traversable1")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) f3) ta4)))))))))
