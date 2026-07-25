#!r6rs
#!chezscheme
(library
  (Data.Functor.Coproduct.Nested lib)
  (export
    at1
    at10
    at2
    at3
    at4
    at5
    at6
    at7
    at8
    at9
    coproduct1
    coproduct10
    coproduct2
    coproduct3
    coproduct4
    coproduct5
    coproduct6
    coproduct7
    coproduct8
    coproduct9
    in1
    in10
    in2
    in3
    in4
    in5
    in6
    in7
    in8
    in9)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Data.Either lib) Data.Either.)
    (prefix (Data.Functor.Coproduct lib) Data.Functor.Coproduct.))

  (scm:define in9
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Left v0)))))))))))

  (scm:define in8
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Left v0))))))))))

  (scm:define in7
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Left v0)))))))))

  (scm:define in6
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Left v0))))))))

  (scm:define in5
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Left v0)))))))

  (scm:define in4
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Left v0))))))

  (scm:define in3
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Left v0)))))

  (scm:define in2
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Left v0))))

  (scm:define in10
    (scm:lambda (v0)
      (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Right (Data.Either.Left v0))))))))))))

  (scm:define in1
    Data.Functor.Coproduct.left)

  (scm:define coproduct9
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (h7)
                    (scm:lambda (i8)
                      (scm:lambda (y9)
                        (scm:cond
                          [(Data.Either.Left? y9) (a0 (Data.Either.Left-value0 y9))]
                          [(Data.Either.Right? y9) (scm:cond
                            [(Data.Either.Left? (Data.Either.Right-value0 y9)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y9)))]
                            [(Data.Either.Right? (Data.Either.Right-value0 y9)) (scm:cond
                              [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y9))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))]
                              [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y9))) (scm:cond
                                [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))) (d3 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))]
                                [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))) (scm:cond
                                  [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))) (e4 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))))]
                                  [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))) (scm:cond
                                    [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))) (f5 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))))]
                                    [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))) (scm:cond
                                      [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))))) (g6 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))))))]
                                      [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))))) (scm:cond
                                        [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))))) (h7 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))))))]
                                        [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))))) (scm:cond
                                          [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))))))) (i8 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))))))))]
                                          [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9))))))))) (scm:letrec ([spin10 (scm:lambda (v11)
                                            (spin10 v11))])
                                            (spin10 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y9)))))))))))]
                                          [scm:else (rt:fail)])]
                                        [scm:else (rt:fail)])]
                                      [scm:else (rt:fail)])]
                                    [scm:else (rt:fail)])]
                                  [scm:else (rt:fail)])]
                                [scm:else (rt:fail)])]
                              [scm:else (rt:fail)])]
                            [scm:else (rt:fail)])]
                          [scm:else (rt:fail)]))))))))))))

  (scm:define coproduct8
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (h7)
                    (scm:lambda (y8)
                      (scm:cond
                        [(Data.Either.Left? y8) (a0 (Data.Either.Left-value0 y8))]
                        [(Data.Either.Right? y8) (scm:cond
                          [(Data.Either.Left? (Data.Either.Right-value0 y8)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y8)))]
                          [(Data.Either.Right? (Data.Either.Right-value0 y8)) (scm:cond
                            [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y8))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))]
                            [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y8))) (scm:cond
                              [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))) (d3 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))))]
                              [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))) (scm:cond
                                [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))) (e4 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))))]
                                [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))) (scm:cond
                                  [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))))) (f5 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))))))]
                                  [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))))) (scm:cond
                                    [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))))) (g6 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))))))]
                                    [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))))) (scm:cond
                                      [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))))))) (h7 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))))))))]
                                      [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8)))))))) (scm:letrec ([spin9 (scm:lambda (v10)
                                        (spin9 v10))])
                                        (spin9 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y8))))))))))]
                                      [scm:else (rt:fail)])]
                                    [scm:else (rt:fail)])]
                                  [scm:else (rt:fail)])]
                                [scm:else (rt:fail)])]
                              [scm:else (rt:fail)])]
                            [scm:else (rt:fail)])]
                          [scm:else (rt:fail)])]
                        [scm:else (rt:fail)])))))))))))

  (scm:define coproduct7
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (g6)
                  (scm:lambda (y7)
                    (scm:cond
                      [(Data.Either.Left? y7) (a0 (Data.Either.Left-value0 y7))]
                      [(Data.Either.Right? y7) (scm:cond
                        [(Data.Either.Left? (Data.Either.Right-value0 y7)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y7)))]
                        [(Data.Either.Right? (Data.Either.Right-value0 y7)) (scm:cond
                          [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y7))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7))))]
                          [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y7))) (scm:cond
                            [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7)))) (d3 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7)))))]
                            [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7)))) (scm:cond
                              [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7))))) (e4 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7))))))]
                              [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7))))) (scm:cond
                                [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7)))))) (f5 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7)))))))]
                                [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7)))))) (scm:cond
                                  [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7))))))) (g6 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7))))))))]
                                  [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7))))))) (scm:letrec ([spin8 (scm:lambda (v9)
                                    (spin8 v9))])
                                    (spin8 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y7)))))))))]
                                  [scm:else (rt:fail)])]
                                [scm:else (rt:fail)])]
                              [scm:else (rt:fail)])]
                            [scm:else (rt:fail)])]
                          [scm:else (rt:fail)])]
                        [scm:else (rt:fail)])]
                      [scm:else (rt:fail)]))))))))))

  (scm:define coproduct6
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (f5)
                (scm:lambda (y6)
                  (scm:cond
                    [(Data.Either.Left? y6) (a0 (Data.Either.Left-value0 y6))]
                    [(Data.Either.Right? y6) (scm:cond
                      [(Data.Either.Left? (Data.Either.Right-value0 y6)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y6)))]
                      [(Data.Either.Right? (Data.Either.Right-value0 y6)) (scm:cond
                        [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y6))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6))))]
                        [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y6))) (scm:cond
                          [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6)))) (d3 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6)))))]
                          [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6)))) (scm:cond
                            [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6))))) (e4 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6))))))]
                            [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6))))) (scm:cond
                              [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6)))))) (f5 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6)))))))]
                              [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6)))))) (scm:letrec ([spin7 (scm:lambda (v8)
                                (spin7 v8))])
                                (spin7 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y6))))))))]
                              [scm:else (rt:fail)])]
                            [scm:else (rt:fail)])]
                          [scm:else (rt:fail)])]
                        [scm:else (rt:fail)])]
                      [scm:else (rt:fail)])]
                    [scm:else (rt:fail)])))))))))

  (scm:define coproduct5
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (e4)
              (scm:lambda (y5)
                (scm:cond
                  [(Data.Either.Left? y5) (a0 (Data.Either.Left-value0 y5))]
                  [(Data.Either.Right? y5) (scm:cond
                    [(Data.Either.Left? (Data.Either.Right-value0 y5)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y5)))]
                    [(Data.Either.Right? (Data.Either.Right-value0 y5)) (scm:cond
                      [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y5))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5))))]
                      [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y5))) (scm:cond
                        [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5)))) (d3 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5)))))]
                        [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5)))) (scm:cond
                          [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5))))) (e4 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5))))))]
                          [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5))))) (scm:letrec ([spin6 (scm:lambda (v7)
                            (spin6 v7))])
                            (spin6 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y5)))))))]
                          [scm:else (rt:fail)])]
                        [scm:else (rt:fail)])]
                      [scm:else (rt:fail)])]
                    [scm:else (rt:fail)])]
                  [scm:else (rt:fail)]))))))))

  (scm:define coproduct4
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (d3)
            (scm:lambda (y4)
              (scm:cond
                [(Data.Either.Left? y4) (a0 (Data.Either.Left-value0 y4))]
                [(Data.Either.Right? y4) (scm:cond
                  [(Data.Either.Left? (Data.Either.Right-value0 y4)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y4)))]
                  [(Data.Either.Right? (Data.Either.Right-value0 y4)) (scm:cond
                    [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y4))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y4))))]
                    [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y4))) (scm:cond
                      [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y4)))) (d3 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y4)))))]
                      [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y4)))) (scm:letrec ([spin5 (scm:lambda (v6)
                        (spin5 v6))])
                        (spin5 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y4))))))]
                      [scm:else (rt:fail)])]
                    [scm:else (rt:fail)])]
                  [scm:else (rt:fail)])]
                [scm:else (rt:fail)])))))))

  (scm:define coproduct3
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (c2)
          (scm:lambda (y3)
            (scm:cond
              [(Data.Either.Left? y3) (a0 (Data.Either.Left-value0 y3))]
              [(Data.Either.Right? y3) (scm:cond
                [(Data.Either.Left? (Data.Either.Right-value0 y3)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y3)))]
                [(Data.Either.Right? (Data.Either.Right-value0 y3)) (scm:cond
                  [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y3))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y3))))]
                  [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y3))) (scm:letrec ([spin4 (scm:lambda (v5)
                    (spin4 v5))])
                    (spin4 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y3)))))]
                  [scm:else (rt:fail)])]
                [scm:else (rt:fail)])]
              [scm:else (rt:fail)]))))))

  (scm:define coproduct2
    (scm:lambda (a0)
      (scm:lambda (b1)
        (scm:lambda (y2)
          (scm:cond
            [(Data.Either.Left? y2) (a0 (Data.Either.Left-value0 y2))]
            [(Data.Either.Right? y2) (scm:cond
              [(Data.Either.Left? (Data.Either.Right-value0 y2)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y2)))]
              [(Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:letrec ([spin3 (scm:lambda (v4)
                (spin3 v4))])
                (spin3 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))]
              [scm:else (rt:fail)])]
            [scm:else (rt:fail)])))))

  (scm:define coproduct10
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
                        (scm:lambda (y10)
                          (scm:cond
                            [(Data.Either.Left? y10) (a0 (Data.Either.Left-value0 y10))]
                            [(Data.Either.Right? y10) (scm:cond
                              [(Data.Either.Left? (Data.Either.Right-value0 y10)) (b1 (Data.Either.Left-value0 (Data.Either.Right-value0 y10)))]
                              [(Data.Either.Right? (Data.Either.Right-value0 y10)) (scm:cond
                                [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y10))) (c2 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))]
                                [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y10))) (scm:cond
                                  [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))) (d3 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))]
                                  [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))) (scm:cond
                                    [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))) (e4 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))]
                                    [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))) (scm:cond
                                      [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))) (f5 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))))]
                                      [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))) (scm:cond
                                        [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))) (g6 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))))]
                                        [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))) (scm:cond
                                          [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))))) (h7 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))))))]
                                          [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))))) (scm:cond
                                            [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))))) (i8 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))))))]
                                            [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))))) (scm:cond
                                              [(Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))))))) (j9 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))))))))]
                                              [(Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10)))))))))) (scm:letrec ([spin11 (scm:lambda (v12)
                                                (spin11 v12))])
                                                (spin11 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y10))))))))))))]
                                              [scm:else (rt:fail)])]
                                            [scm:else (rt:fail)])]
                                          [scm:else (rt:fail)])]
                                        [scm:else (rt:fail)])]
                                      [scm:else (rt:fail)])]
                                    [scm:else (rt:fail)])]
                                  [scm:else (rt:fail)])]
                                [scm:else (rt:fail)])]
                              [scm:else (rt:fail)])]
                            [scm:else (rt:fail)])))))))))))))

  (scm:define coproduct1
    (scm:lambda (y0)
      (scm:cond
        [(Data.Either.Left? y0) (Data.Either.Left-value0 y0)]
        [(Data.Either.Right? y0) (scm:letrec ([spin1 (scm:lambda (v2)
          (spin1 v2))])
          (spin1 (Data.Either.Right-value0 y0)))]
        [scm:else (rt:fail)])))

  (scm:define at9
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))))) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))))))))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))))]
            [scm:else b0])))))

  (scm:define at8
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))))))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))))))]
            [scm:else b0])))))

  (scm:define at7
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))]
            [scm:else b0])))))

  (scm:define at6
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))))]
            [scm:else b0])))))

  (scm:define at5
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))]
            [scm:else b0])))))

  (scm:define at4
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))]
            [scm:else b0])))))

  (scm:define at3
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))]
            [scm:else b0])))))

  (scm:define at2
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (Data.Either.Left? (Data.Either.Right-value0 y2))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 y2)))]
            [scm:else b0])))))

  (scm:define at10
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(scm:and (Data.Either.Right? y2) (scm:and (Data.Either.Right? (Data.Either.Right-value0 y2)) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 y2))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))))) (scm:and (Data.Either.Right? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))) (Data.Either.Left? (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2))))))))))))))))))) (f1 (Data.Either.Left-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 (Data.Either.Right-value0 y2)))))))))))]
            [scm:else b0])))))

  (scm:define at1
    (scm:lambda (b0)
      (scm:lambda (f1)
        (scm:lambda (y2)
          (scm:cond
            [(Data.Either.Left? y2) (f1 (Data.Either.Left-value0 y2))]
            [scm:else b0]))))))
