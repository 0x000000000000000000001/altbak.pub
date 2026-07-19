<?php

namespace Data\Functor\Coproduct;

// ALL IMPORTS: Control.Comonad, Control.Extend, Control.Semigroupoid, Data.Bifunctor, Data.Either, Data.Eq, Data.Functor, Data.Functor.Coproduct, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Prelude, Prim
// TO REQUIRE: Control.Comonad, Control.Extend, Control.Semigroupoid, Data.Bifunctor, Data.Either, Data.Eq, Data.Functor, Data.Functor.Coproduct, Data.Newtype, Data.Ord, Data.Ordering, Data.Semigroup, Data.Show, Prelude
require_once __DIR__ . '/../Control.Comonad/index.php';
require_once __DIR__ . '/../Control.Extend/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Prelude/index.php';

if (!class_exists(__NAMESPACE__ . '\\Phpurs_Data0')) {
  class Phpurs_Data0 { public $tag; public function __construct($t) { $this->tag = $t; } }
  class Phpurs_Data1 { public $tag; public $value0; public function __construct($t, $value0) { $this->tag = $t; $this->value0 = $value0; } }
  class Phpurs_Data2 { public $tag; public $value0, $value1; public function __construct($t, $value0, $value1) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; } }
  class Phpurs_Data3 { public $tag; public $value0, $value1, $value2; public function __construct($t, $value0, $value1, $value2) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; } }
  class Phpurs_Data4 { public $tag; public $value0, $value1, $value2, $value3; public function __construct($t, $value0, $value1, $value2, $value3) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; } }
  class Phpurs_Data5 { public $tag; public $value0, $value1, $value2, $value3, $value4; public function __construct($t, $value0, $value1, $value2, $value3, $value4) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; } }
  class Phpurs_Data6 { public $tag; public $value0, $value1, $value2, $value3, $value4, $value5; public function __construct($t, $value0, $value1, $value2, $value3, $value4, $value5) { $this->tag = $t; $this->value0 = $value0; $this->value1 = $value1; $this->value2 = $value2; $this->value3 = $value3; $this->value4 = $value4; $this->value5 = $value5; } }
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
\PhpursThunks::$thunks['Data_Functor_Coproduct_Coproduct'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_Coproduct"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_showCoproduct'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_showCoproduct"), recVars=[];
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Left"))) {
$__t0 = (("(left " . (($dictShow_0)->show)(($v_2)->value0)) . ")");
} else {
if ((is_object($v_2) && (($v_2)->tag === "Right"))) {
$__t0 = (("(right " . (($dictShow1_1)->show)(($v_2)->value0)) . ")");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_right'] = function() { $v = function($ga_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_right"), recVars=[];
  $__res = new Phpurs_Data1("Right", $ga_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_newtypeCoproduct'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_left'] = function() { $v = function($fa_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_left"), recVars=[];
  $__res = new Phpurs_Data1("Left", $fa_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_functorCoproduct'] = function() { $v = (function() {
  $__fn = function($dictFunctor_0, $dictFunctor1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_functorCoproduct"), recVars=[];
  $__res = (object)["map" => (function() use ($dictFunctor1_1, $dictFunctor_0) {
  $__fn = function($f_2, $v_3 = null) use ($dictFunctor1_1, $dictFunctor_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_0 = (($dictFunctor_0)->map)($f_2);
  $__local_var_5_1 = (($dictFunctor1_1)->map)($f_2);
  if ((is_object($v_3) && (($v_3)->tag === "Left"))) {
$__t2 = new Phpurs_Data1("Left", ($__local_var_4_0)(($v_3)->value0));
} else {
if ((is_object($v_3) && (($v_3)->tag === "Right"))) {
$__t2 = new Phpurs_Data1("Right", ($__local_var_5_1)(($v_3)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_eq1Coproduct'] = function() { $v = (function() {
  $__fn = function($dictEq1_0, $dictEq11_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_eq1Coproduct"), recVars=[];
  $__res = (object)["eq1" => function($dictEq_2) use ($dictEq11_1, $dictEq1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eq12_3_0 = (($dictEq1_0)->eq1)($dictEq_2);
  $eq13_4_1 = (($dictEq11_1)->eq1)($dictEq_2);
  $__res = (function() use ($eq12_3_0, $eq13_4_1) {
  $__fn = function($v_5, $v1_6 = null) use ($eq12_3_0, $eq13_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_5) && (($v_5)->tag === "Left"))) {
$__t2 = ((is_object($v1_6) && (($v1_6)->tag === "Left")) && (($eq12_3_0)(($v_5)->value0))(($v1_6)->value0));
} else {
$__t2 = ((is_object($v_5) && (($v_5)->tag === "Right")) && ((is_object($v1_6) && (($v1_6)->tag === "Right")) && (($eq13_4_1)(($v_5)->value0))(($v1_6)->value0)));
};
  $__res = $__t2;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_eqCoproduct'] = function() { $v = (function() {
  $__fn = function($dictEq1_0, $dictEq11_1 = null, $dictEq_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_eqCoproduct"), recVars=[];
  $eq12_3_0 = (($dictEq1_0)->eq1)($dictEq_2);
  $eq13_4_1 = (($dictEq11_1)->eq1)($dictEq_2);
  $__res = (object)["eq" => (function() use ($eq12_3_0, $eq13_4_1) {
  $__fn = function($v_5, $v1_6 = null) use ($eq12_3_0, $eq13_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_5) && (($v_5)->tag === "Left"))) {
$__t2 = ((is_object($v1_6) && (($v1_6)->tag === "Left")) && (($eq12_3_0)(($v_5)->value0))(($v1_6)->value0));
} else {
$__t2 = ((is_object($v_5) && (($v_5)->tag === "Right")) && ((is_object($v1_6) && (($v1_6)->tag === "Right")) && (($eq13_4_1)(($v_5)->value0))(($v1_6)->value0)));
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_ord1Coproduct'] = function() { $v = function($dictOrd1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_ord1Coproduct"), recVars=[];
  $__local_var_1_0 = (($dictOrd1_0)->Eq10)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($dictOrd11_2) use ($__local_var_1_0, $dictOrd1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_3_1 = (($dictOrd11_2)->Eq10)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eq1Coproduct2_4_2 = (object)["eq1" => function($dictEq_4) use ($__local_var_1_0, $__local_var_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eq12_5_2 = (($__local_var_1_0)->eq1)($dictEq_4);
  $eq13_6_3 = (($__local_var_3_1)->eq1)($dictEq_4);
  $__res = (function() use ($eq12_5_2, $eq13_6_3) {
  $__fn = function($v_7, $v1_8 = null) use ($eq12_5_2, $eq13_6_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_7) && (($v_7)->tag === "Left"))) {
$__t4 = ((is_object($v1_8) && (($v1_8)->tag === "Left")) && (($eq12_5_2)(($v_7)->value0))(($v1_8)->value0));
} else {
$__t4 = ((is_object($v_7) && (($v_7)->tag === "Right")) && ((is_object($v1_8) && (($v1_8)->tag === "Right")) && (($eq13_6_3)(($v_7)->value0))(($v1_8)->value0)));
};
  $__res = $__t4;
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
  $__res = (object)["compare1" => function($dictOrd_5) use ($dictOrd11_2, $dictOrd1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $compare12_6_6 = (($dictOrd1_0)->compare1)($dictOrd_5);
  $compare13_7_7 = (($dictOrd11_2)->compare1)($dictOrd_5);
  $__res = (function() use ($compare12_6_6, $compare13_7_7) {
  $__fn = function($v_8, $v1_9 = null) use ($compare12_6_6, $compare13_7_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_8) && (($v_8)->tag === "Left"))) {
if ((is_object($v1_9) && (($v1_9)->tag === "Left"))) {
$__t9 = (($compare12_6_6)(($v_8)->value0))(($v1_9)->value0);
} else {
$__t9 = new Phpurs_Data0("LT");
};
$__t8 = $__t9;
} else {
if ((is_object($v1_9) && (($v1_9)->tag === "Left"))) {
$__t8 = new Phpurs_Data0("GT");
} else {
if (((is_object($v_8) && (($v_8)->tag === "Right")) && (is_object($v1_9) && (($v1_9)->tag === "Right")))) {
$__t8 = (($compare13_7_7)(($v_8)->value0))(($v1_9)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
};
};
};
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_5) use ($eq1Coproduct2_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eq1Coproduct2_4_2;
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
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_ordCoproduct'] = function() { $v = function($dictOrd1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_ordCoproduct"), recVars=[];
  $ord1Coproduct1_1_0 = (($GLOBALS['Data_Functor_Coproduct_ord1Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_ord1Coproduct')))($dictOrd1_0);
  $__local_var_2_1 = (($dictOrd1_0)->Eq10)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($dictOrd11_3) use ($__local_var_2_1, $ord1Coproduct1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_2 = (($dictOrd11_3)->Eq10)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($dictOrd_5) use ($__local_var_2_1, $__local_var_4_2, $dictOrd11_3, $ord1Coproduct1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_6_3 = (($dictOrd_5)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eq12_7_4 = (($__local_var_2_1)->eq1)($__local_var_6_3);
  $eq13_8_5 = (($__local_var_4_2)->eq1)($__local_var_6_3);
  $eqCoproduct3_7_4 = (object)["eq" => (function() use ($eq12_7_4, $eq13_8_5) {
  $__fn = function($v_9, $v1_10 = null) use ($eq12_7_4, $eq13_8_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_9) && (($v_9)->tag === "Left"))) {
$__t6 = ((is_object($v1_10) && (($v1_10)->tag === "Left")) && (($eq12_7_4)(($v_9)->value0))(($v1_10)->value0));
} else {
$__t6 = ((is_object($v_9) && (($v_9)->tag === "Right")) && ((is_object($v1_10) && (($v1_10)->tag === "Right")) && (($eq13_8_5)(($v_9)->value0))(($v1_10)->value0)));
};
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => ((($ord1Coproduct1_1_0)($dictOrd11_3))->compare1)($dictOrd_5), "Eq0" => function($dollar__unused_8) use ($eqCoproduct3_7_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqCoproduct3_7_4;
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
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_coproduct'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_coproduct"), recVars=[];
  if ((is_object($v2_2) && (($v2_2)->tag === "Left"))) {
$__t0 = ($v_0)(($v2_2)->value0);
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Right"))) {
$__t0 = ($v1_1)(($v2_2)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_extendCoproduct'] = function() { $v = function($dictExtend_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_extendCoproduct"), recVars=[];
  $__local_var_1_0 = (($dictExtend_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($dictExtend1_2) use ($__local_var_1_0, $dictExtend_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_3_1 = (($dictExtend1_2)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorCoproduct2_4_2 = (object)["map" => (function() use ($__local_var_1_0, $__local_var_3_1) {
  $__fn = function($f_4, $v_5 = null) use ($__local_var_1_0, $__local_var_3_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_6_2 = (($__local_var_1_0)->map)($f_4);
  $__local_var_7_3 = (($__local_var_3_1)->map)($f_4);
  if ((is_object($v_5) && (($v_5)->tag === "Left"))) {
$__t4 = new Phpurs_Data1("Left", ($__local_var_6_2)(($v_5)->value0));
} else {
if ((is_object($v_5) && (($v_5)->tag === "Right"))) {
$__t4 = new Phpurs_Data1("Right", ($__local_var_7_3)(($v_5)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["extend" => function($f_5) use ($dictExtend1_2, $dictExtend_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_6_6 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left'))))((($dictExtend_0)->extend)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_5))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left'))))));
  $__local_var_7_7 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))((($dictExtend1_2)->extend)(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($f_5))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))));
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(function($v2_8) use ($__local_var_6_6, $__local_var_7_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_8) && (($v2_8)->tag === "Left"))) {
$__t8 = ($__local_var_6_6)(($v2_8)->value0);
} else {
if ((is_object($v2_8) && (($v2_8)->tag === "Right"))) {
$__t8 = ($__local_var_7_7)(($v2_8)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
};
};
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_5) use ($functorCoproduct2_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $functorCoproduct2_4_2;
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
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_comonadCoproduct'] = function() { $v = function($dictComonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_comonadCoproduct"), recVars=[];
  $extendCoproduct1_1_0 = (($GLOBALS['Data_Functor_Coproduct_extendCoproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_extendCoproduct')))((($dictComonad_0)->Extend0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictComonad1_2) use ($dictComonad_0, $extendCoproduct1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $extendCoproduct2_3_1 = ($extendCoproduct1_1_0)((($dictComonad1_2)->Extend0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["extract" => function($v2_4) use ($dictComonad1_2, $dictComonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v2_4) && (($v2_4)->tag === "Left"))) {
$__t2 = (($dictComonad_0)->extract)(($v2_4)->value0);
} else {
if ((is_object($v2_4) && (($v2_4)->tag === "Right"))) {
$__t2 = (($dictComonad1_2)->extract)(($v2_4)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Extend0" => function($dollar__unused_4) use ($extendCoproduct2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $extendCoproduct2_3_1;
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
}; return $v; };
\PhpursThunks::$thunks['Data_Functor_Coproduct_bihoistCoproduct'] = function() { $v = (function() {
  $__fn = function($natF_0, $natG_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Functor_Coproduct_bihoistCoproduct"), recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Left"))) {
$__t0 = new Phpurs_Data1("Left", ($natF_0)(($v_2)->value0));
} else {
if ((is_object($v_2) && (($v_2)->tag === "Right"))) {
$__t0 = new Phpurs_Data1("Right", ($natG_1)(($v_2)->value0));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
















