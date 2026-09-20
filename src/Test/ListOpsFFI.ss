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
        [(null? l) acc]
        [(p (car l)) (loop (cdr l) (cons (car l) acc))]
        [else (loop (cdr l) acc)])))

  (define (foldl combine initial lst)
    (let loop ([l lst] [acc initial])
      (if (null? l)
          acc
          (loop (cdr l) (combine acc (car l))))))

  (define (runListOpsFFI limit)
    (foldl + 0 (filter-list (lambda (x) (= (modulo x 2) 0)) (range-list 1 limit))))
)
