<?php

namespace Data\List\Types;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Comonad, Control.Extend, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List.Types, Data.Maybe, Data.Monoid, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Comonad, Control.Extend, Control.Monad, Control.MonadPlus, Control.Plus, Control.Semigroupoid, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List.Types, Data.Maybe, Data.Monoid, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.MonadPlus/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semigroup.Traversable/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
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


final class Data_List_Types_Nil { public $tag = 'Nil'; public function __construct() {} }
final class Data_List_Types_Cons { public $tag = 'Cons'; public function __construct(public  $value0, public  $value1) {} }

// Data_List_Types_Nil
$GLOBALS['Data_List_Types_Nil'] = ($GLOBALS['__phpurs_data0_Nil'] ??= new \Data\List\Types\Data_List_Types_Nil());

// Data_List_Types_Cons
$GLOBALS['Data_List_Types_Cons'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_List_Types_NonEmptyList
function majData_majList_majTypes_majNonmajEmptymajList($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_majNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_NonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majTypes_majNonmajEmptymajList';

// Data_List_Types_toList
function majData_majList_majTypes_tomajList($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_tomajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons(($v_0)->{'value0'}, ($v_0)->{'value1'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_toList'] = __NAMESPACE__ . '\\majData_majList_majTypes_tomajList';

// Data_List_Types_newtypeNonEmptyList
$GLOBALS['Data_List_Types_newtypeNonEmptyList'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_nelCons
function majData_majList_majTypes_nelmajCons($a_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_nelmajCons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_0, new \Data\List\Types\Data_List_Types_Cons(($v_1)->{'value0'}, ($v_1)->{'value1'}));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_List_Types_nelCons'] = __NAMESPACE__ . '\\majData_majList_majTypes_nelmajCons';

// Data_List_Types_listMap
function majData_majList_majTypes_listmajMap($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_listmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $chunkedRevMap_1_0 = null;
  $chunkedRevMap_1_0 = (function() use (&$chunkedRevMap_1_0, $f_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$chunkedRevMap_1_0, $f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_chunkedRevMap_1_0_0_v_2 = $v_2;
  $__tco_var_chunkedRevMap_1_0_0_v1_3 = $v1_3;
  tco_loop_chunkedRevMap_1_0_0:;
  $v_2 = $__tco_var_chunkedRevMap_1_0_0_v_2;
  $v1_3 = $__tco_var_chunkedRevMap_1_0_0_v1_3;
  $__t4 = null;;
  if (($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons && (($v1_3)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons && (($v1_3)->{'value1'})->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons))) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons($v1_3, $v_2);
$__tco_6 = ((($v1_3)->{'value1'})->{'value1'})->{'value1'};
$__tco_var_chunkedRevMap_1_0_0_v_2 = $__tco_5;
$__tco_var_chunkedRevMap_1_0_0_v1_3 = $__tco_6;
goto tco_loop_chunkedRevMap_1_0_0;;
$__t4 = null;
goto end_branch_4;;
};
  $reverseUnrolledMap_4_0 = null;
  $reverseUnrolledMap_4_0 = (function() use (&$__tco_var_chunkedRevMap_1_0_0_v_2, &$__tco_var_chunkedRevMap_1_0_0_v1_3, $f_0, &$reverseUnrolledMap_4_0) {
  $__fn = function($v2_5, $v3_6 = null) use (&$__tco_var_chunkedRevMap_1_0_0_v_2, &$__tco_var_chunkedRevMap_1_0_0_v1_3, $f_0, &$reverseUnrolledMap_4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_reverseUnrolledMap_4_0_0_v2_5 = $v2_5;
  $__tco_var_reverseUnrolledMap_4_0_0_v3_6 = $v3_6;
  tco_loop_reverseUnrolledMap_4_0_0:;
  $v2_5 = $__tco_var_reverseUnrolledMap_4_0_0_v2_5;
  $v3_6 = $__tco_var_reverseUnrolledMap_4_0_0_v3_6;
  $__t0 = null;;
  if (($v2_5 instanceof \Data\List\Types\Data_List_Types_Cons && (($v2_5)->{'value0'} instanceof \Data\List\Types\Data_List_Types_Cons && ((($v2_5)->{'value0'})->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons && ((($v2_5)->{'value0'})->{'value1'})->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons)))) {
$__tco_1 = ($v2_5)->{'value1'};
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($f_0)((($v2_5)->{'value0'})->{'value0'}), new \Data\List\Types\Data_List_Types_Cons(($f_0)(((($v2_5)->{'value0'})->{'value1'})->{'value0'}), new \Data\List\Types\Data_List_Types_Cons(($f_0)((((($v2_5)->{'value0'})->{'value1'})->{'value1'})->{'value0'}), $v3_6)));
$__tco_var_reverseUnrolledMap_4_0_0_v2_5 = $__tco_1;
$__tco_var_reverseUnrolledMap_4_0_0_v3_6 = $__tco_2;
goto tco_loop_reverseUnrolledMap_4_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  $__t0 = $v3_6;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__t1 = null;;
  if ($v1_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
if (($v1_3)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t3 = null;;
if ((($v1_3)->{'value1'})->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = new \Data\List\Types\Data_List_Types_Cons(($f_0)(($v1_3)->{'value0'}), new \Data\List\Types\Data_List_Types_Cons(($f_0)((($v1_3)->{'value1'})->{'value0'}), new \Data\List\Types\Data_List_Types_Nil()));
goto end_branch_3;;
};
$__t3 = new \Data\List\Types\Data_List_Types_Nil();
end_branch_3:;
$__t2 = $__t3;
goto end_branch_2;;
};
if (($v1_3)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = new \Data\List\Types\Data_List_Types_Cons(($f_0)(($v1_3)->{'value0'}), new \Data\List\Types\Data_List_Types_Nil());
goto end_branch_2;;
};
$__t2 = new \Data\List\Types\Data_List_Types_Nil();
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  $__t1 = new \Data\List\Types\Data_List_Types_Nil();
  end_branch_1:;
  $__t4 = (($reverseUnrolledMap_4_0)($v_2))($__t1);
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($chunkedRevMap_1_0)(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_listMap'] = __NAMESPACE__ . '\\majData_majList_majTypes_listmajMap';

// Data_List_Types_functorList
$GLOBALS['Data_List_Types_functorList'] = (object)["map" => $GLOBALS['Data_List_Types_listMap']];

// Data_List_Types_functorNonEmptyList
$GLOBALS['Data_List_Types_functorNonEmptyList'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($m_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)(($m_1)->{'value0'}), \Data\List\Types\majData_majList_majTypes_listmajMap($f_0, ($m_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_foldableList
$GLOBALS['Data_List_Types_foldableList'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_v_3 = $v_3;
  $__tco_var_go__go_2_0_0_v1_4 = $v1_4;
  tco_loop_go__go_2_0_0:;
  $v_3 = $__tco_var_go__go_2_0_0_v_3;
  $v1_4 = $__tco_var_go__go_2_0_0_v1_4;
  $__t0 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v_3;
goto end_branch_0;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_2 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_0_0_v_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v1_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($b_2) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($b_2, $f_0) {
  $__num = \func_num_args();
  $__res = (($f_0)($a_3))($b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_1)))(($go__go_2_0)(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $go__go_1_1 = null;
  $go__go_1_1 = (function() use ($f_0, &$go__go_1_1) {
  $__fn = function($b_2, $v_3 = null) use ($f_0, &$go__go_1_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_1_1_b_2 = $b_2;
  $__tco_var_go__go_1_1_1_v_3 = $v_3;
  tco_loop_go__go_1_1_1:;
  $b_2 = $__tco_var_go__go_1_1_1_b_2;
  $v_3 = $__tco_var_go__go_1_1_1_v_3;
  $__t1 = null;;
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_2;
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = (($f_0)($b_2))(($v_3)->{'value0'});
$__tco_3 = ($v_3)->{'value1'};
$__tco_var_go__go_1_1_1_b_2 = $__tco_2;
$__tco_var_go__go_1_1_1_v_3 = $__tco_3;
goto tco_loop_go__go_1_1_1;;
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
  $__res = $go__go_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_3 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_2, $mempty_2_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($acc_4) use ($Semigroup0_1_2, $f_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Semigroup0_1_2)->{'append'})($acc_4)))($f_3);
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

// Data_List_Types_foldableNonEmptyList
$GLOBALS['Data_List_Types_foldableNonEmptyList'] = (object)["foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($f_2) use ($Semigroup0_1_0, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Semigroup0_1_0, $dictMonoid_0, $f_2) {
  $__num = \func_num_args();
  $Semigroup0_4_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $go__go_5_2 = null;
  $go__go_5_2 = (function() use ($Semigroup0_4_1, $f_2, &$go__go_5_2) {
  $__fn = function($b_6, $v_7 = null) use ($Semigroup0_4_1, $f_2, &$go__go_5_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_2_2_b_6 = $b_6;
  $__tco_var_go__go_5_2_2_v_7 = $v_7;
  tco_loop_go__go_5_2_2:;
  $b_6 = $__tco_var_go__go_5_2_2_b_6;
  $v_7 = $__tco_var_go__go_5_2_2_v_7;
  $__t2 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_6;
goto end_branch_2;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_4_1)->{'append'})($b_6), $f_2, ($v_7)->{'value0'});
$__tco_4 = ($v_7)->{'value1'};
$__tco_var_go__go_5_2_2_b_6 = $__tco_3;
$__tco_var_go__go_5_2_2_v_7 = $__tco_4;
goto tco_loop_go__go_5_2_2;;
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
  $__res = ((($Semigroup0_1_0)->{'append'})(($f_2)(($v_3)->{'value0'})))((($go__go_5_2)(($dictMonoid_0)->{'mempty'}))(($v_3)->{'value1'}));
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
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_3 = null;
  $go__go_3_3 = (function() use ($f_0, &$go__go_3_3) {
  $__fn = function($b_4, $v_5 = null) use ($f_0, &$go__go_3_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_3_3_b_4 = $b_4;
  $__tco_var_go__go_3_3_3_v_5 = $v_5;
  tco_loop_go__go_3_3_3:;
  $b_4 = $__tco_var_go__go_3_3_3_b_4;
  $v_5 = $__tco_var_go__go_3_3_3_v_5;
  $__t3 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_4;
goto end_branch_3;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = (($f_0)($b_4))(($v_5)->{'value0'});
$__tco_5 = ($v_5)->{'value1'};
$__tco_var_go__go_3_3_3_b_4 = $__tco_4;
$__tco_var_go__go_3_3_3_v_5 = $__tco_5;
goto tco_loop_go__go_3_3_3;;
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
  $__res = (($go__go_3_3)((($f_0)($b_1))(($v_2)->{'value0'})))(($v_2)->{'value1'});
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
}, "foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_4 = null;
  $go__go_3_4 = (function() use ($f_0, &$go__go_3_4) {
  $__fn = function($b_4, $v_5 = null) use ($f_0, &$go__go_3_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_4_4_b_4 = $b_4;
  $__tco_var_go__go_3_4_4_v_5 = $v_5;
  tco_loop_go__go_3_4_4:;
  $b_4 = $__tco_var_go__go_3_4_4_b_4;
  $v_5 = $__tco_var_go__go_3_4_4_v_5;
  $__t4 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $b_4;
goto end_branch_4;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = (($f_0)(($v_5)->{'value0'}))($b_4);
$__tco_6 = ($v_5)->{'value1'};
$__tco_var_go__go_3_4_4_b_4 = $__tco_5;
$__tco_var_go__go_3_4_4_v_5 = $__tco_6;
goto tco_loop_go__go_3_4_4;;
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
  $go__go_3_5 = null;
  $go__go_3_5 = (function() use (&$go__go_3_5) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__go_3_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_5_5_v_4 = $v_4;
  $__tco_var_go__go_3_5_5_v1_5 = $v1_5;
  tco_loop_go__go_3_5_5:;
  $v_4 = $__tco_var_go__go_3_5_5_v_4;
  $v1_5 = $__tco_var_go__go_3_5_5_v1_5;
  $__t5 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $v_4;
goto end_branch_5;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_7 = ($v1_5)->{'value1'};
$__tco_var_go__go_3_5_5_v_4 = $__tco_6;
$__tco_var_go__go_3_5_5_v1_5 = $__tco_7;
goto tco_loop_go__go_3_5_5;;
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
  $__res = (($f_0)(($v_2)->{'value0'}))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_3_4)($b_1), ($go__go_3_5)(new \Data\List\Types\Data_List_Types_Nil()), ($v_2)->{'value1'}));
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

// Data_List_Types_foldableWithIndexList
$GLOBALS['Data_List_Types_foldableWithIndexList'] = (object)["foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($xs_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = (function() use (&$go__go_3_0) {
  $__fn = function($b_4, $v_5 = null) use (&$go__go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_0_0_b_4 = $b_4;
  $__tco_var_go__go_3_0_0_v_5 = $v_5;
  tco_loop_go__go_3_0_0:;
  $b_4 = $__tco_var_go__go_3_0_0_b_4;
  $v_5 = $__tco_var_go__go_3_0_0_v_5;
  $__t0 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_4;
goto end_branch_0;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\Tuple\Data_Tuple_Tuple((($b_4)->{'value0'} + 1), new \Data\List\Types\Data_List_Types_Cons(($v_5)->{'value0'}, ($b_4)->{'value1'}));
$__tco_2 = ($v_5)->{'value1'};
$__tco_var_go__go_3_0_0_b_4 = $__tco_1;
$__tco_var_go__go_3_0_0_v_5 = $__tco_2;
goto tco_loop_go__go_3_0_0;;
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
  $v_3_0 = (($go__go_3_0)(new \Data\Tuple\Data_Tuple_Tuple(0, new \Data\List\Types\Data_List_Types_Nil())))($xs_2);
  $go__go_4_2 = null;
  $go__go_4_2 = (function() use ($f_0, &$go__go_4_2) {
  $__fn = function($b_5, $v_6 = null) use ($f_0, &$go__go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_2_2_b_5 = $b_5;
  $__tco_var_go__go_4_2_2_v_6 = $v_6;
  tco_loop_go__go_4_2_2:;
  $b_5 = $__tco_var_go__go_4_2_2_b_5;
  $v_6 = $__tco_var_go__go_4_2_2_v_6;
  $__t2 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_5;
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\Tuple\Data_Tuple_Tuple((($b_5)->{'value0'} - 1), ((($f_0)((($b_5)->{'value0'} - 1)))(($v_6)->{'value0'}))(($b_5)->{'value1'}));
$__tco_4 = ($v_6)->{'value1'};
$__tco_var_go__go_4_2_2_b_5 = $__tco_3;
$__tco_var_go__go_4_2_2_v_6 = $__tco_4;
goto tco_loop_go__go_4_2_2;;
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
  $__res = ((($go__go_4_2)(new \Data\Tuple\Data_Tuple_Tuple(($v_3_0)->{'value0'}, $b_1)))(($v_3_0)->{'value1'}))->{'value1'};
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($acc_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_3 = null;
  $go__go_2_3 = (function() use ($f_0, &$go__go_2_3) {
  $__fn = function($b_3, $v_4 = null) use ($f_0, &$go__go_2_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_3_3_b_3 = $b_3;
  $__tco_var_go__go_2_3_3_v_4 = $v_4;
  tco_loop_go__go_2_3_3:;
  $b_3 = $__tco_var_go__go_2_3_3_b_3;
  $v_4 = $__tco_var_go__go_2_3_3_v_4;
  $__t3 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_3;
goto end_branch_3;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = new \Data\Tuple\Data_Tuple_Tuple((($b_3)->{'value0'} + 1), ((($f_0)(($b_3)->{'value0'}))(($b_3)->{'value1'}))(($v_4)->{'value0'}));
$__tco_5 = ($v_4)->{'value1'};
$__tco_var_go__go_2_3_3_b_3 = $__tco_4;
$__tco_var_go__go_2_3_3_v_4 = $__tco_5;
goto tco_loop_go__go_2_3_3;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))(($go__go_2_3)(new \Data\Tuple\Data_Tuple_Tuple(0, $acc_1)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_4 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_5 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_4, $mempty_2_5) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_foldableWithIndexList'])->{'foldlWithIndex'})(function($i_4) use ($Semigroup0_1_4, $f_3) {
  $__num = \func_num_args();
  $__res = function($acc_5) use ($Semigroup0_1_4, $f_3, $i_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Semigroup0_1_4)->{'append'})($acc_5)))(($f_3)($i_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_foldableWithIndexNonEmptyList
$GLOBALS['Data_List_Types_foldableWithIndexNonEmptyList'] = (object)["foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictMonoid_0, $f_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_1))(function($v2_3) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = (1 + ($v2_3)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $Semigroup0_4_2 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__local_var_5_3 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_3_0))($GLOBALS['Data_Maybe_Just']);
  $go__go_6_4 = null;
  $go__go_6_4 = (function() use ($Semigroup0_4_2, $__local_var_5_3, &$go__go_6_4) {
  $__fn = function($b_7, $v_8 = null) use ($Semigroup0_4_2, $__local_var_5_3, &$go__go_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_4_4_b_7 = $b_7;
  $__tco_var_go__go_6_4_4_v_8 = $v_8;
  tco_loop_go__go_6_4_4:;
  $b_7 = $__tco_var_go__go_6_4_4_b_7;
  $v_8 = $__tco_var_go__go_6_4_4_v_8;
  $__t4 = null;;
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $b_7;
goto end_branch_4;;
};
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = new \Data\Tuple\Data_Tuple_Tuple((($b_7)->{'value0'} + 1), \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_4_2)->{'append'})(($b_7)->{'value1'}), ($__local_var_5_3)(($b_7)->{'value0'}), ($v_8)->{'value0'}));
$__tco_6 = ($v_8)->{'value1'};
$__tco_var_go__go_6_4_4_b_7 = $__tco_5;
$__tco_var_go__go_6_4_4_v_8 = $__tco_6;
goto tco_loop_go__go_6_4_4;;
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
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})((($__local_var_3_0)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_2)->{'value0'})))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Tuple_snd'], ($go__go_6_4)(new \Data\Tuple\Data_Tuple_Tuple(0, ($dictMonoid_0)->{'mempty'})), ($v_2)->{'value1'}));
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
}, "foldlWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__local_var_3_5 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = 0;
goto end_branch_5;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = (1 + ($v2_3)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_4_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_3_5))($GLOBALS['Data_Maybe_Just']);
  $go__go_5_8 = null;
  $go__go_5_8 = (function() use ($__local_var_4_7, &$go__go_5_8) {
  $__fn = function($b_6, $v_7 = null) use ($__local_var_4_7, &$go__go_5_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_8_8_b_6 = $b_6;
  $__tco_var_go__go_5_8_8_v_7 = $v_7;
  tco_loop_go__go_5_8_8:;
  $b_6 = $__tco_var_go__go_5_8_8_b_6;
  $v_7 = $__tco_var_go__go_5_8_8_v_7;
  $__t8 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t8 = $b_6;
goto end_branch_8;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_9 = new \Data\Tuple\Data_Tuple_Tuple((($b_6)->{'value0'} + 1), ((($__local_var_4_7)(($b_6)->{'value0'}))(($b_6)->{'value1'}))(($v_7)->{'value0'}));
$__tco_10 = ($v_7)->{'value1'};
$__tco_var_go__go_5_8_8_b_6 = $__tco_9;
$__tco_var_go__go_5_8_8_v_7 = $__tco_10;
goto tco_loop_go__go_5_8_8;;
$__t8 = null;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Tuple_snd'], ($go__go_5_8)(new \Data\Tuple\Data_Tuple_Tuple(0, ((($__local_var_3_5)(new \Data\Maybe\Data_Maybe_Nothing()))($b_1))(($v_2)->{'value0'}))), ($v_2)->{'value1'});
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
}, "foldrWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $__local_var_3_9 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t9 = 0;
goto end_branch_9;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t9 = (1 + ($v2_3)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_4_11 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_3_9))($GLOBALS['Data_Maybe_Just']);
  $go__go_5_12 = null;
  $go__go_5_12 = (function() use (&$go__go_5_12) {
  $__fn = function($b_6, $v_7 = null) use (&$go__go_5_12, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_12_12_b_6 = $b_6;
  $__tco_var_go__go_5_12_12_v_7 = $v_7;
  tco_loop_go__go_5_12_12:;
  $b_6 = $__tco_var_go__go_5_12_12_b_6;
  $v_7 = $__tco_var_go__go_5_12_12_v_7;
  $__t12 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t12 = $b_6;
goto end_branch_12;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_13 = new \Data\Tuple\Data_Tuple_Tuple((($b_6)->{'value0'} + 1), new \Data\List\Types\Data_List_Types_Cons(($v_7)->{'value0'}, ($b_6)->{'value1'}));
$__tco_14 = ($v_7)->{'value1'};
$__tco_var_go__go_5_12_12_b_6 = $__tco_13;
$__tco_var_go__go_5_12_12_v_7 = $__tco_14;
goto tco_loop_go__go_5_12_12;;
$__t12 = null;
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $v_5_12 = (($go__go_5_12)(new \Data\Tuple\Data_Tuple_Tuple(0, new \Data\List\Types\Data_List_Types_Nil())))(($v_2)->{'value1'});
  $go__go_6_14 = null;
  $go__go_6_14 = (function() use ($__local_var_4_11, &$go__go_6_14) {
  $__fn = function($b_7, $v_8 = null) use ($__local_var_4_11, &$go__go_6_14, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_14_14_b_7 = $b_7;
  $__tco_var_go__go_6_14_14_v_8 = $v_8;
  tco_loop_go__go_6_14_14:;
  $b_7 = $__tco_var_go__go_6_14_14_b_7;
  $v_8 = $__tco_var_go__go_6_14_14_v_8;
  $__t14 = null;;
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t14 = $b_7;
goto end_branch_14;;
};
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_15 = new \Data\Tuple\Data_Tuple_Tuple((($b_7)->{'value0'} - 1), ((($__local_var_4_11)((($b_7)->{'value0'} - 1)))(($v_8)->{'value0'}))(($b_7)->{'value1'}));
$__tco_16 = ($v_8)->{'value1'};
$__tco_var_go__go_6_14_14_b_7 = $__tco_15;
$__tco_var_go__go_6_14_14_v_8 = $__tco_16;
goto tco_loop_go__go_6_14_14;;
$__t14 = null;
goto end_branch_14;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t14 = null;
  end_branch_14:;
  $__res = $__t14;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($__local_var_3_9)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_2)->{'value0'}))(((($go__go_6_14)(new \Data\Tuple\Data_Tuple_Tuple(($v_5_12)->{'value0'}, $b_1)))(($v_5_12)->{'value1'}))->{'value1'});
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
}, "Foldable0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_functorWithIndexList
$GLOBALS['Data_List_Types_functorWithIndexList'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($xs_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\Tuple\Data_Tuple_Tuple((($b_3)->{'value0'} + 1), new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, ($b_3)->{'value1'}));
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $v_2_0 = (($go__go_2_0)(new \Data\Tuple\Data_Tuple_Tuple(0, new \Data\List\Types\Data_List_Types_Nil())))($xs_1);
  $go__go_3_2 = null;
  $go__go_3_2 = (function() use ($f_0, &$go__go_3_2) {
  $__fn = function($b_4, $v_5 = null) use ($f_0, &$go__go_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_2_2_b_4 = $b_4;
  $__tco_var_go__go_3_2_2_v_5 = $v_5;
  tco_loop_go__go_3_2_2:;
  $b_4 = $__tco_var_go__go_3_2_2_b_4;
  $v_5 = $__tco_var_go__go_3_2_2_v_5;
  $__t2 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_4;
goto end_branch_2;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\Tuple\Data_Tuple_Tuple((($b_4)->{'value0'} - 1), new \Data\List\Types\Data_List_Types_Cons((($f_0)((($b_4)->{'value0'} - 1)))(($v_5)->{'value0'}), ($b_4)->{'value1'}));
$__tco_4 = ($v_5)->{'value1'};
$__tco_var_go__go_3_2_2_b_4 = $__tco_3;
$__tco_var_go__go_3_2_2_v_5 = $__tco_4;
goto tco_loop_go__go_3_2_2;;
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
  $__res = ((($go__go_3_2)(new \Data\Tuple\Data_Tuple_Tuple(($v_2_0)->{'value0'}, new \Data\List\Types\Data_List_Types_Nil())))(($v_2_0)->{'value1'}))->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_functorWithIndexNonEmptyList
$GLOBALS['Data_List_Types_functorWithIndexNonEmptyList'] = (object)["mapWithIndex" => function($fn_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($fn_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($fn_0))(function($v2_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = (1 + ($v2_2)->{'value0'});
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__local_var_3_2 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_2_0))($GLOBALS['Data_Maybe_Just']);
  $go__go_4_3 = null;
  $go__go_4_3 = (function() use (&$go__go_4_3) {
  $__fn = function($b_5, $v_6 = null) use (&$go__go_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_3_3_b_5 = $b_5;
  $__tco_var_go__go_4_3_3_v_6 = $v_6;
  tco_loop_go__go_4_3_3:;
  $b_5 = $__tco_var_go__go_4_3_3_b_5;
  $v_6 = $__tco_var_go__go_4_3_3_v_6;
  $__t3 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_5;
goto end_branch_3;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = new \Data\Tuple\Data_Tuple_Tuple((($b_5)->{'value0'} + 1), new \Data\List\Types\Data_List_Types_Cons(($v_6)->{'value0'}, ($b_5)->{'value1'}));
$__tco_5 = ($v_6)->{'value1'};
$__tco_var_go__go_4_3_3_b_5 = $__tco_4;
$__tco_var_go__go_4_3_3_v_6 = $__tco_5;
goto tco_loop_go__go_4_3_3;;
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
  $v_4_3 = (($go__go_4_3)(new \Data\Tuple\Data_Tuple_Tuple(0, new \Data\List\Types\Data_List_Types_Nil())))(($v_1)->{'value1'});
  $go__go_5_5 = null;
  $go__go_5_5 = (function() use ($__local_var_3_2, &$go__go_5_5) {
  $__fn = function($b_6, $v_7 = null) use ($__local_var_3_2, &$go__go_5_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_5_5_b_6 = $b_6;
  $__tco_var_go__go_5_5_5_v_7 = $v_7;
  tco_loop_go__go_5_5_5:;
  $b_6 = $__tco_var_go__go_5_5_5_b_6;
  $v_7 = $__tco_var_go__go_5_5_5_v_7;
  $__t5 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $b_6;
goto end_branch_5;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = new \Data\Tuple\Data_Tuple_Tuple((($b_6)->{'value0'} - 1), new \Data\List\Types\Data_List_Types_Cons((($__local_var_3_2)((($b_6)->{'value0'} - 1)))(($v_7)->{'value0'}), ($b_6)->{'value1'}));
$__tco_7 = ($v_7)->{'value1'};
$__tco_var_go__go_5_5_5_b_6 = $__tco_6;
$__tco_var_go__go_5_5_5_v_7 = $__tco_7;
goto tco_loop_go__go_5_5_5;;
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
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((($__local_var_2_0)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_1)->{'value0'}), ((($go__go_5_5)(new \Data\Tuple\Data_Tuple_Tuple(($v_4_3)->{'value0'}, new \Data\List\Types\Data_List_Types_Nil())))(($v_4_3)->{'value1'}))->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_semigroupList
$GLOBALS['Data_List_Types_semigroupList'] = (object)["append" => function($xs_0) {
  $__num = \func_num_args();
  $__res = function($ys_1) use ($xs_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $go__go_2_1 = null;
  $go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_v_3 = $v_3;
  $__tco_var_go__go_2_1_1_v1_4 = $v1_4;
  tco_loop_go__go_2_1_1:;
  $v_3 = $__tco_var_go__go_2_1_1_v_3;
  $v1_4 = $__tco_var_go__go_2_1_1_v1_4;
  $__t1 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_3;
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_3 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_1_1_v_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v1_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_0)($ys_1), ($go__go_2_1)(new \Data\List\Types\Data_List_Types_Nil()), $xs_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_monoidList
$GLOBALS['Data_List_Types_monoidList'] = (object)["mempty" => new \Data\List\Types\Data_List_Types_Nil(), "Semigroup0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_semigroupList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_semigroupNonEmptyList
$GLOBALS['Data_List_Types_semigroupNonEmptyList'] = (object)["append" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($as_prime__1) use ($v_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $go__go_2_1 = null;
  $go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_v_3 = $v_3;
  $__tco_var_go__go_2_1_1_v1_4 = $v1_4;
  tco_loop_go__go_2_1_1:;
  $v_3 = $__tco_var_go__go_2_1_1_v_3;
  $v1_4 = $__tco_var_go__go_2_1_1_v1_4;
  $__t1 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_3;
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_3 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_1_1_v_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v1_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_0)->{'value0'}, \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_0)(new \Data\List\Types\Data_List_Types_Cons(($as_prime__1)->{'value0'}, ($as_prime__1)->{'value1'})), ($go__go_2_1)(new \Data\List\Types\Data_List_Types_Nil()), ($v_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_showList
function majData_majList_majTypes_showmajList($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_showmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $show_1_0 = ($dictShow_0)->{'show'};
  $__res = (object)["show" => function($v_2) use ($show_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = "Nil";
goto end_branch_2;;
};
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use (&$go__go_3_1) {
  $__fn = function($b_4, $v_5 = null) use (&$go__go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_1_1_b_4 = $b_4;
  $__tco_var_go__go_3_1_1_v_5 = $v_5;
  tco_loop_go__go_3_1_1:;
  $b_4 = $__tco_var_go__go_3_1_1_b_4;
  $v_5 = $__tco_var_go__go_3_1_1_v_5;
  $__t1 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_4;
goto end_branch_1;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
if (($b_4)->{'init'}) {
$__t2 = (object)["init" => false, "acc" => ($v_5)->{'value0'}];
goto end_branch_2;;
};
$__t2 = (object)["init" => false, "acc" => ((($b_4)->{'acc'} . " : ") . ($v_5)->{'value0'})];
end_branch_2:;
$__tco_3 = $__t2;
$__tco_4 = ($v_5)->{'value1'};
$__tco_var_go__go_3_1_1_b_4 = $__tco_3;
$__tco_var_go__go_3_1_1_v_5 = $__tco_4;
goto tco_loop_go__go_3_1_1;;
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
  $__t2 = (("(" . ((($go__go_3_1)((object)["init" => true, "acc" => ""]))(\Data\List\Types\majData_majList_majTypes_listmajMap($show_1_0, $v_2)))->{'acc'}) . " : Nil)");
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_showList'] = __NAMESPACE__ . '\\majData_majList_majTypes_showmajList';

// Data_List_Types_showNonEmptyList
function majData_majList_majTypes_showmajNonmajEmptymajList($dictShow_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_showmajNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $show_1_0 = ($dictShow_0)->{'show'};
  $__local_var_2_1 = (object)["show" => function($v_2) use ($show_1_0) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = "Nil";
goto end_branch_2;;
};
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use (&$go__go_3_1) {
  $__fn = function($b_4, $v_5 = null) use (&$go__go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_1_1_b_4 = $b_4;
  $__tco_var_go__go_3_1_1_v_5 = $v_5;
  tco_loop_go__go_3_1_1:;
  $b_4 = $__tco_var_go__go_3_1_1_b_4;
  $v_5 = $__tco_var_go__go_3_1_1_v_5;
  $__t1 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_4;
goto end_branch_1;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t2 = null;;
if (($b_4)->{'init'}) {
$__t2 = (object)["init" => false, "acc" => ($v_5)->{'value0'}];
goto end_branch_2;;
};
$__t2 = (object)["init" => false, "acc" => ((($b_4)->{'acc'} . " : ") . ($v_5)->{'value0'})];
end_branch_2:;
$__tco_3 = $__t2;
$__tco_4 = ($v_5)->{'value1'};
$__tco_var_go__go_3_1_1_b_4 = $__tco_3;
$__tco_var_go__go_3_1_1_v_5 = $__tco_4;
goto tco_loop_go__go_3_1_1;;
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
  $__t2 = (("(" . ((($go__go_3_1)((object)["init" => true, "acc" => ""]))(\Data\List\Types\majData_majList_majTypes_listmajMap($show_1_0, $v_2)))->{'acc'}) . " : Nil)");
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $showNonEmpty_1_0 = (object)["show" => function($v_3) use ($__local_var_2_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = (((("(NonEmpty " . (($dictShow_0)->{'show'})(($v_3)->{'value0'})) . " ") . (($__local_var_2_1)->{'show'})(($v_3)->{'value1'})) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["show" => function($v_2) use ($showNonEmpty_1_0) {
  $__num = \func_num_args();
  $__res = (("(NonEmptyList " . (($showNonEmpty_1_0)->{'show'})($v_2)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_showNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majTypes_showmajNonmajEmptymajList';

// Data_List_Types_traversableList
$GLOBALS['Data_List_Types_traversableList'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_2_1 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($f_3) use ($Apply0_2_1, $Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $go__go_4_2 = null;
  $go__go_4_2 = (function() use (&$go__go_4_2) {
  $__fn = function($b_5, $v_6 = null) use (&$go__go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_2_2_b_5 = $b_5;
  $__tco_var_go__go_4_2_2_v_6 = $v_6;
  tco_loop_go__go_4_2_2:;
  $b_5 = $__tco_var_go__go_4_2_2_b_5;
  $v_6 = $__tco_var_go__go_4_2_2_v_6;
  $__t2 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_5;
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v_6)->{'value0'}, $b_5);
$__tco_4 = ($v_6)->{'value1'};
$__tco_var_go__go_4_2_2_b_5 = $__tco_3;
$__tco_var_go__go_4_2_2_v_6 = $__tco_4;
goto tco_loop_go__go_4_2_2;;
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
  $go__go_4_3 = null;
  $go__go_4_3 = (function() use ($Apply0_2_1, $f_3, &$go__go_4_3) {
  $__fn = function($b_5, $v_6 = null) use ($Apply0_2_1, $f_3, &$go__go_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_3_3_b_5 = $b_5;
  $__tco_var_go__go_4_3_3_v_6 = $v_6;
  tco_loop_go__go_4_3_3:;
  $b_5 = $__tco_var_go__go_4_3_3_b_5;
  $v_6 = $__tco_var_go__go_4_3_3_v_6;
  $__t3 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_5;
goto end_branch_3;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_7_4 = (($Apply0_2_1)->{'Functor0'})(null);
$__tco_5 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_8) use ($Apply0_2_1, $Functor0_7_4, $b_5) {
  $__num = \func_num_args();
  $__res = ((($Apply0_2_1)->{'apply'})(((($Functor0_7_4)->{'map'})(function($b_9) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($b_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_10, $b_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_5)))($b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $f_3, ($v_6)->{'value0'});
$__tco_6 = ($v_6)->{'value1'};
$__tco_var_go__go_4_3_3_b_5 = $__tco_5;
$__tco_var_go__go_4_3_3_v_6 = $__tco_6;
goto tco_loop_go__go_4_3_3;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_1_0)->{'map'})(($go__go_4_2)(new \Data\List\Types\Data_List_Types_Nil()))))(($go__go_4_3)((($dictApplicative_0)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_traversableList'])->{'traverse'})($dictApplicative_0))(function($x_1) {
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
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_traversableNonEmptyList
$GLOBALS['Data_List_Types_traversableNonEmptyList'] = (function() use (&$__fn) {
$functorNonEmpty1_0_0 = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($m_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)(($m_1)->{'value0'}), \Data\List\Types\majData_majList_majTypes_listmajMap($f_0, ($m_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
$foldableNonEmpty1_1_1 = (object)["foldMap" => function($dictMonoid_1) {
  $__num = \func_num_args();
  $Semigroup0_2_1 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $__res = function($f_3) use ($Semigroup0_2_1, $dictMonoid_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($Semigroup0_2_1, $dictMonoid_1, $f_3) {
  $__num = \func_num_args();
  $Semigroup0_5_2 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $go__go_6_3 = null;
  $go__go_6_3 = (function() use ($Semigroup0_5_2, $f_3, &$go__go_6_3) {
  $__fn = function($b_7, $v_8 = null) use ($Semigroup0_5_2, $f_3, &$go__go_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_3_3_b_7 = $b_7;
  $__tco_var_go__go_6_3_3_v_8 = $v_8;
  tco_loop_go__go_6_3_3:;
  $b_7 = $__tco_var_go__go_6_3_3_b_7;
  $v_8 = $__tco_var_go__go_6_3_3_v_8;
  $__t3 = null;;
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_7;
goto end_branch_3;;
};
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_5_2)->{'append'})($b_7), $f_3, ($v_8)->{'value0'});
$__tco_5 = ($v_8)->{'value1'};
$__tco_var_go__go_6_3_3_b_7 = $__tco_4;
$__tco_var_go__go_6_3_3_v_8 = $__tco_5;
goto tco_loop_go__go_6_3_3;;
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
  $__res = ((($Semigroup0_2_1)->{'append'})(($f_3)(($v_4)->{'value0'})))((($go__go_6_3)(($dictMonoid_1)->{'mempty'}))(($v_4)->{'value1'}));
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
}, "foldl" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $f_1) {
  $__num = \func_num_args();
  $go__go_4_4 = null;
  $go__go_4_4 = (function() use ($f_1, &$go__go_4_4) {
  $__fn = function($b_5, $v_6 = null) use ($f_1, &$go__go_4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_4_4_b_5 = $b_5;
  $__tco_var_go__go_4_4_4_v_6 = $v_6;
  tco_loop_go__go_4_4_4:;
  $b_5 = $__tco_var_go__go_4_4_4_b_5;
  $v_6 = $__tco_var_go__go_4_4_4_v_6;
  $__t4 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $b_5;
goto end_branch_4;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = (($f_1)($b_5))(($v_6)->{'value0'});
$__tco_6 = ($v_6)->{'value1'};
$__tco_var_go__go_4_4_4_b_5 = $__tco_5;
$__tco_var_go__go_4_4_4_v_6 = $__tco_6;
goto tco_loop_go__go_4_4_4;;
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
  $__res = (($go__go_4_4)((($f_1)($b_2))(($v_3)->{'value0'})))(($v_3)->{'value1'});
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
}, "foldr" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $f_1) {
  $__num = \func_num_args();
  $go__go_4_5 = null;
  $go__go_4_5 = (function() use ($f_1, &$go__go_4_5) {
  $__fn = function($b_5, $v_6 = null) use ($f_1, &$go__go_4_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_5_5_b_5 = $b_5;
  $__tco_var_go__go_4_5_5_v_6 = $v_6;
  tco_loop_go__go_4_5_5:;
  $b_5 = $__tco_var_go__go_4_5_5_b_5;
  $v_6 = $__tco_var_go__go_4_5_5_v_6;
  $__t5 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $b_5;
goto end_branch_5;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = (($f_1)(($v_6)->{'value0'}))($b_5);
$__tco_7 = ($v_6)->{'value1'};
$__tco_var_go__go_4_5_5_b_5 = $__tco_6;
$__tco_var_go__go_4_5_5_v_6 = $__tco_7;
goto tco_loop_go__go_4_5_5;;
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
  $go__go_4_6 = null;
  $go__go_4_6 = (function() use (&$go__go_4_6) {
  $__fn = function($v_5, $v1_6 = null) use (&$go__go_4_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_6_6_v_5 = $v_5;
  $__tco_var_go__go_4_6_6_v1_6 = $v1_6;
  tco_loop_go__go_4_6_6:;
  $v_5 = $__tco_var_go__go_4_6_6_v_5;
  $v1_6 = $__tco_var_go__go_4_6_6_v1_6;
  $__t6 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = $v_5;
goto end_branch_6;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_7 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_8 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_6_6_v_5 = $__tco_7;
$__tco_var_go__go_4_6_6_v1_6 = $__tco_8;
goto tco_loop_go__go_4_6_6;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($f_1)(($v_3)->{'value0'}))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_4_5)($b_2), ($go__go_4_6)(new \Data\List\Types\Data_List_Types_Nil()), ($v_3)->{'value1'}));
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
return (object)["sequence" => function($dictApplicative_2) {
  $__num = \func_num_args();
  $Apply0_3_8 = (($dictApplicative_2)->{'Apply0'})(null);
  $Functor0_4_9 = (((($dictApplicative_2)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_5) use ($Apply0_3_8, $Functor0_4_9, $dictApplicative_2) {
  $__num = \func_num_args();
  $Apply0_6_10 = (($dictApplicative_2)->{'Apply0'})(null);
  $go__go_7_11 = null;
  $go__go_7_11 = (function() use (&$go__go_7_11) {
  $__fn = function($b_8, $v_9 = null) use (&$go__go_7_11, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_11_11_b_8 = $b_8;
  $__tco_var_go__go_7_11_11_v_9 = $v_9;
  tco_loop_go__go_7_11_11:;
  $b_8 = $__tco_var_go__go_7_11_11_b_8;
  $v_9 = $__tco_var_go__go_7_11_11_v_9;
  $__t11 = null;;
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t11 = $b_8;
goto end_branch_11;;
};
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_12 = new \Data\List\Types\Data_List_Types_Cons(($v_9)->{'value0'}, $b_8);
$__tco_13 = ($v_9)->{'value1'};
$__tco_var_go__go_7_11_11_b_8 = $__tco_12;
$__tco_var_go__go_7_11_11_v_9 = $__tco_13;
goto tco_loop_go__go_7_11_11;;
$__t11 = null;
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_7_12 = null;
  $go__go_7_12 = (function() use ($Apply0_6_10, &$go__go_7_12) {
  $__fn = function($b_8, $v_9 = null) use ($Apply0_6_10, &$go__go_7_12, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_12_12_b_8 = $b_8;
  $__tco_var_go__go_7_12_12_v_9 = $v_9;
  tco_loop_go__go_7_12_12:;
  $b_8 = $__tco_var_go__go_7_12_12_b_8;
  $v_9 = $__tco_var_go__go_7_12_12_v_9;
  $__t12 = null;;
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t12 = $b_8;
goto end_branch_12;;
};
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_10_13 = (($Apply0_6_10)->{'Functor0'})(null);
$__tco_14 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_11) use ($Apply0_6_10, $Functor0_10_13, $b_8) {
  $__num = \func_num_args();
  $__res = ((($Apply0_6_10)->{'apply'})(((($Functor0_10_13)->{'map'})(function($b_12) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($b_12) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_13, $b_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_8)))($b_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_9)->{'value0'});
$__tco_15 = ($v_9)->{'value1'};
$__tco_var_go__go_7_12_12_b_8 = $__tco_14;
$__tco_var_go__go_7_12_12_v_9 = $__tco_15;
goto tco_loop_go__go_7_12_12;;
$__t12 = null;
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Apply0_3_8)->{'apply'})(((($Functor0_4_9)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($v_5)->{'value0'})))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((((((($dictApplicative_2)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($go__go_7_11)(new \Data\List\Types\Data_List_Types_Nil())), ($go__go_7_12)((($dictApplicative_2)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())), ($v_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "traverse" => function($dictApplicative_2) {
  $__num = \func_num_args();
  $Apply0_3_13 = (($dictApplicative_2)->{'Apply0'})(null);
  $Functor0_4_14 = (((($dictApplicative_2)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_5) use ($Apply0_3_13, $Functor0_4_14, $dictApplicative_2) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($Apply0_3_13, $Functor0_4_14, $dictApplicative_2, $f_5) {
  $__num = \func_num_args();
  $Apply0_7_15 = (($dictApplicative_2)->{'Apply0'})(null);
  $go__go_8_16 = null;
  $go__go_8_16 = (function() use (&$go__go_8_16) {
  $__fn = function($b_9, $v_10 = null) use (&$go__go_8_16, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_16_16_b_9 = $b_9;
  $__tco_var_go__go_8_16_16_v_10 = $v_10;
  tco_loop_go__go_8_16_16:;
  $b_9 = $__tco_var_go__go_8_16_16_b_9;
  $v_10 = $__tco_var_go__go_8_16_16_v_10;
  $__t16 = null;;
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t16 = $b_9;
goto end_branch_16;;
};
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_17 = new \Data\List\Types\Data_List_Types_Cons(($v_10)->{'value0'}, $b_9);
$__tco_18 = ($v_10)->{'value1'};
$__tco_var_go__go_8_16_16_b_9 = $__tco_17;
$__tco_var_go__go_8_16_16_v_10 = $__tco_18;
goto tco_loop_go__go_8_16_16;;
$__t16 = null;
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_8_17 = null;
  $go__go_8_17 = (function() use ($Apply0_7_15, $f_5, &$go__go_8_17) {
  $__fn = function($b_9, $v_10 = null) use ($Apply0_7_15, $f_5, &$go__go_8_17, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_17_17_b_9 = $b_9;
  $__tco_var_go__go_8_17_17_v_10 = $v_10;
  tco_loop_go__go_8_17_17:;
  $b_9 = $__tco_var_go__go_8_17_17_b_9;
  $v_10 = $__tco_var_go__go_8_17_17_v_10;
  $__t17 = null;;
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t17 = $b_9;
goto end_branch_17;;
};
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_11_18 = (($Apply0_7_15)->{'Functor0'})(null);
$__tco_19 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_12) use ($Apply0_7_15, $Functor0_11_18, $b_9) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_15)->{'apply'})(((($Functor0_11_18)->{'map'})(function($b_13) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($b_13) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_14, $b_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_9)))($b_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $f_5, ($v_10)->{'value0'});
$__tco_20 = ($v_10)->{'value1'};
$__tco_var_go__go_8_17_17_b_9 = $__tco_19;
$__tco_var_go__go_8_17_17_v_10 = $__tco_20;
goto tco_loop_go__go_8_17_17;;
$__t17 = null;
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Apply0_3_13)->{'apply'})(((($Functor0_4_14)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($f_5)(($v_6)->{'value0'}))))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((((((($dictApplicative_2)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($go__go_8_16)(new \Data\List\Types\Data_List_Types_Nil())), ($go__go_8_17)((($dictApplicative_2)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())), ($v_6)->{'value1'}));
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
}, "Functor0" => function($_dollar___unused_2) use ($functorNonEmpty1_0_0) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_2) use ($foldableNonEmpty1_1_1) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_List_Types_traversableWithIndexList
$GLOBALS['Data_List_Types_traversableWithIndexList'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_2_1 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($f_3) use ($Apply0_2_1, $Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $go__go_4_2 = null;
  $go__go_4_2 = (function() use (&$go__go_4_2) {
  $__fn = function($b_5, $v_6 = null) use (&$go__go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_2_2_b_5 = $b_5;
  $__tco_var_go__go_4_2_2_v_6 = $v_6;
  tco_loop_go__go_4_2_2:;
  $b_5 = $__tco_var_go__go_4_2_2_b_5;
  $v_6 = $__tco_var_go__go_4_2_2_v_6;
  $__t2 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_5;
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v_6)->{'value0'}, $b_5);
$__tco_4 = ($v_6)->{'value1'};
$__tco_var_go__go_4_2_2_b_5 = $__tco_3;
$__tco_var_go__go_4_2_2_v_6 = $__tco_4;
goto tco_loop_go__go_4_2_2;;
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
  $go__go_4_3 = null;
  $go__go_4_3 = (function() use ($Apply0_2_1, $f_3, &$go__go_4_3) {
  $__fn = function($b_5, $v_6 = null) use ($Apply0_2_1, $f_3, &$go__go_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_3_3_b_5 = $b_5;
  $__tco_var_go__go_4_3_3_v_6 = $v_6;
  tco_loop_go__go_4_3_3:;
  $b_5 = $__tco_var_go__go_4_3_3_b_5;
  $v_6 = $__tco_var_go__go_4_3_3_v_6;
  $__t3 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_5;
goto end_branch_3;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_7_4 = (($Apply0_2_1)->{'Functor0'})(null);
$__tco_5 = new \Data\Tuple\Data_Tuple_Tuple((($b_5)->{'value0'} + 1), \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_8) use ($Apply0_2_1, $Functor0_7_4, $b_5) {
  $__num = \func_num_args();
  $__res = ((($Apply0_2_1)->{'apply'})(((($Functor0_7_4)->{'map'})(function($b_9) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($b_9) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_10, $b_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($b_5)->{'value1'})))($b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($f_3)(($b_5)->{'value0'}), ($v_6)->{'value0'}));
$__tco_6 = ($v_6)->{'value1'};
$__tco_var_go__go_4_3_3_b_5 = $__tco_5;
$__tco_var_go__go_4_3_3_v_6 = $__tco_6;
goto tco_loop_go__go_4_3_3;;
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_1_0)->{'map'})(($go__go_4_2)(new \Data\List\Types\Data_List_Types_Nil()))))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))(($go__go_4_3)(new \Data\Tuple\Data_Tuple_Tuple(0, (($dictApplicative_0)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_traversableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_traversableWithIndexNonEmptyList
$GLOBALS['Data_List_Types_traversableWithIndexNonEmptyList'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $__local_var_4_1 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))(function($v2_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v2_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($v2_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (1 + ($v2_4)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $Apply0_5_3 = (($dictApplicative_0)->{'Apply0'})(null);
  $__local_var_6_4 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($__local_var_4_1))($GLOBALS['Data_Maybe_Just']);
  $go__go_7_5 = null;
  $go__go_7_5 = (function() use (&$go__go_7_5) {
  $__fn = function($b_8, $v_9 = null) use (&$go__go_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_5_5_b_8 = $b_8;
  $__tco_var_go__go_7_5_5_v_9 = $v_9;
  tco_loop_go__go_7_5_5:;
  $b_8 = $__tco_var_go__go_7_5_5_b_8;
  $v_9 = $__tco_var_go__go_7_5_5_v_9;
  $__t5 = null;;
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $b_8;
goto end_branch_5;;
};
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(($v_9)->{'value0'}, $b_8);
$__tco_7 = ($v_9)->{'value1'};
$__tco_var_go__go_7_5_5_b_8 = $__tco_6;
$__tco_var_go__go_7_5_5_v_9 = $__tco_7;
goto tco_loop_go__go_7_5_5;;
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
  $go__go_7_6 = null;
  $go__go_7_6 = (function() use ($Apply0_5_3, $__local_var_6_4, &$go__go_7_6) {
  $__fn = function($b_8, $v_9 = null) use ($Apply0_5_3, $__local_var_6_4, &$go__go_7_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_6_6_b_8 = $b_8;
  $__tco_var_go__go_7_6_6_v_9 = $v_9;
  tco_loop_go__go_7_6_6:;
  $b_8 = $__tco_var_go__go_7_6_6_b_8;
  $v_9 = $__tco_var_go__go_7_6_6_v_9;
  $__t6 = null;;
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = $b_8;
goto end_branch_6;;
};
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_10_7 = (($Apply0_5_3)->{'Functor0'})(null);
$__tco_8 = new \Data\Tuple\Data_Tuple_Tuple((($b_8)->{'value0'} + 1), \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_11) use ($Apply0_5_3, $Functor0_10_7, $b_8) {
  $__num = \func_num_args();
  $__res = ((($Apply0_5_3)->{'apply'})(((($Functor0_10_7)->{'map'})(function($b_12) {
  $__num = \func_num_args();
  $__res = function($a_13) use ($b_12) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_13, $b_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($b_8)->{'value1'})))($b_11);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($__local_var_6_4)(($b_8)->{'value0'}), ($v_9)->{'value0'}));
$__tco_9 = ($v_9)->{'value1'};
$__tco_var_go__go_7_6_6_b_8 = $__tco_8;
$__tco_var_go__go_7_6_6_v_9 = $__tco_9;
goto tco_loop_go__go_7_6_6;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((((($dictApplicative_0)->{'Apply0'})(null))->{'apply'})(((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))((($__local_var_4_1)(new \Data\Maybe\Data_Maybe_Nothing()))(($v_3)->{'value0'}))))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($go__go_7_5)(new \Data\List\Types\Data_List_Types_Nil())), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))(($go__go_7_6)(new \Data\Tuple\Data_Tuple_Tuple(0, (($dictApplicative_0)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())))), ($v_3)->{'value1'})));
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
}, "FunctorWithIndex0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $functorNonEmpty1_1_7 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($m_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_1)(($m_2)->{'value0'}), \Data\List\Types\majData_majList_majTypes_listmajMap($f_1, ($m_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $foldableNonEmpty1_2_8 = (object)["foldMap" => function($dictMonoid_2) {
  $__num = \func_num_args();
  $Semigroup0_3_8 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $__res = function($f_4) use ($Semigroup0_3_8, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Semigroup0_3_8, $dictMonoid_2, $f_4) {
  $__num = \func_num_args();
  $Semigroup0_6_9 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $go__go_7_10 = null;
  $go__go_7_10 = (function() use ($Semigroup0_6_9, $f_4, &$go__go_7_10) {
  $__fn = function($b_8, $v_9 = null) use ($Semigroup0_6_9, $f_4, &$go__go_7_10, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_10_10_b_8 = $b_8;
  $__tco_var_go__go_7_10_10_v_9 = $v_9;
  tco_loop_go__go_7_10_10:;
  $b_8 = $__tco_var_go__go_7_10_10_b_8;
  $v_9 = $__tco_var_go__go_7_10_10_v_9;
  $__t10 = null;;
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t10 = $b_8;
goto end_branch_10;;
};
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_11 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_6_9)->{'append'})($b_8), $f_4, ($v_9)->{'value0'});
$__tco_12 = ($v_9)->{'value1'};
$__tco_var_go__go_7_10_10_b_8 = $__tco_11;
$__tco_var_go__go_7_10_10_v_9 = $__tco_12;
goto tco_loop_go__go_7_10_10;;
$__t10 = null;
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Semigroup0_3_8)->{'append'})(($f_4)(($v_5)->{'value0'})))((($go__go_7_10)(($dictMonoid_2)->{'mempty'}))(($v_5)->{'value1'}));
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
}, "foldl" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($b_3, $f_2) {
  $__num = \func_num_args();
  $go__go_5_11 = null;
  $go__go_5_11 = (function() use ($f_2, &$go__go_5_11) {
  $__fn = function($b_6, $v_7 = null) use ($f_2, &$go__go_5_11, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_11_11_b_6 = $b_6;
  $__tco_var_go__go_5_11_11_v_7 = $v_7;
  tco_loop_go__go_5_11_11:;
  $b_6 = $__tco_var_go__go_5_11_11_b_6;
  $v_7 = $__tco_var_go__go_5_11_11_v_7;
  $__t11 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t11 = $b_6;
goto end_branch_11;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_12 = (($f_2)($b_6))(($v_7)->{'value0'});
$__tco_13 = ($v_7)->{'value1'};
$__tco_var_go__go_5_11_11_b_6 = $__tco_12;
$__tco_var_go__go_5_11_11_v_7 = $__tco_13;
goto tco_loop_go__go_5_11_11;;
$__t11 = null;
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_5_11)((($f_2)($b_3))(($v_4)->{'value0'})))(($v_4)->{'value1'});
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
}, "foldr" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($b_3, $f_2) {
  $__num = \func_num_args();
  $go__go_5_12 = null;
  $go__go_5_12 = (function() use ($f_2, &$go__go_5_12) {
  $__fn = function($b_6, $v_7 = null) use ($f_2, &$go__go_5_12, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_12_12_b_6 = $b_6;
  $__tco_var_go__go_5_12_12_v_7 = $v_7;
  tco_loop_go__go_5_12_12:;
  $b_6 = $__tco_var_go__go_5_12_12_b_6;
  $v_7 = $__tco_var_go__go_5_12_12_v_7;
  $__t12 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t12 = $b_6;
goto end_branch_12;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_13 = (($f_2)(($v_7)->{'value0'}))($b_6);
$__tco_14 = ($v_7)->{'value1'};
$__tco_var_go__go_5_12_12_b_6 = $__tco_13;
$__tco_var_go__go_5_12_12_v_7 = $__tco_14;
goto tco_loop_go__go_5_12_12;;
$__t12 = null;
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_5_13 = null;
  $go__go_5_13 = (function() use (&$go__go_5_13) {
  $__fn = function($v_6, $v1_7 = null) use (&$go__go_5_13, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_13_13_v_6 = $v_6;
  $__tco_var_go__go_5_13_13_v1_7 = $v1_7;
  tco_loop_go__go_5_13_13:;
  $v_6 = $__tco_var_go__go_5_13_13_v_6;
  $v1_7 = $__tco_var_go__go_5_13_13_v1_7;
  $__t13 = null;;
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t13 = $v_6;
goto end_branch_13;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_14 = new \Data\List\Types\Data_List_Types_Cons(($v1_7)->{'value0'}, $v_6);
$__tco_15 = ($v1_7)->{'value1'};
$__tco_var_go__go_5_13_13_v_6 = $__tco_14;
$__tco_var_go__go_5_13_13_v1_7 = $__tco_15;
goto tco_loop_go__go_5_13_13;;
$__t13 = null;
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($f_2)(($v_4)->{'value0'}))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_5_12)($b_3), ($go__go_5_13)(new \Data\List\Types\Data_List_Types_Nil()), ($v_4)->{'value1'}));
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
  $__res = (object)["sequence" => function($dictApplicative_3) {
  $__num = \func_num_args();
  $Apply0_4_15 = (($dictApplicative_3)->{'Apply0'})(null);
  $Functor0_5_16 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_6) use ($Apply0_4_15, $Functor0_5_16, $dictApplicative_3) {
  $__num = \func_num_args();
  $Apply0_7_17 = (($dictApplicative_3)->{'Apply0'})(null);
  $go__go_8_18 = null;
  $go__go_8_18 = (function() use (&$go__go_8_18) {
  $__fn = function($b_9, $v_10 = null) use (&$go__go_8_18, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_18_18_b_9 = $b_9;
  $__tco_var_go__go_8_18_18_v_10 = $v_10;
  tco_loop_go__go_8_18_18:;
  $b_9 = $__tco_var_go__go_8_18_18_b_9;
  $v_10 = $__tco_var_go__go_8_18_18_v_10;
  $__t18 = null;;
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t18 = $b_9;
goto end_branch_18;;
};
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_19 = new \Data\List\Types\Data_List_Types_Cons(($v_10)->{'value0'}, $b_9);
$__tco_20 = ($v_10)->{'value1'};
$__tco_var_go__go_8_18_18_b_9 = $__tco_19;
$__tco_var_go__go_8_18_18_v_10 = $__tco_20;
goto tco_loop_go__go_8_18_18;;
$__t18 = null;
goto end_branch_18;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t18 = null;
  end_branch_18:;
  $__res = $__t18;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_8_19 = null;
  $go__go_8_19 = (function() use ($Apply0_7_17, &$go__go_8_19) {
  $__fn = function($b_9, $v_10 = null) use ($Apply0_7_17, &$go__go_8_19, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_19_19_b_9 = $b_9;
  $__tco_var_go__go_8_19_19_v_10 = $v_10;
  tco_loop_go__go_8_19_19:;
  $b_9 = $__tco_var_go__go_8_19_19_b_9;
  $v_10 = $__tco_var_go__go_8_19_19_v_10;
  $__t19 = null;;
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t19 = $b_9;
goto end_branch_19;;
};
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_11_20 = (($Apply0_7_17)->{'Functor0'})(null);
$__tco_21 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_12) use ($Apply0_7_17, $Functor0_11_20, $b_9) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_17)->{'apply'})(((($Functor0_11_20)->{'map'})(function($b_13) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($b_13) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_14, $b_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_9)))($b_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_10)->{'value0'});
$__tco_22 = ($v_10)->{'value1'};
$__tco_var_go__go_8_19_19_b_9 = $__tco_21;
$__tco_var_go__go_8_19_19_v_10 = $__tco_22;
goto tco_loop_go__go_8_19_19;;
$__t19 = null;
goto end_branch_19;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t19 = null;
  end_branch_19:;
  $__res = $__t19;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Apply0_4_15)->{'apply'})(((($Functor0_5_16)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($v_6)->{'value0'})))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((((((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($go__go_8_18)(new \Data\List\Types\Data_List_Types_Nil())), ($go__go_8_19)((($dictApplicative_3)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())), ($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "traverse" => function($dictApplicative_3) {
  $__num = \func_num_args();
  $Apply0_4_20 = (($dictApplicative_3)->{'Apply0'})(null);
  $Functor0_5_21 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_6) use ($Apply0_4_20, $Functor0_5_21, $dictApplicative_3) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Apply0_4_20, $Functor0_5_21, $dictApplicative_3, $f_6) {
  $__num = \func_num_args();
  $Apply0_8_22 = (($dictApplicative_3)->{'Apply0'})(null);
  $go__go_9_23 = null;
  $go__go_9_23 = (function() use (&$go__go_9_23) {
  $__fn = function($b_10, $v_11 = null) use (&$go__go_9_23, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_9_23_23_b_10 = $b_10;
  $__tco_var_go__go_9_23_23_v_11 = $v_11;
  tco_loop_go__go_9_23_23:;
  $b_10 = $__tco_var_go__go_9_23_23_b_10;
  $v_11 = $__tco_var_go__go_9_23_23_v_11;
  $__t23 = null;;
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t23 = $b_10;
goto end_branch_23;;
};
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_24 = new \Data\List\Types\Data_List_Types_Cons(($v_11)->{'value0'}, $b_10);
$__tco_25 = ($v_11)->{'value1'};
$__tco_var_go__go_9_23_23_b_10 = $__tco_24;
$__tco_var_go__go_9_23_23_v_11 = $__tco_25;
goto tco_loop_go__go_9_23_23;;
$__t23 = null;
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_9_24 = null;
  $go__go_9_24 = (function() use ($Apply0_8_22, $f_6, &$go__go_9_24) {
  $__fn = function($b_10, $v_11 = null) use ($Apply0_8_22, $f_6, &$go__go_9_24, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_9_24_24_b_10 = $b_10;
  $__tco_var_go__go_9_24_24_v_11 = $v_11;
  tco_loop_go__go_9_24_24:;
  $b_10 = $__tco_var_go__go_9_24_24_b_10;
  $v_11 = $__tco_var_go__go_9_24_24_v_11;
  $__t24 = null;;
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t24 = $b_10;
goto end_branch_24;;
};
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_12_25 = (($Apply0_8_22)->{'Functor0'})(null);
$__tco_26 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_13) use ($Apply0_8_22, $Functor0_12_25, $b_10) {
  $__num = \func_num_args();
  $__res = ((($Apply0_8_22)->{'apply'})(((($Functor0_12_25)->{'map'})(function($b_14) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($b_14) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_15, $b_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_10)))($b_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $f_6, ($v_11)->{'value0'});
$__tco_27 = ($v_11)->{'value1'};
$__tco_var_go__go_9_24_24_b_10 = $__tco_26;
$__tco_var_go__go_9_24_24_v_11 = $__tco_27;
goto tco_loop_go__go_9_24_24;;
$__t24 = null;
goto end_branch_24;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t24 = null;
  end_branch_24:;
  $__res = $__t24;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Apply0_4_20)->{'apply'})(((($Functor0_5_21)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($f_6)(($v_7)->{'value0'}))))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((((((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($go__go_9_23)(new \Data\List\Types\Data_List_Types_Nil())), ($go__go_9_24)((($dictApplicative_3)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())), ($v_7)->{'value1'}));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorNonEmpty1_1_7) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_1_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_3) use ($foldableNonEmpty1_2_8) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_2_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_unfoldable1List
$GLOBALS['Data_List_Types_unfoldable1List'] = (object)["unfoldr1" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($f_0, &$go__go_2_0) {
  $__fn = function($source_3, $memo_4 = null) use ($f_0, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_source_3 = $source_3;
  $__tco_var_go__go_2_0_0_memo_4 = $memo_4;
  tco_loop_go__go_2_0_0:;
  $source_3 = $__tco_var_go__go_2_0_0_source_3;
  $memo_4 = $__tco_var_go__go_2_0_0_memo_4;
  $v_5_0 = ($f_0)($source_3);
  $__t1 = null;;
  if (($v_5_0)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_2 = (($v_5_0)->{'value1'})->{'value0'};
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v_5_0)->{'value0'}, $memo_4);
$__tco_var_go__go_2_0_0_source_3 = $__tco_2;
$__tco_var_go__go_2_0_0_memo_4 = $__tco_3;
goto tco_loop_go__go_2_0_0;;
$__t1 = null;
goto end_branch_1;;
};
  if (($v_5_0)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$go__go_6_4 = null;
$go__go_6_4 = (function() use (&$__tco_var_go__go_2_0_0_source_3, &$__tco_var_go__go_2_0_0_memo_4, &$go__go_6_4) {
  $__fn = function($b_7, $v_8 = null) use (&$__tco_var_go__go_2_0_0_source_3, &$__tco_var_go__go_2_0_0_memo_4, &$go__go_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_4_4_b_7 = $b_7;
  $__tco_var_go__go_6_4_4_v_8 = $v_8;
  tco_loop_go__go_6_4_4:;
  $b_7 = $__tco_var_go__go_6_4_4_b_7;
  $v_8 = $__tco_var_go__go_6_4_4_v_8;
  $__t4 = null;;
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $b_7;
goto end_branch_4;;
};
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons(($v_8)->{'value0'}, $b_7);
$__tco_6 = ($v_8)->{'value1'};
$__tco_var_go__go_6_4_4_b_7 = $__tco_5;
$__tco_var_go__go_6_4_4_v_8 = $__tco_6;
goto tco_loop_go__go_6_4_4;;
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
$__t1 = (($go__go_6_4)(new \Data\List\Types\Data_List_Types_Nil()))(new \Data\List\Types\Data_List_Types_Cons(($v_5_0)->{'value0'}, $memo_4));
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
  $__res = (($go__go_2_0)($b_1))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_unfoldableList
$GLOBALS['Data_List_Types_unfoldableList'] = (object)["unfoldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($f_0, &$go__go_2_0) {
  $__fn = function($source_3, $memo_4 = null) use ($f_0, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_source_3 = $source_3;
  $__tco_var_go__go_2_0_0_memo_4 = $memo_4;
  tco_loop_go__go_2_0_0:;
  $source_3 = $__tco_var_go__go_2_0_0_source_3;
  $memo_4 = $__tco_var_go__go_2_0_0_memo_4;
  $v_5_0 = ($f_0)($source_3);
  $__t1 = null;;
  if ($v_5_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$go__go_6_2 = null;
$go__go_6_2 = (function() use (&$__tco_var_go__go_2_0_0_source_3, &$__tco_var_go__go_2_0_0_memo_4, &$go__go_6_2) {
  $__fn = function($b_7, $v_8 = null) use (&$__tco_var_go__go_2_0_0_source_3, &$__tco_var_go__go_2_0_0_memo_4, &$go__go_6_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_2_2_b_7 = $b_7;
  $__tco_var_go__go_6_2_2_v_8 = $v_8;
  tco_loop_go__go_6_2_2:;
  $b_7 = $__tco_var_go__go_6_2_2_b_7;
  $v_8 = $__tco_var_go__go_6_2_2_v_8;
  $__t2 = null;;
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_7;
goto end_branch_2;;
};
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v_8)->{'value0'}, $b_7);
$__tco_4 = ($v_8)->{'value1'};
$__tco_var_go__go_6_2_2_b_7 = $__tco_3;
$__tco_var_go__go_6_2_2_v_8 = $__tco_4;
goto tco_loop_go__go_6_2_2;;
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
$__t1 = (($go__go_6_2)(new \Data\List\Types\Data_List_Types_Nil()))($memo_4);
goto end_branch_1;;
};
  if ($v_5_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_3 = (($v_5_0)->{'value0'})->{'value1'};
$__tco_4 = new \Data\List\Types\Data_List_Types_Cons((($v_5_0)->{'value0'})->{'value0'}, $memo_4);
$__tco_var_go__go_2_0_0_source_3 = $__tco_3;
$__tco_var_go__go_2_0_0_memo_4 = $__tco_4;
goto tco_loop_go__go_2_0_0;;
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
  $__res = (($go__go_2_0)($b_1))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Unfoldable10" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_unfoldable1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_unfoldable1NonEmptyList
$GLOBALS['Data_List_Types_unfoldable1NonEmptyList'] = (object)["unfoldr1" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = ($f_0)($b_1);
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use ($f_0, &$go__go_3_1) {
  $__fn = function($source_4, $memo_5 = null) use ($f_0, &$go__go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_1_1_source_4 = $source_4;
  $__tco_var_go__go_3_1_1_memo_5 = $memo_5;
  tco_loop_go__go_3_1_1:;
  $source_4 = $__tco_var_go__go_3_1_1_source_4;
  $memo_5 = $__tco_var_go__go_3_1_1_memo_5;
  $__t2 = null;;
  if ($source_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_3 = (($f_0)(($source_4)->{'value0'}))->{'value1'};
$__tco_4 = new \Data\List\Types\Data_List_Types_Cons((($f_0)(($source_4)->{'value0'}))->{'value0'}, $memo_5);
$__tco_var_go__go_3_1_1_source_4 = $__tco_3;
$__tco_var_go__go_3_1_1_memo_5 = $__tco_4;
goto tco_loop_go__go_3_1_1;;
$__t2 = null;
goto end_branch_2;;
};
  $go__go_6_1 = null;
  $go__go_6_1 = (function() use (&$__tco_var_go__go_3_1_1_source_4, &$__tco_var_go__go_3_1_1_memo_5, &$go__go_6_1) {
  $__fn = function($b_7, $v_8 = null) use (&$__tco_var_go__go_3_1_1_source_4, &$__tco_var_go__go_3_1_1_memo_5, &$go__go_6_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_1_1_b_7 = $b_7;
  $__tco_var_go__go_6_1_1_v_8 = $v_8;
  tco_loop_go__go_6_1_1:;
  $b_7 = $__tco_var_go__go_6_1_1_b_7;
  $v_8 = $__tco_var_go__go_6_1_1_v_8;
  $__t1 = null;;
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_7;
goto end_branch_1;;
};
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v_8)->{'value0'}, $b_7);
$__tco_3 = ($v_8)->{'value1'};
$__tco_var_go__go_6_1_1_b_7 = $__tco_2;
$__tco_var_go__go_6_1_1_v_8 = $__tco_3;
goto tco_loop_go__go_6_1_1;;
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
  $__t2 = (($go__go_6_1)(new \Data\List\Types\Data_List_Types_Nil()))($memo_5);
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__local_var_2_0 = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_2_0)->{'value0'}, (($go__go_3_1)(($__local_var_2_0)->{'value1'}))(new \Data\List\Types\Data_List_Types_Nil()));
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($__local_var_2_0)->{'value0'}, ($__local_var_2_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_foldable1NonEmptyList
$GLOBALS['Data_List_Types_foldable1NonEmptyList'] = (function() use (&$__fn) {
$foldableNonEmpty1_0_0 = (object)["foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($f_2) use ($Semigroup0_1_0, $dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Semigroup0_1_0, $dictMonoid_0, $f_2) {
  $__num = \func_num_args();
  $Semigroup0_4_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $go__go_5_2 = null;
  $go__go_5_2 = (function() use ($Semigroup0_4_1, $f_2, &$go__go_5_2) {
  $__fn = function($b_6, $v_7 = null) use ($Semigroup0_4_1, $f_2, &$go__go_5_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_2_2_b_6 = $b_6;
  $__tco_var_go__go_5_2_2_v_7 = $v_7;
  tco_loop_go__go_5_2_2:;
  $b_6 = $__tco_var_go__go_5_2_2_b_6;
  $v_7 = $__tco_var_go__go_5_2_2_v_7;
  $__t2 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_6;
goto end_branch_2;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_4_1)->{'append'})($b_6), $f_2, ($v_7)->{'value0'});
$__tco_4 = ($v_7)->{'value1'};
$__tco_var_go__go_5_2_2_b_6 = $__tco_3;
$__tco_var_go__go_5_2_2_v_7 = $__tco_4;
goto tco_loop_go__go_5_2_2;;
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
  $__res = ((($Semigroup0_1_0)->{'append'})(($f_2)(($v_3)->{'value0'})))((($go__go_5_2)(($dictMonoid_0)->{'mempty'}))(($v_3)->{'value1'}));
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
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_3 = null;
  $go__go_3_3 = (function() use ($f_0, &$go__go_3_3) {
  $__fn = function($b_4, $v_5 = null) use ($f_0, &$go__go_3_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_3_3_b_4 = $b_4;
  $__tco_var_go__go_3_3_3_v_5 = $v_5;
  tco_loop_go__go_3_3_3:;
  $b_4 = $__tco_var_go__go_3_3_3_b_4;
  $v_5 = $__tco_var_go__go_3_3_3_v_5;
  $__t3 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $b_4;
goto end_branch_3;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = (($f_0)($b_4))(($v_5)->{'value0'});
$__tco_5 = ($v_5)->{'value1'};
$__tco_var_go__go_3_3_3_b_4 = $__tco_4;
$__tco_var_go__go_3_3_3_v_5 = $__tco_5;
goto tco_loop_go__go_3_3_3;;
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
  $__res = (($go__go_3_3)((($f_0)($b_1))(($v_2)->{'value0'})))(($v_2)->{'value1'});
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
}, "foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($b_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($b_1, $f_0) {
  $__num = \func_num_args();
  $go__go_3_4 = null;
  $go__go_3_4 = (function() use ($f_0, &$go__go_3_4) {
  $__fn = function($b_4, $v_5 = null) use ($f_0, &$go__go_3_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_4_4_b_4 = $b_4;
  $__tco_var_go__go_3_4_4_v_5 = $v_5;
  tco_loop_go__go_3_4_4:;
  $b_4 = $__tco_var_go__go_3_4_4_b_4;
  $v_5 = $__tco_var_go__go_3_4_4_v_5;
  $__t4 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $b_4;
goto end_branch_4;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = (($f_0)(($v_5)->{'value0'}))($b_4);
$__tco_6 = ($v_5)->{'value1'};
$__tco_var_go__go_3_4_4_b_4 = $__tco_5;
$__tco_var_go__go_3_4_4_v_5 = $__tco_6;
goto tco_loop_go__go_3_4_4;;
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
  $go__go_3_5 = null;
  $go__go_3_5 = (function() use (&$go__go_3_5) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__go_3_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_5_5_v_4 = $v_4;
  $__tco_var_go__go_3_5_5_v1_5 = $v1_5;
  tco_loop_go__go_3_5_5:;
  $v_4 = $__tco_var_go__go_3_5_5_v_4;
  $v1_5 = $__tco_var_go__go_3_5_5_v1_5;
  $__t5 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $v_4;
goto end_branch_5;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_7 = ($v1_5)->{'value1'};
$__tco_var_go__go_3_5_5_v_4 = $__tco_6;
$__tco_var_go__go_3_5_5_v1_5 = $__tco_7;
goto tco_loop_go__go_3_5_5;;
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
  $__res = (($f_0)(($v_2)->{'value0'}))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_3_4)($b_1), ($go__go_3_5)(new \Data\List\Types\Data_List_Types_Nil()), ($v_2)->{'value1'}));
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
return (object)["foldMap1" => function($dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = function($f_2) use ($dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictSemigroup_1, $f_2) {
  $__num = \func_num_args();
  $go__go_4_7 = null;
  $go__go_4_7 = (function() use ($dictSemigroup_1, $f_2, &$go__go_4_7) {
  $__fn = function($b_5, $v_6 = null) use ($dictSemigroup_1, $f_2, &$go__go_4_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_7_7_b_5 = $b_5;
  $__tco_var_go__go_4_7_7_v_6 = $v_6;
  tco_loop_go__go_4_7_7:;
  $b_5 = $__tco_var_go__go_4_7_7_b_5;
  $v_6 = $__tco_var_go__go_4_7_7_v_6;
  $__t7 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t7 = $b_5;
goto end_branch_7;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_8 = ((($dictSemigroup_1)->{'append'})($b_5))(($f_2)(($v_6)->{'value0'}));
$__tco_9 = ($v_6)->{'value1'};
$__tco_var_go__go_4_7_7_b_5 = $__tco_8;
$__tco_var_go__go_4_7_7_v_6 = $__tco_9;
goto tco_loop_go__go_4_7_7;;
$__t7 = null;
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_4_7)(($f_2)(($v_3)->{'value0'})))(($v_3)->{'value1'});
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
}, "foldr1" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__local_var_3_8 = ($f_1)(($v_2)->{'value0'});
  $go__go_4_9 = null;
  $go__go_4_9 = (function() use ($f_1, &$go__go_4_9) {
  $__fn = function($b_5, $v_6 = null) use ($f_1, &$go__go_4_9, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_9_9_b_5 = $b_5;
  $__tco_var_go__go_4_9_9_v_6 = $v_6;
  tco_loop_go__go_4_9_9:;
  $b_5 = $__tco_var_go__go_4_9_9_b_5;
  $v_6 = $__tco_var_go__go_4_9_9_v_6;
  $__t9 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t9 = $b_5;
goto end_branch_9;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_7_10 = ($f_1)(($v_6)->{'value0'});
$__tco_12 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_Just'], function($v2_8) use ($__local_var_7_10, $v_6) {
  $__num = \func_num_args();
  $__t11 = null;;
  if ($v2_8 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t11 = ($v_6)->{'value0'};
goto end_branch_11;;
};
  if ($v2_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t11 = ($__local_var_7_10)(($v2_8)->{'value0'});
goto end_branch_11;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t11 = null;
  end_branch_11:;
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $b_5);
$__tco_13 = ($v_6)->{'value1'};
$__tco_var_go__go_4_9_9_b_5 = $__tco_12;
$__tco_var_go__go_4_9_9_v_6 = $__tco_13;
goto tco_loop_go__go_4_9_9;;
$__t9 = null;
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_4_10 = null;
  $go__go_4_10 = (function() use (&$go__go_4_10) {
  $__fn = function($v_5, $v1_6 = null) use (&$go__go_4_10, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_10_10_v_5 = $v_5;
  $__tco_var_go__go_4_10_10_v1_6 = $v1_6;
  tco_loop_go__go_4_10_10:;
  $v_5 = $__tco_var_go__go_4_10_10_v_5;
  $v1_6 = $__tco_var_go__go_4_10_10_v1_6;
  $__t10 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t10 = $v_5;
goto end_branch_10;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_11 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_12 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_10_10_v_5 = $__tco_11;
$__tco_var_go__go_4_10_10_v1_6 = $__tco_12;
goto tco_loop_go__go_4_10_10;;
$__t10 = null;
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__local_var_4_9 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_4_9)(new \Data\Maybe\Data_Maybe_Nothing()), ($go__go_4_10)(new \Data\List\Types\Data_List_Types_Nil()), ($v_2)->{'value1'});
  $__t12 = null;;
  if ($__local_var_4_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t12 = ($v_2)->{'value0'};
goto end_branch_12;;
};
  if ($__local_var_4_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t12 = ($__local_var_3_8)(($__local_var_4_9)->{'value0'});
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $go__go_3_13 = null;
  $go__go_3_13 = (function() use ($f_1, &$go__go_3_13) {
  $__fn = function($b_4, $v_5 = null) use ($f_1, &$go__go_3_13, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_13_13_b_4 = $b_4;
  $__tco_var_go__go_3_13_13_v_5 = $v_5;
  tco_loop_go__go_3_13_13:;
  $b_4 = $__tco_var_go__go_3_13_13_b_4;
  $v_5 = $__tco_var_go__go_3_13_13_v_5;
  $__t13 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t13 = $b_4;
goto end_branch_13;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_14 = (($f_1)($b_4))(($v_5)->{'value0'});
$__tco_15 = ($v_5)->{'value1'};
$__tco_var_go__go_3_13_13_b_4 = $__tco_14;
$__tco_var_go__go_3_13_13_v_5 = $__tco_15;
goto tco_loop_go__go_3_13_13;;
$__t13 = null;
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_3_13)(($v_2)->{'value0'}))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_1) use ($foldableNonEmpty1_0_0) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_List_Types_extendNonEmptyList
$GLOBALS['Data_List_Types_extendNonEmptyList'] = (object)["extend" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($f_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use ($f_0, &$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use ($f_0, &$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = (object)["val" => new \Data\List\Types\Data_List_Types_Cons(($f_0)(new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_4)->{'value0'}, ($b_3)->{'acc'})), ($b_3)->{'val'}), "acc" => new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, ($b_3)->{'acc'})];
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $go__go_2_1 = null;
  $go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_v_3 = $v_3;
  $__tco_var_go__go_2_1_1_v1_4 = $v1_4;
  tco_loop_go__go_2_1_1:;
  $v_3 = $__tco_var_go__go_2_1_1_v_3;
  $v1_4 = $__tco_var_go__go_2_1_1_v1_4;
  $__t1 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_3;
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_3 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_1_1_v_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v1_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)($v_1), (\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_0)((object)["val" => new \Data\List\Types\Data_List_Types_Nil(), "acc" => new \Data\List\Types\Data_List_Types_Nil()]), ($go__go_2_1)(new \Data\List\Types\Data_List_Types_Nil()), ($v_1)->{'value1'}))->{'val'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_extendList
$GLOBALS['Data_List_Types_extendList'] = (object)["extend" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$go__go_2_1 = null;
$go__go_2_1 = (function() use (&$go__go_2_1, $v_0) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_1, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_b_3 = $b_3;
  $__tco_var_go__go_2_1_1_v_4 = $v_4;
  tco_loop_go__go_2_1_1:;
  $b_3 = $__tco_var_go__go_2_1_1_b_3;
  $v_4 = $__tco_var_go__go_2_1_1_v_4;
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_3;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = (object)["val" => new \Data\List\Types\Data_List_Types_Cons(($v_0)(new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, ($b_3)->{'acc'})), ($b_3)->{'val'}), "acc" => new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, ($b_3)->{'acc'})];
$__tco_3 = ($v_4)->{'value1'};
$__tco_var_go__go_2_1_1_b_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
$go__go_2_2 = null;
$go__go_2_2 = (function() use (&$go__go_2_2) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_2_2_v_3 = $v_3;
  $__tco_var_go__go_2_2_2_v1_4 = $v1_4;
  tco_loop_go__go_2_2_2:;
  $v_3 = $__tco_var_go__go_2_2_2_v_3;
  $v1_4 = $__tco_var_go__go_2_2_2_v1_4;
  $__t2 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_3;
goto end_branch_2;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_4 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_2_2_v_3 = $__tco_3;
$__tco_var_go__go_2_2_2_v1_4 = $__tco_4;
goto tco_loop_go__go_2_2_2;;
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
$__t0 = new \Data\List\Types\Data_List_Types_Cons(($v_0)($v1_1), (\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_1)((object)["val" => new \Data\List\Types\Data_List_Types_Nil(), "acc" => new \Data\List\Types\Data_List_Types_Nil()]), ($go__go_2_2)(new \Data\List\Types\Data_List_Types_Nil()), ($v1_1)->{'value1'}))->{'val'});
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
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_eq1List
$GLOBALS['Data_List_Types_eq1List'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($xs_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($ys_2) use ($dictEq_0, $xs_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__res = function($v2_6) use ($dictEq_0, &$go__go_3_0, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (( ! $v2_6)) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_6);
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_3_0)(($v_4)->{'value1'}))(($v1_5)->{'value1'}))(($v2_6 && ((($dictEq_0)->{'eq'})(($v1_5)->{'value0'}))(($v_4)->{'value0'})))));
  end_branch_1:;
  $__res = $__t1;
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
  $__res = ((($go__go_3_0)($xs_1))($ys_2))(true);
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

// Data_List_Types_eq1NonEmptyList
$GLOBALS['Data_List_Types_eq1NonEmptyList'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($x_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__res = function($v2_6) use ($dictEq_0, &$go__go_3_0, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (( ! $v2_6)) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_6);
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_3_0)(($v_4)->{'value1'}))(($v1_5)->{'value1'}))(($v2_6 && ((($dictEq_0)->{'eq'})(($v1_5)->{'value0'}))(($v_4)->{'value0'})))));
  end_branch_1:;
  $__res = $__t1;
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
  $__res = (((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'}) && ((($go__go_3_0)(($x_1)->{'value1'}))(($y_2)->{'value1'}))(true));
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

// Data_List_Types_eqList
function majData_majList_majTypes_eqmajList($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_eqmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq" => function($xs_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($ys_2) use ($dictEq_0, $xs_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__res = function($v2_6) use ($dictEq_0, &$go__go_3_0, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (( ! $v2_6)) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_6);
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_3_0)(($v_4)->{'value1'}))(($v1_5)->{'value1'}))(($v2_6 && ((($dictEq_0)->{'eq'})(($v1_5)->{'value0'}))(($v_4)->{'value0'})))));
  end_branch_1:;
  $__res = $__t1;
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
  $__res = ((($go__go_3_0)($xs_1))($ys_2))(true);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_eqList'] = __NAMESPACE__ . '\\majData_majList_majTypes_eqmajList';

// Data_List_Types_eqNonEmptyList
function majData_majList_majTypes_eqmajNonmajEmptymajList($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_eqmajNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq" => function($x_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__res = function($v2_6) use ($dictEq_0, &$go__go_3_0, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (( ! $v2_6)) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_6);
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_3_0)(($v_4)->{'value1'}))(($v1_5)->{'value1'}))(($v2_6 && ((($dictEq_0)->{'eq'})(($v1_5)->{'value0'}))(($v_4)->{'value0'})))));
  end_branch_1:;
  $__res = $__t1;
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
  $__res = (((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'}) && ((($go__go_3_0)(($x_1)->{'value1'}))(($y_2)->{'value1'}))(true));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_eqNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majTypes_eqmajNonmajEmptymajList';

// Data_List_Types_ord1List
$GLOBALS['Data_List_Types_ord1List'] = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($xs_1) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($ys_2) use ($dictOrd_0, $xs_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = (function() use ($dictOrd_0, &$go__go_3_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictOrd_0, &$go__go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_0_0_v_4 = $v_4;
  $__tco_var_go__go_3_0_0_v1_5 = $v1_5;
  tco_loop_go__go_3_0_0:;
  $v_4 = $__tco_var_go__go_3_0_0_v_4;
  $v1_5 = $__tco_var_go__go_3_0_0_v1_5;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = null;;
if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_0;;
};
  if (($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_5 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$v2_6_2 = ((($dictOrd_0)->{'compare'})(($v_4)->{'value0'}))(($v1_5)->{'value0'});
$__t3 = null;;
if ($v2_6_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_4 = ($v_4)->{'value1'};
$__tco_5 = ($v1_5)->{'value1'};
$__tco_var_go__go_3_0_0_v_4 = $__tco_4;
$__tco_var_go__go_3_0_0_v1_5 = $__tco_5;
goto tco_loop_go__go_3_0_0;;
$__t3 = null;
goto end_branch_3;;
};
$__t3 = $v2_6_2;
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
  $__res = (($go__go_3_0)($xs_1))($ys_2);
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
}, "Eq10" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_eq1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_ord1NonEmptyList
$GLOBALS['Data_List_Types_ord1NonEmptyList'] = (function() use (&$__fn) {
$eq1NonEmpty1_0_0 = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($x_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = function($y_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $go__go_3_0 = null;
  $go__go_3_0 = function($v_4) use ($dictEq_0, &$go__go_3_0) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($dictEq_0, &$go__go_3_0, $v_4) {
  $__num = \func_num_args();
  $__res = function($v2_6) use ($dictEq_0, &$go__go_3_0, $v1_5, $v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (( ! $v2_6)) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_6);
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_3_0)(($v_4)->{'value1'}))(($v1_5)->{'value1'}))(($v2_6 && ((($dictEq_0)->{'eq'})(($v1_5)->{'value0'}))(($v_4)->{'value0'})))));
  end_branch_1:;
  $__res = $__t1;
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
  $__res = (((($dictEq_0)->{'eq'})(($x_1)->{'value0'}))(($y_2)->{'value0'}) && ((($go__go_3_0)(($x_1)->{'value1'}))(($y_2)->{'value1'}))(true));
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
return (object)["compare1" => function($dictOrd_1) {
  $__num = \func_num_args();
  $__res = function($x_2) use ($dictOrd_1) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictOrd_1, $x_2) {
  $__num = \func_num_args();
  $v_4_3 = ((($dictOrd_1)->{'compare'})(($x_2)->{'value0'}))(($y_3)->{'value0'});
  $__t5 = null;;
  if ($v_4_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t5 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_5;;
};
  if ($v_4_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  $go__go_5_4 = null;
  $go__go_5_4 = (function() use ($dictOrd_1, &$go__go_5_4) {
  $__fn = function($v_6, $v1_7 = null) use ($dictOrd_1, &$go__go_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_4_4_v_6 = $v_6;
  $__tco_var_go__go_5_4_4_v1_7 = $v1_7;
  tco_loop_go__go_5_4_4:;
  $v_6 = $__tco_var_go__go_5_4_4_v_6;
  $v1_7 = $__tco_var_go__go_5_4_4_v1_7;
  $__t4 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = null;;
if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($v_6 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_7 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$v2_8_6 = ((($dictOrd_1)->{'compare'})(($v_6)->{'value0'}))(($v1_7)->{'value0'});
$__t7 = null;;
if ($v2_8_6 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_8 = ($v_6)->{'value1'};
$__tco_9 = ($v1_7)->{'value1'};
$__tco_var_go__go_5_4_4_v_6 = $__tco_8;
$__tco_var_go__go_5_4_4_v1_7 = $__tco_9;
goto tco_loop_go__go_5_4_4;;
$__t7 = null;
goto end_branch_7;;
};
$__t7 = $v2_8_6;
end_branch_7:;
$__t4 = $__t7;
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
  $__t5 = (($go__go_5_4)(($x_2)->{'value1'}))(($y_3)->{'value1'});
  end_branch_5:;
  $__res = $__t5;
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
}, "Eq10" => function($_dollar___unused_1) use ($eq1NonEmpty1_0_0) {
  $__num = \func_num_args();
  $__res = $eq1NonEmpty1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Data_List_Types_ordList
function majData_majList_majTypes_ordmajList($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_ordmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqList1_1_0 = (object)["eq" => function($xs_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($__local_var_1_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = function($v_5) use ($__local_var_1_0, &$go__go_4_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, &$go__go_4_1, $v_5) {
  $__num = \func_num_args();
  $__res = function($v2_7) use ($__local_var_1_0, &$go__go_4_1, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (( ! $v2_7)) {
$__t2 = false;
goto end_branch_2;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_7);
goto end_branch_2;;
};
  $__t2 = ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_4_1)(($v_5)->{'value1'}))(($v1_6)->{'value1'}))(($v2_7 && ((($__local_var_1_0)->{'eq'})(($v1_6)->{'value0'}))(($v_5)->{'value0'})))));
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
  $__res = ((($go__go_4_1)($xs_2))($ys_3))(true);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($xs_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($ys_3) use ($dictOrd_0, $xs_2) {
  $__num = \func_num_args();
  $go__go_4_4 = null;
  $go__go_4_4 = (function() use ($dictOrd_0, &$go__go_4_4) {
  $__fn = function($v_5, $v1_6 = null) use ($dictOrd_0, &$go__go_4_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_4_4_v_5 = $v_5;
  $__tco_var_go__go_4_4_4_v1_6 = $v1_6;
  tco_loop_go__go_4_4_4:;
  $v_5 = $__tco_var_go__go_4_4_4_v_5;
  $v1_6 = $__tco_var_go__go_4_4_4_v1_6;
  $__t4 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = null;;
if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_5;;
};
$__t5 = new \Data\Ordering\Data_Ordering_LT();
end_branch_5:;
$__t4 = $__t5;
goto end_branch_4;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_4;;
};
  if (($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_6 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$v2_7_6 = ((($dictOrd_0)->{'compare'})(($v_5)->{'value0'}))(($v1_6)->{'value0'});
$__t7 = null;;
if ($v2_7_6 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_8 = ($v_5)->{'value1'};
$__tco_9 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_4_4_v_5 = $__tco_8;
$__tco_var_go__go_4_4_4_v1_6 = $__tco_9;
goto tco_loop_go__go_4_4_4;;
$__t7 = null;
goto end_branch_7;;
};
$__t7 = $v2_7_6;
end_branch_7:;
$__t4 = $__t7;
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
  $__res = (($go__go_4_4)($xs_2))($ys_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_2) use ($eqList1_1_0) {
  $__num = \func_num_args();
  $__res = $eqList1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_ordList'] = __NAMESPACE__ . '\\majData_majList_majTypes_ordmajList';

// Data_List_Types_ordNonEmptyList
function majData_majList_majTypes_ordmajNonmajEmptymajList($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_ordmajNonmajEmptymajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictOrd_0)->{'Eq0'})(null);
  $eqNonEmpty2_1_0 = (object)["eq" => function($x_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($__local_var_1_0, $x_2) {
  $__num = \func_num_args();
  $go__go_4_1 = null;
  $go__go_4_1 = function($v_5) use ($__local_var_1_0, &$go__go_4_1) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($__local_var_1_0, &$go__go_4_1, $v_5) {
  $__num = \func_num_args();
  $__res = function($v2_7) use ($__local_var_1_0, &$go__go_4_1, $v1_6, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (( ! $v2_7)) {
$__t2 = false;
goto end_branch_2;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_7);
goto end_branch_2;;
};
  $__t2 = ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__go_4_1)(($v_5)->{'value1'}))(($v1_6)->{'value1'}))(($v2_7 && ((($__local_var_1_0)->{'eq'})(($v1_6)->{'value0'}))(($v_5)->{'value0'})))));
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
  $__res = (((($__local_var_1_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'}) && ((($go__go_4_1)(($x_2)->{'value1'}))(($y_3)->{'value1'}))(true));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare" => function($x_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictOrd_0, $x_2) {
  $__num = \func_num_args();
  $v_4_4 = ((($dictOrd_0)->{'compare'})(($x_2)->{'value0'}))(($y_3)->{'value0'});
  $__t6 = null;;
  if ($v_4_4 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t6 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_6;;
};
  if ($v_4_4 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t6 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_6;;
};
  $go__go_5_5 = null;
  $go__go_5_5 = (function() use ($dictOrd_0, &$go__go_5_5) {
  $__fn = function($v_6, $v1_7 = null) use ($dictOrd_0, &$go__go_5_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_5_5_v_6 = $v_6;
  $__tco_var_go__go_5_5_5_v1_7 = $v1_7;
  tco_loop_go__go_5_5_5:;
  $v_6 = $__tco_var_go__go_5_5_5_v_6;
  $v1_7 = $__tco_var_go__go_5_5_5_v1_7;
  $__t5 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = null;;
if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_6;;
};
$__t6 = new \Data\Ordering\Data_Ordering_LT();
end_branch_6:;
$__t5 = $__t6;
goto end_branch_5;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_5;;
};
  if (($v_6 instanceof \Data\List\Types\Data_List_Types_Cons && $v1_7 instanceof \Data\List\Types\Data_List_Types_Cons)) {
$v2_8_7 = ((($dictOrd_0)->{'compare'})(($v_6)->{'value0'}))(($v1_7)->{'value0'});
$__t8 = null;;
if ($v2_8_7 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__tco_9 = ($v_6)->{'value1'};
$__tco_10 = ($v1_7)->{'value1'};
$__tco_var_go__go_5_5_5_v_6 = $__tco_9;
$__tco_var_go__go_5_5_5_v1_7 = $__tco_10;
goto tco_loop_go__go_5_5_5;;
$__t8 = null;
goto end_branch_8;;
};
$__t8 = $v2_8_7;
end_branch_8:;
$__t5 = $__t8;
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
  $__t6 = (($go__go_5_5)(($x_2)->{'value1'}))(($y_3)->{'value1'});
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq0" => function($_dollar___unused_2) use ($eqNonEmpty2_1_0) {
  $__num = \func_num_args();
  $__res = $eqNonEmpty2_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_ordNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majTypes_ordmajNonmajEmptymajList';

// Data_List_Types_comonadNonEmptyList
$GLOBALS['Data_List_Types_comonadNonEmptyList'] = (object)["extract" => function($v_0) {
  $__num = \func_num_args();
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_extendNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_applyList
$GLOBALS['Data_List_Types_applyList'] = (object)["apply" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$go__go_2_1 = null;
$go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_b_3 = $b_3;
  $__tco_var_go__go_2_1_1_v_4 = $v_4;
  tco_loop_go__go_2_1_1:;
  $b_3 = $__tco_var_go__go_2_1_1_b_3;
  $v_4 = $__tco_var_go__go_2_1_1_v_4;
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_3;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_3 = ($v_4)->{'value1'};
$__tco_var_go__go_2_1_1_b_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
$go__go_2_2 = null;
$go__go_2_2 = (function() use (&$go__go_2_2) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_2_2_v_3 = $v_3;
  $__tco_var_go__go_2_2_2_v1_4 = $v1_4;
  tco_loop_go__go_2_2_2:;
  $v_3 = $__tco_var_go__go_2_2_2_v_3;
  $v1_4 = $__tco_var_go__go_2_2_2_v1_4;
  $__t2 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_3;
goto end_branch_2;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_4 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_2_2_v_3 = $__tco_3;
$__tco_var_go__go_2_2_2_v1_4 = $__tco_4;
goto tco_loop_go__go_2_2_2;;
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
$__t0 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_1)(((($GLOBALS['Data_List_Types_applyList'])->{'apply'})(($v_0)->{'value1'}))($v1_1)), ($go__go_2_2)(new \Data\List\Types\Data_List_Types_Nil()), \Data\List\Types\majData_majList_majTypes_listmajMap(($v_0)->{'value0'}, $v1_1));
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
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_applyNonEmptyList
$GLOBALS['Data_List_Types_applyNonEmptyList'] = (object)["apply" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $__local_var_3_1 = new \Data\List\Types\Data_List_Types_Cons(($v_0)->{'value0'}, ($v_0)->{'value1'});
  $go__go_4_2 = null;
  $go__go_4_2 = (function() use (&$go__go_4_2) {
  $__fn = function($b_5, $v_6 = null) use (&$go__go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_2_2_b_5 = $b_5;
  $__tco_var_go__go_4_2_2_v_6 = $v_6;
  tco_loop_go__go_4_2_2:;
  $b_5 = $__tco_var_go__go_4_2_2_b_5;
  $v_6 = $__tco_var_go__go_4_2_2_v_6;
  $__t2 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_5;
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v_6)->{'value0'}, $b_5);
$__tco_4 = ($v_6)->{'value1'};
$__tco_var_go__go_4_2_2_b_5 = $__tco_3;
$__tco_var_go__go_4_2_2_v_6 = $__tco_4;
goto tco_loop_go__go_4_2_2;;
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
  $go__go_4_3 = null;
  $go__go_4_3 = (function() use (&$go__go_4_3) {
  $__fn = function($v_5, $v1_6 = null) use (&$go__go_4_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_3_3_v_5 = $v_5;
  $__tco_var_go__go_4_3_3_v1_6 = $v1_6;
  tco_loop_go__go_4_3_3:;
  $v_5 = $__tco_var_go__go_4_3_3_v_5;
  $v1_6 = $__tco_var_go__go_4_3_3_v1_6;
  $__t3 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = $v_5;
goto end_branch_3;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_4 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_5 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_3_3_v_5 = $__tco_4;
$__tco_var_go__go_4_3_3_v1_6 = $__tco_5;
goto tco_loop_go__go_4_3_3;;
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
  $go__go_2_4 = null;
  $go__go_2_4 = (function() use (&$go__go_2_4) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_4_4_v_3 = $v_3;
  $__tco_var_go__go_2_4_4_v1_4 = $v1_4;
  tco_loop_go__go_2_4_4:;
  $v_3 = $__tco_var_go__go_2_4_4_v_3;
  $v1_4 = $__tco_var_go__go_2_4_4_v1_4;
  $__t4 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $v_3;
goto end_branch_4;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_6 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_4_4_v_3 = $__tco_5;
$__tco_var_go__go_2_4_4_v1_4 = $__tco_6;
goto tco_loop_go__go_2_4_4;;
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
  $__t5 = null;;
  if (($v_0)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_5;;
};
  if (($v_0)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons) {
$go__go_2_6 = null;
$go__go_2_6 = (function() use (&$go__go_2_6) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_6_6_b_3 = $b_3;
  $__tco_var_go__go_2_6_6_v_4 = $v_4;
  tco_loop_go__go_2_6_6:;
  $b_3 = $__tco_var_go__go_2_6_6_b_3;
  $v_4 = $__tco_var_go__go_2_6_6_v_4;
  $__t6 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = $b_3;
goto end_branch_6;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_7 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_8 = ($v_4)->{'value1'};
$__tco_var_go__go_2_6_6_b_3 = $__tco_7;
$__tco_var_go__go_2_6_6_v_4 = $__tco_8;
goto tco_loop_go__go_2_6_6;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$go__go_2_7 = null;
$go__go_2_7 = (function() use (&$go__go_2_7) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_7_7_v_3 = $v_3;
  $__tco_var_go__go_2_7_7_v1_4 = $v1_4;
  tco_loop_go__go_2_7_7:;
  $v_3 = $__tco_var_go__go_2_7_7_v_3;
  $v1_4 = $__tco_var_go__go_2_7_7_v1_4;
  $__t7 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t7 = $v_3;
goto end_branch_7;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_8 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_9 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_7_7_v_3 = $__tco_8;
$__tco_var_go__go_2_7_7_v1_4 = $__tco_9;
goto tco_loop_go__go_2_7_7;;
$__t7 = null;
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t5 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_6)(((($GLOBALS['Data_List_Types_applyList'])->{'apply'})((($v_0)->{'value1'})->{'value1'}))(new \Data\List\Types\Data_List_Types_Cons(($v1_1)->{'value0'}, new \Data\List\Types\Data_List_Types_Nil()))), ($go__go_2_7)(new \Data\List\Types\Data_List_Types_Nil()), \Data\List\Types\majData_majList_majTypes_listmajMap((($v_0)->{'value1'})->{'value0'}, new \Data\List\Types\Data_List_Types_Cons(($v1_1)->{'value0'}, new \Data\List\Types\Data_List_Types_Nil())));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((($v_0)->{'value0'})(($v1_1)->{'value0'}), \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_0)(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_4_2)(((($GLOBALS['Data_List_Types_applyList'])->{'apply'})(($__local_var_3_1)->{'value1'}))(($v1_1)->{'value1'})), ($go__go_4_3)(new \Data\List\Types\Data_List_Types_Nil()), \Data\List\Types\majData_majList_majTypes_listmajMap(($__local_var_3_1)->{'value0'}, ($v1_1)->{'value1'}))), ($go__go_2_4)(new \Data\List\Types\Data_List_Types_Nil()), $__t5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_bindList
$GLOBALS['Data_List_Types_bindList'] = (object)["bind" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$go__go_2_1 = null;
$go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_b_3 = $b_3;
  $__tco_var_go__go_2_1_1_v_4 = $v_4;
  tco_loop_go__go_2_1_1:;
  $b_3 = $__tco_var_go__go_2_1_1_b_3;
  $v_4 = $__tco_var_go__go_2_1_1_v_4;
  $__t1 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_3;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_3 = ($v_4)->{'value1'};
$__tco_var_go__go_2_1_1_b_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
$go__go_2_2 = null;
$go__go_2_2 = (function() use (&$go__go_2_2) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_2_2_v_3 = $v_3;
  $__tco_var_go__go_2_2_2_v1_4 = $v1_4;
  tco_loop_go__go_2_2_2:;
  $v_3 = $__tco_var_go__go_2_2_2_v_3;
  $v1_4 = $__tco_var_go__go_2_2_2_v1_4;
  $__t2 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $v_3;
goto end_branch_2;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_4 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_2_2_v_3 = $__tco_3;
$__tco_var_go__go_2_2_2_v1_4 = $__tco_4;
goto tco_loop_go__go_2_2_2;;
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
$__t0 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_1)(((($GLOBALS['Data_List_Types_bindList'])->{'bind'})(($v_0)->{'value1'}))($v1_1)), ($go__go_2_2)(new \Data\List\Types\Data_List_Types_Nil()), ($v1_1)(($v_0)->{'value0'}));
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
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_bindNonEmptyList
$GLOBALS['Data_List_Types_bindNonEmptyList'] = (object)["bind" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($v_0) {
  $__num = \func_num_args();
  $v1_2_0 = ($f_1)(($v_0)->{'value0'});
  $go__go_3_1 = null;
  $go__go_3_1 = (function() use (&$go__go_3_1) {
  $__fn = function($b_4, $v_5 = null) use (&$go__go_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_1_1_b_4 = $b_4;
  $__tco_var_go__go_3_1_1_v_5 = $v_5;
  tco_loop_go__go_3_1_1:;
  $b_4 = $__tco_var_go__go_3_1_1_b_4;
  $v_5 = $__tco_var_go__go_3_1_1_v_5;
  $__t1 = null;;
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_4;
goto end_branch_1;;
};
  if ($v_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v_5)->{'value0'}, $b_4);
$__tco_3 = ($v_5)->{'value1'};
$__tco_var_go__go_3_1_1_b_4 = $__tco_2;
$__tco_var_go__go_3_1_1_v_5 = $__tco_3;
goto tco_loop_go__go_3_1_1;;
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
  $__local_var_4_2 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Types_toList']))($f_1);
  $__t3 = null;;
  if (($v_0)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t3 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_3;;
};
  if (($v_0)->{'value1'} instanceof \Data\List\Types\Data_List_Types_Cons) {
$go__go_5_4 = null;
$go__go_5_4 = (function() use (&$go__go_5_4) {
  $__fn = function($b_6, $v_7 = null) use (&$go__go_5_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_4_4_b_6 = $b_6;
  $__tco_var_go__go_5_4_4_v_7 = $v_7;
  tco_loop_go__go_5_4_4:;
  $b_6 = $__tco_var_go__go_5_4_4_b_6;
  $v_7 = $__tco_var_go__go_5_4_4_v_7;
  $__t4 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t4 = $b_6;
goto end_branch_4;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_5 = new \Data\List\Types\Data_List_Types_Cons(($v_7)->{'value0'}, $b_6);
$__tco_6 = ($v_7)->{'value1'};
$__tco_var_go__go_5_4_4_b_6 = $__tco_5;
$__tco_var_go__go_5_4_4_v_7 = $__tco_6;
goto tco_loop_go__go_5_4_4;;
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
$go__go_5_5 = null;
$go__go_5_5 = (function() use (&$go__go_5_5) {
  $__fn = function($v_6, $v1_7 = null) use (&$go__go_5_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_5_5_v_6 = $v_6;
  $__tco_var_go__go_5_5_5_v1_7 = $v1_7;
  tco_loop_go__go_5_5_5:;
  $v_6 = $__tco_var_go__go_5_5_5_v_6;
  $v1_7 = $__tco_var_go__go_5_5_5_v1_7;
  $__t5 = null;;
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $v_6;
goto end_branch_5;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = new \Data\List\Types\Data_List_Types_Cons(($v1_7)->{'value0'}, $v_6);
$__tco_7 = ($v1_7)->{'value1'};
$__tco_var_go__go_5_5_5_v_6 = $__tco_6;
$__tco_var_go__go_5_5_5_v1_7 = $__tco_7;
goto tco_loop_go__go_5_5_5;;
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
$__t3 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_5_4)(((($GLOBALS['Data_List_Types_bindList'])->{'bind'})((($v_0)->{'value1'})->{'value1'}))($__local_var_4_2)), ($go__go_5_5)(new \Data\List\Types\Data_List_Types_Nil()), ($__local_var_4_2)((($v_0)->{'value1'})->{'value0'}));
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $go__go_3_6 = null;
  $go__go_3_6 = (function() use (&$go__go_3_6) {
  $__fn = function($v_4, $v1_5 = null) use (&$go__go_3_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_3_6_6_v_4 = $v_4;
  $__tco_var_go__go_3_6_6_v1_5 = $v1_5;
  tco_loop_go__go_3_6_6:;
  $v_4 = $__tco_var_go__go_3_6_6_v_4;
  $v1_5 = $__tco_var_go__go_3_6_6_v1_5;
  $__t6 = null;;
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = $v_4;
goto end_branch_6;;
};
  if ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_7 = new \Data\List\Types\Data_List_Types_Cons(($v1_5)->{'value0'}, $v_4);
$__tco_8 = ($v1_5)->{'value1'};
$__tco_var_go__go_3_6_6_v_4 = $__tco_7;
$__tco_var_go__go_3_6_6_v1_5 = $__tco_8;
goto tco_loop_go__go_3_6_6;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v1_2_0)->{'value0'}, \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_3_1)($__t3), ($go__go_3_6)(new \Data\List\Types\Data_List_Types_Nil()), ($v1_2_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applyNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_applicativeList
$GLOBALS['Data_List_Types_applicativeList'] = (object)["pure" => function($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_0, new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_monadList
$GLOBALS['Data_List_Types_monadList'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_bindList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_altNonEmptyList
$GLOBALS['Data_List_Types_altNonEmptyList'] = (object)["alt" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($as_prime__1) use ($v_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $go__go_2_1 = null;
  $go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_v_3 = $v_3;
  $__tco_var_go__go_2_1_1_v1_4 = $v1_4;
  tco_loop_go__go_2_1_1:;
  $v_3 = $__tco_var_go__go_2_1_1_v_3;
  $v1_4 = $__tco_var_go__go_2_1_1_v1_4;
  $__t1 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_3;
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_3 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_1_1_v_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v1_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_0)->{'value0'}, \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_0)(new \Data\List\Types\Data_List_Types_Cons(($as_prime__1)->{'value0'}, ($as_prime__1)->{'value1'})), ($go__go_2_1)(new \Data\List\Types\Data_List_Types_Nil()), ($v_0)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_altList
$GLOBALS['Data_List_Types_altList'] = (object)["alt" => function($xs_0) {
  $__num = \func_num_args();
  $__res = function($ys_1) use ($xs_0) {
  $__num = \func_num_args();
  $go__go_2_0 = null;
  $go__go_2_0 = (function() use (&$go__go_2_0) {
  $__fn = function($b_3, $v_4 = null) use (&$go__go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_0_0_b_3 = $b_3;
  $__tco_var_go__go_2_0_0_v_4 = $v_4;
  tco_loop_go__go_2_0_0:;
  $b_3 = $__tco_var_go__go_2_0_0_b_3;
  $v_4 = $__tco_var_go__go_2_0_0_v_4;
  $__t0 = null;;
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $b_3;
goto end_branch_0;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v_4)->{'value0'}, $b_3);
$__tco_2 = ($v_4)->{'value1'};
$__tco_var_go__go_2_0_0_b_3 = $__tco_1;
$__tco_var_go__go_2_0_0_v_4 = $__tco_2;
goto tco_loop_go__go_2_0_0;;
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
  $go__go_2_1 = null;
  $go__go_2_1 = (function() use (&$go__go_2_1) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__go_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_2_1_1_v_3 = $v_3;
  $__tco_var_go__go_2_1_1_v1_4 = $v1_4;
  tco_loop_go__go_2_1_1:;
  $v_3 = $__tco_var_go__go_2_1_1_v_3;
  $v1_4 = $__tco_var_go__go_2_1_1_v1_4;
  $__t1 = null;;
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $v_3;
goto end_branch_1;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_3 = ($v1_4)->{'value1'};
$__tco_var_go__go_2_1_1_v_3 = $__tco_2;
$__tco_var_go__go_2_1_1_v1_4 = $__tco_3;
goto tco_loop_go__go_2_1_1;;
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
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_2_0)($ys_1), ($go__go_2_1)(new \Data\List\Types\Data_List_Types_Nil()), $xs_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_plusList
$GLOBALS['Data_List_Types_plusList'] = (object)["empty" => new \Data\List\Types\Data_List_Types_Nil(), "Alt0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_altList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_alternativeList
$GLOBALS['Data_List_Types_alternativeList'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_plusList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_monadPlusList
$GLOBALS['Data_List_Types_monadPlusList'] = (object)["Monad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_monadList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_alternativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_applicativeNonEmptyList
$GLOBALS['Data_List_Types_applicativeNonEmptyList'] = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_0) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_0, new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applyNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_monadNonEmptyList
$GLOBALS['Data_List_Types_monadNonEmptyList'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applicativeNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_bindNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_traversable1NonEmptyList
$GLOBALS['Data_List_Types_traversable1NonEmptyList'] = (object)["traverse1" => function($dictApply_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (($dictApply_0)->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApply_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApply_0, $f_2) {
  $__num = \func_num_args();
  $go__go_4_2 = null;
  $go__go_4_2 = (function() use ($dictApply_0, $f_2, &$go__go_4_2) {
  $__fn = function($b_5, $v_6 = null) use ($dictApply_0, $f_2, &$go__go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_2_2_b_5 = $b_5;
  $__tco_var_go__go_4_2_2_v_6 = $v_6;
  tco_loop_go__go_4_2_2:;
  $b_5 = $__tco_var_go__go_4_2_2_b_5;
  $v_6 = $__tco_var_go__go_4_2_2_v_6;
  $__t2 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t2 = $b_5;
goto end_branch_2;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_7_3 = (($dictApply_0)->{'Functor0'})(null);
$__tco_4 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_8) use ($Functor0_7_3, $b_5, $dictApply_0) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((($Functor0_7_3)->{'map'})(function($b_9) {
  $__num = \func_num_args();
  $__res = function($a_10) use ($b_9) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_10, new \Data\List\Types\Data_List_Types_Cons(($b_9)->{'value0'}, ($b_9)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_5)))($b_8);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $f_2, ($v_6)->{'value0'});
$__tco_5 = ($v_6)->{'value1'};
$__tco_var_go__go_4_2_2_b_5 = $__tco_4;
$__tco_var_go__go_4_2_2_v_6 = $__tco_5;
goto tco_loop_go__go_4_2_2;;
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
  $__res = ((($Functor0_1_0)->{'map'})(function($v1_4) {
  $__num = \func_num_args();
  $go__go_5_1 = null;
  $go__go_5_1 = (function() use (&$go__go_5_1) {
  $__fn = function($b_6, $v_7 = null) use (&$go__go_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_1_1_b_6 = $b_6;
  $__tco_var_go__go_5_1_1_v_7 = $v_7;
  tco_loop_go__go_5_1_1:;
  $b_6 = $__tco_var_go__go_5_1_1_b_6;
  $v_7 = $__tco_var_go__go_5_1_1_v_7;
  $__t1 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_6;
goto end_branch_1;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_7)->{'value0'}, new \Data\List\Types\Data_List_Types_Cons(($b_6)->{'value0'}, ($b_6)->{'value1'}));
$__tco_3 = ($v_7)->{'value1'};
$__tco_var_go__go_5_1_1_b_6 = $__tco_2;
$__tco_var_go__go_5_1_1_v_7 = $__tco_3;
goto tco_loop_go__go_5_1_1;;
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
  $__res = (($go__go_5_1)(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_6) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_6, new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v1_4)->{'value0'})))(($v1_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($go__go_4_2)(((($Functor0_1_0)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($a_5) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_5, new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($f_2)(($v_3)->{'value0'}))))(($v_3)->{'value1'}));
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
}, "sequence1" => function($dictApply_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_traversable1NonEmptyList'])->{'traverse1'})($dictApply_0))(function($x_1) {
  $__num = \func_num_args();
  $__res = $x_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable10" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $foldableNonEmpty1_1_3 = (object)["foldMap" => function($dictMonoid_1) {
  $__num = \func_num_args();
  $Semigroup0_2_3 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $__res = function($f_3) use ($Semigroup0_2_3, $dictMonoid_1) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($Semigroup0_2_3, $dictMonoid_1, $f_3) {
  $__num = \func_num_args();
  $Semigroup0_5_4 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $go__go_6_5 = null;
  $go__go_6_5 = (function() use ($Semigroup0_5_4, $f_3, &$go__go_6_5) {
  $__fn = function($b_7, $v_8 = null) use ($Semigroup0_5_4, $f_3, &$go__go_6_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_6_5_5_b_7 = $b_7;
  $__tco_var_go__go_6_5_5_v_8 = $v_8;
  tco_loop_go__go_6_5_5:;
  $b_7 = $__tco_var_go__go_6_5_5_b_7;
  $v_8 = $__tco_var_go__go_6_5_5_v_8;
  $__t5 = null;;
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t5 = $b_7;
goto end_branch_5;;
};
  if ($v_8 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_6 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_5_4)->{'append'})($b_7), $f_3, ($v_8)->{'value0'});
$__tco_7 = ($v_8)->{'value1'};
$__tco_var_go__go_6_5_5_b_7 = $__tco_6;
$__tco_var_go__go_6_5_5_v_8 = $__tco_7;
goto tco_loop_go__go_6_5_5;;
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
  $__res = ((($Semigroup0_2_3)->{'append'})(($f_3)(($v_4)->{'value0'})))((($go__go_6_5)(($dictMonoid_1)->{'mempty'}))(($v_4)->{'value1'}));
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
}, "foldl" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $f_1) {
  $__num = \func_num_args();
  $go__go_4_6 = null;
  $go__go_4_6 = (function() use ($f_1, &$go__go_4_6) {
  $__fn = function($b_5, $v_6 = null) use ($f_1, &$go__go_4_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_6_6_b_5 = $b_5;
  $__tco_var_go__go_4_6_6_v_6 = $v_6;
  tco_loop_go__go_4_6_6:;
  $b_5 = $__tco_var_go__go_4_6_6_b_5;
  $v_6 = $__tco_var_go__go_4_6_6_v_6;
  $__t6 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t6 = $b_5;
goto end_branch_6;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_7 = (($f_1)($b_5))(($v_6)->{'value0'});
$__tco_8 = ($v_6)->{'value1'};
$__tco_var_go__go_4_6_6_b_5 = $__tco_7;
$__tco_var_go__go_4_6_6_v_6 = $__tco_8;
goto tco_loop_go__go_4_6_6;;
$__t6 = null;
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_4_6)((($f_1)($b_2))(($v_3)->{'value0'})))(($v_3)->{'value1'});
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
}, "foldr" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($b_2, $f_1) {
  $__num = \func_num_args();
  $go__go_4_7 = null;
  $go__go_4_7 = (function() use ($f_1, &$go__go_4_7) {
  $__fn = function($b_5, $v_6 = null) use ($f_1, &$go__go_4_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_7_7_b_5 = $b_5;
  $__tco_var_go__go_4_7_7_v_6 = $v_6;
  tco_loop_go__go_4_7_7:;
  $b_5 = $__tco_var_go__go_4_7_7_b_5;
  $v_6 = $__tco_var_go__go_4_7_7_v_6;
  $__t7 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t7 = $b_5;
goto end_branch_7;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_8 = (($f_1)(($v_6)->{'value0'}))($b_5);
$__tco_9 = ($v_6)->{'value1'};
$__tco_var_go__go_4_7_7_b_5 = $__tco_8;
$__tco_var_go__go_4_7_7_v_6 = $__tco_9;
goto tco_loop_go__go_4_7_7;;
$__t7 = null;
goto end_branch_7;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t7 = null;
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_4_8 = null;
  $go__go_4_8 = (function() use (&$go__go_4_8) {
  $__fn = function($v_5, $v1_6 = null) use (&$go__go_4_8, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_8_8_v_5 = $v_5;
  $__tco_var_go__go_4_8_8_v1_6 = $v1_6;
  tco_loop_go__go_4_8_8:;
  $v_5 = $__tco_var_go__go_4_8_8_v_5;
  $v1_6 = $__tco_var_go__go_4_8_8_v1_6;
  $__t8 = null;;
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t8 = $v_5;
goto end_branch_8;;
};
  if ($v1_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_9 = new \Data\List\Types\Data_List_Types_Cons(($v1_6)->{'value0'}, $v_5);
$__tco_10 = ($v1_6)->{'value1'};
$__tco_var_go__go_4_8_8_v_5 = $__tco_9;
$__tco_var_go__go_4_8_8_v1_6 = $__tco_10;
goto tco_loop_go__go_4_8_8;;
$__t8 = null;
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($f_1)(($v_3)->{'value0'}))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_4_7)($b_2), ($go__go_4_8)(new \Data\List\Types\Data_List_Types_Nil()), ($v_3)->{'value1'}));
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
  $__res = (object)["foldMap1" => function($dictSemigroup_2) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictSemigroup_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictSemigroup_2, $f_3) {
  $__num = \func_num_args();
  $go__go_5_10 = null;
  $go__go_5_10 = (function() use ($dictSemigroup_2, $f_3, &$go__go_5_10) {
  $__fn = function($b_6, $v_7 = null) use ($dictSemigroup_2, $f_3, &$go__go_5_10, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_10_10_b_6 = $b_6;
  $__tco_var_go__go_5_10_10_v_7 = $v_7;
  tco_loop_go__go_5_10_10:;
  $b_6 = $__tco_var_go__go_5_10_10_b_6;
  $v_7 = $__tco_var_go__go_5_10_10_v_7;
  $__t10 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t10 = $b_6;
goto end_branch_10;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_11 = ((($dictSemigroup_2)->{'append'})($b_6))(($f_3)(($v_7)->{'value0'}));
$__tco_12 = ($v_7)->{'value1'};
$__tco_var_go__go_5_10_10_b_6 = $__tco_11;
$__tco_var_go__go_5_10_10_v_7 = $__tco_12;
goto tco_loop_go__go_5_10_10;;
$__t10 = null;
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_5_10)(($f_3)(($v_4)->{'value0'})))(($v_4)->{'value1'});
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
}, "foldr1" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__local_var_4_11 = ($f_2)(($v_3)->{'value0'});
  $go__go_5_12 = null;
  $go__go_5_12 = (function() use ($f_2, &$go__go_5_12) {
  $__fn = function($b_6, $v_7 = null) use ($f_2, &$go__go_5_12, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_12_12_b_6 = $b_6;
  $__tco_var_go__go_5_12_12_v_7 = $v_7;
  tco_loop_go__go_5_12_12:;
  $b_6 = $__tco_var_go__go_5_12_12_b_6;
  $v_7 = $__tco_var_go__go_5_12_12_v_7;
  $__t12 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t12 = $b_6;
goto end_branch_12;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__local_var_8_13 = ($f_2)(($v_7)->{'value0'});
$__tco_15 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Data_Maybe_Just'], function($v2_9) use ($__local_var_8_13, $v_7) {
  $__num = \func_num_args();
  $__t14 = null;;
  if ($v2_9 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t14 = ($v_7)->{'value0'};
goto end_branch_14;;
};
  if ($v2_9 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t14 = ($__local_var_8_13)(($v2_9)->{'value0'});
goto end_branch_14;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t14 = null;
  end_branch_14:;
  $__res = $__t14;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $b_6);
$__tco_16 = ($v_7)->{'value1'};
$__tco_var_go__go_5_12_12_b_6 = $__tco_15;
$__tco_var_go__go_5_12_12_v_7 = $__tco_16;
goto tco_loop_go__go_5_12_12;;
$__t12 = null;
goto end_branch_12;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t12 = null;
  end_branch_12:;
  $__res = $__t12;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_5_13 = null;
  $go__go_5_13 = (function() use (&$go__go_5_13) {
  $__fn = function($v_6, $v1_7 = null) use (&$go__go_5_13, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_13_13_v_6 = $v_6;
  $__tco_var_go__go_5_13_13_v1_7 = $v1_7;
  tco_loop_go__go_5_13_13:;
  $v_6 = $__tco_var_go__go_5_13_13_v_6;
  $v1_7 = $__tco_var_go__go_5_13_13_v1_7;
  $__t13 = null;;
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t13 = $v_6;
goto end_branch_13;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_14 = new \Data\List\Types\Data_List_Types_Cons(($v1_7)->{'value0'}, $v_6);
$__tco_15 = ($v1_7)->{'value1'};
$__tco_var_go__go_5_13_13_v_6 = $__tco_14;
$__tco_var_go__go_5_13_13_v1_7 = $__tco_15;
goto tco_loop_go__go_5_13_13;;
$__t13 = null;
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__local_var_5_12 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_5_12)(new \Data\Maybe\Data_Maybe_Nothing()), ($go__go_5_13)(new \Data\List\Types\Data_List_Types_Nil()), ($v_3)->{'value1'});
  $__t15 = null;;
  if ($__local_var_5_12 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t15 = ($v_3)->{'value0'};
goto end_branch_15;;
};
  if ($__local_var_5_12 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t15 = ($__local_var_4_11)(($__local_var_5_12)->{'value0'});
goto end_branch_15;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t15 = null;
  end_branch_15:;
  $__res = $__t15;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl1" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $go__go_4_16 = null;
  $go__go_4_16 = (function() use ($f_2, &$go__go_4_16) {
  $__fn = function($b_5, $v_6 = null) use ($f_2, &$go__go_4_16, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_4_16_16_b_5 = $b_5;
  $__tco_var_go__go_4_16_16_v_6 = $v_6;
  tco_loop_go__go_4_16_16:;
  $b_5 = $__tco_var_go__go_4_16_16_b_5;
  $v_6 = $__tco_var_go__go_4_16_16_v_6;
  $__t16 = null;;
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t16 = $b_5;
goto end_branch_16;;
};
  if ($v_6 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_17 = (($f_2)($b_5))(($v_6)->{'value0'});
$__tco_18 = ($v_6)->{'value1'};
$__tco_var_go__go_4_16_16_b_5 = $__tco_17;
$__tco_var_go__go_4_16_16_v_6 = $__tco_18;
goto tco_loop_go__go_4_16_16;;
$__t16 = null;
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_4_16)(($v_3)->{'value0'}))(($v_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar___unused_2) use ($foldableNonEmpty1_1_3) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_1_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $functorNonEmpty1_1_17 = (object)["map" => function($f_1) {
  $__num = \func_num_args();
  $__res = function($m_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_1)(($m_2)->{'value0'}), \Data\List\Types\majData_majList_majTypes_listmajMap($f_1, ($m_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $foldableNonEmpty1_2_18 = (object)["foldMap" => function($dictMonoid_2) {
  $__num = \func_num_args();
  $Semigroup0_3_18 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $__res = function($f_4) use ($Semigroup0_3_18, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Semigroup0_3_18, $dictMonoid_2, $f_4) {
  $__num = \func_num_args();
  $Semigroup0_6_19 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $go__go_7_20 = null;
  $go__go_7_20 = (function() use ($Semigroup0_6_19, $f_4, &$go__go_7_20) {
  $__fn = function($b_8, $v_9 = null) use ($Semigroup0_6_19, $f_4, &$go__go_7_20, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_7_20_20_b_8 = $b_8;
  $__tco_var_go__go_7_20_20_v_9 = $v_9;
  tco_loop_go__go_7_20_20:;
  $b_8 = $__tco_var_go__go_7_20_20_b_8;
  $v_9 = $__tco_var_go__go_7_20_20_v_9;
  $__t20 = null;;
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t20 = $b_8;
goto end_branch_20;;
};
  if ($v_9 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_21 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($Semigroup0_6_19)->{'append'})($b_8), $f_4, ($v_9)->{'value0'});
$__tco_22 = ($v_9)->{'value1'};
$__tco_var_go__go_7_20_20_b_8 = $__tco_21;
$__tco_var_go__go_7_20_20_v_9 = $__tco_22;
goto tco_loop_go__go_7_20_20;;
$__t20 = null;
goto end_branch_20;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t20 = null;
  end_branch_20:;
  $__res = $__t20;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Semigroup0_3_18)->{'append'})(($f_4)(($v_5)->{'value0'})))((($go__go_7_20)(($dictMonoid_2)->{'mempty'}))(($v_5)->{'value1'}));
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
}, "foldl" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($b_3, $f_2) {
  $__num = \func_num_args();
  $go__go_5_21 = null;
  $go__go_5_21 = (function() use ($f_2, &$go__go_5_21) {
  $__fn = function($b_6, $v_7 = null) use ($f_2, &$go__go_5_21, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_21_21_b_6 = $b_6;
  $__tco_var_go__go_5_21_21_v_7 = $v_7;
  tco_loop_go__go_5_21_21:;
  $b_6 = $__tco_var_go__go_5_21_21_b_6;
  $v_7 = $__tco_var_go__go_5_21_21_v_7;
  $__t21 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t21 = $b_6;
goto end_branch_21;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_22 = (($f_2)($b_6))(($v_7)->{'value0'});
$__tco_23 = ($v_7)->{'value1'};
$__tco_var_go__go_5_21_21_b_6 = $__tco_22;
$__tco_var_go__go_5_21_21_v_7 = $__tco_23;
goto tco_loop_go__go_5_21_21;;
$__t21 = null;
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_5_21)((($f_2)($b_3))(($v_4)->{'value0'})))(($v_4)->{'value1'});
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
}, "foldr" => function($f_2) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($b_3, $f_2) {
  $__num = \func_num_args();
  $go__go_5_22 = null;
  $go__go_5_22 = (function() use ($f_2, &$go__go_5_22) {
  $__fn = function($b_6, $v_7 = null) use ($f_2, &$go__go_5_22, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_22_22_b_6 = $b_6;
  $__tco_var_go__go_5_22_22_v_7 = $v_7;
  tco_loop_go__go_5_22_22:;
  $b_6 = $__tco_var_go__go_5_22_22_b_6;
  $v_7 = $__tco_var_go__go_5_22_22_v_7;
  $__t22 = null;;
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t22 = $b_6;
goto end_branch_22;;
};
  if ($v_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_23 = (($f_2)(($v_7)->{'value0'}))($b_6);
$__tco_24 = ($v_7)->{'value1'};
$__tco_var_go__go_5_22_22_b_6 = $__tco_23;
$__tco_var_go__go_5_22_22_v_7 = $__tco_24;
goto tco_loop_go__go_5_22_22;;
$__t22 = null;
goto end_branch_22;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t22 = null;
  end_branch_22:;
  $__res = $__t22;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_5_23 = null;
  $go__go_5_23 = (function() use (&$go__go_5_23) {
  $__fn = function($v_6, $v1_7 = null) use (&$go__go_5_23, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_5_23_23_v_6 = $v_6;
  $__tco_var_go__go_5_23_23_v1_7 = $v1_7;
  tco_loop_go__go_5_23_23:;
  $v_6 = $__tco_var_go__go_5_23_23_v_6;
  $v1_7 = $__tco_var_go__go_5_23_23_v1_7;
  $__t23 = null;;
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t23 = $v_6;
goto end_branch_23;;
};
  if ($v1_7 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_24 = new \Data\List\Types\Data_List_Types_Cons(($v1_7)->{'value0'}, $v_6);
$__tco_25 = ($v1_7)->{'value1'};
$__tco_var_go__go_5_23_23_v_6 = $__tco_24;
$__tco_var_go__go_5_23_23_v1_7 = $__tco_25;
goto tco_loop_go__go_5_23_23;;
$__t23 = null;
goto end_branch_23;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t23 = null;
  end_branch_23:;
  $__res = $__t23;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($f_2)(($v_4)->{'value0'}))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(($go__go_5_22)($b_3), ($go__go_5_23)(new \Data\List\Types\Data_List_Types_Nil()), ($v_4)->{'value1'}));
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
  $__res = (object)["sequence" => function($dictApplicative_3) {
  $__num = \func_num_args();
  $Apply0_4_25 = (($dictApplicative_3)->{'Apply0'})(null);
  $Functor0_5_26 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_6) use ($Apply0_4_25, $Functor0_5_26, $dictApplicative_3) {
  $__num = \func_num_args();
  $Apply0_7_27 = (($dictApplicative_3)->{'Apply0'})(null);
  $go__go_8_28 = null;
  $go__go_8_28 = (function() use (&$go__go_8_28) {
  $__fn = function($b_9, $v_10 = null) use (&$go__go_8_28, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_28_28_b_9 = $b_9;
  $__tco_var_go__go_8_28_28_v_10 = $v_10;
  tco_loop_go__go_8_28_28:;
  $b_9 = $__tco_var_go__go_8_28_28_b_9;
  $v_10 = $__tco_var_go__go_8_28_28_v_10;
  $__t28 = null;;
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t28 = $b_9;
goto end_branch_28;;
};
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_29 = new \Data\List\Types\Data_List_Types_Cons(($v_10)->{'value0'}, $b_9);
$__tco_30 = ($v_10)->{'value1'};
$__tco_var_go__go_8_28_28_b_9 = $__tco_29;
$__tco_var_go__go_8_28_28_v_10 = $__tco_30;
goto tco_loop_go__go_8_28_28;;
$__t28 = null;
goto end_branch_28;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t28 = null;
  end_branch_28:;
  $__res = $__t28;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_8_29 = null;
  $go__go_8_29 = (function() use ($Apply0_7_27, &$go__go_8_29) {
  $__fn = function($b_9, $v_10 = null) use ($Apply0_7_27, &$go__go_8_29, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_8_29_29_b_9 = $b_9;
  $__tco_var_go__go_8_29_29_v_10 = $v_10;
  tco_loop_go__go_8_29_29:;
  $b_9 = $__tco_var_go__go_8_29_29_b_9;
  $v_10 = $__tco_var_go__go_8_29_29_v_10;
  $__t29 = null;;
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t29 = $b_9;
goto end_branch_29;;
};
  if ($v_10 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_11_30 = (($Apply0_7_27)->{'Functor0'})(null);
$__tco_31 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_12) use ($Apply0_7_27, $Functor0_11_30, $b_9) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_27)->{'apply'})(((($Functor0_11_30)->{'map'})(function($b_13) {
  $__num = \func_num_args();
  $__res = function($a_14) use ($b_13) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_14, $b_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_9)))($b_12);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($x_11) {
  $__num = \func_num_args();
  $__res = $x_11;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($v_10)->{'value0'});
$__tco_32 = ($v_10)->{'value1'};
$__tco_var_go__go_8_29_29_b_9 = $__tco_31;
$__tco_var_go__go_8_29_29_v_10 = $__tco_32;
goto tco_loop_go__go_8_29_29;;
$__t29 = null;
goto end_branch_29;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t29 = null;
  end_branch_29:;
  $__res = $__t29;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Apply0_4_25)->{'apply'})(((($Functor0_5_26)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($v_6)->{'value0'})))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((((((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($go__go_8_28)(new \Data\List\Types\Data_List_Types_Nil())), ($go__go_8_29)((($dictApplicative_3)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())), ($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "traverse" => function($dictApplicative_3) {
  $__num = \func_num_args();
  $Apply0_4_30 = (($dictApplicative_3)->{'Apply0'})(null);
  $Functor0_5_31 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_6) use ($Apply0_4_30, $Functor0_5_31, $dictApplicative_3) {
  $__num = \func_num_args();
  $__res = function($v_7) use ($Apply0_4_30, $Functor0_5_31, $dictApplicative_3, $f_6) {
  $__num = \func_num_args();
  $Apply0_8_32 = (($dictApplicative_3)->{'Apply0'})(null);
  $go__go_9_33 = null;
  $go__go_9_33 = (function() use (&$go__go_9_33) {
  $__fn = function($b_10, $v_11 = null) use (&$go__go_9_33, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_9_33_33_b_10 = $b_10;
  $__tco_var_go__go_9_33_33_v_11 = $v_11;
  tco_loop_go__go_9_33_33:;
  $b_10 = $__tco_var_go__go_9_33_33_b_10;
  $v_11 = $__tco_var_go__go_9_33_33_v_11;
  $__t33 = null;;
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t33 = $b_10;
goto end_branch_33;;
};
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_34 = new \Data\List\Types\Data_List_Types_Cons(($v_11)->{'value0'}, $b_10);
$__tco_35 = ($v_11)->{'value1'};
$__tco_var_go__go_9_33_33_b_10 = $__tco_34;
$__tco_var_go__go_9_33_33_v_11 = $__tco_35;
goto tco_loop_go__go_9_33_33;;
$__t33 = null;
goto end_branch_33;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t33 = null;
  end_branch_33:;
  $__res = $__t33;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go__go_9_34 = null;
  $go__go_9_34 = (function() use ($Apply0_8_32, $f_6, &$go__go_9_34) {
  $__fn = function($b_10, $v_11 = null) use ($Apply0_8_32, $f_6, &$go__go_9_34, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_9_34_34_b_10 = $b_10;
  $__tco_var_go__go_9_34_34_v_11 = $v_11;
  tco_loop_go__go_9_34_34:;
  $b_10 = $__tco_var_go__go_9_34_34_b_10;
  $v_11 = $__tco_var_go__go_9_34_34_v_11;
  $__t34 = null;;
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t34 = $b_10;
goto end_branch_34;;
};
  if ($v_11 instanceof \Data\List\Types\Data_List_Types_Cons) {
$Functor0_12_35 = (($Apply0_8_32)->{'Functor0'})(null);
$__tco_36 = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($b_13) use ($Apply0_8_32, $Functor0_12_35, $b_10) {
  $__num = \func_num_args();
  $__res = ((($Apply0_8_32)->{'apply'})(((($Functor0_12_35)->{'map'})(function($b_14) {
  $__num = \func_num_args();
  $__res = function($a_15) use ($b_14) {
  $__num = \func_num_args();
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_15, $b_14);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($b_10)))($b_13);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $f_6, ($v_11)->{'value0'});
$__tco_37 = ($v_11)->{'value1'};
$__tco_var_go__go_9_34_34_b_10 = $__tco_36;
$__tco_var_go__go_9_34_34_v_11 = $__tco_37;
goto tco_loop_go__go_9_34_34;;
$__t34 = null;
goto end_branch_34;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t34 = null;
  end_branch_34:;
  $__res = $__t34;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($Apply0_4_30)->{'apply'})(((($Functor0_5_31)->{'map'})($GLOBALS['Data_NonEmpty_NonEmpty']))(($f_6)(($v_7)->{'value0'}))))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((((((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($go__go_9_33)(new \Data\List\Types\Data_List_Types_Nil())), ($go__go_9_34)((($dictApplicative_3)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())), ($v_7)->{'value1'}));
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
}, "Functor0" => function($_dollar___unused_3) use ($functorNonEmpty1_1_17) {
  $__num = \func_num_args();
  $__res = $functorNonEmpty1_1_17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_3) use ($foldableNonEmpty1_2_18) {
  $__num = \func_num_args();
  $__res = $foldableNonEmpty1_2_18;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

