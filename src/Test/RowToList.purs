module Test.RowToList where

import Prelude
import Prim.RowList (class RowToList, RowList, Cons, Nil)
import Type.Proxy (Proxy(..))
import Effect (Effect)
import Effect.Console (logShow, log)
import Bench as Bench

class RecordKeys (rl :: RowList Type) where
  keysImpl :: Proxy rl -> Int

instance keysNil :: RecordKeys Nil where
  keysImpl _ = 0

instance keysCons :: RecordKeys tail => RecordKeys (Cons sym ty tail) where
  keysImpl _ = 1 + keysImpl (Proxy :: Proxy tail)

-- | @inline always
keys :: forall row rl. RowToList row rl => RecordKeys rl => Record row -> Int
keys _ = keysImpl (Proxy :: Proxy rl)

describe :: Effect Unit
describe = log "RowToList (Keys Count):"

act :: Effect Unit
act = do
  _ <- Bench.opaque 10000
  let rec = { a: 1, b: "two", c: true, d: 4.0, e: "five" }
  logShow (keys rec)
