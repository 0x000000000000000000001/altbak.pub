<?php

namespace Data\HeytingAlgebra;

// ALL IMPORTS: Data.HeytingAlgebra, Data.Symbol, Data.Unit, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.HeytingAlgebra, Data.Symbol, Data.Unit, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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
$ffi_Data_HeytingAlgebra = \call_user_func(function() {
  $exports = [];
$boolConj = function($a, $b = null) use (&$boolConj) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$boolConj) {

            return $boolConj(...\array_merge($__args, $more));
        };
    }
    return $a && $b;
};
$boolDisj = function($a, $b = null) use (&$boolDisj) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$boolDisj) {

            return $boolDisj(...\array_merge($__args, $more));
        };
    }
    return $a || $b;
};
$boolNot = function($a) use (&$boolNot) { return !$a; };

$exports['boolConj'] = $boolConj;
$exports['boolDisj'] = $boolDisj;
$exports['boolNot'] = $boolNot;
return $exports;
  return $exports;
});
function majData_majHeytingmajAlgebra_boolmajConj(bool $v0, $v1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_boolmajConj';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_HeytingAlgebra;
  $f = ($ffi_Data_HeytingAlgebra['boolConj'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_HeytingAlgebra_boolConj'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_boolmajConj';

function majData_majHeytingmajAlgebra_boolmajDisj(bool $v0, $v1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_boolmajDisj';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Data_HeytingAlgebra;
  $f = ($ffi_Data_HeytingAlgebra['boolDisj'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Data_HeytingAlgebra_boolDisj'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_boolmajDisj';

function majData_majHeytingmajAlgebra_boolmajNot(bool $v0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_boolmajNot';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_HeytingAlgebra;
  $f = ($ffi_Data_HeytingAlgebra['boolNot'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_HeytingAlgebra_boolNot'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_boolmajNot';





// Data_HeytingAlgebra_ttRecord
function majData_majHeytingmajAlgebra_ttmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_ttmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'ttRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_ttRecord'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_ttmajRecord';

// Data_HeytingAlgebra_tt
function majData_majHeytingmajAlgebra_tt($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_tt';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'tt'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_tt'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_tt';

// Data_HeytingAlgebra_notRecord
function majData_majHeytingmajAlgebra_notmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_notmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'notRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_notRecord'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_notmajRecord';

// Data_HeytingAlgebra_not
function majData_majHeytingmajAlgebra_not($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_not';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'not'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_not'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_not';

// Data_HeytingAlgebra_impliesRecord
function majData_majHeytingmajAlgebra_impliesmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_impliesmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'impliesRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_impliesRecord'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_impliesmajRecord';

// Data_HeytingAlgebra_implies
function majData_majHeytingmajAlgebra_implies($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_implies';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'implies'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_implies'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_implies';

// Data_HeytingAlgebra_heytingAlgebraUnit
$GLOBALS['Data_HeytingAlgebra_heytingAlgebraUnit'] = (object)["ff" => $GLOBALS['Data_Unit_unit'], "tt" => $GLOBALS['Data_Unit_unit'], "implies" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "conj" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "disj" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "not" => function($v_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Data_HeytingAlgebra_heytingAlgebraRecordNil
$GLOBALS['Data_HeytingAlgebra_heytingAlgebraRecordNil'] = (object)["conjRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "disjRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "ffRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "impliesRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "notRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "ttRecord" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)[];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];

// Data_HeytingAlgebra_heytingAlgebraProxy
$GLOBALS['Data_HeytingAlgebra_heytingAlgebraProxy'] = (object)["conj" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Type\Proxy\Type_Proxy_Proxy();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "disj" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Type\Proxy\Type_Proxy_Proxy();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "implies" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Type\Proxy\Type_Proxy_Proxy();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "ff" => new \Type\Proxy\Type_Proxy_Proxy(), "not" => function($v_0) {
  $__num = \func_num_args();
  $__res = new \Type\Proxy\Type_Proxy_Proxy();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "tt" => new \Type\Proxy\Type_Proxy_Proxy()];

// Data_HeytingAlgebra_ffRecord
function majData_majHeytingmajAlgebra_ffmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_ffmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'ffRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_ffRecord'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_ffmajRecord';

// Data_HeytingAlgebra_ff
function majData_majHeytingmajAlgebra_ff($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_ff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'ff'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_ff'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_ff';

// Data_HeytingAlgebra_disjRecord
function majData_majHeytingmajAlgebra_disjmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_disjmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'disjRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_disjRecord'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_disjmajRecord';

// Data_HeytingAlgebra_disj
function majData_majHeytingmajAlgebra_disj($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_disj';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'disj'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_disj'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_disj';

// Data_HeytingAlgebra_heytingAlgebraBoolean
$GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'] = (object)["ff" => false, "tt" => true, "implies" => (function() {
  $__fn = function($a_0, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'disj'})((($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'})($a_0)))($b_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "conj" => $GLOBALS['Data_HeytingAlgebra_boolConj'], "disj" => $GLOBALS['Data_HeytingAlgebra_boolDisj'], "not" => $GLOBALS['Data_HeytingAlgebra_boolNot']];

// Data_HeytingAlgebra_conjRecord
function majData_majHeytingmajAlgebra_conjmajRecord($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_conjmajRecord';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'conjRecord'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_conjRecord'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_conjmajRecord';

// Data_HeytingAlgebra_heytingAlgebraRecord
function majData_majHeytingmajAlgebra_heytingmajAlgebramajRecord($_dollar__unused_0, $dictHeytingAlgebraRecord_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_heytingmajAlgebramajRecord';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["ff" => ((($dictHeytingAlgebraRecord_1)->{'ffRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "tt" => ((($dictHeytingAlgebraRecord_1)->{'ttRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))(new \Type\Proxy\Type_Proxy_Proxy()), "conj" => (($dictHeytingAlgebraRecord_1)->{'conjRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "disj" => (($dictHeytingAlgebraRecord_1)->{'disjRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "implies" => (($dictHeytingAlgebraRecord_1)->{'impliesRecord'})(new \Type\Proxy\Type_Proxy_Proxy()), "not" => (($dictHeytingAlgebraRecord_1)->{'notRecord'})(new \Type\Proxy\Type_Proxy_Proxy())];
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_heytingAlgebraRecord'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_heytingmajAlgebramajRecord';

// Data_HeytingAlgebra_conj
function majData_majHeytingmajAlgebra_conj($dict_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_conj';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($dict_0)->{'conj'};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_conj'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_conj';

// Data_HeytingAlgebra_heytingAlgebraFunction
function majData_majHeytingmajAlgebra_heytingmajAlgebramajFunction($dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_heytingmajAlgebramajFunction';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $ff1_1_0 = ($dictHeytingAlgebra_0)->{'ff'};
  $tt1_2_1 = ($dictHeytingAlgebra_0)->{'tt'};
  $__res = (object)["ff" => function($v_3) use ($ff1_1_0) {
  $__num = \func_num_args();
  $__res = $ff1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "tt" => function($v_3) use ($tt1_2_1) {
  $__num = \func_num_args();
  $__res = $tt1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "implies" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($f_3, $g_4 = null, $a_5 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictHeytingAlgebra_0)->{'implies'})(($f_3)($a_5)))(($g_4)($a_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "conj" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($f_3, $g_4 = null, $a_5 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictHeytingAlgebra_0)->{'conj'})(($f_3)($a_5)))(($g_4)($a_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "disj" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($f_3, $g_4 = null, $a_5 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictHeytingAlgebra_0)->{'disj'})(($f_3)($a_5)))(($g_4)($a_5));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "not" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($f_3, $a_4 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictHeytingAlgebra_0)->{'not'})(($f_3)($a_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_heytingAlgebraFunction'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_heytingmajAlgebramajFunction';

// Data_HeytingAlgebra_heytingAlgebraRecordCons
function majData_majHeytingmajAlgebra_heytingmajAlgebramajRecordmajCons($dictIsSymbol_0, $_dollar__unused_1 = null, $dictHeytingAlgebraRecord_2 = null, $dictHeytingAlgebra_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majHeytingmajAlgebra_heytingmajAlgebramajRecordmajCons';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $ff1_4_0 = ($dictHeytingAlgebra_3)->{'ff'};
  $tt1_5_1 = ($dictHeytingAlgebra_3)->{'tt'};
  $__res = (object)["conjRecord" => (function() use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0) {
  $__fn = function($v_6, $ra_7 = null, $rb_8 = null) use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_9_2 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_10_3 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_9_2);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_9_2, ((($dictHeytingAlgebra_3)->{'conj'})(($get_10_3)($ra_7)))(($get_10_3)($rb_8)), (((($dictHeytingAlgebraRecord_2)->{'conjRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_7))($rb_8));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "disjRecord" => (function() use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0) {
  $__fn = function($v_6, $ra_7 = null, $rb_8 = null) use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_9_4 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_10_5 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_9_4);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_9_4, ((($dictHeytingAlgebra_3)->{'disj'})(($get_10_5)($ra_7)))(($get_10_5)($rb_8)), (((($dictHeytingAlgebraRecord_2)->{'disjRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_7))($rb_8));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "impliesRecord" => (function() use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0) {
  $__fn = function($v_6, $ra_7 = null, $rb_8 = null) use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $key_9_6 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $get_10_7 = ($GLOBALS['Record_Unsafe_unsafeGet'])($key_9_6);
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_9_6, ((($dictHeytingAlgebra_3)->{'implies'})(($get_10_7)($ra_7)))(($get_10_7)($rb_8)), (((($dictHeytingAlgebraRecord_2)->{'impliesRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($ra_7))($rb_8));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "ffRecord" => (function() use ($dictHeytingAlgebraRecord_2, $dictIsSymbol_0, $ff1_4_0) {
  $__fn = function($v_6, $row_7 = null) use ($dictHeytingAlgebraRecord_2, $dictIsSymbol_0, $ff1_4_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $ff1_4_0, ((($dictHeytingAlgebraRecord_2)->{'ffRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($row_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "notRecord" => (function() use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0) {
  $__fn = function($v_6, $row_7 = null) use ($dictHeytingAlgebraRecord_2, $dictHeytingAlgebra_3, $dictIsSymbol_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $key_8_8 = (($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy());
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet($key_8_8, (($dictHeytingAlgebra_3)->{'not'})(\Record\Unsafe\majRecord_majUnsafe_unsafemajGet($key_8_8, $row_7)), ((($dictHeytingAlgebraRecord_2)->{'notRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($row_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "ttRecord" => (function() use ($dictHeytingAlgebraRecord_2, $dictIsSymbol_0, $tt1_5_1) {
  $__fn = function($v_6, $row_7 = null) use ($dictHeytingAlgebraRecord_2, $dictIsSymbol_0, $tt1_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Record\Unsafe\majRecord_majUnsafe_unsafemajSet((($dictIsSymbol_0)->{'reflectSymbol'})(new \Type\Proxy\Type_Proxy_Proxy()), $tt1_5_1, ((($dictHeytingAlgebraRecord_2)->{'ttRecord'})(new \Type\Proxy\Type_Proxy_Proxy()))($row_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_HeytingAlgebra_heytingAlgebraRecordCons'] = __NAMESPACE__ . '\\majData_majHeytingmajAlgebra_heytingmajAlgebramajRecordmajCons';

