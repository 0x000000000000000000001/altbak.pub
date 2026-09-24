// Audit: simdjson behind a cgo boundary.
//
// Two questions:
//   * what does the fastest native parser cost when the call crosses Go/C
//     (one call per document, no per-field access), and
//   * what does a cgo crossing cost per call, to price a naive per-value
//     conversion.
//
// The parse result is consumed in C (node count) so the work cannot be elided.
#include <simdjson.h>

#include <cstdint>
#include <cstddef>

namespace {

simdjson::dom::parser *g_parser = nullptr;

uint64_t count_nodes(simdjson::dom::element value) {
  using simdjson::dom::array;
  using simdjson::dom::element_type;
  using simdjson::dom::object;
  switch (value.type()) {
    case element_type::OBJECT: {
      uint64_t total = 1;
      for (auto field : object(value)) total += 1 + count_nodes(field.value);
      return total;
    }
    case element_type::ARRAY: {
      uint64_t total = 1;
      for (simdjson::dom::element item : array(value)) total += count_nodes(item);
      return total;
    }
    default:
      return 1;
  }
}

}  // namespace

extern "C" {

// Parses one document and returns its node count, or 0 on a parse error.
uint64_t simdjson_parse_count(const char *data, size_t len) {
  if (g_parser == nullptr) g_parser = new simdjson::dom::parser();
  simdjson::dom::element root;
  auto error = g_parser->parse(simdjson::padded_string(data, len)).get(root);
  if (error) return 0;
  return count_nodes(root);
}

void simdjson_release(void) {
  delete g_parser;
  g_parser = nullptr;
}

int cgo_ping(int value) { return value + 1; }

}  // extern "C"
