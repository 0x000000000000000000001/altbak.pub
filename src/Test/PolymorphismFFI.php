<?php
$exports['runPolymorphismFFI'] = function($limit) {
    $n = (int)$limit;
    
    $lengthString = function($s) { return strlen($s); };
    $lengthArray = function($arr) { return count($arr); };
    
    $computeLength = function($dict) {
        return function($x) use ($dict) {
            return ($dict->length)($x);
        };
    };
    
    $dictString = (object)["length" => $lengthString];
    $dictArray = (object)["length" => $lengthArray];
    
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $sum += $computeLength($dictString)("hello") + $computeLength($dictArray)([1, 2, 3]);
    }
    return $sum;
};
return $exports;
