// Timing/input harness only: both indexed-read loops come from the diagnostic
// kernel (Test.ArrayIndexing). C has no runtime boxing, so the "native" case
// indexes an int64_t array and the "boxed" case indexes a 24-byte tagged Value
// array (the same layout as the Go runtime's Value: type, int, pointer). Only
// the selected element is converted in both loops, as in the PureScript kernel.
#include <inttypes.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

typedef struct {
  uint8_t tag;
  int64_t int_value;
  void *pointer;
} Value;

static volatile int64_t sink;

static inline int64_t native_reads(const int64_t *source, int64_t size, int64_t count, int64_t start) {
  int64_t checksum = 0;
  int64_t index = start;
  for (int64_t remaining = count; remaining != 0; remaining--) {
    checksum += source[index];
    index += 1;
    if (index == size) {
      index = 0;
    }
  }
  return checksum;
}

static inline int64_t boxed_reads(const Value *source, int64_t size, int64_t count, int64_t start) {
  int64_t checksum = 0;
  int64_t index = start;
  for (int64_t remaining = count; remaining != 0; remaining--) {
    checksum += source[index].int_value;
    index += 1;
    if (index == size) {
      index = 0;
    }
  }
  return checksum;
}

static int64_t oracle(const int64_t *values, int64_t size, int64_t count, int64_t start) {
  int64_t total = 0;
  for (int64_t i = 0; i < size; i++) total += values[i];
  int64_t result = total * (count / size);
  for (int64_t i = 0; i < count % size; i++) result += values[(start + i) % size];
  return result;
}

static double now_nanoseconds(void) {
  struct timespec timestamp;
  clock_gettime(CLOCK_MONOTONIC, &timestamp);
  return (double)timestamp.tv_sec * 1e9 + (double)timestamp.tv_nsec;
}

int main(int argc, char **argv) {
  int64_t accesses = 1 << 23;
  int64_t batches = 10;
  int64_t seed = 5;
  for (int index = 1; index < argc; index++) {
    if (strcmp(argv[index], "-accesses") == 0 && index + 1 < argc) {
      accesses = atoll(argv[++index]);
    } else if (strcmp(argv[index], "-batches") == 0 && index + 1 < argc) {
      batches = atoll(argv[++index]);
    } else if (strcmp(argv[index], "-seed") == 0 && index + 1 < argc) {
      seed = atoll(argv[++index]);
    } else {
      fprintf(stderr, "unknown argument: %s\n", argv[index]);
      return 2;
    }
  }
  if (accesses < 16384 || accesses > (1 << 23) || batches < 3 || seed < 0 || seed > 1000) {
    fprintf(stderr, "invalid accesses, batches or seed\n");
    return 2;
  }
  const int64_t sizes[3] = {16, 1024, 16384};
  for (int slot = 0; slot < 3; slot++) {
    int64_t size = sizes[slot];
    int64_t *native = calloc((size_t)size, sizeof(int64_t));
    Value *boxed = calloc((size_t)size, sizeof(Value));
    if (native == NULL || boxed == NULL) return 3;
    for (int64_t i = 0; i < size; i++) {
      native[i] = (i * 17 + seed * 31) % 251 + 1;
      boxed[i].tag = 0;
      boxed[i].int_value = native[i];
      boxed[i].pointer = NULL;
    }
    int64_t start = seed % size;
    int64_t expected = oracle(native, size, accesses, start);
    for (int representation = 0; representation < 2; representation++) {
      int boxed_case = representation == 1;
      for (int warm = 0; warm < 3; warm++) {
        sink = boxed_case ? boxed_reads(boxed, size, accesses, start)
                          : native_reads(native, size, accesses, start);
        if (sink != expected) {
          fprintf(stderr, "warm-up checksum mismatch\n");
          return 4;
        }
      }
      printf("{\"runtime\":\"c\",\"representation\":\"%s\",\"size\":%" PRId64
             ",\"accesses\":%" PRId64 ",\"seed\":%" PRId64 ",\"checksum\":%" PRId64
             ",\"warmups\":3,\"ns_per_access\":[",
             boxed_case ? "boxed" : "native", size, accesses, seed, (int64_t)sink);
      for (int64_t batch = 0; batch < batches; batch++) {
        double begin = now_nanoseconds();
        sink = boxed_case ? boxed_reads(boxed, size, accesses, start)
                          : native_reads(native, size, accesses, start);
        double elapsed = now_nanoseconds() - begin;
        if (sink != expected || elapsed <= 0) {
          fprintf(stderr, "invalid measured result\n");
          return 5;
        }
        printf("%s%.9f", batch ? "," : "", elapsed / (double)accesses);
      }
      printf("],\"bytes_per_access\":null}\n");
    }
    free(native);
    free(boxed);
  }
  return 0;
}
