module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.AstTree as AstTree

main :: Effect Unit
main = void $ runBench AstTree.describe AstTree.act
