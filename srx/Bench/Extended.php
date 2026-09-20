<?php

$consumeResult = function($result) {
    return function() use ($result) {
        $GLOBALS['altbak_extended_result'] = $result;
    };
};

$exports['consumeResult'] = $consumeResult;
return $exports;
