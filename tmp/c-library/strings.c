#include "common.h"

static int high_surrogate(uint16_t unit) { return unit >= 0xd800 && unit <= 0xdbff; }
static int low_surrogate(uint16_t unit) { return unit >= 0xdc00 && unit <= 0xdfff; }

static uint32_t next_code_point(const uint16_t *text, size_t end, size_t *index) {
    uint32_t first = text[(*index)++];
    if (high_surrogate((uint16_t)first) && *index < end && low_surrogate(text[*index])) {
        uint32_t second = text[(*index)++];
        return UINT32_C(0x10000) + ((first - UINT32_C(0xd800)) << 10) + second - UINT32_C(0xdc00);
    }
    return first;
}

static Int checksum(const uint16_t *text, size_t begin, size_t end) {
    uint32_t sum = 0;
    while (begin < end) sum += next_code_point(text, end, &begin);
    return from_bits(sum);
}

Int string_cu_length(const LibraryInput *input) {
    return from_bits((uint32_t)input->text.length);
}

Int string_cp_length(const LibraryInput *input) {
    const CString text = input->text;
    uint32_t count = 0;
    for (size_t index = 0; index < text.length; ++count)
        (void)next_code_point(text.data, text.length, &index);
    return from_bits(count);
}

/* A 256-slot hashed Horspool table keeps setup small for UTF-16 strings.
 * Collisions shorten a skip, so they cannot skip a possible full match. */
typedef struct { CString pattern; size_t shifts[256]; } Search;

static void prepare_search(Search *search, CString pattern) {
    search->pattern = pattern;
    if (pattern.length > 1) {
        for (size_t i = 0; i < 256; ++i) search->shifts[i] = pattern.length;
        for (size_t i = 0; i + 1 < pattern.length; ++i)
            search->shifts[pattern.data[i] & 255] = pattern.length - 1 - i;
    }
}

static size_t find_units(CString text, const Search *search, size_t begin) {
    const CString pattern = search->pattern;
    if (pattern.length == 0) return begin <= text.length ? begin : SIZE_MAX;
    if (begin > text.length || pattern.length > text.length - begin) return SIZE_MAX;
    if (pattern.length == 1) {
        for (size_t i = begin; i < text.length; ++i)
            if (text.data[i] == pattern.data[0]) return i;
        return SIZE_MAX;
    }
    const size_t last = pattern.length - 1;
    const size_t limit = text.length - pattern.length;
    size_t offset = begin;
    while (offset <= limit) {
        const uint16_t tail = text.data[offset + last];
        if (tail == pattern.data[last] && memcmp(text.data + offset, pattern.data, last * sizeof(uint16_t)) == 0)
            return offset;
        const size_t skip = search->shifts[tail & 255];
        if (skip > limit - offset) break;
        offset += skip;
    }
    return SIZE_MAX;
}

Int string_index_of(const LibraryInput *input) {
    Search search;
    prepare_search(&search, input->pattern);
    const size_t index = find_units(input->text, &search, 0);
    return index == SIZE_MAX ? -1 : from_bits((uint32_t)index);
}

Int string_cp_take(const LibraryInput *input) {
    /* Fuse the prefix and its only consumer; count is still read at runtime. */
    const CString text = input->text;
    uint32_t sum = 0;
    size_t index = 0;
    for (Int remaining = input->count; remaining > 0 && index < text.length; --remaining)
        sum += next_code_point(text.data, text.length, &index);
    return from_bits(sum);
}

Int string_split(const LibraryInput *input) {
    const CString text = input->text;
    Search search;
    prepare_search(&search, input->pattern);
    if (search.pattern.length == 0) {
        /* JS split("") yields isolated UTF-16 units, even for surrogate pairs. */
        uint32_t total = (uint32_t)text.length;
        for (size_t i = 0; i < text.length; ++i) total += text.data[i];
        return from_bits(total);
    }
    Int total = 1;
    size_t begin = 0;
    for (;;) {
        const size_t match = find_units(text, &search, begin);
        if (match == SIZE_MAX) return add_int(total, checksum(text.data, begin, text.length));
        /* Decode each fragment separately: removing a separator must not join
         * a high surrogate on its left to a low surrogate on its right. */
        total = add_int(add_int(total, checksum(text.data, begin, match)), 1);
        begin = match + search.pattern.length;
    }
}

Int string_concat(const LibraryInput *input) {
    /* Fuse concatenation with its consumer, retaining a pending high surrogate
     * across fragments: a pair may be formed only after their concatenation. */
    uint32_t sum = 0;
    uint16_t pending = 0;
    for (size_t part = 0; part < input->string_count; ++part) {
        const CString text = input->strings[part];
        for (size_t i = 0; i < text.length; ++i) {
            const uint16_t unit = text.data[i];
            if (pending) {
                if (low_surrogate(unit)) {
                    sum += UINT32_C(0x10000) + (((uint32_t)pending - UINT32_C(0xd800)) << 10) + unit - UINT32_C(0xdc00);
                    pending = 0;
                    continue;
                }
                sum += pending;
                pending = 0;
            }
            if (high_surrogate(unit)) pending = unit;
            else sum += unit;
        }
    }
    return from_bits(sum + pending);
}
