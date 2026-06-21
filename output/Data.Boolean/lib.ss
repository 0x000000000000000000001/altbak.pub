#!r6rs
#!chezscheme
(library
  (Data.Boolean lib)
  (export
    otherwise)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:))

  (scm:define otherwise
    #t))
