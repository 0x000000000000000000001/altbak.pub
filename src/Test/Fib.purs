module Test.Fib where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib n = fib (n - 1) + fib (n - 2)

describe :: Effect Unit
describe = log "Fibonacci:"

act :: Effect Unit
act = logShow $ fib 40
