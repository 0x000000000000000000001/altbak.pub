pub mod PureScript_Data_Exists {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use crate::module_2a7662d2::PureScript_Unsafe_Coerce;
    pub fn Data_Exists_runExists() -> &dyn Any {
        static Data_Exists_runExists: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Exists_runExists.get_or_init(||
                                              &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce())
    }
    pub fn Data_Exists_mkExists() -> &dyn Any {
        static Data_Exists_mkExists: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Exists_mkExists.get_or_init(||
                                             &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce())
    }
}
