module Test.TCOFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runTCOFFI :: Int -> Int

describe :: Effect Unit
describe = log "Tail Call Optimization FFI (100k calls):"

act :: Effect String
act = do
  dummy <- Bench.opaque 100000
  pure (show ( runTCOFFI dummy))
