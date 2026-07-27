<?php

namespace Effect\Exception;

// ALL IMPORTS: Control.Applicative, Control.Semigroupoid, Data.Either, Data.Functor, Data.Maybe, Data.Show, Effect, Effect.Exception, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Semigroupoid, Data.Either, Data.Functor, Data.Maybe, Data.Show, Effect, Effect.Exception, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Show/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Exception/index.php';
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
$ffi_Effect_Exception = \call_user_func(function() {
  $exports = [];
$error = function($msg) { return new \Exception($msg); };
$errorWithCause = function($msg) {
    return function($cause) use ($msg) {
        return new \Exception($msg, 0, $cause instanceof \Throwable ? $cause : null);
    };
};
$errorWithName = function($msg) {
    return function($name) use ($msg) {
        $e = new class($msg) extends \Exception { public $name; };
        $e->name = $name;
        return $e;
    };
};
$message = function($e) { return $e->getMessage() . "\n" . $e->getTraceAsString(); };
$name = function($e) { return isset($e->name) ? $e->name : \get_class($e); };
$stackImpl = function($just) { return function($nothing) use ($just) { return function($e) use($just, $nothing) { return $just($e->getTraceAsString()); }; }; };
$throwException = function($e) { return function() use($e) { throw $e; }; };
$catchException = function($c, $t = null) use (&$catchException) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$catchException) {

            return $catchException(...\array_merge($__args, $more));
        };
    }
    return function() use($c, $t) { try { return $t(); } catch (\Throwable $e) { return $c($e)(); } };
};
$showErrorImpl = function($e) { return (string)$e; };

$exports['error'] = $error;
$exports['errorWithCause'] = $errorWithCause;
$exports['errorWithName'] = $errorWithName;
$exports['message'] = $message;
$exports['name'] = $name;
$exports['stackImpl'] = $stackImpl;
$exports['throwException'] = $throwException;
$exports['catchException'] = $catchException;
$exports['showErrorImpl'] = $showErrorImpl;
return $exports;
  return $exports;
});
function majEffect_majException_catchmajException($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_catchmajException';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['catchException'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Exception_catchException'] = __NAMESPACE__ . '\\majEffect_majException_catchmajException';

function majEffect_majException_error(string $v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_error';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['error'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Exception_error'] = __NAMESPACE__ . '\\majEffect_majException_error';

function majEffect_majException_errormajWithmajCause(string $v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_errormajWithmajCause';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['errorWithCause'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Exception_errorWithCause'] = __NAMESPACE__ . '\\majEffect_majException_errormajWithmajCause';

function majEffect_majException_errormajWithmajName(string $v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_errormajWithmajName';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['errorWithName'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Exception_errorWithName'] = __NAMESPACE__ . '\\majEffect_majException_errormajWithmajName';

function majEffect_majException_message($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_message';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['message'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Exception_message'] = __NAMESPACE__ . '\\majEffect_majException_message';

function majEffect_majException_name($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_name';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['name'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Exception_name'] = __NAMESPACE__ . '\\majEffect_majException_name';

function majEffect_majException_showmajErrormajImpl($v0): string|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_showmajErrormajImpl';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['showErrorImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Exception_showErrorImpl'] = __NAMESPACE__ . '\\majEffect_majException_showmajErrormajImpl';

function majEffect_majException_stackmajImpl($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_stackmajImpl';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['stackImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Effect_Exception_stackImpl'] = __NAMESPACE__ . '\\majEffect_majException_stackmajImpl';

function majEffect_majException_throwmajException($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majException_throwmajException';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Exception;
  $f = ($ffi_Effect_Exception['throwException'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Exception_throwException'] = __NAMESPACE__ . '\\majEffect_majException_throwmajException';





// Effect_Exception_try
function majEffect_majException_try($action_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majException_try';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Effect\Exception\majEffect_majException_catchmajException((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Effect_applicativeEffect'])->{'pure'}))($GLOBALS['Data_Either_Left']), ((($GLOBALS['Effect_functorEffect'])->{'map'})($GLOBALS['Data_Either_Right']))($action_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Exception_try'] = __NAMESPACE__ . '\\majEffect_majException_try';

// Effect_Exception_throw_closure
$GLOBALS['Effect_Exception_throw_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Exception_throwException']))($GLOBALS['Effect_Exception_error']);

// Effect_Exception_throw
function majEffect_majException_throw(string $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majException_throw';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Exception_throw_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Exception_throw'] = __NAMESPACE__ . '\\majEffect_majException_throw';

// Effect_Exception_stack_closure
$GLOBALS['Effect_Exception_stack_closure'] = (($GLOBALS['Effect_Exception_stackImpl'])($GLOBALS['Data_Maybe_Just']))(new \Data\Maybe\Data_Maybe_Nothing());

// Effect_Exception_stack
function majEffect_majException_stack($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majException_stack';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Exception_stack_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Exception_stack'] = __NAMESPACE__ . '\\majEffect_majException_stack';

// Effect_Exception_showError
$GLOBALS['Effect_Exception_showError'] = (object)["show" => $GLOBALS['Effect_Exception_showErrorImpl']];

