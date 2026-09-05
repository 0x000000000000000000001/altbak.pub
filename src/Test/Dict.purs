module Test.Dict where

import Prelude

class MyClass a where
  myMethod :: a -> Int

instance MyClass Int where
  myMethod x = x + 1

foo :: forall a. MyClass a => a -> Int
foo x = myMethod x

test :: Int -> Int
test = foo
