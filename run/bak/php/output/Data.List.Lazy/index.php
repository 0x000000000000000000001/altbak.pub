<?php

namespace Data\List\Lazy;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Lazy, Data.List.Internal, Data.List.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Lazy, Data.List.Internal, Data.List.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Lazy.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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




// Data_List_Lazy_lessThanOrEq
$GLOBALS['Data_List_Lazy_lessThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_List_Lazy_tailRecM2
function majData_majList_majLazy_tailmajRecmajM2($f_0, $a_1 = null, $b_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_tailmajRecmajM2';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_Rec_Class_monadRecMaybe'])->{'tailRecM'})(function($o_3) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($f_0)(($o_3)->{'a'}))(($o_3)->{'b'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_1, "b" => $b_2]);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_tailRecM2'] = __NAMESPACE__ . '\\majData_majList_majLazy_tailmajRecmajM2';

// Data_List_Lazy_lessThan
$GLOBALS['Data_List_Lazy_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_List_Lazy_greaterThan
$GLOBALS['Data_List_Lazy_greaterThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_List_Lazy_greaterThanOrEq
$GLOBALS['Data_List_Lazy_greaterThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_List_Lazy_any
$GLOBALS['Data_List_Lazy_any'] = (function() use (&$__fn) {
$semigroupDisj1_0_0 = (object)["append" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})($v_0))($v1_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
return (($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldMap'})((object)["mempty" => ($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'ff'}, "Semigroup0" => function($_dollar__unused_1) use ($semigroupDisj1_0_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
})();

// Data_List_Lazy_identity
function majData_majList_majLazy_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_identity'] = __NAMESPACE__ . '\\majData_majList_majLazy_identity';

// Data_List_Lazy_Pattern
function majData_majList_majLazy_majPattern($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_majPattern';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_Pattern'] = __NAMESPACE__ . '\\majData_majList_majLazy_majPattern';

// Data_List_Lazy_zipWith
function majData_majList_majLazy_zipmajWith($f_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_zipmajWith';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_zipWith_f_0 = $f_0;
  $__tco_var_Data_List_Lazy_zipWith_xs_1 = $xs_1;
  $__tco_var_Data_List_Lazy_zipWith_ys_2 = $ys_2;
  tco_loop_Data_List_Lazy_zipWith:;
  $f_0 = $__tco_var_Data_List_Lazy_zipWith_f_0;
  $xs_1 = $__tco_var_Data_List_Lazy_zipWith_xs_1;
  $ys_2 = $__tco_var_Data_List_Lazy_zipWith_ys_2;
  $__res = ((($GLOBALS['Data_Lazy_applyLazy'])->{'apply'})(((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})((function() use ($f_0) {
  $__fn = function($v_3, $v1_4 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v1_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if (($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && $v1_4 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons)) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons((($f_0)(($v_3)->{'value0'}))(($v1_4)->{'value0'}), \Data\List\Lazy\majData_majList_majLazy_zipmajWith($f_0, ($v_3)->{'value1'}, ($v1_4)->{'value1'}));
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
})()))($xs_1)))($ys_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_zipWith'] = __NAMESPACE__ . '\\majData_majList_majLazy_zipmajWith';

// Data_List_Lazy_zipWithA
function majData_majList_majLazy_zipmajWithmajA($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_zipmajWithmajA';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequence1_1_0 = (($GLOBALS['Data_List_Lazy_Types_traversableList'])->{'sequence'})($dictApplicative_0);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($sequence1_1_0)(\Data\List\Lazy\majData_majList_majLazy_zipmajWith($f_2, $xs_3, $ys_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_zipWithA'] = __NAMESPACE__ . '\\majData_majList_majLazy_zipmajWithmajA';

// Data_List_Lazy_zip_closure
$GLOBALS['Data_List_Lazy_zip_closure'] = ($GLOBALS['Data_List_Lazy_zipWith'])($GLOBALS['Data_Tuple_Tuple']);

// Data_List_Lazy_zip
function majData_majList_majLazy_zip($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_zip';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_Lazy_zip_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_zip'] = __NAMESPACE__ . '\\majData_majList_majLazy_zip';

// Data_List_Lazy_updateAt
function majData_majList_majLazy_updatemajAt(int $n_0, $x_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_updatemajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_updateAt_n_0 = $n_0;
  $__tco_var_Data_List_Lazy_updateAt_x_1 = $x_1;
  $__tco_var_Data_List_Lazy_updateAt_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_updateAt:;
  $n_0 = $__tco_var_Data_List_Lazy_updateAt_n_0;
  $x_1 = $__tco_var_Data_List_Lazy_updateAt_x_1;
  $xs_2 = $__tco_var_Data_List_Lazy_updateAt_xs_2;
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v1_3) use ($n_0, $x_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = match ($n_0) { 0 => new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_1, ($v1_3)->{'value1'}), default => new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v1_3)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_updatemajAt(($n_0 - 1), $x_1, ($v1_3)->{'value1'})) };
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_updateAt'] = __NAMESPACE__ . '\\majData_majList_majLazy_updatemajAt';

// Data_List_Lazy_unzip_closure
$GLOBALS['Data_List_Lazy_unzip_closure'] = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})(function($v_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($v_0)->{'value0'};
  $__local_var_2_1 = ($v_0)->{'value1'};
  $__res = function($v1_3) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  $__local_var_4_2 = ($v1_3)->{'value0'};
  $__local_var_5_3 = ($v1_3)->{'value1'};
  $__res = new \Data\Tuple\Data_Tuple_Tuple(\Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_1_0, $__local_var_4_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), \Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_2_1, $__local_var_5_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple($GLOBALS['Data_List_Lazy_Types_nil'], $GLOBALS['Data_List_Lazy_Types_nil']));

// Data_List_Lazy_unzip
function majData_majList_majLazy_unzip($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_unzip';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_unzip_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_unzip'] = __NAMESPACE__ . '\\majData_majList_majLazy_unzip';

// Data_List_Lazy_uncons
function majData_majList_majLazy_uncons($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v_1_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_0);
  $__t1 = null;;
  if ($v_1_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = new \Data\Maybe\Data_Maybe_Just((object)["head" => ($v_1_0)->{'value0'}, "tail" => ($v_1_0)->{'value1'}]);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_uncons'] = __NAMESPACE__ . '\\majData_majList_majLazy_uncons';

// Data_List_Lazy_toUnfoldable
function majData_majList_majLazy_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($dictUnfoldable_0)->{'unfoldr'})(function($xs_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($rec_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($rec_2)->{'head'}, ($rec_2)->{'tail'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\Lazy\majData_majList_majLazy_uncons($xs_1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_toUnfoldable'] = __NAMESPACE__ . '\\majData_majList_majLazy_tomajUnfoldable';

// Data_List_Lazy_takeWhile
function majData_majList_majLazy_takemajWhile($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_takemajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_takeWhile_p_0 = $p_0;
  tco_loop_Data_List_Lazy_takeWhile:;
  $p_0 = $__tco_var_Data_List_Lazy_takeWhile_p_0;
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_1) use ($p_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($p_0)(($v_1)->{'value0'}))) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v_1)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_takemajWhile($p_0, ($v_1)->{'value1'}));
goto end_branch_0;;
};
  $__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_takeWhile'] = __NAMESPACE__ . '\\majData_majList_majLazy_takemajWhile';

// Data_List_Lazy_take
function majData_majList_majLazy_take(int $n_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_take';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_take_n_0 = $n_0;
  tco_loop_Data_List_Lazy_take:;
  $n_0 = $__tco_var_Data_List_Lazy_take_n_0;
  $__t1 = null;;
  if ((($GLOBALS['Data_List_Lazy_lessThanOrEq'])($n_0))(0)) {
$__t1 = function($v_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Lazy_Types_nil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
goto end_branch_1;;
};
  $__t1 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v1_1) use ($n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v1_1)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_take(($n_0 - 1), ($v1_1)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_take'] = __NAMESPACE__ . '\\majData_majList_majLazy_take';

// Data_List_Lazy_tail
function majData_majList_majLazy_tail($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = ($v_1)->{'tail'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\Lazy\majData_majList_majLazy_uncons($xs_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_tail'] = __NAMESPACE__ . '\\majData_majList_majLazy_tail';

// Data_List_Lazy_stripPrefix
function majData_majList_majLazy_stripmajPrefix($dictEq_0, $v_1 = null, $s_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_stripmajPrefix';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_tailRecM2'])((function() use ($dictEq_0) {
  $__fn = function($prefix_3, $input_4 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_5_0 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($prefix_3);
  $__t1 = null;;
  if ($v1_5_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done($input_4));
goto end_branch_1;;
};
  if ($v1_5_0 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$v2_6_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($input_4);
$__t3 = null;;
if (($v2_6_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ((($dictEq_0)->{'eq'})(($v1_5_0)->{'value0'}))(($v2_6_2)->{'value0'}))) {
$__t3 = new \Data\Maybe\Data_Maybe_Just(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((object)["a" => ($v1_5_0)->{'value1'}, "b" => ($v2_6_2)->{'value1'}]));
goto end_branch_3;;
};
$__t3 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_3:;
$__t1 = $__t3;
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
})()))($v_1))($s_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_stripPrefix'] = __NAMESPACE__ . '\\majData_majList_majLazy_stripmajPrefix';

// Data_List_Lazy_span
function majData_majList_majLazy_span($p_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_span';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_Lazy_span_p_0 = $p_0;
  $__tco_var_Data_List_Lazy_span_xs_1 = $xs_1;
  tco_loop_Data_List_Lazy_span:;
  $p_0 = $__tco_var_Data_List_Lazy_span_p_0;
  $xs_1 = $__tco_var_Data_List_Lazy_span_xs_1;
  $v_2_0 = \Data\List\Lazy\majData_majList_majLazy_uncons($xs_1);
  $__t1 = null;;
  if (($v_2_0 instanceof \Data\Maybe\Data_Maybe_Just && ($p_0)((($v_2_0)->{'value0'})->{'head'}))) {
$__local_var_3_2 = (($v_2_0)->{'value0'})->{'head'};
$v1_4_3 = \Data\List\Lazy\majData_majList_majLazy_span($p_0, (($v_2_0)->{'value0'})->{'tail'});
$__local_var_5_4 = ($v1_4_3)->{'init'};
$__t1 = (object)["init" => \Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_3_2, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_3_2, $__local_var_5_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "rest" => ($v1_4_3)->{'rest'}];
goto end_branch_1;;
};
  $__t1 = (object)["init" => $GLOBALS['Data_List_Lazy_Types_nil'], "rest" => $xs_1];
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_span'] = __NAMESPACE__ . '\\majData_majList_majLazy_span';

// Data_List_Lazy_snoc
function majData_majList_majLazy_snoc($xs_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_snoc';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})($GLOBALS['Data_List_Lazy_Types_cons']))(\Data\Lazy\majData_majLazy_defer(function($v_2) use ($x_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_1, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($xs_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_snoc'] = __NAMESPACE__ . '\\majData_majList_majLazy_snoc';

// Data_List_Lazy_singleton
function majData_majList_majLazy_singleton($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_1) use ($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_0, $GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_singleton'] = __NAMESPACE__ . '\\majData_majList_majLazy_singleton';

// Data_List_Lazy_showPattern
function majData_majList_majLazy_showmajPattern($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_showmajPattern';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Pattern "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_showmajList($dictShow_0))->{'show'})($v_1)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_showPattern'] = __NAMESPACE__ . '\\majData_majList_majLazy_showmajPattern';

// Data_List_Lazy_scanlLazy
function majData_majList_majLazy_scanlmajLazy($f_0, $acc_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_scanlmajLazy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_scanlLazy_f_0 = $f_0;
  $__tco_var_Data_List_Lazy_scanlLazy_acc_1 = $acc_1;
  $__tco_var_Data_List_Lazy_scanlLazy_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_scanlLazy:;
  $f_0 = $__tco_var_Data_List_Lazy_scanlLazy_f_0;
  $acc_1 = $__tco_var_Data_List_Lazy_scanlLazy_acc_1;
  $xs_2 = $__tco_var_Data_List_Lazy_scanlLazy_xs_2;
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_3) use ($acc_1, $f_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$acc_prime_4_1 = (($f_0)($acc_1))(($v_3)->{'value0'});
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($acc_prime_4_1, \Data\List\Lazy\majData_majList_majLazy_scanlmajLazy($f_0, $acc_prime_4_1, ($v_3)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_scanlLazy'] = __NAMESPACE__ . '\\majData_majList_majLazy_scanlmajLazy';

// Data_List_Lazy_reverse
function majData_majList_majLazy_reverse($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_reverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v_1) use ($xs_0) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($b_2, $a_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_4) use ($a_3, $b_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_3, $b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_List_Lazy_Types_nil']))($xs_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_reverse'] = __NAMESPACE__ . '\\majData_majList_majLazy_reverse';

// Data_List_Lazy_replicateM
function majData_majList_majLazy_replicatemajM($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_replicatemajM';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_replicateM_dictMonad_0 = $dictMonad_0;
  tco_loop_Data_List_Lazy_replicateM:;
  $dictMonad_0 = $__tco_var_Data_List_Lazy_replicateM_dictMonad_0;
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($n_3, $m_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ((($GLOBALS['Data_List_Lazy_lessThan'])($n_3))(1)) {
$__t2 = (($__local_var_1_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_2;;
};
  $__t2 = ((($__local_var_2_1)->{'bind'})($m_4))(function($a_5) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, $m_4, $n_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'bind'})(\Data\List\Lazy\majData_majList_majLazy_replicatemajM($dictMonad_0, ($n_3 - 1), $m_4)))(function($as_6) use ($__local_var_1_0, $a_5) {
  $__num = \func_num_args();
  $__res = (($__local_var_1_0)->{'pure'})(\Data\Lazy\majData_majLazy_defer(function($v_7) use ($a_5, $as_6) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($a_5, $as_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_replicateM'] = __NAMESPACE__ . '\\majData_majList_majLazy_replicatemajM';

// Data_List_Lazy_repeat
function majData_majList_majLazy_repeat($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_repeat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v_2) use (&$go__1_0, $x_0) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_3) use (&$go__1_0, $x_0) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_0, $go__1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = $go__1_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_repeat'] = __NAMESPACE__ . '\\majData_majList_majLazy_repeat';

// Data_List_Lazy_replicate
function majData_majList_majLazy_replicate(int $i_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_replicate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v_3) use (&$go__2_0, $xs_1) {
  $__num = \func_num_args();
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_4) use (&$go__2_0, $xs_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($xs_1, $go__2_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = \Data\List\Lazy\majData_majList_majLazy_take($i_0, $go__2_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_replicate'] = __NAMESPACE__ . '\\majData_majList_majLazy_replicate';

// Data_List_Lazy_range
function majData_majList_majLazy_range(int $start_0, $end_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_range';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((($GLOBALS['Data_List_Lazy_greaterThan'])($start_0))($end_1)) {
$__t1 = ((($GLOBALS['Data_List_Lazy_Types_unfoldableList'])->{'unfoldr'})(function($x_2) use ($end_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((($GLOBALS['Data_List_Lazy_greaterThanOrEq'])($x_2))($end_1)) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple($x_2, ($x_2 - 1)));
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($start_0);
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_List_Lazy_Types_unfoldableList'])->{'unfoldr'})(function($x_2) use ($end_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((($GLOBALS['Data_List_Lazy_lessThanOrEq'])($x_2))($end_1)) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple($x_2, ($x_2 + 1)));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($start_0);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_range'] = __NAMESPACE__ . '\\majData_majList_majLazy_range';

// Data_List_Lazy_partition
function majData_majList_majLazy_partition($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_partition';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldr'})((function() use ($f_0) {
  $__fn = function($x_1, $v_2 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($f_0)($x_1)) {
$__t0 = (object)["yes" => \Data\Lazy\majData_majLazy_defer(function($v_3) use ($v_2, $x_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_1, ($v_2)->{'yes'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "no" => ($v_2)->{'no'}];
goto end_branch_0;;
};
  $__t0 = (object)["yes" => ($v_2)->{'yes'}, "no" => \Data\Lazy\majData_majLazy_defer(function($v_3) use ($v_2, $x_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_1, ($v_2)->{'no'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["yes" => $GLOBALS['Data_List_Lazy_Types_nil'], "no" => $GLOBALS['Data_List_Lazy_Types_nil']]);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_partition'] = __NAMESPACE__ . '\\majData_majList_majLazy_partition';

// Data_List_Lazy_null_closure
$GLOBALS['Data_List_Lazy_null_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_isNothing']))($GLOBALS['Data_List_Lazy_uncons']);

// Data_List_Lazy_null
function majData_majList_majLazy_null($v_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_null';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_null_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_null'] = __NAMESPACE__ . '\\majData_majList_majLazy_null';

// Data_List_Lazy_nubBy
function majData_majList_majLazy_nubmajBy($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_nubmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $goStep_1_0 = null;
  $go__1_0 = null;
  $goStep_1_0 = (function() use (&$go__1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$v2_4_2 = \Data\List\Internal\majData_majList_majInternal_insertmajAndmajLookupmajBy($p_0, ($v1_3)->{'value0'}, $v_2);
$__t3 = null;;
if (($v2_4_2)->{'found'}) {
$__t3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step((($go__1_0)(($v2_4_2)->{'result'}))(($v1_3)->{'value1'}));
goto end_branch_3;;
};
$__t3 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v1_3)->{'value0'}, (($go__1_0)(($v2_4_2)->{'result'}))(($v1_3)->{'value1'}));
end_branch_3:;
$__t1 = $__t3;
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
  $go__1_0 = (function() use (&$goStep_1_0) {
  $__fn = function($s_2, $v_3 = null) use (&$goStep_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(($goStep_1_0)($s_2)))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go__1_0)(new \Data\List\Internal\Data_List_Internal_Leaf());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_nubBy'] = __NAMESPACE__ . '\\majData_majList_majLazy_nubmajBy';

// Data_List_Lazy_nub
function majData_majList_majLazy_nub($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_nub';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_nubBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_nub'] = __NAMESPACE__ . '\\majData_majList_majLazy_nub';

// Data_List_Lazy_newtypePattern
$GLOBALS['Data_List_Lazy_newtypePattern'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Lazy_mapMaybe
function majData_majList_majLazy_mapmajMaybe($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_mapmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_mapMaybe_f_0 = $f_0;
  tco_loop_Data_List_Lazy_mapMaybe:;
  $f_0 = $__tco_var_Data_List_Lazy_mapMaybe_f_0;
  $go__1_0 = null;
  $go__1_0 = function($v_2) use (&$__tco_var_Data_List_Lazy_mapMaybe_f_0, $f_0, &$go__1_0) {
  $__num = \func_num_args();
  $__tco_var_go__1_0_0_v_2 = $v_2;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $__t0 = null;;
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$v1_3_1 = ($f_0)(($v_2)->{'value0'});
$__t2 = null;;
if ($v1_3_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__tco_3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_2)->{'value1'});
$__tco_var_go__1_0_0_v_2 = $__tco_3;
goto tco_loop_go__1_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if ($v1_3_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v1_3_1)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_mapmajMaybe($f_0, ($v_2)->{'value1'}));
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})($go__1_0)))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_mapMaybe'] = __NAMESPACE__ . '\\majData_majList_majLazy_mapmajMaybe';

// Data_List_Lazy_some
function majData_majList_majLazy_some($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_some';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)->{'Applicative0'})(null))->{'Apply0'})(null))->{'apply'})(((((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_List_Lazy_Types_cons']))($v_2)))((($dictLazy_1)->{'defer'})(function($v1_3) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  $__res = \Data\List\Lazy\majData_majList_majLazy_many($dictAlternative_0, $dictLazy_1, $v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_some'] = __NAMESPACE__ . '\\majData_majList_majLazy_some';

// Data_List_Lazy_many
function majData_majList_majLazy_many($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_many';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'alt'})(\Data\List\Lazy\majData_majList_majLazy_some($dictAlternative_0, $dictLazy_1, $v_2)))((((($dictAlternative_0)->{'Applicative0'})(null))->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_many'] = __NAMESPACE__ . '\\majData_majList_majLazy_many';

// Data_List_Lazy_length_closure
$GLOBALS['Data_List_Lazy_length_closure'] = ((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($l_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($l_0 + 1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(0);

// Data_List_Lazy_length
function majData_majList_majLazy_length($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_length_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_length'] = __NAMESPACE__ . '\\majData_majList_majLazy_length';

// Data_List_Lazy_last_closure
$GLOBALS['Data_List_Lazy_last_closure'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = function($v_1) use (&$go__0_0) {
  $__num = \func_num_args();
  $__tco_var_go__0_0_0_v_1 = $v_1;
  tco_loop_go__0_0_0:;
  $v_1 = $__tco_var_go__0_0_0_v_1;
  $__t0 = null;;
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t2 = null;;
if (\Data\List\Lazy\majData_majList_majLazy_null(($v_1)->{'value1'})) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($v_1)->{'value0'});
goto end_branch_2;;
};
$__tco_1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_1)->{'value1'});
$__tco_var_go__0_0_0_v_1 = $__tco_1;
goto tco_loop_go__0_0_0;;
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return (($GLOBALS['Control_Semigroupoid_composeImpl'])($go__0_0))($GLOBALS['Data_List_Lazy_Types_step']);
})();

// Data_List_Lazy_last
function majData_majList_majLazy_last($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_last';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_last_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_last'] = __NAMESPACE__ . '\\majData_majList_majLazy_last';

// Data_List_Lazy_iterate
function majData_majList_majLazy_iterate($f_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_iterate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v_3) use ($f_0, &$go__2_0, $x_1) {
  $__num = \func_num_args();
  $__local_var_4_1 = ((($GLOBALS['Data_List_Lazy_Types_functorList'])->{'map'})($f_0))($go__2_0);
  $__res = \Data\Lazy\majData_majLazy_defer(function($v_5) use ($__local_var_4_1, $x_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_1, $__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = $go__2_0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_iterate'] = __NAMESPACE__ . '\\majData_majList_majLazy_iterate';

// Data_List_Lazy_insertAt
function majData_majList_majLazy_insertmajAt(int $v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_insertmajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_insertAt_v_0 = $v_0;
  $__tco_var_Data_List_Lazy_insertAt_v1_1 = $v1_1;
  $__tco_var_Data_List_Lazy_insertAt_v2_2 = $v2_2;
  tco_loop_Data_List_Lazy_insertAt:;
  $v_0 = $__tco_var_Data_List_Lazy_insertAt_v_0;
  $v1_1 = $__tco_var_Data_List_Lazy_insertAt_v1_1;
  $v2_2 = $__tco_var_Data_List_Lazy_insertAt_v2_2;
  $__res = match ($v_0) { 0 => \Data\Lazy\majData_majLazy_defer(function($v_3) use ($v1_1, $v2_2) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($v1_1, $v2_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), default => ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v3_3) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v3_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($v1_1, $GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_0;;
};
  if ($v3_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v3_3)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_insertmajAt(($v_0 - 1), $v1_1, ($v3_3)->{'value1'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v2_2) };
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_insertAt'] = __NAMESPACE__ . '\\majData_majList_majLazy_insertmajAt';

// Data_List_Lazy_init_closure
$GLOBALS['Data_List_Lazy_init_closure'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = function($v_1) use (&$go__0_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t2 = null;;
if (\Data\List\Lazy\majData_majList_majLazy_null(($v_1)->{'value1'})) {
$__t2 = new \Data\Maybe\Data_Maybe_Just($GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_2;;
};
$__t2 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(($GLOBALS['Data_List_Lazy_Types_cons'])(($v_1)->{'value0'})))(($go__0_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_1)->{'value1'})));
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return (($GLOBALS['Control_Semigroupoid_composeImpl'])($go__0_0))($GLOBALS['Data_List_Lazy_Types_step']);
})();

// Data_List_Lazy_init
function majData_majList_majLazy_init($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_init';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_init_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_init'] = __NAMESPACE__ . '\\majData_majList_majLazy_init';

// Data_List_Lazy_index
function majData_majList_majLazy_index($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_index';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $__t0 = null;;
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t3 = null;;
switch ($v1_3) {
case 0:
$__t3 = new \Data\Maybe\Data_Maybe_Just(($v_2)->{'value0'});
goto end_branch_3;;
break;
default:
;
break;
};
$__tco_1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_2)->{'value1'});
$__tco_2 = ($v1_3 - 1);
$__tco_var_go__1_0_0_v_2 = $__tco_1;
$__tco_var_go__1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__1_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
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
  $__res = ($go__1_0)(\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_index'] = __NAMESPACE__ . '\\majData_majList_majLazy_index';

// Data_List_Lazy_head
function majData_majList_majLazy_head($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = ($v_1)->{'head'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\Lazy\majData_majList_majLazy_uncons($xs_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_head'] = __NAMESPACE__ . '\\majData_majList_majLazy_head';

// Data_List_Lazy_transpose
function majData_majList_majLazy_transpose($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_transpose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_transpose_xs_0 = $xs_0;
  tco_loop_Data_List_Lazy_transpose:;
  $xs_0 = $__tco_var_Data_List_Lazy_transpose_xs_0;
  $v_1_0 = \Data\List\Lazy\majData_majList_majLazy_uncons($xs_0);
  $__t1 = null;;
  if ($v_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $xs_0;
goto end_branch_1;;
};
  if ($v_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$v1_2_2 = \Data\List\Lazy\majData_majList_majLazy_uncons((($v_1_0)->{'value0'})->{'head'});
$__t3 = null;;
if ($v1_2_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__tco_4 = (($v_1_0)->{'value0'})->{'tail'};
$__tco_var_Data_List_Lazy_transpose_xs_0 = $__tco_4;
goto tco_loop_Data_List_Lazy_transpose;;
$__t3 = null;
goto end_branch_3;;
};
if ($v1_2_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_3_5 = (($v1_2_2)->{'value0'})->{'head'};
$__local_var_4_6 = (($v1_2_2)->{'value0'})->{'tail'};
$__local_var_5_7 = \Data\List\Lazy\majData_majList_majLazy_mapmajMaybe($GLOBALS['Data_List_Lazy_head'], (($v_1_0)->{'value0'})->{'tail'});
$__local_var_6_8 = \Data\Lazy\majData_majLazy_defer(function($v_6) use ($__local_var_3_5, $__local_var_5_7) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_3_5, $__local_var_5_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_7_9 = \Data\List\Lazy\majData_majList_majLazy_mapmajMaybe($GLOBALS['Data_List_Lazy_tail'], (($v_1_0)->{'value0'})->{'tail'});
$__local_var_7_9 = \Data\List\Lazy\majData_majList_majLazy_transpose(\Data\Lazy\majData_majLazy_defer(function($v_8) use ($__local_var_4_6, $__local_var_7_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_4_6, $__local_var_7_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__t3 = \Data\Lazy\majData_majLazy_defer(function($v_8) use ($__local_var_6_8, $__local_var_7_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_6_8, $__local_var_7_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_transpose'] = __NAMESPACE__ . '\\majData_majList_majLazy_transpose';

// Data_List_Lazy_groupBy
function majData_majList_majLazy_groupmajBy($eq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_groupmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_groupBy_eq_0 = $eq_0;
  tco_loop_Data_List_Lazy_groupBy:;
  $eq_0 = $__tco_var_Data_List_Lazy_groupBy_eq_0;
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_1) use ($eq_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_2_1 = ($v_1)->{'value0'};
$v1_3_2 = \Data\List\Lazy\majData_majList_majLazy_span(($eq_0)($__local_var_2_1), ($v_1)->{'value1'});
$__local_var_4_3 = ($v1_3_2)->{'init'};
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(\Data\Lazy\majData_majLazy_defer(function($v2_5) use ($__local_var_2_1, $__local_var_4_3) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_2_1, $__local_var_4_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), \Data\List\Lazy\majData_majList_majLazy_groupmajBy($eq_0, ($v1_3_2)->{'rest'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_groupBy'] = __NAMESPACE__ . '\\majData_majList_majLazy_groupmajBy';

// Data_List_Lazy_group
function majData_majList_majLazy_group($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_group';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_groupBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_group'] = __NAMESPACE__ . '\\majData_majList_majLazy_group';

// Data_List_Lazy_fromStep_closure
$GLOBALS['Data_List_Lazy_fromStep_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))(($GLOBALS['Data_Lazy_applicativeLazy'])->{'pure'});

// Data_List_Lazy_fromStep
function majData_majList_majLazy_frommajStep($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_frommajStep';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_fromStep_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_fromStep'] = __NAMESPACE__ . '\\majData_majList_majLazy_frommajStep';

// Data_List_Lazy_insertBy
function majData_majList_majLazy_insertmajBy($cmp_0, $x_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_insertmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_insertBy_cmp_0 = $cmp_0;
  $__tco_var_Data_List_Lazy_insertBy_x_1 = $x_1;
  $__tco_var_Data_List_Lazy_insertBy_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_insertBy:;
  $cmp_0 = $__tco_var_Data_List_Lazy_insertBy_cmp_0;
  $x_1 = $__tco_var_Data_List_Lazy_insertBy_x_1;
  $xs_2 = $__tco_var_Data_List_Lazy_insertBy_xs_2;
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_3) use ($cmp_0, $x_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_1, $GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = null;;
if ((($cmp_0)($x_1))(($v_3)->{'value0'}) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v_3)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_insertmajBy($cmp_0, $x_1, ($v_3)->{'value1'}));
goto end_branch_1;;
};
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($x_1, \Data\List\Lazy\majData_majList_majLazy_frommajStep($v_3));
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_insertBy'] = __NAMESPACE__ . '\\majData_majList_majLazy_insertmajBy';

// Data_List_Lazy_insert
function majData_majList_majLazy_insert($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_insert';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_insertBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_insert'] = __NAMESPACE__ . '\\majData_majList_majLazy_insert';

// Data_List_Lazy_fromFoldable
function majData_majList_majLazy_frommajFoldable($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_frommajFoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($dictFoldable_0)->{'foldr'})($GLOBALS['Data_List_Lazy_Types_cons']))($GLOBALS['Data_List_Lazy_Types_nil']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_fromFoldable'] = __NAMESPACE__ . '\\majData_majList_majLazy_frommajFoldable';

// Data_List_Lazy_foldrLazy
function majData_majList_majLazy_foldrmajLazy($dictLazy_0, $op_1 = null, $z_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_foldrmajLazy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = function($xs_4) use ($dictLazy_0, &$go__3_0, $op_1, $z_2) {
  $__num = \func_num_args();
  $v_5_1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step($xs_4);
  $__t2 = null;;
  if ($v_5_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_6_3 = ($v_5_1)->{'value0'};
$__local_var_7_4 = ($v_5_1)->{'value1'};
$__t2 = (($dictLazy_0)->{'defer'})(function($v1_8) use ($__local_var_6_3, $__local_var_7_4, &$go__3_0, $op_1) {
  $__num = \func_num_args();
  $__res = (($op_1)($__local_var_6_3))(($go__3_0)($__local_var_7_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_2;;
};
  if ($v_5_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t2 = $z_2;
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
  $__res = $go__3_0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_foldrLazy'] = __NAMESPACE__ . '\\majData_majList_majLazy_foldrmajLazy';

// Data_List_Lazy_foldM
function majData_majList_majLazy_foldmajM($dictMonad_0, $f_1 = null, $b_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_foldmajM';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_List_Lazy_foldM_dictMonad_0 = $dictMonad_0;
  $__tco_var_Data_List_Lazy_foldM_f_1 = $f_1;
  $__tco_var_Data_List_Lazy_foldM_b_2 = $b_2;
  $__tco_var_Data_List_Lazy_foldM_xs_3 = $xs_3;
  tco_loop_Data_List_Lazy_foldM:;
  $dictMonad_0 = $__tco_var_Data_List_Lazy_foldM_dictMonad_0;
  $f_1 = $__tco_var_Data_List_Lazy_foldM_f_1;
  $b_2 = $__tco_var_Data_List_Lazy_foldM_b_2;
  $xs_3 = $__tco_var_Data_List_Lazy_foldM_xs_3;
  $v_4_0 = \Data\List\Lazy\majData_majList_majLazy_uncons($xs_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})($b_2);
goto end_branch_1;;
};
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_5_2 = (($v_4_0)->{'value0'})->{'tail'};
$__t1 = ((((($dictMonad_0)->{'Bind1'})(null))->{'bind'})((($f_1)($b_2))((($v_4_0)->{'value0'})->{'head'})))(function($b_prime_6) use ($__local_var_5_2, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = \Data\List\Lazy\majData_majList_majLazy_foldmajM($dictMonad_0, $f_1, $b_prime_6, $__local_var_5_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_List_Lazy_foldM'] = __NAMESPACE__ . '\\majData_majList_majLazy_foldmajM';

// Data_List_Lazy_findIndex
function majData_majList_majLazy_findmajIndex($fn_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_findmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use ($fn_0, &$go__1_0) {
  $__fn = function($n_2, $list_3 = null) use ($fn_0, &$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_bindMaybe'])->{'bind'})(\Data\List\Lazy\majData_majList_majLazy_uncons($list_3)))(function($o_4) use ($fn_0, &$go__1_0, $n_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($fn_0)(($o_4)->{'head'})) {
$__t1 = (($GLOBALS['Data_Maybe_applicativeMaybe'])->{'pure'})($n_2);
goto end_branch_1;;
};
  $__t1 = (($go__1_0)(($n_2 + 1)))(($o_4)->{'tail'});
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go__1_0)(0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_findIndex'] = __NAMESPACE__ . '\\majData_majList_majLazy_findmajIndex';

// Data_List_Lazy_findLastIndex
function majData_majList_majLazy_findmajLastmajIndex($fn_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_findmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_2) use ($xs_1) {
  $__num = \func_num_args();
  $__res = ((\Data\List\Lazy\majData_majList_majLazy_length($xs_1) - 1) - $v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\Lazy\majData_majList_majLazy_findmajIndex($fn_0, \Data\List\Lazy\majData_majList_majLazy_reverse($xs_1)));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_findLastIndex'] = __NAMESPACE__ . '\\majData_majList_majLazy_findmajLastmajIndex';

// Data_List_Lazy_filterM
function majData_majList_majLazy_filtermajM($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_filtermajM';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_filterM_dictMonad_0 = $dictMonad_0;
  tco_loop_Data_List_Lazy_filterM:;
  $dictMonad_0 = $__tco_var_Data_List_Lazy_filterM_dictMonad_0;
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($p_3, $list_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_5_2 = \Data\List\Lazy\majData_majList_majLazy_uncons($list_4);
  $__t3 = null;;
  if ($v_5_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = (($__local_var_1_0)->{'pure'})($GLOBALS['Data_List_Lazy_Types_nil']);
goto end_branch_3;;
};
  if ($v_5_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_6_4 = (($v_5_2)->{'value0'})->{'head'};
$__local_var_7_5 = (($v_5_2)->{'value0'})->{'tail'};
$__t3 = ((($__local_var_2_1)->{'bind'})(($p_3)($__local_var_6_4)))(function($b_8) use ($__local_var_1_0, $__local_var_2_1, $__local_var_6_4, $__local_var_7_5, $dictMonad_0, $p_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'bind'})(\Data\List\Lazy\majData_majList_majLazy_filtermajM($dictMonad_0, $p_3, $__local_var_7_5)))(function($xs_prime_9) use ($__local_var_1_0, $__local_var_6_4, $b_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($b_8) {
$__t6 = \Data\Lazy\majData_majLazy_defer(function($v_10) use ($__local_var_6_4, $xs_prime_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_6_4, $xs_prime_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_6;;
};
  $__t6 = $xs_prime_9;
  end_branch_6:;
  $__res = (($__local_var_1_0)->{'pure'})($__t6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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
$GLOBALS['Data_List_Lazy_filterM'] = __NAMESPACE__ . '\\majData_majList_majLazy_filtermajM';

// Data_List_Lazy_filter
function majData_majList_majLazy_filter($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_filter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_filter_p_0 = $p_0;
  tco_loop_Data_List_Lazy_filter:;
  $p_0 = $__tco_var_Data_List_Lazy_filter_p_0;
  $go__1_0 = null;
  $go__1_0 = function($v_2) use (&$__tco_var_Data_List_Lazy_filter_p_0, &$go__1_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go__1_0_0_v_2 = $v_2;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $__t0 = null;;
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t2 = null;;
if (($p_0)(($v_2)->{'value0'})) {
$__t2 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v_2)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_filter($p_0, ($v_2)->{'value1'}));
goto end_branch_2;;
};
$__tco_1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_2)->{'value1'});
$__tco_var_go__1_0_0_v_2 = $__tco_1;
goto tco_loop_go__1_0_0;;
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})($go__1_0)))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_filter'] = __NAMESPACE__ . '\\majData_majList_majLazy_filter';

// Data_List_Lazy_intersectBy
function majData_majList_majLazy_intersectmajBy($eq_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_intersectmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Data\List\Lazy\majData_majList_majLazy_filter(function($x_3) use ($eq_0, $ys_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_List_Lazy_any'])(($eq_0)($x_3)))($ys_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $xs_1);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_intersectBy'] = __NAMESPACE__ . '\\majData_majList_majLazy_intersectmajBy';

// Data_List_Lazy_intersect
function majData_majList_majLazy_intersect($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_intersect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_intersectBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_intersect'] = __NAMESPACE__ . '\\majData_majList_majLazy_intersect';

// Data_List_Lazy_nubByEq
function majData_majList_majLazy_nubmajBymajEq($eq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_nubmajBymajEq';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_Lazy_nubByEq_eq_0 = $eq_0;
  tco_loop_Data_List_Lazy_nubByEq:;
  $eq_0 = $__tco_var_Data_List_Lazy_nubByEq_eq_0;
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_1) use ($eq_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_1 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__local_var_2_1 = ($v_1)->{'value0'};
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons($__local_var_2_1, \Data\List\Lazy\majData_majList_majLazy_nubmajBymajEq($eq_0, \Data\List\Lazy\majData_majList_majLazy_filter(function($y_3) use ($__local_var_2_1, $eq_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'})((($eq_0)($__local_var_2_1))($y_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_1)->{'value1'})));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_nubByEq'] = __NAMESPACE__ . '\\majData_majList_majLazy_nubmajBymajEq';

// Data_List_Lazy_nubEq
function majData_majList_majLazy_nubmajEq($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_nubmajEq';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_nubByEq'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_nubEq'] = __NAMESPACE__ . '\\majData_majList_majLazy_nubmajEq';

// Data_List_Lazy_eqPattern
function majData_majList_majLazy_eqmajPattern($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_eqmajPattern';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_eq1List'])->{'eq1'})($dictEq_0))($x_1))($y_2);
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
$GLOBALS['Data_List_Lazy_eqPattern'] = __NAMESPACE__ . '\\majData_majList_majLazy_eqmajPattern';

// Data_List_Lazy_ordPattern
function majData_majList_majLazy_ordmajPattern($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_ordmajPattern';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqPattern1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Lazy_Types_eq1List'])->{'eq1'})($__local_var_1_0))($x_2))($y_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($x_3, $y_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((\Data\List\Lazy\Types\majData_majList_majLazy_majTypes_ordmajList($dictOrd_0))->{'compare'})($x_3))($y_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_3) use ($eqPattern1_2_1) {
  $__num = \func_num_args();
  $__res = $eqPattern1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_ordPattern'] = __NAMESPACE__ . '\\majData_majList_majLazy_ordmajPattern';

// Data_List_Lazy_elemLastIndex
function majData_majList_majLazy_elemmajLastmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_elemmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_Lazy_findLastIndex'])(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_elemLastIndex'] = __NAMESPACE__ . '\\majData_majList_majLazy_elemmajLastmajIndex';

// Data_List_Lazy_elemIndex
function majData_majList_majLazy_elemmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_elemmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_Lazy_findIndex'])(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_elemIndex'] = __NAMESPACE__ . '\\majData_majList_majLazy_elemmajIndex';

// Data_List_Lazy_dropWhile
function majData_majList_majLazy_dropmajWhile($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_dropmajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = function($v_2) use (&$go__1_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go__1_0_0_v_2 = $v_2;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $__t0 = null;;
  if (($v_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons && ($p_0)(($v_2)->{'value0'}))) {
$__tco_1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_2)->{'value1'});
$__tco_var_go__1_0_0_v_2 = $__tco_1;
goto tco_loop_go__1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  $__t0 = \Data\List\Lazy\majData_majList_majLazy_frommajStep($v_2);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($go__1_0))($GLOBALS['Data_List_Lazy_Types_step']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_dropWhile'] = __NAMESPACE__ . '\\majData_majList_majLazy_dropmajWhile';

// Data_List_Lazy_drop
function majData_majList_majLazy_drop(int $n_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_drop';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $__t0 = null;;
  switch ($v_2) {
case 0:
$__t0 = $v1_3;
goto end_branch_0;;
break;
default:
;
break;
};
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__tco_1 = ($v_2 - 1);
$__tco_2 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_3)->{'value1'});
$__tco_var_go__1_0_0_v_2 = $__tco_1;
$__tco_var_go__1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__1_0_0;;
$__t0 = null;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Lazy_Types_List']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(($go__1_0)($n_0))))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_drop'] = __NAMESPACE__ . '\\majData_majList_majLazy_drop';

// Data_List_Lazy_slice
function majData_majList_majLazy_slice(int $start_0, $end_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_slice';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Data\List\Lazy\majData_majList_majLazy_take(($end_1 - $start_0), \Data\List\Lazy\majData_majList_majLazy_drop($start_0, $xs_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_slice'] = __NAMESPACE__ . '\\majData_majList_majLazy_slice';

// Data_List_Lazy_deleteBy
function majData_majList_majLazy_deletemajBy($eq_0, $x_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_deletemajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_deleteBy_eq_0 = $eq_0;
  $__tco_var_Data_List_Lazy_deleteBy_x_1 = $x_1;
  $__tco_var_Data_List_Lazy_deleteBy_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_deleteBy:;
  $eq_0 = $__tco_var_Data_List_Lazy_deleteBy_eq_0;
  $x_1 = $__tco_var_Data_List_Lazy_deleteBy_x_1;
  $xs_2 = $__tco_var_Data_List_Lazy_deleteBy_xs_2;
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v_3) use ($eq_0, $x_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = null;;
if ((($eq_0)($x_1))(($v_3)->{'value0'})) {
$__t1 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v_3)->{'value1'});
goto end_branch_1;;
};
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v_3)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_deletemajBy($eq_0, $x_1, ($v_3)->{'value1'}));
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_deleteBy'] = __NAMESPACE__ . '\\majData_majList_majLazy_deletemajBy';

// Data_List_Lazy_unionBy
function majData_majList_majLazy_unionmajBy($eq_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_unionmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})($xs_1))((((($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() use ($eq_0) {
  $__fn = function($b_3, $a_4 = null) use ($eq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\List\Lazy\majData_majList_majLazy_deletemajBy($eq_0, $a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(\Data\List\Lazy\majData_majList_majLazy_nubmajBymajEq($eq_0, $ys_2)))($xs_1));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_unionBy'] = __NAMESPACE__ . '\\majData_majList_majLazy_unionmajBy';

// Data_List_Lazy_union
function majData_majList_majLazy_union($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_unionBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_union'] = __NAMESPACE__ . '\\majData_majList_majLazy_union';

// Data_List_Lazy_deleteAt
function majData_majList_majLazy_deletemajAt(int $n_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_deletemajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_Lazy_deleteAt_n_0 = $n_0;
  $__tco_var_Data_List_Lazy_deleteAt_xs_1 = $xs_1;
  tco_loop_Data_List_Lazy_deleteAt:;
  $n_0 = $__tco_var_Data_List_Lazy_deleteAt_n_0;
  $xs_1 = $__tco_var_Data_List_Lazy_deleteAt_xs_1;
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v1_2) use ($n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t0 = match ($n_0) { 0 => \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_2)->{'value1'}), default => new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v1_2)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_deletemajAt(($n_0 - 1), ($v1_2)->{'value1'})) };
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_deleteAt'] = __NAMESPACE__ . '\\majData_majList_majLazy_deletemajAt';

// Data_List_Lazy_delete
function majData_majList_majLazy_delete($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_delete';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_deleteBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_delete'] = __NAMESPACE__ . '\\majData_majList_majLazy_delete';

// Data_List_Lazy_difference
function majData_majList_majLazy_difference($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_difference';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_Lazy_Types_foldableList'])->{'foldl'})((function() use ($dictEq_0) {
  $__fn = function($b_1, $a_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\List\Lazy\majData_majList_majLazy_deletemajBy(($dictEq_0)->{'eq'}, $a_2, $b_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_difference'] = __NAMESPACE__ . '\\majData_majList_majLazy_difference';

// Data_List_Lazy_cycle
function majData_majList_majLazy_cycle($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_cycle';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (($GLOBALS['Data_List_Lazy_Types_lazyList'])->{'defer'})(function($v_2) use (&$go__1_0, $xs_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Lazy_Types_semigroupList'])->{'append'})($xs_0))($go__1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = $go__1_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_cycle'] = __NAMESPACE__ . '\\majData_majList_majLazy_cycle';

// Data_List_Lazy_concatMap
function majData_majList_majLazy_concatmajMap($b_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_concatmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})($a_1))($b_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_concatMap'] = __NAMESPACE__ . '\\majData_majList_majLazy_concatmajMap';

// Data_List_Lazy_concat
function majData_majList_majLazy_concat($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_concat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_List_Lazy_Types_bindList'])->{'bind'})($v_0))($GLOBALS['Data_List_Lazy_identity']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_concat'] = __NAMESPACE__ . '\\majData_majList_majLazy_concat';

// Data_List_Lazy_catMaybes_closure
$GLOBALS['Data_List_Lazy_catMaybes_closure'] = ($GLOBALS['Data_List_Lazy_mapMaybe'])($GLOBALS['Data_List_Lazy_identity']);

// Data_List_Lazy_catMaybes
function majData_majList_majLazy_catmajMaybes($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_catmajMaybes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_Lazy_catMaybes_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Lazy_catMaybes'] = __NAMESPACE__ . '\\majData_majList_majLazy_catmajMaybes';

// Data_List_Lazy_alterAt
function majData_majList_majLazy_altermajAt(int $n_0, $f_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_altermajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_Lazy_alterAt_n_0 = $n_0;
  $__tco_var_Data_List_Lazy_alterAt_f_1 = $f_1;
  $__tco_var_Data_List_Lazy_alterAt_xs_2 = $xs_2;
  tco_loop_Data_List_Lazy_alterAt:;
  $n_0 = $__tco_var_Data_List_Lazy_alterAt_n_0;
  $f_1 = $__tco_var_Data_List_Lazy_alterAt_f_1;
  $xs_2 = $__tco_var_Data_List_Lazy_alterAt_xs_2;
  $__res = ((($GLOBALS['Data_Lazy_functorLazy'])->{'map'})(function($v1_3) use ($f_1, $n_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil) {
$__t0 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Nil();
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons) {
$__t1 = null;;
switch ($n_0) {
case 0:
$v2_4_2 = ($f_1)(($v1_3)->{'value0'});
$__t3 = null;;
if ($v2_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = \Data\List\Lazy\Types\majData_majList_majLazy_majTypes_step(($v1_3)->{'value1'});
goto end_branch_3;;
};
if ($v2_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v2_4_2)->{'value0'}, ($v1_3)->{'value1'});
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
break;
default:
;
break;
};
$__t1 = new \Data\List\Lazy\Types\Data_List_Lazy_Types_Cons(($v1_3)->{'value0'}, \Data\List\Lazy\majData_majList_majLazy_altermajAt(($n_0 - 1), $f_1, ($v1_3)->{'value1'}));
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_Lazy_alterAt'] = __NAMESPACE__ . '\\majData_majList_majLazy_altermajAt';

// Data_List_Lazy_modifyAt
function majData_majList_majLazy_modifymajAt(int $n_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majLazy_modifymajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_Lazy_alterAt'])($n_0))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($f_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Lazy_modifyAt'] = __NAMESPACE__ . '\\majData_majList_majLazy_modifymajAt';

