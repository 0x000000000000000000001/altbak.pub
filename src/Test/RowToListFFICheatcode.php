<?php
$exports['runRowToListFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $rec = (object)["a" => 1, "b" => "hello", "c" => true];
        $sum += $rec->a;
    }
    return $sum;
};
return $exports;
