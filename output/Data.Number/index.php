<?php

namespace Data\Number;

// ALL IMPORTS: Data.Function.Uncurried, Data.Maybe, Data.Number, Prim
// TO REQUIRE: Data.Function.Uncurried, Data.Maybe, Data.Number
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Number/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
  class Phpurs_Data7 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; } }
  class Phpurs_Data8 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; } }
  class Phpurs_Data9 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; } }
  class Phpurs_Data10 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; } }
  class Phpurs_Data11 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; $this->value10 = $value10; } }
  class Phpurs_Data12 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10, $value11; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5, $value6, $value7, $value8, $value9, $value10, $value11) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; $this->value6 = $value6; $this->value7 = $value7; $this->value8 = $value8; $this->value9 = $value9; $this->value10 = $value10; $this->value11 = $value11; } }
}
if (!\function_exists(__NAMESPACE__ . '\\phpurs_curry_fallback')) {
  function phpurs_curry_fallback($fn, $args, $expected) {
    $missing = $expected - \count($args);
    if ($missing === 1) {
      return function($a) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num > 1) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a;
        return $fn(...$args);
      };
    }
    if ($missing === 2) {
      return function($a, $b = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 2) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b;
        return $fn(...$args);
      };
    }
    if ($missing === 3) {
      return function($a, $b = null, $c = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 2) { $args[] = $a; $args[] = $b; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 3) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b; $args[] = $c;
        return $fn(...$args);
      };
    }
    if ($missing === 4) {
      return function($a, $b = null, $c = null, $d = null) use ($fn, $args, $expected) {
        $num = \func_num_args();
        if ($num === 1) { $args[] = $a; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 2) { $args[] = $a; $args[] = $b; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num === 3) { $args[] = $a; $args[] = $b; $args[] = $c; return phpurs_curry_fallback($fn, $args, $expected); }
        if ($num > 4) {
          $merged = \array_merge($args, \func_get_args());
          $res = $fn(...\array_slice($merged, 0, $expected));
          return $res(...\array_slice($merged, $expected));
        }
        $args[] = $a; $args[] = $b; $args[] = $c; $args[] = $d;
        return $fn(...$args);
      };
    }
    return function(...$more) use ($fn, $args, $expected) {
      $merged = \array_merge($args, $more);
      if (\count($merged) >= $expected) {
        $res = $fn(...\array_slice($merged, 0, $expected));
        if (\count($merged) > $expected) {
          return $res(...\array_slice($merged, $expected));
        }
        return $res;
      }
      return phpurs_curry_fallback($fn, $merged, $expected);
    };
  }
}
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
function majData_majNumber_abs(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_abs';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_abs'] = __NAMESPACE__ . '\\majData_majNumber_abs';

function majData_majNumber_acos(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_acos';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_acos'] = __NAMESPACE__ . '\\majData_majNumber_acos';

function majData_majNumber_asin(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_asin';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_asin'] = __NAMESPACE__ . '\\majData_majNumber_asin';

function majData_majNumber_atan(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_atan';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_atan'] = __NAMESPACE__ . '\\majData_majNumber_atan';

function majData_majNumber_atan2(float $v0, $v1 = null): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_atan2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_atan2'] = __NAMESPACE__ . '\\majData_majNumber_atan2';

function majData_majNumber_ceil(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_ceil';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_ceil'] = __NAMESPACE__ . '\\majData_majNumber_ceil';

function majData_majNumber_cos(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_cos';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_cos'] = __NAMESPACE__ . '\\majData_majNumber_cos';

function majData_majNumber_exp(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_exp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_exp'] = __NAMESPACE__ . '\\majData_majNumber_exp';

function majData_majNumber_floor(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_floor';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_floor'] = __NAMESPACE__ . '\\majData_majNumber_floor';

$GLOBALS['Data_Number_fromStringImpl'] = new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Number_infinity'] = new class { public function __invoke(...$args) { return $this; } };
function majData_majNumber_ismajFinite(float $v0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_ismajFinite';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_isFinite'] = __NAMESPACE__ . '\\majData_majNumber_ismajFinite';

function majData_majNumber_ismajNamajN(float $v0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_ismajNamajN';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_isNaN'] = __NAMESPACE__ . '\\majData_majNumber_ismajNamajN';

function majData_majNumber_log(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_log';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_log'] = __NAMESPACE__ . '\\majData_majNumber_log';

function majData_majNumber_max(float $v0, $v1 = null): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_max';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_max'] = __NAMESPACE__ . '\\majData_majNumber_max';

function majData_majNumber_min(float $v0, $v1 = null): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_min';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_min'] = __NAMESPACE__ . '\\majData_majNumber_min';

$GLOBALS['Data_Number_nan'] = new class { public function __invoke(...$args) { return $this; } };
function majData_majNumber_pow(float $v0, $v1 = null): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_pow';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_pow'] = __NAMESPACE__ . '\\majData_majNumber_pow';

function majData_majNumber_remainder(float $v0, $v1 = null): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_remainder';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0, $v1);
}
$GLOBALS['Data_Number_remainder'] = __NAMESPACE__ . '\\majData_majNumber_remainder';

function majData_majNumber_round(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_round';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_round'] = __NAMESPACE__ . '\\majData_majNumber_round';

function majData_majNumber_sign(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_sign';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_sign'] = __NAMESPACE__ . '\\majData_majNumber_sign';

function majData_majNumber_sin(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_sin';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_sin'] = __NAMESPACE__ . '\\majData_majNumber_sin';

function majData_majNumber_sqrt(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_sqrt';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_sqrt'] = __NAMESPACE__ . '\\majData_majNumber_sqrt';

function majData_majNumber_tan(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_tan';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_tan'] = __NAMESPACE__ . '\\majData_majNumber_tan';

function majData_majNumber_trunc(float $v0): float|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majNumber_trunc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $f = new class { public function __invoke(...$args) { return $this; } };
  return $f($v0);
}
$GLOBALS['Data_Number_trunc'] = __NAMESPACE__ . '\\majData_majNumber_trunc';





// Data_Number_tau
$GLOBALS['Data_Number_tau'] = 6.283185307179586;

// Data_Number_sqrt2
$GLOBALS['Data_Number_sqrt2'] = 1.4142135623730951;

// Data_Number_sqrt1_2
$GLOBALS['Data_Number_sqrt1_2'] = 0.7071067811865476;

// Data_Number_pi
$GLOBALS['Data_Number_pi'] = 3.141592653589793;

// Data_Number_log2e
$GLOBALS['Data_Number_log2e'] = 1.4426950408889634;

// Data_Number_log10e
$GLOBALS['Data_Number_log10e'] = 0.4342944819032518;

// Data_Number_ln2
$GLOBALS['Data_Number_ln2'] = 0.6931471805599453;

// Data_Number_ln10
$GLOBALS['Data_Number_ln10'] = 2.302585092994046;

// Data_Number_fromString
function majData_majNumber_frommajString(string $str_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majNumber_frommajString';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Number_fromStringImpl'])($str_0, $GLOBALS['Data_Number_isFinite'], $GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Number_fromString'] = __NAMESPACE__ . '\\majData_majNumber_frommajString';

// Data_Number_e
$GLOBALS['Data_Number_e'] = 2.718281828459045;

