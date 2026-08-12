module TestDPE where

import Prelude

class Monoidish a where
  mempty_ :: a
  mappend_ :: a -> a -> a

instance intMonoidish :: Monoidish Int where
  mempty_ = 0
  mappend_ a b = a + b

polyLoop :: forall a. Monoidish a => Int -> a -> a
polyLoop 0 acc = acc
polyLoop n acc = polyLoop (n - 1) (mappend_ acc mempty_)

test :: Int
test = polyLoop 10000000 0
