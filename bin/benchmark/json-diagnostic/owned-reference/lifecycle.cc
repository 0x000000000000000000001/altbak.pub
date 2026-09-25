// Fresh-process construction/two retained decodes/destruction audit.
#define main reference_benchmark_main
#ifdef REFERENCE_TAST
#include "../typed-ast.cc"
#else
#include "../decoding.cc"
#endif
#undef main

int main() {
  std::string text;
  if (!read_text(getenv("DIAG_CORPUS"), text)) return 2;
  simdjson::dom::parser corpus_parser;
  std::vector<Case> cases;
  for (element raw : array_of(corpus_parser.parse(text).value())) {
    object input = object_of(raw);
#ifndef REFERENCE_TAST
    if (!bool_of(require_field(input, "benchmark"))) continue;
#endif
    Case item;
    item.name = string_of(require_field(input, "name"));
    item.contents = string_of(require_field(input, "contents"));
    cases.push_back(std::move(item));
  }
#ifdef REFERENCE_TAST
  using Result = Module;
#else
  using Result = Payload;
#endif
  auto decode = [](const std::string &input) {
    simdjson::dom::parser parser;
    element root = parser.parse(input).value();
#ifdef REFERENCE_TAST
    return decode_module(root, "");
#else
    Payload value;
    decode_payload(root, value);
    return value;
#endif
  };
  // There is no decoder/getter state to preconstruct in this C++ API. Parser
  // construction/allocation/destruction is included in each full decode.
  std::vector<Result> first(cases.size()), second(cases.size());
  auto pass = [&](std::vector<Result> &outputs) {
    double begin = now_microseconds();
    for (size_t index = 0; index < cases.size(); index++) outputs[index] = decode(cases[index].contents);
    return now_microseconds() - begin;
  };
  double first_us = pass(first), second_us = pass(second);
  auto hashes = [](const std::vector<Result> &outputs) {
    std::string out = "[";
    for (size_t index = 0; index < outputs.size(); index++) {
      if (index) out += ',';
#ifdef REFERENCE_TAST
      out += json_quoted(sha256_hex(module_fingerprint(outputs[index])));
#else
      out += json_quoted(payload_fingerprint(outputs[index]));
#endif
    }
    return out + ']';
  };
  std::string first_hashes = hashes(first), second_hashes = hashes(second);
  double begin = now_microseconds();
  first.clear(); second.clear();
  double release_us = now_microseconds() - begin;
  printf("{\"construction\":{\"time_us\":0},\"first_combined\":{\"time_us\":%.9f},"
         "\"second_combined\":{\"time_us\":%.9f},\"retained_collection\":{\"time_us\":0},"
         "\"release_collection\":{\"time_us\":%.9f},"
         "\"first_fingerprints\":%s,\"second_fingerprints\":%s}\n",
         first_us, second_us, release_us, first_hashes.c_str(), second_hashes.c_str());
}
