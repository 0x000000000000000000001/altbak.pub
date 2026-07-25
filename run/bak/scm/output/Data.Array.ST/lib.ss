#!r6rs
#!chezscheme
(library
  (Data.Array.ST lib)
  (export
    freeze
    freezeImpl
    length
    lengthImpl
    modify
    new
    peek
    peekImpl
    poke
    pokeImpl
    pop
    popImpl
    push
    pushAll
    pushAllImpl
    pushImpl
    run
    shift
    shiftImpl
    sort
    sortBy
    sortByImpl
    sortWith
    splice
    spliceImpl
    thaw
    thawImpl
    toAssocArray
    toAssocArrayImpl
    unsafeFreeze
    unsafeFreezeImpl
    unsafeThaw
    unsafeThawImpl
    unshift
    unshiftAll
    unshiftAllImpl
    withArray)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Control.Monad.ST.Internal lib) Control.Monad.ST.Internal.)
    (prefix (Control.Monad.ST.Uncurried lib) Control.Monad.ST.Uncurried.)
    (prefix (Data.Maybe lib) Data.Maybe.)
    (prefix (Data.Ordering lib) Data.Ordering.)
    (Data.Array.ST foreign))

  (scm:define unshiftAll
    (Control.Monad.ST.Uncurried.runSTFn2 unshiftAllImpl))

  (scm:define unshift
    (scm:lambda (a0)
      ((Control.Monad.ST.Uncurried.runSTFn2 unshiftAllImpl) (rt:make-array a0))))

  (scm:define unsafeThaw
    (Control.Monad.ST.Uncurried.runSTFn1 unsafeThawImpl))

  (scm:define unsafeFreeze
    (Control.Monad.ST.Uncurried.runSTFn1 unsafeFreezeImpl))

  (scm:define toAssocArray
    (Control.Monad.ST.Uncurried.runSTFn1 toAssocArrayImpl))

  (scm:define thaw
    (Control.Monad.ST.Uncurried.runSTFn1 thawImpl))

  (scm:define withArray
    (scm:lambda (f0)
      (scm:lambda (xs1)
        (scm:lambda ()
          (scm:let*
            ([result2 (thawImpl xs1)]
             [_ ((f0 result2))])
              ((scm:lambda ()
                (unsafeFreezeImpl result2))))))))

  (scm:define splice
    (Control.Monad.ST.Uncurried.runSTFn4 spliceImpl))

  (scm:define sortBy
    (scm:lambda (comp0)
      (((Control.Monad.ST.Uncurried.runSTFn3 sortByImpl) comp0) (scm:lambda (v1)
        (scm:cond
          [(Data.Ordering.GT? v1) 1]
          [(Data.Ordering.EQ? v1) 0]
          [(Data.Ordering.LT? v1) -1]
          [scm:else (rt:fail)])))))

  (scm:define sortWith
    (scm:lambda (dictOrd0)
      (scm:lambda (f1)
        (sortBy (scm:lambda (x2)
          (scm:lambda (y3)
            (((rt:record-ref dictOrd0 (scm:string->symbol "compare")) (f1 x2)) (f1 y3))))))))

  (scm:define sort
    (scm:lambda (dictOrd0)
      (sortBy (rt:record-ref dictOrd0 (scm:string->symbol "compare")))))

  (scm:define shift
    (((Control.Monad.ST.Uncurried.runSTFn3 shiftImpl) Data.Maybe.Just) Data.Maybe.Nothing))

  (scm:define run
    (scm:lambda (st0)
      (Control.Monad.ST.Internal.run (scm:lambda ()
        (scm:let ([_1 (st0)])
          ((scm:lambda ()
            (unsafeFreezeImpl _1))))))))

  (scm:define pushAll
    (Control.Monad.ST.Uncurried.runSTFn2 pushAllImpl))

  (scm:define push
    (Control.Monad.ST.Uncurried.runSTFn2 pushImpl))

  (scm:define pop
    (((Control.Monad.ST.Uncurried.runSTFn3 popImpl) Data.Maybe.Just) Data.Maybe.Nothing))

  (scm:define poke
    (Control.Monad.ST.Uncurried.runSTFn3 pokeImpl))

  (scm:define peek
    (((Control.Monad.ST.Uncurried.runSTFn4 peekImpl) Data.Maybe.Just) Data.Maybe.Nothing))

  (scm:define modify
    (scm:lambda (i0)
      (scm:lambda (f1)
        (scm:lambda (xs2)
          (scm:lambda ()
            (scm:let ([entry3 (peekImpl Data.Maybe.Just Data.Maybe.Nothing i0 xs2)])
              ((scm:cond
                [(Data.Maybe.Just? entry3) (scm:lambda ()
                  (pokeImpl i0 (f1 (Data.Maybe.Just-value0 entry3)) xs2))]
                [(Data.Maybe.Nothing? entry3) (scm:lambda ()
                  #f)]
                [scm:else (rt:fail)]))))))))

  (scm:define length
    (Control.Monad.ST.Uncurried.runSTFn1 lengthImpl))

  (scm:define freeze
    (Control.Monad.ST.Uncurried.runSTFn1 freezeImpl)))
