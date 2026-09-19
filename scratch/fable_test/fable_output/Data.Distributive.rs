pub mod PureScript_Data_Distributive {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_6b8c0d95::PureScript_Type_Equality;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_Distributive_unwrap() -> &dyn Any {
        static Data_Distributive_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_unwrap.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Distributive_unwrap1() -> &dyn Any {
        static Data_Distributive_unwrap1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_unwrap1.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Distributive_identity() -> &dyn Any {
        static Data_Distributive_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_identity.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                    &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Distributive_Distributiveusd_Dict() -> &dyn Any {
        static Data_Distributive_Distributiveusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_Distributiveusd_Dict.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Data_Distributive_distributiveIdentity() -> &dyn Any {
        static Data_Distributive_distributiveIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_distributiveIdentity.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_Distributiveusd_Dict(),
                                                                                                &&&add(string("distribute"),
                                                                                                       &&Func1::new(move
                                                                                                                        |dictFunctor|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                            &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                               dictFunctor),
                                                                                                                                                                                            &&&PureScript_Data_Distributive::Data_Distributive_unwrap()))),
                                                                                                       add(string("collect"),
                                                                                                           &&Func1::new(move
                                                                                                                            |dictFunctor_1|
                                                                                                                            &Func1::new({
                                                                                                                                            let dictFunctor_1
                                                                                                                                                =
                                                                                                                                                dictFunctor_1.clone();
                                                                                                                                            move
                                                                                                                                                |f|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                    &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                       &&&dictFunctor_1),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Distributive::Data_Distributive_unwrap1()),
                                                                                                                                                                                                                                                       f)))
                                                                                                                                        })),
                                                                                                           add(string("Functor0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Identity::Data_Identity_functorIdentity()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>())))))
    }
    pub fn Data_Distributive_distribute() -> &dyn Any {
        static Data_Distributive_distribute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_distribute.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("distribute"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Distributive_distributiveFunction_004019() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_Distributiveusd_Dict(),
                                         &&&add(string("distribute"),
                                                &&Func1::new(move
                                                                 |dictFunctor|
                                                                 &Func1::new({
                                                                                 let dictFunctor
                                                                                     =
                                                                                     dictFunctor.clone();
                                                                                 move
                                                                                     |a|
                                                                                     &Func1::new({
                                                                                                     let a
                                                                                                         =
                                                                                                         a.clone();
                                                                                                     move
                                                                                                         |e|
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                &&&dictFunctor),
                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                               let e
                                                                                                                                                                                                   =
                                                                                                                                                                                                   e.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |v|
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                       v),
                                                                                                                                                                                                                                    &&&e)
                                                                                                                                                                                           })),
                                                                                                                                          &&&a)
                                                                                                 })
                                                                             })),
                                                add(string("collect"),
                                                    &&Func1::new({
                                                                     let Data_Distributive_distributiveFunction_004019_002d1
                                                                         =
                                                                         Data_Distributive_distributiveFunction_004019_002d1.clone();
                                                                     move
                                                                         |dictFunctor_1|
                                                                         &Func1::new({
                                                                                         let Data_Distributive_distributiveFunction_004019_002d1
                                                                                             =
                                                                                             Data_Distributive_distributiveFunction_004019_002d1.clone();
                                                                                         let dictFunctor_1
                                                                                             =
                                                                                             dictFunctor_1.clone();
                                                                                         move
                                                                                             |f|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_distribute(),
                                                                                                                                                                                                                                       &&&Data_Distributive_distributiveFunction_004019_002d1.Value),
                                                                                                                                                                                                    &&&dictFunctor_1)),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                    &&&dictFunctor_1),
                                                                                                                                                                 f))
                                                                                     })
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Functor::Data_Functor_functorFn()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Distributive_distributiveFunction_004019_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Distributive_distributiveFunction_004019_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Distributive_distributiveFunction_004019_002d1.get_or_init(||
                                                                            Lazy(Data_Distributive_distributiveFunction_004019.clone()))
    }
    pub fn Data_Distributive_distributiveFunction() -> &dyn Any {
        static Data_Distributive_distributiveFunction:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_distributiveFunction.get_or_init(||
                                                               Data_Distributive_distributiveFunction_004019_002d1.Value)
    }
    pub fn Data_Distributive_cotraverse() -> &dyn Any {
        static Data_Distributive_cotraverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_cotraverse.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictDistributive|
                                                                     {
                                                                         let Functor0 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictDistributive)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                         let distribute1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_distribute(),
                                                                                                              dictDistributive);
                                                                         &Func1::new({
                                                                                         let Functor0
                                                                                             =
                                                                                             Functor0.clone();
                                                                                         let distribute1
                                                                                             =
                                                                                             distribute1.clone();
                                                                                         move
                                                                                             |dictFunctor|
                                                                                             {
                                                                                                 let distribute2 =
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&distribute1,
                                                                                                                                      dictFunctor);
                                                                                                 &Func1::new({
                                                                                                                 let distribute2
                                                                                                                     =
                                                                                                                     distribute2.clone();
                                                                                                                 move
                                                                                                                     |f|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                               &&&Functor0),
                                                                                                                                                                                                                            f)),
                                                                                                                                                      &&&distribute2)
                                                                                                             })
                                                                                             }
                                                                                     })
                                                                     }))
    }
    pub fn Data_Distributive_collectDefault() -> &dyn Any {
        static Data_Distributive_collectDefault: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_collectDefault.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictDistributive|
                                                                         {
                                                                             let distribute1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_distribute(),
                                                                                                                  dictDistributive);
                                                                             &Func1::new({
                                                                                             let distribute1
                                                                                                 =
                                                                                                 distribute1.clone();
                                                                                             move
                                                                                                 |dictFunctor|
                                                                                                 {
                                                                                                     let distribute2 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&distribute1,
                                                                                                                                          dictFunctor);
                                                                                                     &Func1::new({
                                                                                                                     let dictFunctor
                                                                                                                         =
                                                                                                                         dictFunctor.clone();
                                                                                                                     let distribute2
                                                                                                                         =
                                                                                                                         distribute2.clone();
                                                                                                                     move
                                                                                                                         |f|
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                             &&&distribute2),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                &&&dictFunctor),
                                                                                                                                                                                             f))
                                                                                                                 })
                                                                                                 }
                                                                                         })
                                                                         }))
    }
    pub fn Data_Distributive_distributiveTuple_004027() -> &dyn Any {
        &Func1::new(move |dictTypeEquals|
                        PureScript_Data_Distributive::Data_Distributive_distributiveTuple_tco(dictTypeEquals))
    }
    pub fn Data_Distributive_distributiveTuple_004027_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Distributive_distributiveTuple_004027_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Distributive_distributiveTuple_004027_002d1.get_or_init(||
                                                                         Lazy(Data_Distributive_distributiveTuple_004027.clone()))
    }
    pub fn Data_Distributive_distributiveTuple_tco(dictTypeEquals: &dyn Any)
     -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_Distributiveusd_Dict(),
                                         &&&add(string("collect"),
                                                &&Func1::new({
                                                                 let dictTypeEquals
                                                                     =
                                                                     dictTypeEquals.clone();
                                                                 move
                                                                     |dictFunctor|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_collectDefault(),
                                                                                                                                         &&PureScript_Data_Distributive::Data_Distributive_distributiveTuple_tco(&&dictTypeEquals)),
                                                                                                      dictFunctor)
                                                             }),
                                                add(string("distribute"),
                                                    &&Func1::new({
                                                                     let dictTypeEquals
                                                                         =
                                                                         dictTypeEquals.clone();
                                                                     move
                                                                         |dictFunctor_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                  |usd__arg1|
                                                                                                                                                                                                  Func1::new({
                                                                                                                                                                                                                 let usd__arg1
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     usd__arg1.clone();
                                                                                                                                                                                                                 move
                                                                                                                                                                                                                     |usd__arg2|
                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                             usd__arg2.clone()))
                                                                                                                                                                                                             })),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Type_Equality::Type_Equality_from(),
                                                                                                                                                                                                                                                      &&&dictTypeEquals),
                                                                                                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                dictFunctor_1),
                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_snd()))
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Distributive_distributiveTuple() -> &dyn Any {
        static Data_Distributive_distributiveTuple: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Distributive_distributiveTuple.get_or_init(||
                                                            Data_Distributive_distributiveTuple_004027_002d1.Value)
    }
    pub fn Data_Distributive_collect() -> &dyn Any {
        static Data_Distributive_collect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Distributive_collect.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("collect"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Distributive_distributeDefault() -> &dyn Any {
        static Data_Distributive_distributeDefault: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Distributive_distributeDefault.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictDistributive|
                                                                            &Func1::new({
                                                                                            let dictDistributive
                                                                                                =
                                                                                                dictDistributive.clone();
                                                                                            move
                                                                                                |dictFunctor|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_collect(),
                                                                                                                                                                                                       &&&dictDistributive),
                                                                                                                                                                    dictFunctor),
                                                                                                                                 &&&PureScript_Data_Distributive::Data_Distributive_identity())
                                                                                        })))
    }
}
