<?php

namespace Control\Comonad\Store\Class;

// ALL IMPORTS: Control.Comonad, Control.Comonad.Env.Trans, Control.Comonad.Store.Class, Control.Comonad.Store.Trans, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Extend, Control.Semigroupoid, Data.Function, Data.Functor, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Comonad, Control.Comonad.Env.Trans, Control.Comonad.Store.Class, Control.Comonad.Store.Trans, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Extend, Control.Semigroupoid, Data.Function, Data.Functor, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Comonad.Env.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Store.Class/index.php';
require_once __DIR__ . '/../Control.Comonad.Store.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Traced.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Control_Comonad_Store_Class_pos
function majControl_majComonad_majStore_majClass_pos($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_pos';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'pos'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_pos'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_pos';

// Control_Comonad_Store_Class_peek
function majControl_majComonad_majStore_majClass_peek($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_peek';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'peek'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_peek'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_peek';

// Control_Comonad_Store_Class_peeks
function majControl_majComonad_majStore_majClass_peeks($dictComonadStore_0, $f_1 = null, $x_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_peeks';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictComonadStore_0)->{'peek'})(($f_1)((($dictComonadStore_0)->{'pos'})($x_2))))($x_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_peeks'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_peeks';

// Control_Comonad_Store_Class_seeks
function majControl_majComonad_majStore_majClass_seeks($dictComonadStore_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_seeks';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $duplicate_1_0 = (((((($dictComonadStore_0)->{'Comonad0'})(null))->{'Extend0'})(null))->{'extend'})(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($f_2) use ($dictComonadStore_0, $duplicate_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3) use ($dictComonadStore_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictComonadStore_0)->{'peek'})(($f_2)((($dictComonadStore_0)->{'pos'})($x_3))))($x_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($duplicate_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_seeks'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_seeks';

// Control_Comonad_Store_Class_seek
function majControl_majComonad_majStore_majClass_seek($dictComonadStore_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_seek';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $duplicate_1_0 = (((((($dictComonadStore_0)->{'Comonad0'})(null))->{'Extend0'})(null))->{'extend'})(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = function($s_2) use ($dictComonadStore_0, $duplicate_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadStore_0)->{'peek'})($s_2)))($duplicate_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_seek'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_seek';

// Control_Comonad_Store_Class_experiment
function majControl_majComonad_majStore_majClass_experiment($dictComonadStore_0, $dictFunctor_1 = null, $f_2 = null, $x_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_experiment';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictFunctor_1)->{'map'})(function($a_4) use ($dictComonadStore_0, $x_3) {
  $__num = \func_num_args();
  $__res = ((($dictComonadStore_0)->{'peek'})($a_4))($x_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)((($dictComonadStore_0)->{'pos'})($x_3)));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_experiment'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_experiment';

// Control_Comonad_Store_Class_comonadStoreTracedT
function majControl_majComonad_majStore_majClass_comonadmajStoremajTracedmajT($dictComonadStore_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_comonadmajStoremajTracedmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $pos1_1_0 = ($dictComonadStore_0)->{'pos'};
  $Comonad0_2_1 = (($dictComonadStore_0)->{'Comonad0'})(null);
  $__local_var_3_2 = (($Comonad0_2_1)->{'Extend0'})(null);
  $Functor0_4_3 = (($__local_var_3_2)->{'Functor0'})(null);
  $__local_var_5_4 = (($__local_var_3_2)->{'Functor0'})(null);
  $functorTracedT1_5_4 = (object)["map" => function($f_6) use ($__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_5_4)->{'map'})(function($g_8) use ($f_6) {
  $__num = \func_num_args();
  $__res = function($t_9) use ($f_6, $g_8) {
  $__num = \func_num_args();
  $__res = ($f_6)(($g_8)($t_9));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = function($dictMonoid_6) use ($Comonad0_2_1, $Functor0_4_3, $__local_var_3_2, $dictComonadStore_0, $functorTracedT1_5_4, $pos1_1_0) {
  $__num = \func_num_args();
  $Functor0_7_6 = (((($Comonad0_2_1)->{'Extend0'})(null))->{'Functor0'})(null);
  $lower1_7_6 = function($v_8) use ($Functor0_7_6, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = ((($Functor0_7_6)->{'map'})(function($f_9) use ($dictMonoid_6) {
  $__num = \func_num_args();
  $__res = ($f_9)(($dictMonoid_6)->{'mempty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__local_var_8_8 = (($dictMonoid_6)->{'Semigroup0'})(null);
  $extendTracedT2_8_8 = (object)["extend" => function($f_9) use ($Functor0_4_3, $__local_var_3_2, $__local_var_8_8) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($Functor0_4_3, $__local_var_3_2, $__local_var_8_8, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'extend'})(function($w_prime__11) use ($Functor0_4_3, $__local_var_8_8, $f_9) {
  $__num = \func_num_args();
  $__res = function($t_12) use ($Functor0_4_3, $__local_var_8_8, $f_9, $w_prime__11) {
  $__num = \func_num_args();
  $__res = ($f_9)(((($Functor0_4_3)->{'map'})(function($h_13) use ($__local_var_8_8, $t_12) {
  $__num = \func_num_args();
  $__res = function($t_prime__14) use ($__local_var_8_8, $h_13, $t_12) {
  $__num = \func_num_args();
  $__res = ($h_13)(((($__local_var_8_8)->{'append'})($t_12))($t_prime__14));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($w_prime__11));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_9) use ($functorTracedT1_5_4) {
  $__num = \func_num_args();
  $__res = $functorTracedT1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadTracedT1_8_8 = (object)["extract" => function($v_9) use ($Comonad0_2_1, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = ((($Comonad0_2_1)->{'extract'})($v_9))(($dictMonoid_6)->{'mempty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_9) use ($extendTracedT2_8_8) {
  $__num = \func_num_args();
  $__res = $extendTracedT2_8_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pos" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($pos1_1_0))($lower1_7_6), "peek" => function($s_9) use ($dictComonadStore_0, $lower1_7_6) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadStore_0)->{'peek'})($s_9)))($lower1_7_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar___unused_9) use ($comonadTracedT1_8_8) {
  $__num = \func_num_args();
  $__res = $comonadTracedT1_8_8;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_comonadStoreTracedT'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_comonadmajStoremajTracedmajT';

// Control_Comonad_Store_Class_comonadStoreStoreT
function majControl_majComonad_majStore_majClass_comonadmajStoremajStoremajT($dictComonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_comonadmajStoremajStoremajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonad_0)->{'Extend0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorStoreT1_2_1 = (object)["map" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_2_1)->{'map'})(function($h_5) use ($f_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_3))($h_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value0'}), ($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $extendStoreT1_1_0 = (object)["extend" => function($f_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($__local_var_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'extend'})(function($w_prime__5) use ($f_3) {
  $__num = \func_num_args();
  $__res = function($s_prime__6) use ($f_3, $w_prime__5) {
  $__num = \func_num_args();
  $__res = ($f_3)(new \Data\Tuple\Data_Tuple_Tuple($w_prime__5, $s_prime__6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value0'}), ($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorStoreT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorStoreT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadStoreT_1_0 = (object)["extract" => function($v_2) use ($dictComonad_0) {
  $__num = \func_num_args();
  $__res = ((($dictComonad_0)->{'extract'})(($v_2)->{'value0'}))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_2) use ($extendStoreT1_1_0) {
  $__num = \func_num_args();
  $__res = $extendStoreT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pos" => function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "peek" => function($s_2) use ($dictComonad_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictComonad_0, $s_2) {
  $__num = \func_num_args();
  $__res = ((($dictComonad_0)->{'extract'})(($v_3)->{'value0'}))($s_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar___unused_2) use ($comonadStoreT_1_0) {
  $__num = \func_num_args();
  $__res = $comonadStoreT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_comonadStoreStoreT'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_comonadmajStoremajStoremajT';

// Control_Comonad_Store_Class_comonadStoreEnvT
function majControl_majComonad_majStore_majClass_comonadmajStoremajEnvmajT($dictComonadStore_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majStore_majClass_comonadmajStoremajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Comonad0_1_0 = (($dictComonadStore_0)->{'Comonad0'})(null);
  $__local_var_2_1 = (($Comonad0_1_0)->{'Extend0'})(null);
  $Functor0_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $__local_var_4_3 = (($__local_var_2_1)->{'Functor0'})(null);
  $functorEnvT1_4_3 = (object)["map" => function($f_5) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_3, $f_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_6)->{'value0'}, ((($__local_var_4_3)->{'map'})($f_5))(($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $extendEnvT1_2_1 = (object)["extend" => function($f_5) use ($Functor0_3_2, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($Functor0_3_2, $__local_var_2_1, $f_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_6)->{'value0'}, ((($Functor0_3_2)->{'map'})($f_5))(((($__local_var_2_1)->{'extend'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Tuple_Tuple'])(($v_6)->{'value0'}))))(($v_6)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorEnvT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorEnvT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadEnvT_2_1 = (object)["extract" => function($v_3) use ($Comonad0_1_0) {
  $__num = \func_num_args();
  $__res = (($Comonad0_1_0)->{'extract'})(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_3) use ($extendEnvT1_2_1) {
  $__num = \func_num_args();
  $__res = $extendEnvT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["pos" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictComonadStore_0)->{'pos'}))(function($v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "peek" => function($s_3) use ($dictComonadStore_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadStore_0)->{'peek'})($s_3)))(function($v_4) {
  $__num = \func_num_args();
  $__res = ($v_4)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar___unused_3) use ($comonadEnvT_2_1) {
  $__num = \func_num_args();
  $__res = $comonadEnvT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Store_Class_comonadStoreEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majStore_majClass_comonadmajStoremajEnvmajT';

