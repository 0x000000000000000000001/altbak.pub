<?php

$consumeResult = function($expected) {
    return function($result) use ($expected) {
        return function() use ($expected, $result) {
            $GLOBALS['altbak_extended_result'] = $result;
            if ($result !== $expected) throw new \RuntimeException('Unstable extended benchmark result');
        };
    };
};

$exports['consumeResult'] = $consumeResult;
return $exports;
