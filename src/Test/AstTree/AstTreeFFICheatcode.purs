module Test.AstTreeFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runAstTreeFFICheatcode :: Int -> Int

describe :: Effect Unit
describe = log "AST Evaluation FFICheatcode:"

act :: Effect Int
act = do
  dummy <- Bench.opaque 3
  pure ( runAstTreeFFICheatcode dummy)
