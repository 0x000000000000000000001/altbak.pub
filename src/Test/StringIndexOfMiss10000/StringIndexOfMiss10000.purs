module Test.StringIndexOfMiss10000 (prepare) where

import Prelude

import Bench as Bench
import Bench.StringFixtures (searchInput)
import Data.Maybe (fromMaybe)
import Data.String.CodeUnits as CU
import Data.String.Pattern (Pattern(..))
import Effect (Effect)

prepare :: Effect (Effect Int)
prepare = do
  fixture <- searchInput
  pure do
    input <- Bench.opaque fixture
    pattern <- Bench.opaque (Pattern "absent")
    pure (fromMaybe (-1) (CU.indexOf pattern input))
