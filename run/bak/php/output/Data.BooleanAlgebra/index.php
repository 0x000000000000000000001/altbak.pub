<?php

namespace Data\BooleanAlgebra;

// ALL IMPORTS: Data.BooleanAlgebra, Data.HeytingAlgebra, Data.Symbol, Data.Unit, Prim, Prim.Row, Prim.RowList, Type.Proxy
// TO REQUIRE: Data.BooleanAlgebra, Data.HeytingAlgebra, Data.Symbol, Data.Unit, Type.Proxy
require_once __DIR__ . '/../Data.BooleanAlgebra/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Type.Proxy/index.php';

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




// Data_BooleanAlgebra_booleanAlgebraUnit
$GLOBALS['Data_BooleanAlgebra_booleanAlgebraUnit'] = (object)["HeytingAlgebra0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_HeytingAlgebra_heytingAlgebraUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_BooleanAlgebra_booleanAlgebraRecordNil
$GLOBALS['Data_BooleanAlgebra_booleanAlgebraRecordNil'] = (object)["HeytingAlgebraRecord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_HeytingAlgebra_heytingAlgebraRecordNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_BooleanAlgebra_booleanAlgebraRecordCons
function majData_majBooleanmajAlgebra_booleanmajAlgebramajRecordmajCons($dictIsSymbol_0, $_dollar__unused_1 = null, $dictBooleanAlgebraRecord_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBooleanmajAlgebra_booleanmajAlgebramajRecordmajCons';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $heytingAlgebraRecordCons1_3_0 = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraRecordCons'])($dictIsSymbol_0))(null))((($dictBooleanAlgebraRecord_2)->{'HeytingAlgebraRecord0'})(null));
  $__res = function($dictBooleanAlgebra_4) use ($heytingAlgebraRecordCons1_3_0) {
  $__num = \func_num_args();
  $heytingAlgebraRecordCons2_5_1 = ($heytingAlgebraRecordCons1_3_0)((($dictBooleanAlgebra_4)->{'HeytingAlgebra0'})(null));
  $__res = (object)["HeytingAlgebraRecord0" => function($_dollar__unused_6) use ($heytingAlgebraRecordCons2_5_1) {
  $__num = \func_num_args();
  $__res = $heytingAlgebraRecordCons2_5_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_BooleanAlgebra_booleanAlgebraRecordCons'] = __NAMESPACE__ . '\\majData_majBooleanmajAlgebra_booleanmajAlgebramajRecordmajCons';

// Data_BooleanAlgebra_booleanAlgebraRecord
function majData_majBooleanmajAlgebra_booleanmajAlgebramajRecord($_dollar__unused_0, $dictBooleanAlgebraRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBooleanmajAlgebra_booleanmajAlgebramajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictBooleanAlgebraRecord_1)->{'HeytingAlgebraRecord0'})(null);
  $heytingAlgebraRecord1_3_1 = (object)["ff" => ((($__local_var_2_0)->{'ffRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "tt" => ((($__local_var_2_0)->{'ttRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "conj" => (($__local_var_2_0)->{'conjRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "disj" => (($__local_var_2_0)->{'disjRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "implies" => (($__local_var_2_0)->{'impliesRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "not" => (($__local_var_2_0)->{'notRecord'})(new \Type\Proxy\Type_Proxy_Proxy())];
  $__res = (object)["HeytingAlgebra0" => function($_dollar__unused_4) use ($heytingAlgebraRecord1_3_1) {
  $__num = \func_num_args();
  $__res = $heytingAlgebraRecord1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_BooleanAlgebra_booleanAlgebraRecord'] = __NAMESPACE__ . '\\majData_majBooleanmajAlgebra_booleanmajAlgebramajRecord';

// Data_BooleanAlgebra_booleanAlgebraProxy
$GLOBALS['Data_BooleanAlgebra_booleanAlgebraProxy'] = (object)["HeytingAlgebra0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_HeytingAlgebra_heytingAlgebraProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_BooleanAlgebra_booleanAlgebraFn
function majData_majBooleanmajAlgebra_booleanmajAlgebramajFn($dictBooleanAlgebra_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBooleanmajAlgebra_booleanmajAlgebramajFn';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictBooleanAlgebra_0)->{'HeytingAlgebra0'})(null);
  $ff1_2_1 = ($__local_var_1_0)->{'ff'};
  $tt1_3_2 = ($__local_var_1_0)->{'tt'};
  $heytingAlgebraFunction_3_2 = (object)["ff" => function($v_4) use ($ff1_2_1) {
  $__num = \func_num_args();
  $__res = $ff1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "tt" => function($v_4) use ($tt1_3_2) {
  $__num = \func_num_args();
  $__res = $tt1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "implies" => (function() use ($__local_var_1_0) {
  $__fn = function($f_4, $g_5 = null, $a_6 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)->{'implies'})(($f_4)($a_6)))(($g_5)($a_6));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "conj" => (function() use ($__local_var_1_0) {
  $__fn = function($f_4, $g_5 = null, $a_6 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)->{'conj'})(($f_4)($a_6)))(($g_5)($a_6));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "disj" => (function() use ($__local_var_1_0) {
  $__fn = function($f_4, $g_5 = null, $a_6 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)->{'disj'})(($f_4)($a_6)))(($g_5)($a_6));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "not" => (function() use ($__local_var_1_0) {
  $__fn = function($f_4, $a_5 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_1_0)->{'not'})(($f_4)($a_5));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["HeytingAlgebra0" => function($_dollar__unused_4) use ($heytingAlgebraFunction_3_2) {
  $__num = \func_num_args();
  $__res = $heytingAlgebraFunction_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_BooleanAlgebra_booleanAlgebraFn'] = __NAMESPACE__ . '\\majData_majBooleanmajAlgebra_booleanmajAlgebramajFn';

// Data_BooleanAlgebra_booleanAlgebraBoolean
$GLOBALS['Data_BooleanAlgebra_booleanAlgebraBoolean'] = (object)["HeytingAlgebra0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

