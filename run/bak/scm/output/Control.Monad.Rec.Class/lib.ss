#!r6rs
#!chezscheme
(library
  (Control.Monad.Rec.Class lib)
  (export
    Done
    Done-value0
    Done?
    Loop
    Loop-value0
    Loop?
    bifunctorStep
    forever
    functorStep
    loop2
    loop3
    monadRecEffect
    monadRecEither
    monadRecFunction
    monadRecIdentity
    monadRecMaybe
    tailRec
    tailRec2
    tailRec3
    tailRecM
    tailRecM2
    tailRecM3
    untilJust
    whileJust)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Monad lib) Control.Monad.)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Identity lib) Data.Identity.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Effect lib) Effect.)
    (prefix (Effect.Ref lib) Effect.Ref.))

  (scm:define-record-type (Loop$ Loop Loop?)
    (scm:fields (scm:immutable value0 Loop-value0)))

  (scm:define-record-type (Done$ Done Done?)
    (scm:fields (scm:immutable value0 Done-value0)))

  (scm:define tailRecM
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "tailRecM"))))

  (scm:define tailRecM2
    (scm:lambda (dictMonadRec0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (((rt:record-ref dictMonadRec0 (scm:string->symbol "tailRecM")) (scm:lambda (o4)
              ((f1 (rt:record-ref o4 (scm:string->symbol "a"))) (rt:record-ref o4 (scm:string->symbol "b"))))) (scm:list (scm:cons (scm:string->symbol "a") a2) (scm:cons (scm:string->symbol "b") b3))))))))

  (scm:define tailRecM3
    (scm:lambda (dictMonadRec0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (((rt:record-ref dictMonadRec0 (scm:string->symbol "tailRecM")) (scm:lambda (o5)
                (((f1 (rt:record-ref o5 (scm:string->symbol "a"))) (rt:record-ref o5 (scm:string->symbol "b"))) (rt:record-ref o5 (scm:string->symbol "c"))))) (scm:list (scm:cons (scm:string->symbol "a") a2) (scm:cons (scm:string->symbol "b") b3) (scm:cons (scm:string->symbol "c") c4)))))))))

  (scm:define untilJust
    (scm:lambda (dictMonadRec0)
      (scm:let ([_1 ((rt:record-ref ((rt:record-ref ((rt:record-ref ((rt:record-ref dictMonadRec0 (scm:string->symbol "Monad0")) (scm:quote undefined)) (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (m2)
          (((rt:record-ref dictMonadRec0 (scm:string->symbol "tailRecM")) (scm:lambda (v3)
            (((rt:record-ref _1 (scm:string->symbol "map")) (scm:lambda (v14)
              (scm:cond
                [(Data.Maybe.Nothing? v14) (Loop Data.Unit.unit)]
                [(Data.Maybe.Just? v14) (Done (Data.Maybe.Just-value0 v14))]
                [scm:else (rt:fail)]))) m2))) Data.Unit.unit)))))

  (scm:define whileJust
    (scm:lambda (dictMonoid0)
      (scm:let ([mempty1 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
        (scm:lambda (dictMonadRec2)
          (scm:let ([_3 ((rt:record-ref ((rt:record-ref ((rt:record-ref ((rt:record-ref dictMonadRec2 (scm:string->symbol "Monad0")) (scm:quote undefined)) (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))])
            (scm:lambda (m4)
              (((rt:record-ref dictMonadRec2 (scm:string->symbol "tailRecM")) (scm:lambda (v5)
                (((rt:record-ref _3 (scm:string->symbol "map")) (scm:lambda (v16)
                  (scm:cond
                    [(Data.Maybe.Nothing? v16) (Done v5)]
                    [(Data.Maybe.Just? v16) (Loop (((rt:record-ref ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined)) (scm:string->symbol "append")) v5) (Data.Maybe.Just-value0 v16)))]
                    [scm:else (rt:fail)]))) m4))) mempty1)))))))

  (scm:define tailRec
    (scm:lambda (f0)
      (scm:letrec ([go1 (scm:lambda (v2)
        (scm:cond
          [(Loop? v2) (go1 (f0 (Loop-value0 v2)))]
          [(Done? v2) (Done-value0 v2)]
          [scm:else (rt:fail)]))])
        (scm:lambda (x2)
          (go1 (f0 x2))))))

  (scm:define tailRec2
    (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (scm:letrec ([go3 (scm:lambda (v4)
            (scm:cond
              [(Loop? v4) (go3 ((f0 (rt:record-ref (Loop-value0 v4) (scm:string->symbol "a"))) (rt:record-ref (Loop-value0 v4) (scm:string->symbol "b"))))]
              [(Done? v4) (Done-value0 v4)]
              [scm:else (rt:fail)]))])
            (go3 ((f0 a1) b2)))))))

  (scm:define tailRec3
    (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:lambda (b2)
          (scm:lambda (c3)
            (scm:letrec ([go4 (scm:lambda (v5)
              (scm:cond
                [(Loop? v5) (go4 (((f0 (rt:record-ref (Loop-value0 v5) (scm:string->symbol "a"))) (rt:record-ref (Loop-value0 v5) (scm:string->symbol "b"))) (rt:record-ref (Loop-value0 v5) (scm:string->symbol "c"))))]
                [(Done? v5) (Done-value0 v5)]
                [scm:else (rt:fail)]))])
              (go4 (((f0 a1) b2) c3))))))))

  (scm:define monadRecMaybe
    (scm:list (scm:cons (scm:string->symbol "tailRecM") (scm:lambda (f0)
      (scm:lambda (a01)
        (scm:let ([_2 (scm:lambda (v2)
          (scm:cond
            [(Data.Maybe.Nothing? v2) (Done Data.Maybe.Nothing)]
            [(Data.Maybe.Just? v2) (scm:cond
              [(Loop? (Data.Maybe.Just-value0 v2)) (Loop (f0 (Loop-value0 (Data.Maybe.Just-value0 v2))))]
              [(Done? (Data.Maybe.Just-value0 v2)) (Done (Data.Maybe.Just (Done-value0 (Data.Maybe.Just-value0 v2))))]
              [scm:else (rt:fail)])]
            [scm:else (rt:fail)]))])
          (scm:letrec ([go3 (scm:lambda (v4)
            (scm:cond
              [(Loop? v4) (go3 (_2 (Loop-value0 v4)))]
              [(Done? v4) (Done-value0 v4)]
              [scm:else (rt:fail)]))])
            (go3 (_2 (f0 a01)))))))) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Data.Maybe.monadMaybe))))

  (scm:define monadRecIdentity
    (scm:list (scm:cons (scm:string->symbol "tailRecM") (scm:lambda (f0)
      (scm:letrec ([go1 (scm:lambda (v2)
        (scm:cond
          [(Loop? v2) (go1 (f0 (Loop-value0 v2)))]
          [(Done? v2) (Done-value0 v2)]
          [scm:else (rt:fail)]))])
        (scm:lambda (x2)
          (go1 (f0 x2)))))) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Data.Identity.monadIdentity))))

  (scm:define monadRecFunction
    (scm:list (scm:cons (scm:string->symbol "tailRecM") (scm:lambda (f0)
      (scm:lambda (a01)
        (scm:lambda (e2)
          (scm:letrec ([go3 (scm:lambda (v4)
            (scm:cond
              [(Loop? v4) (go3 ((f0 (Loop-value0 v4)) e2))]
              [(Done? v4) (Done-value0 v4)]
              [scm:else (rt:fail)]))])
            (go3 ((f0 a01) e2))))))) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Control.Monad.monadFn))))

  (scm:define monadRecEither
    (scm:list (scm:cons (scm:string->symbol "tailRecM") (scm:lambda (f0)
      (scm:lambda (a01)
        (scm:let ([_2 (scm:lambda (v2)
          (scm:cond
            [(Data.Either.Left? v2) (Done (Data.Either.Left (Data.Either.Left-value0 v2)))]
            [(Data.Either.Right? v2) (scm:cond
              [(Loop? (Data.Either.Right-value0 v2)) (Loop (f0 (Loop-value0 (Data.Either.Right-value0 v2))))]
              [(Done? (Data.Either.Right-value0 v2)) (Done (Data.Either.Right (Done-value0 (Data.Either.Right-value0 v2))))]
              [scm:else (rt:fail)])]
            [scm:else (rt:fail)]))])
          (scm:letrec ([go3 (scm:lambda (v4)
            (scm:cond
              [(Loop? v4) (go3 (_2 (Loop-value0 v4)))]
              [(Done? v4) (Done-value0 v4)]
              [scm:else (rt:fail)]))])
            (go3 (_2 (f0 a01)))))))) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Data.Either.monadEither))))

  (scm:define monadRecEffect
    (scm:list (scm:cons (scm:string->symbol "tailRecM") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:let ([_2 (f0 a1)])
          (scm:lambda ()
            (scm:let*
              ([_3 (_2)]
               [r4 ((Effect.Ref._new _3))]
               [_ ((Effect.untilE (scm:let ([_5 (Effect.Ref.read r4)])
                (scm:lambda ()
                  (scm:let ([v6 (_5)])
                    ((scm:cond
                      [(Loop? v6) (scm:lambda ()
                        (scm:let*
                          ([e7 ((f0 (Loop-value0 v6)))]
                           [_ (((Effect.Ref.write e7) r4))])
                            #f))]
                      [(Done? v6) (scm:lambda ()
                        #t)]
                      [scm:else (rt:fail)])))))))]
               [a$p6 ((Effect.Ref.read r4))])
                (scm:cond
                  [(Done? a$p6) (Done-value0 a$p6)]
                  [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Effect.monadEffect))))

  (scm:define loop3
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (Loop (scm:list (scm:cons (scm:string->symbol "a") a0) (scm:cons (scm:string->symbol "b") b1) (scm:cons (scm:string->symbol "c") c2)))))))

  (scm:define loop2
    (scm:lambda (a0)
      (scm:lambda (b1)
        (Loop (scm:list (scm:cons (scm:string->symbol "a") a0) (scm:cons (scm:string->symbol "b") b1))))))

  (scm:define functorStep
    (scm:list (scm:cons (scm:string->symbol "map") (scm:lambda (f0)
      (scm:lambda (m1)
        (scm:cond
          [(Loop? m1) (Loop (Loop-value0 m1))]
          [(Done? m1) (Done (f0 (Done-value0 m1)))]
          [scm:else (rt:fail)]))))))

  (scm:define forever
    (scm:lambda (dictMonadRec0)
      (scm:let ([_1 ((rt:record-ref ((rt:record-ref ((rt:record-ref ((rt:record-ref dictMonadRec0 (scm:string->symbol "Monad0")) (scm:quote undefined)) (scm:string->symbol "Bind1")) (scm:quote undefined)) (scm:string->symbol "Apply0")) (scm:quote undefined)) (scm:string->symbol "Functor0")) (scm:quote undefined))])
        (scm:lambda (ma2)
          (((rt:record-ref dictMonadRec0 (scm:string->symbol "tailRecM")) (scm:lambda (u3)
            (((rt:record-ref _1 (scm:string->symbol "map")) (scm:lambda (v4)
              (Loop u3))) ma2))) Data.Unit.unit)))))

  (scm:define bifunctorStep
    (scm:list (scm:cons (scm:string->symbol "bimap") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Loop? v22) (Loop (v0 (Loop-value0 v22)))]
            [(Done? v22) (Done (v11 (Done-value0 v22)))]
            [scm:else (rt:fail)]))))))))
