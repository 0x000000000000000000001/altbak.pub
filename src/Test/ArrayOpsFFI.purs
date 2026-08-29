module Test.ArrayOpsFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runArrayOpsFFI :: Int -> Int

describe :: Effect Unit
describe = log "Array Processing FFI (900 elements):"

act :: Effect String
act = do
  dummy <- Bench.opaque 900
  pure (show ( runArrayOpsFFI dummy))
