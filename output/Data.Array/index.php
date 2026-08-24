<?php

namespace Data\Array;

// ALL IMPORTS: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Monad.ST, Control.Monad.ST.Internal, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty.Internal, Data.Array.ST, Data.Array.ST.Iterator, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude, Prim
// TO REQUIRE: Control.Alt, Control.Alternative, Control.Applicative, Control.Apply, Control.Bind, Control.Category, Control.Lazy, Control.Monad.Rec.Class, Control.Monad.ST, Control.Monad.ST.Internal, Control.Semigroupoid, Data.Array, Data.Array.NonEmpty.Internal, Data.Array.ST, Data.Array.ST.Iterator, Data.Boolean, Data.Eq, Data.Foldable, Data.Function, Data.Function.Uncurried, Data.Functor, Data.FunctorWithIndex, Data.HeytingAlgebra, Data.Maybe, Data.Ord, Data.Ordering, Data.Ring, Data.Semigroup, Data.Semiring, Data.Traversable, Data.Tuple, Data.Unfoldable, Partial.Unsafe, Prelude
require_once __DIR__ . '/../Control.Alt/index.php';
require_once __DIR__ . '/../Control.Alternative/index.php';
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Bind/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Lazy/index.php';
require_once __DIR__ . '/../Control.Monad.Rec.Class/index.php';
require_once __DIR__ . '/../Control.Monad.ST/index.php';
require_once __DIR__ . '/../Control.Monad.ST.Internal/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Array/index.php';
require_once __DIR__ . '/../Data.Array.NonEmpty.Internal/index.php';
require_once __DIR__ . '/../Data.Array.ST/index.php';
require_once __DIR__ . '/../Data.Array.ST.Iterator/index.php';
require_once __DIR__ . '/../Data.Boolean/index.php';
require_once __DIR__ . '/../Data.Eq/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Function.Uncurried/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.FunctorWithIndex/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Ord/index.php';
require_once __DIR__ . '/../Data.Ordering/index.php';
require_once __DIR__ . '/../Data.Ring/index.php';
require_once __DIR__ . '/../Data.Semigroup/index.php';
require_once __DIR__ . '/../Data.Semiring/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
require_once __DIR__ . '/../Data.Unfoldable/index.php';
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
if (!\function_exists(__NAMESPACE__ . '\\phpurs_execute_effect')) {
  function phpurs_execute_effect($val) {
    if (\is_callable($val)) {
      return $val($GLOBALS['Data_Unit_unit']);
    }
    return $val;
  }
}

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Array = \call_user_func(function() {
  $exports = [];
$rangeImpl = function($start, $end) use (&$rangeImpl) {
    $step = $start > $end ? -1 : 1;
    $result = [];
    $i = $start;
    while ($i !== $end) {
        $result[] = $i;
        $i += $step;
    }
    $result[] = $i;
    return $result;
};

$replicateImpl = function($count, $value) use (&$replicateImpl) {
    if ($count < 1) return [];
    return array_fill(0, $count, $value);
};

$fromFoldableImpl = function($foldr, $xs) use (&$fromFoldableImpl) {
    $step = function($x) {
        return function($acc) use ($x) {
            $acc[] = $x;
            return $acc;
        };
    };
    $arr = $foldr($step)([])($xs);
    return array_reverse($arr);
};

$length = function($xs) use (&$length) {
    return \count($xs);
};

$unconsImpl = function($empty, $next, $xs) use (&$unconsImpl) {
    if (\count($xs) === 0) return $empty((object)[]);
    return $next($xs[0])(\array_slice($xs, 1));
};

$indexImpl = function($just, $nothing, $xs, $i) use (&$indexImpl) {
    return ($i < 0 || $i >= \count($xs)) ? $nothing : $just($xs[$i]);
};

$findMapImpl = function($nothing, $isJust, $f, $xs) use (&$findMapImpl) {
    foreach ($xs as $x) {
        $result = $f($x);
        if ($isJust($result)) return $result;
    }
    return $nothing;
};

$findIndexImpl = function($just, $nothing, $f, $xs) use (&$findIndexImpl) {
    foreach ($xs as $i => $x) {
        if ($f($x)) return $just($i);
    }
    return $nothing;
};

$findLastIndexImpl = function($just, $nothing, $f, $xs) use (&$findLastIndexImpl) {
    for ($i = \count($xs) - 1; $i >= 0; $i--) {
        if ($f($xs[$i])) return $just($i);
    }
    return $nothing;
};

$_insertAt = function($just, $nothing, $i, $a, $l) use (&$_insertAt) {
    if ($i < 0 || $i > \count($l)) return $nothing;
    $l1 = $l;
    array_splice($l1, $i, 0, [$a]);
    return $just($l1);
};

$_deleteAt = function($just, $nothing, $i, $l) use (&$_deleteAt) {
    if ($i < 0 || $i >= \count($l)) return $nothing;
    $l1 = $l;
    array_splice($l1, $i, 1);
    return $just($l1);
};

$_updateAt = function($just, $nothing, $i, $a, $l) use (&$_updateAt) {
    if ($i < 0 || $i >= \count($l)) return $nothing;
    $l1 = $l;
    $l1[$i] = $a;
    return $just($l1);
};

$reverse = function($l) use (&$reverse) {
    return array_reverse($l);
};

$concat = function($xss) use (&$concat) {
    if (\count($xss) === 0) return [];
    return \array_merge(...$xss);
};

$filterImpl = function($f, $xs) use (&$filterImpl) {
    $res = [];
    foreach ($xs as $x) {
        if ($f($x)) $res[] = $x;
    }
    return $res;
};

$partitionImpl = function($f, $xs) use (&$partitionImpl) {
    $yes = [];
    $no = [];
    foreach ($xs as $x) {
        if ($f($x)) $yes[] = $x;
        else $no[] = $x;
    }
    return (object)["yes" => $yes, "no" => $no];
};

$scanlImpl = function($f, $b, $xs) use (&$scanlImpl) {
    $acc = $b;
    $out = [];
    foreach ($xs as $x) {
        $acc = $f($acc)($x);
        $out[] = $acc;
    }
    return $out;
};

$scanrImpl = function($f, $b, $xs) use (&$scanrImpl) {
    $len = \count($xs);
    $acc = $b;
    $out = array_fill(0, $len, null);
    for ($i = $len - 1; $i >= 0; $i--) {
        $acc = $f($xs[$i])($acc);
        $out[$i] = $acc;
    }
    return $out;
};

$sortByImpl = function($compare, $fromOrdering, $xs) use (&$sortByImpl) {
    $out = $xs;
    \usort($out, function($a, $b) use ($compare, $fromOrdering) {
        return $fromOrdering($compare($a)($b));
    });
    return $out;
};

$sliceImpl = function($s, $e, $l) use (&$sliceImpl) {
    return \array_slice($l, $s, $e - $s);
};

$zipWithImpl = function($f, $xs, $ys) use (&$zipWithImpl) {
    $l = \min(\count($xs), \count($ys));
    $result = [];
    for ($i = 0; $i < $l; $i++) {
        $result[] = $f($xs[$i])($ys[$i]);
    }
    return $result;
};

$anyImpl = function($p, $xs) use (&$anyImpl) {
    foreach ($xs as $x) {
        if ($p($x)) return true;
    }
    return false;
};

$allImpl = function($p, $xs) use (&$allImpl) {
    foreach ($xs as $x) {
        if (!$p($x)) return false;
    }
    return true;
};

$unsafeIndexImpl = function($xs, $n) use (&$unsafeIndexImpl) {
    return $xs[$n];
};

$exports['rangeImpl'] = $rangeImpl;
$exports['replicateImpl'] = $replicateImpl;
$exports['fromFoldableImpl'] = $fromFoldableImpl;
$exports['length'] = $length;
$exports['unconsImpl'] = $unconsImpl;
$exports['indexImpl'] = $indexImpl;
$exports['findMapImpl'] = $findMapImpl;
$exports['findIndexImpl'] = $findIndexImpl;
$exports['findLastIndexImpl'] = $findLastIndexImpl;
$exports['_insertAt'] = $_insertAt;
$exports['_deleteAt'] = $_deleteAt;
$exports['_updateAt'] = $_updateAt;
$exports['reverse'] = $reverse;
$exports['concat'] = $concat;
$exports['filterImpl'] = $filterImpl;
$exports['partitionImpl'] = $partitionImpl;
$exports['scanlImpl'] = $scanlImpl;
$exports['scanrImpl'] = $scanrImpl;
$exports['sortByImpl'] = $sortByImpl;
$exports['sliceImpl'] = $sliceImpl;
$exports['zipWithImpl'] = $zipWithImpl;
$exports['anyImpl'] = $anyImpl;
$exports['allImpl'] = $allImpl;
$exports['unsafeIndexImpl'] = $unsafeIndexImpl;
return $exports;
  return $exports;
});
$GLOBALS['Data_Array__deleteAt'] = (\array_key_exists('_deleteAt', $ffi_Data_Array) ? $ffi_Data_Array['_deleteAt'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array__insertAt'] = (\array_key_exists('_insertAt', $ffi_Data_Array) ? $ffi_Data_Array['_insertAt'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array__updateAt'] = (\array_key_exists('_updateAt', $ffi_Data_Array) ? $ffi_Data_Array['_updateAt'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_allImpl'] = (\array_key_exists('allImpl', $ffi_Data_Array) ? $ffi_Data_Array['allImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_anyImpl'] = (\array_key_exists('anyImpl', $ffi_Data_Array) ? $ffi_Data_Array['anyImpl'] : new class { public function __invoke(...$args) { return $this; } });
function majData_majArray_concat($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majArray_concat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Array;
  $f = (\array_key_exists('concat', $ffi_Data_Array) ? $ffi_Data_Array['concat'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Array_concat'] = __NAMESPACE__ . '\\majData_majArray_concat';

$GLOBALS['Data_Array_filterImpl'] = (\array_key_exists('filterImpl', $ffi_Data_Array) ? $ffi_Data_Array['filterImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_findIndexImpl'] = (\array_key_exists('findIndexImpl', $ffi_Data_Array) ? $ffi_Data_Array['findIndexImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_findLastIndexImpl'] = (\array_key_exists('findLastIndexImpl', $ffi_Data_Array) ? $ffi_Data_Array['findLastIndexImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_findMapImpl'] = (\array_key_exists('findMapImpl', $ffi_Data_Array) ? $ffi_Data_Array['findMapImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_fromFoldableImpl'] = (\array_key_exists('fromFoldableImpl', $ffi_Data_Array) ? $ffi_Data_Array['fromFoldableImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_indexImpl'] = (\array_key_exists('indexImpl', $ffi_Data_Array) ? $ffi_Data_Array['indexImpl'] : new class { public function __invoke(...$args) { return $this; } });
function majData_majArray_length($v0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majArray_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Array;
  $f = (\array_key_exists('length', $ffi_Data_Array) ? $ffi_Data_Array['length'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Array_length'] = __NAMESPACE__ . '\\majData_majArray_length';

$GLOBALS['Data_Array_partitionImpl'] = (\array_key_exists('partitionImpl', $ffi_Data_Array) ? $ffi_Data_Array['partitionImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_rangeImpl'] = (\array_key_exists('rangeImpl', $ffi_Data_Array) ? $ffi_Data_Array['rangeImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_replicateImpl'] = (\array_key_exists('replicateImpl', $ffi_Data_Array) ? $ffi_Data_Array['replicateImpl'] : new class { public function __invoke(...$args) { return $this; } });
function majData_majArray_reverse($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majArray_reverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Array;
  $f = (\array_key_exists('reverse', $ffi_Data_Array) ? $ffi_Data_Array['reverse'] : new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Array_reverse'] = __NAMESPACE__ . '\\majData_majArray_reverse';

$GLOBALS['Data_Array_scanlImpl'] = (\array_key_exists('scanlImpl', $ffi_Data_Array) ? $ffi_Data_Array['scanlImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_scanrImpl'] = (\array_key_exists('scanrImpl', $ffi_Data_Array) ? $ffi_Data_Array['scanrImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_sliceImpl'] = (\array_key_exists('sliceImpl', $ffi_Data_Array) ? $ffi_Data_Array['sliceImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_sortByImpl'] = (\array_key_exists('sortByImpl', $ffi_Data_Array) ? $ffi_Data_Array['sortByImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_unconsImpl'] = (\array_key_exists('unconsImpl', $ffi_Data_Array) ? $ffi_Data_Array['unconsImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_unsafeIndexImpl'] = (\array_key_exists('unsafeIndexImpl', $ffi_Data_Array) ? $ffi_Data_Array['unsafeIndexImpl'] : new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_zipWithImpl'] = (\array_key_exists('zipWithImpl', $ffi_Data_Array) ? $ffi_Data_Array['zipWithImpl'] : new class { public function __invoke(...$args) { return $this; } });




// Data_Array_zipWith
function majData_majArray_zipmajWith($__local_var_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_zipmajWith';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_zipWithImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_zipWith'] = __NAMESPACE__ . '\\majData_majArray_zipmajWith';

// Data_Array_zipWithA
function majData_majArray_zipmajWithmajA($dictApplicative_0, $f_1 = null, $xs_2 = null, $ys_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_zipmajWithmajA';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $Apply0_4_0 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = \Data\Traversable\majData_majTraversable_traversemajArraymajImpl(($Apply0_4_0)->{'apply'}, ((($Apply0_4_0)->{'Functor0'})(null))->{'map'}, ($dictApplicative_0)->{'pure'}, function($x_5) {
  $__num = \func_num_args();
  $__res = $x_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, ($GLOBALS['Data_Array_zipWithImpl'])($f_1, $xs_2, $ys_3));
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Array_zipWithA'] = __NAMESPACE__ . '\\majData_majArray_zipmajWithmajA';

// Data_Array_zip_closure
$GLOBALS['Data_Array_zip_closure'] = ($GLOBALS['Data_Array_zipWith'])($GLOBALS['Data_Tuple_Tuple']);

// Data_Array_zip
function majData_majArray_zip($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_zip';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_zip_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_zip'] = __NAMESPACE__ . '\\majData_majArray_zip';

// Data_Array_updateAtIndices
function majData_majArray_updatemajAtmajIndices($dictFoldable_0, $us_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_updatemajAtmajIndices';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($dictFoldable_0, $us_1, $xs_2, &$__fn) {
$result_3_0 = ($GLOBALS['Data_Array_ST_thawImpl'])($xs_2);
$_dollar___unused_4_1 = phpurs_execute_effect((((($dictFoldable_0)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_4) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($a_4) {
  $__num = \func_num_args();
  $__local_var_6_1 = phpurs_execute_effect($a_4);
  $f_prime__6_1 = phpurs_execute_effect(function($x_7) {
  $__num = \func_num_args();
  $__res = $x_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $a_prime__7_3 = phpurs_execute_effect($b_5);
  $__res = phpurs_execute_effect(phpurs_execute_effect(($f_prime__6_1)($a_prime__7_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($v_4) use ($result_3_0) {
  $__num = \func_num_args();
  $__local_var_5_4 = ($v_4)->{'value1'};
  $__local_var_6_5 = ($v_4)->{'value0'};
  $__res = ($GLOBALS['Data_Array_ST_pokeImpl'])($__local_var_6_5, $__local_var_5_4, $result_3_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Unit_unit']))($us_1));
return phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_3_0)));
});
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_updateAtIndices'] = __NAMESPACE__ . '\\majData_majArray_updatemajAtmajIndices';

// Data_Array_updateAt
function majData_majArray_updatemajAt(int $__local_var_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_updatemajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array__updateAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_updateAt'] = __NAMESPACE__ . '\\majData_majArray_updatemajAt';

// Data_Array_unsafeIndex
function majData_majArray_unsafemajIndex($_dollar___unused_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_unsafemajIndex';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($__local_var_1)[$__local_var_2];
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_unsafeIndex'] = __NAMESPACE__ . '\\majData_majArray_unsafemajIndex';

// Data_Array_uncons
function majData_majArray_uncons($__local_var_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_uncons';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_1) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($x_1) {
  $__num = \func_num_args();
  $__res = function($xs_2) use ($x_1) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Just((object)["head" => $x_1, "tail" => $xs_2]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__local_var_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_uncons'] = __NAMESPACE__ . '\\majData_majArray_uncons';

// Data_Array_toUnfoldable
function majData_majArray_tomajUnfoldable($dictUnfoldable_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_tomajUnfoldable';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $len_2_0 = count($xs_1);
  $__res = ((($dictUnfoldable_0)->{'unfoldr'})(function($i_3) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($i_3 < $len_2_0)) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($xs_1)[$i_3], ($i_3 + 1)));
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_toUnfoldable'] = __NAMESPACE__ . '\\majData_majArray_tomajUnfoldable';

// Data_Array_tail
function majData_majArray_tail($__local_var_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_tail';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_1) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Nothing();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($v_1) {
  $__num = \func_num_args();
  $__res = function($xs_2) {
  $__num = \func_num_args();
  $__res = new \Data\Maybe\Data_Maybe_Just($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__local_var_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_tail'] = __NAMESPACE__ . '\\majData_majArray_tail';

// Data_Array_sortBy
function majData_majArray_sortmajBy($comp_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_sortmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_sortByImpl'])($comp_0, function($v_2) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v_2 instanceof \Data\Ordering\Data_Ordering_GT) {
$__t0 = 1;
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Ordering\Data_Ordering_EQ) {
$__t0 = 0;
goto end_branch_0;;
};
  if ($v_2 instanceof \Data\Ordering\Data_Ordering_LT) {
$__t0 = -1;
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_sortBy'] = __NAMESPACE__ . '\\majData_majArray_sortmajBy';

// Data_Array_sortWith
function majData_majArray_sortmajWith($dictOrd_0, $f_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_sortmajWith';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_sortBy'])(function($x_2) use ($dictOrd_0, $f_1) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($dictOrd_0, $f_1, $x_2) {
  $__num = \func_num_args();
  $__res = ((($dictOrd_0)->{'compare'})(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_sortWith'] = __NAMESPACE__ . '\\majData_majArray_sortmajWith';

// Data_Array_sort
function majData_majArray_sort($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_sort';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $compare_1_0 = ($dictOrd_0)->{'compare'};
  $__res = function($xs_2) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_sortmajBy($compare_1_0, $xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_sort'] = __NAMESPACE__ . '\\majData_majArray_sort';

// Data_Array_snoc
function majData_majArray_snoc($xs_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_snoc';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($GLOBALS['Data_Array_ST_push'])($x_1);
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($__local_var_2_0, $xs_0, &$__fn) {
$result_3_1 = ($GLOBALS['Data_Array_ST_thawImpl'])($xs_0);
$_dollar___unused_4_2 = phpurs_execute_effect(($__local_var_2_0)($result_3_1));
return phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_3_1)));
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_snoc'] = __NAMESPACE__ . '\\majData_majArray_snoc';

// Data_Array_slice
function majData_majArray_slice(int $__local_var_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_slice';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_sliceImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_slice'] = __NAMESPACE__ . '\\majData_majArray_slice';

// Data_Array_splitAt
function majData_majArray_splitmajAt(int $v_0, $v1_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_splitmajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($v_0 <= 0)) {
$__t0 = (object)["before" => [], "after" => $v1_1];
goto end_branch_0;;
};
  $__t0 = (object)["before" => ($GLOBALS['Data_Array_sliceImpl'])(0, $v_0, $v1_1), "after" => ($GLOBALS['Data_Array_sliceImpl'])($v_0, count($v1_1), $v1_1)];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_splitAt'] = __NAMESPACE__ . '\\majData_majArray_splitmajAt';

// Data_Array_take
function majData_majArray_take(int $n_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_take';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($n_0 < 1)) {
$__t0 = [];
goto end_branch_0;;
};
  $__t0 = ($GLOBALS['Data_Array_sliceImpl'])(0, $n_0, $xs_1);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_take'] = __NAMESPACE__ . '\\majData_majArray_take';

// Data_Array_singleton
function majData_majArray_singleton($a_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_singleton';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = [$a_0];
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_singleton'] = __NAMESPACE__ . '\\majData_majArray_singleton';

// Data_Array_scanr
function majData_majArray_scanr($__local_var_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_scanr';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_scanrImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_scanr'] = __NAMESPACE__ . '\\majData_majArray_scanr';

// Data_Array_scanl
function majData_majArray_scanl($__local_var_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_scanl';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_scanlImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_scanl'] = __NAMESPACE__ . '\\majData_majArray_scanl';

// Data_Array_replicate
function majData_majArray_replicate(int $__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_replicate';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_replicateImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_replicate'] = __NAMESPACE__ . '\\majData_majArray_replicate';

// Data_Array_range
function majData_majArray_range(int $__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_range';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_rangeImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_range'] = __NAMESPACE__ . '\\majData_majArray_range';

// Data_Array_partition
function majData_majArray_partition($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_partition';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_partitionImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_partition'] = __NAMESPACE__ . '\\majData_majArray_partition';

// Data_Array_null
function majData_majArray_null($xs_0): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_null';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (count($xs_0) === 0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_null'] = __NAMESPACE__ . '\\majData_majArray_null';

// Data_Array_modifyAtIndices
function majData_majArray_modifymajAtmajIndices($dictFoldable_0, $is_1 = null, $f_2 = null, $xs_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_modifymajAtmajIndices';
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($dictFoldable_0, $f_2, $is_1, $xs_3, &$__fn) {
$result_4_0 = ($GLOBALS['Data_Array_ST_thawImpl'])($xs_3);
$_dollar___unused_5_1 = phpurs_execute_effect((((($dictFoldable_0)->{'foldr'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($a_5) {
  $__num = \func_num_args();
  $__res = function($b_6) use ($a_5) {
  $__num = \func_num_args();
  $__local_var_7_1 = phpurs_execute_effect($a_5);
  $f_prime__7_1 = phpurs_execute_effect(function($x_8) {
  $__num = \func_num_args();
  $__res = $x_8;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  $a_prime__8_3 = phpurs_execute_effect($b_6);
  $__res = phpurs_execute_effect(phpurs_execute_effect(($f_prime__7_1)($a_prime__8_3)));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(function($i_5) use ($f_2, $result_4_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\ST\majData_majArray_majSmajT_modify($i_5, $f_2, $result_4_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))($GLOBALS['Data_Unit_unit']))($is_1));
return phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_4_0)));
});
  goto __end;;
  __end:
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Array_modifyAtIndices'] = __NAMESPACE__ . '\\majData_majArray_modifymajAtmajIndices';

// Data_Array_mapWithIndex_closure
$GLOBALS['Data_Array_mapWithIndex_closure'] = $GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'];

// Data_Array_mapWithIndex
function majData_majArray_mapmajWithmajIndex($v_0, $v_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_mapmajWithmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_mapWithIndex_closure'])($v_0, $v_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_mapWithIndex'] = __NAMESPACE__ . '\\majData_majArray_mapmajWithmajIndex';

// Data_Array_intersperse
function majData_majArray_intersperse($a_0, $arr_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_intersperse';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = count($arr_1);
  $__t7 = null;;
  if (($v_2_0 < 2)) {
$__t7 = $arr_1;
goto end_branch_7;;
};
  $__t7 = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($a_0, $arr_1, $v_2_0, &$__fn) {
$out_3_1 = phpurs_execute_effect($GLOBALS['Data_Array_ST_new']);
$_dollar___unused_4_2 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_pushImpl'])(($arr_1)[0], $out_3_1));
$_dollar___unused_5_3 = phpurs_execute_effect(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_for(1, $v_2_0, function($idx_5) use ($a_0, $arr_1, $out_3_1) {
  $__num = \func_num_args();
  $_dollar___unused_6_3 = ($GLOBALS['Data_Array_ST_pushImpl'])($a_0, $out_3_1);
  $__local_var_7_4 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_pushImpl'])(($arr_1)[$idx_5], $out_3_1));
  $__res = phpurs_execute_effect($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$__local_var_3_1 = phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect($out_3_1))));
return phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($__local_var_3_1));
});
  end_branch_7:;
  $__res = $__t7;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_intersperse'] = __NAMESPACE__ . '\\majData_majArray_intersperse';

// Data_Array_intercalate
function majData_majArray_intercalate($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_intercalate';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_1 = ($dictMonoid_0)->{'mempty'};
  $__res = function($sep_3) use ($Semigroup0_1_0, $mempty_2_1) {
  $__num = \func_num_args();
  $__res = function($xs_4) use ($Semigroup0_1_0, $mempty_2_1, $sep_3) {
  $__num = \func_num_args();
  $__res = (\Data\Foldable\majData_majFoldable_foldlmajArray(function($v_5) use ($Semigroup0_1_0, $sep_3) {
  $__num = \func_num_args();
  $__res = function($v1_6) use ($Semigroup0_1_0, $sep_3, $v_5) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($v_5)->{'init'}) {
$__t2 = (object)["init" => false, "acc" => $v1_6];
goto end_branch_2;;
};
  $__t2 = (object)["init" => false, "acc" => ((($Semigroup0_1_0)->{'append'})(($v_5)->{'acc'}))(((($Semigroup0_1_0)->{'append'})($sep_3))($v1_6))];
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (object)["init" => true, "acc" => $mempty_2_1], $xs_4))->{'acc'};
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_intercalate'] = __NAMESPACE__ . '\\majData_majArray_intercalate';

// Data_Array_insertAt
function majData_majArray_insertmajAt(int $__local_var_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_insertmajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array__insertAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_insertAt'] = __NAMESPACE__ . '\\majData_majArray_insertmajAt';

// Data_Array_init
function majData_majArray_init($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_init';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = match (count($xs_0)) { 0 => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Array_sliceImpl'])(0, (count($xs_0) - 1), $xs_0)) };
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_init'] = __NAMESPACE__ . '\\majData_majArray_init';

// Data_Array_index
function majData_majArray_index($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_index';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_index'] = __NAMESPACE__ . '\\majData_majArray_index';

// Data_Array_last
function majData_majArray_last($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_last';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_0, (count($xs_0) - 1));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_last'] = __NAMESPACE__ . '\\majData_majArray_last';

// Data_Array_unsnoc
function majData_majArray_unsnoc($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_unsnoc';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t2 = null;;
  switch (count($xs_0)) {
case 0:
$__t2 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_2;;
break;
default:
;
break;
};
  $__local_var_1_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_0, (count($xs_0) - 1));
  $__t1 = null;;
  if ($__local_var_1_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just((object)["init" => ($GLOBALS['Data_Array_sliceImpl'])(0, (count($xs_0) - 1), $xs_0), "last" => ($__local_var_1_0)->{'value0'}]);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__t2 = $__t1;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_unsnoc'] = __NAMESPACE__ . '\\majData_majArray_unsnoc';

// Data_Array_modifyAt
function majData_majArray_modifymajAt(int $i_0, $f_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_modifymajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_2, $i_0);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = ($GLOBALS['Data_Array__updateAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $i_0, ($f_1)(($__local_var_3_0)->{'value0'}), $xs_2);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_modifyAt'] = __NAMESPACE__ . '\\majData_majArray_modifymajAt';

// Data_Array_span
function majData_majArray_span($p_0, $arr_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_span';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go__go_2_0 = null;
  $go__go_2_0 = function(int $i_3) use ($arr_1, &$go__go_2_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go__go_2_0_0_i_3 = $i_3;
  tco_loop_go__go_2_0_0:;
  $i_3 = $__tco_var_go__go_2_0_0_i_3;
  $v_4_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $arr_1, $i_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = null;;
if (($p_0)(($v_4_0)->{'value0'})) {
$__tco_3 = ($i_3 + 1);
$__tco_var_go__go_2_0_0_i_3 = $__tco_3;
goto tco_loop_go__go_2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
$__t2 = new \Data\Maybe\Data_Maybe_Just($i_3);
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  $breakIndex_3_1 = ($go__go_2_0)(0);
  $__t2 = null;;
  if ($breakIndex_3_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = match (($breakIndex_3_1)->{'value0'}) { 0 => (object)["init" => [], "rest" => $arr_1], default => (object)["init" => ($GLOBALS['Data_Array_sliceImpl'])(0, ($breakIndex_3_1)->{'value0'}, $arr_1), "rest" => ($GLOBALS['Data_Array_sliceImpl'])(($breakIndex_3_1)->{'value0'}, count($arr_1), $arr_1)] };
goto end_branch_2;;
};
  if ($breakIndex_3_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = (object)["init" => $arr_1, "rest" => []];
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_span'] = __NAMESPACE__ . '\\majData_majArray_span';

// Data_Array_takeWhile
function majData_majArray_takemajWhile($p_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_takemajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (\Data\Array\majData_majArray_span($p_0, $xs_1))->{'init'};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_takeWhile'] = __NAMESPACE__ . '\\majData_majArray_takemajWhile';

// Data_Array_unzip
function majData_majArray_unzip($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_unzip';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($xs_0, &$__fn) {
$fsts_1_0 = phpurs_execute_effect($GLOBALS['Data_Array_ST_new']);
$snds_2_1 = phpurs_execute_effect($GLOBALS['Data_Array_ST_new']);
$__local_var_3_2 = phpurs_execute_effect("TODO_PrimEffect");
$iter_3_2 = phpurs_execute_effect(new \Data\Array\ST\Iterator\Data_Array_ST_Iterator_Iterator(function($v_4) use ($xs_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_0, $v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__local_var_3_2));
$_dollar___unused_4_4 = phpurs_execute_effect(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_iterate($iter_3_2, function($v_4) use ($fsts_1_0, $snds_2_1) {
  $__num = \func_num_args();
  $__local_var_5_4 = ($v_4)->{'value0'};
  $__local_var_6_5 = ($v_4)->{'value1'};
  $__local_var_7_6 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_pushImpl'])($__local_var_5_4, $fsts_1_0));
  $__local_var_8_7 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_pushImpl'])($__local_var_6_5, $snds_2_1));
  $__res = phpurs_execute_effect($GLOBALS['Data_Unit_unit']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
$fsts_prime__5_9 = ($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($fsts_1_0);
$snds_prime__6_10 = ($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($snds_2_1);
return phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(new \Data\Tuple\Data_Tuple_Tuple($fsts_prime__5_9, $snds_prime__6_10)))))));
});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_unzip'] = __NAMESPACE__ . '\\majData_majArray_unzip';

// Data_Array_head
function majData_majArray_head($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_head';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_0, 0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_head'] = __NAMESPACE__ . '\\majData_majArray_head';

// Data_Array_nubBy
function majData_majArray_nubmajBy($comp_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_nubmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $indexedAndSorted_2_0 = \Data\Array\majData_majArray_sortmajBy(function($x_2) use ($comp_0) {
  $__num = \func_num_args();
  $__res = function($y_3) use ($comp_0, $x_2) {
  $__num = \func_num_args();
  $__res = (($comp_0)(($x_2)->{'value1'}))(($y_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, \Data\FunctorWithIndex\majData_majFunctormajWithmajIndex_mapmajWithmajIndexmajArray($GLOBALS['Data_Tuple_Tuple'], $xs_1));
  $v_3_1 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $indexedAndSorted_2_0, 0);
  $__t2 = null;;
  if ($v_3_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = [];
goto end_branch_2;;
};
  if ($v_3_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_4_3 = ($v_3_1)->{'value0'};
$__t2 = \Data\Functor\majData_majFunctor_arraymajMap($GLOBALS['Data_Tuple_snd'], \Data\Array\majData_majArray_sortmajBy(function($x_5) {
  $__num = \func_num_args();
  $__res = function($y_6) use ($x_5) {
  $__num = \func_num_args();
  $__res = \Data\Ord\majData_majOrd_ordmajIntmajImpl(new \Data\Ordering\Data_Ordering_LT(), new \Data\Ordering\Data_Ordering_EQ(), new \Data\Ordering\Data_Ordering_GT(), ($x_5)->{'value0'}, ($y_6)->{'value0'});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($__local_var_4_3, $comp_0, $indexedAndSorted_2_0, &$__fn) {
$result_5_4 = ($GLOBALS['Data_Array_ST_unsafeThawImpl'])([$__local_var_4_3]);
$_dollar___unused_6_5 = phpurs_execute_effect(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_foreach($indexedAndSorted_2_0, function($v1_6) use ($comp_0, $result_5_4) {
  $__num = \func_num_args();
  $__local_var_7_5 = ($v1_6)->{'value1'};
  $__local_var_8_6 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v_8) {
  $__num = \func_num_args();
  $__t6 = null;;
  if ($v_8 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t6 = ($v_8)->{'value0'};
goto end_branch_6;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t6 = null;
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($GLOBALS['Data_Array_last']));
  $__local_var_9_8 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_5_4));
  $lst_9_8 = phpurs_execute_effect(($__local_var_8_6)($__local_var_9_8));
  $__t10 = null;;
  if (( ! (($comp_0)($lst_9_8))($__local_var_7_5) instanceof \Data\Ordering\Data_Ordering_EQ)) {
$__local_var_10_11 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_pushImpl'])($v1_6, $result_5_4));
$__t10 = $GLOBALS['Data_Unit_unit'];
goto end_branch_10;;
};
  $__t10 = $GLOBALS['Data_Unit_unit'];
  end_branch_10:;
  $__res = phpurs_execute_effect($__t10);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
return phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_5_4)));
})));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_nubBy'] = __NAMESPACE__ . '\\majData_majArray_nubmajBy';

// Data_Array_nub
function majData_majArray_nub($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_nub';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_nubBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_nub'] = __NAMESPACE__ . '\\majData_majArray_nub';

// Data_Array_groupBy
function majData_majArray_groupmajBy($op_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_groupmajBy';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($op_0, $xs_1, &$__fn) {
$result_2_0 = phpurs_execute_effect($GLOBALS['Data_Array_ST_new']);
$__local_var_3_1 = phpurs_execute_effect("TODO_PrimEffect");
$iter_3_1 = phpurs_execute_effect(new \Data\Array\ST\Iterator\Data_Array_ST_Iterator_Iterator(function($v_4) use ($xs_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_1, $v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $__local_var_3_1));
$_dollar___unused_4_3 = phpurs_execute_effect(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_iterate($iter_3_1, function($x_4) use ($iter_3_1, $op_0, $result_2_0) {
  $__num = \func_num_args();
  $sub_5_3 = phpurs_execute_effect($GLOBALS['Data_Array_ST_new']);
  $_dollar___unused_6_4 = ($GLOBALS['Data_Array_ST_pushImpl'])($x_4, $sub_5_3);
  $_dollar___unused_7_5 = phpurs_execute_effect(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_pushmajWhile(($op_0)($x_4), $iter_3_1, $sub_5_3));
  $grp_8_6 = ($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($sub_5_3);
  $__local_var_5_3 = phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_pushImpl'])($grp_8_6, $result_2_0))))));
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
return phpurs_execute_effect(phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($result_2_0))));
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_groupBy'] = __NAMESPACE__ . '\\majData_majArray_groupmajBy';

// Data_Array_groupAllBy
function majData_majArray_groupmajAllmajBy($cmp_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_groupmajAllmajBy';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupBy'])(function($x_1) use ($cmp_0) {
  $__num = \func_num_args();
  $__res = function($y_2) use ($cmp_0, $x_1) {
  $__num = \func_num_args();
  $__res = (($cmp_0)($x_1))($y_2) instanceof \Data\Ordering\Data_Ordering_EQ;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Data_Array_sortBy'])($cmp_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_groupAllBy'] = __NAMESPACE__ . '\\majData_majArray_groupmajAllmajBy';

// Data_Array_groupAll
function majData_majArray_groupmajAll($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_groupmajAll';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_groupAllBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_groupAll'] = __NAMESPACE__ . '\\majData_majArray_groupmajAll';

// Data_Array_group
function majData_majArray_group($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_group';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $eq_1_0 = ($dictEq_0)->{'eq'};
  $__res = function($xs_2) use ($eq_1_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_groupmajBy($eq_1_0, $xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_group'] = __NAMESPACE__ . '\\majData_majArray_group';

// Data_Array_fromFoldable
function majData_majArray_frommajFoldable($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_frommajFoldable';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = ($dictFoldable_0)->{'foldr'};
  $__res = function($__local_var_2) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_fromFoldableImpl'])($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_fromFoldable'] = __NAMESPACE__ . '\\majData_majArray_frommajFoldable';

// Data_Array_foldr_closure
$GLOBALS['Data_Array_foldr_closure'] = $GLOBALS['Data_Foldable_foldrArray'];

// Data_Array_foldr
function majData_majArray_foldr($v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_foldr';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_foldr_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_foldr'] = __NAMESPACE__ . '\\majData_majArray_foldr';

// Data_Array_foldl_closure
$GLOBALS['Data_Array_foldl_closure'] = $GLOBALS['Data_Foldable_foldlArray'];

// Data_Array_foldl
function majData_majArray_foldl($v_0, $v_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_foldl';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_foldl_closure'])($v_0, $v_1, $v_2);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_foldl'] = __NAMESPACE__ . '\\majData_majArray_foldl';

// Data_Array_transpose
function majData_majArray_transpose($xs_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_transpose';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $go__go_1_0 = null;
  $go__go_1_0 = (function() use (&$go__go_1_0, $xs_0) {
  $__fn = function(int $idx_2, $allArrays_3 = null) use (&$go__go_1_0, $xs_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__go_1_0_0_idx_2 = $idx_2;
  $__tco_var_go__go_1_0_0_allArrays_3 = $allArrays_3;
  tco_loop_go__go_1_0_0:;
  $idx_2 = $__tco_var_go__go_1_0_0_idx_2;
  $allArrays_3 = $__tco_var_go__go_1_0_0_allArrays_3;
  $v_4_0 = \Data\Foldable\majData_majFoldable_foldlmajArray(function($acc_4) use ($idx_2) {
  $__num = \func_num_args();
  $__res = function($nextArr_5) use ($acc_4, $idx_2) {
  $__num = \func_num_args();
  $__local_var_6_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $nextArr_5, $idx_2);
  $__t1 = null;;
  if ($__local_var_6_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $acc_4;
goto end_branch_1;;
};
  if ($__local_var_6_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = null;;
if ($acc_4 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = [($__local_var_6_0)->{'value0'}];
goto end_branch_2;;
};
if ($acc_4 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = \Data\Array\majData_majArray_snoc(($acc_4)->{'value0'}, ($__local_var_6_0)->{'value0'});
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t1 = new \Data\Maybe\Data_Maybe_Just($__t2);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, new \Data\Maybe\Data_Maybe_Nothing(), $xs_0);
  $__t4 = null;;
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = $allArrays_3;
goto end_branch_4;;
};
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_5 = ($idx_2 + 1);
$__tco_6 = \Data\Array\majData_majArray_snoc($allArrays_3, ($v_4_0)->{'value0'});
$__tco_var_go__go_1_0_0_idx_2 = $__tco_5;
$__tco_var_go__go_1_0_0_allArrays_3 = $__tco_6;
goto tco_loop_go__go_1_0_0;;
$__t4 = null;
goto end_branch_4;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t4 = null;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  $__res = (($go__go_1_0)(0))([]);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_transpose'] = __NAMESPACE__ . '\\majData_majArray_transpose';

// Data_Array_foldRecM
function majData_majArray_foldmajRecmajM($dictMonadRec_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_foldmajRecmajM';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Monad0_1_0 = (($dictMonadRec_0)->{'Monad0'})(null);
  $Applicative0_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $Bind1_3_2 = (($Monad0_1_0)->{'Bind1'})(null);
  $__res = function($f_4) use ($Applicative0_2_1, $Bind1_3_2, $dictMonadRec_0) {
  $__num = \func_num_args();
  $__res = function($b_5) use ($Applicative0_2_1, $Bind1_3_2, $dictMonadRec_0, $f_4) {
  $__num = \func_num_args();
  $__res = function($array_6) use ($Applicative0_2_1, $Bind1_3_2, $b_5, $dictMonadRec_0, $f_4) {
  $__num = \func_num_args();
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($o_7) use ($Applicative0_2_1, $Bind1_3_2, $array_6, $f_4) {
  $__num = \func_num_args();
  $__t3 = null;;
  if ((($o_7)->{'b'} >= count($array_6))) {
$__t3 = (($Applicative0_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($o_7)->{'a'}));
goto end_branch_3;;
};
  $__t3 = ((($Bind1_3_2)->{'bind'})((($f_4)(($o_7)->{'a'}))(($array_6)[($o_7)->{'b'}])))(function($res_prime__8) use ($Applicative0_2_1, $o_7) {
  $__num = \func_num_args();
  $__res = (($Applicative0_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((object)["a" => $res_prime__8, "b" => (($o_7)->{'b'} + 1)]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $b_5, "b" => 0]);
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_foldRecM'] = __NAMESPACE__ . '\\majData_majArray_foldmajRecmajM';

// Data_Array_foldMap
function majData_majArray_foldmajMap($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_foldmajMap';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_1 = ($dictMonoid_0)->{'mempty'};
  $__res = function($f_3) use ($Semigroup0_1_0, $mempty_2_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldrArray'])(function($x_4) use ($Semigroup0_1_0, $f_3) {
  $__num = \func_num_args();
  $__res = function($acc_5) use ($Semigroup0_1_0, $f_3, $x_4) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_0)->{'append'})(($f_3)($x_4)))($acc_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($mempty_2_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_foldMap'] = __NAMESPACE__ . '\\majData_majArray_foldmajMap';

// Data_Array_foldM
function majData_majArray_foldmajM($dictMonad_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_foldmajM';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__tco_var_Data_Array_foldM_dictMonad_0 = $dictMonad_0;
  tco_loop_Data_Array_foldM:;
  $dictMonad_0 = $__tco_var_Data_Array_foldM_dictMonad_0;
  $Applicative0_1_0 = (($dictMonad_0)->{'Applicative0'})(null);
  $Bind1_2_1 = (($dictMonad_0)->{'Bind1'})(null);
  $__res = function($f_3) use ($Applicative0_1_0, $Bind1_2_1, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = function($b_4) use ($Applicative0_1_0, $Bind1_2_1, $dictMonad_0, $f_3) {
  $__num = \func_num_args();
  $__res = function($__local_var_5) use ($Applicative0_1_0, $Bind1_2_1, $b_4, $dictMonad_0, $f_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_6) use ($Applicative0_1_0, $b_4) {
  $__num = \func_num_args();
  $__res = (($Applicative0_1_0)->{'pure'})($b_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, function($a_6) use ($Bind1_2_1, $b_4, $dictMonad_0, $f_3) {
  $__num = \func_num_args();
  $__res = function($as_7) use ($Bind1_2_1, $a_6, $b_4, $dictMonad_0, $f_3) {
  $__num = \func_num_args();
  $__res = ((($Bind1_2_1)->{'bind'})((($f_3)($b_4))($a_6)))(function($b_prime__8) use ($as_7, $dictMonad_0, $f_3) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Array_foldM'])($dictMonad_0))($f_3))($b_prime__8))($as_7);
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
}, $__local_var_5);
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
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_foldM'] = __NAMESPACE__ . '\\majData_majArray_foldmajM';

// Data_Array_fold
function majData_majArray_fold($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_fold';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Semigroup0_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $__res = (($GLOBALS['Data_Foldable_foldrArray'])(function($x_2) use ($Semigroup0_1_0) {
  $__num = \func_num_args();
  $__res = function($acc_3) use ($Semigroup0_1_0, $x_2) {
  $__num = \func_num_args();
  $__res = ((($Semigroup0_1_0)->{'append'})($x_2))($acc_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($dictMonoid_0)->{'mempty'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_fold'] = __NAMESPACE__ . '\\majData_majArray_fold';

// Data_Array_findMap
function majData_majArray_findmajMap($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_findmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findMapImpl'])(new \Data\Maybe\Data_Maybe_Nothing(), $GLOBALS['Data_Maybe_isJust'], $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_findMap'] = __NAMESPACE__ . '\\majData_majArray_findmajMap';

// Data_Array_findLastIndex
function majData_majArray_findmajLastmajIndex($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_findmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findLastIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_findLastIndex'] = __NAMESPACE__ . '\\majData_majArray_findmajLastmajIndex';

// Data_Array_insertBy
function majData_majArray_insertmajBy($cmp_0, $x_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_insertmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_findLastIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), function($y_3) use ($cmp_0, $x_1) {
  $__num = \func_num_args();
  $__res = (($cmp_0)($x_1))($y_3) instanceof \Data\Ordering\Data_Ordering_GT;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $ys_2);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = 0;
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = (($__local_var_3_0)->{'value0'} + 1);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__local_var_3_0 = ($GLOBALS['Data_Array__insertAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $__t1, $x_1, $ys_2);
  $__t3 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = ($__local_var_3_0)->{'value0'};
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_insertBy'] = __NAMESPACE__ . '\\majData_majArray_insertmajBy';

// Data_Array_insert
function majData_majArray_insert($dictOrd_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_insert';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_insertBy'])(($dictOrd_0)->{'compare'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_insert'] = __NAMESPACE__ . '\\majData_majArray_insert';

// Data_Array_findIndex
function majData_majArray_findmajIndex($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_findmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_findIndex'] = __NAMESPACE__ . '\\majData_majArray_findmajIndex';

// Data_Array_find
function majData_majArray_find($f_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_find';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $f_0, $xs_1);
  $__t1 = null;;
  if ($__local_var_2_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($xs_1)[($__local_var_2_0)->{'value0'}]);
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_find'] = __NAMESPACE__ . '\\majData_majArray_find';

// Data_Array_filter
function majData_majArray_filter($__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_filter';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_filterImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_filter'] = __NAMESPACE__ . '\\majData_majArray_filter';

// Data_Array_intersectBy
function majData_majArray_intersectmajBy($eq_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_intersectmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_filterImpl'])(function($x_3) use ($eq_0, $ys_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), ($eq_0)($x_3), $ys_2);
  $__t1 = null;;
  if ($__local_var_4_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = false;
goto end_branch_1;;
};
  if ($__local_var_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = true;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $xs_1);
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_intersectBy'] = __NAMESPACE__ . '\\majData_majArray_intersectmajBy';

// Data_Array_intersect
function majData_majArray_intersect($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_intersect';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_intersectBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_intersect'] = __NAMESPACE__ . '\\majData_majArray_intersect';

// Data_Array_elemLastIndex
function majData_majArray_elemmajLastmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_elemmajLastmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findLastIndex'])(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_elemLastIndex'] = __NAMESPACE__ . '\\majData_majArray_elemmajLastmajIndex';

// Data_Array_elemIndex
function majData_majArray_elemmajIndex($dictEq_0, $x_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_elemmajIndex';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findIndex'])(function($v_2) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_elemIndex'] = __NAMESPACE__ . '\\majData_majArray_elemmajIndex';

// Data_Array_notElem
function majData_majArray_notmajElem($dictEq_0, $a_1 = null, $arr_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_notmajElem';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), function($v_3) use ($a_1, $dictEq_0) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_3))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $arr_2);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = true;
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = false;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_notElem'] = __NAMESPACE__ . '\\majData_majArray_notmajElem';

// Data_Array_elem
function majData_majArray_elem($dictEq_0, $a_1 = null, $arr_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_elem';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), function($v_3) use ($a_1, $dictEq_0) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)->{'eq'})($v_3))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $arr_2);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = false;
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t1 = true;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_elem'] = __NAMESPACE__ . '\\majData_majArray_elem';

// Data_Array_dropWhile
function majData_majArray_dropmajWhile($p_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_dropmajWhile';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (\Data\Array\majData_majArray_span($p_0, $xs_1))->{'rest'};
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_dropWhile'] = __NAMESPACE__ . '\\majData_majArray_dropmajWhile';

// Data_Array_dropEnd
function majData_majArray_dropmajEnd(int $n_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_dropmajEnd';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (count($xs_1) - $n_0);
  $__t1 = null;;
  if (($__local_var_2_0 < 1)) {
$__t1 = [];
goto end_branch_1;;
};
  $__t1 = ($GLOBALS['Data_Array_sliceImpl'])(0, $__local_var_2_0, $xs_1);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_dropEnd'] = __NAMESPACE__ . '\\majData_majArray_dropmajEnd';

// Data_Array_drop
function majData_majArray_drop(int $n_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_drop';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($n_0 < 1)) {
$__t0 = $xs_1;
goto end_branch_0;;
};
  $__t0 = ($GLOBALS['Data_Array_sliceImpl'])($n_0, count($xs_1), $xs_1);
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_drop'] = __NAMESPACE__ . '\\majData_majArray_drop';

// Data_Array_takeEnd
function majData_majArray_takemajEnd(int $n_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_takemajEnd';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = (count($xs_1) - $n_0);
  $__t1 = null;;
  if (($__local_var_2_0 < 1)) {
$__t1 = $xs_1;
goto end_branch_1;;
};
  $__t1 = ($GLOBALS['Data_Array_sliceImpl'])($__local_var_2_0, count($xs_1), $xs_1);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_takeEnd'] = __NAMESPACE__ . '\\majData_majArray_takemajEnd';

// Data_Array_deleteAt
function majData_majArray_deletemajAt(int $__local_var_0, $__local_var_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_deletemajAt';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array__deleteAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_deleteAt'] = __NAMESPACE__ . '\\majData_majArray_deletemajAt';

// Data_Array_deleteBy
function majData_majArray_deletemajBy($v_0, $v1_1 = null, $v2_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_deletemajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t4 = null;;
  switch (count($v2_2)) {
case 0:
$__t4 = [];
goto end_branch_4;;
break;
default:
;
break;
};
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), ($v_0)($v1_1), $v2_2);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = $v2_2;
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__local_var_4_2 = ($GLOBALS['Data_Array__deleteAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), ($__local_var_3_0)->{'value0'}, $v2_2);
$__t3 = null;;
if ($__local_var_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = ($__local_var_4_2)->{'value0'};
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__t4 = $__t1;
  end_branch_4:;
  $__res = $__t4;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_deleteBy'] = __NAMESPACE__ . '\\majData_majArray_deletemajBy';

// Data_Array_delete
function majData_majArray_delete($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_delete';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_deleteBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_delete'] = __NAMESPACE__ . '\\majData_majArray_delete';

// Data_Array_difference
function majData_majArray_difference($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_difference';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Foldable_foldrArray'])(($GLOBALS['Data_Array_deleteBy'])(($dictEq_0)->{'eq'}));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_difference'] = __NAMESPACE__ . '\\majData_majArray_difference';

// Data_Array_cons
function majData_majArray_cons($x_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_cons';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Semigroup\majData_majSemigroup_concatmajArray([$x_0], $xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_cons'] = __NAMESPACE__ . '\\majData_majArray_cons';

// Data_Array_some
function majData_majArray_some($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_some';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Apply0_1_0 = (((($dictAlternative_0)->{'Applicative0'})(null))->{'Apply0'})(null);
  $Functor0_2_1 = (((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null);
  $__res = function($dictLazy_3) use ($Apply0_1_0, $Functor0_2_1, $dictAlternative_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($Apply0_1_0, $Functor0_2_1, $dictAlternative_0, $dictLazy_3) {
  $__num = \func_num_args();
  $__res = ((($Apply0_1_0)->{'apply'})(((($Functor0_2_1)->{'map'})($GLOBALS['Data_Array_cons']))($v_4)))((($dictLazy_3)->{'defer'})(function($v1_5) use ($dictAlternative_0, $dictLazy_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Array_many'])($dictAlternative_0))($dictLazy_3))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_some'] = __NAMESPACE__ . '\\majData_majArray_some';

// Data_Array_many
function majData_majArray_many($dictAlternative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_many';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Alt0_1_0 = (((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null);
  $Applicative0_2_1 = (($dictAlternative_0)->{'Applicative0'})(null);
  $__res = function($dictLazy_3) use ($Alt0_1_0, $Applicative0_2_1, $dictAlternative_0) {
  $__num = \func_num_args();
  $__res = function($v_4) use ($Alt0_1_0, $Applicative0_2_1, $dictAlternative_0, $dictLazy_3) {
  $__num = \func_num_args();
  $__res = ((($Alt0_1_0)->{'alt'})(((((((($dictAlternative_0)->{'Applicative0'})(null))->{'Apply0'})(null))->{'apply'})(((((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Array_cons']))($v_4)))((($dictLazy_3)->{'defer'})(function($v1_5) use ($dictAlternative_0, $dictLazy_3, $v_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Array_many'])($dictAlternative_0))($dictLazy_3))($v_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((($Applicative0_2_1)->{'pure'})([]));
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_many'] = __NAMESPACE__ . '\\majData_majArray_many';

// Data_Array_concatMap
function majData_majArray_concatmajMap($b_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_concatmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Control\Bind\majControl_majBind_arraymajBind($a_1, $b_0);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_concatMap'] = __NAMESPACE__ . '\\majData_majArray_concatmajMap';

// Data_Array_mapMaybe
function majData_majArray_mapmajMaybe($f_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_mapmajMaybe';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_concatMap'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v2_1) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ($v2_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t0 = [];
goto end_branch_0;;
};
  if ($v2_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = [($v2_1)->{'value0'}];
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($f_0));
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_mapMaybe'] = __NAMESPACE__ . '\\majData_majArray_mapmajMaybe';

// Data_Array_filterA
function majData_majArray_filtermajA($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_filtermajA';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $Functor0_1_0 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($p_2) use ($Functor0_1_0, $dictApplicative_0) {
  $__num = \func_num_args();
  $Apply0_3_2 = (($dictApplicative_0)->{'Apply0'})(null);
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($Functor0_1_0)->{'map'})(($GLOBALS['Data_Array_mapMaybe'])(function($v_3) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($v_3)->{'value1'}) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(($v_3)->{'value0'});
goto end_branch_1;;
};
  $__t1 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))((((($GLOBALS['Data_Traversable_traverseArrayImpl'])(($Apply0_3_2)->{'apply'}))(((($Apply0_3_2)->{'Functor0'})(null))->{'map'}))(($dictApplicative_0)->{'pure'}))(function($x_4) use ($Functor0_1_0, $p_2) {
  $__num = \func_num_args();
  $__res = ((($Functor0_1_0)->{'map'})(($GLOBALS['Data_Tuple_Tuple'])($x_4)))(($p_2)($x_4));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_filterA'] = __NAMESPACE__ . '\\majData_majArray_filtermajA';

// Data_Array_catMaybes_closure
$GLOBALS['Data_Array_catMaybes_closure'] = ($GLOBALS['Data_Array_mapMaybe'])(function($x_0) {
  $__num = \func_num_args();
  $__res = $x_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Array_catMaybes
function majData_majArray_catmajMaybes($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_catmajMaybes';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_catMaybes_closure'])($v_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_catMaybes'] = __NAMESPACE__ . '\\majData_majArray_catmajMaybes';

// Data_Array_any
function majData_majArray_any($__local_var_0, $__local_var_1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_any';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_anyImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_any'] = __NAMESPACE__ . '\\majData_majArray_any';

// Data_Array_nubByEq
function majData_majArray_nubmajBymajEq($eq_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_nubmajBymajEq';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(function() use ($eq_0, $xs_1, &$__fn) {
$arr_2_0 = phpurs_execute_effect($GLOBALS['Data_Array_ST_new']);
$_dollar___unused_3_1 = phpurs_execute_effect(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_foreach($xs_1, function($x_3) use ($arr_2_0, $eq_0) {
  $__num = \func_num_args();
  $__local_var_4_1 = (($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_HeytingAlgebra_boolNot']))(($GLOBALS['Data_Array_any'])(function($v_4) use ($eq_0, $x_3) {
  $__num = \func_num_args();
  $__res = (($eq_0)($v_4))($x_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  $__local_var_5_2 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($arr_2_0));
  $e_5_2 = phpurs_execute_effect(($__local_var_4_1)($__local_var_5_2));
  $__t4 = null;;
  if ($e_5_2) {
$__local_var_6_5 = phpurs_execute_effect(($GLOBALS['Data_Array_ST_pushImpl'])($x_3, $arr_2_0));
$__t4 = $GLOBALS['Data_Unit_unit'];
goto end_branch_4;;
};
  $__t4 = $GLOBALS['Data_Unit_unit'];
  end_branch_4:;
  $__res = phpurs_execute_effect($__t4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
return phpurs_execute_effect(phpurs_execute_effect(($GLOBALS['Data_Array_ST_unsafeFreezeImpl'])($arr_2_0)));
});
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_nubByEq'] = __NAMESPACE__ . '\\majData_majArray_nubmajBymajEq';

// Data_Array_nubEq
function majData_majArray_nubmajEq($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_nubmajEq';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_nubByEq'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_nubEq'] = __NAMESPACE__ . '\\majData_majArray_nubmajEq';

// Data_Array_unionBy
function majData_majArray_unionmajBy($eq_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_unionmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Data\Semigroup\majData_majSemigroup_concatmajArray($xs_1, \Data\Foldable\majData_majFoldable_foldlmajArray(function($b_3) use ($eq_0) {
  $__num = \func_num_args();
  $__res = function($a_4) use ($b_3, $eq_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_deletemajBy($eq_0, $a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, \Data\Array\majData_majArray_nubmajBymajEq($eq_0, $ys_2), $xs_1));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_unionBy'] = __NAMESPACE__ . '\\majData_majArray_unionmajBy';

// Data_Array_union
function majData_majArray_union($dictEq_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_union';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ($GLOBALS['Data_Array_unionBy'])(($dictEq_0)->{'eq'});
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_union'] = __NAMESPACE__ . '\\majData_majArray_union';

// Data_Array_alterAt
function majData_majArray_altermajAt(int $i_0, $f_1 = null, $xs_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_altermajAt';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_2, $i_0);
  $__t1 = null;;
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t1 = new \Data\Maybe\Data_Maybe_Nothing();
goto end_branch_1;;
};
  if ($__local_var_3_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$v_4_2 = ($f_1)(($__local_var_3_0)->{'value0'});
$__t3 = null;;
if ($v_4_2 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t3 = ($GLOBALS['Data_Array__deleteAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $i_0, $xs_2);
goto end_branch_3;;
};
if ($v_4_2 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t3 = ($GLOBALS['Data_Array__updateAt'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $i_0, ($v_4_2)->{'value0'}, $xs_2);
goto end_branch_3;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t3 = null;
end_branch_3:;
$__t1 = $__t3;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_alterAt'] = __NAMESPACE__ . '\\majData_majArray_altermajAt';

// Data_Array_all
function majData_majArray_all($__local_var_0, $__local_var_1 = null): bool|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_all';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_allImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_all'] = __NAMESPACE__ . '\\majData_majArray_all';

