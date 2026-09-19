pub mod PureScript_Data_Char {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use crate::module_6a1c5ce6::PureScript_Data_Enum;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Char_toCharCode() -> &dyn Any {
        static Data_Char_toCharCode: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Char_toCharCode.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                              &&&PureScript_Data_Enum::Data_Enum_boundedEnumChar()))
    }
    pub fn Data_Char_fromCharCode() -> &dyn Any {
        static Data_Char_fromCharCode: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Char_fromCharCode.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                &&&PureScript_Data_Enum::Data_Enum_boundedEnumChar()))
    }
}
