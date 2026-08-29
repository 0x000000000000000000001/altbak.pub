module Test.LazyEvaluationFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runLazyEvaluationFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "Lazy Evaluation FFICheatcode (1M Thunks Forced, 1k Depth):"

act :: Effect String
act = do
  dummy <- Bench.opaque 1000
  pure (show ( runLazyEvaluationFFICheatcode dummy))
