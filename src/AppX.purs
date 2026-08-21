module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.Primes as Primes

main :: Effect Unit
main = void $ runBench Primes.describe Primes.act
