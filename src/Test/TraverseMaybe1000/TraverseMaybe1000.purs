module Test.TraverseMaybe1000 (prepare) where

import Prelude

import Bench as Bench
import Bench.ArrayFixtures (prepareRange, weightedChecksum)
import Data.Maybe (Maybe(..))
import Data.Traversable (traverse)
import Effect (Effect)

resultChecksum :: Maybe (Array Int) -> Int
resultChecksum Nothing = -1
resultChecksum (Just values) = weightedChecksum values

prepare :: Effect (Effect Int)
prepare = do
  input <- prepareRange 1000
  pure do
    values <- Bench.opaque input
    pure (resultChecksum (traverse (\value -> Just (value + 1)) values))
