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
function majTest_majFilemajOps_loopmajE(int $v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majTest_majFilemajOps_loopmajE';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Test_FileOps;
  $f = ($ffi_Test_FileOps['loopE'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Test_FileOps_loopE'] = __NAMESPACE__ . '\\majTest_majFilemajOps_loopmajE';

function majTest_majFilemajOps_readmajFilemajSync(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majTest_majFilemajOps_readmajFilemajSync';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Test_FileOps;
  $f = ($ffi_Test_FileOps['readFileSync'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Test_FileOps_readFileSync'] = __NAMESPACE__ . '\\majTest_majFilemajOps_readmajFilemajSync';

function majTest_majFilemajOps_writemajFilemajSync(string $v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majTest_majFilemajOps_writemajFilemajSync';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Test_FileOps;
  $f = ($ffi_Test_FileOps['writeFileSync'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Test_FileOps_writeFileSync'] = __NAMESPACE__ . '\\majTest_majFilemajOps_writemajFilemajSync';





// Test_FileOps_discard
$GLOBALS['Test_FileOps_discard'] = (($GLOBALS['Control_Bind_discardUnit'])->{'discard'})($GLOBALS['Effect_bindEffect']);

// Test_FileOps_loopIO
function majTest_majFilemajOps_loopmajImajO(int $n_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majTest_majFilemajOps_loopmajImajO';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Test\FileOps\majTest_majFilemajOps_loopmajE($n_0, (($GLOBALS['Test_FileOps_discard'])(\Test\FileOps\majTest_majFilemajOps_writemajFilemajSync("var/iotest.txt", "Hello IO Benchmarks!")))(function($_dollar__unused_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])->{'bind'})(\Test\FileOps\majTest_majFilemajOps_readmajFilemajSync("var/iotest.txt")))(function($_dollar__unused_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_applicativeEffect'])->{'pure'})($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Test_FileOps_loopIO'] = __NAMESPACE__ . '\\majTest_majFilemajOps_loopmajImajO';

// Test_FileOps_describe
$GLOBALS['Test_FileOps_describe'] = \Effect\Console\majEffect_majConsole_log("File I/O (10k writes/reads):");

// Test_FileOps_act
$GLOBALS['Test_FileOps_act'] = \Test\FileOps\majTest_majFilemajOps_loopmajImajO(10000);

