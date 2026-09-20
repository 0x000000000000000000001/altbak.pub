module Test.RowToListFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runRowToListFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "RowToList FFICheatcode (Keys Count):"

act :: Effect Int
act = do
  dummy <- Bench.opaque 0
  pure ( runRowToListFFICheatcode dummy)
