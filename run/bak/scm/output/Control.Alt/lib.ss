#!r6rs
#!chezscheme
(library
  (Control.Alt lib)
  (export
    alt
    altArray)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Functor lib) Data.Functor.)
    (prefix (Data.Semigroup lib) Data.Semigroup.))

  (scm:define altArray
    (scm:list (scm:cons (scm:string->symbol "alt") Data.Semigroup.concatArray) (scm:cons (scm:string->symbol "Functor0") (scm:lambda (_)
      Data.Functor.functorArray))))

  (scm:define alt
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "alt")))))
