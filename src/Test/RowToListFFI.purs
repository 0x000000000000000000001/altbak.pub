module Test.RowToListFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runRowToListFFI :: Int -> Int

describe :: Effect Unit
describe = log "RowToList FFI (Keys Count):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 0
  logShow $ runRowToListFFI dummy
