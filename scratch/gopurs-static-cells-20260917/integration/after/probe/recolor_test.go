package main

import (
 "testing"
 ps "gopurs/output/purescript"
)

func TestRecolorCopiesFieldsIntoDistinctDonor(t *testing.T) {
 for _, withDonor := range []bool{false, true} {
  input := &tree{Rc: 1, V0: red(), V1: &tree{Rc: 1, V0: black(), V2: 7}, V2: 42,
    V3: &tree{Rc: 1, V0: black(), V2: 77}}
  var donor *tree
  if withDonor { donor = &tree{Rc: 99, V0: red(), V2: -999} }
  result := ps.Call_Test_RBTree___gopurs_owned_makeBlack_0_consume(input, donor)
  if result != input && result != donor { t.Fatal("result must reuse an available cell") }
  if result.Rc != 1 || result.V0 != black() || result.V2 != 42 ||
      result.V1 == nil || result.V1.V2 != 7 || result.V3 == nil || result.V3.V2 != 77 {
   t.Fatal("recoloration did not preserve every field")
  }
  validate(t, result, []int64{7,42,77})
 }
}

func TestInsertWithDistinctDonorPreservesShape(t *testing.T) {
 for _, key := range []int64{0, 12, 32} {
  input := ps.Call_Test_RBTree_buildTree(31, nil)
  reference := ps.Call_Test_RBTree_buildTree(31, nil)
  reference = ps.Call_Test_RBTree_insert(key, reference)
  donor := &tree{Rc: 99, V0: red(), V2: -999}
  result := ps.Call_Test_RBTree___gopurs_owned_insert_0_consume(key, input, donor)
  if render(result) != render(reference) { t.Fatal("donor insertion changed shape or colors") }
 }
}
