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
$ffi_Effect_Console = \call_user_func(function() {
  $exports = [];
$log = function($s) { return function() use($s) { echo $s . "\n"; }; };
$warn = function($s) { return function() use($s) { file_put_contents('php://stderr', $s . "\n"); }; };
$error = function($s) { return function() use($s) { file_put_contents('php://stderr', $s . "\n"); }; };
$info = function($s) { return function() use($s) { echo $s . "\n"; }; };
$debug = function($s) { return function() use($s) { echo $s . "\n"; }; };

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
$GLOBALS['Effect_Console_clear'] = ($ffi_Effect_Console['clear'] ?? new class { public function __invoke(...$args) { return $this; } });
function majEffect_majConsole_debug(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_debug';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['debug'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_debug'] = __NAMESPACE__ . '\\majEffect_majConsole_debug';

function majEffect_majConsole_error(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_error';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['error'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_error'] = __NAMESPACE__ . '\\majEffect_majConsole_error';

function majEffect_majConsole_group(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_group';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['group'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_group'] = __NAMESPACE__ . '\\majEffect_majConsole_group';

function majEffect_majConsole_groupmajCollapsed(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_groupmajCollapsed';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['groupCollapsed'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_groupCollapsed'] = __NAMESPACE__ . '\\majEffect_majConsole_groupmajCollapsed';

$GLOBALS['Effect_Console_groupEnd'] = ($ffi_Effect_Console['groupEnd'] ?? new class { public function __invoke(...$args) { return $this; } });
function majEffect_majConsole_info(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_info';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['info'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_info'] = __NAMESPACE__ . '\\majEffect_majConsole_info';

function majEffect_majConsole_log(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_log';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['log'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_log'] = __NAMESPACE__ . '\\majEffect_majConsole_log';

function majEffect_majConsole_time(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_time';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['time'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_time'] = __NAMESPACE__ . '\\majEffect_majConsole_time';

function majEffect_majConsole_timemajEnd(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_timemajEnd';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['timeEnd'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_timeEnd'] = __NAMESPACE__ . '\\majEffect_majConsole_timemajEnd';

function majEffect_majConsole_timemajLog(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_timemajLog';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['timeLog'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_timeLog'] = __NAMESPACE__ . '\\majEffect_majConsole_timemajLog';

function majEffect_majConsole_warn(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majConsole_warn';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Console;
  $f = ($ffi_Effect_Console['warn'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Console_warn'] = __NAMESPACE__ . '\\majEffect_majConsole_warn';





// Effect_Console_discard
$GLOBALS['Effect_Console_discard'] = (($GLOBALS['Control_Bind_discardUnit'])->{'discard'})($GLOBALS['Effect_bindEffect']);

// Effect_Console_warnShow
function majEffect_majConsole_warnmajShow($dictShow_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majConsole_warnmajShow';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Console\majEffect_majConsole_warn((($dictShow_0)->{'show'})($a_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Console_warnShow'] = __NAMESPACE__ . '\\majEffect_majConsole_warnmajShow';

// Effect_Console_logShow
function majEffect_majConsole_logmajShow($dictShow_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majConsole_logmajShow';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Console\majEffect_majConsole_log((($dictShow_0)->{'show'})($a_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Console_logShow'] = __NAMESPACE__ . '\\majEffect_majConsole_logmajShow';

// Effect_Console_infoShow
function majEffect_majConsole_infomajShow($dictShow_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majConsole_infomajShow';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Console\majEffect_majConsole_info((($dictShow_0)->{'show'})($a_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Console_infoShow'] = __NAMESPACE__ . '\\majEffect_majConsole_infomajShow';

// Effect_Console_grouped
function majEffect_majConsole_grouped(string $name_0, $inner_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majConsole_grouped';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Console_discard'])(\Effect\Console\majEffect_majConsole_group($name_0)))(function($_dollar__unused_2) use ($inner_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Effect_bindEffect'])->{'bind'})($inner_1))(function($result_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Console_discard'])($GLOBALS['Effect_Console_groupEnd']))(function($_dollar__unused_4) use ($result_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_applicativeEffect'])->{'pure'})($result_3);
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
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Console_grouped'] = __NAMESPACE__ . '\\majEffect_majConsole_grouped';

// Effect_Console_errorShow
function majEffect_majConsole_errormajShow($dictShow_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majConsole_errormajShow';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Console\majEffect_majConsole_error((($dictShow_0)->{'show'})($a_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Console_errorShow'] = __NAMESPACE__ . '\\majEffect_majConsole_errormajShow';

// Effect_Console_debugShow
function majEffect_majConsole_debugmajShow($dictShow_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majConsole_debugmajShow';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Console\majEffect_majConsole_debug((($dictShow_0)->{'show'})($a_1));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Console_debugShow'] = __NAMESPACE__ . '\\majEffect_majConsole_debugmajShow';

