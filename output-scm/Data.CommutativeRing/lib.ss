#!r6rs
#!chezscheme
(library
  (Data.CommutativeRing lib)
  (export
    commutativeRingFn
    commutativeRingInt
    commutativeRingNumber
    commutativeRingProxy
    commutativeRingRecord
    commutativeRingRecordCons
    commutativeRingRecordNil
    commutativeRingUnit)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Ring lib) Data.Ring.)
    (prefix (Type.Proxy lib) Type.Proxy.))

  (scm:define commutativeRingUnit
    (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
      Data.Ring.ringUnit))))

  (scm:define commutativeRingRecordNil
    (scm:list (scm:cons (scm:string->symbol "RingRecord0") (scm:lambda (_)
      Data.Ring.ringRecordNil))))

  (scm:define commutativeRingRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (_)
        (scm:lambda (dictCommutativeRingRecord2)
          (scm:let ([ringRecordCons13 (((Data.Ring.ringRecordCons dictIsSymbol0) (scm:quote undefined)) ((rt:record-ref dictCommutativeRingRecord2 (scm:string->symbol "RingRecord0")) (scm:quote undefined)))])
            (scm:lambda (dictCommutativeRing4)
              (scm:let ([ringRecordCons25 (ringRecordCons13 ((rt:record-ref dictCommutativeRing4 (scm:string->symbol "Ring0")) (scm:quote undefined)))])
                (scm:list (scm:cons (scm:string->symbol "RingRecord0") (scm:lambda (_)
                  ringRecordCons25))))))))))

  (scm:define commutativeRingRecord
    (scm:lambda (_)
      (scm:lambda (dictCommutativeRingRecord1)
        (scm:let*
          ([_2 ((rt:record-ref dictCommutativeRingRecord1 (scm:string->symbol "RingRecord0")) (scm:quote undefined))]
           [_3 ((rt:record-ref _2 (scm:string->symbol "SemiringRecord0")) (scm:quote undefined))]
           [ringRecord14 (scm:let ([semiringRecord14 (scm:list (scm:cons (scm:string->symbol "add") ((rt:record-ref _3 (scm:string->symbol "addRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "mul") ((rt:record-ref _3 (scm:string->symbol "mulRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "one") (((rt:record-ref _3 (scm:string->symbol "oneRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "zero") (((rt:record-ref _3 (scm:string->symbol "zeroRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)))])
            (scm:list (scm:cons (scm:string->symbol "sub") ((rt:record-ref _2 (scm:string->symbol "subRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
              semiringRecord14))))])
            (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
              ringRecord14)))))))

  (scm:define commutativeRingProxy
    (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
      Data.Ring.ringProxy))))

  (scm:define commutativeRingNumber
    (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
      Data.Ring.ringNumber))))

  (scm:define commutativeRingInt
    (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
      Data.Ring.ringInt))))

  (scm:define commutativeRingFn
    (scm:lambda (dictCommutativeRing0)
      (scm:let*
        ([_1 ((rt:record-ref dictCommutativeRing0 (scm:string->symbol "Ring0")) (scm:quote undefined))]
         [_2 ((rt:record-ref _1 (scm:string->symbol "Semiring0")) (scm:quote undefined))]
         [ringFn3 (scm:let*
          ([zero13 (rt:record-ref _2 (scm:string->symbol "zero"))]
           [one14 (rt:record-ref _2 (scm:string->symbol "one"))]
           [semiringFn5 (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (f5)
            (scm:lambda (g6)
              (scm:lambda (x7)
                (((rt:record-ref _2 (scm:string->symbol "add")) (f5 x7)) (g6 x7)))))) (scm:cons (scm:string->symbol "zero") (scm:lambda (v5)
            zero13)) (scm:cons (scm:string->symbol "mul") (scm:lambda (f5)
            (scm:lambda (g6)
              (scm:lambda (x7)
                (((rt:record-ref _2 (scm:string->symbol "mul")) (f5 x7)) (g6 x7)))))) (scm:cons (scm:string->symbol "one") (scm:lambda (v5)
            one14)))])
            (scm:list (scm:cons (scm:string->symbol "sub") (scm:lambda (f6)
              (scm:lambda (g7)
                (scm:lambda (x8)
                  (((rt:record-ref _1 (scm:string->symbol "sub")) (f6 x8)) (g7 x8)))))) (scm:cons (scm:string->symbol "Semiring0") (scm:lambda (_)
              semiringFn5))))])
          (scm:list (scm:cons (scm:string->symbol "Ring0") (scm:lambda (_)
            ringFn3)))))))
