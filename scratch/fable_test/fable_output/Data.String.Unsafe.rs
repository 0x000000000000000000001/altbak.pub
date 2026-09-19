pub mod PureScript_Data_String_Unsafe {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    pub fn Data_String_Unsafe_char() -> &dyn Any {
        static Data_String_Unsafe_char: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Unsafe_char.get_or_init(||
                                                &Func1::new(move |arg0|
                                                                panic!("{}",
                                                                       1_i32.get_Message(),)))
    }
    pub fn Data_String_Unsafe_charAt() -> &dyn Any {
        static Data_String_Unsafe_charAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Unsafe_charAt.get_or_init(||
                                                  &Func1::new(move |arg0|
                                                                  panic!("{}",
                                                                         1_i32.get_Message(),)))
    }
}
