module Test.AckermannFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runAckermannFFI :: Int -> Int

describe :: Effect Unit
describe = log "Ackermann FFI (3, 4):"

act :: Effect String
act = do
  dummy <- Bench.opaque 0
  pure (show ( runAckermannFFI dummy))
