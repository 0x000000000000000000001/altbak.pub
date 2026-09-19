pub mod PureScript_Data_Int_Bits {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    pub mod Data_Int_Bits_FFI {
        use super::*;
        pub fn and(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &(n1.clone() & n2.clone())
        }
        pub fn or(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &(n1.clone() | n2.clone())
        }
        pub fn xor(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &(n1.clone() ^ n2.clone())
        }
        pub fn shl(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &(n1.clone() << n2.clone())
        }
        pub fn shr(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &(n1.clone() >> n2.clone())
        }
        pub fn zshr(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &((n1.clone() >> n2.clone()) as i32)
        }
        pub fn complement(n: &dyn Any) -> &dyn Any { &!n }
    }
    pub fn Data_Int_Bits_and() -> &dyn Any {
        static Data_Int_Bits_and: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_Bits_and.get_or_init(||
                                          &Func1::new(move |n|
                                                          Func1::new({
                                                                         let n
                                                                             =
                                                                             n.clone();
                                                                         move
                                                                             |n_1|
                                                                             PureScript_Data_Int_Bits::Data_Int_Bits_FFI::and(&n,
                                                                                                                              n_1)
                                                                     })))
    }
    pub fn Data_Int_Bits_complement() -> &dyn Any {
        static Data_Int_Bits_complement: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_Bits_complement.get_or_init(||
                                                 &Func1::new(move |n|
                                                                 PureScript_Data_Int_Bits::Data_Int_Bits_FFI::complement(n)))
    }
    pub fn Data_Int_Bits_or() -> &dyn Any {
        static Data_Int_Bits_or: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_Bits_or.get_or_init(||
                                         &Func1::new(move |n|
                                                         Func1::new({
                                                                        let n
                                                                            =
                                                                            n.clone();
                                                                        move
                                                                            |n_1|
                                                                            PureScript_Data_Int_Bits::Data_Int_Bits_FFI::or(&n,
                                                                                                                            n_1)
                                                                    })))
    }
    pub fn Data_Int_Bits_shl() -> &dyn Any {
        static Data_Int_Bits_shl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_Bits_shl.get_or_init(||
                                          &Func1::new(move |n|
                                                          Func1::new({
                                                                         let n
                                                                             =
                                                                             n.clone();
                                                                         move
                                                                             |n_1|
                                                                             PureScript_Data_Int_Bits::Data_Int_Bits_FFI::shl(&n,
                                                                                                                              n_1)
                                                                     })))
    }
    pub fn Data_Int_Bits_shr() -> &dyn Any {
        static Data_Int_Bits_shr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_Bits_shr.get_or_init(||
                                          &Func1::new(move |n|
                                                          Func1::new({
                                                                         let n
                                                                             =
                                                                             n.clone();
                                                                         move
                                                                             |n_1|
                                                                             PureScript_Data_Int_Bits::Data_Int_Bits_FFI::shr(&n,
                                                                                                                              n_1)
                                                                     })))
    }
    pub fn Data_Int_Bits_xor() -> &dyn Any {
        static Data_Int_Bits_xor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_Bits_xor.get_or_init(||
                                          &Func1::new(move |n|
                                                          Func1::new({
                                                                         let n
                                                                             =
                                                                             n.clone();
                                                                         move
                                                                             |n_1|
                                                                             PureScript_Data_Int_Bits::Data_Int_Bits_FFI::xor(&n,
                                                                                                                              n_1)
                                                                     })))
    }
    pub fn Data_Int_Bits_zshr() -> &dyn Any {
        static Data_Int_Bits_zshr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_Bits_zshr.get_or_init(||
                                           &Func1::new(move |n|
                                                           Func1::new({
                                                                          let n
                                                                              =
                                                                              n.clone();
                                                                          move
                                                                              |n_1|
                                                                              PureScript_Data_Int_Bits::Data_Int_Bits_FFI::zshr(&n,
                                                                                                                                n_1)
                                                                      })))
    }
}
