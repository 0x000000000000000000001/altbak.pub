module Test.ChurchFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runChurchFFI :: Int -> Int

describe :: Effect Unit
describe = log "Church Numerals FFI (100k Closure Applications):"

act :: Effect String
act = do
  dummy <- Bench.opaque 10
  pure (show ( runChurchFFI dummy))
