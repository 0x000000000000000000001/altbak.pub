module Test.RecordsFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runRecordsFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Deep Record Updates FFICheatcode (10k iterations):"

act :: Effect String
act = do
  dummy <- Bench.opaque 10000
  pure (show ( runRecordsFFICheatcode dummy))
