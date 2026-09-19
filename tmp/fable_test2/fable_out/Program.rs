#![allow(dead_code,)]
#![allow(non_camel_case_types,)]
#![allow(non_snake_case,)]
#![allow(non_upper_case_globals,)]
#![allow(unreachable_code,)]
#![allow(unused_attributes,)]
#![allow(unused_imports,)]
#![allow(unused_macros,)]
#![allow(unused_parens,)]
#![allow(unused_variables,)]
use fable_library_rust::NativeArray_::array_from;
use fable_library_rust::String_::fromString;
pub mod Program {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::defaultOf;
    use fable_library_rust::NativeArray_::Array;
    use fable_library_rust::String_::string;
    pub fn main(argv: Array<string>) -> i32 {
        println!("Hello world");
        defaultOf::<&dyn Any>();
        println!("{} ms", defaultOf::<&dyn Any>());
        0_i32
    }
}
pub fn main() {
    let args = std::env::args().skip(1).map(fromString).collect();
    Program::main(array_from(args));
}
