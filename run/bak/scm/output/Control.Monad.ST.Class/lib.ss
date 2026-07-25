#!r6rs
#!chezscheme
(library
  (Control.Monad.ST.Class lib)
  (export
    liftST
    monadSTEffect
    monadSTST)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Monad.ST.Internal lib) Control.Monad.ST.Internal.)
    (prefix (Effect lib) Effect.)
    (prefix (Unsafe.Coerce lib) Unsafe.Coerce.))

  (scm:define monadSTST
    (scm:list (scm:cons (scm:string->symbol "liftST") (scm:lambda (x0)
      x0)) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Control.Monad.ST.Internal.monadST))))

  (scm:define monadSTEffect
    (scm:list (scm:cons (scm:string->symbol "liftST") Unsafe.Coerce.unsafeCoerce) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Effect.monadEffect))))

  (scm:define liftST
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "liftST")))))
