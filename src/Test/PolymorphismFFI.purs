module Test.PolymorphismFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runPolymorphismFFI :: Int -> Int

describe :: Effect Unit
describe = log "Polymorphism FFI (10M Type Class Dict Lookups):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 10000000
  logShow $ runPolymorphismFFI dummy
