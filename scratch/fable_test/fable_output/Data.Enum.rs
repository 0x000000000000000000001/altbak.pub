pub mod PureScript_Data_Enum {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_Enum_fromCharCode() -> &dyn Any {
        static Data_Enum_fromCharCode: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_fromCharCode.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               panic!("{}",
                                                                      1_i32.get_Message(),)))
    }
    pub fn Data_Enum_toCharCode() -> &dyn Any {
        static Data_Enum_toCharCode: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_toCharCode.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             panic!("{}",
                                                                    1_i32.get_Message(),)))
    }
    pub fn Data_Enum_bottom() -> &dyn Any {
        static Data_Enum_bottom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_bottom.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                          &&&PureScript_Data_Bounded::Data_Bounded_boundedInt()))
    }
    pub fn Data_Enum_fromJust() -> &dyn Any {
        static Data_Enum_fromJust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_fromJust.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                            &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Enum_bottom1() -> &dyn Any {
        static Data_Enum_bottom1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_bottom1.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                           &&&PureScript_Data_Bounded::Data_Bounded_boundedChar()))
    }
    pub fn Data_Enum_top() -> &dyn Any {
        static Data_Enum_top: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Enum_top.get_or_init(||
                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                       &&&PureScript_Data_Bounded::Data_Bounded_boundedChar()))
    }
    pub fn Data_Enum_Enumusd_Dict() -> &dyn Any {
        static Data_Enum_Enumusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_Enumusd_Dict.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Enum_Cardinality() -> &dyn Any {
        static Data_Enum_Cardinality: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_Cardinality.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Enum_BoundedEnumusd_Dict() -> &dyn Any {
        static Data_Enum_BoundedEnumusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_BoundedEnumusd_Dict.get_or_init(||
                                                      &Func1::new(move |x|
                                                                      x.clone()))
    }
    pub fn Data_Enum_toEnum() -> &dyn Any {
        static Data_Enum_toEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_toEnum.get_or_init(||
                                         &Func1::new(move |dict|
                                                         find(string("toEnum"),
                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Enum_succ() -> &dyn Any {
        static Data_Enum_succ: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Enum_succ.get_or_init(||
                                       &Func1::new(move |dict|
                                                       find(string("succ"),
                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Enum_upFromIncluding() -> &dyn Any {
        static Data_Enum_upFromIncluding: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_upFromIncluding.get_or_init(||
                                                  &Func1::new(move |dictEnum|
                                                                  {
                                                                      let succ1 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                           dictEnum);
                                                                      &Func1::new({
                                                                                      let succ1
                                                                                          =
                                                                                          succ1.clone();
                                                                                      move
                                                                                          |dictUnfoldable1|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                              dictUnfoldable1),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                    &&&PureScript_Control_Apply::Control_Apply_applyFn()),
                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                   |usd__arg1|
                                                                                                                                                                                                                   Func1::new({
                                                                                                                                                                                                                                  let usd__arg1
                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                      usd__arg1.clone();
                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                      |usd__arg2|
                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                              usd__arg2.clone()))
                                                                                                                                                                                                                              }))),
                                                                                                                                                              &&&succ1))
                                                                                  })
                                                                  }))
    }
    pub fn Data_Enum_showCardinality() -> &dyn Any {
        static Data_Enum_showCardinality: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_showCardinality.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                   &&&add(string("show"),
                                                                                          &&Func1::new(move
                                                                                                           |v|
                                                                                                           {
                                                                                                               let n =
                                                                                                                   Sharpurs_Prelude::unbox(v);
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                   &&&string("(Cardinality ")),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                                                                                                         &&&n)),
                                                                                                                                                                                   &&&string(")")))
                                                                                                           }),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Enum_pred() -> &dyn Any {
        static Data_Enum_pred: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Enum_pred.get_or_init(||
                                       &Func1::new(move |dict|
                                                       find(string("pred"),
                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Enum_ordCardinality() -> &dyn Any {
        static Data_Enum_ordCardinality: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_ordCardinality.get_or_init(||
                                                 &PureScript_Data_Ord::Data_Ord_ordInt())
    }
    pub fn Data_Enum_newtypeCardinality() -> &dyn Any {
        static Data_Enum_newtypeCardinality: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_newtypeCardinality.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                      &&&add(string("Coercible0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &Sharpurs_Prelude::Prim_undefined()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Enum_fromEnum() -> &dyn Any {
        static Data_Enum_fromEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_fromEnum.get_or_init(||
                                           &Func1::new(move |dict|
                                                           find(string("fromEnum"),
                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Enum_toEnumWithDefaults() -> &dyn Any {
        static Data_Enum_toEnumWithDefaults: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_toEnumWithDefaults.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBoundedEnum|
                                                                     {
                                                                         let bottom2 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bounded0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                         &Func1::new({
                                                                                         let bottom2
                                                                                             =
                                                                                             bottom2.clone();
                                                                                         let dictBoundedEnum
                                                                                             =
                                                                                             dictBoundedEnum.clone();
                                                                                         move
                                                                                             |low|
                                                                                             &Func1::new({
                                                                                                             let low
                                                                                                                 =
                                                                                                                 low.clone();
                                                                                                             move
                                                                                                                 |high|
                                                                                                                 &Func1::new({
                                                                                                                                 let high
                                                                                                                                     =
                                                                                                                                     high.clone();
                                                                                                                                 move
                                                                                                                                     |x|
                                                                                                                                     {
                                                                                                                                         let matchValue:
                                                                                                                                                 LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                           &&&dictBoundedEnum),
                                                                                                                                                                                                        x));
                                                                                                                                         match matchValue.as_ref()
                                                                                                                                             {
                                                                                                                                             Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                             =>
                                                                                                                                             {
                                                                                                                                                 let matchValue_1 =
                                                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                  x),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                     &&&dictBoundedEnum),
                                                                                                                                                                                                                                                  &&&bottom2)));
                                                                                                                                                 match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                  &matchValue_1)
                                                                                                                                                     {
                                                                                                                                                     0_i32
                                                                                                                                                     =>
                                                                                                                                                     &low,
                                                                                                                                                     _
                                                                                                                                                     =>
                                                                                                                                                     &high,
                                                                                                                                                 }
                                                                                                                                             }
                                                                                                                                             Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                             =>
                                                                                                                                             &match matchValue.as_ref()
                                                                                                                                                  {
                                                                                                                                                  Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                  =>
                                                                                                                                                  x.clone(),
                                                                                                                                                  _
                                                                                                                                                  =>
                                                                                                                                                  unreachable!(),
                                                                                                                                              },
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                             })
                                                                                                         })
                                                                                     })
                                                                     }))
    }
    pub fn Data_Enum_eqCardinality() -> &dyn Any {
        static Data_Enum_eqCardinality: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_eqCardinality.get_or_init(||
                                                &PureScript_Data_Eq::Data_Eq_eqInt())
    }
    pub fn Data_Enum_enumUnit() -> &dyn Any {
        static Data_Enum_enumUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumUnit.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                            &&&add(string("succ"),
                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                     &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                   add(string("pred"),
                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                         &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                       add(string("Ord0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Ord::Data_Ord_ordUnit()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>())))))
    }
    pub fn Data_Enum_enumTuple() -> &dyn Any {
        static Data_Enum_enumTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumTuple.get_or_init(||
                                            &Func1::new(move |dictEnum|
                                                            {
                                                                let ordTuple =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_ordTuple(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                               Sharpurs_Prelude::unbox(dictEnum)),
                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                &Func1::new({
                                                                                let dictEnum
                                                                                    =
                                                                                    dictEnum.clone();
                                                                                let ordTuple
                                                                                    =
                                                                                    ordTuple.clone();
                                                                                move
                                                                                    |dictBoundedEnum|
                                                                                    {
                                                                                        let Bounded0 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bounded0"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let Enum1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let ordTuple1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&ordTuple,
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                                                                         &&&add(string("succ"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let Bounded0
                                                                                                                                                     =
                                                                                                                                                     Bounded0.clone();
                                                                                                                                                 let Enum1
                                                                                                                                                     =
                                                                                                                                                     Enum1.clone();
                                                                                                                                                 move
                                                                                                                                                     |v|
                                                                                                                                                     {
                                                                                                                                                         let matchValue:
                                                                                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                                                         let a =
                                                                                                                                                             match matchValue.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                            _)
                                                                                                                                                                 =>
                                                                                                                                                                 x.clone(),
                                                                                                                                                             };
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                              |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                                                                              Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                             let usd__arg1
                                                                                                                                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                                                                                                                                 usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                                                                                                                                 |usd__arg2|
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                                                                                                                                         }))),
                                                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                                                                                                                                                                                                            &&&Bounded0))),
                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                                                                                                                                                                         &&&dictEnum),
                                                                                                                                                                                                                                                                                                                                      &&&a))),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                     |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                                                     |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                                                                                                                                    let usd__arg1_2
                                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                                        usd__arg1_2.clone();
                                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                                        |usd__arg2_1|
                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1_2,
                                                                                                                                                                                                                                                                                                                                                                                                usd__arg2_1.clone()))
                                                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                                                                                                   &&&a))),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                                                                &&&Enum1),
                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                }))
                                                                                                                                                     }
                                                                                                                                             }),
                                                                                                                                add(string("pred"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let Bounded0
                                                                                                                                                         =
                                                                                                                                                         Bounded0.clone();
                                                                                                                                                     let Enum1
                                                                                                                                                         =
                                                                                                                                                         Enum1.clone();
                                                                                                                                                     move
                                                                                                                                                         |v_1|
                                                                                                                                                         {
                                                                                                                                                             let matchValue_1:
                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                 Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                             let a_1 =
                                                                                                                                                                 match matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                _)
                                                                                                                                                                     =>
                                                                                                                                                                     x.clone(),
                                                                                                                                                                 };
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                  |usd__arg1_3|
                                                                                                                                                                                                                                                                                                                                                                                                                                  Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                 let usd__arg1_3
                                                                                                                                                                                                                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                                                                                                                                                                                                                     usd__arg1_3.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                                                                                                                                                                                                                     |usd__arg2_2|
                                                                                                                                                                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1_3,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             usd__arg2_2.clone()))
                                                                                                                                                                                                                                                                                                                                                                                                                                             }))),
                                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&Bounded0))),
                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                                                                                                                                                                                                             &&&dictEnum),
                                                                                                                                                                                                                                                                                                                                          &&&a_1))),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                         |usd__arg1_4|
                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_4.clone())))),
                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                                                         |usd__arg1_5|
                                                                                                                                                                                                                                                                                                                         Func1::new({
                                                                                                                                                                                                                                                                                                                                        let usd__arg1_5
                                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                                            usd__arg1_5.clone();
                                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                                            |usd__arg2_3|
                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1_5,
                                                                                                                                                                                                                                                                                                                                                                                                    usd__arg2_3.clone()))
                                                                                                                                                                                                                                                                                                                                    })),
                                                                                                                                                                                                                                                                                                       &&&a_1))),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                                                                                                    &&&Enum1),
                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }))
                                                                                                                                                         }
                                                                                                                                                 }),
                                                                                                                                    add(string("Ord0"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let ordTuple1
                                                                                                                                                             =
                                                                                                                                                             ordTuple1.clone();
                                                                                                                                                         move
                                                                                                                                                             |usd__unused|
                                                                                                                                                             &ordTuple1
                                                                                                                                                     }),
                                                                                                                                        empty::<string,
                                                                                                                                                &dyn Any>()))))
                                                                                    }
                                                                            })
                                                            }))
    }
    pub fn Data_Enum_enumOrdering() -> &dyn Any {
        static Data_Enum_enumOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumOrdering.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                                &&&add(string("succ"),
                                                                                       &&Func1::new(move
                                                                                                        |v|
                                                                                                        {
                                                                                                            let matchValue:
                                                                                                                    LrcPtr<Data_Ordering_Ordering> =
                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                            match matchValue.as_ref()
                                                                                                                {
                                                                                                                Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                =>
                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor))),
                                                                                                                Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                =>
                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                _
                                                                                                                =>
                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                                            }
                                                                                                        }),
                                                                                       add(string("pred"),
                                                                                           &&Func1::new(move
                                                                                                            |v_1|
                                                                                                            {
                                                                                                                let matchValue_1:
                                                                                                                        LrcPtr<Data_Ordering_Ordering> =
                                                                                                                    Sharpurs_Prelude::unbox(v_1);
                                                                                                                match matchValue_1.as_ref()
                                                                                                                    {
                                                                                                                    Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                    =>
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor))),
                                                                                                                    Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                    =>
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                }
                                                                                                            }),
                                                                                           add(string("Ord0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Ord::Data_Ord_ordOrdering()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))))
    }
    pub fn Data_Enum_enumMaybe() -> &dyn Any {
        static Data_Enum_enumMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumMaybe.get_or_init(||
                                            &Func1::new(move |dictBoundedEnum|
                                                            {
                                                                let Bounded0 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Bounded0"),
                                                                                                            Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                let Enum1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                            Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                let ordMaybe =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_ordMaybe(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                                                 &&&add(string("succ"),
                                                                                                        &&Func1::new({
                                                                                                                         let Bounded0
                                                                                                                             =
                                                                                                                             Bounded0.clone();
                                                                                                                         let Enum1
                                                                                                                             =
                                                                                                                             Enum1.clone();
                                                                                                                         move
                                                                                                                             |v|
                                                                                                                             {
                                                                                                                                 let matchValue:
                                                                                                                                         LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                                 match matchValue.as_ref()
                                                                                                                                     {
                                                                                                                                     Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                     =>
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                            &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                           |usd__arg1|
                                                                                                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                                            &&&Enum1),
                                                                                                                                                                                                         &&matchValue_1_0)),
                                                                                                                                     _
                                                                                                                                     =>
                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                                                                                    &&&Bounded0))))),
                                                                                                                                 }
                                                                                                                             }
                                                                                                                     }),
                                                                                                        add(string("pred"),
                                                                                                            &&Func1::new({
                                                                                                                             let Enum1
                                                                                                                                 =
                                                                                                                                 Enum1.clone();
                                                                                                                             move
                                                                                                                                 |v_1|
                                                                                                                                 {
                                                                                                                                     let matchValue_1:
                                                                                                                                             LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                         Sharpurs_Prelude::unbox(v_1);
                                                                                                                                     match matchValue_1.as_ref()
                                                                                                                                         {
                                                                                                                                         Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                         =>
                                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                                                                                                    &&&Enum1),
                                                                                                                                                                                                                                 &&matchValue_1_1_0))),
                                                                                                                                         _
                                                                                                                                         =>
                                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                         }),
                                                                                                            add(string("Ord0"),
                                                                                                                &&Func1::new({
                                                                                                                                 let ordMaybe
                                                                                                                                     =
                                                                                                                                     ordMaybe.clone();
                                                                                                                                 move
                                                                                                                                     |usd__unused|
                                                                                                                                     &ordMaybe
                                                                                                                             }),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
                                                            }))
    }
    pub fn Data_Enum_enumInt() -> &dyn Any {
        static Data_Enum_enumInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumInt.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                           &&&add(string("succ"),
                                                                                  &&Func1::new(move
                                                                                                   |n|
                                                                                                   {
                                                                                                       let matchValue =
                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                        n),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                        &&&PureScript_Data_Bounded::Data_Bounded_boundedInt())));
                                                                                                       match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                        &matchValue)
                                                                                                           {
                                                                                                           0_i32
                                                                                                           =>
                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                      n),
                                                                                                                                                                                                   &&&1_i32))),
                                                                                                           _
                                                                                                           =>
                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                       }
                                                                                                   }),
                                                                                  add(string("pred"),
                                                                                      &&Func1::new(move
                                                                                                       |n_1|
                                                                                                       {
                                                                                                           let matchValue_1 =
                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                               &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                            n_1),
                                                                                                                                                                         &&&PureScript_Data_Enum::Data_Enum_bottom()));
                                                                                                           match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                            &matchValue_1)
                                                                                                               {
                                                                                                               0_i32
                                                                                                               =>
                                                                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                          n_1),
                                                                                                                                                                                                       &&&1_i32))),
                                                                                                               _
                                                                                                               =>
                                                                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                           }
                                                                                                       }),
                                                                                      add(string("Ord0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))))
    }
    pub fn Data_Enum_enumFromTo() -> &dyn Any {
        static Data_Enum_enumFromTo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumFromTo.get_or_init(||
                                             &Func1::new(move |dictEnum|
                                                             {
                                                                 let Ord0 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                             Sharpurs_Prelude::unbox(dictEnum)),
                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                 let Eq0 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                             Sharpurs_Prelude::unbox(&&Ord0)),
                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                 let Ord01 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                             Sharpurs_Prelude::unbox(dictEnum)),
                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                 let succ1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                      dictEnum);
                                                                 let lessThanOrEq =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                      &&&Ord0);
                                                                 let pred1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                      dictEnum);
                                                                 let greaterThanOrEq =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                      &&&Ord0);
                                                                 &Func1::new({
                                                                                 let Eq0
                                                                                     =
                                                                                     Eq0.clone();
                                                                                 let Ord01
                                                                                     =
                                                                                     Ord01.clone();
                                                                                 let greaterThanOrEq
                                                                                     =
                                                                                     greaterThanOrEq.clone();
                                                                                 let lessThanOrEq
                                                                                     =
                                                                                     lessThanOrEq.clone();
                                                                                 let pred1
                                                                                     =
                                                                                     pred1.clone();
                                                                                 let succ1
                                                                                     =
                                                                                     succ1.clone();
                                                                                 move
                                                                                     |dictUnfoldable1|
                                                                                     {
                                                                                         let go =
                                                                                             &Func1::new(move
                                                                                                             |step|
                                                                                                             &Func1::new({
                                                                                                                             let step
                                                                                                                                 =
                                                                                                                                 step.clone();
                                                                                                                             move
                                                                                                                                 |op|
                                                                                                                                 &Func1::new({
                                                                                                                                                 let op
                                                                                                                                                     =
                                                                                                                                                     op.clone();
                                                                                                                                                 move
                                                                                                                                                     |to_var|
                                                                                                                                                     &Func1::new({
                                                                                                                                                                     let to_var
                                                                                                                                                                         =
                                                                                                                                                                         to_var.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |a|
                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&step,
                                                                                                                                                                                                                                                                                                                                        a)),
                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                    |a_prime|
                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_voidLeft(),
                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_guard(),
                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Maybe::Data_Maybe_alternativeMaybe()),
                                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&op,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 a_prime),
                                                                                                                                                                                                                                                                                                                                                                                                                              &&&to_var))),
                                                                                                                                                                                                                                                                                                                     a_prime)))))
                                                                                                                                                                 })
                                                                                                                                             })
                                                                                                                         }));
                                                                                         &Func1::new({
                                                                                                         let dictUnfoldable1
                                                                                                             =
                                                                                                             dictUnfoldable1.clone();
                                                                                                         let go
                                                                                                             =
                                                                                                             go.clone();
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
                                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                                     let matchValue_1 =
                                                                                                                                         Sharpurs_Prelude::unbox(v1);
                                                                                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                        &&&Eq0),
                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                  &&&matchValue_1))
                                                                                                                                        {
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_singleton(),
                                                                                                                                                                                                             &&&dictUnfoldable1),
                                                                                                                                                                          &&&matchValue)
                                                                                                                                     } else {
                                                                                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                                            &&&Ord01),
                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                                            {
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                                                                    &&&dictUnfoldable1),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                                                                                                                                                                                          &&&succ1),
                                                                                                                                                                                                                                                                                       &&&lessThanOrEq),
                                                                                                                                                                                                                                                    &&&matchValue_1)),
                                                                                                                                                                              &&&matchValue)
                                                                                                                                         } else {
                                                                                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                {
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                                                                        &&&dictUnfoldable1),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                                                                                                                                                                                              &&&pred1),
                                                                                                                                                                                                                                                                                           &&&greaterThanOrEq),
                                                                                                                                                                                                                                                        &&&matchValue_1)),
                                                                                                                                                                                  &&&matchValue)
                                                                                                                                             } else {
                                                                                                                                                 panic!("{}",
                                                                                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Enum.fs"),
                                  Data1: 55_i32,
                                  Data2: 1688_i32,}).get_Message(),)
                                                                                                                                             }
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                         })
                                                                                                     })
                                                                                     }
                                                                             })
                                                             }))
    }
    pub fn Data_Enum_enumFromThenTo() -> &dyn Any {
        static Data_Enum_enumFromThenTo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumFromThenTo.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictUnfoldable|
                                                                 &Func1::new({
                                                                                 let dictUnfoldable
                                                                                     =
                                                                                     dictUnfoldable.clone();
                                                                                 move
                                                                                     |dictFunctor|
                                                                                     &Func1::new({
                                                                                                     let dictFunctor
                                                                                                         =
                                                                                                         dictFunctor.clone();
                                                                                                     move
                                                                                                         |dictBoundedEnum|
                                                                                                         {
                                                                                                             let toEnum1 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                  dictBoundedEnum);
                                                                                                             let go =
                                                                                                                 &Func1::new(move
                                                                                                                                 |step|
                                                                                                                                 &Func1::new({
                                                                                                                                                 let step
                                                                                                                                                     =
                                                                                                                                                     step.clone();
                                                                                                                                                 move
                                                                                                                                                     |to_var|
                                                                                                                                                     &Func1::new({
                                                                                                                                                                     let to_var
                                                                                                                                                                         =
                                                                                                                                                                         to_var.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |e|
                                                                                                                                                                         {
                                                                                                                                                                             let matchValue =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&step);
                                                                                                                                                                             let matchValue_1 =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&to_var);
                                                                                                                                                                             let matchValue_2 =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(e);
                                                                                                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                             &&&matchValue_2),
                                                                                                                                                                                                                                          &&&matchValue_1))
                                                                                                                                                                                {
                                                                                                                                                                                 let e1_2 =
                                                                                                                                                                                     matchValue_2;
                                                                                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e1_2,
                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                                    &&&e1_2),
                                                                                                                                                                                                                                                                                                                                 &&&matchValue)))))
                                                                                                                                                                             } else {
                                                                                                                                                                                 if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                                                    {
                                                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                                                                                 } else {
                                                                                                                                                                                     panic!("{}",
                                                                                                                                                                                            LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Enum.fs"),
                                  Data1: 57_i32,
                                  Data2: 321_i32,}).get_Message(),)
                                                                                                                                                                                 }
                                                                                                                                                                             }
                                                                                                                                                                         }
                                                                                                                                                                 })
                                                                                                                                             }));
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                              &&&Func1::new({
                                                                                                                                                                let dictBoundedEnum
                                                                                                                                                                    =
                                                                                                                                                                    dictBoundedEnum.clone();
                                                                                                                                                                let go
                                                                                                                                                                    =
                                                                                                                                                                    go.clone();
                                                                                                                                                                let toEnum1
                                                                                                                                                                    =
                                                                                                                                                                    toEnum1.clone();
                                                                                                                                                                move
                                                                                                                                                                    |usd__unused|
                                                                                                                                                                    &Func1::new(move
                                                                                                                                                                                    |a|
                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                    let a
                                                                                                                                                                                                        =
                                                                                                                                                                                                        a.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |b|
                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                        let b
                                                                                                                                                                                                                            =
                                                                                                                                                                                                                            b.clone();
                                                                                                                                                                                                                        move
                                                                                                                                                                                                                            |c|
                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                let c_prime =
                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                                        &&&dictBoundedEnum),
                                                                                                                                                                                                                                                                     c);
                                                                                                                                                                                                                                let b_prime =
                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                                        &&&dictBoundedEnum),
                                                                                                                                                                                                                                                                     &&&b);
                                                                                                                                                                                                                                let a_prime =
                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                                        &&&dictBoundedEnum),
                                                                                                                                                                                                                                                                     &&&a);
                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                       &&&dictFunctor),
                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                          &&&toEnum1),
                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Enum::Data_Enum_fromJust())),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                                                                                                                                                                                                                          &&&dictUnfoldable),
                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&b_prime),
                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&a_prime)),
                                                                                                                                                                                                                                                                                                                                                                          &&&c_prime)),
                                                                                                                                                                                                                                                                                                    &&&a_prime))
                                                                                                                                                                                                                            }
                                                                                                                                                                                                                    })
                                                                                                                                                                                                }))
                                                                                                                                                            }))
                                                                                                         }
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Enum_enumEither() -> &dyn Any {
        static Data_Enum_enumEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumEither.get_or_init(||
                                             &Func1::new(move
                                                             |dictBoundedEnum|
                                                             {
                                                                 let Enum1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                             Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                 let Bounded0 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Bounded0"),
                                                                                                             Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                 let ordEither =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_ordEither(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictBoundedEnum)),
                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 &Func1::new({
                                                                                 let Bounded0
                                                                                     =
                                                                                     Bounded0.clone();
                                                                                 let Enum1
                                                                                     =
                                                                                     Enum1.clone();
                                                                                 let ordEither
                                                                                     =
                                                                                     ordEither.clone();
                                                                                 move
                                                                                     |dictBoundedEnum1|
                                                                                     {
                                                                                         let Bounded01 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Bounded0"),
                                                                                                                                     Sharpurs_Prelude::unbox(dictBoundedEnum1)),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         let Enum11 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                                                     Sharpurs_Prelude::unbox(dictBoundedEnum1)),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         let ordEither1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&ordEither,
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Enum1"),
                                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictBoundedEnum1)),
                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                                                                          &&&add(string("succ"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let Bounded01
                                                                                                                                                      =
                                                                                                                                                      Bounded01.clone();
                                                                                                                                                  let Enum11
                                                                                                                                                      =
                                                                                                                                                      Enum11.clone();
                                                                                                                                                  move
                                                                                                                                                      |v|
                                                                                                                                                      {
                                                                                                                                                          let matchValue:
                                                                                                                                                                  LrcPtr<Data_Either_Either> =
                                                                                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                                                                                          match matchValue.as_ref()
                                                                                                                                                              {
                                                                                                                                                              Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                              =>
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                     &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                          |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                          &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_2.clone())))),
                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                       |usd__arg1_3|
                                                                                                                                                                                                                                                                                       &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_3.clone()))))),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                                                                     &&&Enum11),
                                                                                                                                                                                                                                  &&matchValue_1_0)),
                                                                                                                                                              Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                                              =>
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                     &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                                                                                                                                                                                                                          &&&Bounded01)))))),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                          |usd__arg1|
                                                                                                                                                                                                                                                                                                                          &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                       |usd__arg1_1|
                                                                                                                                                                                                                                                                                       &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_1.clone()))))),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                                                                     &&&Enum1),
                                                                                                                                                                                                                                  &&matchValue_0_0)),
                                                                                                                                                          }
                                                                                                                                                      }
                                                                                                                                              }),
                                                                                                                                 add(string("pred"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let Enum11
                                                                                                                                                          =
                                                                                                                                                          Enum11.clone();
                                                                                                                                                      move
                                                                                                                                                          |v_1|
                                                                                                                                                          {
                                                                                                                                                              let matchValue_1:
                                                                                                                                                                      LrcPtr<Data_Either_Either> =
                                                                                                                                                                  Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                              match matchValue_1.as_ref()
                                                                                                                                                                  {
                                                                                                                                                                  Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                  =>
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                         &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                                                                                                                                                                                                                             &&&Bounded0)))))),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                              |usd__arg1_6|
                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_6.clone())))),
                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                           |usd__arg1_7|
                                                                                                                                                                                                                                                                                           &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_7.clone()))))),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                                                                                                         &&&Enum11),
                                                                                                                                                                                                                                      &&matchValue_1_1_0)),
                                                                                                                                                                  Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                  =>
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                         &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                              |usd__arg1_4|
                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_4.clone())))),
                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                           |usd__arg1_5|
                                                                                                                                                                                                                                                                                           &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_5.clone()))))),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                                                                                                         &&&Enum1),
                                                                                                                                                                                                                                      &&matchValue_1_0_0)),
                                                                                                                                                              }
                                                                                                                                                          }
                                                                                                                                                  }),
                                                                                                                                     add(string("Ord0"),
                                                                                                                                         &&Func1::new({
                                                                                                                                                          let ordEither1
                                                                                                                                                              =
                                                                                                                                                              ordEither1.clone();
                                                                                                                                                          move
                                                                                                                                                              |usd__unused|
                                                                                                                                                              &ordEither1
                                                                                                                                                      }),
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>()))))
                                                                                     }
                                                                             })
                                                             }))
    }
    pub fn Data_Enum_enumBoolean() -> &dyn Any {
        static Data_Enum_enumBoolean: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumBoolean.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                               &&&add(string("succ"),
                                                                                      &&Func1::new(move
                                                                                                       |v|
                                                                                                       {
                                                                                                           let matchValue =
                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                           match &Sharpurs_Prelude::_007cLitBool_007c__007c(false,
                                                                                                                                                            &matchValue)
                                                                                                               {
                                                                                                               0_i32
                                                                                                               =>
                                                                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&true)),
                                                                                                               _
                                                                                                               =>
                                                                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                           }
                                                                                                       }),
                                                                                      add(string("pred"),
                                                                                          &&Func1::new(move
                                                                                                           |v_1|
                                                                                                           {
                                                                                                               let matchValue_1 =
                                                                                                                   Sharpurs_Prelude::unbox(v_1);
                                                                                                               match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                &matchValue_1)
                                                                                                                   {
                                                                                                                   0_i32
                                                                                                                   =>
                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&false)),
                                                                                                                   _
                                                                                                                   =>
                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                               }
                                                                                                           }),
                                                                                          add(string("Ord0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Ord::Data_Ord_ordBoolean()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))))
    }
    pub fn Data_Enum_downFromIncluding() -> &dyn Any {
        static Data_Enum_downFromIncluding: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_downFromIncluding.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictEnum|
                                                                    {
                                                                        let pred1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                             dictEnum);
                                                                        &Func1::new({
                                                                                        let pred1
                                                                                            =
                                                                                            pred1.clone();
                                                                                        move
                                                                                            |dictUnfoldable1|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                dictUnfoldable1),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Apply::Control_Apply_applyFn()),
                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                     |usd__arg1|
                                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                                    let usd__arg1
                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                        usd__arg1.clone();
                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                        |usd__arg2|
                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                usd__arg2.clone()))
                                                                                                                                                                                                                                }))),
                                                                                                                                                                &&&pred1))
                                                                                    })
                                                                    }))
    }
    pub fn Data_Enum_diag() -> &dyn Any {
        static Data_Enum_diag: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Enum_diag.get_or_init(||
                                       &Func1::new(move |a|
                                                       &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                               a.clone()))))
    }
    pub fn Data_Enum_downFrom() -> &dyn Any {
        static Data_Enum_downFrom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_downFrom.get_or_init(||
                                           &Func1::new(move |dictEnum|
                                                           {
                                                               let pred1 =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                    dictEnum);
                                                               &Func1::new({
                                                                               let pred1
                                                                                   =
                                                                                   pred1.clone();
                                                                               move
                                                                                   |dictUnfoldable|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                       dictUnfoldable),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                             &&&PureScript_Data_Enum::Data_Enum_diag())),
                                                                                                                                                       &&&pred1))
                                                                           })
                                                           }))
    }
    pub fn Data_Enum_upFrom() -> &dyn Any {
        static Data_Enum_upFrom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_upFrom.get_or_init(||
                                         &Func1::new(move |dictEnum|
                                                         {
                                                             let succ1 =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                  dictEnum);
                                                             &Func1::new({
                                                                             let succ1
                                                                                 =
                                                                                 succ1.clone();
                                                                             move
                                                                                 |dictUnfoldable|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                     dictUnfoldable),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                           &&&PureScript_Data_Enum::Data_Enum_diag())),
                                                                                                                                                     &&&succ1))
                                                                         })
                                                         }))
    }
    pub fn Data_Enum_defaultToEnum() -> &dyn Any {
        static Data_Enum_defaultToEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_defaultToEnum.get_or_init(||
                                                &Func1::new(move |dictBounded|
                                                                &Func1::new({
                                                                                let dictBounded
                                                                                    =
                                                                                    dictBounded.clone();
                                                                                move
                                                                                    |dictEnum|
                                                                                    &Func1::new({
                                                                                                    let dictEnum
                                                                                                        =
                                                                                                        dictEnum.clone();
                                                                                                    move
                                                                                                        |i_prime|
                                                                                                        {
                                                                                                            let go_2 =
                                                                                                                Func0::new({
                                                                                                                               let go_tco
                                                                                                                                   =
                                                                                                                                   go_tco.clone();
                                                                                                                               move
                                                                                                                                   ||
                                                                                                                                   &Func1::new({
                                                                                                                                                   let go_tco
                                                                                                                                                       =
                                                                                                                                                       go_tco.clone();
                                                                                                                                                   move
                                                                                                                                                       |i|
                                                                                                                                                       Func1::new({
                                                                                                                                                                      let go_tco
                                                                                                                                                                          =
                                                                                                                                                                          go_tco.clone();
                                                                                                                                                                      let i
                                                                                                                                                                          =
                                                                                                                                                                          i.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |x|
                                                                                                                                                                          go_tco(i)(x.clone())
                                                                                                                                                                  })
                                                                                                                                               })
                                                                                                                           });
                                                                                                            let go_1 =
                                                                                                                Lazy(go_2);
                                                                                                            fn go_tco(i_1:
                                                                                                                          _)
                                                                                                             ->
                                                                                                                 Func1<&dyn Any,
                                                                                                                       &dyn Any> {
                                                                                                                Func1::new({
                                                                                                                               let go_tco
                                                                                                                                   =
                                                                                                                                   go_tco.clone();
                                                                                                                               let i_1
                                                                                                                                   =
                                                                                                                                   i_1.clone();
                                                                                                                               move
                                                                                                                                   |x_1|
                                                                                                                                   {
                                                                                                                                       let matchValue =
                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                           &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                        &&&i_1),
                                                                                                                                                                                                     &&&0_i32));
                                                                                                                                       match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                        &matchValue)
                                                                                                                                           {
                                                                                                                                           0_i32
                                                                                                                                           =>
                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x_1.clone())),
                                                                                                                                           _
                                                                                                                                           =>
                                                                                                                                           {
                                                                                                                                               let matchValue_1:
                                                                                                                                                       LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                                                 &&&dictEnum),
                                                                                                                                                                                                              x_1));
                                                                                                                                               match matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                   =>
                                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                                   =>
                                                                                                                                                   go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                              &&&i_1),
                                                                                                                                                                                           &&&1_i32))(&match matchValue_1.as_ref()
                                                                                                                                                                                                           {
                                                                                                                                                                                                           Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                           _
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                       }),
                                                                                                                                               }
                                                                                                                                           }
                                                                                                                                       }
                                                                                                                                   }
                                                                                                                           })
                                                                                                            }
                                                                                                            let go =
                                                                                                                go_1.Value;
                                                                                                            {
                                                                                                                let matchValue_2 =
                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                    &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                 i_prime),
                                                                                                                                                                              &&&0_i32));
                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                 &matchValue_2)
                                                                                                                    {
                                                                                                                    0_i32
                                                                                                                    =>
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    go_tco(i_prime.clone())(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                             &&&dictBounded)),
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                })
                                                                            })))
    }
    pub fn Data_Enum_defaultSucc() -> &dyn Any {
        static Data_Enum_defaultSucc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_defaultSucc.get_or_init(||
                                              &Func1::new(move |toEnum_prime|
                                                              &Func1::new({
                                                                              let toEnum_prime
                                                                                  =
                                                                                  toEnum_prime.clone();
                                                                              move
                                                                                  |fromEnum_prime|
                                                                                  &Func1::new({
                                                                                                  let fromEnum_prime
                                                                                                      =
                                                                                                      fromEnum_prime.clone();
                                                                                                  move
                                                                                                      |a|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&toEnum_prime,
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&fromEnum_prime,
                                                                                                                                                                                                                                                a)),
                                                                                                                                                                          &&&1_i32))
                                                                                              })
                                                                          })))
    }
    pub fn Data_Enum_defaultPred() -> &dyn Any {
        static Data_Enum_defaultPred: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_defaultPred.get_or_init(||
                                              &Func1::new(move |toEnum_prime|
                                                              &Func1::new({
                                                                              let toEnum_prime
                                                                                  =
                                                                                  toEnum_prime.clone();
                                                                              move
                                                                                  |fromEnum_prime|
                                                                                  &Func1::new({
                                                                                                  let fromEnum_prime
                                                                                                      =
                                                                                                      fromEnum_prime.clone();
                                                                                                  move
                                                                                                      |a|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&toEnum_prime,
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&fromEnum_prime,
                                                                                                                                                                                                                                                a)),
                                                                                                                                                                          &&&1_i32))
                                                                                              })
                                                                          })))
    }
    pub fn Data_Enum_defaultFromEnum() -> &dyn Any {
        static Data_Enum_defaultFromEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_defaultFromEnum.get_or_init(||
                                                  &Func1::new(move |dictEnum|
                                                                  {
                                                                      let go_2 =
                                                                          Func0::new({
                                                                                         let go_tco
                                                                                             =
                                                                                             go_tco.clone();
                                                                                         move
                                                                                             ||
                                                                                             &Func1::new({
                                                                                                             let go_tco
                                                                                                                 =
                                                                                                                 go_tco.clone();
                                                                                                             move
                                                                                                                 |i|
                                                                                                                 Func1::new({
                                                                                                                                let go_tco
                                                                                                                                    =
                                                                                                                                    go_tco.clone();
                                                                                                                                let i
                                                                                                                                    =
                                                                                                                                    i.clone();
                                                                                                                                move
                                                                                                                                    |x|
                                                                                                                                    go_tco(i)(x.clone())
                                                                                                                            })
                                                                                                         })
                                                                                     });
                                                                      let go_1 =
                                                                          Lazy(go_2);
                                                                      let go_tco =
                                                                          Func1::new({
                                                                                         let dictEnum
                                                                                             =
                                                                                             dictEnum.clone();
                                                                                         move
                                                                                             |i_1|
                                                                                             fix1(&(move
                                                                                                        |go_tco,
                                                                                                         i_1|
                                                                                                        Func1::new({
                                                                                                                       let go_tco
                                                                                                                           =
                                                                                                                           go_tco.clone();
                                                                                                                       let i_1
                                                                                                                           =
                                                                                                                           i_1.clone();
                                                                                                                       move
                                                                                                                           |x_1|
                                                                                                                           {
                                                                                                                               let matchValue:
                                                                                                                                       LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                                                                 &&&dictEnum),
                                                                                                                                                                                              x_1));
                                                                                                                               match matchValue.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                   =>
                                                                                                                                   &i_1,
                                                                                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                   =>
                                                                                                                                   go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                              &&&i_1),
                                                                                                                                                                           &&&1_i32))(&match matchValue.as_ref()
                                                                                                                                                                                           {
                                                                                                                                                                                           Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                           =>
                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                           _
                                                                                                                                                                                           =>
                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                       }),
                                                                                                                               }
                                                                                                                           }
                                                                                                                   })),
                                                                                                  i_1.clone())
                                                                                     });
                                                                      let go =
                                                                          go_1.Value;
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                       &&&0_i32)
                                                                  }))
    }
    pub fn Data_Enum_defaultCardinality() -> &dyn Any {
        static Data_Enum_defaultCardinality: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_defaultCardinality.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBounded|
                                                                     {
                                                                         let bottom2 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                              dictBounded);
                                                                         &Func1::new({
                                                                                         let bottom2
                                                                                             =
                                                                                             bottom2.clone();
                                                                                         move
                                                                                             |dictEnum|
                                                                                             {
                                                                                                 let go_2 =
                                                                                                     Func0::new({
                                                                                                                    let go_tco
                                                                                                                        =
                                                                                                                        go_tco.clone();
                                                                                                                    move
                                                                                                                        ||
                                                                                                                        &Func1::new({
                                                                                                                                        let go_tco
                                                                                                                                            =
                                                                                                                                            go_tco.clone();
                                                                                                                                        move
                                                                                                                                            |i|
                                                                                                                                            Func1::new({
                                                                                                                                                           let go_tco
                                                                                                                                                               =
                                                                                                                                                               go_tco.clone();
                                                                                                                                                           let i
                                                                                                                                                               =
                                                                                                                                                               i.clone();
                                                                                                                                                           move
                                                                                                                                                               |x|
                                                                                                                                                               go_tco(i)(x.clone())
                                                                                                                                                       })
                                                                                                                                    })
                                                                                                                });
                                                                                                 let go_1 =
                                                                                                     Lazy(go_2);
                                                                                                 let go_tco =
                                                                                                     Func1::new({
                                                                                                                    let dictEnum
                                                                                                                        =
                                                                                                                        dictEnum.clone();
                                                                                                                    move
                                                                                                                        |i_1|
                                                                                                                        fix1(&(move
                                                                                                                                   |go_tco,
                                                                                                                                    i_1|
                                                                                                                                   Func1::new({
                                                                                                                                                  let go_tco
                                                                                                                                                      =
                                                                                                                                                      go_tco.clone();
                                                                                                                                                  let i_1
                                                                                                                                                      =
                                                                                                                                                      i_1.clone();
                                                                                                                                                  move
                                                                                                                                                      |x_1|
                                                                                                                                                      {
                                                                                                                                                          let matchValue:
                                                                                                                                                                  LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                                                            &&&dictEnum),
                                                                                                                                                                                                                         x_1));
                                                                                                                                                          match matchValue.as_ref()
                                                                                                                                                              {
                                                                                                                                                              Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                              =>
                                                                                                                                                              &i_1,
                                                                                                                                                              Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                              =>
                                                                                                                                                              go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                         &&&i_1),
                                                                                                                                                                                                      &&&1_i32))(&match matchValue.as_ref()
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                      Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                      _
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                  }),
                                                                                                                                                          }
                                                                                                                                                      }
                                                                                                                                              })),
                                                                                                                             i_1.clone())
                                                                                                                });
                                                                                                 let go =
                                                                                                     go_1.Value;
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                     &&&PureScript_Data_Enum::Data_Enum_Cardinality()),
                                                                                                                                  &&go_tco(&1_i32)(&bottom2))
                                                                                             }
                                                                                     })
                                                                     }))
    }
    pub fn Data_Enum_charToEnum() -> &dyn Any {
        static Data_Enum_charToEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_charToEnum.get_or_init(||
                                             &Func1::new(move |v|
                                                             {
                                                                 let matchValue =
                                                                     Sharpurs_Prelude::unbox(v);
                                                                 if {
                                                                        let n =
                                                                            matchValue;
                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                        &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                                              &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                           &&&n),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toCharCode(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Enum::Data_Enum_bottom1()))),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                        &&&n),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toCharCode(),
                                                                                                                                                                                                        &&&PureScript_Data_Enum::Data_Enum_top()))))
                                                                    } {
                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromCharCode(),
                                                                                                                                                             &&&matchValue)))
                                                                 } else {
                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                 }
                                                             }))
    }
    pub fn Data_Enum_enumChar() -> &dyn Any {
        static Data_Enum_enumChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_enumChar.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                            &&&add(string("succ"),
                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_defaultSucc(),
                                                                                                                                                        &&&PureScript_Data_Enum::Data_Enum_charToEnum()),
                                                                                                                     &&&PureScript_Data_Enum::Data_Enum_toCharCode()),
                                                                                   add(string("pred"),
                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_defaultPred(),
                                                                                                                                                            &&&PureScript_Data_Enum::Data_Enum_charToEnum()),
                                                                                                                         &&&PureScript_Data_Enum::Data_Enum_toCharCode()),
                                                                                       add(string("Ord0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Ord::Data_Ord_ordChar()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>())))))
    }
    pub fn Data_Enum_cardinality() -> &dyn Any {
        static Data_Enum_cardinality: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_cardinality.get_or_init(||
                                              &Func1::new(move |dict|
                                                              find(string("cardinality"),
                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Enum_boundedEnumUnit() -> &dyn Any {
        static Data_Enum_boundedEnumUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_boundedEnumUnit.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                                                                   &&&add(string("cardinality"),
                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                                                            &&&1_i32),
                                                                                          add(string("toEnum"),
                                                                                              &&Func1::new(move
                                                                                                               |v|
                                                                                                               {
                                                                                                                   let matchValue =
                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                   match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                                                   &matchValue)
                                                                                                                       {
                                                                                                                       0_i32
                                                                                                                       =>
                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                       _
                                                                                                                       =>
                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                   }
                                                                                                               }),
                                                                                              add(string("fromEnum"),
                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                    &&&0_i32),
                                                                                                  add(string("Bounded0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_Bounded::Data_Bounded_boundedUnit()),
                                                                                                      add(string("Enum1"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused_1|
                                                                                                                           &PureScript_Data_Enum::Data_Enum_enumUnit()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))))))
    }
    pub fn Data_Enum_boundedEnumOrdering() -> &dyn Any {
        static Data_Enum_boundedEnumOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_boundedEnumOrdering.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                                                                       &&&add(string("cardinality"),
                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                                                                &&&3_i32),
                                                                                              add(string("toEnum"),
                                                                                                  &&Func1::new(move
                                                                                                                   |v|
                                                                                                                   {
                                                                                                                       let matchValue =
                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                       match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                                                       &matchValue)
                                                                                                                           {
                                                                                                                           0_i32
                                                                                                                           =>
                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor))),
                                                                                                                           _
                                                                                                                           =>
                                                                                                                           match &Sharpurs_Prelude::_007cLitInt_007c__007c(1_i32,
                                                                                                                                                                           &matchValue)
                                                                                                                               {
                                                                                                                               0_i32
                                                                                                                               =>
                                                                                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                                                               _
                                                                                                                               =>
                                                                                                                               match &Sharpurs_Prelude::_007cLitInt_007c__007c(2_i32,
                                                                                                                                                                               &matchValue)
                                                                                                                                   {
                                                                                                                                   0_i32
                                                                                                                                   =>
                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor))),
                                                                                                                                   _
                                                                                                                                   =>
                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                               },
                                                                                                                           },
                                                                                                                       }
                                                                                                                   }),
                                                                                                  add(string("fromEnum"),
                                                                                                      &&Func1::new(move
                                                                                                                       |v_1|
                                                                                                                       {
                                                                                                                           let matchValue_1:
                                                                                                                                   LrcPtr<Data_Ordering_Ordering> =
                                                                                                                               Sharpurs_Prelude::unbox(v_1);
                                                                                                                           match matchValue_1.as_ref()
                                                                                                                               {
                                                                                                                               Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                               =>
                                                                                                                               &1_i32,
                                                                                                                               Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                               =>
                                                                                                                               &2_i32,
                                                                                                                               _
                                                                                                                               =>
                                                                                                                               &0_i32,
                                                                                                                           }
                                                                                                                       }),
                                                                                                      add(string("Bounded0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Bounded::Data_Bounded_boundedOrdering()),
                                                                                                          add(string("Enum1"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused_1|
                                                                                                                               &PureScript_Data_Enum::Data_Enum_enumOrdering()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))))))
    }
    pub fn Data_Enum_boundedEnumChar() -> &dyn Any {
        static Data_Enum_boundedEnumChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_boundedEnumChar.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                                                                   &&&add(string("cardinality"),
                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toCharCode(),
                                                                                                                                                                                                                                     &&&PureScript_Data_Enum::Data_Enum_top())),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toCharCode(),
                                                                                                                                                                                                  &&&PureScript_Data_Enum::Data_Enum_bottom1()))),
                                                                                          add(string("toEnum"),
                                                                                              &&PureScript_Data_Enum::Data_Enum_charToEnum(),
                                                                                              add(string("fromEnum"),
                                                                                                  &&PureScript_Data_Enum::Data_Enum_toCharCode(),
                                                                                                  add(string("Bounded0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_Bounded::Data_Bounded_boundedChar()),
                                                                                                      add(string("Enum1"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused_1|
                                                                                                                           &PureScript_Data_Enum::Data_Enum_enumChar()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))))))
    }
    pub fn Data_Enum_boundedEnumBoolean() -> &dyn Any {
        static Data_Enum_boundedEnumBoolean: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Enum_boundedEnumBoolean.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                                                                      &&&add(string("cardinality"),
                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                                                               &&&2_i32),
                                                                                             add(string("toEnum"),
                                                                                                 &&Func1::new(move
                                                                                                                  |v|
                                                                                                                  {
                                                                                                                      let matchValue =
                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                      match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                                                      &matchValue)
                                                                                                                          {
                                                                                                                          0_i32
                                                                                                                          =>
                                                                                                                          &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&false)),
                                                                                                                          _
                                                                                                                          =>
                                                                                                                          match &Sharpurs_Prelude::_007cLitInt_007c__007c(1_i32,
                                                                                                                                                                          &matchValue)
                                                                                                                              {
                                                                                                                              0_i32
                                                                                                                              =>
                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&true)),
                                                                                                                              _
                                                                                                                              =>
                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                          },
                                                                                                                      }
                                                                                                                  }),
                                                                                                 add(string("fromEnum"),
                                                                                                     &&Func1::new(move
                                                                                                                      |v_1|
                                                                                                                      {
                                                                                                                          let matchValue_1 =
                                                                                                                              Sharpurs_Prelude::unbox(v_1);
                                                                                                                          match &Sharpurs_Prelude::_007cLitBool_007c__007c(false,
                                                                                                                                                                           &matchValue_1)
                                                                                                                              {
                                                                                                                              0_i32
                                                                                                                              =>
                                                                                                                              &0_i32,
                                                                                                                              _
                                                                                                                              =>
                                                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                               &matchValue_1)
                                                                                                                                  {
                                                                                                                                  0_i32
                                                                                                                                  =>
                                                                                                                                  &1_i32,
                                                                                                                                  _
                                                                                                                                  =>
                                                                                                                                  panic!("{}",
                                                                                                                                         LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Enum.fs"),
                                  Data1: 111_i32,
                                  Data2: 515_i32,}).get_Message(),),
                                                                                                                              },
                                                                                                                          }
                                                                                                                      }),
                                                                                                     add(string("Bounded0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Bounded::Data_Bounded_boundedBoolean()),
                                                                                                         add(string("Enum1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused_1|
                                                                                                                              &PureScript_Data_Enum::Data_Enum_enumBoolean()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))))))
    }
}
