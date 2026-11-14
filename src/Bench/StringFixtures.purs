module Bench.StringFixtures (checksum, unicodeInput, searchInput) where

import Prelude

import Bench as Bench
import Data.Array as Array
import Data.Enum (fromEnum)
import Data.String.CodePoints as CP
import Data.String.Common as String
import Effect (Effect)

-- Consume all code points of constructed results. This traversal is included
-- in the measured action; these cases are not timings of allocation alone.
checksum :: String -> Int
checksum = Array.foldl (\total cp -> total + fromEnum cp) 0 <<< CP.toCodePointArray

unicodeInput :: Effect String
unicodeInput = do
  count <- Bench.opaque 1000
  Bench.opaque (String.joinWith "" (Array.replicate count "a😀é"))

searchInput :: Effect String
searchInput = do
  count <- Bench.opaque 10000
  Bench.opaque (String.joinWith "" (Array.replicate count "a") <> "needle")
