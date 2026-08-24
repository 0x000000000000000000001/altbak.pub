<?php
$exports['runArrayOpsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $arr = [];
    for ($i = 1; $i <= $n; $i++) {
        $arr[] = $i;
    }
    $sum = 0;
    foreach ($arr as $v) {
        if ($v % 2 === 0) $sum += $v;
    }
    return $sum;
};
return $exports;
