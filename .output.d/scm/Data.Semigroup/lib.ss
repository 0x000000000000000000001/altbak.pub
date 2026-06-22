#!r6rs
#!chezscheme
(library
  (Data.Semigroup lib)
  (export
    append
    appendRecord
    concatArray
    concatString
    semigroupArray
    semigroupFn
    semigroupProxy
    semigroupRecord
    semigroupRecordCons
    semigroupRecordNil
    semigroupString
    semigroupUnit
    semigroupVoid)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Data.Void lib) Data.Void.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Semigroup foreign))

  (scm:define semigroupVoid
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      Data.Void.absurd))))

  (scm:define semigroupUnit
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        Data.Unit.unit)))))

  (scm:define semigroupString
    (scm:list (scm:cons (scm:string->symbol "append") concatString)))

  (scm:define semigroupRecordNil
    (scm:list (scm:cons (scm:string->symbol "appendRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:list)))))))

  (scm:define semigroupProxy
    (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (v0)
      (scm:lambda (v11)
        Type.Proxy.Proxy)))))

  (scm:define semigroupArray
    (scm:list (scm:cons (scm:string->symbol "append") concatArray)))

  (scm:define appendRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "appendRecord"))))

  (scm:define semigroupRecord
    (scm:lambda (_)
      (scm:lambda (dictSemigroupRecord1)
        (scm:list (scm:cons (scm:string->symbol "append") ((rt:record-ref dictSemigroupRecord1 (scm:string->symbol "appendRecord")) Type.Proxy.Proxy))))))

  (scm:define append
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "append"))))

  (scm:define semigroupFn
    (scm:lambda (dictSemigroup0)
      (scm:list (scm:cons (scm:string->symbol "append") (scm:lambda (f1)
        (scm:lambda (g2)
          (scm:lambda (x3)
            (((rt:record-ref dictSemigroup0 (scm:string->symbol "append")) (f1 x3)) (g2 x3)))))))))

  (scm:define semigroupRecordCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (_)
        (scm:lambda (dictSemigroupRecord2)
          (scm:lambda (dictSemigroup3)
            (scm:list (scm:cons (scm:string->symbol "appendRecord") (scm:lambda (v4)
              (scm:lambda (ra5)
                (scm:lambda (rb6)
                  (scm:let*
                    ([key7 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
                     [get8 (Record.Unsafe.unsafeGet key7)])
                      (((Record.Unsafe.unsafeSet key7) (((rt:record-ref dictSemigroup3 (scm:string->symbol "append")) (get8 ra5)) (get8 rb6))) ((((rt:record-ref dictSemigroupRecord2 (scm:string->symbol "appendRecord")) Type.Proxy.Proxy) ra5) rb6))))))))))))))
