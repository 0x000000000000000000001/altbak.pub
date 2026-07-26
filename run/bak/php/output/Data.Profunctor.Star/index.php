<?php

namespace Data\Profunctor\Star;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Either, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Newtype, Data.Profunctor, Data.Profunctor.Choice, Data.Profunctor.Closed, Data.Profunctor.Star, Data.Profunctor.Strong, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Distributive, Data.Either, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Newtype, Data.Profunctor, Data.Profunctor.Choice, Data.Profunctor.Closed, Data.Profunctor.Star, Data.Profunctor.Strong, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Distributive/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Invariant/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Profunctor/index.php';
require_once __DIR__ . '/../Data.Profunctor.Choice/index.php';
require_once __DIR__ . '/../Data.Profunctor.Closed/index.php';
require_once __DIR__ . '/../Data.Profunctor.Star/index.php';
require_once __DIR__ . '/../Data.Profunctor.Strong/index.php';
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


// Data_Profunctor_Star_Star
$GLOBALS['Data_Profunctor_Star_Star'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_semigroupoidStar
$GLOBALS['Data_Profunctor_Star_semigroupoidStar'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $__res = ["compose" => (function() use ($dictBind_0) {
  $__fn = function($v_1 = null, $v1_2 = null, $x_3 = null) use ($dictBind_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictBind_0)['bind'])(($v1_2)($x_3)))($v_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_profunctorStar
$GLOBALS['Data_Profunctor_Star_profunctorStar'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["dimap" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1 = null, $g_2 = null, $v_3 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictFunctor_0)['map'])($g_2)))($v_3)))($f_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_strongStar
$GLOBALS['Data_Profunctor_Star_strongStar'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $profunctorStar1_1_0 = ($GLOBALS['Data_Profunctor_Star_profunctorStar'])($dictFunctor_0);
  $__res = ["first" => (function() use ($dictFunctor_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_4_1 = ($v1_3)->{'value1'};
  $__res = ((($dictFunctor_0)['map'])(function($v2_5 = null) use ($__local_var_4_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", $v2_5, $__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_2)(($v1_3)->{'value0'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "second" => (function() use ($dictFunctor_0) {
  $__fn = function($v_2 = null, $v1_3 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctor_0)['map'])(($GLOBALS['Data_Tuple_Tuple'])(($v1_3)->{'value0'})))(($v_2)(($v1_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Profunctor0" => function($dollar__unused_2 = null) use ($profunctorStar1_1_0) {
  $__num = \func_num_args();
  $__res = $profunctorStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_newtypeStar
$GLOBALS['Data_Profunctor_Star_newtypeStar'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Profunctor_Star_invariantStar
$GLOBALS['Data_Profunctor_Star_invariantStar'] = function($dictInvariant_0 = null) {
  $__num = \func_num_args();
  $__res = ["imap" => (function() use ($dictInvariant_0) {
  $__fn = function($f_1 = null, $g_2 = null, $v_3 = null) use ($dictInvariant_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($dictInvariant_0)['imap'])($f_1))($g_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_hoistStar
$GLOBALS['Data_Profunctor_Star_hoistStar'] = (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Profunctor_Star_functorStar
$GLOBALS['Data_Profunctor_Star_functorStar'] = function($dictFunctor_0 = null) {
  $__num = \func_num_args();
  $__res = ["map" => (function() use ($dictFunctor_0) {
  $__fn = function($f_1 = null, $v_2 = null) use ($dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictFunctor_0)['map'])($f_1)))($v_2);
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

// Data_Profunctor_Star_distributiveStar
$GLOBALS['Data_Profunctor_Star_distributiveStar'] = function($dictDistributive_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_Profunctor_Star_distributiveStar_dictDistributive_0 = $dictDistributive_0;
  tco_loop_Data_Profunctor_Star_distributiveStar:;
  $dictDistributive_0 = $__tco_var_Data_Profunctor_Star_distributiveStar_dictDistributive_0;
  $__local_var_1_0 = (($dictDistributive_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $functorStar1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_1_0)['map'])($f_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["distribute" => function($dictFunctor_3 = null) use ($dictDistributive_0) {
  $__num = \func_num_args();
  $collect1_4_2 = (($dictDistributive_0)['collect'])($dictFunctor_3);
  $__res = (function() use ($collect1_4_2) {
  $__fn = function($f_5 = null, $a_6 = null) use ($collect1_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($collect1_4_2)(function($v_7 = null) use ($a_6) {
  $__num = \func_num_args();
  $__res = ($v_7)($a_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "collect" => (function() use ($dictDistributive_0) {
  $__fn = function($dictFunctor_3 = null, $f_4 = null) use ($dictDistributive_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_Profunctor_Star_distributiveStar'])($dictDistributive_0))['distribute'])($dictFunctor_3)))((($dictFunctor_3)['map'])($f_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_3 = null) use ($functorStar1_2_1) {
  $__num = \func_num_args();
  $__res = $functorStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_closedStar
$GLOBALS['Data_Profunctor_Star_closedStar'] = function($dictDistributive_0 = null) {
  $__num = \func_num_args();
  $distribute_1_0 = (($dictDistributive_0)['distribute'])($GLOBALS['Data_Functor_functorFn']);
  $profunctorStar1_2_1 = ($GLOBALS['Data_Profunctor_Star_profunctorStar'])((($dictDistributive_0)['Functor0'])($GLOBALS['Prim_undefined']));
  $__res = ["closed" => (function() use ($distribute_1_0) {
  $__fn = function($v_3 = null, $g_4 = null) use ($distribute_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($distribute_1_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($v_3))($g_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Profunctor0" => function($dollar__unused_3 = null) use ($profunctorStar1_2_1) {
  $__num = \func_num_args();
  $__res = $profunctorStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_choiceStar
$GLOBALS['Data_Profunctor_Star_choiceStar'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']);
  $pure_2_1 = ($dictApplicative_0)['pure'];
  $profunctorStar1_3_2 = ($GLOBALS['Data_Profunctor_Star_profunctorStar'])($Functor0_1_0);
  $__res = ["left" => function($v_4 = null) use ($Functor0_1_0, $pure_2_1) {
  $__num = \func_num_args();
  $__local_var_5_3 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_1_0)['map'])($GLOBALS['Data_Either_Left'])))($v_4);
  $__local_var_6_4 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_2_1))($GLOBALS['Data_Either_Right']);
  $__res = function($v2_7 = null) use ($__local_var_5_3, $__local_var_6_4) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ((is_object($v2_7) && (($v2_7)->{'tag'} === "Left"))) {
$__t5 = ($__local_var_5_3)(($v2_7)->{'value0'});
goto end_branch_5;;
};
  if ((is_object($v2_7) && (($v2_7)->{'tag'} === "Right"))) {
$__t5 = ($__local_var_6_4)(($v2_7)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "right" => function($v_4 = null) use ($Functor0_1_0, $pure_2_1) {
  $__num = \func_num_args();
  $__local_var_5_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($pure_2_1))($GLOBALS['Data_Either_Left']);
  $__local_var_6_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_1_0)['map'])($GLOBALS['Data_Either_Right'])))($v_4);
  $__res = function($v2_7 = null) use ($__local_var_5_6, $__local_var_6_7) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ((is_object($v2_7) && (($v2_7)->{'tag'} === "Left"))) {
$__t8 = ($__local_var_5_6)(($v2_7)->{'value0'});
goto end_branch_8;;
};
  if ((is_object($v2_7) && (($v2_7)->{'tag'} === "Right"))) {
$__t8 = ($__local_var_6_7)(($v2_7)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Profunctor0" => function($dollar__unused_4 = null) use ($profunctorStar1_3_2) {
  $__num = \func_num_args();
  $__res = $profunctorStar1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_categoryStar
$GLOBALS['Data_Profunctor_Star_categoryStar'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']);
  $semigroupoidStar1_2_1 = ["compose" => (function() use ($__local_var_1_0) {
  $__fn = function($v_2 = null, $v1_3 = null, $x_4 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($__local_var_1_0)['bind'])(($v1_3)($x_4)))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["identity" => ((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'], "Semigroupoid0" => function($dollar__unused_3 = null) use ($semigroupoidStar1_2_1) {
  $__num = \func_num_args();
  $__res = $semigroupoidStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_applyStar
$GLOBALS['Data_Profunctor_Star_applyStar'] = function($dictApply_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictApply_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $functorStar1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_1_0)['map'])($f_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["apply" => (function() use ($dictApply_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $a_5 = null) use ($dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictApply_0)['apply'])(($v_3)($a_5)))(($v1_4)($a_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_3 = null) use ($functorStar1_2_1) {
  $__num = \func_num_args();
  $__res = $functorStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_bindStar
$GLOBALS['Data_Profunctor_Star_bindStar'] = function($dictBind_0 = null) {
  $__num = \func_num_args();
  $applyStar1_1_0 = ($GLOBALS['Data_Profunctor_Star_applyStar'])((($dictBind_0)['Apply0'])($GLOBALS['Prim_undefined']));
  $__res = ["bind" => (function() use ($dictBind_0) {
  $__fn = function($v_2 = null, $f_3 = null, $x_4 = null) use ($dictBind_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictBind_0)['bind'])(($v_2)($x_4)))(function($a_5 = null) use ($f_3, $x_4) {
  $__num = \func_num_args();
  $__res = (($f_3)($a_5))($x_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_2 = null) use ($applyStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applyStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_applicativeStar
$GLOBALS['Data_Profunctor_Star_applicativeStar'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $applyStar1_1_0 = ($GLOBALS['Data_Profunctor_Star_applyStar'])((($dictApplicative_0)['Apply0'])($GLOBALS['Prim_undefined']));
  $__res = ["pure" => (function() use ($dictApplicative_0) {
  $__fn = function($a_2 = null, $v_3 = null) use ($dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictApplicative_0)['pure'])($a_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_2 = null) use ($applyStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applyStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_monadStar
$GLOBALS['Data_Profunctor_Star_monadStar'] = function($dictMonad_0 = null) {
  $__num = \func_num_args();
  $applicativeStar1_1_0 = ($GLOBALS['Data_Profunctor_Star_applicativeStar'])((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']));
  $bindStar1_2_1 = ($GLOBALS['Data_Profunctor_Star_bindStar'])((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']));
  $__res = ["Applicative0" => function($dollar__unused_3 = null) use ($applicativeStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_3 = null) use ($bindStar1_2_1) {
  $__num = \func_num_args();
  $__res = $bindStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_altStar
$GLOBALS['Data_Profunctor_Star_altStar'] = function($dictAlt_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictAlt_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $functorStar1_2_1 = ["map" => (function() use ($__local_var_1_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_1_0)['map'])($f_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["alt" => (function() use ($dictAlt_0) {
  $__fn = function($v_3 = null, $v1_4 = null, $a_5 = null) use ($dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictAlt_0)['alt'])(($v_3)($a_5)))(($v1_4)($a_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_3 = null) use ($functorStar1_2_1) {
  $__num = \func_num_args();
  $__res = $functorStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_plusStar
$GLOBALS['Data_Profunctor_Star_plusStar'] = function($dictPlus_0 = null) {
  $__num = \func_num_args();
  $empty_1_0 = ($dictPlus_0)['empty'];
  $altStar1_2_1 = ($GLOBALS['Data_Profunctor_Star_altStar'])((($dictPlus_0)['Alt0'])($GLOBALS['Prim_undefined']));
  $__res = ["empty" => function($v_3 = null) use ($empty_1_0) {
  $__num = \func_num_args();
  $__res = $empty_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alt0" => function($dollar__unused_3 = null) use ($altStar1_2_1) {
  $__num = \func_num_args();
  $__res = $altStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_alternativeStar
$GLOBALS['Data_Profunctor_Star_alternativeStar'] = function($dictAlternative_0 = null) {
  $__num = \func_num_args();
  $applicativeStar1_1_0 = ($GLOBALS['Data_Profunctor_Star_applicativeStar'])((($dictAlternative_0)['Applicative0'])($GLOBALS['Prim_undefined']));
  $plusStar1_2_1 = ($GLOBALS['Data_Profunctor_Star_plusStar'])((($dictAlternative_0)['Plus1'])($GLOBALS['Prim_undefined']));
  $__res = ["Applicative0" => function($dollar__unused_3 = null) use ($applicativeStar1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($dollar__unused_3 = null) use ($plusStar1_2_1) {
  $__num = \func_num_args();
  $__res = $plusStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Profunctor_Star_monadPlusStar
$GLOBALS['Data_Profunctor_Star_monadPlusStar'] = function($dictMonadPlus_0 = null) {
  $__num = \func_num_args();
  $monadStar1_1_0 = ($GLOBALS['Data_Profunctor_Star_monadStar'])((($dictMonadPlus_0)['Monad0'])($GLOBALS['Prim_undefined']));
  $alternativeStar1_2_1 = ($GLOBALS['Data_Profunctor_Star_alternativeStar'])((($dictMonadPlus_0)['Alternative1'])($GLOBALS['Prim_undefined']));
  $__res = ["Monad0" => function($dollar__unused_3 = null) use ($monadStar1_1_0) {
  $__num = \func_num_args();
  $__res = $monadStar1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($dollar__unused_3 = null) use ($alternativeStar1_2_1) {
  $__num = \func_num_args();
  $__res = $alternativeStar1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

