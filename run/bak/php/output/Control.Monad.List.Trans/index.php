<?php

namespace Control\Monad\List\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.List.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.Trans.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Lazy, Data.Maybe, Data.Monoid, Data.Newtype, Data.Ring, Data.Semigroup, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Data.Unit, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.List.Trans, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.Trans.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Lazy, Data.Maybe, Data.Monoid, Data.Newtype, Data.Ring, Data.Semigroup, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Data.Unit, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.List.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
require_once __DIR__ . '/../Prelude/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
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
\PhpursThunks::$thunks['Control_Monad_List_Trans_Yield'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Yield", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_Skip'] = function() { $v = function($value0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Skip", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_Done'] = function() { $v = ($GLOBALS['__phpurs_data0_Done'] ??= new Phpurs_Data0("Done")); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_ListT'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_ListT"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_wrapLazy'] = function() { $v = (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_wrapLazy"), recVars=[];
  $__res = (($dictApplicative_0)->pure)(new Phpurs_Data1("Skip", $v_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_wrapEffect'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_wrapEffect"), recVars=[];
  $__res = ((($dictFunctor_0)->map)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_List_Trans_Skip'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_Skip'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_unfold'] = function() { $v = (function() {
  $__fn = function($dictMonad_0, $f_1 = null, $z_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_unfold"), recVars=["Control_Monad_List_Trans_unfold"];
  while (true) {
$__res = ((((((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(function($v_3) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_unfold"];
  if ((is_object($v_3) && (($v_3)->tag === "Just"))) {
$__local_var_4_1 = (($v_3)->value0)->value0;
$__t0 = new Phpurs_Data2("Yield", (($v_3)->value0)->value1, (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v1_5) use ($__local_var_4_1, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_unfold"];
  $__res = (((($GLOBALS['Control_Monad_List_Trans_unfold'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_unfold')))($dictMonad_0))($f_1))($__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Nothing"))) {
$__t0 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_1)($z_2));
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_uncons'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_uncons"), recVars=["Control_Monad_List_Trans_uncons"];
  while (true) {
$__local_var_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = function($v_2) use ($__local_var_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_uncons"];
  $__res = ((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)($v_2))(function($v1_3) use ($__local_var_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_uncons"];
  if ((is_object($v1_3) && (($v1_3)->tag === "Yield"))) {
$__t1 = (($__local_var_1_0)->pure)(new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($v1_3)->value0, (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v1_3)->value1))));
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Skip"))) {
$__t1 = ((($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($dictMonad_0))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v1_3)->value0));
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Done"))) {
$__t1 = (($__local_var_1_0)->pure)(new Phpurs_Data0("Nothing"));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_tail'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_tail"), recVars=[];
  $uncons1_1_0 = (($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($dictMonad_0);
  $__res = function($l_2) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(function($v1_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($v1_3)->value0)->value1);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($uncons1_1_0)($l_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_takeWhile'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_takeWhile"), recVars=["Control_Monad_List_Trans_takeWhile"];
  while (true) {
$__local_var_1_0 = (((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($f_2, $v_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_takeWhile"];
  $__res = ((($__local_var_1_0)->map)(function($v_4) use ($dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_takeWhile"];
  if ((is_object($v_4) && (($v_4)->tag === "Yield"))) {
if (($f_2)(($v_4)->value0)) {
$__local_var_5_3 = ((($GLOBALS['Control_Monad_List_Trans_takeWhile'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_takeWhile')))($dictApplicative_0))($f_2);
$__t2 = new Phpurs_Data2("Yield", ($v_4)->value0, (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_5_3, $v_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_takeWhile"];
  $__res = ($__local_var_5_3)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v_4)->value1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
$__t2 = new Phpurs_Data0("Done");
};
$__t1 = $__t2;
} else {
if ((is_object($v_4) && (($v_4)->tag === "Skip"))) {
$__local_var_5_4 = ($v_4)->value0;
$__local_var_6_5 = ((($GLOBALS['Control_Monad_List_Trans_takeWhile'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_takeWhile')))($dictApplicative_0))($f_2);
$__t1 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7) use ($__local_var_5_4, $__local_var_6_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_takeWhile"];
  $__res = ($__local_var_6_5)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_4) && (($v_4)->tag === "Done"))) {
$__t1 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_scanl'] = function() { $v = (function() {
  $__fn = function($dictMonad_0, $f_1 = null, $b_2 = null, $l_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_scanl"), recVars=[];
  $__res = (((($GLOBALS['Control_Monad_List_Trans_unfold'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_unfold')))($dictMonad_0))(function($v_4) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_5_0 = ($v_4)->value0;
  $__res = ((((((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(function($v1_6) use ($__local_var_5_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_6) && (($v1_6)->tag === "Yield"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", (($f_1)($__local_var_5_0))(($v1_6)->value0), (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v1_6)->value1)), $__local_var_5_0));
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Skip"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", $__local_var_5_0, (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v1_6)->value0)), $__local_var_5_0));
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Done"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)->value1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", $b_2, $l_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_prepend__prime__'] = function() { $v = (function() {
  $__fn = function($dictApplicative_0, $h_1 = null, $t_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_prepend'"), recVars=[];
  $__res = (($dictApplicative_0)->pure)(new Phpurs_Data2("Yield", $h_1, $t_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_prepend'] = function() { $v = (function() {
  $__fn = function($dictApplicative_0, $h_1 = null, $t_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_prepend"), recVars=[];
  $__res = (($dictApplicative_0)->pure)(new Phpurs_Data2("Yield", $h_1, (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use ($t_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $t_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_nil'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_nil"), recVars=[];
  $__res = (($dictApplicative_0)->pure)(new Phpurs_Data0("Done"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_singleton'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_singleton"), recVars=[];
  $nil1_1_0 = (($dictApplicative_0)->pure)(new Phpurs_Data0("Done"));
  $__res = function($a_2) use ($dictApplicative_0, $nil1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictApplicative_0)->pure)(new Phpurs_Data2("Yield", $a_2, (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use ($nil1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $nil1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_take'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_take"), recVars=["Control_Monad_List_Trans_take"];
  while (true) {
$nil1_1_0 = (($dictApplicative_0)->pure)(new Phpurs_Data0("Done"));
$__local_var_2_1 = (((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_2_1, $dictApplicative_0, $nil1_1_0) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_2_1, $dictApplicative_0, $nil1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_take"];
  switch ($v_3) {
case 0:
$__t7 = $nil1_1_0;
break;
default:
$__t7 = ((($__local_var_2_1)->map)(function($v2_5) use ($dictApplicative_0, $v_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_take"];
  if ((is_object($v2_5) && (($v2_5)->tag === "Yield"))) {
$__local_var_6_3 = ($v2_5)->value1;
$__local_var_7_4 = ((($GLOBALS['Control_Monad_List_Trans_take'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_take')))($dictApplicative_0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_3))(1));
$__t2 = new Phpurs_Data2("Yield", ($v2_5)->value0, (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_8) use ($__local_var_6_3, $__local_var_7_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_take"];
  $__res = ($__local_var_7_4)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_6_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v2_5) && (($v2_5)->tag === "Skip"))) {
$__local_var_6_5 = ($v2_5)->value0;
$__local_var_7_6 = ((($GLOBALS['Control_Monad_List_Trans_take'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_take')))($dictApplicative_0))($v_3);
$__t2 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_8) use ($__local_var_6_5, $__local_var_7_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_take"];
  $__res = ($__local_var_7_6)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_6_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v2_5) && (($v2_5)->tag === "Done"))) {
$__t2 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v1_4);
break;
};
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_zipWith__prime__'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_zipWith'"), recVars=["Control_Monad_List_Trans_zipWith'"];
  while (true) {
$Applicative0_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$nil1_2_1 = (($Applicative0_1_0)->pure)(new Phpurs_Data0("Done"));
$Bind1_3_2 = (($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$Functor0_4_3 = (((($Bind1_3_2)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$uncons1_5_4 = (($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($dictMonad_0);
$__res = (function() use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $nil1_2_1, $uncons1_5_4) {
  $__fn = function($f_6, $fa_7 = null, $fb_8 = null) use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $nil1_2_1, $uncons1_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_zipWith'"];
  $__res = ((($Functor0_4_3)->map)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_List_Trans_Skip'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_Skip'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer'))))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))))(((($Bind1_3_2)->bind)(($uncons1_5_4)($fa_7)))(function($ua_9) use ($Applicative0_1_0, $Bind1_3_2, $Functor0_4_3, $dictMonad_0, $f_6, $fb_8, $nil1_2_1, $uncons1_5_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_zipWith'"];
  $__res = ((($Bind1_3_2)->bind)(($uncons1_5_4)($fb_8)))(function($ub_10) use ($Applicative0_1_0, $Functor0_4_3, $dictMonad_0, $f_6, $nil1_2_1, $ua_9) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_zipWith'"];
  if ((is_object($ub_10) && (($ub_10)->tag === "Nothing"))) {
$__t5 = (($Applicative0_1_0)->pure)($nil1_2_1);
} else {
if ((is_object($ua_9) && (($ua_9)->tag === "Nothing"))) {
$__t5 = (($Applicative0_1_0)->pure)($nil1_2_1);
} else {
if (((is_object($ua_9) && (($ua_9)->tag === "Just")) && (is_object($ub_10) && (($ub_10)->tag === "Just")))) {
$__local_var_11_6 = (($ua_9)->value0)->value1;
$__local_var_12_7 = (($ub_10)->value0)->value1;
$__local_var_13_8 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v2_13) use ($__local_var_11_6, $__local_var_12_7, $dictMonad_0, $f_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_zipWith'"];
  $__res = ((((($GLOBALS['Control_Monad_List_Trans_zipWith__prime__'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_zipWith__prime__')))($dictMonad_0))($f_6))($__local_var_11_6))($__local_var_12_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t5 = ((($Functor0_4_3)->map)(function($a_14) use ($Applicative0_1_0, $__local_var_13_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_zipWith'"];
  $__res = (($Applicative0_1_0)->pure)(new Phpurs_Data2("Yield", $a_14, $__local_var_13_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($f_6)((($ua_9)->value0)->value0))((($ub_10)->value0)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
};
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_zipWith'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_zipWith"), recVars=[];
  $zipWith__prime__1_1_0 = (($GLOBALS['Control_Monad_List_Trans_zipWith__prime__'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_zipWith__prime__')))($dictMonad_0);
  $__res = function($f_2) use ($dictMonad_0, $zipWith__prime__1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($zipWith__prime__1_1_0)((function() use ($dictMonad_0, $f_2) {
  $__fn = function($a_3, $b_4 = null) use ($dictMonad_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)((($f_2)($a_3))($b_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_newtypeListT'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_mapMaybe'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_mapMaybe"), recVars=["Control_Monad_List_Trans_mapMaybe"];
  while (true) {
$__res = ((($dictFunctor_0)->map)(function($v_3) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_mapMaybe"];
  if ((is_object($v_3) && (($v_3)->tag === "Yield"))) {
$__local_var_4_1 = ($f_1)(($v_3)->value0);
if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Just"))) {
$__local_var_5_4 = ((($GLOBALS['Control_Monad_List_Trans_mapMaybe'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_mapMaybe')))($dictFunctor_0))($f_1);
$__t3 = new Phpurs_Data2("Yield", ($__local_var_4_1)->value0, (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_5_4, $v_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_mapMaybe"];
  $__res = ($__local_var_5_4)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v_3)->value1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
$__local_var_5_2 = ((($GLOBALS['Control_Monad_List_Trans_mapMaybe'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_mapMaybe')))($dictFunctor_0))($f_1);
$__t3 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_5_2, $v_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_mapMaybe"];
  $__res = ($__local_var_5_2)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v_3)->value1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
};
$__t0 = $__t3;
} else {
if ((is_object($v_3) && (($v_3)->tag === "Skip"))) {
$__local_var_4_5 = ($v_3)->value0;
$__local_var_5_6 = ((($GLOBALS['Control_Monad_List_Trans_mapMaybe'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_mapMaybe')))($dictFunctor_0))($f_1);
$__t0 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_4_5, $__local_var_5_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_mapMaybe"];
  $__res = ($__local_var_5_6)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_4_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Done"))) {
$__t0 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_iterate'] = function() { $v = (function() {
  $__fn = function($dictMonad_0, $f_1 = null, $a_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_iterate"), recVars=[];
  $__res = (((($GLOBALS['Control_Monad_List_Trans_unfold'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_unfold')))($dictMonad_0))(function($x_3) use ($dictMonad_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($f_1)($x_3), $x_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_repeat'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_repeat"), recVars=[];
  $__res = ((($GLOBALS['Control_Monad_List_Trans_iterate'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_iterate')))($dictMonad_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_head'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_head"), recVars=[];
  $uncons1_1_0 = (($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($dictMonad_0);
  $__res = function($l_2) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(function($v1_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($v1_3)->value0)->value0);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($uncons1_1_0)($l_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_functorListT'] = function() { $v = function($dictFunctor_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_functorListT"), recVars=["Control_Monad_List_Trans_functorListT"];
  while (true) {
$__res = (object)["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1, $v_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_functorListT"];
  $__res = ((($dictFunctor_0)->map)(function($v_3) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_functorListT"];
  if ((is_object($v_3) && (($v_3)->tag === "Yield"))) {
$__local_var_4_1 = ($v_3)->value1;
$__local_var_5_2 = (((($GLOBALS['Control_Monad_List_Trans_functorListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_functorListT')))($dictFunctor_0))->map)($f_1);
$__t0 = new Phpurs_Data2("Yield", ($f_1)(($v_3)->value0), (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_4_1, $__local_var_5_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_functorListT"];
  $__res = ($__local_var_5_2)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_4_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Skip"))) {
$__local_var_4_3 = ($v_3)->value0;
$__local_var_5_4 = (((($GLOBALS['Control_Monad_List_Trans_functorListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_functorListT')))($dictFunctor_0))->map)($f_1);
$__t0 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_4_3, $__local_var_5_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_functorListT"];
  $__res = ($__local_var_5_4)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_4_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Done"))) {
$__t0 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_fromEffect'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_fromEffect"), recVars=[];
  $nil1_1_0 = (($dictApplicative_0)->pure)(new Phpurs_Data0("Done"));
  $__res = function($fa_2) use ($dictApplicative_0, $nil1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_3_1 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use ($nil1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $nil1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = ((((((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(function($a_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Yield", $a_4, $__local_var_3_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($fa_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_monadTransListT'] = function() { $v = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_fromEffect'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_fromEffect')))((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_foldlRec__prime__'] = function() { $v = function($dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_foldlRec'"), recVars=[];
  $Monad0_1_0 = (($dictMonadRec_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_3_2 = (($Monad0_1_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $uncons1_4_3 = (($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($Monad0_1_0);
  $__res = (function() use ($__local_var_2_1, $__local_var_3_2, $dictMonadRec_0, $uncons1_4_3) {
  $__fn = function($f_5, $a_6 = null, $b_7 = null) use ($__local_var_2_1, $__local_var_3_2, $dictMonadRec_0, $uncons1_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictMonadRec_0)->tailRecM)(function($o_8) use ($__local_var_2_1, $__local_var_3_2, $f_5, $uncons1_4_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_9_4 = ($o_8)->a;
  $__res = ((($__local_var_3_2)->bind)(($uncons1_4_3)(($o_8)->b)))(function($v_10) use ($__local_var_2_1, $__local_var_3_2, $__local_var_9_4, $f_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_10) && (($v_10)->tag === "Nothing"))) {
$__t5 = (($__local_var_2_1)->pure)(new Phpurs_Data1("Done", $__local_var_9_4));
} else {
if ((is_object($v_10) && (($v_10)->tag === "Just"))) {
$__local_var_11_6 = (($v_10)->value0)->value1;
$__t5 = ((($__local_var_3_2)->bind)((($f_5)($__local_var_9_4))((($v_10)->value0)->value0)))(function($b__prime___12) use ($__local_var_11_6, $__local_var_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($__local_var_2_1)->pure)(new Phpurs_Data1("Loop", (object)["a" => $b__prime___12, "b" => $__local_var_11_6]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_6, "b" => $b_7]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_runListTRec'] = function() { $v = function($dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_runListTRec"), recVars=[];
  $__res = (((($GLOBALS['Control_Monad_List_Trans_foldlRec__prime__'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_foldlRec__prime__')))($dictMonadRec_0))((function() use ($dictMonadRec_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((((($dictMonadRec_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_foldlRec'] = function() { $v = function($dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_foldlRec"), recVars=[];
  $Monad0_1_0 = (($dictMonadRec_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $uncons1_3_2 = (($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($Monad0_1_0);
  $__res = (function() use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, $uncons1_3_2) {
  $__fn = function($f_4, $a_5 = null, $b_6 = null) use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, $uncons1_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictMonadRec_0)->tailRecM)(function($o_7) use ($Monad0_1_0, $__local_var_2_1, $f_4, $uncons1_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_8_3 = ($o_7)->a;
  $__res = ((((($Monad0_1_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)(($uncons1_3_2)(($o_7)->b)))(function($v_9) use ($__local_var_2_1, $__local_var_8_3, $f_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_9) && (($v_9)->tag === "Nothing"))) {
$__t4 = (($__local_var_2_1)->pure)(new Phpurs_Data1("Done", $__local_var_8_3));
} else {
if ((is_object($v_9) && (($v_9)->tag === "Just"))) {
$__t4 = (($__local_var_2_1)->pure)(new Phpurs_Data1("Loop", (object)["a" => (($f_4)($__local_var_8_3))((($v_9)->value0)->value0), "b" => (($v_9)->value0)->value1]));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_5, "b" => $b_6]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_foldl__prime__'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_foldl'"), recVars=[];
  $__local_var_1_0 = (($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $uncons1_2_1 = (($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($dictMonad_0);
  $__res = function($f_3) use ($__local_var_1_0, $dictMonad_0, $uncons1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $loop_4_2 = null;
  $loop_4_2 = (function() use ($__local_var_1_0, $dictMonad_0, $f_3, &$loop_4_2, $uncons1_2_1) {
  $__fn = function($b_5, $l_6 = null) use ($__local_var_1_0, $dictMonad_0, $f_3, &$loop_4_2, $uncons1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "loop_4_2"), recVars=["loop_4_2"];
  while (true) {
$__res = ((($__local_var_1_0)->bind)(($uncons1_2_1)($l_6)))(function($v_7) use ($__local_var_1_0, $b_5, $dictMonad_0, $f_3, &$loop_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["loop_4_2"];
  if ((is_object($v_7) && (($v_7)->tag === "Nothing"))) {
$__t3 = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($b_5);
} else {
if ((is_object($v_7) && (($v_7)->tag === "Just"))) {
$__local_var_8_4 = (($v_7)->value0)->value1;
$__t3 = ((($__local_var_1_0)->bind)((($f_3)($b_5))((($v_7)->value0)->value0)))(function($a_9) use ($__local_var_8_4, &$loop_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["loop_4_2"];
  $__res = (($loop_4_2)($a_9))($__local_var_8_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = $loop_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_runListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_runListT"), recVars=[];
  $__res = (((($GLOBALS['Control_Monad_List_Trans_foldl__prime__'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_foldl__prime__')))($dictMonad_0))((function() use ($dictMonad_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_foldl'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_foldl"), recVars=[];
  $uncons1_1_0 = (($GLOBALS['Control_Monad_List_Trans_uncons'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_uncons')))($dictMonad_0);
  $__res = function($f_2) use ($dictMonad_0, $uncons1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $loop_3_1 = null;
  $loop_3_1 = (function() use ($dictMonad_0, $f_2, &$loop_3_1, $uncons1_1_0) {
  $__fn = function($b_4, $l_5 = null) use ($dictMonad_0, $f_2, &$loop_3_1, $uncons1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "loop_3_1"), recVars=["loop_3_1"];
  while (true) {
$__res = ((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)(($uncons1_1_0)($l_5)))(function($v_6) use ($b_4, $dictMonad_0, $f_2, &$loop_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["loop_3_1"];
  if ((is_object($v_6) && (($v_6)->tag === "Nothing"))) {
$__t2 = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($b_4);
} else {
if ((is_object($v_6) && (($v_6)->tag === "Just"))) {
$__t2 = (($loop_3_1)((($f_2)($b_4))((($v_6)->value0)->value0)))((($v_6)->value0)->value1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = $loop_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_filter'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_filter"), recVars=["Control_Monad_List_Trans_filter"];
  while (true) {
$__res = ((($dictFunctor_0)->map)(function($v_3) use ($dictFunctor_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_filter"];
  if ((is_object($v_3) && (($v_3)->tag === "Yield"))) {
$__local_var_4_1 = ($v_3)->value1;
$__local_var_5_2 = ((($GLOBALS['Control_Monad_List_Trans_filter'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_filter')))($dictFunctor_0))($f_1);
$s__prime___6_3 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_4_1, $__local_var_5_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_filter"];
  $__res = ($__local_var_5_2)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_4_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
if (($f_1)(($v_3)->value0)) {
$__t4 = new Phpurs_Data2("Yield", ($v_3)->value0, $s__prime___6_3);
} else {
$__t4 = new Phpurs_Data1("Skip", $s__prime___6_3);
};
$__t0 = $__t4;
} else {
if ((is_object($v_3) && (($v_3)->tag === "Skip"))) {
$__local_var_4_5 = ($v_3)->value0;
$__local_var_5_6 = ((($GLOBALS['Control_Monad_List_Trans_filter'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_filter')))($dictFunctor_0))($f_1);
$__t0 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_4_5, $__local_var_5_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_filter"];
  $__res = ($__local_var_5_6)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_4_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Done"))) {
$__t0 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_dropWhile'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_dropWhile"), recVars=["Control_Monad_List_Trans_dropWhile"];
  while (true) {
$__local_var_1_0 = (((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($f_2, $v_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_dropWhile"];
  $__res = ((($__local_var_1_0)->map)(function($v_4) use ($dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_dropWhile"];
  if ((is_object($v_4) && (($v_4)->tag === "Yield"))) {
if (($f_2)(($v_4)->value0)) {
$__local_var_5_3 = ((($GLOBALS['Control_Monad_List_Trans_dropWhile'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_dropWhile')))($dictApplicative_0))($f_2);
$__t2 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_5_3, $v_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_dropWhile"];
  $__res = ($__local_var_5_3)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))(($v_4)->value1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
$__t2 = new Phpurs_Data2("Yield", ($v_4)->value0, ($v_4)->value1);
};
$__t1 = $__t2;
} else {
if ((is_object($v_4) && (($v_4)->tag === "Skip"))) {
$__local_var_5_4 = ($v_4)->value0;
$__local_var_6_5 = ((($GLOBALS['Control_Monad_List_Trans_dropWhile'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_dropWhile')))($dictApplicative_0))($f_2);
$__t1 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7) use ($__local_var_5_4, $__local_var_6_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_dropWhile"];
  $__res = ($__local_var_6_5)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_4) && (($v_4)->tag === "Done"))) {
$__t1 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_drop'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_drop"), recVars=["Control_Monad_List_Trans_drop"];
  while (true) {
$__local_var_1_0 = (((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($v_2, $v1_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_drop"];
  switch ($v_2) {
case 0:
$__t6 = $v1_3;
break;
default:
$__t6 = ((($__local_var_1_0)->map)(function($v2_4) use ($dictApplicative_0, $v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_drop"];
  if ((is_object($v2_4) && (($v2_4)->tag === "Yield"))) {
$__local_var_5_2 = ($v2_4)->value1;
$__local_var_6_3 = ((($GLOBALS['Control_Monad_List_Trans_drop'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_drop')))($dictApplicative_0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_2))(1));
$__t1 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7) use ($__local_var_5_2, $__local_var_6_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_drop"];
  $__res = ($__local_var_6_3)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v2_4) && (($v2_4)->tag === "Skip"))) {
$__local_var_5_4 = ($v2_4)->value0;
$__local_var_6_5 = ((($GLOBALS['Control_Monad_List_Trans_drop'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_drop')))($dictApplicative_0))($v_2);
$__t1 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7) use ($__local_var_5_4, $__local_var_6_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_drop"];
  $__res = ($__local_var_6_5)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v2_4) && (($v2_4)->tag === "Done"))) {
$__t1 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v1_3);
break;
};
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_cons'] = function() { $v = (function() {
  $__fn = function($dictApplicative_0, $lh_1 = null, $t_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_cons"), recVars=[];
  $__res = (($dictApplicative_0)->pure)(new Phpurs_Data2("Yield", (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($lh_1), $t_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_unfoldable1ListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_unfoldable1ListT"), recVars=[];
  $Applicative0_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $singleton1_2_1 = (($GLOBALS['Control_Monad_List_Trans_singleton'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_singleton')))($Applicative0_1_0);
  $__res = (object)["unfoldr1" => (function() use ($Applicative0_1_0, $singleton1_2_1) {
  $__fn = function($f_3, $b_4 = null) use ($Applicative0_1_0, $singleton1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_5_2 = null;
  $go_5_2 = function($v_6) use ($Applicative0_1_0, $f_3, &$go_5_2, $singleton1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_2"), recVars=["go_5_2"];
  while (true) {
if ((is_object(($v_6)->value1) && ((($v_6)->value1)->tag === "Nothing"))) {
$__t3 = ($singleton1_2_1)(($v_6)->value0);
} else {
if ((is_object(($v_6)->value1) && ((($v_6)->value1)->tag === "Just"))) {
$__local_var_7_4 = ($v_6)->value0;
$__local_var_8_5 = (($v_6)->value1)->value0;
$__t3 = (($Applicative0_1_0)->pure)(new Phpurs_Data2("Yield", (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_9) use ($__local_var_7_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_5_2"];
  $__res = $__local_var_7_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})), (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v1_9) use ($__local_var_8_5, $f_3, &$go_5_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_5_2"];
  $__res = ($go_5_2)(($f_3)($__local_var_8_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($go_5_2)(($f_3)($b_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_unfoldableListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_unfoldableListT"), recVars=[];
  $Applicative0_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $nil1_2_1 = (($Applicative0_1_0)->pure)(new Phpurs_Data0("Done"));
  $unfoldable1ListT1_3_2 = (($GLOBALS['Control_Monad_List_Trans_unfoldable1ListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_unfoldable1ListT')))($dictMonad_0);
  $__res = (object)["unfoldr" => (function() use ($Applicative0_1_0, $nil1_2_1) {
  $__fn = function($f_4, $b_5 = null) use ($Applicative0_1_0, $nil1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_6_3 = null;
  $go_6_3 = function($v_7) use ($Applicative0_1_0, $f_4, &$go_6_3, $nil1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_3"), recVars=["go_6_3"];
  while (true) {
if ((is_object($v_7) && (($v_7)->tag === "Nothing"))) {
$__t4 = $nil1_2_1;
} else {
if ((is_object($v_7) && (($v_7)->tag === "Just"))) {
$__local_var_8_5 = (($v_7)->value0)->value0;
$__local_var_9_6 = (($v_7)->value0)->value1;
$__t4 = (($Applicative0_1_0)->pure)(new Phpurs_Data2("Yield", (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_10) use ($__local_var_8_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_6_3"];
  $__res = $__local_var_8_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})), (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v1_10) use ($__local_var_9_6, $f_4, &$go_6_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_6_3"];
  $__res = ($go_6_3)(($f_4)($__local_var_9_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__res = $__t4;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ($go_6_3)(($f_4)($b_5));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Unfoldable10" => function($dollar__unused_4) use ($unfoldable1ListT1_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $unfoldable1ListT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_semigroupListT'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_semigroupListT"), recVars=["Control_Monad_List_Trans_semigroupListT","Control_Monad_List_Trans_concat"];
  while (true) {
$__res = (object)["append" => (($GLOBALS['Control_Monad_List_Trans_concat'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_concat')))($dictApplicative_0)];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_concat'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_concat"), recVars=["Control_Monad_List_Trans_semigroupListT","Control_Monad_List_Trans_concat"];
  while (true) {
$__local_var_1_0 = (((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($x_2, $y_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_semigroupListT","Control_Monad_List_Trans_concat"];
  $__res = ((($__local_var_1_0)->map)(function($v_4) use ($dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_semigroupListT","Control_Monad_List_Trans_concat"];
  if ((is_object($v_4) && (($v_4)->tag === "Yield"))) {
$__local_var_5_2 = ($v_4)->value1;
$__t1 = new Phpurs_Data2("Yield", ($v_4)->value0, (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_5_2, $dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_semigroupListT","Control_Monad_List_Trans_concat"];
  $__res = (((($GLOBALS['Control_Monad_List_Trans_concat'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_concat')))($dictApplicative_0))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_2)))($y_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_4) && (($v_4)->tag === "Skip"))) {
$__local_var_5_3 = ($v_4)->value0;
$__t1 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_5_3, $dictApplicative_0, $y_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_semigroupListT","Control_Monad_List_Trans_concat"];
  $__res = (((($GLOBALS['Control_Monad_List_Trans_concat'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_concat')))($dictApplicative_0))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_3)))($y_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_4) && (($v_4)->tag === "Done"))) {
$__t1 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5) use ($y_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_semigroupListT","Control_Monad_List_Trans_concat"];
  $__res = $y_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($x_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_monoidListT'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_monoidListT"), recVars=[];
  $semigroupListT1_1_0 = (object)["append" => (($GLOBALS['Control_Monad_List_Trans_concat'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_concat')))($dictApplicative_0)];
  $__res = (object)["mempty" => (($dictApplicative_0)->pure)(new Phpurs_Data0("Done")), "Semigroup0" => function($dollar__unused_2) use ($semigroupListT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $semigroupListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_catMaybes'] = function() { $v = function($dictFunctor_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_catMaybes"), recVars=[];
  $__res = ((($GLOBALS['Control_Monad_List_Trans_mapMaybe'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_mapMaybe')))($dictFunctor_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_monadListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_monadListT"), recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  while (true) {
$__res = (object)["Applicative0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applicativeListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applicativeListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = (($GLOBALS['Control_Monad_List_Trans_bindListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_bindListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_bindListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_bindListT"), recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  while (true) {
$append_1_0 = (($GLOBALS['Control_Monad_List_Trans_concat'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_concat')))((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
$__local_var_2_1 = (((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (object)["bind" => (function() use ($__local_var_2_1, $append_1_0, $dictMonad_0) {
  $__fn = function($fa_3, $f_4 = null) use ($__local_var_2_1, $append_1_0, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = ((($__local_var_2_1)->map)(function($v_5) use ($append_1_0, $dictMonad_0, $f_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  if ((is_object($v_5) && (($v_5)->tag === "Yield"))) {
$__local_var_6_3 = ($v_5)->value0;
$__local_var_7_4 = ($v_5)->value1;
$__t2 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_8) use ($__local_var_6_3, $__local_var_7_4, $append_1_0, $dictMonad_0, $f_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = (($append_1_0)(($f_4)($__local_var_6_3)))(((((($GLOBALS['Control_Monad_List_Trans_bindListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_bindListT')))($dictMonad_0))->bind)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_7_4)))($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_5) && (($v_5)->tag === "Skip"))) {
$__local_var_6_5 = ($v_5)->value0;
$__t2 = new Phpurs_Data1("Skip", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7) use ($__local_var_6_5, $dictMonad_0, $f_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = ((((($GLOBALS['Control_Monad_List_Trans_bindListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_bindListT')))($dictMonad_0))->bind)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_6_5)))($f_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
if ((is_object($v_5) && (($v_5)->tag === "Done"))) {
$__t2 = new Phpurs_Data0("Done");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($fa_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applyListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applyListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_applyListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_applyListT"), recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  while (true) {
$functorListT1_1_0 = (($GLOBALS['Control_Monad_List_Trans_functorListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_functorListT')))((((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
$__res = (object)["apply" => (($GLOBALS['Control_Monad_ap'] ?? \PhpursThunks::eval('Control_Monad_ap')))((object)["Applicative0" => function($dollar__unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applicativeListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applicativeListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = (($GLOBALS['Control_Monad_List_Trans_bindListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_bindListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]), "Functor0" => function($dollar__unused_2) use ($functorListT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = $functorListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_applicativeListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_applicativeListT"), recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  while (true) {
$__res = (object)["pure" => (($GLOBALS['Control_Monad_List_Trans_singleton'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_singleton')))((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))), "Apply0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_List_Trans_monadListT","Control_Monad_List_Trans_bindListT","Control_Monad_List_Trans_applyListT","Control_Monad_List_Trans_applicativeListT"];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applyListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applyListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_monadEffectListT'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_monadEffectListT"), recVars=[];
  $Monad0_1_0 = (($dictMonadEffect_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadListT1_2_1 = (object)["Applicative0" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["pure" => (($GLOBALS['Control_Monad_List_Trans_singleton'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_singleton')))((($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))), "Apply0" => function($dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applyListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applyListT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_bindListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_bindListT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftEffect" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Control_Monad_List_Trans_fromEffect'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_fromEffect')))((($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))))(($dictMonadEffect_0)->liftEffect), "Monad0" => function($dollar__unused_3) use ($monadListT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_monadSTListT'] = function() { $v = function($dictMonadST_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_monadSTListT"), recVars=[];
  $Monad0_1_0 = (($dictMonadST_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadListT1_2_1 = (object)["Applicative0" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["pure" => (($GLOBALS['Control_Monad_List_Trans_singleton'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_singleton')))((($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))), "Apply0" => function($dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applyListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applyListT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_bindListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_bindListT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftST" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Control_Monad_List_Trans_fromEffect'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_fromEffect')))((($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))))(($dictMonadST_0)->liftST), "Monad0" => function($dollar__unused_3) use ($monadListT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_altListT'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_altListT"), recVars=[];
  $functorListT1_1_0 = (($GLOBALS['Control_Monad_List_Trans_functorListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_functorListT')))((((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["alt" => (($GLOBALS['Control_Monad_List_Trans_concat'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_concat')))($dictApplicative_0), "Functor0" => function($dollar__unused_2) use ($functorListT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $functorListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_plusListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_plusListT"), recVars=[];
  $Applicative0_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $altListT1_2_1 = (($GLOBALS['Control_Monad_List_Trans_altListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_altListT')))($Applicative0_1_0);
  $__res = (object)["empty" => (($Applicative0_1_0)->pure)(new Phpurs_Data0("Done")), "Alt0" => function($dollar__unused_3) use ($altListT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $altListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_alternativeListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_alternativeListT"), recVars=[];
  $applicativeListT1_1_0 = (object)["pure" => (($GLOBALS['Control_Monad_List_Trans_singleton'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_singleton')))((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))), "Apply0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applyListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applyListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusListT1_2_1 = (($GLOBALS['Control_Monad_List_Trans_plusListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_plusListT')))($dictMonad_0);
  $__res = (object)["Applicative0" => function($dollar__unused_3) use ($applicativeListT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applicativeListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_3) use ($plusListT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $plusListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_List_Trans_monadPlusListT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_List_Trans_monadPlusListT"), recVars=[];
  $monadListT1_1_0 = (object)["Applicative0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["pure" => (($GLOBALS['Control_Monad_List_Trans_singleton'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_singleton')))((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))), "Apply0" => function($dollar__unused_2) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_applyListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_applyListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_List_Trans_bindListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_bindListT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeListT1_2_1 = (($GLOBALS['Control_Monad_List_Trans_alternativeListT'] ?? \PhpursThunks::eval('Control_Monad_List_Trans_alternativeListT')))($dictMonad_0);
  $__res = (object)["Monad0" => function($dollar__unused_3) use ($monadListT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadListT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_3) use ($alternativeListT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $alternativeListT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






















































