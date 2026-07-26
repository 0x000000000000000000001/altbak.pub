<?php

namespace Control\Monad\Rec\Class;

// ALL IMPORTS: Control.Applicative, Control.Bind, Control.Monad, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Either, Data.Function, Data.Functor, Data.Identity, Data.Maybe, Data.Monoid, Data.Semigroup, Data.Unit, Effect, Effect.Ref, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Control.Monad, Control.Monad.Rec.Class, Control.Semigroupoid, Data.Bifunctor, Data.Either, Data.Function, Data.Functor, Data.Identity, Data.Maybe, Data.Monoid, Data.Semigroup, Data.Unit, Effect, Effect.Ref, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Bifunctor/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Ref/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';

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
\PhpursThunks::$thunks['Control_Monad_Rec_Class_Loop'] = function() { $v = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Loop", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_Done'] = function() { $v = function($value0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Done", $value0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_tailRecM'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'tailRecM'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_tailRecM2'] = function() { $v = (function() {
  $__fn = function($dictMonadRec_0 = null, $f_1 = null, $a_2 = null, $b_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($o_4 = null) use ($f_1) {
  $__num = \func_num_args();
  $__res = (($f_1)(($o_4)->{'a'}))(($o_4)->{'b'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_2, "b" => $b_3]);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_tailRecM3'] = function() { $v = (function() {
  $__fn = function($dictMonadRec_0 = null, $f_1 = null, $a_2 = null, $b_3 = null, $c_4 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($o_5 = null) use ($f_1) {
  $__num = \func_num_args();
  $__res = ((($f_1)(($o_5)->{'a'}))(($o_5)->{'b'}))(($o_5)->{'c'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_2, "b" => $b_3, "c" => $c_4]);
  goto __end;;
  __end:
  return $__num > 5 ? $__res(...\array_slice(\func_get_args(), 5)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_untilJust'] = function() { $v = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((((((($dictMonadRec_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($m_2 = null) use ($__local_var_1_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($v_3 = null) use ($__local_var_1_0, $m_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})(function($v1_4 = null) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data1("Loop", ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
goto end_branch_1;;
};
  if ((is_object($v1_4) && (($v1_4)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Done", ($v1_4)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($m_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_whileJust'] = function() { $v = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $mempty_1_0 = ($dictMonoid_0)->{'mempty'};
  $__res = function($dictMonadRec_2 = null) use ($dictMonoid_0, $mempty_1_0) {
  $__num = \func_num_args();
  $__local_var_3_1 = (((((((($dictMonadRec_2)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($m_4 = null) use ($__local_var_3_1, $dictMonadRec_2, $dictMonoid_0, $mempty_1_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_2)->{'tailRecM'})(function($v_5 = null) use ($__local_var_3_1, $dictMonoid_0, $m_4) {
  $__num = \func_num_args();
  $__res = ((($__local_var_3_1)->{'map'})(function($v1_6 = null) use ($dictMonoid_0, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Nothing"))) {
$__t2 = new Phpurs_Data1("Done", $v_5);
goto end_branch_2;;
};
  if ((is_object($v1_6) && (($v1_6)->{'tag'} === "Just"))) {
$__t2 = new Phpurs_Data1("Loop", ((((($dictMonoid_0)->{'Semigroup0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'append'})($v_5))(($v1_6)->{'value0'}));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($m_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_1_0);
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
\PhpursThunks::$thunks['Control_Monad_Rec_Class_tailRec'] = function() { $v = function($f_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = function($v_2 = null) use ($f_0, &$go_1_0) {
  $__num = \func_num_args();
  $__tco_var_go_1_0_0_v_2 = $v_2;
  $__tco_done_go_1_0_0 = false;
  $__tco_res_go_1_0_0 = null;
  $__tco_loop_go_1_0_0 = function($v_2 = null) use (&$__tco_done_go_1_0_0, &$__tco_var_go_1_0_0_v_2, $f_0, &$go_1_0) {
  $__num = \func_num_args();
  $__tco_done_go_1_0_0 = true;
  $v_2 = $__tco_var_go_1_0_0_v_2;
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Loop"))) {
$__tco_1 = ($f_0)(($v_2)->{'value0'});
$__tco_var_go_1_0_0_v_2 = $__tco_1;
$__tco_done_go_1_0_0 = false;
$__res = null;
goto __end;;
$__t0 = null;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Done"))) {
$__t0 = ($v_2)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  while (($__tco_done_go_1_0_0 === false)) {
$__tco_res_go_1_0_0 = ($__tco_loop_go_1_0_0)($__tco_var_go_1_0_0_v_2);
};
  $__res = $__tco_res_go_1_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))($go_1_0))($f_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_tailRec2'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $a_1 = null, $b_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_Rec_Class_tailRec'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_tailRec')))(function($o_3 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($f_0)(($o_3)->{'a'}))(($o_3)->{'b'});
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
\PhpursThunks::$thunks['Control_Monad_Rec_Class_tailRec3'] = function() { $v = (function() {
  $__fn = function($f_0 = null, $a_1 = null, $b_2 = null, $c_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = ((($GLOBALS['Control_Monad_Rec_Class_tailRec'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_tailRec')))(function($o_4 = null) use ($f_0) {
  $__num = \func_num_args();
  $__res = ((($f_0)(($o_4)->{'a'}))(($o_4)->{'b'}))(($o_4)->{'c'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $a_1, "b" => $b_2, "c" => $c_3]);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_monadRecMaybe'] = function() { $v = (object)["tailRecM" => (function() {
  $__fn = function($f_0 = null, $a0_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Monad_Rec_Class_tailRec'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_tailRec')))(function($v_2 = null) use ($f_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Nothing"))) {
$__t0 = new Phpurs_Data1("Done", new Phpurs_Data0("Nothing"));
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Just"))) {
$__t1 = null;;
if ((is_object(($v_2)->{'value0'}) && ((($v_2)->{'value0'})->{'tag'} === "Loop"))) {
$__t1 = new Phpurs_Data1("Loop", ($f_0)((($v_2)->{'value0'})->{'value0'}));
goto end_branch_1;;
};
if ((is_object(($v_2)->{'value0'}) && ((($v_2)->{'value0'})->{'tag'} === "Done"))) {
$__t1 = new Phpurs_Data1("Done", new Phpurs_Data1("Just", (($v_2)->{'value0'})->{'value0'}));
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_0)($a0_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_monadMaybe'] ?? \PhpursThunks::eval('Data_Maybe_monadMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_monadRecIdentity'] = function() { $v = (object)["tailRecM" => function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Identity_Identity'] ?? \PhpursThunks::eval('Data_Identity_Identity'))))((($GLOBALS['Control_Monad_Rec_Class_tailRec'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_tailRec')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = $v_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Identity_monadIdentity'] ?? \PhpursThunks::eval('Data_Identity_monadIdentity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_monadRecFunction'] = function() { $v = (object)["tailRecM" => (function() {
  $__fn = function($f_0 = null, $a0_1 = null, $e_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_Rec_Class_tailRec'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_tailRec')))(function($a_3 = null) use ($e_2, $f_0) {
  $__num = \func_num_args();
  $__res = (($f_0)($a_3))($e_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($a0_1);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_monadFn'] ?? \PhpursThunks::eval('Control_Monad_monadFn'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_monadRecEither'] = function() { $v = (object)["tailRecM" => (function() {
  $__fn = function($f_0 = null, $a0_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Monad_Rec_Class_tailRec'] ?? \PhpursThunks::eval('Control_Monad_Rec_Class_tailRec')))(function($v_2 = null) use ($f_0) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Left"))) {
$__t0 = new Phpurs_Data1("Done", new Phpurs_Data1("Left", ($v_2)->{'value0'}));
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "Right"))) {
$__t1 = null;;
if ((is_object(($v_2)->{'value0'}) && ((($v_2)->{'value0'})->{'tag'} === "Loop"))) {
$__t1 = new Phpurs_Data1("Loop", ($f_0)((($v_2)->{'value0'})->{'value0'}));
goto end_branch_1;;
};
if ((is_object(($v_2)->{'value0'}) && ((($v_2)->{'value0'})->{'tag'} === "Done"))) {
$__t1 = new Phpurs_Data1("Done", new Phpurs_Data1("Right", (($v_2)->{'value0'})->{'value0'}));
goto end_branch_1;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t1 = null;
end_branch_1:;
$__t0 = $__t1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($f_0)($a0_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Either_monadEither'] ?? \PhpursThunks::eval('Data_Either_monadEither'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_monadRecEffect'] = function() { $v = (object)["tailRecM" => (function() {
  $__fn = function($f_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($f_0)($a_1);
  $__local_var_3_1 = $__local_var_2_0;
  $r_4_2 = "TODO_PrimEffect";
  $v_5_3 = "TODO_PrimEffect";
  $__t4 = null;;
  if ((is_object($v_5_3) && (($v_5_3)->{'tag'} === "Loop"))) {
$e_6_5 = ($f_0)(($v_5_3)->{'value0'});
$dollar__unused_7_6 = "TODO_PrimEffect";
$__t4 = false;
goto end_branch_4;;
};
  if ((is_object($v_5_3) && (($v_5_3)->{'tag'} === "Done"))) {
$__t4 = true;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $dollar__unused_5_3 = (($GLOBALS['Effect_untilE'] ?? \PhpursThunks::eval('Effect_untilE')))($__t4);
  $__res = (((($GLOBALS['Effect_applyEffect'] ?? \PhpursThunks::eval('Effect_applyEffect')))->{'apply'})(function($v_6 = null) {
  $__num = \func_num_args();
  $__t8 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Done"))) {
$__t8 = ($v_6)->{'value0'};
goto end_branch_8;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t8 = null;
  end_branch_8:;
  $__res = $__t8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))("TODO_PrimEffect");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_monadEffect'] ?? \PhpursThunks::eval('Effect_monadEffect'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_loop3'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null, $c_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = new Phpurs_Data1("Loop", (object)["a" => $a_0, "b" => $b_1, "c" => $c_2]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_loop2'] = function() { $v = (function() {
  $__fn = function($a_0 = null, $b_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data1("Loop", (object)["a" => $a_0, "b" => $b_1]);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_functorStep'] = function() { $v = (object)["map" => (function() {
  $__fn = function($f_0 = null, $m_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if ((is_object($m_1) && (($m_1)->{'tag'} === "Loop"))) {
$__t0 = new Phpurs_Data1("Loop", ($m_1)->{'value0'});
goto end_branch_0;;
};
  if ((is_object($m_1) && (($m_1)->{'tag'} === "Done"))) {
$__t0 = new Phpurs_Data1("Done", ($f_0)(($m_1)->{'value0'}));
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
})()]; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_forever'] = function() { $v = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (((((((($dictMonadRec_0)->{'Monad0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Bind1'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Apply0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->{'Functor0'})(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($ma_2 = null) use ($__local_var_1_0, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($u_3 = null) use ($__local_var_1_0, $ma_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_1_0)->{'map'})(function($v_4 = null) use ($u_3) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data1("Loop", $u_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($ma_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_Rec_Class_bifunctorStep'] = function() { $v = (object)["bimap" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Loop"))) {
$__t0 = new Phpurs_Data1("Loop", ($v_0)(($v2_2)->{'value0'}));
goto end_branch_0;;
};
  if ((is_object($v2_2) && (($v2_2)->{'tag'} === "Done"))) {
$__t0 = new Phpurs_Data1("Done", ($v1_1)(($v2_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})()]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };






















