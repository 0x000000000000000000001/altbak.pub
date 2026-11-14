module Test.StringConcat1000 (prepare) where

import Prelude

import Bench as Bench
import Bench.StringFixtures (checksum)
import Data.Array as Array
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  count <- Bench.opaque 1000
  fixture <- Bench.opaque (Array.replicate count "ab")
  pure do
    input <- Bench.opaque fixture
    pure (checksum (Array.foldl (<>) "" input))
