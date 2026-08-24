(library (Test.PrimesFFI foreign)
  (export runPrimesFFI)
  (import (chezscheme))

  (define (range-list start end)
    (let loop ([curr end] [acc '()])
      (if (< curr start)
          acc
          (loop (- curr 1) (cons curr acc)))))

  (define (filter-list p lst)
    (let loop ([l lst] [acc '()])
      (cond
        [(null? l) (reverse acc)]
        [(p (car l)) (loop (cdr l) (cons (car l) acc))]
        [else (loop (cdr l) acc)])))

  (define (sieve lst)
    (if (null? lst)
        '()
        (let ([p (car lst)])
          (cons p (sieve (filter-list (lambda (x) (not (= (modulo x p) 0))) (cdr lst)))))))

  (define (sum-list lst)
    (let loop ([l lst] [acc 0])
      (if (null? l)
          acc
          (loop (cdr l) (+ acc (car l))))))

  (define (runPrimesFFI limit)
    (sum-list (sieve (range-list 2 limit))))
)
