<?php
class AckermannNative {
    public static function evaluate($m, $n) {
        if ($m === 0) return $n + 1;
        if ($n === 0) return self::evaluate($m - 1, 1);
        return self::evaluate($m - 1, self::evaluate($m, $n - 1));
    }
}
$exports['runAckermannFFICheatcode'] = function($m) {
    return AckermannNative::evaluate((int)$m, 4);
};
return $exports;
