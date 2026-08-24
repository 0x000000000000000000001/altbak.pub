<?php
$exports['runLazyEvaluationFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    for ($i = 0; $i < $n; $i++) {
        $acc += 1000;
    }
    return $acc;
};
return $exports;
