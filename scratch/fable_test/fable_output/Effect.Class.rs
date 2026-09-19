pub mod PureScript_Effect_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Effect_Class_MonadEffectusd_Dict() -> &dyn Any {
        static Effect_Class_MonadEffectusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_MonadEffectusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Effect_Class_monadEffectEffect() -> &dyn Any {
        static Effect_Class_monadEffectEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_monadEffectEffect.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Class::Effect_Class_MonadEffectusd_Dict(),
                                                                                        &&&add(string("liftEffect"),
                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                               add(string("Monad0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Effect::Effect_monadEffect()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Effect_Class_liftEffect() -> &dyn Any {
        static Effect_Class_liftEffect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Class_liftEffect.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("liftEffect"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}
