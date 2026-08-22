module Test.AstTreeFFI where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

foreign import runAstTreeFFI :: Int -> Int

describe :: Effect Unit
describe = log "AST Evaluation FFI:"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 3
  logShow $ runAstTreeFFI dummy
