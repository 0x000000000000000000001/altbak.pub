module Test.PrimesFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runPrimesFFI :: Int -> Int

describe :: Effect Unit
describe = log "Prime Sieve FFI (sum primes up to 500):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 500
  logShow $ runPrimesFFI dummy
