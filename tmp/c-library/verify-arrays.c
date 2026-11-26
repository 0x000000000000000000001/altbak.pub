#include "common.h"

#include <assert.h>
#include <limits.h>
#include <stdbool.h>

/* An independent materializing oracle for the fused benchmark kernels. */
static Int checksum(const Int *values, size_t length) {
    uint32_t total = 0;
    for (size_t i = 0; i < length; ++i)
        total += (uint32_t)values[i] * (uint32_t)(1 + i % 2);
    return from_bits(total);
}

static int compare_int(const void *left, const void *right) {
    Int a = *(const Int *)left, b = *(const Int *)right;
    return (a > b) - (a < b);
}

static uint32_t random_state = 0xc01dcafe;
static uint32_t next_random(void) {
    random_state ^= random_state << 13;
    random_state ^= random_state >> 17;
    random_state ^= random_state << 5;
    return random_state;
}

static void verify_variant(Int *left, size_t left_length, Int *right, size_t right_length,
                           Int rejected) {
    LibraryInput input = {.left = {left, left_length}, .right = {right, right_length}, .count = rejected};
    Int *left_before = allocate_array(left_length, sizeof(Int));
    Int *right_before = allocate_array(right_length, sizeof(Int));
    Int *result = allocate_array(left_length + right_length + 1, sizeof(Int));
    if (left_length) memcpy(left_before, left, left_length * sizeof(Int));
    if (right_length) memcpy(right_before, right, right_length * sizeof(Int));

    uint32_t sum = 0;
    bool rejected_found = false;
    for (size_t i = 0; i < left_length; ++i) {
        result[i] = from_bits((uint32_t)left[i] + 1);
        sum += (uint32_t)left[i];
        if (left[i] == rejected) rejected_found = true;
    }
    Int mapped = checksum(result, left_length);
    assert(array_map(&input) == mapped);
    assert(traverse_maybe(&input) == mapped);
    assert(traverse_maybe_fail(&input) == (rejected_found ? -1 : mapped));
    assert(array_foldl(&input) == from_bits(sum));

    size_t selected = 0;
    for (size_t i = 0; i < left_length; ++i)
        if (left[i] % 2 == 0) result[selected++] = left[i];
    assert(array_filter(&input) == checksum(result, selected));

    if (left_length) memcpy(result, left, left_length * sizeof(Int));
    qsort(result, left_length, sizeof(Int), compare_int);
    assert(array_sort(&input) == checksum(result, left_length));

    bool equal = left_length == right_length;
    if (equal) for (size_t i = 0; i < left_length; ++i) equal = equal && left[i] == right[i];
    assert(array_eq(&input) == equal);
    if (left_length) memcpy(result, left, left_length * sizeof(Int));
    if (right_length) memcpy(result + left_length, right, right_length * sizeof(Int));
    assert(array_append(&input) == checksum(result, left_length + right_length));

    assert(left_length == 0 || memcmp(left, left_before, left_length * sizeof(Int)) == 0);
    assert(right_length == 0 || memcmp(right, right_before, right_length * sizeof(Int)) == 0);
    free(left_before);
    free(right_before);
    free(result);
}

int main(void) {
    Int left[10000], right[10000];
    for (size_t iteration = 0; iteration < 2500; ++iteration) {
        size_t length = iteration < 258 ? iteration : next_random() % 258;
        size_t right_length = iteration % 2 ? length : next_random() % 258;
        for (size_t i = 0; i < length; ++i) {
            switch (iteration % 7) {
            case 0: left[i] = from_bits(next_random()); break;
            case 1: left[i] = (Int)(next_random() % 9) - 4; break;
            case 2: left[i] = (Int)i; break;
            case 3: left[i] = (Int)(length - i); break;
            case 4: left[i] = i % 2 ? INT32_MIN : INT32_MAX; break;
            case 5: left[i] = 17; break;
            default: left[i] = (Int)((i / 8) * 8 + (7 - i % 8)); break;
            }
        }
        for (size_t i = 0; i < right_length; ++i)
            right[i] = i < length ? left[i] : from_bits(next_random());
        if (right_length && iteration % 3 == 0) right[right_length - 1] ^= 1;
        Int rejected = length && iteration % 2 ? left[next_random() % length] : from_bits(next_random());
        verify_variant(left, length, right, right_length, rejected);
    }

    for (size_t i = 0; i < 10000; ++i) left[i] = right[i] = (Int)i + 1;
    LibraryInput input = {.left = {left, 10000}, .right = {right, 10000}};
    assert(array_map(&input) == 75025000);
    assert(array_filter(&input) == 37510000);
    assert(array_foldl(&input) == 50005000);
    right[9999] = 10001;
    assert(array_eq(&input) == 0);
    for (size_t i = 0; i < 10000; ++i) left[i] = 10000 - (Int)i;
    assert(array_sort(&input) == 75010000);
    for (size_t i = 0; i < 1000; ++i) { left[i] = (Int)i + 1; right[i] = (Int)i + 1001; }
    input.left.length = input.right.length = 1000;
    assert(array_append(&input) == 3002000);
    assert(traverse_maybe(&input) == 752500);
    input.count = 1;
    assert(traverse_maybe_fail(&input) == -1);
    input.count = 1000;
    assert(traverse_maybe_fail(&input) == -1);
    puts("Array verification passed: 9 official oracles and 2500 variants, including empty arrays, equality, rejected values, arbitrary sort runs, duplicates, Int32 extremes, and immutable inputs.");
    return 0;
}
