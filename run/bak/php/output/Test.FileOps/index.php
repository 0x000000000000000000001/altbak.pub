<?php

namespace Test\FileOps;

// ALL IMPORTS: Control.Applicative, Control.Bind, Data.Unit, Effect, Effect.Console, Prelude, Prim, Test.FileOps
// TO REQUIRE: Control.Applicative, Control.Bind, Data.Unit, Effect, Effect.Console, Prelude, Test.FileOps
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Test.FileOps/index.php';

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
\PhpursThunks::$thunks['Test_FileOps_loopIO'] = function() { $v = function($n_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Test_FileOps_loopIO"), recVars=[];
  $__local_var_1_0 = ((($GLOBALS['Test_FileOps_writeFileSync'] ?? \PhpursThunks::eval('Test_FileOps_writeFileSync')))("var/iotest.txt"))("Hello IO Benchmarks!");
  $dollar__unused_2_1 = ($__local_var_1_0)();
  $dollar__unused_3_2 = ((($GLOBALS['Test_FileOps_readFileSync'] ?? \PhpursThunks::eval('Test_FileOps_readFileSync')))("var/iotest.txt"))();
  $__res = ((($GLOBALS['Test_FileOps_loopE'] ?? \PhpursThunks::eval('Test_FileOps_loopE')))($n_0))(($GLOBALS['Data_Unit_unit'] ?? \PhpursThunks::eval('Data_Unit_unit')));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Test_FileOps_describe'] = function() { $v = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))("File I/O (10k writes/reads):"); return $v; };
\PhpursThunks::$thunks['Test_FileOps_act'] = function() { $v = (($GLOBALS['Test_FileOps_loopIO'] ?? \PhpursThunks::eval('Test_FileOps_loopIO')))(10000); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Test_FileOps = \call_user_func(function() {
  $exports = [];
$writeFileSync = function($path, $content = null) use (&$writeFileSync) {
    if (func_num_args() < 2) {
        $__args = func_get_args();
        return function(...$more) use ($__args, &$writeFileSync) {
            return $writeFileSync(...array_merge($__args, $more));
        };
    }
    return function() use ($path, $content) {
        file_put_contents($path, $content);
    };
};

$readFileSync = function($path) {
    return function() use ($path) {
        return file_get_contents($path);
    };
};

$loopE = function($n, $action = null) use (&$loopE) {
    if (func_num_args() < 2) {
        $__args = func_get_args();
        return function(...$more) use ($__args, &$loopE) {
            return $loopE(...array_merge($__args, $more));
        };
    }
    return function() use ($n, $action) {
        for ($i = 0; $i < $n; $i++) {
            $action();
        }
    };
};

$exports['writeFileSync'] = $writeFileSync;
$exports['readFileSync'] = $readFileSync;
$exports['loopE'] = $loopE;

return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Test_FileOps_loopE'] = function() use (&$ffi_Test_FileOps) { return $ffi_Test_FileOps['loopE']; };
\PhpursThunks::$thunks['Test_FileOps_readFileSync'] = function() use (&$ffi_Test_FileOps) { return $ffi_Test_FileOps['readFileSync']; };
\PhpursThunks::$thunks['Test_FileOps_writeFileSync'] = function() use (&$ffi_Test_FileOps) { return $ffi_Test_FileOps['writeFileSync']; };





