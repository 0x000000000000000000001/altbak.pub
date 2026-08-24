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
  runBench AstTreeFFI.describe AstTreeFFI.act
  runBench FibFFI.describe FibFFI.act
  runBench ListOpsFFI.describe ListOpsFFI.act
  runBench TCOFFI.describe TCOFFI.act
  runBench RecordsFFI.describe RecordsFFI.act
  runBench AckermannFFI.describe AckermannFFI.act
  runBench ChurchFFI.describe ChurchFFI.act
  runBench PrimesFFI.describe PrimesFFI.act
  runBench RBTreeFFI.describe RBTreeFFI.act
  runBench PolymorphismFFI.describe PolymorphismFFI.act
  runBench StateMonadFFI.describe StateMonadFFI.act
  runBench LazyEvaluationFFI.describe LazyEvaluationFFI.act
  runBench ArrayOpsFFI.describe ArrayOpsFFI.act
  runBench RowToListFFI.describe RowToListFFI.act

  runBench AstTreeFFICheatcode.describe AstTreeFFICheatcode.act
  runBench FibFFICheatcode.describe FibFFICheatcode.act
  runBench ListOpsFFICheatcode.describe ListOpsFFICheatcode.act
  runBench TCOFFICheatcode.describe TCOFFICheatcode.act
  runBench RecordsFFICheatcode.describe RecordsFFICheatcode.act
  runBench AckermannFFICheatcode.describe AckermannFFICheatcode.act
  runBench ChurchFFICheatcode.describe ChurchFFICheatcode.act
  runBench PrimesFFICheatcode.describe PrimesFFICheatcode.act
  runBench RBTreeFFICheatcode.describe RBTreeFFICheatcode.act
  runBench PolymorphismFFICheatcode.describe PolymorphismFFICheatcode.act
  runBench StateMonadFFICheatcode.describe StateMonadFFICheatcode.act
  runBench LazyEvaluationFFICheatcode.describe LazyEvaluationFFICheatcode.act
  runBench ArrayOpsFFICheatcode.describe ArrayOpsFFICheatcode.act
  runBench RowToListFFICheatcode.describe RowToListFFICheatcode.act
