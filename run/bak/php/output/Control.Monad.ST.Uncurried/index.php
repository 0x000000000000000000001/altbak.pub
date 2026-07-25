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
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn1'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn1']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn10'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn10']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn2'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn2']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn3'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn3']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn4'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn4']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn5'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn5']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn6'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn6']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn7'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn7']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn8'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn8']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_mkSTFn9'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['mkSTFn9']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn1'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn1']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn10'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn10']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn2'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn2']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn3'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn3']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn4'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn4']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn5'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn5']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn6'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn6']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn7'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn7']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn8'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn8']; };
\PhpursThunks::$thunks['Control_Monad_ST_Uncurried_runSTFn9'] = function() use (&$ffi_Control_Monad_ST_Uncurried) { return $ffi_Control_Monad_ST_Uncurried['runSTFn9']; };



