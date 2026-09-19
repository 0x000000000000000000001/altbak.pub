pub mod PureScript_Data_String_Regex_Unsafe {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_db031083::PureScript_Data_String_Regex;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_String_Regex_Unsafe_identity() -> &dyn Any {
        static Data_String_Regex_Unsafe_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Regex_Unsafe_identity.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                           &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_String_Regex_Unsafe_unsafeRegex() -> &dyn Any {
        static Data_String_Regex_Unsafe_unsafeRegex: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_Regex_Unsafe_unsafeRegex.get_or_init(||
                                                             &Func1::new(move
                                                                             |s|
                                                                             &Func1::new({
                                                                                             let s
                                                                                                 =
                                                                                                 s.clone();
                                                                                             move
                                                                                                 |f|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                        &&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafeCrashWith()),
                                                                                                                                                                     &&&PureScript_Data_String_Regex_Unsafe::Data_String_Regex_Unsafe_identity()),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Regex::Data_String_Regex_regex(),
                                                                                                                                                                                                        &&&s),
                                                                                                                                                                     f))
                                                                                         })))
    }
}
