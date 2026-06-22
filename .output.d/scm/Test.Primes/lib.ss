#!r6rs
#!chezscheme
(library
  (Test.Primes lib)
  (export
    Cons
    Cons*
    Cons-value0
    Cons-value1
    Cons?
    Nil
    Nil?
    act
    describe
    filter
    range
    reverse
    sieve
    sumList)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.EuclideanRing lib) Data.EuclideanRing.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define Nil
    (scm:quote Nil))

  (scm:define Nil?
    (scm:lambda (v)
      (scm:eq? (scm:quote Nil) v)))

  (scm:define-record-type (Cons$ Cons* Cons?)
    (scm:fields (scm:immutable value0 Cons-value0) (scm:immutable value1 Cons-value1)))

  (scm:define Cons
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Cons* value0 value1))))

  (scm:define sumList
    (scm:lambda (lst0)
      (scm:letrec ([go1 (scm:lambda (v2)
        (scm:lambda (v13)
          (scm:cond
            [(Nil? v2) v13]
            [(Cons? v2) ((go1 (Cons-value1 v2)) (scm:fx+ v13 (Cons-value0 v2)))]
            [scm:else (rt:fail)])))])
        ((go1 lst0) 0))))

  (scm:define reverse
    (scm:lambda (lst0)
      (scm:letrec ([go1 (scm:lambda (v2)
        (scm:lambda (v13)
          (scm:cond
            [(Nil? v2) v13]
            [(Cons? v2) ((go1 (Cons-value1 v2)) (Cons* (Cons-value0 v2) v13))]
            [scm:else (rt:fail)])))])
        ((go1 lst0) Nil))))

  (scm:define range
    (scm:lambda (start0)
      (scm:lambda (end1)
        (scm:letrec ([go2 (scm:lambda (curr3)
          (scm:lambda (acc4)
            (scm:cond
              [(scm:fx<? curr3 start0) acc4]
              [scm:else ((go2 (scm:fx- curr3 1)) (Cons* curr3 acc4))])))])
          ((go2 end1) Nil)))))

  (scm:define filter
    (scm:lambda (p0)
      (scm:lambda (lst1)
        (scm:letrec ([go2 (scm:lambda (v3)
          (scm:lambda (v14)
            (scm:cond
              [(Nil? v3) (scm:letrec ([go5 (scm:lambda (v6)
                (scm:lambda (v17)
                  (scm:cond
                    [(Nil? v6) v17]
                    [(Cons? v6) ((go5 (Cons-value1 v6)) (Cons* (Cons-value0 v6) v17))]
                    [scm:else (rt:fail)])))])
                ((go5 v14) Nil))]
              [(Cons? v3) (scm:cond
                [(p0 (Cons-value0 v3)) ((go2 (Cons-value1 v3)) (Cons* (Cons-value0 v3) v14))]
                [scm:else ((go2 (Cons-value1 v3)) v14)])]
              [scm:else (rt:fail)])))])
          ((go2 lst1) Nil)))))

  (scm:define sieve
    (scm:lambda (v0)
      (scm:cond
        [(Nil? v0) Nil]
        [(Cons? v0) (scm:let ([_1 (Cons-value0 v0)])
          (Cons* _1 (sieve (scm:letrec ([go2 (scm:lambda (v3)
            (scm:lambda (v14)
              (scm:cond
                [(Nil? v3) (scm:letrec ([go5 (scm:lambda (v6)
                  (scm:lambda (v17)
                    (scm:cond
                      [(Nil? v6) v17]
                      [(Cons? v6) ((go5 (Cons-value1 v6)) (Cons* (Cons-value0 v6) v17))]
                      [scm:else (rt:fail)])))])
                  ((go5 v14) Nil))]
                [(Cons? v3) (scm:cond
                  [(scm:not (scm:fx=? ((Data.EuclideanRing.intMod (Cons-value0 v3)) _1) 0)) ((go2 (Cons-value1 v3)) (Cons* (Cons-value0 v3) v14))]
                  [scm:else ((go2 (Cons-value1 v3)) v14)])]
                [scm:else (rt:fail)])))])
            ((go2 (Cons-value1 v0)) Nil)))))]
        [scm:else (rt:fail)])))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Prime Sieve (sum primes up to 20k):")))

  (scm:define act
    (scm:letrec ([go0 (scm:lambda (v1)
      (scm:lambda (v12)
        (scm:cond
          [(Nil? v1) v12]
          [(Cons? v1) ((go0 (Cons-value1 v1)) (scm:fx+ v12 (Cons-value0 v1)))]
          [scm:else (rt:fail)])))])
      (Effect.Console.log (Data.Show.showIntImpl ((go0 (sieve (scm:letrec ([go1 (scm:lambda (curr2)
        (scm:lambda (acc3)
          (scm:cond
            [(scm:fx<? curr2 2) acc3]
            [scm:else ((go1 (scm:fx- curr2 1)) (Cons* curr2 acc3))])))])
        ((go1 20000) Nil)))) 0))))))
