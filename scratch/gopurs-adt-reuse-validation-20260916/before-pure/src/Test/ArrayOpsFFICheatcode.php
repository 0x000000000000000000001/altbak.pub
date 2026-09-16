<?php
$exports['runArrayOpsFFICheatcode'] = function($limit) {
    $end = (int)$limit;
    $step = $end >= 1 ? 1 : -1;
    $sum = 0;
    for ($value = 1; ; $value += $step) {
        if ($value % 2 === 0) $sum += $value;
        if ($value === $end) break;
    }
    return $sum;
};
return $exports;
