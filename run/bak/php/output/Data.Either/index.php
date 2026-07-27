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


final class Data_Either_Left { public $tag = 'Left'; public function __construct(public  $value0) {} }
final class Data_Either_Right { public $tag = 'Right'; public function __construct(public  $value0) {} }

// Data_Either_Left
$GLOBALS['Data_Either_Left'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Either\Data_Either_Left($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Either_Right
$GLOBALS['Data_Either_Right'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Either\Data_Either_Right($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Either_showEither
function majData_majEither_showmajEither($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_showmajEither';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Left "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_2)->{'value0'})))(")"));
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Right "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow1_1)->{'show'})(($v_2)->{'value0'})))(")"));
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
$GLOBALS['Data_Either_showEither'] = __NAMESPACE__ . '\\majData_majEither_showmajEither';

// Data_Either_note'
function majData_majEither_note__prime__($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_note__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_Maybe_maybe__prime__'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Either_Left']))($f_0)))($GLOBALS['Data_Either_Right']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_note__prime__'] = __NAMESPACE__ . '\\majData_majEither_note__prime__';

// Data_Either_note
function majData_majEither_note($a_0, $v2_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_note';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Either\Data_Either_Left($a_0);
goto end_branch_0;;
};
  if ($v2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = new \Data\Either\Data_Either_Right(($v2_1)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Either_note'] = __NAMESPACE__ . '\\majData_majEither_note';

// Data_Either_genericEither
$GLOBALS['Data_Either_genericEither'] = (object)["to" => function($x_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($x_0 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inl) {
$__t0 = new \Data\Either\Data_Either_Left(($x_0)->{'value0'});
goto end_branch_0;;
};
  if ($x_0 instanceof \Data\Generic\Rep\Data_Generic_Rep_Inr) {
$__t0 = new \Data\Either\Data_Either_Right(($x_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "from" => function($x_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($x_0 instanceof \Data\Either\Data_Either_Left) {
$__t1 = new \Data\Generic\Rep\Data_Generic_Rep_Inl(($x_0)->{'value0'});
goto end_branch_1;;
};
  if ($x_0 instanceof \Data\Either\Data_Either_Right) {
$__t1 = new \Data\Generic\Rep\Data_Generic_Rep_Inr(($x_0)->{'value0'});
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
$GLOBALS['Data_Either_functorEither'] = (object)["map" => (function() {
  $__fn = function($f_0, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($m_1 instanceof \Data\Either\Data_Either_Left) {
$__t0 = new \Data\Either\Data_Either_Left(($m_1)->{'value0'});
goto end_branch_0;;
};
  if ($m_1 instanceof \Data\Either\Data_Either_Right) {
$__t0 = new \Data\Either\Data_Either_Right(($f_0)(($m_1)->{'value0'}));
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
$GLOBALS['Data_Either_invariantEither'] = (object)["imap" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Either_functorEither'])->{'map'})($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Either_fromRight'
function majData_majEither_frommajRight__prime__($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_frommajRight__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = ($v_0)($GLOBALS['Data_Unit_unit']);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Either_fromRight__prime__'] = __NAMESPACE__ . '\\majData_majEither_frommajRight__prime__';

// Data_Either_fromRight
function majData_majEither_frommajRight($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_frommajRight';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = $v_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Either_fromRight'] = __NAMESPACE__ . '\\majData_majEither_frommajRight';

// Data_Either_fromLeft'
function majData_majEither_frommajLeft__prime__($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_frommajLeft__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = ($v_0)($GLOBALS['Data_Unit_unit']);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Either_fromLeft__prime__'] = __NAMESPACE__ . '\\majData_majEither_frommajLeft__prime__';

// Data_Either_fromLeft
function majData_majEither_frommajLeft($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_frommajLeft';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($v1_1)->{'value0'};
goto end_branch_0;;
};
  $__t0 = $v_0;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Either_fromLeft'] = __NAMESPACE__ . '\\majData_majEither_frommajLeft';

// Data_Either_extendEither
$GLOBALS['Data_Either_extendEither'] = (object)["extend" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\Either\Data_Either_Left) {
$__t0 = new \Data\Either\Data_Either_Left(($v1_1)->{'value0'});
goto end_branch_0;;
};
  $__t0 = new \Data\Either\Data_Either_Right(($v_0)($v1_1));
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_eqEither
function majData_majEither_eqmajEither($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_eqmajEither';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($x_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($y_3 instanceof \Data\Either\Data_Either_Left && ((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = ($x_2 instanceof \Data\Either\Data_Either_Right && ($y_3 instanceof \Data\Either\Data_Either_Right && ((($dictEq1_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})));
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Either_eqEither'] = __NAMESPACE__ . '\\majData_majEither_eqmajEither';

// Data_Either_ordEither
function majData_majEither_ordmajEither($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_ordmajEither';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $__res = function($dictOrd1_2) use ($__local_var_1_0, $dictOrd_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictOrd1_2)->{'Eq0'})(null);
  $eqEither2_4_2 = (object)["eq" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($x_4, $y_5 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ($x_4 instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($y_5 instanceof \Data\Either\Data_Either_Left && ((($__local_var_1_0)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = ($x_4 instanceof \Data\Either\Data_Either_Right && ($y_5 instanceof \Data\Either\Data_Either_Right && ((($__local_var_3_1)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'})));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($dictOrd1_2, $dictOrd_0) {
  $__fn = function($x_5, $y_6 = null) use ($dictOrd1_2, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t4 = null;;
  if ($x_5 instanceof \Data\Either\Data_Either_Left) {
$__t5 = null;;
if ($y_6 instanceof \Data\Either\Data_Either_Left) {
$__t5 = ((($dictOrd_0)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($y_6 instanceof \Data\Either\Data_Either_Left) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($x_5 instanceof \Data\Either\Data_Either_Right && $y_6 instanceof \Data\Either\Data_Either_Right)) {
$__t4 = ((($dictOrd1_2)->{'compare'})(($x_5)->{'value0'}))(($y_6)->{'value0'});
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
})(), "Eq0" => function($_dollar__unused_5) use ($eqEither2_4_2) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_ordEither'] = __NAMESPACE__ . '\\majData_majEither_ordmajEither';

// Data_Either_eq1Either
function majData_majEither_eq1majEither($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_eq1majEither';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq1" => (function() use ($dictEq_0) {
  $__fn = function($dictEq1_1, $x_2 = null, $y_3 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($x_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($y_3 instanceof \Data\Either\Data_Either_Left && ((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}));
goto end_branch_0;;
};
  $__t0 = ($x_2 instanceof \Data\Either\Data_Either_Right && ($y_3 instanceof \Data\Either\Data_Either_Right && ((($dictEq1_1)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_eq1Either'] = __NAMESPACE__ . '\\majData_majEither_eq1majEither';

// Data_Either_ord1Either
function majData_majEither_ord1majEither($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_ord1majEither';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ordEither1_1_0 = ($GLOBALS['Data_Either_ordEither'])($dictOrd_0);
  $__local_var_2_1 = (($dictOrd_0)->{'Eq0'})(null);
  $eq1Either1_3_2 = (object)["eq1" => (function() use ($__local_var_2_1) {
  $__fn = function($dictEq1_3, $x_4 = null, $y_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t2 = null;;
  if ($x_4 instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($y_5 instanceof \Data\Either\Data_Either_Left && ((($__local_var_2_1)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'}));
goto end_branch_2;;
};
  $__t2 = ($x_4 instanceof \Data\Either\Data_Either_Right && ($y_5 instanceof \Data\Either\Data_Either_Right && ((($dictEq1_3)->{'eq'})(($x_4)->{'value0'}))(($y_5)->{'value0'})));
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare1" => function($dictOrd1_4) use ($ordEither1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordEither1_1_0)($dictOrd1_4))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_4) use ($eq1Either1_3_2) {
  $__num = \func_num_args();
  $__res = $eq1Either1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_ord1Either'] = __NAMESPACE__ . '\\majData_majEither_ord1majEither';

// Data_Either_either
function majData_majEither_either($v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_either';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($v_0)(($v2_2)->{'value0'});
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ($v1_1)(($v2_2)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Either_either'] = __NAMESPACE__ . '\\majData_majEither_either';

// Data_Either_hush
function majData_majEither_hush($v2_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_hush';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v2_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v2_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($v2_0)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_hush'] = __NAMESPACE__ . '\\majData_majEither_hush';

// Data_Either_isLeft
function majData_majEither_ismajLeft($v2_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_ismajLeft';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v2_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = true;
goto end_branch_0;;
};
  if ($v2_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = false;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_isLeft'] = __NAMESPACE__ . '\\majData_majEither_ismajLeft';

// Data_Either_isRight
function majData_majEither_ismajRight($v2_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_ismajRight';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v2_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = false;
goto end_branch_0;;
};
  if ($v2_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = true;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_isRight'] = __NAMESPACE__ . '\\majData_majEither_ismajRight';

// Data_Either_choose
function majData_majEither_choose($dictAlt_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_choose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictAlt_0)->{'Functor0'})(null);
  $__res = (function() use ($__local_var_1_0, $dictAlt_0) {
  $__fn = function($a_2, $b_3 = null) use ($__local_var_1_0, $dictAlt_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictAlt_0)->{'alt'})(((($__local_var_1_0)->{'map'})($GLOBALS['Data_Either_Left']))($a_2)))(((($__local_var_1_0)->{'map'})($GLOBALS['Data_Either_Right']))($b_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_choose'] = __NAMESPACE__ . '\\majData_majEither_choose';

// Data_Either_boundedEither
function majData_majEither_boundedmajEither($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_boundedmajEither';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $bottom_1_0 = ($dictBounded_0)->{'bottom'};
  $ordEither1_2_1 = ($GLOBALS['Data_Either_ordEither'])((($dictBounded_0)->{'Ord0'})(null));
  $__res = function($dictBounded1_3) use ($bottom_1_0, $ordEither1_2_1) {
  $__num = \func_num_args();
  $ordEither2_4_2 = ($ordEither1_2_1)((($dictBounded1_3)->{'Ord0'})(null));
  $__res = (object)["top" => new \Data\Either\Data_Either_Right(($dictBounded1_3)->{'top'}), "bottom" => new \Data\Either\Data_Either_Left($bottom_1_0), "Ord0" => function($_dollar__unused_5) use ($ordEither2_4_2) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_boundedEither'] = __NAMESPACE__ . '\\majData_majEither_boundedmajEither';

// Data_Either_blush
function majData_majEither_blush($v2_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_blush';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v2_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($v2_0)->{'value0'});
goto end_branch_0;;
};
  if ($v2_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_blush'] = __NAMESPACE__ . '\\majData_majEither_blush';

// Data_Either_applyEither
$GLOBALS['Data_Either_applyEither'] = (object)["apply" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = new \Data\Either\Data_Either_Left(($v_0)->{'value0'});
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ((($GLOBALS['Data_Either_functorEither'])->{'map'})(($v_0)->{'value0'}))($v1_1);
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
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_bindEither
$GLOBALS['Data_Either_bindEither'] = (object)["bind" => function($v2_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_0 instanceof \Data\Either\Data_Either_Left) {
$__local_var_1_1 = ($v2_0)->{'value0'};
$__t0 = function($v_2) use ($__local_var_1_1) {
  $__num = \func_num_args();
  $__res = new \Data\Either\Data_Either_Left($__local_var_1_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_0;;
};
  if ($v2_0 instanceof \Data\Either\Data_Either_Right) {
$__local_var_1_2 = ($v2_0)->{'value0'};
$__t0 = function($f_2) use ($__local_var_1_2) {
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
}, "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_applyEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_semigroupEither
function majData_majEither_semigroupmajEither($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majEither_semigroupmajEither';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $append1_1_0 = ($dictSemigroup_0)->{'append'};
  $__res = (object)["append" => (function() use ($append1_1_0) {
  $__fn = function($x_2, $y_3 = null) use ($append1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Either_applyEither'])->{'apply'})(((($GLOBALS['Data_Either_functorEither'])->{'map'})($append1_1_0))($x_2)))($y_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Either_semigroupEither'] = __NAMESPACE__ . '\\majData_majEither_semigroupmajEither';

// Data_Either_applicativeEither
$GLOBALS['Data_Either_applicativeEither'] = (object)["pure" => $GLOBALS['Data_Either_Right'], "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_applyEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_monadEither
$GLOBALS['Data_Either_monadEither'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_applicativeEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_bindEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Either_altEither
$GLOBALS['Data_Either_altEither'] = (object)["alt" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Either\Data_Either_Left) {
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
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

