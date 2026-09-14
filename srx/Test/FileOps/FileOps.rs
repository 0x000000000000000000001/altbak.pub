pub fn Test_FileOps_writeFileSync(path: String, content: String) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        let path = purust_core::purust_string_to_utf8_lossy(&path);
        let content = purust_core::purust_string_to_utf8_lossy(&content);
        std::fs::write(&path, content.as_bytes())
            .unwrap_or_else(|error| panic!("writeFileSync({path:?}): {error}"));
        crate::Value::Unit
    })))
}

pub fn Test_FileOps_readFileSync(path: String) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        let path = purust_core::purust_string_to_utf8_lossy(&path);
        let bytes = std::fs::read(&path)
            .unwrap_or_else(|error| panic!("readFileSync({path:?}): {error}"));
        // Match the JavaScript FFI's UTF-8 decoding, including malformed bytes.
        crate::Value::String(purust_core::purust_string_from_utf8(
            &String::from_utf8_lossy(&bytes),
        ))
    })))
}

pub fn Test_FileOps_loopE(count: i64, action: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        for _ in 0..count {
            action.unwrap_func1()(crate::Value::Unit);
        }
        crate::Value::Unit
    })))
}
