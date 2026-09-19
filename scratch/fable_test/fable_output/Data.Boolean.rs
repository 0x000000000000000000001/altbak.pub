pub mod PureScript_Data_Boolean {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    pub fn Data_Boolean_otherwise() -> &dyn Any {
        static Data_Boolean_otherwise: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Boolean_otherwise.get_or_init(|| &true)
    }
}
