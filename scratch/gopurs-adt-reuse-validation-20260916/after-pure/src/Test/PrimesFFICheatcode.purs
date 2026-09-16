module Test.PrimesFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runPrimesFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Prime Sieve FFICheatcode (sum primes up to 500):"

act :: Effect String
act = do
  dummy <- Bench.opaque 500
  pure (show ( runPrimesFFICheatcode dummy))
