<?php

namespace Data\Traversable;

// ALL IMPORTS: Control.Applicative, Control.Apply, Control.Category, Control.Semigroupoid, Data.Const, Data.Either, Data.Foldable, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Traversable, Data.Traversable.Accum, Data.Traversable.Accum.Internal, Data.Tuple, Prelude, Prim
// TO REQUIRE: Control.Applicative, Control.Apply, Control.Category, Control.Semigroupoid, Data.Const, Data.Either, Data.Foldable, Data.Function, Data.Functor, Data.Functor.App, Data.Functor.Compose, Data.Functor.Coproduct, Data.Functor.Product, Data.Identity, Data.Maybe, Data.Maybe.First, Data.Maybe.Last, Data.Monoid.Additive, Data.Monoid.Conj, Data.Monoid.Disj, Data.Monoid.Dual, Data.Monoid.Multiplicative, Data.Traversable, Data.Traversable.Accum, Data.Traversable.Accum.Internal, Data.Tuple, Prelude
require_once __DIR__ . '/../Control.Applicative/index.php';
require_once __DIR__ . '/../Control.Apply/index.php';
require_once __DIR__ . '/../Control.Category/index.php';
require_once __DIR__ . '/../Control.Semigroupoid/index.php';
require_once __DIR__ . '/../Data.Const/index.php';
require_once __DIR__ . '/../Data.Either/index.php';
require_once __DIR__ . '/../Data.Foldable/index.php';
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Functor/index.php';
require_once __DIR__ . '/../Data.Functor.App/index.php';
require_once __DIR__ . '/../Data.Functor.Compose/index.php';
require_once __DIR__ . '/../Data.Functor.Coproduct/index.php';
require_once __DIR__ . '/../Data.Functor.Product/index.php';
require_once __DIR__ . '/../Data.Identity/index.php';
require_once __DIR__ . '/../Data.Maybe/index.php';
require_once __DIR__ . '/../Data.Maybe.First/index.php';
require_once __DIR__ . '/../Data.Maybe.Last/index.php';
require_once __DIR__ . '/../Data.Monoid.Additive/index.php';
require_once __DIR__ . '/../Data.Monoid.Conj/index.php';
require_once __DIR__ . '/../Data.Monoid.Disj/index.php';
require_once __DIR__ . '/../Data.Monoid.Dual/index.php';
require_once __DIR__ . '/../Data.Monoid.Multiplicative/index.php';
require_once __DIR__ . '/../Data.Traversable/index.php';
require_once __DIR__ . '/../Data.Traversable.Accum/index.php';
require_once __DIR__ . '/../Data.Traversable.Accum.Internal/index.php';
require_once __DIR__ . '/../Data.Tuple/index.php';
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
\PhpursThunks::$thunks['Data_Traversable_traverse'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['traverse'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableTuple'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])((($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple')))(($v_2)->{'value0'})))(($f_1)(($v_2)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])((($GLOBALS['Data_Tuple_Tuple'] ?? \PhpursThunks::eval('Data_Tuple_Tuple')))(($v_1)->{'value0'})))(($v_1)->{'value1'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Tuple_functorTuple'] ?? \PhpursThunks::eval('Data_Tuple_functorTuple'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableTuple'] ?? \PhpursThunks::eval('Data_Foldable_foldableTuple'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableMultiplicative'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Multiplicative_Multiplicative'] ?? \PhpursThunks::eval('Data_Monoid_Multiplicative_Multiplicative'))))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Multiplicative_Multiplicative'] ?? \PhpursThunks::eval('Data_Monoid_Multiplicative_Multiplicative'))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Monoid_Multiplicative_functorMultiplicative'] ?? \PhpursThunks::eval('Data_Monoid_Multiplicative_functorMultiplicative'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableMultiplicative'] ?? \PhpursThunks::eval('Data_Foldable_foldableMultiplicative'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableMaybe'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Nothing"))) {
$__t0 = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Nothing"));
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Just"))) {
$__t0 = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(($v_1)(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Nothing"))) {
$__t1 = (($dictApplicative_0)['pure'])(new Phpurs_Data0("Nothing"));
goto end_branch_1;;
};
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Just"))) {
$__t1 = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Maybe_Just'] ?? \PhpursThunks::eval('Data_Maybe_Just'))))(($v_1)->{'value0'});
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
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_functorMaybe'] ?? \PhpursThunks::eval('Data_Maybe_functorMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableMaybe'] ?? \PhpursThunks::eval('Data_Foldable_foldableMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableIdentity'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Identity_Identity'] ?? \PhpursThunks::eval('Data_Identity_Identity'))))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Identity_Identity'] ?? \PhpursThunks::eval('Data_Identity_Identity'))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Identity_functorIdentity'] ?? \PhpursThunks::eval('Data_Identity_functorIdentity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableIdentity'] ?? \PhpursThunks::eval('Data_Foldable_foldableIdentity'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableEither'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__t0 = null;;
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Left"))) {
$__t0 = (($dictApplicative_0)['pure'])(new Phpurs_Data1("Left", ($v1_2)->{'value0'}));
goto end_branch_0;;
};
  if ((is_object($v1_2) && (($v1_2)->{'tag'} === "Right"))) {
$__t0 = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))(($v_1)(($v1_2)->{'value0'}));
goto end_branch_0;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t0 = null;
  end_branch_0:;
  $__res = $__t0;
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__t1 = null;;
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Left"))) {
$__t1 = (($dictApplicative_0)['pure'])(new Phpurs_Data1("Left", ($v_1)->{'value0'}));
goto end_branch_1;;
};
  if ((is_object($v_1) && (($v_1)->{'tag'} === "Right"))) {
$__t1 = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))(($v_1)->{'value0'});
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
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Either_functorEither'] ?? \PhpursThunks::eval('Data_Either_functorEither'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableEither'] ?? \PhpursThunks::eval('Data_Foldable_foldableEither'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableDual'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Dual_Dual'] ?? \PhpursThunks::eval('Data_Monoid_Dual_Dual'))))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Dual_Dual'] ?? \PhpursThunks::eval('Data_Monoid_Dual_Dual'))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Monoid_Dual_functorDual'] ?? \PhpursThunks::eval('Data_Monoid_Dual_functorDual'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableDual'] ?? \PhpursThunks::eval('Data_Foldable_foldableDual'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableDisj'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Disj_Disj'] ?? \PhpursThunks::eval('Data_Monoid_Disj_Disj'))))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Disj_Disj'] ?? \PhpursThunks::eval('Data_Monoid_Disj_Disj'))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Monoid_Disj_functorDisj'] ?? \PhpursThunks::eval('Data_Monoid_Disj_functorDisj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableDisj'] ?? \PhpursThunks::eval('Data_Foldable_foldableDisj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableConst'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null, $v1_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictApplicative_0)['pure'])($v1_2);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictApplicative_0)['pure'])($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Const_functorConst'] ?? \PhpursThunks::eval('Data_Const_functorConst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableConst'] ?? \PhpursThunks::eval('Data_Foldable_foldableConst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableConj'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Conj_Conj'] ?? \PhpursThunks::eval('Data_Monoid_Conj_Conj'))))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Conj_Conj'] ?? \PhpursThunks::eval('Data_Monoid_Conj_Conj'))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Monoid_Conj_functorConj'] ?? \PhpursThunks::eval('Data_Monoid_Conj_functorConj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableConj'] ?? \PhpursThunks::eval('Data_Foldable_foldableConj'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableCompose'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $__tco_var_Data_Traversable_traversableCompose_dictTraversable_0 = $dictTraversable_0;
  tco_loop_Data_Traversable_traversableCompose:;
  $dictTraversable_0 = $__tco_var_Data_Traversable_traversableCompose_dictTraversable_0;
  $__local_var_1_0 = (($dictTraversable_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($dictTraversable_0)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = function($dictTraversable1_3 = null) use ($__local_var_1_0, $__local_var_2_1, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictTraversable1_3)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorCompose1_5_3 = ["map" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5 = null, $v_6 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($__local_var_1_0)['map'])((($__local_var_4_2)['map'])($f_5)))($v_6);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $__local_var_6_4 = (($dictTraversable1_3)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $foldableCompose1_7_5 = ["foldr" => (function() use ($__local_var_2_1, $__local_var_6_4) {
  $__fn = function($f_7 = null, $i_8 = null, $v_9 = null) use ($__local_var_2_1, $__local_var_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__local_var_10_5 = (($__local_var_6_4)['foldr'])($f_7);
  $__res = (((($__local_var_2_1)['foldr'])((function() use ($__local_var_10_5) {
  $__fn = function($b_11 = null, $a_12 = null) use ($__local_var_10_5, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($__local_var_10_5)($a_12))($b_11);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($i_8))($v_9);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($__local_var_2_1, $__local_var_6_4) {
  $__fn = function($f_7 = null, $i_8 = null, $v_9 = null) use ($__local_var_2_1, $__local_var_6_4, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_2_1)['foldl'])((($__local_var_6_4)['foldl'])($f_7)))($i_8))($v_9);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_7 = null) use ($__local_var_2_1, $__local_var_6_4) {
  $__num = \func_num_args();
  $foldMap4_8_6 = (($__local_var_2_1)['foldMap'])($dictMonoid_7);
  $foldMap5_9_7 = (($__local_var_6_4)['foldMap'])($dictMonoid_7);
  $__res = (function() use ($foldMap4_8_6, $foldMap5_9_7) {
  $__fn = function($f_10 = null, $v_11 = null) use ($foldMap4_8_6, $foldMap5_9_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($foldMap4_8_6)(($foldMap5_9_7)($f_10)))($v_11);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["traverse" => function($dictApplicative_8 = null) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $traverse4_9_9 = (($dictTraversable_0)['traverse'])($dictApplicative_8);
  $traverse5_10_10 = (($dictTraversable1_3)['traverse'])($dictApplicative_8);
  $__res = (function() use ($dictApplicative_8, $traverse4_9_9, $traverse5_10_10) {
  $__fn = function($f_11 = null, $v_12 = null) use ($dictApplicative_8, $traverse4_9_9, $traverse5_10_10, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_8)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Functor_Compose_Compose'] ?? \PhpursThunks::eval('Data_Functor_Compose_Compose'))))((($traverse4_9_9)(($traverse5_10_10)($f_11)))($v_12));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_8 = null) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__res = (((((($GLOBALS['Data_Traversable_traversableCompose'] ?? \PhpursThunks::eval('Data_Traversable_traversableCompose')))($dictTraversable_0))($dictTraversable1_3))['traverse'])($dictApplicative_8))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_8 = null) use ($functorCompose1_5_3) {
  $__num = \func_num_args();
  $__res = $functorCompose1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_8 = null) use ($foldableCompose1_7_5) {
  $__num = \func_num_args();
  $__res = $foldableCompose1_7_5;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableAdditive'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Additive_Additive'] ?? \PhpursThunks::eval('Data_Monoid_Additive_Additive'))))(($f_1)($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Monoid_Additive_Additive'] ?? \PhpursThunks::eval('Data_Monoid_Additive_Additive'))))($v_1);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Monoid_Additive_functorAdditive'] ?? \PhpursThunks::eval('Data_Monoid_Additive_functorAdditive'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableAdditive'] ?? \PhpursThunks::eval('Data_Foldable_foldableAdditive'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_sequenceDefault'] = function() { $v = (function() {
  $__fn = function($dictTraversable_0 = null, $dictApplicative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictTraversable_0)['traverse'])($dictApplicative_1))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity']);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableArray'] = function() { $v = ["traverse" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $Apply0_1_0 = (($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__res = (((($GLOBALS['Data_Traversable_traverseArrayImpl'] ?? \PhpursThunks::eval('Data_Traversable_traverseArrayImpl')))(($Apply0_1_0)['apply']))(((($Apply0_1_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map']))(($dictApplicative_0)['pure']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_0 = null) {
  $__num = \func_num_args();
  $__res = (((($GLOBALS['Data_Traversable_traversableArray'] ?? \PhpursThunks::eval('Data_Traversable_traversableArray')))['traverse'])($dictApplicative_0))((($GLOBALS['Control_Category_categoryFn'] ?? \PhpursThunks::eval('Control_Category_categoryFn')))['identity']);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Functor_functorArray'] ?? \PhpursThunks::eval('Data_Functor_functorArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableArray'] ?? \PhpursThunks::eval('Data_Foldable_foldableArray'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_sequence'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)['sequence'];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableApp'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictTraversable_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_2_1 = (($dictTraversable_0)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $foldableApp_3_2 = ["foldr" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3 = null, $i_4 = null, $v_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_2_1)['foldr'])($f_3))($i_4))($v_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldl" => (function() use ($__local_var_2_1) {
  $__fn = function($f_3 = null, $i_4 = null, $v_5 = null) use ($__local_var_2_1, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($__local_var_2_1)['foldl'])($f_3))($i_4))($v_5);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "foldMap" => function($dictMonoid_3 = null) use ($__local_var_2_1) {
  $__num = \func_num_args();
  $__res = (($__local_var_2_1)['foldMap'])($dictMonoid_3);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  $__res = ["traverse" => function($dictApplicative_4 = null) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $traverse3_5_3 = (($dictTraversable_0)['traverse'])($dictApplicative_4);
  $__res = (function() use ($dictApplicative_4, $traverse3_5_3) {
  $__fn = function($f_6 = null, $v_7 = null) use ($dictApplicative_4, $traverse3_5_3, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_4)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Functor_App_App'] ?? \PhpursThunks::eval('Data_Functor_App_App'))))((($traverse3_5_3)($f_6))($v_7));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_4 = null) use ($dictTraversable_0) {
  $__num = \func_num_args();
  $sequence3_5_4 = (($dictTraversable_0)['sequence'])($dictApplicative_4);
  $__res = function($v_6 = null) use ($dictApplicative_4, $sequence3_5_4) {
  $__num = \func_num_args();
  $__res = ((((((($dictApplicative_4)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Functor_App_App'] ?? \PhpursThunks::eval('Data_Functor_App_App'))))(($sequence3_5_4)($v_6));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_4 = null) use ($__local_var_1_0) {
  $__num = \func_num_args();
  $__res = $__local_var_1_0;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_4 = null) use ($foldableApp_3_2) {
  $__num = \func_num_args();
  $__res = $foldableApp_3_2;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableCoproduct'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictTraversable_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $foldableCoproduct_2_1 = (($GLOBALS['Data_Foldable_foldableCoproduct'] ?? \PhpursThunks::eval('Data_Foldable_foldableCoproduct')))((($dictTraversable_0)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictTraversable1_3 = null) use ($__local_var_1_0, $dictTraversable_0, $foldableCoproduct_2_1) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictTraversable1_3)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorCoproduct1_5_3 = ["map" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5 = null, $v_6 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__local_var_7_3 = (($__local_var_1_0)['map'])($f_5);
  $__local_var_8_4 = (($__local_var_4_2)['map'])($f_5);
  $__t5 = null;;
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Left"))) {
$__t5 = new Phpurs_Data1("Left", ($__local_var_7_3)(($v_6)->{'value0'}));
goto end_branch_5;;
};
  if ((is_object($v_6) && (($v_6)->{'tag'} === "Right"))) {
$__t5 = new Phpurs_Data1("Right", ($__local_var_8_4)(($v_6)->{'value0'}));
goto end_branch_5;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t5 = null;
  end_branch_5:;
  $__res = $__t5;
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $foldableCoproduct1_6_7 = ($foldableCoproduct_2_1)((($dictTraversable1_3)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["traverse" => function($dictApplicative_7 = null) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_8_8 = (((($dictApplicative_7)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $traverse4_9_9 = (($dictTraversable_0)['traverse'])($dictApplicative_7);
  $traverse5_10_10 = (($dictTraversable1_3)['traverse'])($dictApplicative_7);
  $__res = function($f_11 = null) use ($__local_var_8_8, $traverse4_9_9, $traverse5_10_10) {
  $__num = \func_num_args();
  $__local_var_12_11 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($__local_var_8_8)['map'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left'))))))(($traverse4_9_9)($f_11));
  $__local_var_13_12 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($__local_var_8_8)['map'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))))(($traverse5_10_10)($f_11));
  $__res = function($v2_14 = null) use ($__local_var_12_11, $__local_var_13_12) {
  $__num = \func_num_args();
  $__t13 = null;;
  if ((is_object($v2_14) && (($v2_14)->{'tag'} === "Left"))) {
$__t13 = ($__local_var_12_11)(($v2_14)->{'value0'});
goto end_branch_13;;
};
  if ((is_object($v2_14) && (($v2_14)->{'tag'} === "Right"))) {
$__t13 = ($__local_var_13_12)(($v2_14)->{'value0'});
goto end_branch_13;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t13 = null;
  end_branch_13:;
  $__res = $__t13;
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
}, "sequence" => function($dictApplicative_7 = null) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $__local_var_8_14 = (((($dictApplicative_7)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $__local_var_9_15 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($__local_var_8_14)['map'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Left'] ?? \PhpursThunks::eval('Data_Either_Left'))))))((($dictTraversable_0)['sequence'])($dictApplicative_7));
  $__local_var_10_16 = ((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))((($__local_var_8_14)['map'])(((($GLOBALS['Control_Semigroupoid_composeImpl'] ?? \PhpursThunks::eval('Control_Semigroupoid_composeImpl')))(($GLOBALS['Data_Functor_Coproduct_Coproduct'] ?? \PhpursThunks::eval('Data_Functor_Coproduct_Coproduct'))))(($GLOBALS['Data_Either_Right'] ?? \PhpursThunks::eval('Data_Either_Right'))))))((($dictTraversable1_3)['sequence'])($dictApplicative_7));
  $__res = function($v2_11 = null) use ($__local_var_10_16, $__local_var_9_15) {
  $__num = \func_num_args();
  $__t17 = null;;
  if ((is_object($v2_11) && (($v2_11)->{'tag'} === "Left"))) {
$__t17 = ($__local_var_9_15)(($v2_11)->{'value0'});
goto end_branch_17;;
};
  if ((is_object($v2_11) && (($v2_11)->{'tag'} === "Right"))) {
$__t17 = ($__local_var_10_16)(($v2_11)->{'value0'});
goto end_branch_17;;
};
  throw new \Exception("Failed pattern match at " . __FILE__ . ":" . __LINE__);
  $__t17 = null;
  end_branch_17:;
  $__res = $__t17;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_7 = null) use ($functorCoproduct1_5_3) {
  $__num = \func_num_args();
  $__res = $functorCoproduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_7 = null) use ($foldableCoproduct1_6_7) {
  $__num = \func_num_args();
  $__res = $foldableCoproduct1_6_7;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableFirst'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Maybe_First_First'] ?? \PhpursThunks::eval('Data_Maybe_First_First'))))(((((($GLOBALS['Data_Traversable_traversableMaybe'] ?? \PhpursThunks::eval('Data_Traversable_traversableMaybe')))['traverse'])($dictApplicative_0))($f_1))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Maybe_First_First'] ?? \PhpursThunks::eval('Data_Maybe_First_First'))))((((($GLOBALS['Data_Traversable_traversableMaybe'] ?? \PhpursThunks::eval('Data_Traversable_traversableMaybe')))['sequence'])($dictApplicative_0))($v_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_functorMaybe'] ?? \PhpursThunks::eval('Data_Maybe_functorMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableFirst'] ?? \PhpursThunks::eval('Data_Foldable_foldableFirst'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableLast'] = function() { $v = ["traverse" => (function() {
  $__fn = function($dictApplicative_0 = null, $f_1 = null, $v_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Maybe_Last_Last'] ?? \PhpursThunks::eval('Data_Maybe_Last_Last'))))(((((($GLOBALS['Data_Traversable_traversableMaybe'] ?? \PhpursThunks::eval('Data_Traversable_traversableMaybe')))['traverse'])($dictApplicative_0))($f_1))($v_2));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(), "sequence" => (function() {
  $__fn = function($dictApplicative_0 = null, $v_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((((((($dictApplicative_0)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])(($GLOBALS['Data_Maybe_Last_Last'] ?? \PhpursThunks::eval('Data_Maybe_Last_Last'))))((((($GLOBALS['Data_Traversable_traversableMaybe'] ?? \PhpursThunks::eval('Data_Traversable_traversableMaybe')))['sequence'])($dictApplicative_0))($v_1));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "Functor0" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Maybe_functorMaybe'] ?? \PhpursThunks::eval('Data_Maybe_functorMaybe'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_0 = null) {
  $__num = \func_num_args();
  $__res = ($GLOBALS['Data_Foldable_foldableLast'] ?? \PhpursThunks::eval('Data_Foldable_foldableLast'));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traversableProduct'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $__local_var_1_0 = (($dictTraversable_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $foldableProduct_2_1 = (($GLOBALS['Data_Foldable_foldableProduct'] ?? \PhpursThunks::eval('Data_Foldable_foldableProduct')))((($dictTraversable_0)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = function($dictTraversable1_3 = null) use ($__local_var_1_0, $dictTraversable_0, $foldableProduct_2_1) {
  $__num = \func_num_args();
  $__local_var_4_2 = (($dictTraversable1_3)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined')));
  $functorProduct1_5_3 = ["map" => (function() use ($__local_var_1_0, $__local_var_4_2) {
  $__fn = function($f_5 = null, $v_6 = null) use ($__local_var_1_0, $__local_var_4_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Tuple", ((($__local_var_1_0)['map'])($f_5))(($v_6)->{'value0'}), ((($__local_var_4_2)['map'])($f_5))(($v_6)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()];
  $foldableProduct1_6_4 = ($foldableProduct_2_1)((($dictTraversable1_3)['Foldable1'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $__res = ["traverse" => function($dictApplicative_7 = null) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $lift2_8_5 = (($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))((($dictApplicative_7)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $traverse4_9_6 = (($dictTraversable_0)['traverse'])($dictApplicative_7);
  $traverse5_10_7 = (($dictTraversable1_3)['traverse'])($dictApplicative_7);
  $__res = (function() use ($lift2_8_5, $traverse4_9_6, $traverse5_10_7) {
  $__fn = function($f_11 = null, $v_12 = null) use ($lift2_8_5, $traverse4_9_6, $traverse5_10_7, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($lift2_8_5)(($GLOBALS['Data_Functor_Product_product'] ?? \PhpursThunks::eval('Data_Functor_Product_product'))))((($traverse4_9_6)($f_11))(($v_12)->{'value0'})))((($traverse5_10_7)($f_11))(($v_12)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "sequence" => function($dictApplicative_7 = null) use ($dictTraversable1_3, $dictTraversable_0) {
  $__num = \func_num_args();
  $lift2_8_8 = (($GLOBALS['Control_Apply_lift2'] ?? \PhpursThunks::eval('Control_Apply_lift2')))((($dictApplicative_7)['Apply0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))));
  $sequence4_9_9 = (($dictTraversable_0)['sequence'])($dictApplicative_7);
  $sequence5_10_10 = (($dictTraversable1_3)['sequence'])($dictApplicative_7);
  $__res = function($v_11 = null) use ($lift2_8_8, $sequence4_9_9, $sequence5_10_10) {
  $__num = \func_num_args();
  $__res = ((($lift2_8_8)(($GLOBALS['Data_Functor_Product_product'] ?? \PhpursThunks::eval('Data_Functor_Product_product'))))(($sequence4_9_9)(($v_11)->{'value0'})))(($sequence5_10_10)(($v_11)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Functor0" => function($dollar__unused_7 = null) use ($functorProduct1_5_3) {
  $__num = \func_num_args();
  $__res = $functorProduct1_5_3;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}, "Foldable1" => function($dollar__unused_7 = null) use ($foldableProduct1_6_4) {
  $__num = \func_num_args();
  $__res = $foldableProduct1_6_4;
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_traverseDefault'] = function() { $v = (function() {
  $__fn = function($dictTraversable_0 = null, $dictApplicative_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $sequence3_2_0 = (($dictTraversable_0)['sequence'])($dictApplicative_1);
  $__res = (function() use ($dictTraversable_0, $sequence3_2_0) {
  $__fn = function($f_3 = null, $ta_4 = null) use ($dictTraversable_0, $sequence3_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ($sequence3_2_0)(((((($dictTraversable_0)['Functor0'])(($GLOBALS['Prim_undefined'] ?? \PhpursThunks::eval('Prim_undefined'))))['map'])($f_3))($ta_4));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_Traversable_mapAccumR'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $traverse2_1_0 = (($dictTraversable_0)['traverse'])(($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateR'] ?? \PhpursThunks::eval('Data_Traversable_Accum_Internal_applicativeStateR')));
  $__res = (function() use ($traverse2_1_0) {
  $__fn = function($f_2 = null, $s0_3 = null, $xs_4 = null) use ($traverse2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverse2_1_0)((function() use ($f_2) {
  $__fn = function($a_5 = null, $s_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_2)($s_6))($a_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($xs_4))($s0_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_scanr'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $mapAccumR1_1_0 = (($GLOBALS['Data_Traversable_mapAccumR'] ?? \PhpursThunks::eval('Data_Traversable_mapAccumR')))($dictTraversable_0);
  $__res = (function() use ($mapAccumR1_1_0) {
  $__fn = function($f_2 = null, $b0_3 = null, $xs_4 = null) use ($mapAccumR1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumR1_1_0)((function() use ($f_2) {
  $__fn = function($b_5 = null, $a_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $b__prime___7_1 = (($f_2)($a_6))($b_5);
  $__res = ["accum" => $b__prime___7_1, "value" => $b__prime___7_1];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))['value'];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_mapAccumL'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $traverse2_1_0 = (($dictTraversable_0)['traverse'])(($GLOBALS['Data_Traversable_Accum_Internal_applicativeStateL'] ?? \PhpursThunks::eval('Data_Traversable_Accum_Internal_applicativeStateL')));
  $__res = (function() use ($traverse2_1_0) {
  $__fn = function($f_2 = null, $s0_3 = null, $xs_4 = null) use ($traverse2_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = ((($traverse2_1_0)((function() use ($f_2) {
  $__fn = function($a_5 = null, $s_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($f_2)($s_6))($a_5);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($xs_4))($s0_3);
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_scanl'] = function() { $v = function($dictTraversable_0 = null) {
  $__num = \func_num_args();
  $mapAccumL1_1_0 = (($GLOBALS['Data_Traversable_mapAccumL'] ?? \PhpursThunks::eval('Data_Traversable_mapAccumL')))($dictTraversable_0);
  $__res = (function() use ($mapAccumL1_1_0) {
  $__fn = function($f_2 = null, $b0_3 = null, $xs_4 = null) use ($mapAccumL1_1_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (((($mapAccumL1_1_0)((function() use ($f_2) {
  $__fn = function($b_5 = null, $a_6 = null) use ($f_2, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $b__prime___7_1 = (($f_2)($b_5))($a_6);
  $__res = ["accum" => $b__prime___7_1, "value" => $b__prime___7_1];
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})()))($b0_3))($xs_4))['value'];
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_Traversable_for'] = function() { $v = (function() {
  $__fn = function($dictApplicative_0 = null, $dictTraversable_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $traverse2_2_0 = (($dictTraversable_1)['traverse'])($dictApplicative_0);
  $__res = (function() use ($traverse2_2_0) {
  $__fn = function($x_3 = null, $f_4 = null) use ($traverse2_2_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($traverse2_2_0)($f_4))($x_3);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})();
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };
$ffi_Data_Traversable = \call_user_func(function() {
  $exports = [];
$traverseArrayImpl = function($apply, $map = null, $pure = null, $f = null, $array = null) use (&$traverseArrayImpl) {
    if (\func_num_args() < 5) {
        $__args = \func_get_args();
        return function(...$more) use ($__args, &$traverseArrayImpl) {
            return $traverseArrayImpl(...\array_merge($__args, $more));
        };
    }

    $array1 = function ($a) { return [$a]; };
    $array2 = function ($a) { return function ($b) use ($a) { return [$a, $b]; }; };
    $array3 = function ($a) { return function ($b) use ($a) { return function ($c) use ($a, $b) { return [$a, $b, $c]; }; }; };
    $concat2 = function ($xs) { return function ($ys) use ($xs) { return \array_merge($xs, $ys); }; };
    
    $go = function ($bot, $top) use (&$go, $array, $apply, $map, $pure, $f, $array1, $array2, $array3, $concat2) {
        switch ($top - $bot) {
            case 0:
                return $pure([]);
            case 1:
                $f1 = $f($array[$bot]);
                $map1 = $map($array1);
                return $map1($f1);
            case 2:
                $f1 = $f($array[$bot]);
                $f2 = $f($array[$bot + 1]);
                $map1 = $map($array2);
                $map2 = $map1($f1);
                return $apply($map2)($f2);
            case 3:
                $f1 = $f($array[$bot]);
                $f2 = $f($array[$bot + 1]);
                $f3 = $f($array[$bot + 2]);
                $map1 = $map($array3);
                $map2 = $map1($f1);
                $app1 = $apply($map2)($f2);
                return $apply($app1)($f3);
            default:
                $pivot = $bot + floor(($top - $bot) / 4) * 2;
                $go1 = $go($bot, $pivot);
                $go2 = $go($pivot, $top);
                $map1 = $map($concat2);
                $map2 = $map1($go1);
                return $apply($map2)($go2);
        }
    };
    return $go(0, \count($array));
};
$exports['traverseArrayImpl'] = $traverseArrayImpl;

return $exports;
  return $exports;
});
\PhpursThunks::$thunks['Data_Traversable_traverseArrayImpl'] = function() use (&$ffi_Data_Traversable) { return $ffi_Data_Traversable['traverseArrayImpl']; };




























