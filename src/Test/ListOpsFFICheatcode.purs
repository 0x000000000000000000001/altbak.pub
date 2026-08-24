module Test.ListOpsFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runListOpsFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "List Processing FFICheatcode (900 elements):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 900
  logShow $ runListOpsFFICheatcode dummy
