pub mod PureScript_Data_Unit {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    pub mod Data_Unit_FFI {
        use super::*;
        use crate::module_aa21d1e7::Sharpurs_Prelude;
        pub fn unit() -> &dyn Any {
            static unit: MutCell<Option<&dyn Any>> = MutCell::new(None);
            unit.get_or_init(|| Sharpurs_Prelude::undefined())
        }
    }
    pub fn Data_Unit_unit() -> &dyn Any {
        static Data_Unit_unit: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Unit_unit.get_or_init(||
                                       &PureScript_Data_Unit::Data_Unit_FFI::unit())
    }
}
