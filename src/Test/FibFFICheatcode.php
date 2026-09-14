<?php
class FibonacciNative {
    public static function evaluate($n) {
        if ($n <= 1) return $n;
        return self::evaluate($n - 1) + self::evaluate($n - 2);
    }
}
$exports['runFibFFICheatcode'] = function($limit) {
    return FibonacciNative::evaluate((int)$limit);
};
return $exports;
