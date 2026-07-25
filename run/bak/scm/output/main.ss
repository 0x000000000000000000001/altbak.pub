#!chezscheme
(import (chezscheme)
        (only (purescm stack-trace) print-stack-trace)
        (only (purescm finalizers) run-finalizers)
        (only (App lib) main))
(base-exception-handler (lambda (e) (print-stack-trace e) (exit -1)))
(collect-request-handler (lambda () (collect) (run-finalizers)))
(main)