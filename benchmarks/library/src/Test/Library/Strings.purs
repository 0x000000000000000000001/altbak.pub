module Test.Library.Strings where

import Prelude

import Bench as Bench
import Data.Array as Array
import Data.Enum (fromEnum)
import Data.Maybe (fromMaybe)
import Data.String.CodePoints as CP
import Data.String.CodeUnits as CU
import Data.String.Common as String
import Data.String.Pattern (Pattern(..))
import Effect (Effect)

-- Consume all code points of constructed results. This traversal is included
-- in the measured action; these cases are not timings of allocation alone.
checksum :: String -> Int
checksum = Array.foldl (\total cp -> total + fromEnum cp) 0 <<< CP.toCodePointArray

unicodeInput :: Effect String
unicodeInput = do
  count <- Bench.opaque 1000
  Bench.opaque (String.joinWith "" (Array.replicate count "a😀é"))

prepareCodeUnitsLength :: Effect (Effect Int)
prepareCodeUnitsLength = do
  fixture <- unicodeInput
  pure do
    input <- Bench.opaque fixture
    pure (CU.length input)

prepareCodePointsLength :: Effect (Effect Int)
prepareCodePointsLength = do
  fixture <- unicodeInput
  pure do
    input <- Bench.opaque fixture
    pure (CP.length input)

searchInput :: Effect String
searchInput = do
  count <- Bench.opaque 10000
  Bench.opaque (String.joinWith "" (Array.replicate count "a") <> "needle")

prepareIndexHit :: Effect (Effect Int)
prepareIndexHit = do
  fixture <- searchInput
  pure do
    input <- Bench.opaque fixture
    pattern <- Bench.opaque (Pattern "needle")
    pure (fromMaybe (-1) (CU.indexOf pattern input))

prepareIndexMiss :: Effect (Effect Int)
prepareIndexMiss = do
  fixture <- searchInput
  pure do
    input <- Bench.opaque fixture
    pattern <- Bench.opaque (Pattern "absent")
    pure (fromMaybe (-1) (CU.indexOf pattern input))

prepareTakeCodePoints :: Effect (Effect Int)
prepareTakeCodePoints = do
  fixture <- unicodeInput
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1500
    pure (checksum (CP.take count input))

prepareSplit :: Effect (Effect Int)
prepareSplit = do
  count <- Bench.opaque 1000
  fixture <- Bench.opaque (String.joinWith "," (Array.replicate count "abc"))
  pure do
    input <- Bench.opaque fixture
    separator <- Bench.opaque (Pattern ",")
    let parts = String.split separator input
    pure (Array.foldl (\total part -> total + checksum part) (Array.length parts) parts)

prepareConcat :: Effect (Effect Int)
prepareConcat = do
  count <- Bench.opaque 1000
  fixture <- Bench.opaque (Array.replicate count "ab")
  pure do
    input <- Bench.opaque fixture
    pure (checksum (Array.foldl (<>) "" input))
