module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.Polymorphism as Polymorphism

main :: Effect Unit
main = void $ runBench Polymorphism.describe Polymorphism.act
