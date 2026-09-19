pub mod PureScript_Control_Comonad_Store {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use crate::module_a7b6ae02::PureScript_Control_Comonad_Store_Trans;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Store_unwrap() -> &dyn Any {
        static Control_Comonad_Store_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_unwrap.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Comonad_Store_store() -> &dyn Any {
        static Control_Comonad_Store_store: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_store.get_or_init(||
                                                    &Func1::new(move |f|
                                                                    &Func1::new({
                                                                                    let f
                                                                                        =
                                                                                        f.clone();
                                                                                    move
                                                                                        |x|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                            &&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_StoreT()),
                                                                                                                         &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                                                                                                                                    &&&f),
                                                                                                                                                                                   x.clone())))
                                                                                })))
    }
    pub fn Control_Comonad_Store_runStore() -> &dyn Any {
        static Control_Comonad_Store_runStore: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_runStore.get_or_init(||
                                                       &Func1::new(move |v|
                                                                       {
                                                                           let s =
                                                                               Sharpurs_Prelude::unbox(v);
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_swap(),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                     &&&PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                                                                                                                  &&&PureScript_Control_Comonad_Store::Control_Comonad_Store_unwrap()),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_swap(),
                                                                                                                                                                                  &&&s)))
                                                                       }))
    }
}
