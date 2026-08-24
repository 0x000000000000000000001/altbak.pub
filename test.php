<?php
$GLOBALS['Data_Unit_unit'] = "UNIT";
function phpurs_execute_effect($val) {
  if (\is_callable($val)) {
    return $val($GLOBALS['Data_Unit_unit']);
  }
  return $val;
}
$log = function($s) { return function() use($s) { echo $s . "\n"; }; };
phpurs_execute_effect($log("125"));
