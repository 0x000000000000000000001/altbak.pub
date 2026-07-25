#!r6rs
#!chezscheme
(library
  (Control.Monad.ST.Global lib)
  (export
    toEffect)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define toEffect
    Unsafe.Coerce.unsafeCoerce))
