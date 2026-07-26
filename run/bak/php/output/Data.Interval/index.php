<?php

namespace Data\Interval;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Category, Control.Extend, Control.Semigroupoid, Data.Bifoldable, Data.Bifunctor, Data.Bitraversable, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Interval, Data.Interval.Duration, Data.Maybe, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Traversable, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Category, Control.Extend, Control.Semigroupoid, Data.Bifoldable, Data.Bifunctor, Data.Bitraversable, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Interval, Data.Interval.Duration, Data.Maybe, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Traversable, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifoldable/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Bitraversable/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Interval/index.php';
require_once __DIR__ . '/../Data.Interval.Duration/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
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


// Data_Interval_show
$GLOBALS['Data_Interval_show'] = (($GLOBALS['Data_Maybe_showMaybe'])($GLOBALS['Data_Show_showInt']))['show'];

// Data_Interval_compare
$GLOBALS['Data_Interval_compare'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT"));
return (function() use ($__local_var_0_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($x_1) && (($x_1)->{'tag'} === "Nothing"))) {
$__t2 = null;;
if ((is_object($y_2) && (($y_2)->{'tag'} === "Nothing"))) {
$__t2 = new Phpurs_Data0("EQ");
goto end_branch_2;;
};
$__t2 = new Phpurs_Data0("LT");
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ((is_object($y_2) && (($y_2)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data0("GT");
goto end_branch_1;;
};
  if (((is_object($x_1) && (($x_1)->{'tag'} === "Just")) && (is_object($y_2) && (($y_2)->{'tag'} === "Just")))) {
$__t1 = (($__local_var_0_0)(($x_1)->{'value0'}))(($y_2)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Interval_StartEnd
$GLOBALS['Data_Interval_StartEnd'] = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("StartEnd", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_DurationEnd
$GLOBALS['Data_Interval_DurationEnd'] = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("DurationEnd", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_StartDuration
$GLOBALS['Data_Interval_StartDuration'] = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("StartDuration", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_DurationOnly
$GLOBALS['Data_Interval_DurationOnly'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("DurationOnly", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Interval_RecurringInterval
$GLOBALS['Data_Interval_RecurringInterval'] = (function() {
  $__fn = function($value0 = null, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("RecurringInterval", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_showInterval
$GLOBALS['Data_Interval_showInterval'] = (function() {
  $__fn = function($dictShow_0 = null, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["show" => function($v_2 = null) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "StartEnd"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(StartEnd "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow1_1)['show'])(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow1_1)['show'])(($v_2)->{'value1'})))(")"))));
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "DurationEnd"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(DurationEnd "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow_0)['show'])(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow1_1)['show'])(($v_2)->{'value1'})))(")"))));
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "StartDuration"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(StartDuration "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow1_1)['show'])(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow_0)['show'])(($v_2)->{'value1'})))(")"))));
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "DurationOnly"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(DurationOnly "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow_0)['show'])(($v_2)->{'value0'})))(")"));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
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

// Data_Interval_showRecurringInterval
$GLOBALS['Data_Interval_showRecurringInterval'] = (function() {
  $__fn = function($dictShow_0 = null, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["show" => function($v_2 = null) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(RecurringInterval "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(($GLOBALS['Data_Interval_show'])(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((((($GLOBALS['Data_Interval_showInterval'])($dictShow_0))($dictShow1_1))['show'])(($v_2)->{'value1'})))(")"))));
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

// Data_Interval_over
$GLOBALS['Data_Interval_over'] = (function() {
  $__fn = function($dictFunctor_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictFunctor_0)['map'])(($GLOBALS['Data_Interval_RecurringInterval'])(($v_2)->{'value0'})))(($f_1)(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_interval
$GLOBALS['Data_Interval_interval'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Interval_foldableInterval
$GLOBALS['Data_Interval_foldableInterval'] = ["foldl" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "StartEnd"))) {
$__t0 = (($v_0)((($v_0)($v1_1))(($v2_2)->{'value0'})))(($v2_2)->{'value1'});
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "DurationEnd"))) {
$__t0 = (($v_0)($v1_1))(($v2_2)->{'value1'});
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "StartDuration"))) {
$__t0 = (($v_0)($v1_1))(($v2_2)->{'value0'});
goto end_branch_0;;
};
  $__t0 = $v1_1;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => function($x_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldrDefault'])($GLOBALS['Data_Interval_foldableInterval']))($x_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_1 = ($dictMonoid_0)['mempty'];
  $__res = function($f_2 = null) use ($dictMonoid_0, $mempty_1_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_foldableInterval'])['foldl'])((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($acc_3 = null, $x_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)['Semigroup0'])(null))['append'])($acc_3))(($f_2)($x_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_foldableRecurringInterval
$GLOBALS['Data_Interval_foldableRecurringInterval'] = ["foldl" => (function() {
  $__fn = function($f_0 = null, $i_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_Interval_foldableInterval'])['foldl'])($f_0))($i_1)))($GLOBALS['Data_Interval_interval']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() {
  $__fn = function($f_0 = null, $i_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_Interval_foldableInterval'])['foldr'])($f_0))($i_1)))($GLOBALS['Data_Interval_interval']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)['mempty'];
  $__res = function($f_2 = null) use ($dictMonoid_0, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_foldableRecurringInterval'])['foldl'])((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($acc_3 = null, $x_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)['Semigroup0'])(null))['append'])($acc_3))(($f_2)($x_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_eqInterval
$GLOBALS['Data_Interval_eqInterval'] = (function() {
  $__fn = function($dictEq_0 = null, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($x_2) && (($x_2)->{'tag'} === "StartEnd"))) {
$__t0 = ((is_object($y_3) && (($y_3)->{'tag'} === "StartEnd")) && ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($dictEq1_1)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'})))(((($dictEq1_1)['eq'])(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  if ((is_object($x_2) && (($x_2)->{'tag'} === "DurationEnd"))) {
$__t0 = ((is_object($y_3) && (($y_3)->{'tag'} === "DurationEnd")) && ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($dictEq_0)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'})))(((($dictEq1_1)['eq'])(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  if ((is_object($x_2) && (($x_2)->{'tag'} === "StartDuration"))) {
$__t0 = ((is_object($y_3) && (($y_3)->{'tag'} === "StartDuration")) && ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])(((($dictEq1_1)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'})))(((($dictEq_0)['eq'])(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_2) && (($x_2)->{'tag'} === "DurationOnly")) && ((is_object($y_3) && (($y_3)->{'tag'} === "DurationOnly")) && ((($dictEq_0)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_eqRecurringInterval
$GLOBALS['Data_Interval_eqRecurringInterval'] = (function() {
  $__fn = function($dictEq_0 = null, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object(($x_2)->{'value0'}) && ((($x_2)->{'value0'})->{'tag'} === "Nothing"))) {
$__t0 = (is_object(($y_3)->{'value0'}) && ((($y_3)->{'value0'})->{'tag'} === "Nothing"));
goto end_branch_0;;
};
  $__t0 = ((is_object(($x_2)->{'value0'}) && ((($x_2)->{'value0'})->{'tag'} === "Just")) && ((is_object(($y_3)->{'value0'}) && ((($y_3)->{'value0'})->{'tag'} === "Just")) && ((($x_2)->{'value0'})->{'value0'} === (($y_3)->{'value0'})->{'value0'})));
  end_branch_0:;
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])['conj'])($__t0))(((((($GLOBALS['Data_Interval_eqInterval'])($dictEq_0))($dictEq1_1))['eq'])(($x_2)->{'value1'}))(($y_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_ordInterval
$GLOBALS['Data_Interval_ordInterval'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $eqInterval1_1_0 = ($GLOBALS['Data_Interval_eqInterval'])((($dictOrd_0)['Eq0'])(null));
  $__res = function($dictOrd1_2 = null) use ($dictOrd_0, $eqInterval1_1_0) {
  $__num = \func_num_args();
  $eqInterval2_3_1 = ($eqInterval1_1_0)((($dictOrd1_2)['Eq0'])(null));
  $__res = ["compare" => (function() use ($dictOrd1_2, $dictOrd_0) {
  $__fn = function($x_4 = null, $y_5 = null) use ($dictOrd1_2, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ((is_object($x_4) && (($x_4)->{'tag'} === "StartEnd"))) {
$__t3 = null;;
if ((is_object($y_5) && (($y_5)->{'tag'} === "StartEnd"))) {
$v_6_4 = ((($dictOrd1_2)['compare'])(($x_4)->{'value0'}))(($y_5)->{'value0'});
$__t5 = null;;
if ((is_object($v_6_4) && (($v_6_4)->{'tag'} === "LT"))) {
$__t5 = new Phpurs_Data0("LT");
goto end_branch_5;;
};
if ((is_object($v_6_4) && (($v_6_4)->{'tag'} === "GT"))) {
$__t5 = new Phpurs_Data0("GT");
goto end_branch_5;;
};
$__t5 = ((($dictOrd1_2)['compare'])(($x_4)->{'value1'}))(($y_5)->{'value1'});
end_branch_5:;
$__t3 = $__t5;
goto end_branch_3;;
};
$__t3 = new Phpurs_Data0("LT");
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
  if ((is_object($y_5) && (($y_5)->{'tag'} === "StartEnd"))) {
$__t2 = new Phpurs_Data0("GT");
goto end_branch_2;;
};
  if ((is_object($x_4) && (($x_4)->{'tag'} === "DurationEnd"))) {
$__t6 = null;;
if ((is_object($y_5) && (($y_5)->{'tag'} === "DurationEnd"))) {
$v_6_7 = ((($dictOrd_0)['compare'])(($x_4)->{'value0'}))(($y_5)->{'value0'});
$__t8 = null;;
if ((is_object($v_6_7) && (($v_6_7)->{'tag'} === "LT"))) {
$__t8 = new Phpurs_Data0("LT");
goto end_branch_8;;
};
if ((is_object($v_6_7) && (($v_6_7)->{'tag'} === "GT"))) {
$__t8 = new Phpurs_Data0("GT");
goto end_branch_8;;
};
$__t8 = ((($dictOrd1_2)['compare'])(($x_4)->{'value1'}))(($y_5)->{'value1'});
end_branch_8:;
$__t6 = $__t8;
goto end_branch_6;;
};
$__t6 = new Phpurs_Data0("LT");
end_branch_6:;
$__t2 = $__t6;
goto end_branch_2;;
};
  if ((is_object($y_5) && (($y_5)->{'tag'} === "DurationEnd"))) {
$__t2 = new Phpurs_Data0("GT");
goto end_branch_2;;
};
  if ((is_object($x_4) && (($x_4)->{'tag'} === "StartDuration"))) {
$__t9 = null;;
if ((is_object($y_5) && (($y_5)->{'tag'} === "StartDuration"))) {
$v_6_10 = ((($dictOrd1_2)['compare'])(($x_4)->{'value0'}))(($y_5)->{'value0'});
$__t11 = null;;
if ((is_object($v_6_10) && (($v_6_10)->{'tag'} === "LT"))) {
$__t11 = new Phpurs_Data0("LT");
goto end_branch_11;;
};
if ((is_object($v_6_10) && (($v_6_10)->{'tag'} === "GT"))) {
$__t11 = new Phpurs_Data0("GT");
goto end_branch_11;;
};
$__t11 = ((($dictOrd_0)['compare'])(($x_4)->{'value1'}))(($y_5)->{'value1'});
end_branch_11:;
$__t9 = $__t11;
goto end_branch_9;;
};
$__t9 = new Phpurs_Data0("LT");
end_branch_9:;
$__t2 = $__t9;
goto end_branch_2;;
};
  if ((is_object($y_5) && (($y_5)->{'tag'} === "StartDuration"))) {
$__t2 = new Phpurs_Data0("GT");
goto end_branch_2;;
};
  if (((is_object($x_4) && (($x_4)->{'tag'} === "DurationOnly")) && (is_object($y_5) && (($y_5)->{'tag'} === "DurationOnly")))) {
$__t2 = ((($dictOrd_0)['compare'])(($x_4)->{'value0'}))(($y_5)->{'value0'});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_4 = null) use ($eqInterval2_3_1) {
  $__num = \func_num_args();
  $__res = $eqInterval2_3_1;
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

// Data_Interval_ordRecurringInterval
$GLOBALS['Data_Interval_ordRecurringInterval'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $ordInterval1_1_0 = ($GLOBALS['Data_Interval_ordInterval'])($dictOrd_0);
  $eqRecurringInterval1_2_1 = ($GLOBALS['Data_Interval_eqRecurringInterval'])((($dictOrd_0)['Eq0'])(null));
  $__res = function($dictOrd1_3 = null) use ($eqRecurringInterval1_2_1, $ordInterval1_1_0) {
  $__num = \func_num_args();
  $eqRecurringInterval2_4_2 = ($eqRecurringInterval1_2_1)((($dictOrd1_3)['Eq0'])(null));
  $__res = ["compare" => (function() use ($dictOrd1_3, $ordInterval1_1_0) {
  $__fn = function($x_5 = null, $y_6 = null) use ($dictOrd1_3, $ordInterval1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_7_3 = (($GLOBALS['Data_Interval_compare'])(($x_5)->{'value0'}))(($y_6)->{'value0'});
  $__t4 = null;;
  if ((is_object($v_7_3) && (($v_7_3)->{'tag'} === "LT"))) {
$__t4 = new Phpurs_Data0("LT");
goto end_branch_4;;
};
  if ((is_object($v_7_3) && (($v_7_3)->{'tag'} === "GT"))) {
$__t4 = new Phpurs_Data0("GT");
goto end_branch_4;;
};
  $__t4 = (((($ordInterval1_1_0)($dictOrd1_3))['compare'])(($x_5)->{'value1'}))(($y_6)->{'value1'});
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_5 = null) use ($eqRecurringInterval2_4_2) {
  $__num = \func_num_args();
  $__res = $eqRecurringInterval2_4_2;
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

// Data_Interval_bifunctorInterval
$GLOBALS['Data_Interval_bifunctorInterval'] = ["bimap" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "StartEnd"))) {
$__t0 = new Phpurs_Data2("StartEnd", ($v1_1)(($v2_2)->{'value0'}), ($v1_1)(($v2_2)->{'value1'}));
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "DurationEnd"))) {
$__t0 = new Phpurs_Data2("DurationEnd", ($v_0)(($v2_2)->{'value0'}), ($v1_1)(($v2_2)->{'value1'}));
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "StartDuration"))) {
$__t0 = new Phpurs_Data2("StartDuration", ($v1_1)(($v2_2)->{'value0'}), ($v_0)(($v2_2)->{'value1'}));
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "DurationOnly"))) {
$__t0 = new Phpurs_Data1("DurationOnly", ($v_0)(($v2_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_Interval_bifunctorRecurringInterval
$GLOBALS['Data_Interval_bifunctorRecurringInterval'] = ["bimap" => (function() {
  $__fn = function($f_0 = null, $g_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data2("RecurringInterval", ($v_2)->{'value0'}, (((($GLOBALS['Data_Interval_bifunctorInterval'])['bimap'])($f_0))($g_1))(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_Interval_functorInterval
$GLOBALS['Data_Interval_functorInterval'] = ["map" => (($GLOBALS['Data_Interval_bifunctorInterval'])['bimap'])(function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];

// Data_Interval_extendInterval
$GLOBALS['Data_Interval_extendInterval'] = ["extend" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "StartEnd"))) {
$__t0 = new Phpurs_Data2("StartEnd", ($v_0)($v1_1), ($v_0)($v1_1));
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "DurationEnd"))) {
$__t0 = new Phpurs_Data2("DurationEnd", ($v1_1)->{'value0'}, ($v_0)($v1_1));
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "StartDuration"))) {
$__t0 = new Phpurs_Data2("StartDuration", ($v_0)($v1_1), ($v1_1)->{'value1'});
goto end_branch_0;;
};
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "DurationOnly"))) {
$__t0 = new Phpurs_Data1("DurationOnly", ($v1_1)->{'value0'});
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
})(), "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_functorRecurringInterval
$GLOBALS['Data_Interval_functorRecurringInterval'] = ["map" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("RecurringInterval", ($v_1)->{'value0'}, ((($GLOBALS['Data_Interval_functorInterval'])['map'])($f_0))(($v_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Interval_extendRecurringInterval
$GLOBALS['Data_Interval_extendRecurringInterval'] = ["extend" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($f_0)($v_1);
  $__res = new Phpurs_Data2("RecurringInterval", ($v_1)->{'value0'}, ((($GLOBALS['Data_Interval_extendInterval'])['extend'])(function($v_3 = null) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = $__local_var_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_traversableInterval
$GLOBALS['Data_Interval_traversableInterval'] = ["traverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)['Apply0'])(null);
  $Functor0_2_1 = (($Apply0_1_0)['Functor0'])(null);
  $__res = (function() use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0) {
  $__fn = function($v_3 = null, $v1_4 = null) use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "StartEnd"))) {
$__t2 = ((($Apply0_1_0)['apply'])(((($Functor0_2_1)['map'])($GLOBALS['Data_Interval_StartEnd']))(($v_3)(($v1_4)->{'value0'}))))(($v_3)(($v1_4)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "DurationEnd"))) {
$__t2 = ((($Functor0_2_1)['map'])(($GLOBALS['Data_Interval_DurationEnd'])(($v1_4)->{'value0'})))(($v_3)(($v1_4)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "StartDuration"))) {
$__local_var_5_3 = ($v1_4)->{'value1'};
$__t2 = ((($Functor0_2_1)['map'])(function($v2_6 = null) use ($__local_var_5_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("StartDuration", $v2_6, $__local_var_5_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_3)(($v1_4)->{'value0'}));
goto end_branch_2;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "DurationOnly"))) {
$__t2 = (($dictApplicative_0)['pure'])(new Phpurs_Data1("DurationOnly", ($v1_4)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_traversableInterval'])['traverse'])($dictApplicative_0))($GLOBALS['Data_Traversable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_foldableInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_traversableRecurringInterval
$GLOBALS['Data_Interval_traversableRecurringInterval'] = ["traverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $over1_1_0 = ($GLOBALS['Data_Interval_over'])((((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null));
  $traverse1_2_1 = (($GLOBALS['Data_Interval_traversableInterval'])['traverse'])($dictApplicative_0);
  $__res = (function() use ($over1_1_0, $traverse1_2_1) {
  $__fn = function($f_3 = null, $i_4 = null) use ($over1_1_0, $traverse1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($over1_1_0)(($traverse1_2_1)($f_3)))($i_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_traversableRecurringInterval'])['traverse'])($dictApplicative_0))($GLOBALS['Data_Traversable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_foldableRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_bifoldableInterval
$GLOBALS['Data_Interval_bifoldableInterval'] = ["bifoldl" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null, $v3_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ((is_object($v3_3) && (($v3_3)->{'tag'} === "StartEnd"))) {
$__t0 = (($v1_1)((($v1_1)($v2_2))(($v3_3)->{'value0'})))(($v3_3)->{'value1'});
goto end_branch_0;;
};
  if ((is_object($v3_3) && (($v3_3)->{'tag'} === "DurationEnd"))) {
$__t0 = (($v1_1)((($v_0)($v2_2))(($v3_3)->{'value0'})))(($v3_3)->{'value1'});
goto end_branch_0;;
};
  if ((is_object($v3_3) && (($v3_3)->{'tag'} === "StartDuration"))) {
$__t0 = (($v1_1)((($v_0)($v2_2))(($v3_3)->{'value1'})))(($v3_3)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($v3_3) && (($v3_3)->{'tag'} === "DurationOnly"))) {
$__t0 = (($v_0)($v2_2))(($v3_3)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldr" => function($x_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Bifoldable_bifoldrDefault'])($GLOBALS['Data_Interval_bifoldableInterval']))($x_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bifoldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_1 = (($dictMonoid_0)['Semigroup0'])(null);
  $mempty_2_2 = ($dictMonoid_0)['mempty'];
  $__res = (function() use ($__local_var_1_1, $mempty_2_2) {
  $__fn = function($f_3 = null, $g_4 = null) use ($__local_var_1_1, $mempty_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Interval_bifoldableInterval'])['bifoldl'])((function() use ($__local_var_1_1, $f_3) {
  $__fn = function($m_5 = null, $a_6 = null) use ($__local_var_1_1, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_1)['append'])($m_5))(($f_3)($a_6));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((function() use ($__local_var_1_1, $g_4) {
  $__fn = function($m_5 = null, $b_6 = null) use ($__local_var_1_1, $g_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_1)['append'])($m_5))(($g_4)($b_6));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_2_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_bifoldableRecurringInterval
$GLOBALS['Data_Interval_bifoldableRecurringInterval'] = ["bifoldl" => (function() {
  $__fn = function($f_0 = null, $g_1 = null, $i_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((((($GLOBALS['Data_Interval_bifoldableInterval'])['bifoldl'])($f_0))($g_1))($i_2)))($GLOBALS['Data_Interval_interval']);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "bifoldr" => (function() {
  $__fn = function($f_0 = null, $g_1 = null, $i_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((((($GLOBALS['Data_Interval_bifoldableInterval'])['bifoldr'])($f_0))($g_1))($i_2)))($GLOBALS['Data_Interval_interval']);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictMonoid_0)['Semigroup0'])(null);
  $mempty_2_1 = ($dictMonoid_0)['mempty'];
  $__res = (function() use ($__local_var_1_0, $mempty_2_1) {
  $__fn = function($f_3 = null, $g_4 = null) use ($__local_var_1_0, $mempty_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Interval_bifoldableRecurringInterval'])['bifoldl'])((function() use ($__local_var_1_0, $f_3) {
  $__fn = function($m_5 = null, $a_6 = null) use ($__local_var_1_0, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['append'])($m_5))(($f_3)($a_6));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((function() use ($__local_var_1_0, $g_4) {
  $__fn = function($m_5 = null, $b_6 = null) use ($__local_var_1_0, $g_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['append'])($m_5))(($g_4)($b_6));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_2_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_bitraversableInterval
$GLOBALS['Data_Interval_bitraversableInterval'] = ["bitraverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)['Apply0'])(null);
  $__local_var_2_1 = (($Apply0_1_0)['Functor0'])(null);
  $__res = (function() use ($Apply0_1_0, $__local_var_2_1) {
  $__fn = function($v_3 = null, $v1_4 = null, $v2_5 = null) use ($Apply0_1_0, $__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t2 = null;;
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "StartEnd"))) {
$__t2 = ((($Apply0_1_0)['apply'])(((($__local_var_2_1)['map'])($GLOBALS['Data_Interval_StartEnd']))(($v1_4)(($v2_5)->{'value0'}))))(($v1_4)(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "DurationEnd"))) {
$__t2 = ((($Apply0_1_0)['apply'])(((($__local_var_2_1)['map'])($GLOBALS['Data_Interval_DurationEnd']))(($v_3)(($v2_5)->{'value0'}))))(($v1_4)(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "StartDuration"))) {
$__t2 = ((($Apply0_1_0)['apply'])(((($__local_var_2_1)['map'])($GLOBALS['Data_Interval_StartDuration']))(($v1_4)(($v2_5)->{'value0'}))))(($v_3)(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ((is_object($v2_5) && (($v2_5)->{'tag'} === "DurationOnly"))) {
$__t2 = ((($__local_var_2_1)['map'])($GLOBALS['Data_Interval_DurationOnly']))(($v_3)(($v2_5)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Interval_bitraversableInterval'])['bitraverse'])($dictApplicative_0))($GLOBALS['Data_Bitraversable_identity']))($GLOBALS['Data_Bitraversable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifunctorInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifoldableInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_bitraversableRecurringInterval
$GLOBALS['Data_Interval_bitraversableRecurringInterval'] = ["bitraverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $over1_1_0 = ($GLOBALS['Data_Interval_over'])((((($dictApplicative_0)['Apply0'])(null))['Functor0'])(null));
  $bitraverse1_2_1 = (($GLOBALS['Data_Interval_bitraversableInterval'])['bitraverse'])($dictApplicative_0);
  $__res = (function() use ($bitraverse1_2_1, $over1_1_0) {
  $__fn = function($l_3 = null, $r_4 = null, $i_5 = null) use ($bitraverse1_2_1, $over1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($over1_1_0)((($bitraverse1_2_1)($l_3))($r_4)))($i_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Interval_bitraversableRecurringInterval'])['bitraverse'])($dictApplicative_0))($GLOBALS['Data_Bitraversable_identity']))($GLOBALS['Data_Bitraversable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifunctorRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifoldableRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

