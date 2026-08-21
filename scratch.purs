module Scratch where
import Prelude
import Data.String as String
import Data.String.Pattern (Pattern(..))
test = String.contains (Pattern ".List") "Test.ListOps.List"
