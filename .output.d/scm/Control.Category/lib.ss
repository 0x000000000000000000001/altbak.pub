#!r6rs
#!chezscheme
(library
  (Control.Category lib)
  (export
    categoryFn
    identity)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Semigroupoid lib) Control.Semigroupoid.))

  (scm:define identity
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "identity"))))

  (scm:define categoryFn
    (scm:list (scm:cons (scm:string->symbol "identity") (scm:lambda (x0)
      x0)) (scm:cons (scm:string->symbol "Semigroupoid0") (scm:lambda (_)
      Control.Semigroupoid.semigroupoidFn)))))
