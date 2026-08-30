module TestTypeApp2 where

import Prelude

myId :: forall @a. a -> a
myId x = x

test1 :: Int
test1 = myId 1

test2 :: Int
test2 = myId @Int 1

data Step a = Step a

test3 :: forall @a. Step a -> Step a
test3 (Step x) = Step (myId x)

test4 :: forall @a. Step a -> Step a
test4 (Step x) = Step (myId @a x)
