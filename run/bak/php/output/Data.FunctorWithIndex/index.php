<?php

namespace Data\FunctorWithIndex;

// ALL IMPORTS: Control.Semigroupoid, Data.Bifunctor, Data.Const, Data.Either, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Tuple, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Semigroupoid, Data.Bifunctor, Data.Const, Data.Either, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.FunctorWithIndex, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Tuple, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Maybe.First/index.php';
require_once __DIR__ . '/../Data.Maybe.Last/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Monoid.Conj/index.php';
require_once __DIR__ . '/../Data.Monoid.Disj/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
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
$ffi_Data_FunctorWithIndex = \call_user_func(function() {
  $exports = [];
$mapWithIndexArray = function($f, $xs = null) use (&$mapWithIndexArray) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$mapWithIndexArray) {
            return $mapWithIndexArray(...\array_merge($__args, $more));
        };
    }
    
    $len = \count($xs);
    $result = array_fill(0, $len, null);
    for ($i = 0; $i < $len; $i++) {
        $f1 = $f($i);
        $result[$i] = $f1($xs[$i]);
    }
    return $result;
};

$exports['mapWithIndexArray'] = $mapWithIndexArray;

return $exports;
  return $exports;
});
$GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'] = $ffi_Data_FunctorWithIndex['mapWithIndexArray'] ?? new class { public function __invoke(...$args) { return $this; } };


// Data_FunctorWithIndex_mapWithIndex
$GLOBALS['Data_FunctorWithIndex_mapWithIndex'] = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['mapWithIndex'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_FunctorWithIndex_mapDefault
$GLOBALS['Data_FunctorWithIndex_mapDefault'] = (function() {
  $__fn = function($dictFunctorWithIndex_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictFunctorWithIndex_0)['mapWithIndex'])(function($v_2 = null) use ($f_1) {
  $__num = \func_num_args();
  $__res = $f_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_FunctorWithIndex_functorWithIndexTuple
$GLOBALS['Data_FunctorWithIndex_functorWithIndexTuple'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($m_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Tuple", ($m_2)->{'value0'}, ($__local_var_1_0)(($m_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Tuple_functorTuple'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexProduct
$GLOBALS['Data_FunctorWithIndex_functorWithIndexProduct'] = function($dictFunctorWithIndex_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictFunctorWithIndex_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $__res = function($dictFunctorWithIndex1_2 = null) use ($__local_var_1_0, $dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictFunctorWithIndex1_2)['Functor0'])($GLOBALS['Prim_undefined']);
  $functorProduct1_4_2 = ["map" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($f_4 = null, $v_5 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($__local_var_1_0)['map'])($f_4))(($v_5)->{'value0'}), ((($__local_var_3_1)['map'])($f_4))(($v_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["mapWithIndex" => (function() use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0) {
  $__fn = function($f_5 = null, $v_6 = null) use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($dictFunctorWithIndex_0)['mapWithIndex'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_5))($GLOBALS['Data_Either_Left'])))(($v_6)->{'value0'}), ((($dictFunctorWithIndex1_2)['mapWithIndex'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_5))($GLOBALS['Data_Either_Right'])))(($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_5 = null) use ($functorProduct1_4_2) {
  $__num = \func_num_args();
  $__res = $functorProduct1_4_2;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_FunctorWithIndex_functorWithIndexMultiplicative
$GLOBALS['Data_FunctorWithIndex_functorWithIndexMultiplicative'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($f_0)($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Multiplicative_functorMultiplicative'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexMaybe
$GLOBALS['Data_FunctorWithIndex_functorWithIndexMaybe'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($v1_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_1_0)(($v1_2)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexLast
$GLOBALS['Data_FunctorWithIndex_functorWithIndexLast'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($v1_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_1_0)(($v1_2)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexIdentity
$GLOBALS['Data_FunctorWithIndex_functorWithIndexIdentity'] = ["mapWithIndex" => (function() {
  $__fn = function($f_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_0)($GLOBALS['Data_Unit_unit']))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Identity_functorIdentity'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexFirst
$GLOBALS['Data_FunctorWithIndex_functorWithIndexFirst'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($v1_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($__local_var_1_0)(($v1_2)->{'value0'}));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Maybe_functorMaybe'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexEither
$GLOBALS['Data_FunctorWithIndex_functorWithIndexEither'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($f_0)($GLOBALS['Data_Unit_unit']);
  $__res = function($m_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($m_2) && (($m_2)->{'tag'} === "Left"))) {
$__t1 = new Phpurs_Data1("Left", ($m_2)->{'value0'});
goto end_branch_1;;
};
  if ((is_object($m_2) && (($m_2)->{'tag'} === "Right"))) {
$__t1 = new Phpurs_Data1("Right", ($__local_var_1_0)(($m_2)->{'value0'}));
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
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Either_functorEither'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexDual
$GLOBALS['Data_FunctorWithIndex_functorWithIndexDual'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($f_0)($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Dual_functorDual'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexDisj
$GLOBALS['Data_FunctorWithIndex_functorWithIndexDisj'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($f_0)($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Disj_functorDisj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexCoproduct
$GLOBALS['Data_FunctorWithIndex_functorWithIndexCoproduct'] = function($dictFunctorWithIndex_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictFunctorWithIndex_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $__res = function($dictFunctorWithIndex1_2 = null) use ($__local_var_1_0, $dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (($dictFunctorWithIndex1_2)['Functor0'])($GLOBALS['Prim_undefined']);
  $functorCoproduct1_4_2 = ["map" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($f_4 = null, $v_5 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_2 = (($__local_var_1_0)['map'])($f_4);
  $__local_var_7_3 = (($__local_var_3_1)['map'])($f_4);
  $__t4 = null;;
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Left"))) {
$__t4 = new Phpurs_Data1("Left", ($__local_var_6_2)(($v_5)->{'value0'}));
goto end_branch_4;;
};
  if ((is_object($v_5) && (($v_5)->{'tag'} === "Right"))) {
$__t4 = new Phpurs_Data1("Right", ($__local_var_7_3)(($v_5)->{'value0'}));
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
})()];
  $__res = ["mapWithIndex" => (function() use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0) {
  $__fn = function($f_5 = null, $v_6 = null) use ($dictFunctorWithIndex1_2, $dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_7_6 = (($dictFunctorWithIndex_0)['mapWithIndex'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_5))($GLOBALS['Data_Either_Left']));
  $__local_var_8_7 = (($dictFunctorWithIndex1_2)['mapWithIndex'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($f_5))($GLOBALS['Data_Either_Right']));
  $__t8 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Left"))) {
$__t8 = new Phpurs_Data1("Left", ($__local_var_7_6)(($v_6)->{'value0'}));
goto end_branch_8;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Right"))) {
$__t8 = new Phpurs_Data1("Right", ($__local_var_8_7)(($v_6)->{'value0'}));
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
})(), "Functor0" => function($dollar__unused_5 = null) use ($functorCoproduct1_4_2) {
  $__num = \func_num_args();
  $__res = $functorCoproduct1_4_2;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_FunctorWithIndex_functorWithIndexConst
$GLOBALS['Data_FunctorWithIndex_functorWithIndexConst'] = ["mapWithIndex" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $v1_1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Const_functorConst'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexConj
$GLOBALS['Data_FunctorWithIndex_functorWithIndexConj'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($f_0)($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Conj_functorConj'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexCompose
$GLOBALS['Data_FunctorWithIndex_functorWithIndexCompose'] = function($dictFunctorWithIndex_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictFunctorWithIndex_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $__res = function($dictFunctorWithIndex1_2 = null) use ($__local_var_1_0, $dictFunctorWithIndex_0) {
  $__num = \func_num_args();
  $mapWithIndex2_3_1 = ($dictFunctorWithIndex1_2)['mapWithIndex'];
  $__local_var_4_2 = (($dictFunctorWithIndex1_2)['Functor0'])($GLOBALS['Prim_undefined']);
  $functorCompose1_5_3 = ["map" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5 = null, $v_6 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['map'])((($__local_var_4_2)['map'])($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = ["mapWithIndex" => (function() use ($dictFunctorWithIndex_0, $mapWithIndex2_3_1) {
  $__fn = function($f_6 = null, $v_7 = null) use ($dictFunctorWithIndex_0, $mapWithIndex2_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctorWithIndex_0)['mapWithIndex'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($mapWithIndex2_3_1))((function() use ($f_6) {
  $__fn = function($a_8 = null, $b_9 = null) use ($f_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($f_6)(new Phpurs_Data2("Tuple", $a_8, $b_9));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))($v_7);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_6 = null) use ($functorCompose1_5_3) {
  $__num = \func_num_args();
  $__res = $functorCompose1_5_3;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_FunctorWithIndex_functorWithIndexArray
$GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'] = ["mapWithIndex" => $GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'], "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Functor_functorArray'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_FunctorWithIndex_functorWithIndexApp
$GLOBALS['Data_FunctorWithIndex_functorWithIndexApp'] = function($dictFunctorWithIndex_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictFunctorWithIndex_0)['Functor0'])($GLOBALS['Prim_undefined']);
  $__res = ["mapWithIndex" => (function() use ($dictFunctorWithIndex_0) {
  $__fn = function($f_2 = null, $v_3 = null) use ($dictFunctorWithIndex_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictFunctorWithIndex_0)['mapWithIndex'])($f_2))($v_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_FunctorWithIndex_functorWithIndexAdditive
$GLOBALS['Data_FunctorWithIndex_functorWithIndexAdditive'] = ["mapWithIndex" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($f_0)($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Monoid_Additive_functorAdditive'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

