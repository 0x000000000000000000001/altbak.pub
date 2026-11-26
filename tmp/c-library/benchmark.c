/* Prepared inputs follow the PureScript prepare functions; only their returned action is timed.
 * Compile kernels as separate translation units without LTO. */
#include "common.h"
#include <inttypes.h>
#include <stdatomic.h>
#include <time.h>

typedef struct { const char *name; LibraryAction action; Int expected; LibraryInput input; } Case;
static Case cases[] = {
    {"array/map-10000", array_map, 75025000, {0}},
    {"array/filter-10000", array_filter, 37510000, {0}},
    {"array/foldl-sum-10000", array_foldl, 50005000, {0}},
    {"array/sort-descending-10000", array_sort, 75010000, {0}},
    {"array/eq-neq-last-10000", array_eq, 0, {0}},
    {"array/append-1000", array_append, 3002000, {0}},
    {"foldable/traverse-maybe-1000", traverse_maybe, 752500, {0}},
    {"foldable/traverse-maybe-fail-first-1000", traverse_maybe_fail, -1, {0}},
    {"foldable/traverse-maybe-fail-last-1000", traverse_maybe_fail, -1, {0}},
    {"string/cu-length-unicode-1000", string_cu_length, 4000, {0}},
    {"string/cp-length-unicode-1000", string_cp_length, 3000, {0}},
    {"string/indexOf-hit-10000", string_index_of, 10000, {0}},
    {"string/indexOf-miss-10000", string_index_of, -1, {0}},
    {"string/cp-take-unicode-1500", string_cp_take, 64421000, {0}},
    {"string/split-1000", string_split, 295000, {0}},
    {"string/concat-1000", string_concat, 195000, {0}},
    {"record/set-small-1000", record_small, 1006, {0}},
    {"record/set-medium-1000", record_medium, 1072, {0}},
    {"record/set-large-1000", record_large, 1272, {0}},
    {"st/ref-write-10000", ref_write, 10000, {0}},
    {"st/ref-modify-10000", ref_modify, 10000, {0}},
    {"effect/ref-write-10000", ref_write, 10000, {0}},
    {"effect/ref-modify-10000", ref_modify, 10000, {0}},
};
#define CASE_COUNT (sizeof cases / sizeof *cases)

static Int ascending[10000], descending[10000], unequal[10000];
static uint16_t unicode_text[4000], search_text[10006], split_text[3999];
static const uint16_t needle[] = {'n','e','e','d','l','e'}, absent[] = {'a','b','s','e','n','t'};
static const uint16_t comma[] = {','}, ab[] = {'a','b'};
static CString fragments[1000];

static void prepare(void) {
    for (size_t i = 0; i < 10000; ++i) {
        ascending[i] = (Int)i + 1;
        descending[i] = 10000 - (Int)i;
        unequal[i] = (Int)i + 1;
        search_text[i] = 'a';
    }
    unequal[9999] = 10001;
    memcpy(search_text + 10000, needle, sizeof needle);
    for (size_t i = 0; i < 1000; ++i) {
        unicode_text[i*4] = 'a'; unicode_text[i*4+1] = 0xd83d;
        unicode_text[i*4+2] = 0xde00; unicode_text[i*4+3] = 0xe9;
        split_text[i*4] = 'a'; split_text[i*4+1] = 'b'; split_text[i*4+2] = 'c';
        if (i < 999) split_text[i*4+3] = ',';
        fragments[i] = (CString){ab, 2};
    }
    for (size_t i = 0; i < CASE_COUNT; ++i) {
        LibraryInput *input = &cases[i].input;
        input->left = (CArray){ascending, i < 5 ? 10000 : 1000};
        input->text = (CString){unicode_text, 4000};
        input->count = 1000;
        for (Int field = 0; field < 16; ++field) input->fields[field] = field + 1;
    }
    cases[3].input.left.data = descending;
    cases[4].input.right = (CArray){unequal, 10000};
    cases[5].input.right = (CArray){ascending + 1000, 1000};
    cases[7].input.count = 1; cases[8].input.count = 1000;
    cases[11].input.text = cases[12].input.text = (CString){search_text, 10006};
    cases[11].input.pattern = (CString){needle, 6};
    cases[12].input.pattern = (CString){absent, 6};
    cases[13].input.count = 1500;
    cases[14].input.text = (CString){split_text, 3999};
    cases[14].input.pattern = (CString){comma, 1};
    cases[15].input.strings = fragments; cases[15].input.string_count = 1000;
    for (size_t i = 19; i < CASE_COUNT; ++i) cases[i].input.count = 10000;
}

/* Hide both the address and memory reachable through it on every invocation.
 * The kernels live in other translation units; no LTO or result caching. */
__attribute__((noinline))
static const LibraryInput *opaque_input(const LibraryInput *input) {
    __asm__ volatile("" : "+r"(input) : : "memory");
    return input;
}
static _Atomic Int sink;
static Int invoke(const Case *test) {
    Int result = test->action(opaque_input(&test->input));
    atomic_store_explicit(&sink, result, memory_order_seq_cst);
    return result;
}
static void expect(const Case *test, Int actual) {
    if (actual != test->expected) {
        fprintf(stderr, "%s: got %" PRId32 ", expected %" PRId32 "\n", test->name, actual, test->expected);
        exit(1);
    }
}
static uint64_t now_ns(void) {
    struct timespec now;
    if (clock_gettime(CLOCK_MONOTONIC, &now)) { perror("clock_gettime"); exit(1); }
    return (uint64_t)now.tv_sec * UINT64_C(1000000000) + (uint64_t)now.tv_nsec;
}
__attribute__((noinline))
static uint64_t batch(const Case *test, int iterations) {
    uint64_t start = now_ns();
    for (int i = 0; i < iterations; ++i) invoke(test);
    return now_ns() - start;
}
static void bench(const Case *test) {
    for (int i = 0; i < 3; ++i) expect(test, invoke(test));
    int iterations = 1;
    while (batch(test, iterations) < UINT64_C(10000000) && iterations < 16777216) iterations *= 2;
    uint64_t durations[10], best = UINT64_MAX;
    for (size_t i = 0; i < 10; ++i) {
        durations[i] = batch(test, iterations);
        expect(test, atomic_load_explicit(&sink, memory_order_seq_cst));
        if (durations[i] < best) best = durations[i];
    }
    printf("{\"case\":\"%s\",\"value\":%" PRId32 ",\"iterations\":%d,\"minimum_us\":%.12f,\"batches_ns\":[",
           test->name, test->expected, iterations, (double)best / 1000.0 / iterations);
    for (size_t i = 0; i < 10; ++i) printf("%s%" PRIu64, i ? "," : "", durations[i]);
    puts("]}"); fflush(stdout);
}
int main(int argc, char **argv) {
    prepare();
    if (argc == 2 && !strcmp(argv[1], "--verify")) {
        for (size_t i = 0; i < CASE_COUNT; ++i) { expect(&cases[i], invoke(&cases[i])); printf("Verified %s\n", cases[i].name); }
        /* Sorting and record updates must preserve prepared fixtures. */
        for (size_t i = 0; i < 10000; ++i)
            if (ascending[i] != (Int)i + 1 || descending[i] != 10000 - (Int)i || unequal[i] != (Int)i + (i == 9999 ? 2 : 1)) abort();
        for (size_t i = 16; i < 19; ++i)
            for (Int j = 0; j < 16; ++j) if (cases[i].input.fields[j] != j + 1) abort();
        puts("Verified 23 library oracles and immutable fixtures.");
        return 0;
    }
    if (argc != 1) { fputs("Usage: benchmark-c-library [--verify]\n", stderr); return 2; }
    for (int warmup = 0; warmup < 3; ++warmup)
        for (size_t i = 0; i < CASE_COUNT; ++i) expect(&cases[i], invoke(&cases[i]));
    for (size_t i = 0; i < CASE_COUNT; ++i) bench(&cases[i]);
    return 0;
}
