module Test.ArrayOpsFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runArrayOpsFFI :: Int -> Int

describe :: Effect Unit
describe = log "Array Processing FFI (900 elements):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 900
  logShow $ runArrayOpsFFI dummy
