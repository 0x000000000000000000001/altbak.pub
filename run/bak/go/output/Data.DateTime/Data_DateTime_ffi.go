package Data_DateTime



import (
	"math"
	"time"
)

func createUTC(y, mo, d, h, m, s, ms int) time.Time {
	return time.Date(y, time.Month(mo+1), d, h, m, s, ms*1000000, time.UTC)
}

func getInt(m map[string]interface{}, key string) int {
	return int(m[key].(int64))
}

func CalcDiff(rec1 map[string]interface{}, rec2 map[string]interface{}) float64 {
	msUTC1 := createUTC(getInt(rec1, "year"), getInt(rec1, "month")-1, getInt(rec1, "day"), getInt(rec1, "hour"), getInt(rec1, "minute"), getInt(rec1, "second"), getInt(rec1, "millisecond")).UnixMilli()
	msUTC2 := createUTC(getInt(rec2, "year"), getInt(rec2, "month")-1, getInt(rec2, "day"), getInt(rec2, "hour"), getInt(rec2, "minute"), getInt(rec2, "second"), getInt(rec2, "millisecond")).UnixMilli()
	return float64(msUTC1 - msUTC2)
}

func AdjustImpl(just func(interface{}) interface{}, nothing interface{}, offset float64, rec map[string]interface{}) interface{} {
	t := createUTC(getInt(rec, "year"), getInt(rec, "month")-1, getInt(rec, "day"), getInt(rec, "hour"), getInt(rec, "minute"), getInt(rec, "second"), getInt(rec, "millisecond"))
	dt := t.Add(time.Duration(offset) * time.Millisecond)
	
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
