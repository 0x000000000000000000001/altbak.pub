<?php
$exports['runAckermannFFI'] = function($m) {
    $ack = function($m, $n) use (&$ack) {
        if ($m === 0) return $n + 1;
        if ($m > 0 && $n === 0) return $ack($m - 1, 1);
        return $ack($m - 1, $ack($m, $n - 1));
    };
    return $ack((int)$m, 4);
};
return $exports;
