module AppX where
import Prelude
import Effect (Effect)
import Test.BenchCheck as BenchCheck

main :: Effect Unit
main = void $ BenchCheck.act
