module Test.AstTreeFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runAstTreeFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "AST Evaluation FFICheatcode:"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 3
  logShow $ runAstTreeFFICheatcode dummy
