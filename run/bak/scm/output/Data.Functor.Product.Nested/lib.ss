#!r6rs
#!chezscheme
(library
  (Data.Functor.Product.Nested lib)
  (export
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
    product1
    product10
    product2
    product3
    product4
    product5
    product6
    product7
    product8
    product9)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Tuple lib) Data.Tuple.)
    (prefix (Data.Unit lib) Data.Unit.))

  (scm:define product9
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

  (scm:define product8
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (h7)
                    (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 (Data.Tuple.Tuple* g6 (Data.Tuple.Tuple* h7 Data.Unit.unit)))))))))))))))))

  (scm:define product7
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 (Data.Tuple.Tuple* g6 Data.Unit.unit)))))))))))))))

  (scm:define product6
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 (Data.Tuple.Tuple* f5 Data.Unit.unit)))))))))))))

  (scm:define product5
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 (Data.Tuple.Tuple* e4 Data.Unit.unit)))))))))))

  (scm:define product4
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 (Data.Tuple.Tuple* d3 Data.Unit.unit)))))))))

  (scm:define product3
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 (Data.Tuple.Tuple* c2 Data.Unit.unit)))))))

  (scm:define product2
    (scm:lambda (a0)
      (scm:lambda (b1)
        (Data.Tuple.Tuple* a0 (Data.Tuple.Tuple* b1 Data.Unit.unit)))))

  (scm:define product10
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

  (scm:define product1
    (scm:lambda (a0)
      (Data.Tuple.Tuple* a0 Data.Unit.unit)))

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
      (Data.Tuple.Tuple-value0 v0))))
