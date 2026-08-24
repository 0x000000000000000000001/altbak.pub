<?php

namespace Data\Foldable;

// ALL IMPORTS: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Const, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.HeytingAlgebra, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Const, Data.Either, Data.Eq, Data.Foldable, Data.Function, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.HeytingAlgebra, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Endo, Data.Monoid.Multiplicative, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Semiring, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Maybe.First/index.php';
require_once __DIR__ . '/../Data.Maybe.Last/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Monoid.Conj/index.php';
require_once __DIR__ . '/../Data.Monoid.Disj/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Endo/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Foldable = \call_user_func(function() {
  $exports = [];
$foldrArray = function($f, $init, $xs) use (&$foldrArray) {
    
    $acc = $init;
    for ($i = \count($xs) - 1; $i >= 0; $i--) {
        $f1 = $f($xs[$i]);
        $acc = $f1($acc);
    }
    return $acc;
};
$exports['foldrArray'] = $foldrArray;

$foldlArray = function($f, $init, $xs) use (&$foldlArray) {
    
    $acc = $init;
    for ($i = 0, $len = \count($xs); $i < $len; $i++) {
        $f1 = $f($acc);
        $acc = $f1($xs[$i]);
    }
    return $acc;
};
$exports['foldlArray'] = $foldlArray;

return $exports;
  return $exports;
});
function majData_majFoldable_foldlmajArray($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFoldable_foldlmajArray';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_Foldable;
  $f = (\array_key_exists('foldlArray', $ffi_Data_Foldable) ? $ffi_Data_Foldable['foldlArray'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_Foldable_foldlArray'] = __NAMESPACE__ . '\\majData_majFoldable_foldlmajArray';

function majData_majFoldable_foldrmajArray($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majFoldable_foldrmajArray';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_Foldable;
  $f = (\array_key_exists('foldrArray', $ffi_Data_Foldable) ? $ffi_Data_Foldable['foldrArray'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_Foldable_foldrArray'] = __NAMESPACE__ . '\\majData_majFoldable_foldrmajArray';



final class Data_Foldable_Empty { public $tag = 'Empty'; public function __construct() {} }
final class Data_Foldable_Node { public $tag = 'Node'; public function __construct(public  $value0) {} }
final class Data_Foldable_Append { public $tag = 'Append'; public function __construct(public  $value0, public  $value1) {} }

// Data_Foldable_Empty
$GLOBALS['Data_Foldable_Empty'] = ($GLOBALS['__phpurs_data0_Empty'] ??= new \Data\Foldable\Data_Foldable_Empty());

// Data_Foldable_Node
$GLOBALS['Data_Foldable_Node'] = function($value0) {
  $__num = \func_num_args();
  $__res = new \Data\Foldable\Data_Foldable_Node($value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Foldable_Append
$GLOBALS['Data_Foldable_Append'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Foldable\Data_Foldable_Append($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Foldable_semigroupFreeMonoidTree
$GLOBALS['Data_Foldable_semigroupFreeMonoidTree'] = (object)["append" => $GLOBALS['Data_Foldable_Append']];

// Data_Foldable_monoidFreeMonoidTree
$GLOBALS['Data_Foldable_monoidFreeMonoidTree'] = (object)["mempty" => new \Data\Foldable\Data_Foldable_Empty(), "Semigroup0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_semigroupFreeMonoidTree'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldr
function majData_majFoldable_foldr($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldr';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldr'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_foldr'] = __NAMESPACE__ . '\\majData_majFoldable_foldr';

// Data_Foldable_indexr
function majData_majFoldable_indexr($dictFoldable_0, $idx_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_indexr';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'elem'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictFoldable_0)->{'foldr'})(function($a_2) use ($idx_1) {
  $__num = \func_num_args();
  $__res = function($cursor_3) use ($a_2, $idx_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($cursor_3)->{'elem'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = $cursor_3;
goto end_branch_1;;
};
  $__t1 = match (($cursor_3)->{'pos'}) { $idx_1 => (object)["elem" => new \Data\Maybe\Data_Maybe_Just($a_2), "pos" => ($cursor_3)->{'pos'}], default => (object)["pos" => (($cursor_3)->{'pos'} + 1), "elem" => ($cursor_3)->{'elem'}] };
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["elem" => new \Data\Maybe\Data_Maybe_Nothing(), "pos" => 0]));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_indexr'] = __NAMESPACE__ . '\\majData_majFoldable_indexr';

// Data_Foldable_null
function majData_majFoldable_null($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_null';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($dictFoldable_0)->{'foldr'})(function($v_1) {
  $__num = \func_num_args();
  $__res = function($v1_2) {
  $__num = \func_num_args();
  $__res = false;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(true);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_null'] = __NAMESPACE__ . '\\majData_majFoldable_null';

// Data_Foldable_oneOf
function majData_majFoldable_onemajOf($dictFoldable_0, $dictPlus_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_onemajOf';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldr'})(((($dictPlus_1)->{'Alt0'})(null))->{'alt'}))(($dictPlus_1)->{'empty'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_oneOf'] = __NAMESPACE__ . '\\majData_majFoldable_onemajOf';

// Data_Foldable_oneOfMap
function majData_majFoldable_onemajOfmajMap($dictFoldable_0, $dictPlus_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_onemajOfmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $alt_2_0 = ((($dictPlus_1)->{'Alt0'})(null))->{'alt'};
  $empty_3_1 = ($dictPlus_1)->{'empty'};
  $__res = function($f_4) use ($alt_2_0, $dictFoldable_0, $empty_3_1) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($alt_2_0))($f_4)))($empty_3_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_oneOfMap'] = __NAMESPACE__ . '\\majData_majFoldable_onemajOfmajMap';

// Data_Foldable_traverse_
function majData_majFoldable_traverse_($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_traverse_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $applySecond_1_0 = function($a_3) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_2_1, $__local_var_1_0, $a_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})(function($v_5) {
  $__num = \func_num_args();
  $__res = function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_3)))($b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictFoldable_2) use ($applySecond_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($applySecond_1_0, $dictApplicative_0, $dictFoldable_2) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_2)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($applySecond_1_0))($f_3)))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_Unit_unit']));
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
$GLOBALS['Data_Foldable_traverse_'] = __NAMESPACE__ . '\\majData_majFoldable_traverse_';

// Data_Foldable_for_
function majData_majFoldable_for_($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_for_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_2_1 = (($__local_var_1_0)->{'Functor0'})(null);
  $applySecond_1_0 = function($a_3) use ($Functor0_2_1, $__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Functor0_2_1, $__local_var_1_0, $a_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})(function($v_5) {
  $__num = \func_num_args();
  $__res = function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_3)))($b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictFoldable_2) use ($applySecond_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($applySecond_1_0, $dictApplicative_0, $dictFoldable_2) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($applySecond_1_0, $b_3, $dictApplicative_0, $dictFoldable_2) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_2)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($applySecond_1_0))($a_4)))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_Unit_unit'])))($b_3);
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
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_for_'] = __NAMESPACE__ . '\\majData_majFoldable_for_';

// Data_Foldable_sequence_
function majData_majFoldable_sequence_($dictApplicative_0, $dictFoldable_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_sequence_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $Functor0_3_1 = (($__local_var_2_0)->{'Functor0'})(null);
  $__res = ((($dictFoldable_1)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_4) use ($Functor0_3_1, $__local_var_2_0) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($Functor0_3_1, $__local_var_2_0, $a_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_0)->{'apply'})(((($Functor0_3_1)->{'map'})(function($v_6) {
  $__num = \func_num_args();
  $__res = function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_4)))($b_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_Unit_unit']));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_sequence_'] = __NAMESPACE__ . '\\majData_majFoldable_sequence_';

// Data_Foldable_foldl
function majData_majFoldable_foldl($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldl'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_foldl'] = __NAMESPACE__ . '\\majData_majFoldable_foldl';

// Data_Foldable_indexl
function majData_majFoldable_indexl($dictFoldable_0, $idx_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_indexl';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'elem'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictFoldable_0)->{'foldl'})(function($cursor_2) use ($idx_1) {
  $__num = \func_num_args();
  $__res = function($a_3) use ($cursor_2, $idx_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($cursor_2)->{'elem'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = $cursor_2;
goto end_branch_1;;
};
  $__t1 = match (($cursor_2)->{'pos'}) { $idx_1 => (object)["elem" => new \Data\Maybe\Data_Maybe_Just($a_3), "pos" => ($cursor_2)->{'pos'}], default => (object)["pos" => (($cursor_2)->{'pos'} + 1), "elem" => ($cursor_2)->{'elem'}] };
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["elem" => new \Data\Maybe\Data_Maybe_Nothing(), "pos" => 0]));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_indexl'] = __NAMESPACE__ . '\\majData_majFoldable_indexl';

// Data_Foldable_intercalate
function majData_majFoldable_intercalate($dictFoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_intercalate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = function($sep_4) use ($Semigroup0_2_0, $dictFoldable_0, $mempty_3_1) {
  $__num = \func_num_args();
  $__res = function($xs_5) use ($Semigroup0_2_0, $dictFoldable_0, $mempty_3_1, $sep_4) {
  $__num = \func_num_args();
  $__res = ((((($dictFoldable_0)->{'foldl'})(function($v_6) use ($Semigroup0_2_0, $sep_4) {
  $__num = \func_num_args();
  $__res = function($v1_7) use ($Semigroup0_2_0, $sep_4, $v_6) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($v_6)->{'init'}) {
$__t2 = (object)["init" => false, "acc" => $v1_7];
goto end_branch_2;;
};
  $__t2 = (object)["init" => false, "acc" => ((($Semigroup0_2_0)->{'append'})(($v_6)->{'acc'}))(((($Semigroup0_2_0)->{'append'})($sep_4))($v1_7))];
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["init" => true, "acc" => $mempty_3_1]))($xs_5))->{'acc'};
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_intercalate'] = __NAMESPACE__ . '\\majData_majFoldable_intercalate';

// Data_Foldable_length
function majData_majFoldable_length($dictFoldable_0, $dictSemiring_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_length';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})(function($c_2) use ($dictSemiring_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($c_2, $dictSemiring_1) {
  $__num = \func_num_args();
  $__res = ((($dictSemiring_1)->{'add'})(($dictSemiring_1)->{'one'}))($c_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictSemiring_1)->{'zero'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_length'] = __NAMESPACE__ . '\\majData_majFoldable_length';

// Data_Foldable_maximumBy
function majData_majFoldable_maximummajBy($dictFoldable_0, $cmp_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_maximummajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})(function($v_2) use ($cmp_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($cmp_1, $v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v1_3);
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = null;;
if ((($cmp_1)(($v_2)->{'value0'}))($v1_3) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = ($v_2)->{'value0'};
goto end_branch_1;;
};
$__t1 = $v1_3;
end_branch_1:;
$__t0 = new \Data\Maybe\Data_Maybe_Just($__t1);
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
}))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_maximumBy'] = __NAMESPACE__ . '\\majData_majFoldable_maximummajBy';

// Data_Foldable_maximum
function majData_majFoldable_maximum($dictOrd_0, $dictFoldable_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_maximum';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_1)->{'foldl'})(function($v_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictOrd_0, $v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v1_3);
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = null;;
if (((($dictOrd_0)->{'compare'})(($v_2)->{'value0'}))($v1_3) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = ($v_2)->{'value0'};
goto end_branch_1;;
};
$__t1 = $v1_3;
end_branch_1:;
$__t0 = new \Data\Maybe\Data_Maybe_Just($__t1);
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
}))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_maximum'] = __NAMESPACE__ . '\\majData_majFoldable_maximum';

// Data_Foldable_minimumBy
function majData_majFoldable_minimummajBy($dictFoldable_0, $cmp_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_minimummajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})(function($v_2) use ($cmp_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($cmp_1, $v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v1_3);
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = null;;
if ((($cmp_1)(($v_2)->{'value0'}))($v1_3) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = ($v_2)->{'value0'};
goto end_branch_1;;
};
$__t1 = $v1_3;
end_branch_1:;
$__t0 = new \Data\Maybe\Data_Maybe_Just($__t1);
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
}))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_minimumBy'] = __NAMESPACE__ . '\\majData_majFoldable_minimummajBy';

// Data_Foldable_minimum
function majData_majFoldable_minimum($dictOrd_0, $dictFoldable_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_minimum';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_1)->{'foldl'})(function($v_2) use ($dictOrd_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictOrd_0, $v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v1_3);
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = null;;
if (((($dictOrd_0)->{'compare'})(($v_2)->{'value0'}))($v1_3) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = ($v_2)->{'value0'};
goto end_branch_1;;
};
$__t1 = $v1_3;
end_branch_1:;
$__t0 = new \Data\Maybe\Data_Maybe_Just($__t1);
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
}))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_minimum'] = __NAMESPACE__ . '\\majData_majFoldable_minimum';

// Data_Foldable_product
function majData_majFoldable_product($dictFoldable_0, $dictSemiring_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_product';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})(($dictSemiring_1)->{'mul'}))(($dictSemiring_1)->{'one'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_product'] = __NAMESPACE__ . '\\majData_majFoldable_product';

// Data_Foldable_sum
function majData_majFoldable_sum($dictFoldable_0, $dictSemiring_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_sum';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})(($dictSemiring_1)->{'add'}))(($dictSemiring_1)->{'zero'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_sum'] = __NAMESPACE__ . '\\majData_majFoldable_sum';

// Data_Foldable_foldableTuple
$GLOBALS['Data_Foldable_foldableTuple'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)(($v_2)->{'value1'}))($z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))(($v_2)->{'value1'});
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)(($v_2)->{'value1'});
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

// Data_Foldable_foldableMultiplicative
$GLOBALS['Data_Foldable_foldableMultiplicative'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($v_2))($z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))($v_2);
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)($v_2);
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

// Data_Foldable_foldableMaybe
$GLOBALS['Data_Foldable_foldableMaybe'] = (object)["foldr" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = (($v_0)(($v2_2)->{'value0'}))($v1_1);
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $v1_1;
goto end_branch_1;;
};
  if ($v2_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (($v_0)($v1_1))(($v2_2)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = function($v_2) use ($mempty_1_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($mempty_1_2, $v_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = $mempty_1_2;
goto end_branch_3;;
};
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = ($v_2)(($v1_3)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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

// Data_Foldable_foldableIdentity
$GLOBALS['Data_Foldable_foldableIdentity'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($v_2))($z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))($v_2);
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)($v_2);
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

// Data_Foldable_foldableEither
$GLOBALS['Data_Foldable_foldableEither'] = (object)["foldr" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = $v1_1;
goto end_branch_0;;
};
  if ($v2_2 instanceof \Data\Either\Data_Either_Right) {
$__t0 = (($v_0)(($v2_2)->{'value0'}))($v1_1);
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($v2_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v2_2 instanceof \Data\Either\Data_Either_Left) {
$__t1 = $v1_1;
goto end_branch_1;;
};
  if ($v2_2 instanceof \Data\Either\Data_Either_Right) {
$__t1 = (($v_0)($v1_1))(($v2_2)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = function($v_2) use ($mempty_1_2) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($mempty_1_2, $v_2) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v1_3 instanceof \Data\Either\Data_Either_Left) {
$__t3 = $mempty_1_2;
goto end_branch_3;;
};
  if ($v1_3 instanceof \Data\Either\Data_Either_Right) {
$__t3 = ($v_2)(($v1_3)->{'value0'});
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
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

// Data_Foldable_foldableDual
$GLOBALS['Data_Foldable_foldableDual'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($v_2))($z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))($v_2);
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)($v_2);
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

// Data_Foldable_foldableDisj
$GLOBALS['Data_Foldable_foldableDisj'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($v_2))($z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))($v_2);
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)($v_2);
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

// Data_Foldable_foldableConst
$GLOBALS['Data_Foldable_foldableConst'] = (object)["foldr" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($z_1) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($z_1) {
  $__num = \func_num_args();
  $__res = $z_1;
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
}, "foldl" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($z_1) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($z_1) {
  $__num = \func_num_args();
  $__res = $z_1;
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__res = function($v_2) use ($mempty_1_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($mempty_1_0) {
  $__num = \func_num_args();
  $__res = $mempty_1_0;
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

// Data_Foldable_foldableConj
$GLOBALS['Data_Foldable_foldableConj'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($v_2))($z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))($v_2);
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)($v_2);
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

// Data_Foldable_foldableAdditive
$GLOBALS['Data_Foldable_foldableAdditive'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($v_2))($z_1);
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
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__res = (($f_0)($z_1))($v_2);
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_1) {
  $__num = \func_num_args();
  $__res = ($f_1)($v_2);
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

// Data_Foldable_foldMapDefaultR
function majData_majFoldable_foldmajMapmajDefaultmajR($dictFoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldmajMapmajDefaultmajR';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_4) use ($Semigroup0_2_0, $dictFoldable_0, $mempty_3_1) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)->{'foldr'})(function($x_5) use ($Semigroup0_2_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($acc_6) use ($Semigroup0_2_0, $f_4, $x_5) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_2_0)->{'append'})(($f_4)($x_5)))($acc_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_3_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_foldMapDefaultR'] = __NAMESPACE__ . '\\majData_majFoldable_foldmajMapmajDefaultmajR';

// Data_Foldable_foldableArray
$GLOBALS['Data_Foldable_foldableArray'] = (object)["foldr" => $GLOBALS['Data_Foldable_foldrArray'], "foldl" => $GLOBALS['Data_Foldable_foldlArray'], "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_1 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_0, $mempty_2_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Foldable_foldableArray'])->{'foldr'})(function($x_4) use ($Semigroup0_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = function($acc_5) use ($Semigroup0_1_0, $f_3, $x_4) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_0)->{'append'})(($f_3)($x_4)))($acc_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldMapDefaultL
function majData_majFoldable_foldmajMapmajDefaultmajL($dictFoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldmajMapmajDefaultmajL';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Semigroup0_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_4) use ($Semigroup0_2_0, $dictFoldable_0, $mempty_3_1) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)->{'foldl'})(function($acc_5) use ($Semigroup0_2_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($x_6) use ($Semigroup0_2_0, $acc_5, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_2_0)->{'append'})($acc_5))(($f_4)($x_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_3_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_foldMapDefaultL'] = __NAMESPACE__ . '\\majData_majFoldable_foldmajMapmajDefaultmajL';

// Data_Foldable_foldMap
function majData_majFoldable_foldmajMap($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'foldMap'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_foldMap'] = __NAMESPACE__ . '\\majData_majFoldable_foldmajMap';

// Data_Foldable_foldableApp
function majData_majFoldable_foldablemajApp($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldablemajApp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["foldr" => function($f_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($i_2) use ($dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFoldable_0, $f_1, $i_2) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldr'})($f_1))($i_2))($v_3);
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
}, "foldl" => function($f_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($i_2) use ($dictFoldable_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFoldable_0, $f_1, $i_2) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldl'})($f_1))($i_2))($v_3);
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
}, "foldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($f_2) use ($dictFoldable_0, $dictMonoid_1) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($dictFoldable_0, $dictMonoid_1, $f_2) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldMap'})($dictMonoid_1))($f_2))($v_3);
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
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_foldableApp'] = __NAMESPACE__ . '\\majData_majFoldable_foldablemajApp';

// Data_Foldable_foldableCompose
function majData_majFoldable_foldablemajCompose($dictFoldable_0, $dictFoldable1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldablemajCompose';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["foldr" => function($f_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($i_3) use ($dictFoldable1_1, $dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldable1_1, $dictFoldable_0, $f_2, $i_3) {
  $__num = \func_num_args();
  $__local_var_5_0 = (($dictFoldable1_1)->{'foldr'})($f_2);
  $__res = (((($dictFoldable_0)->{'foldr'})(function($b_6) use ($__local_var_5_0) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($__local_var_5_0, $b_6) {
  $__num = \func_num_args();
  $__res = (($__local_var_5_0)($a_7))($b_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($i_3))($v_4);
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
}, "foldl" => function($f_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($i_3) use ($dictFoldable1_1, $dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldable1_1, $dictFoldable_0, $f_2, $i_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldl'})((($dictFoldable1_1)->{'foldl'})($f_2)))($i_3))($v_4);
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
}, "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictFoldable1_1, $dictFoldable_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldable1_1, $dictFoldable_0, $dictMonoid_2, $f_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldMap'})($dictMonoid_2))(((($dictFoldable1_1)->{'foldMap'})($dictMonoid_2))($f_3)))($v_4);
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
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_foldableCompose'] = __NAMESPACE__ . '\\majData_majFoldable_foldablemajCompose';

// Data_Foldable_foldableCoproduct
function majData_majFoldable_foldablemajCoproduct($dictFoldable_0, $dictFoldable1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldablemajCoproduct';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["foldr" => function($f_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($z_3) use ($dictFoldable1_1, $dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ((($dictFoldable_0)->{'foldr'})($f_2))($z_3);
  $__local_var_5_1 = ((($dictFoldable1_1)->{'foldr'})($f_2))($z_3);
  $__res = function($v2_6) use ($__local_var_4_0, $__local_var_5_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v2_6 instanceof \Data\Either\Data_Either_Left) {
$__t2 = ($__local_var_4_0)(($v2_6)->{'value0'});
goto end_branch_2;;
};
  if ($v2_6 instanceof \Data\Either\Data_Either_Right) {
$__t2 = ($__local_var_5_1)(($v2_6)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($f_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($z_3) use ($dictFoldable1_1, $dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__local_var_4_3 = ((($dictFoldable_0)->{'foldl'})($f_2))($z_3);
  $__local_var_5_4 = ((($dictFoldable1_1)->{'foldl'})($f_2))($z_3);
  $__res = function($v2_6) use ($__local_var_4_3, $__local_var_5_4) {
  $__num = \func_num_args();
  $__t5 = null;;
  if ($v2_6 instanceof \Data\Either\Data_Either_Left) {
$__t5 = ($__local_var_4_3)(($v2_6)->{'value0'});
goto end_branch_5;;
};
  if ($v2_6 instanceof \Data\Either\Data_Either_Right) {
$__t5 = ($__local_var_5_4)(($v2_6)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
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
}, "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($dictFoldable1_1, $dictFoldable_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__local_var_4_6 = ((($dictFoldable_0)->{'foldMap'})($dictMonoid_2))($f_3);
  $__local_var_5_7 = ((($dictFoldable1_1)->{'foldMap'})($dictMonoid_2))($f_3);
  $__res = function($v2_6) use ($__local_var_4_6, $__local_var_5_7) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ($v2_6 instanceof \Data\Either\Data_Either_Left) {
$__t8 = ($__local_var_4_6)(($v2_6)->{'value0'});
goto end_branch_8;;
};
  if ($v2_6 instanceof \Data\Either\Data_Either_Right) {
$__t8 = ($__local_var_5_7)(($v2_6)->{'value0'});
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
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
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_foldableCoproduct'] = __NAMESPACE__ . '\\majData_majFoldable_foldablemajCoproduct';

// Data_Foldable_foldableFirst
$GLOBALS['Data_Foldable_foldableFirst'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $z_1;
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = (($f_0)(($v_2)->{'value0'}))($z_1);
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $z_1;
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (($f_0)($z_1))(($v_2)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictMonoid_0, $f_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = ($dictMonoid_0)->{'mempty'};
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($f_1)(($v_2)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldableLast
$GLOBALS['Data_Foldable_foldableLast'] = (object)["foldr" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = $z_1;
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = (($f_0)(($v_2)->{'value0'}))($z_1);
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldl" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($z_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($f_0, $z_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $z_1;
goto end_branch_1;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (($f_0)($z_1))(($v_2)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
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
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($f_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = function($v_2) use ($dictMonoid_0, $f_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = ($dictMonoid_0)->{'mempty'};
goto end_branch_2;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ($f_1)(($v_2)->{'value0'});
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldableProduct
function majData_majFoldable_foldablemajProduct($dictFoldable_0, $dictFoldable1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldablemajProduct';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["foldr" => function($f_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($z_3) use ($dictFoldable1_1, $dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldable1_1, $dictFoldable_0, $f_2, $z_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable_0)->{'foldr'})($f_2))((((($dictFoldable1_1)->{'foldr'})($f_2))($z_3))(($v_4)->{'value1'})))(($v_4)->{'value0'});
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
}, "foldl" => function($f_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($z_3) use ($dictFoldable1_1, $dictFoldable_0, $f_2) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($dictFoldable1_1, $dictFoldable_0, $f_2, $z_3) {
  $__num = \func_num_args();
  $__res = (((($dictFoldable1_1)->{'foldl'})($f_2))((((($dictFoldable_0)->{'foldl'})($f_2))($z_3))(($v_4)->{'value0'})))(($v_4)->{'value1'});
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
}, "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $Semigroup0_3_0 = (($dictMonoid_2)->{'Semigroup0'})(null);
  $__res = function($f_4) use ($Semigroup0_3_0, $dictFoldable1_1, $dictFoldable_0, $dictMonoid_2) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($Semigroup0_3_0, $dictFoldable1_1, $dictFoldable_0, $dictMonoid_2, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_3_0)->{'append'})((((($dictFoldable_0)->{'foldMap'})($dictMonoid_2))($f_4))(($v_5)->{'value0'})))((((($dictFoldable1_1)->{'foldMap'})($dictMonoid_2))($f_4))(($v_5)->{'value1'}));
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
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_foldableProduct'] = __NAMESPACE__ . '\\majData_majFoldable_foldablemajProduct';

// Data_Foldable_foldlDefault
function majData_majFoldable_foldlmajDefault($dictFoldable_0, $c_1 = null, $u_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldlmajDefault';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $go__go_4_0 = null;
  $go__go_4_0 = (function() use ($c_1, &$go__go_4_0) {
  $__fn = function($acc_5, $lhs_6 = null, $rhs_7 = null) use ($c_1, &$go__go_4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__go_4_0_0_acc_5 = $acc_5;
  $__tco_var_go__go_4_0_0_lhs_6 = $lhs_6;
  $__tco_var_go__go_4_0_0_rhs_7 = $rhs_7;
  tco_loop_go__go_4_0_0:;
  $acc_5 = $__tco_var_go__go_4_0_0_acc_5;
  $lhs_6 = $__tco_var_go__go_4_0_0_lhs_6;
  $rhs_7 = $__tco_var_go__go_4_0_0_rhs_7;
  $__t0 = null;;
  if ($lhs_6 instanceof \Data\Foldable\Data_Foldable_Node) {
$__tco_1 = (($c_1)($acc_5))(($lhs_6)->{'value0'});
$__tco_2 = $rhs_7;
$__tco_3 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_var_go__go_4_0_0_acc_5 = $__tco_1;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_2;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_3;
goto tco_loop_go__go_4_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  if ($lhs_6 instanceof \Data\Foldable\Data_Foldable_Append) {
$__t11 = null;;
if (($lhs_6)->{'value1'} instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_12 = $acc_5;
$__tco_13 = ($lhs_6)->{'value0'};
$__tco_14 = $rhs_7;
$__tco_var_go__go_4_0_0_acc_5 = $__tco_12;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_13;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_14;
goto tco_loop_go__go_4_0_0;;
$__t11 = null;
goto end_branch_11;;
};
$__t7 = null;;
if ($rhs_7 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_8 = $acc_5;
$__tco_9 = ($lhs_6)->{'value0'};
$__tco_10 = ($lhs_6)->{'value1'};
$__tco_var_go__go_4_0_0_acc_5 = $__tco_8;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_9;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_10;
goto tco_loop_go__go_4_0_0;;
$__t7 = null;
goto end_branch_7;;
};
$__tco_4 = $acc_5;
$__tco_5 = ($lhs_6)->{'value0'};
$__tco_6 = new \Data\Foldable\Data_Foldable_Append(($lhs_6)->{'value1'}, $rhs_7);
$__tco_var_go__go_4_0_0_acc_5 = $__tco_4;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_5;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_6;
goto tco_loop_go__go_4_0_0;;
$__t7 = null;
end_branch_7:;
$__t11 = $__t7;
end_branch_11:;
$__t0 = $__t11;
goto end_branch_0;;
};
  if ($lhs_6 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t18 = null;;
if ($rhs_7 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t18 = $acc_5;
goto end_branch_18;;
};
$__tco_15 = $acc_5;
$__tco_16 = $rhs_7;
$__tco_17 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_var_go__go_4_0_0_acc_5 = $__tco_15;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_16;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_17;
goto tco_loop_go__go_4_0_0;;
$__t18 = null;
end_branch_18:;
$__t0 = $__t18;
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
  $__res = ((($go__go_4_0)($u_2))((((($dictFoldable_0)->{'foldMap'})($GLOBALS['Data_Foldable_monoidFreeMonoidTree']))($GLOBALS['Data_Foldable_Node']))($xs_3)))(new \Data\Foldable\Data_Foldable_Empty());
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Foldable_foldlDefault'] = __NAMESPACE__ . '\\majData_majFoldable_foldlmajDefault';

// Data_Foldable_foldrDefault
function majData_majFoldable_foldrmajDefault($dictFoldable_0, $c_1 = null, $u_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldrmajDefault';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $go__go_4_0 = null;
  $go__go_4_0 = (function() use ($c_1, &$go__go_4_0) {
  $__fn = function($acc_5, $lhs_6 = null, $rhs_7 = null) use ($c_1, &$go__go_4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__go_4_0_0_acc_5 = $acc_5;
  $__tco_var_go__go_4_0_0_lhs_6 = $lhs_6;
  $__tco_var_go__go_4_0_0_rhs_7 = $rhs_7;
  tco_loop_go__go_4_0_0:;
  $acc_5 = $__tco_var_go__go_4_0_0_acc_5;
  $lhs_6 = $__tco_var_go__go_4_0_0_lhs_6;
  $rhs_7 = $__tco_var_go__go_4_0_0_rhs_7;
  $__t0 = null;;
  if ($rhs_7 instanceof \Data\Foldable\Data_Foldable_Node) {
$__tco_1 = (($c_1)(($rhs_7)->{'value0'}))($acc_5);
$__tco_2 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_3 = $lhs_6;
$__tco_var_go__go_4_0_0_acc_5 = $__tco_1;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_2;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_3;
goto tco_loop_go__go_4_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  if ($rhs_7 instanceof \Data\Foldable\Data_Foldable_Append) {
$__t11 = null;;
if (($rhs_7)->{'value0'} instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_12 = $acc_5;
$__tco_13 = $lhs_6;
$__tco_14 = ($rhs_7)->{'value1'};
$__tco_var_go__go_4_0_0_acc_5 = $__tco_12;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_13;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_14;
goto tco_loop_go__go_4_0_0;;
$__t11 = null;
goto end_branch_11;;
};
$__t7 = null;;
if ($lhs_6 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_8 = $acc_5;
$__tco_9 = ($rhs_7)->{'value0'};
$__tco_10 = ($rhs_7)->{'value1'};
$__tco_var_go__go_4_0_0_acc_5 = $__tco_8;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_9;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_10;
goto tco_loop_go__go_4_0_0;;
$__t7 = null;
goto end_branch_7;;
};
$__tco_4 = $acc_5;
$__tco_5 = new \Data\Foldable\Data_Foldable_Append($lhs_6, ($rhs_7)->{'value0'});
$__tco_6 = ($rhs_7)->{'value1'};
$__tco_var_go__go_4_0_0_acc_5 = $__tco_4;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_5;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_6;
goto tco_loop_go__go_4_0_0;;
$__t7 = null;
end_branch_7:;
$__t11 = $__t7;
end_branch_11:;
$__t0 = $__t11;
goto end_branch_0;;
};
  if ($rhs_7 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t18 = null;;
if ($lhs_6 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t18 = $acc_5;
goto end_branch_18;;
};
$__tco_15 = $acc_5;
$__tco_16 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_17 = $lhs_6;
$__tco_var_go__go_4_0_0_acc_5 = $__tco_15;
$__tco_var_go__go_4_0_0_lhs_6 = $__tco_16;
$__tco_var_go__go_4_0_0_rhs_7 = $__tco_17;
goto tco_loop_go__go_4_0_0;;
$__t18 = null;
end_branch_18:;
$__t0 = $__t18;
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
  $__res = ((($go__go_4_0)($u_2))(new \Data\Foldable\Data_Foldable_Empty()))((((($dictFoldable_0)->{'foldMap'})($GLOBALS['Data_Foldable_monoidFreeMonoidTree']))($GLOBALS['Data_Foldable_Node']))($xs_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Foldable_foldrDefault'] = __NAMESPACE__ . '\\majData_majFoldable_foldrmajDefault';

// Data_Foldable_lookup
function majData_majFoldable_lookup($dictFoldable_0, $dictEq_1 = null, $a_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_lookup';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(((($dictFoldable_0)->{'foldMap'})($GLOBALS['Data_Maybe_First_monoidFirst']))(function($v_3) use ($a_2, $dictEq_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (((($dictEq_1)->{'eq'})($a_2))(($v_3)->{'value0'})) {
$__t0 = new \Data\Maybe\Data_Maybe_Just(($v_3)->{'value1'});
goto end_branch_0;;
};
  $__t0 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Foldable_lookup'] = __NAMESPACE__ . '\\majData_majFoldable_lookup';

// Data_Foldable_surroundMap
function majData_majFoldable_surroundmajMap($dictFoldable_0, $dictSemigroup_1 = null, $d_2 = null, $t_3 = null, $f_4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_surroundmajMap';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $semigroupEndo1_5_0 = (object)["append" => function($v_5) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($v_5) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_5))($v1_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ((((($dictFoldable_0)->{'foldMap'})((object)["mempty" => function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_6) use ($semigroupEndo1_5_0) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_5_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($a_5) use ($d_2, $dictSemigroup_1, $t_3) {
  $__num = \func_num_args();
  $__res = function($m_6) use ($a_5, $d_2, $dictSemigroup_1, $t_3) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_1)->{'append'})($d_2))(((($dictSemigroup_1)->{'append'})(($t_3)($a_5)))($m_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_4))($d_2);
  goto __end;;
  __end:
  return 5 < $__num ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
}
$GLOBALS['Data_Foldable_surroundMap'] = __NAMESPACE__ . '\\majData_majFoldable_surroundmajMap';

// Data_Foldable_surround
function majData_majFoldable_surround($dictFoldable_0, $dictSemigroup_1 = null, $d_2 = null, $f_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_surround';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $semigroupEndo1_4_0 = (object)["append" => function($v_4) {
  $__num = \func_num_args();
  $__res = function($v1_5) use ($v_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($v_4))($v1_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ((((($dictFoldable_0)->{'foldMap'})((object)["mempty" => function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Semigroup0" => function($_dollar___unused_5) use ($semigroupEndo1_4_0) {
  $__num = \func_num_args();
  $__res = $semigroupEndo1_4_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($a_4) use ($d_2, $dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = function($m_5) use ($a_4, $d_2, $dictSemigroup_1) {
  $__num = \func_num_args();
  $__res = ((($dictSemigroup_1)->{'append'})($d_2))(((($dictSemigroup_1)->{'append'})($a_4))($m_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_3))($d_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Foldable_surround'] = __NAMESPACE__ . '\\majData_majFoldable_surround';

// Data_Foldable_foldM
function majData_majFoldable_foldmajM($dictFoldable_0, $dictMonad_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldmajM';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Bind1_2_0 = (($dictMonad_1)->{'Bind1'})(null);
  $Applicative0_3_1 = (($dictMonad_1)->{'Applicative0'})(null);
  $__res = function($f_4) use ($Applicative0_3_1, $Bind1_2_0, $dictFoldable_0) {
  $__num = \func_num_args();
  $__res = function($b0_5) use ($Applicative0_3_1, $Bind1_2_0, $dictFoldable_0, $f_4) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)->{'foldl'})(function($b_6) use ($Bind1_2_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($a_7) use ($Bind1_2_0, $b_6, $f_4) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_0)->{'bind'})($b_6))(function($a_8) use ($a_7, $f_4) {
  $__num = \func_num_args();
  $__res = (($f_4)($a_8))($a_7);
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
}))((($Applicative0_3_1)->{'pure'})($b0_5));
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_foldM'] = __NAMESPACE__ . '\\majData_majFoldable_foldmajM';

// Data_Foldable_fold
function majData_majFoldable_fold($dictFoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_fold';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldMap'})($dictMonoid_1))(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_fold'] = __NAMESPACE__ . '\\majData_majFoldable_fold';

// Data_Foldable_findMap
function majData_majFoldable_findmajMap($dictFoldable_0, $p_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_findmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})(function($v_2) use ($p_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($p_1, $v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = ($p_1)($v1_3);
goto end_branch_0;;
};
  $__t0 = $v_2;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_findMap'] = __NAMESPACE__ . '\\majData_majFoldable_findmajMap';

// Data_Foldable_find
function majData_majFoldable_find($dictFoldable_0, $p_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_find';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})(function($v_2) use ($p_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($p_1, $v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if (($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing && ($p_1)($v1_3))) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v1_3);
goto end_branch_0;;
};
  $__t0 = $v_2;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_find'] = __NAMESPACE__ . '\\majData_majFoldable_find';

// Data_Foldable_any
function majData_majFoldable_any($dictFoldable_0, $dictHeytingAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_any';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupDisj1_2_0 = (object)["append" => function($v_2) use ($dictHeytingAlgebra_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictHeytingAlgebra_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictHeytingAlgebra_1)->{'disj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'ff'}, "Semigroup0" => function($_dollar___unused_3) use ($semigroupDisj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_any'] = __NAMESPACE__ . '\\majData_majFoldable_any';

// Data_Foldable_elem
function majData_majFoldable_elem($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_elem';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $semigroupDisj1_1_0 = (object)["append" => function($v_1) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($v_1) {
  $__num = \func_num_args();
  $__res = ($v_1 || $v1_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $any1_1_0 = (($dictFoldable_0)->{'foldMap'})((object)["mempty" => false, "Semigroup0" => function($_dollar___unused_2) use ($semigroupDisj1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  $__res = function($dictEq_2) use ($any1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($any1_1_0))(($dictEq_2)->{'eq'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_elem'] = __NAMESPACE__ . '\\majData_majFoldable_elem';

// Data_Foldable_notElem
function majData_majFoldable_notmajElem($dictFoldable_0, $dictEq_1 = null, $x_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_notmajElem';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $semigroupDisj1_3_0 = (object)["append" => function($v_3) {
  $__num = \func_num_args();
  $__res = function($v1_4) use ($v_3) {
  $__num = \func_num_args();
  $__res = ($v_3 || $v1_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_HeytingAlgebra_boolNot']))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl((($dictFoldable_0)->{'foldMap'})((object)["mempty" => false, "Semigroup0" => function($_dollar___unused_4) use ($semigroupDisj1_3_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_3_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]), ($dictEq_1)->{'eq'}, $x_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Foldable_notElem'] = __NAMESPACE__ . '\\majData_majFoldable_notmajElem';

// Data_Foldable_or
function majData_majFoldable_or($dictFoldable_0, $dictHeytingAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_or';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupDisj1_2_0 = (object)["append" => function($v_2) use ($dictHeytingAlgebra_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictHeytingAlgebra_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictHeytingAlgebra_1)->{'disj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ((($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'ff'}, "Semigroup0" => function($_dollar___unused_3) use ($semigroupDisj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_or'] = __NAMESPACE__ . '\\majData_majFoldable_or';

// Data_Foldable_all
function majData_majFoldable_all($dictFoldable_0, $dictHeytingAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_all';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupConj1_2_0 = (object)["append" => function($v_2) use ($dictHeytingAlgebra_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictHeytingAlgebra_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictHeytingAlgebra_1)->{'conj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'tt'}, "Semigroup0" => function($_dollar___unused_3) use ($semigroupConj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupConj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_all'] = __NAMESPACE__ . '\\majData_majFoldable_all';

// Data_Foldable_and
function majData_majFoldable_and($dictFoldable_0, $dictHeytingAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_and';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupConj1_2_0 = (object)["append" => function($v_2) use ($dictHeytingAlgebra_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($dictHeytingAlgebra_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($dictHeytingAlgebra_1)->{'conj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ((($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'tt'}, "Semigroup0" => function($_dollar___unused_3) use ($semigroupConj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupConj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))(function($x_2) {
  $__num = \func_num_args();
  $__res = $x_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_and'] = __NAMESPACE__ . '\\majData_majFoldable_and';

