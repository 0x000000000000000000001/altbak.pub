#!r6rs
#!chezscheme
(library
  (Data.Array.ST.Iterator lib)
  (export
    Iterator
    Iterator*
    Iterator-value0
    Iterator-value1
    Iterator?
    exhausted
    iterate
    iterator
    next
    peek
    pushAll
    pushWhile)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Monad.ST.Internal lib) Control.Monad.ST.Internal.)
    (prefix (Data.Array.ST lib) Data.Array.ST.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define-record-type (Iterator$ Iterator* Iterator?)
    (scm:fields (scm:immutable value0 Iterator-value0) (scm:immutable value1 Iterator-value1)))

  (scm:define Iterator
    (scm:lambda (value0)
      (scm:lambda (value1)
        (Iterator* value0 value1))))

  (scm:define peek
    (scm:lambda (v0)
      (scm:let ([_1 (Iterator-value1 v0)])
        (scm:lambda ()
          (scm:let ([i2 (scm:unbox _1)])
            ((Iterator-value0 v0) i2))))))

  (scm:define next
    (scm:lambda (v0)
      (scm:let ([_1 (Iterator-value1 v0)])
        (scm:lambda ()
          (scm:let*
            ([i2 (scm:unbox _1)]
             [_3 (scm:unbox _1)]
             [_ (scm:begin (scm:set-box! _1 (scm:fx+ _3 1)) (scm:unbox _1))])
              ((Iterator-value0 v0) i2))))))

  (scm:define pushWhile
    (scm:lambda (p0)
      (scm:lambda (iter1)
        (scm:lambda (array2)
          (scm:lambda ()
            (scm:let ([break3 (scm:box #f)])
              (((Control.Monad.ST.Internal.while (scm:lambda ()
                (scm:let ([_4 (scm:unbox break3)])
                  (scm:not _4)))) (scm:let ([_4 (Iterator-value1 iter1)])
                (scm:lambda ()
                  (scm:let*
                    ([i5 (scm:unbox _4)]
                     [mx6 ((Iterator-value0 iter1) i5)])
                      ((scm:cond
                        [(scm:and (Data.Maybe.Just? mx6) (p0 (Data.Maybe.Just-value0 mx6))) (scm:lambda ()
                          (scm:let*
                            ([_ (Data.Array.ST.pushImpl (Data.Maybe.Just-value0 mx6) array2)]
                             [i8 (scm:unbox (Iterator-value1 iter1))]
                             [_9 (scm:unbox (Iterator-value1 iter1))]
                             [_ (scm:begin (scm:set-box! (Iterator-value1 iter1) (scm:fx+ _9 1)) (scm:unbox (Iterator-value1 iter1)))])
                              Data.Unit.unit))]
                        [scm:else (scm:lambda ()
                          (scm:let ([_7 (scm:begin (scm:set-box! break3 #t) (scm:unbox break3))])
                            Data.Unit.unit))])))))))))))))

  (scm:define pushAll
    (pushWhile (scm:lambda (v0)
      #t)))

  (scm:define iterator
    (scm:lambda (f0)
      (scm:let ([_1 (Iterator f0)])
        (scm:lambda ()
          (scm:let ([_2 (scm:box 0)])
            (_1 _2))))))

  (scm:define iterate
    (scm:lambda (iter0)
      (scm:lambda (f1)
        (scm:lambda ()
          (scm:let ([break2 (scm:box #f)])
            (((Control.Monad.ST.Internal.while (scm:lambda ()
              (scm:let ([_3 (scm:unbox break2)])
                (scm:not _3)))) (scm:let ([_3 (Iterator-value1 iter0)])
              (scm:lambda ()
                (scm:let*
                  ([i4 (scm:unbox _3)]
                   [_5 (scm:unbox _3)]
                   [_ (scm:begin (scm:set-box! _3 (scm:fx+ _5 1)) (scm:unbox _3))]
                   [mx7 ((Iterator-value0 iter0) i4)])
                    ((scm:cond
                      [(Data.Maybe.Just? mx7) (f1 (Data.Maybe.Just-value0 mx7))]
                      [(Data.Maybe.Nothing? mx7) (scm:lambda ()
                        (scm:let ([_8 (scm:begin (scm:set-box! break2 #t) (scm:unbox break2))])
                          Data.Unit.unit))]
                      [scm:else (rt:fail)]))))))))))))

  (scm:define exhausted
    (scm:lambda (x0)
      (scm:let ([_1 (Iterator-value1 x0)])
        (scm:lambda ()
          (scm:let*
            ([i2 (scm:unbox _1)]
             [_3 ((Iterator-value0 x0) i2)])
              (scm:cond
                [(Data.Maybe.Nothing? _3) #t]
                [(Data.Maybe.Just? _3) #f]
                [scm:else (rt:fail)])))))))
