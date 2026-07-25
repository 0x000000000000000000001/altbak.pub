#!r6rs
#!chezscheme
(library
  (Data.Profunctor.Closed lib)
  (export
    closed
    closedFunction)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Semigroupoid lib) Control.Semigroupoid.)
    (prefix (Data.Profunctor lib) Data.Profunctor.))

  (scm:define closedFunction
    (scm:list (scm:cons (scm:string->symbol "closed") (rt:record-ref Control.Semigroupoid.semigroupoidFn (scm:string->symbol "compose"))) (scm:cons (scm:string->symbol "Profunctor0") (scm:lambda (_)
      Data.Profunctor.profunctorFn))))

  (scm:define closed
    (scm:lambda (dict0)
      (rt:record-ref dict0 (scm:string->symbol "closed")))))
