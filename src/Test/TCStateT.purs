module Test.TCStateT where
import Prelude
import Control.Monad.State.Trans (StateT, runStateT)
import Data.Identity (Identity)
import Effect (Effect)
import Effect.Console (log)
myState :: StateT Int Identity Unit
myState = do
  _ <- pure unit
  pure unit
describe :: Effect Unit
describe = log "TCStateT"
act :: Effect String
act = pure (show 42)
