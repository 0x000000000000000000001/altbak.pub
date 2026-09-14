module AppCppFFICheatcode where

import Prelude
import Effect (Effect)
import AppFFICheatcode as Entry
import Test.AstTreeFFICheatcode as AstTree
import Test.FibFFICheatcode as Fib
import Test.ListOpsFFICheatcode as ListOps
import Test.TCOFFICheatcode as TCO
import Test.RecordsFFICheatcode as Records
import Test.AckermannFFICheatcode as Ackermann
import Test.ChurchFFICheatcode as Church
import Test.PrimesFFICheatcode as Primes
import Test.RBTreeFFICheatcode as RBTree
import Test.PolymorphismFFICheatcode as Polymorphism
import Test.StateMonadFFICheatcode as StateMonad
import Test.LazyEvaluationFFICheatcode as LazyEvaluation
import Test.ArrayOpsFFICheatcode as ArrayOps
import Test.RowToListFFICheatcode as RowToList

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
