module Test.StringTakeCodePointsUnicode1500 (prepare) where

import Prelude

import Bench as Bench
import Bench.StringFixtures (checksum, unicodeInput)
import Data.String.CodePoints as CP
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  fixture <- unicodeInput
  pure do
    input <- Bench.opaque fixture
    count <- Bench.opaque 1500
    pure (checksum (CP.take count input))
