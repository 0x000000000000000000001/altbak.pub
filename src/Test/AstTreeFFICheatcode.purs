module Test.AstTreeFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runAstTreeFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "AST Evaluation FFICheatcode:"

act :: Effect String
act = do
  dummy <- Bench.opaque 3
  pure (show ( runAstTreeFFICheatcode dummy))
