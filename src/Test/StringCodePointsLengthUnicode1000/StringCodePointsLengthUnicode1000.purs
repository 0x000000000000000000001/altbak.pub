module Test.StringCodePointsLengthUnicode1000 (prepare) where

import Prelude

import Bench as Bench
import Bench.StringFixtures (unicodeInput)
import Data.String.CodePoints as CP
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  fixture <- unicodeInput
  pure do
    input <- Bench.opaque fixture
    pure (CP.length input)
