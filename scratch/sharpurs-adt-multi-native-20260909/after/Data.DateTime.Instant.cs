using System;


namespace Data.DateTime.Instant;

public static class FFI {
    public static double FromDateTimeImpl(int y, int mo, int d, int h, int mi, int s, int ms) {
        var dt = new global::System.DateTime(y, mo, d, h, mi, s, ms, global::System.DateTimeKind.Utc);
        return new global::System.DateTimeOffset(dt).ToUnixTimeMilliseconds();
    }
    
    public static object ToDateTimeImpl(dynamic ctor, double instant) {
        var dt = global::System.DateTimeOffset.FromUnixTimeMilliseconds((long)instant).UtcDateTime;
        return ctor.Invoke((object)dt.Year).Invoke((object)dt.Month).Invoke((object)dt.Day).Invoke((object)dt.Hour).Invoke((object)dt.Minute).Invoke((object)dt.Second).Invoke((object)dt.Millisecond);
    }
}
