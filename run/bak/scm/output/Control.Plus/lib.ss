#!r6rs
#!chezscheme
(library
  (Control.Plus lib)
  (export
    empty
    plusArray)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Alt lib) Control.Alt.))

  (scm:define plusArray
    (scm:list (scm:cons (scm:string->symbol "empty") (rt:make-array)) (scm:cons (scm:string->symbol "Alt0") (scm:lambda (_)
      Control.Alt.altArray))))

  (scm:define empty
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "empty")))))
