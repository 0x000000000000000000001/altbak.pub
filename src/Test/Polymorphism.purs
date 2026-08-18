module Test.Polymorphism where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

-- A custom type class to test runtime Dictionary dispatch
class Monoidish a where
  mempty_ :: a
  mappend_ :: a -> a -> a

instance intMonoidish :: Monoidish Int where
  mempty_ = 1
  mappend_ x y = x + y

polyLoopGo :: forall a. Monoidish a => Int -> a -> a
polyLoopGo 0 acc = acc
polyLoopGo n acc = polyLoopGo (n - 1) (mappend_ acc mempty_)

polyLoop :: forall a. Monoidish a => Int -> a -> a
polyLoop n_init acc_init = polyLoopGo n_init acc_init

describe :: Effect Unit
describe = log "Polymorphism (10M Type Class Dict Lookups):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 10000000
  logShow $ polyLoop dummy 0
