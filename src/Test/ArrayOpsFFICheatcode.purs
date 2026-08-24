module Test.ArrayOpsFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runArrayOpsFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Array Processing FFICheatcode (900 elements):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 900
  logShow $ runArrayOpsFFICheatcode dummy
