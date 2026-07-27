<?php

namespace Control\Monad\ST\Uncurried;

// ALL IMPORTS: Control.Monad.ST.Internal, Prim
// TO REQUIRE: Control.Monad.ST.Internal
require_once __DIR__ . '/../Control.Monad.ST.Internal/index.php';

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
$ffi_Control_Monad_ST_Uncurried = \call_user_func(function() {
  $exports = [];
$mkSTFn1 = function($fn) use (&$mkSTFn1) {
    return function($x) use ($fn) { return $fn($x)(); };
};
$mkSTFn2 = function($fn) use (&$mkSTFn2) {
    return function($a, $b) use ($fn) { return $fn($a)($b)(); };
};
$mkSTFn3 = function($fn) use (&$mkSTFn3) {
    return function($a, $b, $c) use ($fn) { return $fn($a)($b)($c)(); };
};
$mkSTFn4 = function($fn) use (&$mkSTFn4) {
    return function($a, $b, $c, $d) use ($fn) { return $fn($a)($b)($c)($d)(); };
};
$mkSTFn5 = function($fn) use (&$mkSTFn5) {
    return function($a, $b, $c, $d, $e) use ($fn) { return $fn($a)($b)($c)($d)($e)(); };
};
$mkSTFn6 = function($fn) use (&$mkSTFn6) {
    return function($a, $b, $c, $d, $e, $f) use ($fn) { return $fn($a)($b)($c)($d)($e)($f)(); };
};
$mkSTFn7 = function($fn) use (&$mkSTFn7) {
    return function($a, $b, $c, $d, $e, $f, $g) use ($fn) { return $fn($a)($b)($c)($d)($e)($f)($g)(); };
};
$mkSTFn8 = function($fn) use (&$mkSTFn8) {
    return function($a, $b, $c, $d, $e, $f, $g, $h) use ($fn) { return $fn($a)($b)($c)($d)($e)($f)($g)($h)(); };
};
$mkSTFn9 = function($fn) use (&$mkSTFn9) {
    return function($a, $b, $c, $d, $e, $f, $g, $h, $i) use ($fn) { return $fn($a)($b)($c)($d)($e)($f)($g)($h)($i)(); };
};
$mkSTFn10 = function($fn) use (&$mkSTFn10) {
    return function($a, $b, $c, $d, $e, $f, $g, $h, $i, $j) use ($fn) { return $fn($a)($b)($c)($d)($e)($f)($g)($h)($i)($j)(); };
};

$runSTFn1 = function($fn, $a = null) use (&$runSTFn1) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn1) {

            return $runSTFn1(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a) { return $fn($a); };
};
$runSTFn2 = function($fn, $a = null, $b = null) use (&$runSTFn2) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn2) {

            return $runSTFn2(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b) { return $fn($a, $b); };
};
$runSTFn3 = function($fn, $a = null, $b = null, $c = null) use (&$runSTFn3) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn3) {

            return $runSTFn3(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c) { return $fn($a, $b, $c); };
};
$runSTFn4 = function($fn, $a = null, $b = null, $c = null, $d = null) use (&$runSTFn4) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn4) {

            return $runSTFn4(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c, $d) { return $fn($a, $b, $c, $d); };
};
$runSTFn5 = function($fn, $a = null, $b = null, $c = null, $d = null, $e = null) use (&$runSTFn5) {
    if (\func_num_args() < 6) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn5) {

            return $runSTFn5(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c, $d, $e) { return $fn($a, $b, $c, $d, $e); };
};
$runSTFn6 = function($fn, $a = null, $b = null, $c = null, $d = null, $e = null, $f = null) use (&$runSTFn6) {
    if (\func_num_args() < 7) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn6) {

            return $runSTFn6(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c, $d, $e, $f) { return $fn($a, $b, $c, $d, $e, $f); };
};
$runSTFn7 = function($fn, $a = null, $b = null, $c = null, $d = null, $e = null, $f = null, $g = null) use (&$runSTFn7) {
    if (\func_num_args() < 8) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn7) {

            return $runSTFn7(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c, $d, $e, $f, $g) { return $fn($a, $b, $c, $d, $e, $f, $g); };
};
$runSTFn8 = function($fn, $a = null, $b = null, $c = null, $d = null, $e = null, $f = null, $g = null, $h = null) use (&$runSTFn8) {
    if (\func_num_args() < 9) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn8) {

            return $runSTFn8(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c, $d, $e, $f, $g, $h) { return $fn($a, $b, $c, $d, $e, $f, $g, $h); };
};
$runSTFn9 = function($fn, $a = null, $b = null, $c = null, $d = null, $e = null, $f = null, $g = null, $h = null, $i = null) use (&$runSTFn9) {
    if (\func_num_args() < 10) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn9) {

            return $runSTFn9(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c, $d, $e, $f, $g, $h, $i) { return $fn($a, $b, $c, $d, $e, $f, $g, $h, $i); };
};
$runSTFn10 = function($fn, $a = null, $b = null, $c = null, $d = null, $e = null, $f = null, $g = null, $h = null, $i = null, $j = null) use (&$runSTFn10) {
    if (\func_num_args() < 11) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$runSTFn10) {

            return $runSTFn10(...\array_merge($__args, $more));
        };
    }
    return function() use ($fn, $a, $b, $c, $d, $e, $f, $g, $h, $i, $j) { return $fn($a, $b, $c, $d, $e, $f, $g, $h, $i, $j); };
};

$exports['mkSTFn1'] = $mkSTFn1;
$exports['mkSTFn2'] = $mkSTFn2;
$exports['mkSTFn3'] = $mkSTFn3;
$exports['mkSTFn4'] = $mkSTFn4;
$exports['mkSTFn5'] = $mkSTFn5;
$exports['mkSTFn6'] = $mkSTFn6;
$exports['mkSTFn7'] = $mkSTFn7;
$exports['mkSTFn8'] = $mkSTFn8;
$exports['mkSTFn9'] = $mkSTFn9;
$exports['mkSTFn10'] = $mkSTFn10;
$exports['runSTFn1'] = $runSTFn1;
$exports['runSTFn2'] = $runSTFn2;
$exports['runSTFn3'] = $runSTFn3;
$exports['runSTFn4'] = $runSTFn4;
$exports['runSTFn5'] = $runSTFn5;
$exports['runSTFn6'] = $runSTFn6;
$exports['runSTFn7'] = $runSTFn7;
$exports['runSTFn8'] = $runSTFn8;
$exports['runSTFn9'] = $runSTFn9;
$exports['runSTFn10'] = $runSTFn10;
return $exports;
  return $exports;
});
function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn1($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn1'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn1'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn1';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn10($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn10';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn10'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn10'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn10';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn2($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn2';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn2'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn2'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn2';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn3($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn3';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn3'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn3'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn3';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn4($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn4';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn4'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn4'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn4';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn5($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn5';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn5'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn5'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn5';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn6($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn6';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn6'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn6'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn6';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn7($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn7';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn7'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn7'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn7';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn8($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn8';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn8'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn8'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn8';

function majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn9($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn9';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['mkSTFn9'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Control_Monad_ST_Uncurried_mkSTFn9'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_mkmajSmajTmajFn9';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn1'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn1'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn10($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null, $v8 = null, $v9 = null, $v10 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn10';
  if ($__num < 11) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 11);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn10'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7, $v8, $v9, $v10);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn10'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn10';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn2'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn2'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn3($v0, $v1 = null, $v2 = null, $v3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn3';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn3'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn3'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn3';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn4($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn4';
  if ($__num < 5) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 5);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn4'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn4'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn4';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn5($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn5';
  if ($__num < 6) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 6);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn5'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn5'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn5';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn6($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn6';
  if ($__num < 7) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 7);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn6'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn6'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn6';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn7($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn7';
  if ($__num < 8) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 8);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn7'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn7'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn7';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn8($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null, $v8 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn8';
  if ($__num < 9) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 9);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn8'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7, $v8);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn8'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn8';

function majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn9($v0, $v1 = null, $v2 = null, $v3 = null, $v4 = null, $v5 = null, $v6 = null, $v7 = null, $v8 = null, $v9 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn9';
  if ($__num < 10) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 10);
  }
  global $ffi_Control_Monad_ST_Uncurried;
  $f = ($ffi_Control_Monad_ST_Uncurried['runSTFn9'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2, $v3, $v4, $v5, $v6, $v7, $v8, $v9);
}
$GLOBALS['Control_Monad_ST_Uncurried_runSTFn9'] = __NAMESPACE__ . '\\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn9';






