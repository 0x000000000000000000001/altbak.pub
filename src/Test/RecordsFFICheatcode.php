<?php
$exports['runRecordsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $obj = (object)['a' => 1, 'b' => (object)['c' => 2, 'd' => 3]];
    for ($i = 0; $i < $n; $i++) {
        $obj->b->c += 1;
    }
    return $obj->b->c - 2;
};
return $exports;
