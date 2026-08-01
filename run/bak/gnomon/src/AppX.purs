module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.Fib as Fib

main :: Effect Unit
main = void $ runBench Fib.describe Fib.act
