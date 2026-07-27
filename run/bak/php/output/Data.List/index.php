<?php

namespace Data\List;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Internal, Data.List.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Internal, Data.List.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
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




// Data_List_tailRecM2
function majData_majList_tailmajRecmajM2($f_0, $a_1 = null, $b_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_tailmajRecmajM2';
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
$GLOBALS['Data_List_tailRecM2'] = __NAMESPACE__ . '\\majData_majList_tailmajRecmajM2';

// Data_List_lessThan
$GLOBALS['Data_List_lessThan'] = (function() use (&$__fn) {
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

// Data_List_greaterThan
$GLOBALS['Data_List_greaterThan'] = (function() use (&$__fn) {
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

// Data_List_any
$GLOBALS['Data_List_any'] = (function() use (&$__fn) {
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
return (($GLOBALS['Data_List_Types_foldableList'])->{'foldMap'})((object)["mempty" => ($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'ff'}, "Semigroup0" => function($_dollar__unused_1) use ($semigroupDisj1_0_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
})();

// Data_List_identity
function majData_majList_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_identity'] = __NAMESPACE__ . '\\majData_majList_identity';

// Data_List_Pattern
function majData_majList_majPattern($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majPattern';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Pattern'] = __NAMESPACE__ . '\\majData_majList_majPattern';

// Data_List_updateAt
function majData_majList_updatemajAt(int $v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_updatemajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_updateAt_v_0 = $v_0;
  $__tco_var_Data_List_updateAt_v1_1 = $v1_1;
  $__tco_var_Data_List_updateAt_v2_2 = $v2_2;
  tco_loop_Data_List_updateAt:;
  $v_0 = $__tco_var_Data_List_updateAt_v_0;
  $v1_1 = $__tco_var_Data_List_updateAt_v1_1;
  $v2_2 = $__tco_var_Data_List_updateAt_v2_2;
  $__t0 = null;;
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
switch ($v_0) {
case 0:
$__t2 = new \Data\Maybe\Data_Maybe_Just(new \Data\List\Types\Data_List_Types_Cons($v1_1, ($v2_2)->{'value1'}));
goto end_branch_2;;
break;
default:
;
break;
};
$__local_var_3_1 = ($v2_2)->{'value0'};
$__t2 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v3_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($__local_var_3_1, $v3_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\majData_majList_updatemajAt(($v_0 - 1), $v1_1, ($v2_2)->{'value1'}));
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_updateAt'] = __NAMESPACE__ . '\\majData_majList_updatemajAt';

// Data_List_unzip_closure
$GLOBALS['Data_List_unzip_closure'] = ((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})(function($v_0) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($v_0)->{'value0'};
  $__local_var_2_1 = ($v_0)->{'value1'};
  $__res = function($v1_3) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(new \Data\List\Types\Data_List_Types_Cons($__local_var_1_0, ($v1_3)->{'value0'}), new \Data\List\Types\Data_List_Types_Cons($__local_var_2_1, ($v1_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple(new \Data\List\Types\Data_List_Types_Nil(), new \Data\List\Types\Data_List_Types_Nil()));

// Data_List_unzip
function majData_majList_unzip($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_unzip';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_unzip_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_unzip'] = __NAMESPACE__ . '\\majData_majList_unzip';

// Data_List_uncons
function majData_majList_uncons($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\Maybe\Data_Maybe_Just((object)["head" => ($v_0)->{'value0'}, "tail" => ($v_0)->{'value1'}]);
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
$GLOBALS['Data_List_uncons'] = __NAMESPACE__ . '\\majData_majList_uncons';

// Data_List_toUnfoldable
function majData_majList_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($dictUnfoldable_0)->{'unfoldr'})(function($xs_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($xs_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($xs_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\Maybe\Data_Maybe_Just((object)["head" => ($xs_1)->{'value0'}, "tail" => ($xs_1)->{'value1'}]);
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($rec_2) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($rec_2)->{'head'}, ($rec_2)->{'tail'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_toUnfoldable'] = __NAMESPACE__ . '\\majData_majList_tomajUnfoldable';

// Data_List_tail
function majData_majList_tail($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value1'});
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
$GLOBALS['Data_List_tail'] = __NAMESPACE__ . '\\majData_majList_tail';

// Data_List_stripPrefix
function majData_majList_stripmajPrefix($dictEq_0, $v_1 = null, $s_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_stripmajPrefix';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_List_tailRecM2'])((function() use ($dictEq_0) {
  $__fn = function($prefix_3, $input_4 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($input_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if ($prefix_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
if (((($dictEq_0)->{'eq'})(($prefix_3)->{'value0'}))(($input_4)->{'value0'})) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((object)["a" => ($prefix_3)->{'value1'}, "b" => ($input_4)->{'value1'}]));
goto end_branch_2;;
};
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
if ($prefix_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done($input_4));
goto end_branch_1;;
};
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($prefix_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done($input_4));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
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
$GLOBALS['Data_List_stripPrefix'] = __NAMESPACE__ . '\\majData_majList_stripmajPrefix';

// Data_List_span
function majData_majList_span($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_span';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_span_v_0 = $v_0;
  $__tco_var_Data_List_span_v1_1 = $v1_1;
  tco_loop_Data_List_span:;
  $v_0 = $__tco_var_Data_List_span_v_0;
  $v1_1 = $__tco_var_Data_List_span_v1_1;
  $__t0 = null;;
  if (($v1_1 instanceof \Data\List\Types\Data_List_Types_Cons && ($v_0)(($v1_1)->{'value0'}))) {
$v2_2_1 = \Data\List\majData_majList_span($v_0, ($v1_1)->{'value1'});
$__t0 = (object)["init" => new \Data\List\Types\Data_List_Types_Cons(($v1_1)->{'value0'}, ($v2_2_1)->{'init'}), "rest" => ($v2_2_1)->{'rest'}];
goto end_branch_0;;
};
  $__t0 = (object)["init" => new \Data\List\Types\Data_List_Types_Nil(), "rest" => $v1_1];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_span'] = __NAMESPACE__ . '\\majData_majList_span';

// Data_List_snoc
function majData_majList_snoc($xs_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_snoc';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})($GLOBALS['Data_List_Types_Cons']))(new \Data\List\Types\Data_List_Types_Cons($x_1, new \Data\List\Types\Data_List_Types_Nil())))($xs_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_snoc'] = __NAMESPACE__ . '\\majData_majList_snoc';

// Data_List_singleton
function majData_majList_singleton($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_0, new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_singleton'] = __NAMESPACE__ . '\\majData_majList_singleton';

// Data_List_sortBy
function majData_majList_sortmajBy($cmp_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_sortmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $merge_1_0 = null;
  $merge_1_0 = (function() use ($cmp_0, &$merge_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($cmp_0, &$merge_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($v_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t3 = null;;
if (((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_0)(($v_2)->{'value0'}))(($v1_3)->{'value0'})))(new \Data\Ordering\Data_Ordering_GT())) {
$__t3 = new \Data\List\Types\Data_List_Types_Cons(($v1_3)->{'value0'}, (($merge_1_0)($v_2))(($v1_3)->{'value1'}));
goto end_branch_3;;
};
$__t3 = new \Data\List\Types\Data_List_Types_Cons(($v_2)->{'value0'}, (($merge_1_0)(($v_2)->{'value1'}))($v1_3));
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_2;
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v1_3;
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_2;
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
  $mergePairs_2_4 = null;
  $mergePairs_2_4 = function($v_3) use (&$mergePairs_2_4, &$merge_1_0) {
  $__num = \func_num_args();
  $__t5 = null;;
  if (($v_3 instanceof \Data\List\Types\Data_List_Types_Cons && ($v_3)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons)) {
$__t5 = new \Data\List\Types\Data_List_Types_Cons((($merge_1_0)(($v_3)->{'value0'}))((($v_3)->{'value1'})->{'value0'}), ($mergePairs_2_4)((($v_3)->{'value1'})->{'value1'}));
goto end_branch_5;;
};
  $__t5 = $v_3;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $mergeAll_3_6 = null;
  $mergeAll_3_6 = function($v_4) use (&$mergeAll_3_6, &$mergePairs_2_4) {
  $__num = \func_num_args();
  $__tco_var_mergeAll_3_6_6_v_4 = $v_4;
  tco_loop_mergeAll_3_6_6:;
  $v_4 = $__tco_var_mergeAll_3_6_6_v_4;
  $__t7 = null;;
  if (($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && ($v_4)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil)) {
$__t7 = ($v_4)->{'value0'};
goto end_branch_7;;
};
  $__tco_6 = ($mergePairs_2_4)($v_4);
  $__tco_var_mergeAll_3_6_6_v_4 = $__tco_6;
  goto tco_loop_mergeAll_3_6_6;;
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $sequences_4_7 = null;
  $descending_4_7 = null;
  $ascending_4_7 = null;
  $sequences_4_7 = function($v_5) use (&$ascending_4_7, $cmp_0, &$descending_4_7) {
  $__num = \func_num_args();
  $__t8 = null;;
  if (($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && ($v_5)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons)) {
$__t10 = null;;
if (((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_0)(($v_5)->{'value0'}))((($v_5)->{'value1'})->{'value0'})))(new \Data\Ordering\Data_Ordering_GT())) {
$__t10 = ((($descending_4_7)((($v_5)->{'value1'})->{'value0'}))(new \Data\List\Types\Data_List_Types_Cons(($v_5)->{'value0'}, new \Data\List\Types\Data_List_Types_Nil())))((($v_5)->{'value1'})->{'value1'});
goto end_branch_10;;
};
$__local_var_6_9 = ($v_5)->{'value0'};
$__t10 = ((($ascending_4_7)((($v_5)->{'value1'})->{'value0'}))(function($v1_7) use ($__local_var_6_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($__local_var_6_9, $v1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_5)->{'value1'})->{'value1'});
end_branch_10:;
$__t8 = $__t10;
goto end_branch_8;;
};
  $__t8 = new \Data\List\Types\Data_List_Types_Cons($v_5, new \Data\List\Types\Data_List_Types_Nil());
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $descending_4_7 = (function() use ($cmp_0, &$descending_4_7, &$sequences_4_7) {
  $__fn = function($v_5, $v1_6 = null, $v2_7 = null) use ($cmp_0, &$descending_4_7, &$sequences_4_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t11 = null;;
  if (($v2_7 instanceof \Data\List\Types\Data_List_Types_Cons && ((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_0)($v_5))(($v2_7)->{'value0'})))(new \Data\Ordering\Data_Ordering_GT()))) {
$__t11 = ((($descending_4_7)(($v2_7)->{'value0'}))(new \Data\List\Types\Data_List_Types_Cons($v_5, $v1_6)))(($v2_7)->{'value1'});
goto end_branch_11;;
};
  $__t11 = new \Data\List\Types\Data_List_Types_Cons(new \Data\List\Types\Data_List_Types_Cons($v_5, $v1_6), ($sequences_4_7)($v2_7));
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $ascending_4_7 = (function() use (&$ascending_4_7, $cmp_0, &$sequences_4_7) {
  $__fn = function($v_5, $v1_6 = null, $v2_7 = null) use (&$ascending_4_7, $cmp_0, &$sequences_4_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t12 = null;;
  if (($v2_7 instanceof \Data\List\Types\Data_List_Types_Cons && ((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_0)($v_5))(($v2_7)->{'value0'})))(new \Data\Ordering\Data_Ordering_GT())))(false))) {
$__t12 = ((($ascending_4_7)(($v2_7)->{'value0'}))(function($ys_8) use ($v1_6, $v_5) {
  $__num = \func_num_args();
  $__res = ($v1_6)(new \Data\List\Types\Data_List_Types_Cons($v_5, $ys_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v2_7)->{'value1'});
goto end_branch_12;;
};
  $__t12 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)(new \Data\List\Types\Data_List_Types_Cons($v_5, new \Data\List\Types\Data_List_Types_Nil())), ($sequences_4_7)($v2_7));
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($mergeAll_3_6))($sequences_4_7);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_sortBy'] = __NAMESPACE__ . '\\majData_majList_sortmajBy';

// Data_List_sort
function majData_majList_sort($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_sort';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = \Data\List\majData_majList_sortmajBy($compare_1_0, $xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_sort'] = __NAMESPACE__ . '\\majData_majList_sort';

// Data_List_showPattern
function majData_majList_showmajPattern($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_showmajPattern';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Pattern "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((\Data\List\Types\majData_majList_majTypes_showmajList($dictShow_0))->{'show'})($v_1)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_showPattern'] = __NAMESPACE__ . '\\majData_majList_showmajPattern';

// Data_List_reverse_closure
$GLOBALS['Data_List_reverse_closure'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = (function() use (&$go__0_0) {
  $__fn = function($v_1, $v1_2 = null) use (&$go__0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__0_0_0_v_1 = $v_1;
  $__tco_var_go__0_0_0_v1_2 = $v1_2;
  tco_loop_go__0_0_0:;
  $v_1 = $__tco_var_go__0_0_0_v_1;
  $v1_2 = $__tco_var_go__0_0_0_v1_2;
  $__t0 = null;;
  if ($v1_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v_1;
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v1_2)->{'value0'}, $v_1);
$__tco_2 = ($v1_2)->{'value1'};
$__tco_var_go__0_0_0_v_1 = $__tco_1;
$__tco_var_go__0_0_0_v1_2 = $__tco_2;
goto tco_loop_go__0_0_0;;
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
return ($go__0_0)(new \Data\List\Types\Data_List_Types_Nil());
})();

// Data_List_reverse
function majData_majList_reverse($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_reverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_reverse_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_reverse'] = __NAMESPACE__ . '\\majData_majList_reverse';

// Data_List_take_closure
$GLOBALS['Data_List_take_closure'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = (function() use (&$go__0_0) {
  $__fn = function($v_1, $v1_2 = null, $v2_3 = null) use (&$go__0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__0_0_0_v_1 = $v_1;
  $__tco_var_go__0_0_0_v1_2 = $v1_2;
  $__tco_var_go__0_0_0_v2_3 = $v2_3;
  tco_loop_go__0_0_0:;
  $v_1 = $__tco_var_go__0_0_0_v_1;
  $v1_2 = $__tco_var_go__0_0_0_v1_2;
  $v2_3 = $__tco_var_go__0_0_0_v2_3;
  $__t0 = null;;
  if ((($GLOBALS['Data_List_lessThan'])($v1_2))(1)) {
$go__4_1 = null;
$go__4_1 = (function() use (&$__tco_var_go__0_0_0_v_1, &$__tco_var_go__0_0_0_v1_2, &$__tco_var_go__0_0_0_v2_3, &$go__4_1) {
  $__fn = function($v_5, $v1_6 = null) use (&$__tco_var_go__0_0_0_v_1, &$__tco_var_go__0_0_0_v1_2, &$__tco_var_go__0_0_0_v2_3, &$go__4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_1_1_v_5 = $v_5;
  $__tco_var_go__4_1_1_v1_6 = $v1_6;
  tco_loop_go__4_1_1:;
  $v_5 = $__tco_var_go__4_1_1_v_5;
  $v1_6 = $__tco_var_go__4_1_1_v1_6;
  $__t1 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go__4_1_1_v_5 = $__tco_2;
$__tco_var_go__4_1_1_v1_6 = $__tco_3;
goto tco_loop_go__4_1_1;;
$__t1 = null;
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
$__t0 = (($go__4_1)(new \Data\List\Types\Data_List_Types_Nil()))($v_1);
goto end_branch_0;;
};
  if ($v2_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__4_2 = null;
$go__4_2 = (function() use (&$__tco_var_go__0_0_0_v_1, &$__tco_var_go__0_0_0_v1_2, &$__tco_var_go__0_0_0_v2_3, &$go__4_2) {
  $__fn = function($v_5, $v1_6 = null) use (&$__tco_var_go__0_0_0_v_1, &$__tco_var_go__0_0_0_v1_2, &$__tco_var_go__0_0_0_v2_3, &$go__4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_2_2_v_5 = $v_5;
  $__tco_var_go__4_2_2_v1_6 = $v1_6;
  tco_loop_go__4_2_2:;
  $v_5 = $__tco_var_go__4_2_2_v_5;
  $v1_6 = $__tco_var_go__4_2_2_v1_6;
  $__t2 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_5;
goto end_branch_2;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_4 = ($v1_6)->{'value1'};
$__tco_var_go__4_2_2_v_5 = $__tco_3;
$__tco_var_go__4_2_2_v1_6 = $__tco_4;
goto tco_loop_go__4_2_2;;
$__t2 = null;
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
$__t0 = (($go__4_2)(new \Data\List\Types\Data_List_Types_Nil()))($v_1);
goto end_branch_0;;
};
  if ($v2_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v2_3)->{'value0'}, $v_1);
$__tco_4 = ($v1_2 - 1);
$__tco_5 = ($v2_3)->{'value1'};
$__tco_var_go__0_0_0_v_1 = $__tco_3;
$__tco_var_go__0_0_0_v1_2 = $__tco_4;
$__tco_var_go__0_0_0_v2_3 = $__tco_5;
goto tco_loop_go__0_0_0;;
$__t0 = null;
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
return ($go__0_0)(new \Data\List\Types\Data_List_Types_Nil());
})();

// Data_List_take
function majData_majList_take(int $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_take';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_take_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_take'] = __NAMESPACE__ . '\\majData_majList_take';

// Data_List_takeWhile
function majData_majList_takemajWhile($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_takemajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $__t1 = null;;
  if (($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons && ($p_0)(($v1_3)->{'value0'}))) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_3)->{'value0'}, $v_2);
$__tco_3 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_2;
$__tco_var_go__1_0_0_v1_3 = $__tco_3;
goto tco_loop_go__1_0_0;;
$__t1 = null;
goto end_branch_1;;
};
  $go__4_0 = null;
  $go__4_0 = (function() use (&$__tco_var_go__1_0_0_v_2, &$__tco_var_go__1_0_0_v1_3, &$go__4_0) {
  $__fn = function($v_5, $v1_6 = null) use (&$__tco_var_go__1_0_0_v_2, &$__tco_var_go__1_0_0_v1_3, &$go__4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_0_0_v_5 = $v_5;
  $__tco_var_go__4_0_0_v1_6 = $v1_6;
  tco_loop_go__4_0_0:;
  $v_5 = $__tco_var_go__4_0_0_v_5;
  $v1_6 = $__tco_var_go__4_0_0_v1_6;
  $__t0 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v_5;
goto end_branch_0;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_2 = ($v1_6)->{'value1'};
$__tco_var_go__4_0_0_v_5 = $__tco_1;
$__tco_var_go__4_0_0_v1_6 = $__tco_2;
goto tco_loop_go__4_0_0;;
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
  $__t1 = (($go__4_0)(new \Data\List\Types\Data_List_Types_Nil()))($v_2);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go__1_0)(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_takeWhile'] = __NAMESPACE__ . '\\majData_majList_takemajWhile';

// Data_List_unsnoc
function majData_majList_unsnoc($lst_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_unsnoc';
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
  if ($v_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t3 = null;;
if (($v_2)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = new \Data\Maybe\Data_Maybe_Just((object)["revInit" => $v1_3, "last" => ($v_2)->{'value0'}]);
goto end_branch_3;;
};
$__tco_1 = ($v_2)->{'value1'};
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v_2)->{'value0'}, $v1_3);
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
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($h_2) {
  $__num = \func_num_args();
  $go__3_1 = null;
  $go__3_1 = (function() use (&$go__3_1) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_1_1_v_4 = $v_4;
  $__tco_var_go__3_1_1_v1_5 = $v1_5;
  tco_loop_go__3_1_1:;
  $v_4 = $__tco_var_go__3_1_1_v_4;
  $v1_5 = $__tco_var_go__3_1_1_v1_5;
  $__t1 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_4;
goto end_branch_1;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go__3_1_1_v_4 = $__tco_2;
$__tco_var_go__3_1_1_v1_5 = $__tco_3;
goto tco_loop_go__3_1_1;;
$__t1 = null;
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
  $__res = (object)["init" => (($go__3_1)(new \Data\List\Types\Data_List_Types_Nil()))(($h_2)->{'revInit'}), "last" => ($h_2)->{'last'}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($go__1_0)($lst_0))(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_unsnoc'] = __NAMESPACE__ . '\\majData_majList_unsnoc';

// Data_List_zipWith
function majData_majList_zipmajWith($f_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_zipmajWith';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = (function() use ($f_0, &$go__3_0) {
  $__fn = function($v_4, $v1_5 = null, $v2_6 = null) use ($f_0, &$go__3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__3_0_0_v_4 = $v_4;
  $__tco_var_go__3_0_0_v1_5 = $v1_5;
  $__tco_var_go__3_0_0_v2_6 = $v2_6;
  tco_loop_go__3_0_0:;
  $v_4 = $__tco_var_go__3_0_0_v_4;
  $v1_5 = $__tco_var_go__3_0_0_v1_5;
  $v2_6 = $__tco_var_go__3_0_0_v2_6;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v2_6;
goto end_branch_0;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v2_6;
goto end_branch_0;;
};
  if (($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_5 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$__tco_1 = ($v_4)->{'value1'};
$__tco_2 = ($v1_5)->{'value1'};
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons((($f_0)(($v_4)->{'value0'}))(($v1_5)->{'value0'}), $v2_6);
$__tco_var_go__3_0_0_v_4 = $__tco_1;
$__tco_var_go__3_0_0_v1_5 = $__tco_2;
$__tco_var_go__3_0_0_v2_6 = $__tco_3;
goto tco_loop_go__3_0_0;;
$__t0 = null;
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
  $go__4_1 = null;
  $go__4_1 = (function() use (&$go__4_1) {
  $__fn = function($v_5, $v1_6 = null) use (&$go__4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_1_1_v_5 = $v_5;
  $__tco_var_go__4_1_1_v1_6 = $v1_6;
  tco_loop_go__4_1_1:;
  $v_5 = $__tco_var_go__4_1_1_v_5;
  $v1_6 = $__tco_var_go__4_1_1_v1_6;
  $__t1 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go__4_1_1_v_5 = $__tco_2;
$__tco_var_go__4_1_1_v1_6 = $__tco_3;
goto tco_loop_go__4_1_1;;
$__t1 = null;
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
  $__res = (($go__4_1)(new \Data\List\Types\Data_List_Types_Nil()))(((($go__3_0)($xs_1))($ys_2))(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_zipWith'] = __NAMESPACE__ . '\\majData_majList_zipmajWith';

// Data_List_zip
function majData_majList_zip($xs_0, $ys_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_zip';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use (&$go__2_0) {
  $__fn = function($v_3, $v1_4 = null, $v2_5 = null) use (&$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__2_0_0_v_3 = $v_3;
  $__tco_var_go__2_0_0_v1_4 = $v1_4;
  $__tco_var_go__2_0_0_v2_5 = $v2_5;
  tco_loop_go__2_0_0:;
  $v_3 = $__tco_var_go__2_0_0_v_3;
  $v1_4 = $__tco_var_go__2_0_0_v1_4;
  $v2_5 = $__tco_var_go__2_0_0_v2_5;
  $__t0 = null;;
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v2_5;
goto end_branch_0;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v2_5;
goto end_branch_0;;
};
  if (($v_3 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_4 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$__tco_1 = ($v_3)->{'value1'};
$__tco_2 = ($v1_4)->{'value1'};
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(new \Data\Tuple\Data_Tuple_Tuple(($v_3)->{'value0'}, ($v1_4)->{'value0'}), $v2_5);
$__tco_var_go__2_0_0_v_3 = $__tco_1;
$__tco_var_go__2_0_0_v1_4 = $__tco_2;
$__tco_var_go__2_0_0_v2_5 = $__tco_3;
goto tco_loop_go__2_0_0;;
$__t0 = null;
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
  $go__3_1 = null;
  $go__3_1 = (function() use (&$go__3_1) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_1_1_v_4 = $v_4;
  $__tco_var_go__3_1_1_v1_5 = $v1_5;
  tco_loop_go__3_1_1:;
  $v_4 = $__tco_var_go__3_1_1_v_4;
  $v1_5 = $__tco_var_go__3_1_1_v1_5;
  $__t1 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_4;
goto end_branch_1;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go__3_1_1_v_4 = $__tco_2;
$__tco_var_go__3_1_1_v1_5 = $__tco_3;
goto tco_loop_go__3_1_1;;
$__t1 = null;
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
  $__res = (($go__3_1)(new \Data\List\Types\Data_List_Types_Nil()))(((($go__2_0)($xs_0))($ys_1))(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_zip'] = __NAMESPACE__ . '\\majData_majList_zip';

// Data_List_zipWithA
function majData_majList_zipmajWithmajA($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_zipmajWithmajA';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequence1_1_0 = (($GLOBALS['Data_List_Types_traversableList'])->{'sequence'})($dictApplicative_0);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__5_1 = null;
  $go__5_1 = (function() use ($f_2, &$go__5_1) {
  $__fn = function($v_6, $v1_7 = null, $v2_8 = null) use ($f_2, &$go__5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__5_1_1_v_6 = $v_6;
  $__tco_var_go__5_1_1_v1_7 = $v1_7;
  $__tco_var_go__5_1_1_v2_8 = $v2_8;
  tco_loop_go__5_1_1:;
  $v_6 = $__tco_var_go__5_1_1_v_6;
  $v1_7 = $__tco_var_go__5_1_1_v1_7;
  $v2_8 = $__tco_var_go__5_1_1_v2_8;
  $__t1 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v2_8;
goto end_branch_1;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v2_8;
goto end_branch_1;;
};
  if (($v_6 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_7 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$__tco_2 = ($v_6)->{'value1'};
$__tco_3 = ($v1_7)->{'value1'};
$__tco_4 = new \Data\List\Types\Data_List_Types_Cons((($f_2)(($v_6)->{'value0'}))(($v1_7)->{'value0'}), $v2_8);
$__tco_var_go__5_1_1_v_6 = $__tco_2;
$__tco_var_go__5_1_1_v1_7 = $__tco_3;
$__tco_var_go__5_1_1_v2_8 = $__tco_4;
goto tco_loop_go__5_1_1;;
$__t1 = null;
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
  $go__6_2 = null;
  $go__6_2 = (function() use (&$go__6_2) {
  $__fn = function($v_7, $v1_8 = null) use (&$go__6_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__6_2_2_v_7 = $v_7;
  $__tco_var_go__6_2_2_v1_8 = $v1_8;
  tco_loop_go__6_2_2:;
  $v_7 = $__tco_var_go__6_2_2_v_7;
  $v1_8 = $__tco_var_go__6_2_2_v1_8;
  $__t2 = null;;
  if ($v1_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_7;
goto end_branch_2;;
};
  if ($v1_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v1_8)->{'value0'}, $v_7);
$__tco_4 = ($v1_8)->{'value1'};
$__tco_var_go__6_2_2_v_7 = $__tco_3;
$__tco_var_go__6_2_2_v1_8 = $__tco_4;
goto tco_loop_go__6_2_2;;
$__t2 = null;
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
  $__res = ($sequence1_1_0)((($go__6_2)(new \Data\List\Types\Data_List_Types_Nil()))(((($go__5_1)($xs_3))($ys_4))(new \Data\List\Types\Data_List_Types_Nil())));
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
$GLOBALS['Data_List_zipWithA'] = __NAMESPACE__ . '\\majData_majList_zipmajWithmajA';

// Data_List_range
function majData_majList_range(int $start_0, $end_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_range';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if (($start_0 === $end_1)) {
$__t2 = new \Data\List\Types\Data_List_Types_Cons($start_0, new \Data\List\Types\Data_List_Types_Nil());
goto end_branch_2;;
};
  $go__2_0 = null;
  $go__2_0 = (function() use (&$go__2_0) {
  $__fn = function($s_3, $e_4 = null, $step_5 = null, $rest_6 = null) use (&$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_go__2_0_0_s_3 = $s_3;
  $__tco_var_go__2_0_0_e_4 = $e_4;
  $__tco_var_go__2_0_0_step_5 = $step_5;
  $__tco_var_go__2_0_0_rest_6 = $rest_6;
  tco_loop_go__2_0_0:;
  $s_3 = $__tco_var_go__2_0_0_s_3;
  $e_4 = $__tco_var_go__2_0_0_e_4;
  $step_5 = $__tco_var_go__2_0_0_step_5;
  $rest_6 = $__tco_var_go__2_0_0_rest_6;
  $__t4 = null;;
  if (($s_3 === $e_4)) {
$__t4 = new \Data\List\Types\Data_List_Types_Cons($s_3, $rest_6);
goto end_branch_4;;
};
  $__tco_0 = ($s_3 + $step_5);
  $__tco_1 = $e_4;
  $__tco_2 = $step_5;
  $__tco_3 = new \Data\List\Types\Data_List_Types_Cons($s_3, $rest_6);
  $__tco_var_go__2_0_0_s_3 = $__tco_0;
  $__tco_var_go__2_0_0_e_4 = $__tco_1;
  $__tco_var_go__2_0_0_step_5 = $__tco_2;
  $__tco_var_go__2_0_0_rest_6 = $__tco_3;
  goto tco_loop_go__2_0_0;;
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();
  $__t1 = null;;
  if ((($GLOBALS['Data_List_greaterThan'])($start_0))($end_1)) {
$__t1 = 1;
goto end_branch_1;;
};
  $__t1 = -1;
  end_branch_1:;
  $__t2 = (((($go__2_0)($end_1))($start_0))($__t1))(new \Data\List\Types\Data_List_Types_Nil());
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_range'] = __NAMESPACE__ . '\\majData_majList_range';

// Data_List_partition
function majData_majList_partition($p_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_partition';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})((function() use ($p_0) {
  $__fn = function($x_2, $v_3 = null) use ($p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($p_0)($x_2)) {
$__t0 = (object)["no" => ($v_3)->{'no'}, "yes" => new \Data\List\Types\Data_List_Types_Cons($x_2, ($v_3)->{'yes'})];
goto end_branch_0;;
};
  $__t0 = (object)["no" => new \Data\List\Types\Data_List_Types_Cons($x_2, ($v_3)->{'no'}), "yes" => ($v_3)->{'yes'}];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["no" => new \Data\List\Types\Data_List_Types_Nil(), "yes" => new \Data\List\Types\Data_List_Types_Nil()]))($xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_partition'] = __NAMESPACE__ . '\\majData_majList_partition';

// Data_List_null
function majData_majList_null($v_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_null';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0 instanceof \Data\List\Types\Data_List_Types_Nil;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_null'] = __NAMESPACE__ . '\\majData_majList_null';

// Data_List_nubBy
function majData_majList_nubmajBy($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_nubmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null, $v2_4 = null) use (&$go__1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__1_0_0_v_2 = $v_2;
  $__tco_var_go__1_0_0_v1_3 = $v1_3;
  $__tco_var_go__1_0_0_v2_4 = $v2_4;
  tco_loop_go__1_0_0:;
  $v_2 = $__tco_var_go__1_0_0_v_2;
  $v1_3 = $__tco_var_go__1_0_0_v1_3;
  $v2_4 = $__tco_var_go__1_0_0_v2_4;
  $__t0 = null;;
  if ($v2_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v1_3;
goto end_branch_0;;
};
  if ($v2_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$v3_5_1 = \Data\List\Internal\majData_majList_majInternal_insertmajAndmajLookupmajBy($p_0, ($v2_4)->{'value0'}, $v_2);
$__t5 = null;;
if (($v3_5_1)->{'found'}) {
$__tco_6 = ($v3_5_1)->{'result'};
$__tco_7 = $v1_3;
$__tco_8 = ($v2_4)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_6;
$__tco_var_go__1_0_0_v1_3 = $__tco_7;
$__tco_var_go__1_0_0_v2_4 = $__tco_8;
goto tco_loop_go__1_0_0;;
$__t5 = null;
goto end_branch_5;;
};
$__tco_2 = ($v3_5_1)->{'result'};
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v2_4)->{'value0'}, $v1_3);
$__tco_4 = ($v2_4)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_2;
$__tco_var_go__1_0_0_v1_3 = $__tco_3;
$__tco_var_go__1_0_0_v2_4 = $__tco_4;
goto tco_loop_go__1_0_0;;
$__t5 = null;
end_branch_5:;
$__t0 = $__t5;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_reverse']))((($go__1_0)(new \Data\List\Internal\Data_List_Internal_Leaf()))(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_nubBy'] = __NAMESPACE__ . '\\majData_majList_nubmajBy';

// Data_List_nub
function majData_majList_nub($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_nub';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_nubBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_nub'] = __NAMESPACE__ . '\\majData_majList_nub';

// Data_List_newtypePattern
$GLOBALS['Data_List_newtypePattern'] = (object)["Coercible0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_mapMaybe
function majData_majList_mapmajMaybe($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_mapmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use ($f_0, &$go__1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($f_0, &$go__1_0, &$__fn) {
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
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__4_1 = null;
$go__4_1 = (function() use (&$__tco_var_go__1_0_0_v_2, &$__tco_var_go__1_0_0_v1_3, &$go__4_1) {
  $__fn = function($v_5, $v1_6 = null) use (&$__tco_var_go__1_0_0_v_2, &$__tco_var_go__1_0_0_v1_3, &$go__4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_1_1_v_5 = $v_5;
  $__tco_var_go__4_1_1_v1_6 = $v1_6;
  tco_loop_go__4_1_1:;
  $v_5 = $__tco_var_go__4_1_1_v_5;
  $v1_6 = $__tco_var_go__4_1_1_v1_6;
  $__t1 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go__4_1_1_v_5 = $__tco_2;
$__tco_var_go__4_1_1_v1_6 = $__tco_3;
goto tco_loop_go__4_1_1;;
$__t1 = null;
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
$__t0 = (($go__4_1)(new \Data\List\Types\Data_List_Types_Nil()))($v_2);
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$v2_4_2 = ($f_0)(($v1_3)->{'value0'});
$__t3 = null;;
if ($v2_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__tco_4 = $v_2;
$__tco_5 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_4;
$__tco_var_go__1_0_0_v1_3 = $__tco_5;
goto tco_loop_go__1_0_0;;
$__t3 = null;
goto end_branch_3;;
};
if ($v2_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(($v2_4_2)->{'value0'}, $v_2);
$__tco_7 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_6;
$__tco_var_go__1_0_0_v1_3 = $__tco_7;
goto tco_loop_go__1_0_0;;
$__t3 = null;
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
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
  $__res = ($go__1_0)(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_mapMaybe'] = __NAMESPACE__ . '\\majData_majList_mapmajMaybe';

// Data_List_manyRec
function majData_majList_manymajRec($dictMonadRec_0, $dictAlternative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_manymajRec';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Alt0_2_0 = (((($dictAlternative_1)->{'Plus1'})(null))->{'Alt0'})(null);
  $__local_var_3_1 = (($dictAlternative_1)->{'Applicative0'})(null);
  $__res = function($p_4) use ($Alt0_2_0, $__local_var_3_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($acc_5) use ($Alt0_2_0, $__local_var_3_1, $dictMonadRec_0, $p_4) {
  $__num = \func_num_args();
  $__res = ((((((($dictMonadRec_0)->{'Monad0'})(null))->{'Bind1'})(null))->{'bind'})(((($Alt0_2_0)->{'alt'})(((((($Alt0_2_0)->{'Functor0'})(null))->{'map'})($GLOBALS['Control_Monad_Rec_Class_Loop']))($p_4)))((($__local_var_3_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done($GLOBALS['Data_Unit_unit'])))))(function($aa_6) use ($__local_var_3_1, $acc_5) {
  $__num = \func_num_args();
  $__res = (($__local_var_3_1)->{'pure'})((((($GLOBALS['Control_Monad_Rec_Class_bifunctorStep'])->{'bimap'})(function($v_7) use ($acc_5) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($v_7, $acc_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_7) use ($acc_5) {
  $__num = \func_num_args();
  $go__8_2 = null;
  $go__8_2 = (function() use (&$go__8_2) {
  $__fn = function($v_9, $v1_10 = null) use (&$go__8_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__8_2_2_v_9 = $v_9;
  $__tco_var_go__8_2_2_v1_10 = $v1_10;
  tco_loop_go__8_2_2:;
  $v_9 = $__tco_var_go__8_2_2_v_9;
  $v1_10 = $__tco_var_go__8_2_2_v1_10;
  $__t2 = null;;
  if ($v1_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_9;
goto end_branch_2;;
};
  if ($v1_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v1_10)->{'value0'}, $v_9);
$__tco_4 = ($v1_10)->{'value1'};
$__tco_var_go__8_2_2_v_9 = $__tco_3;
$__tco_var_go__8_2_2_v1_10 = $__tco_4;
goto tco_loop_go__8_2_2;;
$__t2 = null;
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
  $__res = (($go__8_2)(new \Data\List\Types\Data_List_Types_Nil()))($acc_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($aa_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_manyRec'] = __NAMESPACE__ . '\\majData_majList_manymajRec';

// Data_List_someRec
function majData_majList_somemajRec($dictMonadRec_0, $dictAlternative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_somemajRec';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $manyRec2_2_0 = (($GLOBALS['Data_List_manyRec'])($dictMonadRec_0))($dictAlternative_1);
  $__res = function($v_3) use ($dictAlternative_1, $manyRec2_2_0) {
  $__num = \func_num_args();
  $__res = ((((((($dictAlternative_1)->{'Applicative0'})(null))->{'Apply0'})(null))->{'apply'})(((((((((($dictAlternative_1)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_List_Types_Cons']))($v_3)))(($manyRec2_2_0)($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_someRec'] = __NAMESPACE__ . '\\majData_majList_somemajRec';

// Data_List_some
function majData_majList_some($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_some';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)->{'Applicative0'})(null))->{'Apply0'})(null))->{'apply'})(((((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_List_Types_Cons']))($v_2)))((($dictLazy_1)->{'defer'})(function($v1_3) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  $__res = \Data\List\majData_majList_many($dictAlternative_0, $dictLazy_1, $v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_some'] = __NAMESPACE__ . '\\majData_majList_some';

// Data_List_many
function majData_majList_many($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_many';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'alt'})(\Data\List\majData_majList_some($dictAlternative_0, $dictLazy_1, $v_2)))((((($dictAlternative_0)->{'Applicative0'})(null))->{'pure'})(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_many'] = __NAMESPACE__ . '\\majData_majList_many';

// Data_List_length_closure
$GLOBALS['Data_List_length_closure'] = ((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($acc_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($acc_0 + 1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(0);

// Data_List_length
function majData_majList_length($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_length_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_length'] = __NAMESPACE__ . '\\majData_majList_length';

// Data_List_last
function majData_majList_last($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_last';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_last_v_0 = $v_0;
  tco_loop_Data_List_last:;
  $v_0 = $__tco_var_Data_List_last_v_0;
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
if (($v_0)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value0'});
goto end_branch_2;;
};
$__tco_1 = ($v_0)->{'value1'};
$__tco_var_Data_List_last_v_0 = $__tco_1;
goto tco_loop_Data_List_last;;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_last'] = __NAMESPACE__ . '\\majData_majList_last';

// Data_List_insertBy
function majData_majList_insertmajBy($v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_insertmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_insertBy_v_0 = $v_0;
  $__tco_var_Data_List_insertBy_v1_1 = $v1_1;
  $__tco_var_Data_List_insertBy_v2_2 = $v2_2;
  tco_loop_Data_List_insertBy:;
  $v_0 = $__tco_var_Data_List_insertBy_v_0;
  $v1_1 = $__tco_var_Data_List_insertBy_v1_1;
  $v2_2 = $__tco_var_Data_List_insertBy_v2_2;
  $__t0 = null;;
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Cons($v1_1, new \Data\List\Types\Data_List_Types_Nil());
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if ((($v_0)($v1_1))(($v2_2)->{'value0'}) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = new \Data\List\Types\Data_List_Types_Cons(($v2_2)->{'value0'}, \Data\List\majData_majList_insertmajBy($v_0, $v1_1, ($v2_2)->{'value1'}));
goto end_branch_1;;
};
$__t1 = new \Data\List\Types\Data_List_Types_Cons($v1_1, $v2_2);
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
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_insertBy'] = __NAMESPACE__ . '\\majData_majList_insertmajBy';

// Data_List_insertAt
function majData_majList_insertmajAt(int $v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_insertmajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_insertAt_v_0 = $v_0;
  $__tco_var_Data_List_insertAt_v1_1 = $v1_1;
  $__tco_var_Data_List_insertAt_v2_2 = $v2_2;
  tco_loop_Data_List_insertAt:;
  $v_0 = $__tco_var_Data_List_insertAt_v_0;
  $v1_1 = $__tco_var_Data_List_insertAt_v1_1;
  $v2_2 = $__tco_var_Data_List_insertAt_v2_2;
  $__t0 = null;;
  switch ($v_0) {
case 0:
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Data\List\Types\Data_List_Types_Cons($v1_1, $v2_2));
goto end_branch_0;;
break;
default:
;
break;
};
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_3_1 = ($v2_2)->{'value0'};
$__t0 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v3_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($__local_var_3_1, $v3_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\majData_majList_insertmajAt(($v_0 - 1), $v1_1, ($v2_2)->{'value1'}));
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_insertAt'] = __NAMESPACE__ . '\\majData_majList_insertmajAt';

// Data_List_insert
function majData_majList_insert($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_insert';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_insertBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_insert'] = __NAMESPACE__ . '\\majData_majList_insert';

// Data_List_init
function majData_majList_init($lst_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_init';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = ($v_1)->{'init'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\majData_majList_unsnoc($lst_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_init'] = __NAMESPACE__ . '\\majData_majList_init';

// Data_List_index
function majData_majList_index($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_index';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_index_v_0 = $v_0;
  $__tco_var_Data_List_index_v1_1 = $v1_1;
  tco_loop_Data_List_index:;
  $v_0 = $__tco_var_Data_List_index_v_0;
  $v1_1 = $__tco_var_Data_List_index_v1_1;
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t3 = null;;
switch ($v1_1) {
case 0:
$__t3 = new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value0'});
goto end_branch_3;;
break;
default:
;
break;
};
$__tco_1 = ($v_0)->{'value1'};
$__tco_2 = ($v1_1 - 1);
$__tco_var_Data_List_index_v_0 = $__tco_1;
$__tco_var_Data_List_index_v1_1 = $__tco_2;
goto tco_loop_Data_List_index;;
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_index'] = __NAMESPACE__ . '\\majData_majList_index';

// Data_List_head
function majData_majList_head($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value0'});
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
$GLOBALS['Data_List_head'] = __NAMESPACE__ . '\\majData_majList_head';

// Data_List_transpose
function majData_majList_transpose($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_transpose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_transpose_v_0 = $v_0;
  tco_loop_Data_List_transpose:;
  $v_0 = $__tco_var_Data_List_transpose_v_0;
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if (($v_0)->{'value0'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__tco_2 = ($v_0)->{'value1'};
$__tco_var_Data_List_transpose_v_0 = $__tco_2;
goto tco_loop_Data_List_transpose;;
$__t1 = null;
goto end_branch_1;;
};
if (($v_0)->{'value0'} instanceof \Data\List\Types\Data_List_Types_Cons) {
$go__1_3 = null;
$go__1_3 = (function() use (&$go__1_3) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__1_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_3_3_v_2 = $v_2;
  $__tco_var_go__1_3_3_v1_3 = $v1_3;
  tco_loop_go__1_3_3:;
  $v_2 = $__tco_var_go__1_3_3_v_2;
  $v1_3 = $__tco_var_go__1_3_3_v1_3;
  $__t3 = null;;
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__4_4 = null;
$go__4_4 = (function() use (&$__tco_var_go__1_3_3_v_2, &$__tco_var_go__1_3_3_v1_3, &$go__4_4) {
  $__fn = function($v_5, $v1_6 = null) use (&$__tco_var_go__1_3_3_v_2, &$__tco_var_go__1_3_3_v1_3, &$go__4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_4_4_v_5 = $v_5;
  $__tco_var_go__4_4_4_v1_6 = $v1_6;
  tco_loop_go__4_4_4:;
  $v_5 = $__tco_var_go__4_4_4_v_5;
  $v1_6 = $__tco_var_go__4_4_4_v1_6;
  $__t4 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $v_5;
goto end_branch_4;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_6 = ($v1_6)->{'value1'};
$__tco_var_go__4_4_4_v_5 = $__tco_5;
$__tco_var_go__4_4_4_v1_6 = $__tco_6;
goto tco_loop_go__4_4_4;;
$__t4 = null;
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
})();
$__t3 = (($go__4_4)(new \Data\List\Types\Data_List_Types_Nil()))($v_2);
goto end_branch_3;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t5 = null;;
if (($v1_3)->{'value0'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__tco_6 = $v_2;
$__tco_7 = ($v1_3)->{'value1'};
$__tco_var_go__1_3_3_v_2 = $__tco_6;
$__tco_var_go__1_3_3_v1_3 = $__tco_7;
goto tco_loop_go__1_3_3;;
$__t5 = null;
goto end_branch_5;;
};
if (($v1_3)->{'value0'} instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_8 = new \Data\List\Types\Data_List_Types_Cons((($v1_3)->{'value0'})->{'value0'}, $v_2);
$__tco_9 = ($v1_3)->{'value1'};
$__tco_var_go__1_3_3_v_2 = $__tco_8;
$__tco_var_go__1_3_3_v1_3 = $__tco_9;
goto tco_loop_go__1_3_3;;
$__t5 = null;
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$__t3 = $__t5;
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
$go__1_4 = null;
$go__1_4 = (function() use (&$go__1_4) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__1_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_4_4_v_2 = $v_2;
  $__tco_var_go__1_4_4_v1_3 = $v1_3;
  tco_loop_go__1_4_4:;
  $v_2 = $__tco_var_go__1_4_4_v_2;
  $v1_3 = $__tco_var_go__1_4_4_v1_3;
  $__t4 = null;;
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__4_5 = null;
$go__4_5 = (function() use (&$__tco_var_go__1_4_4_v_2, &$__tco_var_go__1_4_4_v1_3, &$go__4_5) {
  $__fn = function($v_5, $v1_6 = null) use (&$__tco_var_go__1_4_4_v_2, &$__tco_var_go__1_4_4_v1_3, &$go__4_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_5_5_v_5 = $v_5;
  $__tco_var_go__4_5_5_v1_6 = $v1_6;
  tco_loop_go__4_5_5:;
  $v_5 = $__tco_var_go__4_5_5_v_5;
  $v1_6 = $__tco_var_go__4_5_5_v1_6;
  $__t5 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $v_5;
goto end_branch_5;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_7 = ($v1_6)->{'value1'};
$__tco_var_go__4_5_5_v_5 = $__tco_6;
$__tco_var_go__4_5_5_v1_6 = $__tco_7;
goto tco_loop_go__4_5_5;;
$__t5 = null;
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t4 = (($go__4_5)(new \Data\List\Types\Data_List_Types_Nil()))($v_2);
goto end_branch_4;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t6 = null;;
if (($v1_3)->{'value0'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__tco_7 = $v_2;
$__tco_8 = ($v1_3)->{'value1'};
$__tco_var_go__1_4_4_v_2 = $__tco_7;
$__tco_var_go__1_4_4_v1_3 = $__tco_8;
goto tco_loop_go__1_4_4;;
$__t6 = null;
goto end_branch_6;;
};
if (($v1_3)->{'value0'} instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_9 = new \Data\List\Types\Data_List_Types_Cons((($v1_3)->{'value0'})->{'value1'}, $v_2);
$__tco_10 = ($v1_3)->{'value1'};
$__tco_var_go__1_4_4_v_2 = $__tco_9;
$__tco_var_go__1_4_4_v1_3 = $__tco_10;
goto tco_loop_go__1_4_4;;
$__t6 = null;
goto end_branch_6;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
end_branch_6:;
$__t4 = $__t6;
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
})();
$__t1 = new \Data\List\Types\Data_List_Types_Cons(new \Data\List\Types\Data_List_Types_Cons((($v_0)->{'value0'})->{'value0'}, (($go__1_3)(new \Data\List\Types\Data_List_Types_Nil()))(($v_0)->{'value1'})), \Data\List\majData_majList_transpose(new \Data\List\Types\Data_List_Types_Cons((($v_0)->{'value0'})->{'value1'}, (($go__1_4)(new \Data\List\Types\Data_List_Types_Nil()))(($v_0)->{'value1'}))));
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_transpose'] = __NAMESPACE__ . '\\majData_majList_transpose';

// Data_List_groupBy
function majData_majList_groupmajBy($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_groupmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_groupBy_v_0 = $v_0;
  $__tco_var_Data_List_groupBy_v1_1 = $v1_1;
  tco_loop_Data_List_groupBy:;
  $v_0 = $__tco_var_Data_List_groupBy_v_0;
  $v1_1 = $__tco_var_Data_List_groupBy_v1_1;
  $__t0 = null;;
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$v2_2_1 = \Data\List\majData_majList_span(($v_0)(($v1_1)->{'value0'}), ($v1_1)->{'value1'});
$__t0 = new \Data\List\Types\Data_List_Types_Cons(new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v1_1)->{'value0'}, ($v2_2_1)->{'init'}), \Data\List\majData_majList_groupmajBy($v_0, ($v2_2_1)->{'rest'}));
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
$GLOBALS['Data_List_groupBy'] = __NAMESPACE__ . '\\majData_majList_groupmajBy';

// Data_List_groupAllBy
function majData_majList_groupmajAllmajBy($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_groupmajAllmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_groupBy'])((function() use ($p_0) {
  $__fn = function($x_1, $y_2 = null) use ($p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($p_0)($x_1))($y_2)))(new \Data\Ordering\Data_Ordering_EQ());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))(($GLOBALS['Data_List_sortBy'])($p_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_groupAllBy'] = __NAMESPACE__ . '\\majData_majList_groupmajAllmajBy';

// Data_List_group
function majData_majList_group($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_group';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_groupBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_group'] = __NAMESPACE__ . '\\majData_majList_group';

// Data_List_groupAll
function majData_majList_groupmajAll($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_groupmajAll';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_group'])((($dictOrd_0)->{'Eq0'})(null))))(function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = \Data\List\majData_majList_sortmajBy($compare_1_0, $xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_groupAll'] = __NAMESPACE__ . '\\majData_majList_groupmajAll';

// Data_List_fromFoldable
function majData_majList_frommajFoldable($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_frommajFoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($dictFoldable_0)->{'foldr'})($GLOBALS['Data_List_Types_Cons']))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_fromFoldable'] = __NAMESPACE__ . '\\majData_majList_frommajFoldable';

// Data_List_foldM
function majData_majList_foldmajM($dictMonad_0, $v_1 = null, $v1_2 = null, $v2_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_foldmajM';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_List_foldM_dictMonad_0 = $dictMonad_0;
  $__tco_var_Data_List_foldM_v_1 = $v_1;
  $__tco_var_Data_List_foldM_v1_2 = $v1_2;
  $__tco_var_Data_List_foldM_v2_3 = $v2_3;
  tco_loop_Data_List_foldM:;
  $dictMonad_0 = $__tco_var_Data_List_foldM_dictMonad_0;
  $v_1 = $__tco_var_Data_List_foldM_v_1;
  $v1_2 = $__tco_var_Data_List_foldM_v1_2;
  $v2_3 = $__tco_var_Data_List_foldM_v2_3;
  $__t0 = null;;
  if ($v2_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})($v1_2);
goto end_branch_0;;
};
  if ($v2_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_4_1 = ($v2_3)->{'value1'};
$__t0 = ((((($dictMonad_0)->{'Bind1'})(null))->{'bind'})((($v_1)($v1_2))(($v2_3)->{'value0'})))(function($b_prime_5) use ($__local_var_4_1, $dictMonad_0, $v_1) {
  $__num = \func_num_args();
  $__res = \Data\List\majData_majList_foldmajM($dictMonad_0, $v_1, $b_prime_5, $__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_List_foldM'] = __NAMESPACE__ . '\\majData_majList_foldmajM';

// Data_List_findIndex
function majData_majList_findmajIndex($fn_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_findmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use ($fn_0, &$go__1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($fn_0, &$go__1_0, &$__fn) {
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
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t3 = null;;
if (($fn_0)(($v1_3)->{'value0'})) {
$__t3 = new \Data\Maybe\Data_Maybe_Just($v_2);
goto end_branch_3;;
};
$__tco_1 = ($v_2 + 1);
$__tco_2 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_1;
$__tco_var_go__1_0_0_v1_3 = $__tco_2;
goto tco_loop_go__1_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
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
  $__res = ($go__1_0)(0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_findIndex'] = __NAMESPACE__ . '\\majData_majList_findmajIndex';

// Data_List_findLastIndex
function majData_majList_findmajLastmajIndex($fn_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_findmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($fn_0, &$go__2_0) {
  $__fn = function($v_3, $v1_4 = null) use ($fn_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__2_0_0_v_3 = $v_3;
  $__tco_var_go__2_0_0_v1_4 = $v1_4;
  tco_loop_go__2_0_0:;
  $v_3 = $__tco_var_go__2_0_0_v_3;
  $v1_4 = $__tco_var_go__2_0_0_v1_4;
  $__t0 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t3 = null;;
if (($fn_0)(($v1_4)->{'value0'})) {
$__t3 = new \Data\Maybe\Data_Maybe_Just($v_3);
goto end_branch_3;;
};
$__tco_1 = ($v_3 + 1);
$__tco_2 = ($v1_4)->{'value1'};
$__tco_var_go__2_0_0_v_3 = $__tco_1;
$__tco_var_go__2_0_0_v1_4 = $__tco_2;
goto tco_loop_go__2_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
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
  $go__3_1 = null;
  $go__3_1 = (function() use (&$go__3_1) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_1_1_v_4 = $v_4;
  $__tco_var_go__3_1_1_v1_5 = $v1_5;
  tco_loop_go__3_1_1:;
  $v_4 = $__tco_var_go__3_1_1_v_4;
  $v1_5 = $__tco_var_go__3_1_1_v1_5;
  $__t1 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_4;
goto end_branch_1;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go__3_1_1_v_4 = $__tco_2;
$__tco_var_go__3_1_1_v1_5 = $__tco_3;
goto tco_loop_go__3_1_1;;
$__t1 = null;
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
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_2) use ($xs_1) {
  $__num = \func_num_args();
  $__res = ((\Data\List\majData_majList_length($xs_1) - 1) - $v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($go__2_0)(0))((($go__3_1)(new \Data\List\Types\Data_List_Types_Nil()))($xs_1)));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_findLastIndex'] = __NAMESPACE__ . '\\majData_majList_findmajLastmajIndex';

// Data_List_filterM
function majData_majList_filtermajM($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_filtermajM';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_List_filterM_dictMonad_0 = $dictMonad_0;
  tco_loop_Data_List_filterM:;
  $dictMonad_0 = $__tco_var_Data_List_filterM_dictMonad_0;
  $__local_var_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $__local_var_2_1 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = (($__local_var_1_0)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil());
goto end_branch_2;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_5_3 = ($v1_4)->{'value0'};
$__local_var_6_4 = ($v1_4)->{'value1'};
$__t2 = ((($__local_var_2_1)->{'bind'})(($v_3)($__local_var_5_3)))(function($b_7) use ($__local_var_1_0, $__local_var_2_1, $__local_var_5_3, $__local_var_6_4, $dictMonad_0, $v_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'bind'})(\Data\List\majData_majList_filtermajM($dictMonad_0, $v_3, $__local_var_6_4)))(function($xs_prime_8) use ($__local_var_1_0, $__local_var_5_3, $b_7) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($b_7) {
$__t5 = new \Data\List\Types\Data_List_Types_Cons($__local_var_5_3, $xs_prime_8);
goto end_branch_5;;
};
  $__t5 = $xs_prime_8;
  end_branch_5:;
  $__res = (($__local_var_1_0)->{'pure'})($__t5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_filterM'] = __NAMESPACE__ . '\\majData_majList_filtermajM';

// Data_List_filter
function majData_majList_filter($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_filter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go__1_0, $p_0, &$__fn) {
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
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__4_1 = null;
$go__4_1 = (function() use (&$__tco_var_go__1_0_0_v_2, &$__tco_var_go__1_0_0_v1_3, &$go__4_1) {
  $__fn = function($v_5, $v1_6 = null) use (&$__tco_var_go__1_0_0_v_2, &$__tco_var_go__1_0_0_v1_3, &$go__4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__4_1_1_v_5 = $v_5;
  $__tco_var_go__4_1_1_v1_6 = $v1_6;
  tco_loop_go__4_1_1:;
  $v_5 = $__tco_var_go__4_1_1_v_5;
  $v1_6 = $__tco_var_go__4_1_1_v1_6;
  $__t1 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_5;
goto end_branch_1;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_3 = ($v1_6)->{'value1'};
$__tco_var_go__4_1_1_v_5 = $__tco_2;
$__tco_var_go__4_1_1_v1_6 = $__tco_3;
goto tco_loop_go__4_1_1;;
$__t1 = null;
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
$__t0 = (($go__4_1)(new \Data\List\Types\Data_List_Types_Nil()))($v_2);
goto end_branch_0;;
};
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t4 = null;;
if (($p_0)(($v1_3)->{'value0'})) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons(($v1_3)->{'value0'}, $v_2);
$__tco_6 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_5;
$__tco_var_go__1_0_0_v1_3 = $__tco_6;
goto tco_loop_go__1_0_0;;
$__t4 = null;
goto end_branch_4;;
};
$__tco_2 = $v_2;
$__tco_3 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_2;
$__tco_var_go__1_0_0_v1_3 = $__tco_3;
goto tco_loop_go__1_0_0;;
$__t4 = null;
end_branch_4:;
$__t0 = $__t4;
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
  $__res = ($go__1_0)(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_filter'] = __NAMESPACE__ . '\\majData_majList_filter';

// Data_List_intersectBy
function majData_majList_intersectmajBy($v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_intersectmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_1;;
};
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_1;;
};
  $go__3_0 = null;
  $go__3_0 = (function() use (&$go__3_0, $v2_2, $v_0) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__3_0, $v2_2, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_0_0_v_4 = $v_4;
  $__tco_var_go__3_0_0_v1_5 = $v1_5;
  tco_loop_go__3_0_0:;
  $v_4 = $__tco_var_go__3_0_0_v_4;
  $v1_5 = $__tco_var_go__3_0_0_v1_5;
  $__t0 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__6_1 = null;
$go__6_1 = (function() use (&$__tco_var_go__3_0_0_v_4, &$__tco_var_go__3_0_0_v1_5, &$go__6_1) {
  $__fn = function($v_7, $v1_8 = null) use (&$__tco_var_go__3_0_0_v_4, &$__tco_var_go__3_0_0_v1_5, &$go__6_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__6_1_1_v_7 = $v_7;
  $__tco_var_go__6_1_1_v1_8 = $v1_8;
  tco_loop_go__6_1_1:;
  $v_7 = $__tco_var_go__6_1_1_v_7;
  $v1_8 = $__tco_var_go__6_1_1_v1_8;
  $__t1 = null;;
  if ($v1_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_7;
goto end_branch_1;;
};
  if ($v1_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_8)->{'value0'}, $v_7);
$__tco_3 = ($v1_8)->{'value1'};
$__tco_var_go__6_1_1_v_7 = $__tco_2;
$__tco_var_go__6_1_1_v1_8 = $__tco_3;
goto tco_loop_go__6_1_1;;
$__t1 = null;
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
$__t0 = (($go__6_1)(new \Data\List\Types\Data_List_Types_Nil()))($v_4);
goto end_branch_0;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t4 = null;;
if ((($GLOBALS['Data_List_any'])(($v_0)(($v1_5)->{'value0'})))($v2_2)) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_6 = ($v1_5)->{'value1'};
$__tco_var_go__3_0_0_v_4 = $__tco_5;
$__tco_var_go__3_0_0_v1_5 = $__tco_6;
goto tco_loop_go__3_0_0;;
$__t4 = null;
goto end_branch_4;;
};
$__tco_2 = $v_4;
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go__3_0_0_v_4 = $__tco_2;
$__tco_var_go__3_0_0_v1_5 = $__tco_3;
goto tco_loop_go__3_0_0;;
$__t4 = null;
end_branch_4:;
$__t0 = $__t4;
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
  $__t1 = (($go__3_0)(new \Data\List\Types\Data_List_Types_Nil()))($v1_1);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_intersectBy'] = __NAMESPACE__ . '\\majData_majList_intersectmajBy';

// Data_List_intersect
function majData_majList_intersect($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_intersect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_intersectBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_intersect'] = __NAMESPACE__ . '\\majData_majList_intersect';

// Data_List_nubByEq
function majData_majList_nubmajBymajEq($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_nubmajBymajEq';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_nubByEq_v_0 = $v_0;
  $__tco_var_Data_List_nubByEq_v1_1 = $v1_1;
  tco_loop_Data_List_nubByEq:;
  $v_0 = $__tco_var_Data_List_nubByEq_v_0;
  $v1_1 = $__tco_var_Data_List_nubByEq_v1_1;
  $__t0 = null;;
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_2_1 = ($v1_1)->{'value0'};
$go__3_2 = null;
$go__3_2 = (function() use ($__local_var_2_1, &$go__3_2, $v_0) {
  $__fn = function($v_4, $v1_5 = null) use ($__local_var_2_1, &$go__3_2, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_2_2_v_4 = $v_4;
  $__tco_var_go__3_2_2_v1_5 = $v1_5;
  tco_loop_go__3_2_2:;
  $v_4 = $__tco_var_go__3_2_2_v_4;
  $v1_5 = $__tco_var_go__3_2_2_v1_5;
  $__t2 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__6_3 = null;
$go__6_3 = (function() use (&$__tco_var_go__3_2_2_v_4, &$__tco_var_go__3_2_2_v1_5, &$go__6_3) {
  $__fn = function($v_7, $v1_8 = null) use (&$__tco_var_go__3_2_2_v_4, &$__tco_var_go__3_2_2_v1_5, &$go__6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__6_3_3_v_7 = $v_7;
  $__tco_var_go__6_3_3_v1_8 = $v1_8;
  tco_loop_go__6_3_3:;
  $v_7 = $__tco_var_go__6_3_3_v_7;
  $v1_8 = $__tco_var_go__6_3_3_v1_8;
  $__t3 = null;;
  if ($v1_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $v_7;
goto end_branch_3;;
};
  if ($v1_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = new \Data\List\Types\Data_List_Types_Cons(($v1_8)->{'value0'}, $v_7);
$__tco_5 = ($v1_8)->{'value1'};
$__tco_var_go__6_3_3_v_7 = $__tco_4;
$__tco_var_go__6_3_3_v1_8 = $__tco_5;
goto tco_loop_go__6_3_3;;
$__t3 = null;
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
$__t2 = (($go__6_3)(new \Data\List\Types\Data_List_Types_Nil()))($v_4);
goto end_branch_2;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t6 = null;;
if ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'})((($v_0)($__local_var_2_1))(($v1_5)->{'value0'}))) {
$__tco_7 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_8 = ($v1_5)->{'value1'};
$__tco_var_go__3_2_2_v_4 = $__tco_7;
$__tco_var_go__3_2_2_v1_5 = $__tco_8;
goto tco_loop_go__3_2_2;;
$__t6 = null;
goto end_branch_6;;
};
$__tco_4 = $v_4;
$__tco_5 = ($v1_5)->{'value1'};
$__tco_var_go__3_2_2_v_4 = $__tco_4;
$__tco_var_go__3_2_2_v1_5 = $__tco_5;
goto tco_loop_go__3_2_2;;
$__t6 = null;
end_branch_6:;
$__t2 = $__t6;
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
$__t0 = new \Data\List\Types\Data_List_Types_Cons($__local_var_2_1, \Data\List\majData_majList_nubmajBymajEq($v_0, (($go__3_2)(new \Data\List\Types\Data_List_Types_Nil()))(($v1_1)->{'value1'})));
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
$GLOBALS['Data_List_nubByEq'] = __NAMESPACE__ . '\\majData_majList_nubmajBymajEq';

// Data_List_nubEq
function majData_majList_nubmajEq($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_nubmajEq';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_nubByEq'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_nubEq'] = __NAMESPACE__ . '\\majData_majList_nubmajEq';

// Data_List_eqPattern
function majData_majList_eqmajPattern($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_eqmajPattern';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Types_eq1List'])->{'eq1'})($dictEq_0))($x_1))($y_2);
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
$GLOBALS['Data_List_eqPattern'] = __NAMESPACE__ . '\\majData_majList_eqmajPattern';

// Data_List_ordPattern
function majData_majList_ordmajPattern($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_ordmajPattern';
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
  $__res = (((($GLOBALS['Data_List_Types_eq1List'])->{'eq1'})($__local_var_1_0))($x_2))($y_3);
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
  $__res = (((\Data\List\Types\majData_majList_majTypes_ordmajList($dictOrd_0))->{'compare'})($x_3))($y_4);
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
$GLOBALS['Data_List_ordPattern'] = __NAMESPACE__ . '\\majData_majList_ordmajPattern';

// Data_List_elemLastIndex
function majData_majList_elemmajLastmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_elemmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_findLastIndex'])(function($v_2) use ($dictEq_0, $x_1) {
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
$GLOBALS['Data_List_elemLastIndex'] = __NAMESPACE__ . '\\majData_majList_elemmajLastmajIndex';

// Data_List_elemIndex
function majData_majList_elemmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_elemmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($dictEq_0, &$go__2_0, $x_1) {
  $__fn = function($v_3, $v1_4 = null) use ($dictEq_0, &$go__2_0, $x_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__2_0_0_v_3 = $v_3;
  $__tco_var_go__2_0_0_v1_4 = $v1_4;
  tco_loop_go__2_0_0:;
  $v_3 = $__tco_var_go__2_0_0_v_3;
  $v1_4 = $__tco_var_go__2_0_0_v1_4;
  $__t0 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t3 = null;;
if (((($dictEq_0)->{'eq'})(($v1_4)->{'value0'}))($x_1)) {
$__t3 = new \Data\Maybe\Data_Maybe_Just($v_3);
goto end_branch_3;;
};
$__tco_1 = ($v_3 + 1);
$__tco_2 = ($v1_4)->{'value1'};
$__tco_var_go__2_0_0_v_3 = $__tco_1;
$__tco_var_go__2_0_0_v1_4 = $__tco_2;
goto tco_loop_go__2_0_0;;
$__t3 = null;
end_branch_3:;
$__t0 = $__t3;
goto end_branch_0;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
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
  $__res = ($go__2_0)(0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_elemIndex'] = __NAMESPACE__ . '\\majData_majList_elemmajIndex';

// Data_List_dropWhile
function majData_majList_dropmajWhile($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_dropmajWhile';
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
  if (($v_2 instanceof \Data\List\Types\Data_List_Types_Cons && ($p_0)(($v_2)->{'value0'}))) {
$__tco_1 = ($v_2)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_1;
goto tco_loop_go__1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  $__t0 = $v_2;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__1_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_dropWhile'] = __NAMESPACE__ . '\\majData_majList_dropmajWhile';

// Data_List_dropEnd
function majData_majList_dropmajEnd(int $n_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_dropmajEnd';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\List\majData_majList_take((\Data\List\majData_majList_length($xs_1) - $n_0), $xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_dropEnd'] = __NAMESPACE__ . '\\majData_majList_dropmajEnd';

// Data_List_drop
function majData_majList_drop(int $v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_drop';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_drop_v_0 = $v_0;
  $__tco_var_Data_List_drop_v1_1 = $v1_1;
  tco_loop_Data_List_drop:;
  $v_0 = $__tco_var_Data_List_drop_v_0;
  $v1_1 = $__tco_var_Data_List_drop_v1_1;
  $__t0 = null;;
  if ((($GLOBALS['Data_List_lessThan'])($v_0))(1)) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = ($v_0 - 1);
$__tco_2 = ($v1_1)->{'value1'};
$__tco_var_Data_List_drop_v_0 = $__tco_1;
$__tco_var_Data_List_drop_v1_1 = $__tco_2;
goto tco_loop_Data_List_drop;;
$__t0 = null;
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
$GLOBALS['Data_List_drop'] = __NAMESPACE__ . '\\majData_majList_drop';

// Data_List_slice
function majData_majList_slice(int $start_0, $end_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_slice';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Data\List\majData_majList_take(($end_1 - $start_0), \Data\List\majData_majList_drop($start_0, $xs_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_slice'] = __NAMESPACE__ . '\\majData_majList_slice';

// Data_List_takeEnd
function majData_majList_takemajEnd(int $n_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_takemajEnd';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\List\majData_majList_drop((\Data\List\majData_majList_length($xs_1) - $n_0), $xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_takeEnd'] = __NAMESPACE__ . '\\majData_majList_takemajEnd';

// Data_List_deleteBy
function majData_majList_deletemajBy($v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_deletemajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_deleteBy_v_0 = $v_0;
  $__tco_var_Data_List_deleteBy_v1_1 = $v1_1;
  $__tco_var_Data_List_deleteBy_v2_2 = $v2_2;
  tco_loop_Data_List_deleteBy:;
  $v_0 = $__tco_var_Data_List_deleteBy_v_0;
  $v1_1 = $__tco_var_Data_List_deleteBy_v1_1;
  $v2_2 = $__tco_var_Data_List_deleteBy_v2_2;
  $__t0 = null;;
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if ((($v_0)($v1_1))(($v2_2)->{'value0'})) {
$__t1 = ($v2_2)->{'value1'};
goto end_branch_1;;
};
$__t1 = new \Data\List\Types\Data_List_Types_Cons(($v2_2)->{'value0'}, \Data\List\majData_majList_deletemajBy($v_0, $v1_1, ($v2_2)->{'value1'}));
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
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_deleteBy'] = __NAMESPACE__ . '\\majData_majList_deletemajBy';

// Data_List_unionBy
function majData_majList_unionmajBy($eq2_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_unionmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_List_Types_semigroupList'])->{'append'})($xs_1))((((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() use ($eq2_0) {
  $__fn = function($b_3, $a_4 = null) use ($eq2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\List\majData_majList_deletemajBy($eq2_0, $a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(\Data\List\majData_majList_nubmajBymajEq($eq2_0, $ys_2)))($xs_1));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_unionBy'] = __NAMESPACE__ . '\\majData_majList_unionmajBy';

// Data_List_union
function majData_majList_union($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_unionBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_union'] = __NAMESPACE__ . '\\majData_majList_union';

// Data_List_deleteAt
function majData_majList_deletemajAt(int $v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_deletemajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_Data_List_deleteAt_v_0 = $v_0;
  $__tco_var_Data_List_deleteAt_v1_1 = $v1_1;
  tco_loop_Data_List_deleteAt:;
  $v_0 = $__tco_var_Data_List_deleteAt_v_0;
  $v1_1 = $__tco_var_Data_List_deleteAt_v1_1;
  $__t0 = null;;
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
switch ($v_0) {
case 0:
$__t2 = new \Data\Maybe\Data_Maybe_Just(($v1_1)->{'value1'});
goto end_branch_2;;
break;
default:
;
break;
};
$__local_var_2_1 = ($v1_1)->{'value0'};
$__t2 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v2_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($__local_var_2_1, $v2_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\majData_majList_deletemajAt(($v_0 - 1), ($v1_1)->{'value1'}));
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_deleteAt'] = __NAMESPACE__ . '\\majData_majList_deletemajAt';

// Data_List_delete
function majData_majList_delete($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_delete';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_deleteBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_delete'] = __NAMESPACE__ . '\\majData_majList_delete';

// Data_List_difference
function majData_majList_difference($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_difference';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() use ($dictEq_0) {
  $__fn = function($b_1, $a_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\List\majData_majList_deletemajBy(($dictEq_0)->{'eq'}, $a_2, $b_1);
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
$GLOBALS['Data_List_difference'] = __NAMESPACE__ . '\\majData_majList_difference';

// Data_List_concatMap
function majData_majList_concatmajMap($b_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_concatmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Types_bindList'])->{'bind'})($a_1))($b_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_concatMap'] = __NAMESPACE__ . '\\majData_majList_concatmajMap';

// Data_List_concat
function majData_majList_concat($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_concat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_List_Types_bindList'])->{'bind'})($v_0))($GLOBALS['Data_List_identity']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_concat'] = __NAMESPACE__ . '\\majData_majList_concat';

// Data_List_catMaybes_closure
$GLOBALS['Data_List_catMaybes_closure'] = (function() use (&$__fn) {
$go__0_0 = null;
$go__0_0 = (function() use (&$go__0_0) {
  $__fn = function($v_1, $v1_2 = null) use (&$go__0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__0_0_0_v_1 = $v_1;
  $__tco_var_go__0_0_0_v1_2 = $v1_2;
  tco_loop_go__0_0_0:;
  $v_1 = $__tco_var_go__0_0_0_v_1;
  $v1_2 = $__tco_var_go__0_0_0_v1_2;
  $__t0 = null;;
  if ($v1_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$go__3_1 = null;
$go__3_1 = (function() use (&$__tco_var_go__0_0_0_v_1, &$__tco_var_go__0_0_0_v1_2, &$go__3_1) {
  $__fn = function($v_4, $v1_5 = null) use (&$__tco_var_go__0_0_0_v_1, &$__tco_var_go__0_0_0_v1_2, &$go__3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__3_1_1_v_4 = $v_4;
  $__tco_var_go__3_1_1_v1_5 = $v1_5;
  tco_loop_go__3_1_1:;
  $v_4 = $__tco_var_go__3_1_1_v_4;
  $v1_5 = $__tco_var_go__3_1_1_v1_5;
  $__t1 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_4;
goto end_branch_1;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_3 = ($v1_5)->{'value1'};
$__tco_var_go__3_1_1_v_4 = $__tco_2;
$__tco_var_go__3_1_1_v1_5 = $__tco_3;
goto tco_loop_go__3_1_1;;
$__t1 = null;
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
$__t0 = (($go__3_1)(new \Data\List\Types\Data_List_Types_Nil()))($v_1);
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
if (($v1_2)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__tco_3 = $v_1;
$__tco_4 = ($v1_2)->{'value1'};
$__tco_var_go__0_0_0_v_1 = $__tco_3;
$__tco_var_go__0_0_0_v1_2 = $__tco_4;
goto tco_loop_go__0_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if (($v1_2)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons((($v1_2)->{'value0'})->{'value0'}, $v_1);
$__tco_6 = ($v1_2)->{'value1'};
$__tco_var_go__0_0_0_v_1 = $__tco_5;
$__tco_var_go__0_0_0_v1_2 = $__tco_6;
goto tco_loop_go__0_0_0;;
$__t2 = null;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
return ($go__0_0)(new \Data\List\Types\Data_List_Types_Nil());
})();

// Data_List_catMaybes
function majData_majList_catmajMaybes($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_catmajMaybes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_catMaybes_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_catMaybes'] = __NAMESPACE__ . '\\majData_majList_catmajMaybes';

// Data_List_alterAt
function majData_majList_altermajAt(int $v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_altermajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_Data_List_alterAt_v_0 = $v_0;
  $__tco_var_Data_List_alterAt_v1_1 = $v1_1;
  $__tco_var_Data_List_alterAt_v2_2 = $v2_2;
  tco_loop_Data_List_alterAt:;
  $v_0 = $__tco_var_Data_List_alterAt_v_0;
  $v1_1 = $__tco_var_Data_List_alterAt_v1_1;
  $v2_2 = $__tco_var_Data_List_alterAt_v2_2;
  $__t0 = null;;
  if ($v2_2 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
switch ($v_0) {
case 0:
$v3_3_3 = ($v1_1)(($v2_2)->{'value0'});
$__t4 = null;;
if ($v3_3_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = ($v2_2)->{'value1'};
goto end_branch_4;;
};
if ($v3_3_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t4 = new \Data\List\Types\Data_List_Types_Cons(($v3_3_3)->{'value0'}, ($v2_2)->{'value1'});
goto end_branch_4;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
end_branch_4:;
$__t2 = new \Data\Maybe\Data_Maybe_Just($__t4);
goto end_branch_2;;
break;
default:
;
break;
};
$__local_var_3_1 = ($v2_2)->{'value0'};
$__t2 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v3_4) use ($__local_var_3_1) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($__local_var_3_1, $v3_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\majData_majList_altermajAt(($v_0 - 1), $v1_1, ($v2_2)->{'value1'}));
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_alterAt'] = __NAMESPACE__ . '\\majData_majList_altermajAt';

// Data_List_modifyAt
function majData_majList_modifymajAt(int $n_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_modifymajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_alterAt'])($n_0))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($f_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_modifyAt'] = __NAMESPACE__ . '\\majData_majList_modifymajAt';

