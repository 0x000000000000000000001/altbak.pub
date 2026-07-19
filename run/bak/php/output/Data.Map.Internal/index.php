<?php

namespace Data\Map\Internal;

// ALL IMPORTS: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude, Prim, Prim.TypeError
// TO REQUIRE: Control.Alt, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Plus, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.FoldableWithIndex, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.List, Data.List.Types, Data.Map.Internal, Data.Maybe, Data.Monoid, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.TraversableWithIndex, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.FoldableWithIndex/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Map.Internal/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.TraversableWithIndex/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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
\PhpursThunks::$thunks['Data_Map_Internal_abs'] = function() { $v = ((($GLOBALS['Data_Ord_abs'] ?? \PhpursThunks::eval('Data_Ord_abs')))(($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt'))))(($GLOBALS['Data_Ring_ringInt'] ?? \PhpursThunks::eval('Data_Ring_ringInt'))); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_Leaf'] = function() { $v = ($GLOBALS['__phpurs_data0_Leaf'] ??= new Phpurs_Data0("Leaf")); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_Node'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null, $value3 = null, $value4 = null, $value5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  $__res = new Phpurs_Data6("Node", $value0, $value1, $value2, $value3, $value4, $value5);
  goto __end;;
  __end:
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_IterLeaf'] = function() { $v = ($GLOBALS['__phpurs_data0_IterLeaf'] ??= new Phpurs_Data0("IterLeaf")); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_IterEmit'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("IterEmit", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_IterNode'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("IterNode", $value0, $value1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_IterDone'] = function() { $v = ($GLOBALS['__phpurs_data0_IterDone'] ??= new Phpurs_Data0("IterDone")); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_IterNext'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("IterNext", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_Split'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("Split", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_SplitLast'] = function() { $v = (function() {
  $__fn = function($value0, $value1 = null, $value2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data3("SplitLast", $value0, $value1, $value2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeNode'] = function() { $v = (function() {
  $__fn = function($k_0, $v_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeNode"), recVars=[];
  if ((is_object($l_2) && (($l_2)->tag === "Leaf"))) {
if ((is_object($r_3) && (($r_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_0, $v_1, $l_2, $r_3);
} else {
if ((is_object($r_3) && (($r_3)->tag === "Node"))) {
$__t1 = new Phpurs_Data6("Node", (1 + ($r_3)->value0), (1 + ($r_3)->value1), $k_0, $v_1, $l_2, $r_3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__t0 = $__t1;
} else {
if ((is_object($l_2) && (($l_2)->tag === "Node"))) {
if ((is_object($r_3) && (($r_3)->tag === "Leaf"))) {
$__t2 = new Phpurs_Data6("Node", (1 + ($l_2)->value0), (1 + ($l_2)->value1), $k_0, $v_1, $l_2, $r_3);
} else {
if ((is_object($r_3) && (($r_3)->tag === "Node"))) {
if ((($l_2)->value0 > ($r_3)->value0)) {
$__t3 = (1 + ($l_2)->value0);
} else {
$__t3 = (1 + ($r_3)->value0);
};
$__t2 = new Phpurs_Data6("Node", $__t3, ((1 + ($l_2)->value1) + ($r_3)->value1), $k_0, $v_1, $l_2, $r_3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__t0 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_toMapIter'] = function() { $v = function($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_toMapIter"), recVars=[];
  $__res = new Phpurs_Data2("IterNode", $a_0, new Phpurs_Data0("IterLeaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepWith'] = function() { $v = (function() {
  $__fn = function($f_0, $next_1 = null, $done_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_stepWith"), recVars=[];
  $go_3_0 = null;
  $go_3_0 = function($v_4) use ($done_2, $f_0, &$go_3_0, $next_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "IterLeaf"))) {
$__t1 = ($done_2)(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
} else {
if ((is_object($v_4) && (($v_4)->tag === "IterEmit"))) {
$__t1 = ($next_1)(($v_4)->value0, ($v_4)->value1, ($v_4)->value2);
} else {
if ((is_object($v_4) && (($v_4)->tag === "IterNode"))) {
$__tco_2 = (($f_0)(($v_4)->value1))(($v_4)->value0);
$v_4 = $__tco_2;
continue ;
$__t1 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_size'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_size"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Leaf"))) {
$__t0 = 0;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Node"))) {
$__t0 = ($v_0)->value1;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_singleton'] = function() { $v = (function() {
  $__fn = function($k_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_singleton"), recVars=[];
  $__res = new Phpurs_Data6("Node", 1, 1, $k_0, $v_1, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeBalancedNode'] = function() { $v = (function() {
  $__fn = function($k_0, $v_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeBalancedNode"), recVars=[];
  if ((is_object($l_2) && (($l_2)->tag === "Leaf"))) {
if ((is_object($r_3) && (($r_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_0, $v_1, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
} else {
if (((is_object($r_3) && (($r_3)->tag === "Node")) && (($r_3)->value0 > 1))) {
if ((function() use ($r_3, &$__fn) {
if ((is_object(($r_3)->value5) && ((($r_3)->value5)->tag === "Leaf"))) {
$__t3 = ((($r_3)->value4)->value0 > 0);
} else {
if ((is_object(($r_3)->value5) && ((($r_3)->value5)->tag === "Node"))) {
$__t3 = ((($r_3)->value4)->value0 > (($r_3)->value5)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
return ((is_object(($r_3)->value4) && ((($r_3)->value4)->tag === "Node")) && $__t3);
})()) {
$__t2 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))((($r_3)->value4)->value2, (($r_3)->value4)->value3, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, $l_2, (($r_3)->value4)->value4), (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($r_3)->value2, ($r_3)->value3, (($r_3)->value4)->value5, ($r_3)->value5));
} else {
$__t2 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($r_3)->value2, ($r_3)->value3, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, $l_2, ($r_3)->value4), ($r_3)->value5);
};
$__t1 = $__t2;
} else {
$__t1 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, $l_2, $r_3);
};
};
$__t0 = $__t1;
} else {
if ((is_object($l_2) && (($l_2)->tag === "Node"))) {
if ((is_object($r_3) && (($r_3)->tag === "Node"))) {
if ((($r_3)->value0 > (($l_2)->value0 + 1))) {
if ((function() use ($r_3, &$__fn) {
if ((is_object(($r_3)->value5) && ((($r_3)->value5)->tag === "Leaf"))) {
$__t7 = ((($r_3)->value4)->value0 > 0);
} else {
if ((is_object(($r_3)->value5) && ((($r_3)->value5)->tag === "Node"))) {
$__t7 = ((($r_3)->value4)->value0 > (($r_3)->value5)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
return ((is_object(($r_3)->value4) && ((($r_3)->value4)->tag === "Node")) && $__t7);
})()) {
$__t6 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))((($r_3)->value4)->value2, (($r_3)->value4)->value3, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, $l_2, (($r_3)->value4)->value4), (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($r_3)->value2, ($r_3)->value3, (($r_3)->value4)->value5, ($r_3)->value5));
} else {
$__t6 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($r_3)->value2, ($r_3)->value3, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, $l_2, ($r_3)->value4), ($r_3)->value5);
};
$__t5 = $__t6;
} else {
if ((($l_2)->value0 > (($r_3)->value0 + 1))) {
if ((function() use ($l_2, &$__fn) {
if ((is_object(($l_2)->value4) && ((($l_2)->value4)->tag === "Leaf"))) {
$__t9 = (0 <= (($l_2)->value5)->value0);
} else {
if ((is_object(($l_2)->value4) && ((($l_2)->value4)->tag === "Node"))) {
$__t9 = ((($l_2)->value4)->value0 <= (($l_2)->value5)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
};
};
return ((is_object(($l_2)->value5) && ((($l_2)->value5)->tag === "Node")) && $__t9);
})()) {
$__t8 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))((($l_2)->value5)->value2, (($l_2)->value5)->value3, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($l_2)->value2, ($l_2)->value3, ($l_2)->value4, (($l_2)->value5)->value4), (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, (($l_2)->value5)->value5, $r_3));
} else {
$__t8 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($l_2)->value2, ($l_2)->value3, ($l_2)->value4, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, ($l_2)->value5, $r_3));
};
$__t5 = $__t8;
} else {
$__t5 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, $l_2, $r_3);
};
};
$__t4 = $__t5;
} else {
if (((is_object($r_3) && (($r_3)->tag === "Leaf")) && (($l_2)->value0 > 1))) {
if ((function() use ($l_2, &$__fn) {
if ((is_object(($l_2)->value4) && ((($l_2)->value4)->tag === "Leaf"))) {
$__t11 = (0 <= (($l_2)->value5)->value0);
} else {
if ((is_object(($l_2)->value4) && ((($l_2)->value4)->tag === "Node"))) {
$__t11 = ((($l_2)->value4)->value0 <= (($l_2)->value5)->value0);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
};
};
return ((is_object(($l_2)->value5) && ((($l_2)->value5)->tag === "Node")) && $__t11);
})()) {
$__t10 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))((($l_2)->value5)->value2, (($l_2)->value5)->value3, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($l_2)->value2, ($l_2)->value3, ($l_2)->value4, (($l_2)->value5)->value4), (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, (($l_2)->value5)->value5, $r_3));
} else {
$__t10 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))(($l_2)->value2, ($l_2)->value3, ($l_2)->value4, (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, ($l_2)->value5, $r_3));
};
$__t4 = $__t10;
} else {
$__t4 = (($GLOBALS['Data_Map_Internal_unsafeNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeNode')))($k_0, $v_1, $l_2, $r_3);
};
};
$__t0 = $__t4;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeSplit'] = function() { $v = (function() {
  $__fn = function($comp_0, $k_1 = null, $m_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeSplit"), recVars=["Data_Map_Internal_unsafeSplit"];
  while (true) {
if ((is_object($m_2) && (($m_2)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data3("Split", new Phpurs_Data0("Nothing"), new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
} else {
if ((is_object($m_2) && (($m_2)->tag === "Node"))) {
$v_3_1 = (($comp_0)($k_1))(($m_2)->value2);
if ((is_object($v_3_1) && (($v_3_1)->tag === "LT"))) {
$v1_4_3 = (($GLOBALS['Data_Map_Internal_unsafeSplit'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplit')))($comp_0, $k_1, ($m_2)->value4);
$__t2 = new Phpurs_Data3("Split", ($v1_4_3)->value0, ($v1_4_3)->value1, (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($m_2)->value2, ($m_2)->value3, ($v1_4_3)->value2, ($m_2)->value5));
} else {
if ((is_object($v_3_1) && (($v_3_1)->tag === "GT"))) {
$v1_4_4 = (($GLOBALS['Data_Map_Internal_unsafeSplit'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplit')))($comp_0, $k_1, ($m_2)->value5);
$__t2 = new Phpurs_Data3("Split", ($v1_4_4)->value0, (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($m_2)->value2, ($m_2)->value3, ($m_2)->value4, ($v1_4_4)->value1), ($v1_4_4)->value2);
} else {
if ((is_object($v_3_1) && (($v_3_1)->tag === "EQ"))) {
$__t2 = new Phpurs_Data3("Split", new Phpurs_Data1("Just", ($m_2)->value3), ($m_2)->value4, ($m_2)->value5);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
$__t0 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeSplitLast'] = function() { $v = (function() {
  $__fn = function($k_0, $v_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeSplitLast"), recVars=["Data_Map_Internal_unsafeSplitLast"];
  while (true) {
if ((is_object($r_3) && (($r_3)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data3("SplitLast", $k_0, $v_1, $l_2);
} else {
if ((is_object($r_3) && (($r_3)->tag === "Node"))) {
$v1_4_1 = (($GLOBALS['Data_Map_Internal_unsafeSplitLast'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplitLast')))(($r_3)->value2, ($r_3)->value3, ($r_3)->value4, ($r_3)->value5);
$__t0 = new Phpurs_Data3("SplitLast", ($v1_4_1)->value0, ($v1_4_1)->value1, (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))($k_0, $v_1, $l_2, ($v1_4_1)->value2));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeJoinNodes'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeJoinNodes"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Leaf"))) {
$__t0 = $v1_1;
} else {
if ((is_object($v_0) && (($v_0)->tag === "Node"))) {
$v2_2_1 = (($GLOBALS['Data_Map_Internal_unsafeSplitLast'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplitLast')))(($v_0)->value2, ($v_0)->value3, ($v_0)->value4, ($v_0)->value5);
$__t0 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v2_2_1)->value0, ($v2_2_1)->value1, ($v2_2_1)->value2, $v1_1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeDifference'] = function() { $v = (function() {
  $__fn = function($comp_0, $l_1 = null, $r_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeDifference"), recVars=["Data_Map_Internal_unsafeDifference"];
  while (true) {
if ((is_object($l_1) && (($l_1)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($r_2) && (($r_2)->tag === "Leaf"))) {
$__t0 = $l_1;
} else {
if ((is_object($r_2) && (($r_2)->tag === "Node"))) {
$v_3_1 = (($GLOBALS['Data_Map_Internal_unsafeSplit'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplit')))($comp_0, ($r_2)->value2, $l_1);
$__t0 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))((($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference')))($comp_0, ($v_3_1)->value1, ($r_2)->value4), (($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference')))($comp_0, ($v_3_1)->value2, ($r_2)->value5));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeIntersectionWith'] = function() { $v = (function() {
  $__fn = function($comp_0, $app_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeIntersectionWith"), recVars=["Data_Map_Internal_unsafeIntersectionWith"];
  while (true) {
if ((is_object($l_2) && (($l_2)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($r_3) && (($r_3)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($r_3) && (($r_3)->tag === "Node"))) {
$v_4_1 = (($GLOBALS['Data_Map_Internal_unsafeSplit'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplit')))($comp_0, ($r_3)->value2, $l_2);
$l__prime___5_2 = (($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith')))($comp_0, $app_1, ($v_4_1)->value1, ($r_3)->value4);
$r__prime___6_3 = (($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith')))($comp_0, $app_1, ($v_4_1)->value2, ($r_3)->value5);
if ((is_object(($v_4_1)->value0) && ((($v_4_1)->value0)->tag === "Just"))) {
$__t4 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($r_3)->value2, (($app_1)((($v_4_1)->value0)->value0))(($r_3)->value3), $l__prime___5_2, $r__prime___6_3);
} else {
if ((is_object(($v_4_1)->value0) && ((($v_4_1)->value0)->tag === "Nothing"))) {
$__t4 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))($l__prime___5_2, $r__prime___6_3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__t0 = $__t4;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unsafeUnionWith'] = function() { $v = (function() {
  $__fn = function($comp_0, $app_1 = null, $l_2 = null, $r_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unsafeUnionWith"), recVars=["Data_Map_Internal_unsafeUnionWith"];
  while (true) {
if ((is_object($l_2) && (($l_2)->tag === "Leaf"))) {
$__t0 = $r_3;
} else {
if ((is_object($r_3) && (($r_3)->tag === "Leaf"))) {
$__t0 = $l_2;
} else {
if ((is_object($r_3) && (($r_3)->tag === "Node"))) {
$v_4_1 = (($GLOBALS['Data_Map_Internal_unsafeSplit'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplit')))($comp_0, ($r_3)->value2, $l_2);
$l__prime___5_2 = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($comp_0, $app_1, ($v_4_1)->value1, ($r_3)->value4);
$r__prime___6_3 = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($comp_0, $app_1, ($v_4_1)->value2, ($r_3)->value5);
if ((is_object(($v_4_1)->value0) && ((($v_4_1)->value0)->tag === "Just"))) {
$__t4 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($r_3)->value2, (($app_1)((($v_4_1)->value0)->value0))(($r_3)->value3), $l__prime___5_2, $r__prime___6_3);
} else {
if ((is_object(($v_4_1)->value0) && ((($v_4_1)->value0)->tag === "Nothing"))) {
$__t4 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($r_3)->value2, ($r_3)->value3, $l__prime___5_2, $r__prime___6_3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__t0 = $__t4;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_unionWith'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unionWith"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($app_2, $m1_3 = null, $m2_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, $app_2, $m1_3, $m2_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_union'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_union"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_2, $m2_3);
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
\PhpursThunks::$thunks['Data_Map_Internal_update'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null, $k_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_update"), recVars=[];
  $go_3_0 = null;
  $go_3_0 = function($v_4) use ($dictOrd_0, $f_1, &$go_3_0, $k_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($v_4) && (($v_4)->tag === "Node"))) {
$v1_5_2 = ((($dictOrd_0)->compare)($k_2))(($v_4)->value2);
if ((is_object($v1_5_2) && (($v1_5_2)->tag === "LT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v_4)->value2, ($v_4)->value3, ($go_3_0)(($v_4)->value4), ($v_4)->value5);
} else {
if ((is_object($v1_5_2) && (($v1_5_2)->tag === "GT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v_4)->value2, ($v_4)->value3, ($v_4)->value4, ($go_3_0)(($v_4)->value5));
} else {
if ((is_object($v1_5_2) && (($v1_5_2)->tag === "EQ"))) {
$v2_6_4 = ($f_1)(($v_4)->value3);
if ((is_object($v2_6_4) && (($v2_6_4)->tag === "Nothing"))) {
$__t5 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))(($v_4)->value4, ($v_4)->value5);
} else {
if ((is_object($v2_6_4) && (($v2_6_4)->tag === "Just"))) {
$__t5 = new Phpurs_Data6("Node", ($v_4)->value0, ($v_4)->value1, ($v_4)->value2, ($v2_6_4)->value0, ($v_4)->value4, ($v_4)->value5);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__t3 = $__t5;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_showTree'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_showTree"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use ($dictShow1_1, $dictShow_0, &$go_2_0) {
  $__fn = function($ind_3, $v_4 = null) use ($dictShow1_1, $dictShow_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Leaf"))) {
$__t1 = ($ind_3 . "Leaf");
} else {
if ((is_object($v_4) && (($v_4)->tag === "Node"))) {
$__t1 = (((((((((($ind_3 . "[") . (($GLOBALS['Data_Show_showIntImpl'] ?? \PhpursThunks::eval('Data_Show_showIntImpl')))(($v_4)->value0)) . "] ") . (($dictShow_0)->show)(($v_4)->value2)) . " => ") . (($dictShow1_1)->show)(($v_4)->value3)) . "
") . (($go_2_0)(($ind_3 . "    ")))(($v_4)->value4)) . "
") . (($go_2_0)(($ind_3 . "    ")))(($v_4)->value5));
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
  $__res = ($go_2_0)("");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_semigroupMap'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_semigroupMap"), recVars=[];
  $compare_2_0 = ($dictOrd_1)->compare;
  $__res = function($dictSemigroup_3) use ($compare_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_1 = ($dictSemigroup_3)->append;
  $__res = (object)["append" => (function() use ($__local_var_4_1, $compare_2_0) {
  $__fn = function($m1_5, $m2_6 = null) use ($__local_var_4_1, $compare_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_2_0, $__local_var_4_1, $m1_5, $m2_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
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
\PhpursThunks::$thunks['Data_Map_Internal_pop'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_pop"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($k_2, $m_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_4_1 = (($GLOBALS['Data_Map_Internal_unsafeSplit'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplit')))($compare_1_0, $k_2, $m_3);
  if ((is_object(($v_4_1)->value0) && ((($v_4_1)->value0)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", (($v_4_1)->value0)->value0, (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))(($v_4_1)->value1, ($v_4_1)->value2)));
} else {
$__t2 = new Phpurs_Data0("Nothing");
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
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_member'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_member"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = false;
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$__tco_4 = ($v_3)->value4;
$v_3 = $__tco_4;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$__tco_5 = ($v_3)->value5;
$v_3 = $__tco_5;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = true;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_mapMaybeWithKey'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_mapMaybeWithKey"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($f_1, &$go_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v2_4_2 = (($f_1)(($v_3)->value2))(($v_3)->value3);
if ((is_object($v2_4_2) && (($v2_4_2)->tag === "Just"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v_3)->value2, ($v2_4_2)->value0, ($go_2_0)(($v_3)->value4), ($go_2_0)(($v_3)->value5));
} else {
if ((is_object($v2_4_2) && (($v2_4_2)->tag === "Nothing"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))(($go_2_0)(($v_3)->value4), ($go_2_0)(($v_3)->value5));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_mapMaybe'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_mapMaybe"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Map_Internal_mapMaybeWithKey'] ?? \PhpursThunks::eval('Data_Map_Internal_mapMaybeWithKey')))($dictOrd_0)))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_lookupLE'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_lookupLE"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$__tco_4 = ($v_3)->value4;
$v_3 = $__tco_4;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$v2_5_5 = ($go_2_0)(($v_3)->value5);
if ((is_object($v2_5_5) && (($v2_5_5)->tag === "Nothing"))) {
$__t6 = new Phpurs_Data1("Just", (object)["key" => ($v_3)->value2, "value" => ($v_3)->value3]);
} else {
$__t6 = $v2_5_5;
};
$__t3 = $__t6;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = new Phpurs_Data1("Just", (object)["key" => ($v_3)->value2, "value" => ($v_3)->value3]);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_lookupGE'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_lookupGE"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$v2_5_4 = ($go_2_0)(($v_3)->value4);
if ((is_object($v2_5_4) && (($v2_5_4)->tag === "Nothing"))) {
$__t5 = new Phpurs_Data1("Just", (object)["key" => ($v_3)->value2, "value" => ($v_3)->value3]);
} else {
$__t5 = $v2_5_4;
};
$__t3 = $__t5;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$__tco_6 = ($v_3)->value5;
$v_3 = $__tco_6;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = new Phpurs_Data1("Just", (object)["key" => ($v_3)->value2, "value" => ($v_3)->value3]);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_lookup'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_lookup"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$__tco_4 = ($v_3)->value4;
$v_3 = $__tco_4;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$__tco_5 = ($v_3)->value5;
$v_3 = $__tco_5;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = new Phpurs_Data1("Just", ($v_3)->value3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_iterMapU'] = function() { $v = (function() {
  $__fn = function($iter_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_iterMapU"), recVars=[];
  if ((is_object($v_1) && (($v_1)->tag === "Leaf"))) {
$__t0 = $iter_0;
} else {
if ((is_object($v_1) && (($v_1)->tag === "Node"))) {
if ((is_object(($v_1)->value4) && ((($v_1)->value4)->tag === "Leaf"))) {
if ((is_object(($v_1)->value5) && ((($v_1)->value5)->tag === "Leaf"))) {
$__t2 = new Phpurs_Data3("IterEmit", ($v_1)->value2, ($v_1)->value3, $iter_0);
} else {
$__t2 = new Phpurs_Data3("IterEmit", ($v_1)->value2, ($v_1)->value3, new Phpurs_Data2("IterNode", ($v_1)->value5, $iter_0));
};
$__t1 = $__t2;
} else {
if ((is_object(($v_1)->value5) && ((($v_1)->value5)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data3("IterEmit", ($v_1)->value2, ($v_1)->value3, new Phpurs_Data2("IterNode", ($v_1)->value4, $iter_0));
} else {
$__t1 = new Phpurs_Data3("IterEmit", ($v_1)->value2, ($v_1)->value3, new Phpurs_Data2("IterNode", ($v_1)->value4, new Phpurs_Data2("IterNode", ($v_1)->value5, $iter_0)));
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
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepUnorderedCps'] = function() { $v = (($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapU'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapU'))); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepUnfoldrUnordered'] = function() { $v = (((($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapU'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapU'))))((function() {
  $__fn = function($k_0, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", $k_0, $v_1), $next_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_toUnfoldableUnordered'] = function() { $v = function($dictUnfoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_toUnfoldableUnordered"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($dictUnfoldable_0)->unfoldr)(($GLOBALS['Data_Map_Internal_stepUnfoldrUnordered'] ?? \PhpursThunks::eval('Data_Map_Internal_stepUnfoldrUnordered')))))(($GLOBALS['Data_Map_Internal_toMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_toMapIter')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepUnordered'] = function() { $v = (((($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapU'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapU'))))((function() {
  $__fn = function($k_0, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data3("IterNext", $k_0, $v_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("IterDone");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_iterMapR'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($iter_1, $v_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = $iter_1;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
if ((is_object(($v_2)->value5) && ((($v_2)->value5)->tag === "Leaf"))) {
$__tco_5 = new Phpurs_Data3("IterEmit", ($v_2)->value2, ($v_2)->value3, $iter_1);
$__tco_6 = ($v_2)->value4;
$iter_1 = $__tco_5;
$v_2 = $__tco_6;
continue ;
$__t4 = null;
} else {
$__tco_2 = new Phpurs_Data3("IterEmit", ($v_2)->value2, ($v_2)->value3, new Phpurs_Data2("IterNode", ($v_2)->value4, $iter_1));
$__tco_3 = ($v_2)->value5;
$iter_1 = $__tco_2;
$v_2 = $__tco_3;
continue ;
$__t4 = null;
};
$__t1 = $__t4;
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
return $go_0_0;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepDescCps'] = function() { $v = (($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapR'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapR'))); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepDesc'] = function() { $v = (((($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapR'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapR'))))((function() {
  $__fn = function($k_0, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data3("IterNext", $k_0, $v_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("IterDone");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_iterMapL'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($iter_1, $v_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = $iter_1;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
if ((is_object(($v_2)->value5) && ((($v_2)->value5)->tag === "Leaf"))) {
$__tco_5 = new Phpurs_Data3("IterEmit", ($v_2)->value2, ($v_2)->value3, $iter_1);
$__tco_6 = ($v_2)->value4;
$iter_1 = $__tco_5;
$v_2 = $__tco_6;
continue ;
$__t4 = null;
} else {
$__tco_2 = new Phpurs_Data3("IterEmit", ($v_2)->value2, ($v_2)->value3, new Phpurs_Data2("IterNode", ($v_2)->value5, $iter_1));
$__tco_3 = ($v_2)->value4;
$iter_1 = $__tco_2;
$v_2 = $__tco_3;
continue ;
$__t4 = null;
};
$__t1 = $__t4;
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
return $go_0_0;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepAscCps'] = function() { $v = (($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapL'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapL'))); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_stepAsc'] = function() { $v = (((($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapL'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapL'))))((function() {
  $__fn = function($k_0, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data3("IterNext", $k_0, $v_1, $next_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("IterDone");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_eqMapIter'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_eqMapIter"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use ($dictEq1_1, $dictEq_0, &$go_2_0) {
  $__fn = function($a_3, $b_4 = null) use ($dictEq1_1, $dictEq_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
$v_5_1 = (($GLOBALS['Data_Map_Internal_stepAsc'] ?? \PhpursThunks::eval('Data_Map_Internal_stepAsc')))($a_3);
if ((is_object($v_5_1) && (($v_5_1)->tag === "IterNext"))) {
$v2_6_3 = (($GLOBALS['Data_Map_Internal_stepAsc'] ?? \PhpursThunks::eval('Data_Map_Internal_stepAsc')))($b_4);
$__t2 = ((is_object($v2_6_3) && (($v2_6_3)->tag === "IterNext")) && ((((($dictEq_0)->eq)(($v_5_1)->value0))(($v2_6_3)->value0) && ((($dictEq1_1)->eq)(($v_5_1)->value1))(($v2_6_3)->value1)) && (($go_2_0)(($v_5_1)->value2))(($v2_6_3)->value2)));
} else {
if ((is_object($v_5_1) && (($v_5_1)->tag === "IterDone"))) {
$__t2 = true;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (object)["eq" => $go_2_0];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_ordMapIter'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_ordMapIter"), recVars=[];
  $eqMapIter1_1_0 = (($GLOBALS['Data_Map_Internal_eqMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMapIter')))((($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictOrd1_2) use ($dictOrd_0, $eqMapIter1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eqMapIter2_3_1 = ($eqMapIter1_1_0)((($dictOrd1_2)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $go_4_2 = null;
  $go_4_2 = (function() use ($dictOrd1_2, $dictOrd_0, &$go_4_2) {
  $__fn = function($a_5, $b_6 = null) use ($dictOrd1_2, $dictOrd_0, &$go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_2"), recVars=["go_4_2"];
  while (true) {
$v_7_3 = (($GLOBALS['Data_Map_Internal_stepAsc'] ?? \PhpursThunks::eval('Data_Map_Internal_stepAsc')))($b_6);
$v1_8_4 = (($GLOBALS['Data_Map_Internal_stepAsc'] ?? \PhpursThunks::eval('Data_Map_Internal_stepAsc')))($a_5);
if ((is_object($v1_8_4) && (($v1_8_4)->tag === "IterNext"))) {
if ((is_object($v_7_3) && (($v_7_3)->tag === "IterNext"))) {
$v3_9_7 = ((($dictOrd_0)->compare)(($v1_8_4)->value0))(($v_7_3)->value0);
if ((is_object($v3_9_7) && (($v3_9_7)->tag === "EQ"))) {
$v4_10_9 = ((($dictOrd1_2)->compare)(($v1_8_4)->value1))(($v_7_3)->value1);
if ((is_object($v4_10_9) && (($v4_10_9)->tag === "EQ"))) {
$__tco_11 = ($v1_8_4)->value2;
$__tco_12 = ($v_7_3)->value2;
$a_5 = $__tco_11;
$b_6 = $__tco_12;
continue ;
$__t10 = null;
} else {
$__t10 = $v4_10_9;
};
$__t8 = $__t10;
} else {
$__t8 = $v3_9_7;
};
$__t6 = $__t8;
} else {
if ((is_object($v_7_3) && (($v_7_3)->tag === "IterDone"))) {
$__t6 = new Phpurs_Data0("GT");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__t5 = $__t6;
} else {
if ((is_object($v1_8_4) && (($v1_8_4)->tag === "IterDone"))) {
if ((is_object($v_7_3) && (($v_7_3)->tag === "IterDone"))) {
$__t13 = new Phpurs_Data0("EQ");
} else {
$__t13 = new Phpurs_Data0("LT");
};
$__t5 = $__t13;
} else {
if ((is_object($v_7_3) && (($v_7_3)->tag === "IterDone"))) {
$__t5 = new Phpurs_Data0("GT");
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (object)["compare" => $go_4_2, "Eq0" => function($dollar__unused_4) use ($eqMapIter2_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqMapIter2_3_1;
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
\PhpursThunks::$thunks['Data_Map_Internal_stepUnfoldr'] = function() { $v = (((($GLOBALS['Data_Map_Internal_stepWith'] ?? \PhpursThunks::eval('Data_Map_Internal_stepWith')))(($GLOBALS['Data_Map_Internal_iterMapL'] ?? \PhpursThunks::eval('Data_Map_Internal_iterMapL'))))((function() {
  $__fn = function($k_0, $v_1 = null, $next_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", new Phpurs_Data2("Tuple", $k_0, $v_1), $next_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_toUnfoldable'] = function() { $v = function($dictUnfoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_toUnfoldable"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($dictUnfoldable_0)->unfoldr)(($GLOBALS['Data_Map_Internal_stepUnfoldr'] ?? \PhpursThunks::eval('Data_Map_Internal_stepUnfoldr')))))(($GLOBALS['Data_Map_Internal_toMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_toMapIter')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_toUnfoldable1'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Unfoldable_unfoldableArray'] ?? \PhpursThunks::eval('Data_Unfoldable_unfoldableArray')))->unfoldr)(($GLOBALS['Data_Map_Internal_stepUnfoldr'] ?? \PhpursThunks::eval('Data_Map_Internal_stepUnfoldr')))))(($GLOBALS['Data_Map_Internal_toMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_toMapIter'))); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_showMap'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $dictShow1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_showMap"), recVars=[];
  $show1_2_0 = (($GLOBALS['Data_Show_showArrayImpl'] ?? \PhpursThunks::eval('Data_Show_showArrayImpl')))(function($v_2) use ($dictShow1_1, $dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((("(Tuple " . (($dictShow_0)->show)(($v_2)->value0)) . " ") . (($dictShow1_1)->show)(($v_2)->value1)) . ")");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $__res = (object)["show" => function($as_3) use ($show1_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (("(fromFoldable " . ($show1_2_0)((($GLOBALS['Data_Map_Internal_toUnfoldable1'] ?? \PhpursThunks::eval('Data_Map_Internal_toUnfoldable1')))($as_3))) . ")");
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
\PhpursThunks::$thunks['Data_Map_Internal_isSubmap'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictEq_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_isSubmap"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use ($dictEq_1, $dictOrd_0, &$go_2_0) {
  $__fn = function($m1_3, $m2_4 = null) use ($dictEq_1, $dictOrd_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($m1_3) && (($m1_3)->tag === "Leaf"))) {
$__t1 = true;
} else {
if ((is_object($m1_3) && (($m1_3)->tag === "Node"))) {
$__local_var_5_2 = ($m1_3)->value2;
$go_6_3 = null;
$go_6_3 = function($v_7) use ($__local_var_5_2, $dictOrd_0, &$go_6_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_3"), recVars=["go_2_0","go_6_3"];
  while (true) {
if ((is_object($v_7) && (($v_7)->tag === "Leaf"))) {
$__t4 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_7) && (($v_7)->tag === "Node"))) {
$v1_8_5 = ((($dictOrd_0)->compare)($__local_var_5_2))(($v_7)->value2);
if ((is_object($v1_8_5) && (($v1_8_5)->tag === "LT"))) {
$__tco_7 = ($v_7)->value4;
$v_7 = $__tco_7;
continue ;
$__t6 = null;
} else {
if ((is_object($v1_8_5) && (($v1_8_5)->tag === "GT"))) {
$__tco_8 = ($v_7)->value5;
$v_7 = $__tco_8;
continue ;
$__t6 = null;
} else {
if ((is_object($v1_8_5) && (($v1_8_5)->tag === "EQ"))) {
$__t6 = new Phpurs_Data1("Just", ($v_7)->value3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
};
$__t4 = $__t6;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__res = $__t4;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$v1_7_9 = ($go_6_3)($m2_4);
if ((is_object($v1_7_9) && (($v1_7_9)->tag === "Nothing"))) {
$__t10 = false;
} else {
if ((is_object($v1_7_9) && (($v1_7_9)->tag === "Just"))) {
$__t10 = (((($dictEq_1)->eq)(($m1_3)->value3))(($v1_7_9)->value0) && ((($go_2_0)(($m1_3)->value4))($m2_4) && (($go_2_0)(($m1_3)->value5))($m2_4)));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
};
};
$__t1 = $__t10;
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
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_isEmpty'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_isEmpty"), recVars=[];
  $__res = (is_object($v_0) && (($v_0)->tag === "Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_intersectionWith'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_intersectionWith"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($app_2, $m1_3 = null, $m2_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith')))($compare_1_0, $app_2, $m1_3, $m2_4);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_intersection'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_intersection"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_2, $m2_3);
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
\PhpursThunks::$thunks['Data_Map_Internal_insertWith'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $app_1 = null, $k_2 = null, $v_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_insertWith"), recVars=[];
  $go_4_0 = null;
  $go_4_0 = function($v1_5) use ($app_1, $dictOrd_0, &$go_4_0, $k_2, $v_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_0"), recVars=["go_4_0"];
  while (true) {
if ((is_object($v1_5) && (($v1_5)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_2, $v_3, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Node"))) {
$v2_6_2 = ((($dictOrd_0)->compare)($k_2))(($v1_5)->value2);
if ((is_object($v2_6_2) && (($v2_6_2)->tag === "LT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v1_5)->value2, ($v1_5)->value3, ($go_4_0)(($v1_5)->value4), ($v1_5)->value5);
} else {
if ((is_object($v2_6_2) && (($v2_6_2)->tag === "GT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v1_5)->value2, ($v1_5)->value3, ($v1_5)->value4, ($go_4_0)(($v1_5)->value5));
} else {
if ((is_object($v2_6_2) && (($v2_6_2)->tag === "EQ"))) {
$__t3 = new Phpurs_Data6("Node", ($v1_5)->value0, ($v1_5)->value1, $k_2, (($app_1)(($v1_5)->value3))($v_3), ($v1_5)->value4, ($v1_5)->value5);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_4_0;
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_insert'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_insert"), recVars=[];
  $go_3_0 = null;
  $go_3_0 = function($v1_4) use ($dictOrd_0, &$go_3_0, $k_1, $v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if ((is_object($v1_4) && (($v1_4)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data6("Node", 1, 1, $k_1, $v_2, new Phpurs_Data0("Leaf"), new Phpurs_Data0("Leaf"));
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "Node"))) {
$v2_5_2 = ((($dictOrd_0)->compare)($k_1))(($v1_4)->value2);
if ((is_object($v2_5_2) && (($v2_5_2)->tag === "LT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v1_4)->value2, ($v1_4)->value3, ($go_3_0)(($v1_4)->value4), ($v1_4)->value5);
} else {
if ((is_object($v2_5_2) && (($v2_5_2)->tag === "GT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v1_4)->value2, ($v1_4)->value3, ($v1_4)->value4, ($go_3_0)(($v1_4)->value5));
} else {
if ((is_object($v2_5_2) && (($v2_5_2)->tag === "EQ"))) {
$__t3 = new Phpurs_Data6("Node", ($v1_4)->value0, ($v1_4)->value1, $k_1, $v_2, ($v1_4)->value4, ($v1_4)->value5);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_functorMap'] = function() { $v = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use ($f_0, &$go_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
$__t1 = new Phpurs_Data6("Node", ($v_2)->value0, ($v_2)->value1, ($v_2)->value2, ($f_0)(($v_2)->value3), ($go_1_0)(($v_2)->value4), ($go_1_0)(($v_2)->value5));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_functorWithIndexMap'] = function() { $v = (object)["mapWithIndex" => function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use ($f_0, &$go_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
$__t1 = new Phpurs_Data6("Node", ($v_2)->value0, ($v_2)->value1, ($v_2)->value2, (($f_0)(($v_2)->value2))(($v_2)->value3), ($go_1_0)(($v_2)->value4), ($go_1_0)(($v_2)->value5));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_functorMap'] ?? \PhpursThunks::eval('Data_Map_Internal_functorMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_foldableMap'] = function() { $v = (object)["foldr" => (function() {
  $__fn = function($f_0, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use ($f_0, &$go_2_0) {
  $__fn = function($m__prime___3, $z__prime___4 = null) use ($f_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($m__prime___3) && (($m__prime___3)->tag === "Leaf"))) {
$__t1 = $z__prime___4;
} else {
if ((is_object($m__prime___3) && (($m__prime___3)->tag === "Node"))) {
$__tco_2 = ($m__prime___3)->value4;
$__tco_3 = (($f_0)(($m__prime___3)->value3))(($go_2_0)(($m__prime___3)->value5, $z__prime___4));
$m__prime___3 = $__tco_2;
$z__prime___4 = $__tco_3;
continue ;
$__t1 = null;
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
  $__res = function($m_3) use (&$go_2_0, $z_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_2_0"];
  $__res = ($go_2_0)($m_3, $z_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() {
  $__fn = function($f_0, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_2_4 = null;
  $go_2_4 = (function() use ($f_0, &$go_2_4) {
  $__fn = function($z__prime___3, $m__prime___4 = null) use ($f_0, &$go_2_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_4"), recVars=["go_2_4"];
  while (true) {
if ((is_object($m__prime___4) && (($m__prime___4)->tag === "Leaf"))) {
$__t5 = $z__prime___3;
} else {
if ((is_object($m__prime___4) && (($m__prime___4)->tag === "Node"))) {
$__tco_6 = (($f_0)(($go_2_4)($z__prime___3, ($m__prime___4)->value4)))(($m__prime___4)->value3);
$__tco_7 = ($m__prime___4)->value5;
$z__prime___3 = $__tco_6;
$m__prime___4 = $__tco_7;
continue ;
$__t5 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = function($m_3) use (&$go_2_4, $z_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_2_4"];
  $__res = ($go_2_4)($z_1, $m_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $mempty_1_8 = ($dictMonoid_0)->mempty;
  $__local_var_2_9 = (($dictMonoid_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_3) use ($__local_var_2_9, $mempty_1_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_4_10 = null;
  $go_4_10 = function($v_5) use ($__local_var_2_9, $f_3, &$go_4_10, $mempty_1_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_10"), recVars=["go_4_10"];
  while (true) {
if ((is_object($v_5) && (($v_5)->tag === "Leaf"))) {
$__t11 = $mempty_1_8;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Node"))) {
$__t11 = ((($__local_var_2_9)->append)(($go_4_10)(($v_5)->value4)))(((($__local_var_2_9)->append)(($f_3)(($v_5)->value3)))(($go_4_10)(($v_5)->value5)));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
};
};
$__res = $__t11;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_4_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_foldableWithIndexMap'] = function() { $v = (object)["foldrWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use ($f_0, &$go_2_0) {
  $__fn = function($m__prime___3, $z__prime___4 = null) use ($f_0, &$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($m__prime___3) && (($m__prime___3)->tag === "Leaf"))) {
$__t1 = $z__prime___4;
} else {
if ((is_object($m__prime___3) && (($m__prime___3)->tag === "Node"))) {
$__tco_2 = ($m__prime___3)->value4;
$__tco_3 = ((($f_0)(($m__prime___3)->value2))(($m__prime___3)->value3))(($go_2_0)(($m__prime___3)->value5, $z__prime___4));
$m__prime___3 = $__tco_2;
$z__prime___4 = $__tco_3;
continue ;
$__t1 = null;
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
  $__res = function($m_3) use (&$go_2_0, $z_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_2_0"];
  $__res = ($go_2_0)($m_3, $z_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldlWithIndex" => (function() {
  $__fn = function($f_0, $z_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_2_4 = null;
  $go_2_4 = (function() use ($f_0, &$go_2_4) {
  $__fn = function($z__prime___3, $m__prime___4 = null) use ($f_0, &$go_2_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_4"), recVars=["go_2_4"];
  while (true) {
if ((is_object($m__prime___4) && (($m__prime___4)->tag === "Leaf"))) {
$__t5 = $z__prime___3;
} else {
if ((is_object($m__prime___4) && (($m__prime___4)->tag === "Node"))) {
$__tco_6 = ((($f_0)(($m__prime___4)->value2))(($go_2_4)($z__prime___3, ($m__prime___4)->value4)))(($m__prime___4)->value3);
$__tco_7 = ($m__prime___4)->value5;
$z__prime___3 = $__tco_6;
$m__prime___4 = $__tco_7;
continue ;
$__t5 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = function($m_3) use (&$go_2_4, $z_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_2_4"];
  $__res = ($go_2_4)($z_1, $m_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "foldMapWithIndex" => function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $mempty_1_8 = ($dictMonoid_0)->mempty;
  $__local_var_2_9 = (($dictMonoid_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_3) use ($__local_var_2_9, $mempty_1_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_4_10 = null;
  $go_4_10 = function($v_5) use ($__local_var_2_9, $f_3, &$go_4_10, $mempty_1_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_10"), recVars=["go_4_10"];
  while (true) {
if ((is_object($v_5) && (($v_5)->tag === "Leaf"))) {
$__t11 = $mempty_1_8;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Node"))) {
$__t11 = ((($__local_var_2_9)->append)(($go_4_10)(($v_5)->value4)))(((($__local_var_2_9)->append)((($f_3)(($v_5)->value2))(($v_5)->value3)))(($go_4_10)(($v_5)->value5)));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
};
};
$__res = $__t11;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_4_10;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_foldableMap'] ?? \PhpursThunks::eval('Data_Map_Internal_foldableMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_keys'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($m__prime___1, $z__prime___2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($m__prime___1) && (($m__prime___1)->tag === "Leaf"))) {
$__t1 = $z__prime___2;
} else {
if ((is_object($m__prime___1) && (($m__prime___1)->tag === "Node"))) {
$__tco_2 = ($m__prime___1)->value4;
$__tco_3 = new Phpurs_Data2("Cons", ($m__prime___1)->value2, ($go_0_0)(($m__prime___1)->value5, $z__prime___2));
$m__prime___1 = $__tco_2;
$z__prime___2 = $__tco_3;
continue ;
$__t1 = null;
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
return function($m_1) use (&$go_0_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_0_0"];
  $__res = ($go_0_0)($m_1, new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_traversableMap'] = function() { $v = (object)["traverse" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Map_Internal_traversableMap"];
  $Apply0_1_0 = (($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Map_Internal_traversableMap"];
  $go_3_1 = null;
  $go_3_1 = function($v_4) use ($Apply0_1_0, $dictApplicative_0, $f_2, &$go_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_1"), recVars=["Data_Map_Internal_traversableMap","go_3_1"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Leaf"))) {
$__t2 = (($dictApplicative_0)->pure)(new Phpurs_Data0("Leaf"));
} else {
if ((is_object($v_4) && (($v_4)->tag === "Node"))) {
$__local_var_5_3 = ($v_4)->value0;
$__local_var_6_4 = ($v_4)->value2;
$__local_var_7_5 = ($v_4)->value1;
$__t2 = ((($Apply0_1_0)->apply)(((($Apply0_1_0)->apply)(((((($Apply0_1_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)((function() use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__fn = function($l__prime___8, $v__prime___9 = null, $r__prime___10 = null) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Map_Internal_traversableMap","go_3_1"];
  $__res = new Phpurs_Data6("Node", $__local_var_5_3, $__local_var_7_5, $__local_var_6_4, $v__prime___9, $l__prime___8, $r__prime___10);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(($go_3_1)(($v_4)->value4))))(($f_2)(($v_4)->value3))))(($go_3_1)(($v_4)->value5));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Map_Internal_traversableMap"];
  $__res = (((($GLOBALS['Data_Map_Internal_traversableMap'] ?? \PhpursThunks::eval('Data_Map_Internal_traversableMap')))->traverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Map_Internal_traversableMap"];
  $__res = ($GLOBALS['Data_Map_Internal_functorMap'] ?? \PhpursThunks::eval('Data_Map_Internal_functorMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_Map_Internal_traversableMap"];
  $__res = ($GLOBALS['Data_Map_Internal_foldableMap'] ?? \PhpursThunks::eval('Data_Map_Internal_foldableMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_traversableWithIndexMap'] = function() { $v = (object)["traverseWithIndex" => function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $Apply0_1_0 = (($dictApplicative_0)->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($f_2) use ($Apply0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_3_1 = null;
  $go_3_1 = function($v_4) use ($Apply0_1_0, $dictApplicative_0, $f_2, &$go_3_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_1"), recVars=["go_3_1"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Leaf"))) {
$__t2 = (($dictApplicative_0)->pure)(new Phpurs_Data0("Leaf"));
} else {
if ((is_object($v_4) && (($v_4)->tag === "Node"))) {
$__local_var_5_3 = ($v_4)->value0;
$__local_var_6_4 = ($v_4)->value2;
$__local_var_7_5 = ($v_4)->value1;
$__t2 = ((($Apply0_1_0)->apply)(((($Apply0_1_0)->apply)(((((($Apply0_1_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)((function() use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5) {
  $__fn = function($l__prime___8, $v__prime___9 = null, $r__prime___10 = null) use ($__local_var_5_3, $__local_var_6_4, $__local_var_7_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_3_1"];
  $__res = new Phpurs_Data6("Node", $__local_var_5_3, $__local_var_7_5, $__local_var_6_4, $v__prime___9, $l__prime___8, $r__prime___10);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(($go_3_1)(($v_4)->value4))))((($f_2)($__local_var_6_4))(($v_4)->value3))))(($go_3_1)(($v_4)->value5));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_3_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FunctorWithIndex0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_functorWithIndexMap'] ?? \PhpursThunks::eval('Data_Map_Internal_functorWithIndexMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "FoldableWithIndex1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_foldableWithIndexMap'] ?? \PhpursThunks::eval('Data_Map_Internal_foldableWithIndexMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Traversable2" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_traversableMap'] ?? \PhpursThunks::eval('Data_Map_Internal_traversableMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_values'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($m__prime___1, $z__prime___2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($m__prime___1) && (($m__prime___1)->tag === "Leaf"))) {
$__t1 = $z__prime___2;
} else {
if ((is_object($m__prime___1) && (($m__prime___1)->tag === "Node"))) {
$__tco_2 = ($m__prime___1)->value4;
$__tco_3 = new Phpurs_Data2("Cons", ($m__prime___1)->value3, ($go_0_0)(($m__prime___1)->value5, $z__prime___2));
$m__prime___1 = $__tco_2;
$z__prime___2 = $__tco_3;
continue ;
$__t1 = null;
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
return function($m_1) use (&$go_0_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_0_0"];
  $__res = ($go_0_0)($m_1, new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_foldSubmapBy'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $appendFn_1 = null, $memptyValue_2 = null, $kmin_3 = null, $kmax_4 = null, $f_5 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_foldSubmapBy"), recVars=[];
  if ((is_object($kmin_3) && (($kmin_3)->tag === "Just"))) {
$__local_var_6_1 = ($kmin_3)->value0;
$__t0 = function($k_7) use ($__local_var_6_1, $dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (is_object(((($dictOrd_0)->compare)($k_7))($__local_var_6_1)) && ((((($dictOrd_0)->compare)($k_7))($__local_var_6_1))->tag === "LT"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
if ((is_object($kmin_3) && (($kmin_3)->tag === "Nothing"))) {
$__t0 = function($v_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = false;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $tooSmall_6_0 = $__t0;
  if ((is_object($kmax_4) && (($kmax_4)->tag === "Just"))) {
$__local_var_7_4 = ($kmax_4)->value0;
$__t3 = function($k_8) use ($__local_var_7_4, $dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (is_object(((($dictOrd_0)->compare)($k_8))($__local_var_7_4)) && ((((($dictOrd_0)->compare)($k_8))($__local_var_7_4))->tag === "GT"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
if ((is_object($kmax_4) && (($kmax_4)->tag === "Nothing"))) {
$__t3 = function($v_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = false;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $tooLarge_7_3 = $__t3;
  if ((is_object($kmin_3) && (($kmin_3)->tag === "Just"))) {
if ((is_object($kmax_4) && (($kmax_4)->tag === "Just"))) {
$__local_var_8_8 = ($kmax_4)->value0;
$__local_var_9_9 = ($kmin_3)->value0;
$__t7 = function($k_10) use ($__local_var_8_8, $__local_var_9_9, $dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (( ! (is_object(((($dictOrd_0)->compare)($__local_var_9_9))($k_10)) && ((((($dictOrd_0)->compare)($__local_var_9_9))($k_10))->tag === "GT"))) && ( ! (is_object(((($dictOrd_0)->compare)($k_10))($__local_var_8_8)) && ((((($dictOrd_0)->compare)($k_10))($__local_var_8_8))->tag === "GT"))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
if ((is_object($kmax_4) && (($kmax_4)->tag === "Nothing"))) {
$__local_var_8_10 = ($kmin_3)->value0;
$__t7 = function($k_9) use ($__local_var_8_10, $dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ( ! (is_object(((($dictOrd_0)->compare)($__local_var_8_10))($k_9)) && ((((($dictOrd_0)->compare)($__local_var_8_10))($k_9))->tag === "GT")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
$__t6 = $__t7;
} else {
if ((is_object($kmin_3) && (($kmin_3)->tag === "Nothing"))) {
if ((is_object($kmax_4) && (($kmax_4)->tag === "Just"))) {
$__local_var_8_12 = ($kmax_4)->value0;
$__t11 = function($k_9) use ($__local_var_8_12, $dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ( ! (is_object(((($dictOrd_0)->compare)($k_9))($__local_var_8_12)) && ((((($dictOrd_0)->compare)($k_9))($__local_var_8_12))->tag === "GT")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
if ((is_object($kmax_4) && (($kmax_4)->tag === "Nothing"))) {
$__t11 = function($v_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = true;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t11 = null;
};
};
$__t6 = $__t11;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
  $inBounds_8_6 = $__t6;
  $go_9_14 = null;
  $go_9_14 = function($v_10) use ($appendFn_1, $f_5, &$go_9_14, $inBounds_8_6, $memptyValue_2, $tooLarge_7_3, $tooSmall_6_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_9_14"), recVars=["go_9_14"];
  while (true) {
if ((is_object($v_10) && (($v_10)->tag === "Leaf"))) {
$__t15 = $memptyValue_2;
} else {
if ((is_object($v_10) && (($v_10)->tag === "Node"))) {
if (($tooSmall_6_0)(($v_10)->value2)) {
$__t16 = $memptyValue_2;
} else {
$__t16 = ($go_9_14)(($v_10)->value4);
};
if (($inBounds_8_6)(($v_10)->value2)) {
$__t17 = (($f_5)(($v_10)->value2))(($v_10)->value3);
} else {
$__t17 = $memptyValue_2;
};
if (($tooLarge_7_3)(($v_10)->value2)) {
$__t18 = $memptyValue_2;
} else {
$__t18 = ($go_9_14)(($v_10)->value5);
};
$__t15 = (($appendFn_1)((($appendFn_1)($__t16))($__t17)))($__t18);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t15 = null;
};
};
$__res = $__t15;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_9_14;
  goto __end;;
  __end:
  return $__num > 6 ? $__res(...\array_slice(\func_get_args(), 6)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_foldSubmap'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictMonoid_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_foldSubmap"), recVars=[];
  $__res = (((($GLOBALS['Data_Map_Internal_foldSubmapBy'] ?? \PhpursThunks::eval('Data_Map_Internal_foldSubmapBy')))($dictOrd_0))(((($dictMonoid_1)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->append))(($dictMonoid_1)->mempty);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_findMin'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_findMin"), recVars=["Data_Map_Internal_findMin"];
  while (true) {
if ((is_object($v_0) && (($v_0)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Node"))) {
if ((is_object(($v_0)->value4) && ((($v_0)->value4)->tag === "Leaf"))) {
$__t2 = new Phpurs_Data1("Just", (object)["key" => ($v_0)->value2, "value" => ($v_0)->value3]);
} else {
$__tco_1 = ($v_0)->value4;
$v_0 = $__tco_1;
continue ;
$__t2 = null;
};
$__t0 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_lookupGT'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_lookupGT"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$v2_5_4 = ($go_2_0)(($v_3)->value4);
if ((is_object($v2_5_4) && (($v2_5_4)->tag === "Nothing"))) {
$__t5 = new Phpurs_Data1("Just", (object)["key" => ($v_3)->value2, "value" => ($v_3)->value3]);
} else {
$__t5 = $v2_5_4;
};
$__t3 = $__t5;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$__tco_6 = ($v_3)->value5;
$v_3 = $__tco_6;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_findMin'] ?? \PhpursThunks::eval('Data_Map_Internal_findMin')))(($v_3)->value5);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_findMax'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_findMax"), recVars=["Data_Map_Internal_findMax"];
  while (true) {
if ((is_object($v_0) && (($v_0)->tag === "Leaf"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Node"))) {
if ((is_object(($v_0)->value5) && ((($v_0)->value5)->tag === "Leaf"))) {
$__t2 = new Phpurs_Data1("Just", (object)["key" => ($v_0)->value2, "value" => ($v_0)->value3]);
} else {
$__tco_1 = ($v_0)->value5;
$v_0 = $__tco_1;
continue ;
$__t2 = null;
};
$__t0 = $__t2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_lookupLT'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_lookupLT"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$__tco_4 = ($v_3)->value4;
$v_3 = $__tco_4;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$v2_5_5 = ($go_2_0)(($v_3)->value5);
if ((is_object($v2_5_5) && (($v2_5_5)->tag === "Nothing"))) {
$__t6 = new Phpurs_Data1("Just", (object)["key" => ($v_3)->value2, "value" => ($v_3)->value3]);
} else {
$__t6 = $v2_5_5;
};
$__t3 = $__t6;
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_findMax'] ?? \PhpursThunks::eval('Data_Map_Internal_findMax')))(($v_3)->value4);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_filterWithKey'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_filterWithKey"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($f_1, &$go_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
if ((($f_1)(($v_3)->value2))(($v_3)->value3)) {
$__t2 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v_3)->value2, ($v_3)->value3, ($go_2_0)(($v_3)->value4), ($go_2_0)(($v_3)->value5));
} else {
$__t2 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))(($go_2_0)(($v_3)->value4), ($go_2_0)(($v_3)->value5));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_filterKeys'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_filterKeys"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($f_1, &$go_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
if (($f_1)(($v_3)->value2)) {
$__t2 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v_3)->value2, ($v_3)->value3, ($go_2_0)(($v_3)->value4), ($go_2_0)(($v_3)->value5));
} else {
$__t2 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))(($go_2_0)(($v_3)->value4), ($go_2_0)(($v_3)->value5));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_filter'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_filter"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Map_Internal_filterWithKey'] ?? \PhpursThunks::eval('Data_Map_Internal_filterWithKey')))($dictOrd_0)))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_eqMap'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $dictEq1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_eqMap"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq1_1, $dictEq_0) {
  $__fn = function($xs_2, $ys_3 = null) use ($dictEq1_1, $dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($xs_2) && (($xs_2)->tag === "Leaf"))) {
$__t0 = (is_object($ys_3) && (($ys_3)->tag === "Leaf"));
} else {
if ((is_object($xs_2) && (($xs_2)->tag === "Node"))) {
$__t0 = ((is_object($ys_3) && (($ys_3)->tag === "Node")) && ((($xs_2)->value1 === ($ys_3)->value1) && (((((($GLOBALS['Data_Map_Internal_eqMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMapIter')))($dictEq_0))($dictEq1_1))->eq)(new Phpurs_Data2("IterNode", $xs_2, new Phpurs_Data0("IterLeaf"))))(new Phpurs_Data2("IterNode", $ys_3, new Phpurs_Data0("IterLeaf")))));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
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
\PhpursThunks::$thunks['Data_Map_Internal_ordMap'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_ordMap"), recVars=[];
  $ordMapIter1_1_0 = (($GLOBALS['Data_Map_Internal_ordMapIter'] ?? \PhpursThunks::eval('Data_Map_Internal_ordMapIter')))($dictOrd_0);
  $eqMap1_2_1 = (($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))((($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictOrd1_3) use ($eqMap1_2_1, $ordMapIter1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $eqMap2_4_2 = ($eqMap1_2_1)((($dictOrd1_3)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = (object)["compare" => (function() use ($dictOrd1_3, $ordMapIter1_1_0) {
  $__fn = function($xs_5, $ys_6 = null) use ($dictOrd1_3, $ordMapIter1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($xs_5) && (($xs_5)->tag === "Leaf"))) {
if ((is_object($ys_6) && (($ys_6)->tag === "Leaf"))) {
$__t4 = new Phpurs_Data0("EQ");
} else {
$__t4 = new Phpurs_Data0("LT");
};
$__t3 = $__t4;
} else {
if ((is_object($ys_6) && (($ys_6)->tag === "Leaf"))) {
$__t3 = new Phpurs_Data0("GT");
} else {
$__t3 = (((($ordMapIter1_1_0)($dictOrd1_3))->compare)(new Phpurs_Data2("IterNode", $xs_5, new Phpurs_Data0("IterLeaf"))))(new Phpurs_Data2("IterNode", $ys_6, new Phpurs_Data0("IterLeaf")));
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_5) use ($eqMap2_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqMap2_4_2;
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
\PhpursThunks::$thunks['Data_Map_Internal_eq1Map'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_eq1Map"), recVars=[];
  $__res = (object)["eq1" => function($dictEq1_1) use ($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))($dictEq_0))($dictEq1_1))->eq;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_ord1Map'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_ord1Map"), recVars=[];
  $ordMap1_1_0 = (($GLOBALS['Data_Map_Internal_ordMap'] ?? \PhpursThunks::eval('Data_Map_Internal_ordMap')))($dictOrd_0);
  $__local_var_2_1 = (($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eq1Map1_3_2 = (object)["eq1" => function($dictEq1_3) use ($__local_var_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Map_Internal_eqMap'] ?? \PhpursThunks::eval('Data_Map_Internal_eqMap')))($__local_var_2_1))($dictEq1_3))->eq;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["compare1" => function($dictOrd1_4) use ($ordMap1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($ordMap1_1_0)($dictOrd1_4))->compare;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Eq10" => function($dollar__unused_4) use ($eq1Map1_3_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eq1Map1_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_empty'] = function() { $v = new Phpurs_Data0("Leaf"); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_fromFoldable'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictFoldable_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_fromFoldable"), recVars=[];
  $__res = ((($dictFoldable_1)->foldl)((function() use ($dictOrd_0) {
  $__fn = function($m_2, $v_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_Map_Internal_insert'] ?? \PhpursThunks::eval('Data_Map_Internal_insert')))($dictOrd_0))(($v_3)->value0))(($v_3)->value1))($m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_fromFoldableWith'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictFoldable_1 = null, $f_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_fromFoldableWith"), recVars=[];
  $f__prime___3_0 = ((($GLOBALS['Data_Map_Internal_insertWith'] ?? \PhpursThunks::eval('Data_Map_Internal_insertWith')))($dictOrd_0))((function() use ($f_2) {
  $__fn = function($b_3, $a_4 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_2)($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  $__res = ((($dictFoldable_1)->foldl)((function() use ($f__prime___3_0) {
  $__fn = function($m_4, $v_5 = null) use ($f__prime___3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($f__prime___3_0)(($v_5)->value0))(($v_5)->value1))($m_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_fromFoldableWithIndex'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $dictFoldableWithIndex_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_fromFoldableWithIndex"), recVars=[];
  $__res = ((($dictFoldableWithIndex_1)->foldlWithIndex)((function() use ($dictOrd_0) {
  $__fn = function($k_2, $m_3 = null, $v_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_Map_Internal_insert'] ?? \PhpursThunks::eval('Data_Map_Internal_insert')))($dictOrd_0))($k_2))($v_4))($m_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_monoidSemigroupMap'] = function() { $v = (function() {
  $__fn = function($dollar__unused_0, $dictOrd_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_monoidSemigroupMap"), recVars=[];
  $semigroupMap2_2_0 = ((($GLOBALS['Data_Map_Internal_semigroupMap'] ?? \PhpursThunks::eval('Data_Map_Internal_semigroupMap')))(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))($dictOrd_1);
  $__res = function($dictSemigroup_3) use ($semigroupMap2_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $semigroupMap3_4_1 = ($semigroupMap2_2_0)($dictSemigroup_3);
  $__res = (object)["mempty" => new Phpurs_Data0("Leaf"), "Semigroup0" => function($dollar__unused_5) use ($semigroupMap3_4_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $semigroupMap3_4_1;
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
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_submap'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_submap"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0, $dictOrd_0) {
  $__fn = function($kmin_2, $kmax_3 = null) use ($compare_1_0, $dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((((($GLOBALS['Data_Map_Internal_foldSubmapBy'] ?? \PhpursThunks::eval('Data_Map_Internal_foldSubmapBy')))($dictOrd_0))((function() use ($compare_1_0) {
  $__fn = function($m1_4, $m2_5 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_4, $m2_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf")))($kmin_2))($kmax_3))(($GLOBALS['Data_Map_Internal_singleton'] ?? \PhpursThunks::eval('Data_Map_Internal_singleton')));
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
\PhpursThunks::$thunks['Data_Map_Internal_unions'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_unions"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = function($dictFoldable_2) use ($compare_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictFoldable_2)->foldl)((function() use ($compare_1_0) {
  $__fn = function($m1_3, $m2_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_3, $m2_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_difference'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_difference"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeDifference'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeDifference')))($compare_1_0, $m1_2, $m2_3);
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
\PhpursThunks::$thunks['Data_Map_Internal_delete'] = function() { $v = (function() {
  $__fn = function($dictOrd_0, $k_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_delete"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = function($v_3) use ($dictOrd_0, &$go_2_0, $k_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Leaf"))) {
$__t1 = new Phpurs_Data0("Leaf");
} else {
if ((is_object($v_3) && (($v_3)->tag === "Node"))) {
$v1_4_2 = ((($dictOrd_0)->compare)($k_1))(($v_3)->value2);
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "LT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v_3)->value2, ($v_3)->value3, ($go_2_0)(($v_3)->value4), ($v_3)->value5);
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "GT"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))(($v_3)->value2, ($v_3)->value3, ($v_3)->value4, ($go_2_0)(($v_3)->value5));
} else {
if ((is_object($v1_4_2) && (($v1_4_2)->tag === "EQ"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))(($v_3)->value4, ($v_3)->value5);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_checkValid'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_checkValid"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use ($dictOrd_0, &$go_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = true;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
if ((is_object(($v_2)->value4) && ((($v_2)->value4)->tag === "Leaf"))) {
if ((is_object(($v_2)->value5) && ((($v_2)->value5)->tag === "Leaf"))) {
$__t3 = true;
} else {
if ((is_object(($v_2)->value5) && ((($v_2)->value5)->tag === "Node"))) {
$__t3 = ((($v_2)->value0 === 2) && (((($v_2)->value5)->value0 === 1) && ((($v_2)->value1 > (($v_2)->value5)->value1) && ((is_object(((($dictOrd_0)->compare)((($v_2)->value5)->value2))(($v_2)->value2)) && ((((($dictOrd_0)->compare)((($v_2)->value5)->value2))(($v_2)->value2))->tag === "GT")) && ($go_1_0)(($v_2)->value5)))));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__t2 = $__t3;
} else {
if ((is_object(($v_2)->value4) && ((($v_2)->value4)->tag === "Node"))) {
if ((is_object(($v_2)->value5) && ((($v_2)->value5)->tag === "Leaf"))) {
$__t4 = ((($v_2)->value0 === 2) && (((($v_2)->value4)->value0 === 1) && ((($v_2)->value1 > (($v_2)->value4)->value1) && ((is_object(((($dictOrd_0)->compare)((($v_2)->value4)->value2))(($v_2)->value2)) && ((((($dictOrd_0)->compare)((($v_2)->value4)->value2))(($v_2)->value2))->tag === "LT")) && ($go_1_0)(($v_2)->value4)))));
} else {
if ((is_object(($v_2)->value5) && ((($v_2)->value5)->tag === "Node"))) {
$__t4 = ((($v_2)->value0 > (($v_2)->value5)->value0) && ((is_object(((($dictOrd_0)->compare)((($v_2)->value5)->value2))(($v_2)->value2)) && ((((($dictOrd_0)->compare)((($v_2)->value5)->value2))(($v_2)->value2))->tag === "GT")) && ((($v_2)->value0 > (($v_2)->value4)->value0) && ((is_object(((($dictOrd_0)->compare)((($v_2)->value4)->value2))(($v_2)->value2)) && ((((($dictOrd_0)->compare)((($v_2)->value4)->value2))(($v_2)->value2))->tag === "LT")) && (((($GLOBALS['Data_Map_Internal_abs'] ?? \PhpursThunks::eval('Data_Map_Internal_abs')))(((($v_2)->value5)->value0 - (($v_2)->value4)->value0)) < 2) && (((((($v_2)->value5)->value1 + (($v_2)->value4)->value1) + 1) === ($v_2)->value1) && (($go_1_0)(($v_2)->value4) && ($go_1_0)(($v_2)->value5))))))));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__t2 = $__t4;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_catMaybes'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_catMaybes"), recVars=[];
  $__res = (((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_Map_Internal_mapMaybeWithKey'] ?? \PhpursThunks::eval('Data_Map_Internal_mapMaybeWithKey')))($dictOrd_0)))(($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const'))))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_applyMap'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_applyMap"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (object)["apply" => (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith')))($compare_1_0, (($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity, $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_functorMap'] ?? \PhpursThunks::eval('Data_Map_Internal_functorMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_bindMap'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_bindMap"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $applyMap1_1_0 = (object)["apply" => (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeIntersectionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeIntersectionWith')))($compare_1_0, (($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity, $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_functorMap'] ?? \PhpursThunks::eval('Data_Map_Internal_functorMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["bind" => (function() use ($dictOrd_0) {
  $__fn = function($m_2, $f_3 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_Map_Internal_mapMaybeWithKey'] ?? \PhpursThunks::eval('Data_Map_Internal_mapMaybeWithKey')))($dictOrd_0))(function($k_4) use ($dictOrd_0, $f_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_5_2 = null;
  $go_5_2 = function($v_6) use ($dictOrd_0, &$go_5_2, $k_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_2"), recVars=["go_5_2"];
  while (true) {
if ((is_object($v_6) && (($v_6)->tag === "Leaf"))) {
$__t3 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_6) && (($v_6)->tag === "Node"))) {
$v1_7_4 = ((($dictOrd_0)->compare)($k_4))(($v_6)->value2);
if ((is_object($v1_7_4) && (($v1_7_4)->tag === "LT"))) {
$__tco_6 = ($v_6)->value4;
$v_6 = $__tco_6;
continue ;
$__t5 = null;
} else {
if ((is_object($v1_7_4) && (($v1_7_4)->tag === "GT"))) {
$__tco_7 = ($v_6)->value5;
$v_6 = $__tco_7;
continue ;
$__t5 = null;
} else {
if ((is_object($v1_7_4) && (($v1_7_4)->tag === "EQ"))) {
$__t5 = new Phpurs_Data1("Just", ($v_6)->value3);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
};
$__t3 = $__t5;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($go_5_2))($f_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($m_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Apply0" => function($dollar__unused_2) use ($applyMap1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $applyMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_anyWithKey'] = function() { $v = function($predicate_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_anyWithKey"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use (&$go_1_0, $predicate_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = false;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
$__t1 = ((($predicate_0)(($v_2)->value2))(($v_2)->value3) || (($go_1_0)(($v_2)->value4) || ($go_1_0)(($v_2)->value5)));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_any'] = function() { $v = function($predicate_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_any"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use (&$go_1_0, $predicate_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Leaf"))) {
$__t1 = false;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Node"))) {
$__t1 = (($predicate_0)(($v_2)->value3) || (($go_1_0)(($v_2)->value4) || ($go_1_0)(($v_2)->value5)));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_alter'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_alter"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (function() use ($compare_1_0) {
  $__fn = function($f_2, $k_3 = null, $m_4 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v_5_1 = (($GLOBALS['Data_Map_Internal_unsafeSplit'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeSplit')))($compare_1_0, $k_3, $m_4);
  $v2_6_2 = ($f_2)(($v_5_1)->value0);
  if ((is_object($v2_6_2) && (($v2_6_2)->tag === "Nothing"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeJoinNodes'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeJoinNodes')))(($v_5_1)->value1, ($v_5_1)->value2);
} else {
if ((is_object($v2_6_2) && (($v2_6_2)->tag === "Just"))) {
$__t3 = (($GLOBALS['Data_Map_Internal_unsafeBalancedNode'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeBalancedNode')))($k_3, ($v2_6_2)->value0, ($v_5_1)->value1, ($v_5_1)->value2);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_altMap'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_altMap"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = (object)["alt" => (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_functorMap'] ?? \PhpursThunks::eval('Data_Map_Internal_functorMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Map_Internal_plusMap'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_Map_Internal_plusMap"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $altMap1_1_0 = (object)["alt" => (function() use ($compare_1_0) {
  $__fn = function($m1_2, $m2_3 = null) use ($compare_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Map_Internal_unsafeUnionWith'] ?? \PhpursThunks::eval('Data_Map_Internal_unsafeUnionWith')))($compare_1_0, ($GLOBALS['Data_Function_const'] ?? \PhpursThunks::eval('Data_Function_const')), $m1_2, $m2_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Data_Map_Internal_functorMap'] ?? \PhpursThunks::eval('Data_Map_Internal_functorMap'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["empty" => new Phpurs_Data0("Leaf"), "Alt0" => function($dollar__unused_2) use ($altMap1_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $altMap1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


































































































