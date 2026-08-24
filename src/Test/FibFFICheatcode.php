<?php
$exports['runFibFFICheatcode'] = function($n) use (&$exports) {
    if ($n === 0) return 0;
    if ($n === 1) return 1;
    $f = $exports['runFibFFICheatcode'];
    return $f($n - 1) + $f($n - 2);
};
return $exports;
