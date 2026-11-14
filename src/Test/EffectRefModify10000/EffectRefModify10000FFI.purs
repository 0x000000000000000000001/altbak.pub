module Test.EffectRefModify10000FFI (prepare) where

import Prelude

import Bench as Bench
import Effect (Effect)

foreign import makeAction :: Int -> Effect (Int -> Int)

-- Immutable fixtures are prepared once; the native action recomputes each call.
prepare :: Effect (Effect Int)
prepare = do
  size <- Bench.opaque 10000
  action <- makeAction size
  pure do
    count <- Bench.opaque size
    pure (action count)
