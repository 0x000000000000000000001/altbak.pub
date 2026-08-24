<?php
$exports['runPolymorphismFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    
    $computeLength = function($x) {
        if (is_string($x)) return strlen($x);
        if (is_array($x)) return count($x);
        return 0;
    };
    
    for ($i = 0; $i < $n; $i++) {
        $sum += $computeLength("hello") + $computeLength([1, 2, 3]);
    }
    return $sum;
};
return $exports;
