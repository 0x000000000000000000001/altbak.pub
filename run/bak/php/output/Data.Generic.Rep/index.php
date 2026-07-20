<?php

namespace Data\Generic\Rep;

// ALL IMPORTS: Data.Generic.Rep, Data.Semigroup, Data.Show, Data.Symbol, Data.Void, Prim, Type.Proxy
// TO REQUIRE: Data.Generic.Rep, Data.Semigroup, Data.Show, Data.Symbol, Data.Void, Type.Proxy
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Symbol/index.php';
require_once __DIR__ . '/../Data.Void/index.php';
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
\PhpursThunks::$thunks['Data_Generic_Rep_Inl'] = function() { $v = function($value0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Inl", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_Inr'] = function() { $v = function($value0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Inr", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_Product'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Product", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_NoArguments'] = function() { $v = ($GLOBALS['__phpurs_data0_NoArguments'] ??= new Phpurs_Data0("NoArguments")); return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_Constructor'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_Constructor"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_Argument'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_Argument"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_to'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_to"), recVars=[];
  $__res = ($dict_0)->to;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_showSum'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_showSum"), recVars=[];
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($v_2) && (($v_2)->tag === "Inl"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Inl "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->show)(($v_2)->value0)))(")"));
} else {
if ((is_object($v_2) && (($v_2)->tag === "Inr"))) {
$__t0 = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Inr "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow1_1)->show)(($v_2)->value0)))(")"));
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
\PhpursThunks::$thunks['Data_Generic_Rep_showProduct'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_showProduct"), recVars=[];
  $__res = (object)["show" => function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Product "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->show)(($v_2)->value0)))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow1_1)->show)(($v_2)->value1)))(")"))));
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
\PhpursThunks::$thunks['Data_Generic_Rep_showNoArguments'] = function() { $v = (object)["show" => function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = "NoArguments";
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_showConstructor'] = function() { $v = (function() {
  $__fn = function($dictIsSymbol_0, $dictShow_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_showConstructor"), recVars=[];
  $__res = (object)["show" => function($v_2) use ($dictIsSymbol_0, $dictShow_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Constructor @"))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($GLOBALS['Data_Show_showStringImpl'] ?? \PhpursThunks::eval('Data_Show_showStringImpl')))((($dictIsSymbol_0)->reflectSymbol)(new Phpurs_Data0("Proxy")))))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))(" "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_1)->show)($v_2)))(")"))));
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
\PhpursThunks::$thunks['Data_Generic_Rep_showArgument'] = function() { $v = function($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_showArgument"), recVars=[];
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Argument "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((($dictShow_0)->show)($v_1)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_repOf'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_repOf"), recVars=[];
  $__res = new Phpurs_Data0("Proxy");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Generic_Rep_from'] = function() { $v = function($dict_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Generic_Rep_from"), recVars=[];
  $__res = ($dict_0)->from;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
















