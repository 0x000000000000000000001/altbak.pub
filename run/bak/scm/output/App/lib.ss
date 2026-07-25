#!r6rs
#!chezscheme
(library
  (App lib)
  (export
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Effect.Console lib) Effect.Console.)
    (prefix (Test.Ackermann lib) Test.Ackermann.)
    (prefix (Test.ArrayOps lib) Test.ArrayOps.)
    (prefix (Test.AstTree lib) Test.AstTree.)
    (prefix (Test.Church lib) Test.Church.)
    (prefix (Test.Fib lib) Test.Fib.)
    (prefix (Test.LazyEvaluation lib) Test.LazyEvaluation.)
    (prefix (Test.ListOps lib) Test.ListOps.)
    (prefix (Test.Polymorphism lib) Test.Polymorphism.)
    (prefix (Test.Primes lib) Test.Primes.)
    (prefix (Test.RBTree lib) Test.RBTree.)
    (prefix (Test.Records lib) Test.Records.)
    (prefix (Test.StateMonad lib) Test.StateMonad.)
    (prefix (Test.TCO lib) Test.TCO.))

  (scm:define main
    (scm:let ([_0 ((Bench.runBench Test.AstTree.describe) Test.AstTree.act)])
      (scm:lambda ()
        (scm:let*
          ([t11 (_0)]
           [t22 (((Bench.runBench Test.Fib.describe) Test.Fib.act))]
           [t33 (((Bench.runBench Test.ListOps.describe) Test.ListOps.act))]
           [t44 (((Bench.runBench Test.TCO.describe) Test.TCO.act))]
           [t55 (((Bench.runBench Test.Records.describe) Test.Records.act))]
           [t66 (((Bench.runBench Test.Ackermann.describe) Test.Ackermann.act))]
           [t77 (((Bench.runBench Test.Church.describe) Test.Church.act))]
           [t88 (((Bench.runBench Test.Primes.describe) Test.Primes.act))]
           [t99 (((Bench.runBench Test.RBTree.describe) Test.RBTree.act))]
           [t1010 (((Bench.runBench Test.Polymorphism.describe) Test.Polymorphism.act))]
           [t1111 (((Bench.runBench Test.StateMonad.describe) Test.StateMonad.act))]
           [t1212 (((Bench.runBench Test.LazyEvaluation.describe) Test.LazyEvaluation.act))]
           [t1313 (((Bench.runBench Test.ArrayOps.describe) Test.ArrayOps.act))])
            ((Effect.Console.log (rt:pstring-concat (rt:pstring-concat (rt:string->pstring "Total exec time: ") (Bench.formatNumber (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl+ (scm:fl/ t11 1000.0) (scm:fl/ t22 1000.0)) (scm:fl/ t33 1000.0)) (scm:fl/ t44 1000.0)) (scm:fl/ t55 1000.0)) (scm:fl/ t66 1000.0)) (scm:fl/ t77 1000.0)) (scm:fl/ t88 1000.0)) (scm:fl/ t99 1000.0)) (scm:fl/ t1010 1000.0)) (scm:fl/ t1111 1000.0)) (scm:fl/ t1212 1000.0)) (scm:fl/ t1313 1000.0)))) (rt:string->pstring " ms\n")))))))))
