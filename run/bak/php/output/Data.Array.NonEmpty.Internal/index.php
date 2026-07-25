<?php

namespace Data\Array\NonEmpty\Internal;

// ALL IMPORTS: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Data.Array.NonEmpty.Internal, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Unfoldable1, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Data.Array.NonEmpty.Internal, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.Ord, Data.Semigroup, Data.Semigroup.Foldable, Data.Semigroup.Traversable, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Unfoldable1, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semigroup.Traversable/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
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
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_NonEmptyArray'] = function() { $v = function($x_0 = null) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray'] = function() { $v = ($GLOBALS['Data_Unfoldable1_unfoldable1Array'] ?? \PhpursThunks::eval('Data_Unfoldable1_unfoldable1Array')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray'] = function() { $v = ($GLOBALS['Data_TraversableWithIndex_traversableWithIndexArray'] ?? \PhpursThunks::eval('Data_TraversableWithIndex_traversableWithIndexArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_traversableNonEmptyArray'] = function() { $v = ($GLOBALS['Data_Traversable_traversableArray'] ?? \PhpursThunks::eval('Data_Traversable_traversableArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_showNonEmptyArray'] = function() { $v = function($dictShow_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["show" => function($v_1 = null) use ($dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(NonEmptyArray "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(((($GLOBALS['Data_Show_showArrayImpl'] ?? \PhpursThunks::eval('Data_Show_showArrayImpl')))(($dictShow_0)->{'show'}))($v_1)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_semigroupNonEmptyArray'] = function() { $v = ($GLOBALS['Data_Semigroup_semigroupArray'] ?? \PhpursThunks::eval('Data_Semigroup_semigroupArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_ordNonEmptyArray'] = function() { $v = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Ord_ordArray'] ?? \PhpursThunks::eval('Data_Ord_ordArray')))($dictOrd_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_ord1NonEmptyArray'] = function() { $v = ($GLOBALS['Data_Ord_ord1Array'] ?? \PhpursThunks::eval('Data_Ord_ord1Array')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_monadNonEmptyArray'] = function() { $v = ($GLOBALS['Control_Monad_monadArray'] ?? \PhpursThunks::eval('Control_Monad_monadArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray'] = function() { $v = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'] ?? \PhpursThunks::eval('Data_FunctorWithIndex_functorWithIndexArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_functorNonEmptyArray'] = function() { $v = ($GLOBALS['Data_Functor_functorArray'] ?? \PhpursThunks::eval('Data_Functor_functorArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray'] = function() { $v = ($GLOBALS['Data_FoldableWithIndex_foldableWithIndexArray'] ?? \PhpursThunks::eval('Data_FoldableWithIndex_foldableWithIndexArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_foldableNonEmptyArray'] = function() { $v = ($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] = function() { $v = (object)["foldMap1" => function($dictSemigroup_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Semigroup_Foldable_foldMap1DefaultL'] ?? \PhpursThunks::eval('Data_Semigroup_Foldable_foldMap1DefaultL')))(($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'))))(($GLOBALS['Data_Functor_functorArray'] ?? \PhpursThunks::eval('Data_Functor_functorArray'))))($dictSemigroup_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr1" => (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_NonEmpty_Internal_foldr1Impl'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldr1Impl'))), "foldl1" => (($GLOBALS['Data_Function_Uncurried_runFn2'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn2')))(($GLOBALS['Data_Array_NonEmpty_Internal_foldl1Impl'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldl1Impl'))), "Foldable0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_traversable1NonEmptyArray'] = function() { $v = (object)["traverse1" => function($dictApply_0 = null) {
  $__num = \func_num_args();
  $apply_1_0 = ($dictApply_0)->{'apply'};
  $map_2_1 = ((($dictApply_0)->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'map'};
  $__res = function($f_3 = null) use ($apply_1_0, $map_2_1) {
  $__num = \func_num_args();
  $__res = ((((($GLOBALS['Data_Function_Uncurried_runFn3'] ?? \PhpursThunks::eval('Data_Function_Uncurried_runFn3')))(($GLOBALS['Data_Array_NonEmpty_Internal_traverse1Impl'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_traverse1Impl'))))($apply_1_0))($map_2_1))($f_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence1" => function($dictApply_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Array_NonEmpty_Internal_traversable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_traversable1NonEmptyArray')))->{'traverse1'})($dictApply_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->{'identity'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable10" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'] ?? \PhpursThunks::eval('Data_Array_NonEmpty_Internal_foldable1NonEmptyArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Traversable_traversableArray'] ?? \PhpursThunks::eval('Data_Traversable_traversableArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_eqNonEmptyArray'] = function() { $v = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["eq" => (($GLOBALS['Data_Eq_eqArrayImpl'] ?? \PhpursThunks::eval('Data_Eq_eqArrayImpl')))(($dictEq_0)->{'eq'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_eq1NonEmptyArray'] = function() { $v = ($GLOBALS['Data_Eq_eq1Array'] ?? \PhpursThunks::eval('Data_Eq_eq1Array')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_bindNonEmptyArray'] = function() { $v = ($GLOBALS['Control_Bind_bindArray'] ?? \PhpursThunks::eval('Control_Bind_bindArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_applyNonEmptyArray'] = function() { $v = ($GLOBALS['Control_Apply_applyArray'] ?? \PhpursThunks::eval('Control_Apply_applyArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_applicativeNonEmptyArray'] = function() { $v = ($GLOBALS['Control_Applicative_applicativeArray'] ?? \PhpursThunks::eval('Control_Applicative_applicativeArray')); return $v; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_altNonEmptyArray'] = function() { $v = ($GLOBALS['Control_Alt_altArray'] ?? \PhpursThunks::eval('Control_Alt_altArray')); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Array_NonEmpty_Internal = \call_user_func(function() {
  $exports = [];
$foldr1Impl = function($f, $xs = null) use (&$foldr1Impl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foldr1Impl) {

            return $foldr1Impl(...\array_merge($__args, $more));
        };
    }
    $acc = $xs[\count($xs) - 1];
    for ($i = \count($xs) - 2; $i >= 0; $i--) {
        $acc = $f($xs[$i])($acc);
    }
    return $acc;
};

$foldl1Impl = function($f, $xs = null) use (&$foldl1Impl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foldl1Impl) {

            return $foldl1Impl(...\array_merge($__args, $more));
        };
    }
    $acc = $xs[0];
    $len = \count($xs);
    for ($i = 1; $i < $len; $i++) {
        $acc = $f($acc)($xs[$i]);
    }
    return $acc;
};

$traverse1Impl = function($apply, $map = null, $f = null) use (&$traverse1Impl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$traverse1Impl) {

            return $traverse1Impl(...\array_merge($__args, $more));
        };
    }

    $emptyList = new \stdClass();

    $consList = function($x) {
        return function($xs) use ($x) {
            $obj = new \stdClass();
            $obj->head = $x;
            $obj->tail = $xs;
            return $obj;
        };
    };

    $finalCell = function($head) use ($emptyList) {
        $obj = new \stdClass();
        $obj->head = $head;
        $obj->tail = $emptyList;
        return $obj;
    };

    $listToArray = function($list) use ($emptyList) {
        $arr = [];
        $xs = $list;
        while ($xs !== $emptyList) {
            $arr[] = $xs->head;
            $xs = $xs->tail;
        }
        return $arr;
    };

    $buildFrom = function($x, $ys) use ($apply, $map, $f, $consList) {
        return $apply($map($consList)($f($x)))($ys);
    };

    $go = function($acc, $currentLen, $xs) use (&$go, $buildFrom) {
        if ($currentLen === 0) {
            return $acc;
        } else {
            $last = $xs[$currentLen - 1];
            $cont = new \stdClass();
            $cont->_isCont = true;
            $cont->fn = function() use (&$go, $buildFrom, $last, $acc, $currentLen, $xs) {
                return $go($buildFrom($last, $acc), $currentLen - 1, $xs);
            };
            return $cont;
        }
    };

    return function($array) use ($map, $finalCell, $f, $go, $listToArray) {
        $acc = $map($finalCell)($f($array[\count($array) - 1]));
        $result = $go($acc, \count($array) - 1, $array);
        while (isset($result->_isCont) && $result->_isCont) {
            $fn = $result->fn;
            $result = $fn();
        }
        return $map($listToArray)($result);
    };
};

$exports['foldr1Impl'] = $foldr1Impl;
$exports['foldl1Impl'] = $foldl1Impl;
$exports['traverse1Impl'] = $traverse1Impl;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_foldl1Impl'] = function() use (&$ffi_Data_Array_NonEmpty_Internal) { return $ffi_Data_Array_NonEmpty_Internal['foldl1Impl']; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_foldr1Impl'] = function() use (&$ffi_Data_Array_NonEmpty_Internal) { return $ffi_Data_Array_NonEmpty_Internal['foldr1Impl']; };
\PhpursThunks::$thunks['Data_Array_NonEmpty_Internal_traverse1Impl'] = function() use (&$ffi_Data_Array_NonEmpty_Internal) { return $ffi_Data_Array_NonEmpty_Internal['traverse1Impl']; };























