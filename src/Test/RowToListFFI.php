<?php
$exports['runRowToListFFI'] = function($limit) {
    $dictNil = (object)["keysImpl" => function($_) { return 0; }];
    $dictCons = function($tail) {
        return (object)["keysImpl" => function($_) use ($tail) {
            return 1 + ($tail->keysImpl)(null);
        }];
    };
    $dict = $dictCons($dictCons($dictCons($dictCons($dictCons($dictNil)))));
    return ($dict->keysImpl)(null);
};
return $exports;
