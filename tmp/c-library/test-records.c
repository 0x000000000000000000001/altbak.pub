#include "common.h"
#include <assert.h>
#include <limits.h>

/* The oracle uses the final-value identity, independently of the mutation
 * loops in records.c. Altering every fixture field detects cached answers. */
static Int record_oracle(const LibraryInput *input, size_t count) {
    uint32_t sum = 0;
    for (size_t i = 0; i < count; ++i) sum += (uint32_t)input->fields[i];
    return from_bits(2u * sum + (uint32_t)input->count);
}

int main(void) {
    const Int counts[] = {0, 1, 2, 7, 999, 1000, 1001, 10000};
    LibraryInput input = {0};
    for (size_t i = 0; i < 16; ++i) input.fields[i] = (Int)i + 1;
    input.count = 1000;
    assert(record_small(&input) == 1006);
    assert(record_medium(&input) == 1072);
    assert(record_large(&input) == 1272);
    uint32_t random = 0x9e3779b9u;
    size_t checks = 3;
    for (size_t variant = 0; variant < 128; ++variant) {
        for (size_t field = 0; field < 16; ++field) {
            random = random * 1664525u + 1013904223u;
            input.fields[field] = from_bits(random);
        }
        for (size_t n = 0; n < sizeof counts / sizeof *counts; ++n) {
            input.count = counts[n];
            LibraryInput original = input;
            assert(record_small(&input) == record_oracle(&input, 2));
            assert(record_medium(&input) == record_oracle(&input, 8));
            assert(record_large(&input) == record_oracle(&input, 16));
            assert(memcmp(&input, &original, sizeof input) == 0);
            checks += 4;
        }
    }
    const Int ref_counts[] = {INT32_MIN, -10000, -1, 0, 1, 2, 7, 9999, 10000, 10001};
    for (size_t n = 0; n < sizeof ref_counts / sizeof *ref_counts; ++n) {
        input.count = ref_counts[n];
        Int expected = input.count > 0 ? input.count : 0;
        LibraryInput original = input;
        assert(ref_write(&input) == expected);
        assert(ref_modify(&input) == expected);
        assert(memcmp(&input, &original, sizeof input) == 0);
        checks += 3;
    }
    printf("Record/ref verification passed: %zu checks.\n", checks);
    return 0;
}
