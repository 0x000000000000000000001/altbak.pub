pub mod PureScript_Data_Set {
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
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_d662adf2::PureScript_Data_List_Types;
    use crate::module_843b47b7::PureScript_Data_List;
    use crate::module_ed2bf3e0::PureScript_Data_Map_Internal;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_285149e9::PureScript_Safe_Coerce;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Set_identity() -> &dyn Any {
        static Data_Set_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_identity.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                           &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Set_Set() -> &dyn Any {
        static Data_Set_Set: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Set_Set.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Set_union() -> &dyn Any {
        static Data_Set_union: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Set_union.get_or_init(||
                                       &Func1::new(move |dictOrd|
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_union(),
                                                                                                                           dictOrd))))
    }
    pub fn Data_Set_toggle() -> &dyn Any {
        static Data_Set_toggle: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_toggle.get_or_init(||
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
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_Set(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_alter(),
                                                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                                    &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit()))),
                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                   |v1|
                                                                                                                                                                                                                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))),
                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                        &&&matchValue_1))
                                                                                                }
                                                                                        })
                                                                    })))
    }
    pub fn Data_Set_toMap() -> &dyn Any {
        static Data_Set_toMap: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Set_toMap.get_or_init(||
                                       &Func1::new(move |v|
                                                       &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Set_toList() -> &dyn Any {
        static Data_Set_toList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_toList.get_or_init(||
                                        &Func1::new(move |v|
                                                        {
                                                            let m =
                                                                Sharpurs_Prelude::unbox(v);
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_keys(),
                                                                                             &&&m)
                                                        }))
    }
    pub fn Data_Set_toUnfoldable() -> &dyn Any {
        static Data_Set_toUnfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_toUnfoldable.get_or_init(||
                                              &Func1::new(move
                                                              |dictUnfoldable|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_toUnfoldable(),
                                                                                                                                                                     dictUnfoldable)),
                                                                                               &&&PureScript_Data_Set::Data_Set_toList())))
    }
    pub fn Data_Set_size() -> &dyn Any {
        static Data_Set_size: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Set_size.get_or_init(||
                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                       &&&PureScript_Data_Map_Internal::Data_Map_Internal_size()))
    }
    pub fn Data_Set_singleton() -> &dyn Any {
        static Data_Set_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_singleton.get_or_init(||
                                           &Func1::new(move |a|
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_Set(),
                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_singleton(),
                                                                                                                                                                  a),
                                                                                                                               &&&PureScript_Data_Unit::Data_Unit_unit()))))
    }
    pub fn Data_Set_showSet() -> &dyn Any {
        static Data_Set_showSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_showSet.get_or_init(||
                                         &Func1::new(move |dictShow|
                                                         {
                                                             let showArray =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_showArray(),
                                                                                                  dictShow);
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                              &&&add(string("show"),
                                                                                                     &&Func1::new({
                                                                                                                      let showArray
                                                                                                                          =
                                                                                                                          showArray.clone();
                                                                                                                      move
                                                                                                                          |s|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                              &&&string("(fromFoldable ")),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                       &&&showArray),
                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_toUnfoldable(),
                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldableArray()),
                                                                                                                                                                                                                                                                                                       s))),
                                                                                                                                                                                              &&&string(")")))
                                                                                                                  }),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))
                                                         }))
    }
    pub fn Data_Set_semigroupSet() -> &dyn Any {
        static Data_Set_semigroupSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_semigroupSet.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                               &&&add(string("append"),
                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_union(),
                                                                                                                                        dictOrd),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Set_member() -> &dyn Any {
        static Data_Set_member: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_member.get_or_init(||
                                        &Func1::new(move |dictOrd|
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_member(),
                                                                                                                            dictOrd))))
    }
    pub fn Data_Set_isEmpty() -> &dyn Any {
        static Data_Set_isEmpty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_isEmpty.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()),
                                                                          &&&PureScript_Data_Map_Internal::Data_Map_Internal_isEmpty()))
    }
    pub fn Data_Set_intersection() -> &dyn Any {
        static Data_Set_intersection: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_intersection.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_intersection(),
                                                                                                                                  dictOrd))))
    }
    pub fn Data_Set_insert() -> &dyn Any {
        static Data_Set_insert: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_insert.get_or_init(||
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
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_Set(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_insert(),
                                                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                                           &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                                        &&&matchValue_1))
                                                                                                }
                                                                                        })
                                                                    })))
    }
    pub fn Data_Set_fromMap() -> &dyn Any {
        static Data_Set_fromMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_fromMap.get_or_init(|| &PureScript_Data_Set::Data_Set_Set())
    }
    pub fn Data_Set_foldableSet() -> &dyn Any {
        static Data_Set_foldableSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_foldableSet.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                              &&&add(string("foldMap"),
                                                                                     &&Func1::new(move
                                                                                                      |dictMonoid|
                                                                                                      &Func1::new({
                                                                                                                      let dictMonoid
                                                                                                                          =
                                                                                                                          dictMonoid.clone();
                                                                                                                      move
                                                                                                                          |f|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                                                                                                    &&&dictMonoid),
                                                                                                                                                                                                                                 f)),
                                                                                                                                                           &&&PureScript_Data_Set::Data_Set_toList())
                                                                                                                  })),
                                                                                     add(string("foldl"),
                                                                                         &&Func1::new(move
                                                                                                          |f_1|
                                                                                                          &Func1::new({
                                                                                                                          let f_1
                                                                                                                              =
                                                                                                                              f_1.clone();
                                                                                                                          move
                                                                                                                              |x|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                                                                                                        &&&f_1),
                                                                                                                                                                                                                                     x)),
                                                                                                                                                               &&&PureScript_Data_Set::Data_Set_toList())
                                                                                                                      })),
                                                                                         add(string("foldr"),
                                                                                             &&Func1::new(move
                                                                                                              |f_2|
                                                                                                              &Func1::new({
                                                                                                                              let f_2
                                                                                                                                  =
                                                                                                                                  f_2.clone();
                                                                                                                              move
                                                                                                                                  |x_1|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_List_Types::Data_List_Types_foldableList()),
                                                                                                                                                                                                                                                                            &&&f_2),
                                                                                                                                                                                                                                         x_1)),
                                                                                                                                                                   &&&PureScript_Data_Set::Data_Set_toList())
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))))
    }
    pub fn Data_Set_findMin() -> &dyn Any {
        static Data_Set_findMin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_findMin.get_or_init(||
                                         &Func1::new(move |v|
                                                         {
                                                             let m =
                                                                 Sharpurs_Prelude::unbox(v);
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                    &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                 &&&Func1::new(move
                                                                                                                                                   |v1|
                                                                                                                                                   find(string("key"),
                                                                                                                                                        Sharpurs_Prelude::unbox(v1)))),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_findMin(),
                                                                                                                                 &&&m))
                                                         }))
    }
    pub fn Data_Set_findMax() -> &dyn Any {
        static Data_Set_findMax: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_findMax.get_or_init(||
                                         &Func1::new(move |v|
                                                         {
                                                             let m =
                                                                 Sharpurs_Prelude::unbox(v);
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                    &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                 &&&Func1::new(move
                                                                                                                                                   |v1|
                                                                                                                                                   find(string("key"),
                                                                                                                                                        Sharpurs_Prelude::unbox(v1)))),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_findMax(),
                                                                                                                                 &&&m))
                                                         }))
    }
    pub fn Data_Set_filter() -> &dyn Any {
        static Data_Set_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_filter.get_or_init(||
                                        &Func1::new(move |dictOrd|
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_filterKeys(),
                                                                                                                            dictOrd))))
    }
    pub fn Data_Set_eqSet() -> &dyn Any {
        static Data_Set_eqSet: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Set_eqSet.get_or_init(||
                                       &Func1::new(move |dictEq|
                                                       {
                                                           let eqMap =
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_eqMap(),
                                                                                                                                   dictEq),
                                                                                                &&&PureScript_Data_Eq::Data_Eq_eqUnit());
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                            &&&add(string("eq"),
                                                                                                   &&Func1::new({
                                                                                                                    let eqMap
                                                                                                                        =
                                                                                                                        eqMap.clone();
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
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                       &&&eqMap),
                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                 &&&matchValue_1)
                                                                                                                                            }
                                                                                                                                    })
                                                                                                                }),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))
                                                       }))
    }
    pub fn Data_Set_ordSet() -> &dyn Any {
        static Data_Set_ordSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_ordSet.get_or_init(||
                                        &Func1::new(move |dictOrd|
                                                        {
                                                            let ordList =
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_ordList(),
                                                                                                 dictOrd);
                                                            let eqSet1 =
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_eqSet(),
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
                                                                                                                         |s1|
                                                                                                                         &Func1::new({
                                                                                                                                         let s1
                                                                                                                                             =
                                                                                                                                             s1.clone();
                                                                                                                                         move
                                                                                                                                             |s2|
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                    &&&ordList),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_toList(),
                                                                                                                                                                                                                                                    &&&s1)),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_toList(),
                                                                                                                                                                                                                 s2))
                                                                                                                                     })
                                                                                                                 }),
                                                                                                    add(string("Eq0"),
                                                                                                        &&Func1::new({
                                                                                                                         let eqSet1
                                                                                                                             =
                                                                                                                             eqSet1.clone();
                                                                                                                         move
                                                                                                                             |usd__unused|
                                                                                                                             &eqSet1
                                                                                                                     }),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))
                                                        }))
    }
    pub fn Data_Set_eq1Set() -> &dyn Any {
        static Data_Set_eq1Set: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_eq1Set.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                         &&&add(string("eq1"),
                                                                                &&Func1::new(move
                                                                                                 |dictEq|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_eqSet(),
                                                                                                                                                                     dictEq))),
                                                                                empty::<string,
                                                                                        &dyn Any>())))
    }
    pub fn Data_Set_ord1Set() -> &dyn Any {
        static Data_Set_ord1Set: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_ord1Set.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                          &&&add(string("compare1"),
                                                                                 &&Func1::new(move
                                                                                                  |dictOrd|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_ordSet(),
                                                                                                                                                                      dictOrd))),
                                                                                 add(string("Eq10"),
                                                                                     &&Func1::new(move
                                                                                                      |usd__unused|
                                                                                                      &PureScript_Data_Set::Data_Set_eq1Set()),
                                                                                     empty::<string,
                                                                                             &dyn Any>()))))
    }
    pub fn Data_Set_empty() -> &dyn Any {
        static Data_Set_empty: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Set_empty.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_Set(),
                                                                        &&&PureScript_Data_Map_Internal::Data_Map_Internal_empty()))
    }
    pub fn Data_Set_fromFoldable() -> &dyn Any {
        static Data_Set_fromFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_fromFoldable.get_or_init(||
                                              &Func1::new(move |dictFoldable|
                                                              &Func1::new({
                                                                              let dictFoldable
                                                                                  =
                                                                                  dictFoldable.clone();
                                                                              move
                                                                                  |dictOrd|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                         &&&dictFoldable),
                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                        let dictOrd
                                                                                                                                                                            =
                                                                                                                                                                            dictOrd.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |m|
                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                            let m
                                                                                                                                                                                                =
                                                                                                                                                                                                m.clone();
                                                                                                                                                                                            move
                                                                                                                                                                                                |a|
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_insert(),
                                                                                                                                                                                                                                                                                                       &&&dictOrd),
                                                                                                                                                                                                                                                                    a),
                                                                                                                                                                                                                                 &&&m)
                                                                                                                                                                                        })
                                                                                                                                                                    })),
                                                                                                                   &&&PureScript_Data_Set::Data_Set_empty())
                                                                          })))
    }
    pub fn Data_Set_map() -> &dyn Any {
        static Data_Set_map: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Set_map.get_or_init(||
                                     &Func1::new(move |dictOrd|
                                                     &Func1::new({
                                                                     let dictOrd
                                                                         =
                                                                         dictOrd.clone();
                                                                     move |f|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                &&&PureScript_Data_Set::Data_Set_foldableSet()),
                                                                                                                                             &&&Func1::new({
                                                                                                                                                               let f
                                                                                                                                                                   =
                                                                                                                                                                   f.clone();
                                                                                                                                                               move
                                                                                                                                                                   |m|
                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                   let m
                                                                                                                                                                                       =
                                                                                                                                                                                       m.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |a|
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_insert(),
                                                                                                                                                                                                                                                                                              &&&dictOrd),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                              a)),
                                                                                                                                                                                                                        &&&m)
                                                                                                                                                                               })
                                                                                                                                                           })),
                                                                                                          &&&PureScript_Data_Set::Data_Set_empty())
                                                                 })))
    }
    pub fn Data_Set_mapMaybe() -> &dyn Any {
        static Data_Set_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_mapMaybe.get_or_init(||
                                          &Func1::new(move |dictOrd|
                                                          &Func1::new({
                                                                          let dictOrd
                                                                              =
                                                                              dictOrd.clone();
                                                                          move
                                                                              |f|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                     &&&PureScript_Data_Set::Data_Set_foldableSet()),
                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                    let f
                                                                                                                                                                        =
                                                                                                                                                                        f.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |a|
                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                        let a
                                                                                                                                                                                            =
                                                                                                                                                                                            a.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |acc|
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                                                                                                                                   acc),
                                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                                  let acc
                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                      acc.clone();
                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                      |b|
                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_insert(),
                                                                                                                                                                                                                                                                                                                                                                                             &&&dictOrd),
                                                                                                                                                                                                                                                                                                                                                          b),
                                                                                                                                                                                                                                                                                                                       &&&acc)
                                                                                                                                                                                                                                                                              })),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                &&&a))
                                                                                                                                                                                    })
                                                                                                                                                                })),
                                                                                                               &&&PureScript_Data_Set::Data_Set_empty())
                                                                      })))
    }
    pub fn Data_Set_monoidSet() -> &dyn Any {
        static Data_Set_monoidSet: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_monoidSet.get_or_init(||
                                           &Func1::new(move |dictOrd|
                                                           {
                                                               let semigroupSet1 =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_semigroupSet(),
                                                                                                    dictOrd);
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                &&&add(string("mempty"),
                                                                                                       &&PureScript_Data_Set::Data_Set_empty(),
                                                                                                       add(string("Semigroup0"),
                                                                                                           &&Func1::new({
                                                                                                                            let semigroupSet1
                                                                                                                                =
                                                                                                                                semigroupSet1.clone();
                                                                                                                            move
                                                                                                                                |usd__unused|
                                                                                                                                &semigroupSet1
                                                                                                                        }),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
                                                           }))
    }
    pub fn Data_Set_unions() -> &dyn Any {
        static Data_Set_unions: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_unions.get_or_init(||
                                        &Func1::new(move |dictFoldable|
                                                        &Func1::new({
                                                                        let dictFoldable
                                                                            =
                                                                            dictFoldable.clone();
                                                                        move
                                                                            |dictOrd|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                   &&&dictFoldable),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_union(),
                                                                                                                                                                                   dictOrd)),
                                                                                                             &&&PureScript_Data_Set::Data_Set_empty())
                                                                    })))
    }
    pub fn Data_Set_difference() -> &dyn Any {
        static Data_Set_difference: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_difference.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_difference(),
                                                                                                                                dictOrd))))
    }
    pub fn Data_Set_subset() -> &dyn Any {
        static Data_Set_subset: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_subset.get_or_init(||
                                        &Func1::new(move |dictOrd|
                                                        &Func1::new({
                                                                        let dictOrd
                                                                            =
                                                                            dictOrd.clone();
                                                                        move
                                                                            |s1|
                                                                            &Func1::new({
                                                                                            let s1
                                                                                                =
                                                                                                s1.clone();
                                                                                            move
                                                                                                |s2|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                    &&&PureScript_Data_Set::Data_Set_isEmpty()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_difference(),
                                                                                                                                                                                                                                          &&&dictOrd),
                                                                                                                                                                                                       &&&s1),
                                                                                                                                                                    s2))
                                                                                        })
                                                                    })))
    }
    pub fn Data_Set_properSubset() -> &dyn Any {
        static Data_Set_properSubset: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_properSubset.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              &Func1::new({
                                                                              let dictOrd
                                                                                  =
                                                                                  dictOrd.clone();
                                                                              move
                                                                                  |s1|
                                                                                  &Func1::new({
                                                                                                  let s1
                                                                                                      =
                                                                                                      s1.clone();
                                                                                                  move
                                                                                                      |s2|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                             &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_notEq(),
                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_size(),
                                                                                                                                                                                                                                                                                   &&&s1)),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_size(),
                                                                                                                                                                                                                                                s2))),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_subset(),
                                                                                                                                                                                                                                                &&&dictOrd),
                                                                                                                                                                                                             &&&s1),
                                                                                                                                                                          s2))
                                                                                              })
                                                                          })))
    }
    pub fn Data_Set_delete() -> &dyn Any {
        static Data_Set_delete: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_delete.get_or_init(||
                                        &Func1::new(move |dictOrd|
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_delete(),
                                                                                                                            dictOrd))))
    }
    pub fn Data_Set_checkValid() -> &dyn Any {
        static Data_Set_checkValid: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_checkValid.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_checkValid(),
                                                                                                                                dictOrd))))
    }
    pub fn Data_Set_catMaybes() -> &dyn Any {
        static Data_Set_catMaybes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Set_catMaybes.get_or_init(||
                                           &Func1::new(move |dictOrd|
                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Set::Data_Set_mapMaybe(),
                                                                                                                               dictOrd),
                                                                                            &&&PureScript_Data_Set::Data_Set_identity())))
    }
}
