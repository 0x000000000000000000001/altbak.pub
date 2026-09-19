module BenchPurust where

import Prelude
import Effect (Effect)
import Test.AstTree as AstTree
import Test.Fib as Fib
import Test.ListOps as ListOps
import Test.TCO as TCO
import Test.Records as Records
import Test.Ackermann as Ackermann
import Test.Church as Church
import Test.Primes as Primes
import Test.RBTree as RBTree
import Test.Polymorphism as Polymorphism
import Test.StateMonad as StateMonad
import Test.LazyEvaluation as LazyEvaluation
import Test.ArrayOps as ArrayOps
import Test.RowToList as RowToList

-- Numeric expressions from each Test module's act. The native harness supplies
-- opaque inputs and consumes the results outside this generated code.
runAstTree :: Int -> Int
runAstTree n = AstTree.eval (AstTree.buildTree n)

runFib :: Int -> Int
runFib n = Fib.fib n

runListOps :: Int -> Int
runListOps n = ListOps.sumEvens n

runTCO :: Int -> Int
runTCO n = TCO.deepTailRec n 0

runRecords :: Int -> Int
runRecords n = (Records.updateRec n Records.initial).b.d.f

runAckermann :: Int -> Int
runAckermann n = Ackermann.ackermann n 4

runChurch :: Int -> Int
runChurch n = Church.toInt (Church.c100k n)

runPrimes :: Int -> Int
runPrimes n = Primes.sumList (Primes.sieve (Primes.range 2 n))

runRBTree :: Int -> Int
runRBTree n = RBTree.depth (RBTree.buildTree n RBTree.E)

runPolymorphism :: Int -> Int
runPolymorphism n = Polymorphism.polyLoop n (0 :: Int)

runStateMonad :: Int -> Int
runStateMonad n = StateMonad.runManyTimes n 0

runLazyEvaluation :: Int -> Int
runLazyEvaluation n = LazyEvaluation.runManyTimes n 0

runArrayOps :: Int -> Int
runArrayOps n = ArrayOps.sumEvens n

runRowToList :: Int -> Int
runRowToList _ = RowToList.keys { a: 1, b: "two", c: true, d: 4.0, e: "five" }

main :: Effect Unit
main = pure unit
