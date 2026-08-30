module AppTest where
import Prelude
import Effect (Effect)
import Bench (opaque)
import Test.AckermannFFI as Test.AckermannFFI

main :: Effect Unit
main = void $ opaque keepAlive
  where
  keepAlive = {
    desc: Test.AckermannFFI.describe,
    act: Test.AckermannFFI.act,
    _dummy: unit
  }
