pub mod PureScript_Effect_Unsafe {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    pub mod Effect_Unsafe_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        pub fn unsafePerformEffect(fVal: &dyn Any) -> &dyn Any {
            fVal(defaultOf())
        }
    }
    pub fn Effect_Unsafe_unsafePerformEffect() -> &dyn Any {
        static Effect_Unsafe_unsafePerformEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Unsafe_unsafePerformEffect.get_or_init(||
                                                          &Func1::new(move
                                                                          |fVal|
                                                                          PureScript_Effect_Unsafe::Effect_Unsafe_FFI::unsafePerformEffect(fVal)))
    }
}
