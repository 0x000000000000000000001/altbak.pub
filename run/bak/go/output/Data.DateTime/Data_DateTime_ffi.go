package Data_DateTime

import (
	"gopurs/output/gopurs_runtime"
	"math"
	"time"
)

func createUTC(y, mo, d, h, m, s, ms int) time.Time {
	return time.Date(y, time.Month(mo+1), d, h, m, s, ms*1000000, time.UTC)
}

func getInt(v gopurs_runtime.Value, key string) int {
	m := v.PtrVal.(map[string]gopurs_runtime.Value)
	return int(m[key].IntVal)
}

var CalcDiff = gopurs_runtime.Func(func(rec1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(rec2 gopurs_runtime.Value) gopurs_runtime.Value {
		msUTC1 := createUTC(getInt(rec1, "year"), getInt(rec1, "month")-1, getInt(rec1, "day"), getInt(rec1, "hour"), getInt(rec1, "minute"), getInt(rec1, "second"), getInt(rec1, "millisecond")).UnixMilli()
		msUTC2 := createUTC(getInt(rec2, "year"), getInt(rec2, "month")-1, getInt(rec2, "day"), getInt(rec2, "hour"), getInt(rec2, "minute"), getInt(rec2, "second"), getInt(rec2, "millisecond")).UnixMilli()
		return gopurs_runtime.Float(float64(msUTC1 - msUTC2))
	})
})

var AdjustImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(offset gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(rec gopurs_runtime.Value) gopurs_runtime.Value {
				t := createUTC(getInt(rec, "year"), getInt(rec, "month")-1, getInt(rec, "day"), getInt(rec, "hour"), getInt(rec, "minute"), getInt(rec, "second"), getInt(rec, "millisecond"))
				off := math.Float64frombits(uint64(offset.IntVal))
				dt := t.Add(time.Duration(off) * time.Millisecond)
				
				resMap := make(map[string]gopurs_runtime.Value)
				resMap["year"] = gopurs_runtime.Int(int64(dt.Year()))
				resMap["month"] = gopurs_runtime.Int(int64(dt.Month()))
				resMap["day"] = gopurs_runtime.Int(int64(dt.Day()))
				resMap["hour"] = gopurs_runtime.Int(int64(dt.Hour()))
				resMap["minute"] = gopurs_runtime.Int(int64(dt.Minute()))
				resMap["second"] = gopurs_runtime.Int(int64(dt.Second()))
				resMap["millisecond"] = gopurs_runtime.Int(int64(dt.Nanosecond() / 1000000))
				
				return gopurs_runtime.Apply(just, gopurs_runtime.Record(resMap))
			})
		})
	})
})
