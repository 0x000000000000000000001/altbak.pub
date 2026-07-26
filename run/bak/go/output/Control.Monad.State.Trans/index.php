<?php

namespace Control\Monad\State\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.State.Trans, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.State.Trans, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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
\PhpursThunks::$thunks['Control_Monad_State_Trans_StateT'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_withStateT'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($v_1))($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_runStateT'] = function() { $v = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_newtypeStateT'] = function() { $v = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadTransStateT'] = function() { $v = ["lift" => (function() {
  $__fn = function($dictMonad_0 = null, $m_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictMonad_0)['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['bind'])($m_1))(function($x_3 = null) use ($dictMonad_0, $s_2) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['pure'])(new Phpurs_Data2("Tuple", $x_3, $s_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_mapStateT'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_0))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_lazyStateT'] = function() { $v = ["defer" => (function() {
  $__fn = function($f_0 = null, $s_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'))))($s_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_functorStateT'] = function() { $v = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1 = null, $v_2 = null, $s_3 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(function($v1_4 = null) use ($f_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_1)(($v1_4)->{'value0'}), ($v1_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_2)($s_3));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_execStateT'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0 = null, $v_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(($GLOBALS['Data_Tuple_snd'] ?? \PhpursThunks::eval('Data_Tuple_snd'))))(($v_1)($s_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_evalStateT'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0 = null, $v_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(($GLOBALS['Data_Tuple_fst'] ?? \PhpursThunks::eval('Data_Tuple_fst'))))(($v_1)($s_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ["Applicative0" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_bindStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ["bind" => (function() use ($dictMonad_0) {
  $__fn = function($v_1 = null, $f_2 = null, $s_3 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictMonad_0)['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['bind'])(($v_1)($s_3)))(function($v1_4 = null) use ($f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)(($v1_4)->{'value0'}))(($v1_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applyStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applyStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_applyStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((((($dictMonad_0)['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorStateT1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $v_3 = null, $s_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)['map'])(function($v1_5 = null) use ($f_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_2)(($v1_5)->{'value0'}), ($v1_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_3)($s_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["apply" => (($GLOBALS['Control_Monad_ap'] ?? \PhpursThunks::eval('Control_Monad_ap')))(["Applicative0" => function($dollar__unused_3 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]), "Functor0" => function($dollar__unused_3 = null) use ($functorStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_applicativeStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ["pure" => (function() use ($dictMonad_0) {
  $__fn = function($a_1 = null, $s_2 = null) use ($dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictMonad_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['pure'])(new Phpurs_Data2("Tuple", $a_1, $s_2));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applyStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applyStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_semigroupStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $lift2_1_0 = (($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))((($GLOBALS['Control_Monad_State_Trans_applyStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applyStateT')))($dictMonad_0));
  $__res = function($dictSemigroup_2 = null) use ($lift2_1_0) {
  $__num = \func_num_args();
  $__res = ["append" => ($lift2_1_0)(($dictSemigroup_2)['append'])];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadAskStateT'] = function() { $v = function($dictMonadAsk_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadAsk_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["ask" => (((($GLOBALS['Control_Monad_State_Trans_monadTransStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadTransStateT')))['lift'])($Monad0_1_0))(($dictMonadAsk_0)['ask']), "Monad0" => function($dollar__unused_3 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadReaderStateT'] = function() { $v = function($dictMonadReader_0 = null) {
  $__num = \func_num_args();
  $monadAskStateT1_1_0 = (($GLOBALS['Control_Monad_State_Trans_monadAskStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadAskStateT')))((($dictMonadReader_0)['MonadAsk0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["local" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Control_Monad_State_Trans_mapStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_mapStateT'))))(($dictMonadReader_0)['local']), "MonadAsk0" => function($dollar__unused_2 = null) use ($monadAskStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadAskStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadContStateT'] = function() { $v = function($dictMonadCont_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonadCont_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["callCC" => (function() use ($dictMonadCont_0) {
  $__fn = function($f_3 = null, $s_4 = null) use ($dictMonadCont_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictMonadCont_0)['callCC'])(function($c_5 = null) use ($f_3, $s_4) {
  $__num = \func_num_args();
  $__res = (($f_3)((function() use ($c_5) {
  $__fn = function($a_6 = null, $s__prime___7 = null) use ($c_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($c_5)(new Phpurs_Data2("Tuple", $a_6, $s__prime___7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($s_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_3 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadEffectState'] = function() { $v = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftEffect" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Monad_State_Trans_monadTransStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadTransStateT')))['lift'])($Monad0_1_0)))(($dictMonadEffect_0)['liftEffect']), "Monad0" => function($dollar__unused_3 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadRecStateT'] = function() { $v = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["tailRecM" => (function() use ($Monad0_1_0, $dictMonadRec_0) {
  $__fn = function($f_3 = null, $a_4 = null, $s_5 = null) use ($Monad0_1_0, $dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($v_6 = null) use ($Monad0_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((((($Monad0_1_0)['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['bind'])((($f_3)(($v_6)->{'value0'}))(($v_6)->{'value1'})))(function($v2_7 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object(($v2_7)->{'value0'}) && ((($v2_7)->{'value0'})->{'tag'} === "Loop"))) {
$__t2 = new Phpurs_Data1("Loop", new Phpurs_Data2("Tuple", (($v2_7)->{'value0'})->{'value0'}, ($v2_7)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object(($v2_7)->{'value0'}) && ((($v2_7)->{'value0'})->{'tag'} === "Done"))) {
$__t2 = new Phpurs_Data1("Done", new Phpurs_Data2("Tuple", (($v2_7)->{'value0'})->{'value0'}, ($v2_7)->{'value1'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = (((($Monad0_1_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['pure'])($__t2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", $a_4, $s_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_3 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadStateStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $pure_1_0 = ((($dictMonad_0)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['pure'];
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["state" => function($f_3 = null) use ($pure_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($pure_1_0))($f_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_3 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadTellStateT'] = function() { $v = function($dictMonadTell_0 = null) {
  $__num = \func_num_args();
  $Monad1_1_0 = (($dictMonadTell_0)['Monad1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Semigroup0_2_1 = (($dictMonadTell_0)['Semigroup0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_3_2 = ["Applicative0" => function($dollar__unused_3 = null) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3 = null) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["tell" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Monad_State_Trans_monadTransStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadTransStateT')))['lift'])($Monad1_1_0)))(($dictMonadTell_0)['tell']), "Semigroup0" => function($dollar__unused_4 = null) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($dollar__unused_4 = null) use ($monadStateT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadStateT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadWriterStateT'] = function() { $v = function($dictMonadWriter_0 = null) {
  $__num = \func_num_args();
  $MonadTell1_1_0 = (($dictMonadWriter_0)['MonadTell1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Monad1_2_1 = (($MonadTell1_1_0)['Monad1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_3_2 = (($Monad1_2_1)['Bind1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_4_3 = (($Monad1_2_1)['Applicative0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Monoid0_5_4 = (($dictMonadWriter_0)['Monoid0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadTellStateT1_6_5 = (($GLOBALS['Control_Monad_State_Trans_monadTellStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadTellStateT')))($MonadTell1_1_0);
  $__res = ["listen" => (function() use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__fn = function($m_7 = null, $s_8 = null) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_3_2)['bind'])((($dictMonadWriter_0)['listen'])(($m_7)($s_8))))(function($v_9 = null) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = (($__local_var_4_3)['pure'])(new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", (($v_9)->{'value0'})->{'value0'}, ($v_9)->{'value1'}), (($v_9)->{'value0'})->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "pass" => (function() use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__fn = function($m_7 = null, $s_8 = null) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictMonadWriter_0)['pass'])(((($__local_var_3_2)['bind'])(($m_7)($s_8)))(function($v_9 = null) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__res = (($__local_var_4_3)['pure'])(new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", (($v_9)->{'value0'})->{'value0'}, ($v_9)->{'value1'}), (($v_9)->{'value0'})->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monoid0" => function($dollar__unused_7 = null) use ($Monoid0_5_4) {
  $__num = \func_num_args();
  $__res = $Monoid0_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($dollar__unused_7 = null) use ($monadTellStateT1_6_5) {
  $__num = \func_num_args();
  $__res = $monadTellStateT1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadThrowStateT'] = function() { $v = function($dictMonadThrow_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadThrow_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["throwError" => function($e_3 = null) use ($Monad0_1_0, $dictMonadThrow_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Control_Monad_State_Trans_monadTransStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadTransStateT')))['lift'])($Monad0_1_0))((($dictMonadThrow_0)['throwError'])($e_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_3 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadErrorStateT'] = function() { $v = function($dictMonadError_0 = null) {
  $__num = \func_num_args();
  $monadThrowStateT1_1_0 = (($GLOBALS['Control_Monad_State_Trans_monadThrowStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadThrowStateT')))((($dictMonadError_0)['MonadThrow0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["catchError" => (function() use ($dictMonadError_0) {
  $__fn = function($v_2 = null, $h_3 = null, $s_4 = null) use ($dictMonadError_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadError_0)['catchError'])(($v_2)($s_4)))(function($e_5 = null) use ($h_3, $s_4) {
  $__num = \func_num_args();
  $__res = (($h_3)($e_5))($s_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($dollar__unused_2 = null) use ($monadThrowStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadThrowStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadSTStateT'] = function() { $v = function($dictMonadST_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadST_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftST" => ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Control_Monad_State_Trans_monadTransStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_monadTransStateT')))['lift'])($Monad0_1_0)))(($dictMonadST_0)['liftST']), "Monad0" => function($dollar__unused_3 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_monoidStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $semigroupStateT1_1_0 = (($GLOBALS['Control_Monad_State_Trans_semigroupStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_semigroupStateT')))($dictMonad_0);
  $__res = function($dictMonoid_2 = null) use ($dictMonad_0, $semigroupStateT1_1_0) {
  $__num = \func_num_args();
  $semigroupStateT2_3_1 = ($semigroupStateT1_1_0)((($dictMonoid_2)['Semigroup0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["mempty" => (((($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($dictMonad_0))['pure'])(($dictMonoid_2)['mempty']), "Semigroup0" => function($dollar__unused_4 = null) use ($semigroupStateT2_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupStateT2_3_1;
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
\PhpursThunks::$thunks['Control_Monad_State_Trans_altStateT'] = function() { $v = (function() {
  $__fn = function($dictMonad_0 = null, $dictAlt_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictAlt_1)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorStateT1_3_1 = ["map" => (function() use ($__local_var_2_0) {
  $__fn = function($f_3 = null, $v_4 = null, $s_5 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_2_0)['map'])(function($v1_6 = null) use ($f_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_3)(($v1_6)->{'value0'}), ($v1_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_4)($s_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["alt" => (function() use ($dictAlt_1) {
  $__fn = function($v_4 = null, $v1_5 = null, $s_6 = null) use ($dictAlt_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictAlt_1)['alt'])(($v_4)($s_6)))(($v1_5)($s_6));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_4 = null) use ($functorStateT1_3_1) {
  $__num = \func_num_args();
  $__res = $functorStateT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_plusStateT'] = function() { $v = (function() {
  $__fn = function($dictMonad_0 = null, $dictPlus_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $empty_2_0 = ($dictPlus_1)['empty'];
  $altStateT2_3_1 = ((($GLOBALS['Control_Monad_State_Trans_altStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_altStateT')))($dictMonad_0))((($dictPlus_1)['Alt0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["empty" => function($v_4 = null) use ($empty_2_0) {
  $__num = \func_num_args();
  $__res = $empty_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($dollar__unused_4 = null) use ($altStateT2_3_1) {
  $__num = \func_num_args();
  $__res = $altStateT2_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_State_Trans_alternativeStateT'] = function() { $v = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $applicativeStateT1_1_0 = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($dictMonad_0);
  $__res = function($dictAlternative_2 = null) use ($applicativeStateT1_1_0, $dictMonad_0) {
  $__num = \func_num_args();
  $plusStateT2_3_1 = ((($GLOBALS['Control_Monad_State_Trans_plusStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_plusStateT')))($dictMonad_0))((($dictAlternative_2)['Plus1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["Applicative0" => function($dollar__unused_4 = null) use ($applicativeStateT1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeStateT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_4 = null) use ($plusStateT2_3_1) {
  $__num = \func_num_args();
  $__res = $plusStateT2_3_1;
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
\PhpursThunks::$thunks['Control_Monad_State_Trans_monadPlusStateT'] = function() { $v = function($dictMonadPlus_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadPlus_0)['Monad0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $monadStateT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_applicativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_applicativeStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_State_Trans_bindStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_bindStateT')))($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeStateT1_3_2 = ((($GLOBALS['Control_Monad_State_Trans_alternativeStateT'] ?? \PhpursThunks::eval('Control_Monad_State_Trans_alternativeStateT')))($Monad0_1_0))((($dictMonadPlus_0)['Alternative1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["Monad0" => function($dollar__unused_4 = null) use ($monadStateT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadStateT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_4 = null) use ($alternativeStateT1_3_2) {
  $__num = \func_num_args();
  $__res = $alternativeStateT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };

































