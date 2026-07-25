#!r6rs
#!chezscheme
(library
  (Data.Semiring lib)
  (export
    add
    addRecord
    intAdd
    intMul
    mul
    mulRecord
    numAdd
    numMul
    one
    oneRecord
    semiringFn
    semiringInt
    semiringNumber
    semiringProxy
    semiringRecord
    semiringRecordCons
    semiringRecordNil
    semiringUnit
    zero
    zeroRecord)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Semiring foreign))

  (scm:define zeroRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "zeroRecord"))))

  (scm:define zero
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "zero"))))

  (scm:define semiringUnit
    (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Unit.unit))) (scm:cons (scm:string->symbol "zero") Data.Unit.unit) (scm:cons (scm:string->symbol "mul") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Unit.unit))) (scm:cons (scm:string->symbol "one") Data.Unit.unit)))

  (scm:define semiringRecordNil
    (scm:list (scm:cons (scm:string->symbol "addRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:list))))) (scm:cons (scm:string->symbol "mulRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:list))))) (scm:cons (scm:string->symbol "oneRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:list)))) (scm:cons (scm:string->symbol "zeroRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:list))))))

  (scm:define semiringProxy
    (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "mul") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy))) (scm:cons (scm:string->symbol "one") Type.Proxy.Proxy) (scm:cons (scm:string->symbol "zero") Type.Proxy.Proxy)))

  (scm:define semiringNumber
    (scm:list (scm:cons (scm:string->symbol "add") numAdd) (scm:cons (scm:string->symbol "zero") 0.0) (scm:cons (scm:string->symbol "mul") numMul) (scm:cons (scm:string->symbol "one") 1.0)))

  (scm:define semiringInt
    (scm:list (scm:cons (scm:string->symbol "add") intAdd) (scm:cons (scm:string->symbol "zero") 0) (scm:cons (scm:string->symbol "mul") intMul) (scm:cons (scm:string->symbol "one") 1)))

  (scm:define oneRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "oneRecord"))))

  (scm:define one
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "one"))))

  (scm:define mulRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "mulRecord"))))

  (scm:define mul
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "mul"))))

  (scm:define addRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "addRecord"))))

  (scm:define semiringRecord
    (scm:lambda (_)
      (scm:lambda (dictSemiringRecord1)
        (scm:list (scm:cons (scm:string->symbol "add") ((rt:record-ref dictSemiringRecord1 (scm:string->symbol "addRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "mul") ((rt:record-ref dictSemiringRecord1 (scm:string->symbol "mulRecord")) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "one") (((rt:record-ref dictSemiringRecord1 (scm:string->symbol "oneRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)) (scm:cons (scm:string->symbol "zero") (((rt:record-ref dictSemiringRecord1 (scm:string->symbol "zeroRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy))))))

  (scm:define add
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "add"))))

  (scm:define semiringFn
    (scm:lambda (dictSemiring0)
      (scm:let*
        ([zero11 (rt:record-ref dictSemiring0 (scm:string->symbol "zero"))]
         [one12 (rt:record-ref dictSemiring0 (scm:string->symbol "one"))])
          (scm:list (scm:cons (scm:string->symbol "add") (scm:lambda (f3)
            (scm:lambda (g4)
              (scm:lambda (x5)
                (((rt:record-ref dictSemiring0 (scm:string->symbol "add")) (f3 x5)) (g4 x5)))))) (scm:cons (scm:string->symbol "zero") (scm:lambda (v3)
            zero11)) (scm:cons (scm:string->symbol "mul") (scm:lambda (f3)
            (scm:lambda (g4)
              (scm:lambda (x5)
                (((rt:record-ref dictSemiring0 (scm:string->symbol "mul")) (f3 x5)) (g4 x5)))))) (scm:cons (scm:string->symbol "one") (scm:lambda (v3)
            one12))))))

  (scm:define semiringRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (_)
        (scm:lambda (dictSemiringRecord2)
          (scm:lambda (dictSemiring3)
            (scm:let*
              ([one14 (rt:record-ref dictSemiring3 (scm:string->symbol "one"))]
               [zero15 (rt:record-ref dictSemiring3 (scm:string->symbol "zero"))])
                (scm:list (scm:cons (scm:string->symbol "addRecord") (scm:lambda (v6)
                  (scm:lambda (ra7)
                    (scm:lambda (rb8)
                      (scm:let*
                        ([key9 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                         [get10 (Record.Unsafe.unsafeGet key9)])
                          (((Record.Unsafe.unsafeSet key9) (((rt:record-ref dictSemiring3 (scm:string->symbol "add")) (get10 ra7)) (get10 rb8))) ((((rt:record-ref dictSemiringRecord2 (scm:string->symbol "addRecord")) Type.Proxy.Proxy) ra7) rb8))))))) (scm:cons (scm:string->symbol "mulRecord") (scm:lambda (v6)
                  (scm:lambda (ra7)
                    (scm:lambda (rb8)
                      (scm:let*
                        ([key9 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                         [get10 (Record.Unsafe.unsafeGet key9)])
                          (((Record.Unsafe.unsafeSet key9) (((rt:record-ref dictSemiring3 (scm:string->symbol "mul")) (get10 ra7)) (get10 rb8))) ((((rt:record-ref dictSemiringRecord2 (scm:string->symbol "mulRecord")) Type.Proxy.Proxy) ra7) rb8))))))) (scm:cons (scm:string->symbol "oneRecord") (scm:lambda (v6)
                  (scm:lambda (v17)
                    (((Record.Unsafe.unsafeSet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)) one14) (((rt:record-ref dictSemiringRecord2 (scm:string->symbol "oneRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy))))) (scm:cons (scm:string->symbol "zeroRecord") (scm:lambda (v6)
                  (scm:lambda (v17)
                    (((Record.Unsafe.unsafeSet ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)) zero15) (((rt:record-ref dictSemiringRecord2 (scm:string->symbol "zeroRecord")) Type.Proxy.Proxy) Type.Proxy.Proxy)))))))))))))
