#!r6rs
#!chezscheme
(library
  (Type.Data.Row lib)
  (export
    RProxy
    RProxy?)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define RProxy
    (scm:quote RProxy))

  (scm:define RProxy?
    (scm:lambda (v)
      (scm:eq? (scm:quote RProxy) v))))
