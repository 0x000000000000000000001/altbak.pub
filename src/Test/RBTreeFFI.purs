module Test.RBTreeFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runRBTreeFFI :: Int -> Int

describe :: Effect Unit
describe = log "Red-Black Tree FFI (100k Worst-Case Insertions):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 100000
  logShow $ runRBTreeFFI dummy
