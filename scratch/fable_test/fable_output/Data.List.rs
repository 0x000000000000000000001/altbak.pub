pub mod PureScript_Data_List {
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
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_4df90a9e::PureScript_Data_List_Internal;
    use crate::module_d662adf2::PureScript_Data_List_Types;
    use crate::module_d662adf2::PureScript_Data_List_Types::Data_List_Types_List;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_List_identity() -> &dyn Any {
        static Data_List_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_identity.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                            &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_List_Pattern() -> &dyn Any {
        static Data_List_Pattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Pattern.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_List_updateAt_004012() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           Func1::new({
                                                          let v1 = v1.clone();
                                                          move |v2|
                                                              PureScript_Data_List::Data_List_updateAt_tco(&v,
                                                                                                           &v1,
                                                                                                           v2)
                                                      })
                                   }))
    }
    pub fn Data_List_updateAt_004012_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_updateAt_004012_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_updateAt_004012_002d1.get_or_init(||
                                                        Lazy(Data_List_updateAt_004012.clone()))
    }
    pub fn Data_List_updateAt_tco(v: &dyn Any, v1: &dyn Any, v2: &dyn Any)
     -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        let matchValue_2: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v2);
        if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                    &matchValue).is_some() {
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                      matchValue_2_1_1)
                   = matchValue_2.as_ref() {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue_1,
                                                                                                                                       &match matchValue_2.as_ref()
                                                                                                                                            {
                                                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                               x)
                                                                                                                                            =>
                                                                                                                                            x.clone(),
                                                                                                                                            _
                                                                                                                                            =>
                                                                                                                                            unreachable!(),
                                                                                                                                        }))))
            } else {
                if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                          matchValue_2_1_1)
                       = matchValue_2.as_ref() {
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                        &&&Func1::new({
                                                                                                          let matchValue_2
                                                                                                              =
                                                                                                              matchValue_2.clone();
                                                                                                          move
                                                                                                              |v3|
                                                                                                              &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                      _)
                                                                                                                                                                                   =>
                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                   _
                                                                                                                                                                                   =>
                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                               },
                                                                                                                                                                              v3.clone()))
                                                                                                      })),
                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_updateAt_004012_002d1.Value,
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                       &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                 &&&1_i32)),
                                                                                                                           &&&matchValue_1),
                                                                                        &&&match matchValue_2.as_ref()
                                                                                               {
                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                  x)
                                                                                               =>
                                                                                               x.clone(),
                                                                                               _
                                                                                               =>
                                                                                               unreachable!(),
                                                                                           }))
                } else {
                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                }
            }
        } else {
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                      matchValue_2_1_1)
                   = matchValue_2.as_ref() {
                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                    &&&Func1::new({
                                                                                                      let matchValue_2
                                                                                                          =
                                                                                                          matchValue_2.clone();
                                                                                                      move
                                                                                                          |v3|
                                                                                                          &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                  _)
                                                                                                                                                                               =>
                                                                                                                                                                               x.clone(),
                                                                                                                                                                               _
                                                                                                                                                                               =>
                                                                                                                                                                               unreachable!(),
                                                                                                                                                                           },
                                                                                                                                                                          v3.clone()))
                                                                                                  })),
                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_updateAt_004012_002d1.Value,
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                             &&&1_i32)),
                                                                                                                       &&&matchValue_1),
                                                                                    &&&match matchValue_2.as_ref()
                                                                                           {
                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                              x)
                                                                                           =>
                                                                                           x.clone(),
                                                                                           _
                                                                                           =>
                                                                                           unreachable!(),
                                                                                       }))
            } else {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
            }
        }
    }
    pub fn Data_List_updateAt() -> &dyn Any {
        static Data_List_updateAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_updateAt.get_or_init(||
                                           Data_List_updateAt_004012_002d1.Value)
    }
    pub fn Data_List_unzip() -> &dyn Any {
        static Data_List_unzip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_unzip.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                               &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                            &&&Func1::new(move
                                                                                                                              |v|
                                                                                                                              {
                                                                                                                                  let matchValue:
                                                                                                                                          LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                  &Func1::new(move
                                                                                                                                                  |v1|
                                                                                                                                                  {
                                                                                                                                                      let matchValue_1:
                                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               },
                                                                                                                                                                                                                                                                              &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                              &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               },
                                                                                                                                                                                                                                                                              &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               }))))
                                                                                                                                                  })
                                                                                                                              })),
                                                                         &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
                                                                                                                                   &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)))))
    }
    pub fn Data_List_uncons() -> &dyn Any {
        static Data_List_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_uncons.get_or_init(||
                                         &Func1::new(move |v|
                                                         {
                                                             let matchValue:
                                                                     LrcPtr<Data_List_Types_List> =
                                                                 Sharpurs_Prelude::unbox(v);
                                                             match matchValue.as_ref()
                                                                 {
                                                                 Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                    matchValue_1_1)
                                                                 =>
                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&add(string("head"),
                                                                                                                             &matchValue_1_0,
                                                                                                                             add(string("tail"),
                                                                                                                                 &matchValue_1_1,
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))),
                                                                 _ =>
                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                             }
                                                         }))
    }
    pub fn Data_List_toUnfoldable() -> &dyn Any {
        static Data_List_toUnfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_toUnfoldable.get_or_init(||
                                               &Func1::new(move
                                                               |dictUnfoldable|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                   dictUnfoldable),
                                                                                                &&&Func1::new(move
                                                                                                                  |xs|
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                         &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                        |rec_var|
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(find(string("head"),
                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(rec_var)),
                                                                                                                                                                                                                                                                find(string("tail"),
                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(rec_var)))))),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_uncons(),
                                                                                                                                                                                      xs))))))
    }
    pub fn Data_List_tail() -> &dyn Any {
        static Data_List_tail: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_tail.get_or_init(||
                                       &Func1::new(move |v|
                                                       {
                                                           let matchValue:
                                                                   LrcPtr<Data_List_Types_List> =
                                                               Sharpurs_Prelude::unbox(v);
                                                           match matchValue.as_ref()
                                                               {
                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                  matchValue_1_1)
                                                               =>
                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1)),
                                                               _ =>
                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                           }
                                                       }))
    }
    pub fn Data_List_stripPrefix() -> &dyn Any {
        static Data_List_stripPrefix: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_stripPrefix.get_or_init(||
                                              &Func1::new(move |dictEq|
                                                              &Func1::new({
                                                                              let dictEq
                                                                                  =
                                                                                  dictEq.clone();
                                                                              move
                                                                                  |v|
                                                                                  &Func1::new({
                                                                                                  let v
                                                                                                      =
                                                                                                      v.clone();
                                                                                                  move
                                                                                                      |s|
                                                                                                      {
                                                                                                          let matchValue =
                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                          let matchValue_1 =
                                                                                                              Sharpurs_Prelude::unbox(s);
                                                                                                          let go =
                                                                                                              &Func1::new(move
                                                                                                                              |prefix|
                                                                                                                              &Func1::new({
                                                                                                                                              let prefix
                                                                                                                                                  =
                                                                                                                                                  prefix.clone();
                                                                                                                                              move
                                                                                                                                                  |input|
                                                                                                                                                  {
                                                                                                                                                      let matchValue_3:
                                                                                                                                                              LrcPtr<Data_List_Types_List> =
                                                                                                                                                          Sharpurs_Prelude::unbox(&&prefix);
                                                                                                                                                      let matchValue_4:
                                                                                                                                                              LrcPtr<Data_List_Types_List> =
                                                                                                                                                          Sharpurs_Prelude::unbox(input);
                                                                                                                                                      if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                                                             =
                                                                                                                                                             matchValue_3.as_ref()
                                                                                                                                                         {
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                |usd__arg1_1|
                                                                                                                                                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                           &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&matchValue_4)))
                                                                                                                                                      } else {
                                                                                                                                                          if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_4_1_0,
                                                                                                                                                                                                                    matchValue_4_1_1)
                                                                                                                                                                 =
                                                                                                                                                                 matchValue_4.as_ref()
                                                                                                                                                             {
                                                                                                                                                              if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                 &&&dictEq),
                                                                                                                                                                                                                                                              &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                           &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                  Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                              }))
                                                                                                                                                                 {
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                   &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&add(string("a"),
                                                                                                                                                                                                                                                                                          &&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                                                                            },
                                                                                                                                                                                                                                                                                          add(string("b"),
                                                                                                                                                                                                                                                                                              &&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                    Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                                                },
                                                                                                                                                                                                                                                                                              empty::<string,
                                                                                                                                                                                                                                                                                                      &dyn Any>())))))
                                                                                                                                                              } else {
                                                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                                                              }
                                                                                                                                                          } else {
                                                                                                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                                                          }
                                                                                                                                                      }
                                                                                                                                                  }
                                                                                                                                          }));
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM2(),
                                                                                                                                                                                                                                                    &&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_monadRecMaybe()),
                                                                                                                                                                                                                 &&&go),
                                                                                                                                                                              &&&matchValue),
                                                                                                                                           &&&matchValue_1)
                                                                                                      }
                                                                                              })
                                                                          })))
    }
    pub fn Data_List_span_004026() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Data_List::Data_List_span_tco(&v,
                                                                                    v1)
                                   }))
    }
    pub fn Data_List_span_004026_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_span_004026_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_span_004026_002d1.get_or_init(||
                                                    Lazy(Data_List_span_004026.clone()))
    }
    pub fn Data_List_span_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v1);
        if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                  matchValue_1_1_1)
               = matchValue_1.as_ref() {
            if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                         &&&match matchValue_1.as_ref()
                                                                                {
                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                   _)
                                                                                =>
                                                                                x.clone(),
                                                                                _
                                                                                =>
                                                                                unreachable!(),
                                                                            }))
               {
                let matchValue_3 =
                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_span_004026_002d1.Value,
                                                                                                                  &&&matchValue),
                                                                               &&&match matchValue_1.as_ref()
                                                                                      {
                                                                                      Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                         x)
                                                                                      =>
                                                                                      x.clone(),
                                                                                      _
                                                                                      =>
                                                                                      unreachable!(),
                                                                                  }));
                {
                    let activePatternResult =
                        Sharpurs_Prelude::_007cHasProp_007c__007c(string("init"),
                                                                  &matchValue_3);
                    if activePatternResult.is_some() {
                        let activePatternResult_1 =
                            Sharpurs_Prelude::_007cHasProp_007c__007c(string("rest"),
                                                                      &matchValue_3);
                        if activePatternResult_1.is_some() {
                            let ys = getValue(activePatternResult);
                            let zs = getValue(activePatternResult_1);
                            &add(string("init"),
                                 &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                       {
                                                                                                       Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                          _)
                                                                                                       =>
                                                                                                       x.clone(),
                                                                                                       _
                                                                                                       =>
                                                                                                       unreachable!(),
                                                                                                   },
                                                                                                  &ys)),
                                 add(string("rest"), &&zs,
                                     empty::<string, &dyn Any>()))
                        } else {
                            panic!("{}",
                                   LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 25_i32,
                                  Data2: 354_i32,}).get_Message(),)
                        }
                    } else {
                        panic!("{}",
                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 25_i32,
                                  Data2: 354_i32,}).get_Message(),)
                    }
                }
            } else {
                &add(string("init"),
                     &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
                     add(string("rest"), &&matchValue_1,
                         empty::<string, &dyn Any>()))
            }
        } else {
            &add(string("init"),
                 &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
                 add(string("rest"), &&matchValue_1,
                     empty::<string, &dyn Any>()))
        }
    }
    pub fn Data_List_span() -> &dyn Any {
        static Data_List_span: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_span.get_or_init(|| Data_List_span_004026_002d1.Value)
    }
    pub fn Data_List_snoc() -> &dyn Any {
        static Data_List_snoc: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_snoc.get_or_init(||
                                       &Func1::new(move |xs|
                                                       &Func1::new({
                                                                       let xs
                                                                           =
                                                                           xs.clone();
                                                                       move
                                                                           |x|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                     &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                    |usd__arg1|
                                                                                                                                                                                                    Func1::new({
                                                                                                                                                                                                                   let usd__arg1
                                                                                                                                                                                                                       =
                                                                                                                                                                                                                       usd__arg1.clone();
                                                                                                                                                                                                                   move
                                                                                                                                                                                                                       |usd__arg2|
                                                                                                                                                                                                                       &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                       usd__arg2.clone()))
                                                                                                                                                                                                               }))),
                                                                                                                                               &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(x.clone(),
                                                                                                                                                                                                                 &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)))),
                                                                                                            &&&xs)
                                                                   })))
    }
    pub fn Data_List_singleton() -> &dyn Any {
        static Data_List_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_singleton.get_or_init(||
                                            &Func1::new(move |a|
                                                            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(a.clone(),
                                                                                                                            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)))))
    }
    pub fn Data_List_sortBy() -> &dyn Any {
        static Data_List_sortBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_sortBy.get_or_init(||
                                         &Func1::new(move |cmp|
                                                         {
                                                             let merge_2 =
                                                                 Func0::new({
                                                                                let merge_tco
                                                                                    =
                                                                                    merge_tco.clone();
                                                                                move
                                                                                    ||
                                                                                    &Func1::new({
                                                                                                    let merge_tco
                                                                                                        =
                                                                                                        merge_tco.clone();
                                                                                                    move
                                                                                                        |v|
                                                                                                        Func1::new({
                                                                                                                       let merge_tco
                                                                                                                           =
                                                                                                                           merge_tco.clone();
                                                                                                                       let v
                                                                                                                           =
                                                                                                                           v.clone();
                                                                                                                       move
                                                                                                                           |v1|
                                                                                                                           merge_tco(v)(v1.clone())
                                                                                                                   })
                                                                                                })
                                                                            });
                                                             let merge_1 =
                                                                 Lazy(merge_2);
                                                             let merge_tco =
                                                                 Func1::new({
                                                                                let cmp
                                                                                    =
                                                                                    cmp.clone();
                                                                                move
                                                                                    |v_1|
                                                                                    fix1(&(move
                                                                                               |merge_tco,
                                                                                                v_1|
                                                                                               Func1::new({
                                                                                                              let merge_tco
                                                                                                                  =
                                                                                                                  merge_tco.clone();
                                                                                                              let v_1
                                                                                                                  =
                                                                                                                  v_1.clone();
                                                                                                              move
                                                                                                                  |v1_1|
                                                                                                                  {
                                                                                                                      let matchValue:
                                                                                                                              LrcPtr<Data_List_Types_List> =
                                                                                                                          Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                      let matchValue_1:
                                                                                                                              LrcPtr<Data_List_Types_List> =
                                                                                                                          Sharpurs_Prelude::unbox(v1_1);
                                                                                                                      if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                             =
                                                                                                                             matchValue.as_ref()
                                                                                                                         {
                                                                                                                          &matchValue_1
                                                                                                                      } else {
                                                                                                                          if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                                 =
                                                                                                                                 matchValue_1.as_ref()
                                                                                                                             {
                                                                                                                              &matchValue
                                                                                                                          } else {
                                                                                                                              if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                    })),
                                                                                                                                                                                           &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)))
                                                                                                                                 {
                                                                                                                                  &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                       {
                                                                                                                                                                                                       Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                       =>
                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                       _
                                                                                                                                                                                                       =>
                                                                                                                                                                                                       unreachable!(),
                                                                                                                                                                                                   },
                                                                                                                                                                                                  merge_tco(&matchValue)(&match matchValue_1.as_ref()
                                                                                                                                                                                                                              {
                                                                                                                                                                                                                              Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                              _
                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                                          })))
                                                                                                                              } else {
                                                                                                                                  if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                     {
                                                                                                                                      &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                           {
                                                                                                                                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                           _
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                       },
                                                                                                                                                                                                      merge_tco(&match matchValue.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                     Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                     _
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                 })(&matchValue_1)))
                                                                                                                                  } else {
                                                                                                                                      panic!("{}",
                                                                                                                                             string("Match failure: PureScript_Data_List_Types.Data_List_Types_List"),)
                                                                                                                                  }
                                                                                                                              }
                                                                                                                          }
                                                                                                                      }
                                                                                                                  }
                                                                                                          })),
                                                                                         v_1.clone())
                                                                            });
                                                             let merge =
                                                                 merge_1.Value;
                                                             {
                                                                 let mergePairs_2 =
                                                                     Func0::new({
                                                                                    let mergePairs_tco
                                                                                        =
                                                                                        mergePairs_tco.clone();
                                                                                    move
                                                                                        ||
                                                                                        &Func1::new({
                                                                                                        let mergePairs_tco
                                                                                                            =
                                                                                                            mergePairs_tco.clone();
                                                                                                        move
                                                                                                            |v_2|
                                                                                                            mergePairs_tco(v_2.clone())
                                                                                                    })
                                                                                });
                                                                 let mergePairs_1 =
                                                                     Lazy(mergePairs_2);
                                                                 let mergePairs_tco =
                                                                     Func1::new({
                                                                                    let merge_tco
                                                                                        =
                                                                                        merge_tco.clone();
                                                                                    move
                                                                                        |v_3|
                                                                                        fix1(&(move
                                                                                                   |mergePairs_tco,
                                                                                                    v_3|
                                                                                                   {
                                                                                                       let matchValue_3:
                                                                                                               LrcPtr<Data_List_Types_List> =
                                                                                                           Sharpurs_Prelude::unbox(v_3);
                                                                                                       if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_3_1_0,
                                                                                                                                                                 matchValue_3_1_1)
                                                                                                              =
                                                                                                              matchValue_3.as_ref()
                                                                                                          {
                                                                                                           let activePatternResult:
                                                                                                                   LrcPtr<Data_List_Types_List> =
                                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_3.as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                         x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                      _
                                                                                                                                                      =>
                                                                                                                                                      unreachable!(),
                                                                                                                                                  });
                                                                                                           if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(activePatternResult_1_0,
                                                                                                                                                                     activePatternResult_1_1)
                                                                                                                  =
                                                                                                                  activePatternResult.as_ref()
                                                                                                              {
                                                                                                               &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(merge_tco(&match matchValue_3.as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          })(&match activePatternResult.as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                  _
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                              }),
                                                                                                                                                                               mergePairs_tco(&match activePatternResult.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                   _
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                               })))
                                                                                                           } else {
                                                                                                               &matchValue_3
                                                                                                           }
                                                                                                       } else {
                                                                                                           &matchValue_3
                                                                                                       }
                                                                                                   }),
                                                                                             v_3.clone())
                                                                                });
                                                                 let mergePairs =
                                                                     mergePairs_1.Value;
                                                                 {
                                                                     let mergeAll_2 =
                                                                         Func0::new({
                                                                                        let mergeAll_tco
                                                                                            =
                                                                                            mergeAll_tco.clone();
                                                                                        move
                                                                                            ||
                                                                                            &Func1::new({
                                                                                                            let mergeAll_tco
                                                                                                                =
                                                                                                                mergeAll_tco.clone();
                                                                                                            move
                                                                                                                |v_4|
                                                                                                                mergeAll_tco(v_4.clone())
                                                                                                        })
                                                                                    });
                                                                     let mergeAll_1 =
                                                                         Lazy(mergeAll_2);
                                                                     let mergeAll_tco =
                                                                         Func1::new({
                                                                                        let mergePairs_tco
                                                                                            =
                                                                                            mergePairs_tco.clone();
                                                                                        move
                                                                                            |v_5|
                                                                                            {
                                                                                                let v_5 =
                                                                                                    v_5.clone();
                                                                                                '_mergeAll_tco:
                                                                                                    loop 
                                                                                                         {
                                                                                                        break
                                                                                                            '_mergeAll_tco
                                                                                                            ({
                                                                                                                 let matchValue_4:
                                                                                                                         LrcPtr<Data_List_Types_List> =
                                                                                                                     Sharpurs_Prelude::unbox(&&v_5);
                                                                                                                 if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_4_1_0,
                                                                                                                                                                           matchValue_4_1_1)
                                                                                                                        =
                                                                                                                        matchValue_4.as_ref()
                                                                                                                    {
                                                                                                                     if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                            =
                                                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_4.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                      x)
                                                                                                                                                                   =>
                                                                                                                                                                   x.clone(),
                                                                                                                                                                   _
                                                                                                                                                                   =>
                                                                                                                                                                   unreachable!(),
                                                                                                                                                               }).as_ref()
                                                                                                                        {
                                                                                                                         &match matchValue_4.as_ref()
                                                                                                                              {
                                                                                                                              Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                 _)
                                                                                                                              =>
                                                                                                                              x.clone(),
                                                                                                                              _
                                                                                                                              =>
                                                                                                                              unreachable!(),
                                                                                                                          }
                                                                                                                     } else {
                                                                                                                         let v_5_temp =
                                                                                                                             mergePairs_tco(&matchValue_4);
                                                                                                                         v_5.set(v_5_temp);
                                                                                                                         continue
                                                                                                                             '_mergeAll_tco

                                                                                                                     }
                                                                                                                 } else {
                                                                                                                     let v_5_temp =
                                                                                                                         mergePairs_tco(&matchValue_4);
                                                                                                                     v_5.set(v_5_temp);
                                                                                                                     continue
                                                                                                                         '_mergeAll_tco

                                                                                                                 }
                                                                                                             })
                                                                                                            ;
                                                                                                    }
                                                                                            }
                                                                                    });
                                                                     let mergeAll =
                                                                         mergeAll_1.Value;
                                                                     {
                                                                         let sequences_2 =
                                                                             Func0::new({
                                                                                            let sequences_tco
                                                                                                =
                                                                                                sequences_tco.clone();
                                                                                            move
                                                                                                ||
                                                                                                &Func1::new({
                                                                                                                let sequences_tco
                                                                                                                    =
                                                                                                                    sequences_tco.clone();
                                                                                                                move
                                                                                                                    |v_6|
                                                                                                                    sequences_tco(v_6.clone())
                                                                                                            })
                                                                                        });
                                                                         let sequences_1 =
                                                                             Lazy(sequences_2);
                                                                         let descending_2 =
                                                                             Func0::new({
                                                                                            let descending_tco
                                                                                                =
                                                                                                descending_tco.clone();
                                                                                            move
                                                                                                ||
                                                                                                &Func1::new({
                                                                                                                let descending_tco
                                                                                                                    =
                                                                                                                    descending_tco.clone();
                                                                                                                move
                                                                                                                    |v_7|
                                                                                                                    Func1::new({
                                                                                                                                   let descending_tco
                                                                                                                                       =
                                                                                                                                       descending_tco.clone();
                                                                                                                                   let v_7
                                                                                                                                       =
                                                                                                                                       v_7.clone();
                                                                                                                                   move
                                                                                                                                       |v1_2|
                                                                                                                                       Func1::new({
                                                                                                                                                      let descending_tco
                                                                                                                                                          =
                                                                                                                                                          descending_tco.clone();
                                                                                                                                                      let v1_2
                                                                                                                                                          =
                                                                                                                                                          v1_2.clone();
                                                                                                                                                      move
                                                                                                                                                          |v2|
                                                                                                                                                          descending_tco(v_7)(v1_2)(v2.clone())
                                                                                                                                                  })
                                                                                                                               })
                                                                                                            })
                                                                                        });
                                                                         let descending_1 =
                                                                             Lazy(descending_2);
                                                                         let ascending_2 =
                                                                             Func0::new({
                                                                                            let ascending_tco
                                                                                                =
                                                                                                ascending_tco.clone();
                                                                                            move
                                                                                                ||
                                                                                                &Func1::new({
                                                                                                                let ascending_tco
                                                                                                                    =
                                                                                                                    ascending_tco.clone();
                                                                                                                move
                                                                                                                    |v_8|
                                                                                                                    Func1::new({
                                                                                                                                   let ascending_tco
                                                                                                                                       =
                                                                                                                                       ascending_tco.clone();
                                                                                                                                   let v_8
                                                                                                                                       =
                                                                                                                                       v_8.clone();
                                                                                                                                   move
                                                                                                                                       |v1_3|
                                                                                                                                       Func1::new({
                                                                                                                                                      let ascending_tco
                                                                                                                                                          =
                                                                                                                                                          ascending_tco.clone();
                                                                                                                                                      let v1_3
                                                                                                                                                          =
                                                                                                                                                          v1_3.clone();
                                                                                                                                                      move
                                                                                                                                                          |v2_1|
                                                                                                                                                          ascending_tco(v_8)(v1_3)(v2_1.clone())
                                                                                                                                                  })
                                                                                                                               })
                                                                                                            })
                                                                                        });
                                                                         let ascending_1 =
                                                                             Lazy(ascending_2);
                                                                         let sequences_tco =
                                                                             Func1::new({
                                                                                            let ascending_tco
                                                                                                =
                                                                                                ascending_tco.clone();
                                                                                            let cmp
                                                                                                =
                                                                                                cmp.clone();
                                                                                            let descending_tco
                                                                                                =
                                                                                                descending_tco.clone();
                                                                                            move
                                                                                                |v_9|
                                                                                                {
                                                                                                    let matchValue_5:
                                                                                                            LrcPtr<Data_List_Types_List> =
                                                                                                        Sharpurs_Prelude::unbox(v_9);
                                                                                                    if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_5_1_0,
                                                                                                                                                              matchValue_5_1_1)
                                                                                                           =
                                                                                                           matchValue_5.as_ref()
                                                                                                       {
                                                                                                        let activePatternResult_2:
                                                                                                                LrcPtr<Data_List_Types_List> =
                                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_5.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                      x)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                                   _
                                                                                                                                                   =>
                                                                                                                                                   unreachable!(),
                                                                                                                                               });
                                                                                                        if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(activePatternResult_2_1_0,
                                                                                                                                                                  activePatternResult_2_1_1)
                                                                                                               =
                                                                                                               activePatternResult_2.as_ref()
                                                                                                           {
                                                                                                            if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                               &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                  &&&match matchValue_5.as_ref()
                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                         Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                               &&&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                      Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                      _
                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                                                  })),
                                                                                                                                                                         &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)))
                                                                                                               {
                                                                                                                descending_tco(&match activePatternResult_2.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                       _)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                    _
                                                                                                                                    =>
                                                                                                                                    unreachable!(),
                                                                                                                                })(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                                                    &&&match matchValue_5.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                              _)
                                                                                                                                                                           =>
                                                                                                                                                                           x.clone(),
                                                                                                                                                                           _
                                                                                                                                                                           =>
                                                                                                                                                                           unreachable!(),
                                                                                                                                                                       }))(&match activePatternResult_2.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                   x)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                unreachable!(),
                                                                                                                                                                            })
                                                                                                            } else {
                                                                                                                let activePatternResult_3:
                                                                                                                        LrcPtr<Data_List_Types_List> =
                                                                                                                    Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_5.as_ref()
                                                                                                                                                           {
                                                                                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                              x)
                                                                                                                                                           =>
                                                                                                                                                           x.clone(),
                                                                                                                                                           _
                                                                                                                                                           =>
                                                                                                                                                           unreachable!(),
                                                                                                                                                       });
                                                                                                                if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(activePatternResult_3_1_0,
                                                                                                                                                                          activePatternResult_3_1_1)
                                                                                                                       =
                                                                                                                       activePatternResult_3.as_ref()
                                                                                                                   {
                                                                                                                    if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                       {
                                                                                                                        ascending_tco(&match activePatternResult_3.as_ref()
                                                                                                                                           {
                                                                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                              _)
                                                                                                                                           =>
                                                                                                                                           x.clone(),
                                                                                                                                           _
                                                                                                                                           =>
                                                                                                                                           unreachable!(),
                                                                                                                                       })(&Func1::new({
                                                                                                                                                          let matchValue_5
                                                                                                                                                              =
                                                                                                                                                              matchValue_5.clone();
                                                                                                                                                          move
                                                                                                                                                              |v1_4|
                                                                                                                                                              &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_5.as_ref()
                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                               },
                                                                                                                                                                                                                              v1_4.clone()))
                                                                                                                                                      }))(&match activePatternResult_3.as_ref()
                                                                                                                                                               {
                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                  x)
                                                                                                                                                               =>
                                                                                                                                                               x.clone(),
                                                                                                                                                               _
                                                                                                                                                               =>
                                                                                                                                                               unreachable!(),
                                                                                                                                                           })
                                                                                                                    } else {
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                                         &&&matchValue_5)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                                     &&&matchValue_5)
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            let activePatternResult_4:
                                                                                                                    LrcPtr<Data_List_Types_List> =
                                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_5.as_ref()
                                                                                                                                                       {
                                                                                                                                                       Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                          x)
                                                                                                                                                       =>
                                                                                                                                                       x.clone(),
                                                                                                                                                       _
                                                                                                                                                       =>
                                                                                                                                                       unreachable!(),
                                                                                                                                                   });
                                                                                                            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(activePatternResult_4_1_0,
                                                                                                                                                                      activePatternResult_4_1_1)
                                                                                                                   =
                                                                                                                   activePatternResult_4.as_ref()
                                                                                                               {
                                                                                                                if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                   {
                                                                                                                    ascending_tco(&match activePatternResult_4.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                          _)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                       _
                                                                                                                                       =>
                                                                                                                                       unreachable!(),
                                                                                                                                   })(&Func1::new({
                                                                                                                                                      let matchValue_5
                                                                                                                                                          =
                                                                                                                                                          matchValue_5.clone();
                                                                                                                                                      move
                                                                                                                                                          |v1_4|
                                                                                                                                                          &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_5.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                               _
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                           },
                                                                                                                                                                                                                          v1_4.clone()))
                                                                                                                                                  }))(&match activePatternResult_4.as_ref()
                                                                                                                                                           {
                                                                                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                              x)
                                                                                                                                                           =>
                                                                                                                                                           x.clone(),
                                                                                                                                                           _
                                                                                                                                                           =>
                                                                                                                                                           unreachable!(),
                                                                                                                                                       })
                                                                                                                } else {
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                                     &&&matchValue_5)
                                                                                                                }
                                                                                                            } else {
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                                 &&&matchValue_5)
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                         &&&matchValue_5)
                                                                                                    }
                                                                                                }
                                                                                        });
                                                                         let sequences =
                                                                             sequences_1.Value;
                                                                         let descending_tco =
                                                                             Func1::new({
                                                                                            let cmp
                                                                                                =
                                                                                                cmp.clone();
                                                                                            let sequences_tco
                                                                                                =
                                                                                                sequences_tco.clone();
                                                                                            move
                                                                                                |v_10|
                                                                                                fix1(&(move
                                                                                                           |descending_tco,
                                                                                                            v_10|
                                                                                                           Func1::new({
                                                                                                                          let descending_tco
                                                                                                                              =
                                                                                                                              descending_tco.clone();
                                                                                                                          let sequences_tco
                                                                                                                              =
                                                                                                                              sequences_tco.clone();
                                                                                                                          let v_10
                                                                                                                              =
                                                                                                                              v_10.clone();
                                                                                                                          move
                                                                                                                              |v1_5|
                                                                                                                              Func1::new({
                                                                                                                                             let descending_tco
                                                                                                                                                 =
                                                                                                                                                 descending_tco.clone();
                                                                                                                                             let sequences_tco
                                                                                                                                                 =
                                                                                                                                                 sequences_tco.clone();
                                                                                                                                             let v1_5
                                                                                                                                                 =
                                                                                                                                                 v1_5.clone();
                                                                                                                                             move
                                                                                                                                                 |v2_2|
                                                                                                                                                 {
                                                                                                                                                     let matchValue_6 =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&v_10);
                                                                                                                                                     let matchValue_7 =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&v1_5);
                                                                                                                                                     let matchValue_8:
                                                                                                                                                             LrcPtr<Data_List_Types_List> =
                                                                                                                                                         Sharpurs_Prelude::unbox(v2_2);
                                                                                                                                                     if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_8_1_0,
                                                                                                                                                                                                               matchValue_8_1_1)
                                                                                                                                                            =
                                                                                                                                                            matchValue_8.as_ref()
                                                                                                                                                        {
                                                                                                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                               &&&matchValue_6),
                                                                                                                                                                                                                                                                                            &&&match matchValue_8.as_ref()
                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)))
                                                                                                                                                            {
                                                                                                                                                             descending_tco(&match matchValue_8.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                 Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                    _)
                                                                                                                                                                                 =>
                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                 _
                                                                                                                                                                                 =>
                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                             })(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue_6,
                                                                                                                                                                                                                                                &matchValue_7)))(&match matchValue_8.as_ref()
                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                      Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                      _
                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                                                                  })
                                                                                                                                                         } else {
                                                                                                                                                             &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue_6,
                                                                                                                                                                                                                                                                                             &matchValue_7)),
                                                                                                                                                                                                                             sequences_tco(&matchValue_8)))
                                                                                                                                                         }
                                                                                                                                                     } else {
                                                                                                                                                         &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue_6,
                                                                                                                                                                                                                                                                                         &matchValue_7)),
                                                                                                                                                                                                                         sequences_tco(&matchValue_8)))
                                                                                                                                                     }
                                                                                                                                                 }
                                                                                                                                         })
                                                                                                                      })),
                                                                                                     v_10.clone())
                                                                                        });
                                                                         let descending =
                                                                             descending_1.Value;
                                                                         let ascending_tco =
                                                                             Func1::new({
                                                                                            let cmp
                                                                                                =
                                                                                                cmp.clone();
                                                                                            let sequences_tco
                                                                                                =
                                                                                                sequences_tco.clone();
                                                                                            move
                                                                                                |v_11|
                                                                                                fix1(&(move
                                                                                                           |ascending_tco,
                                                                                                            v_11|
                                                                                                           Func1::new({
                                                                                                                          let ascending_tco
                                                                                                                              =
                                                                                                                              ascending_tco.clone();
                                                                                                                          let sequences_tco
                                                                                                                              =
                                                                                                                              sequences_tco.clone();
                                                                                                                          let v_11
                                                                                                                              =
                                                                                                                              v_11.clone();
                                                                                                                          move
                                                                                                                              |v1_6|
                                                                                                                              Func1::new({
                                                                                                                                             let ascending_tco
                                                                                                                                                 =
                                                                                                                                                 ascending_tco.clone();
                                                                                                                                             let sequences_tco
                                                                                                                                                 =
                                                                                                                                                 sequences_tco.clone();
                                                                                                                                             let v1_6
                                                                                                                                                 =
                                                                                                                                                 v1_6.clone();
                                                                                                                                             move
                                                                                                                                                 |v2_3|
                                                                                                                                                 {
                                                                                                                                                     let matchValue_10 =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&v_11);
                                                                                                                                                     let matchValue_11 =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&v1_6);
                                                                                                                                                     let matchValue_12:
                                                                                                                                                             LrcPtr<Data_List_Types_List> =
                                                                                                                                                         Sharpurs_Prelude::unbox(v2_3);
                                                                                                                                                     if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_12_1_0,
                                                                                                                                                                                                               matchValue_12_1_1)
                                                                                                                                                            =
                                                                                                                                                            matchValue_12.as_ref()
                                                                                                                                                        {
                                                                                                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_notEq(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                                                                                                                                                                                               &&&matchValue_10),
                                                                                                                                                                                                                                                                                            &&&match matchValue_12.as_ref()
                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)))
                                                                                                                                                            {
                                                                                                                                                             ascending_tco(&match matchValue_12.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                   _)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                unreachable!(),
                                                                                                                                                                            })(&Func1::new({
                                                                                                                                                                                               let matchValue_10
                                                                                                                                                                                                   =
                                                                                                                                                                                                   matchValue_10.clone();
                                                                                                                                                                                               let matchValue_11
                                                                                                                                                                                                   =
                                                                                                                                                                                                   matchValue_11.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |ys|
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&matchValue_11,
                                                                                                                                                                                                                                    &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue_10,
                                                                                                                                                                                                                                                                                                      ys.clone())))
                                                                                                                                                                                           }))(&match matchValue_12.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                    _
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                })
                                                                                                                                                         } else {
                                                                                                                                                             &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                 &&&matchValue_11),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                                                                                                                                                                                 &&&matchValue_10)),
                                                                                                                                                                                                                             sequences_tco(&matchValue_12)))
                                                                                                                                                         }
                                                                                                                                                     } else {
                                                                                                                                                         &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                             &&&matchValue_11),
                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                                                                                                                                                                                             &&&matchValue_10)),
                                                                                                                                                                                                                         sequences_tco(&matchValue_12)))
                                                                                                                                                     }
                                                                                                                                                 }
                                                                                                                                         })
                                                                                                                      })),
                                                                                                     v_11.clone())
                                                                                        });
                                                                         let ascending =
                                                                             ascending_1.Value;
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                             &&&mergeAll),
                                                                                                          &&&sequences)
                                                                     }
                                                                 }
                                                             }
                                                         }))
    }
    pub fn Data_List_sort() -> &dyn Any {
        static Data_List_sort: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_sort.get_or_init(||
                                       &Func1::new(move |dictOrd|
                                                       {
                                                           let compare =
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                dictOrd);
                                                           &Func1::new({
                                                                           let compare
                                                                               =
                                                                               compare.clone();
                                                                           move
                                                                               |xs|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_sortBy(),
                                                                                                                                                   &&&compare),
                                                                                                                xs)
                                                                       })
                                                       }))
    }
    pub fn Data_List_tails_004062() -> &dyn Any {
        &Func1::new(move |v| PureScript_Data_List::Data_List_tails_tco(v))
    }
    pub fn Data_List_tails_004062_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_tails_004062_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_tails_004062_002d1.get_or_init(||
                                                     Lazy(Data_List_tails_004062.clone()))
    }
    pub fn Data_List_tails_tco(v: &dyn Any) -> &dyn Any {
        let matchValue: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v);
        match matchValue.as_ref() {
            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                               matchValue_1_1)
            =>
            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue.clone(),
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&Data_List_tails_004062_002d1.Value,
                                                                                                             &&matchValue_1_1))),
            _ =>
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                             &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)),
        }
    }
    pub fn Data_List_tails() -> &dyn Any {
        static Data_List_tails: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_tails.get_or_init(|| Data_List_tails_004062_002d1.Value)
    }
    pub fn Data_List_showPattern() -> &dyn Any {
        static Data_List_showPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_showPattern.get_or_init(||
                                              &Func1::new(move |dictShow|
                                                              {
                                                                  let showList =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_showList(),
                                                                                                       dictShow);
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                   &&&add(string("show"),
                                                                                                          &&Func1::new({
                                                                                                                           let showList
                                                                                                                               =
                                                                                                                               showList.clone();
                                                                                                                           move
                                                                                                                               |v|
                                                                                                                               {
                                                                                                                                   let s =
                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                       &&&string("(Pattern ")),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                &&&showList),
                                                                                                                                                                                                                                                                             &&&s)),
                                                                                                                                                                                                       &&&string(")")))
                                                                                                                               }
                                                                                                                       }),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))
                                                              }))
    }
    pub fn Data_List_reverse() -> &dyn Any {
        static Data_List_reverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_reverse.get_or_init(||
                                          {
                                              let go_2 =
                                                  Func0::new({
                                                                 let go_tco =
                                                                     go_tco.clone();
                                                                 move ||
                                                                     &Func1::new({
                                                                                     let go_tco
                                                                                         =
                                                                                         go_tco.clone();
                                                                                     move
                                                                                         |v|
                                                                                         Func1::new({
                                                                                                        let go_tco
                                                                                                            =
                                                                                                            go_tco.clone();
                                                                                                        let v
                                                                                                            =
                                                                                                            v.clone();
                                                                                                        move
                                                                                                            |v1|
                                                                                                            go_tco(v)(v1.clone())
                                                                                                    })
                                                                                 })
                                                             });
                                              let go_1 = Lazy(go_2);
                                              fn go_tco(v_1: &dyn Any)
                                               -> Func1<&dyn Any, &dyn Any> {
                                                  Func1::new({
                                                                 let go_tco =
                                                                     go_tco.clone();
                                                                 let v_1 =
                                                                     v_1.clone();
                                                                 move |v1_1|
                                                                     {
                                                                         let matchValue =
                                                                             Sharpurs_Prelude::unbox(&&v_1);
                                                                         let matchValue_1:
                                                                                 LrcPtr<Data_List_Types_List> =
                                                                             Sharpurs_Prelude::unbox(v1_1);
                                                                         match matchValue_1.as_ref()
                                                                             {
                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                matchValue_1_1_1)
                                                                             =>
                                                                             go_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                    &matchValue)))(matchValue_1_1_1),
                                                                             _
                                                                             =>
                                                                             &matchValue,
                                                                         }
                                                                     }
                                                             })
                                              }
                                              let go = go_1.Value;
                                              Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                               &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))
                                          })
    }
    pub fn Data_List_take() -> &dyn Any {
        static Data_List_take: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_take.get_or_init(||
                                       {
                                           let go_2 =
                                               Func0::new({
                                                              let go_tco =
                                                                  go_tco.clone();
                                                              move ||
                                                                  &Func1::new({
                                                                                  let go_tco
                                                                                      =
                                                                                      go_tco.clone();
                                                                                  move
                                                                                      |v|
                                                                                      Func1::new({
                                                                                                     let go_tco
                                                                                                         =
                                                                                                         go_tco.clone();
                                                                                                     let v
                                                                                                         =
                                                                                                         v.clone();
                                                                                                     move
                                                                                                         |v1|
                                                                                                         Func1::new({
                                                                                                                        let go_tco
                                                                                                                            =
                                                                                                                            go_tco.clone();
                                                                                                                        let v1
                                                                                                                            =
                                                                                                                            v1.clone();
                                                                                                                        move
                                                                                                                            |v2|
                                                                                                                            go_tco(v)(v1)(v2.clone())
                                                                                                                    })
                                                                                                 })
                                                                              })
                                                          });
                                           let go_1 = Lazy(go_2);
                                           fn go_tco(v_1: &dyn Any)
                                            ->
                                                Func1<&dyn Any,
                                                      Func1<&dyn Any,
                                                            &dyn Any>> {
                                               Func1::new({
                                                              let go_tco =
                                                                  go_tco.clone();
                                                              let v_1 =
                                                                  v_1.clone();
                                                              move |v1_1|
                                                                  Func1::new({
                                                                                 let go_tco
                                                                                     =
                                                                                     go_tco.clone();
                                                                                 let v1_1
                                                                                     =
                                                                                     v1_1.clone();
                                                                                 move
                                                                                     |v2_1|
                                                                                     {
                                                                                         let matchValue =
                                                                                             Sharpurs_Prelude::unbox(&&v_1);
                                                                                         let matchValue_1 =
                                                                                             Sharpurs_Prelude::unbox(&&v1_1);
                                                                                         let matchValue_2:
                                                                                                 LrcPtr<Data_List_Types_List> =
                                                                                             Sharpurs_Prelude::unbox(v2_1);
                                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                            &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                         &&&matchValue_1),
                                                                                                                                                      &&&1_i32))
                                                                                            {
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                              &&&matchValue)
                                                                                         } else {
                                                                                             if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                                                                                                       matchValue_2_1_1)
                                                                                                    =
                                                                                                    matchValue_2.as_ref()
                                                                                                {
                                                                                                 go_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                _)
                                                                                                                                                                             =>
                                                                                                                                                                             x.clone(),
                                                                                                                                                                             _
                                                                                                                                                                             =>
                                                                                                                                                                             unreachable!(),
                                                                                                                                                                         },
                                                                                                                                                                        &matchValue)))(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                           &&&matchValue_1),
                                                                                                                                                                                                                        &&&1_i32))(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                    })
                                                                                             } else {
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                  &&&matchValue)
                                                                                             }
                                                                                         }
                                                                                     }
                                                                             })
                                                          })
                                           }
                                           let go = go_1.Value;
                                           Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                            &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))
                                       })
    }
    pub fn Data_List_takeWhile() -> &dyn Any {
        static Data_List_takeWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_takeWhile.get_or_init(||
                                            &Func1::new(move |p|
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
                                                                                                           |v|
                                                                                                           Func1::new({
                                                                                                                          let go_tco
                                                                                                                              =
                                                                                                                              go_tco.clone();
                                                                                                                          let v
                                                                                                                              =
                                                                                                                              v.clone();
                                                                                                                          move
                                                                                                                              |v1|
                                                                                                                              go_tco(v)(v1.clone())
                                                                                                                      })
                                                                                                   })
                                                                               });
                                                                let go_1 =
                                                                    Lazy(go_2);
                                                                let go_tco =
                                                                    Func1::new({
                                                                                   let p
                                                                                       =
                                                                                       p.clone();
                                                                                   move
                                                                                       |v_1|
                                                                                       fix1(&(move
                                                                                                  |go_tco,
                                                                                                   v_1|
                                                                                                  Func1::new({
                                                                                                                 let go_tco
                                                                                                                     =
                                                                                                                     go_tco.clone();
                                                                                                                 let v_1
                                                                                                                     =
                                                                                                                     v_1.clone();
                                                                                                                 move
                                                                                                                     |v1_1|
                                                                                                                     {
                                                                                                                         let matchValue =
                                                                                                                             Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                         let matchValue_1:
                                                                                                                                 LrcPtr<Data_List_Types_List> =
                                                                                                                             Sharpurs_Prelude::unbox(v1_1);
                                                                                                                         if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                   matchValue_1_1_1)
                                                                                                                                =
                                                                                                                                matchValue_1.as_ref()
                                                                                                                            {
                                                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                 Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                 =>
                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                 _
                                                                                                                                                                                                 =>
                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                             }))
                                                                                                                                {
                                                                                                                                 go_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                             _
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &matchValue)))(&match matchValue_1.as_ref()
                                                                                                                                                                                                                            {
                                                                                                                                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                            _
                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                                        })
                                                                                                                             } else {
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                                  &&&matchValue)
                                                                                                                             }
                                                                                                                         } else {
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                              &&&matchValue)
                                                                                                                         }
                                                                                                                     }
                                                                                                             })),
                                                                                            v_1.clone())
                                                                               });
                                                                let go =
                                                                    go_1.Value;
                                                                Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                 &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))
                                                            }))
    }
    pub fn Data_List_unsnoc() -> &dyn Any {
        static Data_List_unsnoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_unsnoc.get_or_init(||
                                         &Func1::new(move |lst|
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
                                                                                                        |v|
                                                                                                        Func1::new({
                                                                                                                       let go_tco
                                                                                                                           =
                                                                                                                           go_tco.clone();
                                                                                                                       let v
                                                                                                                           =
                                                                                                                           v.clone();
                                                                                                                       move
                                                                                                                           |v1|
                                                                                                                           go_tco(v)(v1.clone())
                                                                                                                   })
                                                                                                })
                                                                            });
                                                             let go_1 =
                                                                 Lazy(go_2);
                                                             fn go_tco(v_1: _)
                                                              ->
                                                                  Func1<&dyn Any,
                                                                        &dyn Any> {
                                                                 Func1::new({
                                                                                let go_tco
                                                                                    =
                                                                                    go_tco.clone();
                                                                                let v_1
                                                                                    =
                                                                                    v_1.clone();
                                                                                move
                                                                                    |v1_1|
                                                                                    {
                                                                                        let matchValue:
                                                                                                LrcPtr<Data_List_Types_List> =
                                                                                            Sharpurs_Prelude::unbox(&&v_1);
                                                                                        let matchValue_1 =
                                                                                            Sharpurs_Prelude::unbox(v1_1);
                                                                                        if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                  matchValue_1_1)
                                                                                               =
                                                                                               matchValue.as_ref()
                                                                                           {
                                                                                            if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                   =
                                                                                                   Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                                          {
                                                                                                                                          Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                             x)
                                                                                                                                          =>
                                                                                                                                          x.clone(),
                                                                                                                                          _
                                                                                                                                          =>
                                                                                                                                          unreachable!(),
                                                                                                                                      }).as_ref()
                                                                                               {
                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&add(string("revInit"),
                                                                                                                                                            &&matchValue_1,
                                                                                                                                                            add(string("last"),
                                                                                                                                                                &&match matchValue.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                         _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                      _
                                                                                                                                                                      =>
                                                                                                                                                                      unreachable!(),
                                                                                                                                                                  },
                                                                                                                                                                empty::<string,
                                                                                                                                                                        &dyn Any>()))))
                                                                                            } else {
                                                                                                go_tco(&match matchValue.as_ref()
                                                                                                            {
                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                               x)
                                                                                                            =>
                                                                                                            x.clone(),
                                                                                                            _
                                                                                                            =>
                                                                                                            unreachable!(),
                                                                                                        })(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                   _)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                unreachable!(),
                                                                                                                                                                            },
                                                                                                                                                                           &matchValue_1)))
                                                                                            }
                                                                                        } else {
                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                        }
                                                                                    }
                                                                            })
                                                             }
                                                             let go =
                                                                 go_1.Value;
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                    &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                 &&&Func1::new(move
                                                                                                                                                   |h|
                                                                                                                                                   &add(string("init"),
                                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                                                          &&find(string("revInit"),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(h))),
                                                                                                                                                        add(string("last"),
                                                                                                                                                            &find(string("last"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(h)),
                                                                                                                                                            empty::<string,
                                                                                                                                                                    &dyn Any>())))),
                                                                                              &&go_tco(lst.clone())(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)))
                                                         }))
    }
    pub fn Data_List_zipWith() -> &dyn Any {
        static Data_List_zipWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_zipWith.get_or_init(||
                                          &Func1::new(move |f|
                                                          &Func1::new({
                                                                          let f
                                                                              =
                                                                              f.clone();
                                                                          move
                                                                              |xs|
                                                                              &Func1::new({
                                                                                              let xs
                                                                                                  =
                                                                                                  xs.clone();
                                                                                              move
                                                                                                  |ys|
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
                                                                                                                                                 |v|
                                                                                                                                                 Func1::new({
                                                                                                                                                                let go_tco
                                                                                                                                                                    =
                                                                                                                                                                    go_tco.clone();
                                                                                                                                                                let v
                                                                                                                                                                    =
                                                                                                                                                                    v.clone();
                                                                                                                                                                move
                                                                                                                                                                    |v1|
                                                                                                                                                                    Func1::new({
                                                                                                                                                                                   let go_tco
                                                                                                                                                                                       =
                                                                                                                                                                                       go_tco.clone();
                                                                                                                                                                                   let v1
                                                                                                                                                                                       =
                                                                                                                                                                                       v1.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |v2|
                                                                                                                                                                                       go_tco(v)(v1)(v2.clone())
                                                                                                                                                                               })
                                                                                                                                                            })
                                                                                                                                         })
                                                                                                                     });
                                                                                                      let go_1 =
                                                                                                          Lazy(go_2);
                                                                                                      let go_tco =
                                                                                                          Func1::new({
                                                                                                                         let go_1
                                                                                                                             =
                                                                                                                             go_1.clone();
                                                                                                                         move
                                                                                                                             |v_1|
                                                                                                                             Func1::new({
                                                                                                                                            let go_1
                                                                                                                                                =
                                                                                                                                                go_1.clone();
                                                                                                                                            let v_1
                                                                                                                                                =
                                                                                                                                                v_1.clone();
                                                                                                                                            move
                                                                                                                                                |v1_1|
                                                                                                                                                Func1::new({
                                                                                                                                                               let go_1
                                                                                                                                                                   =
                                                                                                                                                                   go_1.clone();
                                                                                                                                                               let v1_1
                                                                                                                                                                   =
                                                                                                                                                                   v1_1.clone();
                                                                                                                                                               move
                                                                                                                                                                   |v2_1|
                                                                                                                                                                   {
                                                                                                                                                                       let matchValue:
                                                                                                                                                                               LrcPtr<Data_List_Types_List> =
                                                                                                                                                                           Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                       let matchValue_1:
                                                                                                                                                                               LrcPtr<Data_List_Types_List> =
                                                                                                                                                                           Sharpurs_Prelude::unbox(&&v1_1);
                                                                                                                                                                       let matchValue_2 =
                                                                                                                                                                           Sharpurs_Prelude::unbox(v2_1);
                                                                                                                                                                       if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                 matchValue_1_1)
                                                                                                                                                                              =
                                                                                                                                                                              matchValue.as_ref()
                                                                                                                                                                          {
                                                                                                                                                                           if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                     matchValue_1_1_1)
                                                                                                                                                                                  =
                                                                                                                                                                                  matchValue_1.as_ref()
                                                                                                                                                                              {
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&go_1.Value,
                                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                                      &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                                                         })),
                                                                                                                                                                                                                &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                                                                   &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                          _
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                  &matchValue_2)))
                                                                                                                                                                           } else {
                                                                                                                                                                               &matchValue_2
                                                                                                                                                                           }
                                                                                                                                                                       } else {
                                                                                                                                                                           &matchValue_2
                                                                                                                                                                       }
                                                                                                                                                                   }
                                                                                                                                                           })
                                                                                                                                        })
                                                                                                                     });
                                                                                                      let go =
                                                                                                          go_1.Value;
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                          &&&PureScript_Data_List::Data_List_reverse()),
                                                                                                                                       &&go_tco(&xs)(ys.clone())(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)))
                                                                                                  }
                                                                                          })
                                                                      })))
    }
    pub fn Data_List_zip() -> &dyn Any {
        static Data_List_zip: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_zip.get_or_init(||
                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_zipWith(),
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
                                                                                                    }))))
    }
    pub fn Data_List_zipWithA() -> &dyn Any {
        static Data_List_zipWithA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_zipWithA.get_or_init(||
                                           &Func1::new(move |dictApplicative|
                                                           &Func1::new({
                                                                           let dictApplicative
                                                                               =
                                                                               dictApplicative.clone();
                                                                           move
                                                                               |f|
                                                                               &Func1::new({
                                                                                               let f
                                                                                                   =
                                                                                                   f.clone();
                                                                                               move
                                                                                                   |xs|
                                                                                                   &Func1::new({
                                                                                                                   let xs
                                                                                                                       =
                                                                                                                       xs.clone();
                                                                                                                   move
                                                                                                                       |ys|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                              &&&PureScript_Data_List_Types::Data_List_Types_traversableList()),
                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_zipWith(),
                                                                                                                                                                                                                                                                 &&&f),
                                                                                                                                                                                                                              &&&xs),
                                                                                                                                                                                           ys))
                                                                                                               })
                                                                                           })
                                                                       })))
    }
    pub fn Data_List_range() -> &dyn Any {
        static Data_List_range: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_range.get_or_init(||
                                        &Func1::new(move |start|
                                                        &Func1::new({
                                                                        let start
                                                                            =
                                                                            start.clone();
                                                                        move
                                                                            |end_var|
                                                                            {
                                                                                let matchValue =
                                                                                    Sharpurs_Prelude::unbox(&&start);
                                                                                let matchValue_1 =
                                                                                    Sharpurs_Prelude::unbox(end_var);
                                                                                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                   &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                &&&matchValue),
                                                                                                                                             &&&matchValue_1))
                                                                                   {
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                                                                                                     &&&matchValue)
                                                                                } else {
                                                                                    if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                       {
                                                                                        let start1_3 =
                                                                                            matchValue;
                                                                                        let end1_3 =
                                                                                            matchValue_1;
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
                                                                                                                                       |s|
                                                                                                                                       Func1::new({
                                                                                                                                                      let go_tco
                                                                                                                                                          =
                                                                                                                                                          go_tco.clone();
                                                                                                                                                      let s
                                                                                                                                                          =
                                                                                                                                                          s.clone();
                                                                                                                                                      move
                                                                                                                                                          |e|
                                                                                                                                                          Func1::new({
                                                                                                                                                                         let e
                                                                                                                                                                             =
                                                                                                                                                                             e.clone();
                                                                                                                                                                         let go_tco
                                                                                                                                                                             =
                                                                                                                                                                             go_tco.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |step|
                                                                                                                                                                             Func1::new({
                                                                                                                                                                                            let go_tco
                                                                                                                                                                                                =
                                                                                                                                                                                                go_tco.clone();
                                                                                                                                                                                            let step
                                                                                                                                                                                                =
                                                                                                                                                                                                step.clone();
                                                                                                                                                                                            move
                                                                                                                                                                                                |rest|
                                                                                                                                                                                                go_tco(s)(e)(step)(rest.clone())
                                                                                                                                                                                        })
                                                                                                                                                                     })
                                                                                                                                                  })
                                                                                                                               })
                                                                                                           });
                                                                                            let go_1 =
                                                                                                Lazy(go_2);
                                                                                            fn go_tco(s_1:
                                                                                                          _)
                                                                                             ->
                                                                                                 Func1<&dyn Any,
                                                                                                       Func1<&dyn Any,
                                                                                                             Func1<&dyn Any,
                                                                                                                   &dyn Any>>> {
                                                                                                Func1::new({
                                                                                                               let go_tco
                                                                                                                   =
                                                                                                                   go_tco.clone();
                                                                                                               let s_1
                                                                                                                   =
                                                                                                                   s_1.clone();
                                                                                                               move
                                                                                                                   |e_1|
                                                                                                                   Func1::new({
                                                                                                                                  let e_1
                                                                                                                                      =
                                                                                                                                      e_1.clone();
                                                                                                                                  let go_tco
                                                                                                                                      =
                                                                                                                                      go_tco.clone();
                                                                                                                                  move
                                                                                                                                      |step_1|
                                                                                                                                      Func1::new({
                                                                                                                                                     let go_tco
                                                                                                                                                         =
                                                                                                                                                         go_tco.clone();
                                                                                                                                                     let step_1
                                                                                                                                                         =
                                                                                                                                                         step_1.clone();
                                                                                                                                                     move
                                                                                                                                                         |rest_1|
                                                                                                                                                         {
                                                                                                                                                             let matchValue_3 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&s_1);
                                                                                                                                                             let matchValue_4 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&e_1);
                                                                                                                                                             let matchValue_5 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&step_1);
                                                                                                                                                             let matchValue_6 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(rest_1);
                                                                                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                                             &&&matchValue_3),
                                                                                                                                                                                                                          &&&matchValue_4))
                                                                                                                                                                {
                                                                                                                                                                 &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue_3,
                                                                                                                                                                                                                                 &matchValue_6))
                                                                                                                                                             } else {
                                                                                                                                                                 if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                                    {
                                                                                                                                                                     let step1_3 =
                                                                                                                                                                         matchValue_5;
                                                                                                                                                                     let s1_3 =
                                                                                                                                                                         matchValue_3;
                                                                                                                                                                     go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                &&&s1_3),
                                                                                                                                                                                                             &&&step1_3))(&matchValue_4)(&step1_3)(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&s1_3,
                                                                                                                                                                                                                                                                                                                   &matchValue_6)))
                                                                                                                                                                 } else {
                                                                                                                                                                     panic!("{}",
                                                                                                                                                                            LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 113_i32,
                                  Data2: 273_i32,}).get_Message(),)
                                                                                                                                                                 }
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                 })
                                                                                                                              })
                                                                                                           })
                                                                                            }
                                                                                            let go =
                                                                                                go_1.Value;
                                                                                            go_tco(&end1_3)(&start1_3)({
                                                                                                                           let matchValue_8 =
                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                            &&&start1_3),
                                                                                                                                                                                         &&&end1_3));
                                                                                                                           match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                            &matchValue_8)
                                                                                                                               {
                                                                                                                               0_i32
                                                                                                                               =>
                                                                                                                               &1_i32,
                                                                                                                               _
                                                                                                                               =>
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                &&&1_i32),
                                                                                                                           }
                                                                                                                       })(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))
                                                                                        }
                                                                                    } else {
                                                                                        panic!("{}",
                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 111_i32,
                                  Data2: 84_i32,}).get_Message(),)
                                                                                    }
                                                                                }
                                                                            }
                                                                    })))
    }
    pub fn Data_List_partition() -> &dyn Any {
        static Data_List_partition: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_partition.get_or_init(||
                                            &Func1::new(move |p|
                                                            &Func1::new({
                                                                            let p
                                                                                =
                                                                                p.clone();
                                                                            move
                                                                                |xs|
                                                                                {
                                                                                    let select =
                                                                                        &Func1::new(move
                                                                                                        |x|
                                                                                                        &Func1::new({
                                                                                                                        let x
                                                                                                                            =
                                                                                                                            x.clone();
                                                                                                                        move
                                                                                                                            |v|
                                                                                                                            {
                                                                                                                                let matchValue =
                                                                                                                                    Sharpurs_Prelude::unbox(&&x);
                                                                                                                                let matchValue_1 =
                                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                                {
                                                                                                                                    let activePatternResult =
                                                                                                                                        Sharpurs_Prelude::_007cHasProp_007c__007c(string("no"),
                                                                                                                                                                                  &matchValue_1);
                                                                                                                                    if activePatternResult.is_some()
                                                                                                                                       {
                                                                                                                                        let activePatternResult_1 =
                                                                                                                                            Sharpurs_Prelude::_007cHasProp_007c__007c(string("yes"),
                                                                                                                                                                                      &matchValue_1);
                                                                                                                                        if activePatternResult_1.is_some()
                                                                                                                                           {
                                                                                                                                            let no =
                                                                                                                                                getValue(activePatternResult);
                                                                                                                                            let yes =
                                                                                                                                                getValue(activePatternResult_1);
                                                                                                                                            let matchValue_3 =
                                                                                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                          &&&matchValue));
                                                                                                                                            match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                             &matchValue_3)
                                                                                                                                                {
                                                                                                                                                0_i32
                                                                                                                                                =>
                                                                                                                                                &add(string("no"),
                                                                                                                                                     &&no,
                                                                                                                                                     add(string("yes"),
                                                                                                                                                         &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue,
                                                                                                                                                                                                                          &yes)),
                                                                                                                                                         empty::<string,
                                                                                                                                                                 &dyn Any>())),
                                                                                                                                                _
                                                                                                                                                =>
                                                                                                                                                &add(string("no"),
                                                                                                                                                     &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&matchValue,
                                                                                                                                                                                                                      &no)),
                                                                                                                                                     add(string("yes"),
                                                                                                                                                         &&yes,
                                                                                                                                                         empty::<string,
                                                                                                                                                                 &dyn Any>())),
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            panic!("{}",
                                                                                                                                                   LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 119_i32,
                                  Data2: 137_i32,}).get_Message(),)
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        panic!("{}",
                                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 119_i32,
                                  Data2: 137_i32,}).get_Message(),)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                    }));
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                              &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                           &&&select),
                                                                                                                                                        &&&add(string("no"),
                                                                                                                                                               &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
                                                                                                                                                               add(string("yes"),
                                                                                                                                                                   &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
                                                                                                                                                                   empty::<string,
                                                                                                                                                                           &dyn Any>()))),
                                                                                                                     xs)
                                                                                }
                                                                        })))
    }
    pub fn Data_List_null() -> &dyn Any {
        static Data_List_null: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_null.get_or_init(||
                                       &Func1::new(move |v|
                                                       if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                              =
                                                              Sharpurs_Prelude::unbox(v).as_ref()
                                                          {
                                                           &true
                                                       } else { &false }))
    }
    pub fn Data_List_nubBy() -> &dyn Any {
        static Data_List_nubBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_nubBy.get_or_init(||
                                        &Func1::new(move |p|
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
                                                                                                       |v|
                                                                                                       Func1::new({
                                                                                                                      let go_tco
                                                                                                                          =
                                                                                                                          go_tco.clone();
                                                                                                                      let v
                                                                                                                          =
                                                                                                                          v.clone();
                                                                                                                      move
                                                                                                                          |v1|
                                                                                                                          Func1::new({
                                                                                                                                         let go_tco
                                                                                                                                             =
                                                                                                                                             go_tco.clone();
                                                                                                                                         let v1
                                                                                                                                             =
                                                                                                                                             v1.clone();
                                                                                                                                         move
                                                                                                                                             |v2|
                                                                                                                                             go_tco(v)(v1)(v2.clone())
                                                                                                                                     })
                                                                                                                  })
                                                                                               })
                                                                           });
                                                            let go_1 =
                                                                Lazy(go_2);
                                                            let go_tco =
                                                                Func1::new({
                                                                               let p
                                                                                   =
                                                                                   p.clone();
                                                                               move
                                                                                   |v_1|
                                                                                   fix1(&(move
                                                                                              |go_tco,
                                                                                               v_1|
                                                                                              Func1::new({
                                                                                                             let go_tco
                                                                                                                 =
                                                                                                                 go_tco.clone();
                                                                                                             let v_1
                                                                                                                 =
                                                                                                                 v_1.clone();
                                                                                                             move
                                                                                                                 |v1_1|
                                                                                                                 Func1::new({
                                                                                                                                let go_tco
                                                                                                                                    =
                                                                                                                                    go_tco.clone();
                                                                                                                                let v1_1
                                                                                                                                    =
                                                                                                                                    v1_1.clone();
                                                                                                                                move
                                                                                                                                    |v2_1|
                                                                                                                                    {
                                                                                                                                        let matchValue =
                                                                                                                                            Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                        let matchValue_1 =
                                                                                                                                            Sharpurs_Prelude::unbox(&&v1_1);
                                                                                                                                        let matchValue_2:
                                                                                                                                                LrcPtr<Data_List_Types_List> =
                                                                                                                                            Sharpurs_Prelude::unbox(v2_1);
                                                                                                                                        match matchValue_2.as_ref()
                                                                                                                                            {
                                                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                                                                                                                                               matchValue_2_1_1)
                                                                                                                                            =>
                                                                                                                                            {
                                                                                                                                                let as_var =
                                                                                                                                                    matchValue_2_1_1.clone();
                                                                                                                                                let acc_1 =
                                                                                                                                                    matchValue_1;
                                                                                                                                                let a =
                                                                                                                                                    matchValue_2_1_0.clone();
                                                                                                                                                let matchValue_4 =
                                                                                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Internal::Data_List_Internal_insertAndLookupBy(),
                                                                                                                                                                                                                                                                                     &&&p),
                                                                                                                                                                                                                                                  &&&a),
                                                                                                                                                                                                               &&&matchValue));
                                                                                                                                                {
                                                                                                                                                    let activePatternResult =
                                                                                                                                                        Sharpurs_Prelude::_007cHasProp_007c__007c(string("found"),
                                                                                                                                                                                                  &matchValue_4);
                                                                                                                                                    if activePatternResult.is_some()
                                                                                                                                                       {
                                                                                                                                                        let activePatternResult_1 =
                                                                                                                                                            Sharpurs_Prelude::_007cHasProp_007c__007c(string("result"),
                                                                                                                                                                                                      &matchValue_4);
                                                                                                                                                        if activePatternResult_1.is_some()
                                                                                                                                                           {
                                                                                                                                                            let found =
                                                                                                                                                                getValue(activePatternResult);
                                                                                                                                                            let s_prime =
                                                                                                                                                                getValue(activePatternResult_1);
                                                                                                                                                            let matchValue_5 =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&found);
                                                                                                                                                            match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                             &matchValue_5)
                                                                                                                                                                {
                                                                                                                                                                0_i32
                                                                                                                                                                =>
                                                                                                                                                                go_tco(&s_prime)(&acc_1)(&as_var),
                                                                                                                                                                _
                                                                                                                                                                =>
                                                                                                                                                                go_tco(&s_prime)(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&a,
                                                                                                                                                                                                                                                 &acc_1)))(&as_var),
                                                                                                                                                            }
                                                                                                                                                        } else {
                                                                                                                                                            panic!("{}",
                                                                                                                                                                   LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 125_i32,
                                  Data2: 624_i32,}).get_Message(),)
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        panic!("{}",
                                                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 125_i32,
                                  Data2: 624_i32,}).get_Message(),)
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            _
                                                                                                                                            =>
                                                                                                                                            &matchValue_1,
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                            })
                                                                                                         })),
                                                                                        v_1.clone())
                                                                           });
                                                            let go =
                                                                go_1.Value;
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                &&&PureScript_Data_List::Data_List_reverse()),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                                   &&&PureScript_Data_List_Internal::Data_List_Internal_emptySet()),
                                                                                                                                &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)))
                                                        }))
    }
    pub fn Data_List_nub() -> &dyn Any {
        static Data_List_nub: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_nub.get_or_init(||
                                      &Func1::new(move |dictOrd|
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_nubBy(),
                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                          dictOrd))))
    }
    pub fn Data_List_newtypePattern() -> &dyn Any {
        static Data_List_newtypePattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_newtypePattern.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                  &&&add(string("Coercible0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &Sharpurs_Prelude::Prim_undefined()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>())))
    }
    pub fn Data_List_mapMaybe() -> &dyn Any {
        static Data_List_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_mapMaybe.get_or_init(||
                                           &Func1::new(move |f|
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
                                                                                                          |v|
                                                                                                          Func1::new({
                                                                                                                         let go_tco
                                                                                                                             =
                                                                                                                             go_tco.clone();
                                                                                                                         let v
                                                                                                                             =
                                                                                                                             v.clone();
                                                                                                                         move
                                                                                                                             |v1|
                                                                                                                             go_tco(v)(v1.clone())
                                                                                                                     })
                                                                                                  })
                                                                              });
                                                               let go_1 =
                                                                   Lazy(go_2);
                                                               let go_tco =
                                                                   Func1::new({
                                                                                  let f
                                                                                      =
                                                                                      f.clone();
                                                                                  move
                                                                                      |v_1|
                                                                                      fix1(&(move
                                                                                                 |go_tco,
                                                                                                  v_1|
                                                                                                 Func1::new({
                                                                                                                let go_tco
                                                                                                                    =
                                                                                                                    go_tco.clone();
                                                                                                                let v_1
                                                                                                                    =
                                                                                                                    v_1.clone();
                                                                                                                move
                                                                                                                    |v1_1|
                                                                                                                    {
                                                                                                                        let matchValue =
                                                                                                                            Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                        let matchValue_1:
                                                                                                                                LrcPtr<Data_List_Types_List> =
                                                                                                                            Sharpurs_Prelude::unbox(v1_1);
                                                                                                                        match matchValue_1.as_ref()
                                                                                                                            {
                                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                               matchValue_1_1_1)
                                                                                                                            =>
                                                                                                                            {
                                                                                                                                let xs =
                                                                                                                                    matchValue_1_1_1.clone();
                                                                                                                                let acc_1 =
                                                                                                                                    matchValue;
                                                                                                                                let matchValue_3:
                                                                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                               &&matchValue_1_1_0));
                                                                                                                                match matchValue_3.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_3_1_0)
                                                                                                                                    =>
                                                                                                                                    go_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_3_1_0,
                                                                                                                                                                                                           &acc_1)))(&xs),
                                                                                                                                    _
                                                                                                                                    =>
                                                                                                                                    go_tco(&acc_1)(&xs),
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _
                                                                                                                            =>
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                             &&&matchValue),
                                                                                                                        }
                                                                                                                    }
                                                                                                            })),
                                                                                           v_1.clone())
                                                                              });
                                                               let go =
                                                                   go_1.Value;
                                                               Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))
                                                           }))
    }
    pub fn Data_List_manyRec() -> &dyn Any {
        static Data_List_manyRec: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_manyRec.get_or_init(||
                                          &Func1::new(move |dictMonadRec|
                                                          {
                                                              let Bind1 =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonadRec)),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                              &Func1::new({
                                                                              let Bind1
                                                                                  =
                                                                                  Bind1.clone();
                                                                              let dictMonadRec
                                                                                  =
                                                                                  dictMonadRec.clone();
                                                                              move
                                                                                  |dictAlternative|
                                                                                  {
                                                                                      let Plus1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let Alt0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                  Sharpurs_Prelude::unbox(&&Plus1)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let Functor0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                  Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&Plus1)),
                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let Applicative0 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                  Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                                      let pure_var =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      &Func1::new({
                                                                                                      let Alt0
                                                                                                          =
                                                                                                          Alt0.clone();
                                                                                                      let Applicative0
                                                                                                          =
                                                                                                          Applicative0.clone();
                                                                                                      let Functor0
                                                                                                          =
                                                                                                          Functor0.clone();
                                                                                                      let pure_var
                                                                                                          =
                                                                                                          pure_var.clone();
                                                                                                      move
                                                                                                          |p|
                                                                                                          {
                                                                                                              let go =
                                                                                                                  &Func1::new({
                                                                                                                                  let p
                                                                                                                                      =
                                                                                                                                      p.clone();
                                                                                                                                  move
                                                                                                                                      |acc|
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                             &&&Bind1),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                                                                                                                                   &&&Alt0),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                         &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                                                                                   &&&p)),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                                                   &&&Applicative0),
                                                                                                                                                                                                                                                                                &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit()))))),
                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                         let acc
                                                                                                                                                                                             =
                                                                                                                                                                                             acc.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |aa|
                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                 &&&pure_var),
                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_bifunctorStep()),
                                                                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                         |v|
                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(v.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                         &acc)))),
                                                                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                      |v_1|
                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                                                                                                                                                                                                                       &&&acc))),
                                                                                                                                                                                                                                                                 aa))
                                                                                                                                                                                     }))
                                                                                                                              });
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM(),
                                                                                                                                                                                                                     &&&dictMonadRec),
                                                                                                                                                                                  &&&go),
                                                                                                                                               &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))
                                                                                                          }
                                                                                                  })
                                                                                  }
                                                                          })
                                                          }))
    }
    pub fn Data_List_someRec() -> &dyn Any {
        static Data_List_someRec: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_someRec.get_or_init(||
                                          &Func1::new(move |dictMonadRec|
                                                          &Func1::new({
                                                                          let dictMonadRec
                                                                              =
                                                                              dictMonadRec.clone();
                                                                          move
                                                                              |dictAlternative|
                                                                              {
                                                                                  let Apply0 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                              Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                  let Functor0 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                              Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined());
                                                                                  &Func1::new({
                                                                                                  let Apply0
                                                                                                      =
                                                                                                      Apply0.clone();
                                                                                                  let Functor0
                                                                                                      =
                                                                                                      Functor0.clone();
                                                                                                  let dictAlternative
                                                                                                      =
                                                                                                      dictAlternative.clone();
                                                                                                  move
                                                                                                      |v|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                             &&&Apply0),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                   &&&Functor0),
                                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                                  |usd__arg1|
                                                                                                                                                                                                                                                                  Func1::new({
                                                                                                                                                                                                                                                                                 let usd__arg1
                                                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                                                     usd__arg1.clone();
                                                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                                                     |usd__arg2|
                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                     usd__arg2.clone()))
                                                                                                                                                                                                                                                                             }))),
                                                                                                                                                                                                             v)),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_manyRec(),
                                                                                                                                                                                                                                                &&&dictMonadRec),
                                                                                                                                                                                                             &&&dictAlternative),
                                                                                                                                                                          v))
                                                                                              })
                                                                              }
                                                                      })))
    }
    pub fn Data_List_some_0040148() -> &dyn Any {
        &Func1::new(move |dictAlternative|
                        PureScript_Data_List::Data_List_some_tco(dictAlternative))
    }
    pub fn Data_List_some_0040148_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_some_0040148_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_some_0040148_002d1.get_or_init(||
                                                     Lazy(Data_List_some_0040148.clone()))
    }
    pub fn Data_List_many_0040150() -> &dyn Any {
        &Func1::new(move |dictAlternative|
                        PureScript_Data_List::Data_List_many_tco(dictAlternative))
    }
    pub fn Data_List_many_0040150_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_many_0040150_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_many_0040150_002d1.get_or_init(||
                                                     Lazy(Data_List_many_0040150.clone()))
    }
    pub fn Data_List_some_tco(dictAlternative: &dyn Any) -> &dyn Any {
        let Apply0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Apply0 = Apply0.clone();
                        let Data_List_many_0040150_002d1 =
                            Data_List_many_0040150_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictAlternative = dictAlternative.clone();
                        move |dictLazy|
                            &Func1::new({
                                            let Data_List_many_0040150_002d1 =
                                                Data_List_many_0040150_002d1.clone();
                                            let dictLazy = dictLazy.clone();
                                            move |v|
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                       &&&Apply0),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                             &&&Functor0),
                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                            |usd__arg1|
                                                                                                                                                                                                            Func1::new({
                                                                                                                                                                                                                           let usd__arg1
                                                                                                                                                                                                                               =
                                                                                                                                                                                                                               usd__arg1.clone();
                                                                                                                                                                                                                           move
                                                                                                                                                                                                                               |usd__arg2|
                                                                                                                                                                                                                               &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                               usd__arg2.clone()))
                                                                                                                                                                                                                       }))),
                                                                                                                                                       v)),
                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                                       &&&dictLazy),
                                                                                                                    &&&Func1::new({
                                                                                                                                      let Data_List_many_0040150_002d1
                                                                                                                                          =
                                                                                                                                          Data_List_many_0040150_002d1.clone();
                                                                                                                                      let v
                                                                                                                                          =
                                                                                                                                          v.clone();
                                                                                                                                      move
                                                                                                                                          |v1|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_many_0040150_002d1.Value,
                                                                                                                                                                                                                                                 &&&dictAlternative),
                                                                                                                                                                                                              &&&dictLazy),
                                                                                                                                                                           &&&v)
                                                                                                                                  })))
                                        })
                    })
    }
    pub fn Data_List_some() -> &dyn Any {
        static Data_List_some: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_some.get_or_init(|| Data_List_some_0040148_002d1.Value)
    }
    pub fn Data_List_many_tco(dictAlternative: &dyn Any) -> &dyn Any {
        let Alt0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictAlternative)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Alt0 = Alt0.clone();
                        let Applicative0 = Applicative0.clone();
                        let Data_List_some_0040148_002d1 =
                            Data_List_some_0040148_002d1.clone();
                        let dictAlternative = dictAlternative.clone();
                        move |dictLazy|
                            &Func1::new({
                                            let Data_List_some_0040148_002d1 =
                                                Data_List_some_0040148_002d1.clone();
                                            let dictLazy = dictLazy.clone();
                                            move |v|
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                       &&&Alt0),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_some_0040148_002d1.Value,
                                                                                                                                                                                                                             &&&dictAlternative),
                                                                                                                                                                                          &&&dictLazy),
                                                                                                                                                       v)),
                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                       &&&Applicative0),
                                                                                                                    &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)))
                                        })
                    })
    }
    pub fn Data_List_many() -> &dyn Any {
        static Data_List_many: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_many.get_or_init(|| Data_List_many_0040150_002d1.Value)
    }
    pub fn Data_List_length() -> &dyn Any {
        static Data_List_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_length.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                             &&&Func1::new(move
                                                                                                                               |acc|
                                                                                                                               &Func1::new({
                                                                                                                                               let acc
                                                                                                                                                   =
                                                                                                                                                   acc.clone();
                                                                                                                                               move
                                                                                                                                                   |v|
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                          &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                       &&&acc),
                                                                                                                                                                                    &&&1_i32)
                                                                                                                                           }))),
                                                                          &&&0_i32))
    }
    pub fn Data_List_last_0040156() -> &dyn Any {
        &Func1::new(move |v| PureScript_Data_List::Data_List_last_tco(v))
    }
    pub fn Data_List_last_0040156_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_last_0040156_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_last_0040156_002d1.get_or_init(||
                                                     Lazy(Data_List_last_0040156.clone()))
    }
    pub fn Data_List_last_tco(v: &dyn Any) -> &dyn Any {
        let matchValue: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v);
        if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                  matchValue_1_1)
               = matchValue.as_ref() {
            if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor =
                   Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                          {
                                                          Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                             x)
                                                          => x.clone(),
                                                          _ => unreachable!(),
                                                      }).as_ref() {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue.as_ref()
                                                                            {
                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                               _)
                                                                            =>
                                                                            x.clone(),
                                                                            _
                                                                            =>
                                                                            unreachable!(),
                                                                        }))
            } else {
                Sharpurs_Prelude::sharpurs_apply(&&&Data_List_last_0040156_002d1.Value,
                                                 &&&match matchValue.as_ref()
                                                        {
                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                           x)
                                                        => x.clone(),
                                                        _ => unreachable!(),
                                                    })
            }
        } else { &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor) }
    }
    pub fn Data_List_last() -> &dyn Any {
        static Data_List_last: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_last.get_or_init(|| Data_List_last_0040156_002d1.Value)
    }
    pub fn Data_List_insertBy_0040160() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           Func1::new({
                                                          let v1 = v1.clone();
                                                          move |v2|
                                                              PureScript_Data_List::Data_List_insertBy_tco(&v,
                                                                                                           &v1,
                                                                                                           v2)
                                                      })
                                   }))
    }
    pub fn Data_List_insertBy_0040160_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_insertBy_0040160_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_insertBy_0040160_002d1.get_or_init(||
                                                         Lazy(Data_List_insertBy_0040160.clone()))
    }
    pub fn Data_List_insertBy_tco(v: &dyn Any, v1: &dyn Any, v2: &dyn Any)
     -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        let matchValue_2: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v2);
        match matchValue_2.as_ref() {
            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                               matchValue_2_1_1)
            => {
                let y = matchValue_2_1_0.clone();
                let x_1 = matchValue_1;
                let cmp = matchValue;
                if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor =
                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                     &&&x_1),
                                                                                  &&&y)).as_ref()
                   {
                    &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&y,
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_insertBy_0040160_002d1.Value,
                                                                                                                                                                                           &&&cmp),
                                                                                                                                                        &&&x_1),
                                                                                                                     &&matchValue_2_1_1)))
                } else {
                    &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&x_1,
                                                                                    matchValue_2.clone()))
                }
            }
            _ =>
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_singleton(),
                                             &&&matchValue_1),
        }
    }
    pub fn Data_List_insertBy() -> &dyn Any {
        static Data_List_insertBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_insertBy.get_or_init(||
                                           Data_List_insertBy_0040160_002d1.Value)
    }
    pub fn Data_List_insertAt_0040164() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           Func1::new({
                                                          let v1 = v1.clone();
                                                          move |v2|
                                                              PureScript_Data_List::Data_List_insertAt_tco(&v,
                                                                                                           &v1,
                                                                                                           v2)
                                                      })
                                   }))
    }
    pub fn Data_List_insertAt_0040164_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_insertAt_0040164_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_insertAt_0040164_002d1.get_or_init(||
                                                         Lazy(Data_List_insertAt_0040164.clone()))
    }
    pub fn Data_List_insertAt_tco(v: &dyn Any, v1: &dyn Any, v2: &dyn Any)
     -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        let matchValue_2: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v2);
        match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32, &matchValue) {
            0_i32 =>
            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(x.clone(),
                                                                                                                                   xs.clone())))),
            _ =>
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                      matchValue_2_1_1)
                   = matchValue_2.as_ref() {
                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                    &&&Func1::new({
                                                                                                      let matchValue_2
                                                                                                          =
                                                                                                          matchValue_2.clone();
                                                                                                      move
                                                                                                          |v3|
                                                                                                          &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                  _)
                                                                                                                                                                               =>
                                                                                                                                                                               x.clone(),
                                                                                                                                                                               _
                                                                                                                                                                               =>
                                                                                                                                                                               unreachable!(),
                                                                                                                                                                           },
                                                                                                                                                                          v3.clone()))
                                                                                                  })),
                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_insertAt_0040164_002d1.Value,
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                             &&&1_i32)),
                                                                                                                       &&&matchValue_1),
                                                                                    &&&match matchValue_2.as_ref()
                                                                                           {
                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                              x)
                                                                                           =>
                                                                                           x.clone(),
                                                                                           _
                                                                                           =>
                                                                                           unreachable!(),
                                                                                       }))
            } else {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
            },
        }
    }
    pub fn Data_List_insertAt() -> &dyn Any {
        static Data_List_insertAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_insertAt.get_or_init(||
                                           Data_List_insertAt_0040164_002d1.Value)
    }
    pub fn Data_List_insert() -> &dyn Any {
        static Data_List_insert: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_insert.get_or_init(||
                                         &Func1::new(move |dictOrd|
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_insertBy(),
                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                             dictOrd))))
    }
    pub fn Data_List_init() -> &dyn Any {
        static Data_List_init: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_init.get_or_init(||
                                       &Func1::new(move |lst|
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                              &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                           &&&Func1::new(move
                                                                                                                                             |v|
                                                                                                                                             find(string("init"),
                                                                                                                                                  Sharpurs_Prelude::unbox(v)))),
                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_unsnoc(),
                                                                                                                           lst))))
    }
    pub fn Data_List_index_0040172() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Data_List::Data_List_index_tco(&v,
                                                                                     v1)
                                   }))
    }
    pub fn Data_List_index_0040172_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_index_0040172_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_index_0040172_002d1.get_or_init(||
                                                      Lazy(Data_List_index_0040172.clone()))
    }
    pub fn Data_List_index_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let matchValue: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                  matchValue_1_1)
               = matchValue.as_ref() {
            if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                        &matchValue_1).is_some()
               {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue.as_ref()
                                                                            {
                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                               _)
                                                                            =>
                                                                            x.clone(),
                                                                            _
                                                                            =>
                                                                            unreachable!(),
                                                                        }))
            } else {
                if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                          matchValue_1_1)
                       = matchValue.as_ref() {
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_index_0040172_002d1.Value,
                                                                                        &&&match matchValue.as_ref()
                                                                                               {
                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                  x)
                                                                                               =>
                                                                                               x.clone(),
                                                                                               _
                                                                                               =>
                                                                                               unreachable!(),
                                                                                           }),
                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                              &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                           &&&matchValue_1),
                                                                                        &&&1_i32))
                } else { panic!("{}", string("Match failure"),) }
            }
        } else { &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor) }
    }
    pub fn Data_List_index() -> &dyn Any {
        static Data_List_index: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_index.get_or_init(|| Data_List_index_0040172_002d1.Value)
    }
    pub fn Data_List_head() -> &dyn Any {
        static Data_List_head: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_head.get_or_init(||
                                       &Func1::new(move |v|
                                                       {
                                                           let matchValue:
                                                                   LrcPtr<Data_List_Types_List> =
                                                               Sharpurs_Prelude::unbox(v);
                                                           match matchValue.as_ref()
                                                               {
                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                  matchValue_1_1)
                                                               =>
                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)),
                                                               _ =>
                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                           }
                                                       }))
    }
    pub fn Data_List_transpose_0040178() -> &dyn Any {
        &Func1::new(move |v| PureScript_Data_List::Data_List_transpose_tco(v))
    }
    pub fn Data_List_transpose_0040178_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_transpose_0040178_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_transpose_0040178_002d1.get_or_init(||
                                                          Lazy(Data_List_transpose_0040178.clone()))
    }
    pub fn Data_List_transpose_tco(v: &dyn Any) -> &dyn Any {
        let matchValue: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v);
        if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                  matchValue_1_1)
               = matchValue.as_ref() {
            if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor =
                   Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                          {
                                                          Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                             _)
                                                          => x.clone(),
                                                          _ => unreachable!(),
                                                      }).as_ref() {
                Sharpurs_Prelude::sharpurs_apply(&&&Data_List_transpose_0040178_002d1.Value,
                                                 &&&match matchValue.as_ref()
                                                        {
                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                           x)
                                                        => x.clone(),
                                                        _ => unreachable!(),
                                                    })
            } else {
                let activePatternResult_1: LrcPtr<Data_List_Types_List> =
                    Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                           {
                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                              _)
                                                           => x.clone(),
                                                           _ =>
                                                           unreachable!(),
                                                       });
                if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(activePatternResult_1_1_0,
                                                                          activePatternResult_1_1_1)
                       = activePatternResult_1.as_ref() {
                    &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                         {
                                                                                                                                                         Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                            _)
                                                                                                                                                         =>
                                                                                                                                                         x.clone(),
                                                                                                                                                         _
                                                                                                                                                         =>
                                                                                                                                                         unreachable!(),
                                                                                                                                                     },
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_mapMaybe(),
                                                                                                                                                                                                                        &&&PureScript_Data_List::Data_List_head()),
                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                               x)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        }))),
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&Data_List_transpose_0040178_002d1.Value,
                                                                                                                     &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                               x)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                            _
                                                                                                                                                                                            =>
                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                        },
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_mapMaybe(),
                                                                                                                                                                                                                                                           &&&PureScript_Data_List::Data_List_tail()),
                                                                                                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                               _
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                           }))))))
                } else {
                    panic!("{}",
                           string("Match failure: PureScript_Data_List_Types.Data_List_Types_List"),)
                }
            }
        } else {
            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)
        }
    }
    pub fn Data_List_transpose() -> &dyn Any {
        static Data_List_transpose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_transpose.get_or_init(||
                                            Data_List_transpose_0040178_002d1.Value)
    }
    pub fn Data_List_groupBy_0040182() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Data_List::Data_List_groupBy_tco(&v,
                                                                                       v1)
                                   }))
    }
    pub fn Data_List_groupBy_0040182_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_groupBy_0040182_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_groupBy_0040182_002d1.get_or_init(||
                                                        Lazy(Data_List_groupBy_0040182.clone()))
    }
    pub fn Data_List_groupBy_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v1);
        match matchValue_1.as_ref() {
            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                               matchValue_1_1_1)
            => {
                let x = matchValue_1_1_0.clone();
                let eq = matchValue;
                let matchValue_3 =
                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_span(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&eq,
                                                                                                                                                     &&&x)),
                                                                               &&matchValue_1_1_1));
                {
                    let activePatternResult =
                        Sharpurs_Prelude::_007cHasProp_007c__007c(string("init"),
                                                                  &matchValue_3);
                    if activePatternResult.is_some() {
                        let activePatternResult_1 =
                            Sharpurs_Prelude::_007cHasProp_007c__007c(string("rest"),
                                                                      &matchValue_3);
                        if activePatternResult_1.is_some() {
                            let ys = getValue(activePatternResult);
                            let zs = getValue(activePatternResult_1);
                            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                             &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&x,
                                                                                                                                                                                                   &ys))),
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_groupBy_0040182_002d1.Value,
                                                                                                                                                                &&&eq),
                                                                                                                             &&&zs)))
                        } else {
                            panic!("{}",
                                   LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 181_i32,
                                  Data2: 400_i32,}).get_Message(),)
                        }
                    } else {
                        panic!("{}",
                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.fs"),
                                  Data1: 181_i32,
                                  Data2: 400_i32,}).get_Message(),)
                    }
                }
            }
            _ =>
            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
        }
    }
    pub fn Data_List_groupBy() -> &dyn Any {
        static Data_List_groupBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_groupBy.get_or_init(||
                                          Data_List_groupBy_0040182_002d1.Value)
    }
    pub fn Data_List_groupAllBy() -> &dyn Any {
        static Data_List_groupAllBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_groupAllBy.get_or_init(||
                                             &Func1::new(move |p|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_groupBy(),
                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                      let p
                                                                                                                                                                                          =
                                                                                                                                                                                          p.clone();
                                                                                                                                                                                      move
                                                                                                                                                                                          |x|
                                                                                                                                                                                          &Func1::new({
                                                                                                                                                                                                          let x
                                                                                                                                                                                                              =
                                                                                                                                                                                                              x.clone();
                                                                                                                                                                                                          move
                                                                                                                                                                                                              |y|
                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                                                                                                                                                                        &&&x),
                                                                                                                                                                                                                                                                                                                     y)),
                                                                                                                                                                                                                                               &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))
                                                                                                                                                                                                      })
                                                                                                                                                                                  }))),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_sortBy(),
                                                                                                                                 p))))
    }
    pub fn Data_List_group() -> &dyn Any {
        static Data_List_group: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_group.get_or_init(||
                                        &Func1::new(move |dictEq|
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_groupBy(),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                            dictEq))))
    }
    pub fn Data_List_groupAll() -> &dyn Any {
        static Data_List_groupAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_groupAll.get_or_init(||
                                           &Func1::new(move |dictOrd|
                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_group(),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_sort(),
                                                                                                                               dictOrd))))
    }
    pub fn Data_List_fromFoldable() -> &dyn Any {
        static Data_List_fromFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_fromFoldable.get_or_init(||
                                               &Func1::new(move |dictFoldable|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                      dictFoldable),
                                                                                                                                   &&&Func1::new(move
                                                                                                                                                     |usd__arg1|
                                                                                                                                                     Func1::new({
                                                                                                                                                                    let usd__arg1
                                                                                                                                                                        =
                                                                                                                                                                        usd__arg1.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |usd__arg2|
                                                                                                                                                                        &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                        usd__arg2.clone()))
                                                                                                                                                                }))),
                                                                                                &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))))
    }
    pub fn Data_List_foldM_0040194() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Data_List::Data_List_foldM_tco(dictMonad))
    }
    pub fn Data_List_foldM_0040194_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_foldM_0040194_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_foldM_0040194_002d1.get_or_init(||
                                                      Lazy(Data_List_foldM_0040194.clone()))
    }
    pub fn Data_List_foldM_tco(dictMonad: &dyn Any) -> &dyn Any {
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Data_List_foldM_0040194_002d1 =
                            Data_List_foldM_0040194_002d1.clone();
                        let dictMonad = dictMonad.clone();
                        move |v|
                            &Func1::new({
                                            let Data_List_foldM_0040194_002d1
                                                =
                                                Data_List_foldM_0040194_002d1.clone();
                                            let v = v.clone();
                                            move |v1|
                                                &Func1::new({
                                                                let Data_List_foldM_0040194_002d1
                                                                    =
                                                                    Data_List_foldM_0040194_002d1.clone();
                                                                let v1 =
                                                                    v1.clone();
                                                                move |v2|
                                                                    {
                                                                        let matchValue =
                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                        let matchValue_1 =
                                                                            Sharpurs_Prelude::unbox(&&v1);
                                                                        let matchValue_2:
                                                                                LrcPtr<Data_List_Types_List> =
                                                                            Sharpurs_Prelude::unbox(v2);
                                                                        match matchValue_2.as_ref()
                                                                            {
                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                                                                               matchValue_2_1_1)
                                                                            =>
                                                                            {
                                                                                let f =
                                                                                    matchValue;
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                       &&&Bind1),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                          &&&matchValue_1),
                                                                                                                                                                                       &&matchValue_2_1_0)),
                                                                                                                 &&&Func1::new({
                                                                                                                                   let Data_List_foldM_0040194_002d1
                                                                                                                                       =
                                                                                                                                       Data_List_foldM_0040194_002d1.clone();
                                                                                                                                   let f
                                                                                                                                       =
                                                                                                                                       f.clone();
                                                                                                                                   let matchValue_2
                                                                                                                                       =
                                                                                                                                       matchValue_2.clone();
                                                                                                                                   move
                                                                                                                                       |b_prime|
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_foldM_0040194_002d1.Value,
                                                                                                                                                                                                                                                                                 &&&dictMonad),
                                                                                                                                                                                                                                              &&&f),
                                                                                                                                                                                                           b_prime),
                                                                                                                                                                        &&&match matchValue_2.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                  x)
                                                                                                                                                                               =>
                                                                                                                                                                               x.clone(),
                                                                                                                                                                               _
                                                                                                                                                                               =>
                                                                                                                                                                               unreachable!(),
                                                                                                                                                                           })
                                                                                                                               }))
                                                                            }
                                                                            _
                                                                            =>
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                &&&Applicative0),
                                                                                                             &&&matchValue_1),
                                                                        }
                                                                    }
                                                            })
                                        })
                    })
    }
    pub fn Data_List_foldM() -> &dyn Any {
        static Data_List_foldM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_foldM.get_or_init(|| Data_List_foldM_0040194_002d1.Value)
    }
    pub fn Data_List_findIndex() -> &dyn Any {
        static Data_List_findIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_findIndex.get_or_init(||
                                            &Func1::new(move |r#fn|
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
                                                                                                           |v|
                                                                                                           Func1::new({
                                                                                                                          let go_tco
                                                                                                                              =
                                                                                                                              go_tco.clone();
                                                                                                                          let v
                                                                                                                              =
                                                                                                                              v.clone();
                                                                                                                          move
                                                                                                                              |v1|
                                                                                                                              go_tco(v)(v1.clone())
                                                                                                                      })
                                                                                                   })
                                                                               });
                                                                let go_1 =
                                                                    Lazy(go_2);
                                                                let go_tco =
                                                                    Func1::new({
                                                                                   let r#fn
                                                                                       =
                                                                                       r#fn.clone();
                                                                                   move
                                                                                       |v_1|
                                                                                       fix1(&(move
                                                                                                  |go_tco,
                                                                                                   v_1|
                                                                                                  Func1::new({
                                                                                                                 let go_tco
                                                                                                                     =
                                                                                                                     go_tco.clone();
                                                                                                                 let v_1
                                                                                                                     =
                                                                                                                     v_1.clone();
                                                                                                                 move
                                                                                                                     |v1_1|
                                                                                                                     {
                                                                                                                         let matchValue =
                                                                                                                             Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                         let matchValue_1:
                                                                                                                                 LrcPtr<Data_List_Types_List> =
                                                                                                                             Sharpurs_Prelude::unbox(v1_1);
                                                                                                                         if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                                =
                                                                                                                                matchValue_1.as_ref()
                                                                                                                            {
                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                         } else {
                                                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&r#fn,
                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                 Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                 =>
                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                 _
                                                                                                                                                                                                 =>
                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                             }))
                                                                                                                                {
                                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&matchValue))
                                                                                                                             } else {
                                                                                                                                 if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                    {
                                                                                                                                     go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                             &&&1_i32))(&match matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                x)
                                                                                                                                                                                             =>
                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                             _
                                                                                                                                                                                             =>
                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                         })
                                                                                                                                 } else {
                                                                                                                                     panic!("{}",
                                                                                                                                            string("Match failure: PureScript_Data_List_Types.Data_List_Types_List"),)
                                                                                                                                 }
                                                                                                                             }
                                                                                                                         }
                                                                                                                     }
                                                                                                             })),
                                                                                            v_1.clone())
                                                                               });
                                                                let go =
                                                                    go_1.Value;
                                                                Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                 &&&0_i32)
                                                            }))
    }
    pub fn Data_List_findLastIndex() -> &dyn Any {
        static Data_List_findLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_findLastIndex.get_or_init(||
                                                &Func1::new(move |r#fn|
                                                                &Func1::new({
                                                                                let r#fn
                                                                                    =
                                                                                    r#fn.clone();
                                                                                move
                                                                                    |xs|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                          let xs
                                                                                                                                                                              =
                                                                                                                                                                              xs.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_length(),
                                                                                                                                                                                                                                                                                                                                                           &&&xs)),
                                                                                                                                                                                                                                                                                     &&&1_i32)),
                                                                                                                                                                                                               v)
                                                                                                                                                                      })),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_findIndex(),
                                                                                                                                                                                           &&&r#fn),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                                                           xs)))
                                                                            })))
    }
    pub fn Data_List_filterM_0040208() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Data_List::Data_List_filterM_tco(dictMonad))
    }
    pub fn Data_List_filterM_0040208_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_filterM_0040208_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_filterM_0040208_002d1.get_or_init(||
                                                        Lazy(Data_List_filterM_0040208.clone()))
    }
    pub fn Data_List_filterM_tco(dictMonad: &dyn Any) -> &dyn Any {
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Data_List_filterM_0040208_002d1 =
                            Data_List_filterM_0040208_002d1.clone();
                        let dictMonad = dictMonad.clone();
                        move |v|
                            &Func1::new({
                                            let Data_List_filterM_0040208_002d1
                                                =
                                                Data_List_filterM_0040208_002d1.clone();
                                            let v = v.clone();
                                            move |v1|
                                                {
                                                    let matchValue =
                                                        Sharpurs_Prelude::unbox(&&v);
                                                    let matchValue_1:
                                                            LrcPtr<Data_List_Types_List> =
                                                        Sharpurs_Prelude::unbox(v1);
                                                    match matchValue_1.as_ref()
                                                        {
                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                           matchValue_1_1_1)
                                                        => {
                                                            let x =
                                                                matchValue_1_1_0.clone();
                                                            let p =
                                                                matchValue;
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                   &&&Bind1),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                   &&&x)),
                                                                                             &&&Func1::new({
                                                                                                               let Data_List_filterM_0040208_002d1
                                                                                                                   =
                                                                                                                   Data_List_filterM_0040208_002d1.clone();
                                                                                                               let matchValue_1
                                                                                                                   =
                                                                                                                   matchValue_1.clone();
                                                                                                               let p
                                                                                                                   =
                                                                                                                   p.clone();
                                                                                                               let x
                                                                                                                   =
                                                                                                                   x.clone();
                                                                                                               move
                                                                                                                   |b|
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                          &&&Bind1),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_filterM_0040208_002d1.Value,
                                                                                                                                                                                                                                                                                                &&&dictMonad),
                                                                                                                                                                                                                                                             &&&p),
                                                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                 Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                                                             })),
                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                      let b
                                                                                                                                                                          =
                                                                                                                                                                          b.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |xs_prime|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                              &&&Applicative0),
                                                                                                                                                                                                           &&{
                                                                                                                                                                                                                 let matchValue_3 =
                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&b);
                                                                                                                                                                                                                 match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                  &matchValue_3)
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                     0_i32
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&x,
                                                                                                                                                                                                                                                                                     xs_prime.clone())),
                                                                                                                                                                                                                     _
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     xs_prime.clone(),
                                                                                                                                                                                                                 }
                                                                                                                                                                                                             })
                                                                                                                                                                  }))
                                                                                                           }))
                                                        }
                                                        _ =>
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                            &&&Applicative0),
                                                                                         &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)),
                                                    }
                                                }
                                        })
                    })
    }
    pub fn Data_List_filterM() -> &dyn Any {
        static Data_List_filterM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_filterM.get_or_init(||
                                          Data_List_filterM_0040208_002d1.Value)
    }
    pub fn Data_List_filter() -> &dyn Any {
        static Data_List_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_filter.get_or_init(||
                                         &Func1::new(move |p|
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
                                                                                                        |v|
                                                                                                        Func1::new({
                                                                                                                       let go_tco
                                                                                                                           =
                                                                                                                           go_tco.clone();
                                                                                                                       let v
                                                                                                                           =
                                                                                                                           v.clone();
                                                                                                                       move
                                                                                                                           |v1|
                                                                                                                           go_tco(v)(v1.clone())
                                                                                                                   })
                                                                                                })
                                                                            });
                                                             let go_1 =
                                                                 Lazy(go_2);
                                                             let go_tco =
                                                                 Func1::new({
                                                                                let p
                                                                                    =
                                                                                    p.clone();
                                                                                move
                                                                                    |v_1|
                                                                                    fix1(&(move
                                                                                               |go_tco,
                                                                                                v_1|
                                                                                               Func1::new({
                                                                                                              let go_tco
                                                                                                                  =
                                                                                                                  go_tco.clone();
                                                                                                              let v_1
                                                                                                                  =
                                                                                                                  v_1.clone();
                                                                                                              move
                                                                                                                  |v1_1|
                                                                                                                  {
                                                                                                                      let matchValue =
                                                                                                                          Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                      let matchValue_1:
                                                                                                                              LrcPtr<Data_List_Types_List> =
                                                                                                                          Sharpurs_Prelude::unbox(v1_1);
                                                                                                                      if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                matchValue_1_1_1)
                                                                                                                             =
                                                                                                                             matchValue_1.as_ref()
                                                                                                                         {
                                                                                                                          if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                       &&&match matchValue_1.as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                              Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                              =>
                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                              _
                                                                                                                                                                                              =>
                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                          }))
                                                                                                                             {
                                                                                                                              go_tco(&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                          _
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          unreachable!(),
                                                                                                                                                                                                      },
                                                                                                                                                                                                     &matchValue)))(&match matchValue_1.as_ref()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                         Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                         _
                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                     })
                                                                                                                          } else {
                                                                                                                              if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                 {
                                                                                                                                  go_tco(&matchValue)(&match matchValue_1.as_ref()
                                                                                                                                                           {
                                                                                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                              x)
                                                                                                                                                           =>
                                                                                                                                                           x.clone(),
                                                                                                                                                           _
                                                                                                                                                           =>
                                                                                                                                                           unreachable!(),
                                                                                                                                                       })
                                                                                                                              } else {
                                                                                                                                  panic!("{}",
                                                                                                                                         string("Match failure: PureScript_Data_List_Types.Data_List_Types_List"),)
                                                                                                                              }
                                                                                                                          }
                                                                                                                      } else {
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_reverse(),
                                                                                                                                                           &&&matchValue)
                                                                                                                      }
                                                                                                                  }
                                                                                                          })),
                                                                                         v_1.clone())
                                                                            });
                                                             let go =
                                                                 go_1.Value;
                                                             Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                              &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor))
                                                         }))
    }
    pub fn Data_List_intersectBy() -> &dyn Any {
        static Data_List_intersectBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_intersectBy.get_or_init(||
                                              &Func1::new(move |v|
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
                                                                                                          let matchValue_1:
                                                                                                                  LrcPtr<Data_List_Types_List> =
                                                                                                              Sharpurs_Prelude::unbox(&&v1);
                                                                                                          let matchValue_2:
                                                                                                                  LrcPtr<Data_List_Types_List> =
                                                                                                              Sharpurs_Prelude::unbox(v2);
                                                                                                          if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                 =
                                                                                                                 matchValue_1.as_ref()
                                                                                                             {
                                                                                                              &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)
                                                                                                          } else {
                                                                                                              if let Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                     =
                                                                                                                     matchValue_2.as_ref()
                                                                                                                 {
                                                                                                                  &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)
                                                                                                              } else {
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_filter(),
                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                        let matchValue_2
                                                                                                                                                                                                            =
                                                                                                                                                                                                            matchValue_2.clone();
                                                                                                                                                                                                        move
                                                                                                                                                                                                            |x|
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_any(),
                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                   x)),
                                                                                                                                                                                                                                             &&&matchValue_2)
                                                                                                                                                                                                    })),
                                                                                                                                                   &&&matchValue_1)
                                                                                                              }
                                                                                                          }
                                                                                                      }
                                                                                              })
                                                                          })))
    }
    pub fn Data_List_intersect() -> &dyn Any {
        static Data_List_intersect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_intersect.get_or_init(||
                                            &Func1::new(move |dictEq|
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_intersectBy(),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                dictEq))))
    }
    pub fn Data_List_nubByEq_0040224() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Data_List::Data_List_nubByEq_tco(&v,
                                                                                       v1)
                                   }))
    }
    pub fn Data_List_nubByEq_0040224_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_nubByEq_0040224_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_nubByEq_0040224_002d1.get_or_init(||
                                                        Lazy(Data_List_nubByEq_0040224.clone()))
    }
    pub fn Data_List_nubByEq_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v1);
        match matchValue_1.as_ref() {
            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                               matchValue_1_1_1)
            => {
                let x = matchValue_1_1_0.clone();
                let eq_prime = matchValue;
                &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&x,
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_nubByEq_0040224_002d1.Value,
                                                                                                                                                    &&&eq_prime),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_filter(),
                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                         let eq_prime
                                                                                                                                                                                                             =
                                                                                                                                                                                                             eq_prime.clone();
                                                                                                                                                                                                         let x
                                                                                                                                                                                                             =
                                                                                                                                                                                                             x.clone();
                                                                                                                                                                                                         move
                                                                                                                                                                                                             |y|
                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&eq_prime,
                                                                                                                                                                                                                                                                                                                    &&&x),
                                                                                                                                                                                                                                                                                 y))
                                                                                                                                                                                                     })),
                                                                                                                                                    &&matchValue_1_1_1))))
            }
            _ =>
            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
        }
    }
    pub fn Data_List_nubByEq() -> &dyn Any {
        static Data_List_nubByEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_nubByEq.get_or_init(||
                                          Data_List_nubByEq_0040224_002d1.Value)
    }
    pub fn Data_List_nubEq() -> &dyn Any {
        static Data_List_nubEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_nubEq.get_or_init(||
                                        &Func1::new(move |dictEq|
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_nubByEq(),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                            dictEq))))
    }
    pub fn Data_List_eqPattern() -> &dyn Any {
        static Data_List_eqPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_eqPattern.get_or_init(||
                                            &Func1::new(move |dictEq|
                                                            {
                                                                let eqList =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_eqList(),
                                                                                                     dictEq);
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                 &&&add(string("eq"),
                                                                                                        &&Func1::new({
                                                                                                                         let eqList
                                                                                                                             =
                                                                                                                             eqList.clone();
                                                                                                                         move
                                                                                                                             |x|
                                                                                                                             &Func1::new({
                                                                                                                                             let x
                                                                                                                                                 =
                                                                                                                                                 x.clone();
                                                                                                                                             move
                                                                                                                                                 |y|
                                                                                                                                                 {
                                                                                                                                                     let matchValue =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                     let matchValue_1 =
                                                                                                                                                         Sharpurs_Prelude::unbox(y);
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                            &&&eqList),
                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                      &&&matchValue_1)
                                                                                                                                                 }
                                                                                                                                         })
                                                                                                                     }),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))
                                                            }))
    }
    pub fn Data_List_ordPattern() -> &dyn Any {
        static Data_List_ordPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_ordPattern.get_or_init(||
                                             &Func1::new(move |dictOrd|
                                                             {
                                                                 let ordList =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_ordList(),
                                                                                                      dictOrd);
                                                                 let eqPattern1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_eqPattern(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                  &&&add(string("compare"),
                                                                                                         &&Func1::new({
                                                                                                                          let ordList
                                                                                                                              =
                                                                                                                              ordList.clone();
                                                                                                                          move
                                                                                                                              |x|
                                                                                                                              &Func1::new({
                                                                                                                                              let x
                                                                                                                                                  =
                                                                                                                                                  x.clone();
                                                                                                                                              move
                                                                                                                                                  |y|
                                                                                                                                                  {
                                                                                                                                                      let matchValue =
                                                                                                                                                          Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                      let matchValue_1 =
                                                                                                                                                          Sharpurs_Prelude::unbox(y);
                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                             &&&ordList),
                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                       &&&matchValue_1)
                                                                                                                                                  }
                                                                                                                                          })
                                                                                                                      }),
                                                                                                         add(string("Eq0"),
                                                                                                             &&Func1::new({
                                                                                                                              let eqPattern1
                                                                                                                                  =
                                                                                                                                  eqPattern1.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &eqPattern1
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Data_List_elemLastIndex() -> &dyn Any {
        static Data_List_elemLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_elemLastIndex.get_or_init(||
                                                &Func1::new(move |dictEq|
                                                                &Func1::new({
                                                                                let dictEq
                                                                                    =
                                                                                    dictEq.clone();
                                                                                move
                                                                                    |x|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_findLastIndex(),
                                                                                                                     &&&Func1::new({
                                                                                                                                       let x
                                                                                                                                           =
                                                                                                                                           x.clone();
                                                                                                                                       move
                                                                                                                                           |v|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                  &&&dictEq),
                                                                                                                                                                                                               v),
                                                                                                                                                                            &&&x)
                                                                                                                                   }))
                                                                            })))
    }
    pub fn Data_List_elemIndex() -> &dyn Any {
        static Data_List_elemIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_elemIndex.get_or_init(||
                                            &Func1::new(move |dictEq|
                                                            &Func1::new({
                                                                            let dictEq
                                                                                =
                                                                                dictEq.clone();
                                                                            move
                                                                                |x|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_findIndex(),
                                                                                                                 &&&Func1::new({
                                                                                                                                   let x
                                                                                                                                       =
                                                                                                                                       x.clone();
                                                                                                                                   move
                                                                                                                                       |v|
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                              &&&dictEq),
                                                                                                                                                                                                           v),
                                                                                                                                                                        &&&x)
                                                                                                                               }))
                                                                        })))
    }
    pub fn Data_List_dropWhile() -> &dyn Any {
        static Data_List_dropWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_dropWhile.get_or_init(||
                                            &Func1::new(move |p|
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
                                                                                                           |v|
                                                                                                           go_tco(v.clone())
                                                                                                   })
                                                                               });
                                                                let go_1 =
                                                                    Lazy(go_2);
                                                                let go_tco =
                                                                    Func1::new({
                                                                                   let p
                                                                                       =
                                                                                       p.clone();
                                                                                   move
                                                                                       |v_1|
                                                                                       {
                                                                                           let v_1 =
                                                                                               v_1.clone();
                                                                                           '_go_tco:
                                                                                               loop 
                                                                                                    {
                                                                                                   break
                                                                                                       '_go_tco
                                                                                                       ({
                                                                                                            let matchValue:
                                                                                                                    LrcPtr<Data_List_Types_List> =
                                                                                                                Sharpurs_Prelude::unbox(&&v_1);
                                                                                                            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                      matchValue_1_1)
                                                                                                                   =
                                                                                                                   matchValue.as_ref()
                                                                                                               {
                                                                                                                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                    Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                       _)
                                                                                                                                                                                    =>
                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                    _
                                                                                                                                                                                    =>
                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                }))
                                                                                                                   {
                                                                                                                    let v_1_temp =
                                                                                                                        &match matchValue.as_ref()
                                                                                                                             {
                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                x)
                                                                                                                             =>
                                                                                                                             x.clone(),
                                                                                                                             _
                                                                                                                             =>
                                                                                                                             unreachable!(),
                                                                                                                         };
                                                                                                                    v_1.set(v_1_temp);
                                                                                                                    continue
                                                                                                                        '_go_tco

                                                                                                                } else {
                                                                                                                    &matchValue
                                                                                                                }
                                                                                                            } else {
                                                                                                                &matchValue
                                                                                                            }
                                                                                                        })
                                                                                                       ;
                                                                                               }
                                                                                       }
                                                                               });
                                                                let go =
                                                                    go_1.Value;
                                                                &go
                                                            }))
    }
    pub fn Data_List_dropEnd() -> &dyn Any {
        static Data_List_dropEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_dropEnd.get_or_init(||
                                          &Func1::new(move |n|
                                                          &Func1::new({
                                                                          let n
                                                                              =
                                                                              n.clone();
                                                                          move
                                                                              |xs|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_take(),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_length(),
                                                                                                                                                                                                                                                           xs)),
                                                                                                                                                                                     &&&n)),
                                                                                                               xs)
                                                                      })))
    }
    pub fn Data_List_drop_0040248() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Data_List::Data_List_drop_tco(&v,
                                                                                    v1)
                                   }))
    }
    pub fn Data_List_drop_0040248_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_drop_0040248_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_drop_0040248_002d1.get_or_init(||
                                                     Lazy(Data_List_drop_0040248.clone()))
    }
    pub fn Data_List_drop_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v1);
        if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                           &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                        &&&matchValue),
                                                                     &&&1_i32))
           {
            &matchValue_1
        } else {
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                      matchValue_1_1_1)
                   = matchValue_1.as_ref() {
                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_drop_0040248_002d1.Value,
                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                             &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                          &&&matchValue),
                                                                                                                       &&&1_i32)),
                                                 &&&match matchValue_1.as_ref()
                                                        {
                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                           x)
                                                        => x.clone(),
                                                        _ => unreachable!(),
                                                    })
            } else {
                &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)
            }
        }
    }
    pub fn Data_List_drop() -> &dyn Any {
        static Data_List_drop: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_List_drop.get_or_init(|| Data_List_drop_0040248_002d1.Value)
    }
    pub fn Data_List_slice() -> &dyn Any {
        static Data_List_slice: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_slice.get_or_init(||
                                        &Func1::new(move |start|
                                                        &Func1::new({
                                                                        let start
                                                                            =
                                                                            start.clone();
                                                                        move
                                                                            |end_var|
                                                                            &Func1::new({
                                                                                            let end_var
                                                                                                =
                                                                                                end_var.clone();
                                                                                            move
                                                                                                |xs|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_take(),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                          &&&end_var),
                                                                                                                                                                                                       &&&start)),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_drop(),
                                                                                                                                                                                                       &&&start),
                                                                                                                                                                    xs))
                                                                                        })
                                                                    })))
    }
    pub fn Data_List_takeEnd() -> &dyn Any {
        static Data_List_takeEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_takeEnd.get_or_init(||
                                          &Func1::new(move |n|
                                                          &Func1::new({
                                                                          let n
                                                                              =
                                                                              n.clone();
                                                                          move
                                                                              |xs|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_drop(),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_length(),
                                                                                                                                                                                                                                                           xs)),
                                                                                                                                                                                     &&&n)),
                                                                                                               xs)
                                                                      })))
    }
    pub fn Data_List_deleteBy_0040256() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           Func1::new({
                                                          let v1 = v1.clone();
                                                          move |v2|
                                                              PureScript_Data_List::Data_List_deleteBy_tco(&v,
                                                                                                           &v1,
                                                                                                           v2)
                                                      })
                                   }))
    }
    pub fn Data_List_deleteBy_0040256_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_deleteBy_0040256_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_deleteBy_0040256_002d1.get_or_init(||
                                                         Lazy(Data_List_deleteBy_0040256.clone()))
    }
    pub fn Data_List_deleteBy_tco(v: &dyn Any, v1: &dyn Any, v2: &dyn Any)
     -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        let matchValue_2: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v2);
        if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                  matchValue_2_1_1)
               = matchValue_2.as_ref() {
            if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                            &&&matchValue_1),
                                                                         &&&match matchValue_2.as_ref()
                                                                                {
                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                   _)
                                                                                =>
                                                                                x.clone(),
                                                                                _
                                                                                =>
                                                                                unreachable!(),
                                                                            }))
               {
                &match matchValue_2.as_ref() {
                     Data_List_Types_List::Data_List_Types_Consusd_Ctor(_, x)
                     => x.clone(),
                     _ => unreachable!(),
                 }
            } else {
                &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                     {
                                                                                     Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                        _)
                                                                                     =>
                                                                                     x.clone(),
                                                                                     _
                                                                                     =>
                                                                                     unreachable!(),
                                                                                 },
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_deleteBy_0040256_002d1.Value,
                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                    &&&matchValue_1),
                                                                                                                 &&&match matchValue_2.as_ref()
                                                                                                                        {
                                                                                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                           x)
                                                                                                                        =>
                                                                                                                        x.clone(),
                                                                                                                        _
                                                                                                                        =>
                                                                                                                        unreachable!(),
                                                                                                                    })))
            }
        } else {
            &LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)
        }
    }
    pub fn Data_List_deleteBy() -> &dyn Any {
        static Data_List_deleteBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_deleteBy.get_or_init(||
                                           Data_List_deleteBy_0040256_002d1.Value)
    }
    pub fn Data_List_unionBy() -> &dyn Any {
        static Data_List_unionBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_unionBy.get_or_init(||
                                          &Func1::new(move |eq|
                                                          &Func1::new({
                                                                          let eq
                                                                              =
                                                                              eq.clone();
                                                                          move
                                                                              |xs|
                                                                              &Func1::new({
                                                                                              let xs
                                                                                                  =
                                                                                                  xs.clone();
                                                                                              move
                                                                                                  |ys|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                         &&&PureScript_Data_List_Types::Data_List_Types_semigroupList()),
                                                                                                                                                                      &&&xs),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                               &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_deleteBy(),
                                                                                                                                                                                                                                                                                                                  &&&eq))),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_nubByEq(),
                                                                                                                                                                                                                                                                               &&&eq),
                                                                                                                                                                                                                                            ys)),
                                                                                                                                                                      &&&xs))
                                                                                          })
                                                                      })))
    }
    pub fn Data_List_union() -> &dyn Any {
        static Data_List_union: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_union.get_or_init(||
                                        &Func1::new(move |dictEq|
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_unionBy(),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                            dictEq))))
    }
    pub fn Data_List_deleteAt_0040264() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Data_List::Data_List_deleteAt_tco(&v,
                                                                                        v1)
                                   }))
    }
    pub fn Data_List_deleteAt_0040264_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_deleteAt_0040264_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_deleteAt_0040264_002d1.get_or_init(||
                                                         Lazy(Data_List_deleteAt_0040264.clone()))
    }
    pub fn Data_List_deleteAt_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v1);
        if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                    &matchValue).is_some() {
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                      matchValue_1_1_1)
                   = matchValue_1.as_ref() {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue_1.as_ref()
                                                                            {
                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                               x)
                                                                            =>
                                                                            x.clone(),
                                                                            _
                                                                            =>
                                                                            unreachable!(),
                                                                        }))
            } else {
                if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                          matchValue_1_1_1)
                       = matchValue_1.as_ref() {
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                        &&&Func1::new({
                                                                                                          let matchValue_1
                                                                                                              =
                                                                                                              matchValue_1.clone();
                                                                                                          move
                                                                                                              |v2|
                                                                                                              &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                      _)
                                                                                                                                                                                   =>
                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                   _
                                                                                                                                                                                   =>
                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                               },
                                                                                                                                                                              v2.clone()))
                                                                                                      })),
                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_deleteAt_0040264_002d1.Value,
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                    &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                 &&&matchValue),
                                                                                                                                                              &&&1_i32)),
                                                                                        &&&match matchValue_1.as_ref()
                                                                                               {
                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                  x)
                                                                                               =>
                                                                                               x.clone(),
                                                                                               _
                                                                                               =>
                                                                                               unreachable!(),
                                                                                           }))
                } else {
                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                }
            }
        } else {
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                      matchValue_1_1_1)
                   = matchValue_1.as_ref() {
                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                    &&&Func1::new({
                                                                                                      let matchValue_1
                                                                                                          =
                                                                                                          matchValue_1.clone();
                                                                                                      move
                                                                                                          |v2|
                                                                                                          &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                  _)
                                                                                                                                                                               =>
                                                                                                                                                                               x.clone(),
                                                                                                                                                                               _
                                                                                                                                                                               =>
                                                                                                                                                                               unreachable!(),
                                                                                                                                                                           },
                                                                                                                                                                          v2.clone()))
                                                                                                  })),
                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_deleteAt_0040264_002d1.Value,
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                          &&&1_i32)),
                                                                                    &&&match matchValue_1.as_ref()
                                                                                           {
                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                              x)
                                                                                           =>
                                                                                           x.clone(),
                                                                                           _
                                                                                           =>
                                                                                           unreachable!(),
                                                                                       }))
            } else {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
            }
        }
    }
    pub fn Data_List_deleteAt() -> &dyn Any {
        static Data_List_deleteAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_deleteAt.get_or_init(||
                                           Data_List_deleteAt_0040264_002d1.Value)
    }
    pub fn Data_List_delete() -> &dyn Any {
        static Data_List_delete: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_delete.get_or_init(||
                                         &Func1::new(move |dictEq|
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_deleteBy(),
                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                             dictEq))))
    }
    pub fn Data_List_difference() -> &dyn Any {
        static Data_List_difference: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_difference.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                 &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_delete(),
                                                                                                                                                                    dictEq)))))
    }
    pub fn Data_List_concatMap() -> &dyn Any {
        static Data_List_concatMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_concatMap.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                &&&PureScript_Data_List_Types::Data_List_Types_bindList())))
    }
    pub fn Data_List_concat() -> &dyn Any {
        static Data_List_concat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_concat.get_or_init(||
                                         &Func1::new(move |v|
                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                &&&PureScript_Data_List_Types::Data_List_Types_bindList()),
                                                                                                                             v),
                                                                                          &&&PureScript_Data_List::Data_List_identity())))
    }
    pub fn Data_List_catMaybes() -> &dyn Any {
        static Data_List_catMaybes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_catMaybes.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_mapMaybe(),
                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                &&&PureScript_Control_Category::Control_Category_categoryFn())))
    }
    pub fn Data_List_alterAt_0040278() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           Func1::new({
                                                          let v1 = v1.clone();
                                                          move |v2|
                                                              PureScript_Data_List::Data_List_alterAt_tco(&v,
                                                                                                          &v1,
                                                                                                          v2)
                                                      })
                                   }))
    }
    pub fn Data_List_alterAt_0040278_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_alterAt_0040278_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_alterAt_0040278_002d1.get_or_init(||
                                                        Lazy(Data_List_alterAt_0040278.clone()))
    }
    pub fn Data_List_alterAt_tco(v: &dyn Any, v1: &dyn Any, v2: &dyn Any)
     -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        let matchValue_2: LrcPtr<Data_List_Types_List> =
            Sharpurs_Prelude::unbox(v2);
        if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                    &matchValue).is_some() {
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                      matchValue_2_1_1)
                   = matchValue_2.as_ref() {
                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                    &&&Func1::new(move
                                                                                                      |usd__arg1|
                                                                                                      &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                 &&{
                                                       let matchValue_4:
                                                               LrcPtr<Data_Maybe_Maybe> =
                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                      &&&match matchValue_2.as_ref()
                                                                                                                             {
                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                _)
                                                                                                                             =>
                                                                                                                             x.clone(),
                                                                                                                             _
                                                                                                                             =>
                                                                                                                             unreachable!(),
                                                                                                                         }));
                                                       match matchValue_4.as_ref()
                                                           {
                                                           Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_4_1_0)
                                                           =>
                                                           &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_4_1_0,
                                                                                                                           &match matchValue_2.as_ref()
                                                                                                                                {
                                                                                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                   x)
                                                                                                                                =>
                                                                                                                                x.clone(),
                                                                                                                                _
                                                                                                                                =>
                                                                                                                                unreachable!(),
                                                                                                                            })),
                                                           _ =>
                                                           &match matchValue_2.as_ref()
                                                                {
                                                                Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                   x)
                                                                => x.clone(),
                                                                _ =>
                                                                unreachable!(),
                                                            },
                                                       }
                                                   })
            } else {
                if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                          matchValue_2_1_1)
                       = matchValue_2.as_ref() {
                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                        &&&Func1::new({
                                                                                                          let matchValue_2
                                                                                                              =
                                                                                                              matchValue_2.clone();
                                                                                                          move
                                                                                                              |v3_1|
                                                                                                              &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                   Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                      _)
                                                                                                                                                                                   =>
                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                   _
                                                                                                                                                                                   =>
                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                               },
                                                                                                                                                                              v3_1.clone()))
                                                                                                      })),
                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_alterAt_0040278_002d1.Value,
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                       &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                 &&&1_i32)),
                                                                                                                           &&&matchValue_1),
                                                                                        &&&match matchValue_2.as_ref()
                                                                                               {
                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                  x)
                                                                                               =>
                                                                                               x.clone(),
                                                                                               _
                                                                                               =>
                                                                                               unreachable!(),
                                                                                           }))
                } else {
                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                }
            }
        } else {
            if let Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_2_1_0,
                                                                      matchValue_2_1_1)
                   = matchValue_2.as_ref() {
                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                    &&&Func1::new({
                                                                                                      let matchValue_2
                                                                                                          =
                                                                                                          matchValue_2.clone();
                                                                                                      move
                                                                                                          |v3_1|
                                                                                                          &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                               Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                  _)
                                                                                                                                                                               =>
                                                                                                                                                                               x.clone(),
                                                                                                                                                                               _
                                                                                                                                                                               =>
                                                                                                                                                                               unreachable!(),
                                                                                                                                                                           },
                                                                                                                                                                          v3_1.clone()))
                                                                                                  })),
                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_alterAt_0040278_002d1.Value,
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                             &&&1_i32)),
                                                                                                                       &&&matchValue_1),
                                                                                    &&&match matchValue_2.as_ref()
                                                                                           {
                                                                                           Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                              x)
                                                                                           =>
                                                                                           x.clone(),
                                                                                           _
                                                                                           =>
                                                                                           unreachable!(),
                                                                                       }))
            } else {
                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
            }
        }
    }
    pub fn Data_List_alterAt() -> &dyn Any {
        static Data_List_alterAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_alterAt.get_or_init(||
                                          Data_List_alterAt_0040278_002d1.Value)
    }
    pub fn Data_List_modifyAt() -> &dyn Any {
        static Data_List_modifyAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_modifyAt.get_or_init(||
                                           &Func1::new(move |n|
                                                           &Func1::new({
                                                                           let n
                                                                               =
                                                                               n.clone();
                                                                           move
                                                                               |f|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_alterAt(),
                                                                                                                                                   &&&n),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                   f))
                                                                       })))
    }
}
