#!r6rs
#!chezscheme
(library
  (Test.FileOps lib)
  (export
    act
    describe
    loopE
    loopIO
    readFileSync
    writeFileSync)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Unit lib) Data.Unit.)
    (prefix (Effect.Console lib) Effect.Console.)
    (Test.FileOps foreign))

  (scm:define loopIO
    (scm:lambda (n0)
      ((loopE n0) (scm:let ([_1 ((writeFileSync (rt:string->pstring "var/iotest.txt")) (rt:string->pstring "Hello IO Benchmarks!"))])
        (scm:lambda ()
          (scm:let*
            ([_ (_1)]
             [_ ((readFileSync (rt:string->pstring "var/iotest.txt")))])
              Data.Unit.unit))))))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "File I/O (10k writes/reads):")))

  (scm:define act
    (loopIO 10000)))
