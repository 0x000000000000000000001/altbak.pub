<?php

namespace Data\List\NonEmpty;

// ALL IMPORTS: Control.Bind, Control.Category, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Bind, Control.Category, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.List, Data.List.NonEmpty, Data.List.Types, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.NonEmpty/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semigroup.Traversable/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
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




// Data_List_NonEmpty_identity
function majData_majList_majNonmajEmpty_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_identity'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_identity';

// Data_List_NonEmpty_zipWith
function majData_majList_majNonmajEmpty_zipmajWith($f_0, $v_1 = null, $v1_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_zipmajWith';
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
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((($f_0)(($v_1)->{'value0'}))(($v1_2)->{'value0'}), (($go__4_1)(new \Data\List\Types\Data_List_Types_Nil()))(((($go__3_0)(($v_1)->{'value1'}))(($v1_2)->{'value1'}))(new \Data\List\Types\Data_List_Types_Nil())));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_zipWith'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_zipmajWith';

// Data_List_NonEmpty_zipWithA
function majData_majList_majNonmajEmpty_zipmajWithmajA($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_zipmajWithmajA';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequence11_1_0 = (($GLOBALS['Data_List_Types_traversable1NonEmptyList'])->{'sequence1'})((($dictApplicative_0)->{'Apply0'})(null));
  $__res = (function() use ($sequence11_1_0) {
  $__fn = function($f_2, $xs_3 = null, $ys_4 = null) use ($sequence11_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($sequence11_1_0)(\Data\List\NonEmpty\majData_majList_majNonmajEmpty_zipmajWith($f_2, $xs_3, $ys_4));
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
$GLOBALS['Data_List_NonEmpty_zipWithA'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_zipmajWithmajA';

// Data_List_NonEmpty_zip_closure
$GLOBALS['Data_List_NonEmpty_zip_closure'] = ($GLOBALS['Data_List_NonEmpty_zipWith'])($GLOBALS['Data_Tuple_Tuple']);

// Data_List_NonEmpty_zip
function majData_majList_majNonmajEmpty_zip($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_zip';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_zip_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_zip'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_zip';

// Data_List_NonEmpty_wrappedOperation2
function majData_majList_majNonmajEmpty_wrappedmajOperation2(string $name_0, $f_1 = null, $v_2 = null, $v1_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_wrappedmajOperation2';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $v2_4_0 = (($f_1)(new \Data\List\Types\Data_List_Types_Cons(($v_2)->{'value0'}, ($v_2)->{'value1'})))(new \Data\List\Types\Data_List_Types_Cons(($v1_3)->{'value0'}, ($v1_3)->{'value1'}));
  $__t1 = null;;
  if ($v2_4_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v2_4_0)->{'value0'}, ($v2_4_0)->{'value1'});
goto end_branch_1;;
};
  if ($v2_4_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = \Partial\majPartial__crashmajWith(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("Impossible: empty list in NonEmptyList "))($name_0));
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
$GLOBALS['Data_List_NonEmpty_wrappedOperation2'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_wrappedmajOperation2';

// Data_List_NonEmpty_wrappedOperation
function majData_majList_majNonmajEmpty_wrappedmajOperation(string $name_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_wrappedmajOperation';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v1_3_0 = ($f_1)(new \Data\List\Types\Data_List_Types_Cons(($v_2)->{'value0'}, ($v_2)->{'value1'}));
  $__t1 = null;;
  if ($v1_3_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v1_3_0)->{'value0'}, ($v1_3_0)->{'value1'});
goto end_branch_1;;
};
  if ($v1_3_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = \Partial\majPartial__crashmajWith(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("Impossible: empty list in NonEmptyList "))($name_0));
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_wrappedOperation'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_wrappedmajOperation';

// Data_List_NonEmpty_updateAt
function majData_majList_majNonmajEmpty_updatemajAt(int $i_0, $a_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_updatemajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  switch ($i_0) {
case 0:
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_1, ($v_2)->{'value1'}));
goto end_branch_1;;
break;
default:
;
break;
};
  $__local_var_3_0 = ($v_2)->{'value0'};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($v1_4) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_3_0, $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(\Data\List\majData_majList_updatemajAt(($i_0 - 1), $a_1, ($v_2)->{'value1'}));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_updateAt'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_updatemajAt';

// Data_List_NonEmpty_unzip
function majData_majList_majNonmajEmpty_unzip($ts_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_unzip';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($GLOBALS['Data_List_Types_functorNonEmptyList'])->{'map'})($GLOBALS['Data_Tuple_fst']))($ts_0), ((($GLOBALS['Data_List_Types_functorNonEmptyList'])->{'map'})($GLOBALS['Data_Tuple_snd']))($ts_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_unzip'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_unzip';

// Data_List_NonEmpty_unsnoc
function majData_majList_majNonmajEmpty_unsnoc($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_unsnoc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $v1_1_0 = \Data\List\majData_majList_unsnoc(($v_0)->{'value1'});
  $__t1 = null;;
  if ($v1_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (object)["init" => new \Data\List\Types\Data_List_Types_Nil(), "last" => ($v_0)->{'value0'}];
goto end_branch_1;;
};
  if ($v1_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (object)["init" => new \Data\List\Types\Data_List_Types_Cons(($v_0)->{'value0'}, (($v1_1_0)->{'value0'})->{'init'}), "last" => (($v1_1_0)->{'value0'})->{'last'}];
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
$GLOBALS['Data_List_NonEmpty_unsnoc'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_unsnoc';

// Data_List_NonEmpty_unionBy_closure
$GLOBALS['Data_List_NonEmpty_unionBy_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("unionBy")))($GLOBALS['Data_List_unionBy']);

// Data_List_NonEmpty_unionBy
function majData_majList_majNonmajEmpty_unionmajBy($v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_unionmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_unionBy_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_unionBy'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_unionmajBy';

// Data_List_NonEmpty_union
function majData_majList_majNonmajEmpty_union($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("union"))(($GLOBALS['Data_List_union'])($dictEq_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_union'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_union';

// Data_List_NonEmpty_uncons
function majData_majList_majNonmajEmpty_uncons($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["head" => ($v_0)->{'value0'}, "tail" => ($v_0)->{'value1'}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_uncons'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_uncons';

// Data_List_NonEmpty_toList
function majData_majList_majNonmajEmpty_tomajList($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_tomajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons(($v_0)->{'value0'}, ($v_0)->{'value1'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_toList'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_tomajList';

// Data_List_NonEmpty_toUnfoldable
function majData_majList_majNonmajEmpty_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($dictUnfoldable_0)->{'unfoldr'})(function($xs_1) {
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
})))($GLOBALS['Data_List_NonEmpty_toList']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_toUnfoldable'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_tomajUnfoldable';

// Data_List_NonEmpty_tail
function majData_majList_majNonmajEmpty_tail($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_tail'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_tail';

// Data_List_NonEmpty_sortBy_closure
$GLOBALS['Data_List_NonEmpty_sortBy_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("sortBy")))($GLOBALS['Data_List_sortBy']);

// Data_List_NonEmpty_sortBy
function majData_majList_majNonmajEmpty_sortmajBy($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_sortmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_sortBy_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_sortBy'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_sortmajBy';

// Data_List_NonEmpty_sort
function majData_majList_majNonmajEmpty_sort($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_sort';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = \Data\List\NonEmpty\majData_majList_majNonmajEmpty_sortmajBy($compare_1_0, $xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_sort'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_sort';

// Data_List_NonEmpty_snoc
function majData_majList_majNonmajEmpty_snoc($v_0, $y_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_snoc';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_0)->{'value0'}, (((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})($GLOBALS['Data_List_Types_Cons']))(new \Data\List\Types\Data_List_Types_Cons($y_1, new \Data\List\Types\Data_List_Types_Nil())))(($v_0)->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_snoc'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_snoc';

// Data_List_NonEmpty_singleton_closure
$GLOBALS['Data_List_NonEmpty_singleton_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_0, ($GLOBALS['Data_List_Types_plusList'])->{'empty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_List_NonEmpty_singleton
function majData_majList_majNonmajEmpty_singleton($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_singleton_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_singleton'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_singleton';

// Data_List_NonEmpty_snoc'
function majData_majList_majNonmajEmpty_snoc__prime__($v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_snoc__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_0)->{'value0'}, (((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})($GLOBALS['Data_List_Types_Cons']))(new \Data\List\Types\Data_List_Types_Cons($v1_1, new \Data\List\Types\Data_List_Types_Nil())))(($v_0)->{'value1'}));
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = \Data\List\NonEmpty\majData_majList_majNonmajEmpty_singleton($v1_1);
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
$GLOBALS['Data_List_NonEmpty_snoc__prime__'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_snoc__prime__';

// Data_List_NonEmpty_reverse_closure
$GLOBALS['Data_List_NonEmpty_reverse_closure'] = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("reverse"))($GLOBALS['Data_List_reverse']);

// Data_List_NonEmpty_reverse
function majData_majList_majNonmajEmpty_reverse($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_reverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_reverse_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_reverse'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_reverse';

// Data_List_NonEmpty_nubEq
function majData_majList_majNonmajEmpty_nubmajEq($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_nubmajEq';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nubEq"))(($GLOBALS['Data_List_nubEq'])($dictEq_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_nubEq'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_nubmajEq';

// Data_List_NonEmpty_nubByEq_closure
$GLOBALS['Data_List_NonEmpty_nubByEq_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nubByEq")))($GLOBALS['Data_List_nubByEq']);

// Data_List_NonEmpty_nubByEq
function majData_majList_majNonmajEmpty_nubmajBymajEq($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_nubmajBymajEq';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_nubByEq_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_nubByEq'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_nubmajBymajEq';

// Data_List_NonEmpty_nubBy_closure
$GLOBALS['Data_List_NonEmpty_nubBy_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nubBy")))($GLOBALS['Data_List_nubBy']);

// Data_List_NonEmpty_nubBy
function majData_majList_majNonmajEmpty_nubmajBy($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_nubmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_nubBy_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_nubBy'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_nubmajBy';

// Data_List_NonEmpty_nub
function majData_majList_majNonmajEmpty_nub($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_nub';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("nub"))(($GLOBALS['Data_List_nubBy'])(($dictOrd_0)->{'compare'}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_nub'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_nub';

// Data_List_NonEmpty_modifyAt
function majData_majList_majNonmajEmpty_modifymajAt(int $i_0, $f_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_modifymajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  switch ($i_0) {
case 0:
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_1)(($v_2)->{'value0'}), ($v_2)->{'value1'}));
goto end_branch_1;;
break;
default:
;
break;
};
  $__local_var_3_0 = ($v_2)->{'value0'};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($v1_4) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_3_0, $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(\Data\List\majData_majList_altermajAt(($i_0 - 1), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Maybe_Just']))($f_1), ($v_2)->{'value1'}));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_modifyAt'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_modifymajAt';

// Data_List_NonEmpty_lift
function majData_majList_majNonmajEmpty_lift($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_lift';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_0)(new \Data\List\Types\Data_List_Types_Cons(($v_1)->{'value0'}, ($v_1)->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_lift'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_lift';

// Data_List_NonEmpty_mapMaybe_closure
$GLOBALS['Data_List_NonEmpty_mapMaybe_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_mapMaybe']);

// Data_List_NonEmpty_mapMaybe
function majData_majList_majNonmajEmpty_mapmajMaybe($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_mapmajMaybe';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_mapMaybe_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_mapMaybe'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_mapmajMaybe';

// Data_List_NonEmpty_partition_closure
$GLOBALS['Data_List_NonEmpty_partition_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_partition']);

// Data_List_NonEmpty_partition
function majData_majList_majNonmajEmpty_partition($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_partition';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_partition_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_partition'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_partition';

// Data_List_NonEmpty_span_closure
$GLOBALS['Data_List_NonEmpty_span_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_span']);

// Data_List_NonEmpty_span
function majData_majList_majNonmajEmpty_span($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_span';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_span_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_span'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_span';

// Data_List_NonEmpty_take_closure
$GLOBALS['Data_List_NonEmpty_take_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_take']);

// Data_List_NonEmpty_take
function majData_majList_majNonmajEmpty_take(int $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_take';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_take_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_take'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_take';

// Data_List_NonEmpty_takeWhile_closure
$GLOBALS['Data_List_NonEmpty_takeWhile_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_takeWhile']);

// Data_List_NonEmpty_takeWhile
function majData_majList_majNonmajEmpty_takemajWhile($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_takemajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_takeWhile_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_takeWhile'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_takemajWhile';

// Data_List_NonEmpty_length
function majData_majList_majNonmajEmpty_length($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (1 + \Data\List\majData_majList_length(($v_0)->{'value1'}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_length'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_length';

// Data_List_NonEmpty_last
function majData_majList_majNonmajEmpty_last($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_last';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if (($v_0)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t1 = null;;
if ((($v_0)->{'value1'})->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = (($v_0)->{'value1'})->{'value0'};
goto end_branch_1;;
};
if (\Data\List\majData_majList_last((($v_0)->{'value1'})->{'value1'}) instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = ($v_0)->{'value0'};
goto end_branch_1;;
};
if (\Data\List\majData_majList_last((($v_0)->{'value1'})->{'value1'}) instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (\Data\List\majData_majList_last((($v_0)->{'value1'})->{'value1'}))->{'value0'};
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  $__t0 = ($v_0)->{'value0'};
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_last'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_last';

// Data_List_NonEmpty_intersectBy_closure
$GLOBALS['Data_List_NonEmpty_intersectBy_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("intersectBy")))($GLOBALS['Data_List_intersectBy']);

// Data_List_NonEmpty_intersectBy
function majData_majList_majNonmajEmpty_intersectmajBy($v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_intersectmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_intersectBy_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_intersectBy'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_intersectmajBy';

// Data_List_NonEmpty_intersect
function majData_majList_majNonmajEmpty_intersect($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_intersect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation2'])("intersect"))(($GLOBALS['Data_List_intersect'])($dictEq_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_intersect'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_intersect';

// Data_List_NonEmpty_insertAt
function majData_majList_majNonmajEmpty_insertmajAt(int $i_0, $a_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_insertmajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  switch ($i_0) {
case 0:
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_1, new \Data\List\Types\Data_List_Types_Cons(($v_2)->{'value0'}, ($v_2)->{'value1'})));
goto end_branch_1;;
break;
default:
;
break;
};
  $__local_var_3_0 = ($v_2)->{'value0'};
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($v1_4) use ($__local_var_3_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($__local_var_3_0, $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(\Data\List\majData_majList_insertmajAt(($i_0 - 1), $a_1, ($v_2)->{'value1'}));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_insertAt'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_insertmajAt';

// Data_List_NonEmpty_init
function majData_majList_majNonmajEmpty_init($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_init';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v_1) {
  $__num = \func_num_args();
  $__res = ($v_1)->{'init'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Data\List\majData_majList_unsnoc(($v_0)->{'value1'}));
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_1;;
};
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\List\Types\Data_List_Types_Cons(($v_0)->{'value0'}, ($__local_var_1_0)->{'value0'});
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
$GLOBALS['Data_List_NonEmpty_init'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_init';

// Data_List_NonEmpty_index
function majData_majList_majNonmajEmpty_index($v_0, $i_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_index';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = match ($i_1) { 0 => new \Data\Maybe\Data_Maybe_Just(($v_0)->{'value0'}), default => \Data\List\majData_majList_index(($v_0)->{'value1'}, ($i_1 - 1)) };
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_index'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_index';

// Data_List_NonEmpty_head
function majData_majList_majNonmajEmpty_head($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_head'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_head';

// Data_List_NonEmpty_groupBy_closure
$GLOBALS['Data_List_NonEmpty_groupBy_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("groupBy")))($GLOBALS['Data_List_groupBy']);

// Data_List_NonEmpty_groupBy
function majData_majList_majNonmajEmpty_groupmajBy($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_groupmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_groupBy_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_groupBy'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_groupmajBy';

// Data_List_NonEmpty_groupAllBy_closure
$GLOBALS['Data_List_NonEmpty_groupAllBy_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("groupAllBy")))($GLOBALS['Data_List_groupAllBy']);

// Data_List_NonEmpty_groupAllBy
function majData_majList_majNonmajEmpty_groupmajAllmajBy($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_groupmajAllmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_groupAllBy_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_groupAllBy'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_groupmajAllmajBy';

// Data_List_NonEmpty_groupAll
function majData_majList_majNonmajEmpty_groupmajAll($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_groupmajAll';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("groupAll"))(($GLOBALS['Data_List_groupAll'])($dictOrd_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_groupAll'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_groupmajAll';

// Data_List_NonEmpty_group
function majData_majList_majNonmajEmpty_group($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_group';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Data_List_NonEmpty_wrappedOperation'])("group"))(($GLOBALS['Data_List_group'])($dictEq_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_group'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_group';

// Data_List_NonEmpty_fromList
function majData_majList_majNonmajEmpty_frommajList($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_frommajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_0)->{'value0'}, ($v_0)->{'value1'}));
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
$GLOBALS['Data_List_NonEmpty_fromList'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_frommajList';

// Data_List_NonEmpty_fromFoldable
function majData_majList_majNonmajEmpty_frommajFoldable($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_frommajFoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_fromList']))(((($dictFoldable_0)->{'foldr'})($GLOBALS['Data_List_Types_Cons']))(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_fromFoldable'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_frommajFoldable';

// Data_List_NonEmpty_foldM
function majData_majList_majNonmajEmpty_foldmajM($dictMonad_0, $f_1 = null, $b_2 = null, $v_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_foldmajM';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__local_var_4_0 = ($v_3)->{'value1'};
  $__res = ((((($dictMonad_0)->{'Bind1'})(null))->{'bind'})((($f_1)($b_2))(($v_3)->{'value0'})))(function($b_prime_5) use ($__local_var_4_0, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = \Data\List\majData_majList_foldmajM($dictMonad_0, $f_1, $b_prime_5, $__local_var_4_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_foldM'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_foldmajM';

// Data_List_NonEmpty_findLastIndex
function majData_majList_majNonmajEmpty_findmajLastmajIndex($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_findmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = \Data\List\majData_majList_findmajLastmajIndex($f_0, ($v_1)->{'value1'});
  $__t1 = null;;
  if ($v1_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just((($v1_2_0)->{'value0'} + 1));
goto end_branch_1;;
};
  if ($v1_2_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = null;;
if (($f_0)(($v_1)->{'value0'})) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(0);
goto end_branch_2;;
};
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_findLastIndex'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_findmajLastmajIndex';

// Data_List_NonEmpty_findIndex
function majData_majList_majNonmajEmpty_findmajIndex($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_findmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if (($f_0)(($v_1)->{'value0'})) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(0);
goto end_branch_1;;
};
  $go__2_0 = null;
  $go__2_0 = (function() use ($f_0, &$go__2_0) {
  $__fn = function($v_3, $v1_4 = null) use ($f_0, &$go__2_0, &$__fn) {
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
if (($f_0)(($v1_4)->{'value0'})) {
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
  $__t1 = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($v1_2) {
  $__num = \func_num_args();
  $__res = ($v1_2 + 1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($go__2_0)(0))(($v_1)->{'value1'}));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_findIndex'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_findmajIndex';

// Data_List_NonEmpty_filterM
function majData_majList_majNonmajEmpty_filtermajM($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_filtermajM';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))(($GLOBALS['Data_List_filterM'])($dictMonad_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_filterM'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_filtermajM';

// Data_List_NonEmpty_filter_closure
$GLOBALS['Data_List_NonEmpty_filter_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_filter']);

// Data_List_NonEmpty_filter
function majData_majList_majNonmajEmpty_filter($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_filter';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_filter_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_filter'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_filter';

// Data_List_NonEmpty_elemLastIndex
function majData_majList_majNonmajEmpty_elemmajLastmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_elemmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_findLastIndex'])(function($v_2) use ($dictEq_0, $x_1) {
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
$GLOBALS['Data_List_NonEmpty_elemLastIndex'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_elemmajLastmajIndex';

// Data_List_NonEmpty_elemIndex
function majData_majList_majNonmajEmpty_elemmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_elemmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_findIndex'])(function($v_2) use ($dictEq_0, $x_1) {
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
$GLOBALS['Data_List_NonEmpty_elemIndex'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_elemmajIndex';

// Data_List_NonEmpty_dropWhile_closure
$GLOBALS['Data_List_NonEmpty_dropWhile_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_dropWhile']);

// Data_List_NonEmpty_dropWhile
function majData_majList_majNonmajEmpty_dropmajWhile($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_dropmajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_dropWhile_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_dropWhile'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_dropmajWhile';

// Data_List_NonEmpty_drop_closure
$GLOBALS['Data_List_NonEmpty_drop_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_NonEmpty_lift']))($GLOBALS['Data_List_drop']);

// Data_List_NonEmpty_drop
function majData_majList_majNonmajEmpty_drop(int $v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_drop';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_List_NonEmpty_drop_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_drop'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_drop';

// Data_List_NonEmpty_cons'
function majData_majList_majNonmajEmpty_cons__prime__($x_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_cons__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($x_0, $xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_cons__prime__'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_cons__prime__';

// Data_List_NonEmpty_cons
function majData_majList_majNonmajEmpty_cons($y_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_cons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($y_0, new \Data\List\Types\Data_List_Types_Cons(($v_1)->{'value0'}, ($v_1)->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_cons'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_cons';

// Data_List_NonEmpty_concatMap
function majData_majList_majNonmajEmpty_concatmajMap($b_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_concatmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_List_Types_bindNonEmptyList'])->{'bind'})($a_1))($b_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_concatMap'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_concatmajMap';

// Data_List_NonEmpty_concat
function majData_majList_majNonmajEmpty_concat($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_concat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_List_Types_bindNonEmptyList'])->{'bind'})($v_0))($GLOBALS['Data_List_NonEmpty_identity']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_concat'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_concat';

// Data_List_NonEmpty_catMaybes
function majData_majList_majNonmajEmpty_catmajMaybes($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_catmajMaybes';
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
$__t2 = null;;
if (($v1_3)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__tco_3 = $v_2;
$__tco_4 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_3;
$__tco_var_go__1_0_0_v1_3 = $__tco_4;
goto tco_loop_go__1_0_0;;
$__t2 = null;
goto end_branch_2;;
};
if (($v1_3)->{'value0'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons((($v1_3)->{'value0'})->{'value0'}, $v_2);
$__tco_6 = ($v1_3)->{'value1'};
$__tco_var_go__1_0_0_v_2 = $__tco_5;
$__tco_var_go__1_0_0_v1_3 = $__tco_6;
goto tco_loop_go__1_0_0;;
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
  $__res = (($go__1_0)(new \Data\List\Types\Data_List_Types_Nil()))(new \Data\List\Types\Data_List_Types_Cons(($v_0)->{'value0'}, ($v_0)->{'value1'}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_NonEmpty_catMaybes'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_catmajMaybes';

// Data_List_NonEmpty_appendFoldable
function majData_majList_majNonmajEmpty_appendmajFoldable($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majNonmajEmpty_appendmajFoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $fromFoldable1_1_0 = ((($dictFoldable_0)->{'foldr'})($GLOBALS['Data_List_Types_Cons']))(new \Data\List\Types\Data_List_Types_Nil());
  $__res = (function() use ($fromFoldable1_1_0) {
  $__fn = function($v_2, $ys_3 = null) use ($fromFoldable1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_2)->{'value0'}, ((($GLOBALS['Data_List_Types_semigroupList'])->{'append'})(($v_2)->{'value1'}))(($fromFoldable1_1_0)($ys_3)));
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
$GLOBALS['Data_List_NonEmpty_appendFoldable'] = __NAMESPACE__ . '\\majData_majList_majNonmajEmpty_appendmajFoldable';

