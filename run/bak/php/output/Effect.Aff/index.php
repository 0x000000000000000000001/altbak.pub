<?php

namespace Effect\Aff;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Error.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.ST.Global, Control.Parallel, Control.Parallel.Class, Control.Plus, Control.Semigroupoid, Data.Either, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Time.Duration, Data.Unit, Effect, Effect.Aff, Effect.Class, Effect.Exception, Effect.Unsafe, Partial.Unsafe, Prelude, Prim, Unsafe.Coerce
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Error.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.ST.Global, Control.Parallel, Control.Parallel.Class, Control.Plus, Control.Semigroupoid, Data.Either, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Time.Duration, Data.Unit, Effect, Effect.Aff, Effect.Class, Effect.Exception, Effect.Unsafe, Partial.Unsafe, Prelude, Unsafe.Coerce
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Global/index.php';
require_once __DIR__ . '/../Control.Parallel/index.php';
require_once __DIR__ . '/../Control.Parallel.Class/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Time.Duration/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Aff/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
require_once __DIR__ . '/../Effect.Exception/index.php';
require_once __DIR__ . '/../Effect.Unsafe/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Unsafe.Coerce/index.php';

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
\PhpursThunks::$thunks['Effect_Aff_void'] = function() { $v = ((($GLOBALS['Control_Applicative_liftA1'] ?? \PhpursThunks::eval('Control_Applicative_liftA1')))(($GLOBALS['Effect_applicativeEffect'] ?? \PhpursThunks::eval('Effect_applicativeEffect'))))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Effect_Aff_Canceler'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_unsafeFromRight'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Right"))) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Left"))) {
$__t0 = (($GLOBALS['Partial__crashWith'] ?? \PhpursThunks::eval('Partial__crashWith')))("unsafeFromRight: Left");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_unsafeFromLeft'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Left"))) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Right"))) {
$__t0 = (($GLOBALS['Partial__crashWith'] ?? \PhpursThunks::eval('Partial__crashWith')))("unsafeFromLeft: Right");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_suspendAff'] = function() { $v = (($GLOBALS['Effect_Aff__fork'] ?? \PhpursThunks::eval('Effect_Aff__fork')))(false); return $v; };
\PhpursThunks::$thunks['Effect_Aff_newtypeCanceler'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_isLeft'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Left"))) {
$__t0 = true;
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Right"))) {
$__t0 = false;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_makeAff'] = function() { $v = function($k_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff__makeAff'] ?? \PhpursThunks::eval('Effect_Aff__makeAff')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), $k_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_makeFiber'] = function() { $v = function($aff_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff__makeFiber'] ?? \PhpursThunks::eval('Effect_Aff__makeFiber')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), $aff_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_launchAff'] = function() { $v = function($aff_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Effect_Aff__makeFiber'] ?? \PhpursThunks::eval('Effect_Aff__makeFiber')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), $aff_0);
  $fiber_2_1 = $__local_var_1_0;
  $dollar__unused_3_2 = ($fiber_2_1)->{'run'};
  $__res = $fiber_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_launchAff_'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Aff_void'] ?? \PhpursThunks::eval('Effect_Aff_void'))))(($GLOBALS['Effect_Aff_launchAff'] ?? \PhpursThunks::eval('Effect_Aff_launchAff'))); return $v; };
\PhpursThunks::$thunks['Effect_Aff_launchSuspendedAff'] = function() { $v = ($GLOBALS['Effect_Aff_makeFiber'] ?? \PhpursThunks::eval('Effect_Aff_makeFiber')); return $v; };
\PhpursThunks::$thunks['Effect_Aff_functorParAff'] = function() { $v = (object)["map" => ($GLOBALS['Effect_Aff__parAffMap'] ?? \PhpursThunks::eval('Effect_Aff__parAffMap'))]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_functorAff'] = function() { $v = (object)["map" => ($GLOBALS['Effect_Aff__map'] ?? \PhpursThunks::eval('Effect_Aff__map'))]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_forkAff'] = function() { $v = (($GLOBALS['Effect_Aff__fork'] ?? \PhpursThunks::eval('Effect_Aff__fork')))(true); return $v; };
\PhpursThunks::$thunks['Effect_Aff_delay'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff__delay'] ?? \PhpursThunks::eval('Effect_Aff__delay')))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), $v_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_bracket'] = function() { $v = (function() {
  $__fn = function($acquire_0 = null, $completed_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Effect_Aff_generalBracket'] ?? \PhpursThunks::eval('Effect_Aff_generalBracket')))($acquire_0))((object)["killed" => function($v_2 = null) use ($completed_1) {
  $__num = \func_num_args();
  $__res = $completed_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "failed" => function($v_2 = null) use ($completed_1) {
  $__num = \func_num_args();
  $__res = $completed_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "completed" => function($v_2 = null) use ($completed_1) {
  $__num = \func_num_args();
  $__res = $completed_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_applyParAff'] = function() { $v = (object)["apply" => ($GLOBALS['Effect_Aff__parAffApply'] ?? \PhpursThunks::eval('Effect_Aff__parAffApply')), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_functorParAff'] ?? \PhpursThunks::eval('Effect_Aff_functorParAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_semigroupParAff'] = function() { $v = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["append" => ((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))(($GLOBALS['Effect_Aff_applyParAff'] ?? \PhpursThunks::eval('Effect_Aff_applyParAff'))))(($dictSemigroup_0)->{'append'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_monadAff'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applicativeAff'] ?? \PhpursThunks::eval('Effect_Aff_applicativeAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_bindAff'] ?? \PhpursThunks::eval('Effect_Aff_bindAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_bindAff'] = function() { $v = (object)["bind" => ($GLOBALS['Effect_Aff__bind'] ?? \PhpursThunks::eval('Effect_Aff__bind')), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applyAff'] ?? \PhpursThunks::eval('Effect_Aff_applyAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_applyAff'] = function() { $v = (object)["apply" => (($GLOBALS['Control_Monad_ap'] ?? \PhpursThunks::eval('Control_Monad_ap')))(($GLOBALS['Effect_Aff_monadAff'] ?? \PhpursThunks::eval('Effect_Aff_monadAff'))), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_functorAff'] ?? \PhpursThunks::eval('Effect_Aff_functorAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_applicativeAff'] = function() { $v = (object)["pure" => ($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applyAff'] ?? \PhpursThunks::eval('Effect_Aff_applyAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_cancelWith'] = function() { $v = (function() {
  $__fn = function($aff_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Effect_Aff_generalBracket'] ?? \PhpursThunks::eval('Effect_Aff_generalBracket')))((($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))))((object)["killed" => (function() use ($v_1) {
  $__fn = function($e_2 = null, $v1_3 = null) use ($v_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_1)($e_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "failed" => function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "completed" => function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($v_2 = null) use ($aff_0) {
  $__num = \func_num_args();
  $__res = $aff_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_finally'] = function() { $v = (function() {
  $__fn = function($fin_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Effect_Aff_generalBracket'] ?? \PhpursThunks::eval('Effect_Aff_generalBracket')))((($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))))((object)["killed" => (function() use ($fin_0) {
  $__fn = function($v_2 = null, $v_3 = null) use ($fin_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $fin_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "failed" => (function() use ($fin_0) {
  $__fn = function($v_2 = null, $v_3 = null) use ($fin_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $fin_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "completed" => (function() use ($fin_0) {
  $__fn = function($v_2 = null, $v_3 = null) use ($fin_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $fin_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]))(function($v_2 = null) use ($a_1) {
  $__num = \func_num_args();
  $__res = $a_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_invincible'] = function() { $v = function($a_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  $__res = (((($GLOBALS['Effect_Aff_generalBracket'] ?? \PhpursThunks::eval('Effect_Aff_generalBracket')))($a_0))((object)["killed" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2 = null, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "failed" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2 = null, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "completed" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2 = null, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]))(($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_lazyAff'] = function() { $v = (object)["defer" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_Aff__bind'] ?? \PhpursThunks::eval('Effect_Aff__bind')))((($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))))($f_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_parallelAff'] = function() { $v = (object)["parallel" => ($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')), "sequential" => ($GLOBALS['Effect_Aff__sequential'] ?? \PhpursThunks::eval('Effect_Aff__sequential')), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applyAff'] ?? \PhpursThunks::eval('Effect_Aff_applyAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applyParAff'] ?? \PhpursThunks::eval('Effect_Aff_applyParAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_applicativeParAff'] = function() { $v = (object)["pure" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))))(($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure'))), "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applyParAff'] ?? \PhpursThunks::eval('Effect_Aff_applyParAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_parSequence_'] = function() { $v = (((($GLOBALS['Control_Parallel_parSequence_'] ?? \PhpursThunks::eval('Control_Parallel_parSequence_')))(($GLOBALS['Effect_Aff_parallelAff'] ?? \PhpursThunks::eval('Effect_Aff_parallelAff'))))(($GLOBALS['Effect_Aff_applicativeParAff'] ?? \PhpursThunks::eval('Effect_Aff_applicativeParAff'))))(($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'))); return $v; };
\PhpursThunks::$thunks['Effect_Aff_monoidParAff'] = function() { $v = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["mempty" => ((($GLOBALS['Effect_Aff_applicativeParAff'] ?? \PhpursThunks::eval('Effect_Aff_applicativeParAff')))->{'pure'})(($dictMonoid_0)->{'mempty'}), "Semigroup0" => function($dollar__unused_1 = null) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (object)["append" => ((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))(($GLOBALS['Effect_Aff_applyParAff'] ?? \PhpursThunks::eval('Effect_Aff_applyParAff'))))(((($dictMonoid_0)->{'Semigroup0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'append'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_semigroupCanceler'] = function() { $v = (object)["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $err_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Effect_Aff_parSequence_'] ?? \PhpursThunks::eval('Effect_Aff_parSequence_')))([($v_0)($err_2), ($v1_1)($err_2)]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_semigroupAff'] = function() { $v = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["append" => ((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))(($GLOBALS['Effect_Aff_applyAff'] ?? \PhpursThunks::eval('Effect_Aff_applyAff'))))(($dictSemigroup_0)->{'append'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_monadEffectAff'] = function() { $v = (object)["liftEffect" => ($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect')), "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_monadAff'] ?? \PhpursThunks::eval('Effect_Aff_monadAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_effectCanceler'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Aff_Canceler'] ?? \PhpursThunks::eval('Effect_Aff_Canceler'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))(($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect')))); return $v; };
\PhpursThunks::$thunks['Effect_Aff_joinFiber'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff__makeAff'] ?? \PhpursThunks::eval('Effect_Aff__makeAff')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), function($k_1 = null) use ($v_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Effect_applyEffect'] ?? \PhpursThunks::eval('Effect_applyEffect')))->{'apply'})(($GLOBALS['Effect_Aff_effectCanceler'] ?? \PhpursThunks::eval('Effect_Aff_effectCanceler'))))((($v_0)->{'join'})($k_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_functorFiber'] = function() { $v = (object)["map" => (function() {
  $__fn = function($f_0 = null, $t_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Unsafe_unsafePerformEffect'] ?? \PhpursThunks::eval('Effect_Unsafe_unsafePerformEffect')))((($GLOBALS['Effect_Aff__makeFiber'] ?? \PhpursThunks::eval('Effect_Aff__makeFiber')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), ((($GLOBALS['Effect_Aff__map'] ?? \PhpursThunks::eval('Effect_Aff__map')))($f_0))((($GLOBALS['Effect_Aff_joinFiber'] ?? \PhpursThunks::eval('Effect_Aff_joinFiber')))($t_1))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_applyFiber'] = function() { $v = (object)["apply" => (function() {
  $__fn = function($t1_0 = null, $t2_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Unsafe_unsafePerformEffect'] ?? \PhpursThunks::eval('Effect_Unsafe_unsafePerformEffect')))((($GLOBALS['Effect_Aff__makeFiber'] ?? \PhpursThunks::eval('Effect_Aff__makeFiber')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), (((($GLOBALS['Effect_Aff_applyAff'] ?? \PhpursThunks::eval('Effect_Aff_applyAff')))->{'apply'})((($GLOBALS['Effect_Aff_joinFiber'] ?? \PhpursThunks::eval('Effect_Aff_joinFiber')))($t1_0)))((($GLOBALS['Effect_Aff_joinFiber'] ?? \PhpursThunks::eval('Effect_Aff_joinFiber')))($t2_1))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_functorFiber'] ?? \PhpursThunks::eval('Effect_Aff_functorFiber'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_applicativeFiber'] = function() { $v = (object)["pure" => function($a_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Unsafe_unsafePerformEffect'] ?? \PhpursThunks::eval('Effect_Unsafe_unsafePerformEffect')))((($GLOBALS['Effect_Aff__makeFiber'] ?? \PhpursThunks::eval('Effect_Aff__makeFiber')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), (($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))($a_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applyFiber'] ?? \PhpursThunks::eval('Effect_Aff_applyFiber'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_killFiber'] = function() { $v = (function() {
  $__fn = function($e_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Effect_Aff__bind'] ?? \PhpursThunks::eval('Effect_Aff__bind')))((($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect')))(($v_1)->{'isSuspended'})))(function($suspended_2 = null) use ($e_0, $v_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($suspended_2) {
$__t0 = (($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect')))((((($GLOBALS['Effect_applyEffect'] ?? \PhpursThunks::eval('Effect_applyEffect')))->{'apply'})(function($v_3 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_1)->{'kill'})($e_0, function($v_3 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
goto end_branch_0;;
};
  $__t0 = (($GLOBALS['Effect_Aff__makeAff'] ?? \PhpursThunks::eval('Effect_Aff__makeAff')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), function($k_3 = null) use ($e_0, $v_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Effect_applyEffect'] ?? \PhpursThunks::eval('Effect_applyEffect')))->{'apply'})(($GLOBALS['Effect_Aff_effectCanceler'] ?? \PhpursThunks::eval('Effect_Aff_effectCanceler'))))((($v_1)->{'kill'})($e_0, $k_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_fiberCanceler'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Aff_Canceler'] ?? \PhpursThunks::eval('Effect_Aff_Canceler'))))((function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Effect_Aff_killFiber'] ?? \PhpursThunks::eval('Effect_Aff_killFiber')))($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()); return $v; };
\PhpursThunks::$thunks['Effect_Aff_supervise'] = function() { $v = function($aff_0 = null) {
  $__num = \func_num_args();
  $killError_1_0 = (($GLOBALS['Effect_Exception_error'] ?? \PhpursThunks::eval('Effect_Exception_error')))("[Aff] Child fiber outlived parent");
  $__local_var_2_1 = (($GLOBALS['Effect_Aff__makeSupervisedFiber'] ?? \PhpursThunks::eval('Effect_Aff__makeSupervisedFiber')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), $aff_0);
  $sup_3_2 = $__local_var_2_1;
  $dollar__unused_4_3 = (($sup_3_2)->{'fiber'})->{'run'};
  $__res = (((($GLOBALS['Effect_Aff_generalBracket'] ?? \PhpursThunks::eval('Effect_Aff_generalBracket')))((($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect')))($sup_3_2)))((object)["killed" => (function() {
  $__fn = function($err_2 = null, $sup_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Aff_parSequence_'] ?? \PhpursThunks::eval('Effect_Aff_parSequence_')))([((($GLOBALS['Effect_Aff_killFiber'] ?? \PhpursThunks::eval('Effect_Aff_killFiber')))($err_2))(($sup_3)->{'fiber'}), (($GLOBALS['Effect_Aff__makeAff'] ?? \PhpursThunks::eval('Effect_Aff__makeAff')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), function($k_4 = null) use ($err_2, $sup_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff__killAll'] ?? \PhpursThunks::eval('Effect_Aff__killAll')))($err_2, ($sup_3)->{'supervisor'}, ($k_4)(new Phpurs_Data1("Right", ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})]);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "failed" => (function() use ($killError_1_0) {
  $__fn = function($v_2 = null, $sup_3 = null) use ($killError_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Aff__makeAff'] ?? \PhpursThunks::eval('Effect_Aff__makeAff')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), function($k_4 = null) use ($killError_1_0, $sup_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff__killAll'] ?? \PhpursThunks::eval('Effect_Aff__killAll')))($killError_1_0, ($sup_3)->{'supervisor'}, ($k_4)(new Phpurs_Data1("Right", ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "completed" => (function() use ($killError_1_0) {
  $__fn = function($v_2 = null, $sup_3 = null) use ($killError_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Aff__makeAff'] ?? \PhpursThunks::eval('Effect_Aff__makeAff')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), function($k_4 = null) use ($killError_1_0, $sup_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff__killAll'] ?? \PhpursThunks::eval('Effect_Aff__killAll')))($killError_1_0, ($sup_3)->{'supervisor'}, ($k_4)(new Phpurs_Data1("Right", ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Aff_joinFiber'] ?? \PhpursThunks::eval('Effect_Aff_joinFiber'))))(function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'fiber'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_monadSTAff'] = function() { $v = (object)["liftST" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect'))))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))), "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_monadAff'] ?? \PhpursThunks::eval('Effect_Aff_monadAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_monadThrowAff'] = function() { $v = (object)["throwError" => ($GLOBALS['Effect_Aff__throwError'] ?? \PhpursThunks::eval('Effect_Aff__throwError')), "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_monadAff'] ?? \PhpursThunks::eval('Effect_Aff_monadAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_monadErrorAff'] = function() { $v = (object)["catchError" => ($GLOBALS['Effect_Aff__catchError'] ?? \PhpursThunks::eval('Effect_Aff__catchError')), "MonadThrow0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_monadThrowAff'] ?? \PhpursThunks::eval('Effect_Aff_monadThrowAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_try'] = function() { $v = (($GLOBALS['Control_Monad_Error_Class_try'] ?? \PhpursThunks::eval('Control_Monad_Error_Class_try')))(($GLOBALS['Effect_Aff_monadErrorAff'] ?? \PhpursThunks::eval('Effect_Aff_monadErrorAff'))); return $v; };
\PhpursThunks::$thunks['Effect_Aff_attempt'] = function() { $v = ($GLOBALS['Effect_Aff_try'] ?? \PhpursThunks::eval('Effect_Aff_try')); return $v; };
\PhpursThunks::$thunks['Effect_Aff_runAff'] = function() { $v = (function() {
  $__fn = function($k_0 = null, $aff_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Aff_launchAff'] ?? \PhpursThunks::eval('Effect_Aff_launchAff')))(((($GLOBALS['Effect_Aff__bind'] ?? \PhpursThunks::eval('Effect_Aff__bind')))((($GLOBALS['Effect_Aff_try'] ?? \PhpursThunks::eval('Effect_Aff_try')))($aff_1)))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect'))))($k_0)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_runAff_'] = function() { $v = (function() {
  $__fn = function($k_0 = null, $aff_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Effect_applyEffect'] ?? \PhpursThunks::eval('Effect_applyEffect')))->{'apply'})(function($v_2 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($GLOBALS['Effect_Aff_runAff'] ?? \PhpursThunks::eval('Effect_Aff_runAff')))($k_0))($aff_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_runSuspendedAff'] = function() { $v = (function() {
  $__fn = function($k_0 = null, $aff_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Aff__makeFiber'] ?? \PhpursThunks::eval('Effect_Aff__makeFiber')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), ((($GLOBALS['Effect_Aff__bind'] ?? \PhpursThunks::eval('Effect_Aff__bind')))((($GLOBALS['Effect_Aff_try'] ?? \PhpursThunks::eval('Effect_Aff_try')))($aff_1)))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Aff__liftEffect'] ?? \PhpursThunks::eval('Effect_Aff__liftEffect'))))($k_0)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_monadRecAff'] = function() { $v = (object)["tailRecM" => function($k_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = function($a_2 = null) use (&$go_1_0, $k_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_Aff__bind'] ?? \PhpursThunks::eval('Effect_Aff__bind')))(($k_0)($a_2)))(function($res_3 = null) use (&$go_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($res_3) && (($res_3)->{'tag'} === "Done"))) {
$__t1 = (($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))(($res_3)->{'value0'});
goto end_branch_1;;
};
  if ((is_object($res_3) && (($res_3)->{'tag'} === "Loop"))) {
$__t1 = ($go_1_0)(($res_3)->{'value0'});
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_monadAff'] ?? \PhpursThunks::eval('Effect_Aff_monadAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_monoidAff'] = function() { $v = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["mempty" => (($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))(($dictMonoid_0)->{'mempty'}), "Semigroup0" => function($dollar__unused_1 = null) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (object)["append" => ((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))(($GLOBALS['Effect_Aff_applyAff'] ?? \PhpursThunks::eval('Effect_Aff_applyAff'))))(((($dictMonoid_0)->{'Semigroup0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'append'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Aff_nonCanceler'] = function() { $v = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Effect_Aff__pure'] ?? \PhpursThunks::eval('Effect_Aff__pure')))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
return function($v_1 = null) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__res = $__local_var_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
})(); return $v; };
\PhpursThunks::$thunks['Effect_Aff_monoidCanceler'] = function() { $v = (object)["mempty" => ($GLOBALS['Effect_Aff_nonCanceler'] ?? \PhpursThunks::eval('Effect_Aff_nonCanceler')), "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_semigroupCanceler'] ?? \PhpursThunks::eval('Effect_Aff_semigroupCanceler'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_never'] = function() { $v = (($GLOBALS['Effect_Aff__makeAff'] ?? \PhpursThunks::eval('Effect_Aff__makeAff')))(($GLOBALS['Effect_Aff_isLeft'] ?? \PhpursThunks::eval('Effect_Aff_isLeft')), ($GLOBALS['Effect_Aff_unsafeFromLeft'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromLeft')), ($GLOBALS['Effect_Aff_unsafeFromRight'] ?? \PhpursThunks::eval('Effect_Aff_unsafeFromRight')), ($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')), ($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')), function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_nonCanceler'] ?? \PhpursThunks::eval('Effect_Aff_nonCanceler'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Effect_Aff_apathize'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Effect_Aff__map'] ?? \PhpursThunks::eval('Effect_Aff__map')))(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Effect_Aff_try'] ?? \PhpursThunks::eval('Effect_Aff_try'))); return $v; };
\PhpursThunks::$thunks['Effect_Aff_altParAff'] = function() { $v = (object)["alt" => ($GLOBALS['Effect_Aff__parAffAlt'] ?? \PhpursThunks::eval('Effect_Aff__parAffAlt')), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_functorParAff'] ?? \PhpursThunks::eval('Effect_Aff_functorParAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_altAff'] = function() { $v = (object)["alt" => (function() {
  $__fn = function($a1_0 = null, $a2_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Effect_Aff__catchError'] ?? \PhpursThunks::eval('Effect_Aff__catchError')))($a1_0))(function($v_2 = null) use ($a2_1) {
  $__num = \func_num_args();
  $__res = $a2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_functorAff'] ?? \PhpursThunks::eval('Effect_Aff_functorAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_plusAff'] = function() { $v = (object)["empty" => (($GLOBALS['Effect_Aff__throwError'] ?? \PhpursThunks::eval('Effect_Aff__throwError')))((($GLOBALS['Effect_Exception_error'] ?? \PhpursThunks::eval('Effect_Exception_error')))("Always fails")), "Alt0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_altAff'] ?? \PhpursThunks::eval('Effect_Aff_altAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_plusParAff'] = function() { $v = (object)["empty" => (($GLOBALS['Effect_Aff_plusAff'] ?? \PhpursThunks::eval('Effect_Aff_plusAff')))->{'empty'}, "Alt0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_altParAff'] ?? \PhpursThunks::eval('Effect_Aff_altParAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Effect_Aff_alternativeParAff'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_applicativeParAff'] ?? \PhpursThunks::eval('Effect_Aff_applicativeParAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff_plusParAff'] ?? \PhpursThunks::eval('Effect_Aff_plusParAff'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Effect_Aff = \call_user_func(function() {
  $exports = [];
class PhpursAffBind {
    public $aff;
    public $f;
    public function __construct($aff, $f) {
        $this->aff = $aff;
        $this->f = $f;
    }
}

class PhpursAffMap {
    public $f;
    public $aff;
    public function __construct($f, $aff) {
        $this->f = $f;
        $this->aff = $aff;
    }
}

class PhpursAffCatch {
    public $aff;
    public $f;
    public function __construct($aff, $f) {
        $this->aff = $aff;
        $this->f = $f;
    }
}

class PhpursAffBracket {
    public $acq;
    public $use;
    public function __construct($acq, $use) {
        $this->acq = $acq;
        $this->use = $use;
    }
}

if (!\function_exists('phpursRunAffTrampoline')) {
function phpursRunAffTrampoline($aff) {
    $current = $aff;
    $stack = []; 

    while (true) {
        try {
            if ($current instanceof \Closure) {
                $res = $current();
            } else {
                $res = $current;
            }
            
            if ($res instanceof PhpursAffBind) {
                $stack[] = ['type' => 'bind', 'f' => $res->f];
                $current = $res->aff;
                continue;
            } elseif ($res instanceof PhpursAffMap) {
                $stack[] = ['type' => 'map', 'f' => $res->f];
                $current = $res->aff;
                continue;
            } elseif ($res instanceof PhpursAffCatch) {
                $stack[] = ['type' => 'catch', 'f' => $res->f];
                $current = $res->aff;
                continue;
            } elseif ($res instanceof PhpursAffBracket) {
                $stack[] = ['type' => 'bracket_acq', 'use' => $res->use];
                $current = $res->acq;
                continue;
            }
            
            while (true) {
                if (empty($stack)) {
                    return $res;
                }
                
                $frame = array_pop($stack);
                
                if ($frame['type'] === 'bind') {
                    $f = $frame['f'];
                    $current = $f($res);
                    break;
                } elseif ($frame['type'] === 'map') {
                    $f = $frame['f'];
                    $res = $f($res);
                } elseif ($frame['type'] === 'catch') {
                    // Success value passed through
                } elseif ($frame['type'] === 'bracket_acq') {
                    $use = $frame['use'];
                    $current = $use($res);
                    break;
                }
            }
        } catch (\Throwable $e) {
            if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { 
                echo "\n\n!!! GLOBAL FATAL ERROR CAUGHT IN AFF:\n" . $e->getTraceAsString() . "\n\n"; 
                \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); 
            }
            
            $caught = false;
            while (!empty($stack)) {
                $frame = array_pop($stack);
                if ($frame['type'] === 'catch') {
                    $f = $frame['f'];
                    $current = $f($e);
                    $caught = true;
                    break;
                }
            }
            if (!$caught) {
                throw $e;
            }
        }
    }
}
}

$_pure = function($x) use (&$_pure) { return function() use($x) { return $x; }; };
$_map = function($f, $aff = null) use (&$_map) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_map) {
            return $_map(...\array_merge($__args, $more));
        };
    }
    return function() use($f, $aff) { return new PhpursAffMap($f, $aff); };
};
$_bind = function($aff, $f = null) use (&$_bind) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_bind) {
            return $_bind(...\array_merge($__args, $more));
        };
    }
    return function() use($aff, $f) { return new PhpursAffBind($aff, $f); };
};
$_liftEffect = function($eff) use (&$_liftEffect) { return $eff; };
$_makeFiber = function($isLeft, $unsafeFromLeft, $unsafeFromRight, $Left, $Right, $aff = null) use (&$_makeFiber) { 
    if (\func_num_args() < 6) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_makeFiber) {
            return $_makeFiber(...\array_merge($__args, $more));
        };
    }
    return function() use($aff) { 
        $fiber = new \Fiber(function() use ($aff) { phpursRunAffTrampoline($aff); }); 
        $fiber->start(); 
        return (object)['run' => function() {}, 'join' => function($k) { return function() { return function(){}; }; }]; 
    }; 
};
$_fork = function($immediate, $aff = null) use (&$_fork) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_fork) {
            return $_fork(...\array_merge($__args, $more));
        };
    }
    return function() use($aff) { 
        $fiber = new \Fiber(function() use ($aff) { phpursRunAffTrampoline($aff); }); 
        \Revolt\EventLoop::queue(function() use($fiber) { $fiber->start(); }); 
        return (object)['run' => function() {}, 'join' => function($k){ return function(){ return function(){}; }; }]; 
    };
};
$_delay = function($right, $ms) use (&$_delay) { 
    return function() use($right, $ms) { 
        $fiber = \Fiber::getCurrent(); 
        if ($ms <= 0.0) {
            static $ticks = 0;
            static $lastYield = 0;
            $ticks++;
            
            $shouldYield = false;
            if ($ticks >= 50) {
                $shouldYield = true;
            } elseif ($ticks % 10 === 0) {
                $now = \hrtime(true);
                if ($now - $lastYield > 5000000) { // 5ms in nanoseconds
                    $shouldYield = true;
                }
            }

            if ($shouldYield) {
                $ticks = 0;
                $lastYield = \hrtime(true);
                \Revolt\EventLoop::queue(function() use($fiber) { 
                    if ($fiber) $fiber->resume(); 
                }); 
                if ($fiber) \Fiber::suspend(); 
            }
        } else {
            \Revolt\EventLoop::delay($ms / 1000, function() use($fiber) { 
                if ($fiber) $fiber->resume(); 
            }); 
            if ($fiber) \Fiber::suspend(); 
        }
        return $right(null); 
    }; 
};
$_makeSupervisedFiber = $_makeFiber;
$_killAll = function($err, $sup, $cb) use (&$_killAll) { return function() { return function(){}; }; };

$_makeAff = function($isLeft, $unsafeFromLeft, $unsafeFromRight, $Left, $Right, $k = null) use (&$_makeAff) { 
    if (\func_num_args() < 6) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_makeAff) {
            return $_makeAff(...\array_merge($__args, $more));
        };
    }
    return function() use($k) { 
        $fiber = \Fiber::getCurrent(); 
        $isDone = false;
        $result = null;
        $exception = null;

        $canceler = $k(function($res) use($fiber, &$isDone, &$result, &$exception) { 
            return function() use($fiber, &$isDone, &$result, &$exception, $res) { 
                $isDone = true;
                if (is_object($res) && $res->tag === "Left") {
                    $exception = $res->value0;
                } else {
                    $result = $res->value0;
                }
                
                if ($fiber && $fiber->isSuspended()) { 
                    if ($exception !== null) {
                        \Revolt\EventLoop::queue(function() use($fiber, $exception) {
                            if ($fiber->isSuspended()) $fiber->throw($exception); 
                        });
                    } else {
                        \Revolt\EventLoop::queue(function() use($fiber, $result) {
                            if ($fiber->isSuspended()) $fiber->resume($result); 
                        });
                    }
                } 
            }; 
        })(); 
        
        if (!$isDone) {
            if ($fiber) {
                return \Fiber::suspend(); 
            } else {
                throw new \RuntimeException("makeAff used outside of a fiber");
            }
        } else {
            if ($exception !== null) throw $exception;
            return $result;
        }
    }; 
};

$_throwError = function($err) use (&$_throwError) { return function() use($err) { throw $err; }; };
$_catchError = function($aff, $f = null) use (&$_catchError) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_catchError) {
            return $_catchError(...\array_merge($__args, $more));
        };
    }
    return function() use($aff, $f) { return new PhpursAffCatch($aff, $f); };
};
$generalBracket = function($acq, $cond = null, $use = null) use (&$generalBracket) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$generalBracket) {
            return $generalBracket(...\array_merge($__args, $more));
        };
    }
    return function() use($acq, $use) { return new PhpursAffBracket($acq, $use); }; 
};
$_parAffMap = $_map;

$_parAffApply = function($aff1, $aff2 = null) use (&$_parAffApply) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_parAffApply) {
            return $_parAffApply(...\array_merge($__args, $more));
        };
    }
    return function() use($aff1, $aff2) { 
        $parent = \Fiber::getCurrent();
        $isDone = false; 
        $completed = 0;
        $res1 = null;
        $res2 = null;
        $error = null;

        $f1 = new \Fiber(function() use($aff1, &$isDone, &$completed, &$res1, &$error, $parent) {
            try {
                $res1 = phpursRunAffTrampoline($aff1);
                if (!$isDone) {
                    $completed++;
                    if ($completed === 2) {
                        $isDone = true;
                        if ($parent && $parent->isSuspended()) {
                            \Revolt\EventLoop::queue(function() use($parent) {
                                if ($parent->isSuspended()) $parent->resume();
                            });
                        }
                    }
                }
            } catch (\Throwable $e) { if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { echo "\n\n!!! GLOBAL FATAL ERROR CAUGHT IN AFF:\n" . $e->getTraceAsString() . "\n\n"; } if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                if (!$isDone) {
                    $isDone = true;
                    $error = $e;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $e) {
                            if ($parent->isSuspended()) $parent->throw($e);
                        });
                    }
                }
            }
        });

        $f2 = new \Fiber(function() use($aff2, &$isDone, &$completed, &$res2, &$error, $parent) {
            try {
                $res2 = phpursRunAffTrampoline($aff2);
                if (!$isDone) {
                    $completed++;
                    if ($completed === 2) {
                        $isDone = true;
                        if ($parent && $parent->isSuspended()) {
                            \Revolt\EventLoop::queue(function() use($parent) {
                                if ($parent->isSuspended()) $parent->resume();
                            });
                        }
                    }
                }
            } catch (\Throwable $e) { if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { echo "\n\n!!! GLOBAL FATAL ERROR CAUGHT IN AFF:\n" . $e->getTraceAsString() . "\n\n"; } if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                if (!$isDone) {
                    $isDone = true;
                    $error = $e;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $e) {
                            if ($parent->isSuspended()) $parent->throw($e);
                        });
                    }
                }
            }
        });

        \Revolt\EventLoop::queue(function() use($f1) { $f1->start(); });
        \Revolt\EventLoop::queue(function() use($f2) { $f2->start(); });

        if (!$isDone) {
            \Fiber::suspend();
        }
        
        if ($error !== null) throw $error;
        return $res1($res2); 
    };
};

$_sequential = function($aff) use (&$_sequential) { return $aff; };

$_parAffAlt = function($aff1, $aff2 = null) use (&$_parAffAlt) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_parAffAlt) {
            return $_parAffAlt(...\array_merge($__args, $more));
        };
    }
    return function() use($aff1, $aff2) { 
        $parent = \Fiber::getCurrent();
        $isDone = false;
        $result = null;
        $doneCount = 0;
        $error2 = null;

        $f1 = new \Fiber(function() use($aff1, &$isDone, &$result, &$doneCount, &$error2, $parent) {
            try {
                $res = phpursRunAffTrampoline($aff1);
                if (!$isDone) {
                    $isDone = true;
                    $result = $res;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $result) {
                            if ($parent->isSuspended()) $parent->resume($result);
                        });
                    }
                }
            } catch (\Throwable $e) { if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { echo "\n\n!!! GLOBAL FATAL ERROR CAUGHT IN AFF:\n" . $e->getTraceAsString() . "\n\n"; } if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                $doneCount++;
                if ($doneCount === 2 && !$isDone) {
                    $isDone = true;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $error2) {
                            if ($parent->isSuspended()) $parent->throw($error2); 
                        });
                    }
                }
            }
        });

        $f2 = new \Fiber(function() use($aff2, &$isDone, &$result, &$doneCount, &$error2, $parent) {
            try {
                $res = phpursRunAffTrampoline($aff2);
                if (!$isDone) {
                    $isDone = true;
                    $result = $res;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $result) {
                            if ($parent->isSuspended()) $parent->resume($result);
                        });
                    }
                }
            } catch (\Throwable $e) { if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { echo "\n\n!!! GLOBAL FATAL ERROR CAUGHT IN AFF:\n" . $e->getTraceAsString() . "\n\n"; } if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                $error2 = $e;
                $doneCount++;
                if ($doneCount === 2 && !$isDone) {
                    $isDone = true;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $error2) {
                            if ($parent->isSuspended()) $parent->throw($error2);
                        });
                    }
                }
            }
        });

        \Revolt\EventLoop::queue(function() use($f1) { $f1->start(); });
        \Revolt\EventLoop::queue(function() use($f2) { $f2->start(); });

        if (!$isDone) {
            return \Fiber::suspend();
        } else {
            if ($doneCount === 2) throw $error2;
            return $result;
        }
    };
};

$exports['_pure'] = $_pure;
$exports['_map'] = $_map;
$exports['_bind'] = $_bind;
$exports['_liftEffect'] = $_liftEffect;
$exports['_makeFiber'] = $_makeFiber;
$exports['_fork'] = $_fork;
$exports['_delay'] = $_delay;
$exports['_makeSupervisedFiber'] = $_makeSupervisedFiber;
$exports['_killAll'] = $_killAll;
$exports['_makeAff'] = $_makeAff;
$exports['_throwError'] = $_throwError;
$exports['_catchError'] = $_catchError;
$exports['generalBracket'] = $generalBracket;
$exports['_parAffMap'] = $_parAffMap;
$exports['_parAffApply'] = $_parAffApply;
$exports['_sequential'] = $_sequential;
$exports['_parAffAlt'] = $_parAffAlt;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Effect_Aff__bind'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_bind']; };
\PhpursThunks::$thunks['Effect_Aff__catchError'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_catchError']; };
\PhpursThunks::$thunks['Effect_Aff__delay'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_delay']; };
\PhpursThunks::$thunks['Effect_Aff__fork'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_fork']; };
\PhpursThunks::$thunks['Effect_Aff__killAll'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_killAll']; };
\PhpursThunks::$thunks['Effect_Aff__liftEffect'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_liftEffect']; };
\PhpursThunks::$thunks['Effect_Aff__makeAff'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_makeAff']; };
\PhpursThunks::$thunks['Effect_Aff__makeFiber'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_makeFiber']; };
\PhpursThunks::$thunks['Effect_Aff__makeSupervisedFiber'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_makeSupervisedFiber']; };
\PhpursThunks::$thunks['Effect_Aff__map'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_map']; };
\PhpursThunks::$thunks['Effect_Aff__parAffAlt'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_parAffAlt']; };
\PhpursThunks::$thunks['Effect_Aff__parAffApply'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_parAffApply']; };
\PhpursThunks::$thunks['Effect_Aff__parAffMap'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_parAffMap']; };
\PhpursThunks::$thunks['Effect_Aff__pure'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_pure']; };
\PhpursThunks::$thunks['Effect_Aff__sequential'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_sequential']; };
\PhpursThunks::$thunks['Effect_Aff__throwError'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['_throwError']; };
\PhpursThunks::$thunks['Effect_Aff_generalBracket'] = function() use (&$ffi_Effect_Aff) { return $ffi_Effect_Aff['generalBracket']; };































































