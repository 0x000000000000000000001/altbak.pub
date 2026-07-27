<?php

namespace Control\Comonad\Env\Trans;

// ALL IMPORTS: Control.Comonad, Control.Comonad.Env.Trans, Control.Comonad.Trans.Class, Control.Extend, Control.Semigroupoid, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.Newtype, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Comonad, Control.Comonad.Env.Trans, Control.Comonad.Trans.Class, Control.Extend, Control.Semigroupoid, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.Newtype, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Comonad.Env.Trans/index.php';
require_once __DIR__ . '/../Control.Comonad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };




// Control_Comonad_Env_Trans_EnvT
function majControl_majComonad_majEnv_majTrans_majEnvmajT($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_majEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_EnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_majEnvmajT';

// Control_Comonad_Env_Trans_withEnvT
function majControl_majComonad_majEnv_majTrans_withmajEnvmajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_withmajEnvmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($f_0)(($v_1)->{'value0'}), ($v_1)->{'value1'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_withEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_withmajEnvmajT';

// Control_Comonad_Env_Trans_runEnvT
function majControl_majComonad_majEnv_majTrans_runmajEnvmajT($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_runmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_runEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_runmajEnvmajT';

// Control_Comonad_Env_Trans_newtypeEnvT
$GLOBALS['Control_Comonad_Env_Trans_newtypeEnvT'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Comonad_Env_Trans_mapEnvT
function majControl_majComonad_majEnv_majTrans_mapmajEnvmajT($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_mapmajEnvmajT';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, ($f_0)(($v_1)->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_mapEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_mapmajEnvmajT';

// Control_Comonad_Env_Trans_functorEnvT
function majControl_majComonad_majEnv_majTrans_functormajEnvmajT($dictFunctor_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_functormajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1, $v_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_2)->{'value0'}, ((($dictFunctor_0)->{'map'})($f_1))(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_functorEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_functormajEnvmajT';

// Control_Comonad_Env_Trans_functorWithIndexEnvT
function majControl_majComonad_majEnv_majTrans_functormajWithmajIndexmajEnvmajT($dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_functormajWithmajIndexmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFunctorWithIndex_0)->{'Functor0'})(null);
  $functorEnvT1_2_1 = (object)["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_3)->{'value0'}, ((($__local_var_1_0)->{'map'})($f_2))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["mapWithIndex" => (function() use ($dictFunctorWithIndex_0) {
  $__fn = function($f_3, $v_4 = null) use ($dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_4)->{'value0'}, ((($dictFunctorWithIndex_0)->{'mapWithIndex'})($f_3))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_3) use ($functorEnvT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorEnvT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_functorWithIndexEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_functormajWithmajIndexmajEnvmajT';

// Control_Comonad_Env_Trans_foldableEnvT
function majControl_majComonad_majEnv_majTrans_foldablemajEnvmajT($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_foldablemajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["foldl" => (function() use ($dictFoldable_0) {
  $__fn = function($fn_1, $a_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)->{'foldl'})($fn_1))($a_2))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() use ($dictFoldable_0) {
  $__fn = function($fn_1, $a_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)->{'foldr'})($fn_1))($a_2))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $foldMap1_2_0 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_1);
  $__res = (function() use ($foldMap1_2_0) {
  $__fn = function($fn_3, $v_4 = null) use ($foldMap1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap1_2_0)($fn_3))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_foldableEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_foldablemajEnvmajT';

// Control_Comonad_Env_Trans_foldableWithIndexEnvT
function majControl_majComonad_majEnv_majTrans_foldablemajWithmajIndexmajEnvmajT($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_foldablemajWithmajIndexmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictFoldableWithIndex_0)->{'Foldable0'})(null);
  $foldableEnvT1_2_1 = (object)["foldl" => (function() use ($__local_var_1_0) {
  $__fn = function($fn_2, $a_3 = null, $v_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_1_0)->{'foldl'})($fn_2))($a_3))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() use ($__local_var_1_0) {
  $__fn = function($fn_2, $a_3 = null, $v_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_1_0)->{'foldr'})($fn_2))($a_3))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $foldMap1_3_1 = (($__local_var_1_0)->{'foldMap'})($dictMonoid_2);
  $__res = (function() use ($foldMap1_3_1) {
  $__fn = function($fn_4, $v_5 = null) use ($foldMap1_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap1_3_1)($fn_4))(($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["foldlWithIndex" => (function() use ($dictFoldableWithIndex_0) {
  $__fn = function($f_3, $a_4 = null, $v_5 = null) use ($dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)->{'foldlWithIndex'})($f_3))($a_4))(($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldrWithIndex" => (function() use ($dictFoldableWithIndex_0) {
  $__fn = function($f_3, $a_4 = null, $v_5 = null) use ($dictFoldableWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldableWithIndex_0)->{'foldrWithIndex'})($f_3))($a_4))(($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_3) use ($dictFoldableWithIndex_0) {
  $__num = \func_num_args();
  $foldMapWithIndex1_4_3 = (($dictFoldableWithIndex_0)->{'foldMapWithIndex'})($dictMonoid_3);
  $__res = (function() use ($foldMapWithIndex1_4_3) {
  $__fn = function($f_5, $v_6 = null) use ($foldMapWithIndex1_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMapWithIndex1_4_3)($f_5))(($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_3) use ($foldableEnvT1_2_1) {
  $__num = \func_num_args();
  $__res = $foldableEnvT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_foldableWithIndexEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_foldablemajWithmajIndexmajEnvmajT';

// Control_Comonad_Env_Trans_traversableEnvT
function majControl_majComonad_majEnv_majTrans_traversablemajEnvmajT($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_traversablemajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $functorEnvT1_2_1 = (object)["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_3)->{'value0'}, ((($__local_var_1_0)->{'map'})($f_2))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_3_2 = (($dictTraversable_0)->{'Foldable1'})(null);
  $foldableEnvT1_4_3 = (object)["foldl" => (function() use ($__local_var_3_2) {
  $__fn = function($fn_4, $a_5 = null, $v_6 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_3_2)->{'foldl'})($fn_4))($a_5))(($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() use ($__local_var_3_2) {
  $__fn = function($fn_4, $a_5 = null, $v_6 = null) use ($__local_var_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_3_2)->{'foldr'})($fn_4))($a_5))(($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_4) use ($__local_var_3_2) {
  $__num = \func_num_args();
  $foldMap1_5_3 = (($__local_var_3_2)->{'foldMap'})($dictMonoid_4);
  $__res = (function() use ($foldMap1_5_3) {
  $__fn = function($fn_6, $v_7 = null) use ($foldMap1_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap1_5_3)($fn_6))(($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["sequence" => function($dictApplicative_5) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $sequence1_6_5 = (($dictTraversable_0)->{'sequence'})($dictApplicative_5);
  $__res = function($v_7) use ($dictApplicative_5, $sequence1_6_5) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((($GLOBALS['Data_Functor_functorFn'])->{'map'})($GLOBALS['Control_Comonad_Env_Trans_EnvT']))(($GLOBALS['Data_Tuple_Tuple'])(($v_7)->{'value0'}))))(($sequence1_6_5)(($v_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "traverse" => function($dictApplicative_5) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $traverse1_6_6 = (($dictTraversable_0)->{'traverse'})($dictApplicative_5);
  $__res = (function() use ($dictApplicative_5, $traverse1_6_6) {
  $__fn = function($f_7, $v_8 = null) use ($dictApplicative_5, $traverse1_6_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_5)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((($GLOBALS['Data_Functor_functorFn'])->{'map'})($GLOBALS['Control_Comonad_Env_Trans_EnvT']))(($GLOBALS['Data_Tuple_Tuple'])(($v_8)->{'value0'}))))((($traverse1_6_6)($f_7))(($v_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_5) use ($functorEnvT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorEnvT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_5) use ($foldableEnvT1_4_3) {
  $__num = \func_num_args();
  $__res = $foldableEnvT1_4_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_traversableEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_traversablemajEnvmajT';

// Control_Comonad_Env_Trans_traversableWithIndexEnvT
function majControl_majComonad_majEnv_majTrans_traversablemajWithmajIndexmajEnvmajT($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_traversablemajWithmajIndexmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversableWithIndex_0)->{'FunctorWithIndex0'})(null);
  $__local_var_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $functorEnvT1_3_2 = (object)["map" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3, $v_4 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_4)->{'value0'}, ((($__local_var_2_1)->{'map'})($f_3))(($v_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $functorWithIndexEnvT1_3_2 = (object)["mapWithIndex" => (function() use ($__local_var_1_0) {
  $__fn = function($f_4, $v_5 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_5)->{'value0'}, ((($__local_var_1_0)->{'mapWithIndex'})($f_4))(($v_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4) use ($functorEnvT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorEnvT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $foldableWithIndexEnvT1_4_4 = \Control\Comonad\Env\Trans\majControl_majComonad_majEnv_majTrans_foldablemajWithmajIndexmajEnvmajT((($dictTraversableWithIndex_0)->{'FoldableWithIndex1'})(null));
  $traversableEnvT1_5_5 = \Control\Comonad\Env\Trans\majControl_majComonad_majEnv_majTrans_traversablemajEnvmajT((($dictTraversableWithIndex_0)->{'Traversable2'})(null));
  $__res = (object)["traverseWithIndex" => function($dictApplicative_6) use ($dictTraversableWithIndex_0) {
  $__num = \func_num_args();
  $traverseWithIndex1_7_6 = (($dictTraversableWithIndex_0)->{'traverseWithIndex'})($dictApplicative_6);
  $__res = (function() use ($dictApplicative_6, $traverseWithIndex1_7_6) {
  $__fn = function($f_8, $v_9 = null) use ($dictApplicative_6, $traverseWithIndex1_7_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(((($GLOBALS['Data_Functor_functorFn'])->{'map'})($GLOBALS['Control_Comonad_Env_Trans_EnvT']))(($GLOBALS['Data_Tuple_Tuple'])(($v_9)->{'value0'}))))((($traverseWithIndex1_7_6)($f_8))(($v_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_6) use ($functorWithIndexEnvT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorWithIndexEnvT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_6) use ($foldableWithIndexEnvT1_4_4) {
  $__num = \func_num_args();
  $__res = $foldableWithIndexEnvT1_4_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_6) use ($traversableEnvT1_5_5) {
  $__num = \func_num_args();
  $__res = $traversableEnvT1_5_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_traversableWithIndexEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_traversablemajWithmajIndexmajEnvmajT';

// Control_Comonad_Env_Trans_extendEnvT
function majControl_majComonad_majEnv_majTrans_extendmajEnvmajT($dictExtend_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_extendmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictExtend_0)->{'Functor0'})(null);
  $functorEnvT1_2_1 = (object)["map" => (function() use ($Functor0_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($Functor0_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_3)->{'value0'}, ((($Functor0_1_0)->{'map'})($f_2))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["extend" => (function() use ($Functor0_1_0, $dictExtend_0) {
  $__fn = function($f_3, $v_4 = null) use ($Functor0_1_0, $dictExtend_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_4)->{'value0'}, ((($Functor0_1_0)->{'map'})($f_3))(((($dictExtend_0)->{'extend'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Comonad_Env_Trans_EnvT']))(($GLOBALS['Data_Tuple_Tuple'])(($v_4)->{'value0'}))))(($v_4)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_3) use ($functorEnvT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorEnvT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_extendEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_extendmajEnvmajT';

// Control_Comonad_Env_Trans_comonadTransEnvT
$GLOBALS['Control_Comonad_Env_Trans_comonadTransEnvT'] = (object)["lower" => (function() {
  $__fn = function($dictComonad_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_1)->{'value1'};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Control_Comonad_Env_Trans_comonadEnvT
function majControl_majComonad_majEnv_majTrans_comonadmajEnvmajT($dictComonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majComonad_majEnv_majTrans_comonadmajEnvmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $extendEnvT1_1_0 = \Control\Comonad\Env\Trans\majControl_majComonad_majEnv_majTrans_extendmajEnvmajT((($dictComonad_0)->{'Extend0'})(null));
  $__res = (object)["extract" => function($v_2) use ($dictComonad_0) {
  $__num = \func_num_args();
  $__res = (($dictComonad_0)->{'extract'})(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar__unused_2) use ($extendEnvT1_1_0) {
  $__num = \func_num_args();
  $__res = $extendEnvT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Comonad_Env_Trans_comonadEnvT'] = __NAMESPACE__ . '\\majControl_majComonad_majEnv_majTrans_comonadmajEnvmajT';

