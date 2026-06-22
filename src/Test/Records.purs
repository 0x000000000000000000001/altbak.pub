module Test.Records where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

type DeepRecord = 
  { a :: Int
  , b :: { c :: Int, d :: { e :: Int, f :: Int } }
  }

initial :: DeepRecord
initial = { a: 0, b: { c: 0, d: { e: 0, f: 0 } } }

-- On simule 1 million de mises à jour immuables sur un Record imbriqué
-- PureScript en JS utilise Object.assign, Erlang utilise des Maps natives
updateRec :: Int -> DeepRecord -> DeepRecord
updateRec 0 r = r
updateRec n r = updateRec (n - 1) 
  r { a = r.a + 1
    , b = r.b { c = r.b.c + 2
              , d = r.b.d { e = r.b.d.e + 3
                          , f = r.b.d.f + (n `mod` 5) 
                          }
              }
    }

describe :: Effect Unit
describe = log "Deep Record Updates (1M iterations):"

act :: Effect Unit
act = logShow (updateRec 1000000 initial).b.d.f
