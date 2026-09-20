module Test.FileOps where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Effect.Ref as Ref
import Bench as Bench

foreign import writeFileSync :: String -> String -> Effect Unit
foreign import readFileSync :: String -> Effect String
foreign import loopE :: Int -> Effect Unit -> Effect Unit

loopIO :: Int -> Effect Int
loopIO n = do
  verified <- Ref.new 0
  loopE n do
    writeFileSync "var/iotest.txt" "Hello IO Benchmarks!"
    content <- readFileSync "var/iotest.txt"
    Ref.modify_ (\count -> if content == "Hello IO Benchmarks!" then count + 1 else count) verified
  Ref.read verified

describe :: Effect Unit
describe = log "File I/O (10k writes/reads):"

act :: Effect String
act = do
  iterations <- Bench.opaque 10000
  verified <- loopIO iterations
  pure (show verified)
