#!r6rs
#!chezscheme
(library
  (Effect.Class lib)
  (export
    liftEffect
    monadEffectEffect)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Effect lib) Effect.))

  (scm:define monadEffectEffect
    (scm:list (scm:cons (scm:string->symbol "liftEffect") (scm:lambda (x0)
      x0)) (scm:cons (scm:string->symbol "Monad0") (scm:lambda (_)
      Effect.monadEffect))))

  (scm:define liftEffect
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "liftEffect")))))
