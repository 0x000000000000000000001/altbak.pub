module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.RBTree as Case
main :: Effect Unit
main = void $ runBench Case.describe Case.act
