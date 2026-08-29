module Test.FibFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runFibFFI :: Int -> Int

describe :: Effect Unit
describe = log "Fibonacci FFI:"

act :: Effect String
act = do
  dummy <- Bench.opaque 10
  pure (show ( runFibFFI dummy))
