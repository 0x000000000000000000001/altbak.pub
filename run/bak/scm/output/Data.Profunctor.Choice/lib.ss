#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Choice lib)
  (export
    choiceFn
    fanin
    identity
    left
    right
    splitChoice)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Profunctor lib) Data.Profunctor.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define right
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "right"))))

  (scm:define left
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "left"))))

  (scm:define splitChoice
    (scm:lambda (dictCategory0)
      (scm:let ([_1 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))])
        (scm:lambda (dictChoice2)
          (scm:lambda (l3)
            (scm:lambda (r4)
              (((rt:record-ref _1 (scm:string->symbol "compose")) ((rt:record-ref dictChoice2 (scm:string->symbol "right")) r4)) ((rt:record-ref dictChoice2 (scm:string->symbol "left")) l3))))))))

  (scm:define fanin
    (scm:lambda (dictCategory0)
      (scm:let*
        ([identity11 (rt:record-ref dictCategory0 (scm:string->symbol "identity"))]
         [_2 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))]
         [_3 ((rt:record-ref dictCategory0 (scm:string->symbol "Semigroupoid0")) (scm:quote undefined))])
          (scm:lambda (dictChoice4)
            (scm:lambda (l5)
              (scm:lambda (r6)
                (((rt:record-ref _2 (scm:string->symbol "compose")) ((((rt:record-ref ((rt:record-ref dictChoice4 (scm:string->symbol "Profunctor0")) (scm:quote undefined)) (scm:string->symbol "dimap")) (scm:lambda (v27)
                  (scm:cond
                    [(Data.Either.Left? v27) (Data.Either.Left-value0 v27)]
                    [(Data.Either.Right? v27) (Data.Either.Right-value0 v27)]
                    [scm:else (rt:fail)]))) identity) identity11)) (((rt:record-ref _3 (scm:string->symbol "compose")) ((rt:record-ref dictChoice4 (scm:string->symbol "right")) r6)) ((rt:record-ref dictChoice4 (scm:string->symbol "left")) l5)))))))))

  (scm:define choiceFn
    (scm:list (scm:cons (scm:string->symbol "left") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(Data.Either.Left? v11) (Data.Either.Left (v0 (Data.Either.Left-value0 v11)))]
          [(Data.Either.Right? v11) (Data.Either.Right (Data.Either.Right-value0 v11))]
          [scm:else (rt:fail)])))) (scm:cons (scm:string->symbol "right") (rt:record-ref Data.Either.functorEither (scm:string->symbol "map"))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
      Data.Profunctor.profunctorFn)))))
