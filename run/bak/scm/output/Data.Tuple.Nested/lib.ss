#!r6rs
#!chezscheme
(library
  (Data.Tuple.Nested lib)
  (export
    curry1
    curry10
    curry2
    curry3
    curry4
    curry5
    curry6
    curry7
    curry8
    curry9
    get1
    get10
    get2
    get3
    get4
    get5
    get6
    get7
    get8
    get9
    over1
    over10
    over2
    over3
    over4
    over5
    over6
    over7
    over8
    over9
    tuple1
    tuple10
    tuple2
    tuple3
    tuple4
    tuple5
    tuple6
    tuple7
    tuple8
    tuple9
    uncurry1
    uncurry10
    uncurry2
    uncurry3
    uncurry4
    uncurry5
    uncurry6
    uncurry7
    uncurry8
    uncurry9)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define uncurry9
    (scm:lambda (f$p0)
      (scm:lambda (v1)
        (((((((((f$p0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))))

  (scm:define uncurry8
    (scm:lambda (f$p0)
      (scm:lambda (v1)
        ((((((((f$p0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))))))

  (scm:define uncurry7
    (scm:lambda (f$p0)
      (scm:lambda (v1)
        (((((((f$p0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))

  (scm:define uncurry6
    (scm:lambda (f$p0)
      (scm:lambda (v1)
        ((((((f$p0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))))

  (scm:define uncurry5
    (scm:lambda (f0)
      (scm:lambda (v1)
        (((((f0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))

  (scm:define uncurry4
    (scm:lambda (f0)
      (scm:lambda (v1)
        ((((f0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))

  (scm:define uncurry3
    (scm:lambda (f0)
      (scm:lambda (v1)
        (((f0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))

  (scm:define uncurry2
    (scm:lambda (f0)
      (scm:lambda (v1)
        ((f0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))))))

  (scm:define uncurry10
    (scm:lambda (f$p0)
      (scm:lambda (v1)
        ((((((((((f$p0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))) (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))))))))

  (scm:define uncurry1
    (scm:lambda (f0)
      (scm:lambda (v1)
        (f0 (Data.Tuple.Tuple-value0 v1)))))

  (scm:define tuple9
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (h7)
                    (scm:lambda (i8)
                      (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 (Data.Tuple.Tuple* g6 (Data.Tuple.Tuple* h7 (Data.Tuple.Tuple* i8 Data.Unit.unit)))))))))))))))))))

  (scm:define tuple8
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (h7)
                    (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 (Data.Tuple.Tuple* g6 (Data.Tuple.Tuple* h7 Data.Unit.unit)))))))))))))))))

  (scm:define tuple7
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 (Data.Tuple.Tuple* g6 Data.Unit.unit)))))))))))))))

  (scm:define tuple6
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 Data.Unit.unit)))))))))))))

  (scm:define tuple5
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 Data.Unit.unit)))))))))))

  (scm:define tuple4
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 Data.Unit.unit)))))))))

  (scm:define tuple3
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 Data.Unit.unit)))))))

  (scm:define tuple2
    (scm:lambda (a0)
      (scm:lambda (b1)
        (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 Data.Unit.unit)))))

  (scm:define tuple10
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (h7)
                    (scm:lambda (i8)
                      (scm:lambda (j9)
                        (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 (Data.Tuple.Tuple* g6 (Data.Tuple.Tuple* h7 (Data.Tuple.Tuple* i8 (Data.Tuple.Tuple* j9 Data.Unit.unit)))))))))))))))))))))

  (scm:define tuple1
    (scm:lambda (a0)
      (Data.Tuple.Tuple* a0 Data.Unit.unit)))

  (scm:define over9
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))))))))))))

  (scm:define over8
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))))))))))

  (scm:define over7
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))))))))

  (scm:define over6
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))))))

  (scm:define over5
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))))

  (scm:define over4
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))

  (scm:define over3
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))

  (scm:define over2
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))

  (scm:define over10
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 v1) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v1)) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))) (Data.Tuple.Tuple* (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))) (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1))))))))))) (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v1)))))))))))))))))))))))

  (scm:define over1
    (scm:lambda (o0)
      (scm:lambda (v1)
        (Data.Tuple.Tuple* (o0 (Data.Tuple.Tuple-value0 v1)) (Data.Tuple.Tuple-value1 v1)))))

  (scm:define get9
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0)))))))))))

  (scm:define get8
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0))))))))))

  (scm:define get7
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0)))))))))

  (scm:define get6
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0))))))))

  (scm:define get5
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0)))))))

  (scm:define get4
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0))))))

  (scm:define get3
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0)))))

  (scm:define get2
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 v0))))

  (scm:define get10
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 (Data.Tuple.Tuple-value1 v0))))))))))))

  (scm:define get1
    (scm:lambda (v0)
      (Data.Tuple.Tuple-value0 v0)))

  (scm:define curry9
    (scm:lambda (z0)
      (scm:lambda (f$p1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (scm:lambda (e6)
                  (scm:lambda (f7)
                    (scm:lambda (g8)
                      (scm:lambda (h9)
                        (scm:lambda (i10)
                          (f$p1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 (Data.Tuple.Tuple* d5 (Data.Tuple.Tuple* e6 (Data.Tuple.Tuple* f7 (Data.Tuple.Tuple* g8 (Data.Tuple.Tuple* h9 (Data.Tuple.Tuple* i10 z0))))))))))))))))))))))

  (scm:define curry8
    (scm:lambda (z0)
      (scm:lambda (f$p1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (scm:lambda (e6)
                  (scm:lambda (f7)
                    (scm:lambda (g8)
                      (scm:lambda (h9)
                        (f$p1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 (Data.Tuple.Tuple* d5 (Data.Tuple.Tuple* e6 (Data.Tuple.Tuple* f7 (Data.Tuple.Tuple* g8 (Data.Tuple.Tuple* h9 z0))))))))))))))))))))

  (scm:define curry7
    (scm:lambda (z0)
      (scm:lambda (f$p1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (scm:lambda (e6)
                  (scm:lambda (f7)
                    (scm:lambda (g8)
                      (f$p1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 (Data.Tuple.Tuple* d5 (Data.Tuple.Tuple* e6 (Data.Tuple.Tuple* f7 (Data.Tuple.Tuple* g8 z0))))))))))))))))))

  (scm:define curry6
    (scm:lambda (z0)
      (scm:lambda (f$p1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (scm:lambda (e6)
                  (scm:lambda (f7)
                    (f$p1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 (Data.Tuple.Tuple* d5 (Data.Tuple.Tuple* e6 (Data.Tuple.Tuple* f7 z0))))))))))))))))

  (scm:define curry5
    (scm:lambda (z0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (scm:lambda (e6)
                  (f1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 (Data.Tuple.Tuple* d5 (Data.Tuple.Tuple* e6 z0))))))))))))))

  (scm:define curry4
    (scm:lambda (z0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (f1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 (Data.Tuple.Tuple* d5 z0))))))))))))

  (scm:define curry3
    (scm:lambda (z0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (f1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 z0))))))))))

  (scm:define curry2
    (scm:lambda (z0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (f1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 z0))))))))

  (scm:define curry10
    (scm:lambda (z0)
      (scm:lambda (f$p1)
        (scm:lambda (a2)
          (scm:lambda (b3)
            (scm:lambda (c4)
              (scm:lambda (d5)
                (scm:lambda (e6)
                  (scm:lambda (f7)
                    (scm:lambda (g8)
                      (scm:lambda (h9)
                        (scm:lambda (i10)
                          (scm:lambda (j11)
                            (f$p1 (Data.Tuple.Tuple* a2 (Data.Tuple.Tuple* b3 (Data.Tuple.Tuple* c4 (Data.Tuple.Tuple* d5 (Data.Tuple.Tuple* e6 (Data.Tuple.Tuple* f7 (Data.Tuple.Tuple* g8 (Data.Tuple.Tuple* h9 (Data.Tuple.Tuple* i10 (Data.Tuple.Tuple* j11 z0))))))))))))))))))))))))

  (scm:define curry1
    (scm:lambda (z0)
      (scm:lambda (f1)
        (scm:lambda (a2)
          (f1 (Data.Tuple.Tuple* a2 z0)))))))
