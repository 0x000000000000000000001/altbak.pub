#!r6rs
#!chezscheme
(library
  (Test.ArrayOps lib)
  (export
    act
    describe
    filterEvens
    range
    sumEvens)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Array lib) Data.Array.)
    (prefix (Data.EuclideanRing lib) Data.EuclideanRing.)
    (prefix (Data.Foldable lib) Data.Foldable.)
    (prefix (Data.Semiring lib) Data.Semiring.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define range
    (scm:lambda (start0)
      (scm:lambda (end1)
        (Data.Array.rangeImpl start0 end1))))

  (scm:define filterEvens
    (scm:lambda (arr0)
      (Data.Array.filterImpl (scm:lambda (x1)
        (scm:fx=? ((Data.EuclideanRing.intMod x1) 2) 0)) arr0)))

  (scm:define sumEvens
    (scm:lambda (n0)
      (((Data.Foldable.foldlArray Data.Semiring.intAdd) 0) (Data.Array.filterImpl (scm:lambda (x1)
        (scm:fx=? ((Data.EuclideanRing.intMod x1) 2) 0)) (Data.Array.rangeImpl 1 n0)))))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Array Processing (900 elements):")))

  (scm:define act
    (scm:let ([_0 (Bench.opaque 900)])
      (scm:lambda ()
        (scm:let ([dummy1 (_0)])
          ((Effect.Console.log (Data.Show.showIntImpl (sumEvens dummy1)))))))))
