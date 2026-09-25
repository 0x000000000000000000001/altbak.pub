// JSON Decoding reference (C/C++), same corpus and protocol as the Go and
// JavaScript diagnostics driven by bin/benchmark/json-diagnostic.py.
//
// Scope of the comparison:
//   * parsing uses a fresh simdjson parser for each owned input document;
//   * concrete results own ordinary std::string/std::vector/optional storage;
//     construction and final destruction are reported separately;
//   * fingerprints are computed canonically in this driver, independently of
//     the other runtimes, and validated by the caller against the frozen
//     oracle.
// Error-message formatting is not reproduced: malformed inputs only have to
// fail (the exact Argonaut messages are a PureScript contract and the error
// cases are excluded from timing by the fixture).
#include <CommonCrypto/CommonDigest.h>
#include <simdjson.h>

#include <algorithm>
#include <chrono>
#include <cmath>
#include <cstdint>
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <memory>
#include <optional>
#include <string>
#include <string_view>
#include <vector>

namespace {

using simdjson::dom::array;
using simdjson::dom::element;
using simdjson::dom::object;

struct decode_error {};

element require_field(object obj, const char *key) { return obj[key].value(); }
std::string_view string_of(element value) { return value.get_string().value(); }
int64_t int_of(element value) { return value.get_int64().value(); }
double double_of(element value) { return value.get_double().value(); }
bool bool_of(element value) { return value.get_bool().value(); }
object object_of(element value) { return value.get_object().value(); }
array array_of(element value) { return value.get_array().value(); }

element optional_field(object obj, const char *key, bool &present) {
  element value;
  auto error = obj[key].get(value);
  if (error == simdjson::NO_SUCH_FIELD) {
    present = false;
    return element();
  }
  if (error) throw decode_error{};
  present = true;
  return value;
}

// --------------------------------------------------------------------------
// Canonical JSON writing (same conventions as the fixture oracle).
// --------------------------------------------------------------------------

void append_escaped(std::string &out, std::string_view text) {
  out.push_back('"');
  for (size_t index = 0; index < text.size();) {
    unsigned char byte = static_cast<unsigned char>(text[index]);
    if (byte == 0xe2 && index + 2 < text.size() &&
        static_cast<unsigned char>(text[index + 1]) == 0x80) {
      unsigned char third = static_cast<unsigned char>(text[index + 2]);
      if (third == 0xa8 || third == 0xa9) {
        out += third == 0xa8 ? "\\u2028" : "\\u2029";
        index += 3;
        continue;
      }
    }
    index += 1;
    switch (byte) {
      case '"': out += "\\\""; break;
      case '\\': out += "\\\\"; break;
      case '\n': out += "\\n"; break;
      case '\r': out += "\\r"; break;
      case '\t': out += "\\t"; break;
      case '<': out += "\\u003c"; break;
      case '>': out += "\\u003e"; break;
      case '&': out += "\\u0026"; break;
      default:
        if (byte < 0x20) {
          char scratch[8];
          snprintf(scratch, sizeof(scratch), "\\u%04x", byte);
          out += scratch;
        } else {
          out.push_back(static_cast<char>(byte));
        }
    }
  }
  out.push_back('"');
}

bool is_integral(double value) {
  return std::isfinite(value) && value == std::floor(value) && std::fabs(value) < 9.0e15;
}

void append_number(std::string &out, double value) {
  if (is_integral(value)) {
    out += std::to_string(static_cast<long long>(value));
    return;
  }
  char scratch[40];
  for (int precision = 15; precision <= 17; precision++) {
    snprintf(scratch, sizeof(scratch), "%.*g", precision, value);
    if (strtod(scratch, nullptr) == value) break;
  }
  out += scratch;
}

void append_canonical(std::string &out, element value) {
  switch (value.type()) {
    case simdjson::dom::element_type::OBJECT: {
      std::vector<std::pair<std::string_view, element>> fields;
      for (auto field : object(value)) fields.emplace_back(field.key, field.value);
      std::sort(fields.begin(), fields.end(),
                [](const auto &left, const auto &right) { return left.first < right.first; });
      out.push_back('{');
      bool first = true;
      for (const auto &field : fields) {
        if (!first) out.push_back(',');
        first = false;
        append_escaped(out, field.first);
        out.push_back(':');
        append_canonical(out, field.second);
      }
      out.push_back('}');
      break;
    }
    case simdjson::dom::element_type::ARRAY: {
      out.push_back('[');
      bool first = true;
      for (element item : array(value)) {
        if (!first) out.push_back(',');
        first = false;
        append_canonical(out, item);
      }
      out.push_back(']');
      break;
    }
    case simdjson::dom::element_type::STRING:
      append_escaped(out, string_of(value));
      break;
    case simdjson::dom::element_type::INT64:
      out += std::to_string(value.get_int64().value());
      break;
    case simdjson::dom::element_type::UINT64:
      out += std::to_string(value.get_uint64().value());
      break;
    case simdjson::dom::element_type::DOUBLE:
      append_number(out, double_of(value));
      break;
    case simdjson::dom::element_type::BOOL:
      out += bool_of(value) ? "true" : "false";
      break;
    default:
      out += "null";
  }
}

std::string sha256_hex(std::string_view text) {
  unsigned char digest[CC_SHA256_DIGEST_LENGTH];
  CC_SHA256(text.data(), static_cast<CC_LONG>(text.size()), digest);
  static const char hex[] = "0123456789abcdef";
  std::string out;
  out.reserve(CC_SHA256_DIGEST_LENGTH * 2);
  for (unsigned char byte : digest) {
    out.push_back(hex[byte >> 4]);
    out.push_back(hex[byte & 0x0f]);
  }
  return out;
}

std::string fingerprint_of(element value) {
  std::string text;
  append_canonical(text, value);
  return sha256_hex(text);
}

// --------------------------------------------------------------------------
// Owned results. Every container uses its ordinary standard-library allocator.
// --------------------------------------------------------------------------

using Str = std::string;

Str copy_string(std::string_view text) { return Str(text); }

struct Profile {
  Str city;
  bool has_note = false;
  Str note;
  std::vector<double> scores;
  size_t score_count = 0;
};

struct User {
  int64_t id = 0;
  Str name;
  bool active = false;
  std::optional<Profile> profile;
  std::vector<Str> tags;
  size_t tag_count = 0;
};

struct Item {
  Str sku;
  int64_t quantity = 0;
  double price = 0;
};

struct Event {
  int kind = 0;  // 0 = view, 1 = purchase
  Str path;
  bool has_duration = false;
  int64_t duration = 0;
  int64_t order_id = 0;
  std::vector<Item> items;
  size_t item_count = 0;
};

struct Payload {
  int64_t version = 0;
  bool has_next = false;
  Str next;
  std::vector<User> users;
  size_t user_count = 0;
  std::vector<Event> events;
  size_t event_count = 0;
};

void decode_profile(element raw, Profile &out) {
  object obj = object_of(raw);
  out.city = copy_string(string_of(require_field(obj, "city")));
  array scores = array_of(require_field(obj, "scores"));
  out.score_count = scores.size();
  out.scores.resize(out.score_count);
  size_t slot = 0;
  for (element score : scores) out.scores[slot++] = double_of(score);
  bool present = false;
  element note = optional_field(obj, "note", present);
  out.has_note = present && !note.is_null();
  if (out.has_note) out.note = copy_string(string_of(note));
}

void decode_user(element raw, User &out) {
  object obj = object_of(raw);
  out.id = int_of(require_field(obj, "id"));
  out.name = copy_string(string_of(require_field(obj, "name")));
  out.active = bool_of(require_field(obj, "active"));
  array tags = array_of(require_field(obj, "tags"));
  out.tag_count = tags.size();
  out.tags.resize(out.tag_count);
  size_t slot = 0;
  for (element tag : tags) out.tags[slot++] = copy_string(string_of(tag));
  bool present = false;
  element profile = optional_field(obj, "profile", present);
  if (present && !profile.is_null()) {
    out.profile.emplace();
    decode_profile(profile, *out.profile);
  }
}

void decode_item(element raw, Item &out) {
  object obj = object_of(raw);
  out.sku = copy_string(string_of(require_field(obj, "sku")));
  out.quantity = int_of(require_field(obj, "quantity"));
  out.price = double_of(require_field(obj, "price"));
}

void decode_event(element raw, Event &out) {
  object obj = object_of(raw);
  std::string_view tag = string_of(require_field(obj, "tag"));
  if (tag == "view") {
    out.kind = 0;
    out.path = copy_string(string_of(require_field(obj, "path")));
    bool present = false;
    element duration = optional_field(obj, "duration", present);
    out.has_duration = present && !duration.is_null();
    if (out.has_duration) out.duration = int_of(duration);
    return;
  }
  if (tag == "purchase") {
    out.kind = 1;
    out.order_id = int_of(require_field(obj, "orderId"));
    array items = array_of(require_field(obj, "items"));
    out.item_count = items.size();
    out.items.resize(out.item_count);
    size_t slot = 0;
    for (element item : items) decode_item(item, out.items[slot++]);
    return;
  }
  throw decode_error{};
}

void decode_payload(element root, Payload &out) {
  object obj = object_of(root);
  out.version = int_of(require_field(obj, "version"));
  array users = array_of(require_field(obj, "users"));
  out.user_count = users.size();
  out.users.resize(out.user_count);
  size_t slot = 0;
  for (element user : users) decode_user(user, out.users[slot++]);
  array events = array_of(require_field(obj, "events"));
  out.event_count = events.size();
  out.events.resize(out.event_count);
  slot = 0;
  for (element event : events) decode_event(event, out.events[slot++]);
  bool present = false;
  element next = optional_field(obj, "next", present);
  out.has_next = present && !next.is_null();
  if (out.has_next) out.next = copy_string(string_of(next));
}

// --------------------------------------------------------------------------
// Canonical payload serialization: the normalized shape the oracle hashes,
// with object keys sorted.
// --------------------------------------------------------------------------

void append_score_list(std::string &out, const Profile &profile) {
  out.push_back('[');
  bool first = true;
  for (size_t i = 0; i < profile.score_count; i++) {
    double score = profile.scores[i];
    if (!first) out.push_back(',');
    first = false;
    append_number(out, score);
  }
  out.push_back(']');
}

void append_profile(std::string &out, const Profile &profile) {
  out += "{\"city\":";
  append_escaped(out, profile.city);
  out += ",\"note\":";
  if (profile.has_note) {
    append_escaped(out, profile.note);
  } else {
    out += "null";
  }
  out += ",\"scores\":";
  append_score_list(out, profile);
  out.push_back('}');
}

void append_user(std::string &out, const User &user) {
  out += "{\"active\":";
  out += user.active ? "true" : "false";
  out += ",\"id\":";
  out += std::to_string(user.id);
  out += ",\"name\":";
  append_escaped(out, user.name);
  out += ",\"profile\":";
  if (user.profile) {
    append_profile(out, *user.profile);
  } else {
    out += "null";
  }
  out += ",\"tags\":[";
  bool first = true;
  for (size_t i = 0; i < user.tag_count; i++) {
    if (!first) out.push_back(',');
    first = false;
    append_escaped(out, user.tags[i]);
  }
  out += "]}";
}

void append_item(std::string &out, const Item &item) {
  out += "{\"price\":";
  append_number(out, item.price);
  out += ",\"quantity\":";
  out += std::to_string(item.quantity);
  out += ",\"sku\":";
  append_escaped(out, item.sku);
  out.push_back('}');
}

void append_event(std::string &out, const Event &event) {
  if (event.kind == 0) {
    out += "{\"duration\":";
    if (event.has_duration) {
      out += std::to_string(event.duration);
    } else {
      out += "null";
    }
    out += ",\"path\":";
    append_escaped(out, event.path);
    out += ",\"tag\":\"view\"}";
    return;
  }
  out += "{\"items\":[";
  bool first = true;
  for (size_t i = 0; i < event.item_count; i++) {
    if (!first) out.push_back(',');
    first = false;
    append_item(out, event.items[i]);
  }
  out += "],\"orderId\":";
  out += std::to_string(event.order_id);
  out += ",\"tag\":\"purchase\"}";
}

std::string payload_fingerprint(const Payload &payload) {
  std::string text = "{\"value\":{\"events\":[";
  bool first = true;
  for (size_t i = 0; i < payload.event_count; i++) {
    if (!first) text.push_back(',');
    first = false;
    append_event(text, payload.events[i]);
  }
  text += "],\"next\":";
  if (payload.has_next) {
    append_escaped(text, payload.next);
  } else {
    text += "null";
  }
  text += ",\"users\":[";
  first = true;
  for (size_t i = 0; i < payload.user_count; i++) {
    if (!first) text.push_back(',');
    first = false;
    append_user(text, payload.users[i]);
  }
  text += "],\"version\":";
  text += std::to_string(payload.version);
  text += "}}";
  return sha256_hex(text);
}

// --------------------------------------------------------------------------
// Driver: corpus loading, the three phases and the report.
// --------------------------------------------------------------------------

struct Case {
  std::string name;
  std::string contents;
  bool benchmark = false;
};

double now_microseconds() {
  using clock = std::chrono::steady_clock;
  return std::chrono::duration<double, std::micro>(clock::now().time_since_epoch()).count();
}

std::string json_quoted(std::string_view text) {
  std::string out;
  append_escaped(out, text);
  return out;
}

bool read_text(const char *path, std::string &out) {
  FILE *handle = fopen(path, "rb");
  if (handle == nullptr) return false;
  char buffer[1 << 16];
  size_t read = 0;
  while ((read = fread(buffer, 1, sizeof(buffer), handle)) > 0) out.append(buffer, read);
  fclose(handle);
  return true;
}

std::vector<int> phase_order() {
  const char *setting = getenv("DIAG_PHASES");
  if (!setting) return {0, 1, 2};
  std::vector<int> result;
  std::string text(setting);
  size_t start = 0;
  do {
    size_t end = text.find(',', start);
    std::string name = text.substr(start, end == std::string::npos ? end : end - start);
    if (name == "parse") result.push_back(0);
    else if (name == "decode") result.push_back(1);
    else if (name == "combined") result.push_back(2);
    else throw decode_error{};
    if (end == std::string::npos) break;
    start = end + 1;
  } while (true);
  return result;
}

}  // namespace

int main() {
  const char *corpus_path = getenv("DIAG_CORPUS");
  if (corpus_path == nullptr) {
    fprintf(stderr, "DIAG_CORPUS is not set\n");
    return 2;
  }
  std::string corpus_text;
  if (!read_text(corpus_path, corpus_text)) {
    fprintf(stderr, "cannot read corpus\n");
    return 2;
  }

  simdjson::dom::parser corpus_parser;
  std::vector<Case> cases;
  try {
    array entries = corpus_parser.parse(corpus_text).get_array().value();
    for (element entry : entries) {
      object obj = object_of(entry);
      Case item;
      std::string_view name = string_of(require_field(obj, "name"));
      std::string_view contents = string_of(require_field(obj, "contents"));
      item.name.assign(name.data(), name.size());
      item.contents.assign(contents.data(), contents.size());
      bool present = false;
      element benchmark = optional_field(obj, "benchmark", present);
      item.benchmark = present && bool_of(benchmark);
      cases.push_back(std::move(item));
    }
  } catch (...) {
    fprintf(stderr, "invalid corpus\n");
    return 2;
  }

  std::vector<int> timed;
  for (size_t index = 0; index < cases.size(); index++) {
    if (cases[index].benchmark) timed.push_back(static_cast<int>(index));
  }
  if (timed.empty()) {
    fprintf(stderr, "no timed cases\n");
    return 2;
  }

  // Every case keeps one parsed document alive for the decode-only control.
  std::vector<std::unique_ptr<simdjson::dom::parser>> parsers;
  std::vector<element> documents;
  parsers.reserve(cases.size());
  documents.reserve(cases.size());
  try {
    for (const auto &item : cases) {
      parsers.push_back(std::make_unique<simdjson::dom::parser>());
      documents.push_back(parsers.back()->parse(item.contents).value());
    }
  } catch (...) {
    fprintf(stderr, "invalid corpus entry\n");
    return 2;
  }

  // Fingerprints are computed once, outside every timed interval.
  std::vector<std::string> json_fingerprints;
  std::vector<std::string> fingerprints;
  std::vector<bool> decoded;
  json_fingerprints.reserve(cases.size());
  fingerprints.reserve(cases.size());
  decoded.reserve(cases.size());
  for (size_t index = 0; index < cases.size(); index++) {
    json_fingerprints.push_back(fingerprint_of(documents[index]));
    Payload payload;
    bool ok = true;
    try {
      decode_payload(documents[index], payload);
    } catch (...) {
      ok = false;
    }
    decoded.push_back(ok);
    fingerprints.push_back(ok ? payload_fingerprint(payload) : std::string());
  }

  // Two warm-up passes then five samples; the minimum is reported, as in the
  // Go and JavaScript drivers.
  const char *phase_names[3] = {"parse", "decode", "combined"};
  std::string phases = "{";
  auto order = phase_order();
  for (size_t phase_slot = 0; phase_slot < order.size(); phase_slot++) {
    int phase = order[phase_slot];
    double best = 0;
    bool have_best = false;
    bool first_entry = true;
    std::string samples = "[";
    for (int pass = 0; pass < 7; pass++) {
      std::vector<Payload> payloads;
      std::vector<std::unique_ptr<simdjson::dom::parser>> parsed_results;
      std::vector<element> parsed_values;
      payloads.reserve(timed.size());
      parsed_results.reserve(timed.size());
      parsed_values.reserve(timed.size());
      double begin = now_microseconds();
      if (phase == 0) {
        for (int index : timed) {
          auto parser = std::make_unique<simdjson::dom::parser>();
          parsed_values.push_back(parser->parse(cases[index].contents).value());
          parsed_results.push_back(std::move(parser));
        }
      } else if (phase == 1) {
        for (int index : timed) {
          payloads.emplace_back();
          decode_payload(documents[index], payloads.back());
        }
      } else {
        for (int index : timed) {
          simdjson::dom::parser parser;
          element root = parser.parse(cases[index].contents).value();
          payloads.emplace_back();
          decode_payload(root, payloads.back());
        }
      }
      double elapsed = now_microseconds() - begin;
      if (phase == 0) {
        for (size_t slot = 0; slot < timed.size(); slot++) {
          int index = timed[slot];
          element root = parsed_values[slot];
          if (fingerprint_of(root) != json_fingerprints[index]) {
            fprintf(stderr, "unstable parse output\n");
            return 3;
          }
        }
      } else {
        for (size_t slot = 0; slot < timed.size(); slot++) {
          int index = timed[slot];
          if (!decoded[index] || payload_fingerprint(payloads[slot]) != fingerprints[index]) {
            fprintf(stderr, "unstable %s output\n", phase_names[phase]);
            return 3;
          }
        }
      }
      double release_begin = now_microseconds();
      payloads.clear();
      parsed_results.clear();
      double release = now_microseconds() - release_begin;
      if (pass >= 2) {
        if (!have_best || elapsed < best) {
          best = elapsed;
          have_best = true;
        }
        if (!first_entry) samples += ",";
        first_entry = false;
        samples += "{\"time_us\":" + std::to_string(elapsed) +
                   ",\"release_us\":" + std::to_string(release) +
                   ",\"lifecycle_us\":" + std::to_string(elapsed + release) + "}";
      }
    }
    samples += "]";
    phases += json_quoted(phase_names[phase]) + ":{\"samples\":" + samples +
              ",\"time_us\":" + std::to_string(best) + "}";
    if (phase_slot + 1 < order.size()) phases += ",";
  }
  phases += "}";

  std::string report = "{\"backend\":\"c\",\"modules\":" + std::to_string(cases.size()) +
                       ",\"timed_cases\":" + std::to_string(timed.size()) + ",\"names\":[";
  for (size_t index = 0; index < cases.size(); index++) {
    if (index) report += ",";
    report += json_quoted(cases[index].name);
  }
  report += "],\"fingerprints\":[";
  for (size_t index = 0; index < cases.size(); index++) {
    if (index) report += ",";
    report += decoded[index] ? json_quoted(fingerprints[index]) : std::string("null");
  }
  report += "],\"json_fingerprints\":[";
  for (size_t index = 0; index < cases.size(); index++) {
    if (index) report += ",";
    report += json_quoted(json_fingerprints[index]);
  }
  report += "],\"phases\":" + phases + "}";

  printf("%s\n", report.c_str());
  return 0;
}
