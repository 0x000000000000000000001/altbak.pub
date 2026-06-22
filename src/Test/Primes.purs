module Test.Primes where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

-- Implementation of the Sieve of Eratosthenes using strict Lists.
-- This benchmark heavily tests:
-- 1. Non-tail recursion (sieve depth is equal to the number of primes found).
-- 2. Massive Garbage Collection: 'filter' creates entirely new lists thousands of times.
-- 3. Closure invocation: The lambda passed to 'filter' captures the prime 'p'.

data List a = Nil | Cons a (List a)

range :: Int -> Int -> List Int
range start end = go end Nil
  where
  go curr acc
    | curr < start = acc
    | otherwise = go (curr - 1) (Cons curr acc)

filter :: forall a. (a -> Boolean) -> List a -> List a
filter p lst = go lst Nil
  where
  go Nil acc = reverse acc
  go (Cons x xs) acc =
    if p x
    then go xs (Cons x acc)
    else go xs acc

reverse :: forall a. List a -> List a
reverse lst = go lst Nil
  where
  go Nil acc = acc
  go (Cons x xs) acc = go xs (Cons x acc)

-- Strict, non-tail-recursive Sieve
sieve :: List Int -> List Int
sieve Nil = Nil
sieve (Cons p xs) = Cons p (sieve (filter (\x -> x `mod` p /= 0) xs))

sumList :: List Int -> Int
sumList lst = go lst 0
  where
  go Nil acc = acc
  go (Cons x xs) acc = go xs (acc + x)

-- Sum of all primes up to 20,000
-- There are 2262 primes, so sieve recurses 2262 times (safe for JS stack).
-- The total iterations across all filters will be around ~22 million.
describe :: Effect Unit
describe = log "Prime Sieve (sum primes up to 20k):"

act :: Effect Unit
act = logShow $ sumList (sieve (range 2 20000))
