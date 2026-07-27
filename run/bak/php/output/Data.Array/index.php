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
$GLOBALS['Data_Array__deleteAt'] = ($ffi_Data_Array['_deleteAt'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array__insertAt'] = ($ffi_Data_Array['_insertAt'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array__updateAt'] = ($ffi_Data_Array['_updateAt'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_allImpl'] = ($ffi_Data_Array['allImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_anyImpl'] = ($ffi_Data_Array['anyImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
function majData_majArray_concat($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majArray_concat';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Array;
  $f = ($ffi_Data_Array['concat'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Array_concat'] = __NAMESPACE__ . '\\majData_majArray_concat';

$GLOBALS['Data_Array_filterImpl'] = ($ffi_Data_Array['filterImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_findIndexImpl'] = ($ffi_Data_Array['findIndexImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_findLastIndexImpl'] = ($ffi_Data_Array['findLastIndexImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_findMapImpl'] = ($ffi_Data_Array['findMapImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_fromFoldableImpl'] = ($ffi_Data_Array['fromFoldableImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_indexImpl'] = ($ffi_Data_Array['indexImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
function majData_majArray_length($v0): int|\Closure {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majArray_length';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Array;
  $f = ($ffi_Data_Array['length'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Array_length'] = __NAMESPACE__ . '\\majData_majArray_length';

$GLOBALS['Data_Array_partitionImpl'] = ($ffi_Data_Array['partitionImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_rangeImpl'] = ($ffi_Data_Array['rangeImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_replicateImpl'] = ($ffi_Data_Array['replicateImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
function majData_majArray_reverse($v0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\majData_majArray_reverse';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  global $ffi_Data_Array;
  $f = ($ffi_Data_Array['reverse'] ?? new class { public function __invoke(...$args) { return $this; } });
  return $f($v0);
}
$GLOBALS['Data_Array_reverse'] = __NAMESPACE__ . '\\majData_majArray_reverse';

$GLOBALS['Data_Array_scanlImpl'] = ($ffi_Data_Array['scanlImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_scanrImpl'] = ($ffi_Data_Array['scanrImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_sliceImpl'] = ($ffi_Data_Array['sliceImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_sortByImpl'] = ($ffi_Data_Array['sortByImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_unconsImpl'] = ($ffi_Data_Array['unconsImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_unsafeIndexImpl'] = ($ffi_Data_Array['unsafeIndexImpl'] ?? new class { public function __invoke(...$args) { return $this; } });
$GLOBALS['Data_Array_zipWithImpl'] = ($ffi_Data_Array['zipWithImpl'] ?? new class { public function __invoke(...$args) { return $this; } });




// Data_Array_traverse_
$GLOBALS['Data_Array_traverse_'] = ($GLOBALS['Data_Foldable_traverse_'])($GLOBALS['Control_Monad_ST_Internal_applicativeST']);

// Data_Array_lessThan
$GLOBALS['Data_Array_lessThan'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Array_lessThanOrEq
$GLOBALS['Data_Array_lessThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_GT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

// Data_Array_discard
$GLOBALS['Data_Array_discard'] = (($GLOBALS['Control_Bind_discardUnit'])->{'discard'})($GLOBALS['Control_Monad_ST_Internal_bindST']);

// Data_Array_void
$GLOBALS['Data_Array_void'] = (($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(function($v_0) {
  $__num = \func_num_args();
  $__res = $GLOBALS['Data_Unit_unit'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});

// Data_Array_intercalate1
function majData_majArray_intercalate1($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_intercalate1';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__local_var_1_0 = (($dictMonoid_0)->{'Semigroup0'})(null);
  $mempty_2_1 = ($dictMonoid_0)->{'mempty'};
  $__res = (function() use ($__local_var_1_0, $mempty_2_1) {
  $__fn = function($sep_3, $xs_4 = null) use ($__local_var_1_0, $mempty_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($GLOBALS['Data_Foldable_foldableArray'])->{'foldl'})((function() use ($__local_var_1_0, $sep_3) {
  $__fn = function($v_5, $v1_6 = null) use ($__local_var_1_0, $sep_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t2 = null;;
  if (($v_5)->{'init'}) {
$__t2 = (object)["init" => false, "acc" => $v1_6];
goto end_branch_2;;
};
  $__t2 = (object)["init" => false, "acc" => ((($__local_var_1_0)->{'append'})(($v_5)->{'acc'}))(((($__local_var_1_0)->{'append'})($sep_3))($v1_6))];
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))((object)["init" => true, "acc" => $mempty_2_1]))($xs_4))->{'acc'};
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_intercalate1'] = __NAMESPACE__ . '\\majData_majArray_intercalate1';

// Data_Array_fromJust
function majData_majArray_frommajJust($v_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_frommajJust';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__t0 = null;;
  if ($v_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t0 = ($v_0)->{'value0'};
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
$GLOBALS['Data_Array_fromJust'] = __NAMESPACE__ . '\\majData_majArray_frommajJust';

// Data_Array_greaterThanOrEq
$GLOBALS['Data_Array_greaterThanOrEq'] = (function() use (&$__fn) {
$__local_var_0_0 = ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT());
return (function() use ($__local_var_0_0) {
  $__fn = function($a1_1, $a2_2 = null) use ($__local_var_0_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ( ! (($__local_var_0_0)($a1_1))($a2_2) instanceof \Data\Ordering\Data_Ordering_LT);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
})();

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
function majData_majArray_zipmajWithmajA($dictApplicative_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_zipmajWithmajA';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $sequence1_1_0 = (($GLOBALS['Data_Traversable_traversableArray'])->{'sequence'})($dictApplicative_0);
  $__res = (function() use ($sequence1_1_0) {
  $__fn = function($f_2, $xs_3 = null, $ys_4 = null) use ($sequence1_1_0, &$__fn) {
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
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
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
function majData_majArray_updatemajAtmajIndices($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_updatemajAtmajIndices';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverse_1_1_0 = ($GLOBALS['Data_Array_traverse_'])($dictFoldable_0);
  $__res = (function() use ($traverse_1_1_0) {
  $__fn = function($us_2, $xs_3 = null) use ($traverse_1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(\Data\Array\ST\majData_majArray_majSmajT_withmajArray(function($res_4) use ($traverse_1_1_0, $us_2) {
  $__num = \func_num_args();
  $__res = (($traverse_1_1_0)(function($v_5) use ($res_4) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn3($GLOBALS['Data_Array_ST_pokeImpl'], ($v_5)->{'value0'}, ($v_5)->{'value1'}, $res_4);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($us_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $xs_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
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
function majData_majArray_unsafemajIndex($_dollar__unused_0, $__local_var_1 = null, $__local_var_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_unsafemajIndex';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($__local_var_1)[0];
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
}, (function() {
  $__fn = function($x_1, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Maybe\Data_Maybe_Just((object)["head" => $x_1, "tail" => $xs_2]);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), $__local_var_0);
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
  if ((($GLOBALS['Data_Array_lessThan'])($i_3))($len_2_0)) {
$__t1 = new \Data\Maybe\Data_Maybe_Just(new \Data\Tuple\Data_Tuple_Tuple(($xs_1)[0], ($i_3 + 1)));
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
}, (function() {
  $__fn = function($v_1, $xs_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new \Data\Maybe\Data_Maybe_Just($xs_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), $__local_var_0);
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
  $__res = ($GLOBALS['Data_Array_sortBy'])((function() use ($dictOrd_0, $f_1) {
  $__fn = function($x_2, $y_3 = null) use ($dictOrd_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictOrd_0)->{'compare'})(($f_1)($x_2)))(($f_1)($y_3));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})());
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
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(\Data\Array\ST\majData_majArray_majSmajT_withmajArray(($GLOBALS['Data_Array_ST_push'])($x_1), $xs_0));
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
  if ((($GLOBALS['Data_Array_lessThanOrEq'])($v_0))(0)) {
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
  if ((($GLOBALS['Data_Array_lessThan'])($n_0))(1)) {
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
function majData_majArray_modifymajAtmajIndices($dictFoldable_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_modifymajAtmajIndices';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $traverse_1_1_0 = ($GLOBALS['Data_Array_traverse_'])($dictFoldable_0);
  $__res = (function() use ($traverse_1_1_0) {
  $__fn = function($is_2, $f_3 = null, $xs_4 = null) use ($traverse_1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(\Data\Array\ST\majData_majArray_majSmajT_withmajArray(function($res_5) use ($f_3, $is_2, $traverse_1_1_0) {
  $__num = \func_num_args();
  $__res = (($traverse_1_1_0)(function($i_6) use ($f_3, $res_5) {
  $__num = \func_num_args();
  $__res = \Data\Array\ST\majData_majArray_majSmajT_modify($i_6, $f_3, $res_5);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))($is_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, $xs_4));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_modifyAtIndices'] = __NAMESPACE__ . '\\majData_majArray_modifymajAtmajIndices';

// Data_Array_mapWithIndex_closure
$GLOBALS['Data_Array_mapWithIndex_closure'] = ($GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'])->{'mapWithIndex'};

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
  $__t1 = null;;
  if ((($GLOBALS['Data_Array_lessThan'])($v_2_0))(2)) {
$__t1 = $arr_1;
goto end_branch_1;;
};
  $__t1 = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})($GLOBALS['Data_Array_ST_new']))(function($out_3) use ($a_0, $arr_1, $v_2_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], ($arr_1)[0], $out_3)))(function($_dollar__unused_4) use ($a_0, $arr_1, $out_3, $v_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_discard'])(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_for(1, $v_2_0, function($idx_5) use ($a_0, $arr_1, $out_3) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], $a_0, $out_3)))(function($_dollar__unused_6) use ($arr_1, $idx_5, $out_3) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_void'])(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], ($arr_1)[0], $out_3));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar__unused_5) use ($out_3) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})($out_3);
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
  $__res = ($GLOBALS['Data_Array_intercalate1'])($dictMonoid_0);
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
  $__res = ((($GLOBALS['Data_Maybe_applyMaybe'])->{'apply'})(((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})((function() {
  $__fn = function($v_1, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (object)["init" => $v_1, "last" => $v1_2];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(match (count($xs_0)) { 0 => new \Data\Maybe\Data_Maybe_Nothing(), default => new \Data\Maybe\Data_Maybe_Just(($GLOBALS['Data_Array_sliceImpl'])(0, (count($xs_0) - 1), $xs_0)) })))(($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_0, (count($xs_0) - 1)));
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
  $go__2_0 = null;
  $go__2_0 = function($i_3) use ($arr_1, &$go__2_0, $p_0) {
  $__num = \func_num_args();
  $__tco_var_go__2_0_0_i_3 = $i_3;
  tco_loop_go__2_0_0:;
  $i_3 = $__tco_var_go__2_0_0_i_3;
  $v_4_0 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $arr_1, $i_3);
  $__t1 = null;;
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = null;;
if (($p_0)(($v_4_0)->{'value0'})) {
$__tco_3 = ($i_3 + 1);
$__tco_var_go__2_0_0_i_3 = $__tco_3;
goto tco_loop_go__2_0_0;;
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
  $breakIndex_3_1 = ($go__2_0)(0);
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
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})($GLOBALS['Data_Array_ST_new']))(function($fsts_1) use ($xs_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})($GLOBALS['Data_Array_ST_new']))(function($snds_2) use ($fsts_1, $xs_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(($GLOBALS['Data_Array_ST_Iterator_Iterator'])(function($v_3) use ($xs_0) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_0, $v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_new(0))))(function($iter_3) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_discard'])(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_iterate($iter_3, function($v_4) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  $__local_var_5_0 = ($v_4)->{'value1'};
  $__res = (($GLOBALS['Data_Array_discard'])(($GLOBALS['Data_Array_void'])(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], ($v_4)->{'value0'}, $fsts_1))))(function($_dollar__unused_6) use ($__local_var_5_0, $snds_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_void'])(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], $__local_var_5_0, $snds_2));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar__unused_4) use ($fsts_1, $snds_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $fsts_1)))(function($fsts_prime_5) use ($snds_2) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $snds_2)))(function($snds_prime_6) use ($fsts_prime_5) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})(new \Data\Tuple\Data_Tuple_Tuple($fsts_prime_5, $snds_prime_6));
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
  $indexedAndSorted_2_0 = \Data\Array\majData_majArray_sortmajBy((function() use ($comp_0) {
  $__fn = function($x_2, $y_3 = null) use ($comp_0, &$__fn) {
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
})(), ((($GLOBALS['Data_FunctorWithIndex_functorWithIndexArray'])->{'mapWithIndex'})($GLOBALS['Data_Tuple_Tuple']))($xs_1));
  $v_3_1 = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $indexedAndSorted_2_0, 0);
  $__t2 = null;;
  if ($v_3_1 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t2 = [];
goto end_branch_2;;
};
  if ($v_3_1 instanceof \Data\Maybe\Data_Maybe_Just) {
$__t2 = ((($GLOBALS['Data_Functor_functorArray'])->{'map'})($GLOBALS['Data_Tuple_snd']))(\Data\Array\majData_majArray_sortmajWith((object)["compare" => ((($GLOBALS['Data_Ord_ordIntImpl'])(new \Data\Ordering\Data_Ordering_LT()))(new \Data\Ordering\Data_Ordering_EQ()))(new \Data\Ordering\Data_Ordering_GT()), "Eq0" => function($_dollar__unused_4) {
  $__num = \func_num_args();
  $__res = (object)["eq" => $GLOBALS['Data_Eq_eqIntImpl']];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}], $GLOBALS['Data_Tuple_fst'], \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeThawImpl'], [($v_3_1)->{'value0'}])))(function($result_4) use ($comp_0, $indexedAndSorted_2_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_discard'])(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_foreach($indexedAndSorted_2_0, function($v1_5) use ($comp_0, $result_4) {
  $__num = \func_num_args();
  $__local_var_6_3 = ($v1_5)->{'value1'};
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Tuple_snd']))((($GLOBALS['Control_Semigroupoid_composeImpl'])($GLOBALS['Data_Array_fromJust']))($GLOBALS['Data_Array_last']))))(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $result_4))))(function($lst_7) use ($__local_var_6_3, $comp_0, $result_4, $v1_5) {
  $__num = \func_num_args();
  $__local_var_8_4 = ($GLOBALS['Data_Array_void'])(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], $v1_5, $result_4));
  $__t5 = null;;
  if (((($GLOBALS['Data_Eq_eqBoolean'])->{'eq'})(((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($comp_0)($lst_7))($__local_var_6_3)))(new \Data\Ordering\Data_Ordering_EQ())))(false)) {
$__t5 = $__local_var_8_4;
goto end_branch_5;;
};
  $__t5 = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})($GLOBALS['Data_Unit_unit']);
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar__unused_5) use ($result_4) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $result_4);
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
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})($GLOBALS['Data_Array_ST_new']))(function($result_2) use ($op_0, $xs_1) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})(($GLOBALS['Data_Array_ST_Iterator_Iterator'])(function($v_3) use ($xs_1) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_indexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $xs_1, $v_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_new(0))))(function($iter_3) use ($op_0, $result_2) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_discard'])(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_iterate($iter_3, function($x_4) use ($iter_3, $op_0, $result_2) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Array_void'])(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})($GLOBALS['Data_Array_ST_new']))(function($sub1_5) use ($iter_3, $op_0, $result_2, $x_4) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], $x_4, $sub1_5)))(function($_dollar__unused_6) use ($iter_3, $op_0, $result_2, $sub1_5, $x_4) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_discard'])(\Data\Array\ST\Iterator\majData_majArray_majSmajT_majIterator_pushmajWhile(($op_0)($x_4), $iter_3, $sub1_5)))(function($_dollar__unused_7) use ($result_2, $sub1_5) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $sub1_5)))(function($grp_8) use ($result_2) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], $grp_8, $result_2);
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
})))(function($_dollar__unused_4) use ($result_2) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $result_2);
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
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_Array_groupBy'])((function() use ($cmp_0) {
  $__fn = function($x_1, $y_2 = null) use ($cmp_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_0)($x_1))($y_2)))(new \Data\Ordering\Data_Ordering_EQ());
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})())))(($GLOBALS['Data_Array_sortBy'])($cmp_0));
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
  $eq2_1_0 = ($dictEq_0)->{'eq'};
  $__res = function($xs_2) use ($eq2_1_0) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_groupmajBy($eq2_1_0, $xs_2);
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
$GLOBALS['Data_Array_foldr_closure'] = ($GLOBALS['Data_Foldable_foldableArray'])->{'foldr'};

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
$GLOBALS['Data_Array_foldl_closure'] = ($GLOBALS['Data_Foldable_foldableArray'])->{'foldl'};

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
  $go__1_0 = null;
  $go__1_0 = (function() use (&$go__1_0, $xs_0) {
  $__fn = function($idx_2, $allArrays_3 = null) use (&$go__1_0, $xs_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__tco_var_go__1_0_0_idx_2 = $idx_2;
  $__tco_var_go__1_0_0_allArrays_3 = $allArrays_3;
  tco_loop_go__1_0_0:;
  $idx_2 = $__tco_var_go__1_0_0_idx_2;
  $allArrays_3 = $__tco_var_go__1_0_0_allArrays_3;
  $v_4_0 = (((($GLOBALS['Data_Foldable_foldableArray'])->{'foldl'})((function() use ($idx_2) {
  $__fn = function($acc_4, $nextArr_5 = null) use ($idx_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
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
$__t2 = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(\Data\Array\ST\majData_majArray_majSmajT_withmajArray(($GLOBALS['Data_Array_ST_push'])(($__local_var_6_0)->{'value0'}), ($acc_4)->{'value0'}));
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
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(new \Data\Maybe\Data_Maybe_Nothing()))($xs_0);
  $__t4 = null;;
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Nothing) {
$__t4 = $allArrays_3;
goto end_branch_4;;
};
  if ($v_4_0 instanceof \Data\Maybe\Data_Maybe_Just) {
$__tco_5 = ($idx_2 + 1);
$__tco_6 = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(\Data\Array\ST\majData_majArray_majSmajT_withmajArray(($GLOBALS['Data_Array_ST_push'])(($v_4_0)->{'value0'}), $allArrays_3));
$__tco_var_go__1_0_0_idx_2 = $__tco_5;
$__tco_var_go__1_0_0_allArrays_3 = $__tco_6;
goto tco_loop_go__1_0_0;;
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
  $__res = (($go__1_0)(0))([]);
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
  $__local_var_2_1 = (($Monad0_1_0)->{'Applicative0'})(null);
  $__res = (function() use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0) {
  $__fn = function($f_3, $b_4 = null, $array_5 = null) use ($Monad0_1_0, $__local_var_2_1, $dictMonadRec_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($dictMonadRec_0)->{'tailRecM'})(function($o_6) use ($Monad0_1_0, $__local_var_2_1, $array_5, $f_3) {
  $__num = \func_num_args();
  $__t2 = null;;
  if ((($GLOBALS['Data_Array_greaterThanOrEq'])(($o_6)->{'b'}))(count($array_5))) {
$__t2 = (($__local_var_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Done(($o_6)->{'a'}));
goto end_branch_2;;
};
  $__t2 = ((((($Monad0_1_0)->{'Bind1'})(null))->{'bind'})((($f_3)(($o_6)->{'a'}))(($array_5)[0])))(function($res_prime_7) use ($__local_var_2_1, $o_6) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)->{'pure'})(new \Control\Monad\Rec\Class\Control_Monad_Rec_Class_Loop((object)["a" => $res_prime_7, "b" => (($o_6)->{'b'} + 1)]));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))((object)["a" => $b_4, "b" => 0]);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
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
  $__res = (($GLOBALS['Data_Foldable_foldableArray'])->{'foldMap'})($dictMonoid_0);
  goto __end;;
  __end:
  return 1 < $__num ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}
$GLOBALS['Data_Array_foldMap'] = __NAMESPACE__ . '\\majData_majArray_foldmajMap';

// Data_Array_foldM
function majData_majArray_foldmajM($dictMonad_0, $f_1 = null, $b_2 = null, $__local_var_3 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_foldmajM';
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
  $__res = ($GLOBALS['Data_Array_unconsImpl'])(function($v_4) use ($b_2, $dictMonad_0) {
  $__num = \func_num_args();
  $__res = (((($dictMonad_0)->{'Applicative0'})(null))->{'pure'})($b_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, (function() use ($b_2, $dictMonad_0, $f_1) {
  $__fn = function($a_4, $as_5 = null) use ($b_2, $dictMonad_0, $f_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((($dictMonad_0)->{'Bind1'})(null))->{'bind'})((($f_1)($b_2))($a_4)))(function($b_prime_6) use ($as_5, $dictMonad_0, $f_1) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_foldmajM($dictMonad_0, $f_1, $b_prime_6, $as_5);
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
  return 4 < $__num ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
}
$GLOBALS['Data_Array_foldM'] = __NAMESPACE__ . '\\majData_majArray_foldmajM';

// Data_Array_fold
function majData_majArray_fold($dictMonoid_0) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_fold';
  if ($__num < 1) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 1);
  }
  $__res = ((($GLOBALS['Data_Foldable_foldableArray'])->{'foldMap'})($dictMonoid_0))($GLOBALS['Data_Foldable_identity']);
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
  $__res = ((($GLOBALS['Data_Ordering_eqOrdering'])->{'eq'})((($cmp_0)($x_1))($y_3)))(new \Data\Ordering\Data_Ordering_GT());
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
  $__res = ((($GLOBALS['Data_Maybe_functorMaybe'])->{'map'})(function($__local_var_2) use ($xs_1) {
  $__num = \func_num_args();
  $__res = ($xs_1)[0];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))(($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), $f_0, $xs_1));
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
function majData_majArray_intersectmajBy($eq2_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_intersectmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ($GLOBALS['Data_Array_filterImpl'])(function($x_3) use ($eq2_0, $ys_2) {
  $__num = \func_num_args();
  $__local_var_4_0 = ($GLOBALS['Data_Array_findIndexImpl'])($GLOBALS['Data_Maybe_Just'], new \Data\Maybe\Data_Maybe_Nothing(), ($eq2_0)($x_3), $ys_2);
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
function majData_majArray_notmajElem($dictEq_0, $a_1 = null, $arr_2 = null): bool|\Closure {
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
function majData_majArray_elem($dictEq_0, $a_1 = null, $arr_2 = null): bool|\Closure {
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
  if ((($GLOBALS['Data_Array_lessThan'])($__local_var_2_0))(1)) {
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
  if ((($GLOBALS['Data_Array_lessThan'])($n_0))(1)) {
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
  if ((($GLOBALS['Data_Array_lessThan'])($__local_var_2_0))(1)) {
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
  $__res = (($GLOBALS['Data_Foldable_foldableArray'])->{'foldr'})(($GLOBALS['Data_Array_delete'])($dictEq_0));
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
  $__res = ((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})([$x_0]))($xs_1);
  goto __end;;
  __end:
  return 2 < $__num ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
}
$GLOBALS['Data_Array_cons'] = __NAMESPACE__ . '\\majData_majArray_cons';

// Data_Array_some
function majData_majArray_some($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_some';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)->{'Applicative0'})(null))->{'Apply0'})(null))->{'apply'})(((((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'Functor0'})(null))->{'map'})($GLOBALS['Data_Array_cons']))($v_2)))((($dictLazy_1)->{'defer'})(function($v1_3) use ($dictAlternative_0, $dictLazy_1, $v_2) {
  $__num = \func_num_args();
  $__res = \Data\Array\majData_majArray_many($dictAlternative_0, $dictLazy_1, $v_2);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_some'] = __NAMESPACE__ . '\\majData_majArray_some';

// Data_Array_many
function majData_majArray_many($dictAlternative_0, $dictLazy_1 = null, $v_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_many';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictAlternative_0)->{'Plus1'})(null))->{'Alt0'})(null))->{'alt'})(\Data\Array\majData_majArray_some($dictAlternative_0, $dictLazy_1, $v_2)))((((($dictAlternative_0)->{'Applicative0'})(null))->{'pure'})([]));
  goto __end;;
  __end:
  return 3 < $__num ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
}
$GLOBALS['Data_Array_many'] = __NAMESPACE__ . '\\majData_majArray_many';

// Data_Array_concatMap
function majData_majArray_concatmajMap($b_0, $a_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_concatmajMap';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($GLOBALS['Control_Bind_bindArray'])->{'bind'})($a_1))($b_0);
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
  $traverse1_1_0 = (($GLOBALS['Data_Traversable_traversableArray'])->{'traverse'})($dictApplicative_0);
  $__local_var_2_1 = (((($dictApplicative_0)->{'Apply0'})(null))->{'Functor0'})(null);
  $__res = function($p_3) use ($__local_var_2_1, $traverse1_1_0) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Control_Semigroupoid_composeImpl'])((($__local_var_2_1)->{'map'})(($GLOBALS['Data_Array_mapMaybe'])(function($v_4) {
  $__num = \func_num_args();
  $__t2 = null;;
  if (($v_4)->{'value1'}) {
$__t2 = new \Data\Maybe\Data_Maybe_Just(($v_4)->{'value0'});
goto end_branch_2;;
};
  $__t2 = new \Data\Maybe\Data_Maybe_Nothing();
  end_branch_2:;
  $__res = $__t2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(($traverse1_1_0)(function($x_4) use ($__local_var_2_1, $p_3) {
  $__num = \func_num_args();
  $__res = ((($__local_var_2_1)->{'map'})(($GLOBALS['Data_Tuple_Tuple'])($x_4)))(($p_3)($x_4));
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
function majData_majArray_nubmajBymajEq($eq2_0, $xs_1 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_nubmajBymajEq';
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_run(((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})($GLOBALS['Data_Array_ST_new']))(function($arr_2) use ($eq2_0, $xs_1) {
  $__num = \func_num_args();
  $__res = (($GLOBALS['Data_Array_discard'])(\Control\Monad\ST\Internal\majControl_majMonad_majSmajT_majInternal_foreach($xs_1, function($x_3) use ($arr_2, $eq2_0) {
  $__num = \func_num_args();
  $__res = ((($GLOBALS['Control_Monad_ST_Internal_bindST'])->{'bind'})(((($GLOBALS['Control_Monad_ST_Internal_functorST'])->{'map'})((($GLOBALS['Control_Semigroupoid_composeImpl'])(($GLOBALS['Data_HeytingAlgebra_heytingAlgebraBoolean'])->{'not'}))(($GLOBALS['Data_Array_any'])(function($v_4) use ($eq2_0, $x_3) {
  $__num = \func_num_args();
  $__res = (($eq2_0)($v_4))($x_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}))))(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $arr_2))))(function($e_4) use ($arr_2, $x_3) {
  $__num = \func_num_args();
  $__local_var_5_0 = ($GLOBALS['Data_Array_void'])(\Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn2($GLOBALS['Data_Array_ST_pushImpl'], $x_3, $arr_2));
  $__t1 = null;;
  if ($e_4) {
$__t1 = $__local_var_5_0;
goto end_branch_1;;
};
  $__t1 = (($GLOBALS['Control_Monad_ST_Internal_applicativeST'])->{'pure'})($GLOBALS['Data_Unit_unit']);
  end_branch_1:;
  $__res = $__t1;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
});
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
})))(function($_dollar__unused_3) use ($arr_2) {
  $__num = \func_num_args();
  $__res = \Control\Monad\ST\Uncurried\majControl_majMonad_majSmajT_majUncurried_runmajSmajTmajFn1($GLOBALS['Data_Array_ST_unsafeFreezeImpl'], $arr_2);
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
function majData_majArray_unionmajBy($eq2_0, $xs_1 = null, $ys_2 = null) {
  $__num = \func_num_args();
  $__fn = __NAMESPACE__ . '\\' . 'majData_majArray_unionmajBy';
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($GLOBALS['Data_Semigroup_semigroupArray'])->{'append'})($xs_1))((((($GLOBALS['Data_Foldable_foldableArray'])->{'foldl'})((function() use ($eq2_0) {
  $__fn = function($b_3, $a_4 = null) use ($eq2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = \Data\Array\majData_majArray_deletemajBy($eq2_0, $a_4, $b_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))(\Data\Array\majData_majArray_nubmajBymajEq($eq2_0, $ys_2)))($xs_1));
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

