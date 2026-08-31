module Test.TCOADT where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench as Bench

data MyList a = Nil | Cons a (MyList a)

foldlMy :: forall a b. (b -> a -> b) -> b -> MyList a -> b
foldlMy _ acc Nil = acc
foldlMy f acc (Cons x xs) = foldlMy f (f acc x) xs

mkList :: Int -> MyList Int -> MyList Int
mkList 0 acc = acc
mkList n acc = mkList (n - 1) (Cons n acc)

describe :: Effect Unit
describe = log "Tail Call Optimization on ADT (100k calls):"

act :: Effect String
act = do
  dummy <- Bench.opaque 100000
  let lst = mkList dummy Nil
  pure (show (foldlMy (+) 0 lst))
