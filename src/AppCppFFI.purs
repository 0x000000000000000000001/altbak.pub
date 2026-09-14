module AppCppFFI where

import Prelude
import Effect (Effect)
import AppFFI as Entry
import Test.AstTreeFFI as AstTree
import Test.FibFFI as Fib
import Test.ListOpsFFI as ListOps
import Test.TCOFFI as TCO
import Test.RecordsFFI as Records
import Test.AckermannFFI as Ackermann
import Test.ChurchFFI as Church
import Test.PrimesFFI as Primes
import Test.RBTreeFFI as RBTree
import Test.PolymorphismFFI as Polymorphism
import Test.StateMonadFFI as StateMonad
import Test.LazyEvaluationFFI as LazyEvaluation
import Test.ArrayOpsFFI as ArrayOps
import Test.RowToListFFI as RowToList

-- Match App: three complete warm-up passes of the selected suite.
warmup :: Effect Unit
warmup = do
  void AstTree.act
  void Fib.act
  void ListOps.act
  void TCO.act
  void Records.act
  void Ackermann.act
  void Church.act
  void Primes.act
  void RBTree.act
  void Polymorphism.act
  void StateMonad.act
  void LazyEvaluation.act
  void ArrayOps.act
  void RowToList.act

main :: Effect Unit
main = do
  warmup
  warmup
  warmup
  Entry.main
