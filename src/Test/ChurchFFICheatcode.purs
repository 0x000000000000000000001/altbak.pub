module Test.ChurchFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runChurchFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Church Numerals FFICheatcode (100k Closure Applications):"

act :: Effect String
act = do
  dummy <- Bench.opaque 10
  pure (show ( runChurchFFICheatcode dummy))
