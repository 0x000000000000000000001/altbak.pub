<?php
$exports['runRowToListFFI'] = function($limit) {
    $dictNil = (object)["keysImpl" => function($_) { return 0; }];
    $dictCons = function($tail) {
        return (object)["keysImpl" => function($_) use ($tail) {
            return 1 + ($tail->keysImpl)(null);
        }];
    };
    $dict = $dictCons($dictCons($dictCons($dictCons($dictCons($dictNil)))));
    // Dynamic PHP represents the heterogeneous record and its recursive
    // dictionary explicitly; keys, like PureScript, does not inspect values.
    $keys = function($dictionary, $record) { return ($dictionary->keysImpl)(null); };
    $record = (object)['a' => 1, 'b' => 'two', 'c' => true, 'd' => 4.0, 'e' => 'five'];
    return $keys($dict, $record);
};
return $exports;
