using System;

namespace Data.Enum;

public static class FFI {
    public static char FromCharCode(int c) => (char)c;
    public static int ToCharCode(char c) => (int)c;
}
