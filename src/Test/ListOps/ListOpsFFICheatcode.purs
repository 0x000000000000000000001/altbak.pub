module Test.ListOpsFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runListOpsFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "List Processing FFICheatcode (900 elements):"

act :: Effect Int
act = do
  dummy <- Bench.opaque 900
  pure ( runListOpsFFICheatcode dummy)
