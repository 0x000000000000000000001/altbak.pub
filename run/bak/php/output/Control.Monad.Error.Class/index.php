<?php

namespace Control\Monad\Error\Class;

// ALL IMPORTS: Control.Applicative, Control.Bind, Control.Monad.Error.Class, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Maybe, Data.Unit, Effect, Effect.Exception, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Control.Monad.Error.Class, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Maybe, Data.Unit, Effect, Effect.Exception, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Exception/index.php';
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


// Control_Monad_Error_Class_throwError
$GLOBALS['Control_Monad_Error_Class_throwError'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['throwError'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Error_Class_monadThrowMaybe
$GLOBALS['Control_Monad_Error_Class_monadThrowMaybe'] = ["throwError" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_monadMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Error_Class_monadThrowEither
$GLOBALS['Control_Monad_Error_Class_monadThrowEither'] = ["throwError" => $GLOBALS['Data_Either_Left'], "Monad0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_monadEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Error_Class_monadThrowEffect
$GLOBALS['Control_Monad_Error_Class_monadThrowEffect'] = ["throwError" => $GLOBALS['Effect_Exception_throwException'], "Monad0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_monadEffect'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Error_Class_monadErrorMaybe
$GLOBALS['Control_Monad_Error_Class_monadErrorMaybe'] = ["catchError" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Nothing"))) {
$__t0 = ($v1_1)($GLOBALS['Data_Unit_unit']);
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Just"))) {
$__t0 = new Phpurs_Data1("Just", ($v_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_Error_Class_monadThrowMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Error_Class_monadErrorEither
$GLOBALS['Control_Monad_Error_Class_monadErrorEither'] = ["catchError" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Left"))) {
$__t0 = ($v1_1)(($v_0)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Right"))) {
$__t0 = new Phpurs_Data1("Right", ($v_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_Error_Class_monadThrowEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Error_Class_monadErrorEffect
$GLOBALS['Control_Monad_Error_Class_monadErrorEffect'] = ["catchError" => (function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Exception_catchException'])($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_Error_Class_monadThrowEffect'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Error_Class_liftMaybe
$GLOBALS['Control_Monad_Error_Class_liftMaybe'] = function($dictMonadThrow_0 = null) {
  $__num = \func_num_args();
  $pure_1_0 = ((((($dictMonadThrow_0)['Monad0'])(null))['Applicative0'])(null))['pure'];
  $__res = function($error_2 = null) use ($dictMonadThrow_0, $pure_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictMonadThrow_0)['throwError'])($error_2);
  $__res = function($v2_4 = null) use ($__local_var_3_1, $pure_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Nothing"))) {
$__t2 = $__local_var_3_1;
goto end_branch_2;;
};
  if ((is_object($v2_4) && (($v2_4)->{'tag'} === "Just"))) {
$__t2 = ($pure_1_0)(($v2_4)->{'value0'});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
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

// Control_Monad_Error_Class_liftEither
$GLOBALS['Control_Monad_Error_Class_liftEither'] = function($dictMonadThrow_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ((((($dictMonadThrow_0)['Monad0'])(null))['Applicative0'])(null))['pure'];
  $__res = function($v2_2 = null) use ($__local_var_1_0, $dictMonadThrow_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Left"))) {
$__t1 = (($dictMonadThrow_0)['throwError'])(($v2_2)->{'value0'});
goto end_branch_1;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Right"))) {
$__t1 = ($__local_var_1_0)(($v2_2)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Error_Class_catchError
$GLOBALS['Control_Monad_Error_Class_catchError'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['catchError'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Error_Class_catchJust
$GLOBALS['Control_Monad_Error_Class_catchJust'] = (function() {
  $__fn = function($dictMonadError_0 = null, $p_1 = null, $act_2 = null, $handler_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictMonadError_0)['catchError'])($act_2))(function($e_4 = null) use ($dictMonadError_0, $handler_3, $p_1) {
  $__num = \func_num_args();
  $v_5_0 = ($p_1)($e_4);
  $__t1 = null;;
  if ((is_object($v_5_0) && (($v_5_0)->{'tag'} === "Nothing"))) {
$__t1 = (((($dictMonadError_0)['MonadThrow0'])(null))['throwError'])($e_4);
goto end_branch_1;;
};
  if ((is_object($v_5_0) && (($v_5_0)->{'tag'} === "Just"))) {
$__t1 = ($handler_3)(($v_5_0)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Error_Class_try
$GLOBALS['Control_Monad_Error_Class_try'] = function($dictMonadError_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (((($dictMonadError_0)['MonadThrow0'])(null))['Monad0'])(null);
  $pure_2_1 = ((($Monad0_1_0)['Applicative0'])(null))['pure'];
  $__res = function($a_3 = null) use ($Monad0_1_0, $dictMonadError_0, $pure_2_1) {
  $__num = \func_num_args();
  $__res = ((($dictMonadError_0)['catchError'])(((((((((($Monad0_1_0)['Bind1'])(null))['Apply0'])(null))['Functor0'])(null))['map'])($GLOBALS['Data_Either_Right']))($a_3)))((($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_2_1))($GLOBALS['Data_Either_Left']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Error_Class_withResource
$GLOBALS['Control_Monad_Error_Class_withResource'] = function($dictMonadError_0 = null) {
  $__num = \func_num_args();
  $MonadThrow0_1_0 = (($dictMonadError_0)['MonadThrow0'])(null);
  $Monad0_2_1 = (($MonadThrow0_1_0)['Monad0'])(null);
  $Bind1_3_2 = (($Monad0_2_1)['Bind1'])(null);
  $try1_4_3 = ($GLOBALS['Control_Monad_Error_Class_try'])($dictMonadError_0);
  $discard1_5_4 = (($GLOBALS['Control_Bind_discardUnit'])['discard'])($Bind1_3_2);
  $__res = (function() use ($Bind1_3_2, $Monad0_2_1, $MonadThrow0_1_0, $discard1_5_4, $try1_4_3) {
  $__fn = function($acquire_6 = null, $release_7 = null, $kleisli_8 = null) use ($Bind1_3_2, $Monad0_2_1, $MonadThrow0_1_0, $discard1_5_4, $try1_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($Bind1_3_2)['bind'])($acquire_6))(function($resource_9 = null) use ($Bind1_3_2, $Monad0_2_1, $MonadThrow0_1_0, $discard1_5_4, $kleisli_8, $release_7, $try1_4_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_3_2)['bind'])(($try1_4_3)(($kleisli_8)($resource_9))))(function($result_10 = null) use ($Monad0_2_1, $MonadThrow0_1_0, $discard1_5_4, $release_7, $resource_9) {
  $__num = \func_num_args();
  $__res = (($discard1_5_4)(($release_7)($resource_9)))(function($_dollar__unused_11 = null) use ($Monad0_2_1, $MonadThrow0_1_0, $result_10) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ((is_object($result_10) && (($result_10)->{'tag'} === "Left"))) {
$__t5 = (($MonadThrow0_1_0)['throwError'])(($result_10)->{'value0'});
goto end_branch_5;;
};
  if ((is_object($result_10) && (($result_10)->{'tag'} === "Right"))) {
$__t5 = (((($Monad0_2_1)['Applicative0'])(null))['pure'])(($result_10)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

