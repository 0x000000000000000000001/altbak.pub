module AppTest where
import Prelude
import Effect (Effect)
import Bench (opaque)
import Test.RBTreeFFI as Test.RBTreeFFI

main :: Effect Unit
main = void $ opaque keepAlive
  where
  keepAlive = {
    desc: Test.RBTreeFFI.describe,
    act: Test.RBTreeFFI.act,
    _dummy: unit
  }
