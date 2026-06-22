module App where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Effect.Now (now)

fib :: Int -> Int
fib 0 = 0

fib 1 = 1

fib n = fib (n - 1) + fib (n - 2)

main :: Effect Unit
main = do
  time <- now
  log $ "Now: " <> show time
  logShow $ fib 40
