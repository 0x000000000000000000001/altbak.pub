module Test.LazyEvaluationFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runLazyEvaluationFFI :: Int -> Int

describe :: Effect Unit
describe = log "Lazy Evaluation FFI (1M Thunks Forced, 1k Depth):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 1000
  logShow $ runLazyEvaluationFFI dummy
