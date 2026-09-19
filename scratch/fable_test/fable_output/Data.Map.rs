pub mod PureScript_Data_Map {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_ed2bf3e0::PureScript_Data_Map_Internal;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_e534597::PureScript_Data_Set;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Map_SemigroupMap() -> &dyn Any {
        static Data_Map_SemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_SemigroupMap.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Map_traversableWithIndexSemigroupMap() -> &dyn Any {
        static Data_Map_traversableWithIndexSemigroupMap:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_traversableWithIndexSemigroupMap.get_or_init(||
                                                                  &PureScript_Data_Map_Internal::Data_Map_Internal_traversableWithIndexMap())
    }
    pub fn Data_Map_traversableSemigroupMap() -> &dyn Any {
        static Data_Map_traversableSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_traversableSemigroupMap.get_or_init(||
                                                         &PureScript_Data_Map_Internal::Data_Map_Internal_traversableMap())
    }
    pub fn Data_Map_showSemigroupMap() -> &dyn Any {
        static Data_Map_showSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_showSemigroupMap.get_or_init(||
                                                  &Func1::new(move |dictShow|
                                                                  {
                                                                      let showMap =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_showMap(),
                                                                                                           dictShow);
                                                                      &Func1::new({
                                                                                      let showMap
                                                                                          =
                                                                                          showMap.clone();
                                                                                      move
                                                                                          |dictShow1|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&showMap,
                                                                                                                           dictShow1)
                                                                                  })
                                                                  }))
    }
    pub fn Data_Map_semigroupSemigroupMap() -> &dyn Any {
        static Data_Map_semigroupSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_semigroupSemigroupMap.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictOrd|
                                                                       &Func1::new({
                                                                                       let dictOrd
                                                                                           =
                                                                                           dictOrd.clone();
                                                                                       move
                                                                                           |dictSemigroup|
                                                                                           {
                                                                                               let append =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                    dictSemigroup);
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                                &&&add(string("append"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let append
                                                                                                                                                            =
                                                                                                                                                            append.clone();
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
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map::Data_Map_SemigroupMap(),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_unionWith(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                                                                                                                                              &&&append),
                                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                                        &&&matchValue_1))
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    }),
                                                                                                                                       empty::<string,
                                                                                                                                               &dyn Any>()))
                                                                                           }
                                                                                   })))
    }
    pub fn Data_Map_plusSemigroupMap() -> &dyn Any {
        static Data_Map_plusSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_plusSemigroupMap.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_plusMap(),
                                                                                                   dictOrd)))
    }
    pub fn Data_Map_ordSemigroupMap() -> &dyn Any {
        static Data_Map_ordSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_ordSemigroupMap.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 {
                                                                     let ordMap =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_ordMap(),
                                                                                                          dictOrd);
                                                                     &Func1::new({
                                                                                     let ordMap
                                                                                         =
                                                                                         ordMap.clone();
                                                                                     move
                                                                                         |dictOrd1|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&ordMap,
                                                                                                                          dictOrd1)
                                                                                 })
                                                                 }))
    }
    pub fn Data_Map_ord1SemigroupMap() -> &dyn Any {
        static Data_Map_ord1SemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_ord1SemigroupMap.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_ord1Map(),
                                                                                                   dictOrd)))
    }
    pub fn Data_Map_newtypeSemigroupMap() -> &dyn Any {
        static Data_Map_newtypeSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_newtypeSemigroupMap.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                      &&&add(string("Coercible0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &Sharpurs_Prelude::Prim_undefined()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Map_monoidSemigroupMap() -> &dyn Any {
        static Data_Map_monoidSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_monoidSemigroupMap.get_or_init(||
                                                    &Func1::new(move |dictOrd|
                                                                    {
                                                                        let semigroupSemigroupMap1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map::Data_Map_semigroupSemigroupMap(),
                                                                                                             dictOrd);
                                                                        &Func1::new({
                                                                                        let semigroupSemigroupMap1
                                                                                            =
                                                                                            semigroupSemigroupMap1.clone();
                                                                                        move
                                                                                            |dictSemigroup|
                                                                                            {
                                                                                                let semigroupSemigroupMap2 =
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&semigroupSemigroupMap1,
                                                                                                                                     dictSemigroup);
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                 &&&add(string("mempty"),
                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map::Data_Map_SemigroupMap(),
                                                                                                                                                                          &&&PureScript_Data_Map_Internal::Data_Map_Internal_empty()),
                                                                                                                                        add(string("Semigroup0"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let semigroupSemigroupMap2
                                                                                                                                                                 =
                                                                                                                                                                 semigroupSemigroupMap2.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused|
                                                                                                                                                                 &semigroupSemigroupMap2
                                                                                                                                                         }),
                                                                                                                                            empty::<string,
                                                                                                                                                    &dyn Any>())))
                                                                                            }
                                                                                    })
                                                                    }))
    }
    pub fn Data_Map_keys() -> &dyn Any {
        static Data_Map_keys: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Map_keys.get_or_init(||
                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                          &&&PureScript_Data_Set::Data_Set_fromMap()),
                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                                                          &&&PureScript_Data_Map_Internal::Data_Map_Internal_functorMap())))
    }
    pub fn Data_Map_functorWithIndexSemigroupMap() -> &dyn Any {
        static Data_Map_functorWithIndexSemigroupMap:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_functorWithIndexSemigroupMap.get_or_init(||
                                                              &PureScript_Data_Map_Internal::Data_Map_Internal_functorWithIndexMap())
    }
    pub fn Data_Map_functorSemigroupMap() -> &dyn Any {
        static Data_Map_functorSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_functorSemigroupMap.get_or_init(||
                                                     &PureScript_Data_Map_Internal::Data_Map_Internal_functorMap())
    }
    pub fn Data_Map_foldableWithIndexSemigroupMap() -> &dyn Any {
        static Data_Map_foldableWithIndexSemigroupMap:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_foldableWithIndexSemigroupMap.get_or_init(||
                                                               &PureScript_Data_Map_Internal::Data_Map_Internal_foldableWithIndexMap())
    }
    pub fn Data_Map_foldableSemigroupMap() -> &dyn Any {
        static Data_Map_foldableSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_foldableSemigroupMap.get_or_init(||
                                                      &PureScript_Data_Map_Internal::Data_Map_Internal_foldableMap())
    }
    pub fn Data_Map_eqSemigroupMap() -> &dyn Any {
        static Data_Map_eqSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_eqSemigroupMap.get_or_init(||
                                                &Func1::new(move |dictEq|
                                                                {
                                                                    let eqMap =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_eqMap(),
                                                                                                         dictEq);
                                                                    &Func1::new({
                                                                                    let eqMap
                                                                                        =
                                                                                        eqMap.clone();
                                                                                    move
                                                                                        |dictEq1|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&eqMap,
                                                                                                                         dictEq1)
                                                                                })
                                                                }))
    }
    pub fn Data_Map_eq1SemigroupMap() -> &dyn Any {
        static Data_Map_eq1SemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_eq1SemigroupMap.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_eq1Map(),
                                                                                                  dictEq)))
    }
    pub fn Data_Map_bindSemigroupMap() -> &dyn Any {
        static Data_Map_bindSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_bindSemigroupMap.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_bindMap(),
                                                                                                   dictOrd)))
    }
    pub fn Data_Map_applySemigroupMap() -> &dyn Any {
        static Data_Map_applySemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_applySemigroupMap.get_or_init(||
                                                   &Func1::new(move |dictOrd|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_applyMap(),
                                                                                                    dictOrd)))
    }
    pub fn Data_Map_altSemigroupMap() -> &dyn Any {
        static Data_Map_altSemigroupMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Map_altSemigroupMap.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_altMap(),
                                                                                                  dictOrd)))
    }
}
