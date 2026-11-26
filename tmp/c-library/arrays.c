#include "common.h"

#include <string.h>

/* The source results are immediately consumed by a weighted checksum. Fuse
 * those pure operations while retaining the result index (not the source
 * index) as the weight for filter and append. All arithmetic wraps as Int32. */
static Int weighted_checksum(const Int *values, size_t length) {
    Int result = 0;
    for (size_t i = 0; i < length; ++i)
        result = add_int(result, mul_int(values[i], (Int)(1 + (i & 1))));
    return result;
}

Int array_map(const LibraryInput *input) {
    Int result = 0;
    for (size_t i = 0; i < input->left.length; ++i)
        result = add_int(result, mul_int(add_int(input->left.data[i], 1),
                                        (Int)(1 + (i & 1))));
    return result;
}

Int array_filter(const LibraryInput *input) {
    Int result = 0;
    size_t output_index = 0;
    for (size_t i = 0; i < input->left.length; ++i) {
        Int value = input->left.data[i];
        if (value % 2 == 0) {
            result = add_int(result, mul_int(value, (Int)(1 + (output_index & 1))));
            ++output_index;
        }
    }
    return result;
}

Int array_foldl(const LibraryInput *input) {
    Int result = 0;
    for (size_t i = 0; i < input->left.length; ++i)
        result = add_int(result, input->left.data[i]);
    return result;
}

/* Discover a natural ordered run. A decreasing run is reversed only after
 * inspecting its elements; arbitrary data is handled by the merge passes.
 * Strict decreasing runs preserve stable ordering of equal elements. */
static size_t ordered_run(Int *values, size_t start, size_t length) {
    size_t end = start + 1;
    if (end == length) return end;
    if (values[start] > values[end]) {
        while (++end < length && values[end - 1] > values[end]) {}
        size_t left = start, right = end;
        while (left < --right) {
            Int value = values[left];
            values[left++] = values[right];
            values[right] = value;
        }
    } else {
        while (++end < length && values[end - 1] <= values[end]) {}
    }
    return end;
}

Int array_sort(const LibraryInput *input) {
    size_t length = input->left.length;
    if (length < 2) return weighted_checksum(input->left.data, length);
    Int *values = allocate_array(length, sizeof *values);
    memcpy(values, input->left.data, length * sizeof *values);
    size_t first_end = ordered_run(values, 0, length);
    Int *scratch = NULL;
    if (first_end != length) scratch = allocate_array(length, sizeof *scratch);
    while (first_end != length) {
        size_t start = 0;
        while (start < length) {
            size_t middle = start == 0 ? first_end : ordered_run(values, start, length);
            if (middle == length) {
                memcpy(scratch + start, values + start, (length - start) * sizeof *values);
                break;
            }
            size_t end = ordered_run(values, middle, length);
            size_t left = start, right = middle, out = start;
            while (left < middle && right < end)
                scratch[out++] = values[left] <= values[right] ? values[left++] : values[right++];
            while (left < middle) scratch[out++] = values[left++];
            while (right < end) scratch[out++] = values[right++];
            start = end;
        }
        Int *previous = values;
        values = scratch;
        scratch = previous;
        first_end = ordered_run(values, 0, length);
    }
    Int result = weighted_checksum(values, length);
    free(scratch);
    free(values);
    return result;
}

Int array_eq(const LibraryInput *input) {
    if (input->left.length != input->right.length) return 0;
    return input->left.length == 0 ||
        memcmp(input->left.data, input->right.data, input->left.length * sizeof(Int)) == 0;
}

Int array_append(const LibraryInput *input) {
    Int result = weighted_checksum(input->left.data, input->left.length);
    for (size_t i = 0; i < input->right.length; ++i)
        result = add_int(result, mul_int(input->right.data[i],
                         (Int)(1 + ((input->left.length + i) & 1))));
    return result;
}

Int traverse_maybe(const LibraryInput *input) {
    return array_map(input);
}

Int traverse_maybe_fail(const LibraryInput *input) {
    Int result = 0;
    for (size_t i = 0; i < input->left.length; ++i) {
        Int value = input->left.data[i];
        /* The callback is pure, so later callback evaluations cannot change
         * Nothing or produce observable effects once rejection occurs. */
        if (value == input->count) return -1;
        result = add_int(result, mul_int(add_int(value, 1), (Int)(1 + (i & 1))));
    }
    return result;
}
