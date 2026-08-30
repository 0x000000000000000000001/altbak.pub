module AppTest where
import Prelude
import Effect (Effect)
import Bench (opaque)
import Test.ChurchFFICheatcode as Test.ChurchFFICheatcode

main :: Effect Unit
main = void $ opaque keepAlive
  where
  keepAlive = {
    desc: Test.ChurchFFICheatcode.describe,
    act: Test.ChurchFFICheatcode.act,
    _dummy: unit
  }
