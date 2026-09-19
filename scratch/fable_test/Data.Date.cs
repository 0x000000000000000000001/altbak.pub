using System;


namespace Data.Date;

public static class FFI {
    private static global::System.DateTime CreateDate(int y, int m, int d) {
        return new global::System.DateTime((int)y, (int)m, (int)d, 0, 0, 0, global::System.DateTimeKind.Utc);
    }
    
    public static object CanonicalDateImpl(dynamic ctor, int y, int m, int d) {
        var date = CreateDate(y, m, d);
        return ctor.Invoke((object)date.Year).Invoke((object)date.Month).Invoke((object)date.Day);
    }
    
    public static int CalcWeekday(int y, int m, int d) {
        var date = CreateDate(y, m, d);
        return (int)date.DayOfWeek;
    }
    
    public static double CalcDiff(int y1, int m1, int d1, int y2, int m2, int d2) {
        var dt1 = CreateDate(y1, m1, d1);
        var dt2 = CreateDate(y2, m2, d2);
        return (dt1 - dt2).TotalMilliseconds;
    }
}
