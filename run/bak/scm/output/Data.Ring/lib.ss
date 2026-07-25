#!r6rs
#!chezscheme
(library
  (Data.Ring lib)
  (export
    intSub
    negate
    numSub
    ringFn
    ringInt
    ringNumber
    ringProxy
    ringRecord
    ringRecordCons
    ringRecordNil
    ringUnit
    sub
    subRecord)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Semiring lib) Data.Semiring.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Ring foreign))

  (scm:define subRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "subRecord"))))

  (scm:define sub
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "sub"))))

  (scm:define ringUnit
    (scm:list (scm:cons (scm:string->symbol "sub") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Unit.unit))) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
      Data.Semiring.semiringUnit))))

  (scm:define ringRecordNil
    (scm:list (scm:cons (scm:string->symbol "subRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:list))))) (scm:cons (scm:string->symbol "SemiringRecord0") (scm:lambda (_)
      Data.Semiring.semiringRecordNil))))

  (scm:define ringRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (_)
        (scm:lambda (dictRingRecord2)
          (scm:let ([semiringRecordCons13 (((Data.Semiring.semiringRecordCons dictIsSymbol0) (scm:quote undefined)) ((rt:record-ref dictRingRecord2 (scm:string->symbol "SemiringRecord0")) (scm:quote undefined)))])
            (scm:lambda (dictRing4)
              (scm:let ([semiringRecordCons25 (semiringRecordCons13 ((rt:record-ref dictRing4 (scm:string->symbol "Semiring0")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "subRecord") (scm:lambda (v6)
                  (scm:lambda (ra7)
                    (scm:lambda (rb8)
                      (scm:let*
                        ([key9 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                         [get10 (Record.Unsafe.unsafeGet key9)])
                          (((Record.Unsafe.unsafeSet key9) (((rt:record-ref dictRing4 (scm:string->symbol "sub")) (get10 ra7)) (get10 rb8))) ((((rt:record-ref dictRingRecord2 (scm:string->symbol "subRecord")) Type.Proxy.Proxy) ra7) rb8))))))) (scm:cons (scm:string->symbol "SemiringRecord0") (scm:lambda (_)
                  semiringRecordCons25))))))))))

  (scm:define ringRecord
    (scm:lambda (_)
      (scm:lambda (dictRingRecord1)
        (scm:let*
          ([_2 ((rt:record-ref dictRingRecord1 (scm:string->symbol "SemiringRecord0")) (scm:quote undefined))]
           [semiringRecord13 (scm:list (scm:cons (scm:string->symbol "add") ((rt:record-ref _2 (scm:string->symbol "addRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "mul") ((rt:record-ref _2 (scm:string->symbol "mulRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "one") (((rt:record-ref _2 (scm:string->symbol "oneRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "zero") (((rt:record-ref _2 (scm:string->symbol "zeroRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)))])
            (scm:list (scm:cons (scm:string->symbol "sub") ((rt:record-ref dictRingRecord1 (scm:string->symbol "subRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
              semiringRecord13)))))))

  (scm:define ringProxy
    (scm:list (scm:cons (scm:string->symbol "sub") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
      Data.Semiring.semiringProxy))))

  (scm:define ringNumber
    (scm:list (scm:cons (scm:string->symbol "sub") numSub) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
      Data.Semiring.semiringNumber))))

  (scm:define ringInt
    (scm:list (scm:cons (scm:string->symbol "sub") intSub) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
      Data.Semiring.semiringInt))))

  (scm:define ringFn
    (scm:lambda (dictRing0)
      (scm:let*
        ([_1 ((rt:record-ref dictRing0 (scm:string->symbol "Semiring0")) (scm:quote undefined))]
         [zero12 (rt:record-ref _1 (scm:string->symbol "zero"))]
         [semiringFn3 (scm:let ([one13 (rt:record-ref _1 (scm:string->symbol "one"))])
          (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (f4)
            (scm:lambda (g5)
              (scm:lambda (x6)
                (((rt:record-ref _1 (scm:string->symbol "add")) (f4 x6)) (g5 x6)))))) (scm:cons (scm:string->symbol "zero") (scm:lambda (v4)
            zero12)) (scm:cons (scm:string->symbol "mul") (scm:lambda (f4)
            (scm:lambda (g5)
              (scm:lambda (x6)
                (((rt:record-ref _1 (scm:string->symbol "mul")) (f4 x6)) (g5 x6)))))) (scm:cons (scm:string->symbol "one") (scm:lambda (v4)
            one13))))])
          (scm:list (scm:cons (scm:string->symbol "sub") (scm:lambda (f4)
            (scm:lambda (g5)
              (scm:lambda (x6)
                (((rt:record-ref dictRing0 (scm:string->symbol "sub")) (f4 x6)) (g5 x6)))))) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
            semiringFn3))))))

  (scm:define negate
    (scm:lambda (dictRing0)
      (scm:let ([zero1 (rt:record-ref ((rt:record-ref dictRing0 (scm:string->symbol "Semiring0")) (scm:quote undefined)) (scm:string->symbol "zero"))])
        (scm:lambda (a2)
          (((rt:record-ref dictRing0 (scm:string->symbol "sub")) zero1) a2))))))
