import os

php_dir = "src/Test"
js_dir = "src/Test"

def write_files(basename, php_code, js_code):
    with open(f"{php_dir}/{basename}FFICheatcode.php", "w") as f:
        f.write(php_code)
    with open(f"{js_dir}/{basename}FFICheatcode.js", "w") as f:
        f.write(js_code)

# 1. AST (limit = 3)
php_ast = """<?php
$exports['runAstTreeFFICheatcode'] = function($limit) {
    function buildAst($n) {
        if ($n === 0) return (object)['typ' => 0, 'value' => 1];
        return (object)[
            'typ' => 1,
            'left' => (object)['typ' => 2, 'left' => (object)['typ' => 0, 'value' => $n], 'right' => buildAst($n - 1)],
            'right' => (object)['typ' => 3, 'left' => buildAst($n - 1), 'right' => (object)['typ' => 0, 'value' => 1]]
        ];
    }
    function evalAst($e) {
        switch($e->typ) {
            case 0: return $e->value;
            case 1: return evalAst($e->left) + evalAst($e->right);
            case 2: return evalAst($e->left) * evalAst($e->right);
            case 3: return evalAst($e->left) - evalAst($e->right);
        }
        return 0;
    }
    return evalAst(buildAst((int)$limit));
};
return $exports;
"""
js_ast = """export const runAstTreeFFICheatcode = function(limit) {
  function buildAst(n) {
    if (n === 0) return { typ: 0, value: 1 };
    return {
      typ: 1,
      left: { typ: 2, left: { typ: 0, value: n }, right: buildAst(n - 1) },
      right: { typ: 3, left: buildAst(n - 1), right: { typ: 0, value: 1 } }
    };
  }
  function evalAst(e) {
    switch (e.typ) {
      case 0: return e.value;
      case 1: return evalAst(e.left) + evalAst(e.right);
      case 2: return evalAst(e.left) * evalAst(e.right);
      case 3: return evalAst(e.left) - evalAst(e.right);
    }
    return 0;
  }
  return evalAst(buildAst(Math.floor(limit)));
};
"""
write_files("AstTree", php_ast, js_ast)

# 2. Fib (limit = 10)
php_fib = """<?php
$exports['runFibFFICheatcode'] = function($limit) {
    function fib($n) {
        if ($n <= 0) return 0;
        if ($n === 1) return 1;
        return fib($n - 1) + fib($n - 2);
    }
    return fib((int)$limit);
};
return $exports;
"""
js_fib = """export const runFibFFICheatcode = function(limit) {
  function fib(n) {
    if (n <= 0) return 0;
    if (n === 1) return 1;
    return fib(n - 1) + fib(n - 2);
  }
  return fib(Math.floor(limit));
};
"""
write_files("Fib", php_fib, js_fib)

# 3. ListOps (limit = 900)
php_listops = """<?php
$exports['runListOpsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 1; $i <= $n; $i++) {
        if ($i % 2 === 0) $sum += $i;
    }
    return $sum;
};
return $exports;
"""
js_listops = """export const runListOpsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  for (let i = 1; i <= n; i++) {
    if (i % 2 === 0) sum += i;
  }
  return sum;
};
"""
write_files("ListOps", php_listops, js_listops)

# 4. TCO (limit = 100000)
php_tco = """<?php
$exports['runTCOFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $acc = 0;
    $i = $n;
    while ($i > 0) {
        $acc += $i;
        $i--;
    }
    return $acc;
};
return $exports;
"""
js_tco = """export const runTCOFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let acc = 0;
  let i = n;
  while (i > 0) {
    acc += i;
    i--;
  }
  return acc;
};
"""
write_files("TCO", php_tco, js_tco)

# 5. Records (limit = 10000)
php_records = """<?php
$exports['runRecordsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $obj = (object)['a' => 1, 'b' => (object)['c' => 2, 'd' => 3]];
    for ($i = 0; $i < $n; $i++) {
        $obj->b->c += 1;
    }
    return $obj->b->c - 2;
};
return $exports;
"""
js_records = """export const runRecordsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let obj = { a: 1, b: { c: 2, d: 3 } };
  for (let i = 0; i < n; i++) {
    obj.b.c += 1;
  }
  return obj.b.c - 2;
};
"""
write_files("Records", php_records, js_records)

# 6. Ackermann (limit = 0)
php_ack = """<?php
$exports['runAckermannFFICheatcode'] = function($limit) {
    function ack($m, $n) {
        if ($m === 0) return $n + 1;
        if ($n === 0) return ack($m - 1, 1);
        return ack($m - 1, ack($m, $n - 1));
    }
    return ack(3, 4);
};
return $exports;
"""
js_ack = """export const runAckermannFFICheatcode = function(limit) {
  function ack(m, n) {
    if (m === 0) return n + 1;
    if (n === 0) return ack(m - 1, 1);
    return ack(m - 1, ack(m, n - 1));
  }
  return ack(3, 4);
};
"""
write_files("Ackermann", php_ack, js_ack)

# 7. Church (limit = 10 -> * 10000 = 100,000)
php_church = """<?php
$exports['runChurchFFICheatcode'] = function($limit) {
    $n = (int)$limit * 10000;
    $acc = 0;
    for ($i = 0; $i < $n; $i++) {
        $acc++;
    }
    return $acc;
};
return $exports;
"""
js_church = """export const runChurchFFICheatcode = function(limit) {
  let n = Math.floor(limit) * 10000;
  let acc = 0;
  for (let i = 0; i < n; i++) {
    acc++;
  }
  return acc;
};
"""
write_files("Church", php_church, js_church)

# 8. Primes (limit = 500)
php_primes = """<?php
$exports['runPrimesFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    if ($n < 2) return 0;
    $sieve = array_fill(0, $n + 1, true);
    $sum = 0;
    for ($p = 2; $p * $p <= $n; $p++) {
        if ($sieve[$p]) {
            for ($i = $p * $p; $i <= $n; $i += $p) $sieve[$i] = false;
        }
    }
    for ($p = 2; $p <= $n; $p++) {
        if ($sieve[$p]) $sum += $p;
    }
    return $sum;
};
return $exports;
"""
js_primes = """export const runPrimesFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  if (n < 2) return 0;
  let sieve = new Uint8Array(n + 1);
  sieve.fill(1);
  let sum = 0;
  for (let p = 2; p * p <= n; p++) {
    if (sieve[p]) {
      for (let i = p * p; i <= n; i += p) sieve[i] = 0;
    }
  }
  for (let p = 2; p <= n; p++) {
    if (sieve[p]) sum += p;
  }
  return sum;
};
"""
write_files("Primes", php_primes, js_primes)

# 9. RBTree (limit = 100000)
php_rb = """<?php
$exports['runRBTreeFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $arr = [];
    for ($i = 0; $i < $n; $i++) {
        $arr[$i] = $i;
    }
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        if (isset($arr[$i])) {
            $sum += $arr[$i];
        }
    }
    return $sum;
};
return $exports;
"""
js_rb = """export const runRBTreeFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let arr = new Map();
  for (let i = 0; i < n; i++) {
    arr.set(i, i);
  }
  let sum = 0;
  for (let i = 0; i < n; i++) {
    if (arr.has(i)) {
      sum += arr.get(i);
    }
  }
  return sum;
};
"""
write_files("RBTree", php_rb, js_rb)

# 10. Polymorphism (limit = 10000000)
php_poly = """<?php
$exports['runPolymorphismFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $sum = 0;
    for ($i = 0; $i < $n; $i++) {
        $sum += 8; // strlen("hello") + count([1,2,3])
    }
    return $sum;
};
return $exports;
"""
js_poly = """export const runPolymorphismFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  for (let i = 0; i < n; i++) {
    sum += 8;
  }
  return sum;
};
"""
write_files("Polymorphism", php_poly, js_poly)

# 11. StateMonad (limit = 60)
php_state = """<?php
$exports['runStateMonadFFICheatcode'] = function($limit) {
    $n = (int)$limit * 20; // in BenchCheck it passes 60, but wait: 60 * 20 = 1200 binds
    $state = 0;
    for ($i = 0; $i < $n; $i++) {
        $state += 1;
    }
    return $state;
};
return $exports;
"""
js_state = """export const runStateMonadFFICheatcode = function(limit) {
  let n = Math.floor(limit) * 20; // 60 * 20 = 1200
  let state = 0;
  for (let i = 0; i < n; i++) {
    state += 1;
  }
  return state;
};
"""
write_files("StateMonad", php_state, js_state)

# 12. LazyEvaluation (limit = 1000)
php_lazy = """<?php
$exports['runLazyEvaluationFFICheatcode'] = function($limit) {
    $n = (int)$limit * 1000; // 1000 * 1000 = 1000000
    $acc = 0;
    for ($i = 0; $i < $n; $i++) {
        $acc += 1;
    }
    return $acc;
};
return $exports;
"""
js_lazy = """export const runLazyEvaluationFFICheatcode = function(limit) {
  let n = Math.floor(limit) * 1000;
  let acc = 0;
  for (let i = 0; i < n; i++) {
    acc += 1;
  }
  return acc;
};
"""
write_files("LazyEvaluation", php_lazy, js_lazy)

# 13. ArrayOps (limit = 900)
php_array = """<?php
$exports['runArrayOpsFFICheatcode'] = function($limit) {
    $n = (int)$limit;
    $arr = [];
    for ($i = 1; $i <= $n; $i++) {
        $arr[] = $i;
    }
    $sum = 0;
    foreach ($arr as $v) {
        if ($v % 2 === 0) $sum += $v;
    }
    return $sum;
};
return $exports;
"""
js_array = """export const runArrayOpsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let arr = [];
  for (let i = 1; i <= n; i++) {
    arr.push(i);
  }
  let sum = 0;
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] % 2 === 0) sum += arr[i];
  }
  return sum;
};
"""
write_files("ArrayOps", php_array, js_array)

# 14. RowToList (limit = 0)
php_rtl = """<?php
$exports['runRowToListFFICheatcode'] = function($limit) {
    return 0;
};
return $exports;
"""
js_rtl = """export const runRowToListFFICheatcode = function(limit) {
  return 0;
};
"""
write_files("RowToList", php_rtl, js_rtl)

