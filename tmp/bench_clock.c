#include <stdint.h>
#include <time.h>

#include <caml/alloc.h>
#include <caml/fail.h>
#include <caml/mlvalues.h>

CAMLprim value altbak_monotonic_ns(value unit)
{
    struct timespec now;
    (void)unit;
    if (clock_gettime(CLOCK_MONOTONIC, &now) != 0) {
        caml_failwith("clock_gettime(CLOCK_MONOTONIC) failed");
    }
    return caml_copy_int64((int64_t)now.tv_sec * INT64_C(1000000000) + now.tv_nsec);
}
