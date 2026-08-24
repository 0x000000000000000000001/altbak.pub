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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_Interval_StartEnd { public $tag = 'StartEnd'; public function __construct(public  $value0, public  $value1) {} }
final class Data_Interval_DurationEnd { public $tag = 'DurationEnd'; public function __construct(public  $value0, public  $value1) {} }
final class Data_Interval_StartDuration { public $tag = 'StartDuration'; public function __construct(public  $value0, public  $value1) {} }
final class Data_Interval_DurationOnly { public $tag = 'DurationOnly'; public function __construct(public  $value0) {} }
final class Data_Interval_RecurringInterval { public $tag = 'RecurringInterval'; public function __construct(public  $value0, public  $value1) {} }

// Data_Interval_showMaybe
$GLOBALS['Data_Interval_showMaybe'] = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = (("(Just " . \Data\Show\majData_majShow_showmajIntmajImpl(($v_0)->{'value0'})) . ")");
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = "Nothing";
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

// Data_Interval_StartEnd
$GLOBALS['Data_Interval_StartEnd'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Interval\Data_Interval_StartEnd($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_DurationEnd
$GLOBALS['Data_Interval_DurationEnd'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Interval\Data_Interval_DurationEnd($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_StartDuration
$GLOBALS['Data_Interval_StartDuration'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Interval\Data_Interval_StartDuration($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_DurationOnly
$GLOBALS['Data_Interval_DurationOnly'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Interval\Data_Interval_DurationOnly($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Interval_RecurringInterval
$GLOBALS['Data_Interval_RecurringInterval'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Interval\Data_Interval_RecurringInterval($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Interval_showInterval
function majData_majInterval_showmajInterval($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_showmajInterval';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = (((("(StartEnd " . (($dictShow1_1)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow1_1)->{'show'})(($v_2)->{'value1'})) . ")");
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = (((("(DurationEnd " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow1_1)->{'show'})(($v_2)->{'value1'})) . ")");
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = (((("(StartDuration " . (($dictShow1_1)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow_0)->{'show'})(($v_2)->{'value1'})) . ")");
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = (("(DurationOnly " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . ")");
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Interval_showInterval'] = __NAMESPACE__ . '\\majData_majInterval_showmajInterval';

// Data_Interval_showRecurringInterval
function majData_majInterval_showmajRecurringmajInterval($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_showmajRecurringmajInterval';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $showInterval2_2_0 = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = (((("(StartEnd " . (($dictShow1_1)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow1_1)->{'show'})(($v_2)->{'value1'})) . ")");
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = (((("(DurationEnd " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow1_1)->{'show'})(($v_2)->{'value1'})) . ")");
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = (((("(StartDuration " . (($dictShow1_1)->{'show'})(($v_2)->{'value0'})) . " ") . (($dictShow_0)->{'show'})(($v_2)->{'value1'})) . ")");
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = (("(DurationOnly " . (($dictShow_0)->{'show'})(($v_2)->{'value0'})) . ")");
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
  $__res = (object)["show" => function($v_3) use ($showInterval2_2_0) {
  $__num = \func_num_args();
  $__res = (((("(RecurringInterval " . (($GLOBALS['Data_Interval_showMaybe'])->{'show'})(($v_3)->{'value0'})) . " ") . (($showInterval2_2_0)->{'show'})(($v_3)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Interval_showRecurringInterval'] = __NAMESPACE__ . '\\majData_majInterval_showmajRecurringmajInterval';

// Data_Interval_interval
function majData_majInterval_interval($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_interval';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_interval'] = __NAMESPACE__ . '\\majData_majInterval_interval';

// Data_Interval_foldableInterval
$GLOBALS['Data_Interval_foldableInterval'] = (object)["foldl" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = (($v_0)((($v_0)($v1_1))(($v2_2)->{'value0'})))(($v2_2)->{'value1'});
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = (($v_0)($v1_1))(($v2_2)->{'value1'});
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = (($v_0)($v1_1))(($v2_2)->{'value0'});
goto end_branch_0;;
};
  $__t0 = $v1_1;
  end_branch_0:;
  $__res = $__t0;
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
}, "foldr" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($u_1) use ($x_0) {
  $__num = \func_num_args();
  $__res = function($xs_2) use ($u_1, $x_0) {
  $__num = \func_num_args();
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use (&$go__go_3_1, $x_0) {
  $__fn = function($acc_4, $lhs_5 = null, $rhs_6 = null) use (&$go__go_3_1, $x_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__go_3_1_1_acc_4 = $acc_4;
  $__tco_var_go__go_3_1_1_lhs_5 = $lhs_5;
  $__tco_var_go__go_3_1_1_rhs_6 = $rhs_6;
  tco_loop_go__go_3_1_1:;
  $acc_4 = $__tco_var_go__go_3_1_1_acc_4;
  $lhs_5 = $__tco_var_go__go_3_1_1_lhs_5;
  $rhs_6 = $__tco_var_go__go_3_1_1_rhs_6;
  $__t1 = null;;
  if ($rhs_6 instanceof \Data\Foldable\Data_Foldable_Node) {
$__tco_2 = (($x_0)(($rhs_6)->{'value0'}))($acc_4);
$__tco_3 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_4 = $lhs_5;
$__tco_var_go__go_3_1_1_acc_4 = $__tco_2;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_3;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_4;
goto tco_loop_go__go_3_1_1;;
$__t1 = null;
goto end_branch_1;;
};
  if ($rhs_6 instanceof \Data\Foldable\Data_Foldable_Append) {
$__t12 = null;;
if (($rhs_6)->{'value0'} instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_13 = $acc_4;
$__tco_14 = $lhs_5;
$__tco_15 = ($rhs_6)->{'value1'};
$__tco_var_go__go_3_1_1_acc_4 = $__tco_13;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_14;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_15;
goto tco_loop_go__go_3_1_1;;
$__t12 = null;
goto end_branch_12;;
};
$__t8 = null;;
if ($lhs_5 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_9 = $acc_4;
$__tco_10 = ($rhs_6)->{'value0'};
$__tco_11 = ($rhs_6)->{'value1'};
$__tco_var_go__go_3_1_1_acc_4 = $__tco_9;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_10;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_11;
goto tco_loop_go__go_3_1_1;;
$__t8 = null;
goto end_branch_8;;
};
$__tco_5 = $acc_4;
$__tco_6 = new \Data\Foldable\Data_Foldable_Append($lhs_5, ($rhs_6)->{'value0'});
$__tco_7 = ($rhs_6)->{'value1'};
$__tco_var_go__go_3_1_1_acc_4 = $__tco_5;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_6;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_7;
goto tco_loop_go__go_3_1_1;;
$__t8 = null;
end_branch_8:;
$__t12 = $__t8;
end_branch_12:;
$__t1 = $__t12;
goto end_branch_1;;
};
  if ($rhs_6 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t19 = null;;
if ($lhs_5 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t19 = $acc_4;
goto end_branch_19;;
};
$__tco_16 = $acc_4;
$__tco_17 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_18 = $lhs_5;
$__tco_var_go__go_3_1_1_acc_4 = $__tco_16;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_17;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_18;
goto tco_loop_go__go_3_1_1;;
$__t19 = null;
end_branch_19:;
$__t1 = $__t19;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($go__go_3_1)($u_1))(new \Data\Foldable\Data_Foldable_Empty()))((((($GLOBALS['Data_Interval_foldableInterval'])->{'foldMap'})($GLOBALS['Data_Foldable_monoidFreeMonoidTree']))($GLOBALS['Data_Foldable_Node']))($xs_2));
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_3 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_2, $mempty_2_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_foldableInterval'])->{'foldl'})(function($acc_4) use ($Semigroup0_1_2, $f_3) {
  $__num = \func_num_args();
  $__res = function($x_5) use ($Semigroup0_1_2, $acc_4, $f_3) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_2)->{'append'})($acc_4))(($f_3)($x_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_foldableRecurringInterval
$GLOBALS['Data_Interval_foldableRecurringInterval'] = (object)["foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($i_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v2_2) use ($f_0, $i_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = (($f_0)((($f_0)($i_1))(($v2_2)->{'value0'})))(($v2_2)->{'value1'});
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = (($f_0)($i_1))(($v2_2)->{'value1'});
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = (($f_0)($i_1))(($v2_2)->{'value0'});
goto end_branch_0;;
};
  $__t0 = $i_1;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Interval_interval']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($i_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_2) use ($f_0, $i_1) {
  $__num = \func_num_args();
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use ($f_0, &$go__go_3_1) {
  $__fn = function($acc_4, $lhs_5 = null, $rhs_6 = null) use ($f_0, &$go__go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__go_3_1_1_acc_4 = $acc_4;
  $__tco_var_go__go_3_1_1_lhs_5 = $lhs_5;
  $__tco_var_go__go_3_1_1_rhs_6 = $rhs_6;
  tco_loop_go__go_3_1_1:;
  $acc_4 = $__tco_var_go__go_3_1_1_acc_4;
  $lhs_5 = $__tco_var_go__go_3_1_1_lhs_5;
  $rhs_6 = $__tco_var_go__go_3_1_1_rhs_6;
  $__t1 = null;;
  if ($rhs_6 instanceof \Data\Foldable\Data_Foldable_Node) {
$__tco_2 = (($f_0)(($rhs_6)->{'value0'}))($acc_4);
$__tco_3 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_4 = $lhs_5;
$__tco_var_go__go_3_1_1_acc_4 = $__tco_2;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_3;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_4;
goto tco_loop_go__go_3_1_1;;
$__t1 = null;
goto end_branch_1;;
};
  if ($rhs_6 instanceof \Data\Foldable\Data_Foldable_Append) {
$__t12 = null;;
if (($rhs_6)->{'value0'} instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_13 = $acc_4;
$__tco_14 = $lhs_5;
$__tco_15 = ($rhs_6)->{'value1'};
$__tco_var_go__go_3_1_1_acc_4 = $__tco_13;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_14;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_15;
goto tco_loop_go__go_3_1_1;;
$__t12 = null;
goto end_branch_12;;
};
$__t8 = null;;
if ($lhs_5 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_9 = $acc_4;
$__tco_10 = ($rhs_6)->{'value0'};
$__tco_11 = ($rhs_6)->{'value1'};
$__tco_var_go__go_3_1_1_acc_4 = $__tco_9;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_10;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_11;
goto tco_loop_go__go_3_1_1;;
$__t8 = null;
goto end_branch_8;;
};
$__tco_5 = $acc_4;
$__tco_6 = new \Data\Foldable\Data_Foldable_Append($lhs_5, ($rhs_6)->{'value0'});
$__tco_7 = ($rhs_6)->{'value1'};
$__tco_var_go__go_3_1_1_acc_4 = $__tco_5;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_6;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_7;
goto tco_loop_go__go_3_1_1;;
$__t8 = null;
end_branch_8:;
$__t12 = $__t8;
end_branch_12:;
$__t1 = $__t12;
goto end_branch_1;;
};
  if ($rhs_6 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t19 = null;;
if ($lhs_5 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t19 = $acc_4;
goto end_branch_19;;
};
$__tco_16 = $acc_4;
$__tco_17 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_18 = $lhs_5;
$__tco_var_go__go_3_1_1_acc_4 = $__tco_16;
$__tco_var_go__go_3_1_1_lhs_5 = $__tco_17;
$__tco_var_go__go_3_1_1_rhs_6 = $__tco_18;
goto tco_loop_go__go_3_1_1;;
$__t19 = null;
end_branch_19:;
$__t1 = $__t19;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__t2 = null;;
  if ($xs_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t2 = new \Data\Foldable\Data_Foldable_Append(new \Data\Foldable\Data_Foldable_Append(new \Data\Foldable\Data_Foldable_Empty(), new \Data\Foldable\Data_Foldable_Node(($xs_2)->{'value0'})), new \Data\Foldable\Data_Foldable_Node(($xs_2)->{'value1'}));
goto end_branch_2;;
};
  if ($xs_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t2 = new \Data\Foldable\Data_Foldable_Append(new \Data\Foldable\Data_Foldable_Empty(), new \Data\Foldable\Data_Foldable_Node(($xs_2)->{'value1'}));
goto end_branch_2;;
};
  if ($xs_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t2 = new \Data\Foldable\Data_Foldable_Append(new \Data\Foldable\Data_Foldable_Empty(), new \Data\Foldable\Data_Foldable_Node(($xs_2)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = new \Data\Foldable\Data_Foldable_Empty();
  end_branch_2:;
  $__res = ((($go__go_3_1)($i_1))(new \Data\Foldable\Data_Foldable_Empty()))($__t2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Interval_interval']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_3 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_4 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_3, $mempty_2_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_foldableRecurringInterval'])->{'foldl'})(function($acc_4) use ($Semigroup0_1_3, $f_3) {
  $__num = \func_num_args();
  $__res = function($x_5) use ($Semigroup0_1_3, $acc_4, $f_3) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_3)->{'append'})($acc_4))(($f_3)($x_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_eqInterval
function majData_majInterval_eqmajInterval($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_eqmajInterval';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => function($x_2) use ($dictEq1_1, $dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictEq1_1, $dictEq_0, $x_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($x_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = ($y_3 instanceof \Data\Interval\Data_Interval_StartEnd && (((($dictEq1_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && ((($dictEq1_1)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  if ($x_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = ($y_3 instanceof \Data\Interval\Data_Interval_DurationEnd && (((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && ((($dictEq1_1)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  if ($x_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = ($y_3 instanceof \Data\Interval\Data_Interval_StartDuration && (((($dictEq1_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && ((($dictEq_0)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  $__t0 = ($x_2 instanceof \Data\Interval\Data_Interval_DurationOnly && ($y_3 instanceof \Data\Interval\Data_Interval_DurationOnly && ((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Interval_eqInterval'] = __NAMESPACE__ . '\\majData_majInterval_eqmajInterval';

// Data_Interval_eqRecurringInterval
function majData_majInterval_eqmajRecurringmajInterval($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_eqmajRecurringmajInterval';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $eqInterval2_2_0 = (object)["eq" => function($x_2) use ($dictEq1_1, $dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictEq1_1, $dictEq_0, $x_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($x_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = ($y_3 instanceof \Data\Interval\Data_Interval_StartEnd && (((($dictEq1_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && ((($dictEq1_1)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  if ($x_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = ($y_3 instanceof \Data\Interval\Data_Interval_DurationEnd && (((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && ((($dictEq1_1)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  if ($x_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = ($y_3 instanceof \Data\Interval\Data_Interval_StartDuration && (((($dictEq1_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && ((($dictEq_0)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'})));
goto end_branch_0;;
};
  $__t0 = ($x_2 instanceof \Data\Interval\Data_Interval_DurationOnly && ($y_3 instanceof \Data\Interval\Data_Interval_DurationOnly && ((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["eq" => function($x_3) use ($eqInterval2_2_0) {
  $__num = \func_num_args();
  $__res = function($y_4) use ($eqInterval2_2_0, $x_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($x_3)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = ($y_4)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing;
goto end_branch_2;;
};
  $__t2 = (($x_3)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just && (($y_4)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just && ((($x_3)->{'value0'})->{'value0'} === (($y_4)->{'value0'})->{'value0'})));
  end_branch_2:;
  $__res = ($__t2 && ((($eqInterval2_2_0)->{'eq'})(($x_3)->{'value1'}))(($y_4)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Interval_eqRecurringInterval'] = __NAMESPACE__ . '\\majData_majInterval_eqmajRecurringmajInterval';

// Data_Interval_ordInterval
function majData_majInterval_ordmajInterval($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_ordmajInterval';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $__res = function($dictOrd1_2) use ($__local_var_1_0, $dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd1_2)->{'Eq0'})(null);
  $eqInterval2_3_1 = (object)["eq" => function($x_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  $__res = function($y_5) use ($__local_var_1_0, $__local_var_3_1, $x_4) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($x_4 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t2 = ($y_5 instanceof \Data\Interval\Data_Interval_StartEnd && (((($__local_var_3_1)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'}) && ((($__local_var_3_1)->{'eq'})(($x_4)->{'value1'}))(($y_5)->{'value1'})));
goto end_branch_2;;
};
  if ($x_4 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t2 = ($y_5 instanceof \Data\Interval\Data_Interval_DurationEnd && (((($__local_var_1_0)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'}) && ((($__local_var_3_1)->{'eq'})(($x_4)->{'value1'}))(($y_5)->{'value1'})));
goto end_branch_2;;
};
  if ($x_4 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t2 = ($y_5 instanceof \Data\Interval\Data_Interval_StartDuration && (((($__local_var_3_1)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'}) && ((($__local_var_1_0)->{'eq'})(($x_4)->{'value1'}))(($y_5)->{'value1'})));
goto end_branch_2;;
};
  $__t2 = ($x_4 instanceof \Data\Interval\Data_Interval_DurationOnly && ($y_5 instanceof \Data\Interval\Data_Interval_DurationOnly && ((($__local_var_1_0)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'})));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($x_4) use ($dictOrd1_2, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($y_5) use ($dictOrd1_2, $dictOrd_0, $x_4) {
  $__num = \func_num_args();
  $__t4 = null;;
  if ($x_4 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t5 = null;;
if ($y_5 instanceof \Data\Interval\Data_Interval_StartEnd) {
$v_6_6 = ((($dictOrd1_2)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
$__t7 = null;;
if ($v_6_6 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t7 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_7;;
};
if ($v_6_6 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t7 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_7;;
};
$__t7 = ((($dictOrd1_2)->{'compare'})(($x_4)->{'value1'}))(($y_5)->{'value1'});
end_branch_7:;
$__t5 = $__t7;
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($y_5 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if ($x_4 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t8 = null;;
if ($y_5 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$v_6_9 = ((($dictOrd_0)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
$__t10 = null;;
if ($v_6_9 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t10 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_10;;
};
if ($v_6_9 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t10 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_10;;
};
$__t10 = ((($dictOrd1_2)->{'compare'})(($x_4)->{'value1'}))(($y_5)->{'value1'});
end_branch_10:;
$__t8 = $__t10;
goto end_branch_8;;
};
$__t8 = new \Data\Ordering\Data_Ordering_LT();
end_branch_8:;
$__t4 = $__t8;
goto end_branch_4;;
};
  if ($y_5 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if ($x_4 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t11 = null;;
if ($y_5 instanceof \Data\Interval\Data_Interval_StartDuration) {
$v_6_12 = ((($dictOrd1_2)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
$__t13 = null;;
if ($v_6_12 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t13 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_13;;
};
if ($v_6_12 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t13 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_13;;
};
$__t13 = ((($dictOrd_0)->{'compare'})(($x_4)->{'value1'}))(($y_5)->{'value1'});
end_branch_13:;
$__t11 = $__t13;
goto end_branch_11;;
};
$__t11 = new \Data\Ordering\Data_Ordering_LT();
end_branch_11:;
$__t4 = $__t11;
goto end_branch_4;;
};
  if ($y_5 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($x_4 instanceof \Data\Interval\Data_Interval_DurationOnly && $y_5 instanceof \Data\Interval\Data_Interval_DurationOnly)) {
$__t4 = ((($dictOrd_0)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_4) use ($eqInterval2_3_1) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_ordInterval'] = __NAMESPACE__ . '\\majData_majInterval_ordmajInterval';

// Data_Interval_ordRecurringInterval
function majData_majInterval_ordmajRecurringmajInterval($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majInterval_ordmajRecurringmajInterval';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $__local_var_2_1 = (($dictOrd_0)->{'Eq0'})(null);
  $__res = function($dictOrd1_3) use ($__local_var_1_0, $__local_var_2_1, $dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictOrd1_3)->{'Eq0'})(null);
  $eqInterval2_4_2 = (object)["eq" => function($x_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($y_6) use ($__local_var_1_0, $__local_var_4_2, $x_5) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($x_5 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t3 = ($y_6 instanceof \Data\Interval\Data_Interval_StartEnd && (((($__local_var_4_2)->{'eq'})(($x_5)->{'value0'}))(($y_6)->{'value0'}) && ((($__local_var_4_2)->{'eq'})(($x_5)->{'value1'}))(($y_6)->{'value1'})));
goto end_branch_3;;
};
  if ($x_5 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t3 = ($y_6 instanceof \Data\Interval\Data_Interval_DurationEnd && (((($__local_var_1_0)->{'eq'})(($x_5)->{'value0'}))(($y_6)->{'value0'}) && ((($__local_var_4_2)->{'eq'})(($x_5)->{'value1'}))(($y_6)->{'value1'})));
goto end_branch_3;;
};
  if ($x_5 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t3 = ($y_6 instanceof \Data\Interval\Data_Interval_StartDuration && (((($__local_var_4_2)->{'eq'})(($x_5)->{'value0'}))(($y_6)->{'value0'}) && ((($__local_var_1_0)->{'eq'})(($x_5)->{'value1'}))(($y_6)->{'value1'})));
goto end_branch_3;;
};
  $__t3 = ($x_5 instanceof \Data\Interval\Data_Interval_DurationOnly && ($y_6 instanceof \Data\Interval\Data_Interval_DurationOnly && ((($__local_var_1_0)->{'eq'})(($x_5)->{'value0'}))(($y_6)->{'value0'})));
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $ordInterval2_4_2 = (object)["compare" => function($x_5) use ($dictOrd1_3, $dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($y_6) use ($dictOrd1_3, $dictOrd_0, $x_5) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($x_5 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t6 = null;;
if ($y_6 instanceof \Data\Interval\Data_Interval_StartEnd) {
$v_7_7 = ((($dictOrd1_3)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
$__t8 = null;;
if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t8 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_8;;
};
if ($v_7_7 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t8 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_8;;
};
$__t8 = ((($dictOrd1_3)->{'compare'})(($x_5)->{'value1'}))(($y_6)->{'value1'});
end_branch_8:;
$__t6 = $__t8;
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
  if ($y_6 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  if ($x_5 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t9 = null;;
if ($y_6 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$v_7_10 = ((($dictOrd_0)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
$__t11 = null;;
if ($v_7_10 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t11 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_11;;
};
if ($v_7_10 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t11 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_11;;
};
$__t11 = ((($dictOrd1_3)->{'compare'})(($x_5)->{'value1'}))(($y_6)->{'value1'});
end_branch_11:;
$__t9 = $__t11;
goto end_branch_9;;
};
$__t9 = new \Data\Ordering\Data_Ordering_LT();
end_branch_9:;
$__t5 = $__t9;
goto end_branch_5;;
};
  if ($y_6 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  if ($x_5 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t12 = null;;
if ($y_6 instanceof \Data\Interval\Data_Interval_StartDuration) {
$v_7_13 = ((($dictOrd1_3)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
$__t14 = null;;
if ($v_7_13 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t14 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_14;;
};
if ($v_7_13 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t14 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_14;;
};
$__t14 = ((($dictOrd_0)->{'compare'})(($x_5)->{'value1'}))(($y_6)->{'value1'});
end_branch_14:;
$__t12 = $__t14;
goto end_branch_12;;
};
$__t12 = new \Data\Ordering\Data_Ordering_LT();
end_branch_12:;
$__t5 = $__t12;
goto end_branch_5;;
};
  if ($y_6 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  if (($x_5 instanceof \Data\Interval\Data_Interval_DurationOnly && $y_6 instanceof \Data\Interval\Data_Interval_DurationOnly)) {
$__t5 = ((($dictOrd_0)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
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
}, "Eq0" => function($_dollar___unused_5) use ($eqInterval2_4_2) {
  $__num = \func_num_args();
  $__res = $eqInterval2_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_16 = (($dictOrd1_3)->{'Eq0'})(null);
  $eqInterval2_6_17 = (object)["eq" => function($x_6) use ($__local_var_2_1, $__local_var_5_16) {
  $__num = \func_num_args();
  $__res = function($y_7) use ($__local_var_2_1, $__local_var_5_16, $x_6) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ($x_6 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t17 = ($y_7 instanceof \Data\Interval\Data_Interval_StartEnd && (((($__local_var_5_16)->{'eq'})(($x_6)->{'value0'}))(($y_7)->{'value0'}) && ((($__local_var_5_16)->{'eq'})(($x_6)->{'value1'}))(($y_7)->{'value1'})));
goto end_branch_17;;
};
  if ($x_6 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t17 = ($y_7 instanceof \Data\Interval\Data_Interval_DurationEnd && (((($__local_var_2_1)->{'eq'})(($x_6)->{'value0'}))(($y_7)->{'value0'}) && ((($__local_var_5_16)->{'eq'})(($x_6)->{'value1'}))(($y_7)->{'value1'})));
goto end_branch_17;;
};
  if ($x_6 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t17 = ($y_7 instanceof \Data\Interval\Data_Interval_StartDuration && (((($__local_var_5_16)->{'eq'})(($x_6)->{'value0'}))(($y_7)->{'value0'}) && ((($__local_var_2_1)->{'eq'})(($x_6)->{'value1'}))(($y_7)->{'value1'})));
goto end_branch_17;;
};
  $__t17 = ($x_6 instanceof \Data\Interval\Data_Interval_DurationOnly && ($y_7 instanceof \Data\Interval\Data_Interval_DurationOnly && ((($__local_var_2_1)->{'eq'})(($x_6)->{'value0'}))(($y_7)->{'value0'})));
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $eqRecurringInterval2_5_16 = (object)["eq" => function($x_7) use ($eqInterval2_6_17) {
  $__num = \func_num_args();
  $__res = function($y_8) use ($eqInterval2_6_17, $x_7) {
  $__num = \func_num_args();
  $__t19 = null;;
  if (($x_7)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t19 = ($y_8)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing;
goto end_branch_19;;
};
  $__t19 = (($x_7)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just && (($y_8)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just && ((($x_7)->{'value0'})->{'value0'} === (($y_8)->{'value0'})->{'value0'})));
  end_branch_19:;
  $__res = ($__t19 && ((($eqInterval2_6_17)->{'eq'})(($x_7)->{'value1'}))(($y_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($x_6) use ($ordInterval2_4_2) {
  $__num = \func_num_args();
  $__res = function($y_7) use ($ordInterval2_4_2, $x_6) {
  $__num = \func_num_args();
  $__t21 = null;;
  if (($x_6)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t22 = null;;
if (($y_7)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t22 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_22;;
};
$__t22 = new \Data\Ordering\Data_Ordering_LT();
end_branch_22:;
$__t21 = $__t22;
goto end_branch_21;;
};
  if (($y_7)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t21 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_21;;
};
  if ((($x_6)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just && ($y_7)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just)) {
$__t21 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), (($x_6)->{'value0'})->{'value0'}, (($y_7)->{'value0'})->{'value0'});
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $v_8_21 = $__t21;
  $__t24 = null;;
  if ($v_8_21 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t24 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_24;;
};
  if ($v_8_21 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t24 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_24;;
};
  $__t24 = ((($ordInterval2_4_2)->{'compare'})(($x_6)->{'value1'}))(($y_7)->{'value1'});
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_6) use ($eqRecurringInterval2_5_16) {
  $__num = \func_num_args();
  $__res = $eqRecurringInterval2_5_16;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Interval_ordRecurringInterval'] = __NAMESPACE__ . '\\majData_majInterval_ordmajRecurringmajInterval';

// Data_Interval_bifunctorInterval
$GLOBALS['Data_Interval_bifunctorInterval'] = (object)["bimap" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = new \Data\Interval\Data_Interval_StartEnd(($v1_1)(($v2_2)->{'value0'}), ($v1_1)(($v2_2)->{'value1'}));
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = new \Data\Interval\Data_Interval_DurationEnd(($v_0)(($v2_2)->{'value0'}), ($v1_1)(($v2_2)->{'value1'}));
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = new \Data\Interval\Data_Interval_StartDuration(($v1_1)(($v2_2)->{'value0'}), ($v_0)(($v2_2)->{'value1'}));
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = new \Data\Interval\Data_Interval_DurationOnly(($v_0)(($v2_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
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
}];

// Data_Interval_bifunctorRecurringInterval
$GLOBALS['Data_Interval_bifunctorRecurringInterval'] = (object)["bimap" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($g_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $g_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_2)->{'value1'} instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = new \Data\Interval\Data_Interval_StartEnd(($g_1)((($v_2)->{'value1'})->{'value0'}), ($g_1)((($v_2)->{'value1'})->{'value1'}));
goto end_branch_0;;
};
  if (($v_2)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = new \Data\Interval\Data_Interval_DurationEnd(($f_0)((($v_2)->{'value1'})->{'value0'}), ($g_1)((($v_2)->{'value1'})->{'value1'}));
goto end_branch_0;;
};
  if (($v_2)->{'value1'} instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = new \Data\Interval\Data_Interval_StartDuration(($g_1)((($v_2)->{'value1'})->{'value0'}), ($f_0)((($v_2)->{'value1'})->{'value1'}));
goto end_branch_0;;
};
  if (($v_2)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = new \Data\Interval\Data_Interval_DurationOnly(($f_0)((($v_2)->{'value1'})->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = new \Data\Interval\Data_Interval_RecurringInterval(($v_2)->{'value0'}, $__t0);
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
}];

// Data_Interval_functorInterval
$GLOBALS['Data_Interval_functorInterval'] = (object)["map" => function($v1_0) {
  $__num = \func_num_args();
  $__res = function($v2_1) use ($v1_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_1 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = new \Data\Interval\Data_Interval_StartEnd(($v1_0)(($v2_1)->{'value0'}), ($v1_0)(($v2_1)->{'value1'}));
goto end_branch_0;;
};
  if ($v2_1 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = new \Data\Interval\Data_Interval_DurationEnd(($v2_1)->{'value0'}, ($v1_0)(($v2_1)->{'value1'}));
goto end_branch_0;;
};
  if ($v2_1 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = new \Data\Interval\Data_Interval_StartDuration(($v1_0)(($v2_1)->{'value0'}), ($v2_1)->{'value1'});
goto end_branch_0;;
};
  if ($v2_1 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = new \Data\Interval\Data_Interval_DurationOnly(($v2_1)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_extendInterval
$GLOBALS['Data_Interval_extendInterval'] = (object)["extend" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = new \Data\Interval\Data_Interval_StartEnd(($v_0)($v1_1), ($v_0)($v1_1));
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = new \Data\Interval\Data_Interval_DurationEnd(($v1_1)->{'value0'}, ($v_0)($v1_1));
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = new \Data\Interval\Data_Interval_StartDuration(($v_0)($v1_1), ($v1_1)->{'value1'});
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = new \Data\Interval\Data_Interval_DurationOnly(($v1_1)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_functorRecurringInterval
$GLOBALS['Data_Interval_functorRecurringInterval'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = new \Data\Interval\Data_Interval_StartEnd(($f_0)((($v_1)->{'value1'})->{'value0'}), ($f_0)((($v_1)->{'value1'})->{'value1'}));
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = new \Data\Interval\Data_Interval_DurationEnd((($v_1)->{'value1'})->{'value0'}, ($f_0)((($v_1)->{'value1'})->{'value1'}));
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = new \Data\Interval\Data_Interval_StartDuration(($f_0)((($v_1)->{'value1'})->{'value0'}), (($v_1)->{'value1'})->{'value1'});
goto end_branch_0;;
};
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = new \Data\Interval\Data_Interval_DurationOnly((($v_1)->{'value1'})->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = new \Data\Interval\Data_Interval_RecurringInterval(($v_1)->{'value0'}, $__t0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_extendRecurringInterval
$GLOBALS['Data_Interval_extendRecurringInterval'] = (object)["extend" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = ($f_0)($v_1);
  $__local_var_2_0 = function($v_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = $__local_var_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__t2 = null;;
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t2 = new \Data\Interval\Data_Interval_StartEnd(($__local_var_2_0)(($v_1)->{'value1'}), ($__local_var_2_0)(($v_1)->{'value1'}));
goto end_branch_2;;
};
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t2 = new \Data\Interval\Data_Interval_DurationEnd((($v_1)->{'value1'})->{'value0'}, ($__local_var_2_0)(($v_1)->{'value1'}));
goto end_branch_2;;
};
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t2 = new \Data\Interval\Data_Interval_StartDuration(($__local_var_2_0)(($v_1)->{'value1'}), (($v_1)->{'value1'})->{'value1'});
goto end_branch_2;;
};
  if (($v_1)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t2 = new \Data\Interval\Data_Interval_DurationOnly((($v_1)->{'value1'})->{'value0'});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = new \Data\Interval\Data_Interval_RecurringInterval(($v_1)->{'value0'}, $__t2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_traversableInterval
$GLOBALS['Data_Interval_traversableInterval'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_3) use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($Apply0_1_0, $Functor0_2_1, $dictApplicative_0, $v_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v1_4 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t2 = ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_Interval_StartEnd']))(($v_3)(($v1_4)->{'value0'}))))(($v_3)(($v1_4)->{'value1'}));
goto end_branch_2;;
};
  if ($v1_4 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t2 = ((($Functor0_2_1)->{'map'})(($GLOBALS['Data_Interval_DurationEnd'])(($v1_4)->{'value0'})))(($v_3)(($v1_4)->{'value1'}));
goto end_branch_2;;
};
  if ($v1_4 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__local_var_5_3 = ($v1_4)->{'value1'};
$__t2 = ((($Functor0_2_1)->{'map'})(function($v2_6) use ($__local_var_5_3) {
  $__num = \func_num_args();
  $__res = new \Data\Interval\Data_Interval_StartDuration($v2_6, $__local_var_5_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_3)(($v1_4)->{'value0'}));
goto end_branch_2;;
};
  if ($v1_4 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t2 = (($dictApplicative_0)->{'pure'})(new \Data\Interval\Data_Interval_DurationOnly(($v1_4)->{'value0'}));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_traversableInterval'])->{'traverse'})($dictApplicative_0))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_foldableInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_traversableRecurringInterval
$GLOBALS['Data_Interval_traversableRecurringInterval'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($i_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $Apply0_4_1 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_5_2 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t3 = null;;
  if (($i_3)->{'value1'} instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t3 = ((($Apply0_4_1)->{'apply'})(((($Functor0_5_2)->{'map'})($GLOBALS['Data_Interval_StartEnd']))(($f_2)((($i_3)->{'value1'})->{'value0'}))))(($f_2)((($i_3)->{'value1'})->{'value1'}));
goto end_branch_3;;
};
  if (($i_3)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t3 = ((($Functor0_5_2)->{'map'})(($GLOBALS['Data_Interval_DurationEnd'])((($i_3)->{'value1'})->{'value0'})))(($f_2)((($i_3)->{'value1'})->{'value1'}));
goto end_branch_3;;
};
  if (($i_3)->{'value1'} instanceof \Data\Interval\Data_Interval_StartDuration) {
$__local_var_6_4 = (($i_3)->{'value1'})->{'value1'};
$__t3 = ((($Functor0_5_2)->{'map'})(function($v2_7) use ($__local_var_6_4) {
  $__num = \func_num_args();
  $__res = new \Data\Interval\Data_Interval_StartDuration($v2_7, $__local_var_6_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)((($i_3)->{'value1'})->{'value0'}));
goto end_branch_3;;
};
  if (($i_3)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t3 = (($dictApplicative_0)->{'pure'})(new \Data\Interval\Data_Interval_DurationOnly((($i_3)->{'value1'})->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = ((($Functor0_1_0)->{'map'})(($GLOBALS['Data_Interval_RecurringInterval'])(($i_3)->{'value0'})))($__t3);
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Interval_traversableRecurringInterval'])->{'traverse'})($dictApplicative_0))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_functorRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_foldableRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_bifoldableInterval
$GLOBALS['Data_Interval_bifoldableInterval'] = (object)["bifoldl" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__res = function($v3_3) use ($v1_1, $v2_2, $v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v3_3 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = (($v1_1)((($v1_1)($v2_2))(($v3_3)->{'value0'})))(($v3_3)->{'value1'});
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = (($v1_1)((($v_0)($v2_2))(($v3_3)->{'value0'})))(($v3_3)->{'value1'});
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = (($v1_1)((($v_0)($v2_2))(($v3_3)->{'value1'})))(($v3_3)->{'value0'});
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = (($v_0)($v2_2))(($v3_3)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bifoldr" => function($x_0) {
  $__num = \func_num_args();
  $__res = function($g_1) use ($x_0) {
  $__num = \func_num_args();
  $__res = function($z_2) use ($g_1, $x_0) {
  $__num = \func_num_args();
  $__res = function($p_3) use ($g_1, $x_0, $z_2) {
  $__num = \func_num_args();
  $semigroupEndo1_4_1 = (object)["append" => function($v_4) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($v_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (((((($GLOBALS['Data_Interval_bifoldableInterval'])->{'bifoldMap'})((object)["mempty" => function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_5) use ($semigroupEndo1_4_1) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_4_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($x_0)))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($g_1)))($p_3))($z_2);
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bifoldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_3 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_2, $mempty_2_3) {
  $__num = \func_num_args();
  $__res = function($g_4) use ($Semigroup0_1_2, $f_3, $mempty_2_3) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Interval_bifoldableInterval'])->{'bifoldl'})(function($m_5) use ($Semigroup0_1_2, $f_3) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Semigroup0_1_2, $f_3, $m_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_2)->{'append'})($m_5))(($f_3)($a_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($m_5) use ($Semigroup0_1_2, $g_4) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Semigroup0_1_2, $g_4, $m_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_2)->{'append'})($m_5))(($g_4)($b_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_3);
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
}];

// Data_Interval_bifoldableRecurringInterval
$GLOBALS['Data_Interval_bifoldableRecurringInterval'] = (object)["bifoldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($g_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($i_2) use ($f_0, $g_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v3_3) use ($f_0, $g_1, $i_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v3_3 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t0 = (($g_1)((($g_1)($i_2))(($v3_3)->{'value0'})))(($v3_3)->{'value1'});
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t0 = (($g_1)((($f_0)($i_2))(($v3_3)->{'value0'})))(($v3_3)->{'value1'});
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t0 = (($g_1)((($f_0)($i_2))(($v3_3)->{'value1'})))(($v3_3)->{'value0'});
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t0 = (($f_0)($i_2))(($v3_3)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Interval_interval']);
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
}, "bifoldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($g_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($i_2) use ($f_0, $g_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($p_3) use ($f_0, $g_1, $i_2) {
  $__num = \func_num_args();
  $semigroupEndo1_4_1 = (object)["append" => function($v_4) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($v_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_4_1 = (object)["mempty" => function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_5) use ($semigroupEndo1_4_1) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_4_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $Semigroup0_5_3 = (($__local_var_4_1)->{'Semigroup0'})(null);
  $__local_var_6_4 = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_0);
  $__local_var_7_5 = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($g_1);
  $__t6 = null;;
  if ($p_3 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t6 = (((($Semigroup0_5_3)->{'append'})(((($Semigroup0_5_3)->{'append'})(($__local_var_4_1)->{'mempty'}))(($__local_var_7_5)(($p_3)->{'value0'}))))(($__local_var_7_5)(($p_3)->{'value1'})))($i_2);
goto end_branch_6;;
};
  if ($p_3 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t6 = (((($Semigroup0_5_3)->{'append'})(((($Semigroup0_5_3)->{'append'})(($__local_var_4_1)->{'mempty'}))(($__local_var_6_4)(($p_3)->{'value0'}))))(($__local_var_7_5)(($p_3)->{'value1'})))($i_2);
goto end_branch_6;;
};
  if ($p_3 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t6 = (((($Semigroup0_5_3)->{'append'})(((($Semigroup0_5_3)->{'append'})(($__local_var_4_1)->{'mempty'}))(($__local_var_6_4)(($p_3)->{'value1'}))))(($__local_var_7_5)(($p_3)->{'value0'})))($i_2);
goto end_branch_6;;
};
  if ($p_3 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t6 = (((($Semigroup0_5_3)->{'append'})(($__local_var_4_1)->{'mempty'}))(($__local_var_6_4)(($p_3)->{'value0'})))($i_2);
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Interval_interval']);
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
}, "bifoldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_7 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_8 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_7, $mempty_2_8) {
  $__num = \func_num_args();
  $__res = function($g_4) use ($Semigroup0_1_7, $f_3, $mempty_2_8) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Interval_bifoldableRecurringInterval'])->{'bifoldl'})(function($m_5) use ($Semigroup0_1_7, $f_3) {
  $__num = \func_num_args();
  $__res = function($a_6) use ($Semigroup0_1_7, $f_3, $m_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_7)->{'append'})($m_5))(($f_3)($a_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($m_5) use ($Semigroup0_1_7, $g_4) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($Semigroup0_1_7, $g_4, $m_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_7)->{'append'})($m_5))(($g_4)($b_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_8);
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
}];

// Data_Interval_bitraversableInterval
$GLOBALS['Data_Interval_bitraversableInterval'] = (object)["bitraverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_3) use ($Apply0_1_0, $Functor0_2_1) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($Apply0_1_0, $Functor0_2_1, $v_3) {
  $__num = \func_num_args();
  $__res = function($v2_5) use ($Apply0_1_0, $Functor0_2_1, $v1_4, $v_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v2_5 instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t2 = ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_Interval_StartEnd']))(($v1_4)(($v2_5)->{'value0'}))))(($v1_4)(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ($v2_5 instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t2 = ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_Interval_DurationEnd']))(($v_3)(($v2_5)->{'value0'}))))(($v1_4)(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ($v2_5 instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t2 = ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_Interval_StartDuration']))(($v1_4)(($v2_5)->{'value0'}))))(($v_3)(($v2_5)->{'value1'}));
goto end_branch_2;;
};
  if ($v2_5 instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t2 = ((($Functor0_2_1)->{'map'})($GLOBALS['Data_Interval_DurationOnly']))(($v_3)(($v2_5)->{'value0'}));
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Interval_bitraversableInterval'])->{'bitraverse'})($dictApplicative_0))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifunctorInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifoldableInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Interval_bitraversableRecurringInterval
$GLOBALS['Data_Interval_bitraversableRecurringInterval'] = (object)["bitraverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($l_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($r_3) use ($Functor0_1_0, $dictApplicative_0, $l_2) {
  $__num = \func_num_args();
  $__res = function($i_4) use ($Functor0_1_0, $dictApplicative_0, $l_2, $r_3) {
  $__num = \func_num_args();
  $Apply0_5_1 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_6_2 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t3 = null;;
  if (($i_4)->{'value1'} instanceof \Data\Interval\Data_Interval_StartEnd) {
$__t3 = ((($Apply0_5_1)->{'apply'})(((($Functor0_6_2)->{'map'})($GLOBALS['Data_Interval_StartEnd']))(($r_3)((($i_4)->{'value1'})->{'value0'}))))(($r_3)((($i_4)->{'value1'})->{'value1'}));
goto end_branch_3;;
};
  if (($i_4)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationEnd) {
$__t3 = ((($Apply0_5_1)->{'apply'})(((($Functor0_6_2)->{'map'})($GLOBALS['Data_Interval_DurationEnd']))(($l_2)((($i_4)->{'value1'})->{'value0'}))))(($r_3)((($i_4)->{'value1'})->{'value1'}));
goto end_branch_3;;
};
  if (($i_4)->{'value1'} instanceof \Data\Interval\Data_Interval_StartDuration) {
$__t3 = ((($Apply0_5_1)->{'apply'})(((($Functor0_6_2)->{'map'})($GLOBALS['Data_Interval_StartDuration']))(($r_3)((($i_4)->{'value1'})->{'value0'}))))(($l_2)((($i_4)->{'value1'})->{'value1'}));
goto end_branch_3;;
};
  if (($i_4)->{'value1'} instanceof \Data\Interval\Data_Interval_DurationOnly) {
$__t3 = ((($Functor0_6_2)->{'map'})($GLOBALS['Data_Interval_DurationOnly']))(($l_2)((($i_4)->{'value1'})->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = ((($Functor0_1_0)->{'map'})(($GLOBALS['Data_Interval_RecurringInterval'])(($i_4)->{'value0'})))($__t3);
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "bisequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Interval_bitraversableRecurringInterval'])->{'bitraverse'})($dictApplicative_0))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifunctor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifunctorRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bifoldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Interval_bifoldableRecurringInterval'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

