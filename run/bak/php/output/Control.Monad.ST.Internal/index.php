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
    return function() use($f, $a) { return $f($a()); };
};
$bind_ = function($a, $f = null) use (&$bind_) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$bind_) {

            return $bind_(...\array_merge($__args, $more));
        };
    }
    return function() use($a, $f) { return $f($a())(); };
};
$pure_ = function($a) { return function() use($a) { return $a; }; };
$new = function($val) { return function() use($val) { return (object)['value' => $val]; }; };
$read = function($ref) { return function() use($ref) { return $ref->value; }; };
$modifyImpl = function($f, $ref = null) use (&$modifyImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$modifyImpl) {

            return $modifyImpl(...\array_merge($__args, $more));
        };
    }
    return function() use($f, $ref) { $t = $f($ref->value); $ref->value = $t->state; return $t->value; };
};
$write = function($val, $ref = null) use (&$write) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$write) {

            return $write(...\array_merge($__args, $more));
        };
    }
    return function() use($val, $ref) { $ref->value = $val; return $val; };
};
$run = function($f) { return $f(); };
$while = function($f, $a = null) use (&$while) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$while) {

            return $while(...\array_merge($__args, $more));
        };
    }
    return function() use($f, $a) { while ($f()) { $a(); } return null; };
};
$for = function($lo, $hi = null, $f = null) use (&$for) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$for) {

            return $for(...\array_merge($__args, $more));
        };
    }
    return function() use($lo, $hi, $f) { for ($i = $lo; $i < $hi; $i++) { $f($i)(); } return null; };
};
$foreach = function($as, $f = null) use (&$foreach) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$foreach) {

            return $foreach(...\array_merge($__args, $more));
        };
    }
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
  $f = ($ffi_Control_Monad_ST_Internal['bind_'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['for'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['foreach'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['map_'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['modifyImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['new'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['pure_'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['read'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['run'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['while'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $f = ($ffi_Control_Monad_ST_Internal['write'] ?? new class { public function __invoke(...$args) { return $this; } });
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
  $s_prime_2_0 = ($f_0)($s_1);
  $__res = (object)["state" => $s_prime_2_0, "value" => $s_prime_2_0];
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

// Control_Monad_ST_Internal_void
$GLOBALS['Control_Monad_ST_Internal_void'] = (($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(function($v_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Control_Monad_ST_Internal_monadST
$GLOBALS['Control_Monad_ST_Internal_monadST'] = (object)["Applicative0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_applicativeST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_bindST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_bindST
$GLOBALS['Control_Monad_ST_Internal_bindST'] = (object)["bind" => $GLOBALS['Control_Monad_ST_Internal_bind_'], "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_applyST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_applyST
$GLOBALS['Control_Monad_ST_Internal_applyST'] = (function() use (&$__fn) {
$__local_var_0_0 = (($GLOBALS['Control_Monad_ST_Internal_monadST'])->{'Bind1'})(null);
return (object)["apply" => (function() use ($__local_var_0_0) {
  $__fn = function($f_1, $a_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_0_0)->{'bind'})($f_1))(function($f_prime_3) use ($__local_var_0_0, $a_2) {
  $__num = \func_num_args();
  $__res = ((($__local_var_0_0)->{'bind'})($a_2))(function($a_prime_4) use ($f_prime_3) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Control_Monad_ST_Internal_monadST'])->{'Applicative0'})(null))->{'pure'})(($f_prime_3)($a_prime_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_functorST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
})();

// Control_Monad_ST_Internal_applicativeST
$GLOBALS['Control_Monad_ST_Internal_applicativeST'] = (object)["pure" => $GLOBALS['Control_Monad_ST_Internal_pure_'], "Apply0" => function($_dollar__unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Control_Monad_ST_Internal_applyST'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Control_Monad_ST_Internal_lift2
function majControl_majMonad_majSmajT_majInternal_lift2($f_0, $a_1 = null, $b_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majSmajT_majInternal_lift2';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_applyST'])->{'apply'})(((((($GLOBALS['Control_Monad_ST_Internal_applyST'])->{'Functor0'})(null))->{'map'})($f_0))($a_1)))($b_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Control_Monad_ST_Internal_lift2'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_lift2';

// Control_Monad_ST_Internal_discard
$GLOBALS['Control_Monad_ST_Internal_discard'] = (($GLOBALS['Control_Bind_discardUnit'])->{'discard'})($GLOBALS['Control_Monad_ST_Internal_bindST']);

// Control_Monad_ST_Internal_semigroupST
function majControl_majMonad_majSmajT_majInternal_semigroupmajSmajT($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majControl_majMonad_majSmajT_majInternal_semigroupmajSmajT';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (object)["append" => ($GLOBALS['Control_Monad_ST_Internal_lift2'])(($dictSemigroup_0)->{'append'})];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_ST_Internal_semigroupST'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_semigroupmajSmajT';

// Control_Monad_ST_Internal_monadRecST
$GLOBALS['Control_Monad_ST_Internal_monadRecST'] = (object)["tailRecM" => (function() {
  $__fn = function($f_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(($f_0)($a_1)))($GLOBALS['Control_Monad_ST_Internal_new'])))(function($r_2) use ($f_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_discard'])(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_while(((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(function($v_3) {
  $__num = \func_num_args();
  $__res = $v_3 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_read($r_2)), ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_read($r_2)))(function($v_3) use ($f_0, $r_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_3 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t0 = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(($f_0)(($v_3)->{'value0'})))(function($e_4) use ($r_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_ST_Internal_void'])(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_write($e_4, $r_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
goto end_branch_0;;
};
  if ($v_3 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t0 = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})($GLOBALS['Data_Unit_unit']);
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(function($_dollar__unused_3) use ($r_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(function($v_4) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($v_4 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t1 = ($v_4)->{'value0'};
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_read($r_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Monad0" => function($_dollar__unused_0) {
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
  $__res = (object)["mempty" => (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})(($dictMonoid_0)->{'mempty'}), "Semigroup0" => function($_dollar__unused_1) use ($dictMonoid_0) {
  $__num = \func_num_args();
  $__res = (object)["append" => ($GLOBALS['Control_Monad_ST_Internal_lift2'])(((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'})];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Control_Monad_ST_Internal_monoidST'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majInternal_monoidmajSmajT';

