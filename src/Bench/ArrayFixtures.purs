module Bench.ArrayFixtures (prepareRange, weightedChecksum) where

import Prelude

import Bench as Bench
import Data.Array as Array
import Data.FoldableWithIndex (foldlWithIndex)
import Effect (Effect)

-- Alternating weights make order observable without overflowing a 32-bit Int.
-- Consuming every result element is part of the measured action.
weightedChecksum :: Array Int -> Int
weightedChecksum = foldlWithIndex (\index total value -> total + (1 + index `mod` 2) * value) 0

prepareRange :: Int -> Effect (Array Int)
prepareRange size = do
  count <- Bench.opaque size
  Bench.opaque (Array.range 1 count)
