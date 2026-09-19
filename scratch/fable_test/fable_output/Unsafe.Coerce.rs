pub mod PureScript_Unsafe_Coerce {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    pub mod Unsafe_Coerce_FFI {
        use super::*;
        use fable_library_rust::Native_::Func1;
        pub fn unsafeCoerce() -> &dyn Any {
            static unsafeCoerce: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            unsafeCoerce.get_or_init(|| &Func1::new(move |x| x.clone()))
        }
    }
    pub fn Unsafe_Coerce_unsafeCoerce() -> &dyn Any {
        static Unsafe_Coerce_unsafeCoerce: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Unsafe_Coerce_unsafeCoerce.get_or_init(||
                                                   &PureScript_Unsafe_Coerce::Unsafe_Coerce_FFI::unsafeCoerce())
    }
}
