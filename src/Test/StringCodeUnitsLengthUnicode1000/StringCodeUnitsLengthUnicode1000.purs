module Test.StringCodeUnitsLengthUnicode1000 (prepare) where

import Prelude

import Bench as Bench
import Bench.StringFixtures (unicodeInput)
import Data.String.CodeUnits as CU
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  fixture <- unicodeInput
  pure do
    input <- Bench.opaque fixture
    pure (CU.length input)
