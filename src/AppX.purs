module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.ListOps as ListOps

main :: Effect Unit
main = void $ runBench ListOps.describe ListOps.act
