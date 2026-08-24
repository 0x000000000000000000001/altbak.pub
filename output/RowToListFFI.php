<?php
$exports['runRowToListFFI'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $rec = (object)["a" => 1, "b" => "hello", "c" => true];
        $countKeys = function($r) {
            return count(get_object_vars($r));
        };
        $sum += $countKeys($rec) + ($rec->a * 2);
    }
    return $sum;
};
return $exports;
