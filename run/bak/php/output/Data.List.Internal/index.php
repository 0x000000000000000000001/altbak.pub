<?php

namespace Data\List\Internal;

// ALL IMPORTS: Data.List.Internal, Data.List.Types, Data.Ordering, Prelude, Prim
// TO REQUIRE: Data.List.Internal, Data.List.Types, Data.Ordering, Prelude
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
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
\PhpursThunks::$thunks['Data_List_Internal_Leaf'] = function() { $v = ($GLOBALS['__phpurs_data0_Leaf'] ??= new Phpurs_Data0("Leaf")); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_Two'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("Two", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_Three'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null, $value4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__res = new Phpurs_Data5("Three", $value0, $value1, $value2, $value3, $value4);
  goto __end;;
  __end:
  return $__num > 5 ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_TwoLeft'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("TwoLeft", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_TwoRight'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("TwoRight", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_ThreeLeft'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new Phpurs_Data4("ThreeLeft", $value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_ThreeMiddle'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new Phpurs_Data4("ThreeMiddle", $value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_ThreeRight'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = new Phpurs_Data4("ThreeRight", $value0, $value1, $value2, $value3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_KickUp'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("KickUp", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_fromZipper'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Internal_fromZipper"), recVars=["Data_List_Internal_fromZipper"];
  while (true) {
if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = $v1_1;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "TwoLeft"))) {
$__tco_2 = ($v_0)->value1;
$__tco_3 = new Phpurs_Data3("Two", $v1_1, (($v_0)->value0)->value0, (($v_0)->value0)->value1);
$v_0 = $__tco_2;
$v1_1 = $__tco_3;
continue ;
$__t1 = null;
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "TwoRight"))) {
$__tco_4 = ($v_0)->value1;
$__tco_5 = new Phpurs_Data3("Two", (($v_0)->value0)->value0, (($v_0)->value0)->value1, $v1_1);
$v_0 = $__tco_4;
$v1_1 = $__tco_5;
continue ;
$__t1 = null;
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "ThreeLeft"))) {
$__tco_6 = ($v_0)->value1;
$__tco_7 = new Phpurs_Data5("Three", $v1_1, (($v_0)->value0)->value0, (($v_0)->value0)->value1, (($v_0)->value0)->value2, (($v_0)->value0)->value3);
$v_0 = $__tco_6;
$v1_1 = $__tco_7;
continue ;
$__t1 = null;
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "ThreeMiddle"))) {
$__tco_8 = ($v_0)->value1;
$__tco_9 = new Phpurs_Data5("Three", (($v_0)->value0)->value0, (($v_0)->value0)->value1, $v1_1, (($v_0)->value0)->value2, (($v_0)->value0)->value3);
$v_0 = $__tco_8;
$v1_1 = $__tco_9;
continue ;
$__t1 = null;
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "ThreeRight"))) {
$__tco_10 = ($v_0)->value1;
$__tco_11 = new Phpurs_Data5("Three", (($v_0)->value0)->value0, (($v_0)->value0)->value1, (($v_0)->value0)->value2, (($v_0)->value0)->value3, $v1_1);
$v_0 = $__tco_10;
$v1_1 = $__tco_11;
continue ;
$__t1 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
};
};
$__t0 = $__t1;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_insertAndLookupBy'] = function() { $v = (function() {
  $__fn = function($comp_0, $k_1 = null, $orig_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Internal_insertAndLookupBy"), recVars=[];
  $up_3_0 = null;
  $up_3_0 = (function() use (&$up_3_0) {
  $__fn = function($v_4, $v1_5 = null) use (&$up_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "up_3_0"), recVars=["up_3_0"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
$__t1 = new Phpurs_Data3("Two", ($v1_5)->value0, ($v1_5)->value1, ($v1_5)->value2);
} else {
if ((is_object($v_4) && (($v_4)->tag === "Cons"))) {
if ((is_object(($v_4)->value0) && ((($v_4)->value0)->tag === "TwoLeft"))) {
$__t2 = ((($GLOBALS['Data_List_Internal_fromZipper'] ?? \PhpursThunks::eval('Data_List_Internal_fromZipper')))(($v_4)->value1))(new Phpurs_Data5("Three", ($v1_5)->value0, ($v1_5)->value1, ($v1_5)->value2, (($v_4)->value0)->value0, (($v_4)->value0)->value1));
} else {
if ((is_object(($v_4)->value0) && ((($v_4)->value0)->tag === "TwoRight"))) {
$__t2 = ((($GLOBALS['Data_List_Internal_fromZipper'] ?? \PhpursThunks::eval('Data_List_Internal_fromZipper')))(($v_4)->value1))(new Phpurs_Data5("Three", (($v_4)->value0)->value0, (($v_4)->value0)->value1, ($v1_5)->value0, ($v1_5)->value1, ($v1_5)->value2));
} else {
if ((is_object(($v_4)->value0) && ((($v_4)->value0)->tag === "ThreeLeft"))) {
$__tco_3 = ($v_4)->value1;
$__tco_4 = new Phpurs_Data3("KickUp", new Phpurs_Data3("Two", ($v1_5)->value0, ($v1_5)->value1, ($v1_5)->value2), (($v_4)->value0)->value0, new Phpurs_Data3("Two", (($v_4)->value0)->value1, (($v_4)->value0)->value2, (($v_4)->value0)->value3));
$v_4 = $__tco_3;
$v1_5 = $__tco_4;
continue ;
$__t2 = null;
} else {
if ((is_object(($v_4)->value0) && ((($v_4)->value0)->tag === "ThreeMiddle"))) {
$__tco_5 = ($v_4)->value1;
$__tco_6 = new Phpurs_Data3("KickUp", new Phpurs_Data3("Two", (($v_4)->value0)->value0, (($v_4)->value0)->value1, ($v1_5)->value0), ($v1_5)->value1, new Phpurs_Data3("Two", ($v1_5)->value2, (($v_4)->value0)->value2, (($v_4)->value0)->value3));
$v_4 = $__tco_5;
$v1_5 = $__tco_6;
continue ;
$__t2 = null;
} else {
if ((is_object(($v_4)->value0) && ((($v_4)->value0)->tag === "ThreeRight"))) {
$__tco_7 = ($v_4)->value1;
$__tco_8 = new Phpurs_Data3("KickUp", new Phpurs_Data3("Two", (($v_4)->value0)->value0, (($v_4)->value0)->value1, (($v_4)->value0)->value2), (($v_4)->value0)->value3, new Phpurs_Data3("Two", ($v1_5)->value0, ($v1_5)->value1, ($v1_5)->value2));
$v_4 = $__tco_7;
$v1_5 = $__tco_8;
continue ;
$__t2 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
};
};
$__t1 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $down_4_9 = null;
  $down_4_9 = (function() use ($comp_0, &$down_4_9, $k_1, $orig_2, &$up_3_0) {
  $__fn = function($v_5, $v1_6 = null) use ($comp_0, &$down_4_9, $k_1, $orig_2, &$up_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "down_4_9"), recVars=["up_3_0","down_4_9"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Leaf"))) {
$__t10 = (object)["found" => false, "result" => (($up_3_0)($v_5))(new Phpurs_Data3("KickUp", new Phpurs_Data0("Leaf"), $k_1, new Phpurs_Data0("Leaf")))];
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Two"))) {
$v2_7_11 = (($comp_0)($k_1))(($v1_6)->value1);
if ((is_object($v2_7_11) && (($v2_7_11)->tag === "EQ"))) {
$__t14 = (object)["found" => true, "result" => $orig_2];
} else {
if ((is_object($v2_7_11) && (($v2_7_11)->tag === "LT"))) {
$__tco_15 = new Phpurs_Data2("Cons", new Phpurs_Data2("TwoLeft", ($v1_6)->value1, ($v1_6)->value2), $v_5);
$__tco_16 = ($v1_6)->value0;
$v_5 = $__tco_15;
$v1_6 = $__tco_16;
continue ;
$__t14 = null;
} else {
$__tco_12 = new Phpurs_Data2("Cons", new Phpurs_Data2("TwoRight", ($v1_6)->value0, ($v1_6)->value1), $v_5);
$__tco_13 = ($v1_6)->value2;
$v_5 = $__tco_12;
$v1_6 = $__tco_13;
continue ;
$__t14 = null;
};
};
$__t10 = $__t14;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Three"))) {
$v2_7_17 = (($comp_0)($k_1))(($v1_6)->value1);
if ((is_object($v2_7_17) && (($v2_7_17)->tag === "EQ"))) {
$__t26 = (object)["found" => true, "result" => $orig_2];
} else {
$v3_8_18 = (($comp_0)($k_1))(($v1_6)->value3);
if ((is_object($v3_8_18) && (($v3_8_18)->tag === "EQ"))) {
$__t21 = (object)["found" => true, "result" => $orig_2];
} else {
if ((is_object($v2_7_17) && (($v2_7_17)->tag === "LT"))) {
$__tco_22 = new Phpurs_Data2("Cons", new Phpurs_Data4("ThreeLeft", ($v1_6)->value1, ($v1_6)->value2, ($v1_6)->value3, ($v1_6)->value4), $v_5);
$__tco_23 = ($v1_6)->value0;
$v_5 = $__tco_22;
$v1_6 = $__tco_23;
continue ;
$__t21 = null;
} else {
if (((is_object($v2_7_17) && (($v2_7_17)->tag === "GT")) && (is_object($v3_8_18) && (($v3_8_18)->tag === "LT")))) {
$__tco_24 = new Phpurs_Data2("Cons", new Phpurs_Data4("ThreeMiddle", ($v1_6)->value0, ($v1_6)->value1, ($v1_6)->value3, ($v1_6)->value4), $v_5);
$__tco_25 = ($v1_6)->value2;
$v_5 = $__tco_24;
$v1_6 = $__tco_25;
continue ;
$__t21 = null;
} else {
$__tco_19 = new Phpurs_Data2("Cons", new Phpurs_Data4("ThreeRight", ($v1_6)->value0, ($v1_6)->value1, ($v1_6)->value2, ($v1_6)->value3), $v_5);
$__tco_20 = ($v1_6)->value4;
$v_5 = $__tco_19;
$v1_6 = $__tco_20;
continue ;
$__t21 = null;
};
};
};
$__t26 = $__t21;
};
$__t10 = $__t26;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
};
};
};
$__res = $__t10;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($down_4_9)(new Phpurs_Data0("Nil")))($orig_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Internal_emptySet'] = function() { $v = new Phpurs_Data0("Leaf"); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };














