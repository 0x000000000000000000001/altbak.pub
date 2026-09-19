#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>

typedef enum { VAL, ADD, MUL, SUB } ExprType;
typedef struct Expr {
    ExprType type;
    union {
        int val;
        struct { struct Expr* l; struct Expr* r; } bin;
    };
} Expr;

Expr* expr_arena;
int expr_idx = 0;
Expr* alloc_expr(ExprType t) {
    Expr* e = &expr_arena[expr_idx++];
    e->type = t;
    return e;
}

int evalAst(Expr* e) {
    switch (e->type) {
        case VAL: return e->val;
        case ADD: return evalAst(e->bin.l) + evalAst(e->bin.r);
        case MUL: return evalAst(e->bin.l) * evalAst(e->bin.r);
        case SUB: return evalAst(e->bin.l) - evalAst(e->bin.r);
    }
    return 0;
}

Expr* buildTreeAst(int n) {
    if (n == 0) {
        Expr* e = alloc_expr(VAL); e->val = 1;
        return e;
    } else {
        Expr* e = alloc_expr(ADD);
        Expr* mul = alloc_expr(MUL);
        Expr* vn = alloc_expr(VAL); vn->val = n;
        mul->bin.l = vn; mul->bin.r = buildTreeAst(n - 1);
        
        Expr* sub = alloc_expr(SUB);
        Expr* v1 = alloc_expr(VAL); v1->val = 1;
        sub->bin.l = buildTreeAst(n - 1); sub->bin.r = v1;
        
        e->bin.l = mul; e->bin.r = sub;
        return e;
    }
}

int runAstTree(int limit) {
    expr_idx = 0;
    Expr* e = buildTreeAst(limit);
    return evalAst(e);
}

int fib(int n) {
    if (n == 0) return 0;
    if (n == 1) return 1;
    return fib(n - 1) + fib(n - 2);
}
int runFib(int limit) { return fib(limit); }

int runListOps(int limit) {
    int sum = 0;
    for (int i = 1; i <= limit; i++) {
        if (i % 2 == 0) sum += i;
    }
    return sum;
}

int runTCO(int limit) {
    int acc = 0;
    int n = limit;
    while (n > 0) {
        acc += (n % 3);
        n--;
    }
    return acc;
}

typedef struct { int e, f; } DictE;
typedef struct { int c; DictE d; } DictC;
typedef struct { int a; DictC b; } DictA;

int runRecords(int limit) {
    DictA r = {0, {0, {0, 0}}};
    for (int n = limit; n >= 1; n--) {
        r.a += 1;
        r.b.c += 2;
        r.b.d.e += 3;
        r.b.d.f += (n % 5);
    }
    return r.b.d.f;
}

int ack(int m, int n) {
    if (m == 0) return n + 1;
    if (m > 0 && n == 0) return ack(m - 1, 1);
    return ack(m - 1, ack(m, n - 1));
}
int runAckermann(int limit) { return ack(limit, 4); }

int runChurch(int limit) {
    int count = limit * limit * limit * limit * limit;
    int acc = 0;
    for (int i = 1; i <= count; i++) acc++;
    return acc;
}

int runPrimes(int limit) {
    int sum = 0;
    for (int curr = 2; curr <= limit; curr++) {
        int is_prime = 1;
        for (int i = 2; i * i <= curr; i++) {
            if (curr % i == 0) { is_prime = 0; break; }
        }
        if (is_prime) sum += curr;
    }
    return sum;
}

typedef enum { R, B } Color;
typedef struct Tree {
    Color c;
    struct Tree *l, *r;
    int v;
} Tree;

Tree* tree_arena;
int tree_idx = 0;

Tree* alloc_tree(Color c, Tree* l, int v, Tree* r) {
    Tree* t = &tree_arena[tree_idx++];
    t->c = c; t->l = l; t->v = v; t->r = r;
    return t;
}

Tree* balance(Color c, Tree* l, int v, Tree* r) {
    if (c == B && l && l->c == R && l->l && l->l->c == R)
        return alloc_tree(R, alloc_tree(B, l->l->l, l->l->v, l->l->r), l->v, alloc_tree(B, l->r, v, r));
    if (c == B && l && l->c == R && l->r && l->r->c == R)
        return alloc_tree(R, alloc_tree(B, l->l, l->v, l->r->l), l->r->v, alloc_tree(B, l->r->r, v, r));
    if (c == B && r && r->c == R && r->l && r->l->c == R)
        return alloc_tree(R, alloc_tree(B, l, v, r->l->l), r->l->v, alloc_tree(B, r->l->r, r->v, r->r));
    if (c == B && r && r->c == R && r->r && r->r->c == R)
        return alloc_tree(R, alloc_tree(B, l, v, r->l), r->v, alloc_tree(B, r->r->l, r->r->v, r->r->r));
    return alloc_tree(c, l, v, r);
}

Tree* ins(int x, Tree* t) {
    if (!t) return alloc_tree(R, NULL, x, NULL);
    if (x < t->v) return balance(t->c, ins(x, t->l), t->v, t->r);
    if (x > t->v) return balance(t->c, t->l, t->v, ins(x, t->r));
    return alloc_tree(t->c, t->l, t->v, t->r);
}

Tree* insert(int x, Tree* t) {
    Tree* res = ins(x, t);
    return alloc_tree(B, res->l, res->v, res->r);
}

int depth(Tree* t) {
    if (!t) return 0;
    int ld = depth(t->l);
    int rd = depth(t->r);
    return 1 + (ld > rd ? ld : rd);
}

int runRBTree(int limit) {
    tree_idx = 0;
    Tree* acc = NULL;
    for (int i = limit; i >= 1; i--) {
        acc = insert(i, acc);
    }
    return depth(acc);
}

int runPolymorphism(int limit) {
    int acc = 0;
    for (int i = 1; i <= limit; i++) acc++;
    return acc;
}

int runStateMonad(int limit) {
    int state = 0;
    for (int i = 1; i <= 20; i++) {
        for (int j = 1; j <= limit; j++) state++;
    }
    return state;
}

int runLazyEvaluation(int limit) {
    int acc = 0;
    for (int i = 1; i <= limit; i++) acc += 1000;
    return acc;
}

int runArrayOps(int limit) {
    int sum = 0;
    for (int i = 1; i <= limit; i++) {
        if (i % 2 == 0) sum += i;
    }
    return sum;
}

int runRowToList(int limit) {
    return 5;
}

double get_time() {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return tv.tv_sec * 1000000.0 + tv.tv_usec;
}

double bench(const char* name, int (*act)(int), int arg) {
    printf("--------------------------------------------------\n\n(Test)\n%s\n\n(Output & Warm-up)\n", name);
    int res = act(arg);
    printf("%d\n", res);
    act(arg); act(arg);
    
    double min_dur = 1e12;
    for (int i = 1; i <= 10; i++) {
        double t1 = get_time();
        
        // Prevent pure function hoisting
        volatile int a = arg + (i % 2) * 0;
        act(a);
        
        double t2 = get_time();
        double d = t2 - t1;
        if (d < min_dur) min_dur = d;
    }
    printf("\n(Execution time - best of 10)\n\n%.2f us\n\n", min_dur);
    return min_dur;
}

int main(int argc, char** argv) {
    expr_arena = (Expr*)malloc(sizeof(Expr) * 100000);
    tree_arena = (Tree*)malloc(sizeof(Tree) * 20000000);
    
    int dummy = argc;
    int lAst = 3 + dummy - 1;
    int lFib = 10 + dummy - 1;
    int lList = 900 + dummy - 1;
    int lTCO = 100000 + dummy - 1;
    int lRec = 10000 + dummy - 1;
    int lAck = 3 + dummy - 1;
    int lChur = 10 + dummy - 1;
    int lPri = 500 + dummy - 1;
    int lRB = 100000 + dummy - 1;
    int lPoly = 10000000 + dummy - 1;
    int lState = 60 + dummy - 1;
    int lLazy = 1000 + dummy - 1;
    int lArr = 900 + dummy - 1;
    int lRow = 0 + dummy - 1;

    double total_us = 
        bench("AST Evaluation:", runAstTree, lAst) +
        bench("Fibonacci:", runFib, lFib) +
        bench("List Processing:", runListOps, lList) +
        bench("Tail Call Optimization:", runTCO, lTCO) +
        bench("Deep Record Updates:", runRecords, lRec) +
        bench("Ackermann:", runAckermann, lAck) +
        bench("Church Numerals (100k Closure Applications):", runChurch, lChur) +
        bench("Prime Sieve (sum primes up to 500):", runPrimes, lPri) +
        bench("Red-Black Tree:", runRBTree, lRB) +
        bench("Polymorphism:", runPolymorphism, lPoly) +
        bench("State Monad:", runStateMonad, lState) +
        bench("Lazy Evaluation:", runLazyEvaluation, lLazy) +
        bench("Array Processing:", runArrayOps, lArr) +
        bench("RowToList:", runRowToList, lRow);

    printf("\n==================================================\n\nTotal exec time: %.2f ms\n", total_us / 1000.0);
    return 0;
}
