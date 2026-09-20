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
    $polyLoop = function($dictionary, $count, $acc) {
        while ($count !== 0) {
            $acc = ($dictionary->mappend_)($acc)($dictionary->mempty_);
            $count--;
        }
        return $acc;
    };
    return $polyLoop($dict, $n, 0);
};
return $exports;
