module Test.StateMonadFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runStateMonadFFI :: Int -> Int

describe :: Effect Unit
describe = log "State Monad FFI (1.2k Binds, 60 Stack Depth):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 60
  logShow $ runStateMonadFFI dummy
