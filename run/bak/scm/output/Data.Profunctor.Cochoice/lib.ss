#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Cochoice lib)
  (export
    unleft
    unright)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define unright
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "unright"))))

  (scm:define unleft
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "unleft")))))
