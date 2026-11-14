module Test.TraverseMaybeFailLast1000 (prepare) where

import Prelude

import Bench as Bench
import Bench.ArrayFixtures (prepareRange, weightedChecksum)
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Effect (Effect)

resultChecksum :: Maybe (Array Int) -> Int
resultChecksum Nothing = -1
resultChecksum (Just values) = weightedChecksum values

-- Failure position does not promise that the Array Traversable instance
-- stops evaluating callbacks at that position.
prepare :: Effect (Effect Int)
prepare = do
  input <- prepareRange 1000
  pure do
    values <- Bench.opaque input
    rejected <- Bench.opaque 1000
    pure (resultChecksum (traverse (\value -> if value == rejected then Nothing else Just (value + 1)) values))
