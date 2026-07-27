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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Foldable = \call_user_func(function() {
  $exports = [];
$foldrArray = function($f, $init = null, $xs = null) use (&$foldrArray) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foldrArray) {
            return $foldrArray(...\array_merge($__args, $more));
        };
    }
    
    $acc = $init;
    for ($i = \count($xs) - 1; $i >= 0; $i--) {
        $f1 = $f($xs[$i]);
        $acc = $f1($acc);
    }
    return $acc;
};
$exports['foldrArray'] = $foldrArray;

$foldlArray = function($f, $init = null, $xs = null) use (&$foldlArray) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foldlArray) {
            return $foldlArray(...\array_merge($__args, $more));
        };
    }
    
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
  $f = ($ffi_Data_Foldable['foldlArray'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Data_Foldable['foldrArray'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_Foldable_foldrArray'] = __NAMESPACE__ . '\\majData_majFoldable_foldrmajArray';



final class Data_Foldable_Empty { public $tag = 'Empty'; public function __construct() {} }
final class Data_Foldable_Node { public $tag = 'Node'; public function __construct(public  $value0) {} }
final class Data_Foldable_Append { public $tag = 'Append'; public function __construct(public  $value0, public  $value1) {} }

// Data_Foldable_identity
function majData_majFoldable_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_identity'] = __NAMESPACE__ . '\\majData_majFoldable_identity';

// Data_Foldable_monoidEndo
$GLOBALS['Data_Foldable_monoidEndo'] = (function() use (&$__fn) {
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
$GLOBALS['Data_Foldable_monoidFreeMonoidTree'] = (object)["mempty" => new \Data\Foldable\Data_Foldable_Empty(), "Semigroup0" => function($_dollar__unused_0) {
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
}))(((($dictFoldable_0)->{'foldr'})((function() use ($idx_1) {
  $__fn = function($a_2, $cursor_3 = null) use ($idx_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($cursor_3)->{'elem'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = $cursor_3;
goto end_branch_0;;
};
  if ((($cursor_3)->{'pos'} === $idx_1)) {
$__t0 = (object)["elem" => new \Data\Maybe\Data_Maybe_Just($a_2), "pos" => ($cursor_3)->{'pos'}];
goto end_branch_0;;
};
  $__t0 = (object)["pos" => (($cursor_3)->{'pos'} + 1), "elem" => ($cursor_3)->{'elem'}];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["elem" => new \Data\Maybe\Data_Maybe_Nothing(), "pos" => 0]));
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
  $__res = ((($dictFoldable_0)->{'foldr'})((function() {
  $__fn = function($v_1, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = false;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(true);
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
  $__res = (function() use ($__local_var_1_0, $dictApplicative_0) {
  $__fn = function($dictFoldable_2, $f_3 = null) use ($__local_var_1_0, $dictApplicative_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_2)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])((function() use ($__local_var_1_0) {
  $__fn = function($a_4, $b_5 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'apply'})(((((($__local_var_1_0)->{'Functor0'})(null))->{'map'})(function($v_6) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Apply_identity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a_4)))($b_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($f_3)))((($dictApplicative_0)->{'pure'})($GLOBALS['Data_Unit_unit']));
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
$GLOBALS['Data_Foldable_traverse_'] = __NAMESPACE__ . '\\majData_majFoldable_traverse_';

// Data_Foldable_for_
function majData_majFoldable_for_($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_for_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverse_1_1_0 = ($GLOBALS['Data_Foldable_traverse_'])($dictApplicative_0);
  $__res = function($dictFoldable_2) use ($traverse_1_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = ($traverse_1_1_0)($dictFoldable_2);
  $__res = (function() use ($__local_var_3_1) {
  $__fn = function($b_4, $a_5 = null) use ($__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_3_1)($a_5))($b_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
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
function majData_majFoldable_sequence_($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_sequence_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverse_1_1_0 = ($GLOBALS['Data_Foldable_traverse_'])($dictApplicative_0);
  $__res = function($dictFoldable_2) use ($traverse_1_1_0) {
  $__num = \func_num_args();
  $__res = (($traverse_1_1_0)($dictFoldable_2))($GLOBALS['Data_Foldable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
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
}))(((($dictFoldable_0)->{'foldl'})((function() use ($idx_1) {
  $__fn = function($cursor_2, $a_3 = null) use ($idx_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($cursor_2)->{'elem'} instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = $cursor_2;
goto end_branch_0;;
};
  if ((($cursor_2)->{'pos'} === $idx_1)) {
$__t0 = (object)["elem" => new \Data\Maybe\Data_Maybe_Just($a_3), "pos" => ($cursor_2)->{'pos'}];
goto end_branch_0;;
};
  $__t0 = (object)["pos" => (($cursor_2)->{'pos'} + 1), "elem" => ($cursor_2)->{'elem'}];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["elem" => new \Data\Maybe\Data_Maybe_Nothing(), "pos" => 0]));
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
  $__local_var_2_0 = (($dictMonoid_1)->{'Semigroup0'})(null);
  $mempty_3_1 = ($dictMonoid_1)->{'mempty'};
  $__res = (function() use ($__local_var_2_0, $dictFoldable_0, $mempty_3_1) {
  $__fn = function($sep_4, $xs_5 = null) use ($__local_var_2_0, $dictFoldable_0, $mempty_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictFoldable_0)->{'foldl'})((function() use ($__local_var_2_0, $sep_4) {
  $__fn = function($v_6, $v1_7 = null) use ($__local_var_2_0, $sep_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if (($v_6)->{'init'}) {
$__t2 = (object)["init" => false, "acc" => $v1_7];
goto end_branch_2;;
};
  $__t2 = (object)["init" => false, "acc" => ((($__local_var_2_0)->{'append'})(($v_6)->{'acc'}))(((($__local_var_2_0)->{'append'})($sep_4))($v1_7))];
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["init" => true, "acc" => $mempty_3_1]))($xs_5))->{'acc'};
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
$GLOBALS['Data_Foldable_intercalate'] = __NAMESPACE__ . '\\majData_majFoldable_intercalate';

// Data_Foldable_length
function majData_majFoldable_length($dictFoldable_0, $dictSemiring_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_length';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $one_2_0 = ($dictSemiring_1)->{'one'};
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($dictSemiring_1, $one_2_0) {
  $__fn = function($c_3, $v_4 = null) use ($dictSemiring_1, $one_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictSemiring_1)->{'add'})($one_2_0))($c_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($dictSemiring_1)->{'zero'});
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
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($cmp_1) {
  $__fn = function($v_2, $v1_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v1_3);
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = null;;
if (((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_1)(($v_2)->{'value0'}))($v1_3)))(new \Data\Ordering\Data_Ordering_GT())) {
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_maximumBy'] = __NAMESPACE__ . '\\majData_majFoldable_maximummajBy';

// Data_Foldable_maximum
function majData_majFoldable_maximum($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_maximum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($dictFoldable_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_maximumBy'])($dictFoldable_2))($compare_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Foldable_maximum'] = __NAMESPACE__ . '\\majData_majFoldable_maximum';

// Data_Foldable_minimumBy
function majData_majFoldable_minimummajBy($dictFoldable_0, $cmp_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_minimummajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($cmp_1) {
  $__fn = function($v_2, $v1_3 = null) use ($cmp_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = new \Data\Maybe\Data_Maybe_Just($v1_3);
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = null;;
if (((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_1)(($v_2)->{'value0'}))($v1_3)))(new \Data\Ordering\Data_Ordering_LT())) {
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\Maybe\Data_Maybe_Nothing());
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_minimumBy'] = __NAMESPACE__ . '\\majData_majFoldable_minimummajBy';

// Data_Foldable_minimum
function majData_majFoldable_minimum($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_minimum';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($dictFoldable_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_minimumBy'])($dictFoldable_2))($compare_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
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
$GLOBALS['Data_Foldable_foldableTuple'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)(($v_2)->{'value1'}))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()];

// Data_Foldable_foldableMultiplicative
$GLOBALS['Data_Foldable_foldableMultiplicative'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()];

// Data_Foldable_foldableMaybe
$GLOBALS['Data_Foldable_foldableMaybe'] = (object)["foldr" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = (function() use ($mempty_1_2) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldableIdentity
$GLOBALS['Data_Foldable_foldableIdentity'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()];

// Data_Foldable_foldableEither
$GLOBALS['Data_Foldable_foldableEither'] = (object)["foldr" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = (function() use ($mempty_1_2) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldableDual
$GLOBALS['Data_Foldable_foldableDual'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()];

// Data_Foldable_foldableDisj
$GLOBALS['Data_Foldable_foldableDisj'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()];

// Data_Foldable_foldableConst
$GLOBALS['Data_Foldable_foldableConst'] = (object)["foldr" => (function() {
  $__fn = function($v_0, $z_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = $z_1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($v_0, $z_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = $z_1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__res = (function() use ($mempty_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($mempty_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $mempty_1_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldableConj
$GLOBALS['Data_Foldable_foldableConj'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()];

// Data_Foldable_foldableAdditive
$GLOBALS['Data_Foldable_foldableAdditive'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($v_2))($z_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($f_0)($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => (function() {
  $__fn = function($dictMonoid_0, $f_1 = null, $v_2 = null) use (&$__fn) {
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
})()];

// Data_Foldable_foldMapDefaultR
function majData_majFoldable_foldmajMapmajDefaultmajR($dictFoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldmajMapmajDefaultmajR';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $mempty_2_0 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_3) use ($dictFoldable_0, $dictMonoid_1, $mempty_2_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)->{'foldr'})((function() use ($dictMonoid_1, $f_3) {
  $__fn = function($x_4, $acc_5 = null) use ($dictMonoid_1, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'})(($f_3)($x_4)))($acc_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_2_0);
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
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Foldable_foldableArray'])->{'foldr'})((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($x_3, $acc_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})(($f_2)($x_3)))($acc_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldableFreeMonoidTree
$GLOBALS['Data_Foldable_foldableFreeMonoidTree'] = (object)["foldl" => function($fn_0) {
  $__num = \func_num_args();
  $go__1_0 = null;
  $go__1_0 = (function() use ($fn_0, &$go__1_0) {
  $__fn = function($acc_2, $lhs_3 = null, $rhs_4 = null) use ($fn_0, &$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__1_0_0_acc_2 = $acc_2;
  $__tco_var_go__1_0_0_lhs_3 = $lhs_3;
  $__tco_var_go__1_0_0_rhs_4 = $rhs_4;
  tco_loop_go__1_0_0:;
  $acc_2 = $__tco_var_go__1_0_0_acc_2;
  $lhs_3 = $__tco_var_go__1_0_0_lhs_3;
  $rhs_4 = $__tco_var_go__1_0_0_rhs_4;
  $__t0 = null;;
  if ($lhs_3 instanceof \Data\Foldable\Data_Foldable_Node) {
$__tco_1 = (($fn_0)($acc_2))(($lhs_3)->{'value0'});
$__tco_2 = $rhs_4;
$__tco_3 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_var_go__1_0_0_acc_2 = $__tco_1;
$__tco_var_go__1_0_0_lhs_3 = $__tco_2;
$__tco_var_go__1_0_0_rhs_4 = $__tco_3;
goto tco_loop_go__1_0_0;;
$__t0 = null;
goto end_branch_0;;
};
  if ($lhs_3 instanceof \Data\Foldable\Data_Foldable_Append) {
$__t7 = null;;
if (($lhs_3)->{'value1'} instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_8 = $acc_2;
$__tco_9 = ($lhs_3)->{'value0'};
$__tco_10 = $rhs_4;
$__tco_var_go__1_0_0_acc_2 = $__tco_8;
$__tco_var_go__1_0_0_lhs_3 = $__tco_9;
$__tco_var_go__1_0_0_rhs_4 = $__tco_10;
goto tco_loop_go__1_0_0;;
$__t7 = null;
goto end_branch_7;;
};
if ($rhs_4 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_11 = $acc_2;
$__tco_12 = ($lhs_3)->{'value0'};
$__tco_13 = ($lhs_3)->{'value1'};
$__tco_var_go__1_0_0_acc_2 = $__tco_11;
$__tco_var_go__1_0_0_lhs_3 = $__tco_12;
$__tco_var_go__1_0_0_rhs_4 = $__tco_13;
goto tco_loop_go__1_0_0;;
$__t7 = null;
goto end_branch_7;;
};
$__tco_4 = $acc_2;
$__tco_5 = ($lhs_3)->{'value0'};
$__tco_6 = new \Data\Foldable\Data_Foldable_Append(($lhs_3)->{'value1'}, $rhs_4);
$__tco_var_go__1_0_0_acc_2 = $__tco_4;
$__tco_var_go__1_0_0_lhs_3 = $__tco_5;
$__tco_var_go__1_0_0_rhs_4 = $__tco_6;
goto tco_loop_go__1_0_0;;
$__t7 = null;
end_branch_7:;
$__t0 = $__t7;
goto end_branch_0;;
};
  if ($lhs_3 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t17 = null;;
if ($rhs_4 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t17 = $acc_2;
goto end_branch_17;;
};
$__tco_14 = $acc_2;
$__tco_15 = $rhs_4;
$__tco_16 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_var_go__1_0_0_acc_2 = $__tco_14;
$__tco_var_go__1_0_0_lhs_3 = $__tco_15;
$__tco_var_go__1_0_0_rhs_4 = $__tco_16;
goto tco_loop_go__1_0_0;;
$__t17 = null;
end_branch_17:;
$__t0 = $__t17;
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
  $__res = (function() use (&$go__1_0) {
  $__fn = function($a_2, $b_3 = null) use (&$go__1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($go__1_0)($a_2))($b_3))(new \Data\Foldable\Data_Foldable_Empty());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldr" => function($fn_0) {
  $__num = \func_num_args();
  $go__1_1 = null;
  $go__1_1 = (function() use ($fn_0, &$go__1_1) {
  $__fn = function($acc_2, $lhs_3 = null, $rhs_4 = null) use ($fn_0, &$go__1_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__tco_var_go__1_1_1_acc_2 = $acc_2;
  $__tco_var_go__1_1_1_lhs_3 = $lhs_3;
  $__tco_var_go__1_1_1_rhs_4 = $rhs_4;
  tco_loop_go__1_1_1:;
  $acc_2 = $__tco_var_go__1_1_1_acc_2;
  $lhs_3 = $__tco_var_go__1_1_1_lhs_3;
  $rhs_4 = $__tco_var_go__1_1_1_rhs_4;
  $__t1 = null;;
  if ($rhs_4 instanceof \Data\Foldable\Data_Foldable_Node) {
$__tco_2 = (($fn_0)(($rhs_4)->{'value0'}))($acc_2);
$__tco_3 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_4 = $lhs_3;
$__tco_var_go__1_1_1_acc_2 = $__tco_2;
$__tco_var_go__1_1_1_lhs_3 = $__tco_3;
$__tco_var_go__1_1_1_rhs_4 = $__tco_4;
goto tco_loop_go__1_1_1;;
$__t1 = null;
goto end_branch_1;;
};
  if ($rhs_4 instanceof \Data\Foldable\Data_Foldable_Append) {
$__t8 = null;;
if (($rhs_4)->{'value0'} instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_9 = $acc_2;
$__tco_10 = $lhs_3;
$__tco_11 = ($rhs_4)->{'value1'};
$__tco_var_go__1_1_1_acc_2 = $__tco_9;
$__tco_var_go__1_1_1_lhs_3 = $__tco_10;
$__tco_var_go__1_1_1_rhs_4 = $__tco_11;
goto tco_loop_go__1_1_1;;
$__t8 = null;
goto end_branch_8;;
};
if ($lhs_3 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__tco_12 = $acc_2;
$__tco_13 = ($rhs_4)->{'value0'};
$__tco_14 = ($rhs_4)->{'value1'};
$__tco_var_go__1_1_1_acc_2 = $__tco_12;
$__tco_var_go__1_1_1_lhs_3 = $__tco_13;
$__tco_var_go__1_1_1_rhs_4 = $__tco_14;
goto tco_loop_go__1_1_1;;
$__t8 = null;
goto end_branch_8;;
};
$__tco_5 = $acc_2;
$__tco_6 = new \Data\Foldable\Data_Foldable_Append($lhs_3, ($rhs_4)->{'value0'});
$__tco_7 = ($rhs_4)->{'value1'};
$__tco_var_go__1_1_1_acc_2 = $__tco_5;
$__tco_var_go__1_1_1_lhs_3 = $__tco_6;
$__tco_var_go__1_1_1_rhs_4 = $__tco_7;
goto tco_loop_go__1_1_1;;
$__t8 = null;
end_branch_8:;
$__t1 = $__t8;
goto end_branch_1;;
};
  if ($rhs_4 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t18 = null;;
if ($lhs_3 instanceof \Data\Foldable\Data_Foldable_Empty) {
$__t18 = $acc_2;
goto end_branch_18;;
};
$__tco_15 = $acc_2;
$__tco_16 = new \Data\Foldable\Data_Foldable_Empty();
$__tco_17 = $lhs_3;
$__tco_var_go__1_1_1_acc_2 = $__tco_15;
$__tco_var_go__1_1_1_lhs_3 = $__tco_16;
$__tco_var_go__1_1_1_rhs_4 = $__tco_17;
goto tco_loop_go__1_1_1;;
$__t18 = null;
end_branch_18:;
$__t1 = $__t18;
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
  $__res = (function() use (&$go__1_1) {
  $__fn = function($a_2, $b_3 = null) use (&$go__1_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($go__1_1)($a_2))(new \Data\Foldable\Data_Foldable_Empty()))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $mempty_1_2 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_2) use ($dictMonoid_0, $mempty_1_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Foldable_foldableFreeMonoidTree'])->{'foldr'})((function() use ($dictMonoid_0, $f_2) {
  $__fn = function($x_3, $acc_4 = null) use ($dictMonoid_0, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})(($f_2)($x_3)))($acc_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_1_2);
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
  $mempty_2_0 = ($dictMonoid_1)->{'mempty'};
  $__res = function($f_3) use ($dictFoldable_0, $dictMonoid_1, $mempty_2_0) {
  $__num = \func_num_args();
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($dictMonoid_1, $f_3) {
  $__fn = function($acc_4, $x_5 = null) use ($dictMonoid_1, $f_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_1)->{'Semigroup0'})(null))->{'append'})($acc_4))(($f_3)($x_5));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($mempty_2_0);
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
  $__res = (object)["foldr" => (function() use ($dictFoldable_0) {
  $__fn = function($f_1, $i_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)->{'foldr'})($f_1))($i_2))($v_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable_0) {
  $__fn = function($f_1, $i_2 = null, $v_3 = null) use ($dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)->{'foldl'})($f_1))($i_2))($v_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_1) use ($dictFoldable_0) {
  $__num = \func_num_args();
  $__res = (($dictFoldable_0)->{'foldMap'})($dictMonoid_1);
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
  $__res = (object)["foldr" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $i_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_5_0 = (($dictFoldable1_1)->{'foldr'})($f_2);
  $__res = (((($dictFoldable_0)->{'foldr'})((function() use ($__local_var_5_0) {
  $__fn = function($b_6, $a_7 = null) use ($__local_var_5_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_5_0)($a_7))($b_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($i_3))($v_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $i_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)->{'foldl'})((($dictFoldable1_1)->{'foldl'})($f_2)))($i_3))($v_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $foldMap4_3_1 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_2);
  $foldMap5_4_2 = (($dictFoldable1_1)->{'foldMap'})($dictMonoid_2);
  $__res = (function() use ($foldMap4_3_1, $foldMap5_4_2) {
  $__fn = function($f_5, $v_6 = null) use ($foldMap4_3_1, $foldMap5_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap4_3_1)(($foldMap5_4_2)($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
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
$GLOBALS['Data_Foldable_foldableCompose'] = __NAMESPACE__ . '\\majData_majFoldable_foldablemajCompose';

// Data_Foldable_foldableCoproduct
function majData_majFoldable_foldablemajCoproduct($dictFoldable_0, $dictFoldable1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldablemajCoproduct';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["foldr" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $foldMap4_3_6 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_2);
  $foldMap5_4_7 = (($dictFoldable1_1)->{'foldMap'})($dictMonoid_2);
  $__res = function($f_5) use ($foldMap4_3_6, $foldMap5_4_7) {
  $__num = \func_num_args();
  $__local_var_6_8 = ($foldMap4_3_6)($f_5);
  $__local_var_7_9 = ($foldMap5_4_7)($f_5);
  $__res = function($v2_8) use ($__local_var_6_8, $__local_var_7_9) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($v2_8 instanceof \Data\Either\Data_Either_Left) {
$__t10 = ($__local_var_6_8)(($v2_8)->{'value0'});
goto end_branch_10;;
};
  if ($v2_8 instanceof \Data\Either\Data_Either_Right) {
$__t10 = ($__local_var_7_9)(($v2_8)->{'value0'});
goto end_branch_10;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t10 = null;
  end_branch_10:;
  $__res = $__t10;
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
$GLOBALS['Data_Foldable_foldableFirst'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldr'})($f_0))($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldl'})($f_0))($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldMap'})($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Foldable_foldableLast
$GLOBALS['Data_Foldable_foldableLast'] = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldr'})($f_0))($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldl'})($f_0))($z_1))($v_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldableMaybe'])->{'foldMap'})($dictMonoid_0);
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
  $__res = (object)["foldr" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable_0)->{'foldr'})($f_2))((((($dictFoldable1_1)->{'foldr'})($f_2))($z_3))(($v_4)->{'value1'})))(($v_4)->{'value0'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($dictFoldable1_1, $dictFoldable_0) {
  $__fn = function($f_2, $z_3 = null, $v_4 = null) use ($dictFoldable1_1, $dictFoldable_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($dictFoldable1_1)->{'foldl'})($f_2))((((($dictFoldable_0)->{'foldl'})($f_2))($z_3))(($v_4)->{'value0'})))(($v_4)->{'value1'});
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_2) use ($dictFoldable1_1, $dictFoldable_0) {
  $__num = \func_num_args();
  $foldMap4_3_0 = (($dictFoldable_0)->{'foldMap'})($dictMonoid_2);
  $foldMap5_4_1 = (($dictFoldable1_1)->{'foldMap'})($dictMonoid_2);
  $__res = (function() use ($dictMonoid_2, $foldMap4_3_0, $foldMap5_4_1) {
  $__fn = function($f_5, $v_6 = null) use ($dictMonoid_2, $foldMap4_3_0, $foldMap5_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonoid_2)->{'Semigroup0'})(null))->{'append'})((($foldMap4_3_0)($f_5))(($v_6)->{'value0'})))((($foldMap5_4_1)($f_5))(($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
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
$GLOBALS['Data_Foldable_foldableProduct'] = __NAMESPACE__ . '\\majData_majFoldable_foldablemajProduct';

// Data_Foldable_foldlDefault
function majData_majFoldable_foldlmajDefault($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldlmajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldMap2_1_0 = (($dictFoldable_0)->{'foldMap'})($GLOBALS['Data_Foldable_monoidFreeMonoidTree']);
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($c_2, $u_3 = null, $xs_4 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Foldable_foldableFreeMonoidTree'])->{'foldl'})($c_2))($u_3))((($foldMap2_1_0)($GLOBALS['Data_Foldable_Node']))($xs_4));
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
$GLOBALS['Data_Foldable_foldlDefault'] = __NAMESPACE__ . '\\majData_majFoldable_foldlmajDefault';

// Data_Foldable_foldrDefault
function majData_majFoldable_foldrmajDefault($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldrmajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldMap2_1_0 = (($dictFoldable_0)->{'foldMap'})($GLOBALS['Data_Foldable_monoidFreeMonoidTree']);
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($c_2, $u_3 = null, $xs_4 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($GLOBALS['Data_Foldable_foldableFreeMonoidTree'])->{'foldr'})($c_2))($u_3))((($foldMap2_1_0)($GLOBALS['Data_Foldable_Node']))($xs_4));
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
$GLOBALS['Data_Foldable_foldrDefault'] = __NAMESPACE__ . '\\majData_majFoldable_foldrmajDefault';

// Data_Foldable_lookup
function majData_majFoldable_lookup($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_lookup';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldMap2_1_0 = (($dictFoldable_0)->{'foldMap'})($GLOBALS['Data_Maybe_First_monoidFirst']);
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($dictEq_2, $a_3 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(($foldMap2_1_0)(function($v_4) use ($a_3, $dictEq_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (((($dictEq_2)->{'eq'})($a_3))(($v_4)->{'value0'})) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($v_4)->{'value1'});
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
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
$GLOBALS['Data_Foldable_lookup'] = __NAMESPACE__ . '\\majData_majFoldable_lookup';

// Data_Foldable_surroundMap
function majData_majFoldable_surroundmajMap($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_surroundmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $foldMap2_1_0 = (($dictFoldable_0)->{'foldMap'})($GLOBALS['Data_Foldable_monoidEndo']);
  $__res = (function() use ($foldMap2_1_0) {
  $__fn = function($dictSemigroup_2, $d_3 = null, $t_4 = null, $f_5 = null) use ($foldMap2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($foldMap2_1_0)((function() use ($d_3, $dictSemigroup_2, $t_4) {
  $__fn = function($a_6, $m_7 = null) use ($d_3, $dictSemigroup_2, $t_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictSemigroup_2)->{'append'})($d_3))(((($dictSemigroup_2)->{'append'})(($t_4)($a_6)))($m_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($f_5))($d_3);
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
$GLOBALS['Data_Foldable_surroundMap'] = __NAMESPACE__ . '\\majData_majFoldable_surroundmajMap';

// Data_Foldable_surround
function majData_majFoldable_surround($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_surround';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $surroundMap1_1_0 = ($GLOBALS['Data_Foldable_surroundMap'])($dictFoldable_0);
  $__res = function($dictSemigroup_2) use ($surroundMap1_1_0) {
  $__num = \func_num_args();
  $surroundMap2_3_1 = ($surroundMap1_1_0)($dictSemigroup_2);
  $__res = function($d_4) use ($surroundMap2_3_1) {
  $__num = \func_num_args();
  $__res = (($surroundMap2_3_1)($d_4))($GLOBALS['Data_Foldable_identity']);
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
$GLOBALS['Data_Foldable_surround'] = __NAMESPACE__ . '\\majData_majFoldable_surround';

// Data_Foldable_foldM
function majData_majFoldable_foldmajM($dictFoldable_0, $dictMonad_1 = null, $f_2 = null, $b0_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_foldmajM';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($dictMonad_1, $f_2) {
  $__fn = function($b_4, $a_5 = null) use ($dictMonad_1, $f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonad_1)->{'Bind1'})(null))->{'bind'})($b_4))(function($a_6) use ($a_5, $f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)($a_6))($a_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((((($dictMonad_1)->{'Applicative0'})(null))->{'pure'})($b0_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Foldable_foldM'] = __NAMESPACE__ . '\\majData_majFoldable_foldmajM';

// Data_Foldable_fold
function majData_majFoldable_fold($dictFoldable_0, $dictMonoid_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_fold';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFoldable_0)->{'foldMap'})($dictMonoid_1))($GLOBALS['Data_Foldable_identity']);
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
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($p_1) {
  $__fn = function($v_2, $v1_3 = null) use ($p_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\Maybe\Data_Maybe_Nothing());
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
  $__res = ((($dictFoldable_0)->{'foldl'})((function() use ($p_1) {
  $__fn = function($v_2, $v1_3 = null) use ($p_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\Maybe\Data_Maybe_Nothing());
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
  $semigroupDisj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_1)->{'disj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'ff'}, "Semigroup0" => function($_dollar__unused_3) use ($semigroupDisj1_2_0) {
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
  $semigroupDisj1_1_0 = (object)["append" => (function() {
  $__fn = function($v_1, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $any1_1_0 = (($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'ff'}, "Semigroup0" => function($_dollar__unused_2) use ($semigroupDisj1_1_0) {
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
function majData_majFoldable_notmajElem($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_notmajElem';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $elem1_1_0 = ($GLOBALS['Data_Foldable_elem'])($dictFoldable_0);
  $__res = function($dictEq_2) use ($elem1_1_0) {
  $__num = \func_num_args();
  $elem2_3_1 = ($elem1_1_0)($dictEq_2);
  $__res = function($x_4) use ($elem2_3_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'}))(($elem2_3_1)($x_4));
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
$GLOBALS['Data_Foldable_notElem'] = __NAMESPACE__ . '\\majData_majFoldable_notmajElem';

// Data_Foldable_or
function majData_majFoldable_or($dictFoldable_0, $dictHeytingAlgebra_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majFoldable_or';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $semigroupDisj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_1)->{'disj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ((($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'ff'}, "Semigroup0" => function($_dollar__unused_3) use ($semigroupDisj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupDisj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))($GLOBALS['Data_Foldable_identity']);
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
  $semigroupConj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_1)->{'conj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'tt'}, "Semigroup0" => function($_dollar__unused_3) use ($semigroupConj1_2_0) {
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
  $semigroupConj1_2_0 = (object)["append" => (function() use ($dictHeytingAlgebra_1) {
  $__fn = function($v_2, $v1_3 = null) use ($dictHeytingAlgebra_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_1)->{'conj'})($v_2))($v1_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ((($dictFoldable_0)->{'foldMap'})((object)["mempty" => ($dictHeytingAlgebra_1)->{'tt'}, "Semigroup0" => function($_dollar__unused_3) use ($semigroupConj1_2_0) {
  $__num = \func_num_args();
  $__res = $semigroupConj1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]))($GLOBALS['Data_Foldable_identity']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Foldable_and'] = __NAMESPACE__ . '\\majData_majFoldable_and';

