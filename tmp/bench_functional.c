/*
 * Functional C reference for the 14 core PureScript programs.
 * Contract: tmp/fp_reference_contract.md. Normal compiler optimization is
 * intentional; there are no barriers or volatile accesses in the kernels.
 *
 * Native C mapping: Int32 arithmetic, immutable typed structs/lists/trees,
 * typed function-pointer/environment closures, and typed dictionaries.
 * Captured environments and recursive nodes use malloc. A per-action linked
 * allocation list frees every allocation before the timed action returns.
 * This includes allocation, tracking and reclamation in timings; it is not a
 * claim that malloc/free and a tracing GC have equivalent costs. No arena is
 * preallocated or reused outside the timed interval. Non-escaping records and
 * dictionaries use ordinary C values, which clang may specialize/eliminate.
 *
 * Build: clang -std=c11 -O3 -Wall -Wextra -Werror bench_functional.c -o benchmark
 * Check: ./benchmark --verify (no timed benchmark campaign)
 */
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <inttypes.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdatomic.h>
#include <time.h>

typedef int32_t Int;

static Int from_bits(uint32_t value) {
    Int result;
    memcpy(&result, &value, sizeof result);
    return result;
}
static Int add_int(Int x, Int y) { return from_bits((uint32_t)x + (uint32_t)y); }
static Int sub_int(Int x, Int y) { return from_bits((uint32_t)x - (uint32_t)y); }
static Int mul_int(Int x, Int y) { return from_bits((uint32_t)x * (uint32_t)y); }
/* All divisors used by these kernels are positive. */
static Int mod_int(Int x, Int positive_divisor) {
    Int result = x % positive_divisor;
    return result < 0 ? add_int(result, positive_divisor) : result;
}

typedef union Allocation Allocation;
union Allocation {
    struct { Allocation *next; } link;
    max_align_t alignment;
};
static Allocation *allocations;
#ifdef VERIFY_ALLOCATIONS
static size_t allocation_count;
static size_t allocation_bytes;
#endif

static void *allocate(size_t size) {
    if (size > SIZE_MAX - sizeof(Allocation)) {
        fputs("Allocation size overflow\n", stderr);
        exit(1);
    }
    Allocation *block = malloc(sizeof(Allocation) + size);
    if (!block) {
        fputs("Allocation failed\n", stderr);
        exit(1);
    }
    block->link.next = allocations;
    allocations = block;
#ifdef VERIFY_ALLOCATIONS
    allocation_count++;
    allocation_bytes += size;
#endif
    return block + 1;
}

static void release_allocations(void) {
    while (allocations) {
        Allocation *next = allocations->link.next;
        free(allocations);
        allocations = next;
    }
}

/* AST: construct Val/Add/Mul/Sub, then recursively interpret. */
typedef enum { EXPR_VAL, EXPR_ADD, EXPR_MUL, EXPR_SUB } ExprTag;
typedef struct Expr Expr;
struct Expr {
    ExprTag tag;
    union { Int value; struct { Expr *left, *right; } pair; } data;
};
static Expr *value_expr(Int value) {
    Expr *node = allocate(sizeof *node);
    *node = (Expr){.tag = EXPR_VAL, .data.value = value};
    return node;
}
static Expr *binary_expr(ExprTag tag, Expr *left, Expr *right) {
    Expr *node = allocate(sizeof *node);
    *node = (Expr){.tag = tag, .data.pair = {left, right}};
    return node;
}
static Expr *build_expr(Int n) {
    if (n == 0) return value_expr(1);
    Expr *number = value_expr(n);
    Expr *left_child = build_expr(sub_int(n, 1));
    Expr *left = binary_expr(EXPR_MUL, number, left_child);
    Expr *right_child = build_expr(sub_int(n, 1));
    Expr *one = value_expr(1);
    Expr *right = binary_expr(EXPR_SUB, right_child, one);
    return binary_expr(EXPR_ADD, left, right);
}
static Int eval_expr(const Expr *node) {
    if (node->tag == EXPR_VAL) return node->data.value;
    Int left = eval_expr(node->data.pair.left);
    Int right = eval_expr(node->data.pair.right);
    switch (node->tag) {
        case EXPR_ADD: return add_int(left, right);
        case EXPR_MUL: return mul_int(left, right);
        case EXPR_SUB: return sub_int(left, right);
        case EXPR_VAL: break;
    }
    abort();
}
static Int run_ast(Int n) { return eval_expr(build_expr(n)); }

static Int fibonacci(Int n) {
    if (n == 0) return 0;
    if (n == 1) return 1;
    Int left = fibonacci(sub_int(n, 1));
    Int right = fibonacci(sub_int(n, 2));
    return add_int(left, right);
}

/* Strict lists and generic (Int-instantiated) folds/predicates. */
typedef struct List List;
struct List { Int head; List *tail; };
typedef struct { Int (*apply)(const void *, Int, Int); const void *env; } Fold;
typedef struct { bool (*apply)(const void *, Int); const void *env; } Predicate;

static List *cons(Int head, List *tail) {
    List *node = allocate(sizeof *node);
    *node = (List){head, tail};
    return node;
}
static List *list_range(Int start, Int end) {
    List *result = NULL;
    for (Int current = end; current >= start; current = sub_int(current, 1)) {
        result = cons(current, result);
        if (current == start) break;
    }
    return result;
}
static List *filter_evens(List *input) {
    List *result = NULL;
    for (; input; input = input->tail)
        if (mod_int(input->head, 2) == 0) result = cons(input->head, result);
    return result; /* ListOps intentionally retains reversed order. */
}
static Int list_fold(Fold f, Int initial, const List *input) {
    Int result = initial;
    for (; input; input = input->tail) result = f.apply(f.env, result, input->head);
    return result;
}
static Int add_fold(const void *env, Int x, Int y) {
    (void)env;
    return add_int(x, y);
}
static Int run_list(Int n) {
    List *range = list_range(1, n);
    List *selected = filter_evens(range);
    return list_fold((Fold){add_fold, NULL}, 0, selected);
}

static Int tail_rec(Int n, Int acc) {
    if (n == 0) return acc;
    return tail_rec(sub_int(n, 1), add_int(acc, mod_int(n, 3)));
}
static Int run_tco(Int n) { return tail_rec(n, 0); }

/* Nested value copies implement immutable updates at all three levels. */
typedef struct { Int e, f; } RecordD;
typedef struct { Int c; RecordD d; } RecordB;
typedef struct { Int a; RecordB b; } RecordA;
static RecordA update_record(Int n, RecordA old) {
    if (n == 0) return old;
    RecordD d = {add_int(old.b.d.e, 3), add_int(old.b.d.f, mod_int(n, 5))};
    RecordB b = {add_int(old.b.c, 2), d};
    RecordA next = {add_int(old.a, 1), b};
    return update_record(sub_int(n, 1), next);
}
static Int run_records(Int n) {
    RecordA initial = {0, {0, {0, 0}}};
    return update_record(n, initial).b.d.f;
}

static Int ackermann(Int m, Int n) {
    if (m == 0) return add_int(n, 1);
    if (n == 0) return ackermann(sub_int(m, 1), 1);
    Int inside = ackermann(m, sub_int(n, 1));
    return ackermann(sub_int(m, 1), inside);
}
static Int run_ackermann(Int m) { return ackermann(m, 4); }

/* Fully curried typed Church closures, including partial applications. */
typedef struct { Int (*apply)(const void *, Int); const void *env; } Endo;
typedef struct { Endo (*apply)(const void *, Endo); const void *env; } Church;
typedef struct { Church previous; } SuccessorEnv;
typedef struct { Church previous; Endo step; } SuccessorCall;
typedef struct { Church left, right; } MultiplyEnv;
typedef struct { Church left, right; Endo step; } MultiplyCall;

static Int identity_step(const void *env, Int x) { (void)env; return x; }
static Int increment_step(const void *env, Int x) { (void)env; return add_int(x, 1); }
static Endo zero_church_apply(const void *env, Endo step) {
    (void)env; (void)step;
    return (Endo){identity_step, NULL};
}
static Church zero_church(void) { return (Church){zero_church_apply, NULL}; }
static Int successor_call(const void *context, Int x) {
    const SuccessorCall *env = context;
    Endo previous = env->previous.apply(env->previous.env, env->step);
    Int result = previous.apply(previous.env, x);
    return env->step.apply(env->step.env, result);
}
static Endo successor_apply(const void *context, Endo step) {
    const SuccessorEnv *env = context;
    SuccessorCall *call = allocate(sizeof *call);
    *call = (SuccessorCall){env->previous, step};
    return (Endo){successor_call, call};
}
static Church successor_church(Church previous) {
    SuccessorEnv *env = allocate(sizeof *env);
    *env = (SuccessorEnv){previous};
    return (Church){successor_apply, env};
}
static Church from_int_church(Int n) {
    if (n == 0) return zero_church();
    Church previous = from_int_church(sub_int(n, 1));
    return successor_church(previous);
}
static Int multiply_call(const void *context, Int x) {
    const MultiplyCall *env = context;
    Endo right = env->right.apply(env->right.env, env->step);
    Endo left = env->left.apply(env->left.env, right);
    return left.apply(left.env, x);
}
static Endo multiply_apply(const void *context, Endo step) {
    const MultiplyEnv *env = context;
    MultiplyCall *call = allocate(sizeof *call);
    *call = (MultiplyCall){env->left, env->right, step};
    return (Endo){multiply_call, call};
}
static Church multiply_church(Church left, Church right) {
    MultiplyEnv *env = allocate(sizeof *env);
    *env = (MultiplyEnv){left, right};
    return (Church){multiply_apply, env};
}
static Church church_100(Int n) {
    Church left = from_int_church(n);
    Church right = from_int_church(n);
    return multiply_church(left, right);
}
static Church church_10k(Int n) {
    Church left = church_100(n);
    Church right = church_100(n);
    return multiply_church(left, right);
}
static Int run_church(Int n) {
    Church left = church_10k(n);
    Church right = from_int_church(n);
    Church number = multiply_church(left, right);
    Endo action = number.apply(number.env, (Endo){increment_step, NULL});
    return action.apply(action.env, 0);
}

/* Strict list sieve, including the captured predicate and filter reversal. */
static List *reverse_list(List *input) {
    List *result = NULL;
    for (; input; input = input->tail) result = cons(input->head, result);
    return result;
}
static List *filter_list(Predicate predicate, List *input) {
    List *selected = NULL;
    for (; input; input = input->tail)
        if (predicate.apply(predicate.env, input->head)) selected = cons(input->head, selected);
    return reverse_list(selected);
}
static bool not_multiple(const void *env, Int x) {
    return mod_int(x, *(const Int *)env) != 0;
}
static List *sieve(List *input) {
    if (!input) return NULL;
    Int prime = input->head;
    Predicate predicate = {not_multiple, &prime};
    List *remaining = filter_list(predicate, input->tail);
    List *tail = sieve(remaining);
    return cons(prime, tail);
}
static Int run_primes(Int n) {
    List *range = list_range(2, n);
    List *primes = sieve(range);
    return list_fold((Fold){add_fold, NULL}, 0, primes);
}

/* Persistent Okasaki red-black tree. No mutation of existing nodes. */
typedef enum { RED, BLACK } Color;
typedef struct Tree Tree;
struct Tree { Color color; Tree *left; Int value; Tree *right; };
static Tree *tree(Color color, Tree *left, Int value, Tree *right) {
    Tree *node = allocate(sizeof *node);
    *node = (Tree){color, left, value, right};
    return node;
}
static bool red(const Tree *node) { return node && node->color == RED; }
static Tree *balance(Color color, Tree *left, Int value, Tree *right) {
    if (color == BLACK && red(left) && red(left->left)) {
        Tree *a = tree(BLACK, left->left->left, left->left->value, left->left->right);
        Tree *b = tree(BLACK, left->right, value, right);
        return tree(RED, a, left->value, b);
    }
    if (color == BLACK && red(left) && red(left->right)) {
        Tree *a = tree(BLACK, left->left, left->value, left->right->left);
        Tree *b = tree(BLACK, left->right->right, value, right);
        return tree(RED, a, left->right->value, b);
    }
    if (color == BLACK && red(right) && red(right->left)) {
        Tree *a = tree(BLACK, left, value, right->left->left);
        Tree *b = tree(BLACK, right->left->right, right->value, right->right);
        return tree(RED, a, right->left->value, b);
    }
    if (color == BLACK && red(right) && red(right->right)) {
        Tree *a = tree(BLACK, left, value, right->left);
        Tree *b = tree(BLACK, right->right->left, right->right->value, right->right->right);
        return tree(RED, a, right->value, b);
    }
    return tree(color, left, value, right);
}
static Tree *insert_inner(Int value, Tree *old) {
    if (!old) return tree(RED, NULL, value, NULL);
    if (value < old->value) {
        Tree *left = insert_inner(value, old->left);
        return balance(old->color, left, old->value, old->right);
    }
    if (value > old->value) {
        Tree *right = insert_inner(value, old->right);
        return balance(old->color, old->left, old->value, right);
    }
    return tree(old->color, old->left, old->value, old->right);
}
static Tree *insert_tree(Int value, Tree *old) {
    Tree *result = insert_inner(value, old);
    return tree(BLACK, result->left, result->value, result->right);
}
static Tree *build_tree(Int n) {
    Tree *result = NULL;
    for (; n != 0; n = sub_int(n, 1)) result = insert_tree(n, result);
    return result;
}
static Int tree_depth(const Tree *node) {
    if (!node) return 0;
    Int left = tree_depth(node->left);
    Int right = tree_depth(node->right);
    return add_int(1, left > right ? left : right);
}
static Int run_tree(Int n) { return tree_depth(build_tree(n)); }

/* A generic typed dictionary pattern, instantiated for Int. */
#define DEFINE_MONOIDISH(Type, Prefix) \
    typedef struct { Type empty; Type (*append)(Type, Type); } Prefix##Monoidish; \
    static Type Prefix##_poly_loop(const Prefix##Monoidish *dict, Int n, Type acc) { \
        for (; n != 0; n = sub_int(n, 1)) acc = dict->append(acc, dict->empty); \
        return acc; \
    }
DEFINE_MONOIDISH(Int, Int)
static Int run_poly(Int n) {
    const IntMonoidish dict = {1, add_int};
    return Int_poly_loop(&dict, n, 0);
}

/* State Int a: Int and Unit values share a scalar slot (Unit maps to zero). */
typedef struct { Int value, state; } StateResult;
typedef struct { StateResult (*run)(const void *, Int); const void *env; } State;
typedef struct { State (*apply)(const void *, Int); const void *env; } StateNext;
typedef struct { State first; StateNext next; } BindEnv;
typedef struct { Int value; } StateValue;
typedef struct { Endo function; } ModifyEnv;
typedef struct { Int depth; } ChainEnv;

static StateResult bind_run(const void *context, Int initial) {
    const BindEnv *env = context;
    StateResult first = env->first.run(env->first.env, initial);
    State second = env->next.apply(env->next.env, first.value);
    return second.run(second.env, first.state);
}
static State bind_state(State first, StateNext next) {
    BindEnv *env = allocate(sizeof *env);
    *env = (BindEnv){first, next};
    return (State){bind_run, env};
}
static StateResult pure_run(const void *context, Int initial) {
    return (StateResult){((const StateValue *)context)->value, initial};
}
static State pure_state(Int value) {
    StateValue *env = allocate(sizeof *env);
    *env = (StateValue){value};
    return (State){pure_run, env};
}
static StateResult get_run(const void *env, Int initial) {
    (void)env;
    return (StateResult){initial, initial};
}
static StateResult put_run(const void *context, Int initial) {
    (void)initial;
    return (StateResult){0, ((const StateValue *)context)->value};
}
static State put_state(Int value) {
    StateValue *env = allocate(sizeof *env);
    *env = (StateValue){value};
    return (State){put_run, env};
}
static State modify_next(const void *context, Int initial) {
    const ModifyEnv *env = context;
    return put_state(env->function.apply(env->function.env, initial));
}
static State modify_state(Endo function) {
    ModifyEnv *env = allocate(sizeof *env);
    *env = (ModifyEnv){function};
    State get = {get_run, NULL};
    return bind_state(get, (StateNext){modify_next, env});
}
static State chain_modifications(Int depth);
static State chain_next(const void *context, Int unit) {
    (void)unit;
    const ChainEnv *env = context;
    return chain_modifications(sub_int(env->depth, 1));
}
static State chain_modifications(Int depth) {
    if (depth == 0) return pure_state(0);
    State modification = modify_state((Endo){increment_step, NULL});
    ChainEnv *env = allocate(sizeof *env);
    *env = (ChainEnv){depth};
    return bind_state(modification, (StateNext){chain_next, env});
}
static Int run_state(Int repetitions) {
    Int result = 0;
    for (; repetitions != 0; repetitions = sub_int(repetitions, 1)) {
        State action = chain_modifications(60);
        StateResult outcome = action.run(action.env, 0);
        result = add_int(result, outcome.state);
    }
    return result;
}

/* Explicit non-memoizing Unit -> Int closures, rebuilt each repetition. */
typedef struct { Int (*force)(const void *); const void *env; } Lazy;
typedef struct { Lazy previous; } LazyEnv;
static Lazy defer(Int (*function)(const void *), const void *env) {
    return (Lazy){function, env};
}
static Int force_lazy(Lazy value) { return value.force(value.env); }
static Int zero_thunk(const void *env) { (void)env; return 0; }
static Int increment_thunk(const void *context) {
    const LazyEnv *env = context;
    return add_int(force_lazy(env->previous), 1);
}
static Lazy build_thunks(Int depth, Lazy initial) {
    Lazy result = initial;
    for (; depth != 0; depth = sub_int(depth, 1)) {
        LazyEnv *env = allocate(sizeof *env);
        *env = (LazyEnv){result};
        result = defer(increment_thunk, env);
    }
    return result;
}
static Int run_lazy(Int repetitions) {
    Int result = 0;
    for (; repetitions != 0; repetitions = sub_int(repetitions, 1)) {
        Lazy chain = build_thunks(1000, defer(zero_thunk, NULL));
        result = add_int(result, force_lazy(chain));
    }
    return result;
}

/* Inclusive native arrays: range, filter into a new array, then generic fold. */
typedef struct { size_t length; Int *items; } IntArray;
static IntArray array_range(Int start, Int end) {
    int64_t difference = (int64_t)end - start;
    uint64_t distance = difference < 0 ? (uint64_t)(-difference) : (uint64_t)difference;
    if (distance + 1 > SIZE_MAX / sizeof(Int)) abort();
    size_t length = (size_t)distance + 1;
    Int *items = allocate(length * sizeof *items);
    Int step = start <= end ? 1 : -1;
    Int value = start;
    for (size_t i = 0; i < length; ++i, value = add_int(value, step)) items[i] = value;
    return (IntArray){length, items};
}
static bool even_predicate(const void *env, Int value) {
    (void)env;
    return mod_int(value, 2) == 0;
}
static IntArray array_filter(Predicate predicate, IntArray input) {
    /* Private output buffer; the logical immutable result is its used prefix. */
    Int *items = allocate(input.length * sizeof *items);
    size_t length = 0;
    for (size_t i = 0; i < input.length; ++i)
        if (predicate.apply(predicate.env, input.items[i])) items[length++] = input.items[i];
    return (IntArray){length, items};
}
static Int array_fold(Fold function, Int initial, IntArray input) {
    Int result = initial;
    for (size_t i = 0; i < input.length; ++i)
        result = function.apply(function.env, result, input.items[i]);
    return result;
}
static Int run_array(Int n) {
    IntArray range = array_range(1, n);
    IntArray selected = array_filter((Predicate){even_predicate, NULL}, range);
    return array_fold((Fold){add_fold, NULL}, 0, selected);
}

/* Typed heterogeneous row with a recursive keysNil/keysCons dictionary. */
typedef struct { Int a; const char *b; bool c; double d; const char *e; } Row;
typedef struct { unsigned char unit; } Proxy;
typedef struct Keys Keys;
struct Keys { Int (*keys_impl)(const Keys *, Proxy); const Keys *tail; };
static Int keys_nil(const Keys *dict, Proxy proxy) {
    (void)dict; (void)proxy;
    return 0;
}
static Int keys_cons(const Keys *dict, Proxy proxy) {
    return add_int(1, dict->tail->keys_impl(dict->tail, proxy));
}
static Int record_keys(const Keys *dict, const Row *row) {
    (void)row;
    return dict->keys_impl(dict, (Proxy){0});
}
static Int run_row(Int ignored) {
    (void)ignored;
    const Keys nil = {keys_nil, NULL};
    const Keys e = {keys_cons, &nil};
    const Keys d = {keys_cons, &e};
    const Keys c = {keys_cons, &d};
    const Keys b = {keys_cons, &c};
    const Keys a = {keys_cons, &b};
    const Row row = {1, "two", true, 4.0, "five"};
    return record_keys(&a, &row);
}

typedef Int (*Action)(Int);
typedef struct { const char *label; Action action; Int input, expected; } Case;
static const Case cases[] = {
    {"AST Evaluation:", run_ast, 3, 7},
    {"Fibonacci:", fibonacci, 10, 55},
    {"List Processing (900 elements):", run_list, 900, 202950},
    {"Tail Call Optimization (100k calls):", run_tco, 100000, 100000},
    {"Deep Record Updates (10k iterations):", run_records, 10000, 20000},
    {"Ackermann (3, 4):", run_ackermann, 3, 125},
    {"Church Numerals (100k Closure Applications):", run_church, 10, 100000},
    {"Prime Sieve (sum primes up to 500):", run_primes, 500, 21536},
    {"Red-Black Tree (100k Worst-Case Insertions):", run_tree, 100000, 22},
    {"Polymorphism (10M Type Class Dict Lookups):", run_poly, 10000000, 10000000},
    {"State Monad (1.2k Binds, 60 Stack Depth):", run_state, 20, 1200},
    {"Lazy Evaluation (1M Thunks Forced, 1k Depth):", run_lazy, 1000, 1000000},
    {"Array Processing (900 elements):", run_array, 900, 202950},
    {"RowToList (Keys Count):", run_row, 10000, 5},
};
static const size_t case_count = sizeof cases / sizeof cases[0];

static Int perform(Action action, Int input) {
    if (allocations) abort();
#ifdef VERIFY_ALLOCATIONS
    allocation_count = allocation_bytes = 0;
#endif
    Int result = action(input);
    release_allocations();
    return result;
}

/* Oracles are verification-only and never provide benchmark results. */
static void expect(const char *name, Int actual, Int expected) {
    if (actual != expected) {
        fprintf(stderr, "%s: got %" PRId32 ", expected %" PRId32 "\n", name, actual, expected);
        exit(1);
    }
}
static int check_tree(const Tree *node, int64_t lower, int64_t upper) {
    if (!node) return 1;
    if (node->value <= lower || node->value >= upper ||
        (red(node) && (red(node->left) || red(node->right)))) abort();
    int left = check_tree(node->left, lower, node->value);
    int right = check_tree(node->right, node->value, upper);
    if (left != right) abort();
    return left + (node->color == BLACK);
}
static void verify(void) {
    for (size_t i = 0; i < case_count; ++i) {
        expect(cases[i].label, perform(cases[i].action, cases[i].input), cases[i].expected);
        printf("Verified %s", cases[i].label);
#ifdef VERIFY_ALLOCATIONS
        printf(" allocations=%zu payload_bytes=%zu", allocation_count, allocation_bytes);
#endif
        putchar('\n');
    }
    const struct { Action action; Int input, output; } variants[] = {
        {run_ast,0,1},{run_ast,1,1},{run_ast,2,2},{run_ast,4,34},{run_ast,5,203},
        {fibonacci,0,0},{fibonacci,1,1},{fibonacci,2,1},{fibonacci,20,6765},
        {run_list,0,0},{run_list,1,0},{run_list,2,2},{run_list,5,6},{run_list,6,12},
        {run_tco,0,0},{run_tco,1,1},{run_tco,2,3},{run_tco,3,3},{run_tco,5,6},
        {run_records,0,0},{run_records,1,1},{run_records,4,10},{run_records,5,10},
        {run_records,6,11},{run_records,17,33},
        {run_ackermann,0,5},{run_ackermann,1,6},{run_ackermann,2,11},
        {run_church,0,0},{run_church,1,1},{run_church,2,32},{run_church,3,243},
        {run_primes,0,0},{run_primes,1,0},{run_primes,2,2},{run_primes,3,5},
        {run_primes,10,17},{run_primes,30,129},
        {run_tree,0,0},{run_tree,1,1},{run_tree,2,2},{run_tree,3,2},
        {run_poly,0,0},{run_poly,1,1},{run_poly,17,17},
        {run_state,0,0},{run_state,1,60},{run_state,2,120},
        {run_lazy,0,0},{run_lazy,1,1000},{run_lazy,2,2000},
        {run_array,-2,-2},{run_array,0,0},{run_array,1,0},{run_array,2,2},
        {run_array,5,6},{run_array,6,12},{run_row,0,5},
    };
    for (size_t i = 0; i < sizeof variants / sizeof variants[0]; ++i)
        expect("input variant", perform(variants[i].action, variants[i].input), variants[i].output);
    expect("Int32 wrap add", add_int(INT32_MAX, 1), INT32_MIN);
    expect("Int32 wrap subtract", sub_int(INT32_MIN, 1), INT32_MAX);
    expect("Int32 wrap multiply", mul_int(INT32_MAX, 2), -2);
    const RecordA old = {7,{8,{9,10}}};
    RecordA next = update_record(2, old);
    expect("record old snapshot", old.b.d.f, 10);
    expect("record next a", next.a, 9);
    expect("record next c", next.b.c, 12);
    expect("record next e", next.b.d.e, 15);
    expect("record next f", next.b.d.f, 13);
    Tree *root = NULL;
    const Int order[] = {8,3,10,1,6,14,4,7,13,2,9,5,12,11};
    for (size_t i = 0; i < sizeof order / sizeof order[0]; ++i) {
        Tree *previous = root;
        Int previous_depth = tree_depth(previous);
        root = insert_tree(order[i], root);
        if (root->color != BLACK) abort();
        check_tree(root, INT64_MIN, INT64_MAX);
        check_tree(previous, INT64_MIN, INT64_MAX);
        expect("persistent old tree", tree_depth(previous), previous_depth);
    }
    Tree *duplicate = insert_tree(8, root);
    check_tree(duplicate, INT64_MIN, INT64_MAX);
    expect("duplicate insertion depth", tree_depth(duplicate), tree_depth(root));
    release_allocations();
    const Keys nil = {keys_nil,NULL};
    const Keys one = {keys_cons,&nil};
    const Keys two = {keys_cons,&one};
    const Keys three = {keys_cons,&two};
    expect("keys Nil", nil.keys_impl(&nil,(Proxy){0}),0);
    expect("keys Cons", one.keys_impl(&one,(Proxy){0}),1);
    expect("keys depth3", three.keys_impl(&three,(Proxy){0}),3);
    printf("Verification passed: %zu core oracles, %zu input variants, Int32, immutable records, persistent trees and key dictionaries.\n",
           case_count, sizeof variants / sizeof variants[0]);
}

/* Harness-only opacity: load the input each iteration, consume every result.
 * seq_cst matches the atomic-store intent of the JS/Go harnesses; the native
 * atomic implementation and call-adapter costs still differ across runtimes.
 */
static _Atomic Int benchmark_sink;
static uint64_t now_ns(void) {
    struct timespec time;
    if (clock_gettime(CLOCK_MONOTONIC, &time) != 0) {
        perror("clock_gettime");
        exit(1);
    }
    return (uint64_t)time.tv_sec * UINT64_C(1000000000) + (uint64_t)time.tv_nsec;
}
static uint64_t batch(Action action, Int input, int iterations) {
    volatile Int opaque_input = input;
    uint64_t start = now_ns();
    for (int i = 0; i < iterations; ++i)
        atomic_store_explicit(&benchmark_sink, perform(action, opaque_input), memory_order_seq_cst);
    return now_ns() - start;
}
static double bench(const Case *test) {
    printf("--------------------------------------------------\n\n(Test)\n%s\n\n(Output & Warm-up)\n", test->label);
    volatile Int opaque_input = test->input;
    Int result = perform(test->action, opaque_input);
    expect(test->label, result, test->expected);
    atomic_store_explicit(&benchmark_sink, result, memory_order_seq_cst);
    printf("%" PRId32 "\n", result);
    for (int warmup = 0; warmup < 2; ++warmup)
        atomic_store_explicit(&benchmark_sink, perform(test->action, opaque_input), memory_order_seq_cst);
    int iterations = 1;
    while (batch(test->action, test->input, iterations) < UINT64_C(10000000) && iterations < 16777216)
        iterations *= 2;
    uint64_t best = UINT64_MAX;
    for (int sample = 0; sample < 10; ++sample) {
        uint64_t elapsed = batch(test->action, test->input, iterations);
        expect(test->label, atomic_load_explicit(&benchmark_sink, memory_order_seq_cst), result);
        if (elapsed < best) best = elapsed;
    }
    double us = (double)best / 1000.0 / iterations;
    printf("\n(Execution time - best of 10)\n\n%.9f μs\n\nBatch iterations: %d\n\n", us, iterations);
    return us;
}
int main(int argc, char **argv) {
    if (argc == 2 && strcmp(argv[1], "--verify") == 0) { verify(); return 0; }
    if (argc != 1) { fprintf(stderr, "Usage: %s [--verify]\n", argv[0]); return 2; }
    puts("Global warm-up in progress...");
    for (int warmup = 0; warmup < 3; ++warmup)
        for (size_t i = 0; i < case_count; ++i) {
            volatile Int input = cases[i].input;
            Int result = perform(cases[i].action, input);
            expect(cases[i].label, result, cases[i].expected);
            atomic_store_explicit(&benchmark_sink, result, memory_order_seq_cst);
        }
    double total = 0;
    for (size_t i = 0; i < case_count; ++i) total += bench(&cases[i]);
    printf("\n==================================================\n\nTotal exec time: %.9f ms\n", total / 1000.0);
    return 0;
}
