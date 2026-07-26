<?php

namespace Effect\Class\Console;

// ALL IMPORTS: Control.Applicative, Control.Bind, Control.Semigroupoid, Data.Function, Data.Show, Data.Unit, Effect.Class, Effect.Class.Console, Effect.Console, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Control.Semigroupoid, Data.Function, Data.Show, Data.Unit, Effect.Class, Effect.Class.Console, Effect.Console
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
require_once __DIR__ . '/../Effect.Class.Console/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';

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


// Effect_Class_Console_warnShow
$GLOBALS['Effect_Class_Console_warnShow'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $liftEffect_1_0 = ($dictMonadEffect_0)['liftEffect'];
  $__res = function($dictShow_2 = null) use ($liftEffect_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($liftEffect_1_0))(($GLOBALS['Effect_Console_warnShow'])($dictShow_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_warn
$GLOBALS['Effect_Class_Console_warn'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_warn']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_timeLog
$GLOBALS['Effect_Class_Console_timeLog'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_timeLog']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_timeEnd
$GLOBALS['Effect_Class_Console_timeEnd'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_timeEnd']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_time
$GLOBALS['Effect_Class_Console_time'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_time']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_logShow
$GLOBALS['Effect_Class_Console_logShow'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $liftEffect_1_0 = ($dictMonadEffect_0)['liftEffect'];
  $__res = function($dictShow_2 = null) use ($liftEffect_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($liftEffect_1_0))(($GLOBALS['Effect_Console_logShow'])($dictShow_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_log
$GLOBALS['Effect_Class_Console_log'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_log']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_infoShow
$GLOBALS['Effect_Class_Console_infoShow'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $liftEffect_1_0 = ($dictMonadEffect_0)['liftEffect'];
  $__res = function($dictShow_2 = null) use ($liftEffect_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($liftEffect_1_0))(($GLOBALS['Effect_Console_infoShow'])($dictShow_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_info
$GLOBALS['Effect_Class_Console_info'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_info']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_groupEnd
$GLOBALS['Effect_Class_Console_groupEnd'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = (($dictMonadEffect_0)['liftEffect'])($GLOBALS['Effect_Console_groupEnd']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_groupCollapsed
$GLOBALS['Effect_Class_Console_groupCollapsed'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_groupCollapsed']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_group
$GLOBALS['Effect_Class_Console_group'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_group']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_grouped
$GLOBALS['Effect_Class_Console_grouped'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])(null);
  $Bind1_2_1 = (($Monad0_1_0)['Bind1'])(null);
  $discard1_3_2 = (($GLOBALS['Control_Bind_discardUnit'])['discard'])($Bind1_2_1);
  $group1_4_3 = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_group']);
  $groupEnd1_5_4 = (($dictMonadEffect_0)['liftEffect'])($GLOBALS['Effect_Console_groupEnd']);
  $__res = (function() use ($Bind1_2_1, $Monad0_1_0, $discard1_3_2, $group1_4_3, $groupEnd1_5_4) {
  $__fn = function($name_6 = null, $inner_7 = null) use ($Bind1_2_1, $Monad0_1_0, $discard1_3_2, $group1_4_3, $groupEnd1_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($discard1_3_2)(($group1_4_3)($name_6)))(function($_dollar__unused_8 = null) use ($Bind1_2_1, $Monad0_1_0, $discard1_3_2, $groupEnd1_5_4, $inner_7) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)['bind'])($inner_7))(function($result_9 = null) use ($Monad0_1_0, $discard1_3_2, $groupEnd1_5_4) {
  $__num = \func_num_args();
  $__res = (($discard1_3_2)($groupEnd1_5_4))(function($_dollar__unused_10 = null) use ($Monad0_1_0, $result_9) {
  $__num = \func_num_args();
  $__res = (((($Monad0_1_0)['Applicative0'])(null))['pure'])($result_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_errorShow
$GLOBALS['Effect_Class_Console_errorShow'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $liftEffect_1_0 = ($dictMonadEffect_0)['liftEffect'];
  $__res = function($dictShow_2 = null) use ($liftEffect_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($liftEffect_1_0))(($GLOBALS['Effect_Console_errorShow'])($dictShow_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_error
$GLOBALS['Effect_Class_Console_error'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_error']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_debugShow
$GLOBALS['Effect_Class_Console_debugShow'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $liftEffect_1_0 = ($dictMonadEffect_0)['liftEffect'];
  $__res = function($dictShow_2 = null) use ($liftEffect_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($liftEffect_1_0))(($GLOBALS['Effect_Console_debugShow'])($dictShow_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_debug
$GLOBALS['Effect_Class_Console_debug'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])(($dictMonadEffect_0)['liftEffect']))($GLOBALS['Effect_Console_debug']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Effect_Class_Console_clear
$GLOBALS['Effect_Class_Console_clear'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $__res = (($dictMonadEffect_0)['liftEffect'])($GLOBALS['Effect_Console_clear']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

