pub mod PureScript_Data_Functor_Coproduct_Inject {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_97eca9ab::PureScript_Data_Functor_Coproduct;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Coproduct_Inject_Injectusd_Dict() -> &dyn Any {
        static Data_Functor_Coproduct_Inject_Injectusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_Inject_Injectusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Data_Functor_Coproduct_Inject_prj() -> &dyn Any {
        static Data_Functor_Coproduct_Inject_prj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_Inject_prj.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("prj"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Functor_Coproduct_Inject_injectReflexive() -> &dyn Any {
        static Data_Functor_Coproduct_Inject_injectReflexive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_Inject_injectReflexive.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct_Inject::Data_Functor_Coproduct_Inject_Injectusd_Dict(),
                                                                                                       &&&add(string("inj"),
                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                                              add(string("prj"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |usd__arg1|
                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone()))),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Functor_Coproduct_Inject_injectLeft() -> &dyn Any {
        static Data_Functor_Coproduct_Inject_injectLeft:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_Inject_injectLeft.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct_Inject::Data_Functor_Coproduct_Inject_Injectusd_Dict(),
                                                                                                  &&&add(string("inj"),
                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                              &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                           &&&Func1::new(move
                                                                                                                                                             |usd__arg1|
                                                                                                                                                             &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                         add(string("prj"),
                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                    |usd__arg1_1|
                                                                                                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                  &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Functor_Coproduct_Inject_inj() -> &dyn Any {
        static Data_Functor_Coproduct_Inject_inj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_Inject_inj.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("inj"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Functor_Coproduct_Inject_injectRight() -> &dyn Any {
        static Data_Functor_Coproduct_Inject_injectRight:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_Inject_injectRight.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictInject|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct_Inject::Data_Functor_Coproduct_Inject_Injectusd_Dict(),
                                                                                                                   &&&add(string("inj"),
                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                               &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                    |usd__arg1|
                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct_Inject::Data_Functor_Coproduct_Inject_inj(),
                                                                                                                                                                                                                                  dictInject))),
                                                                                                                          add(string("prj"),
                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct_Inject::Data_Functor_Coproduct_Inject_prj(),
                                                                                                                                                                                                   dictInject)),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))))
    }
}
