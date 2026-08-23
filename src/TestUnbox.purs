module TestUnbox where

import Prelude
import Effect (Effect)
import Effect.Console (logShow)

applyInt :: (Int -> Int) -> Int -> Int
applyInt f x = f x

main :: Effect Unit
main = logShow $ applyInt (\x -> x + 1) 42
