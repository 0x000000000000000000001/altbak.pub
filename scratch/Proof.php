<?php
$logA = function($s) { return function() use($s) { echo $s . "\n"; }; };
$logB = function($s) { echo $s . "\n"; };
$exports['logA'] = $logA;
$exports['logB'] = $logB;
return $exports;
