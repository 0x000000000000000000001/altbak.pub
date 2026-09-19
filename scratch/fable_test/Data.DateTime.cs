using System;
using System.Collections.Generic;

namespace Data.DateTime;

public static class FFI {
    private static global::System.DateTime CreateUTC(int y, int mo, int d, int h, int m, int s, int ms) {
        return new global::System.DateTime(y, mo, d, h, m, s, ms, global::System.DateTimeKind.Utc);
    }

    private static int GetInt(dynamic m, string key) {
        var dict = (IDictionary<string, object>)m;
        return (int)dict[key];
    }
    
    public static double CalcDiff(dynamic rec1, dynamic rec2) {
        var msUTC1 = CreateUTC(GetInt(rec1, "year"), GetInt(rec1, "month"), GetInt(rec1, "day"), GetInt(rec1, "hour"), GetInt(rec1, "minute"), GetInt(rec1, "second"), GetInt(rec1, "millisecond"));
        var msUTC2 = CreateUTC(GetInt(rec2, "year"), GetInt(rec2, "month"), GetInt(rec2, "day"), GetInt(rec2, "hour"), GetInt(rec2, "minute"), GetInt(rec2, "second"), GetInt(rec2, "millisecond"));
        return (msUTC1 - msUTC2).TotalMilliseconds;
    }
    
    public static object AdjustImpl(dynamic mkDateRec, dynamic just, object nothing, double offset, dynamic rec) {
        var t = CreateUTC(GetInt(rec, "year"), GetInt(rec, "month"), GetInt(rec, "day"), GetInt(rec, "hour"), GetInt(rec, "minute"), GetInt(rec, "second"), GetInt(rec, "millisecond"));
        var dt = t.AddMilliseconds(offset);
        
        var resMap = mkDateRec
            .Invoke((object)(int)dt.Year)
            .Invoke((object)(int)dt.Month)
            .Invoke((object)(int)dt.Day)
            .Invoke((object)(int)dt.Hour)
            .Invoke((object)(int)dt.Minute)
            .Invoke((object)(int)dt.Second)
            .Invoke((object)(int)dt.Millisecond);
        
        return just.Invoke(resMap);
    }
}
