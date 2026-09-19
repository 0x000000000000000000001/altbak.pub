pub mod PureScript_Control_Monad_ST_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_5b0d6b47::PureScript_Control_Monad_ST_Global;
    use crate::module_fcf3066b::PureScript_Control_Monad_ST_Internal;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_ST_Class_MonadSTusd_Dict() -> &dyn Any {
        static Control_Monad_ST_Class_MonadSTusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Class_MonadSTusd_Dict.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Control_Monad_ST_Class_monadSTST() -> &dyn Any {
        static Control_Monad_ST_Class_monadSTST: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Class_monadSTST.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                          &&&add(string("liftST"),
                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                   &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                                 add(string("Monad0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_monadST()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Control_Monad_ST_Class_monadSTEffect() -> &dyn Any {
        static Control_Monad_ST_Class_monadSTEffect: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_ST_Class_monadSTEffect.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Class::Control_Monad_ST_Class_MonadSTusd_Dict(),
                                                                                              &&&add(string("liftST"),
                                                                                                     &&PureScript_Control_Monad_ST_Global::Control_Monad_ST_Global_toEffect(),
                                                                                                     add(string("Monad0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Effect::Effect_monadEffect()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Control_Monad_ST_Class_liftST() -> &dyn Any {
        static Control_Monad_ST_Class_liftST: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_ST_Class_liftST.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("liftST"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}
