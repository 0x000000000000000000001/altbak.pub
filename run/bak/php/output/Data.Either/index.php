<?php

namespace Data\Either;

// ALL IMPORTS: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Extend, Control.Monad, Control.Semigroupoid, Data.Bounded, Data.Either, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.Maybe, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Extend, Control.Monad, Control.Semigroupoid, Data.Bounded, Data.Either, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.Maybe, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Invariant/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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


// Data_Either_Left
$GLOBALS['Data_Either_Left'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Left", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Either_Right
$GLOBALS['Data_Either_Right'] = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Right", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Either_showEither
$GLOBALS['Data_Either_showEither'] = (function() {
  $__fn = function($dictShow_0 = null, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ["show" => function($v_2 = null) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Left"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(Left "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow_0)['show'])(($v_2)->{'value0'})))(")"));
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Right"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])['append'])("(Right "))(((($GLOBALS['Data_Semigroup_semigroupString'])['append'])((($dictShow1_1)['show'])(($v_2)->{'value0'})))(")"));
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

// Data_Either_note'
$GLOBALS['Data_Either_note__prime__'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Maybe_maybe__prime__'])(((($GLOBALS['Control_Semigroupoid_semigroupoidFn'])['compose'])($GLOBALS['Data_Either_Left']))($f_0)))($GLOBALS['Data_Either_Right']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Either_note
$GLOBALS['Data_Either_note'] = (function() {
  $__fn = function($a_0 = null, $v2_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v2_1) && (($v2_1)->{'tag'} === "Nothing"))) {
$__t0 = new Phpurs_Data1("Left", $a_0);
goto end_branch_0;;
};
  if ((is_object($v2_1) && (($v2_1)->{'tag'} === "Just"))) {
$__t0 = new Phpurs_Data1("Right", ($v2_1)->{'value0'});
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
})();

// Data_Either_genericEither
$GLOBALS['Data_Either_genericEither'] = ["to" => function($x_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Inl"))) {
$__t0 = new Phpurs_Data1("Left", ($x_0)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Inr"))) {
$__t0 = new Phpurs_Data1("Right", ($x_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "from" => function($x_0 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Left"))) {
$__t1 = new Phpurs_Data1("Inl", ($x_0)->{'value0'});
goto end_branch_1;;
};
  if ((is_object($x_0) && (($x_0)->{'tag'} === "Right"))) {
$__t1 = new Phpurs_Data1("Inr", ($x_0)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_functorEither
$GLOBALS['Data_Either_functorEither'] = ["map" => (function() {
  $__fn = function($f_0 = null, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($m_1) && (($m_1)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data1("Left", ($m_1)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "Right"))) {
$__t0 = new Phpurs_Data1("Right", ($f_0)(($m_1)->{'value0'}));
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
})()];

// Data_Either_invariantEither
$GLOBALS['Data_Either_invariantEither'] = ["imap" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Either_functorEither'])['map'])($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Either_fromRight'
$GLOBALS['Data_Either_fromRight__prime__'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Right"))) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = ($v_0)($GLOBALS['Data_Unit_unit']);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Either_fromRight
$GLOBALS['Data_Either_fromRight'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Right"))) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = $v_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Either_fromLeft'
$GLOBALS['Data_Either_fromLeft__prime__'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Left"))) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = ($v_0)($GLOBALS['Data_Unit_unit']);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Either_fromLeft
$GLOBALS['Data_Either_fromLeft'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Left"))) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = $v_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Either_extendEither
$GLOBALS['Data_Either_extendEither'] = ["extend" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v1_1) && (($v1_1)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data1("Left", ($v1_1)->{'value0'});
goto end_branch_0;;
};
  $__t0 = new Phpurs_Data1("Right", ($v_0)($v1_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_eqEither
$GLOBALS['Data_Either_eqEither'] = (function() {
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
  if ((is_object($x_2) && (($x_2)->{'tag'} === "Left"))) {
$__t0 = ((is_object($y_3) && (($y_3)->{'tag'} === "Left")) && ((($dictEq_0)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_2) && (($x_2)->{'tag'} === "Right")) && ((is_object($y_3) && (($y_3)->{'tag'} === "Right")) && ((($dictEq1_1)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'})));
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

// Data_Either_ordEither
$GLOBALS['Data_Either_ordEither'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictOrd_0)['Eq0'])(null);
  $__res = function($dictOrd1_2 = null) use ($__local_var_1_0, $dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd1_2)['Eq0'])(null);
  $eqEither2_4_2 = ["eq" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($x_4 = null, $y_5 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ((is_object($x_4) && (($x_4)->{'tag'} === "Left"))) {
$__t2 = ((is_object($y_5) && (($y_5)->{'tag'} === "Left")) && ((($__local_var_1_0)['eq'])(($x_4)->{'value0'}))(($y_5)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = ((is_object($x_4) && (($x_4)->{'tag'} === "Right")) && ((is_object($y_5) && (($y_5)->{'tag'} === "Right")) && ((($__local_var_3_1)['eq'])(($x_4)->{'value0'}))(($y_5)->{'value0'})));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["compare" => (function() use ($dictOrd1_2, $dictOrd_0) {
  $__fn = function($x_5 = null, $y_6 = null) use ($dictOrd1_2, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t4 = null;;
  if ((is_object($x_5) && (($x_5)->{'tag'} === "Left"))) {
$__t5 = null;;
if ((is_object($y_6) && (($y_6)->{'tag'} === "Left"))) {
$__t5 = ((($dictOrd_0)['compare'])(($x_5)->{'value0'}))(($y_6)->{'value0'});
goto end_branch_5;;
};
$__t5 = new Phpurs_Data0("LT");
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ((is_object($y_6) && (($y_6)->{'tag'} === "Left"))) {
$__t4 = new Phpurs_Data0("GT");
goto end_branch_4;;
};
  if (((is_object($x_5) && (($x_5)->{'tag'} === "Right")) && (is_object($y_6) && (($y_6)->{'tag'} === "Right")))) {
$__t4 = ((($dictOrd1_2)['compare'])(($x_5)->{'value0'}))(($y_6)->{'value0'});
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_5 = null) use ($eqEither2_4_2) {
  $__num = \func_num_args();
  $__res = $eqEither2_4_2;
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

// Data_Either_eq1Either
$GLOBALS['Data_Either_eq1Either'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ["eq1" => (function() use ($dictEq_0) {
  $__fn = function($dictEq1_1 = null, $x_2 = null, $y_3 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($x_2) && (($x_2)->{'tag'} === "Left"))) {
$__t0 = ((is_object($y_3) && (($y_3)->{'tag'} === "Left")) && ((($dictEq_0)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = ((is_object($x_2) && (($x_2)->{'tag'} === "Right")) && ((is_object($y_3) && (($y_3)->{'tag'} === "Right")) && ((($dictEq1_1)['eq'])(($x_2)->{'value0'}))(($y_3)->{'value0'})));
  end_branch_0:;
  $__res = $__t0;
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

// Data_Either_ord1Either
$GLOBALS['Data_Either_ord1Either'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $ordEither1_1_0 = ($GLOBALS['Data_Either_ordEither'])($dictOrd_0);
  $__local_var_2_1 = (($dictOrd_0)['Eq0'])(null);
  $eq1Either1_3_2 = ["eq1" => (function() use ($__local_var_2_1) {
  $__fn = function($dictEq1_3 = null, $x_4 = null, $y_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t2 = null;;
  if ((is_object($x_4) && (($x_4)->{'tag'} === "Left"))) {
$__t2 = ((is_object($y_5) && (($y_5)->{'tag'} === "Left")) && ((($__local_var_2_1)['eq'])(($x_4)->{'value0'}))(($y_5)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = ((is_object($x_4) && (($x_4)->{'tag'} === "Right")) && ((is_object($y_5) && (($y_5)->{'tag'} === "Right")) && ((($dictEq1_3)['eq'])(($x_4)->{'value0'}))(($y_5)->{'value0'})));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["compare1" => function($dictOrd1_4 = null) use ($ordEither1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordEither1_1_0)($dictOrd1_4))['compare'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_4 = null) use ($eq1Either1_3_2) {
  $__num = \func_num_args();
  $__res = $eq1Either1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Either_either
$GLOBALS['Data_Either_either'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Left"))) {
$__t0 = ($v_0)(($v2_2)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Right"))) {
$__t0 = ($v1_1)(($v2_2)->{'value0'});
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
})();

// Data_Either_hush
$GLOBALS['Data_Either_hush'] = function($v2_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data0("Nothing");
goto end_branch_0;;
};
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Right"))) {
$__t0 = new Phpurs_Data1("Just", ($v2_0)->{'value0'});
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

// Data_Either_isLeft
$GLOBALS['Data_Either_isLeft'] = function($v2_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Left"))) {
$__t0 = true;
goto end_branch_0;;
};
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Right"))) {
$__t0 = false;
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

// Data_Either_isRight
$GLOBALS['Data_Either_isRight'] = function($v2_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Left"))) {
$__t0 = false;
goto end_branch_0;;
};
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Right"))) {
$__t0 = true;
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

// Data_Either_choose
$GLOBALS['Data_Either_choose'] = function($dictAlt_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictAlt_0)['Functor0'])(null);
  $__res = (function() use ($__local_var_1_0, $dictAlt_0) {
  $__fn = function($a_2 = null, $b_3 = null) use ($__local_var_1_0, $dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictAlt_0)['alt'])(((($__local_var_1_0)['map'])($GLOBALS['Data_Either_Left']))($a_2)))(((($__local_var_1_0)['map'])($GLOBALS['Data_Either_Right']))($b_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Either_boundedEither
$GLOBALS['Data_Either_boundedEither'] = function($dictBounded_0 = null) {
  $__num = \func_num_args();
  $bottom_1_0 = ($dictBounded_0)['bottom'];
  $ordEither1_2_1 = ($GLOBALS['Data_Either_ordEither'])((($dictBounded_0)['Ord0'])(null));
  $__res = function($dictBounded1_3 = null) use ($bottom_1_0, $ordEither1_2_1) {
  $__num = \func_num_args();
  $ordEither2_4_2 = ($ordEither1_2_1)((($dictBounded1_3)['Ord0'])(null));
  $__res = ["top" => new Phpurs_Data1("Right", ($dictBounded1_3)['top']), "bottom" => new Phpurs_Data1("Left", $bottom_1_0), "Ord0" => function($_dollar__unused_5 = null) use ($ordEither2_4_2) {
  $__num = \func_num_args();
  $__res = $ordEither2_4_2;
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

// Data_Either_blush
$GLOBALS['Data_Either_blush'] = function($v2_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data1("Just", ($v2_0)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Right"))) {
$__t0 = new Phpurs_Data0("Nothing");
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

// Data_Either_applyEither
$GLOBALS['Data_Either_applyEither'] = ["apply" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data1("Left", ($v_0)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Right"))) {
$__t0 = ((($GLOBALS['Data_Either_functorEither'])['map'])(($v_0)->{'value0'}))($v1_1);
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
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_bindEither
$GLOBALS['Data_Either_bindEither'] = ["bind" => function($v2_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Left"))) {
$__local_var_1_1 = ($v2_0)->{'value0'};
$__t0 = function($v_2 = null) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Left", $__local_var_1_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_0;;
};
  if ((is_object($v2_0) && (($v2_0)->{'tag'} === "Right"))) {
$__local_var_1_2 = ($v2_0)->{'value0'};
$__t0 = function($f_2 = null) use ($__local_var_1_2) {
  $__num = \func_num_args();
  $__res = ($f_2)($__local_var_1_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_applyEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_semigroupEither
$GLOBALS['Data_Either_semigroupEither'] = function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $append1_1_0 = ($dictSemigroup_0)['append'];
  $__res = ["append" => (function() use ($append1_1_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($append1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Either_applyEither'])['apply'])(((($GLOBALS['Data_Either_functorEither'])['map'])($append1_1_0))($x_2)))($y_3);
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

// Data_Either_applicativeEither
$GLOBALS['Data_Either_applicativeEither'] = ["pure" => $GLOBALS['Data_Either_Right'], "Apply0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_applyEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_monadEither
$GLOBALS['Data_Either_monadEither'] = ["Applicative0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_applicativeEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_bindEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_altEither
$GLOBALS['Data_Either_altEither'] = ["alt" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Left"))) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  $__t0 = $v_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

