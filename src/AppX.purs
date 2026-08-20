module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.ArrayOps as ArrayOps

main :: Effect Unit
main = void $ runBench ArrayOps.describe ArrayOps.act
