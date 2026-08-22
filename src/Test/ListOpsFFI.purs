module Test.ListOpsFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runListOpsFFI :: Int -> Int

describe :: Effect Unit
describe = log "List Processing FFI (900 elements):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 900
  logShow $ runListOpsFFI dummy
