<?php

namespace Control\Apply;

// ALL IMPORTS: Control.Apply, Control.Category, Data.Function, Data.Functor, Prim, Type.Proxy
// TO REQUIRE: Control.Apply, Control.Category, Data.Function, Data.Functor, Type.Proxy
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Control_Apply = \call_user_func(function() {
  $exports = [];
$arrayApply = function($fs, $xs) use (&$arrayApply) {
    $r = []; foreach($fs as $f) { foreach($xs as $x) { $r[] = $f($x); } } return $r;
};

$exports['arrayApply'] = $arrayApply;
return $exports;
  return $exports;
});
function majControl_majApply_arraymajApply($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majApply_arraymajApply';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Apply;
  $f = (\array_key_exists('arrayApply', $ffi_Control_Apply) ? $ffi_Control_Apply['arrayApply'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Apply_arrayApply'] = __NAMESPACE__ . '\\majControl_majApply_arraymajApply';





// Control_Apply_applyProxy
$GLOBALS['Control_Apply_applyProxy'] = (object)["apply" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) {
  $__num = \func_num_args();
  $__res = new \Type\Proxy\Type_Proxy_Proxy();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_functorProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Apply_applyFn
$GLOBALS['Control_Apply_applyFn'] = (object)["apply" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($g_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($x_2) use ($f_0, $g_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($x_2))(($g_1)($x_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_functorFn'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Apply_applyArray
$GLOBALS['Control_Apply_applyArray'] = (object)["apply" => $GLOBALS['Control_Apply_arrayApply'], "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_functorArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Apply_apply
function majControl_majApply_apply($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majApply_apply';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'apply'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Apply_apply'] = __NAMESPACE__ . '\\majControl_majApply_apply';

// Control_Apply_applyFirst
function majControl_majApply_applymajFirst($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majApply_applymajFirst';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($a_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($Functor0_1_0, $a_2, $dictApply_0) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($GLOBALS['Data_Function_const']))($a_2)))($b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Apply_applyFirst'] = __NAMESPACE__ . '\\majControl_majApply_applymajFirst';

// Control_Apply_applySecond
function majControl_majApply_applymajSecond($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majApply_applymajSecond';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($a_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($Functor0_1_0, $a_2, $dictApply_0) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})(function($v_4) {
  $__num = \func_num_args();
  $__res = function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_2)))($b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Apply_applySecond'] = __NAMESPACE__ . '\\majControl_majApply_applymajSecond';

// Control_Apply_lift2
function majControl_majApply_lift2($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majApply_lift2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($Functor0_1_0, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_1_0, $a_3, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($f_2))($a_3)))($b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Apply_lift2'] = __NAMESPACE__ . '\\majControl_majApply_lift2';

// Control_Apply_lift3
function majControl_majApply_lift3($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majApply_lift3';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($Functor0_1_0, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_1_0, $a_3, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($c_5) use ($Functor0_1_0, $a_3, $b_4, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($f_2))($a_3)))($b_4)))($c_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Apply_lift3'] = __NAMESPACE__ . '\\majControl_majApply_lift3';

// Control_Apply_lift4
function majControl_majApply_lift4($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majApply_lift4';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($Functor0_1_0, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_1_0, $a_3, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($c_5) use ($Functor0_1_0, $a_3, $b_4, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($d_6) use ($Functor0_1_0, $a_3, $b_4, $c_5, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($dictApply_0)->{'apply'})(((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($f_2))($a_3)))($b_4)))($c_5)))($d_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Apply_lift4'] = __NAMESPACE__ . '\\majControl_majApply_lift4';

// Control_Apply_lift5
function majControl_majApply_lift5($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majApply_lift5';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($Functor0_1_0, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_1_0, $a_3, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($c_5) use ($Functor0_1_0, $a_3, $b_4, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($d_6) use ($Functor0_1_0, $a_3, $b_4, $c_5, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($e_7) use ($Functor0_1_0, $a_3, $b_4, $c_5, $d_6, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($dictApply_0)->{'apply'})(((($dictApply_0)->{'apply'})(((($dictApply_0)->{'apply'})(((($Functor0_1_0)->{'map'})($f_2))($a_3)))($b_4)))($c_5)))($d_6)))($e_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Apply_lift5'] = __NAMESPACE__ . '\\majControl_majApply_lift5';

