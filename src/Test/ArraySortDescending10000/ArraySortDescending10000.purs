module Test.ArraySortDescending10000 (prepare) where

import Prelude

import Bench as Bench
import Bench.ArrayFixtures (weightedChecksum)
import Data.Array as Array
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  count <- Bench.opaque 10000
  input <- Bench.opaque (Array.range count 1)
  pure do
    values <- Bench.opaque input
    pure (weightedChecksum (Array.sort values))
