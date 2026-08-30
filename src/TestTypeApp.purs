module TestTypeApp where

import Prelude

myAdd :: forall @a. Semiring a => a -> a -> a
myAdd x y = x + y

-- 1. Implicite + concret
test1 :: Int
test1 = myAdd 1 2

-- 2. Explicite + concret
test2 :: Int
test2 = myAdd @Int 1 2

data Step a = Step a

-- 3. Implicite + parametre
test3 :: forall @a. Semiring a => Step a -> Step a
test3 (Step x) = Step (myAdd x x)

-- 4. Explicite + parametre
test4 :: forall @a. Semiring a => Step a -> Step a
test4 (Step x) = Step (myAdd @a x x)
