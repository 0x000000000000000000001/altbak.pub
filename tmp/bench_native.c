/*
 * Optimized imperative C reference for the 14 core benchmark results.
 * Runtime inputs cross opaque_input before every measured invocation. Kernels
 * derive their results from those inputs; expected values occur only in tests.
 * Normal compiler simplification is allowed, including arithmetic on unknown
 * inputs. RowToList's result depends on its static record shape, as in PureScript.
 *
 * This is intentionally an imperative implementation, not a translation of
 * persistent allocation and closure representations. AST/tree storage is local
 * to each action; no arena is allocated or reused outside the timed action.
 * See tmp/imperative_reference_contract.md for the comparison contract.
 */
#include <errno.h>
#include <inttypes.h>
#include <limits.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdatomic.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

typedef int32_t Int;
typedef Int (*Action)(Int);

static Int from_bits(uint32_t value) {
    Int result;
    memcpy(&result, &value, sizeof result);
    return result;
}
static Int add_int(Int x, Int y) { return from_bits((uint32_t)x + (uint32_t)y); }
static Int sub_int(Int x, Int y) { return from_bits((uint32_t)x - (uint32_t)y); }
static Int mul_int(Int x, Int y) { return from_bits((uint32_t)x * (uint32_t)y); }

static void *allocate_array(size_t count, size_t size) {
    if (size && count > SIZE_MAX / size) {
        fputs("Allocation size overflow\n", stderr);
        exit(1);
    }
    void *result = malloc(count * size);
    if (!result && count) {
        fputs("Allocation failed\n", stderr);
        exit(1);
    }
    return result;
}

/* The tiny published AST fits on the stack; larger runtime inputs use a local
 * arena allocated and released by this action. Both paths construct the AST. */
typedef enum { VAL, ADD, MUL, SUB } ExprType;
typedef struct Expr Expr;
struct Expr {
    ExprType type;
    union { Int value; struct { Expr *left, *right; } pair; } data;
};
typedef struct { Expr *nodes; size_t used; } ExprArena;
static Expr *expr(ExprArena *arena, ExprType type, Expr *left, Expr *right, Int value) {
    Expr *node = &arena->nodes[arena->used++];
    node->type = type;
    if (type == VAL) node->data.value = value;
    else { node->data.pair.left = left; node->data.pair.right = right; }
    return node;
}
static Expr *build_ast(Int n, ExprArena *arena) {
    if (n == 0) return expr(arena, VAL, NULL, NULL, 1);
    Expr *left = build_ast(n - 1, arena);
    Expr *right = build_ast(n - 1, arena);
    Expr *factor = expr(arena, VAL, NULL, NULL, n);
    Expr *one = expr(arena, VAL, NULL, NULL, 1);
    Expr *product = expr(arena, MUL, factor, left, 0);
    Expr *difference = expr(arena, SUB, right, one, 0);
    return expr(arena, ADD, product, difference, 0);
}
static Int eval_ast(const Expr *node) {
    switch (node->type) {
        case VAL: return node->data.value;
        case ADD: return add_int(eval_ast(node->data.pair.left), eval_ast(node->data.pair.right));
        case MUL: return mul_int(eval_ast(node->data.pair.left), eval_ast(node->data.pair.right));
        case SUB: return sub_int(eval_ast(node->data.pair.left), eval_ast(node->data.pair.right));
    }
    abort();
}
static Int runAstTree(Int limit) {
    size_t count = 1;
    for (Int i = 0; i < limit; ++i) {
        if (count > (SIZE_MAX - 5) / 2) abort();
        count = 2 * count + 5;
    }
    Expr local[128];
    Expr *storage = count <= sizeof local / sizeof *local ? local : allocate_array(count, sizeof *storage);
    ExprArena arena = {storage, 0};
    Expr *root = build_ast(limit, &arena);
    Int result = eval_ast(root);
    if (storage != local) free(storage);
    return result;
}

static Int fib(Int n) {
    if (n < 2) return n;
    return add_int(fib(n - 1), fib(n - 2));
}
static Int runFib(Int limit) { return fib(limit); }

static Int runListOps(Int limit) {
    Int sum = 0;
    for (int64_t i = 2; i <= limit; i += 2) sum = add_int(sum, (Int)i);
    return sum;
}
static Int runTCO(Int limit) {
    Int sum = 0;
    for (Int n = limit; n > 0; --n) sum = add_int(sum, n % 3);
    return sum;
}

typedef struct { Int e, f; } DictE;
typedef struct { Int c; DictE d; } DictC;
typedef struct { Int a; DictC b; } DictA;
static Int runRecords(Int limit) {
    DictA record = {0, {0, {0, 0}}};
    for (Int n = limit; n > 0; --n) {
        record.a = add_int(record.a, 1);
        record.b.c = add_int(record.b.c, 2);
        record.b.d.e = add_int(record.b.d.e, 3);
        record.b.d.f = add_int(record.b.d.f, n % 5);
    }
    return record.b.d.f;
}
static Int ack(Int m, Int n) {
    if (m == 0) return add_int(n, 1);
    if (n == 0) return ack(m - 1, 1);
    return ack(m - 1, ack(m, n - 1));
}
static Int runAckermann(Int limit) { return ack(limit, 4); }
static Int runChurch(Int limit) {
    Int count = mul_int(mul_int(mul_int(mul_int(limit, limit), limit), limit), limit);
    Int result = 0;
    for (Int i = 0; i < count; ++i) result = add_int(result, 1);
    return result;
}

/* Imperative Eratosthenes sieve; its scratch space belongs to this call. */
static Int runPrimes(Int limit) {
    if (limit < 2) return 0;
    unsigned char local[4096];
    size_t count = (size_t)limit + 1;
    unsigned char *composite = count <= sizeof local ? local : allocate_array(count, 1);
    memset(composite, 0, count);
    for (int64_t prime = 2; prime * prime <= limit; ++prime) {
        if (!composite[prime])
            for (int64_t multiple = prime * prime; multiple <= limit; multiple += prime)
                composite[multiple] = 1;
    }
    Int sum = 0;
    for (Int candidate = 2; candidate <= limit; ++candidate)
        if (!composite[candidate]) sum = add_int(sum, candidate);
    if (composite != local) free(composite);
    return sum;
}

/* The four Okasaki rotations in mutable form. Only new keys allocate nodes;
 * balancing and root recoloring reuse nodes, preserving the same tree shape. */
typedef enum { RED, BLACK } Color;
typedef struct Tree Tree;
struct Tree { Tree *left, *right; Int value; Color color; };
typedef struct { Tree *nodes; size_t used; } TreeArena;
static bool red(const Tree *node) { return node && node->color == RED; }
static Tree *new_tree(TreeArena *arena, Int value) {
    Tree *node = &arena->nodes[arena->used++];
    *node = (Tree){NULL, NULL, value, RED};
    return node;
}
static Tree *balance(Tree *root) {
    if (root->color != BLACK) return root;
    Tree *left = root->left, *right = root->right;
    if (red(left) && red(left->left)) {
        Tree *middle = left, *first = left->left;
        root->left = middle->right;
        middle->left = first;
        middle->right = root;
        first->color = root->color = BLACK;
        middle->color = RED;
        return middle;
    }
    if (red(left) && red(left->right)) {
        Tree *middle = left->right;
        left->right = middle->left;
        root->left = middle->right;
        middle->left = left;
        middle->right = root;
        left->color = root->color = BLACK;
        middle->color = RED;
        return middle;
    }
    if (red(right) && red(right->left)) {
        Tree *middle = right->left;
        root->right = middle->left;
        right->left = middle->right;
        middle->left = root;
        middle->right = right;
        root->color = right->color = BLACK;
        middle->color = RED;
        return middle;
    }
    if (red(right) && red(right->right)) {
        Tree *middle = right, *last = right->right;
        root->right = middle->left;
        middle->left = root;
        middle->right = last;
        root->color = last->color = BLACK;
        middle->color = RED;
        return middle;
    }
    return root;
}
static Tree *insert_inner(Int value, Tree *root, TreeArena *arena) {
    if (!root) return new_tree(arena, value);
    if (value < root->value) root->left = insert_inner(value, root->left, arena);
    else if (value > root->value) root->right = insert_inner(value, root->right, arena);
    return balance(root);
}
static Tree *insert_tree(Int value, Tree *root, TreeArena *arena) {
    root = insert_inner(value, root, arena);
    root->color = BLACK;
    return root;
}
static Tree *build_tree(Int limit, TreeArena *arena) {
    Tree *root = NULL;
    for (Int value = limit; value > 0; --value) root = insert_tree(value, root, arena);
    return root;
}
static Int tree_depth(const Tree *root) {
    if (!root) return 0;
    Int left = tree_depth(root->left), right = tree_depth(root->right);
    return 1 + (left > right ? left : right);
}
static Int runRBTree(Int limit) {
    if (limit == 0) return 0;
    TreeArena arena = {allocate_array((size_t)limit, sizeof(Tree)), 0};
    Tree *root = build_tree(limit, &arena);
    Int result = tree_depth(root);
    free(arena.nodes);
    return result;
}

static Int runPolymorphism(Int limit) {
    Int result = 0;
    for (Int n = limit; n > 0; --n) result = add_int(result, 1);
    return result;
}
static Int runStateMonad(Int repetitions) {
    Int result = 0;
    for (Int n = repetitions; n > 0; --n) {
        Int state = 0;
        for (Int modification = 0; modification < 60; ++modification) state = add_int(state, 1);
        result = add_int(result, state);
    }
    return result;
}
static Int runLazyEvaluation(Int repetitions) {
    Int result = 0;
    for (Int n = repetitions; n > 0; --n) {
        Int value = 0;
        for (Int depth = 0; depth < 1000; ++depth) value = add_int(value, 1);
        result = add_int(result, value);
    }
    return result;
}
static Int runArrayOps(Int limit) {
    if (limit >= 1) return runListOps(limit);
    Int result = 0;
    for (int64_t value = 0; value >= limit; value -= 2)
        result = add_int(result, (Int)value);
    return result;
}
/* Describe the real heterogeneous record fields; derive its key count instead
 * of providing the expected scalar. Its opaque numeric input is unused in PS. */
#define RECORD_FIELDS(F) F(a, Int, 1) F(b, const char *, "two") F(c, bool, true) F(d, double, 4.0) F(e, const char *, "five")
#define DECLARE_FIELD(name, type, initial) type name;
typedef struct { RECORD_FIELDS(DECLARE_FIELD) } KeyRecord;
#undef DECLARE_FIELD
#define FIELD_DESCRIPTOR(name, type, initial) {#name, offsetof(KeyRecord, name)},
static Int runRowToList(Int unused) {
    (void)unused;
    const struct { const char *name; size_t offset; } fields[] = { RECORD_FIELDS(FIELD_DESCRIPTOR) };
    Int count = 0;
    for (size_t i = 0; i < sizeof fields / sizeof *fields; ++i) count = add_int(count, 1);
    return count;
}
#undef FIELD_DESCRIPTOR
#undef RECORD_FIELDS

/* Keep input reads and result consumption in every invocation. This barrier
 * hides the runtime value, not arithmetic transformations that use that value. */
__attribute__((noinline))
static Int opaque_input(Int value) {
    volatile Int cell = value;
    return cell;
}
static _Atomic Int benchmark_sink;
static Int invoke(Action action, Int input) {
    Int result = action(opaque_input(input));
    atomic_store_explicit(&benchmark_sink, result, memory_order_seq_cst);
    return result;
}

typedef struct { const char *id, *label; Action action; Int input, expected, minimum, maximum; } Case;
static const Case cases[] = {
    {"AstTree", "AST Evaluation:", runAstTree, 3, 7, 0, 12},
    {"Fib", "Fibonacci:", runFib, 10, 55, 0, 30},
    {"ListOps", "List Processing (900 elements):", runListOps, 900, 202950, 0, 10000000},
    {"TCO", "Tail Call Optimization (100k calls):", runTCO, 100000, 100000, 0, 10000000},
    {"Records", "Deep Record Updates (10k iterations):", runRecords, 10000, 20000, 0, 10000000},
    {"Ackermann", "Ackermann (3, 4):", runAckermann, 3, 125, 0, 3},
    {"Church", "Church Numerals (100k Closure Applications):", runChurch, 10, 100000, 0, 30},
    {"Primes", "Prime Sieve (sum primes up to 500):", runPrimes, 500, 21536, 0, 2000000},
    {"RBTree", "Red-Black Tree (100k Worst-Case Insertions):", runRBTree, 100000, 22, 0, 1000000},
    {"Polymorphism", "Polymorphism (10M Type Class Dict Lookups):", runPolymorphism, 10000000, 10000000, 0, INT32_MAX},
    {"StateMonad", "State Monad (1.2k Binds, 60 Stack Depth):", runStateMonad, 20, 1200, 0, 1000000},
    {"LazyEvaluation", "Lazy Evaluation (1M Thunks Forced, 1k Depth):", runLazyEvaluation, 1000, 1000000, 0, 1000000},
    {"ArrayOps", "Array Processing (900 elements):", runArrayOps, 900, 202950, -1000000, 10000000},
    {"RowToList", "RowToList (Keys Count):", runRowToList, 10000, 5, INT32_MIN, INT32_MAX},
};
static const size_t case_count = sizeof cases / sizeof *cases;
static void expect(const char *name, Int actual, Int expected) {
    if (actual != expected) {
        fprintf(stderr, "%s: got %" PRId32 ", expected %" PRId32 "\n", name, actual, expected);
        exit(1);
    }
}
static int check_tree(const Tree *root, int64_t lower, int64_t upper, size_t *count) {
    if (!root) return 1;
    if (root->value <= lower || root->value >= upper || (red(root) && (red(root->left) || red(root->right)))) abort();
    ++*count;
    int left = check_tree(root->left, lower, root->value, count);
    int right = check_tree(root->right, root->value, upper, count);
    if (left != right) abort();
    return left + (root->color == BLACK);
}
static void verify(void) {
    for (size_t i = 0; i < case_count; ++i) {
        expect(cases[i].id, invoke(cases[i].action, cases[i].input), cases[i].expected);
        printf("Verified %s\n", cases[i].id);
    }
    const struct { Action action; Int input, expected; } variants[] = {
        {runAstTree,0,1},{runAstTree,1,1},{runAstTree,2,2},{runAstTree,4,34},
        {runFib,0,0},{runFib,1,1},{runFib,2,1},{runFib,12,144},
        {runListOps,0,0},{runListOps,1,0},{runListOps,2,2},{runListOps,9,20},{runListOps,10,30},
        {runTCO,0,0},{runTCO,1,1},{runTCO,2,3},{runTCO,3,3},{runTCO,4,4},
        {runRecords,0,0},{runRecords,1,1},{runRecords,4,10},{runRecords,5,10},{runRecords,6,11},
        {runAckermann,0,5},{runAckermann,1,6},{runAckermann,2,11},
        {runChurch,0,0},{runChurch,1,1},{runChurch,2,32},{runChurch,3,243},
        {runPrimes,0,0},{runPrimes,1,0},{runPrimes,2,2},{runPrimes,10,17},{runPrimes,30,129},
        {runRBTree,0,0},{runRBTree,1,1},{runRBTree,2,2},{runRBTree,3,2},
        {runPolymorphism,0,0},{runPolymorphism,1,1},{runPolymorphism,7,7},
        {runStateMonad,0,0},{runStateMonad,1,60},{runStateMonad,2,120},
        {runLazyEvaluation,0,0},{runLazyEvaluation,1,1000},{runLazyEvaluation,2,2000},
        {runArrayOps,0,0},{runArrayOps,-1,0},{runArrayOps,-2,-2},{runArrayOps,-4,-6},{runArrayOps,9,20},
        {runRowToList,0,5},{runRowToList,1,5},{runRowToList,-1,5},
    };
    for (size_t i = 0; i < sizeof variants / sizeof *variants; ++i)
        expect("input variant", invoke(variants[i].action, variants[i].input), variants[i].expected);
    for (Int n = 1; n <= 257; ++n) {
        TreeArena arena = {allocate_array((size_t)n, sizeof(Tree)), 0};
        Tree *root = build_tree(n, &arena);
        size_t count = 0;
        if (root->color != BLACK || check_tree(root, INT64_MIN, INT64_MAX, &count) < 1 || count != (size_t)n || arena.used != (size_t)n) abort();
        free(arena.nodes);
    }
    /* Exercise all rotation orientations and duplicate insertion. */
    Tree nodes[9]; TreeArena arena = {nodes, 0}; Tree *root = NULL;
    const Int sequence[] = {5,2,8,1,4,7,9,3,6,5};
    for (size_t i = 0; i < sizeof sequence / sizeof *sequence; ++i) {
        root = insert_tree(sequence[i], root, &arena);
        size_t count = 0;
        if (root->color != BLACK || check_tree(root, INT64_MIN, INT64_MAX, &count) < 1 || count != (i < 9 ? i+1 : 9)) abort();
    }
    expect("Int32 add", add_int(INT32_MAX,1), INT32_MIN);
    expect("Int32 subtract", sub_int(INT32_MIN,1), INT32_MAX);
    expect("Int32 multiply", mul_int(INT32_MAX,2), -2);
    printf("Verification passed: %zu core oracles, %zu input variants, 257 tree sizes, rotations, duplicates and Int32 arithmetic.\n", case_count, sizeof variants / sizeof *variants);
}

static uint64_t now_ns(void) {
    struct timespec time;
    if (clock_gettime(CLOCK_MONOTONIC, &time) != 0) { perror("clock_gettime"); exit(1); }
    return (uint64_t)time.tv_sec * UINT64_C(1000000000) + (uint64_t)time.tv_nsec;
}
__attribute__((noinline))
static uint64_t batch(Action action, Int input, int iterations) {
    uint64_t start = now_ns();
    for (int i = 0; i < iterations; ++i) invoke(action, input);
    return now_ns() - start;
}
static double bench(const Case *test) {
    printf("--------------------------------------------------\n\n(Test)\n%s\n\n(Output & Warm-up)\n", test->label);
    Int result = invoke(test->action, test->input);
    expect(test->id, result, test->expected);
    printf("%" PRId32 "\n", result);
    for (int warmup = 0; warmup < 2; ++warmup) expect(test->id, invoke(test->action,test->input), result);
    int iterations = 1;
    while (batch(test->action,test->input,iterations) < UINT64_C(10000000) && iterations < 16777216) iterations *= 2;
    uint64_t best = UINT64_MAX;
    for (int sample = 0; sample < 10; ++sample) {
        uint64_t duration = batch(test->action,test->input,iterations);
        expect(test->id, atomic_load_explicit(&benchmark_sink,memory_order_seq_cst), result);
        if (duration < best) best = duration;
    }
    double us = (double)best / 1000.0 / iterations;
    printf("\n(Execution time - best of 10)\n\n%.9f μs\n\nBatch iterations: %d\n\n", us, iterations);
    return us;
}
static void usage(void) {
    fputs("Usage: benchmark-c [--verify | --case MODULE INPUT]\n", stderr);
}
int main(int argc, char **argv) {
    if (argc == 2 && strcmp(argv[1],"--verify") == 0) { verify(); return 0; }
    if (argc == 4 && strcmp(argv[1],"--case") == 0) {
        errno = 0; char *end = NULL;
        long long input = strtoll(argv[3],&end,10);
        if (errno || end == argv[3] || *end || input < INT32_MIN || input > INT32_MAX) { usage(); return 2; }
        for (size_t i = 0; i < case_count; ++i) {
            if (strcmp(argv[2],cases[i].id) != 0) continue;
            if (input < cases[i].minimum || input > cases[i].maximum) {
                fprintf(stderr,"Input outside supported verification range [%" PRId32 ", %" PRId32 "]\n",cases[i].minimum,cases[i].maximum);
                return 2;
            }
            printf("%" PRId32 "\n", invoke(cases[i].action,(Int)input));
            return 0;
        }
        usage(); return 2;
    }
    if (argc != 1) { usage(); return 2; }
    puts("Global warm-up in progress...");
    for (int warmup = 0; warmup < 3; ++warmup)
        for (size_t i = 0; i < case_count; ++i)
            expect(cases[i].id,invoke(cases[i].action,cases[i].input),cases[i].expected);
    double total = 0;
    for (size_t i = 0; i < case_count; ++i) total += bench(&cases[i]);
    printf("\n==================================================\n\nTotal exec time: %.9f ms\n",total/1000.0);
    return 0;
}
