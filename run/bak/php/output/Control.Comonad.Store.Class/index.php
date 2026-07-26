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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// Control_Comonad_Store_Class_pos
$GLOBALS['Control_Comonad_Store_Class_pos'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['pos'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Store_Class_peek
$GLOBALS['Control_Comonad_Store_Class_peek'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['peek'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Store_Class_peeks
$GLOBALS['Control_Comonad_Store_Class_peeks'] = (function() {
  $__fn = function($dictComonadStore_0 = null, $f_1 = null, $x_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictComonadStore_0)['peek'])(($f_1)((($dictComonadStore_0)['pos'])($x_2))))($x_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Comonad_Store_Class_seeks
$GLOBALS['Control_Comonad_Store_Class_seeks'] = function($dictComonadStore_0 = null) {
  $__num = \func_num_args();
  $duplicate_1_0 = (((((($dictComonadStore_0)['Comonad0'])(null))['Extend0'])(null))['extend'])($GLOBALS['Control_Extend_identity']);
  $__res = function($f_2 = null) use ($dictComonadStore_0, $duplicate_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_3 = null) use ($dictComonadStore_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($dictComonadStore_0)['peek'])(($f_2)((($dictComonadStore_0)['pos'])($x_3))))($x_3);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Store_Class_seek
$GLOBALS['Control_Comonad_Store_Class_seek'] = function($dictComonadStore_0 = null) {
  $__num = \func_num_args();
  $duplicate_1_0 = (((((($dictComonadStore_0)['Comonad0'])(null))['Extend0'])(null))['extend'])($GLOBALS['Control_Extend_identity']);
  $__res = function($s_2 = null) use ($dictComonadStore_0, $duplicate_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadStore_0)['peek'])($s_2)))($duplicate_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Store_Class_experiment
$GLOBALS['Control_Comonad_Store_Class_experiment'] = (function() {
  $__fn = function($dictComonadStore_0 = null, $dictFunctor_1 = null, $f_2 = null, $x_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictFunctor_1)['map'])(function($a_4 = null) use ($dictComonadStore_0, $x_3) {
  $__num = \func_num_args();
  $__res = ((($dictComonadStore_0)['peek'])($a_4))($x_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)((($dictComonadStore_0)['pos'])($x_3)));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Control_Comonad_Store_Class_comonadStoreTracedT
$GLOBALS['Control_Comonad_Store_Class_comonadStoreTracedT'] = function($dictComonadStore_0 = null) {
  $__num = \func_num_args();
  $pos1_1_0 = ($dictComonadStore_0)['pos'];
  $Comonad0_2_1 = (($dictComonadStore_0)['Comonad0'])(null);
  $comonadTracedT_3_2 = ($GLOBALS['Control_Comonad_Traced_Trans_comonadTracedT'])($Comonad0_2_1);
  $__res = function($dictMonoid_4 = null) use ($Comonad0_2_1, $comonadTracedT_3_2, $dictComonadStore_0, $pos1_1_0) {
  $__num = \func_num_args();
  $mempty_5_3 = ($dictMonoid_4)['mempty'];
  $lower1_6_4 = function($v_6 = null) use ($Comonad0_2_1, $mempty_5_3) {
  $__num = \func_num_args();
  $__res = ((((((($Comonad0_2_1)['Extend0'])(null))['Functor0'])(null))['map'])(function($f_7 = null) use ($mempty_5_3) {
  $__num = \func_num_args();
  $__res = ($f_7)($mempty_5_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $comonadTracedT1_7_5 = ($comonadTracedT_3_2)($dictMonoid_4);
  $__res = ["pos" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($pos1_1_0))($lower1_6_4), "peek" => function($s_8 = null) use ($dictComonadStore_0, $lower1_6_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadStore_0)['peek'])($s_8)))($lower1_6_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar__unused_8 = null) use ($comonadTracedT1_7_5) {
  $__num = \func_num_args();
  $__res = $comonadTracedT1_7_5;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Store_Class_comonadStoreStoreT
$GLOBALS['Control_Comonad_Store_Class_comonadStoreStoreT'] = function($dictComonad_0 = null) {
  $__num = \func_num_args();
  $comonadStoreT_1_0 = ($GLOBALS['Control_Comonad_Store_Trans_comonadStoreT'])($dictComonad_0);
  $__res = ["pos" => function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "peek" => (function() use ($dictComonad_0) {
  $__fn = function($s_2 = null, $v_3 = null) use ($dictComonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictComonad_0)['extract'])(($v_3)->{'value0'}))($s_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Comonad0" => function($_dollar__unused_2 = null) use ($comonadStoreT_1_0) {
  $__num = \func_num_args();
  $__res = $comonadStoreT_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Comonad_Store_Class_comonadStoreEnvT
$GLOBALS['Control_Comonad_Store_Class_comonadStoreEnvT'] = function($dictComonadStore_0 = null) {
  $__num = \func_num_args();
  $Comonad0_1_0 = (($dictComonadStore_0)['Comonad0'])(null);
  $lower1_2_1 = (($GLOBALS['Control_Comonad_Env_Trans_comonadTransEnvT'])['lower'])($Comonad0_1_0);
  $comonadEnvT_3_2 = ($GLOBALS['Control_Comonad_Env_Trans_comonadEnvT'])($Comonad0_1_0);
  $__res = ["pos" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictComonadStore_0)['pos']))($lower1_2_1), "peek" => function($s_4 = null) use ($dictComonadStore_0, $lower1_2_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictComonadStore_0)['peek'])($s_4)))($lower1_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Comonad0" => function($_dollar__unused_4 = null) use ($comonadEnvT_3_2) {
  $__num = \func_num_args();
  $__res = $comonadEnvT_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

