module Test.TC where
import Prelude
import Effect (Effect)
import Effect.Console (log)
class MyClass a where
  myMethod :: a -> String
instance MyClass Int where
  myMethod _ = "Int"
describe :: Effect Unit
describe = log "TC"
act :: Effect String
act = pure (myMethod 42)
