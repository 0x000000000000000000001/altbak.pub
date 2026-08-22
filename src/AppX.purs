module AppX where
import Prelude
import Effect (Effect)
import Bench (runBench)
import Test.Records as Records

main :: Effect Unit
main = void $ runBench Records.describe Records.act
