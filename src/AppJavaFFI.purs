module AppJavaFFI where

import Prelude
import Effect (Effect)
import AppFFI as AppFFI
import Test.AstTreeFFI as AstTreeFFI
import Test.FibFFI as FibFFI
import Test.ListOpsFFI as ListOpsFFI
import Test.TCOFFI as TCOFFI
import Test.RecordsFFI as RecordsFFI
import Test.AckermannFFI as AckermannFFI
import Test.ChurchFFI as ChurchFFI
import Test.PrimesFFI as PrimesFFI
import Test.RBTreeFFI as RBTreeFFI
import Test.PolymorphismFFI as PolymorphismFFI
import Test.StateMonadFFI as StateMonadFFI
import Test.LazyEvaluationFFI as LazyEvaluationFFI
import Test.ArrayOpsFFI as ArrayOpsFFI
import Test.RowToListFFI as RowToListFFI
import Test.AstTreeFFICheatcode as AstTreeFFICheatcode
import Test.FibFFICheatcode as FibFFICheatcode
import Test.ListOpsFFICheatcode as ListOpsFFICheatcode
import Test.TCOFFICheatcode as TCOFFICheatcode
import Test.RecordsFFICheatcode as RecordsFFICheatcode
import Test.AckermannFFICheatcode as AckermannFFICheatcode
import Test.ChurchFFICheatcode as ChurchFFICheatcode
import Test.PrimesFFICheatcode as PrimesFFICheatcode
import Test.RBTreeFFICheatcode as RBTreeFFICheatcode
import Test.PolymorphismFFICheatcode as PolymorphismFFICheatcode
import Test.StateMonadFFICheatcode as StateMonadFFICheatcode
import Test.LazyEvaluationFFICheatcode as LazyEvaluationFFICheatcode
import Test.ArrayOpsFFICheatcode as ArrayOpsFFICheatcode
import Test.RowToListFFICheatcode as RowToListFFICheatcode

warmupFP :: Effect Unit
warmupFP = do
  void AstTreeFFI.act
  void FibFFI.act
  void ListOpsFFI.act
  void TCOFFI.act
  void RecordsFFI.act
  void AckermannFFI.act
  void ChurchFFI.act
  void PrimesFFI.act
  void RBTreeFFI.act
  void PolymorphismFFI.act
  void StateMonadFFI.act
  void LazyEvaluationFFI.act
  void ArrayOpsFFI.act
  void RowToListFFI.act

warmupOptimized :: Effect Unit
warmupOptimized = do
  void AstTreeFFICheatcode.act
  void FibFFICheatcode.act
  void ListOpsFFICheatcode.act
  void TCOFFICheatcode.act
  void RecordsFFICheatcode.act
  void AckermannFFICheatcode.act
  void ChurchFFICheatcode.act
  void PrimesFFICheatcode.act
  void RBTreeFFICheatcode.act
  void PolymorphismFFICheatcode.act
  void StateMonadFFICheatcode.act
  void LazyEvaluationFFICheatcode.act
  void ArrayOpsFFICheatcode.act
  void RowToListFFICheatcode.act

-- Match the three complete warm-up passes in App before the per-case warm-ups.
main :: Effect Unit
main = do
  warmupFP
  warmupOptimized
  warmupFP
  warmupOptimized
  warmupFP
  warmupOptimized
  AppFFI.main
