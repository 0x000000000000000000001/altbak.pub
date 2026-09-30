use std::time::Instant;

fn driver_json_string(value: &str) -> String {
    Purs_Data_Argonaut_Core::purust_json_quote(value)
}

fn driver_json_number(value: f64) -> String {
    format!("{}", value)
}

fn driver_field(value: &crate::UnknownType, key: &str) -> crate::UnknownType {
    Purs_Data_Argonaut_Core::purust_json_object_get(value, key)
        .unwrap_or_else(|| panic!("Data.JsonDecoding: corpus entry has no {}", key))
}

// Mirrors the Go and JavaScript drivers: the corpus is read from DIAG_CORPUS,
// every module is parsed and decoded once to build the oracle, then each
// phase measures seven passes (two warm-ups, five samples) and validates the
// produced fingerprints outside the timed interval.
pub fn Test_JsonDecoding_drive(
    parse: purust_core::Func1<String, crate::UnknownType>,
    decode: purust_core::Func1<crate::UnknownType, crate::UnknownType>,
    decode_text: purust_core::Func1<String, crate::UnknownType>,
    encode: purust_core::Func1<crate::UnknownType, String>,
    fingerprint: purust_core::Func1<crate::UnknownType, String>,
) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        let path = std::env::var("DIAG_CORPUS").expect("Data.JsonDecoding: DIAG_CORPUS must be set");
        let raw = std::fs::read_to_string(&path)
            .unwrap_or_else(|error| panic!("Data.JsonDecoding: cannot read {}: {}", path, error));
        let text = purust_core::purust_string_from_utf8(&raw);
        let corpus = Purs_Data_Argonaut_Core::purust_json_parse_text(&text)
            .unwrap_or_else(|error| panic!("Data.JsonDecoding: corpus is not JSON: {}", error));
        let entries = Purs_Data_Argonaut_Core::purust_json_array_items(&corpus)
            .expect("Data.JsonDecoding: corpus must be an array");
        let mut names = Vec::with_capacity(entries.len());
        let mut contents = Vec::with_capacity(entries.len());
        let mut benchmark = Vec::with_capacity(entries.len());
        for entry in &entries {
            names.push(driver_field(entry, "name").unwrap_string());
            contents.push(driver_field(entry, "contents").unwrap_string());
            benchmark.push(matches!(
                driver_field(entry, "benchmark").resolve(),
                crate::Value::Bool(true)
            ));
        }
        let modules = entries.len();
        let indices: Vec<usize> = (0..modules).filter(|index| benchmark[*index]).collect();
        assert!(!indices.is_empty(), "Data.JsonDecoding: no timed cases");
        let parsed: Vec<crate::UnknownType> =
            contents.iter().map(|content| parse(content.clone())).collect();
        let expected_json: Vec<String> = contents
            .iter()
            .map(|content| Purs_Data_Argonaut_Core::purust_canonical_hash(content))
            .collect();
        let expected_ast: Vec<String> = parsed
            .iter()
            .map(|value| {
                let raw = fingerprint(decode(value.clone()));
                Purs_Data_Argonaut_Core::purust_canonical_hash(&raw)
            })
            .collect();
        let mut report = String::from("{\"backend\":\"rust\",\"modules\":");
        report.push_str(&modules.to_string());
        report.push_str(",\"timed_cases\":");
        report.push_str(&indices.len().to_string());
        report.push_str(",\"names\":[");
        for (index, name) in names.iter().enumerate() {
            if index > 0 {
                report.push(',');
            }
            report.push_str(&driver_json_string(name));
        }
        report.push_str("],\"fingerprints\":[");
        for (index, hash) in expected_ast.iter().enumerate() {
            if index > 0 {
                report.push(',');
            }
            report.push_str(&driver_json_string(hash));
        }
        report.push_str("],\"json_fingerprints\":[");
        for (index, hash) in expected_json.iter().enumerate() {
            if index > 0 {
                report.push(',');
            }
            report.push_str(&driver_json_string(hash));
        }
        report.push_str("],\"phases\":{");
        let phases = ["parse", "decode", "combined"];
        for (phase_index, phase) in phases.iter().enumerate() {
            if phase_index > 0 {
                report.push(',');
            }
            report.push('"');
            report.push_str(phase);
            report.push_str("\":{\"samples\":[");
            let mut samples: Vec<f64> = Vec::new();
            let mut best = f64::INFINITY;
            for pass in 0..7 {
                let start = Instant::now();
                let mut results: Vec<crate::UnknownType> = Vec::with_capacity(indices.len());
                for index in &indices {
                    let result = match *phase {
                        "parse" => parse(contents[*index].clone()),
                        "decode" => decode(parsed[*index].clone()),
                        _ => decode_text(contents[*index].clone()),
                    };
                    results.push(result);
                }
                let elapsed = start.elapsed().as_nanos() as f64 / 1000.0;
                for (slot, result) in results.into_iter().enumerate() {
                    let index = indices[slot];
                    let (raw, expected) = if *phase == "parse" {
                        (encode(result), &expected_json[index])
                    } else {
                        (fingerprint(result), &expected_ast[index])
                    };
                    let hash = Purs_Data_Argonaut_Core::purust_canonical_hash(&raw);
                    assert_eq!(hash, *expected, "Data.JsonDecoding: unstable {} output", phase);
                }
                if pass >= 2 {
                    best = best.min(elapsed);
                    samples.push(elapsed);
                }
            }
            for (sample_index, sample) in samples.iter().enumerate() {
                if sample_index > 0 {
                    report.push(',');
                }
                report.push_str("{\"time_us\":");
                report.push_str(&driver_json_number(*sample));
                report.push('}');
            }
            report.push_str("],\"time_us\":");
            report.push_str(&driver_json_number(best));
            report.push('}');
        }
        report.push_str("}}");
        println!("{}", purust_core::purust_string_to_utf8_lossy(&report));
        crate::Value::Unit
    })))
}
