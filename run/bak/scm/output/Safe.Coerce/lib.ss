#!r6rs
#!chezscheme
(library
  (Safe.Coerce lib)
  (export
    coerce)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define coerce
    (scm:lambda (_)
      Unsafe.Coerce.unsafeCoerce)))
