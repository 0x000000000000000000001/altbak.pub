module Test.TCOFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runTCOFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Tail Call Optimization FFICheatcode (100k calls):"

act :: Effect Int
act = do
  dummy <- Bench.opaque 100000
  pure ( runTCOFFICheatcode dummy)
