#!r6rs
#!chezscheme
(library
  (Data.Symbol lib)
  (export
    reflectSymbol
    reifySymbol
    unsafeCoerce)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Type.Proxy lib) Type.Proxy.)
    (Data.Symbol foreign))

  (scm:define reifySymbol
    (scm:lambda (s0)
      (scm:lambda (f1)
        (((unsafeCoerce (scm:lambda (dictIsSymbol2)
          (f1 dictIsSymbol2))) (scm:list (scm:cons (scm:string->symbol "reflectSymbol") (scm:lambda (v2)
          s0)))) Type.Proxy.Proxy))))

  (scm:define reflectSymbol
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "reflectSymbol")))))
