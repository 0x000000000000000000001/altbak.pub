<?php
$exports['runPolymorphismFFI'] = function($limit) {
    $n = (int)$limit;
    $dict = (object)[
        "mempty_" => 1,
        "mappend_" => function($x) {
            return function($y) use ($x) {
                return $x + $y;
            };
        }
    ];
    $acc = 0;
    while ($n > 0) {
        $acc = ($dict->mappend_)($acc)($dict->mempty_);
        $n--;
    }
    return $acc;
};
return $exports;
