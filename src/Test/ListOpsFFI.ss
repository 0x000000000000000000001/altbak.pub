(library (Test.ListOpsFFI foreign)
  (export runListOpsFFI)
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

  (define (sum-list lst)
    (let loop ([l lst] [acc 0])
      (if (null? l)
          acc
          (loop (cdr l) (+ acc (car l))))))

  (define (runListOpsFFI limit)
    (sum-list (filter-list (lambda (x) (= (modulo x 2) 0)) (range-list 1 limit))))
)
