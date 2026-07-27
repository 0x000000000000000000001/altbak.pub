<?php

namespace Data\Tuple\Nested;

// ALL IMPORTS: Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Prelude/index.php';

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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Data_Tuple_Nested_uncurry9
function majData_majTuple_majNested_uncurry9($f_prime_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry9';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((((($f_prime_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'}))((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry9'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry9';

// Data_Tuple_Nested_uncurry8
function majData_majTuple_majNested_uncurry8($f_prime_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry8';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((((($f_prime_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'}))((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry8'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry8';

// Data_Tuple_Nested_uncurry7
function majData_majTuple_majNested_uncurry7($f_prime_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry7';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($f_prime_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'}))((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry7'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry7';

// Data_Tuple_Nested_uncurry6
function majData_majTuple_majNested_uncurry6($f_prime_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry6';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((($f_prime_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'}))((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry6'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry6';

// Data_Tuple_Nested_uncurry5
function majData_majTuple_majNested_uncurry5($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry5';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($f_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'}))((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry5'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry5';

// Data_Tuple_Nested_uncurry4
function majData_majTuple_majNested_uncurry4($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry4';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($f_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'}))((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry4'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry4';

// Data_Tuple_Nested_uncurry3
function majData_majTuple_majNested_uncurry3($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry3';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($f_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry3'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry3';

// Data_Tuple_Nested_uncurry2
function majData_majTuple_majNested_uncurry2($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry2'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry2';

// Data_Tuple_Nested_uncurry10
function majData_majTuple_majNested_uncurry10($f_prime_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry10';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((((((((($f_prime_0)(($v_1)->{'value0'}))((($v_1)->{'value1'})->{'value0'}))(((($v_1)->{'value1'})->{'value1'})->{'value0'}))((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))(((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}))((((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry10'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry10';

// Data_Tuple_Nested_uncurry1
function majData_majTuple_majNested_uncurry1($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_uncurry1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)(($v_1)->{'value0'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_uncurry1'] = __NAMESPACE__ . '\\majData_majTuple_majNested_uncurry1';

// Data_Tuple_Nested_tuple9
function majData_majTuple_majNested_tuple9($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $i_8 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple9';
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, new \Data\Tuple\Data_Tuple_Tuple($d_3, new \Data\Tuple\Data_Tuple_Tuple($e_4, new \Data\Tuple\Data_Tuple_Tuple($f_5, new \Data\Tuple\Data_Tuple_Tuple($g_6, new \Data\Tuple\Data_Tuple_Tuple($h_7, new \Data\Tuple\Data_Tuple_Tuple($i_8, $GLOBALS['Data_Unit_unit'])))))))));
  goto __end;;
  __end:
  return 9 < $__num ? $__res(...\array_slice(\func_get_args(), 9)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple9'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple9';

// Data_Tuple_Nested_tuple8
function majData_majTuple_majNested_tuple8($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple8';
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, new \Data\Tuple\Data_Tuple_Tuple($d_3, new \Data\Tuple\Data_Tuple_Tuple($e_4, new \Data\Tuple\Data_Tuple_Tuple($f_5, new \Data\Tuple\Data_Tuple_Tuple($g_6, new \Data\Tuple\Data_Tuple_Tuple($h_7, $GLOBALS['Data_Unit_unit']))))))));
  goto __end;;
  __end:
  return 8 < $__num ? $__res(...\array_slice(\func_get_args(), 8)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple8'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple8';

// Data_Tuple_Nested_tuple7
function majData_majTuple_majNested_tuple7($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple7';
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, new \Data\Tuple\Data_Tuple_Tuple($d_3, new \Data\Tuple\Data_Tuple_Tuple($e_4, new \Data\Tuple\Data_Tuple_Tuple($f_5, new \Data\Tuple\Data_Tuple_Tuple($g_6, $GLOBALS['Data_Unit_unit'])))))));
  goto __end;;
  __end:
  return 7 < $__num ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple7'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple7';

// Data_Tuple_Nested_tuple6
function majData_majTuple_majNested_tuple6($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple6';
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, new \Data\Tuple\Data_Tuple_Tuple($d_3, new \Data\Tuple\Data_Tuple_Tuple($e_4, new \Data\Tuple\Data_Tuple_Tuple($f_5, $GLOBALS['Data_Unit_unit']))))));
  goto __end;;
  __end:
  return 6 < $__num ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple6'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple6';

// Data_Tuple_Nested_tuple5
function majData_majTuple_majNested_tuple5($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple5';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, new \Data\Tuple\Data_Tuple_Tuple($d_3, new \Data\Tuple\Data_Tuple_Tuple($e_4, $GLOBALS['Data_Unit_unit'])))));
  goto __end;;
  __end:
  return 5 < $__num ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple5'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple5';

// Data_Tuple_Nested_tuple4
function majData_majTuple_majNested_tuple4($a_0, $b_1 = null, $c_2 = null, $d_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple4';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, new \Data\Tuple\Data_Tuple_Tuple($d_3, $GLOBALS['Data_Unit_unit']))));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple4'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple4';

// Data_Tuple_Nested_tuple3
function majData_majTuple_majNested_tuple3($a_0, $b_1 = null, $c_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple3';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, $GLOBALS['Data_Unit_unit'])));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple3'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple3';

// Data_Tuple_Nested_tuple2
function majData_majTuple_majNested_tuple2($a_0, $b_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, $GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple2'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple2';

// Data_Tuple_Nested_tuple10
function majData_majTuple_majNested_tuple10($a_0, $b_1 = null, $c_2 = null, $d_3 = null, $e_4 = null, $f_5 = null, $g_6 = null, $h_7 = null, $i_8 = null, $j_9 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple10';
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, new \Data\Tuple\Data_Tuple_Tuple($b_1, new \Data\Tuple\Data_Tuple_Tuple($c_2, new \Data\Tuple\Data_Tuple_Tuple($d_3, new \Data\Tuple\Data_Tuple_Tuple($e_4, new \Data\Tuple\Data_Tuple_Tuple($f_5, new \Data\Tuple\Data_Tuple_Tuple($g_6, new \Data\Tuple\Data_Tuple_Tuple($h_7, new \Data\Tuple\Data_Tuple_Tuple($i_8, new \Data\Tuple\Data_Tuple_Tuple($j_9, $GLOBALS['Data_Unit_unit']))))))))));
  goto __end;;
  __end:
  return 10 < $__num ? $__res(...\array_slice(\func_get_args(), 10)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple10'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple10';

// Data_Tuple_Nested_tuple1
function majData_majTuple_majNested_tuple1($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_tuple1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($a_0, $GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_tuple1'] = __NAMESPACE__ . '\\majData_majTuple_majNested_tuple1';

// Data_Tuple_Nested_over9
function majData_majTuple_majNested_over9($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over9';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((($v_1)->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)(((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}), ((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})))))))));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over9'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over9';

// Data_Tuple_Nested_over8
function majData_majTuple_majNested_over8($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over8';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((($v_1)->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}), (((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'}))))))));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over8'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over8';

// Data_Tuple_Nested_over7
function majData_majTuple_majNested_over7($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over7';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((($v_1)->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}), ((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})))))));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over7'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over7';

// Data_Tuple_Nested_over6
function majData_majTuple_majNested_over6($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over6';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((($v_1)->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}), (((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'}))))));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over6'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over6';

// Data_Tuple_Nested_over5
function majData_majTuple_majNested_over5($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over5';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((($v_1)->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}), ((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})))));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over5'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over5';

// Data_Tuple_Nested_over4
function majData_majTuple_majNested_over4($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over4';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((($v_1)->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}), (((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'}))));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over4'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over4';

// Data_Tuple_Nested_over3
function majData_majTuple_majNested_over3($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over3';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)(((($v_1)->{'value1'})->{'value1'})->{'value0'}), ((($v_1)->{'value1'})->{'value1'})->{'value1'})));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over3'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over3';

// Data_Tuple_Nested_over2
function majData_majTuple_majNested_over2($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)((($v_1)->{'value1'})->{'value0'}), (($v_1)->{'value1'})->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over2'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over2';

// Data_Tuple_Nested_over10
function majData_majTuple_majNested_over10($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over10';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((($v_1)->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((($v_1)->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}, new \Data\Tuple\Data_Tuple_Tuple(($o_0)((((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'}), (((((((((($v_1)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'}))))))))));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over10'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over10';

// Data_Tuple_Nested_over1
function majData_majTuple_majNested_over1($o_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_over1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($o_0)(($v_1)->{'value0'}), ($v_1)->{'value1'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_over1'] = __NAMESPACE__ . '\\majData_majTuple_majNested_over1';

// Data_Tuple_Nested_get9
function majData_majTuple_majNested_get9($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get9';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((((($v_0)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get9'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get9';

// Data_Tuple_Nested_get8
function majData_majTuple_majNested_get8($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get8';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (((((((($v_0)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get8'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get8';

// Data_Tuple_Nested_get7
function majData_majTuple_majNested_get7($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get7';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((((($v_0)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get7'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get7';

// Data_Tuple_Nested_get6
function majData_majTuple_majNested_get6($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get6';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (((((($v_0)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get6'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get6';

// Data_Tuple_Nested_get5
function majData_majTuple_majNested_get5($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get5';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((((($v_0)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get5'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get5';

// Data_Tuple_Nested_get4
function majData_majTuple_majNested_get4($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get4';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (((($v_0)->{'value1'})->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get4'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get4';

// Data_Tuple_Nested_get3
function majData_majTuple_majNested_get3($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get3';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($v_0)->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get3'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get3';

// Data_Tuple_Nested_get2
function majData_majTuple_majNested_get2($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($v_0)->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get2'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get2';

// Data_Tuple_Nested_get10
function majData_majTuple_majNested_get10($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get10';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (((((((((($v_0)->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value1'})->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get10'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get10';

// Data_Tuple_Nested_get1
function majData_majTuple_majNested_get1($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_get1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_get1'] = __NAMESPACE__ . '\\majData_majTuple_majNested_get1';

// Data_Tuple_Nested_curry9
function majData_majTuple_majNested_curry9($z_0, $f_prime_1 = null, $a_2 = null, $b_3 = null, $c_4 = null, $d_5 = null, $e_6 = null, $f_7 = null, $g_8 = null, $h_9 = null, $i_10 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry9';
  if ($__num < 11) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 11);
  }
  $__res = ($f_prime_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, new \Data\Tuple\Data_Tuple_Tuple($d_5, new \Data\Tuple\Data_Tuple_Tuple($e_6, new \Data\Tuple\Data_Tuple_Tuple($f_7, new \Data\Tuple\Data_Tuple_Tuple($g_8, new \Data\Tuple\Data_Tuple_Tuple($h_9, new \Data\Tuple\Data_Tuple_Tuple($i_10, $z_0))))))))));
  goto __end;;
  __end:
  return 11 < $__num ? $__res(...\array_slice(\func_get_args(), 11)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry9'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry9';

// Data_Tuple_Nested_curry8
function majData_majTuple_majNested_curry8($z_0, $f_prime_1 = null, $a_2 = null, $b_3 = null, $c_4 = null, $d_5 = null, $e_6 = null, $f_7 = null, $g_8 = null, $h_9 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry8';
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  $__res = ($f_prime_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, new \Data\Tuple\Data_Tuple_Tuple($d_5, new \Data\Tuple\Data_Tuple_Tuple($e_6, new \Data\Tuple\Data_Tuple_Tuple($f_7, new \Data\Tuple\Data_Tuple_Tuple($g_8, new \Data\Tuple\Data_Tuple_Tuple($h_9, $z_0)))))))));
  goto __end;;
  __end:
  return 10 < $__num ? $__res(...\array_slice(\func_get_args(), 10)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry8'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry8';

// Data_Tuple_Nested_curry7
function majData_majTuple_majNested_curry7($z_0, $f_prime_1 = null, $a_2 = null, $b_3 = null, $c_4 = null, $d_5 = null, $e_6 = null, $f_7 = null, $g_8 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry7';
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  $__res = ($f_prime_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, new \Data\Tuple\Data_Tuple_Tuple($d_5, new \Data\Tuple\Data_Tuple_Tuple($e_6, new \Data\Tuple\Data_Tuple_Tuple($f_7, new \Data\Tuple\Data_Tuple_Tuple($g_8, $z_0))))))));
  goto __end;;
  __end:
  return 9 < $__num ? $__res(...\array_slice(\func_get_args(), 9)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry7'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry7';

// Data_Tuple_Nested_curry6
function majData_majTuple_majNested_curry6($z_0, $f_prime_1 = null, $a_2 = null, $b_3 = null, $c_4 = null, $d_5 = null, $e_6 = null, $f_7 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry6';
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  $__res = ($f_prime_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, new \Data\Tuple\Data_Tuple_Tuple($d_5, new \Data\Tuple\Data_Tuple_Tuple($e_6, new \Data\Tuple\Data_Tuple_Tuple($f_7, $z_0)))))));
  goto __end;;
  __end:
  return 8 < $__num ? $__res(...\array_slice(\func_get_args(), 8)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry6'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry6';

// Data_Tuple_Nested_curry5
function majData_majTuple_majNested_curry5($z_0, $f_1 = null, $a_2 = null, $b_3 = null, $c_4 = null, $d_5 = null, $e_6 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry5';
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  $__res = ($f_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, new \Data\Tuple\Data_Tuple_Tuple($d_5, new \Data\Tuple\Data_Tuple_Tuple($e_6, $z_0))))));
  goto __end;;
  __end:
  return 7 < $__num ? $__res(...\array_slice(\func_get_args(), 7)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry5'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry5';

// Data_Tuple_Nested_curry4
function majData_majTuple_majNested_curry4($z_0, $f_1 = null, $a_2 = null, $b_3 = null, $c_4 = null, $d_5 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry4';
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__res = ($f_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, new \Data\Tuple\Data_Tuple_Tuple($d_5, $z_0)))));
  goto __end;;
  __end:
  return 6 < $__num ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry4'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry4';

// Data_Tuple_Nested_curry3
function majData_majTuple_majNested_curry3($z_0, $f_1 = null, $a_2 = null, $b_3 = null, $c_4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry3';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__res = ($f_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, $z_0))));
  goto __end;;
  __end:
  return 5 < $__num ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry3'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry3';

// Data_Tuple_Nested_curry2
function majData_majTuple_majNested_curry2($z_0, $f_1 = null, $a_2 = null, $b_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry2';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ($f_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, $z_0)));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry2'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry2';

// Data_Tuple_Nested_curry10
function majData_majTuple_majNested_curry10($z_0, $f_prime_1 = null, $a_2 = null, $b_3 = null, $c_4 = null, $d_5 = null, $e_6 = null, $f_7 = null, $g_8 = null, $h_9 = null, $i_10 = null, $j_11 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry10';
  if ($__num < 12) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 12);
  }
  $__res = ($f_prime_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, new \Data\Tuple\Data_Tuple_Tuple($b_3, new \Data\Tuple\Data_Tuple_Tuple($c_4, new \Data\Tuple\Data_Tuple_Tuple($d_5, new \Data\Tuple\Data_Tuple_Tuple($e_6, new \Data\Tuple\Data_Tuple_Tuple($f_7, new \Data\Tuple\Data_Tuple_Tuple($g_8, new \Data\Tuple\Data_Tuple_Tuple($h_9, new \Data\Tuple\Data_Tuple_Tuple($i_10, new \Data\Tuple\Data_Tuple_Tuple($j_11, $z_0)))))))))));
  goto __end;;
  __end:
  return 12 < $__num ? $__res(...\array_slice(\func_get_args(), 12)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry10'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry10';

// Data_Tuple_Nested_curry1
function majData_majTuple_majNested_curry1($z_0, $f_1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_majNested_curry1';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)(new \Data\Tuple\Data_Tuple_Tuple($a_2, $z_0));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Tuple_Nested_curry1'] = __NAMESPACE__ . '\\majData_majTuple_majNested_curry1';

