pub mod PureScript_Data_Array_ST_Partial {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_1d594369::PureScript_Control_Monad_ST_Uncurried;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_Array_ST_Partial_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        pub fn peekImpl(iVal: &dyn Any, xs: &dyn Any) -> &dyn Any {
            xs[iVal].clone()
        }
        pub fn pokeImpl(iVal: &dyn Any, a: &dyn Any, xs: &dyn Any)
         -> &dyn Any {
            xs[iVal] = a.clone();
            defaultOf()
        }
    }
    pub fn Data_Array_ST_Partial_peekImpl() -> &dyn Any {
        static Data_Array_ST_Partial_peekImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Partial_peekImpl.get_or_init(||
                                                       &Func1::new(move |iVal|
                                                                       Func1::new({
                                                                                      let iVal
                                                                                          =
                                                                                          iVal.clone();
                                                                                      move
                                                                                          |xs|
                                                                                          PureScript_Data_Array_ST_Partial::Data_Array_ST_Partial_FFI::peekImpl(&iVal,
                                                                                                                                                                xs)
                                                                                  })))
    }
    pub fn Data_Array_ST_Partial_pokeImpl() -> &dyn Any {
        static Data_Array_ST_Partial_pokeImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Partial_pokeImpl.get_or_init(||
                                                       &Func1::new(move |iVal|
                                                                       Func1::new({
                                                                                      let iVal
                                                                                          =
                                                                                          iVal.clone();
                                                                                      move
                                                                                          |a|
                                                                                          Func1::new({
                                                                                                         let a
                                                                                                             =
                                                                                                             a.clone();
                                                                                                         move
                                                                                                             |xs|
                                                                                                             PureScript_Data_Array_ST_Partial::Data_Array_ST_Partial_FFI::pokeImpl(&iVal,
                                                                                                                                                                                   &a,
                                                                                                                                                                                   xs)
                                                                                                     })
                                                                                  })))
    }
    pub fn Data_Array_ST_Partial_poke() -> &dyn Any {
        static Data_Array_ST_Partial_poke: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Partial_poke.get_or_init(||
                                                   &Func1::new(move
                                                                   |usd__unused|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn3(),
                                                                                                    &&&PureScript_Data_Array_ST_Partial::Data_Array_ST_Partial_pokeImpl())))
    }
    pub fn Data_Array_ST_Partial_peek() -> &dyn Any {
        static Data_Array_ST_Partial_peek: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Partial_peek.get_or_init(||
                                                   &Func1::new(move
                                                                   |usd__unused|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Uncurried::Control_Monad_ST_Uncurried_runSTFn2(),
                                                                                                    &&&PureScript_Data_Array_ST_Partial::Data_Array_ST_Partial_peekImpl())))
    }
}
