#!/usr/bin/env bash
cd output

# Hack Data.Ord
sed -i '' 's/var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)/if (a1_1.IntVal < a2_2.IntVal) { return true } else { return false }/g' Data.Ord/Data_Ord.go
sed -i '' 's/var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1_loop, a2_2_loop)/if (a1_1_loop.IntVal < a2_2_loop.IntVal) { return true } else { return false }/g' Data.Ord/Data_Ord.go

# Hack Data.Semiring (add)
sed -i '' 's/return gopurs_runtime.Apply2(dictSemiring_0.V1, a1_1_loop, a2_2_loop)/return gopurs_runtime.Int(a1_1_loop.IntVal + a2_2_loop.IntVal)/g' Data.Semiring/Data_Semiring.go
# (mul)
sed -i '' 's/return gopurs_runtime.Apply2(dictSemiring_0.V2, a1_1_loop, a2_2_loop)/return gopurs_runtime.Int(a1_1_loop.IntVal \* a2_2_loop.IntVal)/g' Data.Semiring/Data_Semiring.go

# Hack Data.Ring (sub)
sed -i '' 's/return gopurs_runtime.Apply2(dictRing_0.V1, a1_1_loop, a2_2_loop)/return gopurs_runtime.Int(a1_1_loop.IntVal - a2_2_loop.IntVal)/g' Data.Ring/Data_Ring.go

# Hack Data.Eq (eq)
sed -i '' 's/return gopurs_runtime.Apply2(dictEq_0.V1, a1_1_loop, a2_2_loop)/return gopurs_runtime.Bool(a1_1_loop.IntVal == a2_2_loop.IntVal)/g' Data.Eq/Data_Eq.go

go build -o go_app ./App/main
GOGC=1000 ./go_app
