#!r6rs
#!chezscheme
(library
  (Data.Generic.Rep lib)
  (export
    Argument
    Constructor
    Inl
    Inl-value0
    Inl?
    Inr
    Inr-value0
    Inr?
    NoArguments
    NoArguments?
    Product
    Product*
    Product-value0
    Product-value1
    Product?
    from
    repOf
    showArgument
    showConstructor
    showNoArguments
    showProduct
    showSum
    to)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Type.Proxy lib) Type.Proxy.))

  (scm:define-record-type (Inl$ Inl Inl?)
    (scm:fields (scm:immutable value0 Inl-value0)))

  (scm:define-record-type (Inr$ Inr Inr?)
    (scm:fields (scm:immutable value0 Inr-value0)))

  (scm:define-record-type (Product$ Product* Product?)
    (scm:fields (scm:immutable value0 Product-value0) (scm:immutable value1 Product-value1)))

  (scm:define Product
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Product* value0 value1))))

  (scm:define NoArguments
    (scm:quote NoArguments))

  (scm:define NoArguments?
    (scm:lambda (v)
      (scm:eq? (scm:quote NoArguments) v)))

  (scm:define Constructor
    (scm:lambda (x0)
      x0))

  (scm:define Argument
    (scm:lambda (x0)
      x0))

  (scm:define to
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "to"))))

  (scm:define showSum
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (scm:cond
            [(Inl? v2) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Inl ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Inl-value0 v2))) (rt:string->pstring ")"))]
            [(Inr? v2) (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Inr ") ((rt:record-ref dictShow11 (scm:string->symbol "show")) (Inr-value0 v2))) (rt:string->pstring ")"))]
            [scm:else (rt:fail)])))))))

  (scm:define showProduct
    (scm:lambda (dictShow0)
      (scm:lambda (dictShow11)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Product ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) (Product-value0 v2))) (rt:string->pstring " ")) ((rt:record-ref dictShow11 (scm:string->symbol "show")) (Product-value1 v2))) (rt:string->pstring ")"))))))))

  (scm:define showNoArguments
    (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v0)
      (rt:string->pstring "NoArguments")))))

  (scm:define showConstructor
    (scm:lambda (dictIsSymbol0)
      (scm:lambda (dictShow1)
        (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v2)
          (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Constructor @") (Data.Show.showStringImpl ((rt:record-ref dictIsSymbol0 (scm:string->symbol "reflectSymbol")) Type.Proxy.Proxy))) (rt:string->pstring " ")) ((rt:record-ref dictShow1 (scm:string->symbol "show")) v2)) (rt:string->pstring ")"))))))))

  (scm:define showArgument
    (scm:lambda (dictShow0)
      (scm:list (scm:cons (scm:string->symbol "show") (scm:lambda (v1)
        (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "(Argument ") ((rt:record-ref dictShow0 (scm:string->symbol "show")) v1)) (rt:string->pstring ")")))))))

  (scm:define repOf
    (scm:lambda (dictGeneric0)
      (scm:lambda (v1)
        Type.Proxy.Proxy)))

  (scm:define from
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "from")))))
