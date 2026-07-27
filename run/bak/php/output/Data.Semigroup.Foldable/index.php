<?php

namespace Data\Semigroup\Foldable;

// ALL IMPORTS: Control.Apply, Control.Category, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Identity, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord.Max, Data.Ord.Min, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Apply, Control.Category, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Identity, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord.Max, Data.Ord.Min, Data.Ordering, Data.Semigroup, Data.Semigroup.Foldable, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord.Max/index.php';
require_once __DIR__ . '/../Data.Ord.Min/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
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


final class Data_Semigroup_Foldable_FoldRight1 { public $tag = 'FoldRight1'; public function __construct(public  $value0, public  $value1) {} }

// Data_Semigroup_Foldable_identity
function majData_majSemigroup_majFoldable_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_identity'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_identity';

// Data_Semigroup_Foldable_JoinWith
function majData_majSemigroup_majFoldable_majJoinmajWith($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_majJoinmajWith';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_JoinWith'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_majJoinmajWith';

// Data_Semigroup_Foldable_FoldRight1
$GLOBALS['Data_Semigroup_Foldable_FoldRight1'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Semigroup\Foldable\Data_Semigroup_Foldable_FoldRight1($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Semigroup_Foldable_Act
function majData_majSemigroup_majFoldable_majAct($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_majAct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_Act'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_majAct';

// Data_Semigroup_Foldable_semigroupAct
function majData_majSemigroup_majFoldable_semigroupmajAct($dictApply_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_semigroupmajAct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => (function() use ($dictApply_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictApply_0)->{'apply'})(((((($dictApply_0)->{'Functor0'})(null))->{'map'})(function($v_3) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Apply_identity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_1)))($v1_2);
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
$GLOBALS['Data_Semigroup_Foldable_semigroupAct'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_semigroupmajAct';

// Data_Semigroup_Foldable_runFoldRight1
function majData_majSemigroup_majFoldable_runmajFoldmajRight1($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_runmajFoldmajRight1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($v_0)->{'value0'})(($v_0)->{'value1'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_runFoldRight1'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_runmajFoldmajRight1';

// Data_Semigroup_Foldable_mkFoldRight1_closure
$GLOBALS['Data_Semigroup_Foldable_mkFoldRight1_closure'] = ($GLOBALS['Data_Semigroup_Foldable_FoldRight1'])($GLOBALS['Data_Function_const']);

// Data_Semigroup_Foldable_mkFoldRight1
function majData_majSemigroup_majFoldable_mkmajFoldmajRight1($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_mkmajFoldmajRight1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Semigroup_Foldable_mkFoldRight1_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_mkFoldRight1'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_mkmajFoldmajRight1';

// Data_Semigroup_Foldable_foldr1
function majData_majSemigroup_majFoldable_foldr1($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_foldr1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldr1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_foldr1'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_foldr1';

// Data_Semigroup_Foldable_foldl1
function majData_majSemigroup_majFoldable_foldl1($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_foldl1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldl1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_foldl1'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_foldl1';

// Data_Semigroup_Foldable_maximumBy
function majData_majSemigroup_majFoldable_maximummajBy($dictFoldable1_0, $cmp_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_maximummajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFoldable1_0)->{'foldl1'})((function() use ($cmp_1) {
  $__fn = function($x_2, $y_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_1)($x_2))($y_3)))(new \Data\Ordering\Data_Ordering_GT())) {
$__t0 = $x_2;
goto end_branch_0;;
};
  $__t0 = $y_3;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_maximumBy'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_maximummajBy';

// Data_Semigroup_Foldable_minimumBy
function majData_majSemigroup_majFoldable_minimummajBy($dictFoldable1_0, $cmp_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_minimummajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFoldable1_0)->{'foldl1'})((function() use ($cmp_1) {
  $__fn = function($x_2, $y_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_1)($x_2))($y_3)))(new \Data\Ordering\Data_Ordering_LT())) {
$__t0 = $x_2;
goto end_branch_0;;
};
  $__t0 = $y_3;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_minimumBy'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_minimummajBy';

// Data_Semigroup_Foldable_foldableTuple
$GLOBALS['Data_Semigroup_Foldable_foldableTuple'] = (object)["foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v1_1)->{'value1'};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($v1_1)->{'value1'};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Semigroup_Foldable_foldableMultiplicative
$GLOBALS['Data_Semigroup_Foldable_foldableMultiplicative'] = (object)["foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Semigroup_Foldable_foldableIdentity
$GLOBALS['Data_Semigroup_Foldable_foldableIdentity'] = (object)["foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Semigroup_Foldable_foldableDual
$GLOBALS['Data_Semigroup_Foldable_foldableDual'] = (object)["foldr1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl1" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap1" => (function() {
  $__fn = function($dictSemigroup_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_1)($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Semigroup_Foldable_foldRight1Semigroup
$GLOBALS['Data_Semigroup_Foldable_foldRight1Semigroup'] = (object)["append" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($v_0)->{'value1'};
  $__res = new \Data\Semigroup\Foldable\Data_Semigroup_Foldable_FoldRight1((function() use ($__local_var_2_0, $v1_1, $v_0) {
  $__fn = function($a_3, $f_4 = null) use ($__local_var_2_0, $v1_1, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($v_0)->{'value0'})((($f_4)($__local_var_2_0))(((($v1_1)->{'value0'})($a_3))($f_4))))($f_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), ($v1_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Semigroup_Foldable_semigroupDual
$GLOBALS['Data_Semigroup_Foldable_semigroupDual'] = (object)["append" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_Foldable_foldRight1Semigroup'])->{'append'})($v1_1))($v_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Semigroup_Foldable_foldMap1DefaultR
function majData_majSemigroup_majFoldable_foldmajMap1majDefaultmajR($dictFoldable1_0, $dictFunctor_1 = null, $dictSemigroup_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_foldmajMap1majDefaultmajR';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $append_3_0 = ($dictSemigroup_2)->{'append'};
  $__res = function($f_4) use ($append_3_0, $dictFoldable1_0, $dictFunctor_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictFoldable1_0)->{'foldr1'})($append_3_0)))((($dictFunctor_1)->{'map'})($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_foldMap1DefaultR'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_foldmajMap1majDefaultmajR';

// Data_Semigroup_Foldable_foldMap1DefaultL
function majData_majSemigroup_majFoldable_foldmajMap1majDefaultmajL($dictFoldable1_0, $dictFunctor_1 = null, $dictSemigroup_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_foldmajMap1majDefaultmajL';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $append_3_0 = ($dictSemigroup_2)->{'append'};
  $__res = function($f_4) use ($append_3_0, $dictFoldable1_0, $dictFunctor_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictFoldable1_0)->{'foldl1'})($append_3_0)))((($dictFunctor_1)->{'map'})($f_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_foldMap1DefaultL'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_foldmajMap1majDefaultmajL';

// Data_Semigroup_Foldable_foldMap1
function majData_majSemigroup_majFoldable_foldmajMap1($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_foldmajMap1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldMap1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_foldMap1'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_foldmajMap1';

// Data_Semigroup_Foldable_foldl1Default
function majData_majSemigroup_majFoldable_foldl1majDefault($dictFoldable1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_foldl1majDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Semigroup_Foldable_runFoldRight1']))(((($dictFoldable1_0)->{'foldMap1'})($GLOBALS['Data_Semigroup_Foldable_semigroupDual']))($GLOBALS['Data_Semigroup_Foldable_mkFoldRight1']));
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((function() use ($__local_var_1_0) {
  $__fn = function($b_2, $a_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_1_0)($a_3))($b_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_Function_flip']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_foldl1Default'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_foldl1majDefault';

// Data_Semigroup_Foldable_foldr1Default
function majData_majSemigroup_majFoldable_foldr1majDefault($dictFoldable1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_foldr1majDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Semigroup_Foldable_runFoldRight1']))(((($dictFoldable1_0)->{'foldMap1'})($GLOBALS['Data_Semigroup_Foldable_foldRight1Semigroup']))($GLOBALS['Data_Semigroup_Foldable_mkFoldRight1']));
  $__res = (function() use ($__local_var_1_0) {
  $__fn = function($b_2, $a_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_1_0)($a_3))($b_2);
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
$GLOBALS['Data_Semigroup_Foldable_foldr1Default'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_foldr1majDefault';

// Data_Semigroup_Foldable_intercalateMap
function majData_majSemigroup_majFoldable_intercalatemajMap($dictFoldable1_0, $dictSemigroup_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_intercalatemajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $foldMap12_2_0 = (($dictFoldable1_0)->{'foldMap1'})((object)["append" => (function() use ($dictSemigroup_1) {
  $__fn = function($v_2, $v1_3 = null, $j_4 = null) use ($dictSemigroup_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictSemigroup_1)->{'append'})(($v_2)($j_4)))(((($dictSemigroup_1)->{'append'})($j_4))(($v1_3)($j_4)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]);
  $__res = (function() use ($foldMap12_2_0) {
  $__fn = function($j_3, $f_4 = null, $foldable_5 = null) use ($foldMap12_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($foldMap12_2_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Semigroup_Foldable_JoinWith']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))($f_4))))($foldable_5))($j_3);
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
$GLOBALS['Data_Semigroup_Foldable_intercalateMap'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_intercalatemajMap';

// Data_Semigroup_Foldable_intercalate
function majData_majSemigroup_majFoldable_intercalate($dictFoldable1_0, $dictSemigroup_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_intercalate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($GLOBALS['Data_Semigroup_Foldable_intercalateMap'])($dictFoldable1_0))($dictSemigroup_1);
  $__res = function($a_3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_0)($a_3))($GLOBALS['Data_Semigroup_Foldable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_intercalate'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_intercalate';

// Data_Semigroup_Foldable_maximum
function majData_majSemigroup_majFoldable_maximum($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_maximum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $semigroupMax_1_0 = (object)["append" => (function() use ($dictOrd_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_3_0 = ((($dictOrd_0)->{'compare'})($v_1))($v1_2);
  $__t1 = null;;
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $v1_2;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = $v_1;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = $v_1;
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
})()];
  $__res = function($dictFoldable1_2) use ($semigroupMax_1_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable1_2)->{'foldMap1'})($semigroupMax_1_0))($GLOBALS['Unsafe_Coerce_unsafeCoerce']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_maximum'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_maximum';

// Data_Semigroup_Foldable_minimum
function majData_majSemigroup_majFoldable_minimum($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_minimum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $semigroupMin_1_0 = (object)["append" => (function() use ($dictOrd_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_3_0 = ((($dictOrd_0)->{'compare'})($v_1))($v1_2);
  $__t1 = null;;
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $v_1;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = $v_1;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = $v1_2;
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
})()];
  $__res = function($dictFoldable1_2) use ($semigroupMin_1_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable1_2)->{'foldMap1'})($semigroupMin_1_0))($GLOBALS['Unsafe_Coerce_unsafeCoerce']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_minimum'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_minimum';

// Data_Semigroup_Foldable_traverse1_
function majData_majSemigroup_majFoldable_traverse1_($dictFoldable1_0, $dictApply_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_traverse1_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictApply_1)->{'Functor0'})(null);
  $foldMap12_3_1 = (($dictFoldable1_0)->{'foldMap1'})(\Data\Semigroup\Foldable\majData_majSemigroup_majFoldable_semigroupmajAct($dictApply_1));
  $__res = (function() use ($__local_var_2_0, $foldMap12_3_1) {
  $__fn = function($f_4, $t_5 = null) use ($__local_var_2_0, $foldMap12_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_2_0)->{'map'})(function($v_6) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($foldMap12_3_1)((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Semigroup_Foldable_Act']))($f_4)))($t_5));
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
$GLOBALS['Data_Semigroup_Foldable_traverse1_'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_traverse1_';

// Data_Semigroup_Foldable_for1_
function majData_majSemigroup_majFoldable_for1_($dictFoldable1_0, $dictApply_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_for1_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($GLOBALS['Data_Semigroup_Foldable_traverse1_'])($dictFoldable1_0))($dictApply_1);
  $__res = (function() use ($__local_var_2_0) {
  $__fn = function($b_3, $a_4 = null) use ($__local_var_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_2_0)($a_4))($b_3);
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
$GLOBALS['Data_Semigroup_Foldable_for1_'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_for1_';

// Data_Semigroup_Foldable_sequence1_
function majData_majSemigroup_majFoldable_sequence1_($dictFoldable1_0, $dictApply_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_sequence1_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Semigroup_Foldable_traverse1_'])($dictFoldable1_0))($dictApply_1))($GLOBALS['Data_Semigroup_Foldable_identity']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_sequence1_'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_sequence1_';

// Data_Semigroup_Foldable_fold1
function majData_majSemigroup_majFoldable_fold1($dictFoldable1_0, $dictSemigroup_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majSemigroup_majFoldable_fold1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable1_0)->{'foldMap1'})($dictSemigroup_1))($GLOBALS['Data_Semigroup_Foldable_identity']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Semigroup_Foldable_fold1'] = __NAMESPACE__ . '\\majData_majSemigroup_majFoldable_fold1';

