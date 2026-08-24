<?php
interface Monoidish {
    public function mempty_();
    public function mappend_($x, $y);
}
class IntMonoidish implements Monoidish {
    public function mempty_() { return 1; }
    public function mappend_($x, $y) { return $x + $y; }
}
$exports['runPolymorphismFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    $m = new IntMonoidish();
    while ($n > 0) {
        $acc = $m->mappend_($acc, $m->mempty_());
        $n--;
    }
    return $acc;
};
return $exports;
