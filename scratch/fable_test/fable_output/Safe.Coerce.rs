pub mod PureScript_Safe_Coerce {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_2a7662d2::PureScript_Unsafe_Coerce;
    pub fn Safe_Coerce_coerce() -> &dyn Any {
        static Safe_Coerce_coerce: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Safe_Coerce_coerce.get_or_init(||
                                           &Func1::new(move |usd__unused|
                                                           &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce()))
    }
}
