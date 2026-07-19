<?php

namespace Control\Monad\Maybe\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Maybe.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Maybe, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Maybe.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Maybe, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Maybe.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_MaybeT'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_MaybeT"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_runMaybeT'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_runMaybeT"), recVars=[];
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_newtypeMaybeT'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadTransMaybeT'] = function() { $v = (object)["lift" => function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))($dictMonad_0))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_mapMaybeT'] = function() { $v = (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_mapMaybeT"), recVars=[];
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_functorMaybeT'] = function() { $v = function($dictFunctor_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_functorMaybeT"), recVars=[];
  $__res = (object)["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1, $v_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictFunctor_0)->map)(function($v1_3) use ($f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t0 = new Phpurs_Data1("Just", ($f_1)(($v1_3)->value0));
} else {
$__t0 = new Phpurs_Data0("Nothing");
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
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadMaybeT"), recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  while (true) {
$__res = (object)["Applicative0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_bindMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_bindMaybeT"), recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  while (true) {
$__res = (object)["bind" => (function() use ($dictMonad_0) {
  $__fn = function($v_1, $f_2 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = ((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)($v_1))(function($v1_3) use ($dictMonad_0, $f_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  if ((is_object($v1_3) && (($v1_3)->tag === "Nothing"))) {
$__t0 = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(new Phpurs_Data0("Nothing"));
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Just"))) {
$__t0 = ($f_2)(($v1_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
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
})(), "Apply0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applyMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_applyMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_applyMaybeT"), recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  while (true) {
$__local_var_1_0 = (((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$functorMaybeT1_2_1 = (object)["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = ((($__local_var_1_0)->map)(function($v1_4) use ($f_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  if ((is_object($v1_4) && (($v1_4)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($f_2)(($v1_4)->value0));
} else {
$__t1 = new Phpurs_Data0("Nothing");
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
})()];
$__res = (object)["apply" => (($GLOBALS['Control_Monad_ap'] ?? \PhpursThunks::eval('Control_Monad_ap')))((object)["Applicative0" => function($dollar__unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]), "Functor0" => function($dollar__unused_3) use ($functorMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = $functorMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_applicativeMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_applicativeMaybeT"), recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  while (true) {
$__res = (object)["pure" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')))), "Apply0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_Maybe_Trans_monadMaybeT","Control_Monad_Maybe_Trans_bindMaybeT","Control_Monad_Maybe_Trans_applyMaybeT","Control_Monad_Maybe_Trans_applicativeMaybeT"];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applyMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_semigroupMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_semigroupMaybeT"), recVars=[];
  $lift2_1_0 = (($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))((($GLOBALS['Control_Monad_Maybe_Trans_applyMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applyMaybeT')))($dictMonad_0));
  $__res = function($dictSemigroup_2) use ($lift2_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["append" => ($lift2_1_0)(($dictSemigroup_2)->append)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadAskMaybeT'] = function() { $v = function($dictMonadAsk_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadAskMaybeT"), recVars=[];
  $Monad0_1_0 = (($dictMonadAsk_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["ask" => (((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))($Monad0_1_0))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just')))))(($dictMonadAsk_0)->ask), "Monad0" => function($dollar__unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadReaderMaybeT'] = function() { $v = function($dictMonadReader_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadReaderMaybeT"), recVars=[];
  $monadAskMaybeT1_1_0 = (($GLOBALS['Control_Monad_Maybe_Trans_monadAskMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_monadAskMaybeT')))((($dictMonadReader_0)->MonadAsk0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["local" => function($f_2) use ($dictMonadReader_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictMonadReader_0)->local)($f_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($dollar__unused_2) use ($monadAskMaybeT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadAskMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadContMaybeT'] = function() { $v = function($dictMonadCont_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadContMaybeT"), recVars=[];
  $__local_var_1_0 = (($dictMonadCont_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($dollar__unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["callCC" => function($f_3) use ($dictMonadCont_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictMonadCont_0)->callCC)(function($c_4) use ($f_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($f_3)(function($a_5) use ($c_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($c_4)(new Phpurs_Data1("Just", $a_5));
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
}, "Monad0" => function($dollar__unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadEffectMaybe'] = function() { $v = function($dictMonadEffect_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadEffectMaybe"), recVars=[];
  $Monad0_1_0 = (($dictMonadEffect_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftEffect" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))($Monad0_1_0))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))))(($dictMonadEffect_0)->liftEffect), "Monad0" => function($dollar__unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadRecMaybeT'] = function() { $v = function($dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadRecMaybeT"), recVars=[];
  $Monad0_1_0 = (($dictMonadRec_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["tailRecM" => function($f_3) use ($Monad0_1_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))((($dictMonadRec_0)->tailRecM)(function($a_4) use ($Monad0_1_0, $f_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($Monad0_1_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)(($f_3)($a_4)))(function($m__prime___5) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($m__prime___5) && (($m__prime___5)->tag === "Nothing"))) {
$__t2 = new Phpurs_Data1("Done", new Phpurs_Data0("Nothing"));
} else {
if ((is_object($m__prime___5) && (($m__prime___5)->tag === "Just"))) {
if ((is_object(($m__prime___5)->value0) && ((($m__prime___5)->value0)->tag === "Loop"))) {
$__t3 = new Phpurs_Data1("Loop", (($m__prime___5)->value0)->value0);
} else {
if ((is_object(($m__prime___5)->value0) && ((($m__prime___5)->value0)->tag === "Done"))) {
$__t3 = new Phpurs_Data1("Done", new Phpurs_Data1("Just", (($m__prime___5)->value0)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__t2 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = (((($Monad0_1_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($__t2);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadStateMaybeT'] = function() { $v = function($dictMonadState_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadStateMaybeT"), recVars=[];
  $Monad0_1_0 = (($dictMonadState_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $lift1_2_1 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))($Monad0_1_0))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))));
  $monadMaybeT1_3_2 = (object)["Applicative0" => function($dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["state" => function($f_4) use ($dictMonadState_0, $lift1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($lift1_2_1)((($dictMonadState_0)->state)($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_4) use ($monadMaybeT1_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadTellMaybeT'] = function() { $v = function($dictMonadTell_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadTellMaybeT"), recVars=[];
  $Monad1_1_0 = (($dictMonadTell_0)->Monad1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Semigroup0_2_1 = (($dictMonadTell_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadMaybeT1_3_2 = (object)["Applicative0" => function($dollar__unused_3) use ($Monad1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3) use ($Monad1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["tell" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))($Monad1_1_0))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))))(($dictMonadTell_0)->tell), "Semigroup0" => function($dollar__unused_4) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($dollar__unused_4) use ($monadMaybeT1_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadWriterMaybeT'] = function() { $v = function($dictMonadWriter_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadWriterMaybeT"), recVars=[];
  $MonadTell1_1_0 = (($dictMonadWriter_0)->MonadTell1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Monad1_2_1 = (($MonadTell1_1_0)->Monad1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_3_2 = (($Monad1_2_1)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_4_3 = (($Monad1_2_1)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Monoid0_5_4 = (($dictMonadWriter_0)->Monoid0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadTellMaybeT1_6_5 = (($GLOBALS['Control_Monad_Maybe_Trans_monadTellMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_monadTellMaybeT')))($MonadTell1_1_0);
  $__res = (object)["listen" => function($v_7) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($__local_var_3_2)->bind)((($dictMonadWriter_0)->listen)($v_7)))(function($v_8) use ($__local_var_4_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(($v_8)->value0) && ((($v_8)->value0)->tag === "Just"))) {
$__t6 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", (($v_8)->value0)->value0, ($v_8)->value1));
} else {
$__t6 = new Phpurs_Data0("Nothing");
};
  $__res = (($__local_var_4_3)->pure)($__t6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pass" => function($v_7) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($dictMonadWriter_0)->pass)(((($__local_var_3_2)->bind)($v_7))(function($a_8) use ($__local_var_4_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($a_8) && (($a_8)->tag === "Nothing"))) {
$__t7 = new Phpurs_Data2("Tuple", new Phpurs_Data0("Nothing"), (($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
} else {
if ((is_object($a_8) && (($a_8)->tag === "Just"))) {
$__t7 = new Phpurs_Data2("Tuple", new Phpurs_Data1("Just", (($a_8)->value0)->value0), (($a_8)->value0)->value1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
  $__res = (($__local_var_4_3)->pure)($__t7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monoid0" => function($dollar__unused_7) use ($Monoid0_5_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $Monoid0_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($dollar__unused_7) use ($monadTellMaybeT1_6_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadTellMaybeT1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadThrowMaybeT'] = function() { $v = function($dictMonadThrow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadThrowMaybeT"), recVars=[];
  $Monad0_1_0 = (($dictMonadThrow_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $lift1_2_1 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))($Monad0_1_0))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))));
  $monadMaybeT1_3_2 = (object)["Applicative0" => function($dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["throwError" => function($e_4) use ($dictMonadThrow_0, $lift1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($lift1_2_1)((($dictMonadThrow_0)->throwError)($e_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_4) use ($monadMaybeT1_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadErrorMaybeT'] = function() { $v = function($dictMonadError_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadErrorMaybeT"), recVars=[];
  $monadThrowMaybeT1_1_0 = (($GLOBALS['Control_Monad_Maybe_Trans_monadThrowMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_monadThrowMaybeT')))((($dictMonadError_0)->MonadThrow0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["catchError" => (function() use ($dictMonadError_0) {
  $__fn = function($v_2, $h_3 = null) use ($dictMonadError_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictMonadError_0)->catchError)($v_2))(function($a_4) use ($h_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($h_3)($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($dollar__unused_2) use ($monadThrowMaybeT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadThrowMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadSTMaybeT'] = function() { $v = function($dictMonadST_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadSTMaybeT"), recVars=[];
  $Monad0_1_0 = (($dictMonadST_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadMaybeT1_2_1 = (object)["Applicative0" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2) use ($Monad0_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["liftST" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_Maybe_Trans_MaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_MaybeT'))))(((($GLOBALS['Control_Monad_liftM1'] ?? \PhpursThunks::eval('Control_Monad_liftM1')))($Monad0_1_0))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))))(($dictMonadST_0)->liftST), "Monad0" => function($dollar__unused_3) use ($monadMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monoidMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monoidMaybeT"), recVars=[];
  $semigroupMaybeT1_1_0 = (($GLOBALS['Control_Monad_Maybe_Trans_semigroupMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_semigroupMaybeT')))($dictMonad_0);
  $__res = function($dictMonoid_2) use ($dictMonad_0, $semigroupMaybeT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $semigroupMaybeT2_3_1 = ($semigroupMaybeT1_1_0)((($dictMonoid_2)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["mempty" => (((($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($dictMonad_0))->pure)(($dictMonoid_2)->mempty), "Semigroup0" => function($dollar__unused_4) use ($semigroupMaybeT2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $semigroupMaybeT2_3_1;
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
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_altMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_altMaybeT"), recVars=[];
  $Bind1_1_0 = (($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (((($Bind1_1_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorMaybeT1_3_2 = (object)["map" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3, $v_4 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($__local_var_2_1)->map)(function($v1_5) use ($f_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_5) && (($v1_5)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", ($f_3)(($v1_5)->value0));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["alt" => (function() use ($Bind1_1_0, $dictMonad_0) {
  $__fn = function($v_4, $v1_5 = null) use ($Bind1_1_0, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($Bind1_1_0)->bind)($v_4))(function($m_6) use ($dictMonad_0, $v1_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($m_6) && (($m_6)->tag === "Nothing"))) {
$__t4 = $v1_5;
} else {
$__t4 = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($m_6);
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_4) use ($functorMaybeT1_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $functorMaybeT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_plusMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_plusMaybeT"), recVars=[];
  $altMaybeT1_1_0 = (($GLOBALS['Control_Monad_Maybe_Trans_altMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_altMaybeT')))($dictMonad_0);
  $__res = (object)["empty" => (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(new Phpurs_Data0("Nothing")), "Alt0" => function($dollar__unused_2) use ($altMaybeT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $altMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_alternativeMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_alternativeMaybeT"), recVars=[];
  $applicativeMaybeT1_1_0 = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($dictMonad_0);
  $plusMaybeT1_2_1 = (($GLOBALS['Control_Monad_Maybe_Trans_plusMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_plusMaybeT')))($dictMonad_0);
  $__res = (object)["Applicative0" => function($dollar__unused_3) use ($applicativeMaybeT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applicativeMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_3) use ($plusMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $plusMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Maybe_Trans_monadPlusMaybeT'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_Maybe_Trans_monadPlusMaybeT"), recVars=[];
  $monadMaybeT1_1_0 = (object)["Applicative0" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_applicativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_applicativeMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_1) use ($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Control_Monad_Maybe_Trans_bindMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_bindMaybeT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeMaybeT1_2_1 = (($GLOBALS['Control_Monad_Maybe_Trans_alternativeMaybeT'] ?? \PhpursThunks::eval('Control_Monad_Maybe_Trans_alternativeMaybeT')))($dictMonad_0);
  $__res = (object)["Monad0" => function($dollar__unused_3) use ($monadMaybeT1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $monadMaybeT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_3) use ($alternativeMaybeT1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $alternativeMaybeT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };





























