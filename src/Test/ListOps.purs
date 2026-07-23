module Test.ListOps where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

data List a = Nil | Cons a (List a)

range :: Int -> Int -> List Int
range start end = go end Nil
  where
  go curr acc =
    if curr < start then acc
    else go (curr - 1) (Cons curr acc)

filterEvens :: List Int -> List Int
filterEvens lst = go lst Nil
  where
  go Nil acc = acc
  go (Cons x xs) acc =
    if x `mod` 2 == 0 then go xs (Cons x acc)
    else go xs acc

foldl :: ∀ a b. (b -> a -> b) -> b -> List a -> b
foldl _ acc Nil = acc
foldl f acc (Cons x xs) = foldl f (f acc x) xs

sumEvens :: Int -> Int
sumEvens n = foldl (+) 0 (filterEvens (range 1 n))

describe :: Effect Unit
describe = log "List Processing (900 elements):"

act :: Effect Unit
act = do
  dummy <- Bench.opaque 900
  logShow $ sumEvens dummy
