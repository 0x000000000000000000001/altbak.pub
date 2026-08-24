<?php

namespace Control\Comonad\Traced\Class;

// ALL IMPORTS: Control.Comonad, Control.Comonad.Env, Control.Comonad.Env.Trans, Control.Comonad.Store, Control.Comonad.Store.Trans, Control.Comonad.Traced.Class, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Monad.Identity.Trans, Control.Semigroupoid, Data.Function, Data.Functor, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Comonad, Control.Comonad.Env, Control.Comonad.Env.Trans, Control.Comonad.Store, Control.Comonad.Store.Trans, Control.Comonad.Traced.Class, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Monad.Identity.Trans, Control.Semigroupoid, Data.Function, Data.Functor, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Comonad.Env/index.php';
require_once __DIR__ . '/../Control.Comonad.Env.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Store/index.php';
require_once __DIR__ . '/../Control.Comonad.Store.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Traced.Class/index.php';
require_once __DIR__ . '/../Control.Comonad.Traced.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Identity.Trans/index.php';
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




// Control_Comonad_Traced_Class_track
function majControl_majComonad_majTraced_majClass_track($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_track';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'track'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Traced_Class_track'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_track';

// Control_Comonad_Traced_Class_tracks
function majControl_majComonad_majTraced_majClass_tracks($dictComonadTraced_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_tracks';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Comonad0_1_0 = (($dictComonadTraced_0)->{'Comonad0'})(null);
  $__res = function($f_2) use ($Comonad0_1_0, $dictComonadTraced_0) {
  $__num = \func_num_args();
  $__res = function($w_3) use ($Comonad0_1_0, $dictComonadTraced_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictComonadTraced_0)->{'track'})(($f_2)((($Comonad0_1_0)->{'extract'})($w_3))))($w_3);
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
$GLOBALS['Control_Comonad_Traced_Class_tracks'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_tracks';

// Control_Comonad_Traced_Class_listens
function majControl_majComonad_majTraced_majClass_listens($dictFunctor_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_listens';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})(function($g_3) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($t_4) use ($f_1, $g_3) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($g_3)($t_4), ($f_1)($t_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Comonad_Traced_Class_listens'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_listens';

// Control_Comonad_Traced_Class_listen
function majControl_majComonad_majTraced_majClass_listen($dictFunctor_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_listen';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)->{'map'})(function($f_2) {
  $__num = \func_num_args();
  $__res = function($t_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_2)($t_3), $t_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Comonad_Traced_Class_listen'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_listen';

// Control_Comonad_Traced_Class_comonadTracedTracedT
function majControl_majComonad_majTraced_majClass_comonadmajTracedmajTracedmajT($dictComonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_comonadmajTracedmajTracedmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonad_0)->{'Extend0'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorTracedT1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'map'})(function($g_6) use ($f_4) {
  $__num = \func_num_args();
  $__res = function($t_7) use ($f_4, $g_6) {
  $__num = \func_num_args();
  $__res = ($f_4)(($g_6)($t_7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadTracedT_3_2 = function($dictMonoid_4) use ($Functor0_2_1, $__local_var_1_0, $dictComonad_0, $functorTracedT1_3_2) {
  $__num = \func_num_args();
  $__local_var_5_4 = (($dictMonoid_4)->{'Semigroup0'})(null);
  $extendTracedT2_5_4 = (object)["extend" => function($f_6) use ($Functor0_2_1, $__local_var_1_0, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Functor0_2_1, $__local_var_1_0, $__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'extend'})(function($w_prime__8) use ($Functor0_2_1, $__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($t_9) use ($Functor0_2_1, $__local_var_5_4, $f_6, $w_prime__8) {
  $__num = \func_num_args();
  $__res = ($f_6)(((($Functor0_2_1)->{'map'})(function($h_10) use ($__local_var_5_4, $t_9) {
  $__num = \func_num_args();
  $__res = function($t_prime__11) use ($__local_var_5_4, $h_10, $t_9) {
  $__num = \func_num_args();
  $__res = ($h_10)(((($__local_var_5_4)->{'append'})($t_9))($t_prime__11));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($w_prime__8));
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
}, "Functor0" => function($_dollar___unused_6) use ($functorTracedT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorTracedT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["extract" => function($v_6) use ($dictComonad_0, $dictMonoid_4) {
  $__num = \func_num_args();
  $__res = ((($dictComonad_0)->{'extract'})($v_6))(($dictMonoid_4)->{'mempty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_6) use ($extendTracedT2_5_4) {
  $__num = \func_num_args();
  $__res = $extendTracedT2_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_4) use ($comonadTracedT_3_2, $dictComonad_0) {
  $__num = \func_num_args();
  $comonadTracedT1_5_7 = ($comonadTracedT_3_2)($dictMonoid_4);
  $__res = (object)["track" => function($t_6) use ($dictComonad_0) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($dictComonad_0, $t_6) {
  $__num = \func_num_args();
  $__res = ((($dictComonad_0)->{'extract'})($v_7))($t_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar___unused_6) use ($comonadTracedT1_5_7) {
  $__num = \func_num_args();
  $__res = $comonadTracedT1_5_7;
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
$GLOBALS['Control_Comonad_Traced_Class_comonadTracedTracedT'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_comonadmajTracedmajTracedmajT';

// Control_Comonad_Traced_Class_comonadTracedStoreT
function majControl_majComonad_majTraced_majClass_comonadmajTracedmajStoremajT($dictComonadTraced_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_comonadmajTracedmajStoremajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonadTraced_0)->{'Comonad0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Extend0'})(null);
  $__local_var_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $functorStoreT1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_3_2)->{'map'})(function($h_6) use ($f_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))($h_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)->{'value0'}), ($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $extendStoreT1_2_1 = (object)["extend" => function($f_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_1, $f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_2_1)->{'extend'})(function($w_prime__6) use ($f_4) {
  $__num = \func_num_args();
  $__res = function($s_prime__7) use ($f_4, $w_prime__6) {
  $__num = \func_num_args();
  $__res = ($f_4)(new \Data\Tuple\Data_Tuple_Tuple($w_prime__6, $s_prime__7));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)->{'value0'}), ($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorStoreT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorStoreT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadStoreT_1_0 = (object)["extract" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'extract'})(($v_3)->{'value0'}))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_3) use ($extendStoreT1_2_1) {
  $__num = \func_num_args();
  $__res = $extendStoreT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Functor0_2_6 = (((((($dictComonadTraced_0)->{'Comonad0'})(null))->{'Extend0'})(null))->{'Functor0'})(null);
  $lower1_2_6 = function($v_3) use ($Functor0_2_6) {
  $__num = \func_num_args();
  $__local_var_4_7 = ($v_3)->{'value1'};
  $__res = ((($Functor0_2_6)->{'map'})(function($v1_5) use ($__local_var_4_7) {
  $__num = \func_num_args();
  $__res = ($v1_5)($__local_var_4_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_3)->{'value0'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (object)["track" => function($m_3) use ($dictComonadTraced_0, $lower1_2_6) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadTraced_0)->{'track'})($m_3)))($lower1_2_6);
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
$GLOBALS['Control_Comonad_Traced_Class_comonadTracedStoreT'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_comonadmajTracedmajStoremajT';

// Control_Comonad_Traced_Class_comonadTracedIdentityT
function majControl_majComonad_majTraced_majClass_comonadmajTracedmajIdentitymajT($dictComonadTraced_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_comonadmajTracedmajIdentitymajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonadTraced_0)->{'Comonad0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Extend0'})(null);
  $functorIdentityT1_3_2 = (($__local_var_2_1)->{'Functor0'})(null);
  $extendIdentityI1_2_1 = (object)["extend" => function($f_4) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_1, $f_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'extend'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_4))(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($v_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorIdentityT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorIdentityT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadIdentityT_1_0 = (object)["extract" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($__local_var_1_0)->{'extract'}))($GLOBALS['Control_Monad_Identity_Trans_runIdentityT']), "Extend0" => function($_dollar___unused_3) use ($extendIdentityI1_2_1) {
  $__num = \func_num_args();
  $__res = $extendIdentityI1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["track" => function($m_2) use ($dictComonadTraced_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadTraced_0)->{'track'})($m_2)))($GLOBALS['Control_Monad_Identity_Trans_runIdentityT']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar___unused_2) use ($comonadIdentityT_1_0) {
  $__num = \func_num_args();
  $__res = $comonadIdentityT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Traced_Class_comonadTracedIdentityT'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_comonadmajTracedmajIdentitymajT';

// Control_Comonad_Traced_Class_comonadTracedEnvT
function majControl_majComonad_majTraced_majClass_comonadmajTracedmajEnvmajT($dictComonadTraced_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_comonadmajTracedmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonadTraced_0)->{'Comonad0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Extend0'})(null);
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
  $comonadEnvT_1_0 = (object)["extract" => function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)->{'extract'})(($v_3)->{'value1'});
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
  $__res = (object)["track" => function($m_2) use ($dictComonadTraced_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadTraced_0)->{'track'})($m_2)))(function($v_3) {
  $__num = \func_num_args();
  $__res = ($v_3)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar___unused_2) use ($comonadEnvT_1_0) {
  $__num = \func_num_args();
  $__res = $comonadEnvT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Traced_Class_comonadTracedEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_comonadmajTracedmajEnvmajT';

// Control_Comonad_Traced_Class_censor
function majControl_majComonad_majTraced_majClass_censor($dictFunctor_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majTraced_majClass_censor';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)->{'map'})(function($v1_3) use ($f_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v1_3))($f_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Comonad_Traced_Class_censor'] = __NAMESPACE__ . '\\majControl_majComonad_majTraced_majClass_censor';

