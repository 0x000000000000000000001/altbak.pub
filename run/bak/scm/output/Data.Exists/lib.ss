#!r6rs
#!chezscheme
(library
  (Data.Exists lib)
  (export
    mkExists
    runExists)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define runExists
    Unsafe.Coerce.unsafeCoerce)

  (scm:define mkExists
    Unsafe.Coerce.unsafeCoerce))
