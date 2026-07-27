<?php

namespace Data\Bifoldable;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Category, Control.Semigroupoid, Data.Bifoldable, Data.Const, Data.Either, Data.Foldable, Data.Function, Data.Functor.Clown, Data.Functor.Flip, Data.Functor.Joker, Data.Functor.Product2, Data.Monoid, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Category, Control.Semigroupoid, Data.Bifoldable, Data.Const, Data.Either, Data.Foldable, Data.Function, Data.Functor.Clown, Data.Functor.Flip, Data.Functor.Joker, Data.Functor.Product2, Data.Monoid, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Newtype, Data.Semigroup, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifoldable/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor.Clown/index.php';
require_once __DIR__ . '/../Data.Functor.Flip/index.php';
require_once __DIR__ . '/../Data.Functor.Joker/index.php';
require_once __DIR__ . '/../Data.Functor.Product2/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Monoid.Conj/index.php';
require_once __DIR__ . '/../Data.Monoid.Disj/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Endo/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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




// Data_Bifoldable_identity
function majData_majBifoldable_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_identity'] = __NAMESPACE__ . '\\majData_majBifoldable_identity';

// Data_Bifoldable_monoidEndo
$GLOBALS['Data_Bifoldable_monoidEndo'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Control_Category_categoryFn'])->{'Semigroupoid0'})(null);
$semigroupEndo1_1_1 = (object)["append" => (function() use ($__local_var_0_0) {
  $__fn = function($v_1, $v1_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_0_0)->{'compose'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return (object)["mempty" => function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar__unused_2) use ($semigroupEndo1_1_1) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_Bifoldable_monoidDual
$GLOBALS['Data_Bifoldable_monoidDual'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Data_Bifoldable_monoidEndo'])->{'Semigroup0'})(null);
$semigroupDual1_1_1 = (object)["append" => (function() use ($__local_var_0_0) {
  $__fn = function($v_1, $v1_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_0_0)->{'append'})($v1_2))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return (object)["mempty" => ($GLOBALS['Data_Bifoldable_monoidEndo'])->{'mempty'}, "Semigroup0" => function($_dollar__unused_2) use ($semigroupDual1_1_1) {
  $__num = \func_num_args();
  $__res = $semigroupDual1_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_Bifoldable_bifoldr
function majData_majBifoldable_bifoldr($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldr';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'bifoldr'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldr'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldr';

// Data_Bifoldable_bitraverse_
function majData_majBifoldable_bitraverse_($dictBifoldable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bitraverse_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictApplicative_1)->{'Apply0'})(null);
  $applySecond_3_1 = (function() use ($__local_var_2_0) {
  $__fn = function($a_3, $b_4 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_0)->{'apply'})(((((($__local_var_2_0)->{'Functor0'})(null))->{'map'})(function($v_5) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Apply_identity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_3)))($b_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (function() use ($applySecond_3_1, $dictApplicative_1, $dictBifoldable_0) {
  $__fn = function($f_4, $g_5 = null) use ($applySecond_3_1, $dictApplicative_1, $dictBifoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictBifoldable_0)->{'bifoldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($applySecond_3_1))($f_4)))((($GLOBALS['Control_Semigroupoid_composeImpl'])($applySecond_3_1))($g_5)))((($dictApplicative_1)->{'pure'})($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_bitraverse_'] = __NAMESPACE__ . '\\majData_majBifoldable_bitraverse_';

// Data_Bifoldable_bifor_
function majData_majBifoldable_bifor_($dictBifoldable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifor_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $bitraverse_2_2_0 = (($GLOBALS['Data_Bifoldable_bitraverse_'])($dictBifoldable_0))($dictApplicative_1);
  $__res = (function() use ($bitraverse_2_2_0) {
  $__fn = function($t_3, $f_4 = null, $g_5 = null) use ($bitraverse_2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($bitraverse_2_2_0)($f_4))($g_5))($t_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifor_'] = __NAMESPACE__ . '\\majData_majBifoldable_bifor_';

// Data_Bifoldable_bisequence_
function majData_majBifoldable_bisequence_($dictBifoldable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bisequence_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_Bifoldable_bitraverse_'])($dictBifoldable_0))($dictApplicative_1))($GLOBALS['Data_Bifoldable_identity']))($GLOBALS['Data_Bifoldable_identity']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_bisequence_'] = __NAMESPACE__ . '\\majData_majBifoldable_bisequence_';

// Data_Bifoldable_bifoldl
function majData_majBifoldable_bifoldl($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'bifoldl'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldl'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldl';

// Data_Bifoldable_bifoldableTuple
$GLOBALS['Data_Bifoldable_bifoldableTuple'] = (object)["bifoldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $g_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})(($f_1)(($v_3)->{'value0'})))(($g_2)(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldr" => (function() {
  $__fn = function($f_0, $g_1 = null, $z_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($f_0)(($v_3)->{'value0'}))((($g_1)(($v_3)->{'value1'}))($z_2));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() {
  $__fn = function($f_0, $g_1 = null, $z_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($g_1)((($f_0)($z_2))(($v_3)->{'value0'})))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()];

// Data_Bifoldable_bifoldableJoker
function majData_majBifoldable_bifoldablemajJoker($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldablemajJoker';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["bifoldr" => (function() use ($dictFoldable_0) {
  $__fn = function($v_1, $r_2 = null, $u_3 = null, $v1_4 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($dictFoldable_0)->{'foldr'})($r_2))($u_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() use ($dictFoldable_0) {
  $__fn = function($v_1, $r_2 = null, $u_3 = null, $v1_4 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($dictFoldable_0)->{'foldl'})($r_2))($u_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $foldMap1_2_0 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_1);
  $__res = (function() use ($foldMap1_2_0) {
  $__fn = function($v_3, $r_4 = null, $v1_5 = null) use ($foldMap1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($foldMap1_2_0)($r_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldableJoker'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldablemajJoker';

// Data_Bifoldable_bifoldableEither
$GLOBALS['Data_Bifoldable_bifoldableEither'] = (object)["bifoldr" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null, $v3_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if ($v3_3 instanceof \Data\Either\Data_Either_Left) {
$__t0 = (($v_0)(($v3_3)->{'value0'}))($v2_2);
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\Either\Data_Either_Right) {
$__t0 = (($v1_1)(($v3_3)->{'value0'}))($v2_2);
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
})(), "bifoldl" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null, $v3_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t1 = null;;
  if ($v3_3 instanceof \Data\Either\Data_Either_Left) {
$__t1 = (($v_0)($v2_2))(($v3_3)->{'value0'});
goto end_branch_1;;
};
  if ($v3_3 instanceof \Data\Either\Data_Either_Right) {
$__t1 = (($v1_1)($v2_2))(($v3_3)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => (function() {
  $__fn = function($dictMonoid_0, $v_1 = null, $v1_2 = null, $v2_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t2 = null;;
  if ($v2_3 instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($v_1)(($v2_3)->{'value0'});
goto end_branch_2;;
};
  if ($v2_3 instanceof \Data\Either\Data_Either_Right) {
$__t2 = ($v1_2)(($v2_3)->{'value0'});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()];

// Data_Bifoldable_bifoldableConst
$GLOBALS['Data_Bifoldable_bifoldableConst'] = (object)["bifoldr" => (function() {
  $__fn = function($f_0, $v_1 = null, $z_2 = null, $v1_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($f_0)($v1_3))($z_2);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() {
  $__fn = function($f_0, $v_1 = null, $z_2 = null, $v1_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($f_0)($z_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null, $v1_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ($f_1)($v1_3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})()];

// Data_Bifoldable_bifoldableClown
function majData_majBifoldable_bifoldablemajClown($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldablemajClown';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["bifoldr" => (function() use ($dictFoldable_0) {
  $__fn = function($l_1, $v_2 = null, $u_3 = null, $v1_4 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($dictFoldable_0)->{'foldr'})($l_1))($u_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() use ($dictFoldable_0) {
  $__fn = function($l_1, $v_2 = null, $u_3 = null, $v1_4 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($dictFoldable_0)->{'foldl'})($l_1))($u_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $foldMap1_2_0 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_1);
  $__res = (function() use ($foldMap1_2_0) {
  $__fn = function($l_3, $v_4 = null, $v1_5 = null) use ($foldMap1_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($foldMap1_2_0)($l_3))($v1_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldableClown'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldablemajClown';

// Data_Bifoldable_bifoldMapDefaultR
function majData_majBifoldable_bifoldmajMapmajDefaultmajR($dictBifoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldmajMapmajDefaultmajR';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $append_2_0 = ((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'};
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = (function() use ($append_2_0, $dictBifoldable_0, $mempty_3_1) {
  $__fn = function($f_4, $g_5 = null) use ($append_2_0, $dictBifoldable_0, $mempty_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictBifoldable_0)->{'bifoldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($append_2_0))($f_4)))((($GLOBALS['Control_Semigroupoid_composeImpl'])($append_2_0))($g_5)))($mempty_3_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldMapDefaultR'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldmajMapmajDefaultmajR';

// Data_Bifoldable_bifoldMapDefaultL
function majData_majBifoldable_bifoldmajMapmajDefaultmajL($dictBifoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldmajMapmajDefaultmajL';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = (function() use ($__local_var_2_0, $dictBifoldable_0, $mempty_3_1) {
  $__fn = function($f_4, $g_5 = null) use ($__local_var_2_0, $dictBifoldable_0, $mempty_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictBifoldable_0)->{'bifoldl'})((function() use ($__local_var_2_0, $f_4) {
  $__fn = function($m_6, $a_7 = null) use ($__local_var_2_0, $f_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_0)->{'append'})($m_6))(($f_4)($a_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((function() use ($__local_var_2_0, $g_5) {
  $__fn = function($m_6, $b_7 = null) use ($__local_var_2_0, $g_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_0)->{'append'})($m_6))(($g_5)($b_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_3_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldMapDefaultL'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldmajMapmajDefaultmajL';

// Data_Bifoldable_bifoldMap
function majData_majBifoldable_bifoldmajMap($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'bifoldMap'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldMap'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldmajMap';

// Data_Bifoldable_bifoldableFlip
function majData_majBifoldable_bifoldablemajFlip($dictBifoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldablemajFlip';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["bifoldr" => (function() use ($dictBifoldable_0) {
  $__fn = function($r_1, $l_2 = null, $u_3 = null, $v_4 = null) use ($dictBifoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictBifoldable_0)->{'bifoldr'})($l_2))($r_1))($u_3))($v_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() use ($dictBifoldable_0) {
  $__fn = function($r_1, $l_2 = null, $u_3 = null, $v_4 = null) use ($dictBifoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictBifoldable_0)->{'bifoldl'})($l_2))($r_1))($u_3))($v_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_1) use ($dictBifoldable_0) {
  $__num = \func_num_args();
  $bifoldMap2_2_0 = (($dictBifoldable_0)->{'bifoldMap'})($dictMonoid_1);
  $__res = (function() use ($bifoldMap2_2_0) {
  $__fn = function($r_3, $l_4 = null, $v_5 = null) use ($bifoldMap2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($bifoldMap2_2_0)($l_4))($r_3))($v_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldableFlip'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldablemajFlip';

// Data_Bifoldable_bifoldlDefault
function majData_majBifoldable_bifoldlmajDefault($dictBifoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldlmajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $bifoldMap1_1_0 = (($dictBifoldable_0)->{'bifoldMap'})($GLOBALS['Data_Bifoldable_monoidDual']);
  $__res = (function() use ($bifoldMap1_1_0) {
  $__fn = function($f_2, $g_3 = null, $z_4 = null, $p_5 = null) use ($bifoldMap1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($bifoldMap1_1_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Dual_Dual']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Endo_Endo']))((function() use ($f_2) {
  $__fn = function($b_6, $a_7 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_2)($a_7))($b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Dual_Dual']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Endo_Endo']))((function() use ($g_3) {
  $__fn = function($b_6, $a_7 = null) use ($g_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($g_3)($a_7))($b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))))($p_5))($z_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldlDefault'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldlmajDefault';

// Data_Bifoldable_bifoldrDefault
function majData_majBifoldable_bifoldrmajDefault($dictBifoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldrmajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $bifoldMap1_1_0 = (($dictBifoldable_0)->{'bifoldMap'})($GLOBALS['Data_Bifoldable_monoidEndo']);
  $__res = (function() use ($bifoldMap1_1_0) {
  $__fn = function($f_2, $g_3 = null, $z_4 = null, $p_5 = null) use ($bifoldMap1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($bifoldMap1_1_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Endo_Endo']))($f_2)))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Endo_Endo']))($g_3)))($p_5))($z_4);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldrDefault'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldrmajDefault';

// Data_Bifoldable_bifoldableProduct2
function majData_majBifoldable_bifoldablemajProduct2($dictBifoldable_0, $dictBifoldable1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifoldablemajProduct2';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_Bifoldable_bifoldableProduct2_dictBifoldable_0 = $dictBifoldable_0;
  $__tco_var_Data_Bifoldable_bifoldableProduct2_dictBifoldable1_1 = $dictBifoldable1_1;
  tco_loop_Data_Bifoldable_bifoldableProduct2:;
  $dictBifoldable_0 = $__tco_var_Data_Bifoldable_bifoldableProduct2_dictBifoldable_0;
  $dictBifoldable1_1 = $__tco_var_Data_Bifoldable_bifoldableProduct2_dictBifoldable1_1;
  $__res = (object)["bifoldr" => (function() use ($dictBifoldable1_1, $dictBifoldable_0) {
  $__fn = function($l_2, $r_3 = null, $u_4 = null, $m_5 = null) use ($dictBifoldable1_1, $dictBifoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = \Data\Bifoldable\majData_majBifoldable_bifoldrmajDefault(\Data\Bifoldable\majData_majBifoldable_bifoldablemajProduct2($dictBifoldable_0, $dictBifoldable1_1), $l_2, $r_3, $u_4, $m_5);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldl" => (function() use ($dictBifoldable1_1, $dictBifoldable_0) {
  $__fn = function($l_2, $r_3 = null, $u_4 = null, $m_5 = null) use ($dictBifoldable1_1, $dictBifoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = \Data\Bifoldable\majData_majBifoldable_bifoldlmajDefault(\Data\Bifoldable\majData_majBifoldable_bifoldablemajProduct2($dictBifoldable_0, $dictBifoldable1_1), $l_2, $r_3, $u_4, $m_5);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(), "bifoldMap" => function($dictMonoid_2) use ($dictBifoldable1_1, $dictBifoldable_0) {
  $__num = \func_num_args();
  $bifoldMap3_3_0 = (($dictBifoldable_0)->{'bifoldMap'})($dictMonoid_2);
  $bifoldMap4_4_1 = (($dictBifoldable1_1)->{'bifoldMap'})($dictMonoid_2);
  $__res = (function() use ($bifoldMap3_3_0, $bifoldMap4_4_1, $dictMonoid_2) {
  $__fn = function($l_5, $r_6 = null, $v_7 = null) use ($bifoldMap3_3_0, $bifoldMap4_4_1, $dictMonoid_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((($dictMonoid_2)->{'Semigroup0'})(null))->{'append'})(((($bifoldMap3_3_0)($l_5))($r_6))(($v_7)->{'value0'})))(((($bifoldMap4_4_1)($l_5))($r_6))(($v_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifoldableProduct2'] = __NAMESPACE__ . '\\majData_majBifoldable_bifoldablemajProduct2';

// Data_Bifoldable_bifold
function majData_majBifoldable_bifold($dictBifoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_bifold';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($dictBifoldable_0)->{'bifoldMap'})($dictMonoid_1))($GLOBALS['Data_Bifoldable_identity']))($GLOBALS['Data_Bifoldable_identity']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_bifold'] = __NAMESPACE__ . '\\majData_majBifoldable_bifold';

// Data_Bifoldable_biany
function majData_majBifoldable_biany($dictBifoldable_0, $dictBooleanAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_biany';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictBooleanAlgebra_1)->{'HeytingAlgebra0'})(null);
  $semigroupDisj1_3_1 = (object)["append" => (function() use ($__local_var_2_0) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_0)->{'disj'})($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $bifoldMap2_2_0 = (($dictBifoldable_0)->{'bifoldMap'})((object)["mempty" => ($__local_var_2_0)->{'ff'}, "Semigroup0" => function($_dollar__unused_4) use ($semigroupDisj1_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  $__res = (function() use ($bifoldMap2_2_0) {
  $__fn = function($p_3, $q_4 = null) use ($bifoldMap2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($bifoldMap2_2_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Disj_Disj']))($p_3)))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Disj_Disj']))($q_4)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_biany'] = __NAMESPACE__ . '\\majData_majBifoldable_biany';

// Data_Bifoldable_biall
function majData_majBifoldable_biall($dictBifoldable_0, $dictBooleanAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majBifoldable_biall';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictBooleanAlgebra_1)->{'HeytingAlgebra0'})(null);
  $semigroupConj1_3_1 = (object)["append" => (function() use ($__local_var_2_0) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_0)->{'conj'})($v_3))($v1_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $bifoldMap2_2_0 = (($dictBifoldable_0)->{'bifoldMap'})((object)["mempty" => ($__local_var_2_0)->{'tt'}, "Semigroup0" => function($_dollar__unused_4) use ($semigroupConj1_3_1) {
  $__num = \func_num_args();
  $__res = $semigroupConj1_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  $__res = (function() use ($bifoldMap2_2_0) {
  $__fn = function($p_3, $q_4 = null) use ($bifoldMap2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($bifoldMap2_2_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Conj_Conj']))($p_3)))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Monoid_Conj_Conj']))($q_4)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Bifoldable_biall'] = __NAMESPACE__ . '\\majData_majBifoldable_biall';

