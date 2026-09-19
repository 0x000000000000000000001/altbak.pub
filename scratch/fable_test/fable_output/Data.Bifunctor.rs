pub mod PureScript_Data_Bifunctor {
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
    use crate::module_a8445950::PureScript_Data_Const;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Bifunctor_identity() -> &dyn Any {
        static Data_Bifunctor_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_identity.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bifunctor_identity1() -> &dyn Any {
        static Data_Bifunctor_identity1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_identity1.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                  &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Bifunctor_Bifunctorusd_Dict() -> &dyn Any {
        static Data_Bifunctor_Bifunctorusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_Bifunctorusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Bifunctor_bimap() -> &dyn Any {
        static Data_Bifunctor_bimap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_bimap.get_or_init(||
                                             &Func1::new(move |dict|
                                                             find(string("bimap"),
                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bifunctor_bivoid() -> &dyn Any {
        static Data_Bifunctor_bivoid: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_bivoid.get_or_init(||
                                              &Func1::new(move |dictBifunctor|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                     dictBifunctor),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                     &&&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                  &&&PureScript_Data_Unit::Data_Unit_unit()))))
    }
    pub fn Data_Bifunctor_lmap() -> &dyn Any {
        static Data_Bifunctor_lmap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_lmap.get_or_init(||
                                            &Func1::new(move |dictBifunctor|
                                                            &Func1::new({
                                                                            let dictBifunctor
                                                                                =
                                                                                dictBifunctor.clone();
                                                                            move
                                                                                |f|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                       &&&dictBifunctor),
                                                                                                                                                    f),
                                                                                                                 &&&PureScript_Data_Bifunctor::Data_Bifunctor_identity())
                                                                        })))
    }
    pub fn Data_Bifunctor_rmap() -> &dyn Any {
        static Data_Bifunctor_rmap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_rmap.get_or_init(||
                                            &Func1::new(move |dictBifunctor|
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                dictBifunctor),
                                                                                             &&&PureScript_Data_Bifunctor::Data_Bifunctor_identity1())))
    }
    pub fn Data_Bifunctor_bifunctorTuple() -> &dyn Any {
        static Data_Bifunctor_bifunctorTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_bifunctorTuple.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                       &&&add(string("bimap"),
                                                                                              &&Func1::new(move
                                                                                                               |f|
                                                                                                               &Func1::new({
                                                                                                                               let f
                                                                                                                                   =
                                                                                                                                   f.clone();
                                                                                                                               move
                                                                                                                                   |g|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let g
                                                                                                                                                       =
                                                                                                                                                       g.clone();
                                                                                                                                                   move
                                                                                                                                                       |v|
                                                                                                                                                       {
                                                                                                                                                           let matchValue =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                           let matchValue_1 =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&g);
                                                                                                                                                           let matchValue_2:
                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                                           &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                    &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                    &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                       })))
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           })),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Bifunctor_bifunctorEither() -> &dyn Any {
        static Data_Bifunctor_bifunctorEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_bifunctorEither.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                        &&&add(string("bimap"),
                                                                                               &&Func1::new(move
                                                                                                                |v|
                                                                                                                &Func1::new({
                                                                                                                                let v
                                                                                                                                    =
                                                                                                                                    v.clone();
                                                                                                                                move
                                                                                                                                    |v1|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let v1
                                                                                                                                                        =
                                                                                                                                                        v1.clone();
                                                                                                                                                    move
                                                                                                                                                        |v2|
                                                                                                                                                        {
                                                                                                                                                            let matchValue =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                            let matchValue_1 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                                            let matchValue_2:
                                                                                                                                                                    LrcPtr<Data_Either_Either> =
                                                                                                                                                                Sharpurs_Prelude::unbox(v2);
                                                                                                                                                            match matchValue_2.as_ref()
                                                                                                                                                                {
                                                                                                                                                                Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_2_1_0)
                                                                                                                                                                =>
                                                                                                                                                                &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                            &&matchValue_2_1_0))),
                                                                                                                                                                Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_2_0_0)
                                                                                                                                                                =>
                                                                                                                                                                &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                           &&matchValue_2_0_0))),
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                            })),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Bifunctor_bifunctorConst() -> &dyn Any {
        static Data_Bifunctor_bifunctorConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bifunctor_bifunctorConst.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                       &&&add(string("bimap"),
                                                                                              &&Func1::new(move
                                                                                                               |f|
                                                                                                               &Func1::new({
                                                                                                                               let f
                                                                                                                                   =
                                                                                                                                   f.clone();
                                                                                                                               move
                                                                                                                                   |v|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let v
                                                                                                                                                       =
                                                                                                                                                       v.clone();
                                                                                                                                                   move
                                                                                                                                                       |v1|
                                                                                                                                                       {
                                                                                                                                                           let matchValue =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                           let matchValue_1 =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                           let matchValue_2 =
                                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                               &&&matchValue_2))
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           })),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
}
