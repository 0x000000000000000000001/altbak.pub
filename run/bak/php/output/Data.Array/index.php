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

$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Array = \call_user_func(function() {
  $exports = [];
$rangeImpl = function($start, $end = null) use (&$rangeImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$rangeImpl) {

            return $rangeImpl(...\array_merge($__args, $more));
        };
    }
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

$replicateImpl = function($count, $value = null) use (&$replicateImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$replicateImpl) {

            return $replicateImpl(...\array_merge($__args, $more));
        };
    }
    if ($count < 1) return [];
    return array_fill(0, $count, $value);
};

$fromFoldableImpl = function($foldr, $xs = null) use (&$fromFoldableImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$fromFoldableImpl) {

            return $fromFoldableImpl(...\array_merge($__args, $more));
        };
    }
    
    $emptyList = new \stdClass();
    $curryCons = function($head, $tail = null) {
        if (\func_num_args() < 2) {
            return function($tail) use ($head) {
                $obj = new \stdClass();
                $obj->head = $head;
                $obj->tail = $tail;
                return $obj;
            };
        }
        $obj = new \stdClass();
        $obj->head = $head;
        $obj->tail = $tail;
        return $obj;
    };
    
    $listToArray = function($list) use ($emptyList) {
        $result = [];
        $xs = $list;
        while ($xs !== $emptyList && $xs !== null) {
            $result[] = $xs->head;
            $xs = $xs->tail;
        }
        return $result;
    };
    
    return $listToArray($foldr($curryCons)($emptyList)($xs));
};

$length = function($xs) use (&$length) {
    return \count($xs);
};

$unconsImpl = function($empty, $next = null, $xs = null) use (&$unconsImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$unconsImpl) {

            return $unconsImpl(...\array_merge($__args, $more));
        };
    }
    if (\count($xs) === 0) return $empty((object)[]);
    return $next($xs[0])(\array_slice($xs, 1));
};

$indexImpl = function($just, $nothing = null, $xs = null, $i = null) use (&$indexImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$indexImpl) {

            return $indexImpl(...\array_merge($__args, $more));
        };
    }
    return ($i < 0 || $i >= \count($xs)) ? $nothing : $just($xs[$i]);
};

$findMapImpl = function($nothing, $isJust = null, $f = null, $xs = null) use (&$findMapImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$findMapImpl) {

            return $findMapImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $x) {
        $result = $f($x);
        if ($isJust($result)) return $result;
    }
    return $nothing;
};

$findIndexImpl = function($just, $nothing = null, $f = null, $xs = null) use (&$findIndexImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$findIndexImpl) {

            return $findIndexImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $i => $x) {
        if ($f($x)) return $just($i);
    }
    return $nothing;
};

$findLastIndexImpl = function($just, $nothing = null, $f = null, $xs = null) use (&$findLastIndexImpl) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$findLastIndexImpl) {

            return $findLastIndexImpl(...\array_merge($__args, $more));
        };
    }
    for ($i = \count($xs) - 1; $i >= 0; $i--) {
        if ($f($xs[$i])) return $just($i);
    }
    return $nothing;
};

$_insertAt = function($just, $nothing = null, $i = null, $a = null, $l = null) use (&$_insertAt) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_insertAt) {

            return $_insertAt(...\array_merge($__args, $more));
        };
    }
    if ($i < 0 || $i > \count($l)) return $nothing;
    $l1 = $l;
    array_splice($l1, $i, 0, [$a]);
    return $just($l1);
};

$_deleteAt = function($just, $nothing = null, $i = null, $l = null) use (&$_deleteAt) {
    if (\func_num_args() < 4) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_deleteAt) {

            return $_deleteAt(...\array_merge($__args, $more));
        };
    }
    if ($i < 0 || $i >= \count($l)) return $nothing;
    $l1 = $l;
    array_splice($l1, $i, 1);
    return $just($l1);
};

$_updateAt = function($just, $nothing = null, $i = null, $a = null, $l = null) use (&$_updateAt) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$_updateAt) {

            return $_updateAt(...\array_merge($__args, $more));
        };
    }
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

$filterImpl = function($f, $xs = null) use (&$filterImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$filterImpl) {

            return $filterImpl(...\array_merge($__args, $more));
        };
    }
    $res = [];
    foreach ($xs as $x) {
        if ($f($x)) $res[] = $x;
    }
    return $res;
};

$partitionImpl = function($f, $xs = null) use (&$partitionImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$partitionImpl) {

            return $partitionImpl(...\array_merge($__args, $more));
        };
    }
    $yes = [];
    $no = [];
    foreach ($xs as $x) {
        if ($f($x)) $yes[] = $x;
        else $no[] = $x;
    }
    return (object)["yes" => $yes, "no" => $no];
};

$scanlImpl = function($f, $b = null, $xs = null) use (&$scanlImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$scanlImpl) {

            return $scanlImpl(...\array_merge($__args, $more));
        };
    }
    $acc = $b;
    $out = [];
    foreach ($xs as $x) {
        $acc = $f($acc)($x);
        $out[] = $acc;
    }
    return $out;
};

$scanrImpl = function($f, $b = null, $xs = null) use (&$scanrImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$scanrImpl) {

            return $scanrImpl(...\array_merge($__args, $more));
        };
    }
    $len = \count($xs);
    $acc = $b;
    $out = array_fill(0, $len, null);
    for ($i = $len - 1; $i >= 0; $i--) {
        $acc = $f($xs[$i])($acc);
        $out[$i] = $acc;
    }
    return $out;
};

$sortByImpl = function($compare, $fromOrdering = null, $xs = null) use (&$sortByImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$sortByImpl) {

            return $sortByImpl(...\array_merge($__args, $more));
        };
    }
    $out = $xs;
    \usort($out, function($a, $b) use ($compare, $fromOrdering) {
        return $fromOrdering($compare($a)($b));
    });
    return $out;
};

$sliceImpl = function($s, $e = null, $l = null) use (&$sliceImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$sliceImpl) {

            return $sliceImpl(...\array_merge($__args, $more));
        };
    }
    return \array_slice($l, $s, $e - $s);
};

$zipWithImpl = function($f, $xs = null, $ys = null) use (&$zipWithImpl) {
    if (\func_num_args() < 3) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$zipWithImpl) {

            return $zipWithImpl(...\array_merge($__args, $more));
        };
    }
    $l = \min(\count($xs), \count($ys));
    $result = [];
    for ($i = 0; $i < $l; $i++) {
        $result[] = $f($xs[$i])($ys[$i]);
    }
    return $result;
};

$anyImpl = function($p, $xs = null) use (&$anyImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$anyImpl) {

            return $anyImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $x) {
        if ($p($x)) return true;
    }
    return false;
};

$allImpl = function($p, $xs = null) use (&$allImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$allImpl) {

            return $allImpl(...\array_merge($__args, $more));
        };
    }
    foreach ($xs as $x) {
        if (!$p($x)) return false;
    }
    return true;
};

$unsafeIndexImpl = function($xs, $n = null) use (&$unsafeIndexImpl) {
    if (\func_num_args() < 2) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$unsafeIndexImpl) {

            return $unsafeIndexImpl(...\array_merge($__args, $more));
        };
    }
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
$GLOBALS['Data_Array__deleteAt'] = $ffi_Data_Array['_deleteAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array__insertAt'] = $ffi_Data_Array['_insertAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array__updateAt'] = $ffi_Data_Array['_updateAt'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_allImpl'] = $ffi_Data_Array['allImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_anyImpl'] = $ffi_Data_Array['anyImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_concat'] = $ffi_Data_Array['concat'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_filterImpl'] = $ffi_Data_Array['filterImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_findIndexImpl'] = $ffi_Data_Array['findIndexImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_findLastIndexImpl'] = $ffi_Data_Array['findLastIndexImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_findMapImpl'] = $ffi_Data_Array['findMapImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_fromFoldableImpl'] = $ffi_Data_Array['fromFoldableImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_indexImpl'] = $ffi_Data_Array['indexImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_length'] = $ffi_Data_Array['length'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_partitionImpl'] = $ffi_Data_Array['partitionImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_rangeImpl'] = $ffi_Data_Array['rangeImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_replicateImpl'] = $ffi_Data_Array['replicateImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_reverse'] = $ffi_Data_Array['reverse'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_scanlImpl'] = $ffi_Data_Array['scanlImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_scanrImpl'] = $ffi_Data_Array['scanrImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_sliceImpl'] = $ffi_Data_Array['sliceImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_sortByImpl'] = $ffi_Data_Array['sortByImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_unconsImpl'] = $ffi_Data_Array['unconsImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_unsafeIndexImpl'] = $ffi_Data_Array['unsafeIndexImpl'] ?? new class { public function __invoke(...$args) { return $this; } };
$GLOBALS['Data_Array_zipWithImpl'] = $ffi_Data_Array['zipWithImpl'] ?? new class { public function __invoke(...$args) { return $this; } };


// Data_Array_traverse_
$GLOBALS['Data_Array_traverse_'] = ($GLOBALS['Data_Foldable_traverse_'])($GLOBALS['Control_Monad_ST_Internal_applicativeST']);

// Data_Array_void
$GLOBALS['Data_Array_void'] = ($GLOBALS['Control_Monad_ST_Internal_map_'])(function($v_0 = null) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Array_fromJust
$GLOBALS['Data_Array_fromJust'] = function($v_0 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_0) && (($v_0)->{'tag'} === "Just"))) {
$__t0 = ($v_0)->{'value0'};
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_zipWith
$GLOBALS['Data_Array_zipWith'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_zipWithImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_zipWithA
$GLOBALS['Data_Array_zipWithA'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $sequence1_1_0 = ((($GLOBALS['Data_Traversable_traversableArray'])['traverse'])($dictApplicative_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2 = null, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($sequence1_1_0)(($GLOBALS['Data_Array_zipWithImpl'])($f_2, $xs_3, $ys_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_zip
$GLOBALS['Data_Array_zip'] = ($GLOBALS['Data_Array_zipWith'])($GLOBALS['Data_Tuple_Tuple']);

// Data_Array_updateAtIndices
$GLOBALS['Data_Array_updateAtIndices'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $traverse_1_1_0 = ($GLOBALS['Data_Array_traverse_'])($dictFoldable_0);
  $__res = (function() use ($traverse_1_1_0) {
  $__fn = function($us_2 = null, $xs_3 = null) use ($traverse_1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Data_Array_ST_withArray'])(function($res_4 = null) use ($traverse_1_1_0, $us_2) {
  $__num = \func_num_args();
  $__res = (($traverse_1_1_0)(function($v_5 = null) use ($res_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Array_ST_poke'])(($v_5)->{'value0'}))(($v_5)->{'value1'}))($res_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($us_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_updateAt
$GLOBALS['Data_Array_updateAt'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array__updateAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_unsafeIndex
$GLOBALS['Data_Array_unsafeIndex'] = (function() {
  $__fn = function($dollar__unused_0 = null, $__local_var_1 = null, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($__local_var_1)[0];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_uncons
$GLOBALS['Data_Array_uncons'] = function($__local_var_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (function() {
  $__fn = function($x_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data1("Just", ["head" => $x_1, "tail" => $xs_2]);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), $__local_var_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_toUnfoldable
$GLOBALS['Data_Array_toUnfoldable'] = (function() {
  $__fn = function($dictUnfoldable_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $len_2_0 = count($xs_1);
  $__res = ((($dictUnfoldable_0)['unfoldr'])(function($i_3 = null) use ($len_2_0, $xs_1) {
  $__num = \func_num_args();
  $__t1 = null;;
  if (($i_3 < $len_2_0)) {
$__t1 = new Phpurs_Data1("Just", new Phpurs_Data2("Tuple", ($xs_1)[0], ($i_3 + 1)));
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_tail
$GLOBALS['Data_Array_tail'] = function($__local_var_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_1 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("Nothing");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (function() {
  $__fn = function($v_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data1("Just", $xs_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), $__local_var_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_sortBy
$GLOBALS['Data_Array_sortBy'] = (function() {
  $__fn = function($comp_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_sortByImpl'])($comp_0, function($v_2 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v_2) && (($v_2)->{'tag'} === "GT"))) {
$__t0 = 1;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "EQ"))) {
$__t0 = 0;
goto end_branch_0;;
};
  if ((is_object($v_2) && (($v_2)->{'tag'} === "LT"))) {
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_sortWith
$GLOBALS['Data_Array_sortWith'] = (function() {
  $__fn = function($dictOrd_0 = null, $f_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_sortBy'])((function() use ($dictOrd_0, $f_1) {
  $__fn = function($x_2 = null, $y_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictOrd_0)['compare'])(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_sort
$GLOBALS['Data_Array_sort'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $compare_1_0 = ($dictOrd_0)['compare'];
  $__res = function($xs_2 = null) use ($compare_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_sortBy'])($compare_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_snoc
$GLOBALS['Data_Array_snoc'] = (function() {
  $__fn = function($xs_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Data_Array_ST_withArray'])(($GLOBALS['Data_Array_ST_push'])($x_1)))($xs_0));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_slice
$GLOBALS['Data_Array_slice'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_sliceImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_splitAt
$GLOBALS['Data_Array_splitAt'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t0 = null;;
  if (($v_0 <= 0)) {
$__t0 = ["before" => [], "after" => $v1_1];
goto end_branch_0;;
};
  $__t0 = ["before" => ($GLOBALS['Data_Array_sliceImpl'])(0, $v_0, $v1_1), "after" => ($GLOBALS['Data_Array_sliceImpl'])($v_0, count($v1_1), $v1_1)];
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_take
$GLOBALS['Data_Array_take'] = (function() {
  $__fn = function($n_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_singleton
$GLOBALS['Data_Array_singleton'] = function($a_0 = null) {
  $__num = \func_num_args();
  $__res = [$a_0];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_scanr
$GLOBALS['Data_Array_scanr'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_scanrImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_scanl
$GLOBALS['Data_Array_scanl'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_scanlImpl'])($__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_replicate
$GLOBALS['Data_Array_replicate'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_replicateImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_range
$GLOBALS['Data_Array_range'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_rangeImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_partition
$GLOBALS['Data_Array_partition'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_partitionImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_null
$GLOBALS['Data_Array_null'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__res = (count($xs_0) === 0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_modifyAtIndices
$GLOBALS['Data_Array_modifyAtIndices'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $traverse_1_1_0 = ($GLOBALS['Data_Array_traverse_'])($dictFoldable_0);
  $__res = (function() use ($traverse_1_1_0) {
  $__fn = function($is_2 = null, $f_3 = null, $xs_4 = null) use ($traverse_1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Data_Array_ST_withArray'])(function($res_5 = null) use ($f_3, $is_2, $traverse_1_1_0) {
  $__num = \func_num_args();
  $__res = (($traverse_1_1_0)(function($i_6 = null) use ($f_3, $res_5) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Array_ST_modify'])($i_6))($f_3))($res_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($is_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($xs_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_mapWithIndex
$GLOBALS['Data_Array_mapWithIndex'] = $GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'];

// Data_Array_intersperse
$GLOBALS['Data_Array_intersperse'] = (function() {
  $__fn = function($a_0 = null, $arr_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $v_2_0 = count($arr_1);
  $__t1 = null;;
  if (($v_2_0 < 2)) {
$__t1 = $arr_1;
goto end_branch_1;;
};
  $__t1 = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Control_Monad_ST_Internal_bind_'])($GLOBALS['Data_Array_ST_new']))(function($out_3 = null) use ($a_0, $arr_1, $v_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Data_Array_ST_push'])(($arr_1)[0]))($out_3)))(function($dollar__unused_4 = null) use ($a_0, $arr_1, $out_3, $v_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])(((($GLOBALS['Control_Monad_ST_Internal_for'])(1))($v_2_0))(function($idx_5 = null) use ($a_0, $arr_1, $out_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Data_Array_ST_push'])($a_0))($out_3)))(function($dollar__unused_6 = null) use ($arr_1, $idx_5, $out_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_void'])((($GLOBALS['Data_Array_ST_push'])(($arr_1)[0]))($out_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_5 = null) use ($out_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_ST_Internal_pure_'])($out_3);
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
})))($GLOBALS['Data_Array_ST_unsafeFreeze']));
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_intercalate
$GLOBALS['Data_Array_intercalate'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_intercalate'])($GLOBALS['Data_Foldable_foldableArray']))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_insertAt
$GLOBALS['Data_Array_insertAt'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null, $__local_var_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array__insertAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $__local_var_0, $__local_var_1, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_init
$GLOBALS['Data_Array_init'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__res = match (count($xs_0)) { 0 => new Phpurs_Data0("Nothing"), default => new Phpurs_Data1("Just", ($GLOBALS['Data_Array_sliceImpl'])(0, (count($xs_0) - 1), $xs_0)) };
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_index
$GLOBALS['Data_Array_index'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_last
$GLOBALS['Data_Array_last'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $xs_0, (count($xs_0) - 1));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_unsnoc
$GLOBALS['Data_Array_unsnoc'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__t2 = null;;
  switch (count($xs_0)) {
case 0:
$__t2 = new Phpurs_Data0("Nothing");
goto end_branch_2;;
break;
default:
;
break;
};
  $__local_var_1_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $xs_0, (count($xs_0) - 1));
  $__t1 = null;;
  if ((is_object($__local_var_1_0) && (($__local_var_1_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ["init" => ($GLOBALS['Data_Array_sliceImpl'])(0, (count($xs_0) - 1), $xs_0), "last" => ($__local_var_1_0)->{'value0'}]);
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__t2 = $__t1;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_modifyAt
$GLOBALS['Data_Array_modifyAt'] = (function() {
  $__fn = function($i_0 = null, $f_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $xs_2, $i_0);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__t1 = ($GLOBALS['Data_Array__updateAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $i_0, ($f_1)(($__local_var_3_0)->{'value0'}), $xs_2);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_span
$GLOBALS['Data_Array_span'] = (function() {
  $__fn = function($p_0 = null, $arr_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $go_2_0 = null;
  $go_2_0 = function($i_3 = null) use ($arr_1, &$go_2_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go_2_0_0_i_3 = $i_3;
  tco_loop_go_2_0_0:;
  $i_3 = $__tco_var_go_2_0_0_i_3;
  $v_4_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $arr_1, $i_3);
  $__t1 = null;;
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Just"))) {
$__t2 = null;;
if (($p_0)(($v_4_0)->{'value0'})) {
$__tco_3 = ($i_3 + 1);
$__tco_var_go_2_0_0_i_3 = $__tco_3;
goto tco_loop_go_2_0_0;;
$__t2 = null;
goto end_branch_2;;
};
$__t2 = new Phpurs_Data1("Just", $i_3);
end_branch_2:;
$__t1 = $__t2;
goto end_branch_1;;
};
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
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
  $breakIndex_3_1 = ($go_2_0)(0);
  $__t2 = null;;
  if ((is_object($breakIndex_3_1) && (($breakIndex_3_1)->{'tag'} === "Just"))) {
$__t2 = match (($breakIndex_3_1)->{'value0'}) { 0 => ["init" => [], "rest" => $arr_1], default => ["init" => ($GLOBALS['Data_Array_sliceImpl'])(0, ($breakIndex_3_1)->{'value0'}, $arr_1), "rest" => ($GLOBALS['Data_Array_sliceImpl'])(($breakIndex_3_1)->{'value0'}, count($arr_1), $arr_1)] };
goto end_branch_2;;
};
  if ((is_object($breakIndex_3_1) && (($breakIndex_3_1)->{'tag'} === "Nothing"))) {
$__t2 = ["init" => $arr_1, "rest" => []];
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_takeWhile
$GLOBALS['Data_Array_takeWhile'] = (function() {
  $__fn = function($p_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Array_span'])($p_0))($xs_1))['init'];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_unzip
$GLOBALS['Data_Array_unzip'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Control_Monad_ST_Internal_bind_'])($GLOBALS['Data_Array_ST_new']))(function($fsts_1 = null) use ($xs_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])($GLOBALS['Data_Array_ST_new']))(function($snds_2 = null) use ($fsts_1, $xs_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Control_Monad_ST_Internal_map_'])(($GLOBALS['Data_Array_ST_Iterator_Iterator'])(function($v_3 = null) use ($xs_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $xs_0, $v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Control_Monad_ST_Internal_new'])(0))))(function($iter_3 = null) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Data_Array_ST_Iterator_iterate'])($iter_3))(function($v_4 = null) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  $__local_var_5_0 = ($v_4)->{'value1'};
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])(($GLOBALS['Data_Array_void'])((($GLOBALS['Data_Array_ST_push'])(($v_4)->{'value0'}))($fsts_1))))(function($dollar__unused_6 = null) use ($__local_var_5_0, $snds_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_void'])((($GLOBALS['Data_Array_ST_push'])($__local_var_5_0))($snds_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_4 = null) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])(($GLOBALS['Data_Array_ST_unsafeFreeze'])($fsts_1)))(function($fsts__prime___5 = null) use ($snds_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])(($GLOBALS['Data_Array_ST_unsafeFreeze'])($snds_2)))(function($snds__prime___6 = null) use ($fsts__prime___5) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Control_Monad_ST_Internal_pure_'])(new Phpurs_Data2("Tuple", $fsts__prime___5, $snds__prime___6));
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_head
$GLOBALS['Data_Array_head'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $xs_0, 0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_nubBy
$GLOBALS['Data_Array_nubBy'] = (function() {
  $__fn = function($comp_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $indexedAndSorted_2_0 = (($GLOBALS['Data_Array_sortBy'])((function() use ($comp_0) {
  $__fn = function($x_2 = null, $y_3 = null) use ($comp_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($comp_0)(($x_2)->{'value1'}))(($y_3)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($GLOBALS['Data_FunctorWithIndex_mapWithIndexArray'])($GLOBALS['Data_Tuple_Tuple']))($xs_1));
  $v_3_1 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $indexedAndSorted_2_0, 0);
  $__t2 = null;;
  if ((is_object($v_3_1) && (($v_3_1)->{'tag'} === "Nothing"))) {
$__t2 = [];
goto end_branch_2;;
};
  if ((is_object($v_3_1) && (($v_3_1)->{'tag'} === "Just"))) {
$__t2 = (($GLOBALS['Data_Functor_arrayMap'])($GLOBALS['Data_Tuple_snd']))(((($GLOBALS['Data_Array_sortWith'])($GLOBALS['Data_Ord_ordInt']))($GLOBALS['Data_Tuple_fst']))(($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Control_Monad_ST_Internal_bind_'])(($GLOBALS['Data_Array_ST_unsafeThaw'])([($v_3_1)->{'value0'}])))(function($result_4 = null) use ($comp_0, $indexedAndSorted_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Control_Monad_ST_Internal_foreach'])($indexedAndSorted_2_0))(function($v1_5 = null) use ($comp_0, $result_4) {
  $__num = \func_num_args();
  $__local_var_6_3 = ($v1_5)->{'value1'};
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Control_Monad_ST_Internal_map_'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_fromJust']))($GLOBALS['Data_Array_last']))))(($GLOBALS['Data_Array_ST_unsafeFreeze'])($result_4))))(function($lst_7 = null) use ($__local_var_6_3, $comp_0, $result_4, $v1_5) {
  $__num = \func_num_args();
  $__local_var_8_4 = (($comp_0)($lst_7))($__local_var_6_3);
  $__local_var_9_5 = ($GLOBALS['Data_Array_void'])((($GLOBALS['Data_Array_ST_push'])($v1_5))($result_4));
  $__t6 = null;;
  if (((is_object($__local_var_8_4) && (($__local_var_8_4)->{'tag'} === "LT")) || ((is_object($__local_var_8_4) && (($__local_var_8_4)->{'tag'} === "GT")) || ( ! (is_object($__local_var_8_4) && (($__local_var_8_4)->{'tag'} === "EQ")))))) {
$__t6 = $__local_var_9_5;
goto end_branch_6;;
};
  $__t6 = ($GLOBALS['Control_Monad_ST_Internal_pure_'])($GLOBALS['Data_Unit_unit']);
  end_branch_6:;
  $__res = $__t6;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_5 = null) use ($result_4) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_ST_unsafeFreeze'])($result_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))));
goto end_branch_2;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t2 = null;
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_nub
$GLOBALS['Data_Array_nub'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_nubBy'])(($dictOrd_0)['compare']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_groupBy
$GLOBALS['Data_Array_groupBy'] = (function() {
  $__fn = function($op_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Control_Monad_ST_Internal_bind_'])($GLOBALS['Data_Array_ST_new']))(function($result_2 = null) use ($op_0, $xs_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Control_Monad_ST_Internal_map_'])(($GLOBALS['Data_Array_ST_Iterator_Iterator'])(function($v_3 = null) use ($xs_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $xs_1, $v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(($GLOBALS['Control_Monad_ST_Internal_new'])(0))))(function($iter_3 = null) use ($op_0, $result_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Data_Array_ST_Iterator_iterate'])($iter_3))(function($x_4 = null) use ($iter_3, $op_0, $result_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_void'])((($GLOBALS['Control_Monad_ST_Internal_bind_'])($GLOBALS['Data_Array_ST_new']))(function($sub1_5 = null) use ($iter_3, $op_0, $result_2, $x_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Data_Array_ST_push'])($x_4))($sub1_5)))(function($dollar__unused_6 = null) use ($iter_3, $op_0, $result_2, $sub1_5, $x_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])(((($GLOBALS['Data_Array_ST_Iterator_pushWhile'])(($op_0)($x_4)))($iter_3))($sub1_5)))(function($dollar__unused_7 = null) use ($result_2, $sub1_5) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])(($GLOBALS['Data_Array_ST_unsafeFreeze'])($sub1_5)))(function($grp_8 = null) use ($result_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_ST_push'])($grp_8))($result_2);
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_4 = null) use ($result_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_ST_unsafeFreeze'])($result_2);
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
}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_groupAllBy
$GLOBALS['Data_Array_groupAllBy'] = function($cmp_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupBy'])((function() use ($cmp_0) {
  $__fn = function($x_1 = null, $y_2 = null) use ($cmp_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (is_object((($cmp_0)($x_1))($y_2)) && (((($cmp_0)($x_1))($y_2))->{'tag'} === "EQ"));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))(($GLOBALS['Data_Array_sortBy'])($cmp_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_groupAll
$GLOBALS['Data_Array_groupAll'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_groupAllBy'])(($dictOrd_0)['compare']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_group
$GLOBALS['Data_Array_group'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $eq2_1_0 = ($dictEq_0)['eq'];
  $__res = function($xs_2 = null) use ($eq2_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_groupBy'])($eq2_1_0))($xs_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_fromFoldable
$GLOBALS['Data_Array_fromFoldable'] = function($dictFoldable_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = ($dictFoldable_0)['foldr'];
  $__res = function($__local_var_2 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_fromFoldableImpl'])($__local_var_1_0, $__local_var_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_foldr
$GLOBALS['Data_Array_foldr'] = $GLOBALS['Data_Foldable_foldrArray'];

// Data_Array_foldl
$GLOBALS['Data_Array_foldl'] = $GLOBALS['Data_Foldable_foldlArray'];

// Data_Array_transpose
$GLOBALS['Data_Array_transpose'] = function($xs_0 = null) {
  $__num = \func_num_args();
  $go_1_0 = null;
  $go_1_0 = (function() use (&$go_1_0, $xs_0) {
  $__fn = function($idx_2 = null, $allArrays_3 = null) use (&$go_1_0, $xs_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go_1_0_0_idx_2 = $idx_2;
  $__tco_var_go_1_0_0_allArrays_3 = $allArrays_3;
  tco_loop_go_1_0_0:;
  $idx_2 = $__tco_var_go_1_0_0_idx_2;
  $allArrays_3 = $__tco_var_go_1_0_0_allArrays_3;
  $v_4_0 = ((($GLOBALS['Data_Foldable_foldlArray'])((function() use ($idx_2) {
  $__fn = function($acc_4 = null, $nextArr_5 = null) use ($idx_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_6_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $nextArr_5, $idx_2);
  $__t1 = null;;
  if ((is_object($__local_var_6_0) && (($__local_var_6_0)->{'tag'} === "Nothing"))) {
$__t1 = $acc_4;
goto end_branch_1;;
};
  if ((is_object($__local_var_6_0) && (($__local_var_6_0)->{'tag'} === "Just"))) {
$__t2 = null;;
if ((is_object($acc_4) && (($acc_4)->{'tag'} === "Nothing"))) {
$__t2 = [($__local_var_6_0)->{'value0'}];
goto end_branch_2;;
};
if ((is_object($acc_4) && (($acc_4)->{'tag'} === "Just"))) {
$__t2 = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Data_Array_ST_withArray'])(($GLOBALS['Data_Array_ST_push'])(($__local_var_6_0)->{'value0'})))(($acc_4)->{'value0'}));
goto end_branch_2;;
};
throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
$__t2 = null;
end_branch_2:;
$__t1 = new Phpurs_Data1("Just", $__t2);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new Phpurs_Data0("Nothing")))($xs_0);
  $__t4 = null;;
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Nothing"))) {
$__t4 = $allArrays_3;
goto end_branch_4;;
};
  if ((is_object($v_4_0) && (($v_4_0)->{'tag'} === "Just"))) {
$__tco_5 = ($idx_2 + 1);
$__tco_6 = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Data_Array_ST_withArray'])(($GLOBALS['Data_Array_ST_push'])(($v_4_0)->{'value0'})))($allArrays_3));
$__tco_var_go_1_0_0_idx_2 = $__tco_5;
$__tco_var_go_1_0_0_allArrays_3 = $__tco_6;
goto tco_loop_go_1_0_0;;
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
  $__res = (($go_1_0)(0))([]);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_foldRecM
$GLOBALS['Data_Array_foldRecM'] = function($dictMonadRec_0 = null) {
  $__num = \func_num_args();
  $Monad0_1_0 = (($dictMonadRec_0)['Monad0'])($GLOBALS['Prim_undefined']);
  $__local_var_2_1 = (($Monad0_1_0)['Applicative0'])($GLOBALS['Prim_undefined']);
  $__res = (function() use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0) {
  $__fn = function($f_3 = null, $b_4 = null, $array_5 = null) use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)['tailRecM'])(function($o_6 = null) use ($Monad0_1_0, $__local_var_2_1, $array_5, $f_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((($o_6)['b'] >= count($array_5))) {
$__t2 = (($__local_var_2_1)['pure'])(new Phpurs_Data1("Done", ($o_6)['a']));
goto end_branch_2;;
};
  $__t2 = ((((($Monad0_1_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])((($f_3)(($o_6)['a']))(($array_5)[0])))(function($res__prime___7 = null) use ($__local_var_2_1, $o_6) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)['pure'])(new Phpurs_Data1("Loop", ["a" => $res__prime___7, "b" => (($o_6)['b'] + 1)]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(["a" => $b_4, "b" => 0]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_foldMap
$GLOBALS['Data_Array_foldMap'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Foldable_foldMapDefaultR'])($GLOBALS['Data_Foldable_foldableArray']))($dictMonoid_0);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_foldM
$GLOBALS['Data_Array_foldM'] = (function() {
  $__fn = function($dictMonad_0 = null, $f_1 = null, $b_2 = null, $__local_var_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__tco_var_Data_Array_foldM_dictMonad_0 = $dictMonad_0;
  $__tco_var_Data_Array_foldM_f_1 = $f_1;
  $__tco_var_Data_Array_foldM_b_2 = $b_2;
  $__tco_var_Data_Array_foldM___local_var_3 = $__local_var_3;
  tco_loop_Data_Array_foldM:;
  $dictMonad_0 = $__tco_var_Data_Array_foldM_dictMonad_0;
  $f_1 = $__tco_var_Data_Array_foldM_f_1;
  $b_2 = $__tco_var_Data_Array_foldM_b_2;
  $__local_var_3 = $__tco_var_Data_Array_foldM___local_var_3;
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_4 = null) use ($b_2, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])($b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (function() use ($b_2, $dictMonad_0, $f_1) {
  $__fn = function($a_4 = null, $as_5 = null) use ($b_2, $dictMonad_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonad_0)['Bind1'])($GLOBALS['Prim_undefined']))['bind'])((($f_1)($b_2))($a_4)))(function($b__prime___6 = null) use ($as_5, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Array_foldM'])($dictMonad_0))($f_1))($b__prime___6))($as_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), $__local_var_3);
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})();

// Data_Array_fold
$GLOBALS['Data_Array_fold'] = function($dictMonoid_0 = null) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Foldable_foldMapDefaultR'])($GLOBALS['Data_Foldable_foldableArray']))($dictMonoid_0))(($GLOBALS['Control_Category_categoryFn'])['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_findMap
$GLOBALS['Data_Array_findMap'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findMapImpl'])(new Phpurs_Data0("Nothing"), $GLOBALS['Data_Maybe_isJust'], $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_findLastIndex
$GLOBALS['Data_Array_findLastIndex'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findLastIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_insertBy
$GLOBALS['Data_Array_insertBy'] = (function() {
  $__fn = function($cmp_0 = null, $x_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_findLastIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), function($y_3 = null) use ($cmp_0, $x_1) {
  $__num = \func_num_args();
  $__res = (is_object((($cmp_0)($x_1))($y_3)) && (((($cmp_0)($x_1))($y_3))->{'tag'} === "GT"));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $ys_2);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nothing"))) {
$__t1 = 0;
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__t1 = (($__local_var_3_0)->{'value0'} + 1);
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__local_var_3_0 = ($GLOBALS['Data_Array__insertAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $__t1, $x_1, $ys_2);
  $__t3 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__t3 = ($__local_var_3_0)->{'value0'};
goto end_branch_3;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t3 = null;
  end_branch_3:;
  $__res = $__t3;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_insert
$GLOBALS['Data_Array_insert'] = function($dictOrd_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_insertBy'])(($dictOrd_0)['compare']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_findIndex
$GLOBALS['Data_Array_findIndex'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_find
$GLOBALS['Data_Array_find'] = (function() {
  $__fn = function($f_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_2_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $f_0, $xs_1);
  $__t1 = null;;
  if ((is_object($__local_var_2_0) && (($__local_var_2_0)->{'tag'} === "Just"))) {
$__t1 = new Phpurs_Data1("Just", ($xs_1)[0]);
goto end_branch_1;;
};
  $__t1 = new Phpurs_Data0("Nothing");
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_filter
$GLOBALS['Data_Array_filter'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_filterImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_intersectBy
$GLOBALS['Data_Array_intersectBy'] = (function() {
  $__fn = function($eq2_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_filterImpl'])(function($x_3 = null) use ($eq2_0, $ys_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), ($eq2_0)($x_3), $ys_2);
  $__t1 = null;;
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Nothing"))) {
$__t1 = false;
goto end_branch_1;;
};
  if ((is_object($__local_var_4_0) && (($__local_var_4_0)->{'tag'} === "Just"))) {
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_intersect
$GLOBALS['Data_Array_intersect'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_intersectBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_elemLastIndex
$GLOBALS['Data_Array_elemLastIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findLastIndex'])(function($v_2 = null) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)['eq'])($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_elemIndex
$GLOBALS['Data_Array_elemIndex'] = (function() {
  $__fn = function($dictEq_0 = null, $x_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_findIndex'])(function($v_2 = null) use ($dictEq_0, $x_1) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)['eq'])($v_2))($x_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_notElem
$GLOBALS['Data_Array_notElem'] = (function() {
  $__fn = function($dictEq_0 = null, $a_1 = null, $arr_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), function($v_3 = null) use ($a_1, $dictEq_0) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)['eq'])($v_3))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $arr_2);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nothing"))) {
$__t1 = true;
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__t1 = false;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_elem
$GLOBALS['Data_Array_elem'] = (function() {
  $__fn = function($dictEq_0 = null, $a_1 = null, $arr_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), function($v_3 = null) use ($a_1, $dictEq_0) {
  $__num = \func_num_args();
  $__res = ((($dictEq_0)['eq'])($v_3))($a_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $arr_2);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nothing"))) {
$__t1 = false;
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__t1 = true;
goto end_branch_1;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t1 = null;
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_dropWhile
$GLOBALS['Data_Array_dropWhile'] = (function() {
  $__fn = function($p_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Array_span'])($p_0))($xs_1))['rest'];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_dropEnd
$GLOBALS['Data_Array_dropEnd'] = (function() {
  $__fn = function($n_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_drop
$GLOBALS['Data_Array_drop'] = (function() {
  $__fn = function($n_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_takeEnd
$GLOBALS['Data_Array_takeEnd'] = (function() {
  $__fn = function($n_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_deleteAt
$GLOBALS['Data_Array_deleteAt'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array__deleteAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_deleteBy
$GLOBALS['Data_Array_deleteBy'] = (function() {
  $__fn = function($v_0 = null, $v1_1 = null, $v2_2 = null) use (&$__fn) {
  $__num = \func_num_args();
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
  $__local_var_3_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), ($v_0)($v1_1), $v2_2);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nothing"))) {
$__t1 = $v2_2;
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$__local_var_4_2 = ($GLOBALS['Data_Array__deleteAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), ($__local_var_3_0)->{'value0'}, $v2_2);
$__t3 = null;;
if ((is_object($__local_var_4_2) && (($__local_var_4_2)->{'tag'} === "Just"))) {
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_delete
$GLOBALS['Data_Array_delete'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_deleteBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_difference
$GLOBALS['Data_Array_difference'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldrArray'])(($GLOBALS['Data_Array_delete'])($dictEq_0));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_cons
$GLOBALS['Data_Array_cons'] = (function() {
  $__fn = function($x_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Data_Semigroup_concatArray'])([$x_0]))($xs_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_some
$GLOBALS['Data_Array_some'] = (function() {
  $__fn = function($dictAlternative_0 = null, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)['Applicative0'])($GLOBALS['Prim_undefined']))['Apply0'])($GLOBALS['Prim_undefined']))['apply'])(((((((((($dictAlternative_0)['Plus1'])($GLOBALS['Prim_undefined']))['Alt0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']))['map'])($GLOBALS['Data_Array_cons']))($v_2)))((($dictLazy_1)['defer'])(function($v1_3 = null) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Data_Array_many'])($dictAlternative_0))($dictLazy_1))($v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_many
$GLOBALS['Data_Array_many'] = (function() {
  $__fn = function($dictAlternative_0 = null, $dictLazy_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)['Plus1'])($GLOBALS['Prim_undefined']))['Alt0'])($GLOBALS['Prim_undefined']))['alt'])(((($GLOBALS['Data_Array_some'])($dictAlternative_0))($dictLazy_1))($v_2)))((((($dictAlternative_0)['Applicative0'])($GLOBALS['Prim_undefined']))['pure'])([]));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_concatMap
$GLOBALS['Data_Array_concatMap'] = (function() {
  $__fn = function($b_0 = null, $a_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($GLOBALS['Control_Bind_arrayBind'])($a_1))($b_0);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_mapMaybe
$GLOBALS['Data_Array_mapMaybe'] = function($f_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_concatMap'])((($GLOBALS['Control_Semigroupoid_composeImpl'])(function($v2_1 = null) {
  $__num = \func_num_args();
  $__t0 = null;;
  if ((is_object($v2_1) && (($v2_1)->{'tag'} === "Nothing"))) {
$__t0 = [];
goto end_branch_0;;
};
  if ((is_object($v2_1) && (($v2_1)->{'tag'} === "Just"))) {
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
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_filterA
$GLOBALS['Data_Array_filterA'] = function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $traverse1_1_0 = (($GLOBALS['Data_Traversable_traversableArray'])['traverse'])($dictApplicative_0);
  $__local_var_2_1 = (((($dictApplicative_0)['Apply0'])($GLOBALS['Prim_undefined']))['Functor0'])($GLOBALS['Prim_undefined']);
  $__res = function($p_3 = null) use ($__local_var_2_1, $traverse1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_2_1)['map'])(($GLOBALS['Data_Array_mapMaybe'])(function($v_4 = null) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($v_4)->{'value1'}) {
$__t2 = new Phpurs_Data1("Just", ($v_4)->{'value0'});
goto end_branch_2;;
};
  $__t2 = new Phpurs_Data0("Nothing");
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(($traverse1_1_0)(function($x_4 = null) use ($__local_var_2_1, $p_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)['map'])(($GLOBALS['Data_Tuple_Tuple'])($x_4)))(($p_3)($x_4));
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

// Data_Array_catMaybes
$GLOBALS['Data_Array_catMaybes'] = ($GLOBALS['Data_Array_mapMaybe'])(($GLOBALS['Control_Category_categoryFn'])['identity']);

// Data_Array_any
$GLOBALS['Data_Array_any'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_anyImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_nubByEq
$GLOBALS['Data_Array_nubByEq'] = (function() {
  $__fn = function($eq2_0 = null, $xs_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Control_Monad_ST_Internal_run'])((($GLOBALS['Control_Monad_ST_Internal_bind_'])($GLOBALS['Data_Array_ST_new']))(function($arr_2 = null) use ($eq2_0, $xs_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Control_Monad_ST_Internal_foreach'])($xs_1))(function($x_3 = null) use ($arr_2, $eq2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_bind_'])((($GLOBALS['Control_Monad_ST_Internal_map_'])((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_HeytingAlgebra_boolNot']))(($GLOBALS['Data_Array_any'])(function($v_4 = null) use ($eq2_0, $x_3) {
  $__num = \func_num_args();
  $__res = (($eq2_0)($v_4))($x_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(($GLOBALS['Data_Array_ST_unsafeFreeze'])($arr_2))))(function($e_4 = null) use ($arr_2, $x_3) {
  $__num = \func_num_args();
  $__local_var_5_0 = ($GLOBALS['Data_Array_void'])((($GLOBALS['Data_Array_ST_push'])($x_3))($arr_2));
  $__t1 = null;;
  if ($e_4) {
$__t1 = $__local_var_5_0;
goto end_branch_1;;
};
  $__t1 = ($GLOBALS['Control_Monad_ST_Internal_pure_'])($GLOBALS['Data_Unit_unit']);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($dollar__unused_3 = null) use ($arr_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_ST_unsafeFreeze'])($arr_2);
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

// Data_Array_nubEq
$GLOBALS['Data_Array_nubEq'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_nubByEq'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_unionBy
$GLOBALS['Data_Array_unionBy'] = (function() {
  $__fn = function($eq2_0 = null, $xs_1 = null, $ys_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($GLOBALS['Data_Semigroup_concatArray'])($xs_1))(((($GLOBALS['Data_Foldable_foldlArray'])((function() use ($eq2_0) {
  $__fn = function($b_3 = null, $a_4 = null) use ($eq2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Array_deleteBy'])($eq2_0))($a_4))($b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((($GLOBALS['Data_Array_nubByEq'])($eq2_0))($ys_2)))($xs_1));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_union
$GLOBALS['Data_Array_union'] = function($dictEq_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_unionBy'])(($dictEq_0)['eq']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};

// Data_Array_alterAt
$GLOBALS['Data_Array_alterAt'] = (function() {
  $__fn = function($i_0 = null, $f_1 = null, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_3_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $xs_2, $i_0);
  $__t1 = null;;
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Nothing"))) {
$__t1 = new Phpurs_Data0("Nothing");
goto end_branch_1;;
};
  if ((is_object($__local_var_3_0) && (($__local_var_3_0)->{'tag'} === "Just"))) {
$v_4_2 = ($f_1)(($__local_var_3_0)->{'value0'});
$__t3 = null;;
if ((is_object($v_4_2) && (($v_4_2)->{'tag'} === "Nothing"))) {
$__t3 = ($GLOBALS['Data_Array__deleteAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $i_0, $xs_2);
goto end_branch_3;;
};
if ((is_object($v_4_2) && (($v_4_2)->{'tag'} === "Just"))) {
$__t3 = ($GLOBALS['Data_Array__updateAt'])($GLOBALS['Data_Maybe_Just'], new Phpurs_Data0("Nothing"), $i_0, ($v_4_2)->{'value0'}, $xs_2);
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
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();

// Data_Array_all
$GLOBALS['Data_Array_all'] = (function() {
  $__fn = function($__local_var_0 = null, $__local_var_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($GLOBALS['Data_Array_allImpl'])($__local_var_0, $__local_var_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();

