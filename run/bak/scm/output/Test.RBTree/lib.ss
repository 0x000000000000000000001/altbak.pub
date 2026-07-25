#!r6rs
#!chezscheme
(library
  (Test.RBTree lib)
  (export
    B
    B?
    E
    E?
    R
    R?
    T
    T*
    T-value0
    T-value1
    T-value2
    T-value3
    T?
    act
    balance
    buildTree
    depth
    describe
    ins
    insert
    makeBlack
    max)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Bench lib) Bench.)
    (prefix (Data.Show lib) Data.Show.)
    (prefix (Effect.Console lib) Effect.Console.))

  (scm:define R
    (scm:quote R))

  (scm:define R?
    (scm:lambda (v)
      (scm:eq? (scm:quote R) v)))

  (scm:define B
    (scm:quote B))

  (scm:define B?
    (scm:lambda (v)
      (scm:eq? (scm:quote B) v)))

  (scm:define E
    (scm:quote E))

  (scm:define E?
    (scm:lambda (v)
      (scm:eq? (scm:quote E) v)))

  (scm:define-record-type (T$ T* T?)
    (scm:fields (scm:immutable value0 T-value0) (scm:immutable value1 T-value1) (scm:immutable value2 T-value2) (scm:immutable value3 T-value3)))

  (scm:define T
    (scm:lambda (value0)
      (scm:lambda (value1)
        (scm:lambda (value2)
          (scm:lambda (value3)
            (T* value0 value1 value2 value3))))))

  (scm:define max
    (scm:lambda (x0)
      (scm:lambda (y1)
        (scm:cond
          [(scm:fx>? x0 y1) x0]
          [scm:else y1]))))

  (scm:define makeBlack
    (scm:lambda (v0)
      (scm:cond
        [(T? v0) (T* B (T-value1 v0) (T-value2 v0) (T-value3 v0))]
        [(E? v0) E]
        [scm:else (rt:fail)])))

  (scm:define describe
    (Effect.Console.log (rt:string->pstring "Red-Black Tree (100k Worst-Case Insertions):")))

  (scm:define depth
    (scm:lambda (v0)
      (scm:cond
        [(E? v0) 0]
        [(T? v0) (scm:let*
          ([_1 (depth (T-value1 v0))]
           [_2 (depth (T-value3 v0))])
            (scm:cond
              [(scm:fx>? _1 _2) (scm:fx+ 1 _1)]
              [scm:else (scm:fx+ 1 _2)]))]
        [scm:else (rt:fail)])))

  (scm:define balance
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:lambda (v22)
          (scm:lambda (v33)
            (scm:cond
              [(B? v0) (scm:cond
                [(T? v11) (scm:cond
                  [(R? (T-value0 v11)) (scm:cond
                    [(T? (T-value1 v11)) (scm:cond
                      [(R? (T-value0 (T-value1 v11))) (T* R (T* B (T-value1 (T-value1 v11)) (T-value2 (T-value1 v11)) (T-value3 (T-value1 v11))) (T-value2 v11) (T* B (T-value3 v11) v22 v33))]
                      [(T? (T-value3 v11)) (scm:cond
                        [(R? (T-value0 (T-value3 v11))) (T* R (T* B (T-value1 v11) (T-value2 v11) (T-value1 (T-value3 v11))) (T-value2 (T-value3 v11)) (T* B (T-value3 (T-value3 v11)) v22 v33))]
                        [(scm:and (T? v33) (R? (T-value0 v33))) (scm:cond
                          [(T? (T-value1 v33)) (scm:cond
                            [(R? (T-value0 (T-value1 v33))) (T* R (T* B v11 v22 (T-value1 (T-value1 v33))) (T-value2 (T-value1 v33)) (T* B (T-value3 (T-value1 v33)) (T-value2 v33) (T-value3 v33)))]
                            [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                            [scm:else (T* v0 v11 v22 v33)])]
                          [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                          [scm:else (T* v0 v11 v22 v33)])]
                        [scm:else (T* v0 v11 v22 v33)])]
                      [(scm:and (T? v33) (R? (T-value0 v33))) (scm:cond
                        [(T? (T-value1 v33)) (scm:cond
                          [(R? (T-value0 (T-value1 v33))) (T* R (T* B v11 v22 (T-value1 (T-value1 v33))) (T-value2 (T-value1 v33)) (T* B (T-value3 (T-value1 v33)) (T-value2 v33) (T-value3 v33)))]
                          [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                          [scm:else (T* v0 v11 v22 v33)])]
                        [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                        [scm:else (T* v0 v11 v22 v33)])]
                      [scm:else (T* v0 v11 v22 v33)])]
                    [(T? (T-value3 v11)) (scm:cond
                      [(R? (T-value0 (T-value3 v11))) (T* R (T* B (T-value1 v11) (T-value2 v11) (T-value1 (T-value3 v11))) (T-value2 (T-value3 v11)) (T* B (T-value3 (T-value3 v11)) v22 v33))]
                      [(scm:and (T? v33) (R? (T-value0 v33))) (scm:cond
                        [(T? (T-value1 v33)) (scm:cond
                          [(R? (T-value0 (T-value1 v33))) (T* R (T* B v11 v22 (T-value1 (T-value1 v33))) (T-value2 (T-value1 v33)) (T* B (T-value3 (T-value1 v33)) (T-value2 v33) (T-value3 v33)))]
                          [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                          [scm:else (T* v0 v11 v22 v33)])]
                        [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                        [scm:else (T* v0 v11 v22 v33)])]
                      [scm:else (T* v0 v11 v22 v33)])]
                    [(scm:and (T? v33) (R? (T-value0 v33))) (scm:cond
                      [(T? (T-value1 v33)) (scm:cond
                        [(R? (T-value0 (T-value1 v33))) (T* R (T* B v11 v22 (T-value1 (T-value1 v33))) (T-value2 (T-value1 v33)) (T* B (T-value3 (T-value1 v33)) (T-value2 v33) (T-value3 v33)))]
                        [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                        [scm:else (T* v0 v11 v22 v33)])]
                      [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                      [scm:else (T* v0 v11 v22 v33)])]
                    [scm:else (T* v0 v11 v22 v33)])]
                  [(scm:and (T? v33) (R? (T-value0 v33))) (scm:cond
                    [(T? (T-value1 v33)) (scm:cond
                      [(R? (T-value0 (T-value1 v33))) (T* R (T* B v11 v22 (T-value1 (T-value1 v33))) (T-value2 (T-value1 v33)) (T* B (T-value3 (T-value1 v33)) (T-value2 v33) (T-value3 v33)))]
                      [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                      [scm:else (T* v0 v11 v22 v33)])]
                    [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                    [scm:else (T* v0 v11 v22 v33)])]
                  [scm:else (T* v0 v11 v22 v33)])]
                [(scm:and (T? v33) (R? (T-value0 v33))) (scm:cond
                  [(T? (T-value1 v33)) (scm:cond
                    [(R? (T-value0 (T-value1 v33))) (T* R (T* B v11 v22 (T-value1 (T-value1 v33))) (T-value2 (T-value1 v33)) (T* B (T-value3 (T-value1 v33)) (T-value2 v33) (T-value3 v33)))]
                    [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                    [scm:else (T* v0 v11 v22 v33)])]
                  [(scm:and (T? (T-value3 v33)) (R? (T-value0 (T-value3 v33)))) (T* R (T* B v11 v22 (T-value1 v33)) (T-value2 v33) (T* B (T-value1 (T-value3 v33)) (T-value2 (T-value3 v33)) (T-value3 (T-value3 v33))))]
                  [scm:else (T* v0 v11 v22 v33)])]
                [scm:else (T* v0 v11 v22 v33)])]
              [scm:else (T* v0 v11 v22 v33)]))))))

  (scm:define ins
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(E? v11) (T* R E v0 E)]
          [(T? v11) (scm:cond
            [(scm:fx<? v0 (T-value2 v11)) ((((balance (T-value0 v11)) ((ins v0) (T-value1 v11))) (T-value2 v11)) (T-value3 v11))]
            [(scm:fx>? v0 (T-value2 v11)) ((((balance (T-value0 v11)) (T-value1 v11)) (T-value2 v11)) ((ins v0) (T-value3 v11)))]
            [scm:else (T* (T-value0 v11) (T-value1 v11) (T-value2 v11) (T-value3 v11))])]
          [scm:else (rt:fail)]))))

  (scm:define insert
    (scm:lambda (x0)
      (scm:lambda (s1)
        (scm:let ([_2 ((ins x0) s1)])
          (scm:cond
            [(T? _2) (T* B (T-value1 _2) (T-value2 _2) (T-value3 _2))]
            [(E? _2) E]
            [scm:else (rt:fail)])))))

  (scm:define buildTree
    (scm:lambda (v0)
      (scm:lambda (v11)
        (scm:cond
          [(scm:fx=? v0 0) v11]
          [scm:else ((buildTree (scm:fx- v0 1)) ((insert v0) v11))]))))

  (scm:define act
    (scm:let ([_0 (Bench.opaque 100000)])
      (scm:lambda ()
        (scm:let ([dummy1 (_0)])
          ((Effect.Console.log (Data.Show.showIntImpl (depth ((buildTree dummy1) E))))))))))
