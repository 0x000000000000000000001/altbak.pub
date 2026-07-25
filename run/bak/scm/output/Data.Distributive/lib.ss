#!r6rs
#!chezscheme
(library
  (Data.Distributive lib)
  (export
    collect
    collectDefault
    cotraverse
    distribute
    distributeDefault
    distributiveFunction
    distributiveIdentity
    distributiveTuple
    identity)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.Identity lib) Data.Identity.)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define identity
    (scm:lambda (x0)
      x0))

  (scm:define distributiveIdentity
    (scm:list (scm:cons (scm:string->symbol "distribute") (scm:lambda (dictFunctor0)
      ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) Unsafe.Coerce.unsafeCoerce))) (scm:cons (scm:string->symbol "collect") (scm:lambda (dictFunctor0)
      (scm:lambda (f1)
        ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (x2)
          (f1 x2)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Identity.functorIdentity))))

  (scm:define distribute
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "distribute"))))

  (rt:define-lazy $lazy-distributiveFunction "distributiveFunction" "Data.Distributive"
    (scm:list (scm:cons (scm:string->symbol "distribute") (scm:lambda (dictFunctor0)
      (scm:lambda (a1)
        (scm:lambda (e2)
          (((rt:record-ref dictFunctor0 (scm:string->symbol "map")) (scm:lambda (v3)
            (v3 e2))) a1))))) (scm:cons (scm:string->symbol "collect") (scm:lambda (dictFunctor0)
      (scm:lambda (f1)
        (scm:let*
          ([_2 ((rt:record-ref ($lazy-distributiveFunction) (scm:string->symbol "distribute")) dictFunctor0)]
           [_3 ((rt:record-ref dictFunctor0 (scm:string->symbol "map")) f1)])
            (scm:lambda (x4)
              (_2 (_3 x4))))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorFn))))

  (scm:define distributiveFunction
    ($lazy-distributiveFunction))

  (scm:define cotraverse
    (scm:lambda (dictDistributive0)
      (scm:lambda (dictFunctor1)
        (scm:let ([distribute22 ((rt:record-ref dictDistributive0 (scm:string->symbol "distribute")) dictFunctor1)])
          (scm:lambda (f3)
            (scm:let ([_4 ((rt:record-ref ((rt:record-ref dictDistributive0 (scm:string->symbol "Functor0")) (scm:quote undefined)) (scm:string->symbol "map")) f3)])
              (scm:lambda (x5)
                (_4 (distribute22 x5)))))))))

  (scm:define collectDefault
    (scm:lambda (dictDistributive0)
      (scm:lambda (dictFunctor1)
        (scm:let ([distribute22 ((rt:record-ref dictDistributive0 (scm:string->symbol "distribute")) dictFunctor1)])
          (scm:lambda (f3)
            (scm:let ([_4 ((rt:record-ref dictFunctor1 (scm:string->symbol "map")) f3)])
              (scm:lambda (x5)
                (distribute22 (_4 x5)))))))))

  (scm:define distributiveTuple
    (scm:lambda (dictTypeEquals0)
      (scm:let ([from1 ((rt:record-ref dictTypeEquals0 (scm:string->symbol "proof")) (scm:lambda (a1)
        a1))])
        (scm:list (scm:cons (scm:string->symbol "collect") (scm:lambda (dictFunctor2)
          (scm:let ([distribute23 ((rt:record-ref (distributiveTuple dictTypeEquals0) (scm:string->symbol "distribute")) dictFunctor2)])
            (scm:lambda (f4)
              (scm:let ([_5 ((rt:record-ref dictFunctor2 (scm:string->symbol "map")) f4)])
                (scm:lambda (x6)
                  (distribute23 (_5 x6)))))))) (scm:cons (scm:string->symbol "distribute") (scm:lambda (dictFunctor2)
          (scm:let*
            ([_3 (Data.Tuple.Tuple (from1 Data.Unit.unit))]
             [_4 ((rt:record-ref dictFunctor2 (scm:string->symbol "map")) Data.Tuple.snd)])
              (scm:lambda (x5)
                (_3 (_4 x5)))))) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
          Data.Tuple.functorTuple))))))

  (scm:define collect
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "collect"))))

  (scm:define distributeDefault
    (scm:lambda (dictDistributive0)
      (scm:lambda (dictFunctor1)
        (((rt:record-ref dictDistributive0 (scm:string->symbol "collect")) dictFunctor1) identity)))))
