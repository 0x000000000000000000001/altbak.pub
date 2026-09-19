module BenchFable where

import Prelude
import Effect (Effect)
import Test.AstTree as AstTree
import Test.Fib as Fib
import Test.ListOps as ListOps
import Test.TCO as TCO
import Test.Ackermann as Ackermann
import Test.Church as Church
import Test.Primes as Primes
import Test.RBTree as RBTree
import Test.Polymorphism as Polymorphism
import Test.StateMonad as StateMonad
import Test.LazyEvaluation as LazyEvaluation

runAstTree :: Int -> Int
runAstTree n = AstTree.eval (AstTree.buildTree n)

runFib :: Int -> Int
runFib n = Fib.fib n

runListOps :: Int -> Int
runListOps n = ListOps.sumEvens n

runTCO :: Int -> Int
runTCO n = TCO.deepTailRec n 0

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

main :: Effect Unit
main = pure unit
