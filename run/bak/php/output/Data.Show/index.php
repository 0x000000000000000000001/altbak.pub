<?php

namespace Data\Show;

// ALL IMPORTS: Data.Semigroup, Data.Show, Data.Symbol, Data.Unit, Data.Void, Prim, Prim.Row, Prim.RowList, Record.Unsafe, Type.Proxy
// TO REQUIRE: Data.Semigroup, Data.Show, Data.Symbol, Data.Unit, Data.Void, Record.Unsafe, Type.Proxy
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
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
\PhpursThunks::$thunks['Data_Show_showVoid'] = function() { $v = (object)["show" => ($GLOBALS['Data_Void_absurd'] ?? \PhpursThunks::eval('Data_Void_absurd'))]; return $v; };
\PhpursThunks::$thunks['Data_Show_showUnit'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = "unit";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Show_showString'] = function() { $v = (object)["show" => ($GLOBALS['Data_Show_showStringImpl'] ?? \PhpursThunks::eval('Data_Show_showStringImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Show_showRecordFieldsNil'] = function() { $v = (object)["showRecordFields" => (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = "";
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()]; return $v; };
\PhpursThunks::$thunks['Data_Show_showRecordFields'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Show_showRecordFields"), recVars=[];
  $__res = ($dict_0)->showRecordFields;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Show_showRecord'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0, $dollar__unused_1 = null, $dictShowRecordFields_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Show_showRecord"), recVars=[];
  $__res = (object)["show" => function($record_3) use ($dictShowRecordFields_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (("{" . ((($dictShowRecordFields_2)->showRecordFields)(new Phpurs_Data0("Proxy")))($record_3)) . "}");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Show_showProxy'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = "Proxy";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Show_showNumber'] = function() { $v = (object)["show" => ($GLOBALS['Data_Show_showNumberImpl'] ?? \PhpursThunks::eval('Data_Show_showNumberImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Show_showInt'] = function() { $v = (object)["show" => ($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Show_showChar'] = function() { $v = (object)["show" => ($GLOBALS['Data_Show_showCharImpl'] ?? \PhpursThunks::eval('Data_Show_showCharImpl'))]; return $v; };
\PhpursThunks::$thunks['Data_Show_showBoolean'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ($v_0) {
$__t0 = "true";
} else {
$__t0 = "false";
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Show_show'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Show_show"), recVars=[];
  $__res = ($dict_0)->show;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Show_showArray'] = function() { $v = function($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Show_showArray"), recVars=[];
  $__res = (object)["show" => (($GLOBALS['Data_Show_showArrayImpl'] ?? \PhpursThunks::eval('Data_Show_showArrayImpl')))(($dictShow_0)->show)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Show_showRecordFieldsCons'] = function() { $v = (function() {
  $__fn = function($dictIsSymbol_0, $dictShowRecordFields_1 = null, $dictShow_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Show_showRecordFieldsCons"), recVars=[];
  $__res = (object)["showRecordFields" => (function() use ($dictIsSymbol_0, $dictShowRecordFields_1, $dictShow_2) {
  $__fn = function($v_3, $record_4 = null) use ($dictIsSymbol_0, $dictShowRecordFields_1, $dictShow_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $key_5_0 = (($dictIsSymbol_0)->reflectSymbol)(new Phpurs_Data0("Proxy"));
  $__res = (((((" " . $key_5_0) . ": ") . (($dictShow_2)->show)(((($GLOBALS['Record_Unsafe_unsafeGet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeGet')))($key_5_0))($record_4))) . ",") . ((($dictShowRecordFields_1)->showRecordFields)(new Phpurs_Data0("Proxy")))($record_4));
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
\PhpursThunks::$thunks['Data_Show_showRecordFieldsConsNil'] = function() { $v = (function() {
  $__fn = function($dictIsSymbol_0, $dictShow_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Show_showRecordFieldsConsNil"), recVars=[];
  $__res = (object)["showRecordFields" => (function() use ($dictIsSymbol_0, $dictShow_1) {
  $__fn = function($v_2, $record_3 = null) use ($dictIsSymbol_0, $dictShow_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $key_4_0 = (($dictIsSymbol_0)->reflectSymbol)(new Phpurs_Data0("Proxy"));
  $__res = ((((" " . $key_4_0) . ": ") . (($dictShow_1)->show)(((($GLOBALS['Record_Unsafe_unsafeGet'] ?? \PhpursThunks::eval('Record_Unsafe_unsafeGet')))($key_4_0))($record_3))) . " ");
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
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Show = \call_user_func(function() {
  $exports = [];
$showIntImpl = function($i) use (&$showIntImpl) { return (string)$i; };
$showStringImpl = function($s) use (&$showStringImpl) { return $s; };
$showNumberImpl = function($n) use (&$showNumberImpl) { return (string)$n; };
$showCharImpl = function($c) use (&$showCharImpl) { return $c; };
$showArrayImpl = function($f, $xs = null) use (&$showArrayImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$showArrayImpl) {

            return $showArrayImpl(...\array_merge($__args, $more));
        };
    }
    return "[" . implode(",", array_map($f, $xs)) . "]";
};

$exports['showIntImpl'] = $showIntImpl;
$exports['showStringImpl'] = $showStringImpl;
$exports['showNumberImpl'] = $showNumberImpl;
$exports['showCharImpl'] = $showCharImpl;
$exports['showArrayImpl'] = $showArrayImpl;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Show_showArrayImpl'] = function() use (&$ffi_Data_Show) { return $ffi_Data_Show['showArrayImpl']; };
\PhpursThunks::$thunks['Data_Show_showCharImpl'] = function() use (&$ffi_Data_Show) { return $ffi_Data_Show['showCharImpl']; };
\PhpursThunks::$thunks['Data_Show_showIntImpl'] = function() use (&$ffi_Data_Show) { return $ffi_Data_Show['showIntImpl']; };
\PhpursThunks::$thunks['Data_Show_showNumberImpl'] = function() use (&$ffi_Data_Show) { return $ffi_Data_Show['showNumberImpl']; };
\PhpursThunks::$thunks['Data_Show_showStringImpl'] = function() use (&$ffi_Data_Show) { return $ffi_Data_Show['showStringImpl']; };

















