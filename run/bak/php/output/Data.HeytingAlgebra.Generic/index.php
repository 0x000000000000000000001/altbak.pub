<?php

namespace Data\HeytingAlgebra\Generic;

// ALL IMPORTS: Data.Function, Data.Generic.Rep, Data.HeytingAlgebra, Data.HeytingAlgebra.Generic, Prelude, Prim
// TO REQUIRE: Data.Function, Data.Generic.Rep, Data.HeytingAlgebra, Data.HeytingAlgebra.Generic, Prelude
require_once __DIR__ . '/../Data.Function/index.php';
require_once __DIR__ . '/../Data.Generic.Rep/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra/index.php';
require_once __DIR__ . '/../Data.HeytingAlgebra.Generic/index.php';
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
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericTT__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericTT__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericTT'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericHeytingAlgebra_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictGeneric_0)->{'to'})(($dictGenericHeytingAlgebra_1)->{'genericTT__prime__'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericNot__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericNot__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericNot'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericHeytingAlgebra_1 = null, $x_2 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 3) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 3);
  }
  $__res = (($dictGeneric_0)->{'to'})((($dictGenericHeytingAlgebra_1)->{'genericNot__prime__'})((($dictGeneric_0)->{'from'})($x_2)));
  goto __end;;
  __end:
  return $__num > 3 ? $__res(...\array_slice(\func_get_args(), 3)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericImplies__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericImplies__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericImplies'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericHeytingAlgebra_1 = null, $x_2 = null, $y_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($dictGeneric_0)->{'to'})(((($dictGenericHeytingAlgebra_1)->{'genericImplies__prime__'})((($dictGeneric_0)->{'from'})($x_2)))((($dictGeneric_0)->{'from'})($y_3)));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments'] = function() { $v = (object)["genericFF__prime__" => new Phpurs_Data0("NoArguments"), "genericTT__prime__" => new Phpurs_Data0("NoArguments"), "genericImplies__prime__" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data0("NoArguments");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericConj__prime__" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data0("NoArguments");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericDisj__prime__" => (function() {
  $__fn = function($v_0 = null, $v1_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data0("NoArguments");
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericNot__prime__" => function($v_0 = null) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data0("NoArguments");
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}]; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument'] = function() { $v = function($dictHeytingAlgebra_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["genericFF__prime__" => ($dictHeytingAlgebra_0)->{'ff'}, "genericTT__prime__" => ($dictHeytingAlgebra_0)->{'tt'}, "genericImplies__prime__" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_0)->{'implies'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericConj__prime__" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_0)->{'conj'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericDisj__prime__" => (function() use ($dictHeytingAlgebra_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictHeytingAlgebra_0)->{'disj'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericNot__prime__" => function($v_1 = null) use ($dictHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = (($dictHeytingAlgebra_0)->{'not'})($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericFF__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericFF__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericFF'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericHeytingAlgebra_1 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = (($dictGeneric_0)->{'to'})(($dictGenericHeytingAlgebra_1)->{'genericFF__prime__'});
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericDisj__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericDisj__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericDisj'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericHeytingAlgebra_1 = null, $x_2 = null, $y_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($dictGeneric_0)->{'to'})(((($dictGenericHeytingAlgebra_1)->{'genericDisj__prime__'})((($dictGeneric_0)->{'from'})($x_2)))((($dictGeneric_0)->{'from'})($y_3)));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericConj__prime__'] = function() { $v = function($dict_0 = null) {
  $__num = \func_num_args();
  $__res = ($dict_0)->{'genericConj__prime__'};
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor'] = function() { $v = function($dictGenericHeytingAlgebra_0 = null) {
  $__num = \func_num_args();
  $__res = (object)["genericFF__prime__" => ($dictGenericHeytingAlgebra_0)->{'genericFF__prime__'}, "genericTT__prime__" => ($dictGenericHeytingAlgebra_0)->{'genericTT__prime__'}, "genericImplies__prime__" => (function() use ($dictGenericHeytingAlgebra_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictGenericHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictGenericHeytingAlgebra_0)->{'genericImplies__prime__'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericConj__prime__" => (function() use ($dictGenericHeytingAlgebra_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictGenericHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictGenericHeytingAlgebra_0)->{'genericConj__prime__'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericDisj__prime__" => (function() use ($dictGenericHeytingAlgebra_0) {
  $__fn = function($v_1 = null, $v1_2 = null) use ($dictGenericHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = ((($dictGenericHeytingAlgebra_0)->{'genericDisj__prime__'})($v_1))($v1_2);
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericNot__prime__" => function($v_1 = null) use ($dictGenericHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = (($dictGenericHeytingAlgebra_0)->{'genericNot__prime__'})($v_1);
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}];
  goto __end;;
  __end:
  return $__num > 1 ? $__res(...\array_slice(\func_get_args(), 1)) : $__res;
}; return $v; };
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct'] = function() { $v = function($dictGenericHeytingAlgebra_0 = null) {
  $__num = \func_num_args();
  $genericFF__prime__1_1_0 = ($dictGenericHeytingAlgebra_0)->{'genericFF__prime__'};
  $genericTT__prime__1_2_1 = ($dictGenericHeytingAlgebra_0)->{'genericTT__prime__'};
  $__res = function($dictGenericHeytingAlgebra1_3 = null) use ($dictGenericHeytingAlgebra_0, $genericFF__prime__1_1_0, $genericTT__prime__1_2_1) {
  $__num = \func_num_args();
  $__res = (object)["genericFF__prime__" => new Phpurs_Data2("Product", $genericFF__prime__1_1_0, ($dictGenericHeytingAlgebra1_3)->{'genericFF__prime__'}), "genericTT__prime__" => new Phpurs_Data2("Product", $genericTT__prime__1_2_1, ($dictGenericHeytingAlgebra1_3)->{'genericTT__prime__'}), "genericImplies__prime__" => (function() use ($dictGenericHeytingAlgebra1_3, $dictGenericHeytingAlgebra_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictGenericHeytingAlgebra1_3, $dictGenericHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Product", ((($dictGenericHeytingAlgebra_0)->{'genericImplies__prime__'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictGenericHeytingAlgebra1_3)->{'genericImplies__prime__'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericConj__prime__" => (function() use ($dictGenericHeytingAlgebra1_3, $dictGenericHeytingAlgebra_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictGenericHeytingAlgebra1_3, $dictGenericHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Product", ((($dictGenericHeytingAlgebra_0)->{'genericConj__prime__'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictGenericHeytingAlgebra1_3)->{'genericConj__prime__'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericDisj__prime__" => (function() use ($dictGenericHeytingAlgebra1_3, $dictGenericHeytingAlgebra_0) {
  $__fn = function($v_4 = null, $v1_5 = null) use ($dictGenericHeytingAlgebra1_3, $dictGenericHeytingAlgebra_0, &$__fn) {
  $__num = \func_num_args();
  if ($__num < 2) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 2);
  }
  $__res = new Phpurs_Data2("Product", ((($dictGenericHeytingAlgebra_0)->{'genericDisj__prime__'})(($v_4)->{'value0'}))(($v1_5)->{'value0'}), ((($dictGenericHeytingAlgebra1_3)->{'genericDisj__prime__'})(($v_4)->{'value1'}))(($v1_5)->{'value1'}));
  goto __end;;
  __end:
  return $__num > 2 ? $__res(...\array_slice(\func_get_args(), 2)) : $__res;
  };
  return $__fn;
})(), "genericNot__prime__" => function($v_4 = null) use ($dictGenericHeytingAlgebra1_3, $dictGenericHeytingAlgebra_0) {
  $__num = \func_num_args();
  $__res = new Phpurs_Data2("Product", (($dictGenericHeytingAlgebra_0)->{'genericNot__prime__'})(($v_4)->{'value0'}), (($dictGenericHeytingAlgebra1_3)->{'genericNot__prime__'})(($v_4)->{'value1'}));
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
\PhpursThunks::$thunks['Data_HeytingAlgebra_Generic_genericConj'] = function() { $v = (function() {
  $__fn = function($dictGeneric_0 = null, $dictGenericHeytingAlgebra_1 = null, $x_2 = null, $y_3 = null) use (&$__fn) {
  $__num = \func_num_args();
  if ($__num < 4) {
    return phpurs_curry_fallback($__fn, \func_get_args(), 4);
  }
  $__res = (($dictGeneric_0)->{'to'})(((($dictGenericHeytingAlgebra_1)->{'genericConj__prime__'})((($dictGeneric_0)->{'from'})($x_2)))((($dictGeneric_0)->{'from'})($y_3)));
  goto __end;;
  __end:
  return $__num > 4 ? $__res(...\array_slice(\func_get_args(), 4)) : $__res;
  };
  return $__fn;
})(); return $v; };
$GLOBALS['Prim_undefined'] = function() { throw new \Exception("undefined"); };


















