<?php

namespace Data\Array\NonEmpty;

// ALL IMPORTS: Control.Alternative, Control.Bind, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty, Data.Array.NonEmpty.Internal, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semiring, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Partial.Unsafe, Prelude, Prim, Safe.Coerce, Unsafe.Coerce
// TO REQUIRE: Control.Alternative, Control.Bind, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty, Data.Array.NonEmpty.Internal, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.Maybe, Data.NonEmpty, Data.Ord, Data.Ring, Data.Semigroup, Data.Semigroup.Foldable, Data.Semiring, Data.Tuple, Data.Unfoldable, Data.Unfoldable1, Partial.Unsafe, Prelude, Safe.Coerce, Unsafe.Coerce
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semigroup.Foldable/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
require_once __DIR__ . '/../Data.Unfoldable1/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Safe.Coerce/index.php';
require_once __DIR__ . '/../Unsafe.Coerce/index.php';

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




// Data_Array_NonEmpty_unsafeFromArray
function majData_majArray_majNonmajEmpty_unsafemajFrommajArray($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_unsafemajFrommajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_unsafeFromArray'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_unsafemajFrommajArray';

// Data_Array_NonEmpty_transpose_closure
$GLOBALS['Data_Array_NonEmpty_transpose_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_transpose']))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));

// Data_Array_NonEmpty_transpose
function majData_majArray_majNonmajEmpty_transpose($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_transpose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_transpose_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_transpose'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_transpose';

// Data_Array_NonEmpty_toArray
function majData_majArray_majNonmajEmpty_tomajArray($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_tomajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $v_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_toArray'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_tomajArray';

// Data_Array_NonEmpty_unionBy'
function majData_majArray_majNonmajEmpty_unionmajBy__prime__($eq_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_unionmajBy__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Data_Array_unionBy'])($eq_0))($xs_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_unionBy__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_unionmajBy__prime__';

// Data_Array_NonEmpty_union'
function majData_majArray_majNonmajEmpty_union__prime__($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_union__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_unionBy__prime__'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_union__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_union__prime__';

// Data_Array_NonEmpty_unionBy
function majData_majArray_majNonmajEmpty_unionmajBy($eq_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_unionmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Data_Array_unionBy'])($eq_0))($xs_1))))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_unionBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_unionmajBy';

// Data_Array_NonEmpty_union
function majData_majArray_majNonmajEmpty_union($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_unionBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_union'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_union';

// Data_Array_NonEmpty_unzip_closure
$GLOBALS['Data_Array_NonEmpty_unzip_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_0)->{'value0'}, ($v_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_unzip']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_unzip
function majData_majArray_majNonmajEmpty_unzip($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_unzip';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_unzip_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_unzip'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_unzip';

// Data_Array_NonEmpty_updateAt
function majData_majArray_majNonmajEmpty_updatemajAt(int $i_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_updatemajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_updateAt'])($i_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_updateAt'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_updatemajAt';

// Data_Array_NonEmpty_zip
function majData_majArray_majNonmajEmpty_zip($xs_0, $ys_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_zip';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_zipWithImpl'])($GLOBALS['Data_Tuple_Tuple'], $xs_0, $ys_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_zip'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_zip';

// Data_Array_NonEmpty_zipWith
function majData_majArray_majNonmajEmpty_zipmajWith($f_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_zipmajWith';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_zipWithImpl'])($f_0, $xs_1, $ys_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_zipWith'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_zipmajWith';

// Data_Array_NonEmpty_zipWithA
function majData_majArray_majNonmajEmpty_zipmajWithmajA($dictApplicative_0, $f_1 = null, $xs_2 = null, $ys_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_zipmajWithmajA';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $Apply0_4_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = \Data\Traversable\majData_majTraversable_traversemajArraymajImpl(($Apply0_4_0)->{'apply'}, ((($Apply0_4_0)->{'Functor0'})(null))->{'map'}, ($dictApplicative_0)->{'pure'}, function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($GLOBALS['Data_Array_zipWithImpl'])($f_1, $xs_2, $ys_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_zipWithA'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_zipmajWithmajA';

// Data_Array_NonEmpty_splitAt
function majData_majArray_majNonmajEmpty_splitmajAt(int $i_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_splitmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Array\majData_majArray_splitmajAt($i_0, $xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_splitAt'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_splitmajAt';

// Data_Array_NonEmpty_some
function majData_majArray_majNonmajEmpty_some($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_some';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Apply0_1_0 = (((($dictAlternative_0)->{'Applicative0'})(null))->{'Apply0'})(null);
  $Functor0_2_1 = (((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null);
  $__res = function($dictLazy_3) use ($Apply0_1_0, $Functor0_2_1, $dictAlternative_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(function($v_4) use ($Apply0_1_0, $Functor0_2_1, $dictAlternative_0, $dictLazy_3) {
  $__num = \func_num_args();
  $__res = ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_Array_cons']))($v_4)))((($dictLazy_3)->{'defer'})(function($v1_5) use ($dictAlternative_0, $dictLazy_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'alt'})(((((((($dictAlternative_0)->{'Applicative0'})(null))->{'Apply0'})(null))->{'apply'})(((((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Array_cons']))($v_4)))((($dictLazy_3)->{'defer'})(function($v1_6) use ($dictAlternative_0, $dictLazy_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Array_many'])($dictAlternative_0))($dictLazy_3))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((((($dictAlternative_0)->{'Applicative0'})(null))->{'pure'})([]));
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_some'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_some';

// Data_Array_NonEmpty_snoc'
function majData_majArray_majNonmajEmpty_snoc__prime__($xs_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_snoc__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Array\majData_majArray_snoc($xs_0, $x_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_snoc__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_snoc__prime__';

// Data_Array_NonEmpty_snoc
function majData_majArray_majNonmajEmpty_snoc($xs_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_snoc';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Array\majData_majArray_snoc($xs_0, $x_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_snoc'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_snoc';

// Data_Array_NonEmpty_singleton_closure
$GLOBALS['Data_Array_NonEmpty_singleton_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))($GLOBALS['Data_Array_singleton']);

// Data_Array_NonEmpty_singleton
function majData_majArray_majNonmajEmpty_singleton($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_singleton_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_singleton'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_singleton';

// Data_Array_NonEmpty_replicate
function majData_majArray_majNonmajEmpty_replicate(int $i_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_replicate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), 1, $i_0);
  $__t1 = null;;
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $i_0;
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = 1;
goto end_branch_1;;
};
  if ($v_2_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = 1;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = ($GLOBALS['Data_Array_replicateImpl'])($__t1, $x_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_replicate'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_replicate';

// Data_Array_NonEmpty_range
function majData_majArray_majNonmajEmpty_range(int $x_0, $y_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_range';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_rangeImpl'])($x_0, $y_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_range'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_range';

// Data_Array_NonEmpty_prependArray
function majData_majArray_majNonmajEmpty_prependmajArray($xs_0, $ys_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_prependmajArray';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Semigroup\majData_majSemigroup_concatmajArray($xs_0, $ys_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_prependArray'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_prependmajArray';

// Data_Array_NonEmpty_modifyAt
function majData_majArray_majNonmajEmpty_modifymajAt(int $i_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_modifymajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_modifyAt'])($i_0))($f_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_modifyAt'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_modifymajAt';

// Data_Array_NonEmpty_intersectBy'
function majData_majArray_majNonmajEmpty_intersectmajBy__prime__($eq_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_intersectmajBy__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Array_intersectBy'])($eq_0))($xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_intersectmajBy__prime__';

// Data_Array_NonEmpty_intersectBy
function majData_majArray_majNonmajEmpty_intersectmajBy($eq_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_intersectmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'])($eq_0))($xs_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_intersectBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_intersectmajBy';

// Data_Array_NonEmpty_intersect'
function majData_majArray_majNonmajEmpty_intersect__prime__($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_intersect__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_intersectBy__prime__'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_intersect__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_intersect__prime__';

// Data_Array_NonEmpty_intersect
function majData_majArray_majNonmajEmpty_intersect($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_intersect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_intersectBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_intersect'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_intersect';

// Data_Array_NonEmpty_intercalate
function majData_majArray_majNonmajEmpty_intercalate($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_intercalate';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $semigroupJoinWith1_1_0 = (object)["append" => function($v_1) use ($dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($dictSemigroup_0, $v_1) {
  $__num = \func_num_args();
  $__res = function($j_3) use ($dictSemigroup_0, $v1_2, $v_1) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_0)->{'append'})(($v_1)($j_3)))(((($dictSemigroup_0)->{'append'})($j_3))(($v1_2)($j_3)));
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
  $__res = function($a_2) use ($semigroupJoinWith1_1_0) {
  $__num = \func_num_args();
  $__res = function($foldable_3) use ($a_2, $semigroupJoinWith1_1_0) {
  $__num = \func_num_args();
  $append_4_1 = ($semigroupJoinWith1_1_0)->{'append'};
  $__res = (\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl(function($__local_var_5) use ($append_4_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_Internal_foldl1Impl'])($append_4_1, $__local_var_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($GLOBALS['Data_Functor_arrayMap'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))(function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))), $foldable_3))($a_2);
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_intercalate'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_intercalate';

// Data_Array_NonEmpty_insertAt
function majData_majArray_majNonmajEmpty_insertmajAt(int $i_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_insertmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_insertAt'])($i_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_insertAt'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_insertmajAt';

// Data_Array_NonEmpty_fromFoldable1
function majData_majArray_majNonmajEmpty_frommajFoldable1($dictFoldable1_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_frommajFoldable1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ((($dictFoldable1_0)->{'Foldable0'})(null))->{'foldr'};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))(function($__local_var_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_fromFoldableImpl'])($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_fromFoldable1'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_frommajFoldable1';

// Data_Array_NonEmpty_fromArray
function majData_majArray_majNonmajEmpty_frommajArray($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_frommajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ((count($xs_0) > 0)) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($xs_0);
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_fromArray'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_frommajArray';

// Data_Array_NonEmpty_fromFoldable
function majData_majArray_majNonmajEmpty_frommajFoldable($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_frommajFoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($dictFoldable_0)->{'foldr'};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromArray']))(function($__local_var_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_fromFoldableImpl'])($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_fromFoldable'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_frommajFoldable';

// Data_Array_NonEmpty_transpose'_closure
$GLOBALS['Data_Array_NonEmpty_transpose__prime___closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_fromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_transpose']))($GLOBALS['Unsafe_Coerce_unsafeCoerce']));

// Data_Array_NonEmpty_transpose'
function majData_majArray_majNonmajEmpty_transpose__prime__($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_transpose__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_transpose__prime___closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_transpose__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_transpose__prime__';

// Data_Array_NonEmpty_foldr1
function majData_majArray_majNonmajEmpty_foldr1($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_foldr1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_Internal_foldr1Impl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_foldr1'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_foldr1';

// Data_Array_NonEmpty_foldl1
function majData_majArray_majNonmajEmpty_foldl1($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_foldl1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_Internal_foldl1Impl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_foldl1'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_foldl1';

// Data_Array_NonEmpty_foldMap1
function majData_majArray_majNonmajEmpty_foldmajMap1($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_foldmajMap1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $append_1_0 = ($dictSemigroup_0)->{'append'};
  $__res = function($f_2) use ($append_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($__local_var_3) use ($append_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_Internal_foldl1Impl'])($append_1_0, $__local_var_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Functor_arrayMap'])($f_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_foldMap1'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_foldmajMap1';

// Data_Array_NonEmpty_fold1
function majData_majArray_majNonmajEmpty_fold1($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_fold1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $append_1_0 = ($dictSemigroup_0)->{'append'};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($__local_var_2) use ($append_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_NonEmpty_Internal_foldl1Impl'])($append_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Functor_arrayMap'])(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_fold1'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_fold1';

// Data_Array_NonEmpty_difference'
function majData_majArray_majNonmajEmpty_difference__prime__($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_difference__prime__';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $difference1_1_0 = ($GLOBALS['Data_Foldable_foldrArray'])(($GLOBALS['Data_Array_deleteBy'])(($dictEq_0)->{'eq'}));
  $__res = function($xs_2) use ($difference1_1_0) {
  $__num = \func_num_args();
  $__res = ($difference1_1_0)($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_difference__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_difference__prime__';

// Data_Array_NonEmpty_cons'
function majData_majArray_majNonmajEmpty_cons__prime__($x_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_cons__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Semigroup\majData_majSemigroup_concatmajArray([$x_0], $xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_cons__prime__'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_cons__prime__';

// Data_Array_NonEmpty_fromNonEmpty
function majData_majArray_majNonmajEmpty_frommajNonmajEmpty($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_frommajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Data\Semigroup\majData_majSemigroup_concatmajArray([($v_0)->{'value0'}], ($v_0)->{'value1'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_fromNonEmpty'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_frommajNonmajEmpty';

// Data_Array_NonEmpty_concatMap
function majData_majArray_majNonmajEmpty_concatmajMap($b_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_concatmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Control\Bind\majControl_majBind_arraymajBind($a_1, $b_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_concatMap'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_concatmajMap';

// Data_Array_NonEmpty_concat_closure
$GLOBALS['Data_Array_NonEmpty_concat_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_concat']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_toArray']))(($GLOBALS['Data_Functor_arrayMap'])($GLOBALS['Data_Array_NonEmpty_toArray']))));

// Data_Array_NonEmpty_concat
function majData_majArray_majNonmajEmpty_concat($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_concat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_concat_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_concat'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_concat';

// Data_Array_NonEmpty_appendArray
function majData_majArray_majNonmajEmpty_appendmajArray($xs_0, $ys_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_appendmajArray';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Semigroup\majData_majSemigroup_concatmajArray($xs_0, $ys_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_appendArray'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_appendmajArray';

// Data_Array_NonEmpty_alterAt
function majData_majArray_majNonmajEmpty_altermajAt(int $i_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_altermajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_alterAt'])($i_0))($f_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_alterAt'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_altermajAt';

// Data_Array_NonEmpty_head_closure
$GLOBALS['Data_Array_NonEmpty_head_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_head']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_head
function majData_majArray_majNonmajEmpty_head($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_head_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_head'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_head';

// Data_Array_NonEmpty_init_closure
$GLOBALS['Data_Array_NonEmpty_init_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_init']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_init
function majData_majArray_majNonmajEmpty_init($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_init';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_init_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_init'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_init';

// Data_Array_NonEmpty_last_closure
$GLOBALS['Data_Array_NonEmpty_last_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_last']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_last
function majData_majArray_majNonmajEmpty_last($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_last';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_last_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_last'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_last';

// Data_Array_NonEmpty_tail_closure
$GLOBALS['Data_Array_NonEmpty_tail_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_tail']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_tail
function majData_majArray_majNonmajEmpty_tail($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_tail_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_tail'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_tail';

// Data_Array_NonEmpty_uncons_closure
$GLOBALS['Data_Array_NonEmpty_uncons_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_uncons']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_uncons
function majData_majArray_majNonmajEmpty_uncons($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_uncons_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_uncons'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_uncons';

// Data_Array_NonEmpty_toNonEmpty_closure
$GLOBALS['Data_Array_NonEmpty_toNonEmpty_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__res = new \Data\NonEmpty\Data_NonEmpty_NonEmpty(($v_0)->{'head'}, ($v_0)->{'tail'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_uncons']);

// Data_Array_NonEmpty_toNonEmpty
function majData_majArray_majNonmajEmpty_tomajNonmajEmpty($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_tomajNonmajEmpty';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_toNonEmpty_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_toNonEmpty'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_tomajNonmajEmpty';

// Data_Array_NonEmpty_unsnoc_closure
$GLOBALS['Data_Array_NonEmpty_unsnoc_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_unsnoc']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_unsnoc
function majData_majArray_majNonmajEmpty_unsnoc($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_unsnoc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_unsnoc_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_unsnoc'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_unsnoc';

// Data_Array_NonEmpty_all
function majData_majArray_majNonmajEmpty_all($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_all';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_all'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_all'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_all';

// Data_Array_NonEmpty_any
function majData_majArray_majNonmajEmpty_any($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_any';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_any'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_any'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_any';

// Data_Array_NonEmpty_catMaybes_closure
$GLOBALS['Data_Array_NonEmpty_catMaybes_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_catMaybes']))($GLOBALS['Data_Array_NonEmpty_toArray']);

// Data_Array_NonEmpty_catMaybes
function majData_majArray_majNonmajEmpty_catmajMaybes($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_catmajMaybes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_catMaybes_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_catMaybes'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_catmajMaybes';

// Data_Array_NonEmpty_delete
function majData_majArray_majNonmajEmpty_delete($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_delete';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_deleteBy'])(($dictEq_0)->{'eq'}))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_delete'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_delete';

// Data_Array_NonEmpty_deleteAt
function majData_majArray_majNonmajEmpty_deletemajAt(int $i_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_deletemajAt';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_deleteAt'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_deleteAt'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_deletemajAt';

// Data_Array_NonEmpty_deleteBy
function majData_majArray_majNonmajEmpty_deletemajBy($f_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_deletemajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_deleteBy'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_deleteBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_deletemajBy';

// Data_Array_NonEmpty_difference
function majData_majArray_majNonmajEmpty_difference($dictEq_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_difference';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Foldable_foldrArray'])(($GLOBALS['Data_Array_deleteBy'])(($dictEq_0)->{'eq'})))($xs_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_difference'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_difference';

// Data_Array_NonEmpty_drop
function majData_majArray_majNonmajEmpty_drop(int $i_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_drop';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_drop'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_drop'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_drop';

// Data_Array_NonEmpty_dropEnd
function majData_majArray_majNonmajEmpty_dropmajEnd(int $i_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_dropmajEnd';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_dropEnd'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_dropEnd'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_dropmajEnd';

// Data_Array_NonEmpty_dropWhile
function majData_majArray_majNonmajEmpty_dropmajWhile($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_dropmajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_dropWhile'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_dropWhile'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_dropmajWhile';

// Data_Array_NonEmpty_elem
function majData_majArray_majNonmajEmpty_elem($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_elem';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($arr_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), function($v_3) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_3))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $arr_2);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = false;
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = true;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_elem'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_elem';

// Data_Array_NonEmpty_elemIndex
function majData_majArray_majNonmajEmpty_elemmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_elemmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findIndex'])(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_elemIndex'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_elemmajIndex';

// Data_Array_NonEmpty_elemLastIndex
function majData_majArray_majNonmajEmpty_elemmajLastmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_elemmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findLastIndex'])(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_elemLastIndex'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_elemmajLastmajIndex';

// Data_Array_NonEmpty_filter
function majData_majArray_majNonmajEmpty_filter($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_filter';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_filter'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_filter'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_filter';

// Data_Array_NonEmpty_filterA
function majData_majArray_majNonmajEmpty_filtermajA($dictApplicative_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_filtermajA';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Functor0_2_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $Apply0_3_2 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_2_0)->{'map'})(($GLOBALS['Data_Array_mapMaybe'])(function($v_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($v_3)->{'value1'}) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($v_3)->{'value0'});
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((((($GLOBALS['Data_Traversable_traverseArrayImpl'])(($Apply0_3_2)->{'apply'}))(((($Apply0_3_2)->{'Functor0'})(null))->{'map'}))(($dictApplicative_0)->{'pure'}))(function($x_4) use ($Functor0_2_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_2_0)->{'map'})(($GLOBALS['Data_Tuple_Tuple'])($x_4)))(($f_1)($x_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_filterA'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_filtermajA';

// Data_Array_NonEmpty_find
function majData_majArray_majNonmajEmpty_find($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_find';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_find'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_find'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_find';

// Data_Array_NonEmpty_findIndex
function majData_majArray_majNonmajEmpty_findmajIndex($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_findmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findIndex'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_findIndex'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_findmajIndex';

// Data_Array_NonEmpty_findLastIndex
function majData_majArray_majNonmajEmpty_findmajLastmajIndex($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_findmajLastmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findLastIndex'])($x_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_findLastIndex'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_findmajLastmajIndex';

// Data_Array_NonEmpty_findMap
function majData_majArray_majNonmajEmpty_findmajMap($p_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_findmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_findMap'])($p_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_findMap'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_findmajMap';

// Data_Array_NonEmpty_foldM
function majData_majArray_majNonmajEmpty_foldmajM($dictMonad_0, $f_1 = null, $acc_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_foldmajM';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $Applicative0_3_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_4_1 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($__local_var_5) use ($Applicative0_3_0, $Bind1_4_1, $acc_2, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_6) use ($Applicative0_3_0, $acc_2) {
  $__num = \func_num_args();
  $__res = (($Applicative0_3_0)->{'pure'})($acc_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_6) use ($Bind1_4_1, $acc_2, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($as_7) use ($Bind1_4_1, $a_6, $acc_2, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((($Bind1_4_1)->{'bind'})((($f_1)($acc_2))($a_6)))(function($b_prime__8) use ($as_7, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Array_foldM'])($dictMonad_0))($f_1))($b_prime__8))($as_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__local_var_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_foldM'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_foldmajM';

// Data_Array_NonEmpty_foldRecM
function majData_majArray_majNonmajEmpty_foldmajRecmajM($dictMonadRec_0, $f_1 = null, $acc_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_foldmajRecmajM';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $Monad0_3_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $Applicative0_4_1 = (($Monad0_3_0)->{'Applicative0'})(null);
  $Bind1_5_2 = (($Monad0_3_0)->{'Bind1'})(null);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($array_6) use ($Applicative0_4_1, $Bind1_5_2, $acc_2, $dictMonadRec_0, $f_1) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($o_7) use ($Applicative0_4_1, $Bind1_5_2, $array_6, $f_1) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((($o_7)->{'b'} >= count($array_6))) {
$__t3 = (($Applicative0_4_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($o_7)->{'a'}));
goto end_branch_3;;
};
  $__t3 = ((($Bind1_5_2)->{'bind'})((($f_1)(($o_7)->{'a'}))(($array_6)[($o_7)->{'b'}])))(function($res_prime__8) use ($Applicative0_4_1, $o_7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_4_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((object)["a" => $res_prime__8, "b" => (($o_7)->{'b'} + 1)]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $acc_2, "b" => 0]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_foldRecM'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_foldmajRecmajM';

// Data_Array_NonEmpty_index_closure
$GLOBALS['Data_Array_NonEmpty_index_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_index']))($GLOBALS['Data_Array_NonEmpty_toArray']);

// Data_Array_NonEmpty_index
function majData_majArray_majNonmajEmpty_index($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_index';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_index_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_index'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_index';

// Data_Array_NonEmpty_length_closure
$GLOBALS['Data_Array_NonEmpty_length_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_length']))($GLOBALS['Data_Array_NonEmpty_toArray']);

// Data_Array_NonEmpty_length
function majData_majArray_majNonmajEmpty_length($v_0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_length_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_length'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_length';

// Data_Array_NonEmpty_mapMaybe
function majData_majArray_majNonmajEmpty_mapmajMaybe($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_mapmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_mapMaybe'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_mapMaybe'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_mapmajMaybe';

// Data_Array_NonEmpty_notElem
function majData_majArray_majNonmajEmpty_notmajElem($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_notmajElem';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($arr_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), function($v_3) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_3))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $arr_2);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = true;
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = false;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_notElem'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_notmajElem';

// Data_Array_NonEmpty_partition
function majData_majArray_majNonmajEmpty_partition($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_partition';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_partition'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_partition'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_partition';

// Data_Array_NonEmpty_slice
function majData_majArray_majNonmajEmpty_slice(int $start_0, $end_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_slice';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_slice'])($start_0))($end_1)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_slice'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_slice';

// Data_Array_NonEmpty_span
function majData_majArray_majNonmajEmpty_span($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_span';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_span'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_span'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_span';

// Data_Array_NonEmpty_take
function majData_majArray_majNonmajEmpty_take(int $i_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_take';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_take'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_take'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_take';

// Data_Array_NonEmpty_takeEnd
function majData_majArray_majNonmajEmpty_takemajEnd(int $i_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_takemajEnd';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_takeEnd'])($i_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_takeEnd'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_takemajEnd';

// Data_Array_NonEmpty_takeWhile
function majData_majArray_majNonmajEmpty_takemajWhile($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_takemajWhile';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_takeWhile'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_takeWhile'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_takemajWhile';

// Data_Array_NonEmpty_toUnfoldable
function majData_majArray_majNonmajEmpty_tomajUnfoldable($dictUnfoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_tomajUnfoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_1) use ($dictUnfoldable_0) {
  $__num = \func_num_args();
  $len_2_0 = count($xs_1);
  $__res = ((($dictUnfoldable_0)->{'unfoldr'})(function($i_3) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($i_3 < $len_2_0)) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($xs_1)[$i_3], ($i_3 + 1)));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_toUnfoldable'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_tomajUnfoldable';

// Data_Array_NonEmpty_cons
function majData_majArray_majNonmajEmpty_cons($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_cons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_cons'])($x_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_cons'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_cons';

// Data_Array_NonEmpty_group
function majData_majArray_majNonmajEmpty_group($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_group';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eq_1_0 = ($dictEq_0)->{'eq'};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_2) use ($eq_1_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_groupmajBy($eq_1_0, $xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_group'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_group';

// Data_Array_NonEmpty_groupAllBy
function majData_majArray_majNonmajEmpty_groupmajAllmajBy($op_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_groupmajAllmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupAllBy'])($op_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_groupAllBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_groupmajAllmajBy';

// Data_Array_NonEmpty_groupAll
function majData_majArray_majNonmajEmpty_groupmajAll($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_groupmajAll';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupAllBy'])(($dictOrd_0)->{'compare'})))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_groupAll'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_groupmajAll';

// Data_Array_NonEmpty_groupBy
function majData_majArray_majNonmajEmpty_groupmajBy($op_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_groupmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupBy'])($op_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_groupBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_groupmajBy';

// Data_Array_NonEmpty_insert
function majData_majArray_majNonmajEmpty_insert($dictOrd_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_insert';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_insertBy'])(($dictOrd_0)->{'compare'}))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_insert'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_insert';

// Data_Array_NonEmpty_insertBy
function majData_majArray_majNonmajEmpty_insertmajBy($f_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_insertmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_insertBy'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_insertBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_insertmajBy';

// Data_Array_NonEmpty_intersperse
function majData_majArray_majNonmajEmpty_intersperse($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_intersperse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_intersperse'])($x_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_intersperse'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_intersperse';

// Data_Array_NonEmpty_mapWithIndex
function majData_majArray_majNonmajEmpty_mapmajWithmajIndex($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_mapmajWithmajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_mapWithIndex'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_mapmajWithmajIndex';

// Data_Array_NonEmpty_modifyAtIndices
function majData_majArray_majNonmajEmpty_modifymajAtmajIndices($dictFoldable_0, $is_1 = null, $f_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_modifymajAtmajIndices';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_3) use ($dictFoldable_0, $f_2, $is_1) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($dictFoldable_0, $f_2, $is_1, $xs_3, &$__fn) {
$result_4_0 = ($GLOBALS['Data_Array_ST_thawImpl'])($xs_3);
$_dollar___unused_5_1 = phpurs_execute_effect((((($dictFoldable_0)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($a_5) {
  $__num = \func_num_args();
  $__local_var_7_1 = phpurs_execute_effect($a_5);
  $f_prime__7_1 = phpurs_execute_effect(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $a_prime__8_3 = phpurs_execute_effect($b_6);
  $__res = phpurs_execute_effect(phpurs_execute_effect(($f_prime__7_1)($a_prime__8_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($i_5) use ($f_2, $result_4_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\ST\majData_majArray_majSmajT_modify($i_5, $f_2, $result_4_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Unit_unit']))($is_1));
return phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_4_0)));
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_modifyAtIndices'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_modifymajAtmajIndices';

// Data_Array_NonEmpty_nub
function majData_majArray_majNonmajEmpty_nub($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_nub';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nubBy'])(($dictOrd_0)->{'compare'})))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_nub'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_nub';

// Data_Array_NonEmpty_nubBy
function majData_majArray_majNonmajEmpty_nubmajBy($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_nubmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nubBy'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_nubBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_nubmajBy';

// Data_Array_NonEmpty_nubByEq
function majData_majArray_majNonmajEmpty_nubmajBymajEq($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_nubmajBymajEq';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nubByEq'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_nubByEq'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_nubmajBymajEq';

// Data_Array_NonEmpty_nubEq
function majData_majArray_majNonmajEmpty_nubmajEq($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_nubmajEq';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_nubByEq'])(($dictEq_0)->{'eq'})))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_nubEq'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_nubmajEq';

// Data_Array_NonEmpty_reverse_closure
$GLOBALS['Data_Array_NonEmpty_reverse_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_reverse']))($GLOBALS['Data_Array_NonEmpty_toArray']));

// Data_Array_NonEmpty_reverse
function majData_majArray_majNonmajEmpty_reverse($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_reverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_NonEmpty_reverse_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_reverse'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_reverse';

// Data_Array_NonEmpty_scanl
function majData_majArray_majNonmajEmpty_scanl($f_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_scanl';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_scanl'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_scanl'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_scanl';

// Data_Array_NonEmpty_scanr
function majData_majArray_majNonmajEmpty_scanr($f_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_scanr';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])((($GLOBALS['Data_Array_scanr'])($f_0))($x_1)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_scanr'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_scanr';

// Data_Array_NonEmpty_sort
function majData_majArray_majNonmajEmpty_sort($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_sort';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_sortmajBy($compare_1_0, $xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_sort'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_sort';

// Data_Array_NonEmpty_sortBy
function majData_majArray_majNonmajEmpty_sortmajBy($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_sortmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_sortBy'])($f_0)))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_sortBy'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_sortmajBy';

// Data_Array_NonEmpty_sortWith
function majData_majArray_majNonmajEmpty_sortmajWith($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_sortmajWith';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_sortBy'])(function($x_2) use ($dictOrd_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictOrd_0, $f_1, $x_2) {
  $__num = \func_num_args();
  $__res = ((($dictOrd_0)->{'compare'})(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_sortWith'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_sortmajWith';

// Data_Array_NonEmpty_updateAtIndices
function majData_majArray_majNonmajEmpty_updatemajAtmajIndices($dictFoldable_0, $pairs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_updatemajAtmajIndices';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_NonEmpty_unsafeFromArray']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($xs_2) use ($dictFoldable_0, $pairs_1) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($dictFoldable_0, $pairs_1, $xs_2, &$__fn) {
$result_3_0 = ($GLOBALS['Data_Array_ST_thawImpl'])($xs_2);
$_dollar___unused_4_1 = phpurs_execute_effect((((($dictFoldable_0)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_4) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($a_4) {
  $__num = \func_num_args();
  $__local_var_6_1 = phpurs_execute_effect($a_4);
  $f_prime__6_1 = phpurs_execute_effect(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $a_prime__7_3 = phpurs_execute_effect($b_5);
  $__res = phpurs_execute_effect(phpurs_execute_effect(($f_prime__6_1)($a_prime__7_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_4) use ($result_3_0) {
  $__num = \func_num_args();
  $__local_var_5_4 = ($v_4)->{'value1'};
  $__local_var_6_5 = ($v_4)->{'value0'};
  $__res = ($GLOBALS['Data_Array_ST_pokeImpl'])($__local_var_6_5, $__local_var_5_4, $result_3_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Unit_unit']))($pairs_1));
return phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_3_0)));
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_NonEmpty_toArray']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_updateAtIndices'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_updatemajAtmajIndices';

// Data_Array_NonEmpty_unsafeIndex
function majData_majArray_majNonmajEmpty_unsafemajIndex($_dollar___unused_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_unsafemajIndex';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((function() {
  $__fn = function($__local_var_1, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($__local_var_1)[$__local_var_2];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($GLOBALS['Data_Array_NonEmpty_toArray']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_unsafeIndex'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_unsafemajIndex';

// Data_Array_NonEmpty_toUnfoldable1
function majData_majArray_majNonmajEmpty_tomajUnfoldable1($dictUnfoldable1_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_majNonmajEmpty_tomajUnfoldable1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $len_2_0 = \Data\Array\NonEmpty\majData_majArray_majNonmajEmpty_length($xs_1);
  $__res = ((($dictUnfoldable1_0)->{'unfoldr1'})(function($i_3) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($i_3 < ($len_2_0 - 1))) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($i_3 + 1));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = new \Data\Tuple\Data_Tuple_Tuple((\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((function() {
  $__fn = function($__local_var_4, $__local_var_5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($__local_var_4)[$__local_var_5];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), $GLOBALS['Data_Array_NonEmpty_toArray'], $xs_1))($i_3), $__t1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_NonEmpty_toUnfoldable1'] = __NAMESPACE__ . '\\majData_majArray_majNonmajEmpty_tomajUnfoldable1';

