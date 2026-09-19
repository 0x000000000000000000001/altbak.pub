pub mod PureScript_Control_Monad_ST_Global {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use crate::module_2a7662d2::PureScript_Unsafe_Coerce;
    pub fn Control_Monad_ST_Global_toEffect() -> &dyn Any {
        static Control_Monad_ST_Global_toEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Global_toEffect.get_or_init(||
                                                         &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce())
    }
}
