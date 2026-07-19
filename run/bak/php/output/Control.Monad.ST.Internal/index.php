<?php

namespace Control\Monad\ST\Internal;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Rec.Class, Control.Monad.ST.Internal, Data.Functor, Data.Monoid, Data.Semigroup, Data.Unit, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Bind, Control.Monad, Control.Monad.Rec.Class, Control.Monad.ST.Internal, Data.Functor, Data.Monoid, Data.Semigroup, Data.Unit, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Internal/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
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
\PhpursThunks::$thunks['Control_Monad_ST_Internal_modify__prime__'] = function() { $v = ($GLOBALS['Control_Monad_ST_Internal_modifyImpl'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_modifyImpl')); return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_modify'] = function() { $v = function($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_ST_Internal_modify"), recVars=[];
  $__res = (($GLOBALS['Control_Monad_ST_Internal_modifyImpl'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_modifyImpl')))(function($s_1) use ($f_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $s__prime___2_0 = ($f_0)($s_1);
  $__res = (object)["state" => $s__prime___2_0, "value" => $s__prime___2_0];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_functorST'] = function() { $v = (object)["map" => ($GLOBALS['Control_Monad_ST_Internal_map_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_map_'))]; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_monadST'] = function() { $v = (object)["Applicative0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_ST_Internal_monadST","Control_Monad_ST_Internal_bindST","Control_Monad_ST_Internal_applyST","Control_Monad_ST_Internal_applicativeST"];
  $__res = ($GLOBALS['Control_Monad_ST_Internal_applicativeST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_applicativeST'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_ST_Internal_monadST","Control_Monad_ST_Internal_bindST","Control_Monad_ST_Internal_applyST","Control_Monad_ST_Internal_applicativeST"];
  $__res = ($GLOBALS['Control_Monad_ST_Internal_bindST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bindST'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_bindST'] = function() { $v = (object)["bind" => ($GLOBALS['Control_Monad_ST_Internal_bind_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_bind_')), "Apply0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_ST_Internal_monadST","Control_Monad_ST_Internal_bindST","Control_Monad_ST_Internal_applyST","Control_Monad_ST_Internal_applicativeST"];
  $__res = ($GLOBALS['Control_Monad_ST_Internal_applyST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_applyST'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_applyST'] = function() { $v = (object)["apply" => (($GLOBALS['Control_Monad_ap'] ?? \PhpursThunks::eval('Control_Monad_ap')))(($GLOBALS['Control_Monad_ST_Internal_monadST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_monadST'))), "Functor0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_ST_Internal_monadST","Control_Monad_ST_Internal_bindST","Control_Monad_ST_Internal_applyST","Control_Monad_ST_Internal_applicativeST"];
  $__res = ($GLOBALS['Control_Monad_ST_Internal_functorST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_functorST'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_applicativeST'] = function() { $v = (object)["pure" => ($GLOBALS['Control_Monad_ST_Internal_pure_'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_pure_')), "Apply0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=["Control_Monad_ST_Internal_monadST","Control_Monad_ST_Internal_bindST","Control_Monad_ST_Internal_applyST","Control_Monad_ST_Internal_applicativeST"];
  $__res = ($GLOBALS['Control_Monad_ST_Internal_applyST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_applyST'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_semigroupST'] = function() { $v = function($dictSemigroup_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_ST_Internal_semigroupST"), recVars=[];
  $__res = (object)["append" => ((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))(($GLOBALS['Control_Monad_ST_Internal_applyST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_applyST'))))(($dictSemigroup_0)->append)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_monadRecST'] = function() { $v = (object)["tailRecM" => (function() {
  $__fn = function($f_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__local_var_2_0 = ($f_0)($a_1);
  $__local_var_3_1 = ($__local_var_2_0)();
  $r_4_2 = ("TODO_PrimEffect")();
  $__local_var_5_3 = ("TODO_PrimEffect")();
  $v_5_4 = ("TODO_PrimEffect")();
  if ((is_object($v_5_4) && (($v_5_4)->tag === "Loop"))) {
$e_6_6 = (($f_0)(($v_5_4)->value0))();
$__local_var_7_7 = ("TODO_PrimEffect")();
$__t5 = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
} else {
if ((is_object($v_5_4) && (($v_5_4)->tag === "Done"))) {
$__t5 = ($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit'));
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
};
};
  $dollar__unused_5_3 = (((($GLOBALS['Control_Monad_ST_Internal_while'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_while')))((is_object($__local_var_5_3) && (($__local_var_5_3)->tag === "Loop"))))($__t5))();
  $__local_var_6_9 = ("TODO_PrimEffect")();
  if ((is_object($__local_var_6_9) && (($__local_var_6_9)->tag === "Done"))) {
$__t10 = ($__local_var_6_9)->value0;
} else {
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
};
  $__res = $__t10;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($dollar__unused_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ($GLOBALS['Control_Monad_ST_Internal_monadST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_monadST'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_monoidST'] = function() { $v = function($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Control_Monad_ST_Internal_monoidST"), recVars=[];
  $__local_var_1_0 = ($dictMonoid_0)->mempty;
  $__res = (object)["mempty" => $__local_var_1_0, "Semigroup0" => function($dollar__unused_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (object)["append" => ((($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))(($GLOBALS['Control_Monad_ST_Internal_applyST'] ?? \PhpursThunks::eval('Control_Monad_ST_Internal_applyST'))))(((($dictMonoid_0)->Semigroup0)(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))->append)];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Control_Monad_ST_Internal = \call_user_func(function() {
  $exports = [];
$map_ = function($f, $a = null) use (&$map_) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$map_) {

            return $map_(...\array_merge($__args, $more));
        };
    }
    return function() use(&$f, &$a) { return $f($a()); };
};
$bind_ = function($a, $f = null) use (&$bind_) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$bind_) {

            return $bind_(...\array_merge($__args, $more));
        };
    }
    return function() use(&$a, &$f) { return $f($a())(); };
};
$pure_ = function($a) use (&$pure_) { return function() use(&$a) { return $a; }; };
$new = function($val) use (&$new) { return function() use(&$val) { return (object)['value' => $val]; }; };
$read = function($ref) use (&$read) { return function() use(&$ref) { return $ref->value; }; };
$modifyImpl = function($f, $ref = null) use (&$modifyImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$modifyImpl) {

            return $modifyImpl(...\array_merge($__args, $more));
        };
    }
    return function() use(&$f, &$ref) { $t = $f($ref->value); $ref->value = $t->state; return $t->value; };
};
$write = function($val, $ref = null) use (&$write) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$write) {

            return $write(...\array_merge($__args, $more));
        };
    }
    return function() use(&$val, &$ref) { $ref->value = $val; return $val; };
};
$run = function($f) use (&$run) { return $f(); };
$while = function($f, $a = null) use (&$while) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$while) {

            return $while(...\array_merge($__args, $more));
        };
    }
    return function() use(&$f, &$a) { while ($f()) { $a(); } return null; };
};
$for = function($lo, $hi = null, $f = null) use (&$for) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$for) {

            return $for(...\array_merge($__args, $more));
        };
    }
    return function() use(&$lo, &$hi, &$f) { for ($i = $lo; $i < $hi; $i++) { $f($i)(); } return null; };
};
$foreach = function($as, $f = null) use (&$foreach) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foreach) {

            return $foreach(...\array_merge($__args, $more));
        };
    }
    return function() use(&$as, &$f) { foreach ($as as $a) { $f($a)(); } return null; };
};

$exports['map_'] = $map_;
$exports['bind_'] = $bind_;
$exports['pure_'] = $pure_;
$exports['new'] = $new;
$exports['read'] = $read;
$exports['modifyImpl'] = $modifyImpl;
$exports['write'] = $write;
$exports['run'] = $run;
$exports['while'] = $while;
$exports['for'] = $for;
$exports['foreach'] = $foreach;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Control_Monad_ST_Internal_bind_'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['bind_']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_for'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['for']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_foreach'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['foreach']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_map_'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['map_']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_modifyImpl'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['modifyImpl']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_new'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['new']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_pure_'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['pure_']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_read'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['read']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_run'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['run']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_while'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['while']; };
\PhpursThunks::$thunks['Control_Monad_ST_Internal_write'] = function() use (&$ffi_Control_Monad_ST_Internal) { return $ffi_Control_Monad_ST_Internal['write']; };












