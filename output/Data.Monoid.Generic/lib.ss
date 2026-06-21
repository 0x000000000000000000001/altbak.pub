#!r6rs
#!chezscheme
(library
  (Data.Monoid.Generic lib)
  (export
    genericMempty
    genericMempty$p
    genericMonoidArgument
    genericMonoidConstructor
    genericMonoidNoArguments
    genericMonoidProduct)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Generic.Rep lib) Data.Generic.Rep.))

  (scm:define genericMonoidNoArguments
    (scm:list (scm:cons (scm:string->symbol "genericMempty'") Data.Generic.Rep.NoArguments)))

  (scm:define genericMonoidArgument
    (scm:lambda (dictMonoid0)
      (scm:list (scm:cons (scm:string->symbol "genericMempty'") (rt:record-ref dictMonoid0 (scm:string->symbol "mempty"))))))

  (scm:define genericMempty$p
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "genericMempty'"))))

  (scm:define genericMonoidConstructor
    (scm:lambda (dictGenericMonoid0)
      (scm:list (scm:cons (scm:string->symbol "genericMempty'") (rt:record-ref dictGenericMonoid0 (scm:string->symbol "genericMempty'"))))))

  (scm:define genericMonoidProduct
    (scm:lambda (dictGenericMonoid0)
      (scm:let ([genericMempty$p11 (rt:record-ref dictGenericMonoid0 (scm:string->symbol "genericMempty'"))])
        (scm:lambda (dictGenericMonoid12)
          (scm:list (scm:cons (scm:string->symbol "genericMempty'") (Data.Generic.Rep.Product* genericMempty$p11 (rt:record-ref dictGenericMonoid12 (scm:string->symbol "genericMempty'")))))))))

  (scm:define genericMempty
    (scm:lambda (dictGeneric0)
      (scm:lambda (dictGenericMonoid1)
        ((rt:record-ref dictGeneric0 (scm:string->symbol "to")) (rt:record-ref dictGenericMonoid1 (scm:string->symbol "genericMempty'")))))))
