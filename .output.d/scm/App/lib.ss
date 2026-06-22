#!r6rs
#!chezscheme
(library
  (App lib)
  (export
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Test.Ackermann lib) Test.Ackermann.)
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
    (scm:lambda ()
      (scm:let*
        ([_ (Test.AstTree.describe)]
         [_ (Test.AstTree.act)]
         [_ (Test.Fib.describe)]
         [_ (Test.Fib.act)]
         [_ (Test.ListOps.describe)]
         [_ (Test.ListOps.act)]
         [_ (Test.TCO.describe)]
         [_ (Test.TCO.act)]
         [_ (Test.Records.describe)]
         [_ (Test.Records.act)]
         [_ (Test.Ackermann.describe)]
         [_ (Test.Ackermann.act)]
         [_ (Test.Church.describe)]
         [_ (Test.Church.act)]
         [_ (Test.Primes.describe)]
         [_ (Test.Primes.act)]
         [_ (Test.RBTree.describe)]
         [_ (Test.RBTree.act)]
         [_ (Test.Polymorphism.describe)]
         [_ (Test.Polymorphism.act)]
         [_ (Test.StateMonad.describe)]
         [_ (Test.StateMonad.act)]
         [_ (Test.LazyEvaluation.describe)])
          (Test.LazyEvaluation.act)))))
