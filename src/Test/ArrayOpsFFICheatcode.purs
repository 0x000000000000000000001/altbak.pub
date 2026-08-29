module Test.ArrayOpsFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runArrayOpsFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Array Processing FFICheatcode (900 elements):"

act :: Effect String
act = do
  dummy <- Bench.opaque 900
  pure (show ( runArrayOpsFFICheatcode dummy))
