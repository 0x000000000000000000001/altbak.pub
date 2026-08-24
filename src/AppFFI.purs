module AppFFI where

import Prelude
import Effect (Effect)
import Bench (runBench)

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

main :: Effect Unit
main = do
  void $ runBench AstTreeFFI.describe AstTreeFFI.act
  void $ runBench FibFFI.describe FibFFI.act
  void $ runBench ListOpsFFI.describe ListOpsFFI.act
  void $ runBench TCOFFI.describe TCOFFI.act
  void $ runBench RecordsFFI.describe RecordsFFI.act
  void $ runBench AckermannFFI.describe AckermannFFI.act
  void $ runBench ChurchFFI.describe ChurchFFI.act
  void $ runBench PrimesFFI.describe PrimesFFI.act
  void $ runBench RBTreeFFI.describe RBTreeFFI.act
  void $ runBench PolymorphismFFI.describe PolymorphismFFI.act
  void $ runBench StateMonadFFI.describe StateMonadFFI.act
  void $ runBench LazyEvaluationFFI.describe LazyEvaluationFFI.act
  void $ runBench ArrayOpsFFI.describe ArrayOpsFFI.act
  void $ runBench RowToListFFI.describe RowToListFFI.act

  void $ runBench AstTreeFFICheatcode.describe AstTreeFFICheatcode.act
  void $ runBench FibFFICheatcode.describe FibFFICheatcode.act
  void $ runBench ListOpsFFICheatcode.describe ListOpsFFICheatcode.act
  void $ runBench TCOFFICheatcode.describe TCOFFICheatcode.act
  void $ runBench RecordsFFICheatcode.describe RecordsFFICheatcode.act
  void $ runBench AckermannFFICheatcode.describe AckermannFFICheatcode.act
  void $ runBench ChurchFFICheatcode.describe ChurchFFICheatcode.act
  void $ runBench PrimesFFICheatcode.describe PrimesFFICheatcode.act
  void $ runBench RBTreeFFICheatcode.describe RBTreeFFICheatcode.act
  void $ runBench PolymorphismFFICheatcode.describe PolymorphismFFICheatcode.act
  void $ runBench StateMonadFFICheatcode.describe StateMonadFFICheatcode.act
  void $ runBench LazyEvaluationFFICheatcode.describe LazyEvaluationFFICheatcode.act
  void $ runBench ArrayOpsFFICheatcode.describe ArrayOpsFFICheatcode.act
  void $ runBench RowToListFFICheatcode.describe RowToListFFICheatcode.act
