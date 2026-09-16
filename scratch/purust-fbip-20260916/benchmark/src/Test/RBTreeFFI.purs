module Test.RBTreeFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runRBTreeFFI :: Int -> Int

describe :: Effect Unit
describe = log "Red-Black Tree FFI (100k Worst-Case Insertions):"

act :: Effect String
act = do
  dummy <- Bench.opaque 100000
  pure (show ( runRBTreeFFI dummy))
