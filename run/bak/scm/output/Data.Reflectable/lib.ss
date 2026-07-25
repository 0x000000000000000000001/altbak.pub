#!r6rs
#!chezscheme
(library
  (Data.Reflectable lib)
  (export
    reflectType
    reifiableBoolean
    reifiableInt
    reifiableOrdering
    reifiableString
    reifyType
    unsafeCoerce)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Reflectable foreign))

  (scm:define reifiableString
    (scm:list))

  (scm:define reifiableOrdering
    (scm:list))

  (scm:define reifiableInt
    (scm:list))

  (scm:define reifiableBoolean
    (scm:list))

  (scm:define reifyType
    (scm:lambda (_)
      (scm:lambda (s1)
        (scm:lambda (f2)
          (((unsafeCoerce (scm:lambda (dictReflectable3)
            (f2 dictReflectable3))) (scm:list (scm:cons (scm:string->symbol "reflectType") (scm:lambda (v3)
            s1)))) Type.Proxy.Proxy)))))

  (scm:define reflectType
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "reflectType")))))
