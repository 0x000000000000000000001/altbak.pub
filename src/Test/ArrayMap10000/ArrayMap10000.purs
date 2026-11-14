module Test.ArrayMap10000 (prepare) where

import Prelude

import Bench as Bench
import Bench.ArrayFixtures (prepareRange, weightedChecksum)
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  input <- prepareRange 10000
  pure do
    values <- Bench.opaque input
    pure (weightedChecksum (map (_ + 1) values))
