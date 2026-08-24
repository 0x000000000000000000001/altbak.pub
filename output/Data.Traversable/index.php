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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Traversable = \call_user_func(function() {
  $exports = [];
$traverseArrayImpl = function($apply, $map, $pure, $f, $array) use (&$traverseArrayImpl) {

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
                return $map($array1, $f1);
            case 2:
                $f1 = $f($array[$bot]);
                $f2 = $f($array[$bot + 1]);
                $map2 = $map($array2, $f1);
                return $apply($map2, $f2);
            case 3:
                $f1 = $f($array[$bot]);
                $f2 = $f($array[$bot + 1]);
                $f3 = $f($array[$bot + 2]);
                $map2 = $map($array3, $f1);
                $app1 = $apply($map2, $f2);
                return $apply($app1, $f3);
            default:
                $pivot = $bot + floor(($top - $bot) / 4) * 2;
                $go1 = $go($bot, $pivot);
                $go2 = $go($pivot, $top);
                $map2 = $map($concat2, $go1);
                return $apply($map2, $go2);
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
  $f = (\array_key_exists('traverseArrayImpl', $ffi_Data_Traversable) ? $ffi_Data_Traversable['traverseArrayImpl'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Traversable_traverseArrayImpl'] = __NAMESPACE__ . '\\majData_majTraversable_traversemajArraymajImpl';





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
$GLOBALS['Data_Traversable_traversableTuple'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(($GLOBALS['Data_Tuple_Tuple'])(($v_3)->{'value0'})))(($f_2)(($v_3)->{'value1'}));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_1)->{'map'})(($GLOBALS['Data_Tuple_Tuple'])(($v_2)->{'value0'})))(($v_2)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_functorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableMultiplicative
$GLOBALS['Data_Traversable_traversableMultiplicative'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_1)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Multiplicative_functorMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableMaybe
$GLOBALS['Data_Traversable_traversableMaybe'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($Functor0_1_0, $dictApplicative_0, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ((($Functor0_1_0)->{'map'})($GLOBALS['Data_Maybe_Just']))(($v_2)(($v1_3)->{'value0'}));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_2 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_2, $dictApplicative_0) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_3;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = ((($Functor0_1_2)->{'map'})($GLOBALS['Data_Maybe_Just']))(($v_2)->{'value0'});
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
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableIdentity
$GLOBALS['Data_Traversable_traversableIdentity'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_1)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Identity_functorIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableEither
$GLOBALS['Data_Traversable_traversableEither'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($Functor0_1_0, $dictApplicative_0, $v_2) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v1_3 instanceof \Data\Either\Data_Either_Left) {
$__t1 = (($dictApplicative_0)->{'pure'})(new \Data\Either\Data_Either_Left(($v1_3)->{'value0'}));
goto end_branch_1;;
};
  if ($v1_3 instanceof \Data\Either\Data_Either_Right) {
$__t1 = ((($Functor0_1_0)->{'map'})($GLOBALS['Data_Either_Right']))(($v_2)(($v1_3)->{'value0'}));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_2 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_2, $dictApplicative_0) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ($v_2 instanceof \Data\Either\Data_Either_Left) {
$__t3 = (($dictApplicative_0)->{'pure'})(new \Data\Either\Data_Either_Left(($v_2)->{'value0'}));
goto end_branch_3;;
};
  if ($v_2 instanceof \Data\Either\Data_Either_Right) {
$__t3 = ((($Functor0_1_2)->{'map'})($GLOBALS['Data_Either_Right']))(($v_2)->{'value0'});
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
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableDual
$GLOBALS['Data_Traversable_traversableDual'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_1)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Dual_functorDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableDisj
$GLOBALS['Data_Traversable_traversableDisj'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_1)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Disj_functorDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableConst
$GLOBALS['Data_Traversable_traversableConst'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v1_2) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)->{'pure'})($v1_2);
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_1) use ($dictApplicative_0) {
  $__num = \func_num_args();
  $__res = (($dictApplicative_0)->{'pure'})($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Const_functorConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableConj
$GLOBALS['Data_Traversable_traversableConj'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_1)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Conj_functorConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
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
  $functorCompose1_4_2 = (object)["map" => function($f_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_1_0, $__local_var_4_2, $f_5) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})((($__local_var_4_2)->{'map'})($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_4 = (($dictTraversable1_3)->{'Foldable1'})(null);
  $foldableCompose1_5_4 = (object)["foldr" => function($f_6) use ($__local_var_2_1, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($i_7) use ($__local_var_2_1, $__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_2_1, $__local_var_5_4, $f_6, $i_7) {
  $__num = \func_num_args();
  $__local_var_9_5 = (($__local_var_5_4)->{'foldr'})($f_6);
  $__res = (((($__local_var_2_1)->{'foldr'})(function($b_10) use ($__local_var_9_5) {
  $__num = \func_num_args();
  $__res = function($a_11) use ($__local_var_9_5, $b_10) {
  $__num = \func_num_args();
  $__res = (($__local_var_9_5)($a_11))($b_10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($i_7))($v_8);
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
}, "foldl" => function($f_6) use ($__local_var_2_1, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($i_7) use ($__local_var_2_1, $__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_2_1, $__local_var_5_4, $f_6, $i_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_1)->{'foldl'})((($__local_var_5_4)->{'foldl'})($f_6)))($i_7))($v_8);
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
}, "foldMap" => function($dictMonoid_6) use ($__local_var_2_1, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($__local_var_2_1, $__local_var_5_4, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_2_1, $__local_var_5_4, $dictMonoid_6, $f_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_1)->{'foldMap'})($dictMonoid_6))(((($__local_var_5_4)->{'foldMap'})($dictMonoid_6))($f_7)))($v_8);
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
  $__res = (object)["traverse" => function($dictApplicative_6) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $Functor0_7_7 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_8) use ($Functor0_7_7, $dictApplicative_6, $dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($Functor0_7_7, $dictApplicative_6, $dictTraversable1_3, $dictTraversable_0, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Functor0_7_7)->{'map'})(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((((($dictTraversable_0)->{'traverse'})($dictApplicative_6))(((($dictTraversable1_3)->{'traverse'})($dictApplicative_6))($f_8)))($v_9));
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
  $__res = ((((($GLOBALS['Data_Traversable_traversableCompose'])($dictTraversable_0))($dictTraversable1_3))->{'traverse'})($dictApplicative_6))(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorCompose1_4_2) {
  $__num = \func_num_args();
  $__res = $functorCompose1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_6) use ($foldableCompose1_5_4) {
  $__num = \func_num_args();
  $__res = $foldableCompose1_5_4;
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
$GLOBALS['Data_Traversable_traversableAdditive'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $f_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_2)($v_3));
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_1) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_1)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Additive_functorAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
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
  $__res = ((($dictTraversable_0)->{'traverse'})($dictApplicative_1))(function($x_2) {
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
  $__res = ((($GLOBALS['Data_Traversable_traversableArray'])->{'traverse'})($dictApplicative_0))(function($x_1) {
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
  $__res = $GLOBALS['Data_Functor_functorArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
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
  $functorApp_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $__local_var_2_1 = (($dictTraversable_0)->{'Foldable1'})(null);
  $foldableApp_2_1 = (object)["foldr" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($i_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_1, $f_3, $i_4) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_1)->{'foldr'})($f_3))($i_4))($v_5);
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
}, "foldl" => function($f_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($i_4) use ($__local_var_2_1, $f_3) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_1, $f_3, $i_4) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_1)->{'foldl'})($f_3))($i_4))($v_5);
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
}, "foldMap" => function($dictMonoid_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = function($f_4) use ($__local_var_2_1, $dictMonoid_3) {
  $__num = \func_num_args();
  $__res = function($v_5) use ($__local_var_2_1, $dictMonoid_3, $f_4) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_1)->{'foldMap'})($dictMonoid_3))($f_4))($v_5);
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
  $__res = (object)["traverse" => function($dictApplicative_3) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $Functor0_4_3 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_5) use ($Functor0_4_3, $dictApplicative_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($Functor0_4_3, $dictApplicative_3, $dictTraversable_0, $f_5) {
  $__num = \func_num_args();
  $__res = ((($Functor0_4_3)->{'map'})(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((((($dictTraversable_0)->{'traverse'})($dictApplicative_3))($f_5))($v_6));
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
}, "sequence" => function($dictApplicative_3) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $Functor0_4_4 = (((($dictApplicative_3)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_5) use ($Functor0_4_4, $dictApplicative_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = ((($Functor0_4_4)->{'map'})(function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(((($dictTraversable_0)->{'sequence'})($dictApplicative_3))($v_5));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_3) use ($functorApp_1_0) {
  $__num = \func_num_args();
  $__res = $functorApp_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_3) use ($foldableApp_2_1) {
  $__num = \func_num_args();
  $__res = $foldableApp_2_1;
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
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $__local_var_2_1 = (($dictTraversable_0)->{'Foldable1'})(null);
  $__res = function($dictTraversable1_3) use ($__local_var_1_0, $__local_var_2_1, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictTraversable1_3)->{'Functor0'})(null);
  $functorCoproduct1_4_2 = (object)["map" => function($f_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_1_0, $__local_var_4_2, $f_5) {
  $__num = \func_num_args();
  $__local_var_7_3 = (($__local_var_1_0)->{'map'})($f_5);
  $__local_var_8_4 = (($__local_var_4_2)->{'map'})($f_5);
  $__t5 = null;;
  if ($v_6 instanceof \Data\Either\Data_Either_Left) {
$__t5 = new \Data\Either\Data_Either_Left(($__local_var_7_3)(($v_6)->{'value0'}));
goto end_branch_5;;
};
  if ($v_6 instanceof \Data\Either\Data_Either_Right) {
$__t5 = new \Data\Either\Data_Either_Right(($__local_var_8_4)(($v_6)->{'value0'}));
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
}];
  $__local_var_5_7 = (($dictTraversable1_3)->{'Foldable1'})(null);
  $foldableCoproduct1_5_7 = (object)["foldr" => function($f_6) use ($__local_var_2_1, $__local_var_5_7) {
  $__num = \func_num_args();
  $__res = function($z_7) use ($__local_var_2_1, $__local_var_5_7, $f_6) {
  $__num = \func_num_args();
  $__local_var_8_8 = ((($__local_var_2_1)->{'foldr'})($f_6))($z_7);
  $__local_var_9_9 = ((($__local_var_5_7)->{'foldr'})($f_6))($z_7);
  $__res = function($v2_10) use ($__local_var_8_8, $__local_var_9_9) {
  $__num = \func_num_args();
  $__t10 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t10 = ($__local_var_8_8)(($v2_10)->{'value0'});
goto end_branch_10;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t10 = ($__local_var_9_9)(($v2_10)->{'value0'});
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
}, "foldl" => function($f_6) use ($__local_var_2_1, $__local_var_5_7) {
  $__num = \func_num_args();
  $__res = function($z_7) use ($__local_var_2_1, $__local_var_5_7, $f_6) {
  $__num = \func_num_args();
  $__local_var_8_11 = ((($__local_var_2_1)->{'foldl'})($f_6))($z_7);
  $__local_var_9_12 = ((($__local_var_5_7)->{'foldl'})($f_6))($z_7);
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "foldMap" => function($dictMonoid_6) use ($__local_var_2_1, $__local_var_5_7) {
  $__num = \func_num_args();
  $__res = function($f_7) use ($__local_var_2_1, $__local_var_5_7, $dictMonoid_6) {
  $__num = \func_num_args();
  $__local_var_8_14 = ((($__local_var_2_1)->{'foldMap'})($dictMonoid_6))($f_7);
  $__local_var_9_15 = ((($__local_var_5_7)->{'foldMap'})($dictMonoid_6))($f_7);
  $__res = function($v2_10) use ($__local_var_8_14, $__local_var_9_15) {
  $__num = \func_num_args();
  $__t16 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t16 = ($__local_var_8_14)(($v2_10)->{'value0'});
goto end_branch_16;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t16 = ($__local_var_9_15)(($v2_10)->{'value0'});
goto end_branch_16;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t16 = null;
  end_branch_16:;
  $__res = $__t16;
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
  $__res = (object)["traverse" => function($dictApplicative_6) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $Functor0_7_18 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_8) use ($Functor0_7_18, $dictApplicative_6, $dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_9_19 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_7_18)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Left']))))(((($dictTraversable_0)->{'traverse'})($dictApplicative_6))($f_8));
  $__local_var_10_20 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_7_18)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_10) {
  $__num = \func_num_args();
  $__res = $x_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Right']))))(((($dictTraversable1_3)->{'traverse'})($dictApplicative_6))($f_8));
  $__res = function($v2_11) use ($__local_var_10_20, $__local_var_9_19) {
  $__num = \func_num_args();
  $__t21 = null;;
  if ($v2_11 instanceof \Data\Either\Data_Either_Left) {
$__t21 = ($__local_var_9_19)(($v2_11)->{'value0'});
goto end_branch_21;;
};
  if ($v2_11 instanceof \Data\Either\Data_Either_Right) {
$__t21 = ($__local_var_10_20)(($v2_11)->{'value0'});
goto end_branch_21;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t21 = null;
  end_branch_21:;
  $__res = $__t21;
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
  $Functor0_7_22 = (((($dictApplicative_6)->{'Apply0'})(null))->{'Functor0'})(null);
  $__local_var_8_23 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_7_22)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Left']))))((($dictTraversable_0)->{'sequence'})($dictApplicative_6));
  $__local_var_9_24 = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_7_22)->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_9) {
  $__num = \func_num_args();
  $__res = $x_9;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Either_Right']))))((($dictTraversable1_3)->{'sequence'})($dictApplicative_6));
  $__res = function($v2_10) use ($__local_var_8_23, $__local_var_9_24) {
  $__num = \func_num_args();
  $__t25 = null;;
  if ($v2_10 instanceof \Data\Either\Data_Either_Left) {
$__t25 = ($__local_var_8_23)(($v2_10)->{'value0'});
goto end_branch_25;;
};
  if ($v2_10 instanceof \Data\Either\Data_Either_Right) {
$__t25 = ($__local_var_9_24)(($v2_10)->{'value0'});
goto end_branch_25;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t25 = null;
  end_branch_25:;
  $__res = $__t25;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorCoproduct1_4_2) {
  $__num = \func_num_args();
  $__res = $functorCoproduct1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_6) use ($foldableCoproduct1_5_7) {
  $__num = \func_num_args();
  $__res = $foldableCoproduct1_5_7;
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
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $Functor0_4_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t2 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_2;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ((($Functor0_4_1)->{'map'})($GLOBALS['Data_Maybe_Just']))(($f_2)(($v_3)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t2);
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_3 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_3, $dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_3_4 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t5 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ((($Functor0_3_4)->{'map'})($GLOBALS['Data_Maybe_Just']))(($v_2)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = ((($Functor0_1_3)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Foldable_foldableFirst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Traversable_traversableLast
$GLOBALS['Data_Traversable_traversableLast'] = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($f_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($Functor0_1_0, $dictApplicative_0, $f_2) {
  $__num = \func_num_args();
  $Functor0_4_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t2 = null;;
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_2;;
};
  if ($v_3 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ((($Functor0_4_1)->{'map'})($GLOBALS['Data_Maybe_Just']))(($f_2)(($v_3)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = ((($Functor0_1_0)->{'map'})(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t2);
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
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_1_3 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($v_2) use ($Functor0_1_3, $dictApplicative_0) {
  $__num = \func_num_args();
  $Functor0_3_4 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__t5 = null;;
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t5 = (($dictApplicative_0)->{'pure'})(new \Data\Maybe\Data_Maybe_Nothing());
goto end_branch_5;;
};
  if ($v_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t5 = ((($Functor0_3_4)->{'map'})($GLOBALS['Data_Maybe_Just']))(($v_2)->{'value0'});
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = ((($Functor0_1_3)->{'map'})(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($__t5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_0) {
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
  $__local_var_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $__local_var_2_1 = (($dictTraversable_0)->{'Foldable1'})(null);
  $__res = function($dictTraversable1_3) use ($__local_var_1_0, $__local_var_2_1, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictTraversable1_3)->{'Functor0'})(null);
  $functorProduct1_4_2 = (object)["map" => function($f_5) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  $__res = function($v_6) use ($__local_var_1_0, $__local_var_4_2, $f_5) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'map'})($f_5))(($v_6)->{'value0'}), ((($__local_var_4_2)->{'map'})($f_5))(($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_5_4 = (($dictTraversable1_3)->{'Foldable1'})(null);
  $foldableProduct1_5_4 = (object)["foldr" => function($f_6) use ($__local_var_2_1, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($z_7) use ($__local_var_2_1, $__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_2_1, $__local_var_5_4, $f_6, $z_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_2_1)->{'foldr'})($f_6))((((($__local_var_5_4)->{'foldr'})($f_6))($z_7))(($v_8)->{'value1'})))(($v_8)->{'value0'});
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
}, "foldl" => function($f_6) use ($__local_var_2_1, $__local_var_5_4) {
  $__num = \func_num_args();
  $__res = function($z_7) use ($__local_var_2_1, $__local_var_5_4, $f_6) {
  $__num = \func_num_args();
  $__res = function($v_8) use ($__local_var_2_1, $__local_var_5_4, $f_6, $z_7) {
  $__num = \func_num_args();
  $__res = (((($__local_var_5_4)->{'foldl'})($f_6))((((($__local_var_2_1)->{'foldl'})($f_6))($z_7))(($v_8)->{'value0'})))(($v_8)->{'value1'});
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
}, "foldMap" => function($dictMonoid_6) use ($__local_var_2_1, $__local_var_5_4) {
  $__num = \func_num_args();
  $Semigroup0_7_5 = (($dictMonoid_6)->{'Semigroup0'})(null);
  $__res = function($f_8) use ($Semigroup0_7_5, $__local_var_2_1, $__local_var_5_4, $dictMonoid_6) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($Semigroup0_7_5, $__local_var_2_1, $__local_var_5_4, $dictMonoid_6, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_7_5)->{'append'})((((($__local_var_2_1)->{'foldMap'})($dictMonoid_6))($f_8))(($v_9)->{'value0'})))((((($__local_var_5_4)->{'foldMap'})($dictMonoid_6))($f_8))(($v_9)->{'value1'}));
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
  $__res = (object)["traverse" => function($dictApplicative_6) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $Apply0_7_7 = (($dictApplicative_6)->{'Apply0'})(null);
  $__res = function($f_8) use ($Apply0_7_7, $dictApplicative_6, $dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = function($v_9) use ($Apply0_7_7, $dictApplicative_6, $dictTraversable1_3, $dictTraversable_0, $f_8) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_7)->{'apply'})(((((($Apply0_7_7)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))((((($dictTraversable_0)->{'traverse'})($dictApplicative_6))($f_8))(($v_9)->{'value0'}))))((((($dictTraversable1_3)->{'traverse'})($dictApplicative_6))($f_8))(($v_9)->{'value1'}));
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
  $Apply0_7_8 = (($dictApplicative_6)->{'Apply0'})(null);
  $__res = function($v_8) use ($Apply0_7_8, $dictApplicative_6, $dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = ((($Apply0_7_8)->{'apply'})(((((($Apply0_7_8)->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Functor_Product_product']))(((($dictTraversable_0)->{'sequence'})($dictApplicative_6))(($v_8)->{'value0'}))))(((($dictTraversable1_3)->{'sequence'})($dictApplicative_6))(($v_8)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_6) use ($functorProduct1_4_2) {
  $__num = \func_num_args();
  $__res = $functorProduct1_4_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($_dollar___unused_6) use ($foldableProduct1_5_4) {
  $__num = \func_num_args();
  $__res = $foldableProduct1_5_4;
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
function majData_majTraversable_traversemajDefault($dictTraversable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_traversemajDefault';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (($dictTraversable_0)->{'Functor0'})(null);
  $__res = function($dictApplicative_2) use ($Functor0_1_0, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = function($f_3) use ($Functor0_1_0, $dictApplicative_2, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = function($ta_4) use ($Functor0_1_0, $dictApplicative_2, $dictTraversable_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((($dictTraversable_0)->{'sequence'})($dictApplicative_2))(((($Functor0_1_0)->{'map'})($f_3))($ta_4));
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
$GLOBALS['Data_Traversable_traverseDefault'] = __NAMESPACE__ . '\\majData_majTraversable_traversemajDefault';

// Data_Traversable_mapAccumR
function majData_majTraversable_mapmajAccummajR($dictTraversable_0, $f_1 = null, $s0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_mapmajAccummajR';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictTraversable_0)->{'traverse'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR']))(function($a_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)($s_5))($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3))($s0_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Traversable_mapAccumR'] = __NAMESPACE__ . '\\majData_majTraversable_mapmajAccummajR';

// Data_Traversable_scanr
function majData_majTraversable_scanr($dictTraversable_0, $f_1 = null, $b0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_scanr';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((((($dictTraversable_0)->{'traverse'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR']))(function($a_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $f_1) {
  $__num = \func_num_args();
  $b_prime__6_0 = (($f_1)($a_4))($s_5);
  $__res = (object)["accum" => $b_prime__6_0, "value" => $b_prime__6_0];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3))($b0_2))->{'value'};
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Traversable_scanr'] = __NAMESPACE__ . '\\majData_majTraversable_scanr';

// Data_Traversable_mapAccumL
function majData_majTraversable_mapmajAccummajL($dictTraversable_0, $f_1 = null, $s0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_mapmajAccummajL';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((((($dictTraversable_0)->{'traverse'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL']))(function($a_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)($s_5))($a_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3))($s0_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Traversable_mapAccumL'] = __NAMESPACE__ . '\\majData_majTraversable_mapmajAccummajL';

// Data_Traversable_scanl
function majData_majTraversable_scanl($dictTraversable_0, $f_1 = null, $b0_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_scanl';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((((($dictTraversable_0)->{'traverse'})($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL']))(function($a_4) use ($f_1) {
  $__num = \func_num_args();
  $__res = function($s_5) use ($a_4, $f_1) {
  $__num = \func_num_args();
  $b_prime__6_0 = (($f_1)($s_5))($a_4);
  $__res = (object)["accum" => $b_prime__6_0, "value" => $b_prime__6_0];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3))($b0_2))->{'value'};
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Traversable_scanl'] = __NAMESPACE__ . '\\majData_majTraversable_scanl';

// Data_Traversable_for
function majData_majTraversable_for($dictApplicative_0, $dictTraversable_1 = null, $x_2 = null, $f_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTraversable_for';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (((($dictTraversable_1)->{'traverse'})($dictApplicative_0))($f_3))($x_2);
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Traversable_for'] = __NAMESPACE__ . '\\majData_majTraversable_for';

