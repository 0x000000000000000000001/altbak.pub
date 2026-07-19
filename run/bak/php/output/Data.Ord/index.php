<?php

namespace Data\Ord;

// ALL IMPORTS: Data.Eq, Data.Ord, Data.Ordering, Data.Ring, Data.Semiring, Data.Symbol, Data.Unit, Data.Void, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Eq, Data.Ord, Data.Ordering, Data.Ring, Data.Semiring, Data.Symbol, Data.Unit, Data.Void, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Data.Void/index.php';
require_once __DIR__ . '/../Record.Unsafe/index.php';
require_once __DIR__ . '/../Type.Proxy/index.php';

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
\PhpursThunks::$thunks['Data_Ord_ordVoid'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("EQ");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqVoid'] ?? \PhpursThunks::eval('Data_Eq_eqVoid'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordUnit'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("EQ");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqUnit'] ?? \PhpursThunks::eval('Data_Eq_eqUnit'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordString'] = function() { $v = (object)["compare" => (((($GLOBALS['Data_Ord_ordStringImpl'] ?? \PhpursThunks::eval('Data_Ord_ordStringImpl')))(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT")), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqString'] ?? \PhpursThunks::eval('Data_Eq_eqString'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordRecordNil'] = function() { $v = (object)["compareRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("EQ");
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "EqRecord0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqRowNil'] ?? \PhpursThunks::eval('Data_Eq_eqRowNil'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordProxy'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("EQ");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqProxy'] ?? \PhpursThunks::eval('Data_Eq_eqProxy'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordOrdering'] = function() { $v = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "LT"))) {
if ((is_object($v1_1) && (($v1_1)->tag === "LT"))) {
$__t1 = new Phpurs_Data0("EQ");
} else {
$__t1 = new Phpurs_Data0("LT");
};
$__t0 = $__t1;
} else {
if ((is_object($v_0) && (($v_0)->tag === "EQ"))) {
if ((is_object($v1_1) && (($v1_1)->tag === "EQ"))) {
$__t2 = new Phpurs_Data0("EQ");
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "LT"))) {
$__t2 = new Phpurs_Data0("GT");
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "GT"))) {
$__t2 = new Phpurs_Data0("LT");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
$__t0 = $__t2;
} else {
if ((is_object($v_0) && (($v_0)->tag === "GT"))) {
if ((is_object($v1_1) && (($v1_1)->tag === "GT"))) {
$__t3 = new Phpurs_Data0("EQ");
} else {
$__t3 = new Phpurs_Data0("GT");
};
$__t0 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Ordering_eqOrdering'] ?? \PhpursThunks::eval('Data_Ordering_eqOrdering'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordNumber'] = function() { $v = (object)["compare" => (((($GLOBALS['Data_Ord_ordNumberImpl'] ?? \PhpursThunks::eval('Data_Ord_ordNumberImpl')))(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT")), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqNumber'] ?? \PhpursThunks::eval('Data_Eq_eqNumber'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordInt'] = function() { $v = (object)["compare" => (((($GLOBALS['Data_Ord_ordIntImpl'] ?? \PhpursThunks::eval('Data_Ord_ordIntImpl')))(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT")), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqInt'] ?? \PhpursThunks::eval('Data_Eq_eqInt'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordChar'] = function() { $v = (object)["compare" => (((($GLOBALS['Data_Ord_ordCharImpl'] ?? \PhpursThunks::eval('Data_Ord_ordCharImpl')))(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT")), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqChar'] ?? \PhpursThunks::eval('Data_Eq_eqChar'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordBoolean'] = function() { $v = (object)["compare" => (((($GLOBALS['Data_Ord_ordBooleanImpl'] ?? \PhpursThunks::eval('Data_Ord_ordBooleanImpl')))(new Phpurs_Data0("LT")))(new Phpurs_Data0("EQ")))(new Phpurs_Data0("GT")), "Eq0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eqBoolean'] ?? \PhpursThunks::eval('Data_Eq_eqBoolean'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_compareRecord'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_compareRecord"), recVars=[];
  $__res = ($dict_0)->compareRecord;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordRecord'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0, $dictOrdRecord_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_ordRecord"), recVars=[];
  $eqRec1_2_0 = (object)["eq" => (((($dictOrdRecord_1)->EqRecord0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->eqRecord)(new Phpurs_Data0("Proxy"))];
  $__res = (object)["compare" => (($dictOrdRecord_1)->compareRecord)(new Phpurs_Data0("Proxy")), "Eq0" => function($dollar__unused_3) use ($eqRec1_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqRec1_2_0;
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
\PhpursThunks::$thunks['Data_Ord_compare1'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_compare1"), recVars=[];
  $__res = ($dict_0)->compare1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Ord_compare'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_compare"), recVars=[];
  $__res = ($dict_0)->compare;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Ord_comparing'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null, $x_2 = null, $y_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_comparing"), recVars=[];
  $__res = ((($dictOrd_0)->compare)(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_greaterThan'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $a1_1 = null, $a2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_greaterThan"), recVars=[];
  $__res = (is_object(((($dictOrd_0)->compare)($a1_1))($a2_2)) && ((((($dictOrd_0)->compare)($a1_1))($a2_2))->tag === "GT"));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_greaterThanOrEq'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $a1_1 = null, $a2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_greaterThanOrEq"), recVars=[];
  $__res = ( ! (is_object(((($dictOrd_0)->compare)($a1_1))($a2_2)) && ((((($dictOrd_0)->compare)($a1_1))($a2_2))->tag === "LT")));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_lessThan'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $a1_1 = null, $a2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_lessThan"), recVars=[];
  $__res = (is_object(((($dictOrd_0)->compare)($a1_1))($a2_2)) && ((((($dictOrd_0)->compare)($a1_1))($a2_2))->tag === "LT"));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_signum'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictRing_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_signum"), recVars=[];
  $Semiring0_2_0 = (($dictRing_1)->Semiring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $zero_3_1 = ($Semiring0_2_0)->zero;
  $zero_4_2 = ((($dictRing_1)->Semiring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->zero;
  $one_5_3 = ($Semiring0_2_0)->one;
  $__res = function($x_6) use ($dictOrd_0, $dictRing_1, $one_5_3, $zero_3_1, $zero_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object(((($dictOrd_0)->compare)($x_6))($zero_3_1)) && ((((($dictOrd_0)->compare)($x_6))($zero_3_1))->tag === "LT"))) {
$__t4 = ((($dictRing_1)->sub)($zero_4_2))($one_5_3);
} else {
if ((is_object(((($dictOrd_0)->compare)($x_6))($zero_3_1)) && ((((($dictOrd_0)->compare)($x_6))($zero_3_1))->tag === "GT"))) {
$__t4 = $one_5_3;
} else {
$__t4 = $x_6;
};
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_lessThanOrEq'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $a1_1 = null, $a2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_lessThanOrEq"), recVars=[];
  $__res = ( ! (is_object(((($dictOrd_0)->compare)($a1_1))($a2_2)) && ((((($dictOrd_0)->compare)($a1_1))($a2_2))->tag === "GT")));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_max'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_max"), recVars=[];
  $v_3_0 = ((($dictOrd_0)->compare)($x_1))($y_2);
  if ((is_object($v_3_0) && (($v_3_0)->tag === "LT"))) {
$__t1 = $y_2;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "EQ"))) {
$__t1 = $x_1;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "GT"))) {
$__t1 = $x_1;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_min'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $x_1 = null, $y_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_min"), recVars=[];
  $v_3_0 = ((($dictOrd_0)->compare)($x_1))($y_2);
  if ((is_object($v_3_0) && (($v_3_0)->tag === "LT"))) {
$__t1 = $x_1;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "EQ"))) {
$__t1 = $x_1;
} else {
if ((is_object($v_3_0) && (($v_3_0)->tag === "GT"))) {
$__t1 = $y_2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_ordArray'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_ordArray"), recVars=[];
  $eqArray_1_0 = (object)["eq" => (($GLOBALS['Data_Eq_eqArrayImpl'] ?? \PhpursThunks::eval('Data_Eq_eqArrayImpl')))(((($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->eq)];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($xs_2, $ys_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)(0))((((($GLOBALS['Data_Ord_ordArrayImpl'] ?? \PhpursThunks::eval('Data_Ord_ordArrayImpl')))((function() use ($dictOrd_0) {
  $__fn = function($x_4, $y_5 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_6_1 = ((($dictOrd_0)->compare)($x_4))($y_5);
  if ((is_object($v_6_1) && (($v_6_1)->tag === "EQ"))) {
$__t2 = 0;
} else {
if ((is_object($v_6_1) && (($v_6_1)->tag === "LT"))) {
$__t2 = 1;
} else {
if ((is_object($v_6_1) && (($v_6_1)->tag === "GT"))) {
$__t2 = -1;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($xs_2))($ys_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_2) use ($eqArray_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqArray_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Ord_ord1Array'] = function() { $v = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Ord_ordArray'] ?? \PhpursThunks::eval('Data_Ord_ordArray')))($dictOrd_0))->compare;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Eq_eq1Array'] ?? \PhpursThunks::eval('Data_Eq_eq1Array'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Ord_ordRecordCons'] = function() { $v = function($dictOrdRecord_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_ordRecordCons"), recVars=[];
  $eqRowCons_1_0 = ((($GLOBALS['Data_Eq_eqRowCons'] ?? \PhpursThunks::eval('Data_Eq_eqRowCons')))((($dictOrdRecord_0)->EqRecord0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')))))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (function() use ($dictOrdRecord_0, $eqRowCons_1_0) {
  $__fn = function($dollar__unused_2, $dictIsSymbol_3 = null) use ($dictOrdRecord_0, $eqRowCons_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eqRowCons1_4_1 = ($eqRowCons_1_0)($dictIsSymbol_3);
  $__res = function($dictOrd_5) use ($dictIsSymbol_3, $dictOrdRecord_0, $eqRowCons1_4_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eqRowCons2_6_2 = ($eqRowCons1_4_1)((($dictOrd_5)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["compareRecord" => (function() use ($dictIsSymbol_3, $dictOrdRecord_0, $dictOrd_5) {
  $__fn = function($v_7, $ra_8 = null, $rb_9 = null) use ($dictIsSymbol_3, $dictOrdRecord_0, $dictOrd_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $key_10_3 = (($dictIsSymbol_3)->reflectSymbol)(new Phpurs_Data0("Proxy"));
  $left_11_4 = ((($dictOrd_5)->compare)(((($GLOBALS['Record_Unsafe_unsafeGet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeGet')))($key_10_3))($ra_8)))(((($GLOBALS['Record_Unsafe_unsafeGet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeGet')))($key_10_3))($rb_9));
  if (((is_object($left_11_4) && (($left_11_4)->tag === "LT")) || ((is_object($left_11_4) && (($left_11_4)->tag === "GT")) || ( ! (is_object($left_11_4) && (($left_11_4)->tag === "EQ")))))) {
$__t5 = $left_11_4;
} else {
$__t5 = (((($dictOrdRecord_0)->compareRecord)(new Phpurs_Data0("Proxy")))($ra_8))($rb_9);
};
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "EqRecord0" => function($dollar__unused_7) use ($eqRowCons2_6_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqRowCons2_6_2;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Ord_clamp'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $low_1 = null, $hi_2 = null, $x_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_clamp"), recVars=[];
  $v_4_0 = ((($dictOrd_0)->compare)($low_1))($x_3);
  if ((is_object($v_4_0) && (($v_4_0)->tag === "LT"))) {
$__t1 = $x_3;
} else {
if ((is_object($v_4_0) && (($v_4_0)->tag === "EQ"))) {
$__t1 = $low_1;
} else {
if ((is_object($v_4_0) && (($v_4_0)->tag === "GT"))) {
$__t1 = $low_1;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
  $__local_var_5_1 = $__t1;
  $v_6_3 = ((($dictOrd_0)->compare)($hi_2))($__local_var_5_1);
  if ((is_object($v_6_3) && (($v_6_3)->tag === "LT"))) {
$__t4 = $hi_2;
} else {
if ((is_object($v_6_3) && (($v_6_3)->tag === "EQ"))) {
$__t4 = $hi_2;
} else {
if ((is_object($v_6_3) && (($v_6_3)->tag === "GT"))) {
$__t4 = $__local_var_5_1;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
};
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_between'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $low_1 = null, $hi_2 = null, $x_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_between"), recVars=[];
  if ((is_object(((($dictOrd_0)->compare)($x_3))($low_1)) && ((((($dictOrd_0)->compare)($x_3))($low_1))->tag === "LT"))) {
$__t0 = false;
} else {
$__t0 = ( ! (is_object(((($dictOrd_0)->compare)($x_3))($hi_2)) && ((((($dictOrd_0)->compare)($x_3))($hi_2))->tag === "GT")));
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Ord_abs'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictRing_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Ord_abs"), recVars=[];
  $zero_2_0 = ((($dictRing_1)->Semiring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->zero;
  $zero_3_1 = ((($dictRing_1)->Semiring0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->zero;
  $__res = function($x_4) use ($dictOrd_0, $dictRing_1, $zero_2_0, $zero_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (( ! (is_object(((($dictOrd_0)->compare)($x_4))($zero_2_0)) && ((((($dictOrd_0)->compare)($x_4))($zero_2_0))->tag === "LT")))) {
$__t2 = $x_4;
} else {
$__t2 = ((($dictRing_1)->sub)($zero_3_1))($x_4);
};
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
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Ord = \call_user_func(function() {
  $exports = [];
$ordIntImpl = function($lt, $eq = null, $gt = null, $x = null, $y = null) use (&$ordIntImpl) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$ordIntImpl) {

            return $ordIntImpl(...\array_merge($__args, $more));
        };
    }
    return $x < $y ? $lt : ($x === $y ? $eq : $gt);
};
$ordStringImpl = $ordIntImpl;
$ordNumberImpl = $ordIntImpl;
$ordCharImpl = $ordIntImpl;
$ordBooleanImpl = $ordIntImpl;

$exports['ordIntImpl'] = $ordIntImpl;
$exports['ordStringImpl'] = $ordStringImpl;
$exports['ordNumberImpl'] = $ordNumberImpl;
$exports['ordCharImpl'] = $ordCharImpl;
$exports['ordBooleanImpl'] = $ordBooleanImpl;

$ordArrayImpl = function($f, $xs = null, $ys = null) use (&$ordArrayImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$ordArrayImpl) {
            return $ordArrayImpl(...\array_merge($__args, $more));
        };
    }
    
    $i = 0;
    $xlen = \count($xs);
    $ylen = \count($ys);
    while ($i < $xlen && $i < $ylen) {
        $x = $xs[$i];
        $y = $ys[$i];
        $o = $f($x, $y);
        if ($o !== 0) {
            return $o;
        }
        $i++;
    }
    if ($xlen === $ylen) {
        return 0;
    } else if ($xlen > $ylen) {
        return -1;
    } else {
        return 1;
    }
};

$exports['ordArrayImpl'] = $ordArrayImpl;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Ord_ordArrayImpl'] = function() use (&$ffi_Data_Ord) { return $ffi_Data_Ord['ordArrayImpl']; };
\PhpursThunks::$thunks['Data_Ord_ordBooleanImpl'] = function() use (&$ffi_Data_Ord) { return $ffi_Data_Ord['ordBooleanImpl']; };
\PhpursThunks::$thunks['Data_Ord_ordCharImpl'] = function() use (&$ffi_Data_Ord) { return $ffi_Data_Ord['ordCharImpl']; };
\PhpursThunks::$thunks['Data_Ord_ordIntImpl'] = function() use (&$ffi_Data_Ord) { return $ffi_Data_Ord['ordIntImpl']; };
\PhpursThunks::$thunks['Data_Ord_ordNumberImpl'] = function() use (&$ffi_Data_Ord) { return $ffi_Data_Ord['ordNumberImpl']; };
\PhpursThunks::$thunks['Data_Ord_ordStringImpl'] = function() use (&$ffi_Data_Ord) { return $ffi_Data_Ord['ordStringImpl']; };






























