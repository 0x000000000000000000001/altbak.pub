#!r6rs
#!chezscheme
(library
  (AppX lib)
  (export
    main)
  (import
    (prefix (chezscheme) scm:)
    (prefix (purescm runtime) rt:)
    (prefix (Test.RBTree lib) Test.RBTree.))

  (scm:define main
    (scm:lambda ()
      (scm:let ([_ (Test.RBTree.describe)])
        (Test.RBTree.act)))))
