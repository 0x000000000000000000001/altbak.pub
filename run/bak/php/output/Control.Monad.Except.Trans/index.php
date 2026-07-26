<?php

namespace Control\Monad\Except\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Except.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Except.Trans, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Either, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Except.Trans/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
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


// Control_Monad_Except_Trans_ExceptT
$GLOBALS['Control_Monad_Except_Trans_ExceptT'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_withExceptT
$GLOBALS['Control_Monad_Except_Trans_withExceptT'] = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(function($v2_3 = null) use ($f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Right"))) {
$__t0 = new Phpurs_Data1("Right", ($v2_3)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($v2_3) && (($v2_3)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data1("Left", ($f_1)(($v2_3)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Except_Trans_runExceptT
$GLOBALS['Control_Monad_Except_Trans_runExceptT'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_newtypeExceptT
$GLOBALS['Control_Monad_Except_Trans_newtypeExceptT'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Except_Trans_monadTransExceptT
$GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'] = ["lift" => (function() {
  $__fn = function($dictMonad_0 = null, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])($m_1))(function($a_2 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])(new Phpurs_Data1("Right", $a_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Control_Monad_Except_Trans_mapExceptT
$GLOBALS['Control_Monad_Except_Trans_mapExceptT'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Except_Trans_functorExceptT
$GLOBALS['Control_Monad_Except_Trans_functorExceptT'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => function($f_1 = null) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = (($dictFunctor_0)['map'])(function($m_2 = null) use ($f_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($m_2) && (($m_2)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data1("Left", ($m_2)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($m_2) && (($m_2)->{'tag'} === "Right"))) {
$__t0 = new Phpurs_Data1("Right", ($f_1)(($m_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_except
$GLOBALS['Control_Monad_Except_Trans_except'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Except_Trans_ExceptT']))(($dictApplicative_0)['pure']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadExceptT
$GLOBALS['Control_Monad_Except_Trans_monadExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ["Applicative0" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_bindExceptT
$GLOBALS['Control_Monad_Except_Trans_bindExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $pure_1_0 = ((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'];
  $__res = ["bind" => (function() use ($dictMonad_0, $pure_1_0) {
  $__fn = function($v_2 = null, $k_3 = null) use ($dictMonad_0, $pure_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_1 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_1_0))($GLOBALS['Data_Either_Left']);
  $__res = ((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])($v_2))(function($v2_5 = null) use ($__local_var_4_1, $k_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "Left"))) {
$__t2 = ($__local_var_4_1)(($v2_5)->{'value0'});
goto end_branch_2;;
};
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "Right"))) {
$__t2 = ($k_3)(($v2_5)->{'value0'});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_2 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_applyExceptT
$GLOBALS['Control_Monad_Except_Trans_applyExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']);
  $functorExceptT1_2_1 = ["map" => function($f_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)['map'])(function($m_3 = null) use ($f_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($m_3) && (($m_3)->{'tag'} === "Left"))) {
$__t1 = new Phpurs_Data1("Left", ($m_3)->{'value0'});
goto end_branch_1;;
};
  if ((is_object($m_3) && (($m_3)->{'tag'} === "Right"))) {
$__t1 = new Phpurs_Data1("Right", ($f_2)(($m_3)->{'value0'}));
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
}];
  $__res = ["apply" => ($GLOBALS['Control_Monad_ap'])(["Applicative0" => function($dollar__unused_3 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]), "Functor0" => function($dollar__unused_3 = null) use ($functorExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_applicativeExceptT
$GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__res = ["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Except_Trans_ExceptT']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure']))($GLOBALS['Data_Either_Right'])), "Apply0" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_semigroupExceptT
$GLOBALS['Control_Monad_Except_Trans_semigroupExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $lift2_1_0 = ($GLOBALS['Control_Apply_lift2'])(($GLOBALS['Control_Monad_Except_Trans_applyExceptT'])($dictMonad_0));
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
};

// Control_Monad_Except_Trans_monadAskExceptT
$GLOBALS['Control_Monad_Except_Trans_monadAskExceptT'] = function($dictMonadAsk_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadAsk_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadExceptT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["ask" => ((($GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'])['lift'])($Monad0_1_0))(($dictMonadAsk_0)['ask']), "Monad0" => function($dollar__unused_3 = null) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadReaderExceptT
$GLOBALS['Control_Monad_Except_Trans_monadReaderExceptT'] = function($dictMonadReader_0 = null) {
  $__num = \func_num_args();
  $monadAskExceptT1_1_0 = ($GLOBALS['Control_Monad_Except_Trans_monadAskExceptT'])((($dictMonadReader_0)['MonadAsk0'])($GLOBALS['Prim_undefined']));
  $__res = ["local" => function($f_2 = null) use ($dictMonadReader_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadReader_0)['local'])($f_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($dollar__unused_2 = null) use ($monadAskExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadAskExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadContExceptT
$GLOBALS['Control_Monad_Except_Trans_monadContExceptT'] = function($dictMonadCont_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonadCont_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadExceptT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($__local_var_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["callCC" => function($f_3 = null) use ($dictMonadCont_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadCont_0)['callCC'])(function($c_4 = null) use ($f_3) {
  $__num = \func_num_args();
  $__res = ($f_3)(function($a_5 = null) use ($c_4) {
  $__num = \func_num_args();
  $__res = ($c_4)(new Phpurs_Data1("Right", $a_5));
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
}, "Monad0" => function($dollar__unused_3 = null) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadEffectExceptT
$GLOBALS['Control_Monad_Except_Trans_monadEffectExceptT'] = function($dictMonadEffect_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadEffect_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadExceptT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'])['lift'])($Monad0_1_0)))(($dictMonadEffect_0)['liftEffect']), "Monad0" => function($dollar__unused_3 = null) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadRecExceptT
$GLOBALS['Control_Monad_Except_Trans_monadRecExceptT'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadExceptT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["tailRecM" => function($f_3 = null) use ($Monad0_1_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Except_Trans_ExceptT']))((($dictMonadRec_0)['tailRecM'])(function($a_4 = null) use ($Monad0_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((((($Monad0_1_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])(($f_3)($a_4)))(function($m__prime___5 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($m__prime___5) && (($m__prime___5)->{'tag'} === "Left"))) {
$__t2 = new Phpurs_Data1("Done", new Phpurs_Data1("Left", ($m__prime___5)->{'value0'}));
goto end_branch_2;;
};
  if ((is_object($m__prime___5) && (($m__prime___5)->{'tag'} === "Right"))) {
$__t3 = null;;
if ((is_object(($m__prime___5)->{'value0'}) && ((($m__prime___5)->{'value0'})->{'tag'} === "Loop"))) {
$__t3 = new Phpurs_Data1("Loop", (($m__prime___5)->{'value0'})->{'value0'});
goto end_branch_3;;
};
if ((is_object(($m__prime___5)->{'value0'}) && ((($m__prime___5)->{'value0'})->{'tag'} === "Done"))) {
$__t3 = new Phpurs_Data1("Done", new Phpurs_Data1("Right", (($m__prime___5)->{'value0'})->{'value0'}));
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = (((($Monad0_1_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])($__t2);
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
}, "Monad0" => function($dollar__unused_3 = null) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadStateExceptT
$GLOBALS['Control_Monad_Except_Trans_monadStateExceptT'] = function($dictMonadState_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadState_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadExceptT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["state" => function($f_3 = null) use ($Monad0_1_0, $dictMonadState_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'])['lift'])($Monad0_1_0))((($dictMonadState_0)['state'])($f_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_3 = null) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadTellExceptT
$GLOBALS['Control_Monad_Except_Trans_monadTellExceptT'] = function($dictMonadTell_0 = null) {
  $__num = \func_num_args();
  $Monad1_1_0 = (($dictMonadTell_0)['Monad1'])($GLOBALS['Prim_undefined']);
  $Semigroup0_2_1 = (($dictMonadTell_0)['Semigroup0'])($GLOBALS['Prim_undefined']);
  $monadExceptT1_3_2 = ["Applicative0" => function($dollar__unused_3 = null) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3 = null) use ($Monad1_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($Monad1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'])['lift'])($Monad1_1_0)))(($dictMonadTell_0)['tell']), "Semigroup0" => function($dollar__unused_4 = null) use ($Semigroup0_2_1) {
  $__num = \func_num_args();
  $__res = $Semigroup0_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($dollar__unused_4 = null) use ($monadExceptT1_3_2) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadWriterExceptT
$GLOBALS['Control_Monad_Except_Trans_monadWriterExceptT'] = function($dictMonadWriter_0 = null) {
  $__num = \func_num_args();
  $MonadTell1_1_0 = (($dictMonadWriter_0)['MonadTell1'])($GLOBALS['Prim_undefined']);
  $Monad1_2_1 = (($MonadTell1_1_0)['Monad1'])($GLOBALS['Prim_undefined']);
  $__local_var_3_2 = (($Monad1_2_1)['Bind1'])($GLOBALS['Prim_undefined']);
  $__local_var_4_3 = (($Monad1_2_1)['Applicative0'])($GLOBALS['Prim_undefined']);
  $Monoid0_5_4 = (($dictMonadWriter_0)['Monoid0'])($GLOBALS['Prim_undefined']);
  $monadTellExceptT1_6_5 = ($GLOBALS['Control_Monad_Except_Trans_monadTellExceptT'])($MonadTell1_1_0);
  $__res = ["listen" => function($v_7 = null) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_2)['bind'])((($dictMonadWriter_0)['listen'])($v_7)))(function($v_8 = null) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ((is_object(($v_8)->{'value0'}) && ((($v_8)->{'value0'})->{'tag'} === "Left"))) {
$__t6 = new Phpurs_Data1("Left", (($v_8)->{'value0'})->{'value0'});
goto end_branch_6;;
};
  if ((is_object(($v_8)->{'value0'}) && ((($v_8)->{'value0'})->{'tag'} === "Right"))) {
$__t6 = new Phpurs_Data1("Right", new Phpurs_Data2("Tuple", (($v_8)->{'value0'})->{'value0'}, ($v_8)->{'value1'}));
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = (($__local_var_4_3)['pure'])($__t6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pass" => function($v_7 = null) use ($__local_var_3_2, $__local_var_4_3, $dictMonadWriter_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadWriter_0)['pass'])(((($__local_var_3_2)['bind'])($v_7))(function($a_8 = null) use ($__local_var_4_3) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ((is_object($a_8) && (($a_8)->{'tag'} === "Left"))) {
$__t7 = new Phpurs_Data2("Tuple", new Phpurs_Data1("Left", ($a_8)->{'value0'}), ($GLOBALS['Control_Category_categoryFn'])['identity']);
goto end_branch_7;;
};
  if ((is_object($a_8) && (($a_8)->{'tag'} === "Right"))) {
$__t7 = new Phpurs_Data2("Tuple", new Phpurs_Data1("Right", (($a_8)->{'value0'})->{'value0'}), (($a_8)->{'value0'})->{'value1'});
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = (($__local_var_4_3)['pure'])($__t7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monoid0" => function($dollar__unused_7 = null) use ($Monoid0_5_4) {
  $__num = \func_num_args();
  $__res = $Monoid0_5_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($dollar__unused_7 = null) use ($monadTellExceptT1_6_5) {
  $__num = \func_num_args();
  $__res = $monadTellExceptT1_6_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadThrowExceptT
$GLOBALS['Control_Monad_Except_Trans_monadThrowExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $monadExceptT1_1_0 = ["Applicative0" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_1 = null) use ($dictMonad_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($dictMonad_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["throwError" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Except_Trans_ExceptT']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure']))($GLOBALS['Data_Either_Left'])), "Monad0" => function($dollar__unused_2 = null) use ($monadExceptT1_1_0) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadErrorExceptT
$GLOBALS['Control_Monad_Except_Trans_monadErrorExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $pure_1_0 = ((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'];
  $monadThrowExceptT1_2_1 = ($GLOBALS['Control_Monad_Except_Trans_monadThrowExceptT'])($dictMonad_0);
  $__res = ["catchError" => (function() use ($dictMonad_0, $pure_1_0) {
  $__fn = function($v_3 = null, $k_4 = null) use ($dictMonad_0, $pure_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_5_2 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_1_0))($GLOBALS['Data_Either_Right']);
  $__res = ((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])($v_3))(function($v2_6 = null) use ($__local_var_5_2, $k_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object($v2_6) && (($v2_6)->{'tag'} === "Left"))) {
$__t3 = ($k_4)(($v2_6)->{'value0'});
goto end_branch_3;;
};
  if ((is_object($v2_6) && (($v2_6)->{'tag'} === "Right"))) {
$__t3 = ($__local_var_5_2)(($v2_6)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($dollar__unused_3 = null) use ($monadThrowExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadThrowExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monadSTExceptT
$GLOBALS['Control_Monad_Except_Trans_monadSTExceptT'] = function($dictMonadST_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadST_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $monadExceptT1_2_1 = ["Applicative0" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_2 = null) use ($Monad0_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($Monad0_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Monad_Except_Trans_monadTransExceptT'])['lift'])($Monad0_1_0)))(($dictMonadST_0)['liftST']), "Monad0" => function($dollar__unused_3 = null) use ($monadExceptT1_2_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Except_Trans_monoidExceptT
$GLOBALS['Control_Monad_Except_Trans_monoidExceptT'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $semigroupExceptT1_1_0 = ($GLOBALS['Control_Monad_Except_Trans_semigroupExceptT'])($dictMonad_0);
  $__res = function($dictMonoid_2 = null) use ($dictMonad_0, $semigroupExceptT1_1_0) {
  $__num = \func_num_args();
  $semigroupExceptT2_3_1 = ($semigroupExceptT1_1_0)((($dictMonoid_2)['Semigroup0'])($GLOBALS['Prim_undefined']));
  $__res = ["mempty" => ((($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_0))['pure'])(($dictMonoid_2)['mempty']), "Semigroup0" => function($dollar__unused_4 = null) use ($semigroupExceptT2_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupExceptT2_3_1;
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

// Control_Monad_Except_Trans_altExceptT
$GLOBALS['Control_Monad_Except_Trans_altExceptT'] = (function() {
  $__fn = function($dictSemigroup_0 = null, $dictMonad_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Bind1_2_0 = (($dictMonad_1)['Bind1'])($GLOBALS['Prim_undefined']);
  $__local_var_3_1 = (($dictMonad_1)['Applicative0'])($GLOBALS['Prim_undefined']);
  $__local_var_4_2 = (((($Bind1_2_0)['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']);
  $functorExceptT1_5_3 = ["map" => function($f_5 = null) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_4_2)['map'])(function($m_6 = null) use ($f_5) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((is_object($m_6) && (($m_6)->{'tag'} === "Left"))) {
$__t3 = new Phpurs_Data1("Left", ($m_6)->{'value0'});
goto end_branch_3;;
};
  if ((is_object($m_6) && (($m_6)->{'tag'} === "Right"))) {
$__t3 = new Phpurs_Data1("Right", ($f_5)(($m_6)->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["alt" => (function() use ($Bind1_2_0, $__local_var_3_1, $dictSemigroup_0) {
  $__fn = function($v_6 = null, $v1_7 = null) use ($Bind1_2_0, $__local_var_3_1, $dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Bind1_2_0)['bind'])($v_6))(function($rm_8 = null) use ($Bind1_2_0, $__local_var_3_1, $dictSemigroup_0, $v1_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ((is_object($rm_8) && (($rm_8)->{'tag'} === "Right"))) {
$__t5 = (($__local_var_3_1)['pure'])(new Phpurs_Data1("Right", ($rm_8)->{'value0'}));
goto end_branch_5;;
};
  if ((is_object($rm_8) && (($rm_8)->{'tag'} === "Left"))) {
$__local_var_9_6 = ($rm_8)->{'value0'};
$__t5 = ((($Bind1_2_0)['bind'])($v1_7))(function($rn_10 = null) use ($__local_var_3_1, $__local_var_9_6, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__t7 = null;;
  if ((is_object($rn_10) && (($rn_10)->{'tag'} === "Right"))) {
$__t7 = (($__local_var_3_1)['pure'])(new Phpurs_Data1("Right", ($rn_10)->{'value0'}));
goto end_branch_7;;
};
  if ((is_object($rn_10) && (($rn_10)->{'tag'} === "Left"))) {
$__t7 = (($__local_var_3_1)['pure'])(new Phpurs_Data1("Left", ((($dictSemigroup_0)['append'])($__local_var_9_6))(($rn_10)->{'value0'})));
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_6 = null) use ($functorExceptT1_5_3) {
  $__num = \func_num_args();
  $__res = $functorExceptT1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Except_Trans_plusExceptT
$GLOBALS['Control_Monad_Except_Trans_plusExceptT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $altExceptT1_2_1 = ($GLOBALS['Control_Monad_Except_Trans_altExceptT'])((($dictMonoid_0)['Semigroup0'])($GLOBALS['Prim_undefined']));
  $__res = function($dictMonad_3 = null) use ($altExceptT1_2_1, $mempty_1_0) {
  $__num = \func_num_args();
  $altExceptT2_4_2 = ($altExceptT1_2_1)($dictMonad_3);
  $__res = ["empty" => ((($GLOBALS['Control_Monad_Except_Trans_monadThrowExceptT'])($dictMonad_3))['throwError'])($mempty_1_0), "Alt0" => function($dollar__unused_5 = null) use ($altExceptT2_4_2) {
  $__num = \func_num_args();
  $__res = $altExceptT2_4_2;
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

// Control_Monad_Except_Trans_alternativeExceptT
$GLOBALS['Control_Monad_Except_Trans_alternativeExceptT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $plusExceptT1_1_0 = ($GLOBALS['Control_Monad_Except_Trans_plusExceptT'])($dictMonoid_0);
  $__res = function($dictMonad_2 = null) use ($plusExceptT1_1_0) {
  $__num = \func_num_args();
  $applicativeExceptT1_3_1 = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_2);
  $plusExceptT2_4_2 = ($plusExceptT1_1_0)($dictMonad_2);
  $__res = ["Applicative0" => function($dollar__unused_5 = null) use ($applicativeExceptT1_3_1) {
  $__num = \func_num_args();
  $__res = $applicativeExceptT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_5 = null) use ($plusExceptT2_4_2) {
  $__num = \func_num_args();
  $__res = $plusExceptT2_4_2;
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

// Control_Monad_Except_Trans_monadPlusExceptT
$GLOBALS['Control_Monad_Except_Trans_monadPlusExceptT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $alternativeExceptT1_1_0 = ($GLOBALS['Control_Monad_Except_Trans_alternativeExceptT'])($dictMonoid_0);
  $__res = function($dictMonad_2 = null) use ($alternativeExceptT1_1_0) {
  $__num = \func_num_args();
  $monadExceptT1_3_1 = ["Applicative0" => function($dollar__unused_3 = null) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_applicativeExceptT'])($dictMonad_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3 = null) use ($dictMonad_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_Except_Trans_bindExceptT'])($dictMonad_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $alternativeExceptT2_4_2 = ($alternativeExceptT1_1_0)($dictMonad_2);
  $__res = ["Monad0" => function($dollar__unused_5 = null) use ($monadExceptT1_3_1) {
  $__num = \func_num_args();
  $__res = $monadExceptT1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_5 = null) use ($alternativeExceptT2_4_2) {
  $__num = \func_num_args();
  $__res = $alternativeExceptT2_4_2;
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

