package Data_DateTime

import "gopurs/output/gopurs_runtime"



import (
	"math"
	"time"
)

func createUTC(y, mo, d, h, m, s, ms int) time.Time {
	return time.Date(y, time.Month(mo+1), d, h, m, s, ms*1000000, time.UTC)
}

func getInt(m map[string]any, key string) int {
	return int(m[key].(int64))
}

func CalcDiff(rec1 map[string]any, rec2 map[string]any) float64 {
	msUTC1 := createUTC(getInt(rec1, "year"), getInt(rec1, "month")-1, getInt(rec1, "day"), getInt(rec1, "hour"), getInt(rec1, "minute"), getInt(rec1, "second"), getInt(rec1, "millisecond")).UnixMilli()
	msUTC2 := createUTC(getInt(rec2, "year"), getInt(rec2, "month")-1, getInt(rec2, "day"), getInt(rec2, "hour"), getInt(rec2, "minute"), getInt(rec2, "second"), getInt(rec2, "millisecond")).UnixMilli()
	return float64(msUTC1 - msUTC2)
}

func AdjustImpl(just func(any) any, nothing any, offset float64, rec map[string]any) any {
	t := createUTC(getInt(rec, "year"), getInt(rec, "month")-1, getInt(rec, "day"), getInt(rec, "hour"), getInt(rec, "minute"), getInt(rec, "second"), getInt(rec, "millisecond"))
	dt := t.Add(time.Duration(offset) * time.Millisecond)
	
	resMap := make(map[string]any)
	resMap["year"] = dt.Year()
	resMap["month"] = int(dt.Month())
	resMap["day"] = dt.Day()
	resMap["hour"] = dt.Hour()
	resMap["minute"] = dt.Minute()
	resMap["second"] = dt.Second()
	resMap["millisecond"] = dt.Nanosecond() / 1000000
	
	return just(resMap)
}


// --- Auto-generated FFI wrappers ---
func Call_calcDiff(arg0 map[string]any, arg1 map[string]any) float64 {
	return CalcDiff(arg0, arg1)
}
var _Gopurs_CalcDiff = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_map := gopurs_runtime.RecordToMap(arg0)
	go_arg0 := make(map[string]any)
	for k, v := range arg0_map { go_arg0[k] = v }
	arg1_map := gopurs_runtime.RecordToMap(arg1)
	go_arg1 := make(map[string]any)
	for k, v := range arg1_map { go_arg1[k] = v }
	go_res := CalcDiff(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_adjustImpl(arg0 func(any) any, arg1 any, arg2 float64, arg3 map[string]any) any {
	return AdjustImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_AdjustImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[float64](arg2)
	arg3_map := gopurs_runtime.RecordToMap(arg3)
	go_arg3 := make(map[string]any)
	for k, v := range arg3_map { go_arg3[k] = v }
	go_res := AdjustImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
