<?php
$exports['runRecordsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $r = (object)["a" => 0, "b" => (object)["c" => 0, "d" => (object)["e" => 0, "f" => 0]]];
    while ($n > 0) {
        $r->a += 1;
        $r->b->c += 2;
        $r->b->d->e += 3;
        $r->b->d->f += ($n % 5);
        $n--;
    }
    return $r->b->d->f;
};
return $exports;
