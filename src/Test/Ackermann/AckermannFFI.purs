module Test.AckermannFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runAckermannFFI :: Int -> Int

describe :: Effect Unit
describe = log "Ackermann FFI (3, 4):"

act :: Effect Int
act = do
  m <- Bench.opaque 3
  pure (runAckermannFFI m)
