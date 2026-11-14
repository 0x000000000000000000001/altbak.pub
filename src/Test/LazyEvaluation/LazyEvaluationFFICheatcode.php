<?php
$exports['runLazyEvaluationFFICheatcode'] = function($limit) {
    $n = (int)$limit * 1000; // 1000 * 1000 = 1000000
    $acc = 0;
    for ($i = 0; $i < $n; $i++) {
        $acc += 1;
    }
    return $acc;
};
return $exports;
