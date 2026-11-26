#include "common.h"

/* The fixture remains immutable. The invocation owns its updated value;
 * optimizing away the copy and local reference storage is permitted. */
static Int update_record(const LibraryInput *input, size_t field_count) {
    Int result[16];
    memcpy(result, input->fields, field_count * sizeof *result);
    for (Int remaining = input->count; remaining != 0;
         remaining = add_int(remaining, -1)) {
        result[field_count - 1] = add_int(result[field_count - 1], 1);
    }
    Int original_sum = 0, result_sum = 0;
    for (size_t i = 0; i < field_count; ++i) {
        original_sum = add_int(original_sum, input->fields[i]);
        result_sum = add_int(result_sum, result[i]);
    }
    return add_int(original_sum, result_sum);
}

Int record_small(const LibraryInput *input) { return update_record(input, 2); }
Int record_medium(const LibraryInput *input) { return update_record(input, 8); }
Int record_large(const LibraryInput *input) { return update_record(input, 16); }

/* ST and Effect refs do not escape their action and have the same imperative
 * semantics. Keep the allocation action-local; C may scalar-replace it. */
Int ref_write(const LibraryInput *input) {
    struct { Int value; } ref = {0};
    /* As in PureScript, count == INT32_MAX does not terminate: incrementing
     * that final index wraps. It is outside the benchmark's input domain. */
    for (Int i = 1; i <= input->count; i = add_int(i, 1)) ref.value = i;
    return ref.value;
}

Int ref_modify(const LibraryInput *input) {
    struct { Int value; } ref = {0};
    for (Int remaining = input->count; remaining > 0;
         remaining = add_int(remaining, -1)) {
        ref.value = add_int(ref.value, 1);
    }
    return ref.value;
}
