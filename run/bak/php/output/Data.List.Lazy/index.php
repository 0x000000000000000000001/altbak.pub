<?php

namespace Data\List\Lazy;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Lazy, Data.List.Internal, Data.List.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Functor, Data.HeytingAlgebra, Data.Lazy, Data.List.Internal, Data.List.Lazy, Data.List.Lazy.Types, Data.Maybe, Data.Newtype, Data.NonEmpty, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Show, Data.Traversable, Data.Tuple, Data.Unfoldable, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Internal/index.php';
require_once __DIR__ . '/../Data.List.Lazy/index.php';
require_once __DIR__ . '/../Data.List.Lazy.Types/index.php';
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
\PhpursThunks::$thunks['Data_List_Lazy_tailRecM2'] = function() { $v = (function() {
  $__fn = function($f_0, $a_1 = null, $b_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_tailRecM2"), recVars=[];
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
\PhpursThunks::$thunks['Data_List_Lazy_any'] = function() { $v = ((($GLOBALS['Data_Foldable_any'] ?? \PhpursThunks::eval('Data_Foldable_any')))(($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList'))))(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_heytingAlgebraBoolean'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_Pattern'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_Pattern"), recVars=[];
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_zipWith'] = function() { $v = (function() {
  $__fn = function($f_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_zipWith"), recVars=["Data_List_Lazy_zipWith"];
  while (true) {
$__local_var_3_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($xs_1);
$__local_var_4_1 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($__local_var_3_0, $f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_zipWith"];
  $__local_var_5_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_3_0);
  $__res = function($v1_6) use ($__local_var_5_1, $f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_zipWith"];
  if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_6) && (($v1_6)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("Nil");
} else {
if (((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Cons")) && (is_object($v1_6) && (($v1_6)->tag === "Cons")))) {
$__t2 = new Phpurs_Data2("Cons", (($f_0)(($__local_var_5_1)->value0))(($v1_6)->value0), (((($GLOBALS['Data_List_Lazy_zipWith'] ?? \PhpursThunks::eval('Data_List_Lazy_zipWith')))($f_0))(($__local_var_5_1)->value1))(($v1_6)->value1));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_5_4 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($ys_2);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_4_1, $__local_var_5_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_zipWith"];
  $__res = ((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_4_1))((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_5_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_zipWithA'] = function() { $v = function($dictApplicative_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_zipWithA"), recVars=[];
  $sequence1_1_0 = (((($GLOBALS['Data_List_Lazy_Types_traversableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_traversableList')))->traverse)($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($sequence1_1_0)((((($GLOBALS['Data_List_Lazy_zipWith'] ?? \PhpursThunks::eval('Data_List_Lazy_zipWith')))($f_2))($xs_3))($ys_4));
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
\PhpursThunks::$thunks['Data_List_Lazy_zip'] = function() { $v = (($GLOBALS['Data_List_Lazy_zipWith'] ?? \PhpursThunks::eval('Data_List_Lazy_zipWith')))(($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_updateAt'] = function() { $v = (function() {
  $__fn = function($n_0, $x_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_updateAt"), recVars=["Data_List_Lazy_updateAt"];
  while (true) {
$__local_var_3_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($xs_2);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($__local_var_3_0, $n_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_updateAt"];
  $__local_var_5_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_3_0);
  if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("Nil");
} else {
if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Cons"))) {
switch ($n_0) {
case 0:
$__t3 = new Phpurs_Data2("Cons", $x_1, ($__local_var_5_1)->value1);
break;
default:
$__t3 = new Phpurs_Data2("Cons", ($__local_var_5_1)->value0, (((($GLOBALS['Data_List_Lazy_updateAt'] ?? \PhpursThunks::eval('Data_List_Lazy_updateAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($n_0))(1)))($x_1))(($__local_var_5_1)->value1));
break;
};
$__t2 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_unzip'] = function() { $v = (((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))->foldr)(function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_1_0 = ($v_0)->value0;
  $__local_var_2_1 = ($v_0)->value1;
  $__res = function($v1_3) use ($__local_var_1_0, $__local_var_2_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_4_2 = ($v1_3)->value0;
  $__local_var_5_3 = ($v1_3)->value1;
  $__res = new Phpurs_Data2("Tuple", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_1_0, $__local_var_4_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $__local_var_1_0, $__local_var_4_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_2_1, $__local_var_5_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $__local_var_2_1, $__local_var_5_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(new Phpurs_Data2("Tuple", ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')), ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_uncons'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_uncons"), recVars=[];
  $v_1_0 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($xs_0);
  if ((is_object($v_1_0) && (($v_1_0)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("Nothing");
} else {
if ((is_object($v_1_0) && (($v_1_0)->tag === "Cons"))) {
$__t1 = new Phpurs_Data1("Just", (object)["head" => ($v_1_0)->value0, "tail" => ($v_1_0)->value1]);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_toUnfoldable'] = function() { $v = function($dictUnfoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_toUnfoldable"), recVars=[];
  $__res = (($dictUnfoldable_0)->unfoldr)(function($xs_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($xs_1);
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", (($__local_var_2_0)->value0)->head, (($__local_var_2_0)->value0)->tail));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_takeWhile'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_takeWhile"), recVars=["Data_List_Lazy_takeWhile"];
  while (true) {
$__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->map)(function($v_1) use ($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_takeWhile"];
  if (((is_object($v_1) && (($v_1)->tag === "Cons")) && ($p_0)(($v_1)->value0))) {
$__t0 = new Phpurs_Data2("Cons", ($v_1)->value0, ((($GLOBALS['Data_List_Lazy_takeWhile'] ?? \PhpursThunks::eval('Data_List_Lazy_takeWhile')))($p_0))(($v_1)->value1));
} else {
$__t0 = new Phpurs_Data0("Nil");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_take'] = function() { $v = function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_take"), recVars=["Data_List_Lazy_take"];
  while (true) {
if (( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(0)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_0))(0))->tag === "GT")))) {
$__t1 = function($v_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_take"];
  $__res = ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
} else {
$__t1 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->map)(function($v1_1) use ($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_take"];
  if ((is_object($v1_1) && (($v1_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_1) && (($v1_1)->tag === "Cons"))) {
$__t0 = new Phpurs_Data2("Cons", ($v1_1)->value0, ((($GLOBALS['Data_List_Lazy_take'] ?? \PhpursThunks::eval('Data_List_Lazy_take')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($n_0))(1)))(($v1_1)->value1));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_tail'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_tail"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($xs_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->value0)->tail);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_stripPrefix'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $v_1 = null, $s_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_stripPrefix"), recVars=[];
  $__res = (((($GLOBALS['Data_List_Lazy_tailRecM2'] ?? \PhpursThunks::eval('Data_List_Lazy_tailRecM2')))((function() use ($dictEq_0) {
  $__fn = function($prefix_3, $input_4 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $v1_5_0 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($prefix_3);
  if ((is_object($v1_5_0) && (($v1_5_0)->tag === "Nil"))) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data1("Done", $input_4));
} else {
if ((is_object($v1_5_0) && (($v1_5_0)->tag === "Cons"))) {
$v2_6_2 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($input_4);
if (((is_object($v2_6_2) && (($v2_6_2)->tag === "Cons")) && ((($dictEq_0)->eq)(($v1_5_0)->value0))(($v2_6_2)->value0))) {
$__t3 = new Phpurs_Data1("Just", new Phpurs_Data1("Loop", (object)["a" => ($v1_5_0)->value1, "b" => ($v2_6_2)->value1]));
} else {
$__t3 = new Phpurs_Data0("Nothing");
};
$__t1 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
  $__res = $__t1;
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
\PhpursThunks::$thunks['Data_List_Lazy_span'] = function() { $v = (function() {
  $__fn = function($p_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_span"), recVars=["Data_List_Lazy_span"];
  while (true) {
$v_2_0 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($xs_1);
if (((is_object($v_2_0) && (($v_2_0)->tag === "Just")) && ($p_0)((($v_2_0)->value0)->head))) {
$__local_var_3_2 = (($v_2_0)->value0)->head;
$v1_4_3 = ((($GLOBALS['Data_List_Lazy_span'] ?? \PhpursThunks::eval('Data_List_Lazy_span')))($p_0))((($v_2_0)->value0)->tail);
$__local_var_5_4 = ($v1_4_3)->init;
$__t1 = (object)["init" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_3_2, $__local_var_5_4) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_span"];
  $__res = new Phpurs_Data2("Cons", $__local_var_3_2, $__local_var_5_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "rest" => ($v1_4_3)->rest];
} else {
$__t1 = (object)["init" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')), "rest" => $xs_1];
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_snoc'] = function() { $v = (function() {
  $__fn = function($xs_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_snoc"), recVars=[];
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))->foldr)(($GLOBALS['Data_List_Lazy_Types_cons'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_cons'))))((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_2) use ($x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $x_1, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($xs_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_singleton'] = function() { $v = function($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_singleton"), recVars=[];
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_1) use ($a_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $a_0, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_showPattern'] = function() { $v = function($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_showPattern"), recVars=[];
  $__res = (object)["show" => function($v_1) use ($dictShow_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))("(Pattern "))(((($GLOBALS['Data_Semigroup_concatString'] ?? \PhpursThunks::eval('Data_Semigroup_concatString')))((((($GLOBALS['Data_List_Lazy_Types_showList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_showList')))($dictShow_0))->show)($v_1)))(")"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_scanlLazy'] = function() { $v = (function() {
  $__fn = function($f_0, $acc_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_scanlLazy"), recVars=["Data_List_Lazy_scanlLazy"];
  while (true) {
$__local_var_3_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($xs_2);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($__local_var_3_0, $acc_1, $f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_scanlLazy"];
  $__local_var_5_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_3_0);
  if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("Nil");
} else {
if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Cons"))) {
$acc__prime___6_3 = (($f_0)($acc_1))(($__local_var_5_1)->value0);
$__t2 = new Phpurs_Data2("Cons", $acc__prime___6_3, (((($GLOBALS['Data_List_Lazy_scanlLazy'] ?? \PhpursThunks::eval('Data_List_Lazy_scanlLazy')))($f_0))($acc__prime___6_3))(($__local_var_5_1)->value1));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_reverse'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_reverse"), recVars=[];
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))(function($v_1) use ($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))->foldl)((function() {
  $__fn = function($b_2, $a_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($a_3, $b_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $a_3, $b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))))($xs_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_replicateM'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_replicateM"), recVars=["Data_List_Lazy_replicateM"];
  while (true) {
$__local_var_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__local_var_2_1 = (($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($n_3, $m_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_replicateM"];
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_3))(1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($n_3))(1))->tag === "LT"))) {
$__t2 = (($__local_var_1_0)->pure)(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
} else {
$__t2 = ((($__local_var_2_1)->bind)($m_4))(function($a_5) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, $m_4, $n_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_replicateM"];
  $__res = ((($__local_var_2_1)->bind)((((($GLOBALS['Data_List_Lazy_replicateM'] ?? \PhpursThunks::eval('Data_List_Lazy_replicateM')))($dictMonad_0))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($n_3))(1)))($m_4)))(function($as_6) use ($__local_var_1_0, $a_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_replicateM"];
  $__res = (($__local_var_1_0)->pure)((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_7) use ($a_5, $as_6) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_replicateM"];
  $__res = new Phpurs_Data2("Cons", $a_5, $as_6);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
\PhpursThunks::$thunks['Data_List_Lazy_repeat'] = function() { $v = function($x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_repeat"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))(function($v_2) use (&$go_1_0, $x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_1_0"];
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use (&$go_1_0, $x_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_1_0"];
  $__res = new Phpurs_Data2("Cons", $x_0, $go_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_replicate'] = function() { $v = (function() {
  $__fn = function($i_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_replicate"), recVars=[];
  $__res = ((($GLOBALS['Data_List_Lazy_take'] ?? \PhpursThunks::eval('Data_List_Lazy_take')))($i_0))((($GLOBALS['Data_List_Lazy_repeat'] ?? \PhpursThunks::eval('Data_List_Lazy_repeat')))($xs_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_range'] = function() { $v = (function() {
  $__fn = function($start_0, $end_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_range"), recVars=[];
  if ((is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($start_0))($end_1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($start_0))($end_1))->tag === "GT"))) {
$__t1 = (((($GLOBALS['Data_List_Lazy_Types_unfoldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_unfoldableList')))->unfoldr)(function($x_2) use ($end_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($x_2))($end_1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($x_2))($end_1))->tag === "LT")))) {
$__t2 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", $x_2, ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($x_2))(1)));
} else {
$__t2 = new Phpurs_Data0("Nothing");
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($start_0);
} else {
$__t1 = (((($GLOBALS['Data_List_Lazy_Types_unfoldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_unfoldableList')))->unfoldr)(function($x_2) use ($end_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (( ! (is_object((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($x_2))($end_1)) && (((((($GLOBALS['Data_Ord_ordInt'] ?? \PhpursThunks::eval('Data_Ord_ordInt')))->compare)($x_2))($end_1))->tag === "GT")))) {
$__t0 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", $x_2, ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($x_2))(1)));
} else {
$__t0 = new Phpurs_Data0("Nothing");
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($start_0);
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_partition'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_partition"), recVars=[];
  $__res = (((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))->foldr)((function() use ($f_0) {
  $__fn = function($x_1, $v_2 = null) use ($f_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  if (($f_0)($x_1)) {
$__t0 = (object)["yes" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use ($v_2, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $x_1, ($v_2)->yes);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), "no" => ($v_2)->no];
} else {
$__t0 = (object)["yes" => ($v_2)->yes, "no" => (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use ($v_2, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = new Phpurs_Data2("Cons", $x_1, ($v_2)->no);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})];
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["yes" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')), "no" => ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_null'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_isNothing'] ?? \PhpursThunks::eval('Data_Maybe_isNothing'))))(($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons'))); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_nubBy'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_nubBy"), recVars=[];
  $goStep_1_0 = null;
  $go_1_0 = null;
  $goStep_1_0 = (function() use (&$go_1_0, $p_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, $p_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "goStep_1_0"), recVars=["goStep_1_0","go_1_0"];
  while (true) {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
$v2_4_2 = (((($GLOBALS['Data_List_Internal_insertAndLookupBy'] ?? \PhpursThunks::eval('Data_List_Internal_insertAndLookupBy')))($p_0))(($v1_3)->value0))($v_2);
if (($v2_4_2)->found) {
$__t3 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))((($go_1_0)(($v2_4_2)->result))(($v1_3)->value1));
} else {
$__t3 = new Phpurs_Data2("Cons", ($v1_3)->value0, (($go_1_0)(($v2_4_2)->result))(($v1_3)->value1));
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $go_1_0 = (function() use (&$goStep_1_0) {
  $__fn = function($s_2, $v_3 = null) use (&$goStep_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["goStep_1_0","go_1_0"];
  while (true) {
$__local_var_4_4 = ($goStep_1_0)($s_2);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5) use ($__local_var_4_4, $v_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["goStep_1_0","go_1_0"];
  $__res = ($__local_var_4_4)((($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($v_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ($go_1_0)(new Phpurs_Data0("Leaf"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_nub'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_nub"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_nubBy'] ?? \PhpursThunks::eval('Data_List_Lazy_nubBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_newtypePattern'] = function() { $v = (object)["Coercible0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_mapMaybe'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_mapMaybe"), recVars=["Data_List_Lazy_mapMaybe"];
  while (true) {
$go_1_0 = null;
$go_1_0 = function($v_2) use ($f_0, &$go_1_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["Data_List_Lazy_mapMaybe","go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_2) && (($v_2)->tag === "Cons"))) {
$v1_3_2 = ($f_0)(($v_2)->value0);
if ((is_object($v1_3_2) && (($v1_3_2)->tag === "Nothing"))) {
$__tco_4 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_2)->value1);
$v_2 = $__tco_4;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_3_2) && (($v1_3_2)->tag === "Just"))) {
$__t3 = new Phpurs_Data2("Cons", ($v1_3_2)->value0, ((($GLOBALS['Data_List_Lazy_mapMaybe'] ?? \PhpursThunks::eval('Data_List_Lazy_mapMaybe')))($f_0))(($v_2)->value1));
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
$__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->map)($go_1_0)))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_some'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_some"), recVars=["Data_List_Lazy_some","Data_List_Lazy_many"];
  while (true) {
$__res = ((((((($dictAlternative_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Apply0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->apply)(((((((((($dictAlternative_0)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Functor0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->map)(($GLOBALS['Data_List_Lazy_Types_cons'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_cons'))))($v_2)))((($dictLazy_1)->defer)(function($v1_3) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_some","Data_List_Lazy_many"];
  $__res = (((($GLOBALS['Data_List_Lazy_many'] ?? \PhpursThunks::eval('Data_List_Lazy_many')))($dictAlternative_0))($dictLazy_1))($v_2);
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
\PhpursThunks::$thunks['Data_List_Lazy_many'] = function() { $v = (function() {
  $__fn = function($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_many"), recVars=["Data_List_Lazy_some","Data_List_Lazy_many"];
  while (true) {
$__res = ((((((($dictAlternative_0)->Plus1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->Alt0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->alt)((((($GLOBALS['Data_List_Lazy_some'] ?? \PhpursThunks::eval('Data_List_Lazy_some')))($dictAlternative_0))($dictLazy_1))($v_2)))((((($dictAlternative_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil'))));
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_length'] = function() { $v = (((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))->foldl)((function() {
  $__fn = function($l_0, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($l_0))(1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(0); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_last'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = function($v_1) use (&$go_0_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($v_1) && (($v_1)->tag === "Cons"))) {
if ((($GLOBALS['Data_List_Lazy_null'] ?? \PhpursThunks::eval('Data_List_Lazy_null')))(($v_1)->value1)) {
$__t3 = new Phpurs_Data1("Just", ($v_1)->value0);
} else {
$__tco_2 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_1)->value1);
$v_1 = $__tco_2;
continue ;
$__t3 = null;
};
$__t1 = $__t3;
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($go_0_0))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_iterate'] = function() { $v = (function() {
  $__fn = function($f_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_iterate"), recVars=[];
  $go_2_0 = null;
  $go_2_0 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))(function($v_3) use ($f_0, &$go_2_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_2_0"];
  $__local_var_4_1 = (((($GLOBALS['Data_List_Lazy_Types_functorList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_functorList')))->map)($f_0))($go_2_0);
  $__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5) use ($__local_var_4_1, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_2_0"];
  $__res = new Phpurs_Data2("Cons", $x_1, $__local_var_4_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = $go_2_0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_insertAt'] = function() { $v = (function() {
  $__fn = function($v_0, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_insertAt"), recVars=["Data_List_Lazy_insertAt"];
  while (true) {
switch ($v_0) {
case 0:
$__t3 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use ($v1_1, $v2_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_insertAt"];
  $__res = new Phpurs_Data2("Cons", $v1_1, $v2_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
break;
default:
$__local_var_3_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($v2_2);
$__t3 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($__local_var_3_0, $v1_1, $v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_insertAt"];
  $__local_var_5_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_3_0);
  if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data2("Cons", $v1_1, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
} else {
if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Cons"))) {
$__t2 = new Phpurs_Data2("Cons", ($__local_var_5_1)->value0, (((($GLOBALS['Data_List_Lazy_insertAt'] ?? \PhpursThunks::eval('Data_List_Lazy_insertAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_0))(1)))($v1_1))(($__local_var_5_1)->value1));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
break;
};
$__res = $__t3;
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_init'] = function() { $v = (function() use (&$__fn) {
$go_0_0 = null;
$go_0_0 = function($v_1) use (&$go_0_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_0_0"), recVars=["go_0_0"];
  while (true) {
if ((is_object($v_1) && (($v_1)->tag === "Cons"))) {
if ((($GLOBALS['Data_List_Lazy_null'] ?? \PhpursThunks::eval('Data_List_Lazy_null')))(($v_1)->value1)) {
$__t6 = new Phpurs_Data1("Just", ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
} else {
$__local_var_2_2 = ($go_0_0)((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_1)->value1));
if ((is_object($__local_var_2_2) && (($__local_var_2_2)->tag === "Just"))) {
$__local_var_3_4 = ($v_1)->value0;
$__local_var_4_5 = ($__local_var_2_2)->value0;
$__t3 = new Phpurs_Data1("Just", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_5) use ($__local_var_3_4, $__local_var_4_5) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_0_0"];
  $__res = new Phpurs_Data2("Cons", $__local_var_3_4, $__local_var_4_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
} else {
$__t3 = new Phpurs_Data0("Nothing");
};
$__t6 = $__t3;
};
$__t1 = $__t6;
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
return ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($go_0_0))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')));
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_index'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_index"), recVars=[];
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
switch ($v1_3) {
case 0:
$__t4 = new Phpurs_Data1("Just", ($v_2)->value0);
break;
default:
$__tco_2 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_2)->value1);
$__tco_3 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v1_3))(1);
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
continue 2;
$__t4 = null;
break;
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
  $__res = ($go_1_0)((($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($xs_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_head'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_head"), recVars=[];
  $__local_var_1_0 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($xs_0);
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", (($__local_var_1_0)->value0)->head);
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_transpose'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_transpose"), recVars=["Data_List_Lazy_transpose"];
  while (true) {
$v_1_0 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($xs_0);
if ((is_object($v_1_0) && (($v_1_0)->tag === "Nothing"))) {
$__t1 = $xs_0;
} else {
if ((is_object($v_1_0) && (($v_1_0)->tag === "Just"))) {
$v1_2_2 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))((($v_1_0)->value0)->head);
if ((is_object($v1_2_2) && (($v1_2_2)->tag === "Nothing"))) {
$__tco_4 = (($v_1_0)->value0)->tail;
$xs_0 = $__tco_4;
continue ;
$__t3 = null;
} else {
if ((is_object($v1_2_2) && (($v1_2_2)->tag === "Just"))) {
$__local_var_3_5 = (($v1_2_2)->value0)->head;
$__local_var_4_6 = (($v1_2_2)->value0)->tail;
$__local_var_5_7 = ((($GLOBALS['Data_List_Lazy_mapMaybe'] ?? \PhpursThunks::eval('Data_List_Lazy_mapMaybe')))(($GLOBALS['Data_List_Lazy_head'] ?? \PhpursThunks::eval('Data_List_Lazy_head'))))((($v_1_0)->value0)->tail);
$__local_var_6_8 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_6) use ($__local_var_3_5, $__local_var_5_7) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_transpose"];
  $__res = new Phpurs_Data2("Cons", $__local_var_3_5, $__local_var_5_7);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__local_var_7_9 = ((($GLOBALS['Data_List_Lazy_mapMaybe'] ?? \PhpursThunks::eval('Data_List_Lazy_mapMaybe')))(($GLOBALS['Data_List_Lazy_tail'] ?? \PhpursThunks::eval('Data_List_Lazy_tail'))))((($v_1_0)->value0)->tail);
$__local_var_7_9 = (($GLOBALS['Data_List_Lazy_transpose'] ?? \PhpursThunks::eval('Data_List_Lazy_transpose')))((($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_8) use ($__local_var_4_6, $__local_var_7_9) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_transpose"];
  $__res = new Phpurs_Data2("Cons", $__local_var_4_6, $__local_var_7_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__t3 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_8) use ($__local_var_6_8, $__local_var_7_9) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_transpose"];
  $__res = new Phpurs_Data2("Cons", $__local_var_6_8, $__local_var_7_9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_groupBy'] = function() { $v = function($eq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_groupBy"), recVars=["Data_List_Lazy_groupBy"];
  while (true) {
$__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->map)(function($v_1) use ($eq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_groupBy"];
  if ((is_object($v_1) && (($v_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_1) && (($v_1)->tag === "Cons"))) {
$__local_var_2_1 = ($v_1)->value0;
$v1_3_2 = ((($GLOBALS['Data_List_Lazy_span'] ?? \PhpursThunks::eval('Data_List_Lazy_span')))(($eq_0)($__local_var_2_1)))(($v_1)->value1);
$__local_var_4_3 = ($v1_3_2)->init;
$__t0 = new Phpurs_Data2("Cons", (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v2_5) use ($__local_var_2_1, $__local_var_4_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_groupBy"];
  $__res = new Phpurs_Data2("NonEmpty", $__local_var_2_1, $__local_var_4_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}), ((($GLOBALS['Data_List_Lazy_groupBy'] ?? \PhpursThunks::eval('Data_List_Lazy_groupBy')))($eq_0))(($v1_3_2)->rest));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_group'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_group"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_groupBy'] ?? \PhpursThunks::eval('Data_List_Lazy_groupBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_fromStep'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))((($GLOBALS['Data_Lazy_applicativeLazy'] ?? \PhpursThunks::eval('Data_Lazy_applicativeLazy')))->pure); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_insertBy'] = function() { $v = (function() {
  $__fn = function($cmp_0, $x_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_insertBy"), recVars=["Data_List_Lazy_insertBy"];
  while (true) {
$__local_var_3_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($xs_2);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($__local_var_3_0, $cmp_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_insertBy"];
  $__local_var_5_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_3_0);
  if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data2("Cons", $x_1, ($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
} else {
if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Cons"))) {
if ((is_object((($cmp_0)($x_1))(($__local_var_5_1)->value0)) && (((($cmp_0)($x_1))(($__local_var_5_1)->value0))->tag === "GT"))) {
$__t3 = new Phpurs_Data2("Cons", ($__local_var_5_1)->value0, (((($GLOBALS['Data_List_Lazy_insertBy'] ?? \PhpursThunks::eval('Data_List_Lazy_insertBy')))($cmp_0))($x_1))(($__local_var_5_1)->value1));
} else {
$__t3 = new Phpurs_Data2("Cons", $x_1, (($GLOBALS['Data_List_Lazy_fromStep'] ?? \PhpursThunks::eval('Data_List_Lazy_fromStep')))($__local_var_5_1));
};
$__t2 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_insert'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_insert"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_insertBy'] ?? \PhpursThunks::eval('Data_List_Lazy_insertBy')))(($dictOrd_0)->compare);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_fromFoldable'] = function() { $v = function($dictFoldable_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_fromFoldable"), recVars=[];
  $__res = ((($dictFoldable_0)->foldr)(($GLOBALS['Data_List_Lazy_Types_cons'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_cons'))))(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_foldrLazy'] = function() { $v = (function() {
  $__fn = function($dictLazy_0, $op_1 = null, $z_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_foldrLazy"), recVars=[];
  $go_3_0 = null;
  $go_3_0 = function($xs_4) use ($dictLazy_0, &$go_3_0, $op_1, $z_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_3_0"), recVars=["go_3_0"];
  while (true) {
$v_5_1 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))($xs_4);
if ((is_object($v_5_1) && (($v_5_1)->tag === "Cons"))) {
$__local_var_6_3 = ($v_5_1)->value0;
$__local_var_7_4 = ($v_5_1)->value1;
$__t2 = (($dictLazy_0)->defer)(function($v1_8) use ($__local_var_6_3, $__local_var_7_4, &$go_3_0, $op_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_3_0"];
  $__res = (($op_1)($__local_var_6_3))(($go_3_0)($__local_var_7_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
} else {
if ((is_object($v_5_1) && (($v_5_1)->tag === "Nil"))) {
$__t2 = $z_2;
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
  $__res = $go_3_0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_foldM'] = function() { $v = (function() {
  $__fn = function($dictMonad_0, $f_1 = null, $b_2 = null, $xs_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_foldM"), recVars=["Data_List_Lazy_foldM"];
  while (true) {
$v_4_0 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($xs_3);
if ((is_object($v_4_0) && (($v_4_0)->tag === "Nothing"))) {
$__t1 = (((($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->pure)($b_2);
} else {
if ((is_object($v_4_0) && (($v_4_0)->tag === "Just"))) {
$__local_var_5_2 = (($v_4_0)->value0)->tail;
$__t1 = ((((($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->bind)((($f_1)($b_2))((($v_4_0)->value0)->head)))(function($b__prime___6) use ($__local_var_5_2, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_foldM"];
  $__res = ((((($GLOBALS['Data_List_Lazy_foldM'] ?? \PhpursThunks::eval('Data_List_Lazy_foldM')))($dictMonad_0))($f_1))($b__prime___6))($__local_var_5_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
};
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_findIndex'] = function() { $v = function($fn_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_findIndex"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use ($fn_0, &$go_1_0) {
  $__fn = function($n_2, $list_3 = null) use ($fn_0, &$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
$__local_var_4_1 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($list_3);
if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Just"))) {
if (($fn_0)((($__local_var_4_1)->value0)->head)) {
$__t5 = new Phpurs_Data1("Just", $n_2);
} else {
$__tco_3 = ((($GLOBALS['Data_Semiring_intAdd'] ?? \PhpursThunks::eval('Data_Semiring_intAdd')))($n_2))(1);
$__tco_4 = (($__local_var_4_1)->value0)->tail;
$n_2 = $__tco_3;
$list_3 = $__tco_4;
continue ;
$__t5 = null;
};
$__t2 = $__t5;
} else {
if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Nothing"))) {
$__t2 = new Phpurs_Data0("Nothing");
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
  $__res = ($go_1_0)(0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_findLastIndex'] = function() { $v = (function() {
  $__fn = function($fn_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_findLastIndex"), recVars=[];
  $__local_var_2_0 = ((($GLOBALS['Data_List_Lazy_findIndex'] ?? \PhpursThunks::eval('Data_List_Lazy_findIndex')))($fn_0))((($GLOBALS['Data_List_Lazy_reverse'] ?? \PhpursThunks::eval('Data_List_Lazy_reverse')))($xs_1));
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->tag === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))((($GLOBALS['Data_List_Lazy_length'] ?? \PhpursThunks::eval('Data_List_Lazy_length')))($xs_1)))(1)))(($__local_var_2_0)->value0));
} else {
$__t1 = new Phpurs_Data0("Nothing");
};
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_filterM'] = function() { $v = function($dictMonad_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_filterM"), recVars=["Data_List_Lazy_filterM"];
  while (true) {
$__local_var_1_0 = (($dictMonad_0)->Applicative0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__local_var_2_1 = (($dictMonad_0)->Bind1)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
$__res = (function() use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0) {
  $__fn = function($p_3, $list_4 = null) use ($__local_var_1_0, $__local_var_2_1, $dictMonad_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_filterM"];
  $v_5_2 = (($GLOBALS['Data_List_Lazy_uncons'] ?? \PhpursThunks::eval('Data_List_Lazy_uncons')))($list_4);
  if ((is_object($v_5_2) && (($v_5_2)->tag === "Nothing"))) {
$__t3 = (($__local_var_1_0)->pure)(($GLOBALS['Data_List_Lazy_Types_nil'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_nil')));
} else {
if ((is_object($v_5_2) && (($v_5_2)->tag === "Just"))) {
$__local_var_6_4 = (($v_5_2)->value0)->head;
$__local_var_7_5 = (($v_5_2)->value0)->tail;
$__t3 = ((($__local_var_2_1)->bind)(($p_3)($__local_var_6_4)))(function($b_8) use ($__local_var_1_0, $__local_var_2_1, $__local_var_6_4, $__local_var_7_5, $dictMonad_0, $p_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_filterM"];
  $__res = ((($__local_var_2_1)->bind)((((($GLOBALS['Data_List_Lazy_filterM'] ?? \PhpursThunks::eval('Data_List_Lazy_filterM')))($dictMonad_0))($p_3))($__local_var_7_5)))(function($xs__prime___9) use ($__local_var_1_0, $__local_var_6_4, $b_8) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_filterM"];
  if ($b_8) {
$__t6 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_10) use ($__local_var_6_4, $xs__prime___9) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_filterM"];
  $__res = new Phpurs_Data2("Cons", $__local_var_6_4, $xs__prime___9);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
} else {
$__t6 = $xs__prime___9;
};
  $__res = (($__local_var_1_0)->pure)($__t6);
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
$__t3 = null;
};
};
  $__res = $__t3;
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
\PhpursThunks::$thunks['Data_List_Lazy_filter'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_filter"), recVars=["Data_List_Lazy_filter"];
  while (true) {
$go_1_0 = null;
$go_1_0 = function($v_2) use (&$go_1_0, $p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["Data_List_Lazy_filter","go_1_0"];
  while (true) {
if ((is_object($v_2) && (($v_2)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_2) && (($v_2)->tag === "Cons"))) {
if (($p_0)(($v_2)->value0)) {
$__t3 = new Phpurs_Data2("Cons", ($v_2)->value0, ((($GLOBALS['Data_List_Lazy_filter'] ?? \PhpursThunks::eval('Data_List_Lazy_filter')))($p_0))(($v_2)->value1));
} else {
$__tco_2 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_2)->value1);
$v_2 = $__tco_2;
continue ;
$__t3 = null;
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
$__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->map)($go_1_0)))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_intersectBy'] = function() { $v = (function() {
  $__fn = function($eq_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_intersectBy"), recVars=[];
  $__res = ((($GLOBALS['Data_List_Lazy_filter'] ?? \PhpursThunks::eval('Data_List_Lazy_filter')))(function($x_3) use ($eq_0, $ys_2) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Data_List_Lazy_any'] ?? \PhpursThunks::eval('Data_List_Lazy_any')))(($eq_0)($x_3)))($ys_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_intersect'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_intersect"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_intersectBy'] ?? \PhpursThunks::eval('Data_List_Lazy_intersectBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_nubByEq'] = function() { $v = function($eq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_nubByEq"), recVars=["Data_List_Lazy_nubByEq"];
  while (true) {
$__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->map)(function($v_1) use ($eq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_nubByEq"];
  if ((is_object($v_1) && (($v_1)->tag === "Nil"))) {
$__t0 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v_1) && (($v_1)->tag === "Cons"))) {
$__local_var_2_1 = ($v_1)->value0;
$__t0 = new Phpurs_Data2("Cons", $__local_var_2_1, ((($GLOBALS['Data_List_Lazy_nubByEq'] ?? \PhpursThunks::eval('Data_List_Lazy_nubByEq')))($eq_0))(((($GLOBALS['Data_List_Lazy_filter'] ?? \PhpursThunks::eval('Data_List_Lazy_filter')))(function($y_3) use ($__local_var_2_1, $eq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_nubByEq"];
  $__res = (($GLOBALS['Data_HeytingAlgebra_boolNot'] ?? \PhpursThunks::eval('Data_HeytingAlgebra_boolNot')))((($eq_0)($__local_var_2_1))($y_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($v_1)->value1)));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t0 = null;
};
};
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_nubEq'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_nubEq"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_nubByEq'] ?? \PhpursThunks::eval('Data_List_Lazy_nubByEq')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_eqPattern'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_eqPattern"), recVars=[];
  $__res = (object)["eq" => (function() use ($dictEq_0) {
  $__fn = function($x_1, $y_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eq1List')))->eq1)($dictEq_0))($x_1))($y_2);
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
\PhpursThunks::$thunks['Data_List_Lazy_ordPattern'] = function() { $v = function($dictOrd_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_ordPattern"), recVars=[];
  $__local_var_1_0 = (($dictOrd_0)->Eq0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $eqPattern1_2_1 = (object)["eq" => (function() use ($__local_var_1_0) {
  $__fn = function($x_2, $y_3 = null) use ($__local_var_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_eq1List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_eq1List')))->eq1)($__local_var_1_0))($x_2))($y_3);
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
  $__res = ((((($GLOBALS['Data_List_Lazy_Types_ordList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_ordList')))($dictOrd_0))->compare)($x_3))($y_4);
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
\PhpursThunks::$thunks['Data_List_Lazy_elemLastIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_elemLastIndex"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_findLastIndex'] ?? \PhpursThunks::eval('Data_List_Lazy_findLastIndex')))(function($v_2) use ($dictEq_0, $x_1) {
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
\PhpursThunks::$thunks['Data_List_Lazy_elemIndex'] = function() { $v = (function() {
  $__fn = function($dictEq_0, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_elemIndex"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_findIndex'] ?? \PhpursThunks::eval('Data_List_Lazy_findIndex')))(function($v_2) use ($dictEq_0, $x_1) {
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
\PhpursThunks::$thunks['Data_List_Lazy_dropWhile'] = function() { $v = function($p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_dropWhile"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = function($v_2) use (&$go_1_0, $p_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if (((is_object($v_2) && (($v_2)->tag === "Cons")) && ($p_0)(($v_2)->value0))) {
$__tco_2 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v_2)->value1);
$v_2 = $__tco_2;
continue ;
$__t1 = null;
} else {
$__t1 = (($GLOBALS['Data_List_Lazy_fromStep'] ?? \PhpursThunks::eval('Data_List_Lazy_fromStep')))($v_2);
};
$__res = $__t1;
goto __end;;
};
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($go_1_0))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_drop'] = function() { $v = function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_drop"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0) {
  $__fn = function($v_2, $v1_3 = null) use (&$go_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "go_1_0"), recVars=["go_1_0"];
  while (true) {
if (($v_2 === 0)) {
$__t1 = $v1_3;
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Nil"))) {
$__t1 = new Phpurs_Data0("Nil");
} else {
if ((is_object($v1_3) && (($v1_3)->tag === "Cons"))) {
$__tco_2 = ((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($v_2))(1);
$__tco_3 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($v1_3)->value1);
$v_2 = $__tco_2;
$v1_3 = $__tco_3;
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_List'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_List'))))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(((($GLOBALS['Data_Lazy_functorLazy'] ?? \PhpursThunks::eval('Data_Lazy_functorLazy')))->map)(($go_1_0)($n_0))))(($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce'))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_slice'] = function() { $v = (function() {
  $__fn = function($start_0, $end_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_slice"), recVars=[];
  $__res = ((($GLOBALS['Data_List_Lazy_take'] ?? \PhpursThunks::eval('Data_List_Lazy_take')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($end_1))($start_0)))(((($GLOBALS['Data_List_Lazy_drop'] ?? \PhpursThunks::eval('Data_List_Lazy_drop')))($start_0))($xs_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_deleteBy'] = function() { $v = (function() {
  $__fn = function($eq_0, $x_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_deleteBy"), recVars=["Data_List_Lazy_deleteBy"];
  while (true) {
$__local_var_3_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($xs_2);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($__local_var_3_0, $eq_0, $x_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_deleteBy"];
  $__local_var_5_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_3_0);
  if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("Nil");
} else {
if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Cons"))) {
if ((($eq_0)($x_1))(($__local_var_5_1)->value0)) {
$__t3 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($__local_var_5_1)->value1);
} else {
$__t3 = new Phpurs_Data2("Cons", ($__local_var_5_1)->value0, (((($GLOBALS['Data_List_Lazy_deleteBy'] ?? \PhpursThunks::eval('Data_List_Lazy_deleteBy')))($eq_0))($x_1))(($__local_var_5_1)->value1));
};
$__t2 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_unionBy'] = function() { $v = (function() {
  $__fn = function($eq_0, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_unionBy"), recVars=[];
  $__res = (((($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))->append)($xs_1))(((((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))->foldl)((function() use ($eq_0) {
  $__fn = function($b_3, $a_4 = null) use ($eq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_List_Lazy_deleteBy'] ?? \PhpursThunks::eval('Data_List_Lazy_deleteBy')))($eq_0))($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(((($GLOBALS['Data_List_Lazy_nubByEq'] ?? \PhpursThunks::eval('Data_List_Lazy_nubByEq')))($eq_0))($ys_2)))($xs_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_union'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_union"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_unionBy'] ?? \PhpursThunks::eval('Data_List_Lazy_unionBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_deleteAt'] = function() { $v = (function() {
  $__fn = function($n_0, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_deleteAt"), recVars=["Data_List_Lazy_deleteAt"];
  while (true) {
$__local_var_2_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($xs_1);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_3) use ($__local_var_2_0, $n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_deleteAt"];
  $__local_var_4_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_2_0);
  if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("Nil");
} else {
if ((is_object($__local_var_4_1) && (($__local_var_4_1)->tag === "Cons"))) {
switch ($n_0) {
case 0:
$__t3 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($__local_var_4_1)->value1);
break;
default:
$__t3 = new Phpurs_Data2("Cons", ($__local_var_4_1)->value0, ((($GLOBALS['Data_List_Lazy_deleteAt'] ?? \PhpursThunks::eval('Data_List_Lazy_deleteAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($n_0))(1)))(($__local_var_4_1)->value1));
break;
};
$__t2 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_delete'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_delete"), recVars=[];
  $__res = (($GLOBALS['Data_List_Lazy_deleteBy'] ?? \PhpursThunks::eval('Data_List_Lazy_deleteBy')))(($dictEq_0)->eq);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_difference'] = function() { $v = function($dictEq_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_difference"), recVars=[];
  $__res = ((($GLOBALS['Data_List_Lazy_Types_foldableList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_foldableList')))->foldl)((function() use ($dictEq_0) {
  $__fn = function($b_1, $a_2 = null) use ($dictEq_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (((($GLOBALS['Data_List_Lazy_deleteBy'] ?? \PhpursThunks::eval('Data_List_Lazy_deleteBy')))(($dictEq_0)->eq))($a_2))($b_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_cycle'] = function() { $v = function($xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_cycle"), recVars=[];
  $go_1_0 = null;
  $go_1_0 = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step'))))(function($v_2) use (&$go_1_0, $xs_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["go_1_0"];
  $__res = (((($GLOBALS['Data_List_Lazy_Types_semigroupList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_semigroupList')))->append)($xs_0))($go_1_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__res = $go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_concatMap'] = function() { $v = (function() {
  $__fn = function($b_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_concatMap"), recVars=[];
  $__res = (((($GLOBALS['Data_List_Lazy_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_bindList')))->bind)($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_concat'] = function() { $v = function($v_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_concat"), recVars=[];
  $__res = (((($GLOBALS['Data_List_Lazy_Types_bindList'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_bindList')))->bind)($v_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_catMaybes'] = function() { $v = (($GLOBALS['Data_List_Lazy_mapMaybe'] ?? \PhpursThunks::eval('Data_List_Lazy_mapMaybe')))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))->identity); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_alterAt'] = function() { $v = (function() {
  $__fn = function($n_0, $f_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_alterAt"), recVars=["Data_List_Lazy_alterAt"];
  while (true) {
$__local_var_3_0 = (($GLOBALS['Unsafe_Coerce_unsafeCoerce'] ?? \PhpursThunks::eval('Unsafe_Coerce_unsafeCoerce')))($xs_2);
$__res = (($GLOBALS['Data_Lazy_defer'] ?? \PhpursThunks::eval('Data_Lazy_defer')))(function($v_4) use ($__local_var_3_0, $f_1, $n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Data_List_Lazy_alterAt"];
  $__local_var_5_1 = (($GLOBALS['Data_Lazy_force'] ?? \PhpursThunks::eval('Data_Lazy_force')))($__local_var_3_0);
  if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Nil"))) {
$__t2 = new Phpurs_Data0("Nil");
} else {
if ((is_object($__local_var_5_1) && (($__local_var_5_1)->tag === "Cons"))) {
switch ($n_0) {
case 0:
$v2_6_4 = ($f_1)(($__local_var_5_1)->value0);
if ((is_object($v2_6_4) && (($v2_6_4)->tag === "Nothing"))) {
$__t5 = (($GLOBALS['Data_List_Lazy_Types_step'] ?? \PhpursThunks::eval('Data_List_Lazy_Types_step')))(($__local_var_5_1)->value1);
} else {
if ((is_object($v2_6_4) && (($v2_6_4)->tag === "Just"))) {
$__t5 = new Phpurs_Data2("Cons", ($v2_6_4)->value0, ($__local_var_5_1)->value1);
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
$__t3 = $__t5;
break;
default:
$__t3 = new Phpurs_Data2("Cons", ($__local_var_5_1)->value0, (((($GLOBALS['Data_List_Lazy_alterAt'] ?? \PhpursThunks::eval('Data_List_Lazy_alterAt')))(((($GLOBALS['Data_Ring_intSub'] ?? \PhpursThunks::eval('Data_Ring_intSub')))($n_0))(1)))($f_1))(($__local_var_5_1)->value1));
break;
};
$__t2 = $__t3;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
};
};
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto __end;;
};
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_List_Lazy_modifyAt'] = function() { $v = (function() {
  $__fn = function($n_0, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Data_List_Lazy_modifyAt"), recVars=[];
  $__res = ((($GLOBALS['Data_List_Lazy_alterAt'] ?? \PhpursThunks::eval('Data_List_Lazy_alterAt')))($n_0))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))($f_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };













































































