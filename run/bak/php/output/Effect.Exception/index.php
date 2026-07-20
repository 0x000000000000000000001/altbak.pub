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
\PhpursThunks::$thunks['Effect_Exception_try'] = function() { $v = function($action_0) {
  $__num = \func_num_args();
  // DEBUG UncurriedAbs: currentBindingName=(Just "Effect_Exception_try"), recVars=[];
  $__res = ((($GLOBALS['Effect_Exception_catchException'] ?? \PhpursThunks::eval('Effect_Exception_catchException')))(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_pureE'] ?? \PhpursThunks::eval('Effect_pureE'))))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left')))))((((($GLOBALS['Effect_applyEffect'] ?? \PhpursThunks::eval('Effect_applyEffect')))->apply)((($GLOBALS['Effect_pureE'] ?? \PhpursThunks::eval('Effect_pureE')))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right')))))($action_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Effect_Exception_throw'] = function() { $v = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Effect_Exception_throwException'] ?? \PhpursThunks::eval('Effect_Exception_throwException'))))(($GLOBALS['Effect_Exception_error'] ?? \PhpursThunks::eval('Effect_Exception_error'))); return $v; };
\PhpursThunks::$thunks['Effect_Exception_stack'] = function() { $v = ((($GLOBALS['Effect_Exception_stackImpl'] ?? \PhpursThunks::eval('Effect_Exception_stackImpl')))(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(new Phpurs_Data0("Nothing")); return $v; };
\PhpursThunks::$thunks['Effect_Exception_showError'] = function() { $v = (object)["show" => ($GLOBALS['Effect_Exception_showErrorImpl'] ?? \PhpursThunks::eval('Effect_Exception_showErrorImpl'))]; return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Effect_Exception = \call_user_func(function() {
  $exports = [];
$error = function($msg) use (&$error) { return new \Exception($msg); };
$message = function($e) use (&$message) { return $e->getMessage() . "\n" . $e->getTraceAsString(); };
$name = function($e) use (&$name) { return \get_class($e); };
$stackImpl = function($just) use (&$stackImpl) { return function($nothing) { return function($e) use(&$just, &$nothing) { return $just($e->getTraceAsString()); }; }; };
$throwException = function($e) use (&$throwException) { return function() use(&$e) { throw $e; }; };
$catchException = function($c, $t = null) use (&$catchException) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$catchException) {

            return $catchException(...\array_merge($__args, $more));
        };
    }
    return function() use(&$c, &$t) { try { return $t(); } catch (\Throwable $e) { if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { echo "\n\n!!! CAUGHT IN EXCEPTION.PHP:\n" . $e->getTraceAsString() . "\n\n"; } return $c($e)(); } };
};
$showErrorImpl = function($e) use (&$showErrorImpl) { return (string)$e; };

$exports['error'] = $error;
$exports['message'] = $message;
$exports['name'] = $name;
$exports['stackImpl'] = $stackImpl;
$exports['throwException'] = $throwException;
$exports['catchException'] = $catchException;
$exports['showErrorImpl'] = $showErrorImpl;
return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Effect_Exception_catchException'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['catchException']; };
\PhpursThunks::$thunks['Effect_Exception_error'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['error']; };
\PhpursThunks::$thunks['Effect_Exception_errorWithCause'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['errorWithCause']; };
\PhpursThunks::$thunks['Effect_Exception_errorWithName'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['errorWithName']; };
\PhpursThunks::$thunks['Effect_Exception_message'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['message']; };
\PhpursThunks::$thunks['Effect_Exception_name'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['name']; };
\PhpursThunks::$thunks['Effect_Exception_showErrorImpl'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['showErrorImpl']; };
\PhpursThunks::$thunks['Effect_Exception_stackImpl'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['stackImpl']; };
\PhpursThunks::$thunks['Effect_Exception_throwException'] = function() use (&$ffi_Effect_Exception) { return $ffi_Effect_Exception['throwException']; };






