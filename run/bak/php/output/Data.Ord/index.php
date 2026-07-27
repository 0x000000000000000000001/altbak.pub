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
function majData_majOrd_ordmajArraymajImpl($v0, $v1 = null, $v2 = null): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majOrd_ordmajArraymajImpl';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Data_Ord;
  $f = ($ffi_Data_Ord['ordArrayImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Data_Ord_ordArrayImpl'] = __NAMESPACE__ . '\\majData_majOrd_ordmajArraymajImpl';

function majData_majOrd_ordmajBooleanmajImpl($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majOrd_ordmajBooleanmajImpl';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_Ord;
  $f = ($ffi_Data_Ord['ordBooleanImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Ord_ordBooleanImpl'] = __NAMESPACE__ . '\\majData_majOrd_ordmajBooleanmajImpl';

function majData_majOrd_ordmajCharmajImpl($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majOrd_ordmajCharmajImpl';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_Ord;
  $f = ($ffi_Data_Ord['ordCharImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Ord_ordCharImpl'] = __NAMESPACE__ . '\\majData_majOrd_ordmajCharmajImpl';

function majData_majOrd_ordmajIntmajImpl($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majOrd_ordmajIntmajImpl';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_Ord;
  $f = ($ffi_Data_Ord['ordIntImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Ord_ordIntImpl'] = __NAMESPACE__ . '\\majData_majOrd_ordmajIntmajImpl';

function majData_majOrd_ordmajNumbermajImpl($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majOrd_ordmajNumbermajImpl';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_Ord;
  $f = ($ffi_Data_Ord['ordNumberImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Ord_ordNumberImpl'] = __NAMESPACE__ . '\\majData_majOrd_ordmajNumbermajImpl';

function majData_majOrd_ordmajStringmajImpl($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majOrd_ordmajStringmajImpl';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Data_Ord;
  $f = ($ffi_Data_Ord['ordStringImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Data_Ord_ordStringImpl'] = __NAMESPACE__ . '\\majData_majOrd_ordmajStringmajImpl';





// Data_Ord_ordVoid
$GLOBALS['Data_Ord_ordVoid'] = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Ordering\Data_Ordering_EQ();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqVoid'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordUnit
$GLOBALS['Data_Ord_ordUnit'] = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Ordering\Data_Ordering_EQ();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqUnit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordString
$GLOBALS['Data_Ord_ordString'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordStringImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqString'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordRecordNil
$GLOBALS['Data_Ord_ordRecordNil'] = (object)["compareRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new \Data\Ordering\Data_Ordering_EQ();
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "EqRecord0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqRowNil'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordProxy
$GLOBALS['Data_Ord_ordProxy'] = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Ordering\Data_Ordering_EQ();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqProxy'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordOrdering
$GLOBALS['Data_Ord_ordOrdering'] = (object)["compare" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = null;;
if ($v1_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_1;;
};
$__t1 = new \Data\Ordering\Data_Ordering_LT();
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = null;;
if ($v1_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_2;;
};
if ($v1_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t2 = new \Data\Ordering\Data_Ordering_GT();
goto end_branch_2;;
};
if ($v1_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t2 = new \Data\Ordering\Data_Ordering_LT();
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t0 = $__t2;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = null;;
if ($v1_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t3 = new \Data\Ordering\Data_Ordering_EQ();
goto end_branch_3;;
};
$__t3 = new \Data\Ordering\Data_Ordering_GT();
end_branch_3:;
$__t0 = $__t3;
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
})(), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Ordering_eqOrdering'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordNumber
$GLOBALS['Data_Ord_ordNumber'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordNumberImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqNumber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordInt
$GLOBALS['Data_Ord_ordInt'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordChar
$GLOBALS['Data_Ord_ordChar'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordCharImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqChar'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordBoolean
$GLOBALS['Data_Ord_ordBoolean'] = (object)["compare" => ((($GLOBALS['Data_Ord_ordBooleanImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eqBoolean'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_compareRecord
function majData_majOrd_comparemajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_comparemajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'compareRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ord_compareRecord'] = __NAMESPACE__ . '\\majData_majOrd_comparemajRecord';

// Data_Ord_ordRecord
function majData_majOrd_ordmajRecord($_dollar__unused_0, $dictOrdRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_ordmajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $eqRec1_2_0 = (object)["eq" => (((($dictOrdRecord_1)->{'EqRecord0'})(null))->{'eqRecord'})(new \Type\Proxy\Type_Proxy_Proxy())];
  $__res = (object)["compare" => (($dictOrdRecord_1)->{'compareRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "Eq0" => function($_dollar__unused_3) use ($eqRec1_2_0) {
  $__num = \func_num_args();
  $__res = $eqRec1_2_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Ord_ordRecord'] = __NAMESPACE__ . '\\majData_majOrd_ordmajRecord';

// Data_Ord_compare1
function majData_majOrd_compare1($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_compare1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'compare1'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ord_compare1'] = __NAMESPACE__ . '\\majData_majOrd_compare1';

// Data_Ord_compare
function majData_majOrd_compare($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_compare';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'compare'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ord_compare'] = __NAMESPACE__ . '\\majData_majOrd_compare';

// Data_Ord_compare2
$GLOBALS['Data_Ord_compare2'] = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());

// Data_Ord_comparing
function majData_majOrd_comparing($dictOrd_0, $f_1 = null, $x_2 = null, $y_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_comparing';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictOrd_0)->{'compare'})(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Ord_comparing'] = __NAMESPACE__ . '\\majData_majOrd_comparing';

// Data_Ord_greaterThan
function majData_majOrd_greatermajThan($dictOrd_0, $a1_1 = null, $a2_2 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_greatermajThan';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictOrd_0)->{'compare'})($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Ord_greaterThan'] = __NAMESPACE__ . '\\majData_majOrd_greatermajThan';

// Data_Ord_greaterThanOrEq
function majData_majOrd_greatermajThanmajOrmajEq($dictOrd_0, $a1_1 = null, $a2_2 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_greatermajThanmajOrmajEq';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ( ! ((($dictOrd_0)->{'compare'})($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Ord_greaterThanOrEq'] = __NAMESPACE__ . '\\majData_majOrd_greatermajThanmajOrmajEq';

// Data_Ord_lessThan
function majData_majOrd_lessmajThan($dictOrd_0, $a1_1 = null, $a2_2 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_lessmajThan';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictOrd_0)->{'compare'})($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Ord_lessThan'] = __NAMESPACE__ . '\\majData_majOrd_lessmajThan';

// Data_Ord_signum
function majData_majOrd_signum($dictOrd_0, $dictRing_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_signum';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $Semiring0_2_0 = (($dictRing_1)->{'Semiring0'})(null);
  $zero_3_1 = ($Semiring0_2_0)->{'zero'};
  $zero_4_2 = ((($dictRing_1)->{'Semiring0'})(null))->{'zero'};
  $one_5_3 = ($Semiring0_2_0)->{'one'};
  $__res = function($x_6) use ($dictOrd_0, $dictRing_1, $one_5_3, $zero_3_1, $zero_4_2) {
  $__num = \func_num_args();
  $__t4 = null;;
  if (((($dictOrd_0)->{'compare'})($x_6))($zero_3_1) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = ((($dictRing_1)->{'sub'})($zero_4_2))($one_5_3);
goto end_branch_4;;
};
  if (((($dictOrd_0)->{'compare'})($x_6))($zero_3_1) instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = $one_5_3;
goto end_branch_4;;
};
  $__t4 = $x_6;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Ord_signum'] = __NAMESPACE__ . '\\majData_majOrd_signum';

// Data_Ord_lessThanOrEq
function majData_majOrd_lessmajThanmajOrmajEq($dictOrd_0, $a1_1 = null, $a2_2 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_lessmajThanmajOrmajEq';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ( ! ((($dictOrd_0)->{'compare'})($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Ord_lessThanOrEq'] = __NAMESPACE__ . '\\majData_majOrd_lessmajThanmajOrmajEq';

// Data_Ord_max
function majData_majOrd_max($dictOrd_0, $x_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_max';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_3_0 = ((($dictOrd_0)->{'compare'})($x_1))($y_2);
  $__t1 = null;;
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $y_2;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = $x_1;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = $x_1;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Ord_max'] = __NAMESPACE__ . '\\majData_majOrd_max';

// Data_Ord_min
function majData_majOrd_min($dictOrd_0, $x_1 = null, $y_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_min';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $v_3_0 = ((($dictOrd_0)->{'compare'})($x_1))($y_2);
  $__t1 = null;;
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $x_1;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = $x_1;
goto end_branch_1;;
};
  if ($v_3_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = $y_2;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Ord_min'] = __NAMESPACE__ . '\\majData_majOrd_min';

// Data_Ord_ordArray
function majData_majOrd_ordmajArray($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_ordmajArray';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eqArray_1_0 = (object)["eq" => ($GLOBALS['Data_Eq_eqArrayImpl'])(((($dictOrd_0)->{'Eq0'})(null))->{'eq'})];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($xs_2, $ys_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Ord_compare2'])(0))(\Data\Ord\majData_majOrd_ordmajArraymajImpl((function() use ($dictOrd_0) {
  $__fn = function($x_4, $y_5 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_6_1 = ((($dictOrd_0)->{'compare'})($x_4))($y_5);
  $__t2 = null;;
  if ($v_6_1 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t2 = 0;
goto end_branch_2;;
};
  if ($v_6_1 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t2 = 1;
goto end_branch_2;;
};
  if ($v_6_1 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t2 = -1;
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), $xs_2, $ys_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($_dollar__unused_2) use ($eqArray_1_0) {
  $__num = \func_num_args();
  $__res = $eqArray_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ord_ordArray'] = __NAMESPACE__ . '\\majData_majOrd_ordmajArray';

// Data_Ord_ord1Array
$GLOBALS['Data_Ord_ord1Array'] = (object)["compare1" => function($dictOrd_0) {
  $__num = \func_num_args();
  $__res = (\Data\Ord\majData_majOrd_ordmajArray($dictOrd_0))->{'compare'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Eq_eq1Array'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_Ord_ordRecordCons
function majData_majOrd_ordmajRecordmajCons($dictOrdRecord_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_ordmajRecordmajCons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eqRowCons_1_0 = (($GLOBALS['Data_Eq_eqRowCons'])((($dictOrdRecord_0)->{'EqRecord0'})(null)))(null);
  $__res = (function() use ($dictOrdRecord_0, $eqRowCons_1_0) {
  $__fn = function($_dollar__unused_2, $dictIsSymbol_3 = null) use ($dictOrdRecord_0, $eqRowCons_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $eqRowCons1_4_1 = ($eqRowCons_1_0)($dictIsSymbol_3);
  $__res = function($dictOrd_5) use ($dictIsSymbol_3, $dictOrdRecord_0, $eqRowCons1_4_1) {
  $__num = \func_num_args();
  $eqRowCons2_6_2 = ($eqRowCons1_4_1)((($dictOrd_5)->{'Eq0'})(null));
  $__res = (object)["compareRecord" => (function() use ($dictIsSymbol_3, $dictOrdRecord_0, $dictOrd_5) {
  $__fn = function($v_7, $ra_8 = null, $rb_9 = null) use ($dictIsSymbol_3, $dictOrdRecord_0, $dictOrd_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_10_3 = (($dictIsSymbol_3)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $left_11_4 = ((($dictOrd_5)->{'compare'})(\Record\Unsafe\majRecord_majUnsafe_unsafemajGet($key_10_3, $ra_8)))(\Record\Unsafe\majRecord_majUnsafe_unsafemajGet($key_10_3, $rb_9));
  $__t5 = null;;
  if (((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})($left_11_4))(new \Data\Ordering\Data_Ordering_EQ())))(false)) {
$__t5 = $left_11_4;
goto end_branch_5;;
};
  $__t5 = (((($dictOrdRecord_0)->{'compareRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_8))($rb_9);
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "EqRecord0" => function($_dollar__unused_7) use ($eqRowCons2_6_2) {
  $__num = \func_num_args();
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Ord_ordRecordCons'] = __NAMESPACE__ . '\\majData_majOrd_ordmajRecordmajCons';

// Data_Ord_clamp
function majData_majOrd_clamp($dictOrd_0, $low_1 = null, $hi_2 = null, $x_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_clamp';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $v_4_0 = ((($dictOrd_0)->{'compare'})($low_1))($x_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t1 = $x_3;
goto end_branch_1;;
};
  if ($v_4_0 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t1 = $low_1;
goto end_branch_1;;
};
  if ($v_4_0 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t1 = $low_1;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__local_var_5_1 = $__t1;
  $v_6_3 = ((($dictOrd_0)->{'compare'})($hi_2))($__local_var_5_1);
  $__t4 = null;;
  if ($v_6_3 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t4 = $hi_2;
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t4 = $hi_2;
goto end_branch_4;;
};
  if ($v_6_3 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t4 = $__local_var_5_1;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Ord_clamp'] = __NAMESPACE__ . '\\majData_majOrd_clamp';

// Data_Ord_between
function majData_majOrd_between($dictOrd_0, $low_1 = null, $hi_2 = null, $x_3 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_between';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__t0 = null;;
  if (((($dictOrd_0)->{'compare'})($x_3))($low_1) instanceof \Data\Ordering\Data_Ordering_LT) {
$__t0 = false;
goto end_branch_0;;
};
  $__t0 = ( ! ((($dictOrd_0)->{'compare'})($x_3))($hi_2) instanceof \Data\Ordering\Data_Ordering_GT);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Ord_between'] = __NAMESPACE__ . '\\majData_majOrd_between';

// Data_Ord_abs
function majData_majOrd_abs($dictOrd_0, $dictRing_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majOrd_abs';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $zero_2_0 = ((($dictRing_1)->{'Semiring0'})(null))->{'zero'};
  $zero_3_1 = ((($dictRing_1)->{'Semiring0'})(null))->{'zero'};
  $__res = function($x_4) use ($dictOrd_0, $dictRing_1, $zero_2_0, $zero_3_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (( ! ((($dictOrd_0)->{'compare'})($x_4))($zero_2_0) instanceof \Data\Ordering\Data_Ordering_LT)) {
$__t2 = $x_4;
goto end_branch_2;;
};
  $__t2 = ((($dictRing_1)->{'sub'})($zero_3_1))($x_4);
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Ord_abs'] = __NAMESPACE__ . '\\majData_majOrd_abs';

