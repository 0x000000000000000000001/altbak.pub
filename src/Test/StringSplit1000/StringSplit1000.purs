module Test.StringSplit1000 (prepare) where

import Prelude

import Bench as Bench
import Bench.StringFixtures (checksum)
import Data.Array as Array
import Data.String.Common as String
import Data.String.Pattern (Pattern(..))
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  count <- Bench.opaque 1000
  fixture <- Bench.opaque (String.joinWith "," (Array.replicate count "abc"))
  pure do
    input <- Bench.opaque fixture
    separator <- Bench.opaque (Pattern ",")
    let parts = String.split separator input
    pure (Array.foldl (\total part -> total + checksum part) (Array.length parts) parts)
