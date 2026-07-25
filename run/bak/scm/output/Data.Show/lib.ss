#!r6rs
#!chezscheme
(library
  (Data.Show lib)
  (export
    show
    showArray
    showArrayImpl
    showBoolean
    showChar
    showCharImpl
    showInt
    showIntImpl
    showNumber
    showNumberImpl
    showProxy
    showRecord
    showRecordFields
    showRecordFieldsCons
    showRecordFieldsConsNil
    showRecordFieldsNil
    showString
    showStringImpl
    showUnit
    showVoid)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Void lib) Data.Void.)
    (prefix (Record.Unsafe lib) Record.Unsafe.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Show foreign))

  (scm:define showVoid
    (scm:list (scm:cons (scm:string->symbol "show") Data.Void.absurd)))

  (scm:define showUnit
    (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v0)
      (rt:string->pstring "unit")))))

  (scm:define showString
    (scm:list (scm:cons (scm:string->symbol "show") showStringImpl)))

  (scm:define showRecordFieldsNil
    (scm:list (scm:cons (scm:string->symbol "showRecordFields") (scm:lambda (v0)
      (scm:lambda (v11)
        (rt:string->pstring ""))))))

  (scm:define showRecordFields
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "showRecordFields"))))

  (scm:define showRecord
    (scm:lambda (_)
      (scm:lambda (_)
        (scm:lambda (dictShowRecordFields2)
          (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (record3)
            (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "{") (((rt:record-ref dictShowRecordFields2 (scm:string->symbol "showRecordFields")) Type.Proxy.Proxy) record3)) (rt:string->pstring "}")))))))))

  (scm:define showProxy
    (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v0)
      (rt:string->pstring "Proxy")))))

  (scm:define showNumber
    (scm:list (scm:cons (scm:string->symbol "show") showNumberImpl)))

  (scm:define showInt
    (scm:list (scm:cons (scm:string->symbol "show") showIntImpl)))

  (scm:define showChar
    (scm:list (scm:cons (scm:string->symbol "show") showCharImpl)))

  (scm:define showBoolean
    (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v0)
      (scm:cond
        [v0 (rt:string->pstring "true")]
        [scm:else (rt:string->pstring "false")])))))

  (scm:define show
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "show"))))

  (scm:define showArray
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (showArrayImpl (rt:record-ref dictShow0 (scm:string->symbol "show")))))))

  (scm:define showRecordFieldsCons
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (dictShowRecordFields1)
        (scm:lambda (dictShow2)
          (scm:list (scm:cons (scm:string->symbol "showRecordFields") (scm:lambda (v3)
            (scm:lambda (record4)
              (scm:let ([key5 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)])
                (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring " ") key5) (rt:string->pstring ": ")) ((rt:record-ref dictShow2 (scm:string->symbol "show")) ((Record.Unsafe.unsafeGet key5) record4))) (rt:string->pstring ",")) (((rt:record-ref dictShowRecordFields1 (scm:string->symbol "showRecordFields")) Type.Proxy.Proxy) record4)))))))))))

  (scm:define showRecordFieldsConsNil
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (dictShow1)
        (scm:list (scm:cons (scm:string->symbol "showRecordFields") (scm:lambda (v2)
          (scm:lambda (record3)
            (scm:let ([key4 ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)])
              (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring " ") key4) (rt:string->pstring ": ")) ((rt:record-ref dictShow1 (scm:string->symbol "show")) ((Record.Unsafe.unsafeGet key4) record3))) (rt:string->pstring " ")))))))))))
