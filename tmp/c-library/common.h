#ifndef ALTBAK_C_LIBRARY_COMMON_H
#define ALTBAK_C_LIBRARY_COMMON_H
#include <stdint.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef int32_t Int;
typedef struct { const Int *data; size_t length; } CArray;
/* Native sized UTF-16 strings: length counts code units, not code points. */
typedef struct { const uint16_t *data; size_t length; } CString;
typedef struct {
    CArray left, right;
    CString text, pattern;
    const CString *strings;
    size_t string_count;
    Int count;
    Int fields[16];
} LibraryInput;
typedef Int (*LibraryAction)(const LibraryInput *);

static inline Int from_bits(uint32_t bits) {
    Int result;
    memcpy(&result, &bits, sizeof result);
    return result;
}
static inline Int add_int(Int a, Int b) { return from_bits((uint32_t)a + (uint32_t)b); }
static inline Int mul_int(Int a, Int b) { return from_bits((uint32_t)a * (uint32_t)b); }
static inline void *allocate_array(size_t count, size_t size) {
    if (size && count > SIZE_MAX / size) { fputs("Allocation overflow\n", stderr); exit(1); }
    void *result = malloc(count * size);
    if (!result && count && size) { perror("malloc"); exit(1); }
    return result;
}

Int array_map(const LibraryInput *), array_filter(const LibraryInput *);
Int array_foldl(const LibraryInput *), array_sort(const LibraryInput *);
Int array_eq(const LibraryInput *), array_append(const LibraryInput *);
Int traverse_maybe(const LibraryInput *), traverse_maybe_fail(const LibraryInput *);
Int string_cu_length(const LibraryInput *), string_cp_length(const LibraryInput *);
Int string_index_of(const LibraryInput *), string_cp_take(const LibraryInput *);
Int string_split(const LibraryInput *), string_concat(const LibraryInput *);
Int record_small(const LibraryInput *), record_medium(const LibraryInput *);
Int record_large(const LibraryInput *), ref_write(const LibraryInput *);
Int ref_modify(const LibraryInput *);
#endif
