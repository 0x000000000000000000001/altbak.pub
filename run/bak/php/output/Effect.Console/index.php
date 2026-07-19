<?php

namespace Effect\Console;

// ALL IMPORTS: Control.Applicative, Control.Bind, Data.Show, Data.Unit, Effect, Effect.Console, Prim
// TO REQUIRE: Control.Applicative, Control.Bind, Data.Show, Data.Unit, Effect, Effect.Console
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Console/index.php';

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
\PhpursThunks::$thunks['Effect_Console_warnShow'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Effect_Console_warnShow"), recVars=[];
  $__res = (($GLOBALS['Effect_Console_warn'] ?? \PhpursThunks::eval('Effect_Console_warn')))((($dictShow_0)->show)($a_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Console_logShow'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Effect_Console_logShow"), recVars=[];
  $__res = (($GLOBALS['Effect_Console_log'] ?? \PhpursThunks::eval('Effect_Console_log')))((($dictShow_0)->show)($a_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Console_infoShow'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Effect_Console_infoShow"), recVars=[];
  $__res = (($GLOBALS['Effect_Console_info'] ?? \PhpursThunks::eval('Effect_Console_info')))((($dictShow_0)->show)($a_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Console_grouped'] = function() { $v = (function() {
  $__fn = function($name_0, $inner_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Effect_Console_grouped"), recVars=[];
  $__res = ((($GLOBALS['Effect_bindE'] ?? \PhpursThunks::eval('Effect_bindE')))((($GLOBALS['Effect_Console_group'] ?? \PhpursThunks::eval('Effect_Console_group')))($name_0)))(function($dollar__unused_2) use ($inner_1) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Effect_bindE'] ?? \PhpursThunks::eval('Effect_bindE')))($inner_1))(function($result_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = ((($GLOBALS['Effect_bindE'] ?? \PhpursThunks::eval('Effect_bindE')))(($GLOBALS['Effect_Console_groupEnd'] ?? \PhpursThunks::eval('Effect_Console_groupEnd'))))(function($dollar__unused_4) use ($result_3) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=Nothing, recVars=[];
  $__res = (($GLOBALS['Effect_pureE'] ?? \PhpursThunks::eval('Effect_pureE')))($result_3);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Console_errorShow'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Effect_Console_errorShow"), recVars=[];
  $__res = (($GLOBALS['Effect_Console_error'] ?? \PhpursThunks::eval('Effect_Console_error')))((($dictShow_0)->show)($a_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Effect_Console_debugShow'] = function() { $v = (function() {
  $__fn = function($dictShow_0, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  // DEBUG UncurriedAbs: currentBindingName=(Just "Effect_Console_debugShow"), recVars=[];
  $__res = (($GLOBALS['Effect_Console_debug'] ?? \PhpursThunks::eval('Effect_Console_debug')))((($dictShow_0)->show)($a_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Effect_Console = \call_user_func(function() {
  $exports = [];
$log = function($s) { return function() use(&$s) { echo $s . "\n"; }; };
$warn = function($s) { return function() use(&$s) { file_put_contents('php://stderr', $s . "\n"); }; };
$error = function($s) { return function() use(&$s) { file_put_contents('php://stderr', $s . "\n"); }; };
$info = function($s) { return function() use(&$s) { echo $s . "\n"; }; };
$debug = function($s) { return function() use(&$s) { echo $s . "\n"; }; };

$time = function($s) { return function() {}; };
$timeLog = function($s) { return function() {}; };
$timeEnd = function($s) { return function() {}; };
$clear = function() {};
$group = function($s) { return function() {}; };
$groupCollapsed = function($s) { return function() {}; };
$groupEnd = function() {};

$exports['log'] = $log;
$exports['warn'] = $warn;
$exports['error'] = $error;
$exports['info'] = $info;
$exports['debug'] = $debug;
$exports['time'] = $time;
$exports['timeLog'] = $timeLog;
$exports['timeEnd'] = $timeEnd;
$exports['clear'] = $clear;
$exports['group'] = $group;
$exports['groupCollapsed'] = $groupCollapsed;
$exports['groupEnd'] = $groupEnd;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Effect_Console_clear'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['clear']; };
\PhpursThunks::$thunks['Effect_Console_debug'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['debug']; };
\PhpursThunks::$thunks['Effect_Console_error'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['error']; };
\PhpursThunks::$thunks['Effect_Console_group'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['group']; };
\PhpursThunks::$thunks['Effect_Console_groupCollapsed'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['groupCollapsed']; };
\PhpursThunks::$thunks['Effect_Console_groupEnd'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['groupEnd']; };
\PhpursThunks::$thunks['Effect_Console_info'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['info']; };
\PhpursThunks::$thunks['Effect_Console_log'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['log']; };
\PhpursThunks::$thunks['Effect_Console_time'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['time']; };
\PhpursThunks::$thunks['Effect_Console_timeEnd'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['timeEnd']; };
\PhpursThunks::$thunks['Effect_Console_timeLog'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['timeLog']; };
\PhpursThunks::$thunks['Effect_Console_warn'] = function() use (&$ffi_Effect_Console) { return $ffi_Effect_Console['warn']; };








