module Test.Polymorphism where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

-- A custom type class to test runtime Dictionary dispatch
class Monoidish a where
  mempty_ :: a
  mappend_ :: a -> a -> a

instance intMonoidish :: Monoidish Int where
  mempty_ = 1
  mappend_ x y = x + y

polyLoop :: forall a. Monoidish a => Int -> a -> a
polyLoop n_init acc_init = go n_init acc_init
  where
  go 0 acc = acc
  go n acc = go (n - 1) (mappend_ acc mempty_)

describe :: Effect Unit
describe = log "Polymorphism (10M Type Class Dict Lookups):"

act :: Effect String
act = do
  dummy <- Bench.opaque 10000000
  pure (show ( polyLoop dummy 0))
