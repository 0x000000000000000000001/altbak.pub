#!r6rs
#!chezscheme
(library
  (Control.Monad.ST.Internal lib)
  (export
    applicativeST
    applyST
    bindST
    bind_
    for
    foreach
    functorST
    map_
    modify
    modify$p
    modifyImpl
    monadRecST
    monadST
    monoidST
    new
    pure_
    read
    run
    semigroupST
    while
    write)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Monad.Rec.Class lib) Control.Monad.Rec.Class.)
    (prefix (Data.Unit lib) Data.Unit.)
    (Control.Monad.ST.Internal foreign))

  (scm:define modify$p
    modifyImpl)

  (scm:define modify
    (scm:lambda (f0)
      (modifyImpl (scm:lambda (s1)
        (scm:let ([s$p2 (f0 s1)])
          (scm:list (scm:cons (scm:string->symbol "state") s$p2) (scm:cons (scm:string->symbol "value") s$p2)))))))

  (scm:define functorST
    (scm:list (scm:cons (scm:string->symbol "map") map_)))

  (rt:define-lazy $lazy-monadST "monadST" "Control.Monad.ST.Internal"
    (scm:list (scm:cons (scm:string->symbol "Applicative0") (scm:lambda (_)
      ($lazy-applicativeST))) (scm:cons (scm:string->symbol "Bind1") (scm:lambda (_)
      ($lazy-bindST)))))

  (rt:define-lazy $lazy-bindST "bindST" "Control.Monad.ST.Internal"
    (scm:list (scm:cons (scm:string->symbol "bind") bind_) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      ($lazy-applyST)))))

  (rt:define-lazy $lazy-applyST "applyST" "Control.Monad.ST.Internal"
    (scm:list (scm:cons (scm:string->symbol "apply") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:lambda ()
          (scm:let*
            ([f$p2 (f0)]
             [a$p3 (a1)])
              (((rt:record-ref ($lazy-applicativeST) (scm:string->symbol "pure")) (f$p2 a$p3)))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      functorST))))

  (rt:define-lazy $lazy-applicativeST "applicativeST" "Control.Monad.ST.Internal"
    (scm:list (scm:cons (scm:string->symbol "pure") pure_) (scm:cons (scm:string->symbol "Apply0") (scm:lambda (_)
      ($lazy-applyST)))))

  (scm:define monadST
    ($lazy-monadST))

  (scm:define bindST
    ($lazy-bindST))

  (scm:define applyST
    ($lazy-applyST))

  (scm:define applicativeST
    ($lazy-applicativeST))

  (scm:define semigroupST
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (a1)
        (scm:lambda (b2)
          (scm:lambda ()
            (scm:let*
              ([_3 (a1)]
               [a$p4 (b2)])
                (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) _3) a$p4)))))))))

  (scm:define monadRecST
    (scm:list (scm:cons (scm:string->symbol "tailRecM") (scm:lambda (f0)
      (scm:lambda (a1)
        (scm:let ([_2 (f0 a1)])
          (scm:lambda ()
            (scm:let*
              ([_3 (_2)]
               [r4 (scm:box _3)]
               [_ (((while (scm:lambda ()
                (scm:let ([_5 (scm:unbox r4)])
                  (Control.Monad.Rec.Class.Loop? _5)))) (scm:lambda ()
                (scm:let ([v5 (scm:unbox r4)])
                  ((scm:cond
                    [(Control.Monad.Rec.Class.Loop? v5) (scm:lambda ()
                      (scm:let*
                        ([e6 ((f0 (Control.Monad.Rec.Class.Loop-value0 v5)))]
                         [_7 (scm:begin (scm:set-box! r4 e6) (scm:unbox r4))])
                          Data.Unit.unit))]
                    [(Control.Monad.Rec.Class.Done? v5) (scm:lambda ()
                      Data.Unit.unit)]
                    [scm:else (rt:fail)]))))))]
               [_6 (scm:unbox r4)])
                (scm:cond
                  [(Control.Monad.Rec.Class.Done? _6) (Control.Monad.Rec.Class.Done-value0 _6)]
                  [scm:else (rt:fail)]))))))) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      monadST))))

  (scm:define monoidST
    (scm:lambda (dictMonoid0)
      (scm:let*
        ([_1 ((rt:record-ref dictMonoid0 (scm:string->symbol "Semigroup0")) (scm:quote undefined))]
         [semigroupST12 (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda ()
              (scm:let*
                ([_4 (a2)]
                 [a$p5 (b3)])
                  (((rt:record-ref _1 (scm:string->symbol "append")) _4) a$p5)))))))])
          (scm:list (scm:cons (scm:string->symbol "mempty") (scm:let ([_3 (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))])
            (scm:lambda ()
              _3))) (scm:cons (scm:string->symbol "Semigroup0") (scm:lambda (_)
            semigroupST12)))))))
