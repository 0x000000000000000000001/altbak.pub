<?php

namespace Control\Comonad\Env\Class;

// ALL IMPORTS: Control.Comonad, Control.Comonad.Env.Class, Control.Comonad.Env.Trans, Control.Comonad.Store, Control.Comonad.Store.Trans, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Semigroupoid, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Comonad, Control.Comonad.Env.Class, Control.Comonad.Env.Trans, Control.Comonad.Store, Control.Comonad.Store.Trans, Control.Comonad.Traced.Trans, Control.Comonad.Trans.Class, Control.Semigroupoid, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Comonad.Env.Class/index.php';
require_once __DIR__ . '/../Control.Comonad.Env.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Store/index.php';
require_once __DIR__ . '/../Control.Comonad.Store.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Traced.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
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




// Control_Comonad_Env_Class_local
function majControl_majComonad_majEnv_majClass_local($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_local';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'local'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Class_local'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_local';

// Control_Comonad_Env_Class_comonadAskTuple
$GLOBALS['Control_Comonad_Env_Class_comonadAskTuple'] = (object)["ask" => $GLOBALS['Data_Tuple_fst'], "Comonad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_comonadTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Comonad_Env_Class_comonadEnvTuple
$GLOBALS['Control_Comonad_Env_Class_comonadEnvTuple'] = (object)["local" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_0)(($v_1)->{'value0'}), ($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "ComonadAsk0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Comonad_Env_Class_comonadAskTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Comonad_Env_Class_comonadAskEnvT
function majControl_majComonad_majEnv_majClass_comonadmajAskmajEnvmajT($dictComonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_comonadmajAskmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonad_0)->{'Extend0'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorEnvT1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_5)->{'value0'}, ((($__local_var_3_2)->{'map'})($f_4))(($v_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $extendEnvT1_1_0 = (object)["extend" => function($f_4) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Functor0_2_1, $__local_var_1_0, $f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_5)->{'value0'}, ((($Functor0_2_1)->{'map'})($f_4))(((($__local_var_1_0)->{'extend'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Tuple_Tuple'])(($v_5)->{'value0'}))))(($v_5)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorEnvT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorEnvT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadEnvT_1_0 = (object)["extract" => function($v_2) use ($dictComonad_0) {
  $__num = \func_num_args();
  $__res = (($dictComonad_0)->{'extract'})(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_2) use ($extendEnvT1_1_0) {
  $__num = \func_num_args();
  $__res = $extendEnvT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["ask" => function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'value0'};
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
$GLOBALS['Control_Comonad_Env_Class_comonadAskEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_comonadmajAskmajEnvmajT';

// Control_Comonad_Env_Class_comonadEnvEnvT
function majControl_majComonad_majEnv_majClass_comonadmajEnvmajEnvmajT($dictComonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_comonadmajEnvmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonad_0)->{'Extend0'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $__local_var_3_2 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorEnvT1_3_2 = (object)["map" => function($f_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_3_2, $f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_5)->{'value0'}, ((($__local_var_3_2)->{'map'})($f_4))(($v_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $extendEnvT1_1_0 = (object)["extend" => function($f_4) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Functor0_2_1, $__local_var_1_0, $f_4) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_5)->{'value0'}, ((($Functor0_2_1)->{'map'})($f_4))(((($__local_var_1_0)->{'extend'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Tuple_Tuple'])(($v_5)->{'value0'}))))(($v_5)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_4) use ($functorEnvT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorEnvT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadEnvT_1_0 = (object)["extract" => function($v_2) use ($dictComonad_0) {
  $__num = \func_num_args();
  $__res = (($dictComonad_0)->{'extract'})(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_2) use ($extendEnvT1_1_0) {
  $__num = \func_num_args();
  $__res = $extendEnvT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadAskEnvT1_1_0 = (object)["ask" => function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'value0'};
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
  $__res = (object)["local" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_2)(($v_3)->{'value0'}), ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "ComonadAsk0" => function($_dollar___unused_2) use ($comonadAskEnvT1_1_0) {
  $__num = \func_num_args();
  $__res = $comonadAskEnvT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Class_comonadEnvEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_comonadmajEnvmajEnvmajT';

// Control_Comonad_Env_Class_ask
function majControl_majComonad_majEnv_majClass_ask($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_ask';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'ask'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Class_ask'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_ask';

// Control_Comonad_Env_Class_asks
function majControl_majComonad_majEnv_majClass_asks($dictComonadAsk_0, $f_1 = null, $x_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_asks';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)((($dictComonadAsk_0)->{'ask'})($x_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Class_asks'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_asks';

// Control_Comonad_Env_Class_comonadAskStoreT
function majControl_majComonad_majEnv_majClass_comonadmajAskmajStoremajT($dictComonadAsk_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_comonadmajAskmajStoremajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Comonad0_1_0 = (($dictComonadAsk_0)->{'Comonad0'})(null);
  $__local_var_2_1 = (($Comonad0_1_0)->{'Extend0'})(null);
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
  $comonadStoreT_2_1 = (object)["extract" => function($v_3) use ($Comonad0_1_0) {
  $__num = \func_num_args();
  $__res = ((($Comonad0_1_0)->{'extract'})(($v_3)->{'value0'}))(($v_3)->{'value1'});
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
  $Functor0_3_6 = (((($Comonad0_1_0)->{'Extend0'})(null))->{'Functor0'})(null);
  $__res = (object)["ask" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictComonadAsk_0)->{'ask'}))(function($v_4) use ($Functor0_3_6) {
  $__num = \func_num_args();
  $__local_var_5_7 = ($v_4)->{'value1'};
  $__res = ((($Functor0_3_6)->{'map'})(function($v1_6) use ($__local_var_5_7) {
  $__num = \func_num_args();
  $__res = ($v1_6)($__local_var_5_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->{'value0'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Comonad0" => function($_dollar___unused_3) use ($comonadStoreT_2_1) {
  $__num = \func_num_args();
  $__res = $comonadStoreT_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Class_comonadAskStoreT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_comonadmajAskmajStoremajT';

// Control_Comonad_Env_Class_comonadEnvStoreT
function majControl_majComonad_majEnv_majClass_comonadmajEnvmajStoremajT($dictComonadEnv_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_comonadmajEnvmajStoremajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonadEnv_0)->{'ComonadAsk0'})(null);
  $Comonad0_2_1 = (($__local_var_1_0)->{'Comonad0'})(null);
  $__local_var_3_2 = (($Comonad0_2_1)->{'Extend0'})(null);
  $__local_var_4_3 = (($__local_var_3_2)->{'Functor0'})(null);
  $functorStoreT1_4_3 = (object)["map" => function($f_5) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_4_3, $f_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_4_3)->{'map'})(function($h_7) use ($f_5) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_5))($h_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_6)->{'value0'}), ($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $extendStoreT1_3_2 = (object)["extend" => function($f_5) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_3_2, $f_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_3_2)->{'extend'})(function($w_prime__7) use ($f_5) {
  $__num = \func_num_args();
  $__res = function($s_prime__8) use ($f_5, $w_prime__7) {
  $__num = \func_num_args();
  $__res = ($f_5)(new \Data\Tuple\Data_Tuple_Tuple($w_prime__7, $s_prime__8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_6)->{'value0'}), ($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_5) use ($functorStoreT1_4_3) {
  $__num = \func_num_args();
  $__res = $functorStoreT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadStoreT_3_2 = (object)["extract" => function($v_4) use ($Comonad0_2_1) {
  $__num = \func_num_args();
  $__res = ((($Comonad0_2_1)->{'extract'})(($v_4)->{'value0'}))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_4) use ($extendStoreT1_3_2) {
  $__num = \func_num_args();
  $__res = $extendStoreT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Functor0_4_7 = (((($Comonad0_2_1)->{'Extend0'})(null))->{'Functor0'})(null);
  $comonadAskStoreT1_1_0 = (object)["ask" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($__local_var_1_0)->{'ask'}))(function($v_5) use ($Functor0_4_7) {
  $__num = \func_num_args();
  $__local_var_6_8 = ($v_5)->{'value1'};
  $__res = ((($Functor0_4_7)->{'map'})(function($v1_7) use ($__local_var_6_8) {
  $__num = \func_num_args();
  $__res = ($v1_7)($__local_var_6_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_5)->{'value0'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Comonad0" => function($_dollar___unused_4) use ($comonadStoreT_3_2) {
  $__num = \func_num_args();
  $__res = $comonadStoreT_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["local" => function($f_2) use ($dictComonadEnv_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictComonadEnv_0, $f_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictComonadEnv_0)->{'local'})($f_2))(($v_3)->{'value0'}), ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "ComonadAsk0" => function($_dollar___unused_2) use ($comonadAskStoreT1_1_0) {
  $__num = \func_num_args();
  $__res = $comonadAskStoreT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Class_comonadEnvStoreT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_comonadmajEnvmajStoremajT';

// Control_Comonad_Env_Class_comonadAskTracedT
function majControl_majComonad_majEnv_majClass_comonadmajAskmajTracedmajT($dictComonadAsk_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_comonadmajAskmajTracedmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ask1_1_0 = ($dictComonadAsk_0)->{'ask'};
  $Comonad0_2_1 = (($dictComonadAsk_0)->{'Comonad0'})(null);
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
  $__res = function($dictMonoid_6) use ($Comonad0_2_1, $Functor0_4_3, $__local_var_3_2, $ask1_1_0, $functorTracedT1_5_4) {
  $__num = \func_num_args();
  $__local_var_7_6 = (($dictMonoid_6)->{'Semigroup0'})(null);
  $extendTracedT2_7_6 = (object)["extend" => function($f_8) use ($Functor0_4_3, $__local_var_3_2, $__local_var_7_6) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($Functor0_4_3, $__local_var_3_2, $__local_var_7_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)->{'extend'})(function($w_prime__10) use ($Functor0_4_3, $__local_var_7_6, $f_8) {
  $__num = \func_num_args();
  $__res = function($t_11) use ($Functor0_4_3, $__local_var_7_6, $f_8, $w_prime__10) {
  $__num = \func_num_args();
  $__res = ($f_8)(((($Functor0_4_3)->{'map'})(function($h_12) use ($__local_var_7_6, $t_11) {
  $__num = \func_num_args();
  $__res = function($t_prime__13) use ($__local_var_7_6, $h_12, $t_11) {
  $__num = \func_num_args();
  $__res = ($h_12)(((($__local_var_7_6)->{'append'})($t_11))($t_prime__13));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($w_prime__10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_8) use ($functorTracedT1_5_4) {
  $__num = \func_num_args();
  $__res = $functorTracedT1_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadTracedT1_7_6 = (object)["extract" => function($v_8) use ($Comonad0_2_1, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = ((($Comonad0_2_1)->{'extract'})($v_8))(($dictMonoid_6)->{'mempty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_8) use ($extendTracedT2_7_6) {
  $__num = \func_num_args();
  $__res = $extendTracedT2_7_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Functor0_8_9 = (((($Comonad0_2_1)->{'Extend0'})(null))->{'Functor0'})(null);
  $__res = (object)["ask" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($ask1_1_0))(function($v_9) use ($Functor0_8_9, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = ((($Functor0_8_9)->{'map'})(function($f_10) use ($dictMonoid_6) {
  $__num = \func_num_args();
  $__res = ($f_10)(($dictMonoid_6)->{'mempty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Comonad0" => function($_dollar___unused_8) use ($comonadTracedT1_7_6) {
  $__num = \func_num_args();
  $__res = $comonadTracedT1_7_6;
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
$GLOBALS['Control_Comonad_Env_Class_comonadAskTracedT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_comonadmajAskmajTracedmajT';

// Control_Comonad_Env_Class_comonadEnvTracedT
function majControl_majComonad_majEnv_majClass_comonadmajEnvmajTracedmajT($dictComonadEnv_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majClass_comonadmajEnvmajTracedmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictComonadEnv_0)->{'ComonadAsk0'})(null);
  $ask1_2_1 = ($__local_var_1_0)->{'ask'};
  $Comonad0_3_2 = (($__local_var_1_0)->{'Comonad0'})(null);
  $__local_var_4_3 = (($Comonad0_3_2)->{'Extend0'})(null);
  $Functor0_5_4 = (($__local_var_4_3)->{'Functor0'})(null);
  $__local_var_6_5 = (($__local_var_4_3)->{'Functor0'})(null);
  $functorTracedT1_6_5 = (object)["map" => function($f_7) use ($__local_var_6_5) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_6_5, $f_7) {
  $__num = \func_num_args();
  $__res = ((($__local_var_6_5)->{'map'})(function($g_9) use ($f_7) {
  $__num = \func_num_args();
  $__res = function($t_10) use ($f_7, $g_9) {
  $__num = \func_num_args();
  $__res = ($f_7)(($g_9)($t_10));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadAskTracedT1_3_2 = function($dictMonoid_7) use ($Comonad0_3_2, $Functor0_5_4, $__local_var_4_3, $ask1_2_1, $functorTracedT1_6_5) {
  $__num = \func_num_args();
  $__local_var_8_7 = (($dictMonoid_7)->{'Semigroup0'})(null);
  $extendTracedT2_8_7 = (object)["extend" => function($f_9) use ($Functor0_5_4, $__local_var_4_3, $__local_var_8_7) {
  $__num = \func_num_args();
  $__res = function($v_10) use ($Functor0_5_4, $__local_var_4_3, $__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_4_3)->{'extend'})(function($w_prime__11) use ($Functor0_5_4, $__local_var_8_7, $f_9) {
  $__num = \func_num_args();
  $__res = function($t_12) use ($Functor0_5_4, $__local_var_8_7, $f_9, $w_prime__11) {
  $__num = \func_num_args();
  $__res = ($f_9)(((($Functor0_5_4)->{'map'})(function($h_13) use ($__local_var_8_7, $t_12) {
  $__num = \func_num_args();
  $__res = function($t_prime__14) use ($__local_var_8_7, $h_13, $t_12) {
  $__num = \func_num_args();
  $__res = ($h_13)(((($__local_var_8_7)->{'append'})($t_12))($t_prime__14));
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
}, "Functor0" => function($_dollar___unused_9) use ($functorTracedT1_6_5) {
  $__num = \func_num_args();
  $__res = $functorTracedT1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $comonadTracedT1_8_7 = (object)["extract" => function($v_9) use ($Comonad0_3_2, $dictMonoid_7) {
  $__num = \func_num_args();
  $__res = ((($Comonad0_3_2)->{'extract'})($v_9))(($dictMonoid_7)->{'mempty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_9) use ($extendTracedT2_8_7) {
  $__num = \func_num_args();
  $__res = $extendTracedT2_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Functor0_9_10 = (((($Comonad0_3_2)->{'Extend0'})(null))->{'Functor0'})(null);
  $__res = (object)["ask" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($ask1_2_1))(function($v_10) use ($Functor0_9_10, $dictMonoid_7) {
  $__num = \func_num_args();
  $__res = ((($Functor0_9_10)->{'map'})(function($f_11) use ($dictMonoid_7) {
  $__num = \func_num_args();
  $__res = ($f_11)(($dictMonoid_7)->{'mempty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Comonad0" => function($_dollar___unused_9) use ($comonadTracedT1_8_7) {
  $__num = \func_num_args();
  $__res = $comonadTracedT1_8_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictMonoid_4) use ($comonadAskTracedT1_3_2, $dictComonadEnv_0) {
  $__num = \func_num_args();
  $comonadAskTracedT2_5_12 = ($comonadAskTracedT1_3_2)($dictMonoid_4);
  $__res = (object)["local" => function($f_6) use ($dictComonadEnv_0) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($dictComonadEnv_0, $f_6) {
  $__num = \func_num_args();
  $__res = ((($dictComonadEnv_0)->{'local'})($f_6))($v_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "ComonadAsk0" => function($_dollar___unused_6) use ($comonadAskTracedT2_5_12) {
  $__num = \func_num_args();
  $__res = $comonadAskTracedT2_5_12;
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
$GLOBALS['Control_Comonad_Env_Class_comonadEnvTracedT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majClass_comonadmajEnvmajTracedmajT';

