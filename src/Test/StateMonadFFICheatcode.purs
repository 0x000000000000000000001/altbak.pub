module Test.StateMonadFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runStateMonadFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "State Monad FFICheatcode (1.2k Binds, 60 Stack Depth):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 60
  logShow $ runStateMonadFFICheatcode dummy
