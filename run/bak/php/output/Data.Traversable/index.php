<?php

namespace Data\Traversable;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Category, Control.Semigroupoid, Data.Const, Data.Either, Data.Foldable, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Traversable, Data.Traversable.Accum, Data.Traversable.Accum.Internal, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Category, Control.Semigroupoid, Data.Const, Data.Either, Data.Foldable, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Traversable, Data.Traversable.Accum, Data.Traversable.Accum.Internal, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Maybe.First/index.php';
require_once __DIR__ . '/../Data.Maybe.Last/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Monoid.Conj/index.php';
require_once __DIR__ . '/../Data.Monoid.Disj/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Traversable.Accum/index.php';
require_once __DIR__ . '/../Data.Traversable.Accum.Internal/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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
$ffi_Data_Traversable = \call_user_func(function() {
  $exports = [];
$traverseArrayImpl = function($apply, $map = null, $pure = null, $f = null, $array = null) use (&$traverseArrayImpl) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$traverseArrayImpl) {
            return $traverseArrayImpl(...\array_merge($__args, $more));
        };
    }

    $array1 = function ($a) { return [$a]; };
    $array2 = function ($a) { return function ($b) use ($a) { return [$a, $b]; }; };
    $array3 = function ($a) { return function ($b) use ($a) { return function ($c) use ($a, $b) { return [$a, $b, $c]; }; }; };
    $concat2 = function ($xs) { return function ($ys) use ($xs) { return \array_merge($xs, $ys); }; };
    
    $go = function ($bot, $top) use (&$go, $array, $apply, $map, $pure, $f, $array1, $array2, $array3, $concat2) {
        switch ($top - $bot) {
            case 0:
                return $pure([]);
            case 1:
                $f1 = $f($array[$bot]);
                $map1 = $map($array1);
                return $map1($f1);
            case 2:
                $f1 = $f($array[$bot]);
                $f2 = $f($array[$bot + 1]);
                $map1 = $map($array2);
                $map2 = $map1($f1);
                return $apply($map2)($f2);
            case 3:
                $f1 = $f($array[$bot]);
                $f2 = $f($array[$bot + 1]);
                $f3 = $f($array[$bot + 2]);
                $map1 = $map($array3);
                $map2 = $map1($f1);
                $app1 = $apply($map2)($f2);
                return $apply($app1)($f3);
            default:
                $pivot = $bot + floor(($top - $bot) / 4) * 2;
                $go1 = $go($bot, $pivot);
                $go2 = $go($pivot, $top);
                $map1 = $map($concat2);
                $map2 = $map1($go1);
                return $apply($map2)($go2);
        }
    };
    return $go(0, \count($array));
};
$exports['traverseArrayImpl'] = $traverseArrayImpl;

return $exports;
  return $exports;
});
function majData_majTraversable_traversemajArraymajImpl($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majTraversable_traversemajArraymajImpl';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_Traversable;
  $f = ($ffi_Data_Traversable['traverseArrayImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Traversable_traverseArrayImpl'] = __NAMESPACE__ . '\\majData_majTraversable_traversemajArraymajImpl';





// Data_Traversable_identity
function majData_majTraversable_identity($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_identity';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Traversable_identity'] = __NAMESPACE__ . '\\majData_majTraversable_identity';

// Data_Traversable_traverse
function majData_majTraversable_traverse($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_traverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'traverse'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Traversable_traverse'] = __NAMESPACE__ . '\\majData_majTraversable_traverse';

// Data_Traversable_traversableTuple
$GLOBALS['Data_Traversable_traversableTuple'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($GLOBALS['Data_Tuple_Tuple'])(($v_2)->{'value0'})))(($f_1)(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})(($GLOBALS['Data_Tuple_Tuple'])(($v_1)->{'value0'})))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_functorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableMultiplicative
$GLOBALS['Data_Traversable_traversableMultiplicative'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Multiplicative_Multiplicative']))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Multiplicative_Multiplicative']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Multiplicative_functorMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableMaybe
$GLOBALS['Data_Traversable_traversableMaybe'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v1_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Maybe_Just']))(($v_1)(($v1_2)->{'value0'}));
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
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Maybe_Just']))(($v_1)->{'value0'});
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
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableIdentity
$GLOBALS['Data_Traversable_traversableIdentity'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Identity_Identity']))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Identity_Identity']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Identity_functorIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableEither
$GLOBALS['Data_Traversable_traversableEither'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ($v1_2 instanceof \Data\Either\Data_Either_Left) {
$__t0 = (($dictApplicative_0)->{'pure'})(new \Data\Either\Data_Either_Left(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  if ($v1_2 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Either_Right']))(($v_1)(($v1_2)->{'value0'}));
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
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ($v_1 instanceof \Data\Either\Data_Either_Left) {
$__t1 = (($dictApplicative_0)->{'pure'})(new \Data\Either\Data_Either_Left(($v_1)->{'value0'}));
goto end_branch_1;;
};
  if ($v_1 instanceof \Data\Either\Data_Either_Right) {
$__t1 = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Either_Right']))(($v_1)->{'value0'});
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
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableDual
$GLOBALS['Data_Traversable_traversableDual'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Dual_Dual']))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Dual_Dual']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Dual_functorDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableDisj
$GLOBALS['Data_Traversable_traversableDisj'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Disj_Disj']))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Disj_Disj']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Disj_functorDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableConst
$GLOBALS['Data_Traversable_traversableConst'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)->{'pure'})($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictApplicative_0)->{'pure'})($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Const_functorConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableConj
$GLOBALS['Data_Traversable_traversableConj'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Conj_Conj']))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Conj_Conj']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Conj_functorConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableCompose
function majData_majTraversable_traversablemajCompose($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_traversablemajCompose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_Traversable_traversableCompose_dictTraversable_0 = $dictTraversable_0;
  tco_loop_Data_Traversable_traversableCompose:;
  $dictTraversable_0 = $__tco_var_Data_Traversable_traversableCompose_dictTraversable_0;
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $__local_var_2_1 = (($dictTraversable_0)->{'Foldable1'})(null);
  $__res = function($dictTraversable1_3) use ($__local_var_1_0, $__local_var_2_1, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictTraversable1_3)->{'Functor0'})(null);
  $functorCompose1_5_3 = (object)["map" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5, $v_6 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)->{'map'})((($__local_var_4_2)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_6_4 = (($dictTraversable1_3)->{'Foldable1'})(null);
  $foldableCompose1_7_5 = (object)["foldr" => (function() use ($__local_var_2_1, $__local_var_6_4) {
  $__fn = function($f_7, $i_8 = null, $v_9 = null) use ($__local_var_2_1, $__local_var_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_10_5 = (($__local_var_6_4)->{'foldr'})($f_7);
  $__res = (((($__local_var_2_1)->{'foldr'})((function() use ($__local_var_10_5) {
  $__fn = function($b_11, $a_12 = null) use ($__local_var_10_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_10_5)($a_12))($b_11);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($i_8))($v_9);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($__local_var_2_1, $__local_var_6_4) {
  $__fn = function($f_7, $i_8 = null, $v_9 = null) use ($__local_var_2_1, $__local_var_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_2_1)->{'foldl'})((($__local_var_6_4)->{'foldl'})($f_7)))($i_8))($v_9);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_7) use ($__local_var_2_1, $__local_var_6_4) {
  $__num = \func_num_args();
  $foldMap4_8_6 = (($__local_var_2_1)->{'foldMap'})($dictMonoid_7);
  $foldMap5_9_7 = (($__local_var_6_4)->{'foldMap'})($dictMonoid_7);
  $__res = (function() use ($foldMap4_8_6, $foldMap5_9_7) {
  $__fn = function($f_10, $v_11 = null) use ($foldMap4_8_6, $foldMap5_9_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap4_8_6)(($foldMap5_9_7)($f_10)))($v_11);
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
  $__res = (object)["traverse" => function($dictApplicative_8) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $traverse4_9_9 = (($dictTraversable_0)->{'traverse'})($dictApplicative_8);
  $traverse5_10_10 = (($dictTraversable1_3)->{'traverse'})($dictApplicative_8);
  $__res = (function() use ($dictApplicative_8, $traverse4_9_9, $traverse5_10_10) {
  $__fn = function($f_11, $v_12 = null) use ($dictApplicative_8, $traverse4_9_9, $traverse5_10_10, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_8)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Compose_Compose']))((($traverse4_9_9)(($traverse5_10_10)($f_11)))($v_12));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_8) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = (((\Data\Traversable\majData_majTraversable_traversablemajCompose($dictTraversable_0, $dictTraversable1_3))->{'traverse'})($dictApplicative_8))($GLOBALS['Data_Traversable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_8) use ($functorCompose1_5_3) {
  $__num = \func_num_args();
  $__res = $functorCompose1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_8) use ($foldableCompose1_7_5) {
  $__num = \func_num_args();
  $__res = $foldableCompose1_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Traversable_traversableCompose'] = __NAMESPACE__ . '\\majData_majTraversable_traversablemajCompose';

// Data_Traversable_traversableAdditive
$GLOBALS['Data_Traversable_traversableAdditive'] = (object)["traverse" => (function() {
  $__fn = function($dictApplicative_0, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Additive_Additive']))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Monoid_Additive_Additive']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Additive_functorAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_sequenceDefault
function majData_majTraversable_sequencemajDefault($dictTraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_sequencemajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictTraversable_0)->{'traverse'})($dictApplicative_1))($GLOBALS['Data_Traversable_identity']);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Traversable_sequenceDefault'] = __NAMESPACE__ . '\\majData_majTraversable_sequencemajDefault';

// Data_Traversable_traversableArray
$GLOBALS['Data_Traversable_traversableArray'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = ((($GLOBALS['Data_Traversable_traverseArrayImpl'])(($Apply0_1_0)->{'apply'}))(((($Apply0_1_0)->{'Functor0'})(null))->{'map'}))(($dictApplicative_0)->{'pure'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Traversable_traversableArray'])->{'traverse'})($dictApplicative_0))($GLOBALS['Data_Traversable_identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_functorArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_sequence
function majData_majTraversable_sequence($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_sequence';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'sequence'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Traversable_sequence'] = __NAMESPACE__ . '\\majData_majTraversable_sequence';

// Data_Traversable_traversableApp
function majData_majTraversable_traversablemajApp($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_traversablemajApp';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $__local_var_2_1 = (($dictTraversable_0)->{'Foldable1'})(null);
  $foldableApp_3_2 = (object)["foldr" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3, $i_4 = null, $v_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_2_1)->{'foldr'})($f_3))($i_4))($v_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3, $i_4 = null, $v_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_2_1)->{'foldl'})($f_3))($i_4))($v_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)->{'foldMap'})($dictMonoid_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["traverse" => function($dictApplicative_4) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $traverse3_5_3 = (($dictTraversable_0)->{'traverse'})($dictApplicative_4);
  $__res = (function() use ($dictApplicative_4, $traverse3_5_3) {
  $__fn = function($f_6, $v_7 = null) use ($dictApplicative_4, $traverse3_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_4)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_App_App']))((($traverse3_5_3)($f_6))($v_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_4) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $sequence3_5_4 = (($dictTraversable_0)->{'sequence'})($dictApplicative_4);
  $__res = function($v_6) use ($dictApplicative_4, $sequence3_5_4) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_4)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_App_App']))(($sequence3_5_4)($v_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_4) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_4) use ($foldableApp_3_2) {
  $__num = \func_num_args();
  $__res = $foldableApp_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Traversable_traversableApp'] = __NAMESPACE__ . '\\majData_majTraversable_traversablemajApp';

// Data_Traversable_traversableCoproduct
function majData_majTraversable_traversablemajCoproduct($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_traversablemajCoproduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorCoproduct_1_0 = ($GLOBALS['Data_Functor_Coproduct_functorCoproduct'])((($dictTraversable_0)->{'Functor0'})(null));
  $foldableCoproduct_2_1 = ($GLOBALS['Data_Foldable_foldableCoproduct'])((($dictTraversable_0)->{'Foldable1'})(null));
  $__res = function($dictTraversable1_3) use ($dictTraversable_0, $foldableCoproduct_2_1, $functorCoproduct_1_0) {
  $__num = \func_num_args();
  $functorCoproduct1_4_2 = ($functorCoproduct_1_0)((($dictTraversable1_3)->{'Functor0'})(null));
  $foldableCoproduct1_5_3 = ($foldableCoproduct_2_1)((($dictTraversable1_3)->{'Foldable1'})(null));
  $__res = (object)["traverse" => function($dictApplicative_6) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_7_4 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $traverse4_8_5 = (($dictTraversable_0)->{'traverse'})($dictApplicative_6);
  $traverse5_9_6 = (($dictTraversable1_3)->{'traverse'})($dictApplicative_6);
  $__res = function($f_10) use ($__local_var_7_4, $traverse4_8_5, $traverse5_9_6) {
  $__num = \func_num_args();
  $__local_var_11_7 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_7_4)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Functor_Coproduct_Coproduct']))($GLOBALS['Data_Either_Left']))))(($traverse4_8_5)($f_10));
  $__local_var_12_8 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_7_4)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Functor_Coproduct_Coproduct']))($GLOBALS['Data_Either_Right']))))(($traverse5_9_6)($f_10));
  $__res = function($v2_13) use ($__local_var_11_7, $__local_var_12_8) {
  $__num = \func_num_args();
  $__t9 = null;;
  if ($v2_13 instanceof \Data\Either\Data_Either_Left) {
$__t9 = ($__local_var_11_7)(($v2_13)->{'value0'});
goto end_branch_9;;
};
  if ($v2_13 instanceof \Data\Either\Data_Either_Right) {
$__t9 = ($__local_var_12_8)(($v2_13)->{'value0'});
goto end_branch_9;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t9 = null;
  end_branch_9:;
  $__res = $__t9;
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
}, "sequence" => function($dictApplicative_6) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_7_10 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_8_11 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_7_10)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Functor_Coproduct_Coproduct']))($GLOBALS['Data_Either_Left']))))((($dictTraversable_0)->{'sequence'})($dictApplicative_6));
  $__local_var_9_12 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_7_10)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Functor_Coproduct_Coproduct']))($GLOBALS['Data_Either_Right']))))((($dictTraversable1_3)->{'sequence'})($dictApplicative_6));
  $__res = function($v2_10) use ($__local_var_8_11, $__local_var_9_12) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t13 = ($__local_var_8_11)(($v2_10)->{'value0'});
goto end_branch_13;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t13 = ($__local_var_9_12)(($v2_10)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_6) use ($functorCoproduct1_4_2) {
  $__num = \func_num_args();
  $__res = $functorCoproduct1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_6) use ($foldableCoproduct1_5_3) {
  $__num = \func_num_args();
  $__res = $foldableCoproduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Traversable_traversableCoproduct'] = __NAMESPACE__ . '\\majData_majTraversable_traversablemajCoproduct';

// Data_Traversable_traversableFirst
$GLOBALS['Data_Traversable_traversableFirst'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse2_1_0 = (($GLOBALS['Data_Traversable_traversableMaybe'])->{'traverse'})($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverse2_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictApplicative_0, $traverse2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Maybe_First_First']))((($traverse2_1_0)($f_2))($v_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $sequence2_1_1 = (($GLOBALS['Data_Traversable_traversableMaybe'])->{'sequence'})($dictApplicative_0);
  $__res = function($v_2) use ($dictApplicative_0, $sequence2_1_1) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Maybe_First_First']))(($sequence2_1_1)($v_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableLast
$GLOBALS['Data_Traversable_traversableLast'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $traverse2_1_0 = (($GLOBALS['Data_Traversable_traversableMaybe'])->{'traverse'})($dictApplicative_0);
  $__res = (function() use ($dictApplicative_0, $traverse2_1_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictApplicative_0, $traverse2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Maybe_Last_Last']))((($traverse2_1_0)($f_2))($v_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $sequence2_1_1 = (($GLOBALS['Data_Traversable_traversableMaybe'])->{'sequence'})($dictApplicative_0);
  $__res = function($v_2) use ($dictApplicative_0, $sequence2_1_1) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Maybe_Last_Last']))(($sequence2_1_1)($v_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableLast'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableProduct
function majData_majTraversable_traversablemajProduct($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_traversablemajProduct';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $functorProduct_1_0 = ($GLOBALS['Data_Functor_Product_functorProduct'])((($dictTraversable_0)->{'Functor0'})(null));
  $foldableProduct_2_1 = ($GLOBALS['Data_Foldable_foldableProduct'])((($dictTraversable_0)->{'Foldable1'})(null));
  $__res = function($dictTraversable1_3) use ($dictTraversable_0, $foldableProduct_2_1, $functorProduct_1_0) {
  $__num = \func_num_args();
  $functorProduct1_4_2 = ($functorProduct_1_0)((($dictTraversable1_3)->{'Functor0'})(null));
  $foldableProduct1_5_3 = ($foldableProduct_2_1)((($dictTraversable1_3)->{'Foldable1'})(null));
  $__res = (object)["traverse" => function($dictApplicative_6) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_7_4 = (($dictApplicative_6)->{'Apply0'})(null);
  $traverse4_8_5 = (($dictTraversable_0)->{'traverse'})($dictApplicative_6);
  $traverse5_9_6 = (($dictTraversable1_3)->{'traverse'})($dictApplicative_6);
  $__res = (function() use ($__local_var_7_4, $traverse4_8_5, $traverse5_9_6) {
  $__fn = function($f_10, $v_11 = null) use ($__local_var_7_4, $traverse4_8_5, $traverse5_9_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_7_4)->{'apply'})(((((($__local_var_7_4)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))((($traverse4_8_5)($f_10))(($v_11)->{'value0'}))))((($traverse5_9_6)($f_10))(($v_11)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_6) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_7_7 = (($dictApplicative_6)->{'Apply0'})(null);
  $sequence4_8_8 = (($dictTraversable_0)->{'sequence'})($dictApplicative_6);
  $sequence5_9_9 = (($dictTraversable1_3)->{'sequence'})($dictApplicative_6);
  $__res = function($v_10) use ($__local_var_7_7, $sequence4_8_8, $sequence5_9_9) {
  $__num = \func_num_args();
  $__res = ((($__local_var_7_7)->{'apply'})(((((($__local_var_7_7)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))(($sequence4_8_8)(($v_10)->{'value0'}))))(($sequence5_9_9)(($v_10)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar__unused_6) use ($functorProduct1_4_2) {
  $__num = \func_num_args();
  $__res = $functorProduct1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar__unused_6) use ($foldableProduct1_5_3) {
  $__num = \func_num_args();
  $__res = $foldableProduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Traversable_traversableProduct'] = __NAMESPACE__ . '\\majData_majTraversable_traversablemajProduct';

// Data_Traversable_traverseDefault
function majData_majTraversable_traversemajDefault($dictTraversable_0, $dictApplicative_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_traversemajDefault';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $sequence3_2_0 = (($dictTraversable_0)->{'sequence'})($dictApplicative_1);
  $__res = (function() use ($dictTraversable_0, $sequence3_2_0) {
  $__fn = function($f_3, $ta_4 = null) use ($dictTraversable_0, $sequence3_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($sequence3_2_0)(((((($dictTraversable_0)->{'Functor0'})(null))->{'map'})($f_3))($ta_4));
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
$GLOBALS['Data_Traversable_traverseDefault'] = __NAMESPACE__ . '\\majData_majTraversable_traversemajDefault';

// Data_Traversable_mapAccumR
function majData_majTraversable_mapmajAccummajR($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_mapmajAccummajR';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverse2_1_0 = (($dictTraversable_0)->{'traverse'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR']);
  $__res = (function() use ($traverse2_1_0) {
  $__fn = function($f_2, $s0_3 = null, $xs_4 = null) use ($traverse2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverse2_1_0)((function() use ($f_2) {
  $__fn = function($a_5, $s_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_2)($s_6))($a_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($xs_4))($s0_3);
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
$GLOBALS['Data_Traversable_mapAccumR'] = __NAMESPACE__ . '\\majData_majTraversable_mapmajAccummajR';

// Data_Traversable_scanr
function majData_majTraversable_scanr($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_scanr';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mapAccumR1_1_0 = ($GLOBALS['Data_Traversable_mapAccumR'])($dictTraversable_0);
  $__res = (function() use ($mapAccumR1_1_0) {
  $__fn = function($f_2, $b0_3 = null, $xs_4 = null) use ($mapAccumR1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumR1_1_0)((function() use ($f_2) {
  $__fn = function($b_5, $a_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $b_prime_7_1 = (($f_2)($a_6))($b_5);
  $__res = (object)["accum" => $b_prime_7_1, "value" => $b_prime_7_1];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))->{'value'};
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
$GLOBALS['Data_Traversable_scanr'] = __NAMESPACE__ . '\\majData_majTraversable_scanr';

// Data_Traversable_mapAccumL
function majData_majTraversable_mapmajAccummajL($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_mapmajAccummajL';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverse2_1_0 = (($dictTraversable_0)->{'traverse'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL']);
  $__res = (function() use ($traverse2_1_0) {
  $__fn = function($f_2, $s0_3 = null, $xs_4 = null) use ($traverse2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverse2_1_0)((function() use ($f_2) {
  $__fn = function($a_5, $s_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_2)($s_6))($a_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($xs_4))($s0_3);
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
$GLOBALS['Data_Traversable_mapAccumL'] = __NAMESPACE__ . '\\majData_majTraversable_mapmajAccummajL';

// Data_Traversable_scanl
function majData_majTraversable_scanl($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_scanl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mapAccumL1_1_0 = ($GLOBALS['Data_Traversable_mapAccumL'])($dictTraversable_0);
  $__res = (function() use ($mapAccumL1_1_0) {
  $__fn = function($f_2, $b0_3 = null, $xs_4 = null) use ($mapAccumL1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumL1_1_0)((function() use ($f_2) {
  $__fn = function($b_5, $a_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $b_prime_7_1 = (($f_2)($b_5))($a_6);
  $__res = (object)["accum" => $b_prime_7_1, "value" => $b_prime_7_1];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))->{'value'};
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
$GLOBALS['Data_Traversable_scanl'] = __NAMESPACE__ . '\\majData_majTraversable_scanl';

// Data_Traversable_for
function majData_majTraversable_for($dictApplicative_0, $dictTraversable_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_for';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $traverse2_2_0 = (($dictTraversable_1)->{'traverse'})($dictApplicative_0);
  $__res = (function() use ($traverse2_2_0) {
  $__fn = function($x_3, $f_4 = null) use ($traverse2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($traverse2_2_0)($f_4))($x_3);
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
$GLOBALS['Data_Traversable_for'] = __NAMESPACE__ . '\\majData_majTraversable_for';

