module Test.ArrayFoldlSum10000 (prepare) where

import Prelude

import Bench as Bench
import Bench.ArrayFixtures (prepareRange)
import Data.Array as Array
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  input <- prepareRange 10000
  pure do
    values <- Bench.opaque input
    pure (Array.foldl (+) 0 values)
