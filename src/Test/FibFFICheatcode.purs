module Test.FibFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runFibFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Fibonacci FFICheatcode:"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 10
  logShow $ runFibFFICheatcode dummy
