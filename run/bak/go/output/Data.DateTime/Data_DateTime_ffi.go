package Data_DateTime

import "gopurs/output/gopurs_runtime"



import (
	"time"
)

func createUTC(y, mo, d, h, m, s, ms int) time.Time {
	return time.Date(y, time.Month(mo+1), d, h, m, s, ms*1000000, time.UTC)
}

func getInt(m map[string]interface{}, key string) int {
	v := m[key]
	if val, ok := v.(int64); ok {
		return int(val)
	}
	if val, ok := v.(int); ok {
		return val
	}
	if val, ok := v.(float64); ok {
		return int(val)
	}
	// Fallback for gopurs_runtime.Value
	return int(v.(gopurs_runtime.Value).IntVal)
}

func CalcDiff(rec1 map[string]interface{}, rec2 map[string]interface{}) float64 {
	msUTC1 := createUTC(getInt(rec1, "year"), getInt(rec1, "month")-1, getInt(rec1, "day"), getInt(rec1, "hour"), getInt(rec1, "minute"), getInt(rec1, "second"), getInt(rec1, "millisecond")).UnixMilli()
	msUTC2 := createUTC(getInt(rec2, "year"), getInt(rec2, "month")-1, getInt(rec2, "day"), getInt(rec2, "hour"), getInt(rec2, "minute"), getInt(rec2, "second"), getInt(rec2, "millisecond")).UnixMilli()
	return float64(msUTC1 - msUTC2)
}

func AdjustImpl(just func(interface{}) interface{}, nothing interface{}, offset float64, rec map[string]interface{}) interface{} {
	t := createUTC(getInt(rec, "year"), getInt(rec, "month")-1, getInt(rec, "day"), getInt(rec, "hour"), getInt(rec, "minute"), getInt(rec, "second"), getInt(rec, "millisecond"))
	ms := t.UnixMilli() + int64(offset)
	dt := time.UnixMilli(ms).UTC()
	
	resMap := make(map[string]interface{})
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
var _Gopurs_AdjustImpl = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), Number, Any] (ADT ["Data","Maybe","Maybe"] [Any]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[float64](arg2)
	go_arg3 := gopurs_runtime.UnboxObject(arg3)
	go_res := AdjustImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_CalcDiff = // TAST: (ADT ["Data","Function","Uncurried","Fn2"] [Any, Any, Number])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.UnboxObject(arg0)
	go_arg1 := gopurs_runtime.UnboxObject(arg1)
	go_res := CalcDiff(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})