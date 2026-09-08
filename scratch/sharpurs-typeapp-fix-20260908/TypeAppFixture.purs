module TypeAppFixture where

identity :: forall a. a -> a
identity value = value

useInt :: Int -> Int
useInt value = identity value

useString :: String -> String
useString value = identity value

generic :: forall b. b -> b
generic value = identity value
