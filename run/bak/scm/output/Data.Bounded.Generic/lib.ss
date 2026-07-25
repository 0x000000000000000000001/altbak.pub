#!r6rs
#!chezscheme
(library
  (Data.Bounded.Generic lib)
  (export
    genericBottom
    genericBottom$p
    genericBottomArgument
    genericBottomConstructor
    genericBottomNoArguments
    genericBottomProduct
    genericBottomSum
    genericTop
    genericTop$p
    genericTopArgument
    genericTopConstructor
    genericTopNoArguments
    genericTopProduct
    genericTopSum)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.))

  (scm:define genericTopNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericTop'") Data.Generic.Rep.NoArguments)))

  (scm:define genericTopArgument
    (scm:lambda (dictBounded0)
      (scm:list (scm:cons (scm:string->symbol "genericTop'") (rt:record-ref dictBounded0 (scm:string->symbol "top"))))))

  (scm:define genericTop$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericTop'"))))

  (scm:define genericTopConstructor
    (scm:lambda (dictGenericTop0)
      (scm:list (scm:cons (scm:string->symbol "genericTop'") (rt:record-ref dictGenericTop0 (scm:string->symbol "genericTop'"))))))

  (scm:define genericTopProduct
    (scm:lambda (dictGenericTop0)
      (scm:let ([genericTop$p11 (rt:record-ref dictGenericTop0 (scm:string->symbol "genericTop'"))])
        (scm:lambda (dictGenericTop12)
          (scm:list (scm:cons (scm:string->symbol "genericTop'") (Data.Generic.Rep.Product* genericTop$p11 (rt:record-ref dictGenericTop12 (scm:string->symbol "genericTop'")))))))))

  (scm:define genericTopSum
    (scm:lambda (dictGenericTop0)
      (scm:list (scm:cons (scm:string->symbol "genericTop'") (Data.Generic.Rep.Inr (rt:record-ref dictGenericTop0 (scm:string->symbol "genericTop'")))))))

  (scm:define genericTop
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericTop1)
        ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (rt:record-ref dictGenericTop1 (scm:string->symbol "genericTop'"))))))

  (scm:define genericBottomNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericBottom'") Data.Generic.Rep.NoArguments)))

  (scm:define genericBottomArgument
    (scm:lambda (dictBounded0)
      (scm:list (scm:cons (scm:string->symbol "genericBottom'") (rt:record-ref dictBounded0 (scm:string->symbol "bottom"))))))

  (scm:define genericBottom$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericBottom'"))))

  (scm:define genericBottomConstructor
    (scm:lambda (dictGenericBottom0)
      (scm:list (scm:cons (scm:string->symbol "genericBottom'") (rt:record-ref dictGenericBottom0 (scm:string->symbol "genericBottom'"))))))

  (scm:define genericBottomProduct
    (scm:lambda (dictGenericBottom0)
      (scm:let ([genericBottom$p11 (rt:record-ref dictGenericBottom0 (scm:string->symbol "genericBottom'"))])
        (scm:lambda (dictGenericBottom12)
          (scm:list (scm:cons (scm:string->symbol "genericBottom'") (Data.Generic.Rep.Product* genericBottom$p11 (rt:record-ref dictGenericBottom12 (scm:string->symbol "genericBottom'")))))))))

  (scm:define genericBottomSum
    (scm:lambda (dictGenericBottom0)
      (scm:list (scm:cons (scm:string->symbol "genericBottom'") (Data.Generic.Rep.Inl (rt:record-ref dictGenericBottom0 (scm:string->symbol "genericBottom'")))))))

  (scm:define genericBottom
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericBottom1)
        ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (rt:record-ref dictGenericBottom1 (scm:string->symbol "genericBottom'")))))))
