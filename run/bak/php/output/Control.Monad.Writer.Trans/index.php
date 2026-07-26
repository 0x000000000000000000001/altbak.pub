<?php

namespace Control\Monad\Writer\Trans;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.Monad.Writer.Trans, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Cont.Class, Control.Monad.Error.Class, Control.Monad.Reader.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.State.Class, Control.Monad.Trans.Class, Control.Monad.Writer.Class, Control.Monad.Writer.Trans, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Function, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Effect.Class, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Cont.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Reader.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.State.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Trans.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Writer.Trans/index.php';
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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


// Control_Monad_Writer_Trans_WriterT
$GLOBALS['Control_Monad_Writer_Trans_WriterT'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Writer_Trans_runWriterT
$GLOBALS['Control_Monad_Writer_Trans_runWriterT'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $v_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Writer_Trans_newtypeWriterT
$GLOBALS['Control_Monad_Writer_Trans_newtypeWriterT'] = ["Coercible0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_Writer_Trans_monadTransWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadTransWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $__res = ["lift" => (function() use ($mempty_1_0) {
  $__fn = function($dictMonad_2 = null, $m_3 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonad_2)['Bind1'])(null))['bind'])($m_3))(function($a_4 = null) use ($dictMonad_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_2)['Applicative0'])(null))['pure'])(new Phpurs_Data2("Tuple", $a_4, $mempty_1_0));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Writer_Trans_mapWriterT
$GLOBALS['Control_Monad_Writer_Trans_mapWriterT'] = (function() {
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

// Control_Monad_Writer_Trans_functorWriterT
$GLOBALS['Control_Monad_Writer_Trans_functorWriterT'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => function($f_1 = null) use ($dictFunctor_0) {
  $__num = \func_num_args();
  $__res = (($dictFunctor_0)['map'])(function($v_2 = null) use ($f_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_1)(($v_2)->{'value0'}), ($v_2)->{'value1'});
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

// Control_Monad_Writer_Trans_execWriterT
$GLOBALS['Control_Monad_Writer_Trans_execWriterT'] = (function() {
  $__fn = function($dictFunctor_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)['map'])($GLOBALS['Data_Tuple_snd']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Writer_Trans_applyWriterT
$GLOBALS['Control_Monad_Writer_Trans_applyWriterT'] = (function() {
  $__fn = function($dictSemigroup_0 = null, $dictApply_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Functor0_2_0 = (($dictApply_1)['Functor0'])(null);
  $functorWriterT1_3_1 = ["map" => function($f_3 = null) use ($Functor0_2_0) {
  $__num = \func_num_args();
  $__res = (($Functor0_2_0)['map'])(function($v_4 = null) use ($f_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_3)(($v_4)->{'value0'}), ($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["apply" => (function() use ($Functor0_2_0, $dictApply_1, $dictSemigroup_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($Functor0_2_0, $dictApply_1, $dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictApply_1)['apply'])(((($Functor0_2_0)['map'])((function() use ($dictSemigroup_0) {
  $__fn = function($v3_6 = null, $v4_7 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", (($v3_6)->{'value0'})(($v4_7)->{'value0'}), ((($dictSemigroup_0)['append'])(($v3_6)->{'value1'}))(($v4_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($v_4)))($v1_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4 = null) use ($functorWriterT1_3_1) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_3_1;
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

// Control_Monad_Writer_Trans_bindWriterT
$GLOBALS['Control_Monad_Writer_Trans_bindWriterT'] = (function() {
  $__fn = function($dictSemigroup_0 = null, $dictBind_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Apply0_2_0 = (($dictBind_1)['Apply0'])(null);
  $Functor0_3_1 = (($Apply0_2_0)['Functor0'])(null);
  $functorWriterT1_4_2 = ["map" => function($f_4 = null) use ($Functor0_3_1) {
  $__num = \func_num_args();
  $__res = (($Functor0_3_1)['map'])(function($v_5 = null) use ($f_4) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_4)(($v_5)->{'value0'}), ($v_5)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_5_3 = ["apply" => (function() use ($Apply0_2_0, $Functor0_3_1, $dictSemigroup_0) {
  $__fn = function($v_5 = null, $v1_6 = null) use ($Apply0_2_0, $Functor0_3_1, $dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Apply0_2_0)['apply'])(((($Functor0_3_1)['map'])((function() use ($dictSemigroup_0) {
  $__fn = function($v3_7 = null, $v4_8 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", (($v3_7)->{'value0'})(($v4_8)->{'value0'}), ((($dictSemigroup_0)['append'])(($v3_7)->{'value1'}))(($v4_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($v_5)))($v1_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_5 = null) use ($functorWriterT1_4_2) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["bind" => (function() use ($Apply0_2_0, $dictBind_1, $dictSemigroup_0) {
  $__fn = function($v_6 = null, $k_7 = null) use ($Apply0_2_0, $dictBind_1, $dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictBind_1)['bind'])($v_6))(function($v1_8 = null) use ($Apply0_2_0, $dictSemigroup_0, $k_7) {
  $__num = \func_num_args();
  $__local_var_9_4 = ($v1_8)->{'value1'};
  $__res = ((((($Apply0_2_0)['Functor0'])(null))['map'])(function($v3_10 = null) use ($__local_var_9_4, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($v3_10)->{'value0'}, ((($dictSemigroup_0)['append'])($__local_var_9_4))(($v3_10)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($k_7)(($v1_8)->{'value0'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_6 = null) use ($applyWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_5_3;
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

// Control_Monad_Writer_Trans_semigroupWriterT
$GLOBALS['Control_Monad_Writer_Trans_semigroupWriterT'] = (function() {
  $__fn = function($dictApply_0 = null, $dictSemigroup_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Functor0_2_0 = (($dictApply_0)['Functor0'])(null);
  $__res = function($dictSemigroup1_3 = null) use ($Functor0_2_0, $dictApply_0, $dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = ["append" => (function() use ($Functor0_2_0, $dictApply_0, $dictSemigroup1_3, $dictSemigroup_1) {
  $__fn = function($a_4 = null, $b_5 = null) use ($Functor0_2_0, $dictApply_0, $dictSemigroup1_3, $dictSemigroup_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictApply_0)['apply'])(((($Functor0_2_0)['map'])((function() use ($dictSemigroup_1) {
  $__fn = function($v3_6 = null, $v4_7 = null) use ($dictSemigroup_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", (($v3_6)->{'value0'})(($v4_7)->{'value0'}), ((($dictSemigroup_1)['append'])(($v3_6)->{'value1'}))(($v4_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(((($Functor0_2_0)['map'])(function($v_6 = null) use ($dictSemigroup1_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", (($dictSemigroup1_3)['append'])(($v_6)->{'value0'}), ($v_6)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_4))))($b_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Control_Monad_Writer_Trans_applicativeWriterT
$GLOBALS['Control_Monad_Writer_Trans_applicativeWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $__local_var_2_1 = (($dictMonoid_0)['Semigroup0'])(null);
  $__res = function($dictApplicative_3 = null) use ($__local_var_2_1, $mempty_1_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictApplicative_3)['Apply0'])(null);
  $Functor0_5_3 = (($__local_var_4_2)['Functor0'])(null);
  $functorWriterT1_6_4 = ["map" => function($f_6 = null) use ($Functor0_5_3) {
  $__num = \func_num_args();
  $__res = (($Functor0_5_3)['map'])(function($v_7 = null) use ($f_6) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_6)(($v_7)->{'value0'}), ($v_7)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $applyWriterT2_6_4 = ["apply" => (function() use ($Functor0_5_3, $__local_var_2_1, $__local_var_4_2) {
  $__fn = function($v_7 = null, $v1_8 = null) use ($Functor0_5_3, $__local_var_2_1, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_4_2)['apply'])(((($Functor0_5_3)['map'])((function() use ($__local_var_2_1) {
  $__fn = function($v3_9 = null, $v4_10 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", (($v3_9)->{'value0'})(($v4_10)->{'value0'}), ((($__local_var_2_1)['append'])(($v3_9)->{'value1'}))(($v4_10)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($v_7)))($v1_8);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_7 = null) use ($functorWriterT1_6_4) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["pure" => function($a_7 = null) use ($dictApplicative_3, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_3)['pure'])(new Phpurs_Data2("Tuple", $a_7, $mempty_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_7 = null) use ($applyWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $applyWriterT2_6_4;
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

// Control_Monad_Writer_Trans_monadWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $applicativeWriterT1_1_0 = ($GLOBALS['Control_Monad_Writer_Trans_applicativeWriterT'])($dictMonoid_0);
  $bindWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_bindWriterT'])((($dictMonoid_0)['Semigroup0'])(null));
  $__res = function($dictMonad_3 = null) use ($applicativeWriterT1_1_0, $bindWriterT1_2_1) {
  $__num = \func_num_args();
  $applicativeWriterT2_4_2 = ($applicativeWriterT1_1_0)((($dictMonad_3)['Applicative0'])(null));
  $bindWriterT2_5_3 = ($bindWriterT1_2_1)((($dictMonad_3)['Bind1'])(null));
  $__res = ["Applicative0" => function($_dollar__unused_6 = null) use ($applicativeWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_6 = null) use ($bindWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $bindWriterT2_5_3;
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

// Control_Monad_Writer_Trans_monadAskWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadAskWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $monadWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonadAsk_3 = null) use ($mempty_1_0, $monadWriterT1_2_1) {
  $__num = \func_num_args();
  $Monad0_4_2 = (($dictMonadAsk_3)['Monad0'])(null);
  $monadWriterT2_5_3 = ($monadWriterT1_2_1)($Monad0_4_2);
  $__res = ["ask" => ((((($Monad0_4_2)['Bind1'])(null))['bind'])(($dictMonadAsk_3)['ask']))(function($a_6 = null) use ($Monad0_4_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (((($Monad0_4_2)['Applicative0'])(null))['pure'])(new Phpurs_Data2("Tuple", $a_6, $mempty_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Monad0" => function($_dollar__unused_6 = null) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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

// Control_Monad_Writer_Trans_monadReaderWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadReaderWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $monadAskWriterT1_1_0 = ($GLOBALS['Control_Monad_Writer_Trans_monadAskWriterT'])($dictMonoid_0);
  $__res = function($dictMonadReader_2 = null) use ($monadAskWriterT1_1_0) {
  $__num = \func_num_args();
  $monadAskWriterT2_3_1 = ($monadAskWriterT1_1_0)((($dictMonadReader_2)['MonadAsk0'])(null));
  $__res = ["local" => function($f_4 = null) use ($dictMonadReader_2) {
  $__num = \func_num_args();
  $__res = (($dictMonadReader_2)['local'])($f_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadAsk0" => function($_dollar__unused_4 = null) use ($monadAskWriterT2_3_1) {
  $__num = \func_num_args();
  $__res = $monadAskWriterT2_3_1;
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

// Control_Monad_Writer_Trans_monadContWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadContWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $monadWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonadCont_3 = null) use ($mempty_1_0, $monadWriterT1_2_1) {
  $__num = \func_num_args();
  $monadWriterT2_4_2 = ($monadWriterT1_2_1)((($dictMonadCont_3)['Monad0'])(null));
  $__res = ["callCC" => function($f_5 = null) use ($dictMonadCont_3, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (($dictMonadCont_3)['callCC'])(function($c_6 = null) use ($f_5, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ($f_5)(function($a_7 = null) use ($c_6, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ($c_6)(new Phpurs_Data2("Tuple", $a_7, $mempty_1_0));
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
}, "Monad0" => function($_dollar__unused_5 = null) use ($monadWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_4_2;
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

// Control_Monad_Writer_Trans_monadEffectWriter
$GLOBALS['Control_Monad_Writer_Trans_monadEffectWriter'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $monadWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonadEffect_3 = null) use ($mempty_1_0, $monadWriterT1_2_1) {
  $__num = \func_num_args();
  $Monad0_4_2 = (($dictMonadEffect_3)['Monad0'])(null);
  $monadWriterT2_5_3 = ($monadWriterT1_2_1)($Monad0_4_2);
  $__res = ["liftEffect" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_6 = null) use ($Monad0_4_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ((((($Monad0_4_2)['Bind1'])(null))['bind'])($m_6))(function($a_7 = null) use ($Monad0_4_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (((($Monad0_4_2)['Applicative0'])(null))['pure'])(new Phpurs_Data2("Tuple", $a_7, $mempty_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadEffect_3)['liftEffect']), "Monad0" => function($_dollar__unused_6 = null) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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

// Control_Monad_Writer_Trans_monadRecWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadRecWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonoid_0)['Semigroup0'])(null);
  $mempty_2_1 = ($dictMonoid_0)['mempty'];
  $monadWriterT1_3_2 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonadRec_4 = null) use ($__local_var_1_0, $mempty_2_1, $monadWriterT1_3_2) {
  $__num = \func_num_args();
  $Monad0_5_3 = (($dictMonadRec_4)['Monad0'])(null);
  $monadWriterT2_6_4 = ($monadWriterT1_3_2)($Monad0_5_3);
  $__res = ["tailRecM" => (function() use ($Monad0_5_3, $__local_var_1_0, $dictMonadRec_4, $mempty_2_1) {
  $__fn = function($f_7 = null, $a_8 = null) use ($Monad0_5_3, $__local_var_1_0, $dictMonadRec_4, $mempty_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictMonadRec_4)['tailRecM'])(function($v_9 = null) use ($Monad0_5_3, $__local_var_1_0, $f_7) {
  $__num = \func_num_args();
  $__local_var_10_5 = ($v_9)->{'value1'};
  $__res = ((((($Monad0_5_3)['Bind1'])(null))['bind'])(($f_7)(($v_9)->{'value0'})))(function($v2_11 = null) use ($Monad0_5_3, $__local_var_1_0, $__local_var_10_5) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ((is_object(($v2_11)->{'value0'}) && ((($v2_11)->{'value0'})->{'tag'} === "Loop"))) {
$__t6 = new Phpurs_Data1("Loop", new Phpurs_Data2("Tuple", (($v2_11)->{'value0'})->{'value0'}, ((($__local_var_1_0)['append'])($__local_var_10_5))(($v2_11)->{'value1'})));
goto end_branch_6;;
};
  if ((is_object(($v2_11)->{'value0'}) && ((($v2_11)->{'value0'})->{'tag'} === "Done"))) {
$__t6 = new Phpurs_Data1("Done", new Phpurs_Data2("Tuple", (($v2_11)->{'value0'})->{'value0'}, ((($__local_var_1_0)['append'])($__local_var_10_5))(($v2_11)->{'value1'})));
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = (((($Monad0_5_3)['Applicative0'])(null))['pure'])($__t6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", $a_8, $mempty_2_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($_dollar__unused_7 = null) use ($monadWriterT2_6_4) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_6_4;
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

// Control_Monad_Writer_Trans_monadStateWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadStateWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $monadWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonadState_3 = null) use ($mempty_1_0, $monadWriterT1_2_1) {
  $__num = \func_num_args();
  $Monad0_4_2 = (($dictMonadState_3)['Monad0'])(null);
  $monadWriterT2_5_3 = ($monadWriterT1_2_1)($Monad0_4_2);
  $__res = ["state" => function($f_6 = null) use ($Monad0_4_2, $dictMonadState_3, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ((((($Monad0_4_2)['Bind1'])(null))['bind'])((($dictMonadState_3)['state'])($f_6)))(function($a_7 = null) use ($Monad0_4_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (((($Monad0_4_2)['Applicative0'])(null))['pure'])(new Phpurs_Data2("Tuple", $a_7, $mempty_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar__unused_6 = null) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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

// Control_Monad_Writer_Trans_monadTellWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadTellWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $Semigroup0_1_0 = (($dictMonoid_0)['Semigroup0'])(null);
  $monadWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonad_3 = null) use ($Semigroup0_1_0, $monadWriterT1_2_1) {
  $__num = \func_num_args();
  $monadWriterT2_4_2 = ($monadWriterT1_2_1)($dictMonad_3);
  $__res = ["tell" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Control_Monad_Writer_Trans_WriterT']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictMonad_3)['Applicative0'])(null))['pure']))(($GLOBALS['Data_Tuple_Tuple'])($GLOBALS['Data_Unit_unit']))), "Semigroup0" => function($_dollar__unused_5 = null) use ($Semigroup0_1_0) {
  $__num = \func_num_args();
  $__res = $Semigroup0_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad1" => function($_dollar__unused_5 = null) use ($monadWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_4_2;
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

// Control_Monad_Writer_Trans_monadWriterWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadWriterWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $monadTellWriterT1_1_0 = ($GLOBALS['Control_Monad_Writer_Trans_monadTellWriterT'])($dictMonoid_0);
  $__res = function($dictMonad_2 = null) use ($dictMonoid_0, $monadTellWriterT1_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictMonad_2)['Bind1'])(null);
  $__local_var_4_2 = (($dictMonad_2)['Applicative0'])(null);
  $monadTellWriterT2_5_3 = ($monadTellWriterT1_1_0)($dictMonad_2);
  $__res = ["listen" => function($v_6 = null) use ($__local_var_3_1, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)['bind'])($v_6))(function($v1_7 = null) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_4_2)['pure'])(new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", ($v1_7)->{'value0'}, ($v1_7)->{'value1'}), ($v1_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "pass" => function($v_6 = null) use ($__local_var_3_1, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)['bind'])($v_6))(function($v1_7 = null) use ($__local_var_4_2) {
  $__num = \func_num_args();
  $__res = (($__local_var_4_2)['pure'])(new Phpurs_Data2("Tuple", (($v1_7)->{'value0'})->{'value0'}, ((($v1_7)->{'value0'})->{'value1'})(($v1_7)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monoid0" => function($_dollar__unused_6 = null) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = $dictMonoid_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "MonadTell1" => function($_dollar__unused_6 = null) use ($monadTellWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadTellWriterT2_5_3;
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

// Control_Monad_Writer_Trans_monadThrowWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadThrowWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $monadWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonadThrow_3 = null) use ($mempty_1_0, $monadWriterT1_2_1) {
  $__num = \func_num_args();
  $Monad0_4_2 = (($dictMonadThrow_3)['Monad0'])(null);
  $monadWriterT2_5_3 = ($monadWriterT1_2_1)($Monad0_4_2);
  $__res = ["throwError" => function($e_6 = null) use ($Monad0_4_2, $dictMonadThrow_3, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ((((($Monad0_4_2)['Bind1'])(null))['bind'])((($dictMonadThrow_3)['throwError'])($e_6)))(function($a_7 = null) use ($Monad0_4_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (((($Monad0_4_2)['Applicative0'])(null))['pure'])(new Phpurs_Data2("Tuple", $a_7, $mempty_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar__unused_6 = null) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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

// Control_Monad_Writer_Trans_monadErrorWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadErrorWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $monadThrowWriterT1_1_0 = ($GLOBALS['Control_Monad_Writer_Trans_monadThrowWriterT'])($dictMonoid_0);
  $__res = function($dictMonadError_2 = null) use ($monadThrowWriterT1_1_0) {
  $__num = \func_num_args();
  $monadThrowWriterT2_3_1 = ($monadThrowWriterT1_1_0)((($dictMonadError_2)['MonadThrow0'])(null));
  $__res = ["catchError" => (function() use ($dictMonadError_2) {
  $__fn = function($v_4 = null, $h_5 = null) use ($dictMonadError_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictMonadError_2)['catchError'])($v_4))(function($e_6 = null) use ($h_5) {
  $__num = \func_num_args();
  $__res = ($h_5)($e_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "MonadThrow0" => function($_dollar__unused_4 = null) use ($monadThrowWriterT2_3_1) {
  $__num = \func_num_args();
  $__res = $monadThrowWriterT2_3_1;
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

// Control_Monad_Writer_Trans_monadSTWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadSTWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $monadWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $__res = function($dictMonadST_3 = null) use ($mempty_1_0, $monadWriterT1_2_1) {
  $__num = \func_num_args();
  $Monad0_4_2 = (($dictMonadST_3)['Monad0'])(null);
  $monadWriterT2_5_3 = ($monadWriterT1_2_1)($Monad0_4_2);
  $__res = ["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($m_6 = null) use ($Monad0_4_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ((((($Monad0_4_2)['Bind1'])(null))['bind'])($m_6))(function($a_7 = null) use ($Monad0_4_2, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = (((($Monad0_4_2)['Applicative0'])(null))['pure'])(new Phpurs_Data2("Tuple", $a_7, $mempty_1_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonadST_3)['liftST']), "Monad0" => function($_dollar__unused_6 = null) use ($monadWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_5_3;
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

// Control_Monad_Writer_Trans_monoidWriterT
$GLOBALS['Control_Monad_Writer_Trans_monoidWriterT'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictApplicative_0)['Apply0'])(null);
  $__res = function($dictMonoid_2 = null) use ($__local_var_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictMonoid_2)['Semigroup0'])(null);
  $Functor0_4_2 = (($__local_var_1_0)['Functor0'])(null);
  $__res = function($dictMonoid1_5 = null) use ($Functor0_4_2, $__local_var_1_0, $__local_var_3_1, $dictApplicative_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__local_var_6_3 = (($dictMonoid1_5)['Semigroup0'])(null);
  $semigroupWriterT3_7_4 = ["append" => (function() use ($Functor0_4_2, $__local_var_1_0, $__local_var_3_1, $__local_var_6_3) {
  $__fn = function($a_7 = null, $b_8 = null) use ($Functor0_4_2, $__local_var_1_0, $__local_var_3_1, $__local_var_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['apply'])(((($Functor0_4_2)['map'])((function() use ($__local_var_3_1) {
  $__fn = function($v3_9 = null, $v4_10 = null) use ($__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", (($v3_9)->{'value0'})(($v4_10)->{'value0'}), ((($__local_var_3_1)['append'])(($v3_9)->{'value1'}))(($v4_10)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(((($Functor0_4_2)['map'])(function($v_9 = null) use ($__local_var_6_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", (($__local_var_6_3)['append'])(($v_9)->{'value0'}), ($v_9)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_7))))($b_8);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["mempty" => (((($GLOBALS['Control_Monad_Writer_Trans_applicativeWriterT'])($dictMonoid_2))($dictApplicative_0))['pure'])(($dictMonoid1_5)['mempty']), "Semigroup0" => function($_dollar__unused_8 = null) use ($semigroupWriterT3_7_4) {
  $__num = \func_num_args();
  $__res = $semigroupWriterT3_7_4;
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Writer_Trans_altWriterT
$GLOBALS['Control_Monad_Writer_Trans_altWriterT'] = function($dictAlt_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictAlt_0)['Functor0'])(null);
  $functorWriterT1_2_1 = ["map" => function($f_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)['map'])(function($v_3 = null) use ($f_2) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_2)(($v_3)->{'value0'}), ($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["alt" => (function() use ($dictAlt_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictAlt_0)['alt'])($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_3 = null) use ($functorWriterT1_2_1) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Writer_Trans_plusWriterT
$GLOBALS['Control_Monad_Writer_Trans_plusWriterT'] = function($dictPlus_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictPlus_0)['Alt0'])(null);
  $__local_var_2_1 = (($__local_var_1_0)['Functor0'])(null);
  $functorWriterT1_3_2 = ["map" => function($f_3 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)['map'])(function($v_4 = null) use ($f_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_3)(($v_4)->{'value0'}), ($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altWriterT1_3_2 = ["alt" => (function() use ($__local_var_1_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['alt'])($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_4 = null) use ($functorWriterT1_3_2) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["empty" => ($dictPlus_0)['empty'], "Alt0" => function($_dollar__unused_4 = null) use ($altWriterT1_3_2) {
  $__num = \func_num_args();
  $__res = $altWriterT1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Control_Monad_Writer_Trans_alternativeWriterT
$GLOBALS['Control_Monad_Writer_Trans_alternativeWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $applicativeWriterT1_1_0 = ($GLOBALS['Control_Monad_Writer_Trans_applicativeWriterT'])($dictMonoid_0);
  $__res = function($dictAlternative_2 = null) use ($applicativeWriterT1_1_0) {
  $__num = \func_num_args();
  $applicativeWriterT2_3_1 = ($applicativeWriterT1_1_0)((($dictAlternative_2)['Applicative0'])(null));
  $__local_var_4_2 = (($dictAlternative_2)['Plus1'])(null);
  $__local_var_5_3 = (($__local_var_4_2)['Alt0'])(null);
  $__local_var_6_4 = (($__local_var_5_3)['Functor0'])(null);
  $functorWriterT1_7_5 = ["map" => function($f_7 = null) use ($__local_var_6_4) {
  $__num = \func_num_args();
  $__res = (($__local_var_6_4)['map'])(function($v_8 = null) use ($f_7) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($f_7)(($v_8)->{'value0'}), ($v_8)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $altWriterT1_8_6 = ["alt" => (function() use ($__local_var_5_3) {
  $__fn = function($v_8 = null, $v1_9 = null) use ($__local_var_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_5_3)['alt'])($v_8))($v1_9);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_8 = null) use ($functorWriterT1_7_5) {
  $__num = \func_num_args();
  $__res = $functorWriterT1_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $plusWriterT1_6_4 = ["empty" => ($__local_var_4_2)['empty'], "Alt0" => function($_dollar__unused_9 = null) use ($altWriterT1_8_6) {
  $__num = \func_num_args();
  $__res = $altWriterT1_8_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["Applicative0" => function($_dollar__unused_7 = null) use ($applicativeWriterT2_3_1) {
  $__num = \func_num_args();
  $__res = $applicativeWriterT2_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_7 = null) use ($plusWriterT1_6_4) {
  $__num = \func_num_args();
  $__res = $plusWriterT1_6_4;
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

// Control_Monad_Writer_Trans_monadPlusWriterT
$GLOBALS['Control_Monad_Writer_Trans_monadPlusWriterT'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $monadWriterT1_1_0 = ($GLOBALS['Control_Monad_Writer_Trans_monadWriterT'])($dictMonoid_0);
  $alternativeWriterT1_2_1 = ($GLOBALS['Control_Monad_Writer_Trans_alternativeWriterT'])($dictMonoid_0);
  $__res = function($dictMonadPlus_3 = null) use ($alternativeWriterT1_2_1, $monadWriterT1_1_0) {
  $__num = \func_num_args();
  $monadWriterT2_4_2 = ($monadWriterT1_1_0)((($dictMonadPlus_3)['Monad0'])(null));
  $alternativeWriterT2_5_3 = ($alternativeWriterT1_2_1)((($dictMonadPlus_3)['Alternative1'])(null));
  $__res = ["Monad0" => function($_dollar__unused_6 = null) use ($monadWriterT2_4_2) {
  $__num = \func_num_args();
  $__res = $monadWriterT2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar__unused_6 = null) use ($alternativeWriterT2_5_3) {
  $__num = \func_num_args();
  $__res = $alternativeWriterT2_5_3;
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

