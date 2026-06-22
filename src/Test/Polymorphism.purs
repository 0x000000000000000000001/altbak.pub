module Test.Polymorphism where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

-- A custom type class to test runtime Dictionary dispatch
class Monoidish a where
  mempty_ :: a
  mappend_ :: a -> a -> a

instance intMonoidish :: Monoidish Int where
  mempty_ = 1
  mappend_ x y = x + y

-- In PureScript, this polymorphic function receives a hidden "dict" argument.
-- Inside the loop, it forces dynamic property lookups on this dictionary 
-- for `mappend_` and `mempty_`.
-- This will run 10 million times, causing 20 million dictionary lookups.
polyLoop :: forall a. Monoidish a => Int -> a -> a
polyLoop n_init acc_init = go n_init acc_init
  where
  go 0 acc = acc
  go n acc = go (n - 1) (mappend_ acc mempty_)

describe :: Effect Unit
describe = log "Polymorphism (10M Type Class Dict Lookups):"

act :: Effect Unit
act = logShow $ polyLoop 10000000 0
