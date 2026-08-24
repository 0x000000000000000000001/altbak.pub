<?php
$exports['runListOpsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 1; $i <= $n; $i++) {
        if ($i % 2 === 0) $sum += $i;
    }
    return $sum;
};
return $exports;
