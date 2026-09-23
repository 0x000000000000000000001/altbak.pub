// JSON to Typed AST reference (C/C++), same corpus and protocol as the Go and
// JavaScript diagnostics driven by bin/benchmark/json-diagnostic.py.
//
// Scope of the comparison:
//   * parsing uses simdjson (system include), as in the JSON decoding
//     reference;
//   * the typed module is built into one monotonic arena that is reset per
//     pass, instead of per-value allocation;
//   * the canonical fingerprint is written from the decoded module, exactly as
//     the PureScript fingerprint does, and validated by the caller against the
//     frozen oracle for all twelve modules;
//   * annotation spans are empty (`<internal>`, 0/0), as the PureScript
//     decoder deliberately does not decode them;
//   * the pure usage-validation pass of the PureScript decoder is not
//     reproduced: it returns Unit and cannot change the fingerprint. The
//     reference therefore measures parsing plus typed decoding.
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
#include <string>
#include <string_view>
#include <vector>

namespace {

using simdjson::dom::array;
using simdjson::dom::element;
using simdjson::dom::object;

struct decode_error {};

// --------------------------------------------------------------------------
// Arena and strings.
// --------------------------------------------------------------------------

class Arena {
 public:
  explicit Arena(size_t block_size) : block_size_(block_size) {}

  void reset() {
    for (Block &block : blocks_) block.used = 0;
  }

  void *alloc(size_t size) {
    const size_t alignment = alignof(std::max_align_t);
    size = (size + alignment - 1) & ~(alignment - 1);
    for (Block &block : blocks_) {
      if (block.used + size <= block.size) {
        void *pointer = block.data.get() + block.used;
        block.used += size;
        return pointer;
      }
    }
    Block block;
    block.size = size > block_size_ ? size : block_size_;
    block.data = std::make_unique<char[]>(block.size);
    block.used = size;
    void *pointer = block.data.get();
    blocks_.push_back(std::move(block));
    return pointer;
  }

  template <class T>
  T *alloc_array(size_t count) {
    if (count == 0) return nullptr;
    T *items = static_cast<T *>(alloc(sizeof(T) * count));
    for (size_t index = 0; index < count; index++) new (&items[index]) T();
    return items;
  }

  template <class T, class... Args>
  T *make(Args &&...args) {
    void *pointer = alloc(sizeof(T));
    return new (pointer) T(std::forward<Args>(args)...);
  }

 private:
  struct Block {
    std::unique_ptr<char[]> data;
    size_t size = 0;
    size_t used = 0;
  };
  std::vector<Block> blocks_;
  size_t block_size_;
};

struct Str {
  const char *data = "";
  size_t size = 0;
};

Str copy_string(Arena &arena, std::string_view text) {
  char *data = static_cast<char *>(arena.alloc(text.size() == 0 ? 1 : text.size()));
  memcpy(data, text.data(), text.size());
  return Str{data, text.size()};
}

std::string_view view_of(const Str &text) { return std::string_view(text.data, text.size); }

// --------------------------------------------------------------------------
// simdjson accessors (the corpus is valid; failures abort the case).
// --------------------------------------------------------------------------

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

element optional_field(object obj, const char *key) {
  bool present = false;
  element value = optional_field(obj, key, present);
  return present ? value : element();
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
// Typed AST structures.
// --------------------------------------------------------------------------

struct SourcePos {
  int64_t line = 0;
  int64_t column = 0;
};

struct SourceSpan {
  Str path;
  SourcePos start;
  SourcePos end;
};

struct Type;

struct RowField {
  Str label;
  Type *type = nullptr;
};

struct Constraint {
  Str *parts = nullptr;
  size_t part_count = 0;
  Type **args = nullptr;
  size_t arg_count = 0;
};

enum class TypeKind {
  Int, Number, String, Char, Boolean, Unit, Any,
  TypeLevelString, Array, TypeVar, Adt, TypeApp, Func, Row, Record, ForAll, ConstrainedType
};

struct Type {
  TypeKind kind = TypeKind::Any;
  Str text;                                        // TypeLevelString, TypeVar
  Str *parts = nullptr; size_t part_count = 0;     // Adt
  Type **args = nullptr; size_t arg_count = 0;     // Adt, TypeApp, Func
  Type *constructor = nullptr;                     // TypeApp
  Type *element = nullptr;                         // Array
  Type *row = nullptr;                             // Record
  RowField *fields = nullptr; size_t field_count = 0; Type *tail = nullptr;  // Row
  Type *body = nullptr;                            // ForAll, ConstrainedType, Func result
  Str *vars = nullptr; size_t var_count = 0;       // ForAll
  Constraint *constraints = nullptr; size_t constraint_count = 0;
};

struct BindingUsage {
  int64_t binding_id = 0;
  bool has_max_uses = false;
  int64_t max_uses = 0;
  bool has_escaping = false;
  bool escaping = false;
};

struct VariableUse {
  int64_t binding_id = 0;
  bool has_last_local = false;
  bool last_local = false;
};

struct SourceUsage {
  bool has_binding = false;
  BindingUsage binding;
  bool has_variable = false;
  VariableUse variable;
};

struct Meta {
  int kind = 0;  // 0 IsConstructor, 1 IsNewtype, 2 IsTypeClassConstructor, 3 IsForeign, 4 IsWhere, 5 IsSyntheticApp
  int constructor_type = 0;  // 0 ProductType, 1 SumType
  Str *idents = nullptr;
  size_t ident_count = 0;
};

struct Ann {
  SourceSpan span;
  bool has_meta = false;
  Meta meta;
  bool has_type = false;
  Type *type = nullptr;
  bool has_usage = false;
  SourceUsage usage;
};

struct Qualified {
  bool has_module = false;
  Str module;  // dotted
  Str name;
};

struct Expr;
struct Binder;

struct LiteralItem {
  Expr *expr = nullptr;
  Binder *binder = nullptr;
};

struct LiteralProp {
  Str key;
  LiteralItem value;
};

struct Literal {
  int kind = 0;  // 0 Int, 1 Number, 2 String, 3 Char, 4 Boolean, 5 Array, 6 Record
  int64_t int_value = 0;
  double number_value = 0;
  Str text;
  bool bool_value = false;
  LiteralItem *items = nullptr;
  size_t item_count = 0;
  LiteralProp *props = nullptr;
  size_t prop_count = 0;
};

struct Prop {
  Str key;
  Expr *expr = nullptr;
};

struct Guard {
  Expr *condition = nullptr;
  Expr *value = nullptr;
};

struct CaseAlternative {
  Binder **binders = nullptr;
  size_t binder_count = 0;
  bool guarded = false;
  Expr *expr = nullptr;
  Guard *guards = nullptr;
  size_t guard_count = 0;
};

struct Bind;

struct Expr {
  int kind = 0;  // 0 Var, 1 Literal, 2 Constructor, 3 Accessor, 4 Update,
                 // 5 Abs, 6 App, 7 Case, 8 Let, 9 TypeApp
  Ann ann;
  Qualified value;          // Var
  Literal *literal = nullptr;
  Str type_name;            // Constructor
  Str constructor_name;     // Constructor
  Str *fields = nullptr;    // Constructor field names
  size_t field_count = 0;
  Expr *inner = nullptr;    // Accessor/Update/TypeApp expression
  Str key;                  // Accessor field name
  Prop *props = nullptr;    // Update
  size_t prop_count = 0;
  Str name;                 // Abs argument
  Expr *body = nullptr;     // Abs body, Let body
  Expr *fn = nullptr;       // App
  Expr *arg = nullptr;      // App
  Expr **values = nullptr;  // Case expressions
  size_t value_count = 0;
  CaseAlternative *alternatives = nullptr;
  size_t alternative_count = 0;
  Bind *binds = nullptr;    // Let
  size_t bind_count = 0;
  Type *type_arg = nullptr;      // TypeApp
};

struct Binding {
  Ann ann;
  Str name;
  Expr *expr = nullptr;
};

struct Bind {
  int kind = 0;  // 0 NonRec, 1 Rec
  Binding single;
  Binding *group = nullptr;
  size_t group_count = 0;
};

struct Binder {
  int kind = 0;  // 0 Null, 1 Var, 2 Named, 3 Literal, 4 Constructor
  Ann ann;
  Str name;                 // Var, Named
  Binder *inner = nullptr;  // Named
  Literal *literal = nullptr;
  Qualified proper;         // Constructor type name
  Qualified constructor;    // Constructor ident
  Binder **fields = nullptr;
  size_t field_count = 0;
};

struct Import {
  Ann ann;
  Str module;  // dotted
};

struct ReExport {
  Str module;  // dotted
  Str ident;
};

struct DataConstructor {
  Str name;
  Type **fields = nullptr;
  size_t field_count = 0;
};

struct DataDecl {
  Str name;
  Str *vars = nullptr;
  size_t var_count = 0;
  DataConstructor *constructors = nullptr;
  size_t constructor_count = 0;
};

struct ClassMethod {
  Str name;
  Type *type = nullptr;
};

struct ClassDecl {
  Str name;
  Str *vars = nullptr;
  size_t var_count = 0;
  Constraint *superclasses = nullptr;
  size_t superclass_count = 0;
  ClassMethod *methods = nullptr;
  size_t method_count = 0;
};

struct Comment {
  int kind = 0;  // 0 LineComment, 1 BlockComment
  Str text;
};

struct ForeignEntry {
  Str ident;
  bool has_type = false;
  Type *type = nullptr;
};

struct Module {
  Str name;  // dotted
  Str path;
  SourceSpan span;
  Import *imports = nullptr;
  size_t import_count = 0;
  Str *exports = nullptr;
  size_t export_count = 0;
  ReExport *re_exports = nullptr;
  size_t re_export_count = 0;
  DataDecl *data_decls = nullptr;
  size_t data_decl_count = 0;
  ClassDecl *class_decls = nullptr;
  size_t class_decl_count = 0;
  Bind *decls = nullptr;
  size_t decl_count = 0;
  ForeignEntry *foreign = nullptr;
  size_t foreign_count = 0;
  Comment *comments = nullptr;
  size_t comment_count = 0;
};

// --------------------------------------------------------------------------
// Type table: raw entries, then memoized resolution (cycles abort).
// --------------------------------------------------------------------------

struct RawRowField {
  Str label;
  int64_t type = -1;
};

struct RawConstraint {
  Str *parts = nullptr;
  size_t part_count = 0;
  int64_t *args = nullptr;
  size_t arg_count = 0;
};

struct RawType {
  TypeKind kind = TypeKind::Any;
  Str text;
  Str *parts = nullptr; size_t part_count = 0;
  int64_t *args = nullptr; size_t arg_count = 0;
  int64_t constructor = -1;
  int64_t element = -1;
  int64_t row = -1;
  RawRowField *fields = nullptr; size_t field_count = 0;
  int64_t tail = -1;
  bool has_tail = false;
  int64_t body = -1;
  bool has_body = false;
  Str *vars = nullptr; size_t var_count = 0;
  RawConstraint *constraints = nullptr; size_t constraint_count = 0;
};

class TypeTable {
 public:
  TypeTable(Arena &arena, RawType *entries, size_t count)
      : arena_(arena), entries_(entries), count_(count),
        memo_(arena.alloc_array<Type *>(count == 0 ? 1 : count)),
        state_(arena.alloc_array<uint8_t>(count == 0 ? 1 : count)) {}

  size_t size() const { return count_; }

  Type *resolve(int64_t id) {
    if (id < 0 || static_cast<size_t>(id) >= count_) return any_type();
    if (state_[id] == 2) return memo_[id];
    if (state_[id] == 1) throw decode_error{};  // cycle
    state_[id] = 1;
    Type *resolved = build(entries_[id]);
    memo_[id] = resolved;
    state_[id] = 2;
    return resolved;
  }

  Type *resolve_checked(int64_t id) {
    if (id < 0 || static_cast<size_t>(id) >= count_) throw decode_error{};
    return resolve(id);
  }

  Type *any_type() {
    if (any_ == nullptr) {
      any_ = arena_.make<Type>();
      any_->kind = TypeKind::Any;
    }
    return any_;
  }

 private:
  Type *build(const RawType &raw) {
    Type *type = arena_.make<Type>();
    type->kind = raw.kind;
    switch (raw.kind) {
      case TypeKind::TypeLevelString:
      case TypeKind::TypeVar:
        type->text = raw.text;
        break;
      case TypeKind::Adt:
        type->parts = raw.parts;
        type->part_count = raw.part_count;
        type->args = resolve_args(raw.args, raw.arg_count);
        type->arg_count = raw.arg_count;
        break;
      case TypeKind::TypeApp:
        type->constructor = resolve(raw.constructor);
        type->args = resolve_args(raw.args, raw.arg_count);
        type->arg_count = raw.arg_count;
        break;
      case TypeKind::Func:
        type->args = resolve_args(raw.args, raw.arg_count);
        type->arg_count = raw.arg_count;
        type->body = resolve(raw.body);
        break;
      case TypeKind::Array:
        type->element = resolve(raw.element);
        break;
      case TypeKind::Record:
        type->row = resolve(raw.row);
        break;
      case TypeKind::Row: {
        type->field_count = raw.field_count;
        type->fields = arena_.alloc_array<RowField>(raw.field_count);
        for (size_t index = 0; index < raw.field_count; index++) {
          type->fields[index].label = raw.fields[index].label;
          type->fields[index].type = resolve(raw.fields[index].type);
        }
        type->tail = raw.has_tail ? resolve(raw.tail) : nullptr;
        break;
      }
      case TypeKind::ForAll:
        type->vars = raw.vars;
        type->var_count = raw.var_count;
        type->body = resolve(raw.body);
        break;
      case TypeKind::ConstrainedType: {
        type->constraint_count = raw.constraint_count;
        type->constraints = arena_.alloc_array<Constraint>(raw.constraint_count);
        for (size_t index = 0; index < raw.constraint_count; index++) {
          Constraint &constraint = type->constraints[index];
          constraint.parts = raw.constraints[index].parts;
          constraint.part_count = raw.constraints[index].part_count;
          constraint.arg_count = raw.constraints[index].arg_count;
          constraint.args = resolve_args(raw.constraints[index].args, raw.constraints[index].arg_count);
        }
        type->body = resolve(raw.has_body ? raw.body : -1);
        break;
      }
      default:
        break;
    }
    return type;
  }

  Type **resolve_args(const int64_t *ids, size_t count) {
    Type **types = arena_.alloc_array<Type *>(count);
    for (size_t index = 0; index < count; index++) {
      types[index] = (ids == nullptr) ? any_type() : resolve(ids[index]);
    }
    return types;
  }

  Arena &arena_;
  RawType *entries_;
  size_t count_;
  Type **memo_;
  uint8_t *state_;
  Type *any_ = nullptr;
};

// --------------------------------------------------------------------------
// Small helpers.
// --------------------------------------------------------------------------

std::vector<std::string_view> string_list(element raw) {
  std::vector<std::string_view> out;
  for (element item : array_of(raw)) out.push_back(string_of(item));
  return out;
}

std::vector<int64_t> id_list(element raw) {
  std::vector<int64_t> out;
  for (element item : array_of(raw)) out.push_back(int_of(item));
  return out;
}

Str join_parts(Arena &arena, const std::vector<std::string_view> &parts) {
  std::string joined;
  for (size_t index = 0; index < parts.size(); index++) {
    if (index) joined.push_back('.');
    joined += parts[index];
  }
  return copy_string(arena, joined);
}

Str decode_module_name(Arena &arena, element raw) {
  return join_parts(arena, string_list(raw));
}

SourcePos decode_source_pos(element raw) {
  std::vector<int64_t> values = id_list(raw);
  if (values.size() != 2) throw decode_error{};
  return SourcePos{values[0], values[1]};
}

SourceSpan decode_source_span(Arena &arena, std::string_view path, element raw) {
  object obj = object_of(raw);
  return SourceSpan{copy_string(arena, path), decode_source_pos(require_field(obj, "start")),
                    decode_source_pos(require_field(obj, "end"))};
}

SourceSpan empty_span(Arena &arena) {
  SourceSpan span;
  span.path = copy_string(arena, "<internal>");
  return span;
}

// decodeStringLiteral: a string, or an array of code points.
Str decode_string_literal(Arena &arena, element raw) {
  if (raw.is_string()) return copy_string(arena, string_of(raw));
  std::string text;
  for (int64_t code_point : id_list(raw)) {
    if (code_point < 0 || code_point > 0x10ffff) throw decode_error{};
    if (code_point < 0x80) {
      text.push_back(static_cast<char>(code_point));
    } else if (code_point < 0x800) {
      text.push_back(static_cast<char>(0xc0 | (code_point >> 6)));
      text.push_back(static_cast<char>(0x80 | (code_point & 0x3f)));
    } else if (code_point < 0x10000) {
      text.push_back(static_cast<char>(0xe0 | (code_point >> 12)));
      text.push_back(static_cast<char>(0x80 | ((code_point >> 6) & 0x3f)));
      text.push_back(static_cast<char>(0x80 | (code_point & 0x3f)));
    } else {
      text.push_back(static_cast<char>(0xf0 | (code_point >> 18)));
      text.push_back(static_cast<char>(0x80 | ((code_point >> 12) & 0x3f)));
      text.push_back(static_cast<char>(0x80 | ((code_point >> 6) & 0x3f)));
      text.push_back(static_cast<char>(0x80 | (code_point & 0x3f)));
    }
  }
  return copy_string(arena, text);
}

// --------------------------------------------------------------------------
// Decoders.
// --------------------------------------------------------------------------

void decode_expr(Arena &arena, TypeTable &table, element raw, Expr &out);
void decode_binder(Arena &arena, TypeTable &table, element raw, Binder &out);

Meta decode_meta(Arena &arena, element raw) {
  object obj = object_of(raw);
  std::string_view kind = string_of(require_field(obj, "metaType"));
  Meta meta;
  if (kind == "IsConstructor") {
    meta.kind = 0;
    std::string_view constructor = string_of(require_field(obj, "constructorType"));
    meta.constructor_type = constructor == "SumType" ? 1 : 0;
    array idents = array_of(require_field(obj, "identifiers"));
    meta.ident_count = idents.size();
    meta.idents = arena.alloc_array<Str>(meta.ident_count);
    size_t slot = 0;
    for (element ident : idents) meta.idents[slot++] = copy_string(arena, string_of(ident));
    return meta;
  }
  if (kind == "IsNewtype") meta.kind = 1;
  else if (kind == "IsTypeClassConstructor") meta.kind = 2;
  else if (kind == "IsForeign") meta.kind = 3;
  else if (kind == "IsWhere") meta.kind = 4;
  else if (kind == "IsSyntheticApp") meta.kind = 5;
  else throw decode_error{};
  return meta;
}

SourceUsage decode_source_usage(object obj) {
  SourceUsage usage;
  bool present = false;
  element binding = optional_field(obj, "bindingUsage", present);
  if (present && !binding.is_null()) {
    object binding_obj = object_of(binding);
    usage.has_binding = true;
    usage.binding.binding_id = int_of(require_field(binding_obj, "bindingId"));
    bool has_max = false;
    element max_uses = optional_field(binding_obj, "maxUses", has_max);
    if (has_max && !max_uses.is_null()) {
      usage.binding.has_max_uses = true;
      usage.binding.max_uses = int_of(max_uses);
    }
    bool has_escaping = false;
    element escaping = optional_field(binding_obj, "hasEscapingUseContext", has_escaping);
    if (has_escaping && !escaping.is_null()) {
      usage.binding.has_escaping = true;
      usage.binding.escaping = bool_of(escaping);
    }
  }
  bool has_variable = false;
  element variable = optional_field(obj, "variableUse", has_variable);
  if (has_variable && !variable.is_null()) {
    object variable_obj = object_of(variable);
    usage.has_variable = true;
    usage.variable.binding_id = int_of(require_field(variable_obj, "bindingId"));
    bool has_last = false;
    element last_local = optional_field(variable_obj, "lastLocalUse", has_last);
    if (has_last && !last_local.is_null()) {
      usage.variable.has_last_local = true;
      usage.variable.last_local = bool_of(last_local);
    }
  }
  return usage;
}

Ann decode_ann(Arena &arena, TypeTable &table, element raw, bool with_usage) {
  object obj = object_of(raw);
  Ann ann;
  ann.span = empty_span(arena);
  bool has_meta = false;
  element meta = optional_field(obj, "meta", has_meta);
  if (has_meta && !meta.is_null()) {
    ann.has_meta = true;
    ann.meta = decode_meta(arena, meta);
  }
  bool has_type = false;
  element type = optional_field(obj, "type", has_type);
  if (has_type && !type.is_null()) {
    int64_t id = int_of(type);
    if (id >= 0 && static_cast<size_t>(id) < table.size()) {
      ann.has_type = true;
      ann.type = table.resolve(id);
    }
  }
  if (with_usage) {
    SourceUsage usage = decode_source_usage(obj);
    if (usage.has_binding || usage.has_variable) {
      ann.has_usage = true;
      ann.usage = usage;
    }
  }
  return ann;
}

Qualified decode_qualified(Arena &arena, element raw) {
  object obj = object_of(raw);
  Qualified qualified;
  bool has_module = false;
  element module = optional_field(obj, "moduleName", has_module);
  if (has_module && !module.is_null()) {
    qualified.has_module = true;
    qualified.module = decode_module_name(arena, module);
  }
  qualified.name = copy_string(arena, string_of(require_field(obj, "identifier")));
  return qualified;
}

Literal decode_literal(Arena &arena, TypeTable &table, element raw, bool binder_context) {
  object obj = object_of(raw);
  std::string_view kind = string_of(require_field(obj, "literalType"));
  Literal literal;
  if (kind == "IntLiteral") {
    literal.kind = 0;
    literal.int_value = int_of(require_field(obj, "value"));
    return literal;
  }
  if (kind == "NumberLiteral") {
    literal.kind = 1;
    literal.number_value = double_of(require_field(obj, "value"));
    return literal;
  }
  if (kind == "StringLiteral") {
    literal.kind = 2;
    literal.text = decode_string_literal(arena, require_field(obj, "value"));
    return literal;
  }
  if (kind == "CharLiteral") {
    literal.kind = 3;
    literal.text = decode_string_literal(arena, require_field(obj, "value"));
    return literal;
  }
  if (kind == "BooleanLiteral") {
    literal.kind = 4;
    literal.bool_value = bool_of(require_field(obj, "value"));
    return literal;
  }
  if (kind == "ArrayLiteral") {
    literal.kind = 5;
    array items = array_of(require_field(obj, "value"));
    literal.item_count = items.size();
    literal.items = arena.alloc_array<LiteralItem>(literal.item_count);
    size_t slot = 0;
    for (element item : items) {
      if (binder_context) {
        Binder *binder = arena.make<Binder>();
        decode_binder(arena, table, item, *binder);
        literal.items[slot].binder = binder;
      } else {
        Expr *expr = arena.make<Expr>();
        decode_expr(arena, table, item, *expr);
        literal.items[slot].expr = expr;
      }
      slot++;
    }
    return literal;
  }
  if (kind == "ObjectLiteral") {
    literal.kind = 6;
    array pairs = array_of(require_field(obj, "value"));
    literal.prop_count = pairs.size();
    literal.props = arena.alloc_array<LiteralProp>(literal.prop_count);
    size_t slot = 0;
    for (element pair : pairs) {
      std::vector<element> elements;
      for (element item : array_of(pair)) elements.push_back(item);
      if (elements.size() != 2) throw decode_error{};
      literal.props[slot].key = decode_string_literal(arena, elements[0]);
      if (binder_context) {
        Binder *binder = arena.make<Binder>();
        decode_binder(arena, table, elements[1], *binder);
        literal.props[slot].value.binder = binder;
      } else {
        Expr *expr = arena.make<Expr>();
        decode_expr(arena, table, elements[1], *expr);
        literal.props[slot].value.expr = expr;
      }
      slot++;
    }
    return literal;
  }
  throw decode_error{};
}


Bind decode_bind(Arena &arena, TypeTable &table, element raw);

Binding decode_binding(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  Binding binding;
  binding.ann = decode_ann(arena, table, require_field(obj, "annotation"), true);
  binding.name = copy_string(arena, string_of(require_field(obj, "identifier")));
  Expr *expr = arena.make<Expr>();
  decode_expr(arena, table, require_field(obj, "expression"), *expr);
  binding.expr = expr;
  return binding;
}

Bind decode_bind(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  std::string_view kind = string_of(require_field(obj, "bindType"));
  Bind bind;
  if (kind == "NonRec") {
    bind.kind = 0;
    bind.single = decode_binding(arena, table, raw);
    return bind;
  }
  if (kind == "Rec") {
    bind.kind = 1;
    array group = array_of(require_field(obj, "binds"));
    bind.group_count = group.size();
    bind.group = arena.alloc_array<Binding>(bind.group_count);
    size_t slot = 0;
    for (element item : group) bind.group[slot++] = decode_binding(arena, table, item);
    return bind;
  }
  throw decode_error{};
}

Guard decode_guard(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  Guard guard;
  Expr *condition = arena.make<Expr>();
  decode_expr(arena, table, require_field(obj, "guard"), *condition);
  Expr *value = arena.make<Expr>();
  decode_expr(arena, table, require_field(obj, "expression"), *value);
  guard.condition = condition;
  guard.value = value;
  return guard;
}

CaseAlternative decode_case_alternative(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  CaseAlternative alternative;
  array binders = array_of(require_field(obj, "binders"));
  alternative.binder_count = binders.size();
  alternative.binders = arena.alloc_array<Binder *>(alternative.binder_count);
  size_t slot = 0;
  for (element item : binders) {
    Binder *binder = arena.make<Binder>();
    decode_binder(arena, table, item, *binder);
    alternative.binders[slot++] = binder;
  }
  alternative.guarded = bool_of(require_field(obj, "isGuarded"));
  if (alternative.guarded) {
    array guards = array_of(require_field(obj, "expressions"));
    alternative.guard_count = guards.size();
    alternative.guards = arena.alloc_array<Guard>(alternative.guard_count);
    size_t guard_slot = 0;
    for (element item : guards) alternative.guards[guard_slot++] = decode_guard(arena, table, item);
  } else {
    Expr *expr = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "expression"), *expr);
    alternative.expr = expr;
  }
  return alternative;
}

void decode_expr(Arena &arena, TypeTable &table, element raw, Expr &out) {
  object obj = object_of(raw);
  out.ann = decode_ann(arena, table, require_field(obj, "annotation"), true);
  std::string_view kind = string_of(require_field(obj, "type"));
  if (kind == "Var") {
    out.kind = 0;
    out.value = decode_qualified(arena, require_field(obj, "value"));
    return;
  }
  if (kind == "Literal") {
    out.kind = 1;
    out.literal = arena.make<Literal>();
    *out.literal = decode_literal(arena, table, require_field(obj, "value"), false);
    return;
  }
  if (kind == "Constructor") {
    out.kind = 2;
    out.type_name = copy_string(arena, string_of(require_field(obj, "typeName")));
    bool has_name = false;
    element name = optional_field(obj, "name", has_name);
    element constructor = (has_name && !name.is_null()) ? name : require_field(obj, "constructorName");
    out.constructor_name = copy_string(arena, string_of(constructor));
    bool has_fields = false;
    element fields = optional_field(obj, "fields", has_fields);
    if (!has_fields || fields.is_null()) fields = require_field(obj, "fieldNames");
    array names = array_of(fields);
    out.field_count = names.size();
    out.fields = arena.alloc_array<Str>(out.field_count);
    size_t slot = 0;
    for (element item : names) out.fields[slot++] = decode_string_literal(arena, item);
    return;
  }
  if (kind == "Accessor") {
    out.kind = 3;
    Expr *inner = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "expression"), *inner);
    out.inner = inner;
    out.key = decode_string_literal(arena, require_field(obj, "fieldName"));
    return;
  }
  if (kind == "ObjectUpdate") {
    out.kind = 4;
    Expr *inner = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "expression"), *inner);
    out.inner = inner;
    array pairs = array_of(require_field(obj, "updates"));
    out.prop_count = pairs.size();
    out.props = arena.alloc_array<Prop>(out.prop_count);
    size_t slot = 0;
    for (element pair : pairs) {
      std::vector<element> elements;
      for (element item : array_of(pair)) elements.push_back(item);
      if (elements.size() != 2) throw decode_error{};
      out.props[slot].key = decode_string_literal(arena, elements[0]);
      Expr *value = arena.make<Expr>();
      decode_expr(arena, table, elements[1], *value);
      out.props[slot].expr = value;
      slot++;
    }
    return;
  }
  if (kind == "Abs") {
    out.kind = 5;
    out.name = copy_string(arena, string_of(require_field(obj, "argument")));
    Expr *body = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "body"), *body);
    out.body = body;
    return;
  }
  if (kind == "App") {
    out.kind = 6;
    Expr *fn = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "abstraction"), *fn);
    Expr *arg = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "argument"), *arg);
    out.fn = fn;
    out.arg = arg;
    return;
  }
  if (kind == "Case") {
    out.kind = 7;
    array values = array_of(require_field(obj, "caseExpressions"));
    out.value_count = values.size();
    out.values = arena.alloc_array<Expr *>(out.value_count);
    size_t slot = 0;
    for (element item : values) {
      Expr *value = arena.make<Expr>();
      decode_expr(arena, table, item, *value);
      out.values[slot++] = value;
    }
    array alternatives = array_of(require_field(obj, "caseAlternatives"));
    out.alternative_count = alternatives.size();
    out.alternatives = arena.alloc_array<CaseAlternative>(out.alternative_count);
    size_t alternative_slot = 0;
    for (element item : alternatives) {
      out.alternatives[alternative_slot++] = decode_case_alternative(arena, table, item);
    }
    return;
  }
  if (kind == "Let") {
    out.kind = 8;
    array binds = array_of(require_field(obj, "binds"));
    out.bind_count = binds.size();
    out.binds = arena.alloc_array<Bind>(out.bind_count);
    size_t slot = 0;
    for (element item : binds) out.binds[slot++] = decode_bind(arena, table, item);
    Expr *body = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "expression"), *body);
    out.body = body;
    return;
  }
  if (kind == "TypeApp") {
    out.kind = 9;
    Expr *inner = arena.make<Expr>();
    decode_expr(arena, table, require_field(obj, "expression"), *inner);
    out.inner = inner;
    out.type_arg = table.resolve_checked(int_of(require_field(obj, "typeArgument")));
    return;
  }
  throw decode_error{};
}

void decode_binder(Arena &arena, TypeTable &table, element raw, Binder &out) {
  object obj = object_of(raw);
  out.ann = decode_ann(arena, table, require_field(obj, "annotation"), true);
  std::string_view kind = string_of(require_field(obj, "binderType"));
  if (kind == "NullBinder") {
    out.kind = 0;
    return;
  }
  if (kind == "VarBinder") {
    out.kind = 1;
    out.name = copy_string(arena, string_of(require_field(obj, "identifier")));
    return;
  }
  if (kind == "NamedBinder") {
    out.kind = 2;
    out.name = copy_string(arena, string_of(require_field(obj, "identifier")));
    Binder *inner = arena.make<Binder>();
    decode_binder(arena, table, require_field(obj, "binder"), *inner);
    out.inner = inner;
    return;
  }
  if (kind == "LiteralBinder") {
    out.kind = 3;
    out.literal = arena.make<Literal>();
    *out.literal = decode_literal(arena, table, require_field(obj, "literal"), true);
    return;
  }
  if (kind == "ConstructorBinder") {
    out.kind = 4;
    out.proper = decode_qualified(arena, require_field(obj, "typeName"));
    bool has_name = false;
    element name = optional_field(obj, "name", has_name);
    element constructor = (has_name && !name.is_null()) ? name : require_field(obj, "constructorName");
    out.constructor = decode_qualified(arena, constructor);
    array binders = array_of(require_field(obj, "binders"));
    out.field_count = binders.size();
    out.fields = arena.alloc_array<Binder *>(out.field_count);
    size_t slot = 0;
    for (element item : binders) {
      Binder *binder = arena.make<Binder>();
      decode_binder(arena, table, item, *binder);
      out.fields[slot++] = binder;
    }
    return;
  }
  throw decode_error{};
}

Import decode_import(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  Import entry;
  entry.ann = decode_ann(arena, table, require_field(obj, "annotation"), true);
  entry.module = decode_module_name(arena, require_field(obj, "moduleName"));
  return entry;
}

DataConstructor decode_data_constructor(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  DataConstructor constructor;
  bool has_name = false;
  element name = optional_field(obj, "name", has_name);
  element resolved = (has_name && !name.is_null()) ? name : require_field(obj, "constructorName");
  constructor.name = copy_string(arena, string_of(resolved));
  bool has_fields = false;
  element fields = optional_field(obj, "fields", has_fields);
  if (!has_fields || fields.is_null()) fields = require_field(obj, "fieldTypes");
  std::vector<int64_t> ids = id_list(fields);
  constructor.field_count = ids.size();
  constructor.fields = arena.alloc_array<Type *>(constructor.field_count);
  for (size_t index = 0; index < ids.size(); index++) {
    constructor.fields[index] = table.resolve_checked(ids[index]);
  }
  return constructor;
}

DataDecl decode_data_decl(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  DataDecl decl;
  bool has_name = false;
  element name = optional_field(obj, "name", has_name);
  element resolved = (has_name && !name.is_null()) ? name : require_field(obj, "typeName");
  decl.name = copy_string(arena, string_of(resolved));
  bool has_vars = false;
  element vars = optional_field(obj, "vars", has_vars);
  if (!has_vars || vars.is_null()) vars = optional_field(obj, "typeVars");
  if (!vars.is_null()) {
    std::vector<std::string_view> list = string_list(vars);
    decl.var_count = list.size();
    decl.vars = arena.alloc_array<Str>(decl.var_count);
    size_t slot = 0;
    for (std::string_view item : list) decl.vars[slot++] = copy_string(arena, item);
  }
  array constructors = array_of(require_field(obj, "constructors"));
  decl.constructor_count = constructors.size();
  decl.constructors = arena.alloc_array<DataConstructor>(decl.constructor_count);
  size_t slot = 0;
  for (element item : constructors) {
    decl.constructors[slot++] = decode_data_constructor(arena, table, item);
  }
  return decl;
}

Constraint decode_constraint(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  Constraint constraint;
  std::vector<std::string_view> parts = string_list(require_field(obj, "fqn"));
  constraint.part_count = parts.size();
  constraint.parts = arena.alloc_array<Str>(constraint.part_count);
  size_t slot = 0;
  for (std::string_view part : parts) constraint.parts[slot++] = copy_string(arena, part);
  std::vector<int64_t> ids = id_list(require_field(obj, "args"));
  constraint.arg_count = ids.size();
  constraint.args = arena.alloc_array<Type *>(constraint.arg_count);
  for (size_t index = 0; index < ids.size(); index++) {
    constraint.args[index] = table.resolve_checked(ids[index]);
  }
  return constraint;
}

ClassDecl decode_class_decl(Arena &arena, TypeTable &table, element raw) {
  object obj = object_of(raw);
  ClassDecl decl;
  decl.name = copy_string(arena, string_of(require_field(obj, "name")));
  bool has_vars = false;
  element vars = optional_field(obj, "vars", has_vars);
  if (has_vars && !vars.is_null()) {
    std::vector<std::string_view> list = string_list(vars);
    decl.var_count = list.size();
    decl.vars = arena.alloc_array<Str>(decl.var_count);
    size_t slot = 0;
    for (std::string_view item : list) decl.vars[slot++] = copy_string(arena, item);
  }
  array superclasses = array_of(require_field(obj, "superclasses"));
  decl.superclass_count = superclasses.size();
  decl.superclasses = arena.alloc_array<Constraint>(decl.superclass_count);
  size_t superclass_slot = 0;
  for (element item : superclasses) {
    decl.superclasses[superclass_slot++] = decode_constraint(arena, table, item);
  }
  array methods = array_of(require_field(obj, "methods"));
  decl.method_count = methods.size();
  decl.methods = arena.alloc_array<ClassMethod>(decl.method_count);
  size_t method_slot = 0;
  for (element item : methods) {
    object method = object_of(item);
    decl.methods[method_slot].name = copy_string(arena, string_of(require_field(method, "name")));
    decl.methods[method_slot].type = table.resolve_checked(int_of(require_field(method, "type")));
    method_slot++;
  }
  return decl;
}

Comment decode_comment(Arena &arena, element raw) {
  object obj = object_of(raw);
  Comment comment;
  bool has_line = false;
  element line = optional_field(obj, "LineComment", has_line);
  if (has_line && !line.is_null()) {
    comment.kind = 0;
    comment.text = copy_string(arena, string_of(line));
    return comment;
  }
  element block = require_field(obj, "BlockComment");
  comment.kind = 1;
  comment.text = copy_string(arena, string_of(block));
  return comment;
}

// --------------------------------------------------------------------------
// Type table parsing.
// --------------------------------------------------------------------------

RawType parse_raw_type(Arena &arena, element raw) {
  RawType entry;
  if (raw.is_string()) {
    std::string_view name = string_of(raw);
    if (name == "Int") entry.kind = TypeKind::Int;
    else if (name == "Number") entry.kind = TypeKind::Number;
    else if (name == "String") entry.kind = TypeKind::String;
    else if (name == "Char") entry.kind = TypeKind::Char;
    else if (name == "Boolean") entry.kind = TypeKind::Boolean;
    else if (name == "Unit") entry.kind = TypeKind::Unit;
    else if (name == "Any") entry.kind = TypeKind::Any;
    else throw decode_error{};
    return entry;
  }
  object obj = object_of(raw);
  std::string_view kind = string_of(require_field(obj, "type"));
  if (kind == "Adt") {
    entry.kind = TypeKind::Adt;
    std::vector<std::string_view> parts = string_list(require_field(obj, "fqn"));
    entry.part_count = parts.size();
    entry.parts = arena.alloc_array<Str>(entry.part_count);
    size_t slot = 0;
    for (std::string_view part : parts) entry.parts[slot++] = copy_string(arena, part);
    std::vector<int64_t> ids = id_list(require_field(obj, "args"));
    entry.arg_count = ids.size();
    entry.args = arena.alloc_array<int64_t>(entry.arg_count);
    for (size_t index = 0; index < ids.size(); index++) entry.args[index] = ids[index];
    return entry;
  }
  if (kind == "TypeApp") {
    entry.kind = TypeKind::TypeApp;
    entry.constructor = int_of(require_field(obj, "constructor"));
    std::vector<int64_t> ids = id_list(require_field(obj, "args"));
    entry.arg_count = ids.size();
    entry.args = arena.alloc_array<int64_t>(entry.arg_count);
    for (size_t index = 0; index < ids.size(); index++) entry.args[index] = ids[index];
    return entry;
  }
  if (kind == "Func") {
    entry.kind = TypeKind::Func;
    std::vector<int64_t> ids = id_list(require_field(obj, "args"));
    entry.arg_count = ids.size();
    entry.args = arena.alloc_array<int64_t>(entry.arg_count);
    for (size_t index = 0; index < ids.size(); index++) entry.args[index] = ids[index];
    entry.body = int_of(require_field(obj, "ret"));
    entry.has_body = true;
    return entry;
  }
  if (kind == "Array") {
    entry.kind = TypeKind::Array;
    entry.element = int_of(require_field(obj, "element"));
    return entry;
  }
  if (kind == "TypeVar") {
    entry.kind = TypeKind::TypeVar;
    entry.text = copy_string(arena, string_of(require_field(obj, "name")));
    return entry;
  }
  if (kind == "Record") {
    entry.kind = TypeKind::Record;
    entry.row = int_of(require_field(obj, "row"));
    return entry;
  }
  if (kind == "Row") {
    entry.kind = TypeKind::Row;
    array fields = array_of(require_field(obj, "fields"));
    entry.field_count = fields.size();
    entry.fields = arena.alloc_array<RawRowField>(entry.field_count);
    size_t slot = 0;
    for (element item : fields) {
      object field = object_of(item);
      entry.fields[slot].label = copy_string(arena, string_of(require_field(field, "label")));
      entry.fields[slot].type = int_of(require_field(field, "type"));
      slot++;
    }
    bool has_tail = false;
    element tail = optional_field(obj, "tail", has_tail);
    if (has_tail && !tail.is_null()) {
      entry.has_tail = true;
      entry.tail = int_of(tail);
    }
    return entry;
  }
  if (kind == "ForAll") {
    entry.kind = TypeKind::ForAll;
    std::vector<std::string_view> vars = string_list(require_field(obj, "vars"));
    entry.var_count = vars.size();
    entry.vars = arena.alloc_array<Str>(entry.var_count);
    size_t slot = 0;
    for (std::string_view item : vars) entry.vars[slot++] = copy_string(arena, item);
    entry.body = int_of(require_field(obj, "body"));
    entry.has_body = true;
    return entry;
  }
  if (kind == "ConstrainedType") {
    entry.kind = TypeKind::ConstrainedType;
    array constraints = array_of(require_field(obj, "constraints"));
    entry.constraint_count = constraints.size();
    entry.constraints = arena.alloc_array<RawConstraint>(entry.constraint_count);
    size_t constraint_slot = 0;
    for (element item : constraints) {
      object constraint = object_of(item);
      RawConstraint &target = entry.constraints[constraint_slot++];
      std::vector<std::string_view> parts = string_list(require_field(constraint, "fqn"));
      target.part_count = parts.size();
      target.parts = arena.alloc_array<Str>(target.part_count);
      size_t part_slot = 0;
      for (std::string_view part : parts) target.parts[part_slot++] = copy_string(arena, part);
      std::vector<int64_t> ids = id_list(require_field(constraint, "args"));
      target.arg_count = ids.size();
      target.args = arena.alloc_array<int64_t>(target.arg_count);
      for (size_t index = 0; index < ids.size(); index++) target.args[index] = ids[index];
    }
    entry.body = int_of(require_field(obj, "body"));
    entry.has_body = true;
    return entry;
  }
  if (kind == "TypeLevelString") {
    entry.kind = TypeKind::TypeLevelString;
    entry.text = copy_string(arena, string_of(require_field(obj, "value")));
    return entry;
  }
  if (kind == "Int") entry.kind = TypeKind::Int;
  else if (kind == "Number") entry.kind = TypeKind::Number;
  else if (kind == "String") entry.kind = TypeKind::String;
  else if (kind == "Char") entry.kind = TypeKind::Char;
  else if (kind == "Boolean") entry.kind = TypeKind::Boolean;
  else if (kind == "Unit") entry.kind = TypeKind::Unit;
  else if (kind == "Any") entry.kind = TypeKind::Any;
  else throw decode_error{};
  return entry;
}

// --------------------------------------------------------------------------
// Canonical fingerprint writers (exact mirror of Test.JsonTypedAst.moduleJson).
// --------------------------------------------------------------------------

void open_tag(std::string &out, const char *name) {
  out += "[\"";
  out += name;
  out.push_back('"');
}

void close_tag(std::string &out) { out.push_back(']'); }

void write_str(std::string &out, std::string_view text) { append_escaped(out, text); }
void write_str(std::string &out, const Str &text) { append_escaped(out, view_of(text)); }
void write_int(std::string &out, int64_t value) { out += std::to_string(value); }
void write_bool(std::string &out, bool value) { out += value ? "true" : "false"; }

void write_optional_int(std::string &out, bool has, int64_t value) {
  if (has) {
    open_tag(out, "Just");
    out.push_back(',');
    write_int(out, value);
    close_tag(out);
    return;
  }
  open_tag(out, "Nothing");
  close_tag(out);
}

void write_optional_bool(std::string &out, bool has, bool value) {
  if (has) {
    open_tag(out, "Just");
    out.push_back(',');
    write_bool(out, value);
    close_tag(out);
    return;
  }
  open_tag(out, "Nothing");
  close_tag(out);
}

void write_source_pos(std::string &out, const SourcePos &pos) {
  open_tag(out, "SourcePos");
  out.push_back(',');
  write_int(out, pos.line);
  out.push_back(',');
  write_int(out, pos.column);
  close_tag(out);
}

void write_source_span(std::string &out, const SourceSpan &span) {
  open_tag(out, "SourceSpan");
  out.push_back(',');
  write_str(out, span.path);
  out.push_back(',');
  write_source_pos(out, span.start);
  out.push_back(',');
  write_source_pos(out, span.end);
  close_tag(out);
}

void write_expr_type(std::string &out, const Type *type);

void write_meta(std::string &out, const Meta &meta) {
  switch (meta.kind) {
    case 0: {
      open_tag(out, "IsConstructor");
      out.push_back(',');
      open_tag(out, meta.constructor_type == 1 ? "SumType" : "ProductType");
      close_tag(out);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < meta.ident_count; index++) {
        if (index) out.push_back(',');
        write_str(out, meta.idents[index]);
      }
      out.push_back(']');
      close_tag(out);
      break;
    }
    case 1: open_tag(out, "IsNewtype"); close_tag(out); break;
    case 2: open_tag(out, "IsTypeClassConstructor"); close_tag(out); break;
    case 3: open_tag(out, "IsForeign"); close_tag(out); break;
    case 4: open_tag(out, "IsWhere"); close_tag(out); break;
    default: open_tag(out, "IsSyntheticApp"); close_tag(out); break;
  }
}

void write_optional_meta(std::string &out, const Ann &ann) {
  if (ann.has_meta) {
    open_tag(out, "Just");
    out.push_back(',');
    write_meta(out, ann.meta);
    close_tag(out);
    return;
  }
  open_tag(out, "Nothing");
  close_tag(out);
}

void write_optional_type(std::string &out, bool has, const Type *type) {
  if (has && type != nullptr) {
    open_tag(out, "Just");
    out.push_back(',');
    write_expr_type(out, type);
    close_tag(out);
    return;
  }
  open_tag(out, "Nothing");
  close_tag(out);
}

void write_source_binding_id(std::string &out, const Str &module, int64_t binding_id) {
  open_tag(out, "SourceBindingId");
  out.push_back(',');
  write_str(out, module);
  out.push_back(',');
  write_int(out, binding_id);
  close_tag(out);
}

void write_binding_usage(std::string &out, const Str &module, const BindingUsage &usage) {
  open_tag(out, "BindingUsage");
  out.push_back(',');
  write_source_binding_id(out, module, usage.binding_id);
  out.push_back(',');
  write_optional_int(out, usage.has_max_uses, usage.max_uses);
  out.push_back(',');
  write_optional_bool(out, usage.has_escaping, usage.escaping);
  close_tag(out);
}

void write_variable_use(std::string &out, const Str &module, const VariableUse &usage) {
  open_tag(out, "VariableUse");
  out.push_back(',');
  write_source_binding_id(out, module, usage.binding_id);
  out.push_back(',');
  write_optional_bool(out, usage.has_last_local, usage.last_local);
  close_tag(out);
}

void write_optional_source_usage(std::string &out, const Ann &ann, const Str &module) {
  if (!ann.has_usage) {
    open_tag(out, "Nothing");
    close_tag(out);
    return;
  }
  open_tag(out, "Just");
  out.push_back(',');
  open_tag(out, "SourceUsage");
  out.push_back(',');
  if (ann.usage.has_binding) {
    open_tag(out, "Just");
    out.push_back(',');
    write_binding_usage(out, module, ann.usage.binding);
    close_tag(out);
  } else {
    open_tag(out, "Nothing");
    close_tag(out);
  }
  out.push_back(',');
  if (ann.usage.has_variable) {
    open_tag(out, "Just");
    out.push_back(',');
    write_variable_use(out, module, ann.usage.variable);
    close_tag(out);
  } else {
    open_tag(out, "Nothing");
    close_tag(out);
  }
  close_tag(out);
  close_tag(out);
}

void write_ann(std::string &out, const Ann &ann, const Str &module) {
  open_tag(out, "Ann");
  out.push_back(',');
  write_source_span(out, ann.span);
  out.push_back(',');
  write_optional_meta(out, ann);
  out.push_back(',');
  write_optional_type(out, ann.has_type, ann.type);
  out.push_back(',');
  write_optional_source_usage(out, ann, module);
  close_tag(out);
}

void write_qualified(std::string &out, const Qualified &qualified) {
  open_tag(out, "Qualified");
  out.push_back(',');
  if (qualified.has_module) {
    open_tag(out, "Just");
    out.push_back(',');
    write_str(out, qualified.module);
    close_tag(out);
  } else {
    open_tag(out, "Nothing");
    close_tag(out);
  }
  out.push_back(',');
  write_str(out, qualified.name);
  close_tag(out);
}

void write_constraint(std::string &out, const Constraint &constraint) {
  out.push_back('[');
  out.push_back('[');
  for (size_t index = 0; index < constraint.part_count; index++) {
    if (index) out.push_back(',');
    write_str(out, constraint.parts[index]);
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < constraint.arg_count; index++) {
    if (index) out.push_back(',');
    write_expr_type(out, constraint.args[index]);
  }
  out.push_back(']');
  out.push_back(']');
}

void write_expr_type(std::string &out, const Type *type) {
  if (type == nullptr) {
    open_tag(out, "Any");
    close_tag(out);
    return;
  }
  switch (type->kind) {
    case TypeKind::Int: open_tag(out, "Int"); close_tag(out); break;
    case TypeKind::Number: open_tag(out, "Number"); close_tag(out); break;
    case TypeKind::String: open_tag(out, "String"); close_tag(out); break;
    case TypeKind::Char: open_tag(out, "Char"); close_tag(out); break;
    case TypeKind::Boolean: open_tag(out, "Boolean"); close_tag(out); break;
    case TypeKind::Unit: open_tag(out, "Unit"); close_tag(out); break;
    case TypeKind::Any: open_tag(out, "Any"); close_tag(out); break;
    case TypeKind::TypeLevelString:
      open_tag(out, "TypeLevelString");
      out.push_back(',');
      write_str(out, type->text);
      close_tag(out);
      break;
    case TypeKind::Array:
      open_tag(out, "Array");
      out.push_back(',');
      write_expr_type(out, type->element);
      close_tag(out);
      break;
    case TypeKind::TypeVar:
      open_tag(out, "TypeVar");
      out.push_back(',');
      write_str(out, type->text);
      close_tag(out);
      break;
    case TypeKind::Adt: {
      open_tag(out, "ADT");
      out.push_back(',');
      std::string joined;
      for (size_t index = 0; index < type->part_count; index++) {
        if (index) joined.push_back('.');
        joined += view_of(type->parts[index]);
      }
      write_str(out, joined);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < type->part_count; index++) {
        if (index) out.push_back(',');
        write_str(out, type->parts[index]);
      }
      out.push_back(']');
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < type->arg_count; index++) {
        if (index) out.push_back(',');
        write_expr_type(out, type->args[index]);
      }
      out.push_back(']');
      close_tag(out);
      break;
    }
    case TypeKind::TypeApp:
      open_tag(out, "TypeApp");
      out.push_back(',');
      write_expr_type(out, type->constructor);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < type->arg_count; index++) {
        if (index) out.push_back(',');
        write_expr_type(out, type->args[index]);
      }
      out.push_back(']');
      close_tag(out);
      break;
    case TypeKind::Func:
      open_tag(out, "Func");
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < type->arg_count; index++) {
        if (index) out.push_back(',');
        write_expr_type(out, type->args[index]);
      }
      out.push_back(']');
      out.push_back(',');
      write_expr_type(out, type->body);
      close_tag(out);
      break;
    case TypeKind::Row:
      open_tag(out, "Row");
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < type->field_count; index++) {
        if (index) out.push_back(',');
        out.push_back('[');
        write_str(out, type->fields[index].label);
        out.push_back(',');
        write_expr_type(out, type->fields[index].type);
        out.push_back(']');
      }
      out.push_back(']');
      out.push_back(',');
      write_optional_type(out, type->tail != nullptr, type->tail);
      close_tag(out);
      break;
    case TypeKind::Record:
      open_tag(out, "Record");
      out.push_back(',');
      write_expr_type(out, type->row);
      close_tag(out);
      break;
    case TypeKind::ForAll:
      open_tag(out, "ForAll");
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < type->var_count; index++) {
        if (index) out.push_back(',');
        write_str(out, type->vars[index]);
      }
      out.push_back(']');
      out.push_back(',');
      write_expr_type(out, type->body);
      close_tag(out);
      break;
    case TypeKind::ConstrainedType:
      open_tag(out, "ConstrainedType");
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < type->constraint_count; index++) {
        if (index) out.push_back(',');
        write_constraint(out, type->constraints[index]);
      }
      out.push_back(']');
      out.push_back(',');
      write_expr_type(out, type->body);
      close_tag(out);
      break;
  }
}

void write_literal(std::string &out, const Literal &literal, const Str &module, bool binder_context);
void write_expr(std::string &out, const Expr &expr, const Str &module);
void write_binder(std::string &out, const Binder &binder, const Str &module);

void write_literal_item(std::string &out, const LiteralItem &item, const Str &module, bool binder_context) {
  if (binder_context) {
    write_binder(out, *item.binder, module);
  } else {
    write_expr(out, *item.expr, module);
  }
}

void write_literal(std::string &out, const Literal &literal, const Str &module, bool binder_context) {
  switch (literal.kind) {
    case 0: open_tag(out, "LitInt"); out.push_back(','); write_int(out, literal.int_value); close_tag(out); break;
    case 1: open_tag(out, "LitNumber"); out.push_back(','); append_number(out, literal.number_value); close_tag(out); break;
    case 2: open_tag(out, "LitString"); out.push_back(','); write_str(out, literal.text); close_tag(out); break;
    case 3: open_tag(out, "LitChar"); out.push_back(','); write_str(out, literal.text); close_tag(out); break;
    case 4: open_tag(out, "LitBoolean"); out.push_back(','); write_bool(out, literal.bool_value); close_tag(out); break;
    case 5:
      open_tag(out, "LitArray");
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < literal.item_count; index++) {
        if (index) out.push_back(',');
        write_literal_item(out, literal.items[index], module, binder_context);
      }
      out.push_back(']');
      close_tag(out);
      break;
    default:
      open_tag(out, "LitRecord");
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < literal.prop_count; index++) {
        if (index) out.push_back(',');
        open_tag(out, "Prop");
        out.push_back(',');
        write_str(out, literal.props[index].key);
        out.push_back(',');
        write_literal_item(out, literal.props[index].value, module, binder_context);
        close_tag(out);
      }
      out.push_back(']');
      close_tag(out);
      break;
  }
}

void write_prop(std::string &out, const Prop &prop, const Str &module) {
  open_tag(out, "Prop");
  out.push_back(',');
  write_str(out, prop.key);
  out.push_back(',');
  write_expr(out, *prop.expr, module);
  close_tag(out);
}

void write_case_alternative(std::string &out, const CaseAlternative &alternative, const Str &module) {
  open_tag(out, "CaseAlternative");
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < alternative.binder_count; index++) {
    if (index) out.push_back(',');
    write_binder(out, *alternative.binders[index], module);
  }
  out.push_back(']');
  out.push_back(',');
  if (alternative.guarded) {
    open_tag(out, "Guarded");
    out.push_back(',');
    out.push_back('[');
    for (size_t index = 0; index < alternative.guard_count; index++) {
      if (index) out.push_back(',');
      open_tag(out, "Guard");
      out.push_back(',');
      write_expr(out, *alternative.guards[index].condition, module);
      out.push_back(',');
      write_expr(out, *alternative.guards[index].value, module);
      close_tag(out);
    }
    out.push_back(']');
    close_tag(out);
  } else {
    open_tag(out, "Unconditional");
    out.push_back(',');
    write_expr(out, *alternative.expr, module);
    close_tag(out);
  }
  close_tag(out);
}

void write_expr(std::string &out, const Expr &expr, const Str &module) {
  switch (expr.kind) {
    case 0:
      open_tag(out, "ExprVar");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_qualified(out, expr.value);
      close_tag(out);
      break;
    case 1:
      open_tag(out, "ExprLit");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_literal(out, *expr.literal, module, false);
      close_tag(out);
      break;
    case 2:
      open_tag(out, "ExprConstructor");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_str(out, expr.type_name);
      out.push_back(',');
      write_str(out, expr.constructor_name);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < expr.field_count; index++) {
        if (index) out.push_back(',');
        write_str(out, expr.fields[index]);
      }
      out.push_back(']');
      close_tag(out);
      break;
    case 3:
      open_tag(out, "ExprAccessor");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_expr(out, *expr.inner, module);
      out.push_back(',');
      write_str(out, expr.key);
      close_tag(out);
      break;
    case 4:
      open_tag(out, "ExprUpdate");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_expr(out, *expr.inner, module);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < expr.prop_count; index++) {
        if (index) out.push_back(',');
        write_prop(out, expr.props[index], module);
      }
      out.push_back(']');
      close_tag(out);
      break;
    case 5:
      open_tag(out, "ExprAbs");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_str(out, expr.name);
      out.push_back(',');
      write_expr(out, *expr.body, module);
      close_tag(out);
      break;
    case 6:
      open_tag(out, "ExprApp");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_expr(out, *expr.fn, module);
      out.push_back(',');
      write_expr(out, *expr.arg, module);
      close_tag(out);
      break;
    case 7:
      open_tag(out, "ExprCase");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < expr.value_count; index++) {
        if (index) out.push_back(',');
        write_expr(out, *expr.values[index], module);
      }
      out.push_back(']');
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < expr.alternative_count; index++) {
        if (index) out.push_back(',');
        write_case_alternative(out, expr.alternatives[index], module);
      }
      out.push_back(']');
      close_tag(out);
      break;
    case 8:
      open_tag(out, "ExprLet");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < expr.bind_count; index++) {
        if (index) out.push_back(',');
        const Bind &bind = expr.binds[index];
        if (bind.kind == 0) {
          open_tag(out, "NonRec");
          out.push_back(',');
          open_tag(out, "Binding");
          out.push_back(',');
          write_ann(out, bind.single.ann, module);
          out.push_back(',');
          write_str(out, bind.single.name);
          out.push_back(',');
          write_expr(out, *bind.single.expr, module);
          close_tag(out);
          close_tag(out);
        } else {
          open_tag(out, "Rec");
          out.push_back(',');
          out.push_back('[');
          for (size_t slot = 0; slot < bind.group_count; slot++) {
            if (slot) out.push_back(',');
            open_tag(out, "Binding");
            out.push_back(',');
            write_ann(out, bind.group[slot].ann, module);
            out.push_back(',');
            write_str(out, bind.group[slot].name);
            out.push_back(',');
            write_expr(out, *bind.group[slot].expr, module);
            close_tag(out);
          }
          out.push_back(']');
          close_tag(out);
        }
      }
      out.push_back(']');
      out.push_back(',');
      write_expr(out, *expr.body, module);
      close_tag(out);
      break;
    default:
      open_tag(out, "ExprTypeApp");
      out.push_back(',');
      write_ann(out, expr.ann, module);
      out.push_back(',');
      write_expr(out, *expr.inner, module);
      out.push_back(',');
      write_expr_type(out, expr.type_arg);
      close_tag(out);
      break;
  }
}

void write_binder(std::string &out, const Binder &binder, const Str &module) {
  switch (binder.kind) {
    case 0:
      open_tag(out, "BinderNull");
      out.push_back(',');
      write_ann(out, binder.ann, module);
      close_tag(out);
      break;
    case 1:
      open_tag(out, "BinderVar");
      out.push_back(',');
      write_ann(out, binder.ann, module);
      out.push_back(',');
      write_str(out, binder.name);
      close_tag(out);
      break;
    case 2:
      open_tag(out, "BinderNamed");
      out.push_back(',');
      write_ann(out, binder.ann, module);
      out.push_back(',');
      write_str(out, binder.name);
      out.push_back(',');
      write_binder(out, *binder.inner, module);
      close_tag(out);
      break;
    case 3:
      open_tag(out, "BinderLit");
      out.push_back(',');
      write_ann(out, binder.ann, module);
      out.push_back(',');
      write_literal(out, *binder.literal, module, true);
      close_tag(out);
      break;
    default:
      open_tag(out, "BinderConstructor");
      out.push_back(',');
      write_ann(out, binder.ann, module);
      out.push_back(',');
      write_qualified(out, binder.proper);
      out.push_back(',');
      write_qualified(out, binder.constructor);
      out.push_back(',');
      out.push_back('[');
      for (size_t index = 0; index < binder.field_count; index++) {
        if (index) out.push_back(',');
        write_binder(out, *binder.fields[index], module);
      }
      out.push_back(']');
      close_tag(out);
      break;
  }
}

std::string module_fingerprint(const Module &module) {
  std::string out;
  open_tag(out, "Module");
  out.push_back(',');
  write_str(out, module.name);
  out.push_back(',');
  write_str(out, module.path);
  out.push_back(',');
  write_source_span(out, module.span);
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.import_count; index++) {
    if (index) out.push_back(',');
    open_tag(out, "Import");
    out.push_back(',');
    write_ann(out, module.imports[index].ann, module.name);
    out.push_back(',');
    write_str(out, module.imports[index].module);
    close_tag(out);
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.export_count; index++) {
    if (index) out.push_back(',');
    write_str(out, module.exports[index]);
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.re_export_count; index++) {
    if (index) out.push_back(',');
    open_tag(out, "ReExport");
    out.push_back(',');
    write_str(out, module.re_exports[index].module);
    out.push_back(',');
    write_str(out, module.re_exports[index].ident);
    close_tag(out);
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.data_decl_count; index++) {
    if (index) out.push_back(',');
    const DataDecl &decl = module.data_decls[index];
    open_tag(out, "DataDecl");
    out.push_back(',');
    write_str(out, decl.name);
    out.push_back(',');
    out.push_back('[');
    for (size_t slot = 0; slot < decl.var_count; slot++) {
      if (slot) out.push_back(',');
      write_str(out, decl.vars[slot]);
    }
    out.push_back(']');
    out.push_back(',');
    out.push_back('[');
    for (size_t slot = 0; slot < decl.constructor_count; slot++) {
      if (slot) out.push_back(',');
      open_tag(out, "DataConstructor");
      out.push_back(',');
      write_str(out, decl.constructors[slot].name);
      out.push_back(',');
      out.push_back('[');
      for (size_t field = 0; field < decl.constructors[slot].field_count; field++) {
        if (field) out.push_back(',');
        write_expr_type(out, decl.constructors[slot].fields[field]);
      }
      out.push_back(']');
      close_tag(out);
    }
    out.push_back(']');
    close_tag(out);
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.class_decl_count; index++) {
    if (index) out.push_back(',');
    const ClassDecl &decl = module.class_decls[index];
    open_tag(out, "ClassDecl");
    out.push_back(',');
    write_str(out, decl.name);
    out.push_back(',');
    out.push_back('[');
    for (size_t slot = 0; slot < decl.var_count; slot++) {
      if (slot) out.push_back(',');
      write_str(out, decl.vars[slot]);
    }
    out.push_back(']');
    out.push_back(',');
    out.push_back('[');
    for (size_t slot = 0; slot < decl.superclass_count; slot++) {
      if (slot) out.push_back(',');
      write_constraint(out, decl.superclasses[slot]);
    }
    out.push_back(']');
    out.push_back(',');
    out.push_back('[');
    for (size_t slot = 0; slot < decl.method_count; slot++) {
      if (slot) out.push_back(',');
      out.push_back('[');
      write_str(out, decl.methods[slot].name);
      out.push_back(',');
      write_expr_type(out, decl.methods[slot].type);
      out.push_back(']');
    }
    out.push_back(']');
    close_tag(out);
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.decl_count; index++) {
    if (index) out.push_back(',');
    const Bind &bind = module.decls[index];
    if (bind.kind == 0) {
      open_tag(out, "NonRec");
      out.push_back(',');
      open_tag(out, "Binding");
      out.push_back(',');
      write_ann(out, bind.single.ann, module.name);
      out.push_back(',');
      write_str(out, bind.single.name);
      out.push_back(',');
      write_expr(out, *bind.single.expr, module.name);
      close_tag(out);
      close_tag(out);
    } else {
      open_tag(out, "Rec");
      out.push_back(',');
      out.push_back('[');
      for (size_t slot = 0; slot < bind.group_count; slot++) {
        if (slot) out.push_back(',');
        open_tag(out, "Binding");
        out.push_back(',');
        write_ann(out, bind.group[slot].ann, module.name);
        out.push_back(',');
        write_str(out, bind.group[slot].name);
        out.push_back(',');
        write_expr(out, *bind.group[slot].expr, module.name);
        close_tag(out);
      }
      out.push_back(']');
      close_tag(out);
    }
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.foreign_count; index++) {
    if (index) out.push_back(',');
    out.push_back('[');
    write_str(out, module.foreign[index].ident);
    out.push_back(',');
    write_optional_type(out, module.foreign[index].has_type, module.foreign[index].type);
    out.push_back(']');
  }
  out.push_back(']');
  out.push_back(',');
  out.push_back('[');
  for (size_t index = 0; index < module.comment_count; index++) {
    if (index) out.push_back(',');
    open_tag(out, module.comments[index].kind == 0 ? "LineComment" : "BlockComment");
    out.push_back(',');
    write_str(out, module.comments[index].text);
    close_tag(out);
  }
  out.push_back(']');
  close_tag(out);
  return out;
}

// --------------------------------------------------------------------------
// Module decoder.
// --------------------------------------------------------------------------

Module decode_module(Arena &arena, element root, std::string_view debug_name) {
  (void)debug_name;
  object obj = object_of(root);
  Module module;
  module.name = decode_module_name(arena, require_field(obj, "moduleName"));
  module.path = copy_string(arena, string_of(require_field(obj, "modulePath")));
  module.span = decode_source_span(arena, view_of(module.path), require_field(obj, "sourceSpan"));

  // Type table first: annotations resolve against it.
  bool has_table = false;
  element raw_table = optional_field(obj, "typeTable", has_table);
  std::vector<RawType> raw_entries;
  if (has_table && !raw_table.is_null()) {
    array entries = array_of(raw_table);
    raw_entries.reserve(entries.size());
    for (element entry : entries) raw_entries.push_back(parse_raw_type(arena, entry));
  }
  RawType *raw_array = arena.alloc_array<RawType>(raw_entries.size());
  for (size_t index = 0; index < raw_entries.size(); index++) raw_array[index] = raw_entries[index];
  TypeTable table(arena, raw_array, raw_entries.size());

  // Imports.
  array imports = array_of(require_field(obj, "imports"));
  module.import_count = imports.size();
  module.imports = arena.alloc_array<Import>(module.import_count);
  for (size_t index = 0; index < imports.size(); index++) {
    module.imports[index] = decode_import(arena, table, imports.at(index));
  }

  // Exports.
  array exports = array_of(require_field(obj, "exports"));
  module.export_count = exports.size();
  module.exports = arena.alloc_array<Str>(module.export_count);
  {
    size_t slot = 0;
    for (element item : exports) module.exports[slot++] = copy_string(arena, string_of(item));
  }

  // Re-exports: object of module name -> identifiers, sorted by module name.
  {
    object re_exports = object_of(require_field(obj, "reExports"));
    std::vector<std::pair<std::string_view, element>> entries;
    for (auto field : re_exports) entries.emplace_back(field.key, field.value);
    std::sort(entries.begin(), entries.end(),
              [](const auto &left, const auto &right) { return left.first < right.first; });
    size_t total = 0;
    for (const auto &entry : entries) total += array_of(entry.second).size();
    module.re_exports = arena.alloc_array<ReExport>(total);
    module.re_export_count = total;
    size_t slot = 0;
    for (const auto &entry : entries) {
      Str module_name = copy_string(arena, entry.first);
      for (element ident : array_of(entry.second)) {
        module.re_exports[slot].module = module_name;
        module.re_exports[slot].ident = copy_string(arena, string_of(ident));
        slot++;
      }
    }
  }

  // Data declarations.
  {
    bool present = false;
    element raw = optional_field(obj, "dataDecls", present);
    if (present && !raw.is_null()) {
      array decls = array_of(raw);
      module.data_decl_count = decls.size();
      module.data_decls = arena.alloc_array<DataDecl>(module.data_decl_count);
      size_t slot = 0;
      for (element item : decls) module.data_decls[slot++] = decode_data_decl(arena, table, item);
    }
  }

  // Class declarations.
  {
    bool present = false;
    element raw = optional_field(obj, "classDecls", present);
    if (present && !raw.is_null()) {
      array decls = array_of(raw);
      module.class_decl_count = decls.size();
      module.class_decls = arena.alloc_array<ClassDecl>(module.class_decl_count);
      size_t slot = 0;
      for (element item : decls) module.class_decls[slot++] = decode_class_decl(arena, table, item);
    }
  }

  // Declarations.
  array decls = array_of(require_field(obj, "decls"));
  module.decl_count = decls.size();
  module.decls = arena.alloc_array<Bind>(module.decl_count);
  for (size_t index = 0; index < decls.size(); index++) {
    module.decls[index] = decode_bind(arena, table, decls.at(index));
  }

  // Foreign entries, with their annotation type when present. The fingerprint
  // prints the module's foreign map, so entries are ordered by identifier.
  {
    array foreign = array_of(require_field(obj, "foreign"));
    bool has_annotations = false;
    element raw_annotations = optional_field(obj, "foreignAnnotations", has_annotations);
    module.foreign_count = foreign.size();
    module.foreign = arena.alloc_array<ForeignEntry>(module.foreign_count);
    size_t slot = 0;
    for (element item : foreign) {
      std::string_view ident = string_of(item);
      module.foreign[slot].ident = copy_string(arena, ident);
      if (has_annotations && !raw_annotations.is_null()) {
        object annotations = object_of(raw_annotations);
        element annotation;
        if (annotations[ident].get(annotation) == simdjson::SUCCESS) {
          Ann ann = decode_ann(arena, table, annotation, false);
          module.foreign[slot].has_type = ann.has_type;
          module.foreign[slot].type = ann.type;
        }
      }
      slot++;
    }
    std::sort(module.foreign, module.foreign + module.foreign_count,
              [](const ForeignEntry &left, const ForeignEntry &right) {
                return view_of(left.ident) < view_of(right.ident);
              });
  }

  // Comments.
  array comments = array_of(require_field(obj, "comments"));
  module.comment_count = comments.size();
  module.comments = arena.alloc_array<Comment>(module.comment_count);
  size_t comment_slot = 0;
  for (element item : comments) module.comments[comment_slot++] = decode_comment(arena, item);

  return module;
}

// --------------------------------------------------------------------------
// Driver: corpus loading, passes and report.
// --------------------------------------------------------------------------

struct Case {
  std::string name;
  std::string contents;
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
      cases.push_back(std::move(item));
    }
  } catch (...) {
    fprintf(stderr, "invalid corpus\n");
    return 2;
  }
  if (cases.empty()) {
    fprintf(stderr, "no cases\n");
    return 2;
  }

  // One parser per case keeps every parsed document alive for the decode
  // phase; the parse and combined phases reuse a scratch parser.
  std::vector<std::unique_ptr<simdjson::dom::parser>> parsers;
  std::vector<element> documents;
  parsers.reserve(cases.size());
  documents.reserve(cases.size());
  simdjson::dom::parser scratch;
  try {
    for (const auto &item : cases) {
      parsers.push_back(std::make_unique<simdjson::dom::parser>());
      documents.push_back(parsers.back()->parse(item.contents).value());
    }
  } catch (...) {
    fprintf(stderr, "invalid corpus entry\n");
    return 2;
  }

  const bool print_fingerprints = getenv("DIAG_PRINT_FINGERPRINTS") != nullptr;

  // Fingerprints are computed once, outside every timed interval.
  std::vector<std::string> json_fingerprints;
  std::vector<std::string> fingerprints;
  json_fingerprints.reserve(cases.size());
  fingerprints.reserve(cases.size());
  for (size_t index = 0; index < cases.size(); index++) {
    json_fingerprints.push_back(fingerprint_of(documents[index]));
    Arena arena(32 << 20);
    try {
      Module module = decode_module(arena, documents[index], cases[index].name);
      std::string text = module_fingerprint(module);
      if (print_fingerprints) {
        fprintf(stderr, "=== %s ===\n%s\n", cases[index].name.c_str(), text.c_str());
      }
      fingerprints.push_back(sha256_hex(text));
    } catch (...) {
      fprintf(stderr, "decode failed: %s\n", cases[index].name.c_str());
      return 3;
    }
  }

  const char *phase_names[3] = {"parse", "decode", "combined"};
  std::string phases = "{";
  for (int phase = 0; phase < 3; phase++) {
    double best = 0;
    bool have_best = false;
    bool first_entry = true;
    std::string samples = "[";
    for (int pass = 0; pass < 7; pass++) {
      Arena arena(32 << 20);
      std::vector<Module> modules;
      modules.reserve(cases.size());
      double begin = now_microseconds();
      if (phase == 0) {
        for (const auto &item : cases) scratch.parse(item.contents).value();
      } else if (phase == 1) {
        for (auto &item : documents) modules.push_back(decode_module(arena, item, ""));
      } else {
        for (const auto &item : cases) {
          element root = scratch.parse(item.contents).value();
          modules.push_back(decode_module(arena, root, ""));
        }
      }
      double elapsed = now_microseconds() - begin;
      for (size_t index = 0; index < cases.size(); index++) {
        if (phase == 0) {
          element root = scratch.parse(cases[index].contents).value();
          if (fingerprint_of(root) != json_fingerprints[index]) {
            fprintf(stderr, "unstable parse output\n");
            return 3;
          }
        } else if (sha256_hex(module_fingerprint(modules[index])) != fingerprints[index]) {
          fprintf(stderr, "unstable %s output\n", phase_names[phase]);
          return 3;
        }
      }
      if (pass >= 2) {
        if (!have_best || elapsed < best) {
          best = elapsed;
          have_best = true;
        }
        if (!first_entry) samples += ",";
        first_entry = false;
        samples += "{\"time_us\":" + std::to_string(elapsed) + "}";
      }
    }
    samples += "]";
    phases += json_quoted(phase_names[phase]) + ":{\"samples\":" + samples +
              ",\"time_us\":" + std::to_string(best) + "}";
    if (phase < 2) phases += ",";
  }
  phases += "}";

  std::string report = "{\"backend\":\"c\",\"modules\":" + std::to_string(cases.size()) +
                       ",\"fingerprints\":[";
  for (size_t index = 0; index < cases.size(); index++) {
    if (index) report += ",";
    report += json_quoted(fingerprints[index]);
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
