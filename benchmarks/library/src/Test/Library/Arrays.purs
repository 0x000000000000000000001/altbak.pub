module Test.Library.Arrays where

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

prepareMap :: Effect (Effect Int)
prepareMap = do
  input <- prepareRange 10000
  pure do
    values <- Bench.opaque input
    pure (weightedChecksum (map (_ + 1) values))

prepareFilter :: Effect (Effect Int)
prepareFilter = do
  input <- prepareRange 10000
  pure do
    values <- Bench.opaque input
    pure (weightedChecksum (Array.filter (\value -> value `mod` 2 == 0) values))

prepareFoldl :: Effect (Effect Int)
prepareFoldl = do
  input <- prepareRange 10000
  pure do
    values <- Bench.opaque input
    pure (Array.foldl (+) 0 values)

prepareSort :: Effect (Effect Int)
prepareSort = do
  count <- Bench.opaque 10000
  input <- Bench.opaque (Array.range count 1)
  pure do
    values <- Bench.opaque input
    pure (weightedChecksum (Array.sort values))

prepareEqNeq :: Effect (Effect Int)
prepareEqNeq = do
  count <- Bench.opaque 10000
  left <- Bench.opaque (Array.range 1 count)
  -- Separate arrays with equal prefixes; the final element differs.
  right <- Bench.opaque (Array.snoc (Array.range 1 (count - 1)) (count + 1))
  pure do
    xs <- Bench.opaque left
    ys <- Bench.opaque right
    pure (if xs == ys then 1 else 0)

prepareAppend :: Effect (Effect Int)
prepareAppend = do
  count <- Bench.opaque 1000
  left <- Bench.opaque (Array.range 1 count)
  right <- Bench.opaque (Array.range (count + 1) (count * 2))
  pure do
    xs <- Bench.opaque left
    ys <- Bench.opaque right
    pure (weightedChecksum (xs <> ys))
