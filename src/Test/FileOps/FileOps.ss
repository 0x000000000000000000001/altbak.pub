(library (Test.FileOps foreign)
  (export writeFileSync readFileSync loopE)
  (import (chezscheme))

  (define writeFileSync
    (lambda (path)
      (lambda (content)
        (lambda ()
          (when (file-exists? path) (delete-file path))
          (let ((p (open-output-file path)))
            (put-string p content)
            (close-output-port p))
          '#()))))

  (define readFileSync
    (lambda (path)
      (lambda ()
        (let* ((p (open-input-file path))
               (content (get-string-all p)))
          (close-input-port p)
          content))))

  (define loopE
    (lambda (n)
      (lambda (action)
        (lambda ()
          (let loop ((i n))
            (unless (zero? i)
              (action)
              (loop (- i 1))))))))
)
