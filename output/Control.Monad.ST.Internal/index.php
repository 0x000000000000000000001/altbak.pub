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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Control_Monad_ST_Internal = \call_user_func(function() {
  $exports = [];
$map_ = function($f, $a) use (&$map_) {
    return function() use($f, $a) { return $f($a()); };
};
$bind_ = function($a, $f) use (&$bind_) {
    return function() use($a, $f) { return $f($a())(); };
};
$pure_ = function($a) { return function() use($a) { return $a; }; };
$new = function($val) { return function() use($val) { return (object)['value' => $val]; }; };
$read = function($ref) { return function() use($ref) { return $ref->value; }; };
$modifyImpl = function($f, $ref) use (&$modifyImpl) {
    return function() use($f, $ref) { $t = $f($ref->value); $ref->value = $t->state; return $t->value; };
};
$write = function($val, $ref) use (&$write) {
    return function() use($val, $ref) { $ref->value = $val; return $val; };
};
$run = function($f) { return $f(); };
$while = function($f, $a) use (&$while) {
    return function() use($f, $a) { while ($f()) { $a(); } return null; };
};
$for = function($lo, $hi, $f) use (&$for) {
    return function() use($lo, $hi, $f) { for ($i = $lo; $i < $hi; $i++) { $f($i)(); } return null; };
};
$foreach = function($as, $f) use (&$foreach) {
    return function() use($as, $f) { foreach ($as as $a) { $f($a)(); } return null; };
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
function majControl_majMonad_majSmajT_majInternal_bind_($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_bind_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('bind_', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['bind_'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Monad_ST_Internal_bind_'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_bind_';

function majControl_majMonad_majSmajT_majInternal_for(int $v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_for';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('for', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['for'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Control_Monad_ST_Internal_for'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_for';

function majControl_majMonad_majSmajT_majInternal_foreach($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_foreach';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('foreach', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['foreach'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Monad_ST_Internal_foreach'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_foreach';

function majControl_majMonad_majSmajT_majInternal_map_($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_map_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('map_', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['map_'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Monad_ST_Internal_map_'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_map_';

function majControl_majMonad_majSmajT_majInternal_modifymajImpl($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_modifymajImpl';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('modifyImpl', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['modifyImpl'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Monad_ST_Internal_modifyImpl'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_modifymajImpl';

function majControl_majMonad_majSmajT_majInternal_new($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_new';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('new', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['new'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Internal_new'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_new';

function majControl_majMonad_majSmajT_majInternal_pure_($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_pure_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('pure_', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['pure_'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Internal_pure_'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_pure_';

function majControl_majMonad_majSmajT_majInternal_read($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_read';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('read', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['read'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Internal_read'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_read';

function majControl_majMonad_majSmajT_majInternal_run($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_run';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('run', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['run'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Internal_run'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_run';

function majControl_majMonad_majSmajT_majInternal_while($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_while';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('while', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['while'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Monad_ST_Internal_while'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_while';

function majControl_majMonad_majSmajT_majInternal_write($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_write';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Monad_ST_Internal;
  $f = (\array_key_exists('write', $ffi_Control_Monad_ST_Internal) ? $ffi_Control_Monad_ST_Internal['write'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Monad_ST_Internal_write'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_write';





// Control_Monad_ST_Internal_modify'_closure
$GLOBALS['Control_Monad_ST_Internal_modify__prime___closure'] = $GLOBALS['Control_Monad_ST_Internal_modifyImpl'];

// Control_Monad_ST_Internal_modify'
function majControl_majMonad_majSmajT_majInternal_modify__prime__($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majSmajT_majInternal_modify__prime__';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_modify__prime___closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Control_Monad_ST_Internal_modify__prime__'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_modify__prime__';

// Control_Monad_ST_Internal_modify
function majControl_majMonad_majSmajT_majInternal_modify($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majSmajT_majInternal_modify';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_modifyImpl'])(function($s_1) use ($f_0) {
  $__num = \func_num_args();
  $s_prime__2_0 = ($f_0)($s_1);
  $__res = (object)["state" => $s_prime__2_0, "value" => $s_prime__2_0];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_ST_Internal_modify'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_modify';

// Control_Monad_ST_Internal_functorST
$GLOBALS['Control_Monad_ST_Internal_functorST'] = (object)["map" => $GLOBALS['Control_Monad_ST_Internal_map_']];

// Control_Monad_ST_Internal_monadST
$GLOBALS['Control_Monad_ST_Internal_monadST'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_applicativeST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_bindST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_bindST
$GLOBALS['Control_Monad_ST_Internal_bindST'] = (object)["bind" => $GLOBALS['Control_Monad_ST_Internal_bind_'], "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_applyST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_applyST
$GLOBALS['Control_Monad_ST_Internal_applyST'] = (object)["apply" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = function() use ($a_1, $f_0, &$__fn) {
$f_prime__2_0 = phpurs_execute_effect($f_0);
$a_prime__3_1 = phpurs_execute_effect($a_1);
return phpurs_execute_effect(phpurs_execute_effect((($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})(($f_prime__2_0)($a_prime__3_1))));
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_functorST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_applicativeST
$GLOBALS['Control_Monad_ST_Internal_applicativeST'] = (object)["pure" => $GLOBALS['Control_Monad_ST_Internal_pure_'], "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_applyST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_semigroupST
function majControl_majMonad_majSmajT_majInternal_semigroupmajSmajT($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majSmajT_majInternal_semigroupmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => function($a_1) use ($dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function($b_2) use ($a_1, $dictSemigroup_0) {
  $__num = \func_num_args();
  $__res = function() use ($a_1, $b_2, $dictSemigroup_0, &$__fn) {
$__local_var_3_0 = phpurs_execute_effect($a_1);
$f_prime__3_0 = phpurs_execute_effect((($dictSemigroup_0)->{'append'})($__local_var_3_0));
$a_prime__4_2 = phpurs_execute_effect($b_2);
return phpurs_execute_effect(phpurs_execute_effect(($f_prime__3_0)($a_prime__4_2)));
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_ST_Internal_semigroupST'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_semigroupmajSmajT';

// Control_Monad_ST_Internal_monadRecST
$GLOBALS['Control_Monad_ST_Internal_monadRecST'] = (object)["tailRecM" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($f_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = ($f_0)($a_1);
  $__res = function() use ($__local_var_2_0, $f_0, &$__fn) {
$__local_var_3_1 = phpurs_execute_effect($__local_var_2_0);
$r_3_1 = phpurs_execute_effect(phpurs_execute_effect("TODO_PrimEffect"));
$__local_var_4_3 = phpurs_execute_effect("TODO_PrimEffect");
$v_4_4 = "TODO_PrimEffect";
$__t5 = null;;
if ($v_4_4 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$e_5_6 = phpurs_execute_effect(($f_0)(($v_4_4)->{'value0'}));
$__local_var_6_7 = phpurs_execute_effect("TODO_PrimEffect");
$__t5 = phpurs_execute_effect($GLOBALS['Data_Unit_unit']);
goto end_branch_5;;
};
if ($v_4_4 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t5 = $GLOBALS['Data_Unit_unit'];
goto end_branch_5;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t5 = null;
end_branch_5:;
$_dollar___unused_4_3 = phpurs_execute_effect(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_while($__local_var_4_3 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop, phpurs_execute_effect($__t5)));
$__local_var_5_9 = phpurs_execute_effect("TODO_PrimEffect");
$__t10 = null;;
if ($__local_var_5_9 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t10 = ($__local_var_5_9)->{'value0'};
goto end_branch_10;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t10 = null;
end_branch_10:;
return phpurs_execute_effect(phpurs_execute_effect($__t10));
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_monadST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_monoidST
function majControl_majMonad_majSmajT_majInternal_monoidmajSmajT($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majSmajT_majInternal_monoidmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $semigroupST1_1_0 = (object)["append" => function($a_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($__local_var_1_0, $a_2) {
  $__num = \func_num_args();
  $__res = function() use ($__local_var_1_0, $a_2, $b_3, &$__fn) {
$__local_var_4_1 = phpurs_execute_effect($a_2);
$f_prime__4_1 = phpurs_execute_effect((($__local_var_1_0)->{'append'})($__local_var_4_1));
$a_prime__5_3 = phpurs_execute_effect($b_3);
return phpurs_execute_effect(phpurs_execute_effect(($f_prime__4_1)($a_prime__5_3)));
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__local_var_2_5 = ($dictMonoid_0)->{'mempty'};
  $__res = (object)["mempty" => function() use ($__local_var_2_5, &$__fn) {
return $__local_var_2_5;
}, "Semigroup0" => function($_dollar___unused_2) use ($semigroupST1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupST1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_ST_Internal_monoidST'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_monoidmajSmajT';

