#!r6rs
#!chezscheme
(library
  (AppX lib)
  (export
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Test.TCO lib) Test.TCO.))

  (scm:define main
    (scm:let ([_0 ((Bench.runBench Test.TCO.describe) Test.TCO.act)])
      (scm:lambda ()
        (scm:let ([a$p1 (_0)])
          Data.Unit.unit)))))
