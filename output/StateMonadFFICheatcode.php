<?php
$exports['runStateMonadFFICheatcode'] = function($limit) {
    $state = 0;
    for ($i = 0; $i < 60; $i++) {
        for ($j = 0; $j < 20; $j++) {
            $state += 1;
        }
    }
    return $state;
};
return $exports;
