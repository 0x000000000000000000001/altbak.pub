/* Harness only: all algorithms live in the C emitted by Solod. Separate
 * translation units and no LTO prevent cross-call elimination/hoisting. */
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "kernels.h"

int main(int argc, char **argv) {
    if (argc != 4) return 2;
    int64_t count = strtoll(argv[2], NULL, 10);
    int64_t n = strtoll(argv[3], NULL, 10);
    int64_t sum = 0;
    if (strcmp(argv[1], "fib") == 0) {
        for (int64_t i = 0; i < count; i++)
            sum += kernels_Call_Test_Fib_fib(n + (i & 1));
    } else if (strcmp(argv[1], "ackermann") == 0) {
        for (int64_t i = 0; i < count; i++)
            sum += kernels_Call_Test_Ackermann_ackermann(3, n + (i & 1));
    } else if (strcmp(argv[1], "tco") == 0) {
        for (int64_t i = 0; i < count; i++)
            sum += kernels_Call_Test_TCO_deepTailRec(n + (i & 1), 0);
    } else return 2;
    printf("%lld\n", (long long)sum);
    return 0;
}
