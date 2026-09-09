module Test.AckermannFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runAckermannFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Ackermann FFICheatcode (3, 4):"

act :: Effect String
act = do
  m <- Bench.opaque 3
  pure (show (runAckermannFFICheatcode m))
