/* Harness and arena backing allocation only; algorithms are emitted by Solod. */
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "kernels.h"

int main(int argc, char **argv) {
    if (argc != 4) return 2;
    int64_t count = strtoll(argv[2], NULL, 10);
    int64_t n = strtoll(argv[3], NULL, 10);
    if (n < 0 || n > 100001 || count < 1) return 2;
    int64_t bits = 0;
    for (int64_t k = n; k > 0; k >>= 1) bits++;
    int64_t capacity = (n+1)*(6*bits+4);
    kernels_Constructor_Test_RBTree_T *pool = calloc((size_t)capacity, sizeof(*pool));
    if (!pool) return 3;
    kernels_BeginArena((so_Slice){.ptr=pool, .len=capacity, .cap=capacity});
    if (strcmp(argv[1], "verify") == 0) {
        kernels_Constructor_Test_RBTree_T *root = kernels_Call_Test_RBTree_buildTree(n,NULL);
        kernels_Validation v = kernels_Validate(root,n);
        if (v.Depth != kernels_Call_Test_RBTree_depth(root)) abort();
        printf("%lld %lld %lld %lld %lld\n", (long long)v.Count,(long long)v.Sum,
               (long long)v.Depth,(long long)v.BlackHeight,(long long)kernels_NodesUsed());
        free(pool);
        return 0;
    }
    if (strcmp(argv[1], "bench") != 0) { free(pool); return 2; }
    int64_t sum = 0;
    for (int64_t i = 0; i < count; i++) {
        kernels_ResetArena();
        kernels_Constructor_Test_RBTree_T *root = kernels_Call_Test_RBTree_buildTree(n,NULL);
        sum += kernels_Call_Test_RBTree_depth(root);
    }
    printf("%lld\n", (long long)sum);
    free(pool);
    return 0;
}
