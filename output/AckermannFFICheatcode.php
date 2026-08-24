<?php
$exports['runAckermannFFICheatcode'] = function($ignore) {
    $ack = function($m, $n) use (&$ack) {
        if ($m === 0) return $n + 1;
        if ($m > 0 && $n === 0) return $ack($m - 1, 1);
        return $ack($m - 1, $ack($m, $n - 1));
    };
    return $ack(3, 4);
};
return $exports;
