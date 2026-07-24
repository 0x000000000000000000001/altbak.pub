module AppX where

import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.TCO as TCO

main :: Effect Unit
main = void $ runBench TCO.describe TCO.act
