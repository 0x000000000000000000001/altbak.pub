<?php
interface ShowableCheat {
    public function showLen();
}
class StrShowCheat implements ShowableCheat {
    public function showLen() { return 5; }
}
class ArrShowCheat implements ShowableCheat {
    public function showLen() { return 3; }
}

$exports['runPolymorphismFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    $s1 = new StrShowCheat();
    $s2 = new ArrShowCheat();
    for ($i = 0; $i < $n; $i++) {
        $sum += $s1->showLen() + $s2->showLen();
    }
    return $sum;
};
return $exports;
