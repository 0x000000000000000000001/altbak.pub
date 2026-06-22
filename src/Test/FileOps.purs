module Test.FileOps where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

foreign import writeFileSync :: String -> String -> Effect Unit
foreign import readFileSync :: String -> Effect String
foreign import loopE :: Int -> Effect Unit -> Effect Unit

loopIO :: Int -> Effect Unit
loopIO n = loopE n do
  writeFileSync "var/iotest.txt" "Hello IO Benchmarks!"
  _ <- readFileSync "var/iotest.txt"
  pure unit

describe :: Effect Unit
describe = log "File I/O (10k writes/reads):"

act :: Effect Unit
act = do
  loopIO 10000
