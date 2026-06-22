#!r6rs
#!chezscheme
(library
  (Data.Show.Generic lib)
  (export
    genericShow
    genericShow$p
    genericShowArgs
    genericShowArgsArgument
    genericShowArgsNoArguments
    genericShowArgsProduct
    genericShowConstructor
    genericShowNoConstructors
    genericShowSum
    intercalate)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.)
    (prefix (Data.Semigroup lib) Data.Semigroup.)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Show.Generic foreign))

  (scm:define genericShowArgsNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericShowArgs") (scm:lambda (v0)
      (rt:make-array)))))

  (scm:define genericShowArgsArgument
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "genericShowArgs") (scm:lambda (v1)
        (rt:make-array ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)))))))

  (scm:define genericShowArgs
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericShowArgs"))))

  (scm:define genericShowArgsProduct
    (scm:lambda (dictGenericShowArgs0)
      (scm:lambda (dictGenericShowArgs11)
        (scm:list (scm:cons (scm:string->symbol "genericShowArgs") (scm:lambda (v2)
          ((Data.Semigroup.concatArray ((rt:record-ref dictGenericShowArgs0 (scm:string->symbol "genericShowArgs")) (Data.Generic.Rep.Product-value0 v2))) ((rt:record-ref dictGenericShowArgs11 (scm:string->symbol "genericShowArgs")) (Data.Generic.Rep.Product-value1 v2)))))))))

  (scm:define genericShowConstructor
    (scm:lambda (dictGenericShowArgs0)
      (scm:lambda (dictIsSymbol1)
        (scm:list (scm:cons (scm:string->symbol "genericShow'") (scm:lambda (v2)
          (scm:let*
            ([ctor3 ((rt:record-ref dictIsSymbol1 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy)]
             [v14 ((rt:record-ref dictGenericShowArgs0 (scm:string->symbol "genericShowArgs")) v2)])
              (scm:cond
                [(scm:fx=? (rt:array-length v14) 0) ctor3]
                [scm:else (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(") ((intercalate (rt:string->pstring " ")) ((Data.Semigroup.concatArray (rt:make-array ctor3)) v14))) (rt:string->pstring ")"))]))))))))

  (scm:define genericShow$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericShow'"))))

  (rt:define-lazy $lazy-genericShowNoConstructors "genericShowNoConstructors" "Data.Show.Generic"
    (scm:list (scm:cons (scm:string->symbol "genericShow'") (scm:lambda (a0)
      ((rt:record-ref ($lazy-genericShowNoConstructors) (scm:string->symbol "genericShow'")) a0)))))

  (scm:define genericShowNoConstructors
    ($lazy-genericShowNoConstructors))

  (scm:define genericShowSum
    (scm:lambda (dictGenericShow0)
      (scm:lambda (dictGenericShow11)
        (scm:list (scm:cons (scm:string->symbol "genericShow'") (scm:lambda (v2)
          (scm:cond
            [(Data.Generic.Rep.Inl? v2) ((rt:record-ref dictGenericShow0 (scm:string->symbol "genericShow'")) (Data.Generic.Rep.Inl-value0 v2))]
            [(Data.Generic.Rep.Inr? v2) ((rt:record-ref dictGenericShow11 (scm:string->symbol "genericShow'")) (Data.Generic.Rep.Inr-value0 v2))]
            [scm:else (rt:fail)])))))))

  (scm:define genericShow
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericShow1)
        (scm:lambda (x2)
          ((rt:record-ref dictGenericShow1 (scm:string->symbol "genericShow'")) ((rt:record-ref dictGeneric0 (scm:string->symbol "from")) x2)))))))
