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
\PhpursThunks::$thunks['Data_Interval_show'] = function() { $v = ((($GLOBALS['Data_Maybe_showMaybe'] ?? \PhpursThunks::eval('Data_Maybe_showMaybe')))(($GLOBALS['Data_Show_showInt'] ?? \PhpursThunks::eval('Data_Show_showInt'))))->show; return $v; };
\PhpursThunks::$thunks['Data_Interval_eq'] = function() { $v = (function() {
  $__fn = function($x_0, $y_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_eq"), recVars=[];
  if ((is_object($x_0) && (($x_0)->tag === "Nothing"))) {
$__t0 = (is_object($y_1) && (($y_1)->tag === "Nothing"));
} else {
$__t0 = ((is_object($x_0) && (($x_0)->tag === "Just")) && ((is_object($y_1) && (($y_1)->tag === "Just")) && ((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))(($x_0)->value0))(($y_1)->value0)));
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_compare'] = function() { $v = ((($GLOBALS['Data_Maybe_ordMaybe'] ?? \PhpursThunks::eval('Data_Maybe_ordMaybe')))(($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'))))->compare; return $v; };
\PhpursThunks::$thunks['Data_Interval_StartEnd'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_DurationEnd'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_StartDuration'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_DurationOnly'] = function() { $v = function($value0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("DurationOnly", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_RecurringInterval'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_showInterval'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_showInterval"), recVars=[];
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "StartEnd"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(StartEnd "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow1_1)->show)(($v_2)->value0)))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow1_1)->show)(($v_2)->value1)))(")"))));
} else {
if ((is_object($v_2) && (($v_2)->tag === "DurationEnd"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(DurationEnd "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->show)(($v_2)->value0)))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow1_1)->show)(($v_2)->value1)))(")"))));
} else {
if ((is_object($v_2) && (($v_2)->tag === "StartDuration"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(StartDuration "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow1_1)->show)(($v_2)->value0)))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->show)(($v_2)->value1)))(")"))));
} else {
if ((is_object($v_2) && (($v_2)->tag === "DurationOnly"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(DurationOnly "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->show)(($v_2)->value0)))(")"));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_showRecurringInterval'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_showRecurringInterval"), recVars=[];
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(RecurringInterval "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Interval_show'] ?? \PhpursThunks::eval('Data_Interval_show')))(($v_2)->value0)))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((((($GLOBALS['Data_Interval_showInterval'] ?? \PhpursThunks::eval('Data_Interval_showInterval')))($dictShow_0))($dictShow1_1))->show)(($v_2)->value1)))(")"))));
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
\PhpursThunks::$thunks['Data_Interval_over'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_over"), recVars=[];
  $__res = ((($dictFunctor_0)->map)((($GLOBALS['Data_Interval_RecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_RecurringInterval')))(($v_2)->value0)))(($f_1)(($v_2)->value1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_interval'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_interval"), recVars=[];
  $__res = ($v_0)->value1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_foldableInterval'] = function() { $v = (object)["foldl" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_foldableInterval"];
  if ((is_object($v2_2) && (($v2_2)->tag === "StartEnd"))) {
$__t0 = (($v_0)((($v_0)($v1_1))(($v2_2)->value0)))(($v2_2)->value1);
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "DurationEnd"))) {
$__t0 = (($v_0)($v1_1))(($v2_2)->value1);
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "StartDuration"))) {
$__t0 = (($v_0)($v1_1))(($v2_2)->value0);
} else {
$__t0 = $v1_1;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_foldableInterval"];
  $__res = ((($GLOBALS['Data_Foldable_foldrDefault'] ?? \PhpursThunks::eval('Data_Foldable_foldrDefault')))(($GLOBALS['Data_Interval_foldableInterval'] ?? \PhpursThunks::eval('Data_Interval_foldableInterval'))))($x_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_foldableInterval"];
  $__res = ((($GLOBALS['Data_Foldable_foldMapDefaultL'] ?? \PhpursThunks::eval('Data_Foldable_foldMapDefaultL')))(($GLOBALS['Data_Interval_foldableInterval'] ?? \PhpursThunks::eval('Data_Interval_foldableInterval'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_foldableRecurringInterval'] = function() { $v = (object)["foldl" => (function() {
  $__fn = function($f_0, $i_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_foldableRecurringInterval"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v2_2) use ($f_0, $i_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_foldableRecurringInterval"];
  if ((is_object($v2_2) && (($v2_2)->tag === "StartEnd"))) {
$__t0 = (($f_0)((($f_0)($i_1))(($v2_2)->value0)))(($v2_2)->value1);
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "DurationEnd"))) {
$__t0 = (($f_0)($i_1))(($v2_2)->value1);
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "StartDuration"))) {
$__t0 = (($f_0)($i_1))(($v2_2)->value0);
} else {
$__t0 = $i_1;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Interval_interval'] ?? \PhpursThunks::eval('Data_Interval_interval')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() {
  $__fn = function($f_0, $i_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_foldableRecurringInterval"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((((($GLOBALS['Data_Foldable_foldrDefault'] ?? \PhpursThunks::eval('Data_Foldable_foldrDefault')))(($GLOBALS['Data_Interval_foldableInterval'] ?? \PhpursThunks::eval('Data_Interval_foldableInterval'))))($f_0))($i_1)))(($GLOBALS['Data_Interval_interval'] ?? \PhpursThunks::eval('Data_Interval_interval')));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_foldableRecurringInterval"];
  $__res = ((($GLOBALS['Data_Foldable_foldMapDefaultL'] ?? \PhpursThunks::eval('Data_Foldable_foldMapDefaultL')))(($GLOBALS['Data_Interval_foldableRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_foldableRecurringInterval'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_eqInterval'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_eqInterval"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_2) && (($x_2)->tag === "StartEnd"))) {
$__t0 = ((is_object($y_3) && (($y_3)->tag === "StartEnd")) && ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($dictEq1_1)->eq)(($x_2)->value0))(($y_3)->value0)))(((($dictEq1_1)->eq)(($x_2)->value1))(($y_3)->value1)));
} else {
if ((is_object($x_2) && (($x_2)->tag === "DurationEnd"))) {
$__t0 = ((is_object($y_3) && (($y_3)->tag === "DurationEnd")) && ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($dictEq_0)->eq)(($x_2)->value0))(($y_3)->value0)))(((($dictEq1_1)->eq)(($x_2)->value1))(($y_3)->value1)));
} else {
if ((is_object($x_2) && (($x_2)->tag === "StartDuration"))) {
$__t0 = ((is_object($y_3) && (($y_3)->tag === "StartDuration")) && ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($dictEq1_1)->eq)(($x_2)->value0))(($y_3)->value0)))(((($dictEq_0)->eq)(($x_2)->value1))(($y_3)->value1)));
} else {
$__t0 = ((is_object($x_2) && (($x_2)->tag === "DurationOnly")) && ((is_object($y_3) && (($y_3)->tag === "DurationOnly")) && ((($dictEq_0)->eq)(($x_2)->value0))(($y_3)->value0)));
};
};
};
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_eqRecurringInterval'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_eqRecurringInterval"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_HeytingAlgebra_boolConj'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolConj')))(((($GLOBALS['Data_Interval_eq'] ?? \PhpursThunks::eval('Data_Interval_eq')))(($x_2)->value0))(($y_3)->value0)))((((((($GLOBALS['Data_Interval_eqInterval'] ?? \PhpursThunks::eval('Data_Interval_eqInterval')))($dictEq_0))($dictEq1_1))->eq)(($x_2)->value1))(($y_3)->value1));
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Interval_ordInterval'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_ordInterval"), recVars=[];
  $eqInterval1_1_0 = (($GLOBALS['Data_Interval_eqInterval'] ?? \PhpursThunks::eval('Data_Interval_eqInterval')))((($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictOrd1_2) use ($dictOrd_0, $eqInterval1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eqInterval2_3_1 = ($eqInterval1_1_0)((($dictOrd1_2)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["compare" => (function() use ($dictOrd1_2, $dictOrd_0) {
  $__fn = function($x_4, $y_5 = null) use ($dictOrd1_2, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($x_4) && (($x_4)->tag === "StartEnd"))) {
if ((is_object($y_5) && (($y_5)->tag === "StartEnd"))) {
$v_6_4 = ((($dictOrd1_2)->compare)(($x_4)->value0))(($y_5)->value0);
if ((is_object($v_6_4) && (($v_6_4)->tag === "LT"))) {
$__t5 = new Phpurs_Data0("LT");
} else {
if ((is_object($v_6_4) && (($v_6_4)->tag === "GT"))) {
$__t5 = new Phpurs_Data0("GT");
} else {
$__t5 = ((($dictOrd1_2)->compare)(($x_4)->value1))(($y_5)->value1);
};
};
$__t3 = $__t5;
} else {
$__t3 = new Phpurs_Data0("LT");
};
$__t2 = $__t3;
} else {
if ((is_object($y_5) && (($y_5)->tag === "StartEnd"))) {
$__t2 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_4) && (($x_4)->tag === "DurationEnd"))) {
if ((is_object($y_5) && (($y_5)->tag === "DurationEnd"))) {
$v_6_7 = ((($dictOrd_0)->compare)(($x_4)->value0))(($y_5)->value0);
if ((is_object($v_6_7) && (($v_6_7)->tag === "LT"))) {
$__t8 = new Phpurs_Data0("LT");
} else {
if ((is_object($v_6_7) && (($v_6_7)->tag === "GT"))) {
$__t8 = new Phpurs_Data0("GT");
} else {
$__t8 = ((($dictOrd1_2)->compare)(($x_4)->value1))(($y_5)->value1);
};
};
$__t6 = $__t8;
} else {
$__t6 = new Phpurs_Data0("LT");
};
$__t2 = $__t6;
} else {
if ((is_object($y_5) && (($y_5)->tag === "DurationEnd"))) {
$__t2 = new Phpurs_Data0("GT");
} else {
if ((is_object($x_4) && (($x_4)->tag === "StartDuration"))) {
if ((is_object($y_5) && (($y_5)->tag === "StartDuration"))) {
$v_6_10 = ((($dictOrd1_2)->compare)(($x_4)->value0))(($y_5)->value0);
if ((is_object($v_6_10) && (($v_6_10)->tag === "LT"))) {
$__t11 = new Phpurs_Data0("LT");
} else {
if ((is_object($v_6_10) && (($v_6_10)->tag === "GT"))) {
$__t11 = new Phpurs_Data0("GT");
} else {
$__t11 = ((($dictOrd_0)->compare)(($x_4)->value1))(($y_5)->value1);
};
};
$__t9 = $__t11;
} else {
$__t9 = new Phpurs_Data0("LT");
};
$__t2 = $__t9;
} else {
if ((is_object($y_5) && (($y_5)->tag === "StartDuration"))) {
$__t2 = new Phpurs_Data0("GT");
} else {
if (((is_object($x_4) && (($x_4)->tag === "DurationOnly")) && (is_object($y_5) && (($y_5)->tag === "DurationOnly")))) {
$__t2 = ((($dictOrd_0)->compare)(($x_4)->value0))(($y_5)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
};
};
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_4) use ($eqInterval2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
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
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_ordRecurringInterval'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Interval_ordRecurringInterval"), recVars=[];
  $ordInterval1_1_0 = (($GLOBALS['Data_Interval_ordInterval'] ?? \PhpursThunks::eval('Data_Interval_ordInterval')))($dictOrd_0);
  $eqRecurringInterval1_2_1 = (($GLOBALS['Data_Interval_eqRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_eqRecurringInterval')))((($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictOrd1_3) use ($eqRecurringInterval1_2_1, $ordInterval1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eqRecurringInterval2_4_2 = ($eqRecurringInterval1_2_1)((($dictOrd1_3)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["compare" => (function() use ($dictOrd1_3, $ordInterval1_1_0) {
  $__fn = function($x_5, $y_6 = null) use ($dictOrd1_3, $ordInterval1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_7_3 = ((($GLOBALS['Data_Interval_compare'] ?? \PhpursThunks::eval('Data_Interval_compare')))(($x_5)->value0))(($y_6)->value0);
  if ((is_object($v_7_3) && (($v_7_3)->tag === "LT"))) {
$__t4 = new Phpurs_Data0("LT");
} else {
if ((is_object($v_7_3) && (($v_7_3)->tag === "GT"))) {
$__t4 = new Phpurs_Data0("GT");
} else {
$__t4 = (((($ordInterval1_1_0)($dictOrd1_3))->compare)(($x_5)->value1))(($y_6)->value1);
};
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_5) use ($eqRecurringInterval2_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
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
}; return $v; };
\PhpursThunks::$thunks['Data_Interval_bifunctorInterval'] = function() { $v = (object)["bimap" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_2) && (($v2_2)->tag === "StartEnd"))) {
$__t0 = new Phpurs_Data2("StartEnd", ($v1_1)(($v2_2)->value0), ($v1_1)(($v2_2)->value1));
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "DurationEnd"))) {
$__t0 = new Phpurs_Data2("DurationEnd", ($v_0)(($v2_2)->value0), ($v1_1)(($v2_2)->value1));
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "StartDuration"))) {
$__t0 = new Phpurs_Data2("StartDuration", ($v1_1)(($v2_2)->value0), ($v_0)(($v2_2)->value1));
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "DurationOnly"))) {
$__t0 = new Phpurs_Data1("DurationOnly", ($v_0)(($v2_2)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_bifunctorRecurringInterval'] = function() { $v = (object)["bimap" => (function() {
  $__fn = function($f_0, $g_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(($v_2)->value1) && ((($v_2)->value1)->tag === "StartEnd"))) {
$__t0 = new Phpurs_Data2("StartEnd", ($g_1)((($v_2)->value1)->value0), ($g_1)((($v_2)->value1)->value1));
} else {
if ((is_object(($v_2)->value1) && ((($v_2)->value1)->tag === "DurationEnd"))) {
$__t0 = new Phpurs_Data2("DurationEnd", ($f_0)((($v_2)->value1)->value0), ($g_1)((($v_2)->value1)->value1));
} else {
if ((is_object(($v_2)->value1) && ((($v_2)->value1)->tag === "StartDuration"))) {
$__t0 = new Phpurs_Data2("StartDuration", ($g_1)((($v_2)->value1)->value0), ($f_0)((($v_2)->value1)->value1));
} else {
if ((is_object(($v_2)->value1) && ((($v_2)->value1)->tag === "DurationOnly"))) {
$__t0 = new Phpurs_Data1("DurationOnly", ($f_0)((($v_2)->value1)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = new Phpurs_Data2("RecurringInterval", ($v_2)->value0, $__t0);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_functorInterval'] = function() { $v = (object)["map" => (function() {
  $__fn = function($v1_0, $v2_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_1) && (($v2_1)->tag === "StartEnd"))) {
$__t0 = new Phpurs_Data2("StartEnd", ($v1_0)(($v2_1)->value0), ($v1_0)(($v2_1)->value1));
} else {
if ((is_object($v2_1) && (($v2_1)->tag === "DurationEnd"))) {
$__t0 = new Phpurs_Data2("DurationEnd", ($v2_1)->value0, ($v1_0)(($v2_1)->value1));
} else {
if ((is_object($v2_1) && (($v2_1)->tag === "StartDuration"))) {
$__t0 = new Phpurs_Data2("StartDuration", ($v1_0)(($v2_1)->value0), ($v2_1)->value1);
} else {
if ((is_object($v2_1) && (($v2_1)->tag === "DurationOnly"))) {
$__t0 = new Phpurs_Data1("DurationOnly", ($v2_1)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_extendInterval'] = function() { $v = (object)["extend" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v1_1) && (($v1_1)->tag === "StartEnd"))) {
$__t0 = new Phpurs_Data2("StartEnd", ($v_0)($v1_1), ($v_0)($v1_1));
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "DurationEnd"))) {
$__t0 = new Phpurs_Data2("DurationEnd", ($v1_1)->value0, ($v_0)($v1_1));
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "StartDuration"))) {
$__t0 = new Phpurs_Data2("StartDuration", ($v_0)($v1_1), ($v1_1)->value1);
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "DurationOnly"))) {
$__t0 = new Phpurs_Data1("DurationOnly", ($v1_1)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Interval_functorInterval'] ?? \PhpursThunks::eval('Data_Interval_functorInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_functorRecurringInterval'] = function() { $v = (object)["map" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "StartEnd"))) {
$__t0 = new Phpurs_Data2("StartEnd", ($f_0)((($v_1)->value1)->value0), ($f_0)((($v_1)->value1)->value1));
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "DurationEnd"))) {
$__t0 = new Phpurs_Data2("DurationEnd", (($v_1)->value1)->value0, ($f_0)((($v_1)->value1)->value1));
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "StartDuration"))) {
$__t0 = new Phpurs_Data2("StartDuration", ($f_0)((($v_1)->value1)->value0), (($v_1)->value1)->value1);
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "DurationOnly"))) {
$__t0 = new Phpurs_Data1("DurationOnly", (($v_1)->value1)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = new Phpurs_Data2("RecurringInterval", ($v_1)->value0, $__t0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Interval_extendRecurringInterval'] = function() { $v = (object)["extend" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = ($f_0)($v_1);
  if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "StartEnd"))) {
$__t1 = new Phpurs_Data2("StartEnd", $__local_var_2_0, $__local_var_2_0);
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "DurationEnd"))) {
$__t1 = new Phpurs_Data2("DurationEnd", (($v_1)->value1)->value0, $__local_var_2_0);
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "StartDuration"))) {
$__t1 = new Phpurs_Data2("StartDuration", $__local_var_2_0, (($v_1)->value1)->value1);
} else {
if ((is_object(($v_1)->value1) && ((($v_1)->value1)->tag === "DurationOnly"))) {
$__t1 = new Phpurs_Data1("DurationOnly", (($v_1)->value1)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
};
  $__res = new Phpurs_Data2("RecurringInterval", ($v_1)->value0, $__t1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Interval_functorRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_functorRecurringInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_traversableInterval'] = function() { $v = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableInterval"];
  $Apply0_1_0 = (($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $Functor0_2_1 = (($Apply0_1_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (function() use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0) {
  $__fn = function($v_3, $v1_4 = null) use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableInterval"];
  if ((is_object($v1_4) && (($v1_4)->tag === "StartEnd"))) {
$__t2 = ((($Apply0_1_0)->apply)(((($Functor0_2_1)->map)(($GLOBALS['Data_Interval_StartEnd'] ?? \PhpursThunks::eval('Data_Interval_StartEnd'))))(($v_3)(($v1_4)->value0))))(($v_3)(($v1_4)->value1));
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "DurationEnd"))) {
$__t2 = ((($Functor0_2_1)->map)((($GLOBALS['Data_Interval_DurationEnd'] ?? \PhpursThunks::eval('Data_Interval_DurationEnd')))(($v1_4)->value0)))(($v_3)(($v1_4)->value1));
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "StartDuration"))) {
$__local_var_5_3 = ($v1_4)->value1;
$__t2 = ((($Functor0_2_1)->map)(function($v2_6) use ($__local_var_5_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableInterval"];
  $__res = new Phpurs_Data2("StartDuration", $v2_6, $__local_var_5_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_3)(($v1_4)->value0));
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "DurationOnly"))) {
$__t2 = (($dictApplicative_0)->pure)(new Phpurs_Data1("DurationOnly", ($v1_4)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
};
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableInterval"];
  $__res = (((($GLOBALS['Data_Interval_traversableInterval'] ?? \PhpursThunks::eval('Data_Interval_traversableInterval')))->traverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableInterval"];
  $__res = ($GLOBALS['Data_Interval_functorInterval'] ?? \PhpursThunks::eval('Data_Interval_functorInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableInterval"];
  $__res = ($GLOBALS['Data_Interval_foldableInterval'] ?? \PhpursThunks::eval('Data_Interval_foldableInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_traversableRecurringInterval'] = function() { $v = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableRecurringInterval"];
  $over1_1_0 = (($GLOBALS['Data_Interval_over'] ?? \PhpursThunks::eval('Data_Interval_over')))((((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traverse1_2_1 = ((($GLOBALS['Data_Interval_traversableInterval'] ?? \PhpursThunks::eval('Data_Interval_traversableInterval')))->traverse)($dictApplicative_0);
  $__res = (function() use ($over1_1_0, $traverse1_2_1) {
  $__fn = function($f_3, $i_4 = null) use ($over1_1_0, $traverse1_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableRecurringInterval"];
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableRecurringInterval"];
  $__res = (((($GLOBALS['Data_Interval_traversableRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_traversableRecurringInterval')))->traverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableRecurringInterval"];
  $__res = ($GLOBALS['Data_Interval_functorRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_functorRecurringInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_traversableRecurringInterval"];
  $__res = ($GLOBALS['Data_Interval_foldableRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_foldableRecurringInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_bifoldableInterval'] = function() { $v = (object)["bifoldl" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null, $v3_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bifoldableInterval"];
  if ((is_object($v3_3) && (($v3_3)->tag === "StartEnd"))) {
$__t0 = (($v1_1)((($v1_1)($v2_2))(($v3_3)->value0)))(($v3_3)->value1);
} else {
if ((is_object($v3_3) && (($v3_3)->tag === "DurationEnd"))) {
$__t0 = (($v1_1)((($v_0)($v2_2))(($v3_3)->value0)))(($v3_3)->value1);
} else {
if ((is_object($v3_3) && (($v3_3)->tag === "StartDuration"))) {
$__t0 = (($v1_1)((($v_0)($v2_2))(($v3_3)->value1)))(($v3_3)->value0);
} else {
if ((is_object($v3_3) && (($v3_3)->tag === "DurationOnly"))) {
$__t0 = (($v_0)($v2_2))(($v3_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldr" => function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bifoldableInterval"];
  $__res = ((($GLOBALS['Data_Bifoldable_bifoldrDefault'] ?? \PhpursThunks::eval('Data_Bifoldable_bifoldrDefault')))(($GLOBALS['Data_Interval_bifoldableInterval'] ?? \PhpursThunks::eval('Data_Interval_bifoldableInterval'))))($x_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bifoldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bifoldableInterval"];
  $__res = ((($GLOBALS['Data_Bifoldable_bifoldMapDefaultL'] ?? \PhpursThunks::eval('Data_Bifoldable_bifoldMapDefaultL')))(($GLOBALS['Data_Interval_bifoldableInterval'] ?? \PhpursThunks::eval('Data_Interval_bifoldableInterval'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_bifoldableRecurringInterval'] = function() { $v = (object)["bifoldl" => (function() {
  $__fn = function($f_0, $g_1 = null, $i_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bifoldableRecurringInterval"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v3_3) use ($f_0, $g_1, $i_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bifoldableRecurringInterval"];
  if ((is_object($v3_3) && (($v3_3)->tag === "StartEnd"))) {
$__t0 = (($g_1)((($g_1)($i_2))(($v3_3)->value0)))(($v3_3)->value1);
} else {
if ((is_object($v3_3) && (($v3_3)->tag === "DurationEnd"))) {
$__t0 = (($g_1)((($f_0)($i_2))(($v3_3)->value0)))(($v3_3)->value1);
} else {
if ((is_object($v3_3) && (($v3_3)->tag === "StartDuration"))) {
$__t0 = (($g_1)((($f_0)($i_2))(($v3_3)->value1)))(($v3_3)->value0);
} else {
if ((is_object($v3_3) && (($v3_3)->tag === "DurationOnly"))) {
$__t0 = (($f_0)($i_2))(($v3_3)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Interval_interval'] ?? \PhpursThunks::eval('Data_Interval_interval')));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "bifoldr" => (function() {
  $__fn = function($f_0, $g_1 = null, $i_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bifoldableRecurringInterval"];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((((($GLOBALS['Data_Bifoldable_bifoldrDefault'] ?? \PhpursThunks::eval('Data_Bifoldable_bifoldrDefault')))(($GLOBALS['Data_Interval_bifoldableInterval'] ?? \PhpursThunks::eval('Data_Interval_bifoldableInterval'))))($f_0))($g_1))($i_2)))(($GLOBALS['Data_Interval_interval'] ?? \PhpursThunks::eval('Data_Interval_interval')));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bifoldableRecurringInterval"];
  $__res = ((($GLOBALS['Data_Bifoldable_bifoldMapDefaultL'] ?? \PhpursThunks::eval('Data_Bifoldable_bifoldMapDefaultL')))(($GLOBALS['Data_Interval_bifoldableRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_bifoldableRecurringInterval'))))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_bitraversableInterval'] = function() { $v = (object)["bitraverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableInterval"];
  $Apply0_1_0 = (($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($Apply0_1_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (function() use ($Apply0_1_0, $__local_var_2_1) {
  $__fn = function($v_3, $v1_4 = null, $v2_5 = null) use ($Apply0_1_0, $__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableInterval"];
  if ((is_object($v2_5) && (($v2_5)->tag === "StartEnd"))) {
$__t2 = ((($Apply0_1_0)->apply)(((($__local_var_2_1)->map)(($GLOBALS['Data_Interval_StartEnd'] ?? \PhpursThunks::eval('Data_Interval_StartEnd'))))(($v1_4)(($v2_5)->value0))))(($v1_4)(($v2_5)->value1));
} else {
if ((is_object($v2_5) && (($v2_5)->tag === "DurationEnd"))) {
$__t2 = ((($Apply0_1_0)->apply)(((($__local_var_2_1)->map)(($GLOBALS['Data_Interval_DurationEnd'] ?? \PhpursThunks::eval('Data_Interval_DurationEnd'))))(($v_3)(($v2_5)->value0))))(($v1_4)(($v2_5)->value1));
} else {
if ((is_object($v2_5) && (($v2_5)->tag === "StartDuration"))) {
$__t2 = ((($Apply0_1_0)->apply)(((($__local_var_2_1)->map)(($GLOBALS['Data_Interval_StartDuration'] ?? \PhpursThunks::eval('Data_Interval_StartDuration'))))(($v1_4)(($v2_5)->value0))))(($v_3)(($v2_5)->value1));
} else {
if ((is_object($v2_5) && (($v2_5)->tag === "DurationOnly"))) {
$__t2 = ((($__local_var_2_1)->map)(($GLOBALS['Data_Interval_DurationOnly'] ?? \PhpursThunks::eval('Data_Interval_DurationOnly'))))(($v_3)(($v2_5)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
};
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
}, "bisequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableInterval"];
  $__res = ((((($GLOBALS['Data_Interval_bitraversableInterval'] ?? \PhpursThunks::eval('Data_Interval_bitraversableInterval')))->bitraverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableInterval"];
  $__res = ($GLOBALS['Data_Interval_bifunctorInterval'] ?? \PhpursThunks::eval('Data_Interval_bifunctorInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableInterval"];
  $__res = ($GLOBALS['Data_Interval_bifoldableInterval'] ?? \PhpursThunks::eval('Data_Interval_bifoldableInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Interval_bitraversableRecurringInterval'] = function() { $v = (object)["bitraverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableRecurringInterval"];
  $over1_1_0 = (($GLOBALS['Data_Interval_over'] ?? \PhpursThunks::eval('Data_Interval_over')))((((($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $bitraverse1_2_1 = ((($GLOBALS['Data_Interval_bitraversableInterval'] ?? \PhpursThunks::eval('Data_Interval_bitraversableInterval')))->bitraverse)($dictApplicative_0);
  $__res = (function() use ($bitraverse1_2_1, $over1_1_0) {
  $__fn = function($l_3, $r_4 = null, $i_5 = null) use ($bitraverse1_2_1, $over1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableRecurringInterval"];
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
}, "bisequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableRecurringInterval"];
  $__res = ((((($GLOBALS['Data_Interval_bitraversableRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_bitraversableRecurringInterval')))->bitraverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableRecurringInterval"];
  $__res = ($GLOBALS['Data_Interval_bifunctorRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_bifunctorRecurringInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Interval_bitraversableRecurringInterval"];
  $__res = ($GLOBALS['Data_Interval_bifoldableRecurringInterval'] ?? \PhpursThunks::eval('Data_Interval_bifoldableRecurringInterval'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
































