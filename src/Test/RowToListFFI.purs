module Test.RowToListFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runRowToListFFI :: Int -> Int

describe :: Effect Unit
describe = log "RowToList FFI (Keys Count):"

act :: Effect Int
act = do
  dummy <- Bench.opaque 0
  pure ( runRowToListFFI dummy)
