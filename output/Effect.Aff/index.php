<?php

namespace Effect\Aff;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Error.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.ST.Global, Control.Parallel, Control.Parallel.Class, Control.Plus, Control.Semigroupoid, Data.Either, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Time.Duration, Data.Unit, Effect, Effect.Aff, Effect.Class, Effect.Exception, Effect.Unsafe, Partial.Unsafe, Prelude, Prim, Unsafe.Coerce
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Lazy, Control.Monad, Control.Monad.Error.Class, Control.Monad.Rec.Class, Control.Monad.ST.Class, Control.Monad.ST.Global, Control.Parallel, Control.Parallel.Class, Control.Plus, Control.Semigroupoid, Data.Either, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.Monoid, Data.Newtype, Data.Semigroup, Data.Time.Duration, Data.Unit, Effect, Effect.Aff, Effect.Class, Effect.Exception, Effect.Unsafe, Partial.Unsafe, Prelude, Unsafe.Coerce
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad/index.php';
require_once __DIR__ . '/../Control.Monad.Error.Class/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Global/index.php';
require_once __DIR__ . '/../Control.Parallel/index.php';
require_once __DIR__ . '/../Control.Parallel.Class/index.php';
require_once __DIR__ . '/../Control.Plus/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Monoid/index.php';
require_once __DIR__ . '/../Data.Newtype/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Time.Duration/index.php';
require_once __DIR__ . '/../Data.Unit/index.php';
require_once __DIR__ . '/../Effect/index.php';
require_once __DIR__ . '/../Effect.Aff/index.php';
require_once __DIR__ . '/../Effect.Class/index.php';
require_once __DIR__ . '/../Effect.Exception/index.php';
require_once __DIR__ . '/../Effect.Unsafe/index.php';
require_once __DIR__ . '/../Partial.Unsafe/index.php';
require_once __DIR__ . '/../Prelude/index.php';
require_once __DIR__ . '/../Unsafe.Coerce/index.php';

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
$ffi_Effect_Aff = \call_user_func(function() {
  $exports = [];
class PhpursAffBind {
    public $aff;
    public $f;
    public function __construct($aff, $f) {
        $this->aff = $aff;
        $this->f = $f;
    }
}

class PhpursAffMap {
    public $f;
    public $aff;
    public function __construct($f, $aff) {
        $this->f = $f;
        $this->aff = $aff;
    }
}

class PhpursAffCatch {
    public $aff;
    public $f;
    public function __construct($aff, $f) {
        $this->aff = $aff;
        $this->f = $f;
    }
}

class PhpursAffBracket {
    public $acq;
    public $cond;
    public $use;
    public function __construct($acq, $cond, $use) {
        $this->acq = $acq;
        $this->cond = $cond;
        $this->use = $use;
    }
}

class PhpursAffKillException extends \Exception {
    public $error;
    public function __construct($error) {
        $this->error = $error;
        parent::__construct("Fiber killed");
    }
}

class PhpursFiberObj {
    public static $masked = [];
    public static $pendingKills = [];
    public $fiber;
    public $activeCanceler;
    public static $fiberMap = null;
    public $isDone = false;
    public $result = null;
    public $joiners = [];
    public $run;
    public $join;
    public $isSuspended;
    public $onComplete;
    public $kill;

    public function __construct($fiber) {
        $this->fiber = $fiber;
        $this->activeCanceler = null;
        if (self::$fiberMap === null) {
            self::$fiberMap = new \WeakMap();
        }
        if ($fiber) {
            self::$fiberMap[$fiber] = $this;
        }

        $this->run = function() use ($fiber) {
            if ($fiber && $fiber->isStarted() === false) {
                \Revolt\EventLoop::queue(function() use($fiber) { 
                    if ($fiber->isStarted() === false) {
                        $fiber->start();
                    }
                });
            }
        };
        
        $this->join = function($k) {
            return function() use($k) {
                if ($this->isDone) {
                    $cb = $k($this->result);
                    $cb();
                } else {
                    $this->joiners[] = $k;
                }
                ($this->run)();
                return function() {};
            };
        };
        
        $this->onComplete = function($cb) {
            return function() {};
        };
        
        $this->isSuspended = function() {
            return false;
        };
        
        $this->kill = function($error, $cb) {
            return function() use($error, $cb) {
                if ($this->isDone) {
                    $Right = $GLOBALS['Data_Either_Right'] ?? function($x) { return (object)['tag' => 'Right', 'value0' => $x]; };
                    $fn = $cb($Right(null));
                    $fn();
                    return function() {};
                }
                
                $this->activeCanceler = $error;
                
                // Throw kill exception if fiber is suspended
                if ($this->fiber && $this->fiber->isSuspended()) {
                    if (!empty(self::$masked[spl_object_id($this->fiber)])) {
                        self::$pendingKills[spl_object_id($this->fiber)] = $error;
                    } else {
                        \Revolt\EventLoop::queue(function() use($error) {
                            if ($this->fiber && $this->fiber->isSuspended()) {
                                $this->fiber->throw(new PhpursAffKillException($error));
                            }
                        });
                    }
                }
                
                $Right = $GLOBALS['Data_Either_Right'] ?? function($x) { return (object)['tag' => 'Right', 'value0' => $x]; };
                $fn = $cb($Right(null));
                $fn();
                return function() {};
            };
        };
    }
    
    public function finish($either) {
        $this->isDone = true;
        $this->result = $either;
        foreach ($this->joiners as $k) {
            \Revolt\EventLoop::queue(function() use($k, $either) { 
                $cb = $k($either);
                $cb(); 
            });
        }
        $this->joiners = [];
    }

    public function kill($err, $k, $Right, $unit) {
        $cancelFiber = null;
        if ($this->fiber) {
            if ($this->fiber->isSuspended() || $this->fiber->isStarted() === false) {
                if ($this->activeCanceler !== null) {
                    $affCanceler = ($this->activeCanceler)($err);
                    
                    $cancelFiber = new \Fiber(function() use($affCanceler, $err, $k, $Right, $unit) {
                        try {
                            phpursRunAffTrampoline($affCanceler);
                        } catch (\Throwable $e) {}
                        
                        if ($this->fiber && $this->fiber->isSuspended()) {
                            $this->fiber->throw(new PhpursAffKillException($err));
                        }
                        return $k($Right($unit))();
                    });
                } else {
                    $cancelFiber = new \Fiber(function() use($err, $k, $Right, $unit) {
                        if ($this->fiber && $this->fiber->isSuspended()) {
                            $this->fiber->throw(new PhpursAffKillException($err));
                        }
                        return $k($Right($unit))();
                    });
                }
            } else {
                return $k($Right($unit))();
            }
        }
        
        if ($cancelFiber) {
            $cancelFiber->start();
            // Since this runs in a separate fiber, we just let it execute.
        } else {
            $k($Right($unit))();
        }
    }
}

if (!\function_exists('phpursRunAffTrampoline')) {
function phpursRunAffTrampoline($aff) {
    $current = $aff;
    $stack = []; 

    while (true) {
        try {
            if ($current instanceof \Closure) {
                $res = $current();
            } else {
                $res = $current;
            }
            
            if ($res instanceof PhpursAffBind) {
                $stack[] = ['type' => 'bind', 'f' => $res->f];
                $current = $res->aff;
                continue;
            } elseif ($res instanceof PhpursAffMap) {
                $stack[] = ['type' => 'map', 'f' => $res->f];
                $current = $res->aff;
                continue;
            } elseif ($res instanceof PhpursAffCatch) {
                $stack[] = ['type' => 'catch', 'f' => $res->f];
                $current = $res->aff;
                continue;
            } elseif ($res instanceof PhpursAffBracket) {
                $acq = $res->acq;
                $cond = $res->cond;
                $use = $res->use;
                

                $fiberId = spl_object_id(\Fiber::getCurrent());
                if (!isset(PhpursFiberObj::$masked[$fiberId])) PhpursFiberObj::$masked[$fiberId] = 0;
                PhpursFiberObj::$masked[$fiberId]++;
                try {
                    $resource = phpursRunAffTrampoline($acq);
                } catch (\Throwable $e) {
                    PhpursFiberObj::$masked[$fiberId]--;
                    throw $e;
                }
                PhpursFiberObj::$masked[$fiberId]--;
                
                try {
                    if (isset(PhpursFiberObj::$pendingKills[$fiberId])) {
                        $killErr = PhpursFiberObj::$pendingKills[$fiberId];
                        unset(PhpursFiberObj::$pendingKills[$fiberId]);
                        throw new PhpursAffKillException($killErr);
                    }
                    $useResult = phpursRunAffTrampoline($use($resource));
                    
                    // Completed!
                    PhpursFiberObj::$masked[$fiberId]++;
                    try {
                        phpursRunAffTrampoline(($cond->completed)($useResult)($resource));
                    } catch (\Throwable $e) {
                        PhpursFiberObj::$masked[$fiberId]--;
                        throw $e;
                    }
                    PhpursFiberObj::$masked[$fiberId]--;
                    
                    $res = $useResult;
                } catch (\Throwable $err) {
                    if ($err instanceof PhpursAffKillException) {
                        // Killed!
                        PhpursFiberObj::$masked[$fiberId]++;
                        try {
                            phpursRunAffTrampoline(($cond->killed)($err->error)($resource));
                        } catch (\Throwable $e) {
                            PhpursFiberObj::$masked[$fiberId]--;
                            throw $e;
                        }
                        PhpursFiberObj::$masked[$fiberId]--;
                        throw $err;
                    } else {
                        // Failed!
                        PhpursFiberObj::$masked[$fiberId]++;
                        try {
                            phpursRunAffTrampoline(($cond->failed)($err)($resource));
                        } catch (\Throwable $e) {
                            PhpursFiberObj::$masked[$fiberId]--;
                            throw $e;
                        }
                        PhpursFiberObj::$masked[$fiberId]--;
                        throw $err;
                    }
                }
            }
            
            while (true) {
                if (empty($stack)) {
                    return $res;
                }
                
                $frame = array_pop($stack);
                
                if ($frame['type'] === 'bind') {
                    $f = $frame['f'];
                    $current = $f($res);
                    break;
                } elseif ($frame['type'] === 'map') {
                    $f = $frame['f'];
                    $res = $f($res);
                } elseif ($frame['type'] === 'catch') {
                    // Success value passed through
                }
            }
        } catch (\Throwable $e) {
            
            if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { 
                echo "\n\n!!! GLOBAL FATAL ERROR CAUGHT IN AFF:\n" . $e->getTraceAsString() . "\n\n"; 
                \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); 
            }
            
            if ($e instanceof PhpursAffKillException) {
                throw $e;
            }
            
            $caught = false;
            while (!empty($stack)) {
                $frame = array_pop($stack);
                if ($frame['type'] === 'catch') {
                    $f = $frame['f'];
                    $current = $f($e);
                    $caught = true;
                    break;
                }
            }
            if (!$caught) {
                throw $e;
            }
        }
    }
}
}

$_pure = function($x) use (&$_pure) { return function() use($x) { return $x; }; };
$_map = function($f, $aff) use (&$_map) {
    return function() use($f, $aff) { return new PhpursAffMap($f, $aff); };
};
$_bind = function($aff, $f) use (&$_bind) {
    return function() use($aff, $f) { return new PhpursAffBind($aff, $f); };
};
$_liftEffect = function($eff) use (&$_liftEffect) { return $eff; };
$_makeFiber = function($isLeft, $unsafeFromLeft, $unsafeFromRight, $Left, $Right, $aff) use (&$_makeFiber) {
    return function() use($aff, $Left, $Right) { 
        $fiber = new \Fiber(function() use ($aff, &$obj, $Left, $Right) { 
            try {
                $res = phpursRunAffTrampoline($aff);
                $obj->finish($Right($res));
            } catch (\Throwable $e) {
                if ($e instanceof PhpursAffKillException) {
                    // echo "\nFATAL ERROR IN FIBER: " . $e->getMessage() . "\n";
                    if ($e instanceof PhpursAffKillException) { $obj->finish($Left($e->error)); } else { $obj->finish($Left($e)); }
                } else {
                    // echo "\nFATAL ERROR IN FIBER: " . $e->getMessage() . "\n";
                    if ($e instanceof PhpursAffKillException) { $obj->finish($Left($e->error)); } else { $obj->finish($Left($e)); }
                }
            }
        }); 
        $obj = new PhpursFiberObj($fiber);
        $fiber->start(); 
        return $obj; 
    }; 
};
$_fork = function($immediate, $aff) use (&$_fork) {
    return function() use($aff, $immediate) { 
        $Left = $GLOBALS['Data_Either_Left'];
        $Right = $GLOBALS['Data_Either_Right'];
        $fiber = new \Fiber(function() use ($aff, &$obj, $Left, $Right) { 
            try {
                $res = phpursRunAffTrampoline($aff);
                $obj->finish($Right($res));
            } catch (\Throwable $e) {
                if ($e instanceof PhpursAffKillException) {
                    // echo "\nFATAL ERROR IN FIBER: " . $e->getMessage() . "\n";
                    if ($e instanceof PhpursAffKillException) { $obj->finish($Left($e->error)); } else { $obj->finish($Left($e)); }
                } else {
                    // echo "\nFATAL ERROR IN FIBER: " . $e->getMessage() . "\n";
                    if ($e instanceof PhpursAffKillException) { $obj->finish($Left($e->error)); } else { $obj->finish($Left($e)); }
                }
            }
        }); 
        $obj = new PhpursFiberObj($fiber);
        if ($immediate) {
            ($obj->run)();
        }
        return $obj; 
    };
};
$_delay = function($right, $ms) use (&$_delay) { 
    return function() use($right, $ms) { 
        $fiber = \Fiber::getCurrent(); 
        if ($ms <= 0.0) {
            static $ticks = 0;
            static $lastYield = 0;
            $ticks++;
            
            $shouldYield = false;
            if ($ticks >= 50) {
                $shouldYield = true;
            } elseif ($ticks % 10 === 0) {
                $now = \hrtime(true);
                if ($now - $lastYield > 5000000) { // 5ms in nanoseconds
                    $shouldYield = true;
                }
            }

            if ($shouldYield) {
                $ticks = 0;
                $lastYield = \hrtime(true);
                \Revolt\EventLoop::queue(function() use($fiber) { 
                    if ($fiber && $fiber->isSuspended()) $fiber->resume(); 
                }); 
                if ($fiber) \Fiber::suspend(); 
            }
        } else {
            \Revolt\EventLoop::delay($ms / 1000, function() use($fiber) { 
                if ($fiber && $fiber->isSuspended()) $fiber->resume(); 
            }); 
            if ($fiber) \Fiber::suspend(); 
        }
        return null; 
    }; 
};
$_makeSupervisedFiber = function($isLeft, $unsafeFromLeft, $unsafeFromRight, $Left, $Right, $aff) use (&$_makeFiber) {
    return function() use($isLeft, $unsafeFromLeft, $unsafeFromRight, $Left, $Right, $aff, &$_makeFiber) {
        $supervisor = $_makeFiber($isLeft, $unsafeFromLeft, $unsafeFromRight, $Left, $Right, $aff)();
        return (object)[
            "fiber" => $supervisor,
            "supervisor" => $supervisor
        ];
    };
};
$_killAll = function($err, $sup, $cb) use (&$_killAll) { return function() { return function(){}; }; };

$_makeAff = function($isLeft, $unsafeFromLeft, $unsafeFromRight, $Left, $Right, $k) use (&$_makeAff) {
    return function() use($k) { 
        $fiber = \Fiber::getCurrent(); 
        $isDone = false;
        $result;
        $exception;

        $canceler = $k(function($res) use($fiber, &$isDone, &$result, &$exception) { 
            return function() use($fiber, &$isDone, &$result, &$exception, $res) { 
                $isDone = true;
                if (is_object($res) && $res->tag === "Left") {
                    $exception = $res->value0;
                } else {
                    $result = $res->value0;
                }
                
                if ($fiber && $fiber->isSuspended()) { 
                    if ($exception !== null) {
                        \Revolt\EventLoop::queue(function() use($fiber, $exception) {
                            if ($fiber->isSuspended()) $fiber->throw($exception); 
                        });
                    } else {
                        \Revolt\EventLoop::queue(function() use($fiber, $result) {
                            if ($fiber->isSuspended()) $fiber->resume($result); 
                        });
                    }
                } 
            }; 
        })(); 
        
        if (!$isDone) {
            if ($fiber) {
                return \Fiber::suspend(); 
            } else {
                throw new \RuntimeException("makeAff used outside of a fiber");
            }
        } else {
            if ($exception !== null) throw $exception;
            return $result;
        }
    }; 
};

$_throwError = function($err) use (&$_throwError) { return function() use($err) { throw $err; }; };
$_catchError = function($aff, $f) use (&$_catchError) {
    return function() use($aff, $f) { return new PhpursAffCatch($aff, $f); };
};
$generalBracket = function($acq, $cond, $use) use (&$generalBracket) {
    return function() use($acq, $cond, $use) { return new PhpursAffBracket($acq, $cond, $use); }; 
};
$_parAffMap = $_map;

$_parAffApply = function($aff1, $aff2) use (&$_parAffApply) {
    return function() use($aff1, $aff2) { 
        $parent = \Fiber::getCurrent();
        $isDone = false; 
        $completed = 0;
        $res1;
        $res2;
        $error;

        $f1 = new \Fiber(function() use($aff1, &$isDone, &$completed, &$res1, &$error, $parent) {
            try {
                $res1 = phpursRunAffTrampoline($aff1);
                if (!$isDone) {
                    $completed++;
                    if ($completed === 2) {
                        $isDone = true;
                        if ($parent && $parent->isSuspended()) {
                            \Revolt\EventLoop::queue(function() use($parent) {
                                if ($parent->isSuspended()) $parent->resume();
                            });
                        }
                    }
                }
            } catch (\Throwable $e) { 
             
            if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                if (!$isDone) {
                    $isDone = true;
                    $error = $e;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $e) {
                            if ($parent->isSuspended()) $parent->throw($e);
                        });
                    }
                }
            }
        });

        $f2 = new \Fiber(function() use($aff2, &$isDone, &$completed, &$res2, &$error, $parent) {
            try {
                $res2 = phpursRunAffTrampoline($aff2);
                if (!$isDone) {
                    $completed++;
                    if ($completed === 2) {
                        $isDone = true;
                        if ($parent && $parent->isSuspended()) {
                            \Revolt\EventLoop::queue(function() use($parent) {
                                if ($parent->isSuspended()) $parent->resume();
                            });
                        }
                    }
                }
            } catch (\Throwable $e) { 
             
            if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                if (!$isDone) {
                    $isDone = true;
                    $error = $e;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $e) {
                            if ($parent->isSuspended()) $parent->throw($e);
                        });
                    }
                }
            }
        });

        \Revolt\EventLoop::queue(function() use($f1) { $f1->start(); });
        \Revolt\EventLoop::queue(function() use($f2) { $f2->start(); });

        if (!$isDone) {
            \Fiber::suspend();
        }
        
        if ($error !== null) throw $error;
        return $res1($res2); 
    };
};

$_sequential = function($aff) use (&$_sequential) { return $aff; };

$_parAffAlt = function($aff1, $aff2) use (&$_parAffAlt) {
    return function() use($aff1, $aff2) { 
        $parent = \Fiber::getCurrent();
        $isDone = false;
        $result;
        $doneCount = 0;
        $error2;

        $f1 = new \Fiber(function() use($aff1, &$isDone, &$result, &$doneCount, &$error2, $parent) {
            try {
                $res = phpursRunAffTrampoline($aff1);
                if (!$isDone) {
                    $isDone = true;
                    $result = $res;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $result) {
                            if ($parent->isSuspended()) $parent->resume($result);
                        });
                    }
                }
            } catch (\Throwable $e) { 
             
            if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                $doneCount++;
                if ($doneCount === 2 && !$isDone) {
                    $isDone = true;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $error2) {
                            if ($parent->isSuspended()) $parent->throw($error2); 
                        });
                    }
                }
            }
        });

        $f2 = new \Fiber(function() use($aff2, &$isDone, &$result, &$doneCount, &$error2, $parent) {
            try {
                $res = phpursRunAffTrampoline($aff2);
                if (!$isDone) {
                    $isDone = true;
                    $result = $res;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $result) {
                            if ($parent->isSuspended()) $parent->resume($result);
                        });
                    }
                }
            } catch (\Throwable $e) { 
             
            if (strpos($e->getMessage(), 'Object of class stdClass') !== false) { \file_put_contents('/tmp/aff_caught.log', 'CAUGHT: ' . \get_class($e) . ' ' . $e->getMessage() . "\n" . $e->getTraceAsString() . "\n\n", FILE_APPEND); }
                $error2 = $e;
                $doneCount++;
                if ($doneCount === 2 && !$isDone) {
                    $isDone = true;
                    if ($parent && $parent->isSuspended()) {
                        \Revolt\EventLoop::queue(function() use($parent, $error2) {
                            if ($parent->isSuspended()) $parent->throw($error2);
                        });
                    }
                }
            }
        });

        \Revolt\EventLoop::queue(function() use($f1) { $f1->start(); });
        \Revolt\EventLoop::queue(function() use($f2) { $f2->start(); });

        if (!$isDone) {
            return \Fiber::suspend();
        } else {
            if ($doneCount === 2) throw $error2;
            return $result;
        }
    };
};

$exports['_pure'] = $_pure;
$exports['_map'] = $_map;
$exports['_bind'] = $_bind;
$exports['_liftEffect'] = $_liftEffect;
$exports['_makeFiber'] = $_makeFiber;
$exports['_fork'] = $_fork;
$exports['_delay'] = $_delay;
$exports['_makeSupervisedFiber'] = $_makeSupervisedFiber;
$exports['_killAll'] = $_killAll;
$exports['_makeAff'] = $_makeAff;
$exports['_throwError'] = $_throwError;
$exports['_catchError'] = $_catchError;
$exports['generalBracket'] = $generalBracket;
$exports['_parAffMap'] = $_parAffMap;
$exports['_parAffApply'] = $_parAffApply;
$exports['_sequential'] = $_sequential;
$exports['_parAffAlt'] = $_parAffAlt;
return $exports;
  return $exports;
});
function majEffect_majAff__bind($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__bind';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_bind', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_bind'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Aff__bind'] = __NAMESPACE__ . '\\majEffect_majAff__bind';

function majEffect_majAff__catchmajError($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__catchmajError';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_catchError', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_catchError'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Aff__catchError'] = __NAMESPACE__ . '\\majEffect_majAff__catchmajError';

$GLOBALS['Effect_Aff__delay'] = (\array_key_exists('_delay', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_delay'] : new class { public function __invoke(...$args) { return $this; } });
function majEffect_majAff__fork(bool $v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__fork';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_fork', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_fork'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Aff__fork'] = __NAMESPACE__ . '\\majEffect_majAff__fork';

$GLOBALS['Effect_Aff__killAll'] = (\array_key_exists('_killAll', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_killAll'] : new class { public function __invoke(...$args) { return $this; } });
function majEffect_majAff__liftmajEffect($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__liftmajEffect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_liftEffect', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_liftEffect'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Aff__liftEffect'] = __NAMESPACE__ . '\\majEffect_majAff__liftmajEffect';

$GLOBALS['Effect_Aff__makeAff'] = (\array_key_exists('_makeAff', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_makeAff'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Effect_Aff__makeFiber'] = (\array_key_exists('_makeFiber', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_makeFiber'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Effect_Aff__makeSupervisedFiber'] = (\array_key_exists('_makeSupervisedFiber', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_makeSupervisedFiber'] : new class { public function __invoke(...$args) { return $this; } });
function majEffect_majAff__map($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__map';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_map', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_map'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Aff__map'] = __NAMESPACE__ . '\\majEffect_majAff__map';

function majEffect_majAff__parmajAffmajAlt($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__parmajAffmajAlt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_parAffAlt', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_parAffAlt'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Aff__parAffAlt'] = __NAMESPACE__ . '\\majEffect_majAff__parmajAffmajAlt';

function majEffect_majAff__parmajAffmajApply($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__parmajAffmajApply';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_parAffApply', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_parAffApply'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Aff__parAffApply'] = __NAMESPACE__ . '\\majEffect_majAff__parmajAffmajApply';

function majEffect_majAff__parmajAffmajMap($v0, $v1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__parmajAffmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_parAffMap', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_parAffMap'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1);
}
$GLOBALS['Effect_Aff__parAffMap'] = __NAMESPACE__ . '\\majEffect_majAff__parmajAffmajMap';

function majEffect_majAff__pure($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__pure';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_pure', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_pure'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Aff__pure'] = __NAMESPACE__ . '\\majEffect_majAff__pure';

$GLOBALS['Effect_Aff__sequential'] = (\array_key_exists('_sequential', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_sequential'] : new class { public function __invoke(...$args) { return $this; } });
function majEffect_majAff__throwmajError($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff__throwmajError';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('_throwError', $ffi_Effect_Aff) ? $ffi_Effect_Aff['_throwError'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Effect_Aff__throwError'] = __NAMESPACE__ . '\\majEffect_majAff__throwmajError';

function majEffect_majAff_generalmajBracket($v0, $v1 = null, $v2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majEffect_majAff_generalmajBracket';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  global $ffi_Effect_Aff;
  $f = (\array_key_exists('generalBracket', $ffi_Effect_Aff) ? $ffi_Effect_Aff['generalBracket'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0, $v1, $v2);
}
$GLOBALS['Effect_Aff_generalBracket'] = __NAMESPACE__ . '\\majEffect_majAff_generalmajBracket';





// Effect_Aff_Canceler
function majEffect_majAff_majCanceler($x_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_majCanceler';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = $x_0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_Canceler'] = __NAMESPACE__ . '\\majEffect_majAff_majCanceler';

// Effect_Aff_unsafeFromRight
function majEffect_majAff_unsafemajFrommajRight($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_unsafemajFrommajRight';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = \Partial\majPartial__crashmajWith("unsafeFromRight: Left");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_unsafeFromRight'] = __NAMESPACE__ . '\\majEffect_majAff_unsafemajFrommajRight';

// Effect_Aff_unsafeFromLeft
function majEffect_majAff_unsafemajFrommajLeft($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_unsafemajFrommajLeft';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = \Partial\majPartial__crashmajWith("unsafeFromLeft: Right");
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_unsafeFromLeft'] = __NAMESPACE__ . '\\majEffect_majAff_unsafemajFrommajLeft';

// Effect_Aff_suspendAff_closure
$GLOBALS['Effect_Aff_suspendAff_closure'] = ($GLOBALS['Effect_Aff__fork'])(false);

// Effect_Aff_suspendAff
function majEffect_majAff_suspendmajAff($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_suspendmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff_suspendAff_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_suspendAff'] = __NAMESPACE__ . '\\majEffect_majAff_suspendmajAff';

// Effect_Aff_newtypeCanceler
$GLOBALS['Effect_Aff_newtypeCanceler'] = (object)["Coercible0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = null;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_isLeft
function majEffect_majAff_ismajLeft($v_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_ismajLeft';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Either\Data_Either_Left) {
$__t0 = true;
goto end_branch_0;;
};
  if ($v_0 instanceof \Data\Either\Data_Either_Right) {
$__t0 = false;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_isLeft'] = __NAMESPACE__ . '\\majEffect_majAff_ismajLeft';

// Effect_Aff_makeAff
function majEffect_majAff_makemajAff($k_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_makemajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff__makeAff'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], $k_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_makeAff'] = __NAMESPACE__ . '\\majEffect_majAff_makemajAff';

// Effect_Aff_makeFiber
function majEffect_majAff_makemajFiber($aff_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_makemajFiber';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff__makeFiber'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], $aff_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_makeFiber'] = __NAMESPACE__ . '\\majEffect_majAff_makemajFiber';

// Effect_Aff_launchAff
function majEffect_majAff_launchmajAff($aff_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_launchmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($GLOBALS['Effect_Aff__makeFiber'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], $aff_0);
  $__res = function() use ($__local_var_1_0, &$__fn) {
$fiber_2_1 = phpurs_execute_effect($__local_var_1_0);
$_dollar___unused_3_2 = phpurs_execute_effect(($fiber_2_1)->{'run'});
return phpurs_execute_effect(phpurs_execute_effect($fiber_2_1));
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_launchAff'] = __NAMESPACE__ . '\\majEffect_majAff_launchmajAff';

// Effect_Aff_launchAff__closure
$GLOBALS['Effect_Aff_launchAff__closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_0) {
  $__num = \func_num_args();
  $__res = function() use ($a_0, &$__fn) {
$a_prime__1_0 = phpurs_execute_effect($a_0);
return phpurs_execute_effect($GLOBALS['Data_Unit_unit']);
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Effect_Aff_launchAff']);

// Effect_Aff_launchAff_
function majEffect_majAff_launchmajAff_($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_launchmajAff_';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff_launchAff__closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_launchAff_'] = __NAMESPACE__ . '\\majEffect_majAff_launchmajAff_';

// Effect_Aff_launchSuspendedAff_closure
$GLOBALS['Effect_Aff_launchSuspendedAff_closure'] = $GLOBALS['Effect_Aff_makeFiber'];

// Effect_Aff_launchSuspendedAff
function majEffect_majAff_launchmajSuspendedmajAff($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_launchmajSuspendedmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff_launchSuspendedAff_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_launchSuspendedAff'] = __NAMESPACE__ . '\\majEffect_majAff_launchmajSuspendedmajAff';

// Effect_Aff_functorParAff
$GLOBALS['Effect_Aff_functorParAff'] = (object)["map" => $GLOBALS['Effect_Aff__parAffMap']];

// Effect_Aff_functorAff
$GLOBALS['Effect_Aff_functorAff'] = (object)["map" => $GLOBALS['Effect_Aff__map']];

// Effect_Aff_forkAff_closure
$GLOBALS['Effect_Aff_forkAff_closure'] = ($GLOBALS['Effect_Aff__fork'])(true);

// Effect_Aff_forkAff
function majEffect_majAff_forkmajAff($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_forkmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff_forkAff_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_forkAff'] = __NAMESPACE__ . '\\majEffect_majAff_forkmajAff';

// Effect_Aff_delay
function majEffect_majAff_delay(float $v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_delay';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff__delay'])($GLOBALS['Data_Either_Right'], $v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_delay'] = __NAMESPACE__ . '\\majEffect_majAff_delay';

// Effect_Aff_bracket
function majEffect_majAff_bracket($acquire_0, $completed_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_bracket';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Effect_Aff_generalBracket'])($acquire_0))((object)["killed" => function($v_2) use ($completed_1) {
  $__num = \func_num_args();
  $__res = $completed_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "failed" => function($v_2) use ($completed_1) {
  $__num = \func_num_args();
  $__res = $completed_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "completed" => function($v_2) use ($completed_1) {
  $__num = \func_num_args();
  $__res = $completed_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_bracket'] = __NAMESPACE__ . '\\majEffect_majAff_bracket';

// Effect_Aff_applyParAff
$GLOBALS['Effect_Aff_applyParAff'] = (object)["apply" => $GLOBALS['Effect_Aff__parAffApply'], "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_functorParAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_semigroupParAff
function majEffect_majAff_semigroupmajParmajAff($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_semigroupmajParmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($dictSemigroup_0)->{'append'};
  $__res = (object)["append" => function($a_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($__local_var_1_0, $a_2) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__parmajAffmajApply(\Effect\Aff\majEffect_majAff__parmajAffmajMap($__local_var_1_0, $a_2), $b_3);
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
$GLOBALS['Effect_Aff_semigroupParAff'] = __NAMESPACE__ . '\\majEffect_majAff_semigroupmajParmajAff';

// Effect_Aff_monadAff
$GLOBALS['Effect_Aff_monadAff'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applicativeAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Bind1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_bindAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_bindAff
$GLOBALS['Effect_Aff_bindAff'] = (object)["bind" => $GLOBALS['Effect_Aff__bind'], "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applyAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_applyAff
$GLOBALS['Effect_Aff_applyAff'] = (object)["apply" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind($f_0, function($f_prime__2) use ($a_1) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind($a_1, function($a_prime__3) use ($f_prime__2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Effect_Aff_applicativeAff'])->{'pure'})(($f_prime__2)($a_prime__3));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_functorAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_applicativeAff
$GLOBALS['Effect_Aff_applicativeAff'] = (object)["pure" => $GLOBALS['Effect_Aff__pure'], "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applyAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_cancelWith
function majEffect_majAff_cancelmajWith($aff_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_cancelmajWith';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Aff\majEffect_majAff_generalmajBracket(\Effect\Aff\majEffect_majAff__pure($GLOBALS['Data_Unit_unit']), (object)["killed" => function($e_2) use ($v_1) {
  $__num = \func_num_args();
  $__res = function($v1_3) use ($e_2, $v_1) {
  $__num = \func_num_args();
  $__res = ($v_1)($e_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "failed" => function($v_2) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff__pure'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "completed" => function($v_2) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff__pure'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}], function($v_2) use ($aff_0) {
  $__num = \func_num_args();
  $__res = $aff_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_cancelWith'] = __NAMESPACE__ . '\\majEffect_majAff_cancelmajWith';

// Effect_Aff_finally
function majEffect_majAff_finally($fin_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_finally';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Aff\majEffect_majAff_generalmajBracket(\Effect\Aff\majEffect_majAff__pure($GLOBALS['Data_Unit_unit']), (object)["killed" => function($v_2) use ($fin_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($fin_0) {
  $__num = \func_num_args();
  $__res = $fin_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "failed" => function($v_2) use ($fin_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($fin_0) {
  $__num = \func_num_args();
  $__res = $fin_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "completed" => function($v_2) use ($fin_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($fin_0) {
  $__num = \func_num_args();
  $__res = $fin_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}], function($v_2) use ($a_1) {
  $__num = \func_num_args();
  $__res = $a_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_finally'] = __NAMESPACE__ . '\\majEffect_majAff_finally';

// Effect_Aff_invincible
function majEffect_majAff_invincible($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_invincible';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = \Effect\Aff\majEffect_majAff__pure($GLOBALS['Data_Unit_unit']);
  $__res = \Effect\Aff\majEffect_majAff_generalmajBracket($a_0, (object)["killed" => function($v_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "failed" => function($v_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "completed" => function($v_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($v_3) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}], $GLOBALS['Effect_Aff__pure']);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_invincible'] = __NAMESPACE__ . '\\majEffect_majAff_invincible';

// Effect_Aff_lazyAff
$GLOBALS['Effect_Aff_lazyAff'] = (object)["defer" => function($f_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__pure($GLOBALS['Data_Unit_unit']), $f_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_parallelAff
$GLOBALS['Effect_Aff_parallelAff'] = (object)["parallel" => $GLOBALS['Unsafe_Coerce_unsafeCoerce'], "sequential" => $GLOBALS['Effect_Aff__sequential'], "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applyAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applyParAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_applicativeParAff
$GLOBALS['Effect_Aff_applicativeParAff'] = (object)["pure" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))($GLOBALS['Effect_Aff__pure']), "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applyParAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_monoidParAff
function majEffect_majAff_monoidmajParmajAff($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_monoidmajParmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'};
  $semigroupParAff1_1_0 = (object)["append" => function($a_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($__local_var_1_0, $a_2) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__parmajAffmajApply(\Effect\Aff\majEffect_majAff__parmajAffmajMap($__local_var_1_0, $a_2), $b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Unsafe_Coerce_unsafeCoerce'], $GLOBALS['Effect_Aff__pure'], ($dictMonoid_0)->{'mempty'}), "Semigroup0" => function($_dollar___unused_2) use ($semigroupParAff1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupParAff1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_monoidParAff'] = __NAMESPACE__ . '\\majEffect_majAff_monoidmajParmajAff';

// Effect_Aff_semigroupCanceler
$GLOBALS['Effect_Aff_semigroupCanceler'] = (object)["append" => function($v_0) {
  $__num = \func_num_args();
  $__res = function($v1_1) use ($v_0) {
  $__num = \func_num_args();
  $__res = function($err_2) use ($v1_1, $v_0) {
  $__num = \func_num_args();
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Effect_Aff__sequential'], (($GLOBALS['Data_Foldable_foldrArray'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_3) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($a_3) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__parmajAffmajApply(\Effect\Aff\majEffect_majAff__parmajAffmajMap(function($v_5) {
  $__num = \func_num_args();
  $__res = function($x_6) {
  $__num = \func_num_args();
  $__res = $x_6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $a_3), $b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(function($x_3) {
  $__num = \func_num_args();
  $__res = $x_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Unsafe_Coerce_unsafeCoerce'], $GLOBALS['Effect_Aff__pure'], $GLOBALS['Data_Unit_unit'])), [($v_0)($err_2), ($v1_1)($err_2)]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_semigroupAff
function majEffect_majAff_semigroupmajAff($dictSemigroup_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_semigroupmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($dictSemigroup_0)->{'append'};
  $__res = (object)["append" => function($a_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($__local_var_1_0, $a_2) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__map($__local_var_1_0, $a_2), function($f_prime__4) use ($b_3) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind($b_3, function($a_prime__5) use ($f_prime__4) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__pure(($f_prime__4)($a_prime__5));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_semigroupAff'] = __NAMESPACE__ . '\\majEffect_majAff_semigroupmajAff';

// Effect_Aff_monadEffectAff
$GLOBALS['Effect_Aff_monadEffectAff'] = (object)["liftEffect" => $GLOBALS['Effect_Aff__liftEffect'], "Monad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_monadAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_effectCanceler_closure
$GLOBALS['Effect_Aff_effectCanceler_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_0) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Function_const']))($GLOBALS['Effect_Aff__liftEffect']));

// Effect_Aff_effectCanceler
function majEffect_majAff_effectmajCanceler($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_effectmajCanceler';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Aff_effectCanceler_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_effectCanceler'] = __NAMESPACE__ . '\\majEffect_majAff_effectmajCanceler';

// Effect_Aff_joinFiber
function majEffect_majAff_joinmajFiber($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_joinmajFiber';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff__makeAff'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], function($k_1) use ($v_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = (($v_0)->{'join'})($k_1);
  $__res = function() use ($__local_var_2_0, &$__fn) {
$a_prime__3_1 = phpurs_execute_effect($__local_var_2_0);
return phpurs_execute_effect(($GLOBALS['Effect_Aff_effectCanceler'])($a_prime__3_1));
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_joinFiber'] = __NAMESPACE__ . '\\majEffect_majAff_joinmajFiber';

// Effect_Aff_functorFiber
$GLOBALS['Effect_Aff_functorFiber'] = (object)["map" => function($f_0) {
  $__num = \func_num_args();
  $__res = function($t_1) use ($f_0) {
  $__num = \func_num_args();
  $__res = \Effect\Unsafe\majEffect_majUnsafe_unsafemajPerformmajEffect(($GLOBALS['Effect_Aff__makeFiber'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], \Effect\Aff\majEffect_majAff__map($f_0, \Effect\Aff\majEffect_majAff_joinmajFiber($t_1))));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_applyFiber
$GLOBALS['Effect_Aff_applyFiber'] = (object)["apply" => function($t1_0) {
  $__num = \func_num_args();
  $__res = function($t2_1) use ($t1_0) {
  $__num = \func_num_args();
  $__local_var_2_0 = \Effect\Aff\majEffect_majAff_joinmajFiber($t2_1);
  $__res = \Effect\Unsafe\majEffect_majUnsafe_unsafemajPerformmajEffect(($GLOBALS['Effect_Aff__makeFiber'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff_joinmajFiber($t1_0), function($f_prime__3) use ($__local_var_2_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind($__local_var_2_0, function($a_prime__4) use ($f_prime__3) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__pure(($f_prime__3)($a_prime__4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_functorFiber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_applicativeFiber
$GLOBALS['Effect_Aff_applicativeFiber'] = (object)["pure" => function($a_0) {
  $__num = \func_num_args();
  $__res = \Effect\Unsafe\majEffect_majUnsafe_unsafemajPerformmajEffect(($GLOBALS['Effect_Aff__makeFiber'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], \Effect\Aff\majEffect_majAff__pure($a_0)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Apply0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applyFiber'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_killFiber
function majEffect_majAff_killmajFiber($e_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_killmajFiber';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__liftmajEffect(($v_1)->{'isSuspended'}), function($suspended_2) use ($e_0, $v_1) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ($suspended_2) {
$__local_var_3_3 = (($v_1)->{'kill'})($e_0, function($v_3) {
  $__num = \func_num_args();
  $__res = function() use (&$__fn) {
return $GLOBALS['Data_Unit_unit'];
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
$__t2 = \Effect\Aff\majEffect_majAff__liftmajEffect(function() use ($__local_var_3_3, &$__fn) {
$a_prime__4_4 = phpurs_execute_effect($__local_var_3_3);
return phpurs_execute_effect($GLOBALS['Data_Unit_unit']);
});
goto end_branch_2;;
};
  $__t2 = ($GLOBALS['Effect_Aff__makeAff'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], function($k_3) use ($e_0, $v_1) {
  $__num = \func_num_args();
  $__local_var_4_0 = (($v_1)->{'kill'})($e_0, $k_3);
  $__res = function() use ($__local_var_4_0, &$__fn) {
$a_prime__5_1 = phpurs_execute_effect($__local_var_4_0);
return phpurs_execute_effect(($GLOBALS['Effect_Aff_effectCanceler'])($a_prime__5_1));
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_killFiber'] = __NAMESPACE__ . '\\majEffect_majAff_killmajFiber';

// Effect_Aff_fiberCanceler_closure
$GLOBALS['Effect_Aff_fiberCanceler_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(function($x_0) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($b_0) {
  $__num = \func_num_args();
  $__res = function($a_1) use ($b_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff_killmajFiber($a_1, $b_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Effect_Aff_fiberCanceler
function majEffect_majAff_fibermajCanceler($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_fibermajCanceler';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Aff_fiberCanceler_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_fiberCanceler'] = __NAMESPACE__ . '\\majEffect_majAff_fibermajCanceler';

// Effect_Aff_supervise
function majEffect_majAff_supervise($aff_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_supervise';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $killError_1_0 = \Effect\Exception\majEffect_majException_error("[Aff] Child fiber outlived parent");
  $__local_var_2_1 = ($GLOBALS['Effect_Aff__makeSupervisedFiber'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], $aff_0);
  $__res = \Effect\Aff\majEffect_majAff_generalmajBracket(\Effect\Aff\majEffect_majAff__liftmajEffect(function() use ($__local_var_2_1, &$__fn) {
$sup_3_2 = phpurs_execute_effect($__local_var_2_1);
$_dollar___unused_4_3 = phpurs_execute_effect((($sup_3_2)->{'fiber'})->{'run'});
return phpurs_execute_effect(phpurs_execute_effect($sup_3_2));
}), (object)["killed" => function($err_2) {
  $__num = \func_num_args();
  $__res = function($sup_3) use ($err_2) {
  $__num = \func_num_args();
  $__res = \Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Effect_Aff__sequential'], (($GLOBALS['Data_Foldable_foldrArray'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_4) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($a_4) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__parmajAffmajApply(\Effect\Aff\majEffect_majAff__parmajAffmajMap(function($v_6) {
  $__num = \func_num_args();
  $__res = function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $a_4), $b_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Unsafe_Coerce_unsafeCoerce']))(function($x_4) {
  $__num = \func_num_args();
  $__res = $x_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(\Control\Semigroupoid\majControl_majSemigroupoid_composemajImpl($GLOBALS['Unsafe_Coerce_unsafeCoerce'], $GLOBALS['Effect_Aff__pure'], $GLOBALS['Data_Unit_unit'])), [\Effect\Aff\majEffect_majAff_killmajFiber($err_2, ($sup_3)->{'fiber'}), ($GLOBALS['Effect_Aff__makeAff'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], function($k_4) use ($err_2, $sup_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff__killAll'])($err_2, ($sup_3)->{'supervisor'}, ($k_4)(new \Data\Either\Data_Either_Right($GLOBALS['Data_Unit_unit'])));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "failed" => function($v_2) use ($killError_1_0) {
  $__num = \func_num_args();
  $__res = function($sup_3) use ($killError_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff__makeAff'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], function($k_4) use ($killError_1_0, $sup_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff__killAll'])($killError_1_0, ($sup_3)->{'supervisor'}, ($k_4)(new \Data\Either\Data_Either_Right($GLOBALS['Data_Unit_unit'])));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "completed" => function($v_2) use ($killError_1_0) {
  $__num = \func_num_args();
  $__res = function($sup_3) use ($killError_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff__makeAff'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], function($k_4) use ($killError_1_0, $sup_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Effect_Aff__killAll'])($killError_1_0, ($sup_3)->{'supervisor'}, ($k_4)(new \Data\Either\Data_Either_Right($GLOBALS['Data_Unit_unit'])));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}], (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff_joinFiber']))(function($v_2) {
  $__num = \func_num_args();
  $__res = ($v_2)->{'fiber'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_supervise'] = __NAMESPACE__ . '\\majEffect_majAff_supervise';

// Effect_Aff_monadSTAff
$GLOBALS['Effect_Aff_monadSTAff'] = (object)["liftST" => (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff__liftEffect']))($GLOBALS['Unsafe_Coerce_unsafeCoerce']), "Monad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_monadAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_monadThrowAff
$GLOBALS['Effect_Aff_monadThrowAff'] = (object)["throwError" => $GLOBALS['Effect_Aff__throwError'], "Monad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_monadAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_monadErrorAff
$GLOBALS['Effect_Aff_monadErrorAff'] = (object)["catchError" => $GLOBALS['Effect_Aff__catchError'], "MonadThrow0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_monadThrowAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_attempt
function majEffect_majAff_attempt($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_attempt';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Effect\Aff\majEffect_majAff__catchmajError(\Effect\Aff\majEffect_majAff__map($GLOBALS['Data_Either_Right'], $a_0), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff__pure']))($GLOBALS['Data_Either_Left']));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_attempt'] = __NAMESPACE__ . '\\majEffect_majAff_attempt';

// Effect_Aff_runAff
function majEffect_majAff_runmajAff($k_0, $aff_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_runmajAff';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Effect\Aff\majEffect_majAff_launchmajAff(\Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__catchmajError(\Effect\Aff\majEffect_majAff__map($GLOBALS['Data_Either_Right'], $aff_1), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff__pure']))($GLOBALS['Data_Either_Left'])), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff__liftEffect']))($k_0)));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_runAff'] = __NAMESPACE__ . '\\majEffect_majAff_runmajAff';

// Effect_Aff_runAff_
function majEffect_majAff_runmajAff_($k_0, $aff_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_runmajAff_';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = \Effect\Aff\majEffect_majAff_runmajAff($k_0, $aff_1);
  $__res = function() use ($__local_var_2_0, &$__fn) {
$a_prime__3_1 = phpurs_execute_effect($__local_var_2_0);
return phpurs_execute_effect($GLOBALS['Data_Unit_unit']);
};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_runAff_'] = __NAMESPACE__ . '\\majEffect_majAff_runmajAff_';

// Effect_Aff_runSuspendedAff
function majEffect_majAff_runmajSuspendedmajAff($k_0, $aff_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_runmajSuspendedmajAff';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Effect_Aff__makeFiber'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__catchmajError(\Effect\Aff\majEffect_majAff__map($GLOBALS['Data_Either_Right'], $aff_1), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff__pure']))($GLOBALS['Data_Either_Left'])), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff__liftEffect']))($k_0)));
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Effect_Aff_runSuspendedAff'] = __NAMESPACE__ . '\\majEffect_majAff_runmajSuspendedmajAff';

// Effect_Aff_monadRecAff
$GLOBALS['Effect_Aff_monadRecAff'] = (object)["tailRecM" => function($k_0) {
  $__num = \func_num_args();
  $go__go_1_0 = null;
  $go__go_1_0 = function($a_2) use (&$go__go_1_0, $k_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(($k_0)($a_2), function($res_3) use (&$go__go_1_0) {
  $__num = \func_num_args();
  $__t1 = null;;
  if ($res_3 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done) {
$__t1 = \Effect\Aff\majEffect_majAff__pure(($res_3)->{'value0'});
goto end_branch_1;;
};
  if ($res_3 instanceof \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop) {
$__t1 = ($go__go_1_0)(($res_3)->{'value0'});
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $__res = $go__go_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Monad0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_monadAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_monoidAff
function majEffect_majAff_monoidmajAff($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_monoidmajAff';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ((($dictMonoid_0)->{'Semigroup0'})(null))->{'append'};
  $semigroupAff1_1_0 = (object)["append" => function($a_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = function($b_3) use ($__local_var_1_0, $a_2) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind(\Effect\Aff\majEffect_majAff__map($__local_var_1_0, $a_2), function($f_prime__4) use ($b_3) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__bind($b_3, function($a_prime__5) use ($f_prime__4) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__pure(($f_prime__4)($a_prime__5));
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
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = (object)["mempty" => \Effect\Aff\majEffect_majAff__pure(($dictMonoid_0)->{'mempty'}), "Semigroup0" => function($_dollar___unused_2) use ($semigroupAff1_1_0) {
  $__num = \func_num_args();
  $__res = $semigroupAff1_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_monoidAff'] = __NAMESPACE__ . '\\majEffect_majAff_monoidmajAff';

// Effect_Aff_nonCanceler_closure
$GLOBALS['Effect_Aff_nonCanceler_closure'] = (function() use (&$__fn) {
$__local_var_0_0 = \Effect\Aff\majEffect_majAff__pure($GLOBALS['Data_Unit_unit']);
return function($v_1) use ($__local_var_0_0) {
  $__num = \func_num_args();
  $__res = $__local_var_0_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
})();

// Effect_Aff_nonCanceler
function majEffect_majAff_nonmajCanceler($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_nonmajCanceler';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff_nonCanceler_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_nonCanceler'] = __NAMESPACE__ . '\\majEffect_majAff_nonmajCanceler';

// Effect_Aff_monoidCanceler
$GLOBALS['Effect_Aff_monoidCanceler'] = (object)["mempty" => $GLOBALS['Effect_Aff_nonCanceler'], "Semigroup0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_semigroupCanceler'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_never
$GLOBALS['Effect_Aff_never'] = ($GLOBALS['Effect_Aff__makeAff'])($GLOBALS['Effect_Aff_isLeft'], $GLOBALS['Effect_Aff_unsafeFromLeft'], $GLOBALS['Effect_Aff_unsafeFromRight'], $GLOBALS['Data_Either_Left'], $GLOBALS['Data_Either_Right'], function($v_0) {
  $__num = \func_num_args();
  $__res = function() use (&$__fn) {
return $GLOBALS['Effect_Aff_nonCanceler'];
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Effect_Aff_apathize_closure
$GLOBALS['Effect_Aff_apathize_closure'] = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Effect_Aff__map'])(function($v_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($a_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__catchmajError(\Effect\Aff\majEffect_majAff__map($GLOBALS['Data_Either_Right'], $a_0), (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Effect_Aff__pure']))($GLOBALS['Data_Either_Left']));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Effect_Aff_apathize
function majEffect_majAff_apathize($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majEffect_majAff_apathize';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Effect_Aff_apathize_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Effect_Aff_apathize'] = __NAMESPACE__ . '\\majEffect_majAff_apathize';

// Effect_Aff_altParAff
$GLOBALS['Effect_Aff_altParAff'] = (object)["alt" => $GLOBALS['Effect_Aff__parAffAlt'], "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_functorParAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_altAff
$GLOBALS['Effect_Aff_altAff'] = (object)["alt" => function($a1_0) {
  $__num = \func_num_args();
  $__res = function($a2_1) use ($a1_0) {
  $__num = \func_num_args();
  $__res = \Effect\Aff\majEffect_majAff__catchmajError($a1_0, function($v_2) use ($a2_1) {
  $__num = \func_num_args();
  $__res = $a2_1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_functorAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_plusAff
$GLOBALS['Effect_Aff_plusAff'] = (object)["empty" => \Effect\Aff\majEffect_majAff__throwmajError(\Effect\Exception\majEffect_majException_error("Always fails")), "Alt0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_altAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_plusParAff
$GLOBALS['Effect_Aff_plusParAff'] = (object)["empty" => \Effect\Aff\majEffect_majAff__throwmajError(\Effect\Exception\majEffect_majException_error("Always fails")), "Alt0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_altParAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

// Effect_Aff_alternativeParAff
$GLOBALS['Effect_Aff_alternativeParAff'] = (object)["Applicative0" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_applicativeParAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Plus1" => function($_dollar___unused_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Effect_Aff_plusParAff'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];

