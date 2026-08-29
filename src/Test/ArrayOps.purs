module Test.ArrayOps where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench
import Data.Array as Array

range :: Int -> Int -> Array Int
range start end = Array.range start end

filterEvens :: Array Int -> Array Int
filterEvens arr = Array.filter (\x -> mod x 2 == 0) arr

sumEvens :: Int -> Int
sumEvens n = Array.foldl (+) 0 (filterEvens (range 1 n))

describe :: Effect Unit
describe = log "Array Processing (900 elements):"

act :: Effect String
act = do
  dummy <- Bench.opaque 900
  pure (show ( sumEvens dummy))
