module Test.Implicit where

foo :: forall a. a -> a
foo x = x

bar :: Int
bar = foo 42
