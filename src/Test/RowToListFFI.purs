module Test.RowToListFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runRowToListFFI :: Int -> Int

describe :: Effect Unit
describe = log "RowToList FFI (Keys Count):"

act :: Effect String
act = do
  dummy <- Bench.opaque 0
  pure (show ( runRowToListFFI dummy))
