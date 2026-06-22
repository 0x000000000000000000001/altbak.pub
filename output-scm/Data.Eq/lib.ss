#!r6rs
#!chezscheme
(library
  (Data.Eq lib)
  (export
    eq
    eq1
    eq1Array
    eqArray
    eqArrayImpl
    eqBoolean
    eqBooleanImpl
    eqChar
    eqCharImpl
    eqInt
    eqIntImpl
    eqNumber
    eqNumberImpl
    eqProxy
    eqRec
    eqRecord
    eqRowCons
    eqRowNil
    eqString
    eqStringImpl
    eqUnit
    eqVoid
    notEq
    notEq1)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Eq foreign))

  (scm:define eqVoid
    (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (v0)
      (scm:lambda (v11)
        #t)))))

  (scm:define eqUnit
    (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (v0)
      (scm:lambda (v11)
        #t)))))

  (scm:define eqString
    (scm:list (scm:cons (scm:string->symbol "eq") eqStringImpl)))

  (scm:define eqRowNil
    (scm:list (scm:cons (scm:string->symbol "eqRecord") (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          #t))))))

  (scm:define eqRecord
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "eqRecord"))))

  (scm:define eqRec
    (scm:lambda (_)
      (scm:lambda (dictEqRecord1)
        (scm:list (scm:cons (scm:string->symbol "eq") ((rt:record-ref dictEqRecord1 (scm:string->symbol "eqRecord")) Type.Proxy.Proxy))))))

  (scm:define eqProxy
    (scm:list (scm:cons (scm:string->symbol "eq") (scm:lambda (v0)
      (scm:lambda (v11)
        #t)))))

  (scm:define eqNumber
    (scm:list (scm:cons (scm:string->symbol "eq") eqNumberImpl)))

  (scm:define eqInt
    (scm:list (scm:cons (scm:string->symbol "eq") eqIntImpl)))

  (scm:define eqChar
    (scm:list (scm:cons (scm:string->symbol "eq") eqCharImpl)))

  (scm:define eqBoolean
    (scm:list (scm:cons (scm:string->symbol "eq") eqBooleanImpl)))

  (scm:define eq1
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "eq1"))))

  (scm:define eq
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "eq"))))

  (scm:define eqArray
    (scm:lambda (dictEq0)
      (scm:list (scm:cons (scm:string->symbol "eq") (eqArrayImpl (rt:record-ref dictEq0 (scm:string->symbol "eq")))))))

  (scm:define eq1Array
    (scm:list (scm:cons (scm:string->symbol "eq1") (scm:lambda (dictEq0)
      (eqArrayImpl (rt:record-ref dictEq0 (scm:string->symbol "eq")))))))

  (scm:define eqRowCons
    (scm:lambda (dictEqRecord0)
      (scm:lambda (_)
        (scm:lambda (dictIsSymbol2)
          (scm:lambda (dictEq3)
            (scm:list (scm:cons (scm:string->symbol "eqRecord") (scm:lambda (v4)
              (scm:lambda (ra5)
                (scm:lambda (rb6)
                  (scm:let ([get7 (Record.Unsafe.unsafeGet ((rt:record-ref dictIsSymbol2 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy))])
                    (scm:and (((rt:record-ref dictEq3 (scm:string->symbol "eq")) (get7 ra5)) (get7 rb6)) ((((rt:record-ref dictEqRecord0 (scm:string->symbol "eqRecord")) Type.Proxy.Proxy) ra5) rb6)))))))))))))

  (scm:define notEq
    (scm:lambda (dictEq0)
      (scm:lambda (x1)
        (scm:lambda (y2)
          (scm:not (((rt:record-ref dictEq0 (scm:string->symbol "eq")) x1) y2))))))

  (scm:define notEq1
    (scm:lambda (dictEq10)
      (scm:lambda (dictEq1)
        (scm:let ([eq122 ((rt:record-ref dictEq10 (scm:string->symbol "eq1")) dictEq1)])
          (scm:lambda (x3)
            (scm:lambda (y4)
              (scm:not ((eq122 x3) y4)))))))))
