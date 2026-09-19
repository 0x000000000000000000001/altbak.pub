pub mod PureScript_Data_Set_NonEmpty {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_ba715d91::PureScript_Data_Array_NonEmpty_Internal;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_d662adf2::PureScript_Data_List_Types;
    use crate::module_ed2bf3e0::PureScript_Data_Map_Internal;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_2563115d::PureScript_Data_Semigroup_Foldable;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e534597::PureScript_Data_Set;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_285149e9::PureScript_Safe_Coerce;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Set_NonEmpty_coerce() -> &dyn Any {
        static Data_Set_NonEmpty_coerce: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_coerce.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Set_NonEmpty_NonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_NonEmptySet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_NonEmptySet.get_or_init(||
                                                      &Func1::new(move |x|
                                                                      x.clone()))
    }
    pub fn Data_Set_NonEmpty_unionSet() -> &dyn Any {
        static Data_Set_NonEmpty_unionSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_unionSet.get_or_init(||
                                                   &Func1::new(move |dictOrd|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_semigroupSet(),
                                                                                                                                                                          dictOrd)))))
    }
    pub fn Data_Set_NonEmpty_toUnfoldable1() -> &dyn Any {
        static Data_Set_NonEmpty_toUnfoldable1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_toUnfoldable1.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictUnfoldable1|
                                                                        {
                                                                            let stepNext =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_stepAscCps(),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_mkFn3(),
                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                         |k|
                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                         let k
                                                                                                                                                                                                                             =
                                                                                                                                                                                                                             k.clone();
                                                                                                                                                                                                                         move
                                                                                                                                                                                                                             |v|
                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                             |next|
                                                                                                                                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&k,
                                                                                                                                                                                                                                                                                                                                                            next.clone())))))
                                                                                                                                                                                                                     })))),
                                                                                                                 &&&Func1::new(move
                                                                                                                                   |v_1|
                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)));
                                                                            let stepHead =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_stepAscCps(),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_mkFn3(),
                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                         |k_1|
                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                         let k_1
                                                                                                                                                                                                                             =
                                                                                                                                                                                                                             k_1.clone();
                                                                                                                                                                                                                         move
                                                                                                                                                                                                                             |v_2|
                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                             |next_1|
                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&k_1,
                                                                                                                                                                                                                                                                                                     next_1.clone())))
                                                                                                                                                                                                                     })))),
                                                                                                                 &&&Func1::new(move
                                                                                                                                   |v_3|
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafeCrashWith(),
                                                                                                                                                                    &&&string("toUnfoldable1: impossible"))));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                                      dictUnfoldable1),
                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                     let stepNext
                                                                                                                                                                                                         =
                                                                                                                                                                                                         stepNext.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |v_4|
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                                                                                                                                                                                                             &&&stepNext),
                                                                                                                                                                                                                                          v_4)
                                                                                                                                                                                                 }))),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                   &&&stepHead),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                      &&&PureScript_Data_Map_Internal::Data_Map_Internal_toMapIter()),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Set::Data_Set_toMap()),
                                                                                                                                                                                                                      &&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_coerce()))))
                                                                        }))
    }
    pub fn Data_Set_NonEmpty_toUnfoldable11() -> &dyn Any {
        static Data_Set_NonEmpty_toUnfoldable11: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_toUnfoldable11.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_toUnfoldable1(),
                                                                                          &&&PureScript_Data_List_Types::Data_List_Types_unfoldable1NonEmptyList()))
    }
    pub fn Data_Set_NonEmpty_toUnfoldable() -> &dyn Any {
        static Data_Set_NonEmpty_toUnfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_toUnfoldable.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictUnfoldable|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_toUnfoldable(),
                                                                                                                                           dictUnfoldable))))
    }
    pub fn Data_Set_NonEmpty_toSet() -> &dyn Any {
        static Data_Set_NonEmpty_toSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_toSet.get_or_init(||
                                                &Func1::new(move |v|
                                                                &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Set_NonEmpty_subset() -> &dyn Any {
        static Data_Set_NonEmpty_subset: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_subset.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_subset(),
                                                                                                                                     dictOrd))))
    }
    pub fn Data_Set_NonEmpty_size() -> &dyn Any {
        static Data_Set_NonEmpty_size: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_size.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                &&&PureScript_Data_Set::Data_Set_size()))
    }
    pub fn Data_Set_NonEmpty_singleton() -> &dyn Any {
        static Data_Set_NonEmpty_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_singleton.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                     &&&PureScript_Data_Set::Data_Set_singleton()))
    }
    pub fn Data_Set_NonEmpty_showNonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_showNonEmptySet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_showNonEmptySet.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictShow|
                                                                          {
                                                                              let showNonEmptyArray =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_showNonEmptyArray(),
                                                                                                                   dictShow);
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                               &&&add(string("show"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let showNonEmptyArray
                                                                                                                                           =
                                                                                                                                           showNonEmptyArray.clone();
                                                                                                                                       move
                                                                                                                                           |s|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                               &&&string("(fromFoldable1 ")),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                        &&&showNonEmptyArray),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_toUnfoldable1(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray()),
                                                                                                                                                                                                                                                                                                                        s))),
                                                                                                                                                                                                               &&&string(")")))
                                                                                                                                   }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))
                                                                          }))
    }
    pub fn Data_Set_NonEmpty_semigroupNonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_semigroupNonEmptySet:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_semigroupNonEmptySet.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictOrd|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_semigroupSet(),
                                                                                                                dictOrd)))
    }
    pub fn Data_Set_NonEmpty_properSubset() -> &dyn Any {
        static Data_Set_NonEmpty_properSubset: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_properSubset.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictOrd|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_properSubset(),
                                                                                                                                           dictOrd))))
    }
    pub fn Data_Set_NonEmpty_ordNonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_ordNonEmptySet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_ordNonEmptySet.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictOrd|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_ordSet(),
                                                                                                          dictOrd)))
    }
    pub fn Data_Set_NonEmpty_ord1NonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_ord1NonEmptySet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_ord1NonEmptySet.get_or_init(||
                                                          &PureScript_Data_Set::Data_Set_ord1Set())
    }
    pub fn Data_Set_NonEmpty_min() -> &dyn Any {
        static Data_Set_NonEmpty_min: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_min.get_or_init(||
                                              &Func1::new(move |v|
                                                              {
                                                                  let s =
                                                                      Sharpurs_Prelude::unbox(v);
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                   &&&Func1::new({
                                                                                                                     let s
                                                                                                                         =
                                                                                                                         s.clone();
                                                                                                                     move
                                                                                                                         |usd__unused|
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_findMin(),
                                                                                                                                                                                             &&&s))
                                                                                                                 }))
                                                              }))
    }
    pub fn Data_Set_NonEmpty_member() -> &dyn Any {
        static Data_Set_NonEmpty_member: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_member.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_member(),
                                                                                                                                     dictOrd))))
    }
    pub fn Data_Set_NonEmpty_max() -> &dyn Any {
        static Data_Set_NonEmpty_max: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_max.get_or_init(||
                                              &Func1::new(move |v|
                                                              {
                                                                  let s =
                                                                      Sharpurs_Prelude::unbox(v);
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                   &&&Func1::new({
                                                                                                                     let s
                                                                                                                         =
                                                                                                                         s.clone();
                                                                                                                     move
                                                                                                                         |usd__unused|
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_findMax(),
                                                                                                                                                                                             &&&s))
                                                                                                                 }))
                                                              }))
    }
    pub fn Data_Set_NonEmpty_mapMaybe() -> &dyn Any {
        static Data_Set_NonEmpty_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_mapMaybe.get_or_init(||
                                                   &Func1::new(move |dictOrd|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_mapMaybe(),
                                                                                                                                       dictOrd))))
    }
    pub fn Data_Set_NonEmpty_map() -> &dyn Any {
        static Data_Set_NonEmpty_map: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_map.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_map(),
                                                                                                                                  dictOrd))))
    }
    pub fn Data_Set_NonEmpty_insert() -> &dyn Any {
        static Data_Set_NonEmpty_insert: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_insert.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_insert(),
                                                                                                                                     dictOrd))))
    }
    pub fn Data_Set_NonEmpty_fromSet() -> &dyn Any {
        static Data_Set_NonEmpty_fromSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_fromSet.get_or_init(||
                                                  &Func1::new(move |s|
                                                                  {
                                                                      let matchValue =
                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_isEmpty(),
                                                                                                                                    s));
                                                                      match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                       &matchValue)
                                                                          {
                                                                          0_i32
                                                                          =>
                                                                          &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                          _ =>
                                                                          &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_NonEmptySet(),
                                                                                                                                                                  s))),
                                                                      }
                                                                  }))
    }
    pub fn Data_Set_NonEmpty_intersection() -> &dyn Any {
        static Data_Set_NonEmpty_intersection: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_intersection.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictOrd|
                                                                       &Func1::new({
                                                                                       let dictOrd
                                                                                           =
                                                                                           dictOrd.clone();
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
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_fromSet(),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_intersection(),
                                                                                                                                                                                                                                                             &&&dictOrd),
                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                       &&&matchValue_1))
                                                                                                               }
                                                                                                       })
                                                                                   })))
    }
    pub fn Data_Set_NonEmpty_fromFoldable1() -> &dyn Any {
        static Data_Set_NonEmpty_fromFoldable1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_fromFoldable1.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictFoldable1|
                                                                        &Func1::new({
                                                                                        let dictFoldable1
                                                                                            =
                                                                                            dictFoldable1.clone();
                                                                                        move
                                                                                            |dictOrd|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                   &&&dictFoldable1),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_semigroupNonEmptySet(),
                                                                                                                                                                                                   dictOrd)),
                                                                                                                             &&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_singleton())
                                                                                    })))
    }
    pub fn Data_Set_NonEmpty_fromFoldable() -> &dyn Any {
        static Data_Set_NonEmpty_fromFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_fromFoldable.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictFoldable|
                                                                       {
                                                                           let fromFoldable2 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_fromFoldable(),
                                                                                                                dictFoldable);
                                                                           &Func1::new({
                                                                                           let fromFoldable2
                                                                                               =
                                                                                               fromFoldable2.clone();
                                                                                           move
                                                                                               |dictOrd|
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                   &&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_fromSet()),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&fromFoldable2,
                                                                                                                                                                   dictOrd))
                                                                                       })
                                                                       }))
    }
    pub fn Data_Set_NonEmpty_foldableNonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_foldableNonEmptySet:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_foldableNonEmptySet.get_or_init(||
                                                              &PureScript_Data_Set::Data_Set_foldableSet())
    }
    pub fn Data_Set_NonEmpty_foldable1NonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_foldable1NonEmptySet:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_foldable1NonEmptySet.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_Foldable1usd_Dict(),
                                                                                                &&&add(string("foldMap1"),
                                                                                                       &&Func1::new(move
                                                                                                                        |dictSemigroup|
                                                                                                                        &Func1::new({
                                                                                                                                        let dictSemigroup
                                                                                                                                            =
                                                                                                                                            dictSemigroup.clone();
                                                                                                                                        move
                                                                                                                                            |f|
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_List_Types::Data_List_Types_foldable1NonEmptyList()),
                                                                                                                                                                                                                                                                                      &&&dictSemigroup),
                                                                                                                                                                                                                                                   f)),
                                                                                                                                                                             &&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_toUnfoldable11())
                                                                                                                                    })),
                                                                                                       add(string("foldr1"),
                                                                                                           &&Func1::new(move
                                                                                                                            |f_1|
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldr1(),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_List_Types::Data_List_Types_foldable1NonEmptyList()),
                                                                                                                                                                                                                                   f_1)),
                                                                                                                                                             &&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_toUnfoldable11())),
                                                                                                           add(string("foldl1"),
                                                                                                               &&Func1::new(move
                                                                                                                                |f_2|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldl1(),
                                                                                                                                                                                                                                                                          &&&PureScript_Data_List_Types::Data_List_Types_foldable1NonEmptyList()),
                                                                                                                                                                                                                                       f_2)),
                                                                                                                                                                 &&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_toUnfoldable11())),
                                                                                                               add(string("Foldable0"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused|
                                                                                                                                    &PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_foldableNonEmptySet()),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))))
    }
    pub fn Data_Set_NonEmpty_filter() -> &dyn Any {
        static Data_Set_NonEmpty_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_filter.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_filter(),
                                                                                                                                     dictOrd))))
    }
    pub fn Data_Set_NonEmpty_eqNonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_eqNonEmptySet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_eqNonEmptySet.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictEq|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_eqSet(),
                                                                                                         dictEq)))
    }
    pub fn Data_Set_NonEmpty_eq1NonEmptySet() -> &dyn Any {
        static Data_Set_NonEmpty_eq1NonEmptySet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_eq1NonEmptySet.get_or_init(||
                                                         &PureScript_Data_Set::Data_Set_eq1Set())
    }
    pub fn Data_Set_NonEmpty_difference() -> &dyn Any {
        static Data_Set_NonEmpty_difference: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_difference.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictOrd|
                                                                     &Func1::new({
                                                                                     let dictOrd
                                                                                         =
                                                                                         dictOrd.clone();
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
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_fromSet(),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_difference(),
                                                                                                                                                                                                                                                           &&&dictOrd),
                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                     &&&matchValue_1))
                                                                                                             }
                                                                                                     })
                                                                                 })))
    }
    pub fn Data_Set_NonEmpty_delete() -> &dyn Any {
        static Data_Set_NonEmpty_delete: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_delete.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 &Func1::new({
                                                                                 let dictOrd
                                                                                     =
                                                                                     dictOrd.clone();
                                                                                 move
                                                                                     |a|
                                                                                     &Func1::new({
                                                                                                     let a
                                                                                                         =
                                                                                                         a.clone();
                                                                                                     move
                                                                                                         |v|
                                                                                                         {
                                                                                                             let matchValue =
                                                                                                                 Sharpurs_Prelude::unbox(&&a);
                                                                                                             let matchValue_1 =
                                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set_NonEmpty::Data_Set_NonEmpty_fromSet(),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_delete(),
                                                                                                                                                                                                                                                       &&&dictOrd),
                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                 &&&matchValue_1))
                                                                                                         }
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Set_NonEmpty_cons() -> &dyn Any {
        static Data_Set_NonEmpty_cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_NonEmpty_cons.get_or_init(||
                                               &Func1::new(move |dictOrd|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_insert(),
                                                                                                                                   dictOrd))))
    }
}
