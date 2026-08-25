module Test.Visible where
import Prelude
myId :: forall @a. a -> a
myId x = x
foo :: Int -> Int
foo = myId @Int
