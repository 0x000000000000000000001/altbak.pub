// Differential checks against JavaScript's UTF-16 string operations. No timing.
import { mkdtempSync, writeFileSync, rmSync } from "node:fs";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { execFileSync } from "node:child_process";

const directory = dirname(fileURLToPath(import.meta.url));
const scratch = mkdtempSync(join(tmpdir(), "altbak-c-strings-"));
const checksum = text => Array.from(text).reduce((sum, cp) => (sum + cp.codePointAt(0)) | 0, 0);
const units = text => Array.from({ length: text.length }, (_, i) => text.charCodeAt(i));
const declaration = (name, text) => `static const uint16_t ${name}[] = {${units(text).join(",") || "0"}};`;
const cases = [];
const add = (text, pattern, count, strings) => cases.push({ text, pattern, count, strings });

add("a😀é".repeat(1000), "needle", 1500, Array(1000).fill("ab"));
add("a".repeat(10000) + "needle", "needle", 1500, []);
add("a".repeat(10000) + "needle", "absent", 0, []);
add(Array(1000).fill("abc").join(","), ",", 0, []);
for (const text of ["", "abc", "😀", "\ud800", "\udc00", "\ud800,\udc00", "\ud800\ud800\udc00", ",abc,", "aaaaa", "\u0100\u0200\u0300\u0100"])
  for (const pattern of ["", ",", "a", "aa", "😀", "\ud800", "\udc00", "\u0300\u0100"])
    for (const count of [-1, 0, 1, 2, 20])
      add(text, pattern, count, ["\ud800", "", "\udc00", text, "\ud800"]);

let state = 0x7235af1;
const random = bound => { state ^= state << 13; state ^= state >>> 17; state ^= state << 5; return (state >>> 0) % bound; };
const alphabet = [0, 97, 98, 99, 233, 0x0100, 0x0200, 0x0300, 0xd800, 0xdbff, 0xdc00, 0xdfff, 0xd83d, 0xde00];
const randomString = length => String.fromCharCode(...Array.from({ length }, () => alphabet[random(alphabet.length)]));
for (let i = 0; i < 500; ++i) {
  const text = randomString(random(32));
  const pattern = i % 2 ? randomString(random(7)) : text.slice(random(text.length + 1), random(text.length + 1));
  add(text, pattern, random(40) - 3, Array.from({ length: random(8) }, () => randomString(random(8))));
}

let checks = 0;
const source = [
  '#include "common.h"',
  'static void check(const char *name, Int actual, Int expected) { if (actual != expected) { fprintf(stderr, "%s: got %d expected %d\\n", name, actual, expected); exit(1); } }',
  "int main(void) {",
];
for (const [index, test] of cases.entries()) {
  source.push("{", declaration("text", test.text), declaration("pattern", test.pattern));
  test.strings.forEach((part, i) => source.push(declaration(`part${i}`, part)));
  source.push(`const CString parts[] = {${test.strings.map((part, i) => `{part${i},${part.length}}`).join(",") || "{NULL,0}"}};`);
  source.push(`LibraryInput input = {.text={text,${test.text.length}}, .pattern={pattern,${test.pattern.length}}, .strings=parts, .string_count=${test.strings.length}, .count=${test.count}};`);
  const expected = {
    string_cu_length: test.text.length,
    string_cp_length: Array.from(test.text).length,
    string_index_of: test.text.indexOf(test.pattern),
    string_cp_take: checksum(Array.from(test.text).slice(0, Math.max(0, test.count)).join("")),
    string_split: test.text.split(test.pattern).reduce((sum, part) => (sum + checksum(part)) | 0, test.text.split(test.pattern).length),
    string_concat: checksum(test.strings.join("")),
  };
  for (const [action, result] of Object.entries(expected)) {
    source.push(`check("${index}:${action}", ${action}(&input), ${result});`);
    ++checks;
  }
  source.push("}");
}
source.push('puts("String differential checks passed"); return 0; }');

try {
  const testFile = join(scratch, "test.c");
  writeFileSync(testFile, source.join("\n") + "\n");
  for (const [name, flags] of [
    ["optimized", ["-O3"]],
    ["sanitized", ["-O1", "-g", "-fsanitize=address,undefined", "-fno-omit-frame-pointer"]],
  ]) {
    const binary = join(scratch, name);
    execFileSync(process.env.CC || "clang", ["-std=c11", "-Wall", "-Wextra", "-Werror", "-pedantic", ...flags, "-I", directory, join(directory, "strings.c"), testFile, "-o", binary], { stdio: "inherit" });
    execFileSync(binary, [], { stdio: "inherit" });
  }
  console.log(JSON.stringify({ fixtures: cases.length, checks_per_build: checks, builds: ["optimized", "address-undefined-sanitized"] }));
} finally {
  rmSync(scratch, { recursive: true, force: true });
}
