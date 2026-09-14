"""Independent values for the native FFI contracts; no benchmark clocks."""
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
NOMINAL = {
    'AstTree': (3, 7), 'Fib': (10, 55), 'ListOps': (900, 202950),
    'TCO': (100000, 100000), 'Records': (10000, 20000), 'Ackermann': (3, 125),
    'Church': (10, 100000), 'Primes': (500, 21536), 'RBTree': (100000, 22),
    'Polymorphism': (10000000, 10000000), 'StateMonad': (60, 1200),
    'LazyEvaluation': (1000, 1000000), 'ArrayOps': (900, 202950), 'RowToList': (0, 5),
}


def tree_balance(color, left, key, right):
    def red(tree):
        return tree is not None and tree[0] == 'R'
    if color == 'B':
        if red(left):
            _, a, x, b = left
            if red(a):
                _, aa, ax, ab = a
                return ('R', ('B', aa, ax, ab), x, ('B', b, key, right))
            if red(b):
                _, ba, bx, bb = b
                return ('R', ('B', a, x, ba), bx, ('B', bb, key, right))
        if red(right):
            _, a, x, b = right
            if red(a):
                _, aa, ax, ab = a
                return ('R', ('B', left, key, aa), ax, ('B', ab, x, b))
            if red(b):
                _, ba, bx, bb = b
                return ('R', ('B', left, key, a), x, ('B', ba, bx, bb))
    return (color, left, key, right)


def tree_insert(key, tree):
    def insert(value):
        if value is None:
            return ('R', None, key, None)
        color, left, old, right = value
        if key < old:
            return tree_balance(color, insert(left), old, right)
        if key > old:
            return tree_balance(color, left, old, insert(right))
        return value
    _, left, value, right = insert(tree)
    return ('B', left, value, right)


def tree_depth(tree):
    return 0 if tree is None else 1 + max(tree_depth(tree[1]), tree_depth(tree[3]))


def reference(name, n):
    if name == 'AstTree':
        value = 1
        for depth in range(1, n + 1):
            value = (depth + 1) * value - 1
        return value
    if name == 'Fib':
        a, b = 0, 1
        for _ in range(n):
            a, b = b, a + b
        return a
    if name == 'ListOps':
        return sum(range(2, n + 1, 2))
    if name == 'ArrayOps':
        values = range(1, n + 1) if n >= 1 else range(1, n - 1, -1)
        return sum(value for value in values if value % 2 == 0)
    if name == 'TCO':
        return (n // 3) * 3 + (n % 3) * (n % 3 + 1) // 2
    if name == 'Records':
        return (n // 5) * 10 + (n % 5) * (n % 5 + 1) // 2
    if name == 'Ackermann':
        # Closed forms A(m, 4), independent of the recursive native kernels.
        return [5, 6, 11, 125][n]
    if name == 'Church':
        return n ** 5
    if name == 'Primes':
        return sum(p for p in range(2, n + 1)
                   if all(p % d for d in range(2, int(p ** 0.5) + 1)))
    if name == 'RBTree':
        tree = None
        for key in range(n, 0, -1):
            tree = tree_insert(key, tree)
        return tree_depth(tree)
    if name == 'Polymorphism':
        return n
    if name == 'StateMonad':
        return 20 * n  # Shared FFI wrapper passes depth, unlike Test.StateMonad.act.
    if name == 'LazyEvaluation':
        return 1000 * n  # Number of repetitions; fixed depth 1000 in the source.
    if name == 'RowToList':
        return 5
    raise ValueError(name)


def cases():
    result = []
    for name, (argument, expected) in NOMINAL.items():
        arguments = range(4) if name == 'Ackermann' else range(7)
        if name == 'ArrayOps':
            arguments = range(-3, 7)
        for value in arguments:
            if value != argument:
                result.append((name, value, reference(name, value)))
        result.append((name, argument, expected))
    return result
