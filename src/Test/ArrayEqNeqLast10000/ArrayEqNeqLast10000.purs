module Test.ArrayEqNeqLast10000 (prepare) where

import Prelude

import Bench as Bench
import Data.Array as Array
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  count <- Bench.opaque 10000
  left <- Bench.opaque (Array.range 1 count)
  -- Separate arrays with equal prefixes; the final element differs.
  right <- Bench.opaque (Array.snoc (Array.range 1 (count - 1)) (count + 1))
  pure do
    xs <- Bench.opaque left
    ys <- Bench.opaque right
    pure (if xs == ys then 1 else 0)
