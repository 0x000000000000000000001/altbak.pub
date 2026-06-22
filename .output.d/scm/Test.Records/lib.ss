#!r6rs
#!chezscheme
(library
  (Test.Records lib)
  (export
    act
    describe
    initial
    updateRec)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.EuclideanRing lib) Data.EuclideanRing.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define updateRec
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) v11]
          [scm:else ((updateRec (scm:fx- v0 1)) (scm:cons (scm:cons (scm:string->symbol "a") (scm:fx+ (rt:record-ref v11 (scm:string->symbol "a")) 1)) (scm:cons (scm:cons (scm:string->symbol "b") (scm:cons (scm:cons (scm:string->symbol "c") (scm:fx+ (rt:record-ref (rt:record-ref v11 (scm:string->symbol "b")) (scm:string->symbol "c")) 2)) (scm:cons (scm:cons (scm:string->symbol "d") (scm:cons (scm:cons (scm:string->symbol "e") (scm:fx+ (rt:record-ref (rt:record-ref (rt:record-ref v11 (scm:string->symbol "b")) (scm:string->symbol "d")) (scm:string->symbol "e")) 3)) (scm:cons (scm:cons (scm:string->symbol "f") (scm:fx+ (rt:record-ref (rt:record-ref (rt:record-ref v11 (scm:string->symbol "b")) (scm:string->symbol "d")) (scm:string->symbol "f")) ((Data.EuclideanRing.intMod v0) 5))) (rt:record-ref (rt:record-ref v11 (scm:string->symbol "b")) (scm:string->symbol "d"))))) (rt:record-ref v11 (scm:string->symbol "b"))))) v11)))]))))

  (scm:define initial
    (scm:list (scm:cons (scm:string->symbol "a") 0) (scm:cons (scm:string->symbol "b") (scm:list (scm:cons (scm:string->symbol "c") 0) (scm:cons (scm:string->symbol "d") (scm:list (scm:cons (scm:string->symbol "e") 0) (scm:cons (scm:string->symbol "f") 0)))))))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Deep Record Updates (1M iterations):")))

  (scm:define act
    (Effect.Console.log (Data.Show.showIntImpl (rt:record-ref (rt:record-ref (rt:record-ref ((updateRec 1000000) initial) (scm:string->symbol "b")) (scm:string->symbol "d")) (scm:string->symbol "f"))))))
