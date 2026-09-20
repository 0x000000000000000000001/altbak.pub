<?php

$consumeResult = function($expected, $result) {
    return function() use ($expected, $result) {
        $GLOBALS['altbak_extended_result'] = $result;
        if ($result !== $expected) throw new \RuntimeException('Unstable extended benchmark result');
    };
};

$exports['consumeResult'] = $consumeResult;
return $exports;
