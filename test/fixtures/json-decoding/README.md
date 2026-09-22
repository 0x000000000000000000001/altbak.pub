# General JSON decoding fixture

This corpus exercises ordinary Argonaut decoders independently of PBO and TAST.
The PureScript payload uses records, arrays, optional fields and an explicit
tagged `Event` decoder (`view` or `purchase`). Parsing inputs are all valid JSON.

Five successful cases are timed: empty arrays, flat record arrays, nested
records and variants, null/missing optionals, and Unicode/escaping/numbers.
Twelve small cases check correctness separately: ten expected decoding errors
and two successful optional-field cases. Error cases are excluded from timing
so fast early failures cannot make the successful workload appear faster.

`corpus.json` stores `{name, contents, benchmark}` entries. `expected.json`
contains fixed SHA-256 hashes for the complete raw JSON and complete normalized
typed result of every entry, including exact error messages and paths. A typed
success is `{value: payload}` with absent optionals normalized to JSON null;
a failure is `{error: printJsonDecodeError error}`. Object keys are sorted and
HTML characters and Unicode line separators escaped, matching the JS/Go
diagnostic drivers. Numeric inputs avoid negative zero and runtime-specific
exponent formatting; integral floats use the same hash as their integer values.

The oracle comes from explicit input/result constructors and manually specified
Argonaut error paths in `test/native/json_decoding_fixture.py`. It is not
generated from either backend's observed result. Array element errors retain
`Named "Array"` and their index; malformed array containers do not add `Named`.
Record error priority follows alphabetical labels, and the event decoder reads
its tag before the selected variant's fields.

Check the frozen corpus and oracle with:

```sh
python3 -B test/native/json_decoding_fixture.py
```

After an intentional contract change, regenerate and check with:

```sh
python3 -B test/native/json_decoding_fixture.py --write
```

This is a bounded application-shaped workload, not a claim that every JSON
library, malformed JSON parser error, or possible PureScript data type is covered.
