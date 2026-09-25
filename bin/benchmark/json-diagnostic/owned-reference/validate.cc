// Standalone semantic and ownership audit. Build once for each reference with
// -DREFERENCE_TAST for the typed module decoder; sanitizer builds are untimed.
#define main reference_benchmark_main
#ifdef REFERENCE_TAST
#include "../typed-ast.cc"
#else
#include "../decoding.cc"
#endif
#undef main

#include <cassert>

int main() {
  std::string corpus;
  if (!read_text(getenv("DIAG_CORPUS"), corpus)) return 2;
  simdjson::dom::parser corpus_parser;
  std::string result = "[";
  bool first = true;
  size_t retained = 0;
  for (element entry : array_of(corpus_parser.parse(corpus).value())) {
    object input = object_of(entry);
    std::string name(string_of(require_field(input, "name")));
    std::string text(string_of(require_field(input, "contents")));
    std::string hash;
    bool accepted = false;
    try {
#ifdef REFERENCE_TAST
      Module value;
#else
      Payload value;
#endif
      {
        simdjson::dom::parser parser;
        element root = parser.parse(text).value();
#ifdef REFERENCE_TAST
        value = decode_module(root, "");
#else
        decode_payload(root, value);
#endif
      } // Destroy the DOM before reading any of the result.
      std::fill(text.begin(), text.end(), 'x');
      text.clear();
      text.shrink_to_fit();
#ifdef REFERENCE_TAST
      hash = sha256_hex(module_fingerprint(value));
      // A detached executable subtree and its shared types must outlive the
      // original module, independently of every parser and temporary table.
      if (!value.decls.empty()) {
        Binding &binding = value.decls.front().kind == 0
            ? value.decls.front().single : value.decls.front().group.front();
        std::string module_name = value.name;
        std::string expected;
        write_expr(expected, *binding.expr, module_name);
        auto child = std::move(binding.expr);
        value = Module{};
        std::string actual;
        write_expr(actual, *child, module_name);
        assert(actual == expected);
        retained++;
      }
#else
      hash = payload_fingerprint(value);
      if (!value.users.empty()) {
        std::string expected;
        append_user(expected, value.users.front());
        User child = std::move(value.users.front());
        value = Payload{};
        std::string actual;
        append_user(actual, child);
        assert(actual == expected);
        retained++;
      }
#endif
      accepted = true;
    } catch (...) {}
    if (!first) result += ',';
    first = false;
    result += "{\"name\":" + json_quoted(name) + ",\"accepted\":" +
        (accepted ? "true" : "false") + ",\"fingerprint\":" +
        (accepted ? json_quoted(hash) : "null") + "}";
  }
  fprintf(stderr, "%zu detached children validated after parent/input/parser destruction\n", retained);
  printf("%s]\n", result.c_str());
}
