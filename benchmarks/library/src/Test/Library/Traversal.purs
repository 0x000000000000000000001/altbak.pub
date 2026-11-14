module Test.Library.Traversal where

import Prelude

import Bench as Bench
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Effect (Effect)
import Test.Library.Arrays (prepareRange, weightedChecksum)

resultChecksum :: Maybe (Array Int) -> Int
resultChecksum Nothing = -1
resultChecksum (Just values) = weightedChecksum values

prepareSuccess :: Effect (Effect Int)
prepareSuccess = do
  input <- prepareRange 1000
  pure do
    values <- Bench.opaque input
    pure (resultChecksum (traverse (\value -> Just (value + 1)) values))

-- Failure position is a workload distinction, not a promise that the Array
-- Traversable implementation stops evaluating the callback at that position.
prepareFailure :: Int -> Effect (Effect Int)
prepareFailure failure = do
  input <- prepareRange 1000
  pure do
    values <- Bench.opaque input
    rejected <- Bench.opaque failure
    pure (resultChecksum (traverse (\value -> if value == rejected then Nothing else Just (value + 1)) values))

prepareFailFirst :: Effect (Effect Int)
prepareFailFirst = prepareFailure 1

prepareFailLast :: Effect (Effect Int)
prepareFailLast = prepareFailure 1000
