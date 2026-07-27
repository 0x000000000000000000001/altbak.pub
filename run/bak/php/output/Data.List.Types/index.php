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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


final class Data_List_Types_Nil { public $tag = 'Nil'; public function __construct() {} }
final class Data_List_Types_Cons { public $tag = 'Cons'; public function __construct(public  $value0, public  $value1) {} }

// Data_List_Types_identity
function majData_majList_majTypes_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_identity'] = __NAMESPACE__ . '\\majData_majList_majTypes_identity';

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
$GLOBALS['Data_List_Types_newtypeNonEmptyList'] = (object)["Coercible0" => function($_dollar__unused_0) {
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
$GLOBALS['Data_List_Types_functorNonEmptyList'] = (object)["map" => (function() {
  $__fn = function($f_0, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)(($m_1)->{'value0'}), ((($GLOBALS['Data_List_Types_functorList'])->{'map'})($f_0))(($m_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Types_foldableList
$GLOBALS['Data_List_Types_foldableList'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use (&$go__2_0) {
  $__fn = function($v_3, $v1_4 = null) use (&$go__2_0, &$__fn) {
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
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = $v_3;
goto end_branch_0;;
};
  if ($v1_4 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_1 = new \Data\List\Types\Data_List_Types_Cons(($v1_4)->{'value0'}, $v_3);
$__tco_2 = ($v1_4)->{'value1'};
$__tco_var_go__2_0_0_v_3 = $__tco_1;
$__tco_var_go__2_0_0_v1_4 = $__tco_2;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() use ($f_0) {
  $__fn = function($b_2, $a_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)($a_3))($b_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b_1)))(($go__2_0)(new \Data\List\Types\Data_List_Types_Nil()));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl" => function($f_0) {
  $__num = \func_num_args();
  $go__1_1 = null;
  $go__1_1 = (function() use ($f_0, &$go__1_1) {
  $__fn = function($b_2, $v_3 = null) use ($f_0, &$go__1_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_1_1_b_2 = $b_2;
  $__tco_var_go__1_1_1_v_3 = $v_3;
  tco_loop_go__1_1_1:;
  $b_2 = $__tco_var_go__1_1_1_b_2;
  $v_3 = $__tco_var_go__1_1_1_v_3;
  $__t1 = null;;
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = $b_2;
goto end_branch_1;;
};
  if ($v_3 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__tco_2 = (($f_0)($b_2))(($v_3)->{'value0'});
$__tco_3 = ($v_3)->{'value1'};
$__tco_var_go__1_1_1_b_2 = $__tco_2;
$__tco_var_go__1_1_1_v_3 = $__tco_3;
goto tco_loop_go__1_1_1;;
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
  $__res = $go__1_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($acc_3) use ($dictMonoid_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})($acc_3)))($f_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_1_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_intercalate
$GLOBALS['Data_List_Types_intercalate'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Data_Monoid_monoidString'])->{'Semigroup0'})(null);
return (function() use ($__local_var_0_0) {
  $__fn = function($sep_1, $xs_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() use ($__local_var_0_0, $sep_1) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_0_0, $sep_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if (($v_3)->{'init'}) {
$__t1 = (object)["init" => false, "acc" => $v1_4];
goto end_branch_1;;
};
  $__t1 = (object)["init" => false, "acc" => ((($__local_var_0_0)->{'append'})(($v_3)->{'acc'}))(((($__local_var_0_0)->{'append'})($sep_1))($v1_4))];
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["init" => true, "acc" => ($GLOBALS['Data_Monoid_monoidString'])->{'mempty'}]))($xs_2))->{'acc'};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_List_Types_foldableNonEmptyList
$GLOBALS['Data_List_Types_foldableNonEmptyList'] = (object)["foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMap1_1_0 = (($GLOBALS['Data_List_Types_foldableList'])->{'foldMap'})($dictMonoid_0);
  $__res = (function() use ($dictMonoid_0, $foldMap1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictMonoid_0, $foldMap1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})(($f_2)(($v_3)->{'value0'})))((($foldMap1_1_0)($f_2))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})($f_0))((($f_0)($b_1))(($v_2)->{'value0'})))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldr" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)(($v_2)->{'value0'}))((((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})($f_0))($b_1))(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Types_foldableWithIndexList
$GLOBALS['Data_List_Types_foldableWithIndexList'] = (object)["foldrWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_3_0 = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($v1_3) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($v1_3)->{'value1'};
  $__local_var_5_1 = ($v1_3)->{'value0'};
  $__res = function($a_6) use ($__local_var_4_0, $__local_var_5_1) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_5_1 + 1), new \Data\List\Types\Data_List_Types_Cons($a_6, $__local_var_4_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple(0, new \Data\List\Types\Data_List_Types_Nil())))($xs_2);
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($v1_4) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_5_3 = ($v1_4)->{'value1'};
  $__local_var_6_4 = ($v1_4)->{'value0'};
  $__res = function($a_7) use ($__local_var_5_3, $__local_var_6_4, $f_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_6_4 - 1), ((($f_0)(($__local_var_6_4 - 1)))($a_7))($__local_var_5_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple(($v_3_0)->{'value0'}, $b_1)))(($v_3_0)->{'value1'}))->{'value1'};
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0, $acc_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))(((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($v_2) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_3_5 = ($v_2)->{'value1'};
  $__local_var_4_6 = ($v_2)->{'value0'};
  $__res = function($a_5) use ($__local_var_3_5, $__local_var_4_6, $f_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($__local_var_4_6 + 1), ((($f_0)($__local_var_4_6))($__local_var_3_5))($a_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Tuple\Data_Tuple_Tuple(0, $acc_1)));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_7 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_7) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_foldableWithIndexList'])->{'foldlWithIndex'})((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($i_3, $acc_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})($acc_4)))(($f_2)($i_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_foldableWithIndexNonEmpty
$GLOBALS['Data_List_Types_foldableWithIndexNonEmpty'] = \Data\NonEmpty\majData_majNonmajEmpty_foldablemajWithmajIndexmajNonmajEmpty($GLOBALS['Data_List_Types_foldableWithIndexList']);

// Data_List_Types_foldableWithIndexNonEmptyList
$GLOBALS['Data_List_Types_foldableWithIndexNonEmptyList'] = (object)["foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $foldMapWithIndex1_1_0 = (($GLOBALS['Data_List_Types_foldableWithIndexNonEmpty'])->{'foldMapWithIndex'})($dictMonoid_0);
  $__res = (function() use ($foldMapWithIndex1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($foldMapWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMapWithIndex1_1_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))(function($v2_4) {
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
})))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldlWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_Types_foldableWithIndexNonEmpty'])->{'foldlWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = 0;
goto end_branch_2;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = (1 + ($v2_3)->{'value0'});
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($b_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldrWithIndex" => (function() {
  $__fn = function($f_0, $b_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_List_Types_foldableWithIndexNonEmpty'])->{'foldrWithIndex'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_0))(function($v2_3) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = 0;
goto end_branch_3;;
};
  if ($v2_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = (1 + ($v2_3)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($b_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Foldable0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_functorWithIndexList
$GLOBALS['Data_List_Types_functorWithIndexList'] = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_foldableWithIndexList'])->{'foldrWithIndex'})((function() use ($f_0) {
  $__fn = function($i_1, $x_2 = null, $acc_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons((($f_0)($i_1))($x_2), $acc_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_mapWithIndex
$GLOBALS['Data_List_Types_mapWithIndex'] = (\Data\NonEmpty\majData_majNonmajEmpty_functormajWithmajIndex($GLOBALS['Data_List_Types_functorWithIndexList']))->{'mapWithIndex'};

// Data_List_Types_functorWithIndexNonEmptyList
$GLOBALS['Data_List_Types_functorWithIndexNonEmptyList'] = (object)["mapWithIndex" => (function() {
  $__fn = function($fn_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_List_Types_mapWithIndex'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($fn_0))(function($v2_2) {
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
})))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_semigroupList
$GLOBALS['Data_List_Types_semigroupList'] = (object)["append" => (function() {
  $__fn = function($xs_0, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})($GLOBALS['Data_List_Types_Cons']))($ys_1))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Types_monoidList
$GLOBALS['Data_List_Types_monoidList'] = (object)["mempty" => new \Data\List\Types\Data_List_Types_Nil(), "Semigroup0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_semigroupList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_semigroupNonEmptyList
$GLOBALS['Data_List_Types_semigroupNonEmptyList'] = (object)["append" => (function() {
  $__fn = function($v_0, $as_prime_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_0)->{'value0'}, ((($GLOBALS['Data_List_Types_semigroupList'])->{'append'})(($v_0)->{'value1'}))(new \Data\List\Types\Data_List_Types_Cons(($as_prime_1)->{'value0'}, ($as_prime_1)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

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
  $__t1 = null;;
  if ($v_2 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = "Nil";
goto end_branch_1;;
};
  $__t1 = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("("))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($GLOBALS['Data_List_Types_intercalate'])(" : "))(((($GLOBALS['Data_List_Types_functorList'])->{'map'})($show_1_0))($v_2))))(" : Nil)"));
  end_branch_1:;
  $__res = $__t1;
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
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(NonEmptyList "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(((\Data\NonEmpty\majData_majNonmajEmpty_showmajNonmajEmpty($dictShow_0, \Data\List\Types\majData_majList_majTypes_showmajList($dictShow_0)))->{'show'})($v_1)))(")"));
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
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})(((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($b_3, $a_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\List\Types\Data_List_Types_Nil()))))(((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($acc_3) use ($Apply0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($b_4) use ($Apply0_1_0, $acc_3) {
  $__num = \func_num_args();
  $__res = ((($Apply0_1_0)->{'apply'})(((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})((function() {
  $__fn = function($b_5, $a_6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_6, $b_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($acc_3)))($b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($dictApplicative_0)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_traversableList'])->{'traverse'})($dictApplicative_0))($GLOBALS['Data_List_Types_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_traversableNonEmptyList
$GLOBALS['Data_List_Types_traversableNonEmptyList'] = \Data\NonEmpty\majData_majNonmajEmpty_traversablemajNonmajEmpty($GLOBALS['Data_List_Types_traversableList']);

// Data_List_Types_traversableWithIndexList
$GLOBALS['Data_List_Types_traversableWithIndexList'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})(((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($b_3, $a_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\List\Types\Data_List_Types_Nil()))))(((($GLOBALS['Data_List_Types_foldableWithIndexList'])->{'foldlWithIndex'})((function() use ($Apply0_1_0, $f_2) {
  $__fn = function($i_3, $acc_4 = null) use ($Apply0_1_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($b_5) use ($Apply0_1_0, $acc_4) {
  $__num = \func_num_args();
  $__res = ((($Apply0_1_0)->{'apply'})(((((($Apply0_1_0)->{'Functor0'})(null))->{'map'})((function() {
  $__fn = function($b_6, $a_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_7, $b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($acc_4)))($b_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($i_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($dictApplicative_0)->{'pure'})(new \Data\List\Types\Data_List_Types_Nil())));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableWithIndexList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_traversableList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_traverseWithIndex
$GLOBALS['Data_List_Types_traverseWithIndex'] = (\Data\NonEmpty\majData_majNonmajEmpty_traversablemajWithmajIndexmajNonmajEmpty($GLOBALS['Data_List_Types_traversableWithIndexList']))->{'traverseWithIndex'};

// Data_List_Types_traversableWithIndexNonEmptyList
$GLOBALS['Data_List_Types_traversableWithIndexNonEmptyList'] = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverseWithIndex1_1_0 = ($GLOBALS['Data_List_Types_traverseWithIndex'])($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverseWithIndex1_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictApplicative_0, $traverseWithIndex1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_List_Types_NonEmptyList']))((($traverseWithIndex1_1_0)((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_2))(function($v2_4) {
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
})))($v_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldableWithIndexNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_traversableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_unfoldable1List
$GLOBALS['Data_List_Types_unfoldable1List'] = (object)["unfoldr1" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($f_0, &$go__2_0) {
  $__fn = function($source_3, $memo_4 = null) use ($f_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__2_0_0_source_3 = $source_3;
  $__tco_var_go__2_0_0_memo_4 = $memo_4;
  tco_loop_go__2_0_0:;
  $source_3 = $__tco_var_go__2_0_0_source_3;
  $memo_4 = $__tco_var_go__2_0_0_memo_4;
  $v_5_0 = ($f_0)($source_3);
  $__t1 = null;;
  if (($v_5_0)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_2 = (($v_5_0)->{'value1'})->{'value0'};
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons(($v_5_0)->{'value0'}, $memo_4);
$__tco_var_go__2_0_0_source_3 = $__tco_2;
$__tco_var_go__2_0_0_memo_4 = $__tco_3;
goto tco_loop_go__2_0_0;;
$__t1 = null;
goto end_branch_1;;
};
  if (($v_5_0)->{'value1'} instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($b_6, $a_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_7, $b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\List\Types\Data_List_Types_Nil()))(new \Data\List\Types\Data_List_Types_Cons(($v_5_0)->{'value0'}, $memo_4));
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
  $__res = (($go__2_0)($b_1))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Types_unfoldableList
$GLOBALS['Data_List_Types_unfoldableList'] = (object)["unfoldr" => (function() {
  $__fn = function($f_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__2_0 = null;
  $go__2_0 = (function() use ($f_0, &$go__2_0) {
  $__fn = function($source_3, $memo_4 = null) use ($f_0, &$go__2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__2_0_0_source_3 = $source_3;
  $__tco_var_go__2_0_0_memo_4 = $memo_4;
  tco_loop_go__2_0_0:;
  $source_3 = $__tco_var_go__2_0_0_source_3;
  $memo_4 = $__tco_var_go__2_0_0_memo_4;
  $v_5_0 = ($f_0)($source_3);
  $__t1 = null;;
  if ($v_5_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($b_6, $a_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\List\Types\Data_List_Types_Cons($a_7, $b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\List\Types\Data_List_Types_Nil()))($memo_4);
goto end_branch_1;;
};
  if ($v_5_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_2 = (($v_5_0)->{'value0'})->{'value1'};
$__tco_3 = new \Data\List\Types\Data_List_Types_Cons((($v_5_0)->{'value0'})->{'value0'}, $memo_4);
$__tco_var_go__2_0_0_source_3 = $__tco_2;
$__tco_var_go__2_0_0_memo_4 = $__tco_3;
goto tco_loop_go__2_0_0;;
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
  $__res = (($go__2_0)($b_1))(new \Data\List\Types\Data_List_Types_Nil());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Unfoldable10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_unfoldable1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_unfoldable1NonEmptyList
$GLOBALS['Data_List_Types_unfoldable1NonEmptyList'] = \Data\NonEmpty\majData_majNonmajEmpty_unfoldable1majNonmajEmpty($GLOBALS['Data_List_Types_unfoldableList']);

// Data_List_Types_foldable1NonEmptyList
$GLOBALS['Data_List_Types_foldable1NonEmptyList'] = \Data\NonEmpty\majData_majNonmajEmpty_foldable1majNonmajEmpty($GLOBALS['Data_List_Types_foldableList']);

// Data_List_Types_extendNonEmptyList
$GLOBALS['Data_List_Types_extendNonEmptyList'] = (object)["extend" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($f_0)($v_1), ((((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})((function() use ($f_0) {
  $__fn = function($a_2, $v1_3 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["val" => new \Data\List\Types\Data_List_Types_Cons(($f_0)(new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_2, ($v1_3)->{'acc'})), ($v1_3)->{'val'}), "acc" => new \Data\List\Types\Data_List_Types_Cons($a_2, ($v1_3)->{'acc'})];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["val" => new \Data\List\Types\Data_List_Types_Nil(), "acc" => new \Data\List\Types\Data_List_Types_Nil()]))(($v_1)->{'value1'}))->{'val'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_extendList
$GLOBALS['Data_List_Types_extendList'] = (object)["extend" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v1_1 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = new \Data\List\Types\Data_List_Types_Cons(($v_0)($v1_1), ((((($GLOBALS['Data_List_Types_foldableList'])->{'foldr'})((function() use ($v_0) {
  $__fn = function($a_prime_2, $v2_3 = null) use ($v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["val" => new \Data\List\Types\Data_List_Types_Cons(($v_0)(new \Data\List\Types\Data_List_Types_Cons($a_prime_2, ($v2_3)->{'acc'})), ($v2_3)->{'val'}), "acc" => new \Data\List\Types\Data_List_Types_Cons($a_prime_2, ($v2_3)->{'acc'})];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["val" => new \Data\List\Types\Data_List_Types_Nil(), "acc" => new \Data\List\Types\Data_List_Types_Nil()]))(($v1_1)->{'value1'}))->{'val'});
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
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_eq1List
$GLOBALS['Data_List_Types_eq1List'] = (object)["eq1" => (function() {
  $__fn = function($dictEq_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = (function() use ($dictEq_0, &$go__3_0) {
  $__fn = function($v_4, $v1_5 = null, $v2_6 = null) use ($dictEq_0, &$go__3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t1 = null;;
  if (( ! $v2_6)) {
$__t1 = false;
goto end_branch_1;;
};
  if ($v_4 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t1 = ($v1_5 instanceof \Data\List\Types\Data_List_Types_Nil && $v2_6);
goto end_branch_1;;
};
  $__t1 = ($v_4 instanceof \Data\List\Types\Data_List_Types_Cons && ($v1_5 instanceof \Data\List\Types\Data_List_Types_Cons && ((($go__3_0)(($v_4)->{'value1'}))(($v1_5)->{'value1'}))(((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})($v2_6))(((($dictEq_0)->{'eq'})(($v1_5)->{'value0'}))(($v_4)->{'value0'})))));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($go__3_0)($xs_1))($ys_2))(true);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()];

// Data_List_Types_eq1NonEmptyList
$GLOBALS['Data_List_Types_eq1NonEmptyList'] = (object)["eq1" => function($dictEq_0) {
  $__num = \func_num_args();
  $__res = (\Data\NonEmpty\majData_majNonmajEmpty_eqmajNonmajEmpty($GLOBALS['Data_List_Types_eq1List'], $dictEq_0))->{'eq'};
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
  $__res = (object)["eq" => (($GLOBALS['Data_List_Types_eq1List'])->{'eq1'})($dictEq_0)];
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
  $__res = \Data\NonEmpty\majData_majNonmajEmpty_eqmajNonmajEmpty($GLOBALS['Data_List_Types_eq1List'], $dictEq_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_List_Types_eqNonEmptyList'] = __NAMESPACE__ . '\\majData_majList_majTypes_eqmajNonmajEmptymajList';

// Data_List_Types_ord1List
$GLOBALS['Data_List_Types_ord1List'] = (object)["compare1" => (function() {
  $__fn = function($dictOrd_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $go__3_0 = null;
  $go__3_0 = (function() use ($dictOrd_0, &$go__3_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictOrd_0, &$go__3_0, &$__fn) {
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
$__tco_var_go__3_0_0_v_4 = $__tco_4;
$__tco_var_go__3_0_0_v1_5 = $__tco_5;
goto tco_loop_go__3_0_0;;
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
  $__res = (($go__3_0)($xs_1))($ys_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Eq10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_eq1List'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_ordNonEmpty
$GLOBALS['Data_List_Types_ordNonEmpty'] = ($GLOBALS['Data_NonEmpty_ordNonEmpty'])($GLOBALS['Data_List_Types_ord1List']);

// Data_List_Types_ord1NonEmptyList
$GLOBALS['Data_List_Types_ord1NonEmptyList'] = \Data\NonEmpty\majData_majNonmajEmpty_ord1majNonmajEmpty($GLOBALS['Data_List_Types_ord1List']);

// Data_List_Types_ordList
function majData_majList_majTypes_ordmajList($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majList_majTypes_ordmajList';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eqList1_1_0 = (object)["eq" => (($GLOBALS['Data_List_Types_eq1List'])->{'eq1'})((($dictOrd_0)->{'Eq0'})(null))];
  $__res = (object)["compare" => (($GLOBALS['Data_List_Types_ord1List'])->{'compare1'})($dictOrd_0), "Eq0" => function($_dollar__unused_2) use ($eqList1_1_0) {
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
  $__res = ($GLOBALS['Data_List_Types_ordNonEmpty'])($dictOrd_0);
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
}, "Extend0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_extendNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_applyList
$GLOBALS['Data_List_Types_applyList'] = (object)["apply" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = ((($GLOBALS['Data_List_Types_semigroupList'])->{'append'})(((($GLOBALS['Data_List_Types_functorList'])->{'map'})(($v_0)->{'value0'}))($v1_1)))(((($GLOBALS['Data_List_Types_applyList'])->{'apply'})(($v_0)->{'value1'}))($v1_1));
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
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_applyNonEmptyList
$GLOBALS['Data_List_Types_applyNonEmptyList'] = (object)["apply" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty((($v_0)->{'value0'})(($v1_1)->{'value0'}), ((($GLOBALS['Data_List_Types_semigroupList'])->{'append'})(((($GLOBALS['Data_List_Types_applyList'])->{'apply'})(($v_0)->{'value1'}))(new \Data\List\Types\Data_List_Types_Cons(($v1_1)->{'value0'}, new \Data\List\Types\Data_List_Types_Nil()))))(((($GLOBALS['Data_List_Types_applyList'])->{'apply'})(new \Data\List\Types\Data_List_Types_Cons(($v_0)->{'value0'}, ($v_0)->{'value1'})))(($v1_1)->{'value1'})));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_bindList
$GLOBALS['Data_List_Types_bindList'] = (object)["bind" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Nil) {
$__t0 = new \Data\List\Types\Data_List_Types_Nil();
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\List\Types\Data_List_Types_Cons) {
$__t0 = ((($GLOBALS['Data_List_Types_semigroupList'])->{'append'})(($v1_1)(($v_0)->{'value0'})))(((($GLOBALS['Data_List_Types_bindList'])->{'bind'})(($v_0)->{'value1'}))($v1_1));
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
})(), "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_bindNonEmptyList
$GLOBALS['Data_List_Types_bindNonEmptyList'] = (object)["bind" => (function() {
  $__fn = function($v_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_2_0 = ($f_1)(($v_0)->{'value0'});
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v1_2_0)->{'value0'}, ((($GLOBALS['Data_List_Types_semigroupList'])->{'append'})(($v1_2_0)->{'value1'}))(((($GLOBALS['Data_List_Types_bindList'])->{'bind'})(($v_0)->{'value1'}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Types_toList']))($f_1))));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_0) {
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
}, "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_monadList
$GLOBALS['Data_List_Types_monadList'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_bindList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_altNonEmptyList
$GLOBALS['Data_List_Types_altNonEmptyList'] = (object)["alt" => ($GLOBALS['Data_List_Types_semigroupNonEmptyList'])->{'append'}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_altList
$GLOBALS['Data_List_Types_altList'] = (object)["alt" => ($GLOBALS['Data_List_Types_semigroupList'])->{'append'}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_functorList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_plusList
$GLOBALS['Data_List_Types_plusList'] = (object)["empty" => new \Data\List\Types\Data_List_Types_Nil(), "Alt0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_altList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_alternativeList
$GLOBALS['Data_List_Types_alternativeList'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applicativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_plusList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_monadPlusList
$GLOBALS['Data_List_Types_monadPlusList'] = (object)["Monad0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_monadList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Alternative1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_alternativeList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_applicativeNonEmptyList
$GLOBALS['Data_List_Types_applicativeNonEmptyList'] = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_List_Types_NonEmptyList']))(function($a_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_0, ($GLOBALS['Data_List_Types_plusList'])->{'empty'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applyNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_List_Types_monadNonEmptyList
$GLOBALS['Data_List_Types_monadNonEmptyList'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_applicativeNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
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
  $__res = (function() use ($Functor0_1_0, $dictApply_0) {
  $__fn = function($f_2, $v_3 = null) use ($Functor0_1_0, $dictApply_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($Functor0_1_0)->{'map'})(function($v1_4) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})((function() {
  $__fn = function($b_5, $a_6 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_6, new \Data\List\Types\Data_List_Types_Cons(($b_5)->{'value0'}, ($b_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($GLOBALS['Data_List_Types_applicativeNonEmptyList'])->{'pure'})(($v1_4)->{'value0'})))(($v1_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((((($GLOBALS['Data_List_Types_foldableList'])->{'foldl'})(function($acc_4) use ($dictApply_0, $f_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($b_5) use ($acc_4, $dictApply_0) {
  $__num = \func_num_args();
  $__res = ((($dictApply_0)->{'apply'})(((((($dictApply_0)->{'Functor0'})(null))->{'map'})((function() {
  $__fn = function($b_6, $a_7 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty($a_7, new \Data\List\Types\Data_List_Types_Cons(($b_6)->{'value0'}, ($b_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($acc_4)))($b_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($Functor0_1_0)->{'map'})(($GLOBALS['Data_List_Types_applicativeNonEmptyList'])->{'pure'}))(($f_2)(($v_3)->{'value0'}))))(($v_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence1" => function($dictApply_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_List_Types_traversable1NonEmptyList'])->{'traverse1'})($dictApply_0))($GLOBALS['Data_List_Types_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_foldable1NonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_List_Types_traversableNonEmptyList'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

