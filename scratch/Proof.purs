module Proof where

import Prelude
import Effect (Effect)

foreign import logA :: String -> Effect Unit
foreign import logB :: String -> Effect Unit

test :: Effect Unit
test = do
  logA "Hello from A"
  logB "Hello from B"
