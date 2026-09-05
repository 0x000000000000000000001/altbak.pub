module Scratch.TC where

import Prelude
import Effect.Console (logShow)

class MyClass a where
  myMethod :: a -> String

instance MyClass Int where
  myMethod _ = "Int"

main = logShow (myMethod 42)
