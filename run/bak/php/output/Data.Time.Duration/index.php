<?php

namespace Data\Time\Duration;

// ALL IMPORTS: Control.Category, Control.Semigroupoid, Data.Eq, Data.EuclideanRing, Data.Monoid, Data.Newtype, Data.Ord, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time.Duration, Prelude, Prim
// TO REQUIRE: Control.Category, Control.Semigroupoid, Data.Eq, Data.EuclideanRing, Data.Monoid, Data.Newtype, Data.Ord, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Time.Duration, Prelude
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.EuclideanRing/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Time.Duration/index.php';
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


// Data_Time_Duration_negate
$GLOBALS['Data_Time_Duration_negate'] = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = ( - $a_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_Seconds
$GLOBALS['Data_Time_Duration_Seconds'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_Minutes
$GLOBALS['Data_Time_Duration_Minutes'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_Milliseconds
$GLOBALS['Data_Time_Duration_Milliseconds'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_Hours
$GLOBALS['Data_Time_Duration_Hours'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_Days
$GLOBALS['Data_Time_Duration_Days'] = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_toDuration
$GLOBALS['Data_Time_Duration_toDuration'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['toDuration'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_showSeconds
$GLOBALS['Data_Time_Duration_showSeconds'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Seconds " . ($GLOBALS['Data_Show_showNumberImpl'])($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_showMinutes
$GLOBALS['Data_Time_Duration_showMinutes'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Minutes " . ($GLOBALS['Data_Show_showNumberImpl'])($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_showMilliseconds
$GLOBALS['Data_Time_Duration_showMilliseconds'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Milliseconds " . ($GLOBALS['Data_Show_showNumberImpl'])($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_showHours
$GLOBALS['Data_Time_Duration_showHours'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Hours " . ($GLOBALS['Data_Show_showNumberImpl'])($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_showDays
$GLOBALS['Data_Time_Duration_showDays'] = ["show" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = (("(Days " . ($GLOBALS['Data_Show_showNumberImpl'])($v_0)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_semigroupSeconds
$GLOBALS['Data_Time_Duration_semigroupSeconds'] = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0 + $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Time_Duration_semigroupMinutes
$GLOBALS['Data_Time_Duration_semigroupMinutes'] = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0 + $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Time_Duration_semigroupMilliseconds
$GLOBALS['Data_Time_Duration_semigroupMilliseconds'] = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0 + $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Time_Duration_semigroupHours
$GLOBALS['Data_Time_Duration_semigroupHours'] = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0 + $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Time_Duration_semigroupDays
$GLOBALS['Data_Time_Duration_semigroupDays'] = ["append" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v_0 + $v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Time_Duration_ordSeconds
$GLOBALS['Data_Time_Duration_ordSeconds'] = $GLOBALS['Data_Ord_ordNumber'];

// Data_Time_Duration_ordMinutes
$GLOBALS['Data_Time_Duration_ordMinutes'] = $GLOBALS['Data_Ord_ordNumber'];

// Data_Time_Duration_ordMilliseconds
$GLOBALS['Data_Time_Duration_ordMilliseconds'] = $GLOBALS['Data_Ord_ordNumber'];

// Data_Time_Duration_ordHours
$GLOBALS['Data_Time_Duration_ordHours'] = $GLOBALS['Data_Ord_ordNumber'];

// Data_Time_Duration_ordDays
$GLOBALS['Data_Time_Duration_ordDays'] = $GLOBALS['Data_Ord_ordNumber'];

// Data_Time_Duration_newtypeSeconds
$GLOBALS['Data_Time_Duration_newtypeSeconds'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_newtypeMinutes
$GLOBALS['Data_Time_Duration_newtypeMinutes'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_newtypeMilliseconds
$GLOBALS['Data_Time_Duration_newtypeMilliseconds'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_newtypeHours
$GLOBALS['Data_Time_Duration_newtypeHours'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_newtypeDays
$GLOBALS['Data_Time_Duration_newtypeDays'] = ["Coercible0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Prim_undefined'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_monoidSeconds
$GLOBALS['Data_Time_Duration_monoidSeconds'] = ["mempty" => 0.0, "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_Duration_semigroupSeconds'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_monoidMinutes
$GLOBALS['Data_Time_Duration_monoidMinutes'] = ["mempty" => 0.0, "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_Duration_semigroupMinutes'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_monoidMilliseconds
$GLOBALS['Data_Time_Duration_monoidMilliseconds'] = ["mempty" => 0.0, "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_Duration_semigroupMilliseconds'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_monoidHours
$GLOBALS['Data_Time_Duration_monoidHours'] = ["mempty" => 0.0, "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_Duration_semigroupHours'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_monoidDays
$GLOBALS['Data_Time_Duration_monoidDays'] = ["mempty" => 0.0, "Semigroup0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Time_Duration_semigroupDays'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_fromDuration
$GLOBALS['Data_Time_Duration_fromDuration'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['fromDuration'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_negateDuration
$GLOBALS['Data_Time_Duration_negateDuration'] = function($dictDuration_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictDuration_0)['toDuration']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Time_Duration_negate']))(($dictDuration_0)['fromDuration']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Time_Duration_eqSeconds
$GLOBALS['Data_Time_Duration_eqSeconds'] = $GLOBALS['Data_Eq_eqNumber'];

// Data_Time_Duration_eqMinutes
$GLOBALS['Data_Time_Duration_eqMinutes'] = $GLOBALS['Data_Eq_eqNumber'];

// Data_Time_Duration_eqMilliseconds
$GLOBALS['Data_Time_Duration_eqMilliseconds'] = $GLOBALS['Data_Eq_eqNumber'];

// Data_Time_Duration_eqHours
$GLOBALS['Data_Time_Duration_eqHours'] = $GLOBALS['Data_Eq_eqNumber'];

// Data_Time_Duration_eqDays
$GLOBALS['Data_Time_Duration_eqDays'] = $GLOBALS['Data_Eq_eqNumber'];

// Data_Time_Duration_durationSeconds
$GLOBALS['Data_Time_Duration_durationSeconds'] = ["fromDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 * 1000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "toDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 / 1000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_durationMinutes
$GLOBALS['Data_Time_Duration_durationMinutes'] = ["fromDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 * 60000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "toDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 / 60000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_durationMilliseconds
$GLOBALS['Data_Time_Duration_durationMilliseconds'] = ["fromDuration" => ($GLOBALS['Control_Category_categoryFn'])['identity'], "toDuration" => ($GLOBALS['Control_Category_categoryFn'])['identity']];

// Data_Time_Duration_durationHours
$GLOBALS['Data_Time_Duration_durationHours'] = ["fromDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 * 3600000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "toDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 / 3600000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_durationDays
$GLOBALS['Data_Time_Duration_durationDays'] = ["fromDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 * 86400000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "toDuration" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = ($v_0 / 86400000.0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Time_Duration_convertDuration
$GLOBALS['Data_Time_Duration_convertDuration'] = function($dictDuration_0 = null) {
  $__num = \func_num_args();
  $fromDuration1_1_0 = ($dictDuration_0)['fromDuration'];
  $__res = function($dictDuration1_2 = null) use ($fromDuration1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($dictDuration1_2)['toDuration']))($fromDuration1_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

