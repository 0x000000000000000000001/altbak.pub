pub mod PureScript_Control_Comonad_Trans_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_bc5f9127::PureScript_Control_Monad_Identity_Trans;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Trans_Class_ComonadTransusd_Dict() -> &dyn Any {
        static Control_Comonad_Trans_Class_ComonadTransusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Trans_Class_ComonadTransusd_Dict.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |x|
                                                                                         x.clone()))
    }
    pub fn Control_Comonad_Trans_Class_lower() -> &dyn Any {
        static Control_Comonad_Trans_Class_lower: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Trans_Class_lower.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("lower"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Comonad_Trans_Class_comonadTransIdentityT() -> &dyn Any {
        static Control_Comonad_Trans_Class_comonadTransIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Trans_Class_comonadTransIdentityT.get_or_init(||
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_ComonadTransusd_Dict(),
                                                                                                           &&&add(string("lower"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |dictComonad|
                                                                                                                                   &PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_runIdentityT()),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>())))
    }
}
