#!r6rs
#!chezscheme
(library
  (Test.ListOps lib)
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
    filterEvens
    foldl
    range
    sumEvens)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.EuclideanRing lib) Data.EuclideanRing.)
    (prefix (Data.Semiring lib) Data.Semiring.)
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

  (scm:define range
    (scm:lambda (start0)
      (scm:lambda (end1)
        (scm:letrec ([go2 (scm:lambda (curr3)
          (scm:lambda (acc4)
            (scm:cond
              [(scm:fx<? curr3 start0) acc4]
              [scm:else ((go2 (scm:fx- curr3 1)) (Cons* curr3 acc4))])))])
          ((go2 end1) Nil)))))

  (scm:define foldl
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:cond
            [(Nil? v22) v11]
            [(Cons? v22) (((foldl v0) ((v0 v11) (Cons-value0 v22))) (Cons-value1 v22))]
            [scm:else (rt:fail)])))))

  (scm:define filterEvens
    (scm:lambda (lst0)
      (scm:letrec ([go1 (scm:lambda (v2)
        (scm:lambda (v13)
          (scm:cond
            [(Nil? v2) v13]
            [(Cons? v2) (scm:cond
              [(scm:fx=? ((Data.EuclideanRing.intMod (Cons-value0 v2)) 2) 0) ((go1 (Cons-value1 v2)) (Cons* (Cons-value0 v2) v13))]
              [scm:else ((go1 (Cons-value1 v2)) v13)])]
            [scm:else (rt:fail)])))])
        ((go1 lst0) Nil))))

  (scm:define sumEvens
    (((foldl Data.Semiring.intAdd) 0) (filterEvens (scm:letrec ([go0 (scm:lambda (curr1)
      (scm:lambda (acc2)
        (scm:cond
          [(scm:fx<? curr1 1) acc2]
          [scm:else ((go0 (scm:fx- curr1 1)) (Cons* curr1 acc2))])))])
      ((go0 90000) Nil)))))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "List Processing (90k elements):")))

  (scm:define act
    (Effect.Console.log (Data.Show.showIntImpl sumEvens))))
