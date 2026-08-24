<?php
$GLOBALS['Data_Unit_unit'] = "UNIT";
function phpurs_execute_effect($val) {
  if (\is_callable($val)) {
    return $val($GLOBALS['Data_Unit_unit']);
  }
  return $val;
}
$log = function($s) { return function() use($s) { echo $s . "\n"; }; };
$act_1 = function() use ($log) {
    $dummy = phpurs_execute_effect(function() { return 3; });
    return phpurs_execute_effect($log("result is " . $dummy));
};
phpurs_execute_effect($act_1);
