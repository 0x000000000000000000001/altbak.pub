#!r6rs
#!chezscheme
(library
  (Bench lib)
  (export
    benchNow
    formatNumber
    opaque
    runBench)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Effect.Console lib) Effect.Console.)
    (Bench foreign))

  (scm:define runBench
    (scm:lambda (describe0)
      (scm:lambda (act1)
        (scm:lambda ()
          (scm:let*
            ([_ (describe0)]
             [t13 (benchNow)]
             [_ (act1)]
             [t25 (benchNow)]
             [dt6 (scm:fl- t25 t13)]
             [_ ((Effect.Console.log (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "\nExecution time: ") (formatNumber dt6)) (rt:string->pstring " μs\n"))))])
              dt6))))))
