#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Strong lib)
  (export
    fanout
    first
    identity
    second
    splitStrong
    strongFn)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Profunctor lib) Data.Profunctor.)
    (prefix (Data.Tuple lib) Data.Tuple.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define strongFn
    (scm:list (scm:cons (scm:string->symbol "first") (scm:lambda (a2b0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (a2b0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value1 v1))))) (scm:cons (scm:string->symbol "second") (rt:record-ref Data.Tuple.functorTuple (scm:string->symbol "map"))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
      Data.Profunctor.profunctorFn))))

  (scm:define second
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "second"))))

  (scm:define first
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "first"))))

  (scm:define splitStrong
    (scm:lambda (dictCategory0)
      (scm:let ([_1 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))])
        (scm:lambda (dictStrong2)
          (scm:lambda (l3)
            (scm:lambda (r4)
              (((rt:record-ref _1 (scm:string->symbol "compose")) ((rt:record-ref dictStrong2 (scm:string->symbol "second")) r4)) ((rt:record-ref dictStrong2 (scm:string->symbol "first")) l3))))))))

  (scm:define fanout
    (scm:lambda (dictCategory0)
      (scm:let*
        ([identity11 (rt:record-ref dictCategory0 (scm:string->symbol "identity"))]
         [_2 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))]
         [_3 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))])
          (scm:lambda (dictStrong4)
            (scm:lambda (l5)
              (scm:lambda (r6)
                (((rt:record-ref _2 (scm:string->symbol "compose")) (((rt:record-ref _3 (scm:string->symbol "compose")) ((rt:record-ref dictStrong4 (scm:string->symbol "second")) r6)) ((rt:record-ref dictStrong4 (scm:string->symbol "first")) l5))) ((((rt:record-ref ((rt:record-ref dictStrong4 (scm:string->symbol "Profunctor0")) (scm:quote undefined)) (scm:string->symbol "dimap")) identity) (scm:lambda (a7)
                  (Data.Tuple.Tuple* a7 a7))) identity11)))))))))
