#!r6rs
#!chezscheme
(library
  (AppX lib)
  (export
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Test.LazyEvaluation lib) Test.LazyEvaluation.))

  (scm:define main
    (scm:lambda ()
      (scm:let ([_ (Test.LazyEvaluation.describe)])
        (Test.LazyEvaluation.act)))))
