module Test.AstTreeFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

foreign import runAstTreeFFI :: Int -> Int

describe :: Effect Unit
describe = log "AST Evaluation FFI:"

act :: Effect String
act = do
  dummy <- Bench.opaque 3
  pure (show ( runAstTreeFFI dummy))
