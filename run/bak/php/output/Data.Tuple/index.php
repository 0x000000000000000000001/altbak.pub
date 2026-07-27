<?php

namespace Data\Tuple;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.Semigroupoid, Data.BooleanAlgebra, Data.Bounded, Data.CommutativeRing, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.HeytingAlgebra, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Comonad, Control.Extend, Control.Lazy, Control.Monad, Control.Semigroupoid, Data.BooleanAlgebra, Data.Bounded, Data.CommutativeRing, Data.Eq, Data.Function, Data.Functor, Data.Functor.Invariant, Data.Generic.Rep, Data.HeytingAlgebra, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.BooleanAlgebra/index.php';
require_once __DIR__ . '/../Data.Bounded/index.php';
require_once __DIR__ . '/../Data.CommutativeRing/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Invariant/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
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


final class Data_Tuple_Tuple { public $tag = 'Tuple'; public function __construct(public  $value0, public  $value1) {} }

// Data_Tuple_Tuple
$GLOBALS['Data_Tuple_Tuple'] = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple($value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Tuple_uncurry
function majData_majTuple_uncurry($f_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_uncurry';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)(($v_1)->{'value0'}))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_uncurry'] = __NAMESPACE__ . '\\majData_majTuple_uncurry';

// Data_Tuple_swap
function majData_majTuple_swap($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_swap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_0)->{'value1'}, ($v_0)->{'value0'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_swap'] = __NAMESPACE__ . '\\majData_majTuple_swap';

// Data_Tuple_snd
function majData_majTuple_snd($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_snd';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_snd'] = __NAMESPACE__ . '\\majData_majTuple_snd';

// Data_Tuple_showTuple
function majData_majTuple_showmajTuple($dictShow_0, $dictShow1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_showmajTuple';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})("(Tuple "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow_0)->{'show'})(($v_2)->{'value0'})))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})(" "))(((($GLOBALS['Data_Semigroup_semigroupString'])->{'append'})((($dictShow1_1)->{'show'})(($v_2)->{'value1'})))(")"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_showTuple'] = __NAMESPACE__ . '\\majData_majTuple_showmajTuple';

// Data_Tuple_semiringTuple
function majData_majTuple_semiringmajTuple($dictSemiring_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_semiringmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $one_1_0 = ($dictSemiring_0)->{'one'};
  $zero_2_1 = ($dictSemiring_0)->{'zero'};
  $__res = function($dictSemiring1_3) use ($dictSemiring_0, $one_1_0, $zero_2_1) {
  $__num = \func_num_args();
  $__res = (object)["add" => (function() use ($dictSemiring1_3, $dictSemiring_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictSemiring1_3, $dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictSemiring_0)->{'add'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictSemiring1_3)->{'add'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => new \Data\Tuple\Data_Tuple_Tuple($one_1_0, ($dictSemiring1_3)->{'one'}), "mul" => (function() use ($dictSemiring1_3, $dictSemiring_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictSemiring1_3, $dictSemiring_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictSemiring_0)->{'mul'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictSemiring1_3)->{'mul'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "zero" => new \Data\Tuple\Data_Tuple_Tuple($zero_2_1, ($dictSemiring1_3)->{'zero'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_semiringTuple'] = __NAMESPACE__ . '\\majData_majTuple_semiringmajTuple';

// Data_Tuple_semigroupoidTuple
$GLOBALS['Data_Tuple_semigroupoidTuple'] = (object)["compose" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v1_1)->{'value0'}, ($v_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Tuple_semigroupTuple
function majData_majTuple_semigroupmajTuple($dictSemigroup_0, $dictSemigroup1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_semigroupmajTuple';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["append" => (function() use ($dictSemigroup1_1, $dictSemigroup_0) {
  $__fn = function($v_2, $v1_3 = null) use ($dictSemigroup1_1, $dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictSemigroup_0)->{'append'})(($v_2)->{'value0'}))(($v1_3)->{'value0'}), ((($dictSemigroup1_1)->{'append'})(($v_2)->{'value1'}))(($v1_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_semigroupTuple'] = __NAMESPACE__ . '\\majData_majTuple_semigroupmajTuple';

// Data_Tuple_ringTuple
function majData_majTuple_ringmajTuple($dictRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_ringmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictRing_0)->{'Semiring0'})(null);
  $one_2_1 = ($__local_var_1_0)->{'one'};
  $zero_3_2 = ($__local_var_1_0)->{'zero'};
  $semiringTuple1_3_2 = function($dictSemiring1_4) use ($__local_var_1_0, $one_2_1, $zero_3_2) {
  $__num = \func_num_args();
  $__res = (object)["add" => (function() use ($__local_var_1_0, $dictSemiring1_4) {
  $__fn = function($v_5, $v1_6 = null) use ($__local_var_1_0, $dictSemiring1_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'add'})(($v_5)->{'value0'}))(($v1_6)->{'value0'}), ((($dictSemiring1_4)->{'add'})(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "one" => new \Data\Tuple\Data_Tuple_Tuple($one_2_1, ($dictSemiring1_4)->{'one'}), "mul" => (function() use ($__local_var_1_0, $dictSemiring1_4) {
  $__fn = function($v_5, $v1_6 = null) use ($__local_var_1_0, $dictSemiring1_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_1_0)->{'mul'})(($v_5)->{'value0'}))(($v1_6)->{'value0'}), ((($dictSemiring1_4)->{'mul'})(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "zero" => new \Data\Tuple\Data_Tuple_Tuple($zero_3_2, ($dictSemiring1_4)->{'zero'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = function($dictRing1_4) use ($dictRing_0, $semiringTuple1_3_2) {
  $__num = \func_num_args();
  $semiringTuple2_5_4 = ($semiringTuple1_3_2)((($dictRing1_4)->{'Semiring0'})(null));
  $__res = (object)["sub" => (function() use ($dictRing1_4, $dictRing_0) {
  $__fn = function($v_6, $v1_7 = null) use ($dictRing1_4, $dictRing_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictRing_0)->{'sub'})(($v_6)->{'value0'}))(($v1_7)->{'value0'}), ((($dictRing1_4)->{'sub'})(($v_6)->{'value1'}))(($v1_7)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Semiring0" => function($_dollar__unused_6) use ($semiringTuple2_5_4) {
  $__num = \func_num_args();
  $__res = $semiringTuple2_5_4;
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
$GLOBALS['Data_Tuple_ringTuple'] = __NAMESPACE__ . '\\majData_majTuple_ringmajTuple';

// Data_Tuple_monoidTuple
function majData_majTuple_monoidmajTuple($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_monoidmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__local_var_2_1 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = function($dictMonoid1_3) use ($__local_var_2_1, $mempty_1_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictMonoid1_3)->{'Semigroup0'})(null);
  $semigroupTuple2_5_3 = (object)["append" => (function() use ($__local_var_2_1, $__local_var_4_2) {
  $__fn = function($v_5, $v1_6 = null) use ($__local_var_2_1, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($__local_var_2_1)->{'append'})(($v_5)->{'value0'}))(($v1_6)->{'value0'}), ((($__local_var_4_2)->{'append'})(($v_5)->{'value1'}))(($v1_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["mempty" => new \Data\Tuple\Data_Tuple_Tuple($mempty_1_0, ($dictMonoid1_3)->{'mempty'}), "Semigroup0" => function($_dollar__unused_6) use ($semigroupTuple2_5_3) {
  $__num = \func_num_args();
  $__res = $semigroupTuple2_5_3;
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
$GLOBALS['Data_Tuple_monoidTuple'] = __NAMESPACE__ . '\\majData_majTuple_monoidmajTuple';

// Data_Tuple_heytingAlgebraTuple
function majData_majTuple_heytingmajAlgebramajTuple($dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_heytingmajAlgebramajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $tt_1_0 = ($dictHeytingAlgebra_0)->{'tt'};
  $ff_2_1 = ($dictHeytingAlgebra_0)->{'ff'};
  $__res = function($dictHeytingAlgebra1_3) use ($dictHeytingAlgebra_0, $ff_2_1, $tt_1_0) {
  $__num = \func_num_args();
  $__res = (object)["tt" => new \Data\Tuple\Data_Tuple_Tuple($tt_1_0, ($dictHeytingAlgebra1_3)->{'tt'}), "ff" => new \Data\Tuple\Data_Tuple_Tuple($ff_2_1, ($dictHeytingAlgebra1_3)->{'ff'}), "implies" => (function() use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictHeytingAlgebra_0)->{'implies'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictHeytingAlgebra1_3)->{'implies'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "conj" => (function() use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictHeytingAlgebra_0)->{'conj'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictHeytingAlgebra1_3)->{'conj'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "disj" => (function() use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__fn = function($v_4, $v1_5 = null) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictHeytingAlgebra_0)->{'disj'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictHeytingAlgebra1_3)->{'disj'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "not" => function($v_4) use ($dictHeytingAlgebra1_3, $dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($dictHeytingAlgebra_0)->{'not'})(($v_4)->{'value0'}), (($dictHeytingAlgebra1_3)->{'not'})(($v_4)->{'value1'}));
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
$GLOBALS['Data_Tuple_heytingAlgebraTuple'] = __NAMESPACE__ . '\\majData_majTuple_heytingmajAlgebramajTuple';

// Data_Tuple_genericTuple
$GLOBALS['Data_Tuple_genericTuple'] = (object)["to" => function($x_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($x_0)->{'value0'}, ($x_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "from" => function($x_0) {
  $__num = \func_num_args();
  $__res = new \Data\Generic\Rep\Data_Generic_Rep_Product(($x_0)->{'value0'}, ($x_0)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Tuple_functorTuple
$GLOBALS['Data_Tuple_functorTuple'] = (object)["map" => (function() {
  $__fn = function($f_0, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($m_1)->{'value0'}, ($f_0)(($m_1)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Tuple_invariantTuple
$GLOBALS['Data_Tuple_invariantTuple'] = (object)["imap" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Tuple_functorTuple'])->{'map'})($f_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_Tuple_fst
function majData_majTuple_fst($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_fst';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($v_0)->{'value0'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_fst'] = __NAMESPACE__ . '\\majData_majTuple_fst';

// Data_Tuple_lazyTuple
function majData_majTuple_lazymajTuple($dictLazy_0, $dictLazy1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_lazymajTuple';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["defer" => function($f_2) use ($dictLazy1_1, $dictLazy_0) {
  $__num = \func_num_args();
  $__res = new \Data\Tuple\Data_Tuple_Tuple((($dictLazy_0)->{'defer'})(function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)($GLOBALS['Data_Unit_unit']))->{'value0'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($dictLazy1_1)->{'defer'})(function($v_3) use ($f_2) {
  $__num = \func_num_args();
  $__res = (($f_2)($GLOBALS['Data_Unit_unit']))->{'value1'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_lazyTuple'] = __NAMESPACE__ . '\\majData_majTuple_lazymajTuple';

// Data_Tuple_extendTuple
$GLOBALS['Data_Tuple_extendTuple'] = (object)["extend" => (function() {
  $__fn = function($f_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(($v_1)->{'value0'}, ($f_0)($v_1));
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
}];

// Data_Tuple_eqTuple
function majData_majTuple_eqmajTuple($dictEq_0, $dictEq1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_eqmajTuple';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($x_2, $y_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'conj'})(((($dictEq_0)->{'eq'})(($x_2)->{'value0'}))(($y_3)->{'value0'})))(((($dictEq1_1)->{'eq'})(($x_2)->{'value1'}))(($y_3)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Tuple_eqTuple'] = __NAMESPACE__ . '\\majData_majTuple_eqmajTuple';

// Data_Tuple_ordTuple
function majData_majTuple_ordmajTuple($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_ordmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eqTuple1_1_0 = ($GLOBALS['Data_Tuple_eqTuple'])((($dictOrd_0)->{'Eq0'})(null));
  $__res = function($dictOrd1_2) use ($dictOrd_0, $eqTuple1_1_0) {
  $__num = \func_num_args();
  $eqTuple2_3_1 = ($eqTuple1_1_0)((($dictOrd1_2)->{'Eq0'})(null));
  $__res = (object)["compare" => (function() use ($dictOrd1_2, $dictOrd_0) {
  $__fn = function($x_4, $y_5 = null) use ($dictOrd1_2, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_6_2 = ((($dictOrd_0)->{'compare'})(($x_4)->{'value0'}))(($y_5)->{'value0'});
  $__t3 = null;;
  if ($v_6_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t3 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_3;;
};
  if ($v_6_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_3;;
};
  $__t3 = ((($dictOrd1_2)->{'compare'})(($x_4)->{'value1'}))(($y_5)->{'value1'});
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_4) use ($eqTuple2_3_1) {
  $__num = \func_num_args();
  $__res = $eqTuple2_3_1;
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
$GLOBALS['Data_Tuple_ordTuple'] = __NAMESPACE__ . '\\majData_majTuple_ordmajTuple';

// Data_Tuple_eq1Tuple
function majData_majTuple_eq1majTuple($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_eq1majTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["eq1" => function($dictEq1_1) use ($dictEq_0) {
  $__num = \func_num_args();
  $__res = (\Data\Tuple\majData_majTuple_eqmajTuple($dictEq_0, $dictEq1_1))->{'eq'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_eq1Tuple'] = __NAMESPACE__ . '\\majData_majTuple_eq1majTuple';

// Data_Tuple_ord1Tuple
function majData_majTuple_ord1majTuple($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_ord1majTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ordTuple1_1_0 = ($GLOBALS['Data_Tuple_ordTuple'])($dictOrd_0);
  $__local_var_2_1 = (($dictOrd_0)->{'Eq0'})(null);
  $eq1Tuple1_3_2 = (object)["eq1" => function($dictEq1_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (\Data\Tuple\majData_majTuple_eqmajTuple($__local_var_2_1, $dictEq1_3))->{'eq'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare1" => function($dictOrd1_4) use ($ordTuple1_1_0) {
  $__num = \func_num_args();
  $__res = (($ordTuple1_1_0)($dictOrd1_4))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_4) use ($eq1Tuple1_3_2) {
  $__num = \func_num_args();
  $__res = $eq1Tuple1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_ord1Tuple'] = __NAMESPACE__ . '\\majData_majTuple_ord1majTuple';

// Data_Tuple_curry
function majData_majTuple_curry($f_0, $a_1 = null, $b_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_curry';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($f_0)(new \Data\Tuple\Data_Tuple_Tuple($a_1, $b_2));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Tuple_curry'] = __NAMESPACE__ . '\\majData_majTuple_curry';

// Data_Tuple_comonadTuple
$GLOBALS['Data_Tuple_comonadTuple'] = (object)["extract" => $GLOBALS['Data_Tuple_snd'], "Extend0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_extendTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Tuple_commutativeRingTuple
function majData_majTuple_commutativemajRingmajTuple($dictCommutativeRing_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_commutativemajRingmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ringTuple1_1_0 = ($GLOBALS['Data_Tuple_ringTuple'])((($dictCommutativeRing_0)->{'Ring0'})(null));
  $__res = function($dictCommutativeRing1_2) use ($ringTuple1_1_0) {
  $__num = \func_num_args();
  $ringTuple2_3_1 = ($ringTuple1_1_0)((($dictCommutativeRing1_2)->{'Ring0'})(null));
  $__res = (object)["Ring0" => function($_dollar__unused_4) use ($ringTuple2_3_1) {
  $__num = \func_num_args();
  $__res = $ringTuple2_3_1;
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
$GLOBALS['Data_Tuple_commutativeRingTuple'] = __NAMESPACE__ . '\\majData_majTuple_commutativemajRingmajTuple';

// Data_Tuple_boundedTuple
function majData_majTuple_boundedmajTuple($dictBounded_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_boundedmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $top_1_0 = ($dictBounded_0)->{'top'};
  $bottom_2_1 = ($dictBounded_0)->{'bottom'};
  $ordTuple1_3_2 = ($GLOBALS['Data_Tuple_ordTuple'])((($dictBounded_0)->{'Ord0'})(null));
  $__res = function($dictBounded1_4) use ($bottom_2_1, $ordTuple1_3_2, $top_1_0) {
  $__num = \func_num_args();
  $ordTuple2_5_3 = ($ordTuple1_3_2)((($dictBounded1_4)->{'Ord0'})(null));
  $__res = (object)["top" => new \Data\Tuple\Data_Tuple_Tuple($top_1_0, ($dictBounded1_4)->{'top'}), "bottom" => new \Data\Tuple\Data_Tuple_Tuple($bottom_2_1, ($dictBounded1_4)->{'bottom'}), "Ord0" => function($_dollar__unused_6) use ($ordTuple2_5_3) {
  $__num = \func_num_args();
  $__res = $ordTuple2_5_3;
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
$GLOBALS['Data_Tuple_boundedTuple'] = __NAMESPACE__ . '\\majData_majTuple_boundedmajTuple';

// Data_Tuple_booleanAlgebraTuple
function majData_majTuple_booleanmajAlgebramajTuple($dictBooleanAlgebra_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_booleanmajAlgebramajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $heytingAlgebraTuple1_1_0 = ($GLOBALS['Data_Tuple_heytingAlgebraTuple'])((($dictBooleanAlgebra_0)->{'HeytingAlgebra0'})(null));
  $__res = function($dictBooleanAlgebra1_2) use ($heytingAlgebraTuple1_1_0) {
  $__num = \func_num_args();
  $heytingAlgebraTuple2_3_1 = ($heytingAlgebraTuple1_1_0)((($dictBooleanAlgebra1_2)->{'HeytingAlgebra0'})(null));
  $__res = (object)["HeytingAlgebra0" => function($_dollar__unused_4) use ($heytingAlgebraTuple2_3_1) {
  $__num = \func_num_args();
  $__res = $heytingAlgebraTuple2_3_1;
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
$GLOBALS['Data_Tuple_booleanAlgebraTuple'] = __NAMESPACE__ . '\\majData_majTuple_booleanmajAlgebramajTuple';

// Data_Tuple_applyTuple
function majData_majTuple_applymajTuple($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_applymajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["apply" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_1, $v1_2 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictSemigroup_0)->{'append'})(($v_1)->{'value0'}))(($v1_2)->{'value0'}), (($v_1)->{'value1'})(($v1_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_functorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_applyTuple'] = __NAMESPACE__ . '\\majData_majTuple_applymajTuple';

// Data_Tuple_bindTuple
function majData_majTuple_bindmajTuple($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_bindmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applyTuple1_1_0 = \Data\Tuple\majData_majTuple_applymajTuple($dictSemigroup_0);
  $__res = (object)["bind" => (function() use ($dictSemigroup_0) {
  $__fn = function($v_2, $f_3 = null) use ($dictSemigroup_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v1_4_1 = ($f_3)(($v_2)->{'value1'});
  $__res = new \Data\Tuple\Data_Tuple_Tuple(((($dictSemigroup_0)->{'append'})(($v_2)->{'value0'}))(($v1_4_1)->{'value0'}), ($v1_4_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($_dollar__unused_2) use ($applyTuple1_1_0) {
  $__num = \func_num_args();
  $__res = $applyTuple1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_bindTuple'] = __NAMESPACE__ . '\\majData_majTuple_bindmajTuple';

// Data_Tuple_applicativeTuple
function majData_majTuple_applicativemajTuple($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_applicativemajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applyTuple1_1_0 = \Data\Tuple\majData_majTuple_applymajTuple((($dictMonoid_0)->{'Semigroup0'})(null));
  $__res = (object)["pure" => ($GLOBALS['Data_Tuple_Tuple'])(($dictMonoid_0)->{'mempty'}), "Apply0" => function($_dollar__unused_2) use ($applyTuple1_1_0) {
  $__num = \func_num_args();
  $__res = $applyTuple1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_applicativeTuple'] = __NAMESPACE__ . '\\majData_majTuple_applicativemajTuple';

// Data_Tuple_monadTuple
function majData_majTuple_monadmajTuple($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majTuple_monadmajTuple';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $applicativeTuple1_1_0 = \Data\Tuple\majData_majTuple_applicativemajTuple($dictMonoid_0);
  $bindTuple1_2_1 = \Data\Tuple\majData_majTuple_bindmajTuple((($dictMonoid_0)->{'Semigroup0'})(null));
  $__res = (object)["Applicative0" => function($_dollar__unused_3) use ($applicativeTuple1_1_0) {
  $__num = \func_num_args();
  $__res = $applicativeTuple1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_3) use ($bindTuple1_2_1) {
  $__num = \func_num_args();
  $__res = $bindTuple1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Tuple_monadTuple'] = __NAMESPACE__ . '\\majData_majTuple_monadmajTuple';

