#!r6rs
#!chezscheme
(library
  (AppX lib)
  (export
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Test.Fib lib) Test.Fib.))

  (scm:define main
    (scm:lambda ()
      (scm:let ([_ (Test.Fib.describe)])
        (Test.Fib.act)))))
