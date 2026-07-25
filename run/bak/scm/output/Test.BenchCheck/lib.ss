#!r6rs
#!chezscheme
(library
  (Test.BenchCheck lib)
  (export
    act)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define act
    (scm:lambda ()
      (scm:let*
        ([t10 (Bench.benchNow)]
         [t21 (Bench.benchNow)])
          ((Effect.Console.log (rt:pstring-concat (rt:string->pstring "Delta: ") (Data.Show.showNumberImpl (scm:fl- t21 t10)))))))))
