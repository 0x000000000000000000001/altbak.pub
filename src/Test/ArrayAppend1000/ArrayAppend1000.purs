module Test.ArrayAppend1000 (prepare) where

import Prelude

import Bench as Bench
import Bench.ArrayFixtures (weightedChecksum)
import Data.Array as Array
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  count <- Bench.opaque 1000
  left <- Bench.opaque (Array.range 1 count)
  right <- Bench.opaque (Array.range (count + 1) (count * 2))
  pure do
    xs <- Bench.opaque left
    ys <- Bench.opaque right
    pure (weightedChecksum (xs <> ys))
