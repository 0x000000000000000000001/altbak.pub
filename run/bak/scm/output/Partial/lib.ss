#!r6rs
#!chezscheme
(library
  (Partial lib)
  (export
    _crashWith
    crash
    crashWith)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (Partial foreign))

  (scm:define crashWith
    (scm:lambda (_)
      _crashWith))

  (scm:define crash
    (scm:lambda (_)
      (_crashWith (rt:string->pstring "Partial.crash: partial function")))))
