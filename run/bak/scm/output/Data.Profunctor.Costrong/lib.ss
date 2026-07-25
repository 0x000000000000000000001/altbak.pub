#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Costrong lib)
  (export
    unfirst
    unsecond)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define unsecond
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "unsecond"))))

  (scm:define unfirst
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "unfirst")))))
