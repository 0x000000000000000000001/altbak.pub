#!r6rs
#!chezscheme
(library
  (Control.MonadPlus lib)
  (export
    monadPlusArray)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Alternative lib) Control.Alternative.)
    (prefix (Control.Monad lib) Control.Monad.))

  (scm:define monadPlusArray
    (scm:list (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Control.Monad.monadArray)) (scm:cons (scm:string->symbol "Alternative1") (scm:lambda (_)
      Control.Alternative.alternativeArray)))))
