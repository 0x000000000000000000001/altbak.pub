<?php

namespace Data\List;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Internal, Data.List.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.List, Data.List.Internal, Data.List.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Data.Unit, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.List/index.php';
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Types/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.NonEmpty/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
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
\PhpursThunks::$thunks['Data_List_tailRecM2'] = function() { $v = (function() {
  $__fn = function($f_0, $a_1 = null, $b_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_tailRecM2"), recVars=[];
  $__res = (((($GLOBALS['Control_Monad_Rec_Class_monadRecMaybe'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_monadRecMaybe')))->tailRecM)(function($o_3) use ($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($f_0)(($o_3)->a))(($o_3)->b);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_1, "b" => $b_2]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_any'] = function() { $v = ((($GLOBALS['Data_Foldable_any'] ?? \PhpursThunks::eval('Data_Foldable_any')))(($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList'))))(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_heytingAlgebraBoolean'))); return $v; };
\PhpursThunks::$thunks['Data_List_Pattern'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Pattern"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_updateAt'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_updateAt"), recVars=["Data_List_updateAt"];
  while (true) {
if ((is_object($v2_2) && (($v2_2)->tag === "Cons"))) {
switch ($v_0) {
case 0:
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", $v1_1, ($v2_2)->value1));
break;
default:
$__local_var_3_1 = (((($GLOBALS['Data_List_updateAt'] ?? \PhpursThunks::eval('Data_List_updateAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1)))($v1_1))(($v2_2)->value1);
if ((is_object($__local_var_3_1) && (($__local_var_3_1)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v2_2)->value0, ($__local_var_3_1)->value0));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
$__t3 = $__t2;
break;
};
$__t0 = $__t3;
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_unzip'] = function() { $v = (((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_1_0 = ($v_0)->value0;
  $__local_var_2_1 = ($v_0)->value1;
  $__res = function($v1_3) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Tuple", new Phpurs_Data2("Cons", $__local_var_1_0, ($v1_3)->value0), new Phpurs_Data2("Cons", $__local_var_2_1, ($v1_3)->value1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", new Phpurs_Data0("Nil"), new Phpurs_Data0("Nil"))); return $v; };
\PhpursThunks::$thunks['Data_List_uncons'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_uncons"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", (object)["head" => ($v_0)->value0, "tail" => ($v_0)->value1]);
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
\PhpursThunks::$thunks['Data_List_toUnfoldable'] = function() { $v = function($dictUnfoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_toUnfoldable"), recVars=[];
  $__res = (($dictUnfoldable_0)->unfoldr)(function($xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($xs_1) && (($xs_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($xs_1) && (($xs_1)->tag === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($xs_1)->value0, ($xs_1)->value1));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_tail'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_tail"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", ($v_0)->value1);
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
\PhpursThunks::$thunks['Data_List_stripPrefix'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $v_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_stripPrefix"), recVars=[];
  $__res = (((($GLOBALS['Data_List_tailRecM2'] ?? \PhpursThunks::eval('Data_List_tailRecM2')))((function() use ($dictEq_0) {
  $__fn = function($prefix_3, $input_4 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($input_4) && (($input_4)->tag === "Cons"))) {
if ((is_object($prefix_3) && (($prefix_3)->tag === "Cons"))) {
if (((($dictEq_0)->eq)(($prefix_3)->value0))(($input_4)->value0)) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data1("Loop", (object)["a" => ($prefix_3)->value1, "b" => ($input_4)->value1]));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
$__t1 = $__t2;
} else {
if ((is_object($prefix_3) && (($prefix_3)->tag === "Nil"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data1("Done", $input_4));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
};
$__t0 = $__t1;
} else {
if ((is_object($prefix_3) && (($prefix_3)->tag === "Nil"))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data1("Done", $input_4));
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($v_1))($s_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_span'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_span"), recVars=["Data_List_span"];
  while (true) {
if (((is_object($v1_1) && (($v1_1)->tag === "Cons")) && ($v_0)(($v1_1)->value0))) {
$v2_2_1 = ((($GLOBALS['Data_List_span'] ?? \PhpursThunks::eval('Data_List_span')))($v_0))(($v1_1)->value1);
$__t0 = (object)["init" => new Phpurs_Data2("Cons", ($v1_1)->value0, ($v2_2_1)->init), "rest" => ($v2_2_1)->rest];
} else {
$__t0 = (object)["init" => new Phpurs_Data0("Nil"), "rest" => $v1_1];
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_snoc'] = function() { $v = (function() {
  $__fn = function($xs_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_snoc"), recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))(new Phpurs_Data2("Cons", $x_1, new Phpurs_Data0("Nil"))))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_singleton'] = function() { $v = function($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_singleton"), recVars=[];
  $__res = new Phpurs_Data2("Cons", $a_0, new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_sortBy'] = function() { $v = function($cmp_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_sortBy"), recVars=[];
  $merge_1_0 = null;
  $merge_1_0 = (function() use ($cmp_0, &$merge_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($cmp_0, &$merge_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "merge_1_0"), recVars=["merge_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Cons"))) {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
if ((is_object((($cmp_0)(($v_2)->value0))(($v1_3)->value0)) && (((($cmp_0)(($v_2)->value0))(($v1_3)->value0))->tag === "GT"))) {
$__t3 = new Phpurs_Data2("Cons", ($v1_3)->value0, (($merge_1_0)($v_2))(($v1_3)->value1));
} else {
$__t3 = new Phpurs_Data2("Cons", ($v_2)->value0, (($merge_1_0)(($v_2)->value1))($v1_3));
};
$__t2 = $__t3;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$__t2 = $v_2;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
$__t1 = $__t2;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = $v1_3;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$__t1 = $v_2;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $mergePairs_2_4 = null;
  $mergePairs_2_4 = function($v_3) use (&$mergePairs_2_4, &$merge_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "mergePairs_2_4"), recVars=["merge_1_0","mergePairs_2_4"];
  while (true) {
if (((is_object($v_3) && (($v_3)->tag === "Cons")) && (is_object(($v_3)->value1) && ((($v_3)->value1)->tag === "Cons")))) {
$__t5 = new Phpurs_Data2("Cons", (($merge_1_0)(($v_3)->value0))((($v_3)->value1)->value0), ($mergePairs_2_4)((($v_3)->value1)->value1));
} else {
$__t5 = $v_3;
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $mergeAll_3_6 = null;
  $mergeAll_3_6 = function($v_4) use (&$mergeAll_3_6, &$mergePairs_2_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "mergeAll_3_6"), recVars=["merge_1_0","mergePairs_2_4","mergeAll_3_6"];
  while (true) {
if (((is_object($v_4) && (($v_4)->tag === "Cons")) && (is_object(($v_4)->value1) && ((($v_4)->value1)->tag === "Nil")))) {
$__t8 = ($v_4)->value0;
} else {
$__tco_7 = ($mergePairs_2_4)($v_4);
$v_4 = $__tco_7;
continue ;
$__t8 = null;
};
$__res = $__t8;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $sequences_4_9 = null;
  $descending_4_9 = null;
  $ascending_4_9 = null;
  $sequences_4_9 = function($v_5) use (&$ascending_4_9, $cmp_0, &$descending_4_9) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "sequences_4_9"), recVars=["merge_1_0","mergePairs_2_4","mergeAll_3_6","sequences_4_9","descending_4_9","ascending_4_9"];
  while (true) {
if (((is_object($v_5) && (($v_5)->tag === "Cons")) && (is_object(($v_5)->value1) && ((($v_5)->value1)->tag === "Cons")))) {
if ((is_object((($cmp_0)(($v_5)->value0))((($v_5)->value1)->value0)) && (((($cmp_0)(($v_5)->value0))((($v_5)->value1)->value0))->tag === "GT"))) {
$__t12 = ((($descending_4_9)((($v_5)->value1)->value0))(new Phpurs_Data2("Cons", ($v_5)->value0, new Phpurs_Data0("Nil"))))((($v_5)->value1)->value1);
} else {
$__local_var_6_11 = ($v_5)->value0;
$__t12 = ((($ascending_4_9)((($v_5)->value1)->value0))(function($v1_7) use ($__local_var_6_11) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["merge_1_0","mergePairs_2_4","mergeAll_3_6","sequences_4_9","descending_4_9","ascending_4_9"];
  $__res = new Phpurs_Data2("Cons", $__local_var_6_11, $v1_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($v_5)->value1)->value1);
};
$__t10 = $__t12;
} else {
$__t10 = new Phpurs_Data2("Cons", $v_5, new Phpurs_Data0("Nil"));
};
$__res = $__t10;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $descending_4_9 = (function() use ($cmp_0, &$descending_4_9, &$sequences_4_9) {
  $__fn = function($v_5, $v1_6 = null, $v2_7 = null) use ($cmp_0, &$descending_4_9, &$sequences_4_9, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "descending_4_9"), recVars=["merge_1_0","mergePairs_2_4","mergeAll_3_6","sequences_4_9","descending_4_9","ascending_4_9"];
  while (true) {
if (((is_object($v2_7) && (($v2_7)->tag === "Cons")) && (is_object((($cmp_0)($v_5))(($v2_7)->value0)) && (((($cmp_0)($v_5))(($v2_7)->value0))->tag === "GT")))) {
$__tco_14 = ($v2_7)->value0;
$__tco_15 = new Phpurs_Data2("Cons", $v_5, $v1_6);
$__tco_16 = ($v2_7)->value1;
$v_5 = $__tco_14;
$v1_6 = $__tco_15;
$v2_7 = $__tco_16;
continue ;
$__t13 = null;
} else {
$__t13 = new Phpurs_Data2("Cons", new Phpurs_Data2("Cons", $v_5, $v1_6), ($sequences_4_9)($v2_7));
};
$__res = $__t13;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $ascending_4_9 = (function() use (&$ascending_4_9, $cmp_0, &$sequences_4_9) {
  $__fn = function($v_5, $v1_6 = null, $v2_7 = null) use (&$ascending_4_9, $cmp_0, &$sequences_4_9, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "ascending_4_9"), recVars=["merge_1_0","mergePairs_2_4","mergeAll_3_6","sequences_4_9","descending_4_9","ascending_4_9"];
  while (true) {
if (((is_object($v2_7) && (($v2_7)->tag === "Cons")) && ((($GLOBALS['Data_Eq_eqBooleanImpl'] ?? \PhpursThunks::eval('Data_Eq_eqBooleanImpl')))((is_object((($cmp_0)($v_5))(($v2_7)->value0)) && (((($cmp_0)($v_5))(($v2_7)->value0))->tag === "GT"))))(false))) {
$__tco_18 = ($v2_7)->value0;
$__tco_19 = function($ys_8) use ($v1_6, $v_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["merge_1_0","mergePairs_2_4","mergeAll_3_6","sequences_4_9","descending_4_9","ascending_4_9"];
  $__res = ($v1_6)(new Phpurs_Data2("Cons", $v_5, $ys_8));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
$__tco_20 = ($v2_7)->value1;
$v_5 = $__tco_18;
$v1_6 = $__tco_19;
$v2_7 = $__tco_20;
continue ;
$__t17 = null;
} else {
$__t17 = new Phpurs_Data2("Cons", ($v1_6)(new Phpurs_Data2("Cons", $v_5, new Phpurs_Data0("Nil"))), ($sequences_4_9)($v2_7));
};
$__res = $__t17;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($mergeAll_3_6))($sequences_4_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_sort'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_sort"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_List_sortBy'] ?? \PhpursThunks::eval('Data_List_sortBy')))($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_showPattern'] = function() { $v = function($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_showPattern"), recVars=[];
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Pattern "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((((($GLOBALS['Data_List_Types_showList'] ?? \PhpursThunks::eval('Data_List_Types_showList')))($dictShow_0))->show)($v_1)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_reverse'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($v_1, $v1_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($v1_2) && (($v1_2)->tag === "Nil"))) {
$__t1 = $v_1;
} else {
if ((is_object($v1_2) && (($v1_2)->tag === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_2)->value0, $v_1);
$__tco_3 = ($v1_2)->value1;
$v_1 = $__tco_2;
$v1_2 = $__tco_3;
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
return ($go_0_0)(new Phpurs_Data0("Nil"));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_take'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($v_1, $v1_2 = null, $v2_3 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v1_2))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v1_2))(1))->tag === "LT"))) {
$go_4_2 = null;
$go_4_2 = (function() use (&$go_4_2) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_2"), recVars=["go_0_0","go_4_2"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t3 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_5 = ($v1_6)->value1;
$v_5 = $__tco_4;
$v1_6 = $__tco_5;
continue ;
$__t3 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = (($go_4_2)(new Phpurs_Data0("Nil")))($v_1);
} else {
if ((is_object($v2_3) && (($v2_3)->tag === "Nil"))) {
$go_4_6 = null;
$go_4_6 = (function() use (&$go_4_6) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_6"), recVars=["go_0_0","go_4_6"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t7 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_8 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_9 = ($v1_6)->value1;
$v_5 = $__tco_8;
$v1_6 = $__tco_9;
continue ;
$__t7 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
$__res = $__t7;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = (($go_4_6)(new Phpurs_Data0("Nil")))($v_1);
} else {
if ((is_object($v2_3) && (($v2_3)->tag === "Cons"))) {
$__tco_10 = new Phpurs_Data2("Cons", ($v2_3)->value0, $v_1);
$__tco_11 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v1_2))(1);
$__tco_12 = ($v2_3)->value1;
$v_1 = $__tco_10;
$v1_2 = $__tco_11;
$v2_3 = $__tco_12;
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
return ($go_0_0)(new Phpurs_Data0("Nil"));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_takeWhile'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_takeWhile"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if (((is_object($v1_3) && (($v1_3)->tag === "Cons")) && ($p_0)(($v1_3)->value0))) {
$__tco_6 = new Phpurs_Data2("Cons", ($v1_3)->value0, $v_2);
$__tco_7 = ($v1_3)->value1;
$v_2 = $__tco_6;
$v1_3 = $__tco_7;
continue ;
$__t5 = null;
} else {
$go_4_1 = null;
$go_4_1 = (function() use (&$go_4_1) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_1"), recVars=["go_1_0","go_4_1"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t2 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_3 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_4 = ($v1_6)->value1;
$v_5 = $__tco_3;
$v1_6 = $__tco_4;
continue ;
$__t2 = null;
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
$__t5 = (($go_4_1)(new Phpurs_Data0("Nil")))($v_2);
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go_1_0)(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_unsnoc'] = function() { $v = function($lst_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_unsnoc"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_2) && (($v_2)->tag === "Cons"))) {
if ((is_object(($v_2)->value1) && ((($v_2)->value1)->tag === "Nil"))) {
$__t4 = new Phpurs_Data1("Just", (object)["revInit" => $v1_3, "last" => ($v_2)->value0]);
} else {
$__tco_2 = ($v_2)->value1;
$__tco_3 = new Phpurs_Data2("Cons", ($v_2)->value0, $v1_3);
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
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
  $__local_var_2_5 = (($go_1_0)($lst_0))(new Phpurs_Data0("Nil"));
  if ((is_object($__local_var_2_5) && (($__local_var_2_5)->tag === "Just"))) {
$go_3_7 = null;
$go_3_7 = (function() use (&$go_3_7) {
  $__fn = function($v_4, $v1_5 = null) use (&$go_3_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_7"), recVars=["go_1_0","go_3_7"];
  while (true) {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$__t8 = $v_4;
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Cons"))) {
$__tco_9 = new Phpurs_Data2("Cons", ($v1_5)->value0, $v_4);
$__tco_10 = ($v1_5)->value1;
$v_4 = $__tco_9;
$v1_5 = $__tco_10;
continue ;
$__t8 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t8 = null;
};
};
$__res = $__t8;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t6 = new Phpurs_Data1("Just", (object)["init" => (($go_3_7)(new Phpurs_Data0("Nil")))((($__local_var_2_5)->value0)->revInit), "last" => (($__local_var_2_5)->value0)->last]);
} else {
$__t6 = new Phpurs_Data0("Nothing");
};
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_zipWith'] = function() { $v = (function() {
  $__fn = function($f_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_zipWith"), recVars=[];
  $go_3_0 = null;
  $go_3_0 = (function() use ($f_0, &$go_3_0) {
  $__fn = function($v_4, $v1_5 = null, $v2_6 = null) use ($f_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if ((is_object($v_4) && (($v_4)->tag === "Nil"))) {
$__t1 = $v2_6;
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$__t1 = $v2_6;
} else {
if (((is_object($v_4) && (($v_4)->tag === "Cons")) && (is_object($v1_5) && (($v1_5)->tag === "Cons")))) {
$__tco_2 = ($v_4)->value1;
$__tco_3 = ($v1_5)->value1;
$__tco_4 = new Phpurs_Data2("Cons", (($f_0)(($v_4)->value0))(($v1_5)->value0), $v2_6);
$v_4 = $__tco_2;
$v1_5 = $__tco_3;
$v2_6 = $__tco_4;
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $go_4_5 = null;
  $go_4_5 = (function() use (&$go_4_5) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_5"), recVars=["go_3_0","go_4_5"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t6 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_7 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_8 = ($v1_6)->value1;
$v_5 = $__tco_7;
$v1_6 = $__tco_8;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_4_5)(new Phpurs_Data0("Nil")))(((($go_3_0)($xs_1))($ys_2))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_zip'] = function() { $v = (function() {
  $__fn = function($xs_0, $ys_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_zip"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($v_3, $v1_4 = null, $v2_5 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Nil"))) {
$__t1 = $v2_5;
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "Nil"))) {
$__t1 = $v2_5;
} else {
if (((is_object($v_3) && (($v_3)->tag === "Cons")) && (is_object($v1_4) && (($v1_4)->tag === "Cons")))) {
$__tco_2 = ($v_3)->value1;
$__tco_3 = ($v1_4)->value1;
$__tco_4 = new Phpurs_Data2("Cons", new Phpurs_Data2("Tuple", ($v_3)->value0, ($v1_4)->value0), $v2_5);
$v_3 = $__tco_2;
$v1_4 = $__tco_3;
$v2_5 = $__tco_4;
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $go_3_5 = null;
  $go_3_5 = (function() use (&$go_3_5) {
  $__fn = function($v_4, $v1_5 = null) use (&$go_3_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_5"), recVars=["go_2_0","go_3_5"];
  while (true) {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$__t6 = $v_4;
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Cons"))) {
$__tco_7 = new Phpurs_Data2("Cons", ($v1_5)->value0, $v_4);
$__tco_8 = ($v1_5)->value1;
$v_4 = $__tco_7;
$v1_5 = $__tco_8;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go_3_5)(new Phpurs_Data0("Nil")))(((($go_2_0)($xs_0))($ys_1))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_zipWithA'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_zipWithA"), recVars=[];
  $sequence1_1_0 = (((($GLOBALS['Data_List_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Types_traversableList')))->traverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $go_5_1 = null;
  $go_5_1 = (function() use ($f_2, &$go_5_1) {
  $__fn = function($v_6, $v1_7 = null, $v2_8 = null) use ($f_2, &$go_5_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_5_1"), recVars=["go_5_1"];
  while (true) {
if ((is_object($v_6) && (($v_6)->tag === "Nil"))) {
$__t2 = $v2_8;
} else {
if ((is_object($v1_7) && (($v1_7)->tag === "Nil"))) {
$__t2 = $v2_8;
} else {
if (((is_object($v_6) && (($v_6)->tag === "Cons")) && (is_object($v1_7) && (($v1_7)->tag === "Cons")))) {
$__tco_3 = ($v_6)->value1;
$__tco_4 = ($v1_7)->value1;
$__tco_5 = new Phpurs_Data2("Cons", (($f_2)(($v_6)->value0))(($v1_7)->value0), $v2_8);
$v_6 = $__tco_3;
$v1_7 = $__tco_4;
$v2_8 = $__tco_5;
continue ;
$__t2 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
$__res = $__t2;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $go_6_6 = null;
  $go_6_6 = (function() use (&$go_6_6) {
  $__fn = function($v_7, $v1_8 = null) use (&$go_6_6, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_6"), recVars=["go_5_1","go_6_6"];
  while (true) {
if ((is_object($v1_8) && (($v1_8)->tag === "Nil"))) {
$__t7 = $v_7;
} else {
if ((is_object($v1_8) && (($v1_8)->tag === "Cons"))) {
$__tco_8 = new Phpurs_Data2("Cons", ($v1_8)->value0, $v_7);
$__tco_9 = ($v1_8)->value1;
$v_7 = $__tco_8;
$v1_8 = $__tco_9;
continue ;
$__t7 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
$__res = $__t7;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($sequence1_1_0)((($go_6_6)(new Phpurs_Data0("Nil")))(((($go_5_1)($xs_3))($ys_4))(new Phpurs_Data0("Nil"))));
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
\PhpursThunks::$thunks['Data_List_range'] = function() { $v = (function() {
  $__fn = function($start_0, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_range"), recVars=[];
  if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))($start_0))($end_1)) {
$__t7 = new Phpurs_Data2("Cons", $start_0, new Phpurs_Data0("Nil"));
} else {
$go_2_0 = null;
$go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($s_3, $e_4 = null, $step_5 = null, $rest_6 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if (((($GLOBALS['Data_Eq_eqIntImpl'] ?? \PhpursThunks::eval('Data_Eq_eqIntImpl')))($s_3))($e_4)) {
$__t5 = new Phpurs_Data2("Cons", $s_3, $rest_6);
} else {
$__tco_1 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($s_3))($step_5);
$__tco_2 = $e_4;
$__tco_3 = $step_5;
$__tco_4 = new Phpurs_Data2("Cons", $s_3, $rest_6);
$s_3 = $__tco_1;
$e_4 = $__tco_2;
$step_5 = $__tco_3;
$rest_6 = $__tco_4;
continue ;
$__t5 = null;
};
$__res = $__t5;
goto __end;;
};
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();
if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($start_0))($end_1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($start_0))($end_1))->tag === "GT"))) {
$__t6 = 1;
} else {
$__t6 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(0))(1);
};
$__t7 = (((($go_2_0)($end_1))($start_0))($__t6))(new Phpurs_Data0("Nil"));
};
  $__res = $__t7;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_partition'] = function() { $v = (function() {
  $__fn = function($p_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_partition"), recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)((function() use ($p_0) {
  $__fn = function($x_2, $v_3 = null) use ($p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (($p_0)($x_2)) {
$__t0 = (object)["no" => ($v_3)->no, "yes" => new Phpurs_Data2("Cons", $x_2, ($v_3)->yes)];
} else {
$__t0 = (object)["no" => new Phpurs_Data2("Cons", $x_2, ($v_3)->no), "yes" => ($v_3)->yes];
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["no" => new Phpurs_Data0("Nil"), "yes" => new Phpurs_Data0("Nil")]))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_null'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_null"), recVars=[];
  $__res = (is_object($v_0) && (($v_0)->tag === "Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_nubBy'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_nubBy"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null, $v2_4 = null) use (&$go_1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v2_4) && (($v2_4)->tag === "Nil"))) {
$__t1 = $v1_3;
} else {
if ((is_object($v2_4) && (($v2_4)->tag === "Cons"))) {
$v3_5_2 = (((($GLOBALS['Data_List_Internal_insertAndLookupBy'] ?? \PhpursThunks::eval('Data_List_Internal_insertAndLookupBy')))($p_0))(($v2_4)->value0))($v_2);
if (($v3_5_2)->found) {
$__tco_7 = ($v3_5_2)->result;
$__tco_8 = $v1_3;
$__tco_9 = ($v2_4)->value1;
$v_2 = $__tco_7;
$v1_3 = $__tco_8;
$v2_4 = $__tco_9;
continue ;
$__t6 = null;
} else {
$__tco_3 = ($v3_5_2)->result;
$__tco_4 = new Phpurs_Data2("Cons", ($v2_4)->value0, $v1_3);
$__tco_5 = ($v2_4)->value1;
$v_2 = $__tco_3;
$v1_3 = $__tco_4;
$v2_4 = $__tco_5;
continue ;
$__t6 = null;
};
$__t1 = $__t6;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_reverse'] ?? \PhpursThunks::eval('Data_List_reverse'))))((($go_1_0)(new Phpurs_Data0("Leaf")))(new Phpurs_Data0("Nil")));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_nub'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_nub"), recVars=[];
  $__res = (($GLOBALS['Data_List_nubBy'] ?? \PhpursThunks::eval('Data_List_nubBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_newtypePattern'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_mapMaybe'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_mapMaybe"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use ($f_0, &$go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($f_0, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$go_4_2 = null;
$go_4_2 = (function() use (&$go_4_2) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_2"), recVars=["go_1_0","go_4_2"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t3 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_5 = ($v1_6)->value1;
$v_5 = $__tco_4;
$v1_6 = $__tco_5;
continue ;
$__t3 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = (($go_4_2)(new Phpurs_Data0("Nil")))($v_2);
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
$v2_4_6 = ($f_0)(($v1_3)->value0);
if ((is_object($v2_4_6) && (($v2_4_6)->tag === "Nothing"))) {
$__tco_8 = $v_2;
$__tco_9 = ($v1_3)->value1;
$v_2 = $__tco_8;
$v1_3 = $__tco_9;
continue ;
$__t7 = null;
} else {
if ((is_object($v2_4_6) && (($v2_4_6)->tag === "Just"))) {
$__tco_10 = new Phpurs_Data2("Cons", ($v2_4_6)->value0, $v_2);
$__tco_11 = ($v1_3)->value1;
$v_2 = $__tco_10;
$v1_3 = $__tco_11;
continue ;
$__t7 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t7 = null;
};
};
$__t1 = $__t7;
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
  $__res = ($go_1_0)(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_manyRec'] = function() { $v = (function() {
  $__fn = function($dictMonadRec_0, $dictAlternative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_manyRec"), recVars=[];
  $Alt0_2_0 = (((($dictAlternative_1)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_3_1 = (($dictAlternative_1)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($p_4) use ($Alt0_2_0, $__local_var_3_1, $dictMonadRec_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictMonadRec_0)->tailRecM)(function($acc_5) use ($Alt0_2_0, $__local_var_3_1, $dictMonadRec_0, $p_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((((($dictMonadRec_0)->Monad0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)(((($Alt0_2_0)->alt)(((((($Alt0_2_0)->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($GLOBALS['Control_Monad_Rec_Class_Loop'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_Loop'))))($p_4)))((($__local_var_3_1)->pure)(new Phpurs_Data1("Done", ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')))))))(function($aa_6) use ($__local_var_3_1, $acc_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if ((is_object($aa_6) && (($aa_6)->tag === "Loop"))) {
$__t2 = new Phpurs_Data1("Loop", new Phpurs_Data2("Cons", ($aa_6)->value0, $acc_5));
} else {
if ((is_object($aa_6) && (($aa_6)->tag === "Done"))) {
$go_7_3 = null;
$go_7_3 = (function() use (&$go_7_3) {
  $__fn = function($v_8, $v1_9 = null) use (&$go_7_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_7_3"), recVars=["go_7_3"];
  while (true) {
if ((is_object($v1_9) && (($v1_9)->tag === "Nil"))) {
$__t4 = $v_8;
} else {
if ((is_object($v1_9) && (($v1_9)->tag === "Cons"))) {
$__tco_5 = new Phpurs_Data2("Cons", ($v1_9)->value0, $v_8);
$__tco_6 = ($v1_9)->value1;
$v_8 = $__tco_5;
$v1_9 = $__tco_6;
continue ;
$__t4 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__res = $__t4;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t2 = new Phpurs_Data1("Done", (($go_7_3)(new Phpurs_Data0("Nil")))($acc_5));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = (($__local_var_3_1)->pure)($__t2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data0("Nil"));
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
\PhpursThunks::$thunks['Data_List_someRec'] = function() { $v = (function() {
  $__fn = function($dictMonadRec_0, $dictAlternative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_someRec"), recVars=[];
  $manyRec2_2_0 = ((($GLOBALS['Data_List_manyRec'] ?? \PhpursThunks::eval('Data_List_manyRec')))($dictMonadRec_0))($dictAlternative_1);
  $__res = function($v_3) use ($dictAlternative_1, $manyRec2_2_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((((($dictAlternative_1)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->apply)(((((((((($dictAlternative_1)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))($v_3)))(($manyRec2_2_0)($v_3));
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
\PhpursThunks::$thunks['Data_List_some'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_some"), recVars=["Data_List_some","Data_List_many"];
  while (true) {
$__res = ((((((($dictAlternative_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->apply)(((((((((($dictAlternative_0)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))($v_2)))((($dictLazy_1)->defer)(function($v1_3) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_some","Data_List_many"];
  $__res = (((($GLOBALS['Data_List_many'] ?? \PhpursThunks::eval('Data_List_many')))($dictAlternative_0))($dictLazy_1))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_many'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_many"), recVars=["Data_List_some","Data_List_many"];
  while (true) {
$__res = ((((((($dictAlternative_0)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->alt)((((($GLOBALS['Data_List_some'] ?? \PhpursThunks::eval('Data_List_some')))($dictAlternative_0))($dictLazy_1))($v_2)))((((($dictAlternative_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(new Phpurs_Data0("Nil")));
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_length'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($b_1, $v_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = $b_1;
} else {
if ((is_object($v_2) && (($v_2)->tag === "Cons"))) {
$__tco_2 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($b_1))(1);
$__tco_3 = ($v_2)->value1;
$b_1 = $__tco_2;
$v_2 = $__tco_3;
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
return ($go_0_0)(0);
})(); return $v; };
\PhpursThunks::$thunks['Data_List_last'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_last"), recVars=["Data_List_last"];
  while (true) {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
if ((is_object(($v_0)->value1) && ((($v_0)->value1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data1("Just", ($v_0)->value0);
} else {
$__tco_1 = ($v_0)->value1;
$v_0 = $__tco_1;
continue ;
$__t2 = null;
};
$__t0 = $__t2;
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_insertBy'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_insertBy"), recVars=["Data_List_insertBy"];
  while (true) {
if ((is_object($v2_2) && (($v2_2)->tag === "Nil"))) {
$__t0 = new Phpurs_Data2("Cons", $v1_1, new Phpurs_Data0("Nil"));
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Cons"))) {
if ((is_object((($v_0)($v1_1))(($v2_2)->value0)) && (((($v_0)($v1_1))(($v2_2)->value0))->tag === "GT"))) {
$__t1 = new Phpurs_Data2("Cons", ($v2_2)->value0, (((($GLOBALS['Data_List_insertBy'] ?? \PhpursThunks::eval('Data_List_insertBy')))($v_0))($v1_1))(($v2_2)->value1));
} else {
$__t1 = new Phpurs_Data2("Cons", $v1_1, $v2_2);
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_insertAt'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_insertAt"), recVars=["Data_List_insertAt"];
  while (true) {
if (($v_0 === 0)) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", $v1_1, $v2_2));
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Cons"))) {
$__local_var_3_1 = (((($GLOBALS['Data_List_insertAt'] ?? \PhpursThunks::eval('Data_List_insertAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1)))($v1_1))(($v2_2)->value1);
if ((is_object($__local_var_3_1) && (($__local_var_3_1)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v2_2)->value0, ($__local_var_3_1)->value0));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
$__t0 = $__t2;
} else {
$__t0 = new Phpurs_Data0("Nothing");
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
\PhpursThunks::$thunks['Data_List_insert'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_insert"), recVars=[];
  $__res = (($GLOBALS['Data_List_insertBy'] ?? \PhpursThunks::eval('Data_List_insertBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_init'] = function() { $v = function($lst_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_init"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_List_unsnoc'] ?? \PhpursThunks::eval('Data_List_unsnoc')))($lst_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->value0)->init);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_index'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_index"), recVars=["Data_List_index"];
  while (true) {
if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
switch ($v1_1) {
case 0:
$__t3 = new Phpurs_Data1("Just", ($v_0)->value0);
break;
default:
$__tco_1 = ($v_0)->value1;
$__tco_2 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v1_1))(1);
$v_0 = $__tco_1;
$v1_1 = $__tco_2;
continue 2;
$__t3 = null;
break;
};
$__t0 = $__t3;
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
\PhpursThunks::$thunks['Data_List_head'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_head"), recVars=[];
  if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
$__t0 = new Phpurs_Data1("Just", ($v_0)->value0);
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
\PhpursThunks::$thunks['Data_List_transpose'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_transpose"), recVars=["Data_List_transpose"];
  while (true) {
if ((is_object($v_0) && (($v_0)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_0) && (($v_0)->tag === "Cons"))) {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Nil"))) {
$__tco_2 = ($v_0)->value1;
$v_0 = $__tco_2;
continue ;
$__t1 = null;
} else {
if ((is_object(($v_0)->value0) && ((($v_0)->value0)->tag === "Cons"))) {
$go_1_3 = null;
$go_1_3 = (function() use (&$go_1_3) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_3"), recVars=["Data_List_transpose","go_1_3"];
  while (true) {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$go_4_5 = null;
$go_4_5 = (function() use (&$go_4_5) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_5"), recVars=["Data_List_transpose","go_1_3","go_4_5"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t6 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_7 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_8 = ($v1_6)->value1;
$v_5 = $__tco_7;
$v1_6 = $__tco_8;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__res = $__t6;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t4 = (($go_4_5)(new Phpurs_Data0("Nil")))($v_2);
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
if ((is_object(($v1_3)->value0) && ((($v1_3)->value0)->tag === "Nil"))) {
$__tco_10 = $v_2;
$__tco_11 = ($v1_3)->value1;
$v_2 = $__tco_10;
$v1_3 = $__tco_11;
continue ;
$__t9 = null;
} else {
if ((is_object(($v1_3)->value0) && ((($v1_3)->value0)->tag === "Cons"))) {
$__tco_12 = new Phpurs_Data2("Cons", (($v1_3)->value0)->value0, $v_2);
$__tco_13 = ($v1_3)->value1;
$v_2 = $__tco_12;
$v1_3 = $__tco_13;
continue ;
$__t9 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t9 = null;
};
};
$__t4 = $__t9;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t4 = null;
};
};
$__res = $__t4;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$go_1_14 = null;
$go_1_14 = (function() use (&$go_1_14) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_14, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_14"), recVars=["Data_List_transpose","go_1_14"];
  while (true) {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$go_4_16 = null;
$go_4_16 = (function() use (&$go_4_16) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_16, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_16"), recVars=["Data_List_transpose","go_1_14","go_4_16"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t17 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_18 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_19 = ($v1_6)->value1;
$v_5 = $__tco_18;
$v1_6 = $__tco_19;
continue ;
$__t17 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t17 = null;
};
};
$__res = $__t17;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t15 = (($go_4_16)(new Phpurs_Data0("Nil")))($v_2);
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
if ((is_object(($v1_3)->value0) && ((($v1_3)->value0)->tag === "Nil"))) {
$__tco_21 = $v_2;
$__tco_22 = ($v1_3)->value1;
$v_2 = $__tco_21;
$v1_3 = $__tco_22;
continue ;
$__t20 = null;
} else {
if ((is_object(($v1_3)->value0) && ((($v1_3)->value0)->tag === "Cons"))) {
$__tco_23 = new Phpurs_Data2("Cons", (($v1_3)->value0)->value1, $v_2);
$__tco_24 = ($v1_3)->value1;
$v_2 = $__tco_23;
$v1_3 = $__tco_24;
continue ;
$__t20 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t20 = null;
};
};
$__t15 = $__t20;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t15 = null;
};
};
$__res = $__t15;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = new Phpurs_Data2("Cons", new Phpurs_Data2("Cons", (($v_0)->value0)->value0, (($go_1_3)(new Phpurs_Data0("Nil")))(($v_0)->value1)), (($GLOBALS['Data_List_transpose'] ?? \PhpursThunks::eval('Data_List_transpose')))(new Phpurs_Data2("Cons", (($v_0)->value0)->value1, (($go_1_14)(new Phpurs_Data0("Nil")))(($v_0)->value1))));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_groupBy'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_groupBy"), recVars=["Data_List_groupBy"];
  while (true) {
if ((is_object($v1_1) && (($v1_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "Cons"))) {
$v2_2_1 = ((($GLOBALS['Data_List_span'] ?? \PhpursThunks::eval('Data_List_span')))(($v_0)(($v1_1)->value0)))(($v1_1)->value1);
$__t0 = new Phpurs_Data2("Cons", new Phpurs_Data2("NonEmpty", ($v1_1)->value0, ($v2_2_1)->init), ((($GLOBALS['Data_List_groupBy'] ?? \PhpursThunks::eval('Data_List_groupBy')))($v_0))(($v2_2_1)->rest));
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
\PhpursThunks::$thunks['Data_List_groupAllBy'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_groupAllBy"), recVars=[];
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_List_groupBy'] ?? \PhpursThunks::eval('Data_List_groupBy')))((function() use ($p_0) {
  $__fn = function($x_1, $y_2 = null) use ($p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (is_object((($p_0)($x_1))($y_2)) && (((($p_0)($x_1))($y_2))->tag === "EQ"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))((($GLOBALS['Data_List_sortBy'] ?? \PhpursThunks::eval('Data_List_sortBy')))($p_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_group'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_group"), recVars=[];
  $__res = (($GLOBALS['Data_List_groupBy'] ?? \PhpursThunks::eval('Data_List_groupBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_groupAll'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_groupAll"), recVars=[];
  $compare_1_0 = ($dictOrd_0)->compare;
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($GLOBALS['Data_List_group'] ?? \PhpursThunks::eval('Data_List_group')))((($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))))(function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_List_sortBy'] ?? \PhpursThunks::eval('Data_List_sortBy')))($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_fromFoldable'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_fromFoldable"), recVars=[];
  $__res = ((($dictFoldable_0)->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_foldM'] = function() { $v = (function() {
  $__fn = function($dictMonad_0, $v_1 = null, $v1_2 = null, $v2_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_foldM"), recVars=["Data_List_foldM"];
  while (true) {
if ((is_object($v2_3) && (($v2_3)->tag === "Nil"))) {
$__t0 = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($v1_2);
} else {
if ((is_object($v2_3) && (($v2_3)->tag === "Cons"))) {
$__local_var_4_1 = ($v2_3)->value1;
$__t0 = ((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)((($v_1)($v1_2))(($v2_3)->value0)))(function($b__prime___5) use ($__local_var_4_1, $dictMonad_0, $v_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_foldM"];
  $__res = ((((($GLOBALS['Data_List_foldM'] ?? \PhpursThunks::eval('Data_List_foldM')))($dictMonad_0))($v_1))($b__prime___5))($__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
\PhpursThunks::$thunks['Data_List_findIndex'] = function() { $v = function($fn_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_findIndex"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use ($fn_0, &$go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use ($fn_0, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
if (($fn_0)(($v1_3)->value0)) {
$__t4 = new Phpurs_Data1("Just", $v_2);
} else {
$__tco_2 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($v_2))(1);
$__tco_3 = ($v1_3)->value1;
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
continue ;
$__t4 = null;
};
$__t1 = $__t4;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("Nothing");
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
  $__res = ($go_1_0)(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_findLastIndex'] = function() { $v = (function() {
  $__fn = function($fn_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_findLastIndex"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (function() use (&$go_2_0) {
  $__fn = function($v_3, $v1_4 = null) use (&$go_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_2_0"), recVars=["go_2_0"];
  while (true) {
if ((is_object($v1_4) && (($v1_4)->tag === "Nil"))) {
$__t1 = $v_3;
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "Cons"))) {
$__tco_2 = new Phpurs_Data2("Cons", ($v1_4)->value0, $v_3);
$__tco_3 = ($v1_4)->value1;
$v_3 = $__tco_2;
$v1_4 = $__tco_3;
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
  $__local_var_2_0 = ((($GLOBALS['Data_List_findIndex'] ?? \PhpursThunks::eval('Data_List_findIndex')))($fn_0))((($go_2_0)(new Phpurs_Data0("Nil")))($xs_1));
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t5 = new Phpurs_Data1("Just", ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_List_length'] ?? \PhpursThunks::eval('Data_List_length')))($xs_1)))(1)))(($__local_var_2_0)->value0));
} else {
$__t5 = new Phpurs_Data0("Nothing");
};
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_filterM'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_filterM"), recVars=["Data_List_filterM"];
  while (true) {
$__local_var_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__local_var_2_1 = (($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($v_3, $v1_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_filterM"];
  if ((is_object($v1_4) && (($v1_4)->tag === "Nil"))) {
$__t2 = (($__local_var_1_0)->pure)(new Phpurs_Data0("Nil"));
} else {
if ((is_object($v1_4) && (($v1_4)->tag === "Cons"))) {
$__local_var_5_3 = ($v1_4)->value0;
$__local_var_6_4 = ($v1_4)->value1;
$__t2 = ((($__local_var_2_1)->bind)(($v_3)($__local_var_5_3)))(function($b_7) use ($__local_var_1_0, $__local_var_2_1, $__local_var_5_3, $__local_var_6_4, $dictMonad_0, $v_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_filterM"];
  $__res = ((($__local_var_2_1)->bind)((((($GLOBALS['Data_List_filterM'] ?? \PhpursThunks::eval('Data_List_filterM')))($dictMonad_0))($v_3))($__local_var_6_4)))(function($xs__prime___8) use ($__local_var_1_0, $__local_var_5_3, $b_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_filterM"];
  if ($b_7) {
$__t5 = new Phpurs_Data2("Cons", $__local_var_5_3, $xs__prime___8);
} else {
$__t5 = $xs__prime___8;
};
  $__res = (($__local_var_1_0)->pure)($__t5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
})();
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_filter'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_filter"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$go_4_2 = null;
$go_4_2 = (function() use (&$go_4_2) {
  $__fn = function($v_5, $v1_6 = null) use (&$go_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_4_2"), recVars=["go_1_0","go_4_2"];
  while (true) {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t3 = $v_5;
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_6)->value0, $v_5);
$__tco_5 = ($v1_6)->value1;
$v_5 = $__tco_4;
$v1_6 = $__tco_5;
continue ;
$__t3 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = (($go_4_2)(new Phpurs_Data0("Nil")))($v_2);
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
if (($p_0)(($v1_3)->value0)) {
$__tco_9 = new Phpurs_Data2("Cons", ($v1_3)->value0, $v_2);
$__tco_10 = ($v1_3)->value1;
$v_2 = $__tco_9;
$v1_3 = $__tco_10;
continue ;
$__t8 = null;
} else {
$__tco_6 = $v_2;
$__tco_7 = ($v1_3)->value1;
$v_2 = $__tco_6;
$v1_3 = $__tco_7;
continue ;
$__t8 = null;
};
$__t1 = $__t8;
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
  $__res = ($go_1_0)(new Phpurs_Data0("Nil"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_intersectBy'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_intersectBy"), recVars=[];
  if ((is_object($v1_1) && (($v1_1)->tag === "Nil"))) {
$__t11 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Nil"))) {
$__t11 = new Phpurs_Data0("Nil");
} else {
$go_3_0 = null;
$go_3_0 = (function() use (&$go_3_0, $v2_2, $v_0) {
  $__fn = function($v_4, $v1_5 = null) use (&$go_3_0, $v2_2, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$go_6_2 = null;
$go_6_2 = (function() use (&$go_6_2) {
  $__fn = function($v_7, $v1_8 = null) use (&$go_6_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_2"), recVars=["go_3_0","go_6_2"];
  while (true) {
if ((is_object($v1_8) && (($v1_8)->tag === "Nil"))) {
$__t3 = $v_7;
} else {
if ((is_object($v1_8) && (($v1_8)->tag === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_8)->value0, $v_7);
$__tco_5 = ($v1_8)->value1;
$v_7 = $__tco_4;
$v1_8 = $__tco_5;
continue ;
$__t3 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = (($go_6_2)(new Phpurs_Data0("Nil")))($v_4);
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Cons"))) {
if (((($GLOBALS['Data_List_any'] ?? \PhpursThunks::eval('Data_List_any')))(($v_0)(($v1_5)->value0)))($v2_2)) {
$__tco_9 = new Phpurs_Data2("Cons", ($v1_5)->value0, $v_4);
$__tco_10 = ($v1_5)->value1;
$v_4 = $__tco_9;
$v1_5 = $__tco_10;
continue ;
$__t8 = null;
} else {
$__tco_6 = $v_4;
$__tco_7 = ($v1_5)->value1;
$v_4 = $__tco_6;
$v1_5 = $__tco_7;
continue ;
$__t8 = null;
};
$__t1 = $__t8;
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
$__t11 = (($go_3_0)(new Phpurs_Data0("Nil")))($v1_1);
};
};
  $__res = $__t11;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_intersect'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_intersect"), recVars=[];
  $__res = (($GLOBALS['Data_List_intersectBy'] ?? \PhpursThunks::eval('Data_List_intersectBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_nubByEq'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_nubByEq"), recVars=["Data_List_nubByEq"];
  while (true) {
if ((is_object($v1_1) && (($v1_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "Cons"))) {
$__local_var_2_1 = ($v1_1)->value0;
$go_3_2 = null;
$go_3_2 = (function() use ($__local_var_2_1, &$go_3_2, $v_0) {
  $__fn = function($v_4, $v1_5 = null) use ($__local_var_2_1, &$go_3_2, $v_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_2"), recVars=["Data_List_nubByEq","go_3_2"];
  while (true) {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$go_6_4 = null;
$go_6_4 = (function() use (&$go_6_4) {
  $__fn = function($v_7, $v1_8 = null) use (&$go_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_6_4"), recVars=["Data_List_nubByEq","go_3_2","go_6_4"];
  while (true) {
if ((is_object($v1_8) && (($v1_8)->tag === "Nil"))) {
$__t5 = $v_7;
} else {
if ((is_object($v1_8) && (($v1_8)->tag === "Cons"))) {
$__tco_6 = new Phpurs_Data2("Cons", ($v1_8)->value0, $v_7);
$__tco_7 = ($v1_8)->value1;
$v_7 = $__tco_6;
$v1_8 = $__tco_7;
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
$__t3 = (($go_6_4)(new Phpurs_Data0("Nil")))($v_4);
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Cons"))) {
if ((($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot')))((($v_0)($__local_var_2_1))(($v1_5)->value0))) {
$__tco_11 = new Phpurs_Data2("Cons", ($v1_5)->value0, $v_4);
$__tco_12 = ($v1_5)->value1;
$v_4 = $__tco_11;
$v1_5 = $__tco_12;
continue ;
$__t10 = null;
} else {
$__tco_8 = $v_4;
$__tco_9 = ($v1_5)->value1;
$v_4 = $__tco_8;
$v1_5 = $__tco_9;
continue ;
$__t10 = null;
};
$__t3 = $__t10;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t0 = new Phpurs_Data2("Cons", $__local_var_2_1, ((($GLOBALS['Data_List_nubByEq'] ?? \PhpursThunks::eval('Data_List_nubByEq')))($v_0))((($go_3_2)(new Phpurs_Data0("Nil")))(($v1_1)->value1)));
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
\PhpursThunks::$thunks['Data_List_nubEq'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_nubEq"), recVars=[];
  $__res = (($GLOBALS['Data_List_nubByEq'] ?? \PhpursThunks::eval('Data_List_nubByEq')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_eqPattern'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_eqPattern"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Types_eq1List')))->eq1)($dictEq_0))($x_1))($y_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_ordPattern'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_ordPattern"), recVars=[];
  $__local_var_1_0 = (($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eqPattern1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Types_eq1List')))->eq1)($__local_var_1_0))($x_2))($y_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__res = (object)["compare" => (function() use ($dictOrd_0) {
  $__fn = function($x_3, $y_4 = null) use ($dictOrd_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Types_ordList'] ?? \PhpursThunks::eval('Data_List_Types_ordList')))($dictOrd_0))->compare)($x_3))($y_4);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Eq0" => function($dollar__unused_3) use ($eqPattern1_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = $eqPattern1_2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_elemLastIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_elemLastIndex"), recVars=[];
  $__res = (($GLOBALS['Data_List_findLastIndex'] ?? \PhpursThunks::eval('Data_List_findLastIndex')))(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictEq_0)->eq)($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_elemIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_elemIndex"), recVars=[];
  $__res = (($GLOBALS['Data_List_findIndex'] ?? \PhpursThunks::eval('Data_List_findIndex')))(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($dictEq_0)->eq)($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_dropWhile'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_dropWhile"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use (&$go_1_0, $p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if (((is_object($v_2) && (($v_2)->tag === "Cons")) && ($p_0)(($v_2)->value0))) {
$__tco_2 = ($v_2)->value1;
$v_2 = $__tco_2;
continue ;
$__t1 = null;
} else {
$__t1 = $v_2;
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
\PhpursThunks::$thunks['Data_List_dropEnd'] = function() { $v = (function() {
  $__fn = function($n_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_dropEnd"), recVars=[];
  $__res = ((($GLOBALS['Data_List_take'] ?? \PhpursThunks::eval('Data_List_take')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_List_length'] ?? \PhpursThunks::eval('Data_List_length')))($xs_1)))($n_0)))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_drop'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_drop"), recVars=["Data_List_drop"];
  while (true) {
if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($v_0))(1))->tag === "LT"))) {
$__t0 = $v1_1;
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "Cons"))) {
$__tco_1 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1);
$__tco_2 = ($v1_1)->value1;
$v_0 = $__tco_1;
$v1_1 = $__tco_2;
continue ;
$__t0 = null;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_slice'] = function() { $v = (function() {
  $__fn = function($start_0, $end_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_slice"), recVars=[];
  $__res = ((($GLOBALS['Data_List_take'] ?? \PhpursThunks::eval('Data_List_take')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($end_1))($start_0)))(((($GLOBALS['Data_List_drop'] ?? \PhpursThunks::eval('Data_List_drop')))($start_0))($xs_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_takeEnd'] = function() { $v = (function() {
  $__fn = function($n_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_takeEnd"), recVars=[];
  $__res = ((($GLOBALS['Data_List_drop'] ?? \PhpursThunks::eval('Data_List_drop')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_List_length'] ?? \PhpursThunks::eval('Data_List_length')))($xs_1)))($n_0)))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_deleteBy'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_deleteBy"), recVars=["Data_List_deleteBy"];
  while (true) {
if ((is_object($v2_2) && (($v2_2)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v2_2) && (($v2_2)->tag === "Cons"))) {
if ((($v_0)($v1_1))(($v2_2)->value0)) {
$__t1 = ($v2_2)->value1;
} else {
$__t1 = new Phpurs_Data2("Cons", ($v2_2)->value0, (((($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))($v_0))($v1_1))(($v2_2)->value1));
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_unionBy'] = function() { $v = (function() {
  $__fn = function($eq2_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_unionBy"), recVars=[];
  $go_3_0 = null;
  $go_3_0 = (function() use ($eq2_0, &$go_3_0) {
  $__fn = function($b_4, $v_5 = null) use ($eq2_0, &$go_3_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
if ((is_object($v_5) && (($v_5)->tag === "Nil"))) {
$__t1 = $b_4;
} else {
if ((is_object($v_5) && (($v_5)->tag === "Cons"))) {
$__tco_2 = (((($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))($eq2_0))(($v_5)->value0))($b_4);
$__tco_3 = ($v_5)->value1;
$b_4 = $__tco_2;
$v_5 = $__tco_3;
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
  $__res = ((((($GLOBALS['Data_List_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Types_Cons'] ?? \PhpursThunks::eval('Data_List_Types_Cons'))))((($go_3_0)(((($GLOBALS['Data_List_nubByEq'] ?? \PhpursThunks::eval('Data_List_nubByEq')))($eq2_0))($ys_2)))($xs_1)))($xs_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_union'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_union"), recVars=[];
  $__res = (($GLOBALS['Data_List_unionBy'] ?? \PhpursThunks::eval('Data_List_unionBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_deleteAt'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_deleteAt"), recVars=["Data_List_deleteAt"];
  while (true) {
if ((is_object($v1_1) && (($v1_1)->tag === "Cons"))) {
switch ($v_0) {
case 0:
$__t3 = new Phpurs_Data1("Just", ($v1_1)->value1);
break;
default:
$__local_var_2_1 = ((($GLOBALS['Data_List_deleteAt'] ?? \PhpursThunks::eval('Data_List_deleteAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1)))(($v1_1)->value1);
if ((is_object($__local_var_2_1) && (($__local_var_2_1)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v1_1)->value0, ($__local_var_2_1)->value0));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
$__t3 = $__t2;
break;
};
$__t0 = $__t3;
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_delete'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_delete"), recVars=[];
  $__res = (($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_difference'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_difference"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use ($dictEq_0, &$go_1_0) {
  $__fn = function($b_2, $v_3 = null) use ($dictEq_0, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if ((is_object($v_3) && (($v_3)->tag === "Nil"))) {
$__t1 = $b_2;
} else {
if ((is_object($v_3) && (($v_3)->tag === "Cons"))) {
$__tco_2 = (((($GLOBALS['Data_List_deleteBy'] ?? \PhpursThunks::eval('Data_List_deleteBy')))(($dictEq_0)->eq))(($v_3)->value0))($b_2);
$__tco_3 = ($v_3)->value1;
$b_2 = $__tco_2;
$v_3 = $__tco_3;
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
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_concatMap'] = function() { $v = (function() {
  $__fn = function($b_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_concatMap"), recVars=[];
  $__res = (((($GLOBALS['Data_List_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Types_bindList')))->bind)($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_concat'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_concat"), recVars=[];
  $__res = (((($GLOBALS['Data_List_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Types_bindList')))->bind)($v_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_catMaybes'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = (function() use (&$go_0_0) {
  $__fn = function($v_1, $v1_2 = null) use (&$go_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($v1_2) && (($v1_2)->tag === "Nil"))) {
$go_3_2 = null;
$go_3_2 = (function() use (&$go_3_2) {
  $__fn = function($v_4, $v1_5 = null) use (&$go_3_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_2"), recVars=["go_0_0","go_3_2"];
  while (true) {
if ((is_object($v1_5) && (($v1_5)->tag === "Nil"))) {
$__t3 = $v_4;
} else {
if ((is_object($v1_5) && (($v1_5)->tag === "Cons"))) {
$__tco_4 = new Phpurs_Data2("Cons", ($v1_5)->value0, $v_4);
$__tco_5 = ($v1_5)->value1;
$v_4 = $__tco_4;
$v1_5 = $__tco_5;
continue ;
$__t3 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
};
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
$__t1 = (($go_3_2)(new Phpurs_Data0("Nil")))($v_1);
} else {
if ((is_object($v1_2) && (($v1_2)->tag === "Cons"))) {
if ((is_object(($v1_2)->value0) && ((($v1_2)->value0)->tag === "Nothing"))) {
$__tco_7 = $v_1;
$__tco_8 = ($v1_2)->value1;
$v_1 = $__tco_7;
$v1_2 = $__tco_8;
continue ;
$__t6 = null;
} else {
if ((is_object(($v1_2)->value0) && ((($v1_2)->value0)->tag === "Just"))) {
$__tco_9 = new Phpurs_Data2("Cons", (($v1_2)->value0)->value0, $v_1);
$__tco_10 = ($v1_2)->value1;
$v_1 = $__tco_9;
$v1_2 = $__tco_10;
continue ;
$__t6 = null;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t6 = null;
};
};
$__t1 = $__t6;
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
return ($go_0_0)(new Phpurs_Data0("Nil"));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_alterAt'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_alterAt"), recVars=["Data_List_alterAt"];
  while (true) {
if ((is_object($v2_2) && (($v2_2)->tag === "Cons"))) {
switch ($v_0) {
case 0:
$v3_3_4 = ($v1_1)(($v2_2)->value0);
if ((is_object($v3_3_4) && (($v3_3_4)->tag === "Nothing"))) {
$__t5 = ($v2_2)->value1;
} else {
if ((is_object($v3_3_4) && (($v3_3_4)->tag === "Just"))) {
$__t5 = new Phpurs_Data2("Cons", ($v3_3_4)->value0, ($v2_2)->value1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__t3 = new Phpurs_Data1("Just", $__t5);
break;
default:
$__local_var_3_1 = (((($GLOBALS['Data_List_alterAt'] ?? \PhpursThunks::eval('Data_List_alterAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1)))($v1_1))(($v2_2)->value1);
if ((is_object($__local_var_3_1) && (($__local_var_3_1)->tag === "Just"))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Cons", ($v2_2)->value0, ($__local_var_3_1)->value0));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
$__t3 = $__t2;
break;
};
$__t0 = $__t3;
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
$__res = $__t0;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_modifyAt'] = function() { $v = (function() {
  $__fn = function($n_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_modifyAt"), recVars=[];
  $__res = ((($GLOBALS['Data_List_alterAt'] ?? \PhpursThunks::eval('Data_List_alterAt')))($n_0))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))($f_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };














































































